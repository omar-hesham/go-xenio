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
//
// This file provides functions to interact with the XNOUsers contract

package xenio

import (
	"math/big"

	"github.com/xenioplatform/go-xenio/common"
	"github.com/xenioplatform/go-xenio/contracts/xnousers"
	"github.com/xenioplatform/go-xenio/core/types"
	"github.com/xenioplatform/go-xenio/log"
)

type User struct {
	Address  common.Address `json:"address"`
	Name     string         `json:"username"`
	IsServer bool           `json:"isserver"`
	Game     string         `json:"gamename"`
}

/* Free data retrieval calls */

func (api *API) GetAllUsers() ([]User, error) {
	contract, err := getUsersContract()
	callOpts := getFreeTxTransactor()
	userAddresses, err := contract.GetAllUsersAddresses(callOpts)
	var usersList []User
	for i := 0; i < len(userAddresses); i++ {
		userName, isServer, gameName, err := contract.GetUser(callOpts, userAddresses[i])
		if err != nil {
			return nil, err
		}
		user := User{userAddresses[i], userName, isServer, gameName}
		usersList = append(usersList, user)
	}
	return usersList, err
}

func (api *API) GetAllUsersAddresses() ([]common.Address, error) {
	contract, err := getUsersContract()
	callOpts := getFreeTxTransactor()
	result, err := contract.GetAllUsersAddresses(callOpts)
	return result, err
}

func (api *API) GetGamers() ([]User, error) {
	contract, err := getUsersContract()
	callOpts := getFreeTxTransactor()
	gamerAddresses, err := contract.GetGamersAddresses(callOpts)
	var gamersList []User
	for i := 0; i < len(gamerAddresses); i++ {
		userName, isServer, gameName, err := contract.GetUser(callOpts, gamerAddresses[i])
		if err != nil {
			return nil, err
		}
		gamer := User{gamerAddresses[i], userName, isServer, gameName}
		gamersList = append(gamersList, gamer)
	}
	return gamersList, err
}

func (api *API) GetGamersAddresses() ([]common.Address, error) {
	contract, err := getUsersContract()
	callOpts := getFreeTxTransactor()
	result, err := contract.GetGamersAddresses(callOpts)
	return result, err
}

func (api *API) GetGamersNumber() (uint64, error) {
	contract, err := getUsersContract()
	callOpts := getFreeTxTransactor()
	result, err := contract.GetGamersNumber(callOpts)
	return result.Uint64(), err
}

func (api *API) GetServers() ([]User, error) {
	contract, err := getUsersContract()
	callOpts := getFreeTxTransactor()
	serverAddresses, err := contract.GetServersAddresses(callOpts)
	var serversList []User
	for i := 0; i < len(serverAddresses); i++ {
		userName, isServer, gameName, err := contract.GetUser(callOpts, serverAddresses[i])
		if err != nil {
			return nil, err
		}
		server := User{serverAddresses[i], userName, isServer, gameName}
		serversList = append(serversList, server)
	}
	return serversList, err
}

func (api *API) GetServersAddresses() ([]common.Address, error) {
	contract, err := getUsersContract()
	callOpts := getFreeTxTransactor()
	result, err := contract.GetServersAddresses(callOpts)
	return result, err
}

func (api *API) GetServersNumber() (uint64, error) {
	contract, err := getUsersContract()
	callOpts := getFreeTxTransactor()
	result, err := contract.GetServersNumber(callOpts)
	return result.Uint64(), err
}

func (api *API) GetUser(userAddress common.Address) (User, error) {
	contract, err := getUsersContract()
	callOpts := getFreeTxTransactor()
	userName, isServer, gameName, err := contract.GetUser(callOpts, userAddress)
	user := User{userAddress, userName, isServer, gameName}
	return user, err
}

func (api *API) IsServer(userAddress common.Address) (bool, error) {
	contract, err := getUsersContract()
	callOpts := getFreeTxTransactor()
	result, err := contract.IsServer(callOpts, userAddress)
	return result, err
}

/* Paid mutator transaction calls */

func (api *API) RegisterNewUser(name string, isServer bool, game string, gas *big.Int) (*types.Transaction, error) {
	if currentTransactor.contractAuth == nil || currentTransactor.authorizedTransactions == 0 {
		return nil, errTransactorNotSet
	}
	// Set Gas for paid transaction -- use default gas if not provided
	currentTransactor.contractAuth.GasPrice = gas

	contract, err := getUsersContract()
	result, err := contract.RegisterNewUser(currentTransactor.contractAuth, name, isServer, game)
	if err == nil {
		// transaction was successful, deduct from authorized transactions
		evaluateContractTransactorAuth()
	}
	return result, err
}

// Deploy and propose new users contract - under construction
func (api *API) DeployNewUsersContract(gas *big.Int) (common.Address, error) {
	if currentTransactor.contractAuth == nil || currentTransactor.authorizedTransactions == 0 {
		return *new(common.Address), errTransactorNotSet
	}
	// Set Gas for paid transaction -- use default gas if not provided
	currentTransactor.contractAuth.GasPrice = gas

	conn, err := getContractBackend()
	address, _, _, err := xnousers.DeployXNOUsers(currentTransactor.contractAuth, conn)
	if err == nil {
		// propose newly created contract
		api.UsersContractVote(address, true)

		// transaction was successful, deduct from authorized transactions
		evaluateContractTransactorAuth()
	}
	return address, err
}

/* XNOUsers helper functions */

// Get deployed contract object
func getUsersContract() (*xnousers.XNOUsers, error) {
	conn, err := getContractBackend()
	contract, err := xnousers.NewXNOUsers(deployedUsersContract, conn)
	if err != nil {
		log.Error("Failed to instantiate a Users contract: " + err.Error())
	}
	return contract, err
}

func (api *API) getXNOUsersABI() string {
	api.xenio.lock.Lock()
	defer api.xenio.lock.Unlock()
	return xnousers.XNOUsersABI
}
