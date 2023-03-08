// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2023, Berachain Foundation. All rights reserved.
// Use of this software is govered by the Business Source License included
// in the LICENSE file of this repository and at www.mariadb.com/bsl11.
//
// ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
// TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
// VERSIONS OF THE LICENSED WORK.
//
// THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
// LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
// LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
//
// TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
// AN “AS IS” BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.

package governance

import (
	"context"
	"math/big"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	v1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"

	"pkg.berachain.dev/polaris/precompile/contracts/solidity/generated"
)

// `submitProposalHelper` is a helper function for the `SubmitProposal` method of the governance precompile contract.
func (c *Contract) submitProposalHelper(
	ctx context.Context,
	messages []*codectypes.Any,
	initialDeposit []generated.IGovernanceModuleCoin,
	proposer sdk.AccAddress,
	metadata, title, summary string,
	expedited bool,
) ([]any, error) {
	coins := []sdk.Coin{}

	// Convert the initial deposit to sdk.Coin.
	for _, coin := range initialDeposit {
		coins = append(coins, sdk.NewCoin(coin.Denom, sdk.NewIntFromBigInt(coin.Amount)))
	}

	res, err := c.msgServer.SubmitProposal(ctx, &v1.MsgSubmitProposal{
		Messages:       messages,
		InitialDeposit: coins,
		Proposer:       proposer.String(),
		Metadata:       metadata,
		Title:          title,
		Summary:        summary,
		Expedited:      expedited,
	})
	if err != nil {
		return nil, err
	}

	return []any{big.NewInt(int64(res.ProposalId))}, nil
}

// `cancelProposalHelper` is a helper function for the `CancelProposal` method of the governance precompile contract.
func (c *Contract) cancelProposalHelper(
	ctx context.Context,
	proposer sdk.AccAddress,
	proposalID *big.Int,
) ([]any, error) {
	res, err := c.msgServer.CancelProposal(ctx, &v1.MsgCancelProposal{
		ProposalId: proposalID.Uint64(),
		Proposer:   proposer.String(),
	})
	if err != nil {
		return nil, err
	}

	return []any{big.NewInt(res.CanceledTime.Unix()), big.NewInt(int64(res.CanceledHeight))}, nil
}

// `execLegacyContentHelper` is a helper function for the `ExecLegacyContent`
// method of the governance precompile contract.
func (c *Contract) execLegacyContentHelper(
	ctx context.Context,
	content *codectypes.Any,
	authority string,
) ([]any, error) {
	_, err := c.msgServer.ExecLegacyContent(ctx, &v1.MsgExecLegacyContent{
		Content:   content,
		Authority: authority,
	})
	if err != nil {
		return nil, err
	}

	return []any{}, nil
}

// `voteHelper` is a helper function for the `Vote` method of the governance precompile contract.
func (c *Contract) voteHelper(
	ctx context.Context,
	voter sdk.AccAddress,
	proposalID *big.Int,
	option int32,
	metadata string,
) ([]any, error) {
	_, err := c.msgServer.Vote(ctx, &v1.MsgVote{
		ProposalId: proposalID.Uint64(),
		Voter:      voter.String(),
		Option:     v1.VoteOption(option),
		Metadata:   metadata,
	})
	if err != nil {
		return nil, err
	}

	return []any{}, nil
}

// `voteWeighted` is a helper function for the `VoteWeighted` method of the governance precompile contract.
func (c *Contract) voteWeightedHelper(
	ctx context.Context,
	voter sdk.AccAddress,
	proposalID *big.Int,
	options []generated.IGovernanceModuleWeightedVoteOption,
	metadata string,
) ([]any, error) {
	// Convert the options to v1.WeightedVoteOption.
	msgOptions := make([]*v1.WeightedVoteOption, len(options))
	for i, option := range options {
		msgOptions[i] = &v1.WeightedVoteOption{
			Option: v1.VoteOption(option.VoteOption),
			Weight: option.Weight,
		}
	}

	_, err := c.msgServer.VoteWeighted(
		ctx, &v1.MsgVoteWeighted{
			ProposalId: proposalID.Uint64(),
			Voter:      voter.String(),
			Options:    msgOptions,
			Metadata:   metadata,
		},
	)
	if err != nil {
		return nil, err
	}

	return []any{}, nil
}

// `getProposalHelper` is a helper function for the `GetProposal` method of the governance precompile contract.
func (c *Contract) getProposalHelper(ctx context.Context, proposalID *big.Int) ([]any, error) {
	res, err := c.querier.Proposal(ctx, &v1.QueryProposalRequest{
		ProposalId: proposalID.Uint64(),
	})
	if err != nil {
		return nil, err
	}
	return []any{transformProposalToABIProposal(*res.Proposal)}, nil
}

// `getProposalsHelper` is a helper function for the `GetProposal` method of the governance precompile contract.
func (c *Contract) getProposalsHelper(
	ctx context.Context,
	proposalStatus int32,
	voter,
	depositor sdk.AccAddress,
) ([]any, error) {
	res, err := c.querier.Proposals(ctx, &v1.QueryProposalsRequest{
		ProposalStatus: v1.ProposalStatus(proposalStatus),
		Voter:          voter.String(),
		Depositor:      depositor.String(),
		Pagination:     nil,
	})
	if err != nil {
		return nil, err
	}

	var proposals []generated.IGovernanceModuleProposal
	for _, proposal := range res.Proposals {
		proposals = append(proposals, transformProposalToABIProposal(*proposal))
	}

	return []any{proposals}, nil
}

// `transformProposalToABIProposal` is a helper function to transform a `v1.Proposal` to an `IGovernanceModule.Proposal`.
func transformProposalToABIProposal(proposal v1.Proposal) generated.IGovernanceModuleProposal {
	var message []byte
	for _, msg := range proposal.Messages {
		message = append(message, msg.Value...)
	}

	var totalDeposit []generated.IGovernanceModuleCoin
	for _, coin := range proposal.TotalDeposit {
		totalDeposit = append(totalDeposit, generated.IGovernanceModuleCoin{
			Denom:  coin.Denom,
			Amount: coin.Amount.BigInt(),
		})
	}

	return generated.IGovernanceModuleProposal{
		Id:      proposal.Id,
		Message: message,
		Status:  int32(proposal.Status), // Status is an alias for int32.
		FinalTallyResult: generated.IGovernanceModuleTallyResult{
			YesCount:        proposal.FinalTallyResult.YesCount,
			AbstainCount:    proposal.FinalTallyResult.AbstainCount,
			NoCount:         proposal.FinalTallyResult.NoCount,
			NoWithVetoCount: proposal.FinalTallyResult.NoWithVetoCount,
		},
		SubmitTime:      uint64(proposal.SubmitTime.Unix()),
		DepositEndTime:  uint64(proposal.DepositEndTime.Unix()),
		TotalDeposit:    totalDeposit,
		VotingStartTime: uint64(proposal.VotingStartTime.Unix()),
		VotingEndTime:   uint64(proposal.VotingEndTime.Unix()),
		Metadata:        proposal.Metadata,
		Title:           proposal.Title,
		Summary:         proposal.Summary,
		Proposer:        proposal.Proposer,
	}
}
