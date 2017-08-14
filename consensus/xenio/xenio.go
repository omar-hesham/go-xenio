package xenio

import (
	"github.com/xenioplatform/go-xenio/common"
	"github.com/xenioplatform/go-xenio/params"
	"github.com/xenioplatform/go-xenio/ethdb"
	"github.com/xenioplatform/go-xenio/core/types"
	"github.com/xenioplatform/go-xenio/consensus"
	"github.com/xenioplatform/go-xenio/core/state"
	"math/big"
	"sync"
)

var (
	blockReward *big.Int = big.NewInt(5e+18)
	epochLength = uint64(30000) 	// Default number of blocks after which to checkpoint and reset the pending votes
	blockPeriod = uint64(15)    	// Default minimum difference between two consecutive block's timestamps
	superBlockPeriod = uint64(900)	// Default minimum difference between two consecutive super block's timestamps
	// Some weird constants to avoid constant memory allocs for them.

	big8  = big.NewInt(8)
	big32 = big.NewInt(32)
)

type Xenio struct {
	config *params.XenioConfig // Consensus engine configuration parameters
	db     ethdb.Database       // Database to store and retrieve snapshot checkpoints


	stakers map[common.Address]bool // Current list of stakers

	signer common.Address // Ethereum address of the signing key

	lock   sync.RWMutex   // its a mutex
}

// New creates a Clique proof-of-authority consensus engine with the initial
// signers set to the ones provided by the user.
func New(config *params.XenioConfig, db ethdb.Database) *Xenio {
	// Set any missing consensus parameters to their defaults
	conf := *config
	if conf.SuperPeriod == 0 {
		conf.SuperPeriod = superBlockPeriod
	}
	if conf.Period == 0 {
		conf.Period = blockPeriod
	}

	return &Xenio{
		config:     &conf,
		db:         db,
		//recents:    recents,
		//signatures: signatures,
		//proposals:  make(map[common.Address]bool),
	}
}

func (ethash *Xenio) Author(header *types.Header) (common.Address, error) {
	return header.Coinbase, nil
}

func calcDiffucultyXenio(time uint64, parent *types.Header) *big.Int {
	return big.NewInt(0x4000)

}

// Prepare implements consensus.Engine, initializing the difficulty field of a
// header to conform to the ethash protocol. The changes are done inline.
func (ethash *Xenio) Prepare(chain consensus.ChainReader, header *types.Header) error {
	parent := chain.GetHeader(header.ParentHash, header.Number.Uint64()-1)
	if parent == nil {
		return consensus.ErrUnknownAncestor
	}
	header.Difficulty = calcDiffucultyXenio(header.Time.Uint64(), parent)

	return nil
}

// Finalize implements consensus.Engine, accumulating the block and uncle rewards,
// setting the final state and assembling the block.
func (ethash *Xenio) Finalize(chain consensus.ChainReader, header *types.Header, state *state.StateDB, txs []*types.Transaction, uncles []*types.Header, receipts []*types.Receipt) (*types.Block, error) {
	// Accumulate any block and uncle rewards and commit the final state root
	AccumulateRewards(state, header, uncles)
	header.Root = state.IntermediateRoot(chain.Config().IsEIP158(header.Number))

	// Header seems complete, assemble into a block and return
	return types.NewBlock(header, txs, uncles, receipts), nil
}

// AccumulateRewards credits the coinbase of the given block with the mining
// reward. The total reward consists of the static block reward and rewards for
// included uncles. The coinbase of each uncle block is also rewarded.
func AccumulateRewards(state *state.StateDB, header *types.Header, uncles []*types.Header) {
	reward := new(big.Int).Set(blockReward)
	r := new(big.Int)
	for _, uncle := range uncles {
		r.Add(uncle.Number, big8)
		r.Sub(r, header.Number)
		r.Mul(r, blockReward)
		r.Div(r, big8)
		state.AddBalance(uncle.Coinbase, r)

		r.Div(blockReward, big32)
		reward.Add(reward, r)
	}
	state.AddBalance(header.Coinbase, reward)
}