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
	"github.com/xenioplatform/go-xenio/accounts/abi/bind"
	"github.com/xenioplatform/go-xenio/common"
	"github.com/xenioplatform/go-xenio/contracts/xnousers"
	"github.com/xenioplatform/go-xenio/ethclient"
	"github.com/xenioplatform/go-xenio/log"
	//"github.com/xenioplatform/go-xenio/core/types"
)


/* Free data retrieval calls */

func (api *API) GetAllUsersAddresses() ([]common.Address, error) {
	contract, err := getUsersContract()
	callOpts := bind.CallOpts{}
	result, err := contract.GetAllUsersAddresses(&callOpts)
	return result, err
}

//func (api *API) GetGamers() ([][32]byte, [][32]byte, error) {
//	contract, err := getUsersContract()
//	callOpts := bind.CallOpts{}
//	names, games, err := contract.GetGamers(&callOpts)
//	return names, games, err
//}

func (api *API) GetGamersAddresses() ([]common.Address, error) {
	contract, err := getUsersContract()
	callOpts := bind.CallOpts{}
	result, err := contract.GetGamersAddresses(&callOpts)
	return result, err
}

func (api *API) GetGamersNumber() (uint64, error) {
	contract, err := getUsersContract()
	callOpts := bind.CallOpts{}
	result, err := contract.GetGamersNumber(&callOpts)
	return result.Uint64(), err
}

//func (api *API) GetServers() ([][32]byte, [][32]byte, error) {
//	contract, err := getUsersContract()
//	callOpts := bind.CallOpts{}
//	names, games, err := contract.GetServers(&callOpts)
//	return names, games, err
//}

func (api *API) GetServersAddresses() ([]common.Address, error) {
	contract, err := getUsersContract()
	callOpts := bind.CallOpts{}
	result, err := contract.GetServersAddresses(&callOpts)
	return result, err
}

//func (api *API) GetUsers() ([][32]byte, [][32]byte, error) {
//	contract, err := getUsersContract()
//	callOpts := bind.CallOpts{}
//	names, games, err := contract.GetUsers(&callOpts)
//	return names, games, err
//}

func (api *API) IsServer(userAddress common.Address) (bool, error) {
	contract, err := getUsersContract()
	callOpts := bind.CallOpts{}
	result, err := contract.IsServer(&callOpts, userAddress)
	return result, err
}

/* Paid mutator transaction calls */

//// not working - under construction
//func (api *API) RegisterNewUser(name string, isServer bool, game string) (*types.Transaction, error) {
//	contract, err := getUsersContract()
//	transactOps := bind.TransactOpts{From: common.HexToAddress("0xd956e4b845b574c3519509b1c2cd3090b7eb97d4"), GasLimit: big.NewInt(300000)}
//	result, _ := contract.RegisterNewUser(&transactOps, name, isServer, game)
//	return result, err
//}


/* Contract helper functions */

//// not working - under construction
//// Deploy and propose new users contract
//func (api *API) DeployNewUsersContract() {
//	// Create an IPC based RPC connection to a remote node
//	conn, err := ethclient.Dial(currentIPCEndpoint)
//	if err != nil {
//		log.Error("Failed to connect to the Xenio client: " + err.Error())
//	}
//
//bind.NewKeyedTransactor()
//bind.NewTransactor()
//
//	transactOps := bind.TransactOpts{From: common.HexToAddress("0xd956e4b845b574c3519509b1c2cd3090b7eb97d4"), GasLimit: big.NewInt(300000)}
//	//address, transaction, contractObject, err := xnousers.DeployXNOUsers(&transactOps, conn)
//	_, _, _, _ = xnousers.DeployXNOUsers(&transactOps, conn)
//}

// Get deployed contract object
func getUsersContract() (*xnousers.XNOUsers, error) {
	// Create an IPC based RPC connection to a remote node
	conn, err := ethclient.Dial(currentIPCEndpoint)
	if err != nil {
		log.Error("Failed to connect to the Xenio client: " + err.Error())
	}
	// Instantiate the contract and display its name
	contract, err := xnousers.NewXNOUsers(deployedUsersContract, conn)
	if err != nil {
		log.Error("Failed to instantiate a Users contract: " + err.Error())
	}
	return contract, err
}

func (api *API) GetXNOUsersABI() string{
	api.xenio.lock.Lock()
	defer api.xenio.lock.Unlock()
	return xnousers.XNOUsersABI
}