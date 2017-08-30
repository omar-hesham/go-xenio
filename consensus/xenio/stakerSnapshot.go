package xenio

import (
	//"encoding/json"
	"time"

	"github.com/xenioplatform/go-xenio/common"
	//"github.com/xenioplatform/go-xenio/ethdb"
	//"github.com/xenioplatform/go-xenio/log"
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
	// Cast the stakers into an existing or new list
	for _, value := range stakers {
		var newStaker common.Staker
		newStaker.LastSeen = value.LastSeen
		common.StakerSnapShot.Stakers[value.Address] = newStaker
	}
}

func NewStakerSnapshot() *common.StakerSnapshot {
	snap := &common.StakerSnapshot{
		Created: time.Now(),
		Stakers: make(map[common.Address]common.Staker),
	}
	return snap
}

func StakerExists(address common.Address) bool {
	_, exists := common.StakerSnapShot.Stakers[address]
	return exists
}

func DeleteStaker(address common.Address) {
	delete(common.StakerSnapShot.Stakers, address)
}

func (api *API) GetStakerSnapshot() *common.StakerSnapshot {
	if common.StakerSnapShot == nil {
		common.StakerSnapShot = NewStakerSnapshot()
	}
	return common.StakerSnapShot
}

func (api *API) AddStakerToSnapshot(address common.Address) {
	if common.StakerSnapShot == nil {
		api.GetStakerSnapshot()
	}
	var staker common.Staker
	staker.LastSeen = time.Now()
	common.StakerSnapShot.Stakers[address] = staker
}
