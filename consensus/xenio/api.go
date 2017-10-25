// Copyright 2017 The go-xenio Authors
// Copyright 2017 The go-ethereum Authors
//
// This file is part of the go-xenio library.
//
// The go-xenio library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-xenio library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-xenio library. If not, see <http://www.gnu.org/licenses/>.

package xenio

import (
	"github.com/xenioplatform/go-xenio/common"
	"github.com/xenioplatform/go-xenio/consensus"
	"github.com/xenioplatform/go-xenio/contracts/xnogames"
	"github.com/xenioplatform/go-xenio/contracts/xnousers"
	"github.com/xenioplatform/go-xenio/core/types"
	"github.com/xenioplatform/go-xenio/rpc"
	"github.com/xenioplatform/go-xenio/log"
)

// API is a user facing RPC API to allow controlling the signer and voting
// mechanisms of the proof-of-authority scheme.
type API struct {
	chain  consensus.ChainReader
	xenio *Xenio
}

// GetSnapshot retrieves the state snapshot at a given block.
func (api *API) GetSnapshot(number *rpc.BlockNumber) (*Snapshot, error) {
	// Retrieve the requested block number (or current if none requested)
	var header *types.Header
	if number == nil || *number == rpc.LatestBlockNumber {
		header = api.chain.CurrentHeader()
	} else {
		header = api.chain.GetHeaderByNumber(uint64(number.Int64()))
	}
	// Ensure we have an actually valid block and return its snapshot
	if header == nil {
		return nil, errUnknownBlock
	}
	return api.xenio.snapshot(api.chain, header.Number.Uint64(), header.Hash(), nil)
}

// GetSnapshotAtHash retrieves the state snapshot at a given block.
func (api *API) GetSnapshotAtHash(hash common.Hash) (*Snapshot, error) {
	header := api.chain.GetHeaderByHash(hash)
	if header == nil {
		return nil, errUnknownBlock
	}
	return api.xenio.snapshot(api.chain, header.Number.Uint64(), header.Hash(), nil)
}

// GetSigners retrieves the list of authorized signers at the specified block.
func (api *API) GetSigners(number *rpc.BlockNumber) ([]common.Address, error) {
	// Retrieve the requested block number (or current if none requested)
	var header *types.Header
	if number == nil || *number == rpc.LatestBlockNumber {
		header = api.chain.CurrentHeader()
	} else {
		header = api.chain.GetHeaderByNumber(uint64(number.Int64()))
	}
	// Ensure we have an actually valid block and return the signers from its snapshot
	if header == nil {
		return nil, errUnknownBlock
	}
	snap, err := api.xenio.snapshot(api.chain, header.Number.Uint64(), header.Hash(), nil)
	if err != nil {
		return nil, err
	}
	return snap.masterNodes(), nil
}

// GetSignersAtHash retrieves the state snapshot at a given block.
func (api *API) GetSignersAtHash(hash common.Hash) ([]common.Address, error) {
	header := api.chain.GetHeaderByHash(hash)
	if header == nil {
		return nil, errUnknownBlock
	}
	snap, err := api.xenio.snapshot(api.chain, header.Number.Uint64(), header.Hash(), nil)
	if err != nil {
		return nil, err
	}
	return snap.masterNodes(), nil
}

// Proposals returns the current proposals the node tries to uphold and vote on.
func (api *API) Proposals() map[common.Address]bool {
	api.xenio.lock.RLock()
	defer api.xenio.lock.RUnlock()

	proposals := make(map[common.Address]bool)
	for address, auth := range api.xenio.proposals {
		proposals[address] = auth
	}
	return proposals
}

// Proposals returns the current proposals the node tries to uphold and vote on.
func (api *API) Votes() map[string]Vote {
	api.xenio.lock.RLock()
	defer api.xenio.lock.RUnlock()

	votes := make(map[string]Vote)
	for str, vt := range api.xenio.Votes {
		votes[str] = vt
	}
	return votes
}

// Propose injects a new authorization proposal that the signer will attempt to
// push through.
func (api *API) Propose(address common.Address, auth bool) {
	api.xenio.lock.Lock()
	defer api.xenio.lock.Unlock()

	api.xenio.proposals[address] = auth
}

func (api *API) GetXNOGamesABI() string{
	api.xenio.lock.Lock()
	defer api.xenio.lock.Unlock()
	return xnogames.XNOGamesABI
}

// GamesContractPropose injects a new games contract authorization proposal that the signer will attempt to
// push through.
func (api *API) GamesContractVote(address common.Address, vote bool) bool{
	api.xenio.lock.Lock()
	defer api.xenio.lock.Unlock()

	var _vote Vote
	//vote.Signer = ?
	_vote.Authorize = vote
	_vote.VoteType = GamesContract
	_vote.Address = address // contract address
	api.xenio.Votes[address.String()] = _vote

	return true
}

// GameServerVote injects a new game server authorization proposal that the signer will attempt to
// push through.
func (api *API) GameServerVote(address common.Address, vote bool) bool{
	var ca common.Address
	if api.xenio.signer == ca{
		log.Error("coinbase not parsed, check if staker is active")
		return false
	}
	api.xenio.lock.Lock()
	defer api.xenio.lock.Unlock()
	var _vote Vote
	_vote.Signer = api.xenio.signer
	_vote.Authorize = vote
	_vote.VoteType = MasterNode
	_vote.Address = address // server's coinbase address
	vhash := common.GetMD5Hash(_vote.Address.String() + _vote.Signer.String())
	api.xenio.Votes[vhash] = _vote

	return true
}

func (api *API) GetXNOUsersABI() string{
	api.xenio.lock.Lock()
	defer api.xenio.lock.Unlock()
	return xnousers.XNOUsersABI
}

// UsersContractPropose injects a new users contract authorization proposal that the signer will attempt to
// push through.
func (api *API) UsersContractVote(address common.Address, vote bool) bool{
	api.xenio.lock.Lock()
	defer api.xenio.lock.Unlock()

	var _vote Vote
	//vote.Signer = ?
	_vote.Authorize = vote
    _vote.VoteType = UsersContract
    _vote.Address = address // contract address
    //_vote.Block = ?
	api.xenio.Votes[address.String()] = _vote

	return true
}

// Discard drops a currently running proposal, stopping the signer from casting
// further votes (either for or against).
func (api *API) Discard(address common.Address) {
	api.xenio.lock.Lock()
	defer api.xenio.lock.Unlock()

	delete(api.xenio.proposals, address)
}

// GetRewardsList returns list of rewards in block
func (api *API) GetRewardsList(number *rpc.BlockNumber) ([]common.Address, error) {
	// Retrieve the requested block number (or current if none requested)
	var header *types.Header
	if number == nil || *number == rpc.LatestBlockNumber {
		header = api.chain.CurrentHeader()
	} else {
		header = api.chain.GetHeaderByNumber(uint64(number.Int64()))
	}
	// Ensure we have an actually valid block and return its snapshot
	if header == nil {
		return nil, errUnknownBlock
	}

	return header.RewardList, nil
}

func (api *API) GetCompletedTransactions(address common.Address) interface{} {
	header := api.chain.CurrentHeader()

	completedTxs := make([]*types.Transaction, 0)

	for n := uint64(1); n <= header.Number.Uint64(); n++ {
		h := api.chain.GetHeaderByNumber(n)
		b := api.chain.GetBlock(h.Hash(), n)

		txs := b.Transactions()
		for t := range txs {
			to := txs[t].To()
			if to != nil && *to == address {
				completedTxs = append(completedTxs, txs[t])
			}
		}
	}

	return completedTxs
}
