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

package journal

import (
	"math"

	"pkg.berachain.dev/stargazer/eth/common"
	coretypes "pkg.berachain.dev/stargazer/eth/core/types"
	"pkg.berachain.dev/stargazer/lib/ds"
	"pkg.berachain.dev/stargazer/lib/ds/stack"
)

// `logs` is a state plugin that tracks Ethereum logs.
type logs struct {
	// Reset every tx.
	ds.Stack[*coretypes.Log] // journal of tx logs
}

// `NewLogs` returns a new `logs` journal.
//
//nolint:revive // only used as a `state.LogsJournal`.
func NewLogs() *logs {
	return &logs{
		Stack: stack.New[*coretypes.Log](initCapacity),
	}
}

// `RegistryKey` implements `libtypes.Registrable`.
func (l *logs) RegistryKey() string {
	return logsRegistryKey
}

// `AddLog` adds a log to the `Logs` store.
func (l *logs) AddLog(log *coretypes.Log) {
	l.Push(log)
}

// `BuildLogsAndClear` returns the logs for the tx with the given metadata.
func (l *logs) BuildLogsAndClear(
	txHash common.Hash,
	blockHash common.Hash,
	txIndex uint,
	logIndex uint,
) []*coretypes.Log {
	size := l.Size()
	buf := make([]*coretypes.Log, size)
	for i := uint(size) - 1; i < math.MaxUint; i-- {
		buf[i] = l.Pop()
		buf[i].TxHash = txHash
		buf[i].BlockHash = blockHash
		buf[i].TxIndex = txIndex
		buf[i].Index = logIndex + i
	}
	return buf
}

// `Snapshot` takes a snapshot of the `Logs` store.
//
// `Snapshot` implements `libtypes.Snapshottable`.
func (l *logs) Snapshot() int {
	return l.Size()
}

// `RevertToSnapshot` reverts the `Logs` store to a given snapshot id.
//
// `RevertToSnapshot` implements `libtypes.Snapshottable`.
func (l *logs) RevertToSnapshot(id int) {
	l.PopToSize(id)
}

// `Finalize` implements `libtypes.Controllable`.
func (l *logs) Finalize() {
	l.Stack = stack.New[*coretypes.Log](initCapacity)
}
