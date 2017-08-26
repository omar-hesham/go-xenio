package xenio

import (
	"github.com/xenioplatform/go-xenio/common"
	"github.com/xenioplatform/go-xenio/ethdb"
	"encoding/json"
	"math/big"
	"github.com/xenioplatform/go-xenio/log"
)

func newStakerSnapshot(stakers []common.Address, blockNumber *big.Int) *common.StakerSnapshot {
	snap := &common.StakerSnapshot{
		Stakers:  make([]common.Address,len(stakers)),
	}
	for _, staker := range stakers {
		var ca common.Address
		if staker == ca{
			continue
		}
		st, _ := json.Marshal(staker)
		log.Info("adding " + string(st) + " to snapshot")
		snap.Stakers = append(snap.Stakers, staker)
	}
	//snap.BlockNumber.Add(blockNumber, big.NewInt(1))
	snap.BlockNumber = blockNumber
	return snap
}


// loadSnapshot loads an existing snapshot from the database.
// we might don't need this, only memory snapshots might be enough
/*func loadStakerSnapshot(db ethdb.Database) (*common.StakerSnapshot, error) {
	blob, err := db.Get([]byte("xenioStakers-"))
	if err != nil {
		return nil, err
	}
	snap := new(Snapshot)
	if err := json.Unmarshal(blob, snap); err != nil {
		return nil, err
	}

	return snap, nil
}

// store inserts the snapshot into the database.
func (self *API) storeStakersSnapshot(db ethdb.Database) error {
	blob, err := json.Marshal(self)
	if err != nil {
		return err
	}
	return db.Put([]byte("xenioStakers-"), blob)
}*/

// store inserts the snapshot into the database.
func (self *API) deleteStakersSnapshot(db ethdb.Database) error {
	log.Info("deleted")
	return db.Delete([]byte("xenioStakers-"))
}

// cast adds new stakers to the list.
func (self *API) stakerCast(stakers []common.Address) bool {
	// Cast the stakers into an existing or new list
	newStakerList  := make([]common.Address,len(stakers))
	if self.xenio.stakerSnapshot.Stakers != nil{
		newStakerList = make([]common.Address,len(stakers)+len(self.xenio.stakerSnapshot.Stakers))
	}
	for _, staker := range self.xenio.stakerSnapshot.Stakers {
		newStakerList = append(self.xenio.stakerSnapshot.Stakers, staker)
	}
	for _, staker := range stakers {
		newStakerList = append(self.xenio.stakerSnapshot.Stakers, staker)
	}
	self.xenio.stakerSnapshot.Stakers = newStakerList
	return true
}

func (self *API) stakerExists(staker common.Address) bool {
	for _, a := range self.xenio.stakerSnapshot.Stakers {
		if a == staker {
			return true
		}
	}
	return false
}


func (api *API) GetStakerSnapshot() (*common.StakerSnapshot, error) {
	if api.xenio.stakerSnapshot == nil || api.chain.CurrentHeader().Number != api.xenio.stakerSnapshot.BlockNumber{
		signers := make([]common.Address,0)
		api.xenio.stakerSnapshot = newStakerSnapshot(signers,api.chain.CurrentHeader().Number)
	}

	return api.xenio.stakerSnapshot, nil
}

func (api *API) AddStakerToSnapshot(address common.Address) (bool, error) {
	if api.xenio.stakerSnapshot == nil || api.chain.CurrentHeader().Number != api.xenio.stakerSnapshot.BlockNumber {
		api.GetStakerSnapshot()
	}
	for _, st := range api.xenio.stakerSnapshot.Stakers {
		if st == address {
			return false, nil
		}
	}
	api.xenio.stakerSnapshot.Stakers = append(api.xenio.stakerSnapshot.Stakers,address)
	return true, nil
}