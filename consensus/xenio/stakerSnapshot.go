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

// cast adds new stakers to the list.
func StakerCast(stakers []common.StakerTransmit) {
	if common.StakerSnapShot == nil {
		common.StakerSnapShot = NewStakerSnapshot()
	}
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
				if newStaker.FirstSeen.UTC().Unix() == -62135596800 { // default Unix Time
					newStaker.FirstSeen = newStaker.LastSeen.UTC()
				}

				existingStaker, exists := common.StakerSnapShot.Stakers.Load(value.Address)

				if exists {
					// Keep existing firstSeen value if we've seen this staker before time in list, unless existing firstSeen is default Unix Time
					if newStaker.FirstSeen.UTC().After(existingStaker.(common.Staker).FirstSeen.UTC()) && existingStaker.(common.Staker).FirstSeen.UTC().Unix() != -62135596800 {
						newStaker.FirstSeen = existingStaker.(common.Staker).FirstSeen.UTC()
					}
					// Keep existing lastSeen value if we've seen this staker after time in list
					if newStaker.LastSeen.UTC().Before(existingStaker.(common.Staker).LastSeen.UTC()) {
						newStaker.LastSeen = existingStaker.(common.Staker).LastSeen.UTC()
					}
				}
				// Update staker list
				common.StakerSnapShot.Stakers.Store(value.Address, newStaker)
			}
		}
	}
}

func NewStakerSnapshot() *common.StakerSnapshot {
	snap := &common.StakerSnapshot{
		Created: time.Unix(time.Now().UTC().Unix(), 0).UTC(),
	}
	return snap
}

func StakerExists(address common.Address) bool {
	if common.StakerSnapShot != nil {
		_, exists := common.StakerSnapShot.Stakers.Load(address)
		return exists
	}
	return false
}

func StakerExpired(address common.Address) bool {
	if common.StakerSnapShot != nil {
		stakerData, exists := common.StakerSnapShot.Stakers.Load(address)
		if exists == true {
			delta := time.Now().UTC().Sub(stakerData.(common.Staker).LastSeen.UTC())
			if delta.Seconds() >= common.StakerTTL {
				return true
			}
		}
	}
	return false
}

func DeleteStaker(address common.Address) {
	if common.StakerSnapShot != nil {
		common.StakerSnapShot.Stakers.Delete(address)
	}
}

func DeleteExpiredStaker(address common.Address) {
	if common.StakerSnapShot != nil {
		expired := StakerExpired(address)
		if expired {
			common.StakerSnapShot.Stakers.Delete(address)
		}
	}
}

func DeleteAllExpiredStakers() {
	if common.StakerSnapShot != nil {
		common.StakerSnapShot.Stakers.Range(
			func(address, staker interface{}) bool {
				if StakerExpired(address.(common.Address)) == true {
					common.StakerSnapShot.Stakers.Delete(address)
				}
				return true
			})
	}
}

func HasCoins(address common.Address,value int64, state *state.StateDB) bool {
	coins := state.GetBalance(address)
	if coins.Cmp(big.NewInt(value)) == 1 {
		return true
	} else {
		return false
	}
}

func (api *API) GetActiveStakerList() []common.Address {
	var stakerList []common.Address
	if common.StakerSnapShot != nil {
		common.StakerSnapShot.Stakers.Range(
			func(address, staker interface{}) bool {
				if !StakerExpired(address.(common.Address)) {
					stakerList = append(stakerList, address.(common.Address))
				}
				return true
			})
	}
	return stakerList
}

func (api *API) GetStakerSnapshot() *common.StakerSnapshot {
	if common.StakerSnapShot == nil {
		common.StakerSnapShot = NewStakerSnapshot()
	} else {
		DeleteAllExpiredStakers()
	}
	return common.StakerSnapShot
}

func (api *API) GetStakerSnapshotJS() interface{} {
	var (
		snapshot   = api.GetStakerSnapshot()
		stakerList = make(map[common.Address]common.Staker)
	)

	snapshot.Stakers.Range(
		func(address, staker interface{}) bool {
			stakerList[address.(common.Address)] = staker.(common.Staker)
			return true
		})

	return map[string]interface{}{
		"created": common.StakerSnapShot.Created.UTC(),
		"stakers": stakerList,
	}
}

func (api *API) AddStakerToSnapshot(address common.Address) {
	if common.StakerSnapShot == nil {
		api.GetStakerSnapshot()
	}
	var staker common.Staker
	staker.LastSeen = time.Unix(time.Now().UTC().Unix(), 0).UTC()

	existingStaker, exists := common.StakerSnapShot.Stakers.Load(address)

	if exists {
		staker.FirstSeen = existingStaker.(common.Staker).FirstSeen.UTC()
	}

	if !exists || existingStaker.(common.Staker).FirstSeen.UTC().Unix() == -62135596800 { // default Unix Time
		staker.FirstSeen = staker.LastSeen.UTC()
	} else {
		staker.FirstSeen = existingStaker.(common.Staker).FirstSeen.UTC()
	}
	common.StakerSnapShot.Stakers.Store(address, staker)
}
