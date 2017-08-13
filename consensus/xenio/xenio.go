package xenio

import (
	"github.com/xenioplatform/go-xenio/common"
	"github.com/xenioplatform/go-xenio/params"
	"github.com/xenioplatform/go-xenio/ethdb"
)

var (
	epochLength = uint64(30000) 	// Default number of blocks after which to checkpoint and reset the pending votes
	blockPeriod = uint64(15)    	// Default minimum difference between two consecutive block's timestamps
	superBlockPeriod = uint64(900)	// Default minimum difference between two consecutive super block's timestamps
)

type Xenio struct {
	config *params.XenioConfig // Consensus engine configuration parameters
	db     ethdb.Database       // Database to store and retrieve snapshot checkpoints


	proposals map[common.Address]bool // Current list of proposals we are pushing

	signer common.Address // Ethereum address of the signing key
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