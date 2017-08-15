package xenio

import (
	"errors"
	"math/big"
	"time"
	"github.com/xenioplatform/go-xenio/common/hexutil"
	"github.com/xenioplatform/go-xenio/core/types"
)

var(
	nonceAuthVote = hexutil.MustDecode("0xffffffffffffffff") // Magic nonce number to vote on adding a new signer
	nonceDropVote = hexutil.MustDecode("0x0000000000000000") // Magic nonce number to vote on removing a signer.

)

var ( //errors
	errUnknownBlock = errors.New("unknown block")
	errGenesisBlock = errors.New("genesis block")
	errInvalidVotingChain = errors.New("invalid chain")
	// errUnauthorized is returned if a header is signed by a non-authorized entity.
	errUnauthorized = errors.New("unauthorized")
	// errInvalidVote is returned if a nonce value is something else that the two
	// allowed constants of 0x00..0 or 0xff..f.
	errInvalidVote = errors.New("vote nonce not 0x00..0 or 0xff..f")
	errMissingSignature = errors.New("extra-data 65 byte suffix signature missing")
	errInvalidCheckpointBeneficiary = errors.New("beneficiary in checkpoint block non-zero")
	errInvalidCheckpointVote = errors.New("vote nonce in checkpoint block non-zero")
	errMissingVanity = errors.New("extra-data 32 byte vanity prefix missing")
	errExtraSigners = errors.New("non-checkpoint block contains extra signer list")
	errInvalidCheckpointSigners = errors.New("invalid signer list on checkpoint block")
	errInvalidMixDigest = errors.New("non-zero mix digest")
	errInvalidUncleHash = errors.New("non empty uncle hash")
	errInvalidDifficulty = errors.New("invalid difficulty")
	ErrInvalidTimestamp = errors.New("invalid timestamp")
)

var (
	blockReward *big.Int = big.NewInt(5e+18)
	epochLength = uint64(30000) 	// Default number of blocks after which to checkpoint and reset the pending votes
	blockPeriod = uint64(15)    	// Default minimum difference between two consecutive block's timestamps
	superBlockPeriod = uint64(900)	// Default minimum difference between two consecutive super block's timestamps
	extraSeal   = 65
	extraVanity = 32
	big8  = big.NewInt(8) //performance helper
	big32 = big.NewInt(32) //performance helper
	diffInTurn = big.NewInt(2) // Block difficulty for in-turn signatures
	diffNoTurn = big.NewInt(1) // Block difficulty for out-of-turn signatures
	uncleHash = types.CalcUncleHash(nil) // Always Keccak256(RLP([])) as uncles are meaningless outside of PoW.
)

const (
	checkpointInterval = 1024 // Number of blocks after which to save the vote snapshot to the database
	//inmemorySnapshots  = 128  // Number of recent vote snapshots to keep in memory
	//inmemorySignatures = 4096 // Number of recent block signatures to keep in memory

	wiggleTime = 500 * time.Millisecond // Random delay (per signer) to allow concurrent signers
)