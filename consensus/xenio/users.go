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
	"github.com/xenioplatform/go-xenio/accounts/abi/bind"
	"github.com/xenioplatform/go-xenio/common"
	//"github.com/xenioplatform/go-xenio/contracts/xnogames"
	"github.com/xenioplatform/go-xenio/contracts/xnousers"
	"github.com/xenioplatform/go-xenio/ethclient"
	"github.com/xenioplatform/go-xenio/log"
	"github.com/xenioplatform/go-xenio/core/types"
	"math/big"
)

func (api *API) GetUserAddresses(contractAddress common.Address) []common.Address {

	h := bind.CallOpts{}
	// Create an IPC based RPC connection to a remote node
	conn, err := ethclient.Dial(currentIPCEndpoint)
	if err != nil {
		log.Error("Failed to connect to the Ethereum client: " + err.Error())
	}
	// Instantiate the contract and display its name
	token, err := xnousers.NewXNOUsers(contractAddress, conn)
	if err != nil {
		log.Error("Failed to instantiate a Token contract: " + err.Error())
	}
	f1, _ := token.GetGamersAddresses(&h)
	return f1
}

func (api *API) GetServerAddresses(contractAddress common.Address) []common.Address {

	h := bind.CallOpts{}
	// Create an IPC based RPC connection to a remote node
	conn, err := ethclient.Dial(currentIPCEndpoint)
	if err != nil {
		log.Error("Failed to connect to the Ethereum client: " + err.Error())
	}
	// Instantiate the contract and display its name
	token, err := xnousers.NewXNOUsers(contractAddress, conn)
	if err != nil {
		log.Error("Failed to instantiate a Token contract: " + err.Error())
	}
	f2, _ := token.GetServersAddresses(&h)
	return f2
}

// not working
func (api *API) RegisterNewUser(contractAddress common.Address, name string, isServer bool, game string) *types.Transaction {

	h := bind.TransactOpts{From: common.HexToAddress("0xd956e4b845b574c3519509b1c2cd3090b7eb97d4"), GasLimit: big.NewInt(300000)}
	// Create an IPC based RPC connection to a remote node
	conn, err := ethclient.Dial(currentIPCEndpoint)
	if err != nil {
		log.Error("Failed to connect to the Ethereum client: " + err.Error())
	}
	// Instantiate the contract and display its name
	token, err := xnousers.NewXNOUsers(contractAddress, conn)
	if err != nil {
		log.Error("Failed to instantiate a Token contract: " + err.Error())
	}
	result, _ := token.RegisterNewUser(&h, name, isServer, game)
	return result
}