package xenio

import (
	"github.com/xenioplatform/go-xenio/common"
	//"github.com/xenioplatform/go-xenio/ethdb"
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
}

// store inserts the snapshot into the database.
func (self *API) deleteStakersSnapshot(db ethdb.Database) error {
	log.Info("deleted")
	return db.Delete([]byte("xenioStakers-"))
}
*/
// cast adds new stakers to the list.
func (self *API) stakerCast(stakers []common.Address) bool {
	// Cast the stakers into an existing or new list
	newStakerList  := make([]common.Address,len(stakers))
	if common.StakerSnapShot.Stakers != nil{
		newStakerList = make([]common.Address,len(stakers)+len(common.StakerSnapShot.Stakers))
	}
	for _, staker := range common.StakerSnapShot.Stakers {
		newStakerList = append(common.StakerSnapShot.Stakers, staker)
	}
	for _, staker := range stakers {
		newStakerList = append(common.StakerSnapShot.Stakers, staker)
	}
	common.StakerSnapShot.Stakers = newStakerList
	return true
}

func (self *API) stakerExists(staker common.Address) bool {
	for _, a := range common.StakerSnapShot.Stakers {
		if a == staker {
			return true
		}
	}
	return false
}

func (api *API) GetStakerSnapshot() (*common.StakerSnapshot, error) {
	//var currentHeaderNumber *big.Int
	//
	//if api.chain.CurrentHeader() != nil {
	//	//currentHeaderNumber = api.chain.CurrentHeader().Number.Add(api.chain.CurrentHeader().Number, big.NewInt(20))
	//	currentHeaderNumber = api.chain.CurrentHeader().Number
	//}

	if common.StakerSnapShot == nil || api.chain.CurrentHeader().Number.Cmp(common.StakerSnapShot.BlockNumber) != -1 {
		signers := make([]common.Address, 0)

		common.StakerSnapShot = newStakerSnapshot(signers, api.chain.CurrentHeader().Number)
	}

	return common.StakerSnapShot, nil
}

func (api *API) AddStakerToSnapshot(address common.Address) (bool, error) {
	if common.StakerSnapShot == nil || api.chain.CurrentHeader().Number != common.StakerSnapShot.BlockNumber {
		api.GetStakerSnapshot()
	}
	for _, st := range common.StakerSnapShot.Stakers {
		if st == address {
			return false, nil
		}
	}
	common.StakerSnapShot.Stakers = append(common.StakerSnapShot.Stakers,address)
	return true, nil
}