// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package xno

import (
	"math/big"
	"strings"

	"github.com/xenioplatform/go-xenio/common"
	"github.com/xenioplatform/go-xenio/core/types"
	"github.com/xenioplatform/go-xenio/accounts/abi"
	"github.com/xenioplatform/go-xenio/accounts/abi/bind"
)

// EternalStorageABI is the input ABI used to generate the binding from.
const EternalStorageABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"record\",\"type\":\"bytes32\"}],\"name\":\"getBytes32Value\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"record\",\"type\":\"bytes32\"}],\"name\":\"deleteAddressValue\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"record\",\"type\":\"bytes32\"}],\"name\":\"deleteBytesValue\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"record\",\"type\":\"bytes32\"}],\"name\":\"deleteBytes32Value\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"record\",\"type\":\"bytes32\"}],\"name\":\"getBooleanValue\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"record\",\"type\":\"bytes32\"},{\"name\":\"value\",\"type\":\"bytes32\"}],\"name\":\"setBytes32Value\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"record\",\"type\":\"bytes32\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setUIntValue\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"record\",\"type\":\"bytes32\"}],\"name\":\"deleteBooleanValue\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"record\",\"type\":\"bytes32\"},{\"name\":\"value\",\"type\":\"bool\"}],\"name\":\"setBooleanValue\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"record\",\"type\":\"bytes32\"}],\"name\":\"getBytesValue\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"record\",\"type\":\"bytes32\"}],\"name\":\"getAddressValue\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"record\",\"type\":\"bytes32\"},{\"name\":\"value\",\"type\":\"address\"}],\"name\":\"setAddressValue\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"passOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"record\",\"type\":\"bytes32\"}],\"name\":\"deleteIntValue\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"record\",\"type\":\"bytes32\"}],\"name\":\"getIntValue\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"record\",\"type\":\"bytes32\"}],\"name\":\"deleteUIntValue\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"record\",\"type\":\"bytes32\"}],\"name\":\"getStringValue\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"record\",\"type\":\"bytes32\"},{\"name\":\"value\",\"type\":\"int256\"}],\"name\":\"setIntValue\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"record\",\"type\":\"bytes32\"}],\"name\":\"deleteStringValue\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"record\",\"type\":\"bytes32\"}],\"name\":\"getUIntValue\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"record\",\"type\":\"bytes32\"},{\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"setBytesValue\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"record\",\"type\":\"bytes32\"},{\"name\":\"value\",\"type\":\"string\"}],\"name\":\"setStringValue\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// EternalStorageBin is the compiled bytecode used for deploying new contracts.
const EternalStorageBin = `0x6060604052341561000f57600080fd5b60008054600160a060020a033316600160a060020a0319909116179055610b068061003b6000396000f30060606040526004361061013d5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663025ec81a8114610142578063043106c01461016a5780630c55d92514610182578063124f24181461019857806317e7dd22146101ae57806325cf512d146101d85780633562fd20146101f15780633cc1635c1461020a5780633eba9ed21461022057806341c0e1b51461023b57806344bfa56e1461024e5780634c77e5ba146102db5780635a2bf25a1461030d578063620996df1461032f5780638267a9ee1461034e578063893d20e8146103645780639007127b1461037757806393fe42481461038d578063a209a29c146103a3578063a77aa49e146103b9578063ba69fcaa146103d2578063bdc963d8146103e8578063c9a52d2c146103fe578063f586606614610454575b600080fd5b341561014d57600080fd5b6101586004356104aa565b60405190815260200160405180910390f35b341561017557600080fd5b6101806004356104bc565b005b341561018d57600080fd5b610180600435610502565b34156101a357600080fd5b610180600435610537565b34156101b957600080fd5b6101c4600435610563565b604051901515815260200160405180910390f35b34156101e357600080fd5b610180600435602435610578565b34156101fc57600080fd5b6101806004356024356105a5565b341561021557600080fd5b6101806004356105d2565b341561022b57600080fd5b6101806004356024351515610605565b341561024657600080fd5b610180610640565b341561025957600080fd5b610264600435610669565b60405160208082528190810183818151815260200191508051906020019080838360005b838110156102a0578082015183820152602001610288565b50505050905090810190601f1680156102cd5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156102e657600080fd5b6102f160043561072c565b604051600160a060020a03909116815260200160405180910390f35b341561031857600080fd5b610180600435600160a060020a0360243516610747565b341561033a57600080fd5b610180600160a060020a036004351661079d565b341561035957600080fd5b6101806004356107fc565b341561036f57600080fd5b6102f1610828565b341561038257600080fd5b610158600435610838565b341561039857600080fd5b61018060043561084a565b34156103ae57600080fd5b610264600435610876565b34156103c457600080fd5b610180600435602435610902565b34156103dd57600080fd5b61018060043561092f565b34156103f357600080fd5b610158600435610961565b341561040957600080fd5b610180600480359060446024803590810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965061097395505050505050565b341561045f57600080fd5b610180600480359060446024803590810190830135806020601f820181900481020160405190810160405281815292919060208401838380828437509496506109b295505050505050565b60009081526005602052604090205490565b60005433600160a060020a039081169116146104d757600080fd5b6000908152600360205260409020805473ffffffffffffffffffffffffffffffffffffffff19169055565b60005433600160a060020a0390811691161461051d57600080fd5b6000818152600460205260408120610534916109ec565b50565b60005433600160a060020a0390811691161461055257600080fd5b600090815260056020526040812055565b60009081526006602052604090205460ff1690565b60005433600160a060020a0390811691161461059357600080fd5b60009182526005602052604090912055565b60005433600160a060020a039081169116146105c057600080fd5b60009182526001602052604090912055565b60005433600160a060020a039081169116146105ed57600080fd5b6000908152600660205260409020805460ff19169055565b60005433600160a060020a0390811691161461062057600080fd5b600091825260066020526040909120805460ff1916911515919091179055565b60005433600160a060020a0390811691161461065b57600080fd5b600054600160a060020a0316ff5b610671610a30565b6004600083600019166000191681526020019081526020016000208054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156107205780601f106106f557610100808354040283529160200191610720565b820191906000526020600020905b81548152906001019060200180831161070357829003601f168201915b50505050509050919050565b600090815260036020526040902054600160a060020a031690565b60005433600160a060020a0390811691161461076257600080fd5b600091825260036020526040909120805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03909216919091179055565b60005433600160a060020a039081169116146107b857600080fd5b600160a060020a03811615156107cd57600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60005433600160a060020a0390811691161461081757600080fd5b600090815260076020526040812055565b600054600160a060020a03165b90565b60009081526007602052604090205490565b60005433600160a060020a0390811691161461086557600080fd5b600090815260016020526040812055565b61087e610a30565b6002600083600019166000191681526020019081526020016000208054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156107205780601f106106f557610100808354040283529160200191610720565b60005433600160a060020a0390811691161461091d57600080fd5b60009182526007602052604090912055565b60005433600160a060020a0390811691161461094a57600080fd5b6000818152600260205260408120610534916109ec565b60009081526001602052604090205490565b60005433600160a060020a0390811691161461098e57600080fd5b60008281526004602052604090208180516109ad929160200190610a42565b505050565b60005433600160a060020a039081169116146109cd57600080fd5b60008281526002602052604090208180516109ad929160200190610a42565b50805460018160011615610100020316600290046000825580601f10610a125750610534565b601f0160209004906000526020600020908101906105349190610ac0565b60206040519081016040526000815290565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610a8357805160ff1916838001178555610ab0565b82800160010185558215610ab0579182015b82811115610ab0578251825591602001919060010190610a95565b50610abc929150610ac0565b5090565b61083591905b80821115610abc5760008155600101610ac65600a165627a7a7230582089078e369701338df82d3be48ed51f1ec92dc24b78d1b6f8c6cd51c593647c300029`

// DeployEternalStorage deploys a new Ethereum contract, binding an instance of EternalStorage to it.
func DeployEternalStorage(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EternalStorage, error) {
	parsed, err := abi.JSON(strings.NewReader(EternalStorageABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(EternalStorageBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EternalStorage{EternalStorageCaller: EternalStorageCaller{contract: contract}, EternalStorageTransactor: EternalStorageTransactor{contract: contract}}, nil
}

// EternalStorage is an auto generated Go binding around an Ethereum contract.
type EternalStorage struct {
	EternalStorageCaller     // Read-only binding to the contract
	EternalStorageTransactor // Write-only binding to the contract
}

// EternalStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type EternalStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EternalStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EternalStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EternalStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EternalStorageSession struct {
	Contract     *EternalStorage   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EternalStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EternalStorageCallerSession struct {
	Contract *EternalStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// EternalStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EternalStorageTransactorSession struct {
	Contract     *EternalStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// EternalStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type EternalStorageRaw struct {
	Contract *EternalStorage // Generic contract binding to access the raw methods on
}

// EternalStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EternalStorageCallerRaw struct {
	Contract *EternalStorageCaller // Generic read-only contract binding to access the raw methods on
}

// EternalStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EternalStorageTransactorRaw struct {
	Contract *EternalStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEternalStorage creates a new instance of EternalStorage, bound to a specific deployed contract.
func NewEternalStorage(address common.Address, backend bind.ContractBackend) (*EternalStorage, error) {
	contract, err := bindEternalStorage(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EternalStorage{EternalStorageCaller: EternalStorageCaller{contract: contract}, EternalStorageTransactor: EternalStorageTransactor{contract: contract}}, nil
}

// NewEternalStorageCaller creates a new read-only instance of EternalStorage, bound to a specific deployed contract.
func NewEternalStorageCaller(address common.Address, caller bind.ContractCaller) (*EternalStorageCaller, error) {
	contract, err := bindEternalStorage(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &EternalStorageCaller{contract: contract}, nil
}

// NewEternalStorageTransactor creates a new write-only instance of EternalStorage, bound to a specific deployed contract.
func NewEternalStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*EternalStorageTransactor, error) {
	contract, err := bindEternalStorage(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &EternalStorageTransactor{contract: contract}, nil
}

// bindEternalStorage binds a generic wrapper to an already deployed contract.
func bindEternalStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EternalStorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EternalStorage *EternalStorageRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EternalStorage.Contract.EternalStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EternalStorage *EternalStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EternalStorage.Contract.EternalStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EternalStorage *EternalStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EternalStorage.Contract.EternalStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EternalStorage *EternalStorageCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EternalStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EternalStorage *EternalStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EternalStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EternalStorage *EternalStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EternalStorage.Contract.contract.Transact(opts, method, params...)
}

// GetAddressValue is a free data retrieval call binding the contract method 0x4c77e5ba.
//
// Solidity: function getAddressValue(record bytes32) constant returns(address)
func (_EternalStorage *EternalStorageCaller) GetAddressValue(opts *bind.CallOpts, record [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _EternalStorage.contract.Call(opts, out, "getAddressValue", record)
	return *ret0, err
}

// GetAddressValue is a free data retrieval call binding the contract method 0x4c77e5ba.
//
// Solidity: function getAddressValue(record bytes32) constant returns(address)
func (_EternalStorage *EternalStorageSession) GetAddressValue(record [32]byte) (common.Address, error) {
	return _EternalStorage.Contract.GetAddressValue(&_EternalStorage.CallOpts, record)
}

// GetAddressValue is a free data retrieval call binding the contract method 0x4c77e5ba.
//
// Solidity: function getAddressValue(record bytes32) constant returns(address)
func (_EternalStorage *EternalStorageCallerSession) GetAddressValue(record [32]byte) (common.Address, error) {
	return _EternalStorage.Contract.GetAddressValue(&_EternalStorage.CallOpts, record)
}

// GetBooleanValue is a free data retrieval call binding the contract method 0x17e7dd22.
//
// Solidity: function getBooleanValue(record bytes32) constant returns(bool)
func (_EternalStorage *EternalStorageCaller) GetBooleanValue(opts *bind.CallOpts, record [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _EternalStorage.contract.Call(opts, out, "getBooleanValue", record)
	return *ret0, err
}

// GetBooleanValue is a free data retrieval call binding the contract method 0x17e7dd22.
//
// Solidity: function getBooleanValue(record bytes32) constant returns(bool)
func (_EternalStorage *EternalStorageSession) GetBooleanValue(record [32]byte) (bool, error) {
	return _EternalStorage.Contract.GetBooleanValue(&_EternalStorage.CallOpts, record)
}

// GetBooleanValue is a free data retrieval call binding the contract method 0x17e7dd22.
//
// Solidity: function getBooleanValue(record bytes32) constant returns(bool)
func (_EternalStorage *EternalStorageCallerSession) GetBooleanValue(record [32]byte) (bool, error) {
	return _EternalStorage.Contract.GetBooleanValue(&_EternalStorage.CallOpts, record)
}

// GetBytes32Value is a free data retrieval call binding the contract method 0x025ec81a.
//
// Solidity: function getBytes32Value(record bytes32) constant returns(bytes32)
func (_EternalStorage *EternalStorageCaller) GetBytes32Value(opts *bind.CallOpts, record [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _EternalStorage.contract.Call(opts, out, "getBytes32Value", record)
	return *ret0, err
}

// GetBytes32Value is a free data retrieval call binding the contract method 0x025ec81a.
//
// Solidity: function getBytes32Value(record bytes32) constant returns(bytes32)
func (_EternalStorage *EternalStorageSession) GetBytes32Value(record [32]byte) ([32]byte, error) {
	return _EternalStorage.Contract.GetBytes32Value(&_EternalStorage.CallOpts, record)
}

// GetBytes32Value is a free data retrieval call binding the contract method 0x025ec81a.
//
// Solidity: function getBytes32Value(record bytes32) constant returns(bytes32)
func (_EternalStorage *EternalStorageCallerSession) GetBytes32Value(record [32]byte) ([32]byte, error) {
	return _EternalStorage.Contract.GetBytes32Value(&_EternalStorage.CallOpts, record)
}

// GetBytesValue is a free data retrieval call binding the contract method 0x44bfa56e.
//
// Solidity: function getBytesValue(record bytes32) constant returns(bytes)
func (_EternalStorage *EternalStorageCaller) GetBytesValue(opts *bind.CallOpts, record [32]byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _EternalStorage.contract.Call(opts, out, "getBytesValue", record)
	return *ret0, err
}

// GetBytesValue is a free data retrieval call binding the contract method 0x44bfa56e.
//
// Solidity: function getBytesValue(record bytes32) constant returns(bytes)
func (_EternalStorage *EternalStorageSession) GetBytesValue(record [32]byte) ([]byte, error) {
	return _EternalStorage.Contract.GetBytesValue(&_EternalStorage.CallOpts, record)
}

// GetBytesValue is a free data retrieval call binding the contract method 0x44bfa56e.
//
// Solidity: function getBytesValue(record bytes32) constant returns(bytes)
func (_EternalStorage *EternalStorageCallerSession) GetBytesValue(record [32]byte) ([]byte, error) {
	return _EternalStorage.Contract.GetBytesValue(&_EternalStorage.CallOpts, record)
}

// GetIntValue is a free data retrieval call binding the contract method 0x9007127b.
//
// Solidity: function getIntValue(record bytes32) constant returns(int256)
func (_EternalStorage *EternalStorageCaller) GetIntValue(opts *bind.CallOpts, record [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EternalStorage.contract.Call(opts, out, "getIntValue", record)
	return *ret0, err
}

// GetIntValue is a free data retrieval call binding the contract method 0x9007127b.
//
// Solidity: function getIntValue(record bytes32) constant returns(int256)
func (_EternalStorage *EternalStorageSession) GetIntValue(record [32]byte) (*big.Int, error) {
	return _EternalStorage.Contract.GetIntValue(&_EternalStorage.CallOpts, record)
}

// GetIntValue is a free data retrieval call binding the contract method 0x9007127b.
//
// Solidity: function getIntValue(record bytes32) constant returns(int256)
func (_EternalStorage *EternalStorageCallerSession) GetIntValue(record [32]byte) (*big.Int, error) {
	return _EternalStorage.Contract.GetIntValue(&_EternalStorage.CallOpts, record)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_EternalStorage *EternalStorageCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _EternalStorage.contract.Call(opts, out, "getOwner")
	return *ret0, err
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_EternalStorage *EternalStorageSession) GetOwner() (common.Address, error) {
	return _EternalStorage.Contract.GetOwner(&_EternalStorage.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_EternalStorage *EternalStorageCallerSession) GetOwner() (common.Address, error) {
	return _EternalStorage.Contract.GetOwner(&_EternalStorage.CallOpts)
}

// GetStringValue is a free data retrieval call binding the contract method 0xa209a29c.
//
// Solidity: function getStringValue(record bytes32) constant returns(string)
func (_EternalStorage *EternalStorageCaller) GetStringValue(opts *bind.CallOpts, record [32]byte) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _EternalStorage.contract.Call(opts, out, "getStringValue", record)
	return *ret0, err
}

// GetStringValue is a free data retrieval call binding the contract method 0xa209a29c.
//
// Solidity: function getStringValue(record bytes32) constant returns(string)
func (_EternalStorage *EternalStorageSession) GetStringValue(record [32]byte) (string, error) {
	return _EternalStorage.Contract.GetStringValue(&_EternalStorage.CallOpts, record)
}

// GetStringValue is a free data retrieval call binding the contract method 0xa209a29c.
//
// Solidity: function getStringValue(record bytes32) constant returns(string)
func (_EternalStorage *EternalStorageCallerSession) GetStringValue(record [32]byte) (string, error) {
	return _EternalStorage.Contract.GetStringValue(&_EternalStorage.CallOpts, record)
}

// GetUIntValue is a free data retrieval call binding the contract method 0xbdc963d8.
//
// Solidity: function getUIntValue(record bytes32) constant returns(uint256)
func (_EternalStorage *EternalStorageCaller) GetUIntValue(opts *bind.CallOpts, record [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EternalStorage.contract.Call(opts, out, "getUIntValue", record)
	return *ret0, err
}

// GetUIntValue is a free data retrieval call binding the contract method 0xbdc963d8.
//
// Solidity: function getUIntValue(record bytes32) constant returns(uint256)
func (_EternalStorage *EternalStorageSession) GetUIntValue(record [32]byte) (*big.Int, error) {
	return _EternalStorage.Contract.GetUIntValue(&_EternalStorage.CallOpts, record)
}

// GetUIntValue is a free data retrieval call binding the contract method 0xbdc963d8.
//
// Solidity: function getUIntValue(record bytes32) constant returns(uint256)
func (_EternalStorage *EternalStorageCallerSession) GetUIntValue(record [32]byte) (*big.Int, error) {
	return _EternalStorage.Contract.GetUIntValue(&_EternalStorage.CallOpts, record)
}

// DeleteAddressValue is a paid mutator transaction binding the contract method 0x043106c0.
//
// Solidity: function deleteAddressValue(record bytes32) returns()
func (_EternalStorage *EternalStorageTransactor) DeleteAddressValue(opts *bind.TransactOpts, record [32]byte) (*types.Transaction, error) {
	return _EternalStorage.contract.Transact(opts, "deleteAddressValue", record)
}

// DeleteAddressValue is a paid mutator transaction binding the contract method 0x043106c0.
//
// Solidity: function deleteAddressValue(record bytes32) returns()
func (_EternalStorage *EternalStorageSession) DeleteAddressValue(record [32]byte) (*types.Transaction, error) {
	return _EternalStorage.Contract.DeleteAddressValue(&_EternalStorage.TransactOpts, record)
}

// DeleteAddressValue is a paid mutator transaction binding the contract method 0x043106c0.
//
// Solidity: function deleteAddressValue(record bytes32) returns()
func (_EternalStorage *EternalStorageTransactorSession) DeleteAddressValue(record [32]byte) (*types.Transaction, error) {
	return _EternalStorage.Contract.DeleteAddressValue(&_EternalStorage.TransactOpts, record)
}

// DeleteBooleanValue is a paid mutator transaction binding the contract method 0x3cc1635c.
//
// Solidity: function deleteBooleanValue(record bytes32) returns()
func (_EternalStorage *EternalStorageTransactor) DeleteBooleanValue(opts *bind.TransactOpts, record [32]byte) (*types.Transaction, error) {
	return _EternalStorage.contract.Transact(opts, "deleteBooleanValue", record)
}

// DeleteBooleanValue is a paid mutator transaction binding the contract method 0x3cc1635c.
//
// Solidity: function deleteBooleanValue(record bytes32) returns()
func (_EternalStorage *EternalStorageSession) DeleteBooleanValue(record [32]byte) (*types.Transaction, error) {
	return _EternalStorage.Contract.DeleteBooleanValue(&_EternalStorage.TransactOpts, record)
}

// DeleteBooleanValue is a paid mutator transaction binding the contract method 0x3cc1635c.
//
// Solidity: function deleteBooleanValue(record bytes32) returns()
func (_EternalStorage *EternalStorageTransactorSession) DeleteBooleanValue(record [32]byte) (*types.Transaction, error) {
	return _EternalStorage.Contract.DeleteBooleanValue(&_EternalStorage.TransactOpts, record)
}

// DeleteBytes32Value is a paid mutator transaction binding the contract method 0x124f2418.
//
// Solidity: function deleteBytes32Value(record bytes32) returns()
func (_EternalStorage *EternalStorageTransactor) DeleteBytes32Value(opts *bind.TransactOpts, record [32]byte) (*types.Transaction, error) {
	return _EternalStorage.contract.Transact(opts, "deleteBytes32Value", record)
}

// DeleteBytes32Value is a paid mutator transaction binding the contract method 0x124f2418.
//
// Solidity: function deleteBytes32Value(record bytes32) returns()
func (_EternalStorage *EternalStorageSession) DeleteBytes32Value(record [32]byte) (*types.Transaction, error) {
	return _EternalStorage.Contract.DeleteBytes32Value(&_EternalStorage.TransactOpts, record)
}

// DeleteBytes32Value is a paid mutator transaction binding the contract method 0x124f2418.
//
// Solidity: function deleteBytes32Value(record bytes32) returns()
func (_EternalStorage *EternalStorageTransactorSession) DeleteBytes32Value(record [32]byte) (*types.Transaction, error) {
	return _EternalStorage.Contract.DeleteBytes32Value(&_EternalStorage.TransactOpts, record)
}

// DeleteBytesValue is a paid mutator transaction binding the contract method 0x0c55d925.
//
// Solidity: function deleteBytesValue(record bytes32) returns()
func (_EternalStorage *EternalStorageTransactor) DeleteBytesValue(opts *bind.TransactOpts, record [32]byte) (*types.Transaction, error) {
	return _EternalStorage.contract.Transact(opts, "deleteBytesValue", record)
}

// DeleteBytesValue is a paid mutator transaction binding the contract method 0x0c55d925.
//
// Solidity: function deleteBytesValue(record bytes32) returns()
func (_EternalStorage *EternalStorageSession) DeleteBytesValue(record [32]byte) (*types.Transaction, error) {
	return _EternalStorage.Contract.DeleteBytesValue(&_EternalStorage.TransactOpts, record)
}

// DeleteBytesValue is a paid mutator transaction binding the contract method 0x0c55d925.
//
// Solidity: function deleteBytesValue(record bytes32) returns()
func (_EternalStorage *EternalStorageTransactorSession) DeleteBytesValue(record [32]byte) (*types.Transaction, error) {
	return _EternalStorage.Contract.DeleteBytesValue(&_EternalStorage.TransactOpts, record)
}

// DeleteIntValue is a paid mutator transaction binding the contract method 0x8267a9ee.
//
// Solidity: function deleteIntValue(record bytes32) returns()
func (_EternalStorage *EternalStorageTransactor) DeleteIntValue(opts *bind.TransactOpts, record [32]byte) (*types.Transaction, error) {
	return _EternalStorage.contract.Transact(opts, "deleteIntValue", record)
}

// DeleteIntValue is a paid mutator transaction binding the contract method 0x8267a9ee.
//
// Solidity: function deleteIntValue(record bytes32) returns()
func (_EternalStorage *EternalStorageSession) DeleteIntValue(record [32]byte) (*types.Transaction, error) {
	return _EternalStorage.Contract.DeleteIntValue(&_EternalStorage.TransactOpts, record)
}

// DeleteIntValue is a paid mutator transaction binding the contract method 0x8267a9ee.
//
// Solidity: function deleteIntValue(record bytes32) returns()
func (_EternalStorage *EternalStorageTransactorSession) DeleteIntValue(record [32]byte) (*types.Transaction, error) {
	return _EternalStorage.Contract.DeleteIntValue(&_EternalStorage.TransactOpts, record)
}

// DeleteStringValue is a paid mutator transaction binding the contract method 0xba69fcaa.
//
// Solidity: function deleteStringValue(record bytes32) returns()
func (_EternalStorage *EternalStorageTransactor) DeleteStringValue(opts *bind.TransactOpts, record [32]byte) (*types.Transaction, error) {
	return _EternalStorage.contract.Transact(opts, "deleteStringValue", record)
}

// DeleteStringValue is a paid mutator transaction binding the contract method 0xba69fcaa.
//
// Solidity: function deleteStringValue(record bytes32) returns()
func (_EternalStorage *EternalStorageSession) DeleteStringValue(record [32]byte) (*types.Transaction, error) {
	return _EternalStorage.Contract.DeleteStringValue(&_EternalStorage.TransactOpts, record)
}

// DeleteStringValue is a paid mutator transaction binding the contract method 0xba69fcaa.
//
// Solidity: function deleteStringValue(record bytes32) returns()
func (_EternalStorage *EternalStorageTransactorSession) DeleteStringValue(record [32]byte) (*types.Transaction, error) {
	return _EternalStorage.Contract.DeleteStringValue(&_EternalStorage.TransactOpts, record)
}

// DeleteUIntValue is a paid mutator transaction binding the contract method 0x93fe4248.
//
// Solidity: function deleteUIntValue(record bytes32) returns()
func (_EternalStorage *EternalStorageTransactor) DeleteUIntValue(opts *bind.TransactOpts, record [32]byte) (*types.Transaction, error) {
	return _EternalStorage.contract.Transact(opts, "deleteUIntValue", record)
}

// DeleteUIntValue is a paid mutator transaction binding the contract method 0x93fe4248.
//
// Solidity: function deleteUIntValue(record bytes32) returns()
func (_EternalStorage *EternalStorageSession) DeleteUIntValue(record [32]byte) (*types.Transaction, error) {
	return _EternalStorage.Contract.DeleteUIntValue(&_EternalStorage.TransactOpts, record)
}

// DeleteUIntValue is a paid mutator transaction binding the contract method 0x93fe4248.
//
// Solidity: function deleteUIntValue(record bytes32) returns()
func (_EternalStorage *EternalStorageTransactorSession) DeleteUIntValue(record [32]byte) (*types.Transaction, error) {
	return _EternalStorage.Contract.DeleteUIntValue(&_EternalStorage.TransactOpts, record)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_EternalStorage *EternalStorageTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EternalStorage.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_EternalStorage *EternalStorageSession) Kill() (*types.Transaction, error) {
	return _EternalStorage.Contract.Kill(&_EternalStorage.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_EternalStorage *EternalStorageTransactorSession) Kill() (*types.Transaction, error) {
	return _EternalStorage.Contract.Kill(&_EternalStorage.TransactOpts)
}

// PassOwnership is a paid mutator transaction binding the contract method 0x620996df.
//
// Solidity: function passOwnership(_newOwner address) returns()
func (_EternalStorage *EternalStorageTransactor) PassOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _EternalStorage.contract.Transact(opts, "passOwnership", _newOwner)
}

// PassOwnership is a paid mutator transaction binding the contract method 0x620996df.
//
// Solidity: function passOwnership(_newOwner address) returns()
func (_EternalStorage *EternalStorageSession) PassOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _EternalStorage.Contract.PassOwnership(&_EternalStorage.TransactOpts, _newOwner)
}

// PassOwnership is a paid mutator transaction binding the contract method 0x620996df.
//
// Solidity: function passOwnership(_newOwner address) returns()
func (_EternalStorage *EternalStorageTransactorSession) PassOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _EternalStorage.Contract.PassOwnership(&_EternalStorage.TransactOpts, _newOwner)
}

// SetAddressValue is a paid mutator transaction binding the contract method 0x5a2bf25a.
//
// Solidity: function setAddressValue(record bytes32, value address) returns()
func (_EternalStorage *EternalStorageTransactor) SetAddressValue(opts *bind.TransactOpts, record [32]byte, value common.Address) (*types.Transaction, error) {
	return _EternalStorage.contract.Transact(opts, "setAddressValue", record, value)
}

// SetAddressValue is a paid mutator transaction binding the contract method 0x5a2bf25a.
//
// Solidity: function setAddressValue(record bytes32, value address) returns()
func (_EternalStorage *EternalStorageSession) SetAddressValue(record [32]byte, value common.Address) (*types.Transaction, error) {
	return _EternalStorage.Contract.SetAddressValue(&_EternalStorage.TransactOpts, record, value)
}

// SetAddressValue is a paid mutator transaction binding the contract method 0x5a2bf25a.
//
// Solidity: function setAddressValue(record bytes32, value address) returns()
func (_EternalStorage *EternalStorageTransactorSession) SetAddressValue(record [32]byte, value common.Address) (*types.Transaction, error) {
	return _EternalStorage.Contract.SetAddressValue(&_EternalStorage.TransactOpts, record, value)
}

// SetBooleanValue is a paid mutator transaction binding the contract method 0x3eba9ed2.
//
// Solidity: function setBooleanValue(record bytes32, value bool) returns()
func (_EternalStorage *EternalStorageTransactor) SetBooleanValue(opts *bind.TransactOpts, record [32]byte, value bool) (*types.Transaction, error) {
	return _EternalStorage.contract.Transact(opts, "setBooleanValue", record, value)
}

// SetBooleanValue is a paid mutator transaction binding the contract method 0x3eba9ed2.
//
// Solidity: function setBooleanValue(record bytes32, value bool) returns()
func (_EternalStorage *EternalStorageSession) SetBooleanValue(record [32]byte, value bool) (*types.Transaction, error) {
	return _EternalStorage.Contract.SetBooleanValue(&_EternalStorage.TransactOpts, record, value)
}

// SetBooleanValue is a paid mutator transaction binding the contract method 0x3eba9ed2.
//
// Solidity: function setBooleanValue(record bytes32, value bool) returns()
func (_EternalStorage *EternalStorageTransactorSession) SetBooleanValue(record [32]byte, value bool) (*types.Transaction, error) {
	return _EternalStorage.Contract.SetBooleanValue(&_EternalStorage.TransactOpts, record, value)
}

// SetBytes32Value is a paid mutator transaction binding the contract method 0x25cf512d.
//
// Solidity: function setBytes32Value(record bytes32, value bytes32) returns()
func (_EternalStorage *EternalStorageTransactor) SetBytes32Value(opts *bind.TransactOpts, record [32]byte, value [32]byte) (*types.Transaction, error) {
	return _EternalStorage.contract.Transact(opts, "setBytes32Value", record, value)
}

// SetBytes32Value is a paid mutator transaction binding the contract method 0x25cf512d.
//
// Solidity: function setBytes32Value(record bytes32, value bytes32) returns()
func (_EternalStorage *EternalStorageSession) SetBytes32Value(record [32]byte, value [32]byte) (*types.Transaction, error) {
	return _EternalStorage.Contract.SetBytes32Value(&_EternalStorage.TransactOpts, record, value)
}

// SetBytes32Value is a paid mutator transaction binding the contract method 0x25cf512d.
//
// Solidity: function setBytes32Value(record bytes32, value bytes32) returns()
func (_EternalStorage *EternalStorageTransactorSession) SetBytes32Value(record [32]byte, value [32]byte) (*types.Transaction, error) {
	return _EternalStorage.Contract.SetBytes32Value(&_EternalStorage.TransactOpts, record, value)
}

// SetBytesValue is a paid mutator transaction binding the contract method 0xc9a52d2c.
//
// Solidity: function setBytesValue(record bytes32, value bytes) returns()
func (_EternalStorage *EternalStorageTransactor) SetBytesValue(opts *bind.TransactOpts, record [32]byte, value []byte) (*types.Transaction, error) {
	return _EternalStorage.contract.Transact(opts, "setBytesValue", record, value)
}

// SetBytesValue is a paid mutator transaction binding the contract method 0xc9a52d2c.
//
// Solidity: function setBytesValue(record bytes32, value bytes) returns()
func (_EternalStorage *EternalStorageSession) SetBytesValue(record [32]byte, value []byte) (*types.Transaction, error) {
	return _EternalStorage.Contract.SetBytesValue(&_EternalStorage.TransactOpts, record, value)
}

// SetBytesValue is a paid mutator transaction binding the contract method 0xc9a52d2c.
//
// Solidity: function setBytesValue(record bytes32, value bytes) returns()
func (_EternalStorage *EternalStorageTransactorSession) SetBytesValue(record [32]byte, value []byte) (*types.Transaction, error) {
	return _EternalStorage.Contract.SetBytesValue(&_EternalStorage.TransactOpts, record, value)
}

// SetIntValue is a paid mutator transaction binding the contract method 0xa77aa49e.
//
// Solidity: function setIntValue(record bytes32, value int256) returns()
func (_EternalStorage *EternalStorageTransactor) SetIntValue(opts *bind.TransactOpts, record [32]byte, value *big.Int) (*types.Transaction, error) {
	return _EternalStorage.contract.Transact(opts, "setIntValue", record, value)
}

// SetIntValue is a paid mutator transaction binding the contract method 0xa77aa49e.
//
// Solidity: function setIntValue(record bytes32, value int256) returns()
func (_EternalStorage *EternalStorageSession) SetIntValue(record [32]byte, value *big.Int) (*types.Transaction, error) {
	return _EternalStorage.Contract.SetIntValue(&_EternalStorage.TransactOpts, record, value)
}

// SetIntValue is a paid mutator transaction binding the contract method 0xa77aa49e.
//
// Solidity: function setIntValue(record bytes32, value int256) returns()
func (_EternalStorage *EternalStorageTransactorSession) SetIntValue(record [32]byte, value *big.Int) (*types.Transaction, error) {
	return _EternalStorage.Contract.SetIntValue(&_EternalStorage.TransactOpts, record, value)
}

// SetStringValue is a paid mutator transaction binding the contract method 0xf5866066.
//
// Solidity: function setStringValue(record bytes32, value string) returns()
func (_EternalStorage *EternalStorageTransactor) SetStringValue(opts *bind.TransactOpts, record [32]byte, value string) (*types.Transaction, error) {
	return _EternalStorage.contract.Transact(opts, "setStringValue", record, value)
}

// SetStringValue is a paid mutator transaction binding the contract method 0xf5866066.
//
// Solidity: function setStringValue(record bytes32, value string) returns()
func (_EternalStorage *EternalStorageSession) SetStringValue(record [32]byte, value string) (*types.Transaction, error) {
	return _EternalStorage.Contract.SetStringValue(&_EternalStorage.TransactOpts, record, value)
}

// SetStringValue is a paid mutator transaction binding the contract method 0xf5866066.
//
// Solidity: function setStringValue(record bytes32, value string) returns()
func (_EternalStorage *EternalStorageTransactorSession) SetStringValue(record [32]byte, value string) (*types.Transaction, error) {
	return _EternalStorage.Contract.SetStringValue(&_EternalStorage.TransactOpts, record, value)
}

// SetUIntValue is a paid mutator transaction binding the contract method 0x3562fd20.
//
// Solidity: function setUIntValue(record bytes32, value uint256) returns()
func (_EternalStorage *EternalStorageTransactor) SetUIntValue(opts *bind.TransactOpts, record [32]byte, value *big.Int) (*types.Transaction, error) {
	return _EternalStorage.contract.Transact(opts, "setUIntValue", record, value)
}

// SetUIntValue is a paid mutator transaction binding the contract method 0x3562fd20.
//
// Solidity: function setUIntValue(record bytes32, value uint256) returns()
func (_EternalStorage *EternalStorageSession) SetUIntValue(record [32]byte, value *big.Int) (*types.Transaction, error) {
	return _EternalStorage.Contract.SetUIntValue(&_EternalStorage.TransactOpts, record, value)
}

// SetUIntValue is a paid mutator transaction binding the contract method 0x3562fd20.
//
// Solidity: function setUIntValue(record bytes32, value uint256) returns()
func (_EternalStorage *EternalStorageTransactorSession) SetUIntValue(record [32]byte, value *big.Int) (*types.Transaction, error) {
	return _EternalStorage.Contract.SetUIntValue(&_EternalStorage.TransactOpts, record, value)
}

// OwnableABI is the input ABI used to generate the binding from.
const OwnableABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"passOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// OwnableBin is the compiled bytecode used for deploying new contracts.
const OwnableBin = `0x6060604052341561000f57600080fd5b60008054600160a060020a033316600160a060020a03199091161790556101818061003b6000396000f3006060604052600436106100565763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166341c0e1b5811461005b578063620996df14610070578063893d20e81461008f575b600080fd5b341561006657600080fd5b61006e6100be565b005b341561007b57600080fd5b61006e600160a060020a03600435166100e7565b341561009a57600080fd5b6100a2610146565b604051600160a060020a03909116815260200160405180910390f35b60005433600160a060020a039081169116146100d957600080fd5b600054600160a060020a0316ff5b60005433600160a060020a0390811691161461010257600080fd5b600160a060020a038116151561011757600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600054600160a060020a0316905600a165627a7a723058205bb986ccf846572e0a7006270acec7bc9722ef2e4e5c77ffd6d95bc17c6bbed80029`

// DeployOwnable deploys a new Ethereum contract, binding an instance of Ownable to it.
func DeployOwnable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Ownable, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OwnableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}}, nil
}

// Ownable is an auto generated Go binding around an Ethereum contract.
type Ownable struct {
	OwnableCaller     // Read-only binding to the contract
	OwnableTransactor // Write-only binding to the contract
}

// OwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableSession struct {
	Contract     *Ownable          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableCallerSession struct {
	Contract *OwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableTransactorSession struct {
	Contract     *OwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableRaw struct {
	Contract *Ownable // Generic contract binding to access the raw methods on
}

// OwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableCallerRaw struct {
	Contract *OwnableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableTransactorRaw struct {
	Contract *OwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnable creates a new instance of Ownable, bound to a specific deployed contract.
func NewOwnable(address common.Address, backend bind.ContractBackend) (*Ownable, error) {
	contract, err := bindOwnable(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}}, nil
}

// NewOwnableCaller creates a new read-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableCaller(address common.Address, caller bind.ContractCaller) (*OwnableCaller, error) {
	contract, err := bindOwnable(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableCaller{contract: contract}, nil
}

// NewOwnableTransactor creates a new write-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableTransactor, error) {
	contract, err := bindOwnable(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &OwnableTransactor{contract: contract}, nil
}

// bindOwnable binds a generic wrapper to an already deployed contract.
func bindOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.OwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transact(opts, method, params...)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_Ownable *OwnableCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Ownable.contract.Call(opts, out, "getOwner")
	return *ret0, err
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_Ownable *OwnableSession) GetOwner() (common.Address, error) {
	return _Ownable.Contract.GetOwner(&_Ownable.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_Ownable *OwnableCallerSession) GetOwner() (common.Address, error) {
	return _Ownable.Contract.GetOwner(&_Ownable.CallOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Ownable *OwnableTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Ownable *OwnableSession) Kill() (*types.Transaction, error) {
	return _Ownable.Contract.Kill(&_Ownable.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Ownable *OwnableTransactorSession) Kill() (*types.Transaction, error) {
	return _Ownable.Contract.Kill(&_Ownable.TransactOpts)
}

// PassOwnership is a paid mutator transaction binding the contract method 0x620996df.
//
// Solidity: function passOwnership(_newOwner address) returns()
func (_Ownable *OwnableTransactor) PassOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "passOwnership", _newOwner)
}

// PassOwnership is a paid mutator transaction binding the contract method 0x620996df.
//
// Solidity: function passOwnership(_newOwner address) returns()
func (_Ownable *OwnableSession) PassOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.PassOwnership(&_Ownable.TransactOpts, _newOwner)
}

// PassOwnership is a paid mutator transaction binding the contract method 0x620996df.
//
// Solidity: function passOwnership(_newOwner address) returns()
func (_Ownable *OwnableTransactorSession) PassOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.PassOwnership(&_Ownable.TransactOpts, _newOwner)
}

// UsersLibABI is the input ABI used to generate the binding from.
const UsersLibABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_storageContract\",\"type\":\"address\"},{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_imageUrl\",\"type\":\"bytes32\"}],\"name\":\"updateProfilePicture\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_storageContract\",\"type\":\"address\"},{\"name\":\"_name\",\"type\":\"bytes32\"}],\"name\":\"getUserDetailsByName\",\"outputs\":[{\"name\":\"found\",\"type\":\"bool\"},{\"name\":\"userID\",\"type\":\"uint256\"},{\"name\":\"userAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_storageContract\",\"type\":\"address\"},{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_name\",\"type\":\"bytes32\"}],\"name\":\"updateName\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_storageContract\",\"type\":\"address\"},{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_name\",\"type\":\"bytes32\"}],\"name\":\"addUser\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_storageContract\",\"type\":\"address\"},{\"name\":\"_userID\",\"type\":\"uint256\"}],\"name\":\"getUserDetailsByID\",\"outputs\":[{\"name\":\"found\",\"type\":\"bool\"},{\"name\":\"userAddress\",\"type\":\"address\"},{\"name\":\"userName\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_storageContract\",\"type\":\"address\"}],\"name\":\"getUsersCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_storageContract\",\"type\":\"address\"},{\"name\":\"_name\",\"type\":\"bytes32\"}],\"name\":\"getIDByName\",\"outputs\":[{\"name\":\"found\",\"type\":\"bool\"},{\"name\":\"id\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_storageContract\",\"type\":\"address\"},{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_newAddress\",\"type\":\"address\"}],\"name\":\"updateAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_storageContract\",\"type\":\"address\"},{\"name\":\"idx\",\"type\":\"uint256\"}],\"name\":\"removeUser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_storageContract\",\"type\":\"address\"},{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"getIDByAddress\",\"outputs\":[{\"name\":\"found\",\"type\":\"bool\"},{\"name\":\"id\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_storageContract\",\"type\":\"address\"}],\"name\":\"getActiveUsersCount\",\"outputs\":[{\"name\":\"activeUsersCount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_storageContract\",\"type\":\"address\"},{\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"isActive\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"UserAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"UserRemoved\",\"type\":\"event\"}]"

// UsersLibBin is the compiled bytecode used for deploying new contracts.
const UsersLibBin = `0x6060604052341561000f57600080fd5b6114078061001e6000396000f3006060604052600436106100a05763ffffffff60e060020a6000350416630eb1f41181146100a5578063127de333146100d65780631fa768db1461011c578063364e4ddb14610139578063656a19611461015657806378f082681461019b5780637c6e7d45146101c1578063868562b7146101f257806393d3f77414610212578063a838e4511461022b578063c817e2df14610245578063d266e83b14610259575b600080fd5b6100c2600160a060020a0360043581169060243516604435610270565b604051901515815260200160405180910390f35b6100ed600160a060020a036004351660243561032c565b60405192151583526020830191909152600160a060020a03166040808301919091526060909101905180910390f35b6100c2600160a060020a03600435811690602435166044356103e6565b6100c2600160a060020a0360043581169060243516604435610483565b61016d600160a060020a0360043516602435610743565b6040519215158352600160a060020a0390911660208301526040808301919091526060909101905180910390f35b6101af600160a060020a03600435166108ee565b60405190815260200160405180910390f35b6101d8600160a060020a036004351660243561098f565b604051911515825260208201526040908101905180910390f35b6100c2600160a060020a0360043581169060243581169060443516610b24565b610229600160a060020a0360043516602435610bb9565b005b6101d8600160a060020a0360043581169060243516610c7a565b6101af600160a060020a0360043516610e32565b6100c2600160a060020a0360043516602435610eff565b600080600061027f8686610c7a565b9150915081156103235785600160a060020a03166325cf512d826040517f757365725f696d675f6c6f676f000000000000000000000000000000000000008152600d810191909152602d0160405180910390208660405160e060020a63ffffffff851602815260048101929092526024820152604401600060405180830381600087803b151561030e57600080fd5b6102c65a03f1151561031f57600080fd5b5050505b50949350505050565b600080600061033b858561098f565b909350915082151561034c576103df565b84600160a060020a0316634c77e5ba8360405160008051602061139c8339815191528152600c810191909152602c01604051809103902060006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b15156103c257600080fd5b6102c65a03f115156103d357600080fd5b50505060405180519150505b9250925092565b60008060006103f58686610c7a565b9150915081801561040b575061040b8685610f9a565b156103235785600160a060020a03166325cf512d826040516000805160206113bc8339815191528152600981019190915260290160405180910390208660405160e060020a63ffffffff851602815260048101929092526024820152604401600060405180830381600087803b151561030e57600080fd5b6000806104908584610f9a565b151561049b57600080fd5b6104a58585611127565b15156104b057600080fd5b6104bb8560006112db565b905084600160a060020a03166325cf512d826040516000805160206113bc8339815191528152600981019190915260290160405180910390208560405160e060020a63ffffffff851602815260048101929092526024820152604401600060405180830381600087803b151561053057600080fd5b6102c65a03f1151561054157600080fd5b50505084600160a060020a0316635a2bf25a8260405160008051602061139c8339815191528152600c810191909152602c0160405180910390208660405160e060020a63ffffffff85160281526004810192909252600160a060020a03166024820152604401600060405180830381600087803b15156105c057600080fd5b6102c65a03f115156105d157600080fd5b50505084600160a060020a0316633eba9ed28260405160008051602061137c8339815191528152600f810191909152602f016040518091039020600160405160e060020a63ffffffff8516028152600481019290925215156024820152604401600060405180830381600087803b151561064a57600080fd5b6102c65a03f1151561065b57600080fd5b505050610667856108ee565b8160010111156107055784600160a060020a0316633562fd206040517f75736572735f636f756e740000000000000000000000000000000000000000008152600b0160405180910390208360010160405160e060020a63ffffffff851602815260048101929092526024820152604401600060405180830381600087803b15156106f057600080fd5b6102c65a03f1151561070157600080fd5b5050505b7f27412bb19840b11057dee7c7019767312555c4497d8d4c2383fa6b7791bf014b8160405190815260200160405180910390a1506001949350505050565b600080600084600160a060020a03166317e7dd228560405160008051602061137c8339815191528152600f810191909152602f01604051809103902060006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b15156107be57600080fd5b6102c65a03f115156107cf57600080fd5b50505060405180519350508215156107e6576103df565b84600160a060020a0316634c77e5ba8560405160008051602061139c8339815191528152600c810191909152602c01604051809103902060006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b151561085c57600080fd5b6102c65a03f1151561086d57600080fd5b5050506040518051925050600160a060020a03851663025ec81a856040516000805160206113bc83398151915281526009810191909152602901604051809103902060006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b15156103c257600080fd5b600081600160a060020a031663bdc963d86040517f75736572735f636f756e740000000000000000000000000000000000000000008152600b01604051809103902060006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b151561096f57600080fd5b6102c65a03f1151561098057600080fd5b50505060405180519392505050565b60008060008061099e866108ee565b9150600090505b81811015610b135785600160a060020a03166317e7dd228260405160008051602061137c8339815191528152600f810191909152602f01604051809103902060006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b1515610a2357600080fd5b6102c65a03f11515610a3457600080fd5b5050506040518051905015610b0b5784604051908152602001604051908190039020600160a060020a03871663025ec81a836040516000805160206113bc83398151915281526009810191909152602901604051809103902060006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b1515610acc57600080fd5b6102c65a03f11515610add57600080fd5b505050604051805190506040519081526020016040519081900390201415610b0b5760018193509350610b1b565b6001016109a5565b600093508392505b50509250929050565b6000806000610b338686610c7a565b9150915081156103235785600160a060020a0316635a2bf25a8260405160008051602061139c8339815191528152600c810191909152602c0160405180910390208660405160e060020a63ffffffff85160281526004810192909252600160a060020a03166024820152604401600060405180830381600087803b151561030e57600080fd5b81600160a060020a0316633eba9ed28260405160008051602061137c8339815191528152600f810191909152602f016040518091039020600060405160e060020a63ffffffff8516028152600481019290925215156024820152604401600060405180830381600087803b1515610c2f57600080fd5b6102c65a03f11515610c4057600080fd5b5050507fab24b62fd8ab41107c600afaa5e4d6e14f64f253a96e231e221809b6cc7c46a68160405190815260200160405180910390a15050565b600080600080610c89866108ee565b9150600090505b81811015610b135785600160a060020a03166317e7dd228260405160008051602061137c8339815191528152600f810191909152602f01604051809103902060006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b1515610d0e57600080fd5b6102c65a03f11515610d1f57600080fd5b5050506040518051905015610e2a5784604051600160a060020a03919091166c01000000000000000000000000028152601401604051908190039020600160a060020a038716634c77e5ba8360405160008051602061139c8339815191528152600c810191909152602c01604051809103902060006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b1515610dd157600080fd5b6102c65a03f11515610de257600080fd5b50505060405180519050604051600160a060020a03919091166c010000000000000000000000000281526014016040519081900390201415610e2a5760018193509350610b1b565b600101610c90565b6000806000610e40846108ee565b915060009250600090505b81811015610ef85783600160a060020a03166317e7dd228260405160008051602061137c8339815191528152600f810191909152602f01604051809103902060006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b1515610ec957600080fd5b6102c65a03f11515610eda57600080fd5b5050506040518051905015610ef0576001909201915b600101610e4b565b5050919050565b600082600160a060020a03166317e7dd228360405160008051602061137c8339815191528152600f810191909152602f01604051809103902060006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b1515610f7757600080fd5b6102c65a03f11515610f8857600080fd5b50505060405180519150505b92915050565b6000806000610fa8856108ee565b9150600090505b8181101561111a5784600160a060020a03166317e7dd228260405160008051602061137c8339815191528152600f810191909152602f01604051809103902060006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b151561102d57600080fd5b6102c65a03f1151561103e57600080fd5b50505060405180519050156111125783604051908152602001604051908190039020600160a060020a03861663025ec81a836040516000805160206113bc83398151915281526009810191909152602901604051809103902060006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b15156110d657600080fd5b6102c65a03f115156110e757600080fd5b505050604051805190506040519081526020016040519081900390201415611112576000925061111f565b600101610faf565b600192505b505092915050565b6000806000611135856108ee565b9150600090505b8181101561111a5784600160a060020a03166317e7dd228260405160008051602061137c8339815191528152600f810191909152602f01604051809103902060006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b15156111ba57600080fd5b6102c65a03f115156111cb57600080fd5b50505060405180519050156112d35783604051600160a060020a03919091166c01000000000000000000000000028152601401604051908190039020600160a060020a038616634c77e5ba8360405160008051602061139c8339815191528152600c810191909152602c01604051809103902060006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b151561127d57600080fd5b6102c65a03f1151561128e57600080fd5b50505060405180519050604051600160a060020a03919091166c0100000000000000000000000002815260140160405190819003902014156112d3576000925061111f565b60010161113c565b805b82600160a060020a03166317e7dd228260405160008051602061137c8339815191528152600f810191909152602f01604051809103902060006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b151561135357600080fd5b6102c65a03f1151561136457600080fd5b5050506040518051905015610f94576001016112dd5600757365725f726567697374657265640000000000000000000000000000000000757365725f616464726573730000000000000000000000000000000000000000757365725f6e616d650000000000000000000000000000000000000000000000a165627a7a723058209933781bd855924fdf3224c10c35ea1cf60fa4ea3388a068277792452a7b02310029`

// DeployUsersLib deploys a new Ethereum contract, binding an instance of UsersLib to it.
func DeployUsersLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *UsersLib, error) {
	parsed, err := abi.JSON(strings.NewReader(UsersLibABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(UsersLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &UsersLib{UsersLibCaller: UsersLibCaller{contract: contract}, UsersLibTransactor: UsersLibTransactor{contract: contract}}, nil
}

// UsersLib is an auto generated Go binding around an Ethereum contract.
type UsersLib struct {
	UsersLibCaller     // Read-only binding to the contract
	UsersLibTransactor // Write-only binding to the contract
}

// UsersLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type UsersLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UsersLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UsersLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UsersLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UsersLibSession struct {
	Contract     *UsersLib         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UsersLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UsersLibCallerSession struct {
	Contract *UsersLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// UsersLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UsersLibTransactorSession struct {
	Contract     *UsersLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// UsersLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type UsersLibRaw struct {
	Contract *UsersLib // Generic contract binding to access the raw methods on
}

// UsersLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UsersLibCallerRaw struct {
	Contract *UsersLibCaller // Generic read-only contract binding to access the raw methods on
}

// UsersLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UsersLibTransactorRaw struct {
	Contract *UsersLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUsersLib creates a new instance of UsersLib, bound to a specific deployed contract.
func NewUsersLib(address common.Address, backend bind.ContractBackend) (*UsersLib, error) {
	contract, err := bindUsersLib(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UsersLib{UsersLibCaller: UsersLibCaller{contract: contract}, UsersLibTransactor: UsersLibTransactor{contract: contract}}, nil
}

// NewUsersLibCaller creates a new read-only instance of UsersLib, bound to a specific deployed contract.
func NewUsersLibCaller(address common.Address, caller bind.ContractCaller) (*UsersLibCaller, error) {
	contract, err := bindUsersLib(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &UsersLibCaller{contract: contract}, nil
}

// NewUsersLibTransactor creates a new write-only instance of UsersLib, bound to a specific deployed contract.
func NewUsersLibTransactor(address common.Address, transactor bind.ContractTransactor) (*UsersLibTransactor, error) {
	contract, err := bindUsersLib(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &UsersLibTransactor{contract: contract}, nil
}

// bindUsersLib binds a generic wrapper to an already deployed contract.
func bindUsersLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UsersLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UsersLib *UsersLibRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _UsersLib.Contract.UsersLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UsersLib *UsersLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UsersLib.Contract.UsersLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UsersLib *UsersLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UsersLib.Contract.UsersLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UsersLib *UsersLibCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _UsersLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UsersLib *UsersLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UsersLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UsersLib *UsersLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UsersLib.Contract.contract.Transact(opts, method, params...)
}

// GetActiveUsersCount is a free data retrieval call binding the contract method 0xc817e2df.
//
// Solidity: function getActiveUsersCount(_storageContract address) constant returns(activeUsersCount uint256)
func (_UsersLib *UsersLibCaller) GetActiveUsersCount(opts *bind.CallOpts, _storageContract common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _UsersLib.contract.Call(opts, out, "getActiveUsersCount", _storageContract)
	return *ret0, err
}

// GetActiveUsersCount is a free data retrieval call binding the contract method 0xc817e2df.
//
// Solidity: function getActiveUsersCount(_storageContract address) constant returns(activeUsersCount uint256)
func (_UsersLib *UsersLibSession) GetActiveUsersCount(_storageContract common.Address) (*big.Int, error) {
	return _UsersLib.Contract.GetActiveUsersCount(&_UsersLib.CallOpts, _storageContract)
}

// GetActiveUsersCount is a free data retrieval call binding the contract method 0xc817e2df.
//
// Solidity: function getActiveUsersCount(_storageContract address) constant returns(activeUsersCount uint256)
func (_UsersLib *UsersLibCallerSession) GetActiveUsersCount(_storageContract common.Address) (*big.Int, error) {
	return _UsersLib.Contract.GetActiveUsersCount(&_UsersLib.CallOpts, _storageContract)
}

// GetIDByAddress is a free data retrieval call binding the contract method 0xa838e451.
//
// Solidity: function getIDByAddress(_storageContract address, _address address) constant returns(found bool, id uint256)
func (_UsersLib *UsersLibCaller) GetIDByAddress(opts *bind.CallOpts, _storageContract common.Address, _address common.Address) (struct {
	Found bool
	Id    *big.Int
}, error) {
	ret := new(struct {
		Found bool
		Id    *big.Int
	})
	out := ret
	err := _UsersLib.contract.Call(opts, out, "getIDByAddress", _storageContract, _address)
	return *ret, err
}

// GetIDByAddress is a free data retrieval call binding the contract method 0xa838e451.
//
// Solidity: function getIDByAddress(_storageContract address, _address address) constant returns(found bool, id uint256)
func (_UsersLib *UsersLibSession) GetIDByAddress(_storageContract common.Address, _address common.Address) (struct {
	Found bool
	Id    *big.Int
}, error) {
	return _UsersLib.Contract.GetIDByAddress(&_UsersLib.CallOpts, _storageContract, _address)
}

// GetIDByAddress is a free data retrieval call binding the contract method 0xa838e451.
//
// Solidity: function getIDByAddress(_storageContract address, _address address) constant returns(found bool, id uint256)
func (_UsersLib *UsersLibCallerSession) GetIDByAddress(_storageContract common.Address, _address common.Address) (struct {
	Found bool
	Id    *big.Int
}, error) {
	return _UsersLib.Contract.GetIDByAddress(&_UsersLib.CallOpts, _storageContract, _address)
}

// GetIDByName is a free data retrieval call binding the contract method 0x7c6e7d45.
//
// Solidity: function getIDByName(_storageContract address, _name bytes32) constant returns(found bool, id uint256)
func (_UsersLib *UsersLibCaller) GetIDByName(opts *bind.CallOpts, _storageContract common.Address, _name [32]byte) (struct {
	Found bool
	Id    *big.Int
}, error) {
	ret := new(struct {
		Found bool
		Id    *big.Int
	})
	out := ret
	err := _UsersLib.contract.Call(opts, out, "getIDByName", _storageContract, _name)
	return *ret, err
}

// GetIDByName is a free data retrieval call binding the contract method 0x7c6e7d45.
//
// Solidity: function getIDByName(_storageContract address, _name bytes32) constant returns(found bool, id uint256)
func (_UsersLib *UsersLibSession) GetIDByName(_storageContract common.Address, _name [32]byte) (struct {
	Found bool
	Id    *big.Int
}, error) {
	return _UsersLib.Contract.GetIDByName(&_UsersLib.CallOpts, _storageContract, _name)
}

// GetIDByName is a free data retrieval call binding the contract method 0x7c6e7d45.
//
// Solidity: function getIDByName(_storageContract address, _name bytes32) constant returns(found bool, id uint256)
func (_UsersLib *UsersLibCallerSession) GetIDByName(_storageContract common.Address, _name [32]byte) (struct {
	Found bool
	Id    *big.Int
}, error) {
	return _UsersLib.Contract.GetIDByName(&_UsersLib.CallOpts, _storageContract, _name)
}

// GetUserDetailsByID is a free data retrieval call binding the contract method 0x656a1961.
//
// Solidity: function getUserDetailsByID(_storageContract address, _userID uint256) constant returns(found bool, userAddress address, userName bytes32)
func (_UsersLib *UsersLibCaller) GetUserDetailsByID(opts *bind.CallOpts, _storageContract common.Address, _userID *big.Int) (struct {
	Found       bool
	UserAddress common.Address
	UserName    [32]byte
}, error) {
	ret := new(struct {
		Found       bool
		UserAddress common.Address
		UserName    [32]byte
	})
	out := ret
	err := _UsersLib.contract.Call(opts, out, "getUserDetailsByID", _storageContract, _userID)
	return *ret, err
}

// GetUserDetailsByID is a free data retrieval call binding the contract method 0x656a1961.
//
// Solidity: function getUserDetailsByID(_storageContract address, _userID uint256) constant returns(found bool, userAddress address, userName bytes32)
func (_UsersLib *UsersLibSession) GetUserDetailsByID(_storageContract common.Address, _userID *big.Int) (struct {
	Found       bool
	UserAddress common.Address
	UserName    [32]byte
}, error) {
	return _UsersLib.Contract.GetUserDetailsByID(&_UsersLib.CallOpts, _storageContract, _userID)
}

// GetUserDetailsByID is a free data retrieval call binding the contract method 0x656a1961.
//
// Solidity: function getUserDetailsByID(_storageContract address, _userID uint256) constant returns(found bool, userAddress address, userName bytes32)
func (_UsersLib *UsersLibCallerSession) GetUserDetailsByID(_storageContract common.Address, _userID *big.Int) (struct {
	Found       bool
	UserAddress common.Address
	UserName    [32]byte
}, error) {
	return _UsersLib.Contract.GetUserDetailsByID(&_UsersLib.CallOpts, _storageContract, _userID)
}

// GetUserDetailsByName is a free data retrieval call binding the contract method 0x127de333.
//
// Solidity: function getUserDetailsByName(_storageContract address, _name bytes32) constant returns(found bool, userID uint256, userAddress address)
func (_UsersLib *UsersLibCaller) GetUserDetailsByName(opts *bind.CallOpts, _storageContract common.Address, _name [32]byte) (struct {
	Found       bool
	UserID      *big.Int
	UserAddress common.Address
}, error) {
	ret := new(struct {
		Found       bool
		UserID      *big.Int
		UserAddress common.Address
	})
	out := ret
	err := _UsersLib.contract.Call(opts, out, "getUserDetailsByName", _storageContract, _name)
	return *ret, err
}

// GetUserDetailsByName is a free data retrieval call binding the contract method 0x127de333.
//
// Solidity: function getUserDetailsByName(_storageContract address, _name bytes32) constant returns(found bool, userID uint256, userAddress address)
func (_UsersLib *UsersLibSession) GetUserDetailsByName(_storageContract common.Address, _name [32]byte) (struct {
	Found       bool
	UserID      *big.Int
	UserAddress common.Address
}, error) {
	return _UsersLib.Contract.GetUserDetailsByName(&_UsersLib.CallOpts, _storageContract, _name)
}

// GetUserDetailsByName is a free data retrieval call binding the contract method 0x127de333.
//
// Solidity: function getUserDetailsByName(_storageContract address, _name bytes32) constant returns(found bool, userID uint256, userAddress address)
func (_UsersLib *UsersLibCallerSession) GetUserDetailsByName(_storageContract common.Address, _name [32]byte) (struct {
	Found       bool
	UserID      *big.Int
	UserAddress common.Address
}, error) {
	return _UsersLib.Contract.GetUserDetailsByName(&_UsersLib.CallOpts, _storageContract, _name)
}

// GetUsersCount is a free data retrieval call binding the contract method 0x78f08268.
//
// Solidity: function getUsersCount(_storageContract address) constant returns(uint256)
func (_UsersLib *UsersLibCaller) GetUsersCount(opts *bind.CallOpts, _storageContract common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _UsersLib.contract.Call(opts, out, "getUsersCount", _storageContract)
	return *ret0, err
}

// GetUsersCount is a free data retrieval call binding the contract method 0x78f08268.
//
// Solidity: function getUsersCount(_storageContract address) constant returns(uint256)
func (_UsersLib *UsersLibSession) GetUsersCount(_storageContract common.Address) (*big.Int, error) {
	return _UsersLib.Contract.GetUsersCount(&_UsersLib.CallOpts, _storageContract)
}

// GetUsersCount is a free data retrieval call binding the contract method 0x78f08268.
//
// Solidity: function getUsersCount(_storageContract address) constant returns(uint256)
func (_UsersLib *UsersLibCallerSession) GetUsersCount(_storageContract common.Address) (*big.Int, error) {
	return _UsersLib.Contract.GetUsersCount(&_UsersLib.CallOpts, _storageContract)
}

// IsActive is a free data retrieval call binding the contract method 0xd266e83b.
//
// Solidity: function isActive(_storageContract address, _id uint256) constant returns(bool)
func (_UsersLib *UsersLibCaller) IsActive(opts *bind.CallOpts, _storageContract common.Address, _id *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _UsersLib.contract.Call(opts, out, "isActive", _storageContract, _id)
	return *ret0, err
}

// IsActive is a free data retrieval call binding the contract method 0xd266e83b.
//
// Solidity: function isActive(_storageContract address, _id uint256) constant returns(bool)
func (_UsersLib *UsersLibSession) IsActive(_storageContract common.Address, _id *big.Int) (bool, error) {
	return _UsersLib.Contract.IsActive(&_UsersLib.CallOpts, _storageContract, _id)
}

// IsActive is a free data retrieval call binding the contract method 0xd266e83b.
//
// Solidity: function isActive(_storageContract address, _id uint256) constant returns(bool)
func (_UsersLib *UsersLibCallerSession) IsActive(_storageContract common.Address, _id *big.Int) (bool, error) {
	return _UsersLib.Contract.IsActive(&_UsersLib.CallOpts, _storageContract, _id)
}

// AddUser is a paid mutator transaction binding the contract method 0x364e4ddb.
//
// Solidity: function addUser(_storageContract address, _address address, _name bytes32) returns(success bool)
func (_UsersLib *UsersLibTransactor) AddUser(opts *bind.TransactOpts, _storageContract common.Address, _address common.Address, _name [32]byte) (*types.Transaction, error) {
	return _UsersLib.contract.Transact(opts, "addUser", _storageContract, _address, _name)
}

// AddUser is a paid mutator transaction binding the contract method 0x364e4ddb.
//
// Solidity: function addUser(_storageContract address, _address address, _name bytes32) returns(success bool)
func (_UsersLib *UsersLibSession) AddUser(_storageContract common.Address, _address common.Address, _name [32]byte) (*types.Transaction, error) {
	return _UsersLib.Contract.AddUser(&_UsersLib.TransactOpts, _storageContract, _address, _name)
}

// AddUser is a paid mutator transaction binding the contract method 0x364e4ddb.
//
// Solidity: function addUser(_storageContract address, _address address, _name bytes32) returns(success bool)
func (_UsersLib *UsersLibTransactorSession) AddUser(_storageContract common.Address, _address common.Address, _name [32]byte) (*types.Transaction, error) {
	return _UsersLib.Contract.AddUser(&_UsersLib.TransactOpts, _storageContract, _address, _name)
}

// RemoveUser is a paid mutator transaction binding the contract method 0x93d3f774.
//
// Solidity: function removeUser(_storageContract address, idx uint256) returns()
func (_UsersLib *UsersLibTransactor) RemoveUser(opts *bind.TransactOpts, _storageContract common.Address, idx *big.Int) (*types.Transaction, error) {
	return _UsersLib.contract.Transact(opts, "removeUser", _storageContract, idx)
}

// RemoveUser is a paid mutator transaction binding the contract method 0x93d3f774.
//
// Solidity: function removeUser(_storageContract address, idx uint256) returns()
func (_UsersLib *UsersLibSession) RemoveUser(_storageContract common.Address, idx *big.Int) (*types.Transaction, error) {
	return _UsersLib.Contract.RemoveUser(&_UsersLib.TransactOpts, _storageContract, idx)
}

// RemoveUser is a paid mutator transaction binding the contract method 0x93d3f774.
//
// Solidity: function removeUser(_storageContract address, idx uint256) returns()
func (_UsersLib *UsersLibTransactorSession) RemoveUser(_storageContract common.Address, idx *big.Int) (*types.Transaction, error) {
	return _UsersLib.Contract.RemoveUser(&_UsersLib.TransactOpts, _storageContract, idx)
}

// UpdateAddress is a paid mutator transaction binding the contract method 0x868562b7.
//
// Solidity: function updateAddress(_storageContract address, _address address, _newAddress address) returns(bool)
func (_UsersLib *UsersLibTransactor) UpdateAddress(opts *bind.TransactOpts, _storageContract common.Address, _address common.Address, _newAddress common.Address) (*types.Transaction, error) {
	return _UsersLib.contract.Transact(opts, "updateAddress", _storageContract, _address, _newAddress)
}

// UpdateAddress is a paid mutator transaction binding the contract method 0x868562b7.
//
// Solidity: function updateAddress(_storageContract address, _address address, _newAddress address) returns(bool)
func (_UsersLib *UsersLibSession) UpdateAddress(_storageContract common.Address, _address common.Address, _newAddress common.Address) (*types.Transaction, error) {
	return _UsersLib.Contract.UpdateAddress(&_UsersLib.TransactOpts, _storageContract, _address, _newAddress)
}

// UpdateAddress is a paid mutator transaction binding the contract method 0x868562b7.
//
// Solidity: function updateAddress(_storageContract address, _address address, _newAddress address) returns(bool)
func (_UsersLib *UsersLibTransactorSession) UpdateAddress(_storageContract common.Address, _address common.Address, _newAddress common.Address) (*types.Transaction, error) {
	return _UsersLib.Contract.UpdateAddress(&_UsersLib.TransactOpts, _storageContract, _address, _newAddress)
}

// UpdateName is a paid mutator transaction binding the contract method 0x1fa768db.
//
// Solidity: function updateName(_storageContract address, _address address, _name bytes32) returns(bool)
func (_UsersLib *UsersLibTransactor) UpdateName(opts *bind.TransactOpts, _storageContract common.Address, _address common.Address, _name [32]byte) (*types.Transaction, error) {
	return _UsersLib.contract.Transact(opts, "updateName", _storageContract, _address, _name)
}

// UpdateName is a paid mutator transaction binding the contract method 0x1fa768db.
//
// Solidity: function updateName(_storageContract address, _address address, _name bytes32) returns(bool)
func (_UsersLib *UsersLibSession) UpdateName(_storageContract common.Address, _address common.Address, _name [32]byte) (*types.Transaction, error) {
	return _UsersLib.Contract.UpdateName(&_UsersLib.TransactOpts, _storageContract, _address, _name)
}

// UpdateName is a paid mutator transaction binding the contract method 0x1fa768db.
//
// Solidity: function updateName(_storageContract address, _address address, _name bytes32) returns(bool)
func (_UsersLib *UsersLibTransactorSession) UpdateName(_storageContract common.Address, _address common.Address, _name [32]byte) (*types.Transaction, error) {
	return _UsersLib.Contract.UpdateName(&_UsersLib.TransactOpts, _storageContract, _address, _name)
}

// UpdateProfilePicture is a paid mutator transaction binding the contract method 0x0eb1f411.
//
// Solidity: function updateProfilePicture(_storageContract address, _address address, _imageUrl bytes32) returns(bool)
func (_UsersLib *UsersLibTransactor) UpdateProfilePicture(opts *bind.TransactOpts, _storageContract common.Address, _address common.Address, _imageUrl [32]byte) (*types.Transaction, error) {
	return _UsersLib.contract.Transact(opts, "updateProfilePicture", _storageContract, _address, _imageUrl)
}

// UpdateProfilePicture is a paid mutator transaction binding the contract method 0x0eb1f411.
//
// Solidity: function updateProfilePicture(_storageContract address, _address address, _imageUrl bytes32) returns(bool)
func (_UsersLib *UsersLibSession) UpdateProfilePicture(_storageContract common.Address, _address common.Address, _imageUrl [32]byte) (*types.Transaction, error) {
	return _UsersLib.Contract.UpdateProfilePicture(&_UsersLib.TransactOpts, _storageContract, _address, _imageUrl)
}

// UpdateProfilePicture is a paid mutator transaction binding the contract method 0x0eb1f411.
//
// Solidity: function updateProfilePicture(_storageContract address, _address address, _imageUrl bytes32) returns(bool)
func (_UsersLib *UsersLibTransactorSession) UpdateProfilePicture(_storageContract common.Address, _address common.Address, _imageUrl [32]byte) (*types.Transaction, error) {
	return _UsersLib.Contract.UpdateProfilePicture(&_UsersLib.TransactOpts, _storageContract, _address, _imageUrl)
}
