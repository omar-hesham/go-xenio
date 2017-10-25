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

// Package xenio implements the proof-of-authority consensus engine.
package xenio

import (
	"bytes"
	"errors"
	"math/big"
	"sync"
	"time"

	"github.com/xenioplatform/go-xenio/accounts"
	"github.com/xenioplatform/go-xenio/common"
	"github.com/xenioplatform/go-xenio/common/hexutil"
	"github.com/xenioplatform/go-xenio/consensus"
	"github.com/xenioplatform/go-xenio/core/state"
	"github.com/xenioplatform/go-xenio/core/types"
	"github.com/xenioplatform/go-xenio/crypto"
	"github.com/xenioplatform/go-xenio/crypto/sha3"
	"github.com/xenioplatform/go-xenio/ethdb"
	"github.com/xenioplatform/go-xenio/log"
	"github.com/xenioplatform/go-xenio/params"
	"github.com/xenioplatform/go-xenio/rlp"
	"github.com/xenioplatform/go-xenio/rpc"
	lru "github.com/hashicorp/golang-lru"
	"encoding/json"

	"fmt"
)

const (
	checkpointInterval = 1024 // Number of blocks after which to save the vote snapshot to the database
	inmemorySnapshots  = 128  // Number of recent vote snapshots to keep in memory
	inmemorySignatures = 4096 // Number of recent block signatures to keep in memory

	wiggleTime = time.Second * 15 // delay
	noiseScalingFactor = 0.1 // scale of noise

	requiredCoins = 1000 // Voter must hold at least this amount of coins in order to be eligible to vote
	// ReceiptStatusFailed is the status code of a transaction if execution failed.
	ReceiptStatusFailed = uint(0) // should have the same value with receipt.go
	// ReceiptStatusSuccessful is the status code of a transaction if execution succeeded.
	ReceiptStatusSuccessful = uint(1)// should have the same value with receipt.go
)

// xenio proof-of-network protocol constants.
var (
	epochLength = uint64(30000) // Default number of blocks after which to checkpoint and reset the pending votes
	blockPeriod = uint64(15)    // Default minimum difference between two consecutive block's timestamps
	superPeriod = uint64(20*60)

	extraVanity = 32 // Fixed number of extra-data prefix bytes reserved for signer vanity
	extraSeal   = 65 // Fixed number of extra-data suffix bytes reserved for signer seal

	nonceAuthVote = hexutil.MustDecode("0xffffffffffffffff") // Magic nonce number to vote on adding a new signer
	nonceDropVote = hexutil.MustDecode("0x0000000000000000") // Magic nonce number to vote on removing a signer.

	uncleHash = types.CalcUncleHash(nil) // Always Keccak256(RLP([])) as uncles are meaningless outside of PoW.

	diffInTurn = big.NewInt(2) // Block difficulty for in-turn signatures
	diffNoTurn = big.NewInt(1) // Block difficulty for out-of-turn signatures

	blockReward *big.Int = big.NewInt(0) // big.NewInt(5e+18)
	big8  = big.NewInt(8)
	big32 = big.NewInt(32)

	currentState *state.StateDB

	currentIPCEndpoint string // required for interaction with contracts
)

// Various error messages to mark blocks invalid. These should be private to
// prevent engine specific errors from being referenced in the remainder of the
// codebase, inherently breaking if the engine is swapped out. Please put common
// error types into the consensus package.
var (
	// errUnknownBlock is returned when the list of signers is requested for a block
	// that is not part of the local blockchain.
	errUnknownBlock = errors.New("unknown block")

	// errInvalidCheckpointBeneficiary is returned if a checkpoint/epoch transition
	// block has a beneficiary set to non-zeroes.
	errInvalidCheckpointBeneficiary = errors.New("beneficiary in checkpoint block non-zero")

	// errInvalidVote is returned if a nonce value is something else that the two
	// allowed constants of 0x00..0 or 0xff..f.
	errInvalidVote = errors.New("vote nonce not 0x00..0 or 0xff..f")

	// errInvalidCheckpointVote is returned if a checkpoint/epoch transition block
	// has a vote nonce set to non-zeroes.
	errInvalidCheckpointVote = errors.New("vote nonce in checkpoint block non-zero")

	// errMissingVanity is returned if a block's extra-data section is shorter than
	// 32 bytes, which is required to store the signer vanity.
	errMissingVanity = errors.New("extra-data 32 byte vanity prefix missing")

	// errMissingSignature is returned if a block's extra-data section doesn't seem
	// to contain a 65 byte secp256k1 signature.
	errMissingSignature = errors.New("extra-data 65 byte suffix signature missing")

	// errExtraSigners is returned if non-checkpoint block contain signer data in
	// their extra-data fields.
	errExtraSigners = errors.New("non-checkpoint block contains extra signer list")

	// errInvalidCheckpointSigners is returned if a checkpoint block contains an
	// invalid list of signers (i.e. non divisible by 20 bytes, or not the correct
	// ones).
	errInvalidCheckpointSigners = errors.New("invalid signer list on checkpoint block")

	// errInvalidMixDigest is returned if a block's mix digest is non-zero.
	errInvalidMixDigest = errors.New("non-zero mix digest")

	// errInvalidUncleHash is returned if a block contains an non-empty uncle list.
	errInvalidUncleHash = errors.New("non empty uncle hash")

	// errInvalidDifficulty is returned if the difficulty of a block is not either
	// of 1 or 2, or if the value does not match the turn of the signer.
	errInvalidDifficulty = errors.New("invalid difficulty")

	// ErrInvalidTimestamp is returned if the timestamp of a block is lower than
	// the previous block's timestamp + the minimum block period.
	ErrInvalidTimestamp = errors.New("invalid timestamp")

	// errInvalidVotingChain is returned if an authorization list is attempted to
	// be modified via out-of-range or non-contiguous headers.
	errInvalidVotingChain = errors.New("invalid voting chain")

	// errUnauthorized is returned if a header is signed by a non-authorized entity.
	errUnauthorized = errors.New("unauthorized")
	// errOutOfTurn is returned if a header is signed by a authorized but out of turn entity.
	errOutOfTurn = errors.New("invalid turn")

	errMasterNodesTurn = errors.New("normal peer cannot mint in master-nodes block number")

	errInvalidVoteJSON = errors.New("cannot read vote stream")
)

// SignerFn is a signer callback function to request a hash to be signed by a
// backing account.
type SignerFn func(accounts.Account, []byte) ([]byte, error)

// sigHash returns the hash which is used as input for the proof-of-authority
// signing. It is the hash of the entire header apart from the 65 byte signature
// contained at the end of the extra data.
//
// Note, the method requires the extra data to be at least 65 bytes, otherwise it
// panics. This is done to avoid accidentally using both forms (signature present
// or not), which could be abused to produce different hashes for the same header.
func sigHash(header *types.Header) (hash common.Hash) {
	hasher := sha3.NewKeccak256()

	rlp.Encode(hasher, []interface{}{
		header.ParentHash,
		header.UncleHash,
		header.Coinbase,
		header.Staker,
		header.Root,
		header.TxHash,
		header.ReceiptHash,
		header.Bloom,
		header.Difficulty,
		header.Number,
		header.GasLimit,
		header.GasUsed,
		header.Time,
		header.Extra[:len(header.Extra)-65], // Yes, this will panic if extra is too short
		header.MixDigest,
		header.Nonce,
	})
	hasher.Sum(hash[:0])
	return hash
}

// ecrecover extracts the Ethereum account address from a signed header.
func ecrecover(header *types.Header, sigcache *lru.ARCCache) (common.Address, error) {
	// If the signature's already cached, return that
	hash := header.Hash()
	if address, known := sigcache.Get(hash); known {
		return address.(common.Address), nil
	}
	// Retrieve the signature from the header extra-data
	if len(header.Extra) < extraSeal {
		return common.Address{}, errMissingSignature
	}
	signature := header.Extra[len(header.Extra)-extraSeal:]

	// Recover the public key and the Ethereum address
	pubkey, err := crypto.Ecrecover(sigHash(header).Bytes(), signature)
	if err != nil {
		return common.Address{}, err
	}
	var signer common.Address
	copy(signer[:], crypto.Keccak256(pubkey[1:])[12:])

	sigcache.Add(hash, signer)
	return signer, nil
}

// Clique is the proof-of-authority consensus engine proposed to support the
// Ethereum testnet following the Ropsten attacks.
type Xenio struct {
	config *params.XenioConfig // Consensus engine configuration parameters
	db     ethdb.Database       // Database to store and retrieve snapshot checkpoints

	recents    *lru.ARCCache // Snapshots for recent block to speed up reorgs
	signatures *lru.ARCCache // Signatures of recent blocks to speed up mining

	//proposals
	proposals map[common.Address]bool // Current list of node proposals
	//gamesContract common.Address // games contract proposal
	//usersContract common.Address // users contract proposal

	Votes map[string]Vote

	signer common.Address // Ethereum address of the signing key
	signFn SignerFn       // Signer function to authorize hashes with
	lock   sync.RWMutex   // Protects the signer fields

	//stakerSnapshot *common.StakerSnapshot

}

// New creates a Clique proof-of-authority consensus engine with the initial
// signers set to the ones provided by the user.
func New(config *params.XenioConfig, db ethdb.Database) *Xenio {
	// Set any missing consensus parameters to their defaults
	conf := *config
	if conf.Epoch == 0 {
		conf.Epoch = epochLength
	}
	if conf.Period == 0 {
		conf.Period = blockPeriod
	}
	if conf.SuperPeriod == 0 {
		conf.SuperPeriod = superPeriod
	}
	// Allocate the snapshot caches and create the engine
	recents, _ := lru.NewARC(inmemorySnapshots)
	signatures, _ := lru.NewARC(inmemorySignatures)

	return &Xenio{
		config:     &conf,
		db:         db,
		recents:    recents,
		signatures: signatures,
		proposals:  make(map[common.Address]bool),
		Votes:      make(map[string]Vote),
	}
}

// Author implements consensus.Engine, returning the Ethereum address recovered
// from the signature in the header's extra-data section.
func (c *Xenio) Author(header *types.Header) (common.Address, error) {
	return ecrecover(header, c.signatures)
}

// VerifyHeader checks whether a header conforms to the consensus rules.
func (c *Xenio) VerifyHeader(chain consensus.ChainReader, header *types.Header, seal bool) error {
	return c.verifyHeader(chain, header, nil)
}

// VerifyHeaders is similar to VerifyHeader, but verifies a batch of headers. The
// method returns a quit channel to abort the operations and a results channel to
// retrieve the async verifications (the order is that of the input slice).
func (c *Xenio) VerifyHeaders(chain consensus.ChainReader, headers []*types.Header, seals []bool) (chan<- struct{}, <-chan error) {
	abort := make(chan struct{})
	results := make(chan error, len(headers))

	go func() {
		for i, header := range headers {
			err := c.verifyHeader(chain, header, headers[:i])

			select {
			case <-abort:
				return
			case results <- err:
			}
		}
	}()
	return abort, results
}

// verifyHeader checks whether a header conforms to the consensus rules.The
// caller may optionally pass in a batch of parents (ascending order) to avoid
// looking those up from the database. This is useful for concurrently verifying
// a batch of new headers.
func (c *Xenio) verifyHeader(chain consensus.ChainReader, header *types.Header, parents []*types.Header) error {
	if header.Number == nil {
		return errUnknownBlock
	}
	number := header.Number.Uint64()

	// Don't waste time checking blocks from the future
	if header.Time.Cmp(big.NewInt(time.Now().Unix())) > 0 {
		return consensus.ErrFutureBlock
	}
	// Checkpoint blocks need to enforce zero beneficiary
	checkpoint := (number % c.config.Epoch) == 0
	if checkpoint && header.Coinbase != (common.Address{}) {
		return errInvalidCheckpointBeneficiary
	}
	// Nonces must be 0x00..0 or 0xff..f, zeroes enforced on checkpoints
	if !bytes.Equal(header.Nonce[:], nonceAuthVote) && !bytes.Equal(header.Nonce[:], nonceDropVote) {
		return errInvalidVote
	}
	if checkpoint && !bytes.Equal(header.Nonce[:], nonceDropVote) {
		return errInvalidCheckpointVote
	}
	// Check that the extra-data contains both the vanity and signature
	if len(header.Extra) < extraVanity {
		return errMissingVanity
	}
	if len(header.Extra) < extraVanity+extraSeal {
		return errMissingSignature
	}
	// Ensure that the extra-data contains a signer list on checkpoint, but none otherwise
	signersBytes := len(header.Extra) - extraVanity - extraSeal
	if !checkpoint && signersBytes != 0 {
		return errExtraSigners
	}
	if checkpoint && signersBytes%common.AddressLength != 0 {
		return errInvalidCheckpointSigners
	}
	// Ensure that the mix digest is zero as we don't have fork protection currently
	if header.MixDigest != (common.Hash{}) {
		return errInvalidMixDigest
	}
	// Ensure that the block doesn't contain any uncles which are meaningless in PoA
	if header.UncleHash != uncleHash {
		return errInvalidUncleHash
	}
	// Ensure that the block's difficulty is meaningful (may not be correct at this point)
	if number > 0 {
		if header.Difficulty == nil || (header.Difficulty.Cmp(diffInTurn) != 0 && header.Difficulty.Cmp(diffNoTurn) != 0) {
			return errInvalidDifficulty
		}
	}
	// All basic checks passed, verify cascading fields
	return c.verifyCascadingFields(chain, header, parents)
}

// verifyCascadingFields verifies all the header fields that are not standalone,
// rather depend on a batch of previous headers. The caller may optionally pass
// in a batch of parents (ascending order) to avoid looking those up from the
// database. This is useful for concurrently verifying a batch of new headers.
func (c *Xenio) verifyCascadingFields(chain consensus.ChainReader, header *types.Header, parents []*types.Header) error {
	// The genesis block is the always valid dead-end
	number := header.Number.Uint64()
	if number == 0 {
		return nil
	}
	// Ensure that the block's timestamp isn't too close to it's parent
	var parent *types.Header
	if len(parents) > 0 {
		parent = parents[len(parents)-1]
	} else {
		parent = chain.GetHeader(header.ParentHash, number-1)
	}
	if parent == nil || parent.Number.Uint64() != number-1 || parent.Hash() != header.ParentHash {
		return consensus.ErrUnknownAncestor
	}
	if parent.Time.Uint64()+c.config.Period > header.Time.Uint64() {
		return ErrInvalidTimestamp
	}
	// Retrieve the snapshot needed to verify this header and cache it
	snap, err := c.snapshot(chain, number-1, header.ParentHash, parents)
	if err != nil {
		return err
	}
	// If the block is a checkpoint block, verify the signer list
	if number%c.config.Epoch == 0 {
		signers := make([]byte, len(snap.MasterNodes)*common.AddressLength)
		for i, signer := range snap.masterNodes() {
			copy(signers[i*common.AddressLength:], signer[:])
		}
		extraSuffix := len(header.Extra) - extraSeal
		if !bytes.Equal(header.Extra[extraVanity:extraSuffix], signers) {
			return errInvalidCheckpointSigners
		}
	}
	// All basic checks passed, verify the seal and return
	return c.verifySeal(chain, header, parents)
}

// snapshot retrieves the authorization snapshot at a given point in time.
func (c *Xenio) snapshot(chain consensus.ChainReader, number uint64, hash common.Hash, parents []*types.Header) (*Snapshot, error) {
	// Search for a snapshot in memory or on disk for checkpoints
	var (
		headers []*types.Header
		snap    *Snapshot
	)
	for snap == nil {
		// If an in-memory snapshot was found, use that
		if s, ok := c.recents.Get(hash); ok {
			snap = s.(*Snapshot)
			break
		}
		// If an on-disk checkpoint snapshot can be found, use that
		if number%checkpointInterval == 0 {
			if s, err := loadSnapshot(c.config, c.signatures, c.db, hash); err == nil {
				log.Trace("Loaded voting snapshot form disk", "number", number, "hash", hash)
				snap = s
				break
			}
		}
		// If we're at block zero, make a snapshot
		if number == 0 {
			genesis := chain.GetHeaderByNumber(0)
			if err := c.VerifyHeader(chain, genesis, false); err != nil {
				return nil, err
			}
			signers := make([]common.Address, (len(genesis.Extra)-extraVanity-extraSeal)/common.AddressLength)
			for i := 0; i < len(signers); i++ {
				copy(signers[i][:], genesis.Extra[extraVanity+i*common.AddressLength:])
			}
			snap = newSnapshot(c.config, c.signatures, 0, genesis.Hash(), signers)
			if err := snap.store(c.db); err != nil {
				return nil, err
			}
			log.Trace("Stored genesis voting snapshot to disk")
			break
		}
		// No snapshot for this header, gather the header and move backward
		var header *types.Header
		if len(parents) > 0 {
			// If we have explicit parents, pick from there (enforced)
			header = parents[len(parents)-1]
			if header.Hash() != hash || header.Number.Uint64() != number {
				return nil, consensus.ErrUnknownAncestor
			}
			parents = parents[:len(parents)-1]
		} else {
			// No explicit parents (or no more left), reach out to the database
			header = chain.GetHeader(hash, number)
			if header == nil {
				return nil, consensus.ErrUnknownAncestor
			}
		}
		headers = append(headers, header)
		number, hash = number-1, header.ParentHash
	}
	// Previous snapshot found, apply any pending headers on top of it
	for i := 0; i < len(headers)/2; i++ {
		headers[i], headers[len(headers)-1-i] = headers[len(headers)-1-i], headers[i]
	}
	snap, err := snap.apply(headers)
	if err != nil {
		log.Error("snapshot apply headers error (xenio.go snapshot function")
		return nil, err
	}
	c.recents.Add(snap.Hash, snap)

	//// Retrieve and update Signer List
	//genesis := chain.GetHeaderByNumber(0)
	//if err := c.VerifyHeader(chain, genesis, false); err != nil {
	//	return nil, err
	//}
	//signers := make([]common.Address, (len(genesis.Extra)-extraVanity-extraSeal)/common.AddressLength)
	//for i := 0; i < len(signers); i++ {
	//	copy(signers[i][:], genesis.Extra[extraVanity+i*common.AddressLength:])
	//}
	//snap = updateSigners(snap, c.config, number, genesis.Time, signers)

	// If we've generated a new checkpoint snapshot, save to disk
	if snap.Number%checkpointInterval == 0 && len(headers) > 0 {
		if err = snap.store(c.db); err != nil {
			return nil, err
		}
		log.Trace("Stored voting snapshot to disk", "number", snap.Number, "hash", snap.Hash)
	}
	return snap, err
}

// VerifyUncles implements consensus.Engine, always returning an error for any
// uncles as this consensus mechanism doesn't permit uncles.
func (c *Xenio) VerifyUncles(chain consensus.ChainReader, block *types.Block) error {
	if len(block.Uncles()) > 0 {
		return errors.New("uncles not allowed")
	}
	return nil
}

// VerifySeal implements consensus.Engine, checking whether the signature contained
// in the header satisfies the consensus protocol requirements.
func (c *Xenio) VerifySeal(chain consensus.ChainReader, header *types.Header) error {
	return c.verifySeal(chain, header, nil)
}

// verifySeal checks whether the signature contained in the header satisfies the
// consensus protocol requirements. The method accepts an optional list of parent
// headers that aren't yet part of the local blockchain to generate the snapshots
// from.
func (c *Xenio) verifySeal(chain consensus.ChainReader, header *types.Header, parents []*types.Header) error {
	// Verifying the genesis block is not supported
	number := header.Number.Uint64()
	if number == 0 {
		return errUnknownBlock
	}

	// Retrieve the snapshot needed to verify this header and cache it
	snap, err := c.snapshot(chain, number-1, header.ParentHash, parents)
	if err != nil {
		return err
	}

	// Resolve the authorization key and check against signers
	signer, err := ecrecover(header, c.signatures)
	if err != nil {
		return err
	}

	// Get authorised node. Check whether the signer address corresponds to a validated master, or staking node.
	signingNode, isAuthorised := snap.getSigningNode(signer)
	if !isAuthorised {
		return errUnauthorized
	}

	// Check whether the authorised node is next in turn. Give out-of-turn error if the signing node does not contain the next in line block.
	if !signingNode.isInTurn(snap) {
		if !signingNode.IsMasterNode {
			return errOutOfTurn
		} else {
			headerTime := time.Unix(chain.CurrentHeader().Time.Int64(), 0).UTC()
			// During FastSync the header in the chain may be more recent than the block we are checking
			if headerTime.After(time.Unix(header.Time.Int64(), 0).UTC()) {
				headerTime = time.Unix(header.Time.Int64(), 0).UTC()
			}
			// Verify past block's timestamp
			if number != chain.CurrentHeader().Number.Uint64() {
				headerTime = headerTime.Add(time.Duration(chain.Config().Xenio.Period) * time.Second).UTC()
			}
			if headerTime.UTC().Unix() >= time.Now().UTC().Unix() {
				return ErrInvalidTimestamp
			}
		}

	}
	//verify vote stream
	if len(header.Votes) > 0{
		voteData := make(map[string]Vote, 0)
		if err := json.Unmarshal(header.Votes, &voteData); err != nil {
			return errInvalidVoteJSON
		}
	}
	return nil
}

// Prepare implements consensus.Engine, preparing all the consensus fields of the
// header for running the transactions on top.
func (c *Xenio) Prepare(chain consensus.ChainReader, header *types.Header) error {
	// If the block isn't a checkpoint, cast a random vote (good enough for now)
	header.Coinbase = common.Address{}
	header.Nonce = types.BlockNonce{}

	number := header.Number.Uint64()

	// Assemble the voting snapshot to check which votes make sense
	snap, err := c.snapshot(chain, number-1, header.ParentHash, nil)
	if err != nil {
		return err
	}
	/*if number%c.config.Epoch != 0 {
		c.lock.RLock()

		// Gather all the proposals that make sense voting on
		addresses := make([]common.Address, 0, len(c.proposals))
		for address, authorize := range c.proposals {
			if snap.validVote(address, authorize) {
				addresses = append(addresses, address)
			}
		}
		// If there's pending proposals, cast a vote on them
		if len(addresses) > 0 {
			header.Coinbase = addresses[rand.Intn(len(addresses))]
			if c.proposals[header.Coinbase] {
				copy(header.Nonce[:], nonceAuthVote)
			} else {
				copy(header.Nonce[:], nonceDropVote)
			}
		}
		c.lock.RUnlock()
	}*/
	// Set the correct difficulty
	header.Difficulty = diffNoTurn
	if snap.inturn(header.Number.Uint64(), c.signer) {
		header.Difficulty = diffInTurn
	}
	// Ensure the extra data has all it's components
	if len(header.Extra) < extraVanity {
		header.Extra = append(header.Extra, bytes.Repeat([]byte{0x00}, extraVanity-len(header.Extra))...)
	}
	header.Extra = header.Extra[:extraVanity]

	if number%c.config.Epoch == 0 {
		for _, signer := range snap.masterNodes() {
			header.Extra = append(header.Extra, signer[:]...)
		}
	}
	header.Extra = append(header.Extra, make([]byte, extraSeal)...)

	// Mix digest is reserved for now, set to empty
	header.MixDigest = common.Hash{}

	// Ensure the timestamp has the correct delay
	parent := chain.GetHeader(header.ParentHash, number-1)
	if parent == nil {
		return consensus.ErrUnknownAncestor
	}
	header.Time = new(big.Int).Add(parent.Time, new(big.Int).SetUint64(c.config.Period))
	if header.Time.Int64() < time.Now().Unix() {
		header.Time = big.NewInt(time.Now().Unix())
	}
	//PoN Rewards
	if number >= 1 {
		if common.StakerSnapShot != nil {
			common.StakerSnapShot.Stakers.Range(
				func(address, staker interface{}) bool {
					if !StakerExpired(address.(common.Address)) {
						header.RewardList = append(header.RewardList, address.(common.Address))
					}
					return true
				})
				}
			}
	return nil
}

// Finalize implements consensus.Engine, ensuring no uncles are set, nor block
// rewards given, and returns the final block.
func (c *Xenio) Finalize(chain consensus.ChainReader, header *types.Header, state *state.StateDB, txs []*types.Transaction, uncles []*types.Header, receipts []*types.Receipt) (*types.Block, error) {
	AccumulateRewards(state, header, txs, uncles, receipts)
	header.Root = state.IntermediateRoot(chain.Config().IsEIP158(header.Number))
	header.UncleHash = types.CalcUncleHash(nil)

	// Assemble and return the final block for sealing
	return types.NewBlock(header, txs, nil, receipts), nil
}

// Authorize injects a private key into the consensus engine to mint new blocks
// with.
func (c *Xenio) Authorize(signer common.Address, signFn SignerFn) {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.signer = signer
	c.signFn = signFn
}

// Seal implements consensus.Engine, attempting to create a sealed block using
// the local signing credentials.
func (c *Xenio) Seal(chain consensus.ChainReader, block *types.Block, stop <-chan struct{}) (*types.Block, error) {
	header := block.Header()

	// Sealing the genesis block is not supported
	number := header.Number.Uint64()
	if number == 0 {
		return nil, errUnknownBlock
	}
	// Don't hold the signer fields for the entire sealing procedure
	c.lock.RLock()
	signer, signFn := c.signer, c.signFn
	c.lock.RUnlock()

	// Bail out if we're unauthorized to sign a block
	snap, err := c.snapshot(chain, number-1, header.ParentHash, nil)
	if err != nil {
		return nil, err
	}
	ca := common.Address{}
	if signer == ca{
		return nil, errUnauthorized
	}
	// Get authorised node. Check whether the signer address corresponds to a validated master, or staking node.
	signingNode, isAuthorised := snap.getSigningNode(signer)
	if !isAuthorised { return nil, errUnauthorized }

	// Estimate delay time by adding a small amount of noise
	estimatedTime := time.Unix(chain.CurrentHeader().Time.Int64(),0).Add(time.Duration(chain.Config().Xenio.Period)*time.Second)
	// Checks whether the authorised node is next in turn. Gives out-of-turn error if the signing node does not contain the next in line block.
	if !signingNode.isInTurn(snap){

		if !signingNode.IsMasterNode{ return nil,nil }

		//Check whether estimated time exceeds current time
		estimatedTime = estimatedTime.Add(wiggleTime)
		if estimatedTime.Unix() >= time.Now().Unix(){ return nil,nil }
	}

	log.Warn("We are next to mint a block")

	delayTime := estimatedTime.Unix() - time.Now().Unix()
	remainingSeconds, _ := json.Marshal(delayTime)
	if delayTime < 0 {
		remainingSeconds, _ = json.Marshal(-1 * delayTime)
		log.Info("Mining block is delayed by " + string(remainingSeconds) + " seconds")
	}else{
		log.Info("Mining block in " + string(remainingSeconds) + " seconds")
	}

	select {
	case <-stop:
		return nil, nil
	case <-time.After(time.Duration(delayTime)*time.Second):
	}

	isSuperBlock := snap.changeSuperBlockHeaders(signingNode,header)
	if isSuperBlock { //if it is a superblock, update the list

		// create a node list with master and staking nodes
		nodes := make(map[common.Address]Signer, 0) // new list

		// initialise nodes list from snapshot
		// add master nodes to nodes list
		for address, node := range snap.MasterNodes{
			//Deep copy
			var masterNode Signer
			masterNode.BlockNumber = node.BlockNumber
			masterNode.IsMasterNode = node.IsMasterNode
			nodes[address] = masterNode
		}
		// update block numbers of nodes list
		// for master nodes
		master_block_number := common.MasterBlockIncrement * (header.Number.Uint64()/common.MasterBlockIncrement) // intialisation, division rounds it down
		for address, node := range snap.MasterNodes {
			// create a new node and add it to the list
			var masterNode Signer
			masterNode.BlockNumber = node.BlockNumber
			if masterNode.BlockNumber == nil{
				masterNode.BlockNumber = make([]uint64, 1)
			}
			if masterNode.BlockNumber[0] <= header.Number.Uint64() { // pass, if there is already a future block in the list
				for {
					master_block_number += common.MasterBlockIncrement
					if !isDuplicated(master_block_number, nodes) {
						break
					}
				}
				masterNode.BlockNumber[0] = master_block_number
			}
			masterNode.IsMasterNode = true  // mark it as a master node
			nodes[address] = masterNode // add it to the list
		}

		//get max block number
		max_block_number := getMaxBlockNumber(nodes)

		// for staking nodes
		stakerSnap := common.StakerSnapShot // keep order of staker's fixed
		if stakerSnap != nil {
			if snap.stakersExist(stakerSnap) { //checks if the list has only master nodes
				current_block_number := header.Number.Uint64() //initialise block number
				for {
					if !stakerSnap.HasStakers() { break }
					if current_block_number >= max_block_number { break	}

					common.StakerSnapShot.Stakers.Range(
						func(address, staker interface{}) bool {
							// Skip this node, if it is already in the master nodes list
							if _, isMasterNode := snap.MasterNodes[address.(common.Address)]; isMasterNode /* StakerExpired(address.(common.Address)) */ {
								return true
							}
							// Add a new node to the nodes list
							stakingNode := nodes[address.(common.Address)]
							current_block_number++
							if isDuplicated(current_block_number, nodes) {
								current_block_number++
							}
							if current_block_number >= max_block_number {
								return false
							}
							stakingNode.BlockNumber = append(stakingNode.BlockNumber, current_block_number) // update block numbers
							stakingNode.IsMasterNode = false                                                // mark it as regular
							nodes[address.(common.Address)] = stakingNode
							return true
						})
				}
			}else{
				log.Warn("No staking nodes found in the stakerSnapshot")
			}
		}
		if (len(nodes)) > 0 {
			blob, _ := json.Marshal(nodes)
			header.SuperBlock = blob
			//	log.Warn(string(blob))
		}
	}

	api := API{}
	test := api.GetUserAddresses(snap.UsersContractAddress)
	fmt.Println("Users' Addresses in Seal:", test)

	//see whats for voting and autocast our vote
	if len(snap.NewVotes) > 0{
		for hash, vote := range snap.NewVotes {
			if vote.VoteType == MasterNode{
				if !signingNode.IsMasterNode { continue }
				ourhash := common.GetMD5Hash(vote.Address.String() + signer.String())// to see if the vote is ours
				if ourhash == hash {
					continue
				}
				var newVote Vote
				newVote.VoteType = MasterNode
				if HasCoins(vote.Address, requiredCoins, currentState) {
					log.Warn("automated upvote for a gameserver")
					newVote.Authorize = true

				}else{
					log.Warn("automated downvote for a gameserver")
					newVote.Authorize = false
				}
				newVote.Address = vote.Address
				c.Votes[ourhash] = newVote
			}
		}
	}


	//transmit votes to the network
	if len(c.Votes) > 0 && currentState != nil {
		transmitArray := make(map[string]Vote)
		for key, vote := range c.Votes {
			if vote.VoteType == MasterNode {
				if !signingNode.IsMasterNode {
					log.Warn("Discarded vote because normal peers are not allowed to vote for game servers")
					continue
				}
				//1000 coins limit for masternodes
				if !HasCoins(vote.Address, requiredCoins, currentState) {
					log.Warn("Discarded vote because candidate doesn't have enough coin balance")
					continue
				}
			}
			vote.Block = header.Number.Int64()
			transmitArray[key] = vote
		}
		c.Votes = transmitArray //replace arrays
		if len(c.Votes) > 0 {
			blob, verr := json.Marshal(c.Votes)
			if verr == nil {
				header.Votes = blob
			} else {
				log.Warn("Votes List: " + verr.Error())
				c.Votes = make(map[string]Vote)

			}
		}
	}


	// Sign all the things!
	sighash, err := signFn(accounts.Account{Address: signer}, sigHash(header).Bytes())
	if err != nil {
		return nil, err
	}
	//clear votes
	c.Votes = make(map[string]Vote)
	copy(header.Extra[len(header.Extra)-extraSeal:], sighash)

	return block.WithSeal(header), nil
}

// APIs implements consensus.Engine, returning the user facing RPC API to allow
// controlling the signer voting.
func (c *Xenio) APIs(chain consensus.ChainReader) []rpc.API {
	return []rpc.API{{
		Namespace: "xenio",
		Version:   "1.0",
		Service:   &API{chain: chain, xenio: c},
		Public:    false,
	}}
}

func (c *Xenio) State(state *state.StateDB) {
	currentState = state
}

func AccumulateRewards(state *state.StateDB, header *types.Header, txs []*types.Transaction, uncles []*types.Header, receipts []*types.Receipt) {
	reward := calculateReward(txs, receipts)

	//r := new(big.Int)
	/*for _, uncle := range uncles {
		r.Add(uncle.Number, big8)
		r.Sub(r, header.Number)
		r.Mul(r, blockReward)
		r.Div(r, big8)
		state.AddBalance(uncle.Coinbase, r)

		r.Div(blockReward, big32)
		reward.Add(reward, r)
	}*/

	//log.Warn(header.Number.String() + " blockReward " + reward.String() + " weis")

	if len(header.RewardList) >= 1 &&  reward.Int64() != 0 {
		// fee=sum(fees)/stakers
		reward.Div(reward, big.NewInt(int64(len(header.RewardList))))
		for _, address := range header.RewardList {
			//if HasCoins(address, state) {
			state.AddBalance(address, reward)
			//cb, _ := json.Marshal(address)
			//log.Warn(string(cb) + " rewarded " + reward.String() + " weis")
			//}
		}
	}
}

func calculateReward(txs []*types.Transaction, receipts []*types.Receipt) *big.Int {
	reward := new(big.Int)
	if receipts != nil && len(receipts) > 0 {
		for i := range receipts {
			if receipts[i].Status == ReceiptStatusSuccessful {
				reward.Add(reward, receipts[i].GasUsed)
			}
		}
	}
	return reward
}

func CurrentIPCEndpoint(IPCEndpoint string) {
	currentIPCEndpoint = IPCEndpoint
}
