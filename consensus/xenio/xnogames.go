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

// XNOGamesABI is the input ABI used to generate the binding from.
const XNOGamesABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"imageURL\",\"type\":\"string\"},{\"name\":\"SHA256notaryHash\",\"type\":\"bytes32\"}],\"name\":\"addLogo\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"userAddress\",\"type\":\"address\"}],\"name\":\"getGameImages\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"SHA256notaryHash\",\"type\":\"bytes32\"}],\"name\":\"getImage\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAllImages\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"gameAddress\",\"type\":\"address\"}],\"name\":\"getGame\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"gameAddress\",\"type\":\"address\"}],\"name\":\"removeUser\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"imgHash\",\"type\":\"bytes32\"}],\"name\":\"removeImage\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"publisher\",\"type\":\"string\"},{\"name\":\"developer\",\"type\":\"string\"},{\"name\":\"country\",\"type\":\"bytes32\"},{\"name\":\"state\",\"type\":\"bytes32\"}],\"name\":\"registerNewGame\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAllGames\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"}]"

// XNOGamesBin is the compiled bytecode used for deploying new contracts.
const XNOGamesBin = `0x606060405260008054600160a060020a033316600160a060020a0319909116179055610ccc806100306000396000f3006060604052600436106100985763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416632594677581146100a557806334c166941461010c5780636ced1ae91461013d5780637b7284fc146101d25780638d1a4b281461023857806398575188146103ba5780639f6c9358146103d9578063ad456441146103ef578063db1c45f9146104cb575b34156100a357600080fd5b005b34156100b057600080fd5b6100f860046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965050933593506104de92505050565b604051901515815260200160405180910390f35b341561011757600080fd5b61012b600160a060020a03600435166105d7565b60405190815260200160405180910390f35b341561014857600080fd5b6101536004356105f5565b6040518080602001838152602001828103825284818151815260200191508051906020019080838360005b8381101561019657808201518382015260200161017e565b50505050905090810190601f1680156101c35780820380516001836020036101000a031916815260200191505b50935050505060405180910390f35b34156101dd57600080fd5b6101e56106c2565b60405160208082528190810183818151815260200191508051906020019060200280838360005b8381101561022457808201518382015260200161020c565b505050509050019250505060405180910390f35b341561024357600080fd5b610257600160a060020a0360043516610722565b604051606081018490526080810183905260a0810182905260c080825281906020820190604083019083018a818151815260200191508051906020019080838360005b838110156102b257808201518382015260200161029a565b50505050905090810190601f1680156102df5780820380516001836020036101000a031916815260200191505b50848103835289818151815260200191508051906020019080838360005b838110156103155780820151838201526020016102fd565b50505050905090810190601f1680156103425780820380516001836020036101000a031916815260200191505b50848103825288818151815260200191508051906020019080838360005b83811015610378578082015183820152602001610360565b50505050905090810190601f1680156103a55780820380516001836020036101000a031916815260200191505b50995050505050505050505060405180910390f35b34156103c557600080fd5b6100f8600160a060020a036004351661095c565b34156103e457600080fd5b6100f86004356109d6565b34156103fa57600080fd5b6100f860046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284375094965050843594602001359350610a1a92505050565b34156104d657600080fd5b6101e5610b1f565b33600160a060020a0381166000908152600160208190526040822054919291600291811615610100026000190116041561051b57600091506105d0565b83511561052b57600091506105d0565b600083815260036020526040902054600261010060018316150260001901909116041561055b57600091506105d0565b600480546001810161056d8382610b86565b506000918252602080832091909101859055848252600390526040902084805161059b929160200190610baf565b50600083815260036020908152604080832042600191820155600160a060020a03851684529182905290912060050184905591505b5092915050565b600160a060020a031660009081526001602052604090206005015490565b6105fd610c2d565b600082815260036020908152604080832060018082015482549294909385936002610100948216159490940260001901169290920491601f83018290048202909101905190810160405280929190818152602001828054600181600116156101000203166002900480156106b25780601f10610687576101008083540402835291602001916106b2565b820191906000526020600020905b81548152906001019060200180831161069557829003601f168201915b5050505050915091509150915091565b6106ca610c2d565b600480548060200260200160405190810160405280929190818152602001828054801561071757602002820191906000526020600020905b81548152600190910190602001808311610702575b505050505090505b90565b61072a610c2d565b610732610c2d565b61073a610c2d565b600160a060020a03841660009081526001602081815260408084206003810154600482015460058301548354889788979596878201966002808a01979695948a948116156101000260001901160491601f83018290048202909101905190810160405280929190818152602001828054600181600116156101000203166002900480156108085780601f106107dd57610100808354040283529160200191610808565b820191906000526020600020905b8154815290600101906020018083116107eb57829003601f168201915b50505050509550848054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156108a45780601f10610879576101008083540402835291602001916108a4565b820191906000526020600020905b81548152906001019060200180831161088757829003601f168201915b50505050509450838054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156109405780601f1061091557610100808354040283529160200191610940565b820191906000526020600020905b81548152906001019060200180831161092357829003601f168201915b5050505050935095509550955095509550955091939550919395565b6000805433600160a060020a0390811691161461097857600080fd5b600160a060020a03821660009081526001602052604081209061099b8282610c3f565b6109a9600183016000610c3f565b6109b7600283016000610c3f565b5060006003820181905560048201819055600590910155506001919050565b6000805433600160a060020a039081169116146109f257600080fd5b600082815260036020526040812090610a0b8282610c3f565b50600060019182015592915050565b33600160a060020a0381166000908152600160205260408120909190878051610a47929160200190610baf565b50600160a060020a038116600090815260016020819052604090912001868051610a75929160200190610baf565b50600160a060020a0381166000908152600160205260409020600201858051610aa2929160200190610baf565b50600160a060020a0381166000908152600160208190526040909120600481018590556003018590556002805490918101610add8382610b86565b5060009182526020909120018054600160a060020a03831673ffffffffffffffffffffffffffffffffffffffff19909116179055600191505095945050505050565b610b27610c2d565b600280548060200260200160405190810160405280929190818152602001828054801561071757602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610b5f575050505050905090565b815481835581811511610baa57600083815260209020610baa918101908301610c86565b505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610bf057805160ff1916838001178555610c1d565b82800160010185558215610c1d579182015b82811115610c1d578251825591602001919060010190610c02565b50610c29929150610c86565b5090565b60206040519081016040526000815290565b50805460018160011615610100020316600290046000825580601f10610c655750610c83565b601f016020900490600052602060002090810190610c839190610c86565b50565b61071f91905b80821115610c295760008155600101610c8c5600a165627a7a723058202138cba45fda268567b485c791a824c88f688183f5a31ddc4b3c342dac657d9d0029`

// DeployXNOGames deploys a new Ethereum contract, binding an instance of XNOGames to it.
func DeployXNOGames(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *XNOGames, error) {
	parsed, err := abi.JSON(strings.NewReader(XNOGamesABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(XNOGamesBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &XNOGames{XNOGamesCaller: XNOGamesCaller{contract: contract}, XNOGamesTransactor: XNOGamesTransactor{contract: contract}}, nil
}

// XNOGames is an auto generated Go binding around an Ethereum contract.
type XNOGames struct {
	XNOGamesCaller     // Read-only binding to the contract
	XNOGamesTransactor // Write-only binding to the contract
}

// XNOGamesCaller is an auto generated read-only Go binding around an Ethereum contract.
type XNOGamesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// XNOGamesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type XNOGamesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// XNOGamesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type XNOGamesSession struct {
	Contract     *XNOGames         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// XNOGamesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type XNOGamesCallerSession struct {
	Contract *XNOGamesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// XNOGamesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type XNOGamesTransactorSession struct {
	Contract     *XNOGamesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// XNOGamesRaw is an auto generated low-level Go binding around an Ethereum contract.
type XNOGamesRaw struct {
	Contract *XNOGames // Generic contract binding to access the raw methods on
}

// XNOGamesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type XNOGamesCallerRaw struct {
	Contract *XNOGamesCaller // Generic read-only contract binding to access the raw methods on
}

// XNOGamesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type XNOGamesTransactorRaw struct {
	Contract *XNOGamesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewXNOGames creates a new instance of XNOGames, bound to a specific deployed contract.
func NewXNOGames(address common.Address, backend bind.ContractBackend) (*XNOGames, error) {
	contract, err := bindXNOGames(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &XNOGames{XNOGamesCaller: XNOGamesCaller{contract: contract}, XNOGamesTransactor: XNOGamesTransactor{contract: contract}}, nil
}

// NewXNOGamesCaller creates a new read-only instance of XNOGames, bound to a specific deployed contract.
func NewXNOGamesCaller(address common.Address, caller bind.ContractCaller) (*XNOGamesCaller, error) {
	contract, err := bindXNOGames(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &XNOGamesCaller{contract: contract}, nil
}

// NewXNOGamesTransactor creates a new write-only instance of XNOGames, bound to a specific deployed contract.
func NewXNOGamesTransactor(address common.Address, transactor bind.ContractTransactor) (*XNOGamesTransactor, error) {
	contract, err := bindXNOGames(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &XNOGamesTransactor{contract: contract}, nil
}

// bindXNOGames binds a generic wrapper to an already deployed contract.
func bindXNOGames(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(XNOGamesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_XNOGames *XNOGamesRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _XNOGames.Contract.XNOGamesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_XNOGames *XNOGamesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _XNOGames.Contract.XNOGamesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_XNOGames *XNOGamesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _XNOGames.Contract.XNOGamesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_XNOGames *XNOGamesCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _XNOGames.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_XNOGames *XNOGamesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _XNOGames.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_XNOGames *XNOGamesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _XNOGames.Contract.contract.Transact(opts, method, params...)
}

// GetAllGames is a free data retrieval call binding the contract method 0xdb1c45f9.
//
// Solidity: function getAllGames() constant returns(address[])
func (_XNOGames *XNOGamesCaller) GetAllGames(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _XNOGames.contract.Call(opts, out, "getAllGames")
	return *ret0, err
}

// GetAllGames is a free data retrieval call binding the contract method 0xdb1c45f9.
//
// Solidity: function getAllGames() constant returns(address[])
func (_XNOGames *XNOGamesSession) GetAllGames() ([]common.Address, error) {
	return _XNOGames.Contract.GetAllGames(&_XNOGames.CallOpts)
}

// GetAllGames is a free data retrieval call binding the contract method 0xdb1c45f9.
//
// Solidity: function getAllGames() constant returns(address[])
func (_XNOGames *XNOGamesCallerSession) GetAllGames() ([]common.Address, error) {
	return _XNOGames.Contract.GetAllGames(&_XNOGames.CallOpts)
}

// GetAllImages is a free data retrieval call binding the contract method 0x7b7284fc.
//
// Solidity: function getAllImages() constant returns(bytes32[])
func (_XNOGames *XNOGamesCaller) GetAllImages(opts *bind.CallOpts) ([][32]byte, error) {
	var (
		ret0 = new([][32]byte)
	)
	out := ret0
	err := _XNOGames.contract.Call(opts, out, "getAllImages")
	return *ret0, err
}

// GetAllImages is a free data retrieval call binding the contract method 0x7b7284fc.
//
// Solidity: function getAllImages() constant returns(bytes32[])
func (_XNOGames *XNOGamesSession) GetAllImages() ([][32]byte, error) {
	return _XNOGames.Contract.GetAllImages(&_XNOGames.CallOpts)
}

// GetAllImages is a free data retrieval call binding the contract method 0x7b7284fc.
//
// Solidity: function getAllImages() constant returns(bytes32[])
func (_XNOGames *XNOGamesCallerSession) GetAllImages() ([][32]byte, error) {
	return _XNOGames.Contract.GetAllImages(&_XNOGames.CallOpts)
}

// GetGame is a free data retrieval call binding the contract method 0x8d1a4b28.
//
// Solidity: function getGame(gameAddress address) constant returns(string, string, string, bytes32, bytes32, bytes32)
func (_XNOGames *XNOGamesCaller) GetGame(opts *bind.CallOpts, gameAddress common.Address) (string, string, string, [32]byte, [32]byte, [32]byte, error) {
	var (
		ret0 = new(string)
		ret1 = new(string)
		ret2 = new(string)
		ret3 = new([32]byte)
		ret4 = new([32]byte)
		ret5 = new([32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
	}
	err := _XNOGames.contract.Call(opts, out, "getGame", gameAddress)
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, err
}

// GetGame is a free data retrieval call binding the contract method 0x8d1a4b28.
//
// Solidity: function getGame(gameAddress address) constant returns(string, string, string, bytes32, bytes32, bytes32)
func (_XNOGames *XNOGamesSession) GetGame(gameAddress common.Address) (string, string, string, [32]byte, [32]byte, [32]byte, error) {
	return _XNOGames.Contract.GetGame(&_XNOGames.CallOpts, gameAddress)
}

// GetGame is a free data retrieval call binding the contract method 0x8d1a4b28.
//
// Solidity: function getGame(gameAddress address) constant returns(string, string, string, bytes32, bytes32, bytes32)
func (_XNOGames *XNOGamesCallerSession) GetGame(gameAddress common.Address) (string, string, string, [32]byte, [32]byte, [32]byte, error) {
	return _XNOGames.Contract.GetGame(&_XNOGames.CallOpts, gameAddress)
}

// GetGameImages is a free data retrieval call binding the contract method 0x34c16694.
//
// Solidity: function getGameImages(userAddress address) constant returns(bytes32)
func (_XNOGames *XNOGamesCaller) GetGameImages(opts *bind.CallOpts, userAddress common.Address) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _XNOGames.contract.Call(opts, out, "getGameImages", userAddress)
	return *ret0, err
}

// GetGameImages is a free data retrieval call binding the contract method 0x34c16694.
//
// Solidity: function getGameImages(userAddress address) constant returns(bytes32)
func (_XNOGames *XNOGamesSession) GetGameImages(userAddress common.Address) ([32]byte, error) {
	return _XNOGames.Contract.GetGameImages(&_XNOGames.CallOpts, userAddress)
}

// GetGameImages is a free data retrieval call binding the contract method 0x34c16694.
//
// Solidity: function getGameImages(userAddress address) constant returns(bytes32)
func (_XNOGames *XNOGamesCallerSession) GetGameImages(userAddress common.Address) ([32]byte, error) {
	return _XNOGames.Contract.GetGameImages(&_XNOGames.CallOpts, userAddress)
}

// GetImage is a free data retrieval call binding the contract method 0x6ced1ae9.
//
// Solidity: function getImage(SHA256notaryHash bytes32) constant returns(string, uint256)
func (_XNOGames *XNOGamesCaller) GetImage(opts *bind.CallOpts, SHA256notaryHash [32]byte) (string, *big.Int, error) {
	var (
		ret0 = new(string)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _XNOGames.contract.Call(opts, out, "getImage", SHA256notaryHash)
	return *ret0, *ret1, err
}

// GetImage is a free data retrieval call binding the contract method 0x6ced1ae9.
//
// Solidity: function getImage(SHA256notaryHash bytes32) constant returns(string, uint256)
func (_XNOGames *XNOGamesSession) GetImage(SHA256notaryHash [32]byte) (string, *big.Int, error) {
	return _XNOGames.Contract.GetImage(&_XNOGames.CallOpts, SHA256notaryHash)
}

// GetImage is a free data retrieval call binding the contract method 0x6ced1ae9.
//
// Solidity: function getImage(SHA256notaryHash bytes32) constant returns(string, uint256)
func (_XNOGames *XNOGamesCallerSession) GetImage(SHA256notaryHash [32]byte) (string, *big.Int, error) {
	return _XNOGames.Contract.GetImage(&_XNOGames.CallOpts, SHA256notaryHash)
}

// AddLogo is a paid mutator transaction binding the contract method 0x25946775.
//
// Solidity: function addLogo(imageURL string, SHA256notaryHash bytes32) returns(success bool)
func (_XNOGames *XNOGamesTransactor) AddLogo(opts *bind.TransactOpts, imageURL string, SHA256notaryHash [32]byte) (*types.Transaction, error) {
	return _XNOGames.contract.Transact(opts, "addLogo", imageURL, SHA256notaryHash)
}

// AddLogo is a paid mutator transaction binding the contract method 0x25946775.
//
// Solidity: function addLogo(imageURL string, SHA256notaryHash bytes32) returns(success bool)
func (_XNOGames *XNOGamesSession) AddLogo(imageURL string, SHA256notaryHash [32]byte) (*types.Transaction, error) {
	return _XNOGames.Contract.AddLogo(&_XNOGames.TransactOpts, imageURL, SHA256notaryHash)
}

// AddLogo is a paid mutator transaction binding the contract method 0x25946775.
//
// Solidity: function addLogo(imageURL string, SHA256notaryHash bytes32) returns(success bool)
func (_XNOGames *XNOGamesTransactorSession) AddLogo(imageURL string, SHA256notaryHash [32]byte) (*types.Transaction, error) {
	return _XNOGames.Contract.AddLogo(&_XNOGames.TransactOpts, imageURL, SHA256notaryHash)
}

// RegisterNewGame is a paid mutator transaction binding the contract method 0xad456441.
//
// Solidity: function registerNewGame(name string, publisher string, developer string, country bytes32, state bytes32) returns(success bool)
func (_XNOGames *XNOGamesTransactor) RegisterNewGame(opts *bind.TransactOpts, name string, publisher string, developer string, country [32]byte, state [32]byte) (*types.Transaction, error) {
	return _XNOGames.contract.Transact(opts, "registerNewGame", name, publisher, developer, country, state)
}

// RegisterNewGame is a paid mutator transaction binding the contract method 0xad456441.
//
// Solidity: function registerNewGame(name string, publisher string, developer string, country bytes32, state bytes32) returns(success bool)
func (_XNOGames *XNOGamesSession) RegisterNewGame(name string, publisher string, developer string, country [32]byte, state [32]byte) (*types.Transaction, error) {
	return _XNOGames.Contract.RegisterNewGame(&_XNOGames.TransactOpts, name, publisher, developer, country, state)
}

// RegisterNewGame is a paid mutator transaction binding the contract method 0xad456441.
//
// Solidity: function registerNewGame(name string, publisher string, developer string, country bytes32, state bytes32) returns(success bool)
func (_XNOGames *XNOGamesTransactorSession) RegisterNewGame(name string, publisher string, developer string, country [32]byte, state [32]byte) (*types.Transaction, error) {
	return _XNOGames.Contract.RegisterNewGame(&_XNOGames.TransactOpts, name, publisher, developer, country, state)
}

// RemoveImage is a paid mutator transaction binding the contract method 0x9f6c9358.
//
// Solidity: function removeImage(imgHash bytes32) returns(success bool)
func (_XNOGames *XNOGamesTransactor) RemoveImage(opts *bind.TransactOpts, imgHash [32]byte) (*types.Transaction, error) {
	return _XNOGames.contract.Transact(opts, "removeImage", imgHash)
}

// RemoveImage is a paid mutator transaction binding the contract method 0x9f6c9358.
//
// Solidity: function removeImage(imgHash bytes32) returns(success bool)
func (_XNOGames *XNOGamesSession) RemoveImage(imgHash [32]byte) (*types.Transaction, error) {
	return _XNOGames.Contract.RemoveImage(&_XNOGames.TransactOpts, imgHash)
}

// RemoveImage is a paid mutator transaction binding the contract method 0x9f6c9358.
//
// Solidity: function removeImage(imgHash bytes32) returns(success bool)
func (_XNOGames *XNOGamesTransactorSession) RemoveImage(imgHash [32]byte) (*types.Transaction, error) {
	return _XNOGames.Contract.RemoveImage(&_XNOGames.TransactOpts, imgHash)
}

// RemoveUser is a paid mutator transaction binding the contract method 0x98575188.
//
// Solidity: function removeUser(gameAddress address) returns(success bool)
func (_XNOGames *XNOGamesTransactor) RemoveUser(opts *bind.TransactOpts, gameAddress common.Address) (*types.Transaction, error) {
	return _XNOGames.contract.Transact(opts, "removeUser", gameAddress)
}

// RemoveUser is a paid mutator transaction binding the contract method 0x98575188.
//
// Solidity: function removeUser(gameAddress address) returns(success bool)
func (_XNOGames *XNOGamesSession) RemoveUser(gameAddress common.Address) (*types.Transaction, error) {
	return _XNOGames.Contract.RemoveUser(&_XNOGames.TransactOpts, gameAddress)
}

// RemoveUser is a paid mutator transaction binding the contract method 0x98575188.
//
// Solidity: function removeUser(gameAddress address) returns(success bool)
func (_XNOGames *XNOGamesTransactorSession) RemoveUser(gameAddress common.Address) (*types.Transaction, error) {
	return _XNOGames.Contract.RemoveUser(&_XNOGames.TransactOpts, gameAddress)
}
