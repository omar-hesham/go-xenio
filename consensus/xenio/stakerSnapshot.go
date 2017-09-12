// Copyright 2017 The go-xenio Authors
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
			newStaker.FirstSeen = newStaker.LastSeen
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
	staker.FirstSeen = time.Now().UTC()
	staker.LastSeen = staker.FirstSeen
	common.StakerSnapShot.Stakers[address] = staker
}
