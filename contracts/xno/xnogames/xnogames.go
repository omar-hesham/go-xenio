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

// GamesLibABI is the input ABI used to generate the binding from.
const GamesLibABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_storageContract\",\"type\":\"address\"},{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_title\",\"type\":\"bytes32\"}],\"name\":\"updateTitle\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_storageContract\",\"type\":\"address\"},{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_publisher\",\"type\":\"bytes32\"}],\"name\":\"updatePublisher\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_storageContract\",\"type\":\"address\"},{\"name\":\"_gameID\",\"type\":\"uint256\"}],\"name\":\"getGenreListByID\",\"outputs\":[{\"name\":\"genreList\",\"type\":\"bytes32[10]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_storageContract\",\"type\":\"address\"}],\"name\":\"getActiveGamesCount\",\"outputs\":[{\"name\":\"activeGamesCount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_storageContract\",\"type\":\"address\"},{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_imageUrl\",\"type\":\"bytes32\"}],\"name\":\"updateLogo\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_storageContract\",\"type\":\"address\"},{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_release\",\"type\":\"uint256\"}],\"name\":\"updateReleaseDate\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_storageContract\",\"type\":\"address\"}],\"name\":\"getGamesCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_storageContract\",\"type\":\"address\"},{\"name\":\"_title\",\"type\":\"bytes32\"}],\"name\":\"getIDByTitle\",\"outputs\":[{\"name\":\"found\",\"type\":\"bool\"},{\"name\":\"id\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_storageContract\",\"type\":\"address\"},{\"name\":\"idx\",\"type\":\"uint256\"}],\"name\":\"removeGame\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_storageContract\",\"type\":\"address\"},{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_newAddress\",\"type\":\"address\"}],\"name\":\"updateAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_storageContract\",\"type\":\"address\"},{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"updatePrice\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_storageContract\",\"type\":\"address\"},{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_genreID\",\"type\":\"uint256\"}],\"name\":\"removeGenreByID\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_storageContract\",\"type\":\"address\"},{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"getIDByAddress\",\"outputs\":[{\"name\":\"found\",\"type\":\"bool\"},{\"name\":\"id\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_storageContract\",\"type\":\"address\"},{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_developer\",\"type\":\"bytes32\"}],\"name\":\"updateDeveloper\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_storageContract\",\"type\":\"address\"},{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_genreName\",\"type\":\"bytes32\"}],\"name\":\"removeGenreByName\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_storageContract\",\"type\":\"address\"},{\"name\":\"_gameID\",\"type\":\"uint256\"}],\"name\":\"getGameDetailsByID\",\"outputs\":[{\"name\":\"found\",\"type\":\"bool\"},{\"name\":\"gameAddress\",\"type\":\"address\"},{\"name\":\"gameTitle\",\"type\":\"bytes32\"},{\"name\":\"gamePublisher\",\"type\":\"bytes32\"},{\"name\":\"gameDeveloper\",\"type\":\"bytes32\"},{\"name\":\"gameRelease\",\"type\":\"uint256\"},{\"name\":\"gamePrice\",\"type\":\"uint256\"},{\"name\":\"gameLogo\",\"type\":\"bytes32\"},{\"name\":\"gameGenreList\",\"type\":\"bytes32[10]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_storageContract\",\"type\":\"address\"},{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_genre\",\"type\":\"bytes32\"}],\"name\":\"addGenre\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_storageContract\",\"type\":\"address\"},{\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"isActive\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_storageContract\",\"type\":\"address\"},{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_title\",\"type\":\"bytes32\"},{\"name\":\"_genre\",\"type\":\"bytes32\"},{\"name\":\"_publisher\",\"type\":\"bytes32\"},{\"name\":\"_developer\",\"type\":\"bytes32\"},{\"name\":\"_release\",\"type\":\"uint256\"},{\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"addGame\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_storageContract\",\"type\":\"address\"},{\"name\":\"_title\",\"type\":\"bytes32\"}],\"name\":\"getGameDetailsByTitle\",\"outputs\":[{\"name\":\"found\",\"type\":\"bool\"},{\"name\":\"gameID\",\"type\":\"uint256\"},{\"name\":\"gameAddress\",\"type\":\"address\"},{\"name\":\"gamePublisher\",\"type\":\"bytes32\"},{\"name\":\"gameDeveloper\",\"type\":\"bytes32\"},{\"name\":\"gameRelease\",\"type\":\"uint256\"},{\"name\":\"gamePrice\",\"type\":\"uint256\"},{\"name\":\"gameLogo\",\"type\":\"bytes32\"},{\"name\":\"gameGenreList\",\"type\":\"bytes32[10]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"GameAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"GameRemoved\",\"type\":\"event\"}]"

// GamesLibBin is the compiled bytecode used for deploying new contracts.
const GamesLibBin = `0x6060604052341561000f57600080fd5b612d458061001e6000396000f3006060604052600436106100f85763ffffffff60e060020a60003504166309c1946681146100fd578063153b2d2b1461012e57806315f42bd01461014b5780632ec5a38f1461019b5780632f5fe1cf146101c1578063523f15be146101de57806359e32aeb146101fb578063723747d01461020f5780637f7c372414610240578063868562b7146102595780638fbb2a901461027957806391e6589614610296578063a838e451146102b3578063b1ed481c146102cd578063bc5e6df5146102ea578063c053dcda14610307578063ce28ed92146103a1578063d266e83b146103be578063d2dc3597146103d5578063f9580cd814610401575b600080fd5b61011a600160a060020a0360043581169060243516604435610475565b604051901515815260200160405180910390f35b61011a600160a060020a0360043581169060243516604435610534565b610162600160a060020a03600435166024356105c0565b604051808261014080838360005b83811015610188578082015183820152602001610170565b5050505090500191505060405180910390f35b6101af600160a060020a03600435166107c7565b60405190815260200160405180910390f35b61011a600160a060020a0360043581169060243516604435610894565b61011a600160a060020a0360043581169060243516604435610932565b6101af600160a060020a03600435166109be565b610226600160a060020a0360043516602435610a5f565b604051911515825260208201526040908101905180910390f35b610257600160a060020a0360043516602435610bf4565b005b61011a600160a060020a0360043581169060243581169060443516610cb5565b61011a600160a060020a0360043581169060243516604435610d4a565b61011a600160a060020a0360043581169060243516604435610dd6565b610226600160a060020a0360043581169060243516610e6b565b61011a600160a060020a0360043581169060243516604435611023565b61011a600160a060020a03600435811690602435166044356110af565b61031e600160a060020a0360043516602435611189565b6040518915158152600160a060020a038916602082015260408101889052606081018790526080810186905260a0810185905260c0810184905260e0810183905261010081018261014080838360005b8381101561038657808201518382015260200161036e565b50505050905001995050505050505050505060405180910390f35b61011a600160a060020a0360043581169060243516604435611666565b61011a600160a060020a036004351660243561188e565b61011a600160a060020a036004358116906024351660443560643560843560a43560c43560e435611929565b610418600160a060020a0360043516602435611fc3565b604051891515815260208101899052600160a060020a0388166040820152606081018790526080810186905260a0810185905260c0810184905260e08101839052610100810182610140808383600081518382015260200161036e565b60008060006104848686610e6b565b9150915081801561049a575061049a868561237b565b156105275785600160a060020a03166325cf512d82604051600080516020612cda8339815191528152600a810191909152602a0160405180910390208660405160e060020a63ffffffff851602815260048101929092526024820152604401600060405180830381600087803b151561051257600080fd5b6102c65a03f1151561052357600080fd5b5050505b8192505b50509392505050565b60008060006105438686610e6b565b9150915081156105275785600160a060020a03166325cf512d82604051600080516020612c3a8339815191528152600e810191909152602e0160405180910390208660405160e060020a63ffffffff851602815260048101929092526024820152604401600060405180830381600087803b151561051257600080fd5b6105c8612bcf565b600083600160a060020a03166317e7dd2284604051600080516020612c5a8339815191528152600f810191909152602f01604051809103902060006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b151561064057600080fd5b6102c65a03f1151561065157600080fd5b505050604051805190501515610666576107c0565b5060005b600a8110156107c05783600160a060020a03166317e7dd228483604051600080516020612c9a833981519152815260158101929092526035820152605501604051809103902060006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b15156106ef57600080fd5b6102c65a03f1151561070057600080fd5b50505060405180519050156107b85783600160a060020a031663025ec81a8483604051600080516020612bfa8339815191528152600a810192909252602a820152604a01604051809103902060006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b151561078b57600080fd5b6102c65a03f1151561079c57600080fd5b505050604051805190508282600a81106107b257fe5b60200201525b60010161066a565b5092915050565b60008060006107d5846109be565b915060009250600090505b8181101561088d5783600160a060020a03166317e7dd2282604051600080516020612c5a8339815191528152600f810191909152602f01604051809103902060006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b151561085e57600080fd5b6102c65a03f1151561086f57600080fd5b5050506040518051905015610885576001909201915b6001016107e0565b5050919050565b60008060006108a38686610e6b565b9150915081156105275785600160a060020a03166325cf512d826040517f67616d655f696d675f6c6f676f000000000000000000000000000000000000008152600d810191909152602d0160405180910390208660405160e060020a63ffffffff851602815260048101929092526024820152604401600060405180830381600087803b151561051257600080fd5b60008060006109418686610e6b565b9150915081156105275785600160a060020a0316633562fd2082604051600080516020612cba8339815191528152601181019190915260310160405180910390208660405160e060020a63ffffffff851602815260048101929092526024820152604401600060405180830381600087803b151561051257600080fd5b600081600160a060020a031663bdc963d86040517f67616d65735f636f756e740000000000000000000000000000000000000000008152600b01604051809103902060006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b1515610a3f57600080fd5b6102c65a03f11515610a5057600080fd5b50505060405180519392505050565b600080600080610a6e866109be565b9150600090505b81811015610be35785600160a060020a03166317e7dd2282604051600080516020612c5a8339815191528152600f810191909152602f01604051809103902060006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b1515610af357600080fd5b6102c65a03f11515610b0457600080fd5b5050506040518051905015610bdb5784604051908152602001604051908190039020600160a060020a03871663025ec81a83604051600080516020612cda8339815191528152600a810191909152602a01604051809103902060006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b1515610b9c57600080fd5b6102c65a03f11515610bad57600080fd5b505050604051805190506040519081526020016040519081900390201415610bdb5760018193509350610beb565b600101610a75565b600093508392505b50509250929050565b81600160a060020a0316633eba9ed282604051600080516020612c5a8339815191528152600f810191909152602f016040518091039020600060405160e060020a63ffffffff8516028152600481019290925215156024820152604401600060405180830381600087803b1515610c6a57600080fd5b6102c65a03f11515610c7b57600080fd5b5050507fa5b60816e58ebc3aefc8d7d132f4e2ac8884aad58a49e38957c45fe164f0189d8160405190815260200160405180910390a15050565b6000806000610cc48686610e6b565b9150915081156105275785600160a060020a0316635a2bf25a82604051600080516020612cfa8339815191528152600c810191909152602c0160405180910390208660405160e060020a63ffffffff85160281526004810192909252600160a060020a03166024820152604401600060405180830381600087803b151561051257600080fd5b6000806000610d598686610e6b565b9150915081156105275785600160a060020a0316633562fd2082604051600080516020612c1a8339815191528152600a810191909152602a0160405180910390208660405160e060020a63ffffffff851602815260048101929092526024820152604401600060405180830381600087803b151561051257600080fd5b6000806000610de58686610e6b565b9150915081156105275785600160a060020a0316633eba9ed28286604051600080516020612c9a8339815191528152601581019290925260358201526055016040518091039020600060405160e060020a63ffffffff8516028152600481019290925215156024820152604401600060405180830381600087803b151561051257600080fd5b600080600080610e7a866109be565b9150600090505b81811015610be35785600160a060020a03166317e7dd2282604051600080516020612c5a8339815191528152600f810191909152602f01604051809103902060006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b1515610eff57600080fd5b6102c65a03f11515610f1057600080fd5b505050604051805190501561101b5784604051600160a060020a03919091166c01000000000000000000000000028152601401604051908190039020600160a060020a038716634c77e5ba83604051600080516020612cfa8339815191528152600c810191909152602c01604051809103902060006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b1515610fc257600080fd5b6102c65a03f11515610fd357600080fd5b50505060405180519050604051600160a060020a03919091166c01000000000000000000000000028152601401604051908190039020141561101b5760018193509350610beb565b600101610e81565b60008060006110328686610e6b565b9150915081156105275785600160a060020a03166325cf512d82604051600080516020612c7a8339815191528152600e810191909152602e0160405180910390208660405160e060020a63ffffffff851602815260048101929092526024820152604401600060405180830381600087803b151561051257600080fd5b60008060008060006110c18888610e6b565b935093508315611179576110d6888488612508565b9150915081156111795787600160a060020a0316633eba9ed28483604051600080516020612c9a8339815191528152601581019290925260358201526055016040518091039020600060405160e060020a63ffffffff8516028152600481019290925215156024820152604401600060405180830381600087803b151561115c57600080fd5b6102c65a03f1151561116d57600080fd5b5050506001945061117e565b600094505b505050509392505050565b60008060008060008060008061119d612bcf565b8a600160a060020a03166317e7dd228b604051600080516020612c5a8339815191528152600f810191909152602f01604051809103902060006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b151561121357600080fd5b6102c65a03f1151561122457600080fd5b505050604051805199505088151561123b57611659565b8a600160a060020a0316634c77e5ba8b604051600080516020612cfa8339815191528152600c810191909152602c01604051809103902060006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b15156112b157600080fd5b6102c65a03f115156112c257600080fd5b5050506040518051985050600160a060020a038b1663025ec81a8b604051600080516020612cda8339815191528152600a810191909152602a01604051809103902060006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b151561134357600080fd5b6102c65a03f1151561135457600080fd5b5050506040518051975050600160a060020a038b1663025ec81a8b604051600080516020612c3a8339815191528152600e810191909152602e01604051809103902060006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b15156113d557600080fd5b6102c65a03f115156113e657600080fd5b5050506040518051965050600160a060020a038b1663025ec81a8b604051600080516020612c7a8339815191528152600e810191909152602e01604051809103902060006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b151561146757600080fd5b6102c65a03f1151561147857600080fd5b5050506040518051955050600160a060020a038b1663bdc963d88b604051600080516020612cba83398151915281526011810191909152603101604051809103902060006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b15156114f957600080fd5b6102c65a03f1151561150a57600080fd5b5050506040518051945050600160a060020a038b1663bdc963d88b604051600080516020612c1a8339815191528152600a810191909152602a01604051809103902060006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b151561158b57600080fd5b6102c65a03f1151561159c57600080fd5b5050506040518051935050600160a060020a038b1663025ec81a8b6040517f67616d655f696d675f6c6f676f000000000000000000000000000000000000008152600d810191909152602d01604051809103902060006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b151561162f57600080fd5b6102c65a03f1151561164057600080fd5b5050506040518051905091506116568b8b6105c0565b90505b9295985092959850929598565b6000806000806116768787610e6b565b9250925082156118835761168c878360006126ab565b9050611699878387612758565b15156116a457600080fd5b600a6001820111156116b557600080fd5b86600160a060020a03166325cf512d8383604051600080516020612bfa8339815191528152600a810192909252602a820152604a0160405180910390208760405160e060020a63ffffffff851602815260048101929092526024820152604401600060405180830381600087803b151561172e57600080fd5b6102c65a03f1151561173f57600080fd5b50505086600160a060020a0316633eba9ed28383604051600080516020612c9a8339815191528152601581019290925260358201526055016040518091039020600160405160e060020a63ffffffff8516028152600481019290925215156024820152604401600060405180830381600087803b15156117be57600080fd5b6102c65a03f115156117cf57600080fd5b5050506117dc87836128f1565b8160010111156118835786600160a060020a0316633562fd20836040517f67616d655f67656e7265735f636f756e740000000000000000000000000000008152601181019190915260310160405180910390208360010160405160e060020a63ffffffff851602815260048101929092526024820152604401600060405180830381600087803b151561186e57600080fd5b6102c65a03f1151561187f57600080fd5b5050505b509095945050505050565b600082600160a060020a03166317e7dd2283604051600080516020612c5a8339815191528152600f810191909152602f01604051809103902060006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b151561190657600080fd5b6102c65a03f1151561191757600080fd5b50505060405180519150505b92915050565b6000806119368a8961237b565b151561194157600080fd5b61194b8a8a61297b565b151561195657600080fd5b6119618a6000612b2f565b905089600160a060020a03166325cf512d82604051600080516020612cda8339815191528152600a810191909152602a0160405180910390208a60405160e060020a63ffffffff851602815260048101929092526024820152604401600060405180830381600087803b15156119d657600080fd5b6102c65a03f115156119e757600080fd5b50505089600160a060020a0316635a2bf25a82604051600080516020612cfa8339815191528152600c810191909152602c0160405180910390208b60405160e060020a63ffffffff85160281526004810192909252600160a060020a03166024820152604401600060405180830381600087803b1515611a6657600080fd5b6102c65a03f11515611a7757600080fd5b50505089600160a060020a0316633eba9ed282604051600080516020612c5a8339815191528152600f810191909152602f016040518091039020600160405160e060020a63ffffffff8516028152600481019290925215156024820152604401600060405180830381600087803b1515611af057600080fd5b6102c65a03f11515611b0157600080fd5b50505089600160a060020a03166325cf512d826000604051600080516020612bfa8339815191528152600a810192909252602a820152604a0160405180910390208960405160e060020a63ffffffff851602815260048101929092526024820152604401600060405180830381600087803b1515611b7e57600080fd5b6102c65a03f11515611b8f57600080fd5b50505089600160a060020a0316633eba9ed2826000604051600080516020612c9a8339815191528152601581019290925260358201526055016040518091039020600160405160e060020a63ffffffff8516028152600481019290925215156024820152604401600060405180830381600087803b1515611c0f57600080fd5b6102c65a03f11515611c2057600080fd5b50505089600160a060020a0316633562fd20826040517f67616d655f67656e7265735f636f756e74000000000000000000000000000000815260118101919091526031016040518091039020600160405160e060020a63ffffffff851602815260048101929092526024820152604401600060405180830381600087803b1515611ca957600080fd5b6102c65a03f11515611cba57600080fd5b50505089600160a060020a03166325cf512d82604051600080516020612c3a8339815191528152600e810191909152602e0160405180910390208860405160e060020a63ffffffff851602815260048101929092526024820152604401600060405180830381600087803b1515611d3057600080fd5b6102c65a03f11515611d4157600080fd5b50505089600160a060020a03166325cf512d82604051600080516020612c7a8339815191528152600e810191909152602e0160405180910390208760405160e060020a63ffffffff851602815260048101929092526024820152604401600060405180830381600087803b1515611db757600080fd5b6102c65a03f11515611dc857600080fd5b50505089600160a060020a0316633562fd2082604051600080516020612cba8339815191528152601181019190915260310160405180910390208660405160e060020a63ffffffff851602815260048101929092526024820152604401600060405180830381600087803b1515611e3e57600080fd5b6102c65a03f11515611e4f57600080fd5b50505089600160a060020a0316633562fd2082604051600080516020612c1a8339815191528152600a810191909152602a0160405180910390208560405160e060020a63ffffffff851602815260048101929092526024820152604401600060405180830381600087803b1515611ec557600080fd5b6102c65a03f11515611ed657600080fd5b505050611ee28a6109be565b816001011115611f805789600160a060020a0316633562fd206040517f67616d65735f636f756e740000000000000000000000000000000000000000008152600b0160405180910390208360010160405160e060020a63ffffffff851602815260048101929092526024820152604401600060405180830381600087803b1515611f6b57600080fd5b6102c65a03f11515611f7c57600080fd5b5050505b7fc8793439a010fad128652634a5a9a7d6b64af473d533205a069e8c64949910ef8160405190815260200160405180910390a15060019998505050505050505050565b600080600080600080600080611fd7612bcf565b611fe18b8b610a5f565b9099509750881515611ff257611659565b8a600160a060020a0316634c77e5ba89604051600080516020612cfa8339815191528152600c810191909152602c01604051809103902060006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b151561206857600080fd5b6102c65a03f1151561207957600080fd5b5050506040518051975050600160a060020a038b1663025ec81a89604051600080516020612c3a8339815191528152600e810191909152602e01604051809103902060006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b15156120fa57600080fd5b6102c65a03f1151561210b57600080fd5b5050506040518051965050600160a060020a038b1663025ec81a89604051600080516020612c7a8339815191528152600e810191909152602e01604051809103902060006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b151561218c57600080fd5b6102c65a03f1151561219d57600080fd5b5050506040518051955050600160a060020a038b1663bdc963d889604051600080516020612cba83398151915281526011810191909152603101604051809103902060006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b151561221e57600080fd5b6102c65a03f1151561222f57600080fd5b5050506040518051945050600160a060020a038b1663bdc963d889604051600080516020612c1a8339815191528152600a810191909152602a01604051809103902060006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b15156122b057600080fd5b6102c65a03f115156122c157600080fd5b5050506040518051935050600160a060020a038b1663025ec81a896040517f67616d655f696d675f6c6f676f000000000000000000000000000000000000008152600d810191909152602d01604051809103902060006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b151561235457600080fd5b6102c65a03f1151561236557600080fd5b5050506040518051905091506116568b896105c0565b6000806000612389856109be565b9150600090505b818110156124fb5784600160a060020a03166317e7dd2282604051600080516020612c5a8339815191528152600f810191909152602f01604051809103902060006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b151561240e57600080fd5b6102c65a03f1151561241f57600080fd5b50505060405180519050156124f35783604051908152602001604051908190039020600160a060020a03861663025ec81a83604051600080516020612cda8339815191528152600a810191909152602a01604051809103902060006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b15156124b757600080fd5b6102c65a03f115156124c857600080fd5b5050506040518051905060405190815260200160405190819003902014156124f35760009250612500565b600101612390565b600192505b505092915050565b60008060008061251887876128f1565b9150600090505b818110156126995786600160a060020a03166317e7dd228783604051600080516020612c9a833981519152815260158101929092526035820152605501604051809103902060006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b15156125a357600080fd5b6102c65a03f115156125b457600080fd5b50505060405180519050156126915784604051908152602001604051908190039020600160a060020a03881663025ec81a8884604051600080516020612bfa8339815191528152600a810192909252602a820152604a01604051809103902060006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b151561265257600080fd5b6102c65a03f1151561266357600080fd5b50505060405180519050604051908152602001604051908190039020141561269157600181935093506126a1565b60010161251f565b600093508392505b5050935093915050565b805b83600160a060020a03166317e7dd228483604051600080516020612c9a833981519152815260158101929092526035820152605501604051809103902060006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b151561272957600080fd5b6102c65a03f1151561273a57600080fd5b5050506040518051905015612751576001016126ad565b9392505050565b600080600061276786866128f1565b9150600090505b818110156128e55785600160a060020a03166317e7dd228683604051600080516020612c9a833981519152815260158101929092526035820152605501604051809103902060006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b15156127f257600080fd5b6102c65a03f1151561280357600080fd5b50505060405180519050156128dd5783604051908152602001604051908190039020600160a060020a03871663025ec81a8784604051600080516020612bfa8339815191528152600a810192909252602a820152604a01604051809103902060006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b15156128a157600080fd5b6102c65a03f115156128b257600080fd5b5050506040518051905060405190815260200160405190819003902014156128dd576000925061052b565b60010161276e565b50600195945050505050565b600082600160a060020a031663bdc963d8836040517f67616d655f67656e7265735f636f756e7400000000000000000000000000000081526011810191909152603101604051809103902060006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b151561190657600080fd5b6000806000612989856109be565b9150600090505b818110156124fb5784600160a060020a03166317e7dd2282604051600080516020612c5a8339815191528152600f810191909152602f01604051809103902060006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b1515612a0e57600080fd5b6102c65a03f11515612a1f57600080fd5b5050506040518051905015612b275783604051600160a060020a03919091166c01000000000000000000000000028152601401604051908190039020600160a060020a038616634c77e5ba83604051600080516020612cfa8339815191528152600c810191909152602c01604051809103902060006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b1515612ad157600080fd5b6102c65a03f11515612ae257600080fd5b50505060405180519050604051600160a060020a03919091166c010000000000000000000000000281526014016040519081900390201415612b275760009250612500565b600101612990565b805b82600160a060020a03166317e7dd2282604051600080516020612c5a8339815191528152600f810191909152602f01604051809103902060006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b1515612ba757600080fd5b6102c65a03f11515612bb857600080fd5b505050604051805190501561192357600101612b31565b610140604051908101604052600a815b60008152600019919091019060200181612bdf5790505090560067616d655f67656e72650000000000000000000000000000000000000000000067616d655f70726963650000000000000000000000000000000000000000000067616d655f7075626c697368657200000000000000000000000000000000000067616d655f72656769737465726564000000000000000000000000000000000067616d655f646576656c6f70657200000000000000000000000000000000000067616d655f67656e72655f72656769737465726564000000000000000000000067616d655f72656c656173655f6461746500000000000000000000000000000067616d655f7469746c650000000000000000000000000000000000000000000067616d655f616464726573730000000000000000000000000000000000000000a165627a7a72305820509656f6ea2bad952bb201b9e80535918c8eec4373109a41fb15dcdf3d2e2e770029`

// DeployGamesLib deploys a new Ethereum contract, binding an instance of GamesLib to it.
func DeployGamesLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GamesLib, error) {
	parsed, err := abi.JSON(strings.NewReader(GamesLibABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(GamesLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GamesLib{GamesLibCaller: GamesLibCaller{contract: contract}, GamesLibTransactor: GamesLibTransactor{contract: contract}}, nil
}

// GamesLib is an auto generated Go binding around an Ethereum contract.
type GamesLib struct {
	GamesLibCaller     // Read-only binding to the contract
	GamesLibTransactor // Write-only binding to the contract
}

// GamesLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type GamesLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GamesLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GamesLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GamesLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GamesLibSession struct {
	Contract     *GamesLib         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GamesLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GamesLibCallerSession struct {
	Contract *GamesLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// GamesLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GamesLibTransactorSession struct {
	Contract     *GamesLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// GamesLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type GamesLibRaw struct {
	Contract *GamesLib // Generic contract binding to access the raw methods on
}

// GamesLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GamesLibCallerRaw struct {
	Contract *GamesLibCaller // Generic read-only contract binding to access the raw methods on
}

// GamesLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GamesLibTransactorRaw struct {
	Contract *GamesLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGamesLib creates a new instance of GamesLib, bound to a specific deployed contract.
func NewGamesLib(address common.Address, backend bind.ContractBackend) (*GamesLib, error) {
	contract, err := bindGamesLib(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GamesLib{GamesLibCaller: GamesLibCaller{contract: contract}, GamesLibTransactor: GamesLibTransactor{contract: contract}}, nil
}

// NewGamesLibCaller creates a new read-only instance of GamesLib, bound to a specific deployed contract.
func NewGamesLibCaller(address common.Address, caller bind.ContractCaller) (*GamesLibCaller, error) {
	contract, err := bindGamesLib(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &GamesLibCaller{contract: contract}, nil
}

// NewGamesLibTransactor creates a new write-only instance of GamesLib, bound to a specific deployed contract.
func NewGamesLibTransactor(address common.Address, transactor bind.ContractTransactor) (*GamesLibTransactor, error) {
	contract, err := bindGamesLib(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &GamesLibTransactor{contract: contract}, nil
}

// bindGamesLib binds a generic wrapper to an already deployed contract.
func bindGamesLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GamesLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GamesLib *GamesLibRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GamesLib.Contract.GamesLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GamesLib *GamesLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GamesLib.Contract.GamesLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GamesLib *GamesLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GamesLib.Contract.GamesLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GamesLib *GamesLibCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GamesLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GamesLib *GamesLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GamesLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GamesLib *GamesLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GamesLib.Contract.contract.Transact(opts, method, params...)
}

// GetActiveGamesCount is a free data retrieval call binding the contract method 0x2ec5a38f.
//
// Solidity: function getActiveGamesCount(_storageContract address) constant returns(activeGamesCount uint256)
func (_GamesLib *GamesLibCaller) GetActiveGamesCount(opts *bind.CallOpts, _storageContract common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _GamesLib.contract.Call(opts, out, "getActiveGamesCount", _storageContract)
	return *ret0, err
}

// GetActiveGamesCount is a free data retrieval call binding the contract method 0x2ec5a38f.
//
// Solidity: function getActiveGamesCount(_storageContract address) constant returns(activeGamesCount uint256)
func (_GamesLib *GamesLibSession) GetActiveGamesCount(_storageContract common.Address) (*big.Int, error) {
	return _GamesLib.Contract.GetActiveGamesCount(&_GamesLib.CallOpts, _storageContract)
}

// GetActiveGamesCount is a free data retrieval call binding the contract method 0x2ec5a38f.
//
// Solidity: function getActiveGamesCount(_storageContract address) constant returns(activeGamesCount uint256)
func (_GamesLib *GamesLibCallerSession) GetActiveGamesCount(_storageContract common.Address) (*big.Int, error) {
	return _GamesLib.Contract.GetActiveGamesCount(&_GamesLib.CallOpts, _storageContract)
}

// GetGameDetailsByID is a free data retrieval call binding the contract method 0xc053dcda.
//
// Solidity: function getGameDetailsByID(_storageContract address, _gameID uint256) constant returns(found bool, gameAddress address, gameTitle bytes32, gamePublisher bytes32, gameDeveloper bytes32, gameRelease uint256, gamePrice uint256, gameLogo bytes32, gameGenreList bytes32[10])
func (_GamesLib *GamesLibCaller) GetGameDetailsByID(opts *bind.CallOpts, _storageContract common.Address, _gameID *big.Int) (struct {
	Found         bool
	GameAddress   common.Address
	GameTitle     [32]byte
	GamePublisher [32]byte
	GameDeveloper [32]byte
	GameRelease   *big.Int
	GamePrice     *big.Int
	GameLogo      [32]byte
	GameGenreList [10][32]byte
}, error) {
	ret := new(struct {
		Found         bool
		GameAddress   common.Address
		GameTitle     [32]byte
		GamePublisher [32]byte
		GameDeveloper [32]byte
		GameRelease   *big.Int
		GamePrice     *big.Int
		GameLogo      [32]byte
		GameGenreList [10][32]byte
	})
	out := ret
	err := _GamesLib.contract.Call(opts, out, "getGameDetailsByID", _storageContract, _gameID)
	return *ret, err
}

// GetGameDetailsByID is a free data retrieval call binding the contract method 0xc053dcda.
//
// Solidity: function getGameDetailsByID(_storageContract address, _gameID uint256) constant returns(found bool, gameAddress address, gameTitle bytes32, gamePublisher bytes32, gameDeveloper bytes32, gameRelease uint256, gamePrice uint256, gameLogo bytes32, gameGenreList bytes32[10])
func (_GamesLib *GamesLibSession) GetGameDetailsByID(_storageContract common.Address, _gameID *big.Int) (struct {
	Found         bool
	GameAddress   common.Address
	GameTitle     [32]byte
	GamePublisher [32]byte
	GameDeveloper [32]byte
	GameRelease   *big.Int
	GamePrice     *big.Int
	GameLogo      [32]byte
	GameGenreList [10][32]byte
}, error) {
	return _GamesLib.Contract.GetGameDetailsByID(&_GamesLib.CallOpts, _storageContract, _gameID)
}

// GetGameDetailsByID is a free data retrieval call binding the contract method 0xc053dcda.
//
// Solidity: function getGameDetailsByID(_storageContract address, _gameID uint256) constant returns(found bool, gameAddress address, gameTitle bytes32, gamePublisher bytes32, gameDeveloper bytes32, gameRelease uint256, gamePrice uint256, gameLogo bytes32, gameGenreList bytes32[10])
func (_GamesLib *GamesLibCallerSession) GetGameDetailsByID(_storageContract common.Address, _gameID *big.Int) (struct {
	Found         bool
	GameAddress   common.Address
	GameTitle     [32]byte
	GamePublisher [32]byte
	GameDeveloper [32]byte
	GameRelease   *big.Int
	GamePrice     *big.Int
	GameLogo      [32]byte
	GameGenreList [10][32]byte
}, error) {
	return _GamesLib.Contract.GetGameDetailsByID(&_GamesLib.CallOpts, _storageContract, _gameID)
}

// GetGameDetailsByTitle is a free data retrieval call binding the contract method 0xf9580cd8.
//
// Solidity: function getGameDetailsByTitle(_storageContract address, _title bytes32) constant returns(found bool, gameID uint256, gameAddress address, gamePublisher bytes32, gameDeveloper bytes32, gameRelease uint256, gamePrice uint256, gameLogo bytes32, gameGenreList bytes32[10])
func (_GamesLib *GamesLibCaller) GetGameDetailsByTitle(opts *bind.CallOpts, _storageContract common.Address, _title [32]byte) (struct {
	Found         bool
	GameID        *big.Int
	GameAddress   common.Address
	GamePublisher [32]byte
	GameDeveloper [32]byte
	GameRelease   *big.Int
	GamePrice     *big.Int
	GameLogo      [32]byte
	GameGenreList [10][32]byte
}, error) {
	ret := new(struct {
		Found         bool
		GameID        *big.Int
		GameAddress   common.Address
		GamePublisher [32]byte
		GameDeveloper [32]byte
		GameRelease   *big.Int
		GamePrice     *big.Int
		GameLogo      [32]byte
		GameGenreList [10][32]byte
	})
	out := ret
	err := _GamesLib.contract.Call(opts, out, "getGameDetailsByTitle", _storageContract, _title)
	return *ret, err
}

// GetGameDetailsByTitle is a free data retrieval call binding the contract method 0xf9580cd8.
//
// Solidity: function getGameDetailsByTitle(_storageContract address, _title bytes32) constant returns(found bool, gameID uint256, gameAddress address, gamePublisher bytes32, gameDeveloper bytes32, gameRelease uint256, gamePrice uint256, gameLogo bytes32, gameGenreList bytes32[10])
func (_GamesLib *GamesLibSession) GetGameDetailsByTitle(_storageContract common.Address, _title [32]byte) (struct {
	Found         bool
	GameID        *big.Int
	GameAddress   common.Address
	GamePublisher [32]byte
	GameDeveloper [32]byte
	GameRelease   *big.Int
	GamePrice     *big.Int
	GameLogo      [32]byte
	GameGenreList [10][32]byte
}, error) {
	return _GamesLib.Contract.GetGameDetailsByTitle(&_GamesLib.CallOpts, _storageContract, _title)
}

// GetGameDetailsByTitle is a free data retrieval call binding the contract method 0xf9580cd8.
//
// Solidity: function getGameDetailsByTitle(_storageContract address, _title bytes32) constant returns(found bool, gameID uint256, gameAddress address, gamePublisher bytes32, gameDeveloper bytes32, gameRelease uint256, gamePrice uint256, gameLogo bytes32, gameGenreList bytes32[10])
func (_GamesLib *GamesLibCallerSession) GetGameDetailsByTitle(_storageContract common.Address, _title [32]byte) (struct {
	Found         bool
	GameID        *big.Int
	GameAddress   common.Address
	GamePublisher [32]byte
	GameDeveloper [32]byte
	GameRelease   *big.Int
	GamePrice     *big.Int
	GameLogo      [32]byte
	GameGenreList [10][32]byte
}, error) {
	return _GamesLib.Contract.GetGameDetailsByTitle(&_GamesLib.CallOpts, _storageContract, _title)
}

// GetGamesCount is a free data retrieval call binding the contract method 0x59e32aeb.
//
// Solidity: function getGamesCount(_storageContract address) constant returns(uint256)
func (_GamesLib *GamesLibCaller) GetGamesCount(opts *bind.CallOpts, _storageContract common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _GamesLib.contract.Call(opts, out, "getGamesCount", _storageContract)
	return *ret0, err
}

// GetGamesCount is a free data retrieval call binding the contract method 0x59e32aeb.
//
// Solidity: function getGamesCount(_storageContract address) constant returns(uint256)
func (_GamesLib *GamesLibSession) GetGamesCount(_storageContract common.Address) (*big.Int, error) {
	return _GamesLib.Contract.GetGamesCount(&_GamesLib.CallOpts, _storageContract)
}

// GetGamesCount is a free data retrieval call binding the contract method 0x59e32aeb.
//
// Solidity: function getGamesCount(_storageContract address) constant returns(uint256)
func (_GamesLib *GamesLibCallerSession) GetGamesCount(_storageContract common.Address) (*big.Int, error) {
	return _GamesLib.Contract.GetGamesCount(&_GamesLib.CallOpts, _storageContract)
}

// GetGenreListByID is a free data retrieval call binding the contract method 0x15f42bd0.
//
// Solidity: function getGenreListByID(_storageContract address, _gameID uint256) constant returns(genreList bytes32[10])
func (_GamesLib *GamesLibCaller) GetGenreListByID(opts *bind.CallOpts, _storageContract common.Address, _gameID *big.Int) ([10][32]byte, error) {
	var (
		ret0 = new([10][32]byte)
	)
	out := ret0
	err := _GamesLib.contract.Call(opts, out, "getGenreListByID", _storageContract, _gameID)
	return *ret0, err
}

// GetGenreListByID is a free data retrieval call binding the contract method 0x15f42bd0.
//
// Solidity: function getGenreListByID(_storageContract address, _gameID uint256) constant returns(genreList bytes32[10])
func (_GamesLib *GamesLibSession) GetGenreListByID(_storageContract common.Address, _gameID *big.Int) ([10][32]byte, error) {
	return _GamesLib.Contract.GetGenreListByID(&_GamesLib.CallOpts, _storageContract, _gameID)
}

// GetGenreListByID is a free data retrieval call binding the contract method 0x15f42bd0.
//
// Solidity: function getGenreListByID(_storageContract address, _gameID uint256) constant returns(genreList bytes32[10])
func (_GamesLib *GamesLibCallerSession) GetGenreListByID(_storageContract common.Address, _gameID *big.Int) ([10][32]byte, error) {
	return _GamesLib.Contract.GetGenreListByID(&_GamesLib.CallOpts, _storageContract, _gameID)
}

// GetIDByAddress is a free data retrieval call binding the contract method 0xa838e451.
//
// Solidity: function getIDByAddress(_storageContract address, _address address) constant returns(found bool, id uint256)
func (_GamesLib *GamesLibCaller) GetIDByAddress(opts *bind.CallOpts, _storageContract common.Address, _address common.Address) (struct {
	Found bool
	Id    *big.Int
}, error) {
	ret := new(struct {
		Found bool
		Id    *big.Int
	})
	out := ret
	err := _GamesLib.contract.Call(opts, out, "getIDByAddress", _storageContract, _address)
	return *ret, err
}

// GetIDByAddress is a free data retrieval call binding the contract method 0xa838e451.
//
// Solidity: function getIDByAddress(_storageContract address, _address address) constant returns(found bool, id uint256)
func (_GamesLib *GamesLibSession) GetIDByAddress(_storageContract common.Address, _address common.Address) (struct {
	Found bool
	Id    *big.Int
}, error) {
	return _GamesLib.Contract.GetIDByAddress(&_GamesLib.CallOpts, _storageContract, _address)
}

// GetIDByAddress is a free data retrieval call binding the contract method 0xa838e451.
//
// Solidity: function getIDByAddress(_storageContract address, _address address) constant returns(found bool, id uint256)
func (_GamesLib *GamesLibCallerSession) GetIDByAddress(_storageContract common.Address, _address common.Address) (struct {
	Found bool
	Id    *big.Int
}, error) {
	return _GamesLib.Contract.GetIDByAddress(&_GamesLib.CallOpts, _storageContract, _address)
}

// GetIDByTitle is a free data retrieval call binding the contract method 0x723747d0.
//
// Solidity: function getIDByTitle(_storageContract address, _title bytes32) constant returns(found bool, id uint256)
func (_GamesLib *GamesLibCaller) GetIDByTitle(opts *bind.CallOpts, _storageContract common.Address, _title [32]byte) (struct {
	Found bool
	Id    *big.Int
}, error) {
	ret := new(struct {
		Found bool
		Id    *big.Int
	})
	out := ret
	err := _GamesLib.contract.Call(opts, out, "getIDByTitle", _storageContract, _title)
	return *ret, err
}

// GetIDByTitle is a free data retrieval call binding the contract method 0x723747d0.
//
// Solidity: function getIDByTitle(_storageContract address, _title bytes32) constant returns(found bool, id uint256)
func (_GamesLib *GamesLibSession) GetIDByTitle(_storageContract common.Address, _title [32]byte) (struct {
	Found bool
	Id    *big.Int
}, error) {
	return _GamesLib.Contract.GetIDByTitle(&_GamesLib.CallOpts, _storageContract, _title)
}

// GetIDByTitle is a free data retrieval call binding the contract method 0x723747d0.
//
// Solidity: function getIDByTitle(_storageContract address, _title bytes32) constant returns(found bool, id uint256)
func (_GamesLib *GamesLibCallerSession) GetIDByTitle(_storageContract common.Address, _title [32]byte) (struct {
	Found bool
	Id    *big.Int
}, error) {
	return _GamesLib.Contract.GetIDByTitle(&_GamesLib.CallOpts, _storageContract, _title)
}

// IsActive is a free data retrieval call binding the contract method 0xd266e83b.
//
// Solidity: function isActive(_storageContract address, _id uint256) constant returns(bool)
func (_GamesLib *GamesLibCaller) IsActive(opts *bind.CallOpts, _storageContract common.Address, _id *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _GamesLib.contract.Call(opts, out, "isActive", _storageContract, _id)
	return *ret0, err
}

// IsActive is a free data retrieval call binding the contract method 0xd266e83b.
//
// Solidity: function isActive(_storageContract address, _id uint256) constant returns(bool)
func (_GamesLib *GamesLibSession) IsActive(_storageContract common.Address, _id *big.Int) (bool, error) {
	return _GamesLib.Contract.IsActive(&_GamesLib.CallOpts, _storageContract, _id)
}

// IsActive is a free data retrieval call binding the contract method 0xd266e83b.
//
// Solidity: function isActive(_storageContract address, _id uint256) constant returns(bool)
func (_GamesLib *GamesLibCallerSession) IsActive(_storageContract common.Address, _id *big.Int) (bool, error) {
	return _GamesLib.Contract.IsActive(&_GamesLib.CallOpts, _storageContract, _id)
}

// AddGame is a paid mutator transaction binding the contract method 0xd2dc3597.
//
// Solidity: function addGame(_storageContract address, _address address, _title bytes32, _genre bytes32, _publisher bytes32, _developer bytes32, _release uint256, _price uint256) returns(success bool)
func (_GamesLib *GamesLibTransactor) AddGame(opts *bind.TransactOpts, _storageContract common.Address, _address common.Address, _title [32]byte, _genre [32]byte, _publisher [32]byte, _developer [32]byte, _release *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _GamesLib.contract.Transact(opts, "addGame", _storageContract, _address, _title, _genre, _publisher, _developer, _release, _price)
}

// AddGame is a paid mutator transaction binding the contract method 0xd2dc3597.
//
// Solidity: function addGame(_storageContract address, _address address, _title bytes32, _genre bytes32, _publisher bytes32, _developer bytes32, _release uint256, _price uint256) returns(success bool)
func (_GamesLib *GamesLibSession) AddGame(_storageContract common.Address, _address common.Address, _title [32]byte, _genre [32]byte, _publisher [32]byte, _developer [32]byte, _release *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _GamesLib.Contract.AddGame(&_GamesLib.TransactOpts, _storageContract, _address, _title, _genre, _publisher, _developer, _release, _price)
}

// AddGame is a paid mutator transaction binding the contract method 0xd2dc3597.
//
// Solidity: function addGame(_storageContract address, _address address, _title bytes32, _genre bytes32, _publisher bytes32, _developer bytes32, _release uint256, _price uint256) returns(success bool)
func (_GamesLib *GamesLibTransactorSession) AddGame(_storageContract common.Address, _address common.Address, _title [32]byte, _genre [32]byte, _publisher [32]byte, _developer [32]byte, _release *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _GamesLib.Contract.AddGame(&_GamesLib.TransactOpts, _storageContract, _address, _title, _genre, _publisher, _developer, _release, _price)
}

// AddGenre is a paid mutator transaction binding the contract method 0xce28ed92.
//
// Solidity: function addGenre(_storageContract address, _address address, _genre bytes32) returns(bool)
func (_GamesLib *GamesLibTransactor) AddGenre(opts *bind.TransactOpts, _storageContract common.Address, _address common.Address, _genre [32]byte) (*types.Transaction, error) {
	return _GamesLib.contract.Transact(opts, "addGenre", _storageContract, _address, _genre)
}

// AddGenre is a paid mutator transaction binding the contract method 0xce28ed92.
//
// Solidity: function addGenre(_storageContract address, _address address, _genre bytes32) returns(bool)
func (_GamesLib *GamesLibSession) AddGenre(_storageContract common.Address, _address common.Address, _genre [32]byte) (*types.Transaction, error) {
	return _GamesLib.Contract.AddGenre(&_GamesLib.TransactOpts, _storageContract, _address, _genre)
}

// AddGenre is a paid mutator transaction binding the contract method 0xce28ed92.
//
// Solidity: function addGenre(_storageContract address, _address address, _genre bytes32) returns(bool)
func (_GamesLib *GamesLibTransactorSession) AddGenre(_storageContract common.Address, _address common.Address, _genre [32]byte) (*types.Transaction, error) {
	return _GamesLib.Contract.AddGenre(&_GamesLib.TransactOpts, _storageContract, _address, _genre)
}

// RemoveGame is a paid mutator transaction binding the contract method 0x7f7c3724.
//
// Solidity: function removeGame(_storageContract address, idx uint256) returns()
func (_GamesLib *GamesLibTransactor) RemoveGame(opts *bind.TransactOpts, _storageContract common.Address, idx *big.Int) (*types.Transaction, error) {
	return _GamesLib.contract.Transact(opts, "removeGame", _storageContract, idx)
}

// RemoveGame is a paid mutator transaction binding the contract method 0x7f7c3724.
//
// Solidity: function removeGame(_storageContract address, idx uint256) returns()
func (_GamesLib *GamesLibSession) RemoveGame(_storageContract common.Address, idx *big.Int) (*types.Transaction, error) {
	return _GamesLib.Contract.RemoveGame(&_GamesLib.TransactOpts, _storageContract, idx)
}

// RemoveGame is a paid mutator transaction binding the contract method 0x7f7c3724.
//
// Solidity: function removeGame(_storageContract address, idx uint256) returns()
func (_GamesLib *GamesLibTransactorSession) RemoveGame(_storageContract common.Address, idx *big.Int) (*types.Transaction, error) {
	return _GamesLib.Contract.RemoveGame(&_GamesLib.TransactOpts, _storageContract, idx)
}

// RemoveGenreByID is a paid mutator transaction binding the contract method 0x91e65896.
//
// Solidity: function removeGenreByID(_storageContract address, _address address, _genreID uint256) returns(bool)
func (_GamesLib *GamesLibTransactor) RemoveGenreByID(opts *bind.TransactOpts, _storageContract common.Address, _address common.Address, _genreID *big.Int) (*types.Transaction, error) {
	return _GamesLib.contract.Transact(opts, "removeGenreByID", _storageContract, _address, _genreID)
}

// RemoveGenreByID is a paid mutator transaction binding the contract method 0x91e65896.
//
// Solidity: function removeGenreByID(_storageContract address, _address address, _genreID uint256) returns(bool)
func (_GamesLib *GamesLibSession) RemoveGenreByID(_storageContract common.Address, _address common.Address, _genreID *big.Int) (*types.Transaction, error) {
	return _GamesLib.Contract.RemoveGenreByID(&_GamesLib.TransactOpts, _storageContract, _address, _genreID)
}

// RemoveGenreByID is a paid mutator transaction binding the contract method 0x91e65896.
//
// Solidity: function removeGenreByID(_storageContract address, _address address, _genreID uint256) returns(bool)
func (_GamesLib *GamesLibTransactorSession) RemoveGenreByID(_storageContract common.Address, _address common.Address, _genreID *big.Int) (*types.Transaction, error) {
	return _GamesLib.Contract.RemoveGenreByID(&_GamesLib.TransactOpts, _storageContract, _address, _genreID)
}

// RemoveGenreByName is a paid mutator transaction binding the contract method 0xbc5e6df5.
//
// Solidity: function removeGenreByName(_storageContract address, _address address, _genreName bytes32) returns(bool)
func (_GamesLib *GamesLibTransactor) RemoveGenreByName(opts *bind.TransactOpts, _storageContract common.Address, _address common.Address, _genreName [32]byte) (*types.Transaction, error) {
	return _GamesLib.contract.Transact(opts, "removeGenreByName", _storageContract, _address, _genreName)
}

// RemoveGenreByName is a paid mutator transaction binding the contract method 0xbc5e6df5.
//
// Solidity: function removeGenreByName(_storageContract address, _address address, _genreName bytes32) returns(bool)
func (_GamesLib *GamesLibSession) RemoveGenreByName(_storageContract common.Address, _address common.Address, _genreName [32]byte) (*types.Transaction, error) {
	return _GamesLib.Contract.RemoveGenreByName(&_GamesLib.TransactOpts, _storageContract, _address, _genreName)
}

// RemoveGenreByName is a paid mutator transaction binding the contract method 0xbc5e6df5.
//
// Solidity: function removeGenreByName(_storageContract address, _address address, _genreName bytes32) returns(bool)
func (_GamesLib *GamesLibTransactorSession) RemoveGenreByName(_storageContract common.Address, _address common.Address, _genreName [32]byte) (*types.Transaction, error) {
	return _GamesLib.Contract.RemoveGenreByName(&_GamesLib.TransactOpts, _storageContract, _address, _genreName)
}

// UpdateAddress is a paid mutator transaction binding the contract method 0x868562b7.
//
// Solidity: function updateAddress(_storageContract address, _address address, _newAddress address) returns(bool)
func (_GamesLib *GamesLibTransactor) UpdateAddress(opts *bind.TransactOpts, _storageContract common.Address, _address common.Address, _newAddress common.Address) (*types.Transaction, error) {
	return _GamesLib.contract.Transact(opts, "updateAddress", _storageContract, _address, _newAddress)
}

// UpdateAddress is a paid mutator transaction binding the contract method 0x868562b7.
//
// Solidity: function updateAddress(_storageContract address, _address address, _newAddress address) returns(bool)
func (_GamesLib *GamesLibSession) UpdateAddress(_storageContract common.Address, _address common.Address, _newAddress common.Address) (*types.Transaction, error) {
	return _GamesLib.Contract.UpdateAddress(&_GamesLib.TransactOpts, _storageContract, _address, _newAddress)
}

// UpdateAddress is a paid mutator transaction binding the contract method 0x868562b7.
//
// Solidity: function updateAddress(_storageContract address, _address address, _newAddress address) returns(bool)
func (_GamesLib *GamesLibTransactorSession) UpdateAddress(_storageContract common.Address, _address common.Address, _newAddress common.Address) (*types.Transaction, error) {
	return _GamesLib.Contract.UpdateAddress(&_GamesLib.TransactOpts, _storageContract, _address, _newAddress)
}

// UpdateDeveloper is a paid mutator transaction binding the contract method 0xb1ed481c.
//
// Solidity: function updateDeveloper(_storageContract address, _address address, _developer bytes32) returns(bool)
func (_GamesLib *GamesLibTransactor) UpdateDeveloper(opts *bind.TransactOpts, _storageContract common.Address, _address common.Address, _developer [32]byte) (*types.Transaction, error) {
	return _GamesLib.contract.Transact(opts, "updateDeveloper", _storageContract, _address, _developer)
}

// UpdateDeveloper is a paid mutator transaction binding the contract method 0xb1ed481c.
//
// Solidity: function updateDeveloper(_storageContract address, _address address, _developer bytes32) returns(bool)
func (_GamesLib *GamesLibSession) UpdateDeveloper(_storageContract common.Address, _address common.Address, _developer [32]byte) (*types.Transaction, error) {
	return _GamesLib.Contract.UpdateDeveloper(&_GamesLib.TransactOpts, _storageContract, _address, _developer)
}

// UpdateDeveloper is a paid mutator transaction binding the contract method 0xb1ed481c.
//
// Solidity: function updateDeveloper(_storageContract address, _address address, _developer bytes32) returns(bool)
func (_GamesLib *GamesLibTransactorSession) UpdateDeveloper(_storageContract common.Address, _address common.Address, _developer [32]byte) (*types.Transaction, error) {
	return _GamesLib.Contract.UpdateDeveloper(&_GamesLib.TransactOpts, _storageContract, _address, _developer)
}

// UpdateLogo is a paid mutator transaction binding the contract method 0x2f5fe1cf.
//
// Solidity: function updateLogo(_storageContract address, _address address, _imageUrl bytes32) returns(bool)
func (_GamesLib *GamesLibTransactor) UpdateLogo(opts *bind.TransactOpts, _storageContract common.Address, _address common.Address, _imageUrl [32]byte) (*types.Transaction, error) {
	return _GamesLib.contract.Transact(opts, "updateLogo", _storageContract, _address, _imageUrl)
}

// UpdateLogo is a paid mutator transaction binding the contract method 0x2f5fe1cf.
//
// Solidity: function updateLogo(_storageContract address, _address address, _imageUrl bytes32) returns(bool)
func (_GamesLib *GamesLibSession) UpdateLogo(_storageContract common.Address, _address common.Address, _imageUrl [32]byte) (*types.Transaction, error) {
	return _GamesLib.Contract.UpdateLogo(&_GamesLib.TransactOpts, _storageContract, _address, _imageUrl)
}

// UpdateLogo is a paid mutator transaction binding the contract method 0x2f5fe1cf.
//
// Solidity: function updateLogo(_storageContract address, _address address, _imageUrl bytes32) returns(bool)
func (_GamesLib *GamesLibTransactorSession) UpdateLogo(_storageContract common.Address, _address common.Address, _imageUrl [32]byte) (*types.Transaction, error) {
	return _GamesLib.Contract.UpdateLogo(&_GamesLib.TransactOpts, _storageContract, _address, _imageUrl)
}

// UpdatePrice is a paid mutator transaction binding the contract method 0x8fbb2a90.
//
// Solidity: function updatePrice(_storageContract address, _address address, _price uint256) returns(bool)
func (_GamesLib *GamesLibTransactor) UpdatePrice(opts *bind.TransactOpts, _storageContract common.Address, _address common.Address, _price *big.Int) (*types.Transaction, error) {
	return _GamesLib.contract.Transact(opts, "updatePrice", _storageContract, _address, _price)
}

// UpdatePrice is a paid mutator transaction binding the contract method 0x8fbb2a90.
//
// Solidity: function updatePrice(_storageContract address, _address address, _price uint256) returns(bool)
func (_GamesLib *GamesLibSession) UpdatePrice(_storageContract common.Address, _address common.Address, _price *big.Int) (*types.Transaction, error) {
	return _GamesLib.Contract.UpdatePrice(&_GamesLib.TransactOpts, _storageContract, _address, _price)
}

// UpdatePrice is a paid mutator transaction binding the contract method 0x8fbb2a90.
//
// Solidity: function updatePrice(_storageContract address, _address address, _price uint256) returns(bool)
func (_GamesLib *GamesLibTransactorSession) UpdatePrice(_storageContract common.Address, _address common.Address, _price *big.Int) (*types.Transaction, error) {
	return _GamesLib.Contract.UpdatePrice(&_GamesLib.TransactOpts, _storageContract, _address, _price)
}

// UpdatePublisher is a paid mutator transaction binding the contract method 0x153b2d2b.
//
// Solidity: function updatePublisher(_storageContract address, _address address, _publisher bytes32) returns(bool)
func (_GamesLib *GamesLibTransactor) UpdatePublisher(opts *bind.TransactOpts, _storageContract common.Address, _address common.Address, _publisher [32]byte) (*types.Transaction, error) {
	return _GamesLib.contract.Transact(opts, "updatePublisher", _storageContract, _address, _publisher)
}

// UpdatePublisher is a paid mutator transaction binding the contract method 0x153b2d2b.
//
// Solidity: function updatePublisher(_storageContract address, _address address, _publisher bytes32) returns(bool)
func (_GamesLib *GamesLibSession) UpdatePublisher(_storageContract common.Address, _address common.Address, _publisher [32]byte) (*types.Transaction, error) {
	return _GamesLib.Contract.UpdatePublisher(&_GamesLib.TransactOpts, _storageContract, _address, _publisher)
}

// UpdatePublisher is a paid mutator transaction binding the contract method 0x153b2d2b.
//
// Solidity: function updatePublisher(_storageContract address, _address address, _publisher bytes32) returns(bool)
func (_GamesLib *GamesLibTransactorSession) UpdatePublisher(_storageContract common.Address, _address common.Address, _publisher [32]byte) (*types.Transaction, error) {
	return _GamesLib.Contract.UpdatePublisher(&_GamesLib.TransactOpts, _storageContract, _address, _publisher)
}

// UpdateReleaseDate is a paid mutator transaction binding the contract method 0x523f15be.
//
// Solidity: function updateReleaseDate(_storageContract address, _address address, _release uint256) returns(bool)
func (_GamesLib *GamesLibTransactor) UpdateReleaseDate(opts *bind.TransactOpts, _storageContract common.Address, _address common.Address, _release *big.Int) (*types.Transaction, error) {
	return _GamesLib.contract.Transact(opts, "updateReleaseDate", _storageContract, _address, _release)
}

// UpdateReleaseDate is a paid mutator transaction binding the contract method 0x523f15be.
//
// Solidity: function updateReleaseDate(_storageContract address, _address address, _release uint256) returns(bool)
func (_GamesLib *GamesLibSession) UpdateReleaseDate(_storageContract common.Address, _address common.Address, _release *big.Int) (*types.Transaction, error) {
	return _GamesLib.Contract.UpdateReleaseDate(&_GamesLib.TransactOpts, _storageContract, _address, _release)
}

// UpdateReleaseDate is a paid mutator transaction binding the contract method 0x523f15be.
//
// Solidity: function updateReleaseDate(_storageContract address, _address address, _release uint256) returns(bool)
func (_GamesLib *GamesLibTransactorSession) UpdateReleaseDate(_storageContract common.Address, _address common.Address, _release *big.Int) (*types.Transaction, error) {
	return _GamesLib.Contract.UpdateReleaseDate(&_GamesLib.TransactOpts, _storageContract, _address, _release)
}

// UpdateTitle is a paid mutator transaction binding the contract method 0x09c19466.
//
// Solidity: function updateTitle(_storageContract address, _address address, _title bytes32) returns(bool)
func (_GamesLib *GamesLibTransactor) UpdateTitle(opts *bind.TransactOpts, _storageContract common.Address, _address common.Address, _title [32]byte) (*types.Transaction, error) {
	return _GamesLib.contract.Transact(opts, "updateTitle", _storageContract, _address, _title)
}

// UpdateTitle is a paid mutator transaction binding the contract method 0x09c19466.
//
// Solidity: function updateTitle(_storageContract address, _address address, _title bytes32) returns(bool)
func (_GamesLib *GamesLibSession) UpdateTitle(_storageContract common.Address, _address common.Address, _title [32]byte) (*types.Transaction, error) {
	return _GamesLib.Contract.UpdateTitle(&_GamesLib.TransactOpts, _storageContract, _address, _title)
}

// UpdateTitle is a paid mutator transaction binding the contract method 0x09c19466.
//
// Solidity: function updateTitle(_storageContract address, _address address, _title bytes32) returns(bool)
func (_GamesLib *GamesLibTransactorSession) UpdateTitle(_storageContract common.Address, _address common.Address, _title [32]byte) (*types.Transaction, error) {
	return _GamesLib.Contract.UpdateTitle(&_GamesLib.TransactOpts, _storageContract, _address, _title)
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

// SecurityLibABI is the input ABI used to generate the binding from.
const SecurityLibABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_storageContract\",\"type\":\"address\"}],\"name\":\"getAdminsCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_storageContract\",\"type\":\"address\"},{\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"removeAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_storageContract\",\"type\":\"address\"},{\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"isUserAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_storageContract\",\"type\":\"address\"},{\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"AdminAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"AdminRemoved\",\"type\":\"event\"}]"

// SecurityLibBin is the compiled bytecode used for deploying new contracts.
const SecurityLibBin = `0x6060604052341561000f57600080fd5b6107b48061001e6000396000f3006060604052600436106100485763ffffffff60e060020a6000350416631945bb72811461004d578063268959e5146100735780634cbc5e991461008f5780638a19b4d2146100bd575b600080fd5b610061600160a060020a03600435166100d7565b60405190815260200160405180910390f35b61008d600160a060020a0360043581169060243516610166565b005b6100a9600160a060020a0360043581169060243516610419565b604051901515815260200160405180910390f35b61008d600160a060020a03600435811690602435166104c7565b600081600160a060020a031663bdc963d86040516000805160206107698339815191528152600c01604051809103902060006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b151561014657600080fd5b6102c65a03f1151561015757600080fd5b50505060405180519392505050565b60008083600160a060020a03166317e7dd228460405160d960020a6430b236b4b7028152600160a060020a03919091166c01000000000000000000000000026005820152601901604051809103902060006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b15156101f457600080fd5b6102c65a03f1151561020557600080fd5b505050604051805192505081151561021c57600080fd5b83600160a060020a031663bdc963d86040516000805160206107698339815191528152600c01604051809103902060006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b151561028957600080fd5b6102c65a03f1151561029a57600080fd5b505050604051805191505060018114156102b357600080fd5b83600160a060020a0316633eba9ed28460405160d960020a6430b236b4b7028152600160a060020a03919091166c010000000000000000000000000260058201526019016040518091039020600060405160e060020a63ffffffff8516028152600481019290925215156024820152604401600060405180830381600087803b151561033e57600080fd5b6102c65a03f1151561034f57600080fd5b50505060018103905083600160a060020a0316633562fd206040516000805160206107698339815191528152600c0160405180910390208360405160e060020a63ffffffff851602815260048101929092526024820152604401600060405180830381600087803b15156103c257600080fd5b6102c65a03f115156103d357600080fd5b5050507fa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f83604051600160a060020a03909116815260200160405180910390a150505050565b600082600160a060020a03166317e7dd228360405160d960020a6430b236b4b7028152600160a060020a03919091166c01000000000000000000000000026005820152601901604051809103902060006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b15156104a657600080fd5b6102c65a03f115156104b757600080fd5b5050506040518051949350505050565b60008083600160a060020a03166317e7dd228460405160d960020a6430b236b4b7028152600160a060020a03919091166c01000000000000000000000000026005820152601901604051809103902060006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b151561055557600080fd5b6102c65a03f1151561056657600080fd5b5050506040518051925050811561057c57600080fd5b83600160a060020a0316633eba9ed28460405160d960020a6430b236b4b7028152600160a060020a03919091166c010000000000000000000000000260058201526019016040518091039020600160405160e060020a63ffffffff8516028152600481019290925215156024820152604401600060405180830381600087803b151561060757600080fd5b6102c65a03f1151561061857600080fd5b50505083600160a060020a031663bdc963d86040516000805160206107698339815191528152600c01604051809103902060006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b151561068857600080fd5b6102c65a03f1151561069957600080fd5b5050506040518051915050600160a060020a038416633562fd206040516000805160206107698339815191528152600c0160405180910390208360010160405160e060020a63ffffffff851602815260048101929092526024820152604401600060405180830381600087803b151561071157600080fd5b6102c65a03f1151561072257600080fd5b5050507f44d6d25963f097ad14f29f06854a01f575648a1ef82f30e562ccd3889717e33983604051600160a060020a03909116815260200160405180910390a150505050560061646d696e5f636f756e74730000000000000000000000000000000000000000a165627a7a723058201b330c794829a12f8a0d2fdf6f8d020f09c38474bd405a94a378a0856ae8166b0029`

// DeploySecurityLib deploys a new Ethereum contract, binding an instance of SecurityLib to it.
func DeploySecurityLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SecurityLib, error) {
	parsed, err := abi.JSON(strings.NewReader(SecurityLibABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SecurityLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SecurityLib{SecurityLibCaller: SecurityLibCaller{contract: contract}, SecurityLibTransactor: SecurityLibTransactor{contract: contract}}, nil
}

// SecurityLib is an auto generated Go binding around an Ethereum contract.
type SecurityLib struct {
	SecurityLibCaller     // Read-only binding to the contract
	SecurityLibTransactor // Write-only binding to the contract
}

// SecurityLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type SecurityLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SecurityLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SecurityLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SecurityLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SecurityLibSession struct {
	Contract     *SecurityLib      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SecurityLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SecurityLibCallerSession struct {
	Contract *SecurityLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// SecurityLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SecurityLibTransactorSession struct {
	Contract     *SecurityLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// SecurityLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type SecurityLibRaw struct {
	Contract *SecurityLib // Generic contract binding to access the raw methods on
}

// SecurityLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SecurityLibCallerRaw struct {
	Contract *SecurityLibCaller // Generic read-only contract binding to access the raw methods on
}

// SecurityLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SecurityLibTransactorRaw struct {
	Contract *SecurityLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSecurityLib creates a new instance of SecurityLib, bound to a specific deployed contract.
func NewSecurityLib(address common.Address, backend bind.ContractBackend) (*SecurityLib, error) {
	contract, err := bindSecurityLib(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SecurityLib{SecurityLibCaller: SecurityLibCaller{contract: contract}, SecurityLibTransactor: SecurityLibTransactor{contract: contract}}, nil
}

// NewSecurityLibCaller creates a new read-only instance of SecurityLib, bound to a specific deployed contract.
func NewSecurityLibCaller(address common.Address, caller bind.ContractCaller) (*SecurityLibCaller, error) {
	contract, err := bindSecurityLib(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &SecurityLibCaller{contract: contract}, nil
}

// NewSecurityLibTransactor creates a new write-only instance of SecurityLib, bound to a specific deployed contract.
func NewSecurityLibTransactor(address common.Address, transactor bind.ContractTransactor) (*SecurityLibTransactor, error) {
	contract, err := bindSecurityLib(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &SecurityLibTransactor{contract: contract}, nil
}

// bindSecurityLib binds a generic wrapper to an already deployed contract.
func bindSecurityLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SecurityLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SecurityLib *SecurityLibRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SecurityLib.Contract.SecurityLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SecurityLib *SecurityLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SecurityLib.Contract.SecurityLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SecurityLib *SecurityLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SecurityLib.Contract.SecurityLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SecurityLib *SecurityLibCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SecurityLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SecurityLib *SecurityLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SecurityLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SecurityLib *SecurityLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SecurityLib.Contract.contract.Transact(opts, method, params...)
}

// GetAdminsCount is a free data retrieval call binding the contract method 0x1945bb72.
//
// Solidity: function getAdminsCount(_storageContract address) constant returns(uint256)
func (_SecurityLib *SecurityLibCaller) GetAdminsCount(opts *bind.CallOpts, _storageContract common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SecurityLib.contract.Call(opts, out, "getAdminsCount", _storageContract)
	return *ret0, err
}

// GetAdminsCount is a free data retrieval call binding the contract method 0x1945bb72.
//
// Solidity: function getAdminsCount(_storageContract address) constant returns(uint256)
func (_SecurityLib *SecurityLibSession) GetAdminsCount(_storageContract common.Address) (*big.Int, error) {
	return _SecurityLib.Contract.GetAdminsCount(&_SecurityLib.CallOpts, _storageContract)
}

// GetAdminsCount is a free data retrieval call binding the contract method 0x1945bb72.
//
// Solidity: function getAdminsCount(_storageContract address) constant returns(uint256)
func (_SecurityLib *SecurityLibCallerSession) GetAdminsCount(_storageContract common.Address) (*big.Int, error) {
	return _SecurityLib.Contract.GetAdminsCount(&_SecurityLib.CallOpts, _storageContract)
}

// IsUserAdmin is a free data retrieval call binding the contract method 0x4cbc5e99.
//
// Solidity: function isUserAdmin(_storageContract address, _user address) constant returns(bool)
func (_SecurityLib *SecurityLibCaller) IsUserAdmin(opts *bind.CallOpts, _storageContract common.Address, _user common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SecurityLib.contract.Call(opts, out, "isUserAdmin", _storageContract, _user)
	return *ret0, err
}

// IsUserAdmin is a free data retrieval call binding the contract method 0x4cbc5e99.
//
// Solidity: function isUserAdmin(_storageContract address, _user address) constant returns(bool)
func (_SecurityLib *SecurityLibSession) IsUserAdmin(_storageContract common.Address, _user common.Address) (bool, error) {
	return _SecurityLib.Contract.IsUserAdmin(&_SecurityLib.CallOpts, _storageContract, _user)
}

// IsUserAdmin is a free data retrieval call binding the contract method 0x4cbc5e99.
//
// Solidity: function isUserAdmin(_storageContract address, _user address) constant returns(bool)
func (_SecurityLib *SecurityLibCallerSession) IsUserAdmin(_storageContract common.Address, _user common.Address) (bool, error) {
	return _SecurityLib.Contract.IsUserAdmin(&_SecurityLib.CallOpts, _storageContract, _user)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x8a19b4d2.
//
// Solidity: function addAdmin(_storageContract address, _user address) returns()
func (_SecurityLib *SecurityLibTransactor) AddAdmin(opts *bind.TransactOpts, _storageContract common.Address, _user common.Address) (*types.Transaction, error) {
	return _SecurityLib.contract.Transact(opts, "addAdmin", _storageContract, _user)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x8a19b4d2.
//
// Solidity: function addAdmin(_storageContract address, _user address) returns()
func (_SecurityLib *SecurityLibSession) AddAdmin(_storageContract common.Address, _user common.Address) (*types.Transaction, error) {
	return _SecurityLib.Contract.AddAdmin(&_SecurityLib.TransactOpts, _storageContract, _user)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x8a19b4d2.
//
// Solidity: function addAdmin(_storageContract address, _user address) returns()
func (_SecurityLib *SecurityLibTransactorSession) AddAdmin(_storageContract common.Address, _user common.Address) (*types.Transaction, error) {
	return _SecurityLib.Contract.AddAdmin(&_SecurityLib.TransactOpts, _storageContract, _user)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x268959e5.
//
// Solidity: function removeAdmin(_storageContract address, _user address) returns()
func (_SecurityLib *SecurityLibTransactor) RemoveAdmin(opts *bind.TransactOpts, _storageContract common.Address, _user common.Address) (*types.Transaction, error) {
	return _SecurityLib.contract.Transact(opts, "removeAdmin", _storageContract, _user)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x268959e5.
//
// Solidity: function removeAdmin(_storageContract address, _user address) returns()
func (_SecurityLib *SecurityLibSession) RemoveAdmin(_storageContract common.Address, _user common.Address) (*types.Transaction, error) {
	return _SecurityLib.Contract.RemoveAdmin(&_SecurityLib.TransactOpts, _storageContract, _user)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x268959e5.
//
// Solidity: function removeAdmin(_storageContract address, _user address) returns()
func (_SecurityLib *SecurityLibTransactorSession) RemoveAdmin(_storageContract common.Address, _user common.Address) (*types.Transaction, error) {
	return _SecurityLib.Contract.RemoveAdmin(&_SecurityLib.TransactOpts, _storageContract, _user)
}

// XnoGamesABI is the input ABI used to generate the binding from.
const XnoGamesABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"removeXnoAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_title\",\"type\":\"bytes32\"}],\"name\":\"removeGameByTitle\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getGamesCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deleteStorage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_developer\",\"type\":\"bytes32\"}],\"name\":\"changeGameDeveloper\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_title\",\"type\":\"bytes32\"}],\"name\":\"updateGameTitle\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_publisher\",\"type\":\"bytes32\"}],\"name\":\"updateGamePublisher\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_title\",\"type\":\"bytes32\"},{\"name\":\"_genre\",\"type\":\"bytes32\"},{\"name\":\"_publisher\",\"type\":\"bytes32\"},{\"name\":\"_developer\",\"type\":\"bytes32\"},{\"name\":\"_release\",\"type\":\"uint256\"},{\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"registerNewGame\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_release\",\"type\":\"uint256\"}],\"name\":\"updateGameReleaseDate\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_genreID\",\"type\":\"uint256\"}],\"name\":\"deleteGameGenreByID\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deactivate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_imgUrl\",\"type\":\"bytes32\"}],\"name\":\"updateGameLogo\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"passOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_genreID\",\"type\":\"uint256\"}],\"name\":\"removeGameGenreByID\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"createNewStorage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_title\",\"type\":\"bytes32\"}],\"name\":\"getGameIDByTitle\",\"outputs\":[{\"name\":\"found\",\"type\":\"bool\"},{\"name\":\"gameID\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"removeGameByAddress\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getStorageOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"addXnoAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"changeGamePrice\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"updateGamePrice\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"eternalStorage\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_developer\",\"type\":\"bytes32\"}],\"name\":\"updateGameDeveloper\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_currentAddress\",\"type\":\"address\"},{\"name\":\"_newAddress\",\"type\":\"address\"}],\"name\":\"changeGameAddress\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"getGameIDByAddress\",\"outputs\":[{\"name\":\"found\",\"type\":\"bool\"},{\"name\":\"gameID\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_gameID\",\"type\":\"uint256\"}],\"name\":\"getGameByID\",\"outputs\":[{\"name\":\"found\",\"type\":\"bool\"},{\"name\":\"gameAddress\",\"type\":\"address\"},{\"name\":\"gameTitle\",\"type\":\"bytes32\"},{\"name\":\"gamePublisher\",\"type\":\"bytes32\"},{\"name\":\"gameDeveloper\",\"type\":\"bytes32\"},{\"name\":\"gameRelease\",\"type\":\"uint256\"},{\"name\":\"gamePrice\",\"type\":\"uint256\"},{\"name\":\"gameLogo\",\"type\":\"bytes32\"},{\"name\":\"gameGenreList\",\"type\":\"bytes32[10]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_publisher\",\"type\":\"bytes32\"}],\"name\":\"changeGamePublisher\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_genreName\",\"type\":\"bytes32\"}],\"name\":\"deleteGameGenreByName\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_imgUrl\",\"type\":\"bytes32\"}],\"name\":\"changeGameLogo\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_genreName\",\"type\":\"bytes32\"}],\"name\":\"removeGameGenreByName\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAllGameIDs\",\"outputs\":[{\"name\":\"list\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_title\",\"type\":\"bytes32\"}],\"name\":\"getGameByTitle\",\"outputs\":[{\"name\":\"found\",\"type\":\"bool\"},{\"name\":\"gameID\",\"type\":\"uint256\"},{\"name\":\"gameAddress\",\"type\":\"address\"},{\"name\":\"gamePublisher\",\"type\":\"bytes32\"},{\"name\":\"gameDeveloper\",\"type\":\"bytes32\"},{\"name\":\"gameRelease\",\"type\":\"uint256\"},{\"name\":\"gamePrice\",\"type\":\"uint256\"},{\"name\":\"gameLogo\",\"type\":\"bytes32\"},{\"name\":\"gameGenreList\",\"type\":\"bytes32[10]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_title\",\"type\":\"bytes32\"}],\"name\":\"changeGameTitle\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_storageContract\",\"type\":\"address\"}],\"name\":\"attachExistingStorage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_genre\",\"type\":\"bytes32\"}],\"name\":\"insertGameGenre\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOnwer\",\"type\":\"address\"}],\"name\":\"detachStorage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_genre\",\"type\":\"bytes32\"}],\"name\":\"addGameGenre\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deregisterGame\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_release\",\"type\":\"uint256\"}],\"name\":\"changeGameReleaseDate\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_storageAddress\",\"type\":\"address\"}],\"name\":\"StorageCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_storageAddress\",\"type\":\"address\"}],\"name\":\"StorageAttached\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"StorageDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"StorageDetached\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"XnoGamesCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"XnoGamesDeactivated\",\"type\":\"event\"}]"

// XnoGamesBin is the compiled bytecode used for deploying new contracts.
const XnoGamesBin = `0x606060405234156200001057600080fd5b60008054600160a060020a03191633600160a060020a03161790557fa0a81062e1419b8afa72ddd0fa9d4ad760c90ae5f54bd5c699c4f78801da501062000064640100000000620000858102620016981704565b604051600160a060020a03909116815260200160405180910390a162000094565b600054600160a060020a031690565b61383b80620000a46000396000f3006060604052600436106101de5763ffffffff60e060020a600035041662ef809e81146101e35780630bd51cc51461020457806317b7a4761461022e578063192150e6146102535780631b275b2a1461026657806322068654146102885780632c7d55c31461029e57806341c0e1b5146102b45780634de67dac146102c75780634edd5b6f146102ec5780634fb177781461030257806351b42b00146103185780635ab753711461032b578063620996df146103415780636dd8f6cf146103605780636dddf403146103825780637456d6771461039557806383804603146103c5578063893d20e8146103e45780638ae7597714610413578063928d9ef0146104265780639470211e14610445578063969d6abd1461046757806398ff9c541461047d578063a0e5467514610490578063a1bd7051146104a6578063a9fa580c146104cb578063accebdfc146104ea578063add4612014610583578063aedb8d34146105a5578063b59ed827146105bb578063b8ca9200146105dd578063bbed50aa146105ff578063c0f7944114610665578063c5671765146106d8578063c99fe8d8146106fa578063da307fa314610719578063df3db6701461073b578063e5babda11461075a578063f9e1a45d14610770578063f9ea4c3714610783575b600080fd5b34156101ee57600080fd5b610202600160a060020a03600435166107a5565b005b341561020f57600080fd5b61021a60043561085e565b604051901515815260200160405180910390f35b341561023957600080fd5b610241610a40565b60405190815260200160405180910390f35b341561025e57600080fd5b610202610aeb565b341561027157600080fd5b61021a600160a060020a0360043516602435610bbf565b341561029357600080fd5b61021a600435610d1a565b34156102a957600080fd5b61021a600435610dd9565b34156102bf57600080fd5b610202610e77565b34156102d257600080fd5b61021a60043560243560443560643560843560a435610ea0565b34156102f757600080fd5b61021a600435610f89565b341561030d57600080fd5b61021a600435611027565b341561032357600080fd5b6102026110c5565b341561033657600080fd5b61021a600435611130565b341561034c57600080fd5b610202600160a060020a03600435166111ce565b341561036b57600080fd5b61021a600160a060020a036004351660243561122d565b341561038d57600080fd5b610202611366565b34156103a057600080fd5b6103ab6004356114ab565b604051911515825260208201526040908101905180910390f35b34156103d057600080fd5b61021a600160a060020a0360043516611567565b34156103ef57600080fd5b6103f7611698565b604051600160a060020a03909116815260200160405180910390f35b341561041e57600080fd5b6103f76116a7565b341561043157600080fd5b610202600160a060020a036004351661171e565b341561045057600080fd5b61021a600160a060020a03600435166024356117bf565b341561047257600080fd5b61021a6004356118f8565b341561048857600080fd5b6103f7611996565b341561049b57600080fd5b61021a6004356119a5565b34156104b157600080fd5b61021a600160a060020a0360043581169060243516611a43565b34156104d657600080fd5b6103ab600160a060020a0360043516611b7b565b34156104f557600080fd5b610500600435611c0e565b6040518915158152600160a060020a038916602082015260408101889052606081018790526080810186905260a0810185905260c0810184905260e0810183905261010081018261014080838360005b83811015610568578082015183820152602001610550565b50505050905001995050505050505050505060405180910390f35b341561058e57600080fd5b61021a600160a060020a0360043516602435611d20565b34156105b057600080fd5b61021a600435611e59565b34156105c657600080fd5b61021a600160a060020a0360043516602435611ef7565b34156105e857600080fd5b61021a600160a060020a0360043516602435612030565b341561060a57600080fd5b610612612169565b60405160208082528190810183818151815260200191508051906020019060200280838360005b83811015610651578082015183820152602001610639565b505050509050019250505060405180910390f35b341561067057600080fd5b61067b600435612391565b604051891515815260208101899052600160a060020a0388166040820152606081018790526080810186905260a0810185905260c0810184905260e081018390526101008101826101408083836000815183820152602001610550565b34156106e357600080fd5b61021a600160a060020a0360043516602435612434565b341561070557600080fd5b610202600160a060020a036004351661256d565b341561072457600080fd5b61021a600160a060020a03600435166024356126fe565b341561074657600080fd5b610202600160a060020a0360043516612837565b341561076557600080fd5b61021a6004356128d3565b341561077b57600080fd5b610202612971565b341561078e57600080fd5b61021a600160a060020a0360043516602435612ab3565b60018054600160a060020a0316151581146107bf57600080fd5b60005433600160a060020a039081169116146107da57600080fd5b60015473__SecurityLib.sol:SecurityLib___________9063268959e590600160a060020a03168460405160e060020a63ffffffff8516028152600160a060020a0392831660048201529116602482015260440160006040518083038186803b151561084657600080fd5b6102c65a03f4151561085757600080fd5b5050505050565b600180546000918291829190600160a060020a03161515811461088057600080fd5b60015473__SecurityLib.sol:SecurityLib___________90634cbc5e9990600160a060020a03163360006040516020015260405160e060020a63ffffffff8516028152600160a060020a0392831660048201529116602482015260440160206040518083038186803b15156108f557600080fd5b6102c65a03f4151561090657600080fd5b50505060405180519050151561091b57600080fd5b60015473__GamesLib.sol:GamesLib_________________9063723747d090600160a060020a03168760006040516040015260405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401604080518083038186803b151561098d57600080fd5b6102c65a03f4151561099e57600080fd5b50505060405180519060200180519050925092508215610a375760015473__GamesLib.sol:GamesLib_________________90637f7c372490600160a060020a03168460405160e060020a63ffffffff8516028152600160a060020a039092166004830152602482015260440160006040518083038186803b1515610a2257600080fd5b6102c65a03f41515610a3357600080fd5b5050505b50909392505050565b6001805460009190600160a060020a031615158114610a5e57600080fd5b60015473__GamesLib.sol:GamesLib_________________90632ec5a38f90600160a060020a031660006040516020015260405160e060020a63ffffffff8416028152600160a060020a03909116600482015260240160206040518083038186803b1515610acb57600080fd5b6102c65a03f41515610adc57600080fd5b50505060405180519250505090565b60018054600160a060020a031615158114610b0557600080fd5b60005433600160a060020a03908116911614610b2057600080fd5b600154600160a060020a03166341c0e1b56040518163ffffffff1660e060020a028152600401600060405180830381600087803b1515610b5f57600080fd5b6102c65a03f11515610b7057600080fd5b50506001805473ffffffffffffffffffffffffffffffffffffffff19169055507f0f41df7e9c715020a9e0819fd4928247939cd858a13263fd926b1055992e81ae60405160405180910390a150565b6001805460009190600160a060020a031615158114610bdd57600080fd5b60015473__SecurityLib.sol:SecurityLib___________90634cbc5e9990600160a060020a03163360006040516020015260405160e060020a63ffffffff8516028152600160a060020a0392831660048201529116602482015260440160206040518083038186803b1515610c5257600080fd5b6102c65a03f41515610c6357600080fd5b505050604051805190501515610c7857600080fd5b60015473__GamesLib.sol:GamesLib_________________9063b1ed481c90600160a060020a0316868660006040516020015260405160e060020a63ffffffff8616028152600160a060020a039384166004820152919092166024820152604481019190915260640160206040518083038186803b1515610cf857600080fd5b6102c65a03f41515610d0957600080fd5b505050604051805195945050505050565b6001805460009190600160a060020a031615158114610d3857600080fd5b60015473__GamesLib.sol:GamesLib_________________906309c1946690600160a060020a0316338660006040516020015260405160e060020a63ffffffff8616028152600160a060020a039384166004820152919092166024820152604481019190915260640160206040518083038186803b1515610db857600080fd5b6102c65a03f41515610dc957600080fd5b5050506040518051949350505050565b6001805460009190600160a060020a031615158114610df757600080fd5b60015473__GamesLib.sol:GamesLib_________________9063153b2d2b90600160a060020a0316338660006040516020015260405160e060020a63ffffffff8616028152600160a060020a039384166004820152919092166024820152604481019190915260640160206040518083038186803b1515610db857600080fd5b60005433600160a060020a03908116911614610e9257600080fd5b600054600160a060020a0316ff5b6001805460009190600160a060020a031615158114610ebe57600080fd5b60015473__GamesLib.sol:GamesLib_________________9063d2dc359790600160a060020a0316338b8b8b8b8b8b60006040516020015260405160e060020a63ffffffff8b16028152600160a060020a03988916600482015296909716602487015260448601949094526064850192909252608484015260a483015260c482015260e48101919091526101040160206040518083038186803b1515610f6357600080fd5b6102c65a03f41515610f7457600080fd5b50505060405180519998505050505050505050565b6001805460009190600160a060020a031615158114610fa757600080fd5b60015473__GamesLib.sol:GamesLib_________________9063523f15be90600160a060020a0316338660006040516020015260405160e060020a63ffffffff8616028152600160a060020a039384166004820152919092166024820152604481019190915260640160206040518083038186803b1515610db857600080fd5b6001805460009190600160a060020a03161515811461104557600080fd5b60015473__GamesLib.sol:GamesLib_________________906391e6589690600160a060020a0316338660006040516020015260405160e060020a63ffffffff8616028152600160a060020a039384166004820152919092166024820152604481019190915260640160206040518083038186803b1515610db857600080fd5b600154600090600160a060020a0316156110de57600080fd5b60005433600160a060020a039081169116146110f957600080fd5b611101610e77565b7f54a8e6e923e48342ede4a07ca6c6a5768a55a26f13b4cbf99cdbc33ebe6bfc1560405160405180910390a150565b6001805460009190600160a060020a03161515811461114e57600080fd5b60015473__GamesLib.sol:GamesLib_________________90632f5fe1cf90600160a060020a0316338660006040516020015260405160e060020a63ffffffff8616028152600160a060020a039384166004820152919092166024820152604481019190915260640160206040518083038186803b1515610db857600080fd5b60005433600160a060020a039081169116146111e957600080fd5b600160a060020a03811615156111fe57600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b6001805460009190600160a060020a03161515811461124b57600080fd5b60015473__SecurityLib.sol:SecurityLib___________90634cbc5e9990600160a060020a03163360006040516020015260405160e060020a63ffffffff8516028152600160a060020a0392831660048201529116602482015260440160206040518083038186803b15156112c057600080fd5b6102c65a03f415156112d157600080fd5b5050506040518051905015156112e657600080fd5b60015473__GamesLib.sol:GamesLib_________________906391e6589690600160a060020a0316868660006040516020015260405160e060020a63ffffffff8616028152600160a060020a039384166004820152919092166024820152604481019190915260640160206040518083038186803b1515610cf857600080fd5b600154600090600160a060020a03161561137f57600080fd5b60005433600160a060020a0390811691161461139a57600080fd5b6113a2612c82565b604051809103906000f08015156113b857600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03928316179081905573__SecurityLib.sol:SecurityLib___________91638a19b4d29116611408611698565b60405160e060020a63ffffffff8516028152600160a060020a0392831660048201529116602482015260440160006040518083038186803b151561144b57600080fd5b6102c65a03f4151561145c57600080fd5b50506001547f6c9d76c1677b0909ce573fabdda9f1bbbc891f981a7efa98415bdcd0ff9f54119150600160a060020a0316604051600160a060020a03909116815260200160405180910390a150565b600180546000918291600160a060020a0316151581146114ca57600080fd5b60015473__GamesLib.sol:GamesLib_________________9063723747d090600160a060020a03168660006040516040015260405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401604080518083038186803b151561153c57600080fd5b6102c65a03f4151561154d57600080fd5b505050604051805190602001805190509250925050915091565b600180546000918291829190600160a060020a03161515811461158957600080fd5b60015473__SecurityLib.sol:SecurityLib___________90634cbc5e9990600160a060020a03163360006040516020015260405160e060020a63ffffffff8516028152600160a060020a0392831660048201529116602482015260440160206040518083038186803b15156115fe57600080fd5b6102c65a03f4151561160f57600080fd5b50505060405180519050151561162457600080fd5b60015473__GamesLib.sol:GamesLib_________________9063a838e45190600160a060020a03168760006040516040015260405160e060020a63ffffffff8516028152600160a060020a03928316600482015291166024820152604401604080518083038186803b151561098d57600080fd5b600054600160a060020a031690565b6001805460009190600160a060020a0316151581146116c557600080fd5b600154600160a060020a031663893d20e86000604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b151561170d57600080fd5b6102c65a03f11515610adc57600080fd5b60018054600160a060020a03161515811461173857600080fd5b60005433600160a060020a0390811691161461175357600080fd5b60015473__SecurityLib.sol:SecurityLib___________90638a19b4d290600160a060020a03168460405160e060020a63ffffffff8516028152600160a060020a0392831660048201529116602482015260440160006040518083038186803b151561084657600080fd5b6001805460009190600160a060020a0316151581146117dd57600080fd5b60015473__SecurityLib.sol:SecurityLib___________90634cbc5e9990600160a060020a03163360006040516020015260405160e060020a63ffffffff8516028152600160a060020a0392831660048201529116602482015260440160206040518083038186803b151561185257600080fd5b6102c65a03f4151561186357600080fd5b50505060405180519050151561187857600080fd5b60015473__GamesLib.sol:GamesLib_________________90638fbb2a9090600160a060020a0316868660006040516020015260405160e060020a63ffffffff8616028152600160a060020a039384166004820152919092166024820152604481019190915260640160206040518083038186803b1515610cf857600080fd5b6001805460009190600160a060020a03161515811461191657600080fd5b60015473__GamesLib.sol:GamesLib_________________90638fbb2a9090600160a060020a0316338660006040516020015260405160e060020a63ffffffff8616028152600160a060020a039384166004820152919092166024820152604481019190915260640160206040518083038186803b1515610db857600080fd5b600154600160a060020a031681565b6001805460009190600160a060020a0316151581146119c357600080fd5b60015473__GamesLib.sol:GamesLib_________________9063b1ed481c90600160a060020a0316338660006040516020015260405160e060020a63ffffffff8616028152600160a060020a039384166004820152919092166024820152604481019190915260640160206040518083038186803b1515610db857600080fd5b6001805460009190600160a060020a031615158114611a6157600080fd5b60015473__SecurityLib.sol:SecurityLib___________90634cbc5e9990600160a060020a03163360006040516020015260405160e060020a63ffffffff8516028152600160a060020a0392831660048201529116602482015260440160206040518083038186803b1515611ad657600080fd5b6102c65a03f41515611ae757600080fd5b505050604051805190501515611afc57600080fd5b60015473__GamesLib.sol:GamesLib_________________9063868562b790600160a060020a0316868660006040516020015260405160e060020a63ffffffff8616028152600160a060020a0393841660048201529183166024830152909116604482015260640160206040518083038186803b1515610cf857600080fd5b600180546000918291600160a060020a031615158114611b9a57600080fd5b60015473__GamesLib.sol:GamesLib_________________9063a838e45190600160a060020a03168660006040516040015260405160e060020a63ffffffff8516028152600160a060020a03928316600482015291166024820152604401604080518083038186803b151561153c57600080fd5b600080600080600080600080611c22612c92565b60018054600160a060020a031615158114611c3c57600080fd5b60015473__GamesLib.sol:GamesLib_________________9063c053dcda90600160a060020a03168d6000604051610240015260405160e060020a63ffffffff8516028152600160a060020a03909216600483015260248201526044016102406040518083038186803b1515611cb157600080fd5b6102c65a03f41515611cc257600080fd5b5050506040518051906020018051906020018051906020018051906020018051906020018051906020018051906020018051906020018061014001604052995099509950995099509950995099509950509193959799909294969850565b6001805460009190600160a060020a031615158114611d3e57600080fd5b60015473__SecurityLib.sol:SecurityLib___________90634cbc5e9990600160a060020a03163360006040516020015260405160e060020a63ffffffff8516028152600160a060020a0392831660048201529116602482015260440160206040518083038186803b1515611db357600080fd5b6102c65a03f41515611dc457600080fd5b505050604051805190501515611dd957600080fd5b60015473__GamesLib.sol:GamesLib_________________9063153b2d2b90600160a060020a0316868660006040516020015260405160e060020a63ffffffff8616028152600160a060020a039384166004820152919092166024820152604481019190915260640160206040518083038186803b1515610cf857600080fd5b6001805460009190600160a060020a031615158114611e7757600080fd5b60015473__GamesLib.sol:GamesLib_________________9063bc5e6df590600160a060020a0316338660006040516020015260405160e060020a63ffffffff8616028152600160a060020a039384166004820152919092166024820152604481019190915260640160206040518083038186803b1515610db857600080fd5b6001805460009190600160a060020a031615158114611f1557600080fd5b60015473__SecurityLib.sol:SecurityLib___________90634cbc5e9990600160a060020a03163360006040516020015260405160e060020a63ffffffff8516028152600160a060020a0392831660048201529116602482015260440160206040518083038186803b1515611f8a57600080fd5b6102c65a03f41515611f9b57600080fd5b505050604051805190501515611fb057600080fd5b60015473__GamesLib.sol:GamesLib_________________90632f5fe1cf90600160a060020a0316868660006040516020015260405160e060020a63ffffffff8616028152600160a060020a039384166004820152919092166024820152604481019190915260640160206040518083038186803b1515610cf857600080fd5b6001805460009190600160a060020a03161515811461204e57600080fd5b60015473__SecurityLib.sol:SecurityLib___________90634cbc5e9990600160a060020a03163360006040516020015260405160e060020a63ffffffff8516028152600160a060020a0392831660048201529116602482015260440160206040518083038186803b15156120c357600080fd5b6102c65a03f415156120d457600080fd5b5050506040518051905015156120e957600080fd5b60015473__GamesLib.sol:GamesLib_________________9063bc5e6df590600160a060020a0316868660006040516020015260405160e060020a63ffffffff8616028152600160a060020a039384166004820152919092166024820152604481019190915260640160206040518083038186803b1515610cf857600080fd5b612171612cbc565b600180546000918291600160a060020a03161515811461219057600080fd5b6001546000935073__GamesLib.sol:GamesLib_________________90632ec5a38f90600160a060020a0316856040516020015260405160e060020a63ffffffff8416028152600160a060020a03909116600482015260240160206040518083038186803b151561220057600080fd5b6102c65a03f4151561221157600080fd5b505050604051805190506040518059106122285750595b90808252806020026020018201604052509350600091505b60015473__GamesLib.sol:GamesLib_________________906359e32aeb90600160a060020a031660006040516020015260405160e060020a63ffffffff8416028152600160a060020a03909116600482015260240160206040518083038186803b15156122ad57600080fd5b6102c65a03f415156122be57600080fd5b5050506040518051905082101561238b5760015473__GamesLib.sol:GamesLib_________________9063d266e83b90600160a060020a03168460006040516020015260405160e060020a63ffffffff8516028152600160a060020a039092166004830152602482015260440160206040518083038186803b151561234257600080fd5b6102c65a03f4151561235357600080fd5b5050506040518051905015612380578184848060010195508151811061237557fe5b602090810290910101525b600190910190612240565b50505090565b6000806000806000806000806123a5612c92565b60018054600160a060020a0316151581146123bf57600080fd5b60015473__GamesLib.sol:GamesLib_________________9063f9580cd890600160a060020a03168d6000604051610240015260405160e060020a63ffffffff8516028152600160a060020a03909216600483015260248201526044016102406040518083038186803b1515611cb157600080fd5b6001805460009190600160a060020a03161515811461245257600080fd5b60015473__SecurityLib.sol:SecurityLib___________90634cbc5e9990600160a060020a03163360006040516020015260405160e060020a63ffffffff8516028152600160a060020a0392831660048201529116602482015260440160206040518083038186803b15156124c757600080fd5b6102c65a03f415156124d857600080fd5b5050506040518051905015156124ed57600080fd5b60015473__GamesLib.sol:GamesLib_________________906309c1946690600160a060020a0316868660006040516020015260405160e060020a63ffffffff8616028152600160a060020a039384166004820152919092166024820152604481019190915260640160206040518083038186803b1515610cf857600080fd5b60005433600160a060020a0390811691161461258857600080fd5b30600160a060020a031681600160a060020a031663893d20e86000604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b15156125d857600080fd5b6102c65a03f115156125e957600080fd5b50505060405180519050600160a060020a031614151561260857600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03838116919091179182905573__SecurityLib.sol:SecurityLib___________91638a19b4d2911661265b611698565b60405160e060020a63ffffffff8516028152600160a060020a0392831660048201529116602482015260440160006040518083038186803b151561269e57600080fd5b6102c65a03f415156126af57600080fd5b50506001547f4fde8bbd254ef19fbcc2cc839169e2d44ffaf6a0091074d7e180c1484709e98b9150600160a060020a0316604051600160a060020a03909116815260200160405180910390a150565b6001805460009190600160a060020a03161515811461271c57600080fd5b60015473__SecurityLib.sol:SecurityLib___________90634cbc5e9990600160a060020a03163360006040516020015260405160e060020a63ffffffff8516028152600160a060020a0392831660048201529116602482015260440160206040518083038186803b151561279157600080fd5b6102c65a03f415156127a257600080fd5b5050506040518051905015156127b757600080fd5b60015473__GamesLib.sol:GamesLib_________________9063ce28ed9290600160a060020a0316868660006040516020015260405160e060020a63ffffffff8616028152600160a060020a039384166004820152919092166024820152604481019190915260640160206040518083038186803b1515610cf857600080fd5b60018054600160a060020a03161515811461285157600080fd5b60005433600160a060020a0390811691161461286c57600080fd5b61287582612bec565b6001805473ffffffffffffffffffffffffffffffffffffffff191690557ff536614c5aed5495ea6a3a70d294685beeef2dc10128c2639c35d0228a8f545c82604051600160a060020a03909116815260200160405180910390a15050565b6001805460009190600160a060020a0316151581146128f157600080fd5b60015473__GamesLib.sol:GamesLib_________________9063ce28ed9290600160a060020a0316338660006040516020015260405160e060020a63ffffffff8616028152600160a060020a039384166004820152919092166024820152604481019190915260640160206040518083038186803b1515610db857600080fd5b600180546000918291600160a060020a03161515811461299057600080fd5b60015473__GamesLib.sol:GamesLib_________________9063a838e45190600160a060020a03163360006040516040015260405160e060020a63ffffffff8516028152600160a060020a03928316600482015291166024820152604401604080518083038186803b1515612a0457600080fd5b6102c65a03f41515612a1557600080fd5b50505060405180519060200180519050925092508215612aae5760015473__GamesLib.sol:GamesLib_________________90637f7c372490600160a060020a03168460405160e060020a63ffffffff8516028152600160a060020a039092166004830152602482015260440160006040518083038186803b1515612a9957600080fd5b6102c65a03f41515612aaa57600080fd5b5050505b505050565b6001805460009190600160a060020a031615158114612ad157600080fd5b60015473__SecurityLib.sol:SecurityLib___________90634cbc5e9990600160a060020a03163360006040516020015260405160e060020a63ffffffff8516028152600160a060020a0392831660048201529116602482015260440160206040518083038186803b1515612b4657600080fd5b6102c65a03f41515612b5757600080fd5b505050604051805190501515612b6c57600080fd5b60015473__GamesLib.sol:GamesLib_________________9063523f15be90600160a060020a0316868660006040516020015260405160e060020a63ffffffff8616028152600160a060020a039384166004820152919092166024820152604481019190915260640160206040518083038186803b1515610cf857600080fd5b60018054600160a060020a031615158114612c0657600080fd5b60005433600160a060020a03908116911614612c2157600080fd5b600154600160a060020a031663620996df8360405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401600060405180830381600087803b1515612c7157600080fd5b6102c65a03f1151561085757600080fd5b604051610b4180612ccf83390190565b610140604051908101604052600a815b60008152600019919091019060200181612ca25790505090565b6020604051908101604052600081529056006060604052341561000f57600080fd5b60008054600160a060020a033316600160a060020a0319909116179055610b068061003b6000396000f30060606040526004361061013d5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663025ec81a8114610142578063043106c01461016a5780630c55d92514610182578063124f24181461019857806317e7dd22146101ae57806325cf512d146101d85780633562fd20146101f15780633cc1635c1461020a5780633eba9ed21461022057806341c0e1b51461023b57806344bfa56e1461024e5780634c77e5ba146102db5780635a2bf25a1461030d578063620996df1461032f5780638267a9ee1461034e578063893d20e8146103645780639007127b1461037757806393fe42481461038d578063a209a29c146103a3578063a77aa49e146103b9578063ba69fcaa146103d2578063bdc963d8146103e8578063c9a52d2c146103fe578063f586606614610454575b600080fd5b341561014d57600080fd5b6101586004356104aa565b60405190815260200160405180910390f35b341561017557600080fd5b6101806004356104bc565b005b341561018d57600080fd5b610180600435610502565b34156101a357600080fd5b610180600435610537565b34156101b957600080fd5b6101c4600435610563565b604051901515815260200160405180910390f35b34156101e357600080fd5b610180600435602435610578565b34156101fc57600080fd5b6101806004356024356105a5565b341561021557600080fd5b6101806004356105d2565b341561022b57600080fd5b6101806004356024351515610605565b341561024657600080fd5b610180610640565b341561025957600080fd5b610264600435610669565b60405160208082528190810183818151815260200191508051906020019080838360005b838110156102a0578082015183820152602001610288565b50505050905090810190601f1680156102cd5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156102e657600080fd5b6102f160043561072c565b604051600160a060020a03909116815260200160405180910390f35b341561031857600080fd5b610180600435600160a060020a0360243516610747565b341561033a57600080fd5b610180600160a060020a036004351661079d565b341561035957600080fd5b6101806004356107fc565b341561036f57600080fd5b6102f1610828565b341561038257600080fd5b610158600435610838565b341561039857600080fd5b61018060043561084a565b34156103ae57600080fd5b610264600435610876565b34156103c457600080fd5b610180600435602435610902565b34156103dd57600080fd5b61018060043561092f565b34156103f357600080fd5b610158600435610961565b341561040957600080fd5b610180600480359060446024803590810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965061097395505050505050565b341561045f57600080fd5b610180600480359060446024803590810190830135806020601f820181900481020160405190810160405281815292919060208401838380828437509496506109b295505050505050565b60009081526005602052604090205490565b60005433600160a060020a039081169116146104d757600080fd5b6000908152600360205260409020805473ffffffffffffffffffffffffffffffffffffffff19169055565b60005433600160a060020a0390811691161461051d57600080fd5b6000818152600460205260408120610534916109ec565b50565b60005433600160a060020a0390811691161461055257600080fd5b600090815260056020526040812055565b60009081526006602052604090205460ff1690565b60005433600160a060020a0390811691161461059357600080fd5b60009182526005602052604090912055565b60005433600160a060020a039081169116146105c057600080fd5b60009182526001602052604090912055565b60005433600160a060020a039081169116146105ed57600080fd5b6000908152600660205260409020805460ff19169055565b60005433600160a060020a0390811691161461062057600080fd5b600091825260066020526040909120805460ff1916911515919091179055565b60005433600160a060020a0390811691161461065b57600080fd5b600054600160a060020a0316ff5b610671610a30565b6004600083600019166000191681526020019081526020016000208054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156107205780601f106106f557610100808354040283529160200191610720565b820191906000526020600020905b81548152906001019060200180831161070357829003601f168201915b50505050509050919050565b600090815260036020526040902054600160a060020a031690565b60005433600160a060020a0390811691161461076257600080fd5b600091825260036020526040909120805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03909216919091179055565b60005433600160a060020a039081169116146107b857600080fd5b600160a060020a03811615156107cd57600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60005433600160a060020a0390811691161461081757600080fd5b600090815260076020526040812055565b600054600160a060020a03165b90565b60009081526007602052604090205490565b60005433600160a060020a0390811691161461086557600080fd5b600090815260016020526040812055565b61087e610a30565b6002600083600019166000191681526020019081526020016000208054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156107205780601f106106f557610100808354040283529160200191610720565b60005433600160a060020a0390811691161461091d57600080fd5b60009182526007602052604090912055565b60005433600160a060020a0390811691161461094a57600080fd5b6000818152600260205260408120610534916109ec565b60009081526001602052604090205490565b60005433600160a060020a0390811691161461098e57600080fd5b60008281526004602052604090208180516109ad929160200190610a42565b505050565b60005433600160a060020a039081169116146109cd57600080fd5b60008281526002602052604090208180516109ad929160200190610a42565b50805460018160011615610100020316600290046000825580601f10610a125750610534565b601f0160209004906000526020600020908101906105349190610ac0565b60206040519081016040526000815290565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610a8357805160ff1916838001178555610ab0565b82800160010185558215610ab0579182015b82811115610ab0578251825591602001919060010190610a95565b50610abc929150610ac0565b5090565b61083591905b80821115610abc5760008155600101610ac65600a165627a7a7230582089078e369701338df82d3be48ed51f1ec92dc24b78d1b6f8c6cd51c593647c300029a165627a7a723058200894e1982a1b9e67ddeea8d4062ca40863c49c6ab9ce095275187a001dfb31ca0029`

// DeployXnoGames deploys a new Ethereum contract, binding an instance of XnoGames to it.
func DeployXnoGames(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *XnoGames, error) {
	parsed, err := abi.JSON(strings.NewReader(XnoGamesABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(XnoGamesBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &XnoGames{XnoGamesCaller: XnoGamesCaller{contract: contract}, XnoGamesTransactor: XnoGamesTransactor{contract: contract}}, nil
}

// XnoGames is an auto generated Go binding around an Ethereum contract.
type XnoGames struct {
	XnoGamesCaller     // Read-only binding to the contract
	XnoGamesTransactor // Write-only binding to the contract
}

// XnoGamesCaller is an auto generated read-only Go binding around an Ethereum contract.
type XnoGamesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// XnoGamesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type XnoGamesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// XnoGamesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type XnoGamesSession struct {
	Contract     *XnoGames         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// XnoGamesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type XnoGamesCallerSession struct {
	Contract *XnoGamesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// XnoGamesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type XnoGamesTransactorSession struct {
	Contract     *XnoGamesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// XnoGamesRaw is an auto generated low-level Go binding around an Ethereum contract.
type XnoGamesRaw struct {
	Contract *XnoGames // Generic contract binding to access the raw methods on
}

// XnoGamesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type XnoGamesCallerRaw struct {
	Contract *XnoGamesCaller // Generic read-only contract binding to access the raw methods on
}

// XnoGamesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type XnoGamesTransactorRaw struct {
	Contract *XnoGamesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewXnoGames creates a new instance of XnoGames, bound to a specific deployed contract.
func NewXnoGames(address common.Address, backend bind.ContractBackend) (*XnoGames, error) {
	contract, err := bindXnoGames(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &XnoGames{XnoGamesCaller: XnoGamesCaller{contract: contract}, XnoGamesTransactor: XnoGamesTransactor{contract: contract}}, nil
}

// NewXnoGamesCaller creates a new read-only instance of XnoGames, bound to a specific deployed contract.
func NewXnoGamesCaller(address common.Address, caller bind.ContractCaller) (*XnoGamesCaller, error) {
	contract, err := bindXnoGames(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &XnoGamesCaller{contract: contract}, nil
}

// NewXnoGamesTransactor creates a new write-only instance of XnoGames, bound to a specific deployed contract.
func NewXnoGamesTransactor(address common.Address, transactor bind.ContractTransactor) (*XnoGamesTransactor, error) {
	contract, err := bindXnoGames(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &XnoGamesTransactor{contract: contract}, nil
}

// bindXnoGames binds a generic wrapper to an already deployed contract.
func bindXnoGames(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(XnoGamesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_XnoGames *XnoGamesRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _XnoGames.Contract.XnoGamesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_XnoGames *XnoGamesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _XnoGames.Contract.XnoGamesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_XnoGames *XnoGamesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _XnoGames.Contract.XnoGamesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_XnoGames *XnoGamesCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _XnoGames.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_XnoGames *XnoGamesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _XnoGames.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_XnoGames *XnoGamesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _XnoGames.Contract.contract.Transact(opts, method, params...)
}

// EternalStorage is a free data retrieval call binding the contract method 0x98ff9c54.
//
// Solidity: function eternalStorage() constant returns(address)
func (_XnoGames *XnoGamesCaller) EternalStorage(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _XnoGames.contract.Call(opts, out, "eternalStorage")
	return *ret0, err
}

// EternalStorage is a free data retrieval call binding the contract method 0x98ff9c54.
//
// Solidity: function eternalStorage() constant returns(address)
func (_XnoGames *XnoGamesSession) EternalStorage() (common.Address, error) {
	return _XnoGames.Contract.EternalStorage(&_XnoGames.CallOpts)
}

// EternalStorage is a free data retrieval call binding the contract method 0x98ff9c54.
//
// Solidity: function eternalStorage() constant returns(address)
func (_XnoGames *XnoGamesCallerSession) EternalStorage() (common.Address, error) {
	return _XnoGames.Contract.EternalStorage(&_XnoGames.CallOpts)
}

// GetAllGameIDs is a free data retrieval call binding the contract method 0xbbed50aa.
//
// Solidity: function getAllGameIDs() constant returns(list uint256[])
func (_XnoGames *XnoGamesCaller) GetAllGameIDs(opts *bind.CallOpts) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _XnoGames.contract.Call(opts, out, "getAllGameIDs")
	return *ret0, err
}

// GetAllGameIDs is a free data retrieval call binding the contract method 0xbbed50aa.
//
// Solidity: function getAllGameIDs() constant returns(list uint256[])
func (_XnoGames *XnoGamesSession) GetAllGameIDs() ([]*big.Int, error) {
	return _XnoGames.Contract.GetAllGameIDs(&_XnoGames.CallOpts)
}

// GetAllGameIDs is a free data retrieval call binding the contract method 0xbbed50aa.
//
// Solidity: function getAllGameIDs() constant returns(list uint256[])
func (_XnoGames *XnoGamesCallerSession) GetAllGameIDs() ([]*big.Int, error) {
	return _XnoGames.Contract.GetAllGameIDs(&_XnoGames.CallOpts)
}

// GetGameByID is a free data retrieval call binding the contract method 0xaccebdfc.
//
// Solidity: function getGameByID(_gameID uint256) constant returns(found bool, gameAddress address, gameTitle bytes32, gamePublisher bytes32, gameDeveloper bytes32, gameRelease uint256, gamePrice uint256, gameLogo bytes32, gameGenreList bytes32[10])
func (_XnoGames *XnoGamesCaller) GetGameByID(opts *bind.CallOpts, _gameID *big.Int) (struct {
	Found         bool
	GameAddress   common.Address
	GameTitle     [32]byte
	GamePublisher [32]byte
	GameDeveloper [32]byte
	GameRelease   *big.Int
	GamePrice     *big.Int
	GameLogo      [32]byte
	GameGenreList [10][32]byte
}, error) {
	ret := new(struct {
		Found         bool
		GameAddress   common.Address
		GameTitle     [32]byte
		GamePublisher [32]byte
		GameDeveloper [32]byte
		GameRelease   *big.Int
		GamePrice     *big.Int
		GameLogo      [32]byte
		GameGenreList [10][32]byte
	})
	out := ret
	err := _XnoGames.contract.Call(opts, out, "getGameByID", _gameID)
	return *ret, err
}

// GetGameByID is a free data retrieval call binding the contract method 0xaccebdfc.
//
// Solidity: function getGameByID(_gameID uint256) constant returns(found bool, gameAddress address, gameTitle bytes32, gamePublisher bytes32, gameDeveloper bytes32, gameRelease uint256, gamePrice uint256, gameLogo bytes32, gameGenreList bytes32[10])
func (_XnoGames *XnoGamesSession) GetGameByID(_gameID *big.Int) (struct {
	Found         bool
	GameAddress   common.Address
	GameTitle     [32]byte
	GamePublisher [32]byte
	GameDeveloper [32]byte
	GameRelease   *big.Int
	GamePrice     *big.Int
	GameLogo      [32]byte
	GameGenreList [10][32]byte
}, error) {
	return _XnoGames.Contract.GetGameByID(&_XnoGames.CallOpts, _gameID)
}

// GetGameByID is a free data retrieval call binding the contract method 0xaccebdfc.
//
// Solidity: function getGameByID(_gameID uint256) constant returns(found bool, gameAddress address, gameTitle bytes32, gamePublisher bytes32, gameDeveloper bytes32, gameRelease uint256, gamePrice uint256, gameLogo bytes32, gameGenreList bytes32[10])
func (_XnoGames *XnoGamesCallerSession) GetGameByID(_gameID *big.Int) (struct {
	Found         bool
	GameAddress   common.Address
	GameTitle     [32]byte
	GamePublisher [32]byte
	GameDeveloper [32]byte
	GameRelease   *big.Int
	GamePrice     *big.Int
	GameLogo      [32]byte
	GameGenreList [10][32]byte
}, error) {
	return _XnoGames.Contract.GetGameByID(&_XnoGames.CallOpts, _gameID)
}

// GetGameByTitle is a free data retrieval call binding the contract method 0xc0f79441.
//
// Solidity: function getGameByTitle(_title bytes32) constant returns(found bool, gameID uint256, gameAddress address, gamePublisher bytes32, gameDeveloper bytes32, gameRelease uint256, gamePrice uint256, gameLogo bytes32, gameGenreList bytes32[10])
func (_XnoGames *XnoGamesCaller) GetGameByTitle(opts *bind.CallOpts, _title [32]byte) (struct {
	Found         bool
	GameID        *big.Int
	GameAddress   common.Address
	GamePublisher [32]byte
	GameDeveloper [32]byte
	GameRelease   *big.Int
	GamePrice     *big.Int
	GameLogo      [32]byte
	GameGenreList [10][32]byte
}, error) {
	ret := new(struct {
		Found         bool
		GameID        *big.Int
		GameAddress   common.Address
		GamePublisher [32]byte
		GameDeveloper [32]byte
		GameRelease   *big.Int
		GamePrice     *big.Int
		GameLogo      [32]byte
		GameGenreList [10][32]byte
	})
	out := ret
	err := _XnoGames.contract.Call(opts, out, "getGameByTitle", _title)
	return *ret, err
}

// GetGameByTitle is a free data retrieval call binding the contract method 0xc0f79441.
//
// Solidity: function getGameByTitle(_title bytes32) constant returns(found bool, gameID uint256, gameAddress address, gamePublisher bytes32, gameDeveloper bytes32, gameRelease uint256, gamePrice uint256, gameLogo bytes32, gameGenreList bytes32[10])
func (_XnoGames *XnoGamesSession) GetGameByTitle(_title [32]byte) (struct {
	Found         bool
	GameID        *big.Int
	GameAddress   common.Address
	GamePublisher [32]byte
	GameDeveloper [32]byte
	GameRelease   *big.Int
	GamePrice     *big.Int
	GameLogo      [32]byte
	GameGenreList [10][32]byte
}, error) {
	return _XnoGames.Contract.GetGameByTitle(&_XnoGames.CallOpts, _title)
}

// GetGameByTitle is a free data retrieval call binding the contract method 0xc0f79441.
//
// Solidity: function getGameByTitle(_title bytes32) constant returns(found bool, gameID uint256, gameAddress address, gamePublisher bytes32, gameDeveloper bytes32, gameRelease uint256, gamePrice uint256, gameLogo bytes32, gameGenreList bytes32[10])
func (_XnoGames *XnoGamesCallerSession) GetGameByTitle(_title [32]byte) (struct {
	Found         bool
	GameID        *big.Int
	GameAddress   common.Address
	GamePublisher [32]byte
	GameDeveloper [32]byte
	GameRelease   *big.Int
	GamePrice     *big.Int
	GameLogo      [32]byte
	GameGenreList [10][32]byte
}, error) {
	return _XnoGames.Contract.GetGameByTitle(&_XnoGames.CallOpts, _title)
}

// GetGameIDByAddress is a free data retrieval call binding the contract method 0xa9fa580c.
//
// Solidity: function getGameIDByAddress(_address address) constant returns(found bool, gameID uint256)
func (_XnoGames *XnoGamesCaller) GetGameIDByAddress(opts *bind.CallOpts, _address common.Address) (struct {
	Found  bool
	GameID *big.Int
}, error) {
	ret := new(struct {
		Found  bool
		GameID *big.Int
	})
	out := ret
	err := _XnoGames.contract.Call(opts, out, "getGameIDByAddress", _address)
	return *ret, err
}

// GetGameIDByAddress is a free data retrieval call binding the contract method 0xa9fa580c.
//
// Solidity: function getGameIDByAddress(_address address) constant returns(found bool, gameID uint256)
func (_XnoGames *XnoGamesSession) GetGameIDByAddress(_address common.Address) (struct {
	Found  bool
	GameID *big.Int
}, error) {
	return _XnoGames.Contract.GetGameIDByAddress(&_XnoGames.CallOpts, _address)
}

// GetGameIDByAddress is a free data retrieval call binding the contract method 0xa9fa580c.
//
// Solidity: function getGameIDByAddress(_address address) constant returns(found bool, gameID uint256)
func (_XnoGames *XnoGamesCallerSession) GetGameIDByAddress(_address common.Address) (struct {
	Found  bool
	GameID *big.Int
}, error) {
	return _XnoGames.Contract.GetGameIDByAddress(&_XnoGames.CallOpts, _address)
}

// GetGameIDByTitle is a free data retrieval call binding the contract method 0x7456d677.
//
// Solidity: function getGameIDByTitle(_title bytes32) constant returns(found bool, gameID uint256)
func (_XnoGames *XnoGamesCaller) GetGameIDByTitle(opts *bind.CallOpts, _title [32]byte) (struct {
	Found  bool
	GameID *big.Int
}, error) {
	ret := new(struct {
		Found  bool
		GameID *big.Int
	})
	out := ret
	err := _XnoGames.contract.Call(opts, out, "getGameIDByTitle", _title)
	return *ret, err
}

// GetGameIDByTitle is a free data retrieval call binding the contract method 0x7456d677.
//
// Solidity: function getGameIDByTitle(_title bytes32) constant returns(found bool, gameID uint256)
func (_XnoGames *XnoGamesSession) GetGameIDByTitle(_title [32]byte) (struct {
	Found  bool
	GameID *big.Int
}, error) {
	return _XnoGames.Contract.GetGameIDByTitle(&_XnoGames.CallOpts, _title)
}

// GetGameIDByTitle is a free data retrieval call binding the contract method 0x7456d677.
//
// Solidity: function getGameIDByTitle(_title bytes32) constant returns(found bool, gameID uint256)
func (_XnoGames *XnoGamesCallerSession) GetGameIDByTitle(_title [32]byte) (struct {
	Found  bool
	GameID *big.Int
}, error) {
	return _XnoGames.Contract.GetGameIDByTitle(&_XnoGames.CallOpts, _title)
}

// GetGamesCount is a free data retrieval call binding the contract method 0x17b7a476.
//
// Solidity: function getGamesCount() constant returns(uint256)
func (_XnoGames *XnoGamesCaller) GetGamesCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _XnoGames.contract.Call(opts, out, "getGamesCount")
	return *ret0, err
}

// GetGamesCount is a free data retrieval call binding the contract method 0x17b7a476.
//
// Solidity: function getGamesCount() constant returns(uint256)
func (_XnoGames *XnoGamesSession) GetGamesCount() (*big.Int, error) {
	return _XnoGames.Contract.GetGamesCount(&_XnoGames.CallOpts)
}

// GetGamesCount is a free data retrieval call binding the contract method 0x17b7a476.
//
// Solidity: function getGamesCount() constant returns(uint256)
func (_XnoGames *XnoGamesCallerSession) GetGamesCount() (*big.Int, error) {
	return _XnoGames.Contract.GetGamesCount(&_XnoGames.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_XnoGames *XnoGamesCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _XnoGames.contract.Call(opts, out, "getOwner")
	return *ret0, err
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_XnoGames *XnoGamesSession) GetOwner() (common.Address, error) {
	return _XnoGames.Contract.GetOwner(&_XnoGames.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_XnoGames *XnoGamesCallerSession) GetOwner() (common.Address, error) {
	return _XnoGames.Contract.GetOwner(&_XnoGames.CallOpts)
}

// GetStorageOwner is a free data retrieval call binding the contract method 0x8ae75977.
//
// Solidity: function getStorageOwner() constant returns(address)
func (_XnoGames *XnoGamesCaller) GetStorageOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _XnoGames.contract.Call(opts, out, "getStorageOwner")
	return *ret0, err
}

// GetStorageOwner is a free data retrieval call binding the contract method 0x8ae75977.
//
// Solidity: function getStorageOwner() constant returns(address)
func (_XnoGames *XnoGamesSession) GetStorageOwner() (common.Address, error) {
	return _XnoGames.Contract.GetStorageOwner(&_XnoGames.CallOpts)
}

// GetStorageOwner is a free data retrieval call binding the contract method 0x8ae75977.
//
// Solidity: function getStorageOwner() constant returns(address)
func (_XnoGames *XnoGamesCallerSession) GetStorageOwner() (common.Address, error) {
	return _XnoGames.Contract.GetStorageOwner(&_XnoGames.CallOpts)
}

// AddGameGenre is a paid mutator transaction binding the contract method 0xe5babda1.
//
// Solidity: function addGameGenre(_genre bytes32) returns(success bool)
func (_XnoGames *XnoGamesTransactor) AddGameGenre(opts *bind.TransactOpts, _genre [32]byte) (*types.Transaction, error) {
	return _XnoGames.contract.Transact(opts, "addGameGenre", _genre)
}

// AddGameGenre is a paid mutator transaction binding the contract method 0xe5babda1.
//
// Solidity: function addGameGenre(_genre bytes32) returns(success bool)
func (_XnoGames *XnoGamesSession) AddGameGenre(_genre [32]byte) (*types.Transaction, error) {
	return _XnoGames.Contract.AddGameGenre(&_XnoGames.TransactOpts, _genre)
}

// AddGameGenre is a paid mutator transaction binding the contract method 0xe5babda1.
//
// Solidity: function addGameGenre(_genre bytes32) returns(success bool)
func (_XnoGames *XnoGamesTransactorSession) AddGameGenre(_genre [32]byte) (*types.Transaction, error) {
	return _XnoGames.Contract.AddGameGenre(&_XnoGames.TransactOpts, _genre)
}

// AddXnoAdmin is a paid mutator transaction binding the contract method 0x928d9ef0.
//
// Solidity: function addXnoAdmin(_admin address) returns()
func (_XnoGames *XnoGamesTransactor) AddXnoAdmin(opts *bind.TransactOpts, _admin common.Address) (*types.Transaction, error) {
	return _XnoGames.contract.Transact(opts, "addXnoAdmin", _admin)
}

// AddXnoAdmin is a paid mutator transaction binding the contract method 0x928d9ef0.
//
// Solidity: function addXnoAdmin(_admin address) returns()
func (_XnoGames *XnoGamesSession) AddXnoAdmin(_admin common.Address) (*types.Transaction, error) {
	return _XnoGames.Contract.AddXnoAdmin(&_XnoGames.TransactOpts, _admin)
}

// AddXnoAdmin is a paid mutator transaction binding the contract method 0x928d9ef0.
//
// Solidity: function addXnoAdmin(_admin address) returns()
func (_XnoGames *XnoGamesTransactorSession) AddXnoAdmin(_admin common.Address) (*types.Transaction, error) {
	return _XnoGames.Contract.AddXnoAdmin(&_XnoGames.TransactOpts, _admin)
}

// AttachExistingStorage is a paid mutator transaction binding the contract method 0xc99fe8d8.
//
// Solidity: function attachExistingStorage(_storageContract address) returns()
func (_XnoGames *XnoGamesTransactor) AttachExistingStorage(opts *bind.TransactOpts, _storageContract common.Address) (*types.Transaction, error) {
	return _XnoGames.contract.Transact(opts, "attachExistingStorage", _storageContract)
}

// AttachExistingStorage is a paid mutator transaction binding the contract method 0xc99fe8d8.
//
// Solidity: function attachExistingStorage(_storageContract address) returns()
func (_XnoGames *XnoGamesSession) AttachExistingStorage(_storageContract common.Address) (*types.Transaction, error) {
	return _XnoGames.Contract.AttachExistingStorage(&_XnoGames.TransactOpts, _storageContract)
}

// AttachExistingStorage is a paid mutator transaction binding the contract method 0xc99fe8d8.
//
// Solidity: function attachExistingStorage(_storageContract address) returns()
func (_XnoGames *XnoGamesTransactorSession) AttachExistingStorage(_storageContract common.Address) (*types.Transaction, error) {
	return _XnoGames.Contract.AttachExistingStorage(&_XnoGames.TransactOpts, _storageContract)
}

// ChangeGameAddress is a paid mutator transaction binding the contract method 0xa1bd7051.
//
// Solidity: function changeGameAddress(_currentAddress address, _newAddress address) returns(success bool)
func (_XnoGames *XnoGamesTransactor) ChangeGameAddress(opts *bind.TransactOpts, _currentAddress common.Address, _newAddress common.Address) (*types.Transaction, error) {
	return _XnoGames.contract.Transact(opts, "changeGameAddress", _currentAddress, _newAddress)
}

// ChangeGameAddress is a paid mutator transaction binding the contract method 0xa1bd7051.
//
// Solidity: function changeGameAddress(_currentAddress address, _newAddress address) returns(success bool)
func (_XnoGames *XnoGamesSession) ChangeGameAddress(_currentAddress common.Address, _newAddress common.Address) (*types.Transaction, error) {
	return _XnoGames.Contract.ChangeGameAddress(&_XnoGames.TransactOpts, _currentAddress, _newAddress)
}

// ChangeGameAddress is a paid mutator transaction binding the contract method 0xa1bd7051.
//
// Solidity: function changeGameAddress(_currentAddress address, _newAddress address) returns(success bool)
func (_XnoGames *XnoGamesTransactorSession) ChangeGameAddress(_currentAddress common.Address, _newAddress common.Address) (*types.Transaction, error) {
	return _XnoGames.Contract.ChangeGameAddress(&_XnoGames.TransactOpts, _currentAddress, _newAddress)
}

// ChangeGameDeveloper is a paid mutator transaction binding the contract method 0x1b275b2a.
//
// Solidity: function changeGameDeveloper(_address address, _developer bytes32) returns(success bool)
func (_XnoGames *XnoGamesTransactor) ChangeGameDeveloper(opts *bind.TransactOpts, _address common.Address, _developer [32]byte) (*types.Transaction, error) {
	return _XnoGames.contract.Transact(opts, "changeGameDeveloper", _address, _developer)
}

// ChangeGameDeveloper is a paid mutator transaction binding the contract method 0x1b275b2a.
//
// Solidity: function changeGameDeveloper(_address address, _developer bytes32) returns(success bool)
func (_XnoGames *XnoGamesSession) ChangeGameDeveloper(_address common.Address, _developer [32]byte) (*types.Transaction, error) {
	return _XnoGames.Contract.ChangeGameDeveloper(&_XnoGames.TransactOpts, _address, _developer)
}

// ChangeGameDeveloper is a paid mutator transaction binding the contract method 0x1b275b2a.
//
// Solidity: function changeGameDeveloper(_address address, _developer bytes32) returns(success bool)
func (_XnoGames *XnoGamesTransactorSession) ChangeGameDeveloper(_address common.Address, _developer [32]byte) (*types.Transaction, error) {
	return _XnoGames.Contract.ChangeGameDeveloper(&_XnoGames.TransactOpts, _address, _developer)
}

// ChangeGameLogo is a paid mutator transaction binding the contract method 0xb59ed827.
//
// Solidity: function changeGameLogo(_address address, _imgUrl bytes32) returns(success bool)
func (_XnoGames *XnoGamesTransactor) ChangeGameLogo(opts *bind.TransactOpts, _address common.Address, _imgUrl [32]byte) (*types.Transaction, error) {
	return _XnoGames.contract.Transact(opts, "changeGameLogo", _address, _imgUrl)
}

// ChangeGameLogo is a paid mutator transaction binding the contract method 0xb59ed827.
//
// Solidity: function changeGameLogo(_address address, _imgUrl bytes32) returns(success bool)
func (_XnoGames *XnoGamesSession) ChangeGameLogo(_address common.Address, _imgUrl [32]byte) (*types.Transaction, error) {
	return _XnoGames.Contract.ChangeGameLogo(&_XnoGames.TransactOpts, _address, _imgUrl)
}

// ChangeGameLogo is a paid mutator transaction binding the contract method 0xb59ed827.
//
// Solidity: function changeGameLogo(_address address, _imgUrl bytes32) returns(success bool)
func (_XnoGames *XnoGamesTransactorSession) ChangeGameLogo(_address common.Address, _imgUrl [32]byte) (*types.Transaction, error) {
	return _XnoGames.Contract.ChangeGameLogo(&_XnoGames.TransactOpts, _address, _imgUrl)
}

// ChangeGamePrice is a paid mutator transaction binding the contract method 0x9470211e.
//
// Solidity: function changeGamePrice(_address address, _price uint256) returns(success bool)
func (_XnoGames *XnoGamesTransactor) ChangeGamePrice(opts *bind.TransactOpts, _address common.Address, _price *big.Int) (*types.Transaction, error) {
	return _XnoGames.contract.Transact(opts, "changeGamePrice", _address, _price)
}

// ChangeGamePrice is a paid mutator transaction binding the contract method 0x9470211e.
//
// Solidity: function changeGamePrice(_address address, _price uint256) returns(success bool)
func (_XnoGames *XnoGamesSession) ChangeGamePrice(_address common.Address, _price *big.Int) (*types.Transaction, error) {
	return _XnoGames.Contract.ChangeGamePrice(&_XnoGames.TransactOpts, _address, _price)
}

// ChangeGamePrice is a paid mutator transaction binding the contract method 0x9470211e.
//
// Solidity: function changeGamePrice(_address address, _price uint256) returns(success bool)
func (_XnoGames *XnoGamesTransactorSession) ChangeGamePrice(_address common.Address, _price *big.Int) (*types.Transaction, error) {
	return _XnoGames.Contract.ChangeGamePrice(&_XnoGames.TransactOpts, _address, _price)
}

// ChangeGamePublisher is a paid mutator transaction binding the contract method 0xadd46120.
//
// Solidity: function changeGamePublisher(_address address, _publisher bytes32) returns(success bool)
func (_XnoGames *XnoGamesTransactor) ChangeGamePublisher(opts *bind.TransactOpts, _address common.Address, _publisher [32]byte) (*types.Transaction, error) {
	return _XnoGames.contract.Transact(opts, "changeGamePublisher", _address, _publisher)
}

// ChangeGamePublisher is a paid mutator transaction binding the contract method 0xadd46120.
//
// Solidity: function changeGamePublisher(_address address, _publisher bytes32) returns(success bool)
func (_XnoGames *XnoGamesSession) ChangeGamePublisher(_address common.Address, _publisher [32]byte) (*types.Transaction, error) {
	return _XnoGames.Contract.ChangeGamePublisher(&_XnoGames.TransactOpts, _address, _publisher)
}

// ChangeGamePublisher is a paid mutator transaction binding the contract method 0xadd46120.
//
// Solidity: function changeGamePublisher(_address address, _publisher bytes32) returns(success bool)
func (_XnoGames *XnoGamesTransactorSession) ChangeGamePublisher(_address common.Address, _publisher [32]byte) (*types.Transaction, error) {
	return _XnoGames.Contract.ChangeGamePublisher(&_XnoGames.TransactOpts, _address, _publisher)
}

// ChangeGameReleaseDate is a paid mutator transaction binding the contract method 0xf9ea4c37.
//
// Solidity: function changeGameReleaseDate(_address address, _release uint256) returns(success bool)
func (_XnoGames *XnoGamesTransactor) ChangeGameReleaseDate(opts *bind.TransactOpts, _address common.Address, _release *big.Int) (*types.Transaction, error) {
	return _XnoGames.contract.Transact(opts, "changeGameReleaseDate", _address, _release)
}

// ChangeGameReleaseDate is a paid mutator transaction binding the contract method 0xf9ea4c37.
//
// Solidity: function changeGameReleaseDate(_address address, _release uint256) returns(success bool)
func (_XnoGames *XnoGamesSession) ChangeGameReleaseDate(_address common.Address, _release *big.Int) (*types.Transaction, error) {
	return _XnoGames.Contract.ChangeGameReleaseDate(&_XnoGames.TransactOpts, _address, _release)
}

// ChangeGameReleaseDate is a paid mutator transaction binding the contract method 0xf9ea4c37.
//
// Solidity: function changeGameReleaseDate(_address address, _release uint256) returns(success bool)
func (_XnoGames *XnoGamesTransactorSession) ChangeGameReleaseDate(_address common.Address, _release *big.Int) (*types.Transaction, error) {
	return _XnoGames.Contract.ChangeGameReleaseDate(&_XnoGames.TransactOpts, _address, _release)
}

// ChangeGameTitle is a paid mutator transaction binding the contract method 0xc5671765.
//
// Solidity: function changeGameTitle(_address address, _title bytes32) returns(success bool)
func (_XnoGames *XnoGamesTransactor) ChangeGameTitle(opts *bind.TransactOpts, _address common.Address, _title [32]byte) (*types.Transaction, error) {
	return _XnoGames.contract.Transact(opts, "changeGameTitle", _address, _title)
}

// ChangeGameTitle is a paid mutator transaction binding the contract method 0xc5671765.
//
// Solidity: function changeGameTitle(_address address, _title bytes32) returns(success bool)
func (_XnoGames *XnoGamesSession) ChangeGameTitle(_address common.Address, _title [32]byte) (*types.Transaction, error) {
	return _XnoGames.Contract.ChangeGameTitle(&_XnoGames.TransactOpts, _address, _title)
}

// ChangeGameTitle is a paid mutator transaction binding the contract method 0xc5671765.
//
// Solidity: function changeGameTitle(_address address, _title bytes32) returns(success bool)
func (_XnoGames *XnoGamesTransactorSession) ChangeGameTitle(_address common.Address, _title [32]byte) (*types.Transaction, error) {
	return _XnoGames.Contract.ChangeGameTitle(&_XnoGames.TransactOpts, _address, _title)
}

// CreateNewStorage is a paid mutator transaction binding the contract method 0x6dddf403.
//
// Solidity: function createNewStorage() returns()
func (_XnoGames *XnoGamesTransactor) CreateNewStorage(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _XnoGames.contract.Transact(opts, "createNewStorage")
}

// CreateNewStorage is a paid mutator transaction binding the contract method 0x6dddf403.
//
// Solidity: function createNewStorage() returns()
func (_XnoGames *XnoGamesSession) CreateNewStorage() (*types.Transaction, error) {
	return _XnoGames.Contract.CreateNewStorage(&_XnoGames.TransactOpts)
}

// CreateNewStorage is a paid mutator transaction binding the contract method 0x6dddf403.
//
// Solidity: function createNewStorage() returns()
func (_XnoGames *XnoGamesTransactorSession) CreateNewStorage() (*types.Transaction, error) {
	return _XnoGames.Contract.CreateNewStorage(&_XnoGames.TransactOpts)
}

// Deactivate is a paid mutator transaction binding the contract method 0x51b42b00.
//
// Solidity: function deactivate() returns()
func (_XnoGames *XnoGamesTransactor) Deactivate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _XnoGames.contract.Transact(opts, "deactivate")
}

// Deactivate is a paid mutator transaction binding the contract method 0x51b42b00.
//
// Solidity: function deactivate() returns()
func (_XnoGames *XnoGamesSession) Deactivate() (*types.Transaction, error) {
	return _XnoGames.Contract.Deactivate(&_XnoGames.TransactOpts)
}

// Deactivate is a paid mutator transaction binding the contract method 0x51b42b00.
//
// Solidity: function deactivate() returns()
func (_XnoGames *XnoGamesTransactorSession) Deactivate() (*types.Transaction, error) {
	return _XnoGames.Contract.Deactivate(&_XnoGames.TransactOpts)
}

// DeleteGameGenreByID is a paid mutator transaction binding the contract method 0x4fb17778.
//
// Solidity: function deleteGameGenreByID(_genreID uint256) returns(success bool)
func (_XnoGames *XnoGamesTransactor) DeleteGameGenreByID(opts *bind.TransactOpts, _genreID *big.Int) (*types.Transaction, error) {
	return _XnoGames.contract.Transact(opts, "deleteGameGenreByID", _genreID)
}

// DeleteGameGenreByID is a paid mutator transaction binding the contract method 0x4fb17778.
//
// Solidity: function deleteGameGenreByID(_genreID uint256) returns(success bool)
func (_XnoGames *XnoGamesSession) DeleteGameGenreByID(_genreID *big.Int) (*types.Transaction, error) {
	return _XnoGames.Contract.DeleteGameGenreByID(&_XnoGames.TransactOpts, _genreID)
}

// DeleteGameGenreByID is a paid mutator transaction binding the contract method 0x4fb17778.
//
// Solidity: function deleteGameGenreByID(_genreID uint256) returns(success bool)
func (_XnoGames *XnoGamesTransactorSession) DeleteGameGenreByID(_genreID *big.Int) (*types.Transaction, error) {
	return _XnoGames.Contract.DeleteGameGenreByID(&_XnoGames.TransactOpts, _genreID)
}

// DeleteGameGenreByName is a paid mutator transaction binding the contract method 0xaedb8d34.
//
// Solidity: function deleteGameGenreByName(_genreName bytes32) returns(success bool)
func (_XnoGames *XnoGamesTransactor) DeleteGameGenreByName(opts *bind.TransactOpts, _genreName [32]byte) (*types.Transaction, error) {
	return _XnoGames.contract.Transact(opts, "deleteGameGenreByName", _genreName)
}

// DeleteGameGenreByName is a paid mutator transaction binding the contract method 0xaedb8d34.
//
// Solidity: function deleteGameGenreByName(_genreName bytes32) returns(success bool)
func (_XnoGames *XnoGamesSession) DeleteGameGenreByName(_genreName [32]byte) (*types.Transaction, error) {
	return _XnoGames.Contract.DeleteGameGenreByName(&_XnoGames.TransactOpts, _genreName)
}

// DeleteGameGenreByName is a paid mutator transaction binding the contract method 0xaedb8d34.
//
// Solidity: function deleteGameGenreByName(_genreName bytes32) returns(success bool)
func (_XnoGames *XnoGamesTransactorSession) DeleteGameGenreByName(_genreName [32]byte) (*types.Transaction, error) {
	return _XnoGames.Contract.DeleteGameGenreByName(&_XnoGames.TransactOpts, _genreName)
}

// DeleteStorage is a paid mutator transaction binding the contract method 0x192150e6.
//
// Solidity: function deleteStorage() returns()
func (_XnoGames *XnoGamesTransactor) DeleteStorage(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _XnoGames.contract.Transact(opts, "deleteStorage")
}

// DeleteStorage is a paid mutator transaction binding the contract method 0x192150e6.
//
// Solidity: function deleteStorage() returns()
func (_XnoGames *XnoGamesSession) DeleteStorage() (*types.Transaction, error) {
	return _XnoGames.Contract.DeleteStorage(&_XnoGames.TransactOpts)
}

// DeleteStorage is a paid mutator transaction binding the contract method 0x192150e6.
//
// Solidity: function deleteStorage() returns()
func (_XnoGames *XnoGamesTransactorSession) DeleteStorage() (*types.Transaction, error) {
	return _XnoGames.Contract.DeleteStorage(&_XnoGames.TransactOpts)
}

// DeregisterGame is a paid mutator transaction binding the contract method 0xf9e1a45d.
//
// Solidity: function deregisterGame() returns()
func (_XnoGames *XnoGamesTransactor) DeregisterGame(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _XnoGames.contract.Transact(opts, "deregisterGame")
}

// DeregisterGame is a paid mutator transaction binding the contract method 0xf9e1a45d.
//
// Solidity: function deregisterGame() returns()
func (_XnoGames *XnoGamesSession) DeregisterGame() (*types.Transaction, error) {
	return _XnoGames.Contract.DeregisterGame(&_XnoGames.TransactOpts)
}

// DeregisterGame is a paid mutator transaction binding the contract method 0xf9e1a45d.
//
// Solidity: function deregisterGame() returns()
func (_XnoGames *XnoGamesTransactorSession) DeregisterGame() (*types.Transaction, error) {
	return _XnoGames.Contract.DeregisterGame(&_XnoGames.TransactOpts)
}

// DetachStorage is a paid mutator transaction binding the contract method 0xdf3db670.
//
// Solidity: function detachStorage(_newOnwer address) returns()
func (_XnoGames *XnoGamesTransactor) DetachStorage(opts *bind.TransactOpts, _newOnwer common.Address) (*types.Transaction, error) {
	return _XnoGames.contract.Transact(opts, "detachStorage", _newOnwer)
}

// DetachStorage is a paid mutator transaction binding the contract method 0xdf3db670.
//
// Solidity: function detachStorage(_newOnwer address) returns()
func (_XnoGames *XnoGamesSession) DetachStorage(_newOnwer common.Address) (*types.Transaction, error) {
	return _XnoGames.Contract.DetachStorage(&_XnoGames.TransactOpts, _newOnwer)
}

// DetachStorage is a paid mutator transaction binding the contract method 0xdf3db670.
//
// Solidity: function detachStorage(_newOnwer address) returns()
func (_XnoGames *XnoGamesTransactorSession) DetachStorage(_newOnwer common.Address) (*types.Transaction, error) {
	return _XnoGames.Contract.DetachStorage(&_XnoGames.TransactOpts, _newOnwer)
}

// InsertGameGenre is a paid mutator transaction binding the contract method 0xda307fa3.
//
// Solidity: function insertGameGenre(_address address, _genre bytes32) returns(success bool)
func (_XnoGames *XnoGamesTransactor) InsertGameGenre(opts *bind.TransactOpts, _address common.Address, _genre [32]byte) (*types.Transaction, error) {
	return _XnoGames.contract.Transact(opts, "insertGameGenre", _address, _genre)
}

// InsertGameGenre is a paid mutator transaction binding the contract method 0xda307fa3.
//
// Solidity: function insertGameGenre(_address address, _genre bytes32) returns(success bool)
func (_XnoGames *XnoGamesSession) InsertGameGenre(_address common.Address, _genre [32]byte) (*types.Transaction, error) {
	return _XnoGames.Contract.InsertGameGenre(&_XnoGames.TransactOpts, _address, _genre)
}

// InsertGameGenre is a paid mutator transaction binding the contract method 0xda307fa3.
//
// Solidity: function insertGameGenre(_address address, _genre bytes32) returns(success bool)
func (_XnoGames *XnoGamesTransactorSession) InsertGameGenre(_address common.Address, _genre [32]byte) (*types.Transaction, error) {
	return _XnoGames.Contract.InsertGameGenre(&_XnoGames.TransactOpts, _address, _genre)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_XnoGames *XnoGamesTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _XnoGames.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_XnoGames *XnoGamesSession) Kill() (*types.Transaction, error) {
	return _XnoGames.Contract.Kill(&_XnoGames.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_XnoGames *XnoGamesTransactorSession) Kill() (*types.Transaction, error) {
	return _XnoGames.Contract.Kill(&_XnoGames.TransactOpts)
}

// PassOwnership is a paid mutator transaction binding the contract method 0x620996df.
//
// Solidity: function passOwnership(_newOwner address) returns()
func (_XnoGames *XnoGamesTransactor) PassOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _XnoGames.contract.Transact(opts, "passOwnership", _newOwner)
}

// PassOwnership is a paid mutator transaction binding the contract method 0x620996df.
//
// Solidity: function passOwnership(_newOwner address) returns()
func (_XnoGames *XnoGamesSession) PassOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _XnoGames.Contract.PassOwnership(&_XnoGames.TransactOpts, _newOwner)
}

// PassOwnership is a paid mutator transaction binding the contract method 0x620996df.
//
// Solidity: function passOwnership(_newOwner address) returns()
func (_XnoGames *XnoGamesTransactorSession) PassOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _XnoGames.Contract.PassOwnership(&_XnoGames.TransactOpts, _newOwner)
}

// RegisterNewGame is a paid mutator transaction binding the contract method 0x4de67dac.
//
// Solidity: function registerNewGame(_title bytes32, _genre bytes32, _publisher bytes32, _developer bytes32, _release uint256, _price uint256) returns(success bool)
func (_XnoGames *XnoGamesTransactor) RegisterNewGame(opts *bind.TransactOpts, _title [32]byte, _genre [32]byte, _publisher [32]byte, _developer [32]byte, _release *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _XnoGames.contract.Transact(opts, "registerNewGame", _title, _genre, _publisher, _developer, _release, _price)
}

// RegisterNewGame is a paid mutator transaction binding the contract method 0x4de67dac.
//
// Solidity: function registerNewGame(_title bytes32, _genre bytes32, _publisher bytes32, _developer bytes32, _release uint256, _price uint256) returns(success bool)
func (_XnoGames *XnoGamesSession) RegisterNewGame(_title [32]byte, _genre [32]byte, _publisher [32]byte, _developer [32]byte, _release *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _XnoGames.Contract.RegisterNewGame(&_XnoGames.TransactOpts, _title, _genre, _publisher, _developer, _release, _price)
}

// RegisterNewGame is a paid mutator transaction binding the contract method 0x4de67dac.
//
// Solidity: function registerNewGame(_title bytes32, _genre bytes32, _publisher bytes32, _developer bytes32, _release uint256, _price uint256) returns(success bool)
func (_XnoGames *XnoGamesTransactorSession) RegisterNewGame(_title [32]byte, _genre [32]byte, _publisher [32]byte, _developer [32]byte, _release *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _XnoGames.Contract.RegisterNewGame(&_XnoGames.TransactOpts, _title, _genre, _publisher, _developer, _release, _price)
}

// RemoveGameByAddress is a paid mutator transaction binding the contract method 0x83804603.
//
// Solidity: function removeGameByAddress(_address address) returns(success bool)
func (_XnoGames *XnoGamesTransactor) RemoveGameByAddress(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _XnoGames.contract.Transact(opts, "removeGameByAddress", _address)
}

// RemoveGameByAddress is a paid mutator transaction binding the contract method 0x83804603.
//
// Solidity: function removeGameByAddress(_address address) returns(success bool)
func (_XnoGames *XnoGamesSession) RemoveGameByAddress(_address common.Address) (*types.Transaction, error) {
	return _XnoGames.Contract.RemoveGameByAddress(&_XnoGames.TransactOpts, _address)
}

// RemoveGameByAddress is a paid mutator transaction binding the contract method 0x83804603.
//
// Solidity: function removeGameByAddress(_address address) returns(success bool)
func (_XnoGames *XnoGamesTransactorSession) RemoveGameByAddress(_address common.Address) (*types.Transaction, error) {
	return _XnoGames.Contract.RemoveGameByAddress(&_XnoGames.TransactOpts, _address)
}

// RemoveGameByTitle is a paid mutator transaction binding the contract method 0x0bd51cc5.
//
// Solidity: function removeGameByTitle(_title bytes32) returns(success bool)
func (_XnoGames *XnoGamesTransactor) RemoveGameByTitle(opts *bind.TransactOpts, _title [32]byte) (*types.Transaction, error) {
	return _XnoGames.contract.Transact(opts, "removeGameByTitle", _title)
}

// RemoveGameByTitle is a paid mutator transaction binding the contract method 0x0bd51cc5.
//
// Solidity: function removeGameByTitle(_title bytes32) returns(success bool)
func (_XnoGames *XnoGamesSession) RemoveGameByTitle(_title [32]byte) (*types.Transaction, error) {
	return _XnoGames.Contract.RemoveGameByTitle(&_XnoGames.TransactOpts, _title)
}

// RemoveGameByTitle is a paid mutator transaction binding the contract method 0x0bd51cc5.
//
// Solidity: function removeGameByTitle(_title bytes32) returns(success bool)
func (_XnoGames *XnoGamesTransactorSession) RemoveGameByTitle(_title [32]byte) (*types.Transaction, error) {
	return _XnoGames.Contract.RemoveGameByTitle(&_XnoGames.TransactOpts, _title)
}

// RemoveGameGenreByID is a paid mutator transaction binding the contract method 0x6dd8f6cf.
//
// Solidity: function removeGameGenreByID(_address address, _genreID uint256) returns(success bool)
func (_XnoGames *XnoGamesTransactor) RemoveGameGenreByID(opts *bind.TransactOpts, _address common.Address, _genreID *big.Int) (*types.Transaction, error) {
	return _XnoGames.contract.Transact(opts, "removeGameGenreByID", _address, _genreID)
}

// RemoveGameGenreByID is a paid mutator transaction binding the contract method 0x6dd8f6cf.
//
// Solidity: function removeGameGenreByID(_address address, _genreID uint256) returns(success bool)
func (_XnoGames *XnoGamesSession) RemoveGameGenreByID(_address common.Address, _genreID *big.Int) (*types.Transaction, error) {
	return _XnoGames.Contract.RemoveGameGenreByID(&_XnoGames.TransactOpts, _address, _genreID)
}

// RemoveGameGenreByID is a paid mutator transaction binding the contract method 0x6dd8f6cf.
//
// Solidity: function removeGameGenreByID(_address address, _genreID uint256) returns(success bool)
func (_XnoGames *XnoGamesTransactorSession) RemoveGameGenreByID(_address common.Address, _genreID *big.Int) (*types.Transaction, error) {
	return _XnoGames.Contract.RemoveGameGenreByID(&_XnoGames.TransactOpts, _address, _genreID)
}

// RemoveGameGenreByName is a paid mutator transaction binding the contract method 0xb8ca9200.
//
// Solidity: function removeGameGenreByName(_address address, _genreName bytes32) returns(success bool)
func (_XnoGames *XnoGamesTransactor) RemoveGameGenreByName(opts *bind.TransactOpts, _address common.Address, _genreName [32]byte) (*types.Transaction, error) {
	return _XnoGames.contract.Transact(opts, "removeGameGenreByName", _address, _genreName)
}

// RemoveGameGenreByName is a paid mutator transaction binding the contract method 0xb8ca9200.
//
// Solidity: function removeGameGenreByName(_address address, _genreName bytes32) returns(success bool)
func (_XnoGames *XnoGamesSession) RemoveGameGenreByName(_address common.Address, _genreName [32]byte) (*types.Transaction, error) {
	return _XnoGames.Contract.RemoveGameGenreByName(&_XnoGames.TransactOpts, _address, _genreName)
}

// RemoveGameGenreByName is a paid mutator transaction binding the contract method 0xb8ca9200.
//
// Solidity: function removeGameGenreByName(_address address, _genreName bytes32) returns(success bool)
func (_XnoGames *XnoGamesTransactorSession) RemoveGameGenreByName(_address common.Address, _genreName [32]byte) (*types.Transaction, error) {
	return _XnoGames.Contract.RemoveGameGenreByName(&_XnoGames.TransactOpts, _address, _genreName)
}

// RemoveXnoAdmin is a paid mutator transaction binding the contract method 0x00ef809e.
//
// Solidity: function removeXnoAdmin(_admin address) returns()
func (_XnoGames *XnoGamesTransactor) RemoveXnoAdmin(opts *bind.TransactOpts, _admin common.Address) (*types.Transaction, error) {
	return _XnoGames.contract.Transact(opts, "removeXnoAdmin", _admin)
}

// RemoveXnoAdmin is a paid mutator transaction binding the contract method 0x00ef809e.
//
// Solidity: function removeXnoAdmin(_admin address) returns()
func (_XnoGames *XnoGamesSession) RemoveXnoAdmin(_admin common.Address) (*types.Transaction, error) {
	return _XnoGames.Contract.RemoveXnoAdmin(&_XnoGames.TransactOpts, _admin)
}

// RemoveXnoAdmin is a paid mutator transaction binding the contract method 0x00ef809e.
//
// Solidity: function removeXnoAdmin(_admin address) returns()
func (_XnoGames *XnoGamesTransactorSession) RemoveXnoAdmin(_admin common.Address) (*types.Transaction, error) {
	return _XnoGames.Contract.RemoveXnoAdmin(&_XnoGames.TransactOpts, _admin)
}

// UpdateGameDeveloper is a paid mutator transaction binding the contract method 0xa0e54675.
//
// Solidity: function updateGameDeveloper(_developer bytes32) returns(success bool)
func (_XnoGames *XnoGamesTransactor) UpdateGameDeveloper(opts *bind.TransactOpts, _developer [32]byte) (*types.Transaction, error) {
	return _XnoGames.contract.Transact(opts, "updateGameDeveloper", _developer)
}

// UpdateGameDeveloper is a paid mutator transaction binding the contract method 0xa0e54675.
//
// Solidity: function updateGameDeveloper(_developer bytes32) returns(success bool)
func (_XnoGames *XnoGamesSession) UpdateGameDeveloper(_developer [32]byte) (*types.Transaction, error) {
	return _XnoGames.Contract.UpdateGameDeveloper(&_XnoGames.TransactOpts, _developer)
}

// UpdateGameDeveloper is a paid mutator transaction binding the contract method 0xa0e54675.
//
// Solidity: function updateGameDeveloper(_developer bytes32) returns(success bool)
func (_XnoGames *XnoGamesTransactorSession) UpdateGameDeveloper(_developer [32]byte) (*types.Transaction, error) {
	return _XnoGames.Contract.UpdateGameDeveloper(&_XnoGames.TransactOpts, _developer)
}

// UpdateGameLogo is a paid mutator transaction binding the contract method 0x5ab75371.
//
// Solidity: function updateGameLogo(_imgUrl bytes32) returns(success bool)
func (_XnoGames *XnoGamesTransactor) UpdateGameLogo(opts *bind.TransactOpts, _imgUrl [32]byte) (*types.Transaction, error) {
	return _XnoGames.contract.Transact(opts, "updateGameLogo", _imgUrl)
}

// UpdateGameLogo is a paid mutator transaction binding the contract method 0x5ab75371.
//
// Solidity: function updateGameLogo(_imgUrl bytes32) returns(success bool)
func (_XnoGames *XnoGamesSession) UpdateGameLogo(_imgUrl [32]byte) (*types.Transaction, error) {
	return _XnoGames.Contract.UpdateGameLogo(&_XnoGames.TransactOpts, _imgUrl)
}

// UpdateGameLogo is a paid mutator transaction binding the contract method 0x5ab75371.
//
// Solidity: function updateGameLogo(_imgUrl bytes32) returns(success bool)
func (_XnoGames *XnoGamesTransactorSession) UpdateGameLogo(_imgUrl [32]byte) (*types.Transaction, error) {
	return _XnoGames.Contract.UpdateGameLogo(&_XnoGames.TransactOpts, _imgUrl)
}

// UpdateGamePrice is a paid mutator transaction binding the contract method 0x969d6abd.
//
// Solidity: function updateGamePrice(_price uint256) returns(success bool)
func (_XnoGames *XnoGamesTransactor) UpdateGamePrice(opts *bind.TransactOpts, _price *big.Int) (*types.Transaction, error) {
	return _XnoGames.contract.Transact(opts, "updateGamePrice", _price)
}

// UpdateGamePrice is a paid mutator transaction binding the contract method 0x969d6abd.
//
// Solidity: function updateGamePrice(_price uint256) returns(success bool)
func (_XnoGames *XnoGamesSession) UpdateGamePrice(_price *big.Int) (*types.Transaction, error) {
	return _XnoGames.Contract.UpdateGamePrice(&_XnoGames.TransactOpts, _price)
}

// UpdateGamePrice is a paid mutator transaction binding the contract method 0x969d6abd.
//
// Solidity: function updateGamePrice(_price uint256) returns(success bool)
func (_XnoGames *XnoGamesTransactorSession) UpdateGamePrice(_price *big.Int) (*types.Transaction, error) {
	return _XnoGames.Contract.UpdateGamePrice(&_XnoGames.TransactOpts, _price)
}

// UpdateGamePublisher is a paid mutator transaction binding the contract method 0x2c7d55c3.
//
// Solidity: function updateGamePublisher(_publisher bytes32) returns(success bool)
func (_XnoGames *XnoGamesTransactor) UpdateGamePublisher(opts *bind.TransactOpts, _publisher [32]byte) (*types.Transaction, error) {
	return _XnoGames.contract.Transact(opts, "updateGamePublisher", _publisher)
}

// UpdateGamePublisher is a paid mutator transaction binding the contract method 0x2c7d55c3.
//
// Solidity: function updateGamePublisher(_publisher bytes32) returns(success bool)
func (_XnoGames *XnoGamesSession) UpdateGamePublisher(_publisher [32]byte) (*types.Transaction, error) {
	return _XnoGames.Contract.UpdateGamePublisher(&_XnoGames.TransactOpts, _publisher)
}

// UpdateGamePublisher is a paid mutator transaction binding the contract method 0x2c7d55c3.
//
// Solidity: function updateGamePublisher(_publisher bytes32) returns(success bool)
func (_XnoGames *XnoGamesTransactorSession) UpdateGamePublisher(_publisher [32]byte) (*types.Transaction, error) {
	return _XnoGames.Contract.UpdateGamePublisher(&_XnoGames.TransactOpts, _publisher)
}

// UpdateGameReleaseDate is a paid mutator transaction binding the contract method 0x4edd5b6f.
//
// Solidity: function updateGameReleaseDate(_release uint256) returns(success bool)
func (_XnoGames *XnoGamesTransactor) UpdateGameReleaseDate(opts *bind.TransactOpts, _release *big.Int) (*types.Transaction, error) {
	return _XnoGames.contract.Transact(opts, "updateGameReleaseDate", _release)
}

// UpdateGameReleaseDate is a paid mutator transaction binding the contract method 0x4edd5b6f.
//
// Solidity: function updateGameReleaseDate(_release uint256) returns(success bool)
func (_XnoGames *XnoGamesSession) UpdateGameReleaseDate(_release *big.Int) (*types.Transaction, error) {
	return _XnoGames.Contract.UpdateGameReleaseDate(&_XnoGames.TransactOpts, _release)
}

// UpdateGameReleaseDate is a paid mutator transaction binding the contract method 0x4edd5b6f.
//
// Solidity: function updateGameReleaseDate(_release uint256) returns(success bool)
func (_XnoGames *XnoGamesTransactorSession) UpdateGameReleaseDate(_release *big.Int) (*types.Transaction, error) {
	return _XnoGames.Contract.UpdateGameReleaseDate(&_XnoGames.TransactOpts, _release)
}

// UpdateGameTitle is a paid mutator transaction binding the contract method 0x22068654.
//
// Solidity: function updateGameTitle(_title bytes32) returns(success bool)
func (_XnoGames *XnoGamesTransactor) UpdateGameTitle(opts *bind.TransactOpts, _title [32]byte) (*types.Transaction, error) {
	return _XnoGames.contract.Transact(opts, "updateGameTitle", _title)
}

// UpdateGameTitle is a paid mutator transaction binding the contract method 0x22068654.
//
// Solidity: function updateGameTitle(_title bytes32) returns(success bool)
func (_XnoGames *XnoGamesSession) UpdateGameTitle(_title [32]byte) (*types.Transaction, error) {
	return _XnoGames.Contract.UpdateGameTitle(&_XnoGames.TransactOpts, _title)
}

// UpdateGameTitle is a paid mutator transaction binding the contract method 0x22068654.
//
// Solidity: function updateGameTitle(_title bytes32) returns(success bool)
func (_XnoGames *XnoGamesTransactorSession) UpdateGameTitle(_title [32]byte) (*types.Transaction, error) {
	return _XnoGames.Contract.UpdateGameTitle(&_XnoGames.TransactOpts, _title)
}
