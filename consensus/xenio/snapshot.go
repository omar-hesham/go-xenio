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
	"bytes"
	"encoding/json"
	"time"

	"github.com/xenioplatform/go-xenio/common"
	"github.com/xenioplatform/go-xenio/core/types"
	"github.com/xenioplatform/go-xenio/consensus"
	"github.com/xenioplatform/go-xenio/ethdb"
	"github.com/xenioplatform/go-xenio/params"
	"github.com/xenioplatform/go-xenio/log"
	lru "github.com/hashicorp/golang-lru"
)

// Vote represents a single vote that an authorized signer made to modify the
// list of authorizations.
type Vote struct {
	Signer    common.Address `json:"signer"`    // Authorized signer that cast this vote
	Block     uint64         `json:"block"`     // Block number the vote was cast in (expire old votes)
	Address   common.Address `json:"address"`   // Account being voted on to change its authorization
	Authorize bool           `json:"authorize"` // Whether to authorize or deauthorize the voted account
}

// Tally is a simple vote tally to keep the current score of votes. Votes that
// go against the proposal aren't counted since it's equivalent to not voting.
type Tally struct {
	Authorize bool `json:"authorize"` // Whether the vote is about authorizing or kicking someone
	Votes     int  `json:"votes"`     // Number of votes until now wanting to pass the proposal
}

// Snapshot is the state of the authorization voting at a given point in time.
type Snapshot struct {
	config         *params.XenioConfig // Consensus engine parameters to fine tune behavior
	sigcache       *lru.ARCCache        // Cache of recent block signatures to speed up ecrecover

	Number         uint64                      `json:"number"`      // Block number where the snapshot was created
	Hash           common.Hash                 `json:"hash"`        // Block hash where the snapshot was created
	MasterNodes    map[common.Address]Signer   `json:"masternodes"` // Set of authorized super signing nodes at this moment
	StakingNodes   map[common.Address]Signer   `json:"stakingnodes"`// Set of normal signers
	Recents        map[uint64]common.Address   `json:"recents"`     // Set of recent signers for spam protections
	Votes          []*Vote                     `json:"votes"`       // List of votes cast in chronological order
	Tally          map[common.Address]Tally    `json:"tally"`       // Current vote tally to avoid recalculating
	LastSuperBlock time.Time                   `json:"lastsuperbloctime"` // the time of the last known superblock
}

type Signer struct {
	BlockNumber  []uint64					`json:"blocknumber"` // Block number assigned for signing
	SignDate 	 time.Time					`json:"signdate"`    // Date Time for Block signing
	IsMasterNode bool                       `json:"ismasternode"`// indicates of the node is regular or not
}

// newSnapshot creates a new snapshot with the specified startup parameters. This
// method does not initialize the set of recent signers, so only ever use if for
// the genesis block.
func newSnapshot(config *params.XenioConfig, sigcache *lru.ARCCache, number uint64, hash common.Hash, signers []common.Address) *Snapshot {
	snap := &Snapshot{
		config:       config,
		sigcache:     sigcache,
		Number:       number,
		Hash:         hash,
		MasterNodes:  make(map[common.Address]Signer),
		StakingNodes: make(map[common.Address]Signer),
		Recents:      make(map[uint64]common.Address),
		Tally:        make(map[common.Address]Tally),
	}
	var newSigner Signer
	for i, signer := range signers {
		newSigner.IsMasterNode = true
		newSigner.BlockNumber = make([]uint64, 1)
		newSigner.BlockNumber[0] = number + uint64(i + 1) // Block Zero not in play!
		newSigner.SignDate = time.Unix(time.Now().UTC().Unix()+int64(i)*int64(config.Period), 0).UTC()
		snap.MasterNodes[signer] = newSigner
	}
	return snap
}

// loadSnapshot loads an existing snapshot from the database.
func loadSnapshot(config *params.XenioConfig, sigcache *lru.ARCCache, db ethdb.Database, hash common.Hash) (*Snapshot, error) {
	blob, err := db.Get(append([]byte("xenio-"), hash[:]...))
	if err != nil {
		return nil, err
	}
	snap := new(Snapshot)
	if err := json.Unmarshal(blob, snap); err != nil {
		return nil, err
	}
	snap.config = config
	snap.sigcache = sigcache

	return snap, nil
}

// store inserts the snapshot into the database.
func (s *Snapshot) store(db ethdb.Database) error {
	blob, err := json.Marshal(s)
	if err != nil {
		return err
	}
	return db.Put(append([]byte("xenio-"), s.Hash[:]...), blob)
}

// copy creates a deep copy of the snapshot, though not the individual votes.
func (s *Snapshot) copy() *Snapshot {
	cpy := &Snapshot{
		config:       s.config,
		sigcache:     s.sigcache,
		Number:       s.Number,
		Hash:         s.Hash,
		MasterNodes:  make(map[common.Address]Signer),
		StakingNodes: make(map[common.Address]Signer),
		Recents:      make(map[uint64]common.Address),
		Votes:        make([]*Vote, len(s.Votes)),
		Tally:        make(map[common.Address]Tally),
	}
	for address, signerData := range s.MasterNodes {
		cpy.MasterNodes[address] = signerData
	}
	for address, signerData := range s.StakingNodes {
		cpy.StakingNodes[address] = signerData
	}
	for block, signer := range s.Recents {
		cpy.Recents[block] = signer
	}
	for address, tally := range s.Tally {
		cpy.Tally[address] = tally
	}
	copy(cpy.Votes, s.Votes)

	return cpy
}

// validVote returns whether it makes sense to cast the specified vote in the
// given snapshot context (e.g. don't try to add an already authorized signer).
func (s *Snapshot) validVote(address common.Address, authorize bool) bool {
	_, signer := s.MasterNodes[address]
	return (signer && !authorize) || (!signer && authorize)
}

// cast adds a new vote into the tally.
func (s *Snapshot) cast(address common.Address, authorize bool) bool {
	// Ensure the vote is meaningful
	if !s.validVote(address, authorize) {
		return false
	}
	// Cast the vote into an existing or new tally
	if old, ok := s.Tally[address]; ok {
		old.Votes++
		s.Tally[address] = old
	} else {
		s.Tally[address] = Tally{Authorize: authorize, Votes: 1}
	}
	return true
}

// uncast removes a previously cast vote from the tally.
func (s *Snapshot) uncast(address common.Address, authorize bool) bool {
	// If there's no tally, it's a dangling vote, just drop
	tally, ok := s.Tally[address]
	if !ok {
		return false
	}
	// Ensure we only revert counted votes
	if tally.Authorize != authorize {
		return false
	}
	// Otherwise revert the vote
	if tally.Votes > 1 {
		tally.Votes--
		s.Tally[address] = tally
	} else {
		delete(s.Tally, address)
	}
	return true
}

// apply creates a new authorization snapshot by applying the given headers to
// the original one.
func (s *Snapshot) apply(headers []*types.Header) (*Snapshot, error) {
	// Allow passing in no headers for cleaner code
	if len(headers) == 0 {
		return s, nil
	}
	// Sanity check that the headers can be applied
	for i := 0; i < len(headers)-1; i++ {
		if headers[i+1].Number.Uint64() != headers[i].Number.Uint64()+1 {
			return nil, errInvalidVotingChain
		}
	}
	if headers[0].Number.Uint64() != s.Number+1 {
		return nil, errInvalidVotingChain
	}
	// Iterate through the headers and create a new snapshot
	snap := s.copy()

	for _, header := range headers {
		// Remove any votes on checkpoint blocks
		number := header.Number.Uint64()
		if number%s.config.Epoch == 0 {
			snap.Votes = nil
			snap.Tally = make(map[common.Address]Tally)
		}
		// Delete the oldest signer from the recent list to allow it signing again
		if limit := uint64(len(snap.MasterNodes)/2 + 1); number >= limit {
			delete(snap.Recents, number-limit)
		}
		// Resolve the authorization key and check against signers
		signer, err := ecrecover(header, s.sigcache)
		if err != nil {
			return nil, err
		}

		snap.Recents[number] = signer

		// Header authorized, discard any previous votes from the signer
		for i, vote := range snap.Votes {
			if vote.Signer == signer && vote.Address == header.Coinbase {
				// Uncast the vote from the cached tally
				snap.uncast(vote.Address, vote.Authorize)

				// Uncast the vote from the chronological list
				snap.Votes = append(snap.Votes[:i], snap.Votes[i+1:]...)
				break // only one vote allowed
			}
		}
		// Tally up the new vote from the signer
		var authorize bool
		switch {
		case bytes.Equal(header.Nonce[:], nonceAuthVote):
			authorize = true
		case bytes.Equal(header.Nonce[:], nonceDropVote):
			authorize = false
		default:
			return nil, errInvalidVote
		}
		if snap.cast(header.Coinbase, authorize) {
			snap.Votes = append(snap.Votes, &Vote{
				Signer:    signer,
				Block:     number,
				Address:   header.Coinbase,
				Authorize: authorize,
			})
		}
		// If the vote passed, update the list of master nodes
		if tally := snap.Tally[header.Coinbase]; tally.Votes > len(snap.MasterNodes)/2 {
			if tally.Authorize {
				var newSigner Signer
				snap.MasterNodes[header.Coinbase] = newSigner
			} else {
				delete(snap.MasterNodes, header.Coinbase)

				// Signer list shrunk, delete any leftover recent caches
				if limit := uint64(len(snap.MasterNodes)/2 + 1); number >= limit {
					delete(snap.Recents, number-limit)
				}
				// Discard any previous votes the deauthorized signer cast
				for i := 0; i < len(snap.Votes); i++ {
					if snap.Votes[i].Signer == header.Coinbase {
						// Uncast the vote from the cached tally
						snap.uncast(snap.Votes[i].Address, snap.Votes[i].Authorize)

						// Uncast the vote from the chronological list
						snap.Votes = append(snap.Votes[:i], snap.Votes[i+1:]...)

						i--
					}
				}
			}
			// Discard any previous votes around the just changed account
			for i := 0; i < len(snap.Votes); i++ {
				if snap.Votes[i].Address == header.Coinbase {
					snap.Votes = append(snap.Votes[:i], snap.Votes[i+1:]...)
					i--
				}
			}
			delete(snap.Tally, header.Coinbase)
		}
		// Retrieve and update regular Signer List
		if len(header.SuperBlock) > 0{
			superBlockData := make(map[common.Address]Signer,0)

			if err := json.Unmarshal(header.SuperBlock,&superBlockData); err != nil {
				log.Trace(err.Error())
			}else{
				snap.StakingNodes = nil //clear old list
				snap.StakingNodes = make(map[common.Address]Signer,0)
				for key, node := range superBlockData{
					if node.IsMasterNode {
						snap.MasterNodes[key] = node
					}else{
						snap.StakingNodes[key] = node
					}
				}
				snap.LastSuperBlock =time.Unix(header.Time.Int64(),0)
			}
		}

	}
	snap.Number += uint64(len(headers))
	snap.Hash = headers[len(headers)-1].Hash()

	//genesis := chain.GetHeaderByNumber(0)
	//if err := c.VerifyHeader(chain, genesis, false); err != nil {
	//	return nil, err
	//}
	//signers := make([]common.Address, (len(genesis.Extra)-extraVanity-extraSeal)/common.AddressLength)
	//for i := 0; i < len(signers); i++ {
	//	copy(signers[i][:], genesis.Extra[extraVanity+i*common.AddressLength:])
	//}
	//snap = updateSigners(snap, nil, snap.Number, nil, snap.MasterNodes)

	//blob,_ := json.Marshal(snap)
	//log.Warn(string(blob))
	return snap, nil
}


// signers retrieves the list of authorized signers in ascending order.
func (s *Snapshot) masterNodes() []common.Address {
	signers := make([]common.Address, 0, len(s.MasterNodes))
	for signer := range s.MasterNodes {
		signers = append(signers, signer)
	}
	for i := 0; i < len(signers); i++ {
		for j := i + 1; j < len(signers); j++ {
			if bytes.Compare(signers[i][:], signers[j][:]) > 0 {
				signers[i], signers[j] = signers[j], signers[i]
			}
		}
	}
	return signers
}
/*
func updateSigners(snap *Snapshot, config *params.XenioConfig, number uint64, SignDate *big.Int, signers map[common.Address]Signer) *Snapshot {
	var newSigner Signer
	i := 0
	for address := range signers {
		newSigner.BlockNumber = number + uint64(i+1) // Assign next Block
		//newSigner.SignDate = time.Unix(SignDate.Int64()+int64(i+1)*int64(config.Period), 0).UTC()
		snap.MasterNodes[address] = newSigner
		i++
	}
	return snap
}
*/
//func updateSigners(snap *Snapshot, config *params.XenioConfig, number uint64, SignDate *big.Int, signers []common.Address) *Snapshot {
//	var newSigner Signer
//	for i, signer := range signers {
//		newSigner.BlockNumber = number + uint64(i+1) // Assign next Block
//		newSigner.SignDate = time.Unix(SignDate.Int64()+int64(i+1)*int64(config.Period), 0).UTC()
//		snap.Signers[signer] = newSigner
//	}
//	return snap
//}

// inturn returns if a signer at a given block height is in-turn or not.
func (s *Snapshot) inturn(number uint64, signer common.Address) bool {
	signers, offset := s.masterNodes(), 0
	for offset < len(signers) && signers[offset] != signer {
		offset++
	}
	return (number % uint64(len(signers))) == uint64(offset)
}

// Checks whether the node is authorised
//func isAuthorisedNode(snap *Snapshot, signer common.Address) bool {
//	_, isAuthorizedMaster := snap.MasterNodes[signer]
//	_, isAuthorizedStaking := snap.StakingNodes[signer]
//	return isAuthorizedMaster || isAuthorizedStaking
//}

// Locates the singer in the snapshot and returns the signing node
func (s *Snapshot) getSigningNode(signer common.Address) (Signer, bool) {
	if signingNode, isAuthorized := s.MasterNodes[signer]; isAuthorized { return signingNode, isAuthorized } //Check master nodes
	if signingNode, isAuthorized := s.StakingNodes[signer]; isAuthorized { return signingNode, isAuthorized }//Check staking nodes
	return Signer{}, false	//Not authorised node
}

// Checks whether the signing node is next in turn
func (signingNode Signer) isInTurn (s *Snapshot) bool{
	// Check if the authorised node already contains the next in turn block
	for _, turn := range signingNode.BlockNumber{
		if turn == s.Number + 1 { return true }
	}
	return false
}

// Count how many nodes in the snapshot are prior to the signing node
func (s *Snapshot) getPriorNodes(signingNode Signer, currentBlockNumber uint64) int {
	priorNodes := 0
	minBlock := uint64(0)
	for _, numb := range signingNode.BlockNumber{
		if numb < currentBlockNumber{ continue }
		if minBlock < numb { minBlock = numb}
	}
	if minBlock > 0{
		for _, node := range s.StakingNodes{
			//assuming that block arrays are in order
			for _, numb := range node.BlockNumber{
				if numb < currentBlockNumber{ continue }
				if minBlock < numb { priorNodes++ }
			}
		}
	}
	blob, _ := json.Marshal(priorNodes)
	log.Warn(string(blob))
	return priorNodes
}

// Estimate the delivery time of previous blocks of all nodes prior to the signing node
func (s *Snapshot) estimatePriorDelayTime(chain consensus.ChainReader, signingNode Signer, wiggleTime time.Duration) time.Time {
	// count nodes prior to the signing node
	priorNodes := s.getPriorNodes(signingNode,chain.CurrentHeader().Number.Uint64())
	// add 2mins for each prior node
	dt := time.Unix(chain.CurrentHeader().Time.Int64(), 0).Add(wiggleTime*time.Duration(2 * priorNodes))
	return dt
}

// Checks whether a staking node tries to over-turn a master node
func (signingNode Signer) isOverTurner (s *Snapshot) bool {
	if !signingNode.IsMasterNode { //if staking node
		for _, node := range s.MasterNodes {
			for _, turn := range node.BlockNumber {
				if turn == s.Number+1 {
					return true
				}
			}
		}
	}
	return false
}

// Checks whether the given block number already exists in the nodes list
func isDuplicated (block_number uint64, nodes map[common.Address]Signer) bool {
	for _, node := range nodes {
		for _, turn := range node.BlockNumber {
			if turn == block_number { return true }
		}
	}
	return false
}



