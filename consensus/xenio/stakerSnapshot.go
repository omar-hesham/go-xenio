package xenio

import (
	"encoding/json"
	"time"

	"github.com/xenioplatform/go-xenio/common"
	//"github.com/xenioplatform/go-xenio/ethdb"
	"github.com/xenioplatform/go-xenio/log"
)

func newStakerSnapshot(stakers []common.Staker) *common.StakerSnapshot {
	snap := &common.StakerSnapshot{
		Stakers: make([]common.Staker, len(stakers)),
	}
	for _, staker := range stakers {
		var ca common.Address
		if staker.Address == ca {
			continue
		}
		st, _ := json.Marshal(staker.Address)
		log.Info("adding " + string(st) + " to snapshot")
		snap.Stakers = append(snap.Stakers, staker)
	}
	snap.Created = time.Now()
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
func (self *API) stakerCast(stakers []common.Staker) bool {
	// Cast the stakers into an existing or new list
	newStakerList := make([]common.Staker, len(stakers))
	if common.StakerSnapShot.Stakers != nil {
		newStakerList = make([]common.Staker, len(stakers)+len(common.StakerSnapShot.Stakers))
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

func (self *API) stakerExists(staker common.Staker) bool {
	for _, a := range common.StakerSnapShot.Stakers {
		if a.Address == staker.Address {
			return true
		}
	}
	return false
}

func (api *API) GetStakerSnapshot() *common.StakerSnapshot {
	if common.StakerSnapShot == nil {
		signers := make([]common.Staker, 0)
		common.StakerSnapShot = newStakerSnapshot(signers)
	}
	return common.StakerSnapShot
}

func (api *API) AddStakerToSnapshot(address common.Address) (bool, error) {
	if common.StakerSnapShot == nil{
		api.GetStakerSnapshot()
	}
	for _, st := range common.StakerSnapShot.Stakers {
		if st.Address == address {
			return false, nil
		}
	}
	var staker common.Staker
	staker.Address = address
	staker.LastSeen = time.Now()
	common.StakerSnapShot.Stakers = append(common.StakerSnapShot.Stakers, staker)
	return true, nil
}
