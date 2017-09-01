package xenio

import (
	//"encoding/json"
	"time"

	"github.com/xenioplatform/go-xenio/common"
	//"github.com/xenioplatform/go-xenio/ethdb"
	//"github.com/xenioplatform/go-xenio/log"
	"strconv"
)

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

// cast adds new stakers to the list.
func StakerCast(stakers map[common.Address]common.Staker) bool {
	// Cast the stakers into an existing or new list
	newStakerList := make(map[common.Address]common.Staker, len(stakers))

	for key, value := range common.StakerSnapShot.Stakers {
		newStakerList[key] = value
	}
	for key, value := range stakers {
		newStakerList[key] = value
	}
	common.StakerSnapShot.Stakers = newStakerList
	return true
}
*/

// cast adds new stakers to the list.
func StakerCast(stakers []common.StakerTransmit) {
	now := time.Now().UTC()
	for _, value := range stakers {
		var newStaker common.Staker
		_time, err := strconv.ParseInt(value.LastSeen, 10, 64)
		if err == nil {
			newStaker.LastSeen = time.Unix(_time, 0).UTC()
			delta := now.Sub(newStaker.LastSeen)
			// only add non expired stakers to list
			if delta.Seconds() <= common.StakerTTL {
				stakerExists := StakerExists(value.Address)
				if !stakerExists || newStaker.LastSeen.After(common.StakerSnapShot.Stakers[value.Address].LastSeen) {
					common.StakerSnapShot.Stakers[value.Address] = newStaker
				}
			}
		}
	}
}

func NewStakerSnapshot() *common.StakerSnapshot {
	snap := &common.StakerSnapshot{
		Created: time.Now().UTC(),
		Stakers: make(map[common.Address]common.Staker),
	}
	return snap
}

func StakerExists(address common.Address) bool {
	_, exists := common.StakerSnapShot.Stakers[address]
	return exists
}

func StakerExpired(address common.Address) bool {
	stakerData, exists := common.StakerSnapShot.Stakers[address]
	if exists == true {
		delta := time.Now().Sub(stakerData.LastSeen)
		if delta.Seconds() >= common.StakerTTL {
			return true
		}
	}
	return false
}

func DeleteStaker(address common.Address) {
	delete(common.StakerSnapShot.Stakers, address)
}

func DeleteExpiredStaker(address common.Address) {
	expired := StakerExpired(address)
	if expired {
		delete(common.StakerSnapShot.Stakers, address)
	}
}

func DeleteAllExpiredStakers() {
	for address := range common.StakerSnapShot.Stakers {
		if StakerExpired(address) == true {
			delete(common.StakerSnapShot.Stakers, address)
		}
	}
}

func (api *API) GetActiveStakerList() []common.Address {
	var stakerList []common.Address
	if common.StakerSnapShot != nil {
		for address := range common.StakerSnapShot.Stakers {
			if !StakerExpired(address) {
				stakerList = append(stakerList, address)
			}
		}
	}
	return stakerList
}

func (api *API) GetStakerSnapshot() *common.StakerSnapshot {
	if common.StakerSnapShot == nil {
		common.StakerSnapShot = NewStakerSnapshot()
	}
	//DeleteAllExpiredStakers()
	return common.StakerSnapShot
}

func (api *API) AddStakerToSnapshot(address common.Address) {
	if common.StakerSnapShot == nil {
		api.GetStakerSnapshot()
	}
	var staker common.Staker
	staker.LastSeen = time.Now().UTC()
	common.StakerSnapShot.Stakers[address] = staker
}
