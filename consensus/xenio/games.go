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
// This file provides functions to interact with the XNOGames contract

package xenio

import (
	//"github.com/xenioplatform/go-xenio/accounts/abi/bind"
	//"github.com/xenioplatform/go-xenio/common"
	"github.com/xenioplatform/go-xenio/contracts/xnogames"
	"github.com/xenioplatform/go-xenio/ethclient"
	"github.com/xenioplatform/go-xenio/log"
)


/* Free data retrieval calls */

//func (api *API) GetAllImages() ([][32]byte, error) {
//	contract, err := getGamesContract()
//	callOpts := bind.CallOpts{}
//	result, err := contract.GetAllImages(&callOpts)
//	return result, err
//}
//
//func (api *API) GetGame(gameAddress common.Address) (string, string, string, [32]byte, [32]byte, [32]byte, error) {
//	contract, err := getGamesContract()
//	callOpts := bind.CallOpts{}
//	name, publisher, developer, country, state, logoImg, err := contract.GetGame(&callOpts, gameAddress)
//	return name, publisher, developer, country, state, logoImg, err
//}
//
//func (api *API) GetGameImages(userAddress common.Address) ([32]byte, error) {
//	contract, err := getGamesContract()
//	callOpts := bind.CallOpts{}
//	result, err := contract.GetGameImages(&callOpts, userAddress)
//	return result, err
//}
//
//func (api *API) GetGames() ([][32]byte, [][32]byte, [][32]byte, [][32]byte, [][32]byte, error) {
//	contract, err := getGamesContract()
//	callOpts := bind.CallOpts{}
//	names, publishers, developers, countries, states, err := contract.GetGames(&callOpts)
//	return names, publishers, developers, countries, states, err
//}
//
//func (api *API) GetGamesAddresses() ([]common.Address, error) {
//	contract, err := getGamesContract()
//	callOpts := bind.CallOpts{}
//	result, err := contract.GetGamesAddresses(&callOpts)
//	return result, err
//}
//
//func (api *API) GetImage(SHA256notaryHash [32]byte) (string, *big.Int, error) {
//	contract, err := getGamesContract()
//	callOpts := bind.CallOpts{}
//	imageURL, timestamp, err := contract.GetImage(&callOpts, SHA256notaryHash)
//	return imageURL, timestamp, err
//}


///* Contract helper functions */

// Get deployed contract object
func getGamesContract() (*xnogames.XNOGames, error) {
	// Create an IPC based RPC connection to a remote node
	conn, err := ethclient.Dial(currentIPCEndpoint)
	if err != nil {
		log.Error("Failed to connect to the Xenio client: " + err.Error())
	}
	// Instantiate the contract and display its name
	contract, err := xnogames.NewXNOGames(deployedGamesContract, conn)
	if err != nil {
		log.Error("Failed to instantiate a Users contract: " + err.Error())
	}
	return contract, err
}

func (api *API) GetXNOGamesABI() string{
	api.xenio.lock.Lock()
	defer api.xenio.lock.Unlock()
	return xnogames.XNOGamesABI
}

//// Instantiate the contract and display its name
//token2, err := xnogames.NewXNOGames(snap.GamesContractAddress, conn)
//if err != nil {
//	log.Error("Failed to instantiate a Token contract: " + err.Error())
//}
//g1, g2, g3, g4, g5, _ := token2.GetGames(&h)
////if err != nil {
////	log.Error("Failed to retrieve token name: " + err.Error())
////}
//fmt.Println("All Games:", g1, g2, g3, g4, g5)
