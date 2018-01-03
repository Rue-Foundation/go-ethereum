// Copyright 2016 The go-rue Authors
// This file is part of the go-rue library.
//
// The go-rue library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-rue library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-rue library. If not, see <http://www.gnu.org/licenses/>.

package rueclient

import "github.com/Rue-Foundation/go-rue"

// Verify that Client implements the rue interfaces.
var (
	_ = rue.ChainReader(&Client{})
	_ = rue.TransactionReader(&Client{})
	_ = rue.ChainStateReader(&Client{})
	_ = rue.ChainSyncReader(&Client{})
	_ = rue.ContractCaller(&Client{})
	_ = rue.GasEstimator(&Client{})
	_ = rue.GasPricer(&Client{})
	_ = rue.LogFilterer(&Client{})
	_ = rue.PendingStateReader(&Client{})
	// _ = rue.PendingStateEventer(&Client{})
	_ = rue.PendingContractCaller(&Client{})
)
