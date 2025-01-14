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

package ethrlp

import (
	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"

	"pkg.berachain.dev/polaris/eth/common"
)

// EthereumRlpEncoded is an interface that should be used to work with all ethereum rlp encoded data.
type EthereumRlpEncoded interface {
	Hash() common.Hash
	MarshalBinary() ([]byte, error)
	UnmarshalBinary([]byte) error
}

// EthereumRlpStore is a wrapper around the parent store that allows to store and retrieve,
// implement `EthereumRlpEncoded` interface.
type EthereumRlpStore[T EthereumRlpEncoded] struct {
	storetypes.KVStore
	parent storetypes.KVStore
}

// NewRlpEncodedStore creates a new instance of `EthereumRlpStore` from provided parent store and key prefix.
func NewRlpEncodedStore[T EthereumRlpEncoded](parent storetypes.KVStore, keyPrefix []byte) *EthereumRlpStore[T] {
	return &EthereumRlpStore[T]{parent: prefix.NewStore(parent, keyPrefix)}
}

// Set stores the provided data in the parent store.
func (rlps *EthereumRlpStore[T]) Set(key []byte, data T) {
	bz, err := data.MarshalBinary()
	if err != nil {
		// ctx.Logger().Error("MarshalBinary for block. Failed to update offchain storagae", "err", err)
		return
	}
	rlps.parent.Set(key, bz)
}

// Get retrieves the unmarshalled data from the parent store.
func (rlps *EthereumRlpStore[T]) Get(key []byte) (T, bool) {
	var t T
	bz := rlps.parent.Get(key)
	if bz == nil {
		return t, false
	}

	err := t.UnmarshalBinary(bz)
	if err != nil {
		return t, false
	}
	return t, true
}
