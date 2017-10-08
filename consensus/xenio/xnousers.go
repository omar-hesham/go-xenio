// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package xenio

import (
	"math/big"
	"strings"
	"github.com/xenioplatform/go-xenio/common"
	"github.com/xenioplatform/go-xenio/core/types"
	"github.com/xenioplatform/go-xenio/accounts/abi"
	"github.com/xenioplatform/go-xenio/accounts/abi/bind"
)

// XNOUsersABI is the input ABI used to generate the binding from.
const XNOUsersABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"imageURL\",\"type\":\"string\"},{\"name\":\"SHA256notaryHash\",\"type\":\"bytes32\"}],\"name\":\"addLogo\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"userAddress\",\"type\":\"address\"}],\"name\":\"getGameImages\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"SHA256notaryHash\",\"type\":\"bytes32\"}],\"name\":\"getImage\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAllImages\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"string\"},{\"name\":\"isServer\",\"type\":\"bool\"},{\"name\":\"country\",\"type\":\"bytes32\"},{\"name\":\"state\",\"type\":\"bytes32\"}],\"name\":\"registerNewUser\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"gameAddress\",\"type\":\"address\"}],\"name\":\"getGame\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"gameAddress\",\"type\":\"address\"}],\"name\":\"removeUser\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"imgHash\",\"type\":\"bytes32\"}],\"name\":\"removeImage\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAllGames\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"}]"

// XNOUsersBin is the compiled bytecode used for deploying new contracts.
const XNOUsersBin = `0x606060405260008054600160a060020a033316600160a060020a03199091161790556109f0806100306000396000f300606060405236156100965763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416632594677581146100a357806334c166941461010a5780636ced1ae91461013b5780637b7284fc146101d05780638b907bbc146102365780638d1a4b2814610295578063985751881461034c5780639f6c93581461036b578063db1c45f914610381575b34156100a157600080fd5b005b34156100ae57600080fd5b6100f660046024813581810190830135806020601f82018190048102016040519081016040528181529291906020840183838082843750949650509335935061039492505050565b604051901515815260200160405180910390f35b341561011557600080fd5b610129600160a060020a036004351661048d565b60405190815260200160405180910390f35b341561014657600080fd5b6101516004356104ab565b6040518080602001838152602001828103825284818151815260200191508051906020019080838360005b8381101561019457808201518382015260200161017c565b50505050905090810190601f1680156101c15780820380516001836020036101000a031916815260200191505b50935050505060405180910390f35b34156101db57600080fd5b6101e3610578565b60405160208082528190810183818151815260200191508051906020019060200280838360005b8381101561022257808201518382015260200161020a565b505050509050019250505060405180910390f35b341561024157600080fd5b6100f660046024813581810190830135806020601f820181900481020160405190810160405281815292919060208401838380828437509496505050508235151592602081013592506040013590506105d8565b34156102a057600080fd5b6102b4600160a060020a0360043516610691565b604051841515602082015260408101849052606081018390526080810182905260a08082528190810187818151815260200191508051906020019080838360005b8381101561030d5780820151838201526020016102f5565b50505050905090810190601f16801561033a5780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390f35b341561035757600080fd5b6100f6600160a060020a036004351661078c565b341561037657600080fd5b6100f66004356107ff565b341561038c57600080fd5b6101e3610843565b33600160a060020a038116600090815260016020819052604082205491929160029181161561010002600019011604156103d15760009150610486565b8351156103e15760009150610486565b60008381526003602052604090205460026101006001831615026000190190911604156104115760009150610486565b600480546001810161042383826108aa565b50600091825260208083209190910185905584825260039052604090208480516104519291602001906108d3565b50600083815260036020908152604080832042600191820155600160a060020a03851684529182905290912060040184905591505b5092915050565b600160a060020a031660009081526001602052604090206004015490565b6104b3610951565b600082815260036020908152604080832060018082015482549294909385936002610100948216159490940260001901169290920491601f83018290048202909101905190810160405280929190818152602001828054600181600116156101000203166002900480156105685780601f1061053d57610100808354040283529160200191610568565b820191906000526020600020905b81548152906001019060200180831161054b57829003601f168201915b5050505050915091509150915091565b610580610951565b60048054806020026020016040519081016040528092919081815260200182805480156105cd57602002820191906000526020600020905b815481526001909101906020018083116105b8575b505050505090505b90565b33600160a060020a03811660009081526001602052604081209091908680516106059291602001906108d3565b50600160a060020a0381166000908152600160208190526040909120808201805460ff19168815151790556002808201879055600390910185905580549091810161065083826108aa565b5060009182526020909120018054600160a060020a03831673ffffffffffffffffffffffffffffffffffffffff199091161790556001915050949350505050565b610699610951565b600160a060020a0382166000908152600160208181526040808420808401546002808301546003840154600485015485548a998a998a9960ff9098169795969495939489946000199385161561010002939093019093169290920491601f8301819004810201905190810160405280929190818152602001828054600181600116156101000203166002900480156107725780601f1061074757610100808354040283529160200191610772565b820191906000526020600020905b81548152906001019060200180831161075557829003601f168201915b505050505094509450945094509450945091939590929450565b6000805433600160a060020a039081169116146107a857600080fd5b600160a060020a0382166000908152600160205260408120906107cb8282610963565b506001818101805460ff19169055600060028301819055600383018190556004830181905560059092019190915592915050565b6000805433600160a060020a0390811691161461081b57600080fd5b6000828152600360205260408120906108348282610963565b50600060019182015592915050565b61084b610951565b60028054806020026020016040519081016040528092919081815260200182805480156105cd57602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610883575050505050905090565b8154818355818115116108ce576000838152602090206108ce9181019083016109aa565b505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061091457805160ff1916838001178555610941565b82800160010185558215610941579182015b82811115610941578251825591602001919060010190610926565b5061094d9291506109aa565b5090565b60206040519081016040526000815290565b50805460018160011615610100020316600290046000825580601f1061098957506109a7565b601f0160209004906000526020600020908101906109a791906109aa565b50565b6105d591905b8082111561094d57600081556001016109b05600a165627a7a72305820cb7dacc9fc9ef5ba0509de1bd6d796f7cbb6f034ce5780b8ef4ad5eb9e214c090029`

// DeployXNOUsers deploys a new Ethereum contract, binding an instance of XNOUsers to it.
func DeployXNOUsers(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *XNOUsers, error) {
	parsed, err := abi.JSON(strings.NewReader(XNOUsersABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(XNOUsersBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &XNOUsers{XNOUsersCaller: XNOUsersCaller{contract: contract}, XNOUsersTransactor: XNOUsersTransactor{contract: contract}}, nil
}

// XNOUsers is an auto generated Go binding around an Ethereum contract.
type XNOUsers struct {
	XNOUsersCaller     // Read-only binding to the contract
	XNOUsersTransactor // Write-only binding to the contract
}

// XNOUsersCaller is an auto generated read-only Go binding around an Ethereum contract.
type XNOUsersCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// XNOUsersTransactor is an auto generated write-only Go binding around an Ethereum contract.
type XNOUsersTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// XNOUsersSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type XNOUsersSession struct {
	Contract     *XNOUsers         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// XNOUsersCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type XNOUsersCallerSession struct {
	Contract *XNOUsersCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// XNOUsersTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type XNOUsersTransactorSession struct {
	Contract     *XNOUsersTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// XNOUsersRaw is an auto generated low-level Go binding around an Ethereum contract.
type XNOUsersRaw struct {
	Contract *XNOUsers // Generic contract binding to access the raw methods on
}

// XNOUsersCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type XNOUsersCallerRaw struct {
	Contract *XNOUsersCaller // Generic read-only contract binding to access the raw methods on
}

// XNOUsersTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type XNOUsersTransactorRaw struct {
	Contract *XNOUsersTransactor // Generic write-only contract binding to access the raw methods on
}

// NewXNOUsers creates a new instance of XNOUsers, bound to a specific deployed contract.
func NewXNOUsers(address common.Address, backend bind.ContractBackend) (*XNOUsers, error) {
	contract, err := bindXNOUsers(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &XNOUsers{XNOUsersCaller: XNOUsersCaller{contract: contract}, XNOUsersTransactor: XNOUsersTransactor{contract: contract}}, nil
}

// NewXNOUsersCaller creates a new read-only instance of XNOUsers, bound to a specific deployed contract.
func NewXNOUsersCaller(address common.Address, caller bind.ContractCaller) (*XNOUsersCaller, error) {
	contract, err := bindXNOUsers(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &XNOUsersCaller{contract: contract}, nil
}

// NewXNOUsersTransactor creates a new write-only instance of XNOUsers, bound to a specific deployed contract.
func NewXNOUsersTransactor(address common.Address, transactor bind.ContractTransactor) (*XNOUsersTransactor, error) {
	contract, err := bindXNOUsers(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &XNOUsersTransactor{contract: contract}, nil
}

// bindXNOUsers binds a generic wrapper to an already deployed contract.
func bindXNOUsers(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(XNOUsersABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_XNOUsers *XNOUsersRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _XNOUsers.Contract.XNOUsersCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_XNOUsers *XNOUsersRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _XNOUsers.Contract.XNOUsersTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_XNOUsers *XNOUsersRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _XNOUsers.Contract.XNOUsersTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_XNOUsers *XNOUsersCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _XNOUsers.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_XNOUsers *XNOUsersTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _XNOUsers.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_XNOUsers *XNOUsersTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _XNOUsers.Contract.contract.Transact(opts, method, params...)
}

// GetAllGames is a free data retrieval call binding the contract method 0xdb1c45f9.
//
// Solidity: function getAllGames() constant returns(address[])
func (_XNOUsers *XNOUsersCaller) GetAllGames(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _XNOUsers.contract.Call(opts, out, "getAllGames")
	return *ret0, err
}

// GetAllGames is a free data retrieval call binding the contract method 0xdb1c45f9.
//
// Solidity: function getAllGames() constant returns(address[])
func (_XNOUsers *XNOUsersSession) GetAllGames() ([]common.Address, error) {
	return _XNOUsers.Contract.GetAllGames(&_XNOUsers.CallOpts)
}

// GetAllGames is a free data retrieval call binding the contract method 0xdb1c45f9.
//
// Solidity: function getAllGames() constant returns(address[])
func (_XNOUsers *XNOUsersCallerSession) GetAllGames() ([]common.Address, error) {
	return _XNOUsers.Contract.GetAllGames(&_XNOUsers.CallOpts)
}

// GetAllImages is a free data retrieval call binding the contract method 0x7b7284fc.
//
// Solidity: function getAllImages() constant returns(bytes32[])
func (_XNOUsers *XNOUsersCaller) GetAllImages(opts *bind.CallOpts) ([][32]byte, error) {
	var (
		ret0 = new([][32]byte)
	)
	out := ret0
	err := _XNOUsers.contract.Call(opts, out, "getAllImages")
	return *ret0, err
}

// GetAllImages is a free data retrieval call binding the contract method 0x7b7284fc.
//
// Solidity: function getAllImages() constant returns(bytes32[])
func (_XNOUsers *XNOUsersSession) GetAllImages() ([][32]byte, error) {
	return _XNOUsers.Contract.GetAllImages(&_XNOUsers.CallOpts)
}

// GetAllImages is a free data retrieval call binding the contract method 0x7b7284fc.
//
// Solidity: function getAllImages() constant returns(bytes32[])
func (_XNOUsers *XNOUsersCallerSession) GetAllImages() ([][32]byte, error) {
	return _XNOUsers.Contract.GetAllImages(&_XNOUsers.CallOpts)
}

// GetGame is a free data retrieval call binding the contract method 0x8d1a4b28.
//
// Solidity: function getGame(gameAddress address) constant returns(string, bool, bytes32, bytes32, bytes32)
func (_XNOUsers *XNOUsersCaller) GetGame(opts *bind.CallOpts, gameAddress common.Address) (string, bool, [32]byte, [32]byte, [32]byte, error) {
	var (
		ret0 = new(string)
		ret1 = new(bool)
		ret2 = new([32]byte)
		ret3 = new([32]byte)
		ret4 = new([32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
	}
	err := _XNOUsers.contract.Call(opts, out, "getGame", gameAddress)
	return *ret0, *ret1, *ret2, *ret3, *ret4, err
}

// GetGame is a free data retrieval call binding the contract method 0x8d1a4b28.
//
// Solidity: function getGame(gameAddress address) constant returns(string, bool, bytes32, bytes32, bytes32)
func (_XNOUsers *XNOUsersSession) GetGame(gameAddress common.Address) (string, bool, [32]byte, [32]byte, [32]byte, error) {
	return _XNOUsers.Contract.GetGame(&_XNOUsers.CallOpts, gameAddress)
}

// GetGame is a free data retrieval call binding the contract method 0x8d1a4b28.
//
// Solidity: function getGame(gameAddress address) constant returns(string, bool, bytes32, bytes32, bytes32)
func (_XNOUsers *XNOUsersCallerSession) GetGame(gameAddress common.Address) (string, bool, [32]byte, [32]byte, [32]byte, error) {
	return _XNOUsers.Contract.GetGame(&_XNOUsers.CallOpts, gameAddress)
}

// GetGameImages is a free data retrieval call binding the contract method 0x34c16694.
//
// Solidity: function getGameImages(userAddress address) constant returns(bytes32)
func (_XNOUsers *XNOUsersCaller) GetGameImages(opts *bind.CallOpts, userAddress common.Address) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _XNOUsers.contract.Call(opts, out, "getGameImages", userAddress)
	return *ret0, err
}

// GetGameImages is a free data retrieval call binding the contract method 0x34c16694.
//
// Solidity: function getGameImages(userAddress address) constant returns(bytes32)
func (_XNOUsers *XNOUsersSession) GetGameImages(userAddress common.Address) ([32]byte, error) {
	return _XNOUsers.Contract.GetGameImages(&_XNOUsers.CallOpts, userAddress)
}

// GetGameImages is a free data retrieval call binding the contract method 0x34c16694.
//
// Solidity: function getGameImages(userAddress address) constant returns(bytes32)
func (_XNOUsers *XNOUsersCallerSession) GetGameImages(userAddress common.Address) ([32]byte, error) {
	return _XNOUsers.Contract.GetGameImages(&_XNOUsers.CallOpts, userAddress)
}

// GetImage is a free data retrieval call binding the contract method 0x6ced1ae9.
//
// Solidity: function getImage(SHA256notaryHash bytes32) constant returns(string, uint256)
func (_XNOUsers *XNOUsersCaller) GetImage(opts *bind.CallOpts, SHA256notaryHash [32]byte) (string, *big.Int, error) {
	var (
		ret0 = new(string)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _XNOUsers.contract.Call(opts, out, "getImage", SHA256notaryHash)
	return *ret0, *ret1, err
}

// GetImage is a free data retrieval call binding the contract method 0x6ced1ae9.
//
// Solidity: function getImage(SHA256notaryHash bytes32) constant returns(string, uint256)
func (_XNOUsers *XNOUsersSession) GetImage(SHA256notaryHash [32]byte) (string, *big.Int, error) {
	return _XNOUsers.Contract.GetImage(&_XNOUsers.CallOpts, SHA256notaryHash)
}

// GetImage is a free data retrieval call binding the contract method 0x6ced1ae9.
//
// Solidity: function getImage(SHA256notaryHash bytes32) constant returns(string, uint256)
func (_XNOUsers *XNOUsersCallerSession) GetImage(SHA256notaryHash [32]byte) (string, *big.Int, error) {
	return _XNOUsers.Contract.GetImage(&_XNOUsers.CallOpts, SHA256notaryHash)
}

// AddLogo is a paid mutator transaction binding the contract method 0x25946775.
//
// Solidity: function addLogo(imageURL string, SHA256notaryHash bytes32) returns(success bool)
func (_XNOUsers *XNOUsersTransactor) AddLogo(opts *bind.TransactOpts, imageURL string, SHA256notaryHash [32]byte) (*types.Transaction, error) {
	return _XNOUsers.contract.Transact(opts, "addLogo", imageURL, SHA256notaryHash)
}

// AddLogo is a paid mutator transaction binding the contract method 0x25946775.
//
// Solidity: function addLogo(imageURL string, SHA256notaryHash bytes32) returns(success bool)
func (_XNOUsers *XNOUsersSession) AddLogo(imageURL string, SHA256notaryHash [32]byte) (*types.Transaction, error) {
	return _XNOUsers.Contract.AddLogo(&_XNOUsers.TransactOpts, imageURL, SHA256notaryHash)
}

// AddLogo is a paid mutator transaction binding the contract method 0x25946775.
//
// Solidity: function addLogo(imageURL string, SHA256notaryHash bytes32) returns(success bool)
func (_XNOUsers *XNOUsersTransactorSession) AddLogo(imageURL string, SHA256notaryHash [32]byte) (*types.Transaction, error) {
	return _XNOUsers.Contract.AddLogo(&_XNOUsers.TransactOpts, imageURL, SHA256notaryHash)
}

// RegisterNewUser is a paid mutator transaction binding the contract method 0x8b907bbc.
//
// Solidity: function registerNewUser(id string, isServer bool, country bytes32, state bytes32) returns(success bool)
func (_XNOUsers *XNOUsersTransactor) RegisterNewUser(opts *bind.TransactOpts, id string, isServer bool, country [32]byte, state [32]byte) (*types.Transaction, error) {
	return _XNOUsers.contract.Transact(opts, "registerNewUser", id, isServer, country, state)
}

// RegisterNewUser is a paid mutator transaction binding the contract method 0x8b907bbc.
//
// Solidity: function registerNewUser(id string, isServer bool, country bytes32, state bytes32) returns(success bool)
func (_XNOUsers *XNOUsersSession) RegisterNewUser(id string, isServer bool, country [32]byte, state [32]byte) (*types.Transaction, error) {
	return _XNOUsers.Contract.RegisterNewUser(&_XNOUsers.TransactOpts, id, isServer, country, state)
}

// RegisterNewUser is a paid mutator transaction binding the contract method 0x8b907bbc.
//
// Solidity: function registerNewUser(id string, isServer bool, country bytes32, state bytes32) returns(success bool)
func (_XNOUsers *XNOUsersTransactorSession) RegisterNewUser(id string, isServer bool, country [32]byte, state [32]byte) (*types.Transaction, error) {
	return _XNOUsers.Contract.RegisterNewUser(&_XNOUsers.TransactOpts, id, isServer, country, state)
}

// RemoveImage is a paid mutator transaction binding the contract method 0x9f6c9358.
//
// Solidity: function removeImage(imgHash bytes32) returns(success bool)
func (_XNOUsers *XNOUsersTransactor) RemoveImage(opts *bind.TransactOpts, imgHash [32]byte) (*types.Transaction, error) {
	return _XNOUsers.contract.Transact(opts, "removeImage", imgHash)
}

// RemoveImage is a paid mutator transaction binding the contract method 0x9f6c9358.
//
// Solidity: function removeImage(imgHash bytes32) returns(success bool)
func (_XNOUsers *XNOUsersSession) RemoveImage(imgHash [32]byte) (*types.Transaction, error) {
	return _XNOUsers.Contract.RemoveImage(&_XNOUsers.TransactOpts, imgHash)
}

// RemoveImage is a paid mutator transaction binding the contract method 0x9f6c9358.
//
// Solidity: function removeImage(imgHash bytes32) returns(success bool)
func (_XNOUsers *XNOUsersTransactorSession) RemoveImage(imgHash [32]byte) (*types.Transaction, error) {
	return _XNOUsers.Contract.RemoveImage(&_XNOUsers.TransactOpts, imgHash)
}

// RemoveUser is a paid mutator transaction binding the contract method 0x98575188.
//
// Solidity: function removeUser(gameAddress address) returns(success bool)
func (_XNOUsers *XNOUsersTransactor) RemoveUser(opts *bind.TransactOpts, gameAddress common.Address) (*types.Transaction, error) {
	return _XNOUsers.contract.Transact(opts, "removeUser", gameAddress)
}

// RemoveUser is a paid mutator transaction binding the contract method 0x98575188.
//
// Solidity: function removeUser(gameAddress address) returns(success bool)
func (_XNOUsers *XNOUsersSession) RemoveUser(gameAddress common.Address) (*types.Transaction, error) {
	return _XNOUsers.Contract.RemoveUser(&_XNOUsers.TransactOpts, gameAddress)
}

// RemoveUser is a paid mutator transaction binding the contract method 0x98575188.
//
// Solidity: function removeUser(gameAddress address) returns(success bool)
func (_XNOUsers *XNOUsersTransactorSession) RemoveUser(gameAddress common.Address) (*types.Transaction, error) {
	return _XNOUsers.Contract.RemoveUser(&_XNOUsers.TransactOpts, gameAddress)
}
