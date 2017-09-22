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
	"math/big"
	"strconv"
	"time"

	"github.com/xenioplatform/go-xenio/common"
	"github.com/xenioplatform/go-xenio/core/state"
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

*/

// cast adds new stakers to the list.
func StakerCast(stakers []common.StakerTransmit) {
	now := time.Now().UTC()
	for _, value := range stakers {
		var newStaker common.Staker
		lastSeen, err := strconv.ParseInt(value.LastSeen, 10, 64)
		firstSeen, err := strconv.ParseInt(value.FirstSeen, 10, 64)
		if err == nil {
			delta := now.Sub(time.Unix(lastSeen, 0).UTC())
			// only add non expired stakers to list
			if delta.Seconds() <= common.StakerTTL {
				newStaker.LastSeen = time.Unix(lastSeen, 0).UTC()
				newStaker.FirstSeen = time.Unix(firstSeen, 0).UTC()

				// Handle case where clients haven't recorded firstSeen
				if newStaker.FirstSeen.Unix() == 0 {
					newStaker.FirstSeen = newStaker.LastSeen
				}

				if StakerExists(value.Address) {
					// Keep existing firstSeen value if we've seen this staker before time in list, unless existing firstSeen is default Unix Time
					if newStaker.FirstSeen.After(common.StakerSnapShot.Stakers[value.Address].FirstSeen) && common.StakerSnapShot.Stakers[value.Address].FirstSeen.Unix() != 0 {
						newStaker.FirstSeen = common.StakerSnapShot.Stakers[value.Address].FirstSeen
					}
					// Keep existing lastSeen value if we've seen this staker after time in list
					if newStaker.LastSeen.Before(common.StakerSnapShot.Stakers[value.Address].LastSeen) {
						newStaker.LastSeen = common.StakerSnapShot.Stakers[value.Address].LastSeen
					}
				}
				// Update staker list
				common.StakerSnapShot.Stakers[value.Address] = newStaker
			}
		}
	}
}

func NewStakerSnapshot() *common.StakerSnapshot {
	snap := &common.StakerSnapshot{
		Created: time.Unix(time.Now().UTC().Unix(), 0).UTC(),
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

func HasCoins(address common.Address, state *state.StateDB) bool {
	coins := state.GetBalance(address)
	if coins.Cmp(big.NewInt(0)) == 1 {
		return true
	} else {
		return false
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
	staker.LastSeen = time.Unix(time.Now().UTC().Unix(), 0).UTC()
	if !StakerExists(address) || common.StakerSnapShot.Stakers[address].FirstSeen.Unix() == 0 {
		staker.FirstSeen = staker.LastSeen
	}
	common.StakerSnapShot.Stakers[address] = staker
}
