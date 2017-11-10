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
	"math/big"
	"strings"

	"github.com/xenioplatform/go-xenio/common"
	"github.com/xenioplatform/go-xenio/contracts/xnogames"
	"github.com/xenioplatform/go-xenio/core/types"
	"github.com/xenioplatform/go-xenio/log"
)

type Game struct {
	Address   string `json:"address"`
	Name      string `json:"title"`
	Publisher string `json:"publisher"`
	Developer string `json:"developer"`
	Country   string `json:"country"`
	State     string `json:"state"`
	LogoImg   string `json:"logoimg"`
}

/* Free data retrieval calls */

func (api *API) GetAllImages() ([][]byte, error) {
	contract, err := getGamesContract()
	callOpts := getFreeTxTransactor()
	result, err := contract.GetAllImages(callOpts)
	return result, err
}

func (api *API) GetAllGames() ([]Game, error) {
	contract, err := getGamesContract()
	callOpts := getFreeTxTransactor()
	gamesAddresses, err := contract.GetGamesAddresses(callOpts)
	var games []Game
	if err.Error() == "abi: unmarshalling empty output" {
		return nil, nil
	}
	for i := 0; i < len(gamesAddresses); i++ {
		name, publisher, developer, country, state, logoImg, err := contract.GetGame(callOpts, gamesAddresses[i])
		if err != nil {
			return nil, err
		}
		game := Game{gamesAddresses[i].String(), name, publisher, developer, strings.Replace(string(country[:]), "\x00", "", -1), strings.Replace(string(state[:]), "\x00", "", -1), strings.Replace(string(logoImg[:]), "\x00", "", -1)}
		games = append(games, game)
	}
	return games, err
}

func (api *API) GetGame(gameAddress common.Address) (Game, error) {
	contract, err := getGamesContract()
	callOpts := getFreeTxTransactor()
	name, publisher, developer, country, state, logoImg, err := contract.GetGame(callOpts, gameAddress)
	game := Game{gameAddress.String(), name, publisher, developer, strings.Replace(string(country[:]), "\x00", "", -1), strings.Replace(string(state[:]), "\x00", "", -1), strings.Replace(string(logoImg[:]), "\x00", "", -1)}
	return game, err
}

func (api *API) GetGameImages(userAddress common.Address) ([32]byte, error) {
	contract, err := getGamesContract()
	callOpts := getFreeTxTransactor()
	result, err := contract.GetGameImages(callOpts, userAddress)
	return result, err
}

func (api *API) GetGamesAddresses() ([]common.Address, error) {
	contract, err := getGamesContract()
	callOpts := getFreeTxTransactor()
	result, err := contract.GetGamesAddresses(callOpts)
	if err.Error() == "abi: unmarshalling empty output" {
		return nil, nil
	}
	return result, err
}

//func (api *API) GetImage(SHA256notaryHash [32]byte) (string, *big.Int, error) {
//	contract, err := getGamesContract()
//	callOpts := bind.CallOpts{}
//	imageURL, timestamp, err := contract.GetImage(&callOpts, SHA256notaryHash)
//	return imageURL, timestamp, err
//}

/* Paid mutator transaction calls */

func (api *API) RegisterNewGame(name string, publisher string, developer string, country [32]byte, state [32]byte, gas *big.Int) (*types.Transaction, error) {
	if currentTransactor.contractAuth == nil || currentTransactor.authorizedTransactions == 0 {
		return nil, errTransactorNotSet
	}
	// Set Gas for paid transaction -- use default gas if not provided
	currentTransactor.contractAuth.GasPrice = gas

	contract, err := getGamesContract()
	result, err := contract.RegisterNewGame(currentTransactor.contractAuth, name, publisher, developer, country, state)
	if err == nil {
		// transaction was successful, deduct from authorized transactions
		resetContractTransactor()
	}
	return result, err
}

// Deploy and propose new users contract - under construction
func (api *API) DeployNewGamesContract(gas *big.Int) (common.Address, error) {
	if currentTransactor.contractAuth == nil || currentTransactor.authorizedTransactions == 0 {
		return *new(common.Address), errTransactorNotSet
	}
	// Set Gas for paid transaction -- use default gas if not provided
	currentTransactor.contractAuth.GasPrice = gas

	conn, err := getContractBackend()
	address, _, _, err := xnogames.DeployXNOGames(currentTransactor.contractAuth, conn)
	if err == nil {
		// propose newly created contract
		api.GamesContractVote(address, true)

		// transaction was successful, deduct from authorized transactions
		resetContractTransactor()
	}
	return address, err
}

///* XNOGames contract helper functions */

// Get deployed contract object
func getGamesContract() (*xnogames.XNOGames, error) {
	conn, err := getContractBackend()
	contract, err := xnogames.NewXNOGames(deployedGamesContract, conn)
	if err != nil {
		log.Error("Failed to instantiate a Games contract: " + err.Error())
	}
	return contract, err
}

func (api *API) GetXNOGamesABI() string {
	api.xenio.lock.Lock()
	defer api.xenio.lock.Unlock()
	return xnogames.XNOGamesABI
}
