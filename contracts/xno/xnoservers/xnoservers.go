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

// ServersLibABI is the input ABI used to generate the binding from.
const ServersLibABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_storageContract\",\"type\":\"address\"},{\"name\":\"idx\",\"type\":\"uint256\"}],\"name\":\"removeServer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_storageContract\",\"type\":\"address\"},{\"name\":\"_serverID\",\"type\":\"uint256\"}],\"name\":\"getServerDetailsByID\",\"outputs\":[{\"name\":\"found\",\"type\":\"bool\"},{\"name\":\"serverAddress\",\"type\":\"address\"},{\"name\":\"serverIp\",\"type\":\"bytes32\"},{\"name\":\"serverType\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_storageContract\",\"type\":\"address\"},{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_ip\",\"type\":\"bytes32\"},{\"name\":\"_type\",\"type\":\"bytes32\"}],\"name\":\"addServer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_storageContract\",\"type\":\"address\"}],\"name\":\"getActiveServersCount\",\"outputs\":[{\"name\":\"activeServersCount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_storageContract\",\"type\":\"address\"},{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_ip\",\"type\":\"bytes32\"}],\"name\":\"updateIP\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_storageContract\",\"type\":\"address\"},{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_newAddress\",\"type\":\"address\"}],\"name\":\"updateAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_storageContract\",\"type\":\"address\"},{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"getIDByAddress\",\"outputs\":[{\"name\":\"found\",\"type\":\"bool\"},{\"name\":\"id\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_storageContract\",\"type\":\"address\"},{\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"isActive\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_storageContract\",\"type\":\"address\"},{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_type\",\"type\":\"bytes32\"}],\"name\":\"updateType\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_storageContract\",\"type\":\"address\"}],\"name\":\"getServersCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"ServerAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"ServerRemoved\",\"type\":\"event\"}]"

// ServersLibBin is the compiled bytecode used for deploying new contracts.
const ServersLibBin = `0x6060604052341561000f57600080fd5b61110e8061001e6000396000f30060606040526004361061008a5763ffffffff60e060020a6000350416631d3d4f01811461008f578063239c0663146100a85780633331f7ea146100f55780633a5ed5fa1461012957806379b0a7331461014f578063868562b71461016c578063a838e4511461018c578063d266e83b146101c0578063de14c6bf146101d7578063de44ef86146101f4575b600080fd5b6100a6600160a060020a0360043516602435610208565b005b6100bf600160a060020a03600435166024356102c9565b6040519315158452600160a060020a03909216602084015260408084019190915260608301919091526080909101905180910390f35b610115600160a060020a036004358116906024351660443560643561053f565b604051901515815260200160405180910390f35b61013d600160a060020a0360043516610872565b60405190815260200160405180910390f35b610115600160a060020a036004358116906024351660443561093f565b610115600160a060020a03600435811690602435811690604435166109e9565b6101a6600160a060020a0360043581169060243516610a7e565b604051911515825260208201526040908101905180910390f35b610115600160a060020a0360043516602435610c47565b610115600160a060020a0360043581169060243516604435610ce2565b61013d600160a060020a0360043516610d80565b81600160a060020a0316633eba9ed2826040516000805160206110a3833981519152815260118101919091526031016040518091039020600060405160e060020a63ffffffff8516028152600481019290925215156024820152604401600060405180830381600087803b151561027e57600080fd5b6102c65a03f1151561028f57600080fd5b5050507f7f1d7ebcd3d300931ae6b9c6bbb7cc141e437b1551c0902c77e5af368b574d878160405190815260200160405180910390a15050565b60008060008085600160a060020a03166317e7dd22866040516000805160206110a383398151915281526011810191909152603101604051809103902060006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b151561034557600080fd5b6102c65a03f1151561035657600080fd5b505050604051805194505083151561036d57610536565b85600160a060020a0316634c77e5ba866040516000805160206110c38339815191528152600e810191909152602e01604051809103902060006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b15156103e357600080fd5b6102c65a03f115156103f457600080fd5b5050506040518051935050600160a060020a03861663025ec81a8660405160008051602061108383398151915281526009810191909152602901604051809103902060006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b151561047557600080fd5b6102c65a03f1151561048657600080fd5b5050506040518051925050600160a060020a03861663025ec81a866040517f7365727665725f747970650000000000000000000000000000000000000000008152600b810191909152602b01604051809103902060006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b151561051957600080fd5b6102c65a03f1151561052a57600080fd5b50505060405180519150505b92959194509250565b60008061054c8686610e21565b151561055757600080fd5b610562866000610fe2565b905085600160a060020a03166325cf512d826040516000805160206110838339815191528152600981019190915260290160405180910390208660405160e060020a63ffffffff851602815260048101929092526024820152604401600060405180830381600087803b15156105d757600080fd5b6102c65a03f115156105e857600080fd5b50505085600160a060020a03166325cf512d826040516000805160206110838339815191528152600981019190915260290160405180910390208560405160e060020a63ffffffff851602815260048101929092526024820152604401600060405180830381600087803b151561065e57600080fd5b6102c65a03f1151561066f57600080fd5b50505085600160a060020a0316635a2bf25a826040516000805160206110c38339815191528152600e810191909152602e0160405180910390208760405160e060020a63ffffffff85160281526004810192909252600160a060020a03166024820152604401600060405180830381600087803b15156106ee57600080fd5b6102c65a03f115156106ff57600080fd5b50505085600160a060020a0316633eba9ed2826040516000805160206110a3833981519152815260118101919091526031016040518091039020600160405160e060020a63ffffffff8516028152600481019290925215156024820152604401600060405180830381600087803b151561077857600080fd5b6102c65a03f1151561078957600080fd5b50505061079586610d80565b8160010111156108335785600160a060020a0316633562fd206040517f736572766572735f636f756e74000000000000000000000000000000000000008152600d0160405180910390208360010160405160e060020a63ffffffff851602815260048101929092526024820152604401600060405180830381600087803b151561081e57600080fd5b6102c65a03f1151561082f57600080fd5b5050505b7f8bea6666bfd0398f355b74f6475c4640cf560269f6d7f69123f354631a2fe9a08160405190815260200160405180910390a150600195945050505050565b600080600061088084610d80565b915060009250600090505b818110156109385783600160a060020a03166317e7dd22826040516000805160206110a383398151915281526011810191909152603101604051809103902060006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b151561090957600080fd5b6102c65a03f1151561091a57600080fd5b5050506040518051905015610930576001909201915b60010161088b565b5050919050565b600080600061094e8686610a7e565b9150915081156109e05785600160a060020a03166325cf512d826040516000805160206110838339815191528152600981019190915260290160405180910390208660405160e060020a63ffffffff851602815260048101929092526024820152604401600060405180830381600087803b15156109cb57600080fd5b6102c65a03f115156109dc57600080fd5b5050505b50949350505050565b60008060006109f88686610a7e565b9150915081156109e05785600160a060020a0316635a2bf25a826040516000805160206110c38339815191528152600e810191909152602e0160405180910390208660405160e060020a63ffffffff85160281526004810192909252600160a060020a03166024820152604401600060405180830381600087803b15156109cb57600080fd5b600080600080610a8d86610d80565b9150600090505b81811015610c365785600160a060020a03166317e7dd22826040516000805160206110a383398151915281526011810191909152603101604051809103902060006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b1515610b1257600080fd5b6102c65a03f11515610b2357600080fd5b5050506040518051905015610c2e5784604051600160a060020a03919091166c01000000000000000000000000028152601401604051908190039020600160a060020a038716634c77e5ba836040516000805160206110c38339815191528152600e810191909152602e01604051809103902060006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b1515610bd557600080fd5b6102c65a03f11515610be657600080fd5b50505060405180519050604051600160a060020a03919091166c010000000000000000000000000281526014016040519081900390201415610c2e5760018193509350610c3e565b600101610a94565b600093508392505b50509250929050565b600082600160a060020a03166317e7dd22836040516000805160206110a383398151915281526011810191909152603101604051809103902060006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b1515610cbf57600080fd5b6102c65a03f11515610cd057600080fd5b50505060405180519150505b92915050565b6000806000610cf18686610a7e565b9150915081156109e05785600160a060020a03166325cf512d826040517f7365727665725f747970650000000000000000000000000000000000000000008152600b810191909152602b0160405180910390208660405160e060020a63ffffffff851602815260048101929092526024820152604401600060405180830381600087803b15156109cb57600080fd5b600081600160a060020a031663bdc963d86040517f736572766572735f636f756e74000000000000000000000000000000000000008152600d01604051809103902060006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b1515610e0157600080fd5b6102c65a03f11515610e1257600080fd5b50505060405180519392505050565b6000806000610e2f85610d80565b9150600090505b81811015610fd55784600160a060020a03166317e7dd22826040516000805160206110a383398151915281526011810191909152603101604051809103902060006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b1515610eb457600080fd5b6102c65a03f11515610ec557600080fd5b5050506040518051905015610fcd5783604051600160a060020a03919091166c01000000000000000000000000028152601401604051908190039020600160a060020a038616634c77e5ba836040516000805160206110c38339815191528152600e810191909152602e01604051809103902060006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b1515610f7757600080fd5b6102c65a03f11515610f8857600080fd5b50505060405180519050604051600160a060020a03919091166c010000000000000000000000000281526014016040519081900390201415610fcd5760009250610fda565b600101610e36565b600192505b505092915050565b805b82600160a060020a03166317e7dd22826040516000805160206110a383398151915281526011810191909152603101604051809103902060006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b151561105a57600080fd5b6102c65a03f1151561106b57600080fd5b5050506040518051905015610cdc57600101610fe456007365727665725f697000000000000000000000000000000000000000000000007365727665725f726567697374657265640000000000000000000000000000007365727665725f61646472657373000000000000000000000000000000000000a165627a7a72305820d05fc2f1628488ebc2d5fdda90e386ca2a1fe92b8ae36974de51b81f982253490029`

// DeployServersLib deploys a new Ethereum contract, binding an instance of ServersLib to it.
func DeployServersLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ServersLib, error) {
	parsed, err := abi.JSON(strings.NewReader(ServersLibABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ServersLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ServersLib{ServersLibCaller: ServersLibCaller{contract: contract}, ServersLibTransactor: ServersLibTransactor{contract: contract}}, nil
}

// ServersLib is an auto generated Go binding around an Ethereum contract.
type ServersLib struct {
	ServersLibCaller     // Read-only binding to the contract
	ServersLibTransactor // Write-only binding to the contract
}

// ServersLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type ServersLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ServersLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ServersLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ServersLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ServersLibSession struct {
	Contract     *ServersLib       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ServersLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ServersLibCallerSession struct {
	Contract *ServersLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ServersLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ServersLibTransactorSession struct {
	Contract     *ServersLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ServersLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type ServersLibRaw struct {
	Contract *ServersLib // Generic contract binding to access the raw methods on
}

// ServersLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ServersLibCallerRaw struct {
	Contract *ServersLibCaller // Generic read-only contract binding to access the raw methods on
}

// ServersLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ServersLibTransactorRaw struct {
	Contract *ServersLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewServersLib creates a new instance of ServersLib, bound to a specific deployed contract.
func NewServersLib(address common.Address, backend bind.ContractBackend) (*ServersLib, error) {
	contract, err := bindServersLib(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ServersLib{ServersLibCaller: ServersLibCaller{contract: contract}, ServersLibTransactor: ServersLibTransactor{contract: contract}}, nil
}

// NewServersLibCaller creates a new read-only instance of ServersLib, bound to a specific deployed contract.
func NewServersLibCaller(address common.Address, caller bind.ContractCaller) (*ServersLibCaller, error) {
	contract, err := bindServersLib(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &ServersLibCaller{contract: contract}, nil
}

// NewServersLibTransactor creates a new write-only instance of ServersLib, bound to a specific deployed contract.
func NewServersLibTransactor(address common.Address, transactor bind.ContractTransactor) (*ServersLibTransactor, error) {
	contract, err := bindServersLib(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &ServersLibTransactor{contract: contract}, nil
}

// bindServersLib binds a generic wrapper to an already deployed contract.
func bindServersLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ServersLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ServersLib *ServersLibRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ServersLib.Contract.ServersLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ServersLib *ServersLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ServersLib.Contract.ServersLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ServersLib *ServersLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ServersLib.Contract.ServersLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ServersLib *ServersLibCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ServersLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ServersLib *ServersLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ServersLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ServersLib *ServersLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ServersLib.Contract.contract.Transact(opts, method, params...)
}

// GetActiveServersCount is a free data retrieval call binding the contract method 0x3a5ed5fa.
//
// Solidity: function getActiveServersCount(_storageContract address) constant returns(activeServersCount uint256)
func (_ServersLib *ServersLibCaller) GetActiveServersCount(opts *bind.CallOpts, _storageContract common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ServersLib.contract.Call(opts, out, "getActiveServersCount", _storageContract)
	return *ret0, err
}

// GetActiveServersCount is a free data retrieval call binding the contract method 0x3a5ed5fa.
//
// Solidity: function getActiveServersCount(_storageContract address) constant returns(activeServersCount uint256)
func (_ServersLib *ServersLibSession) GetActiveServersCount(_storageContract common.Address) (*big.Int, error) {
	return _ServersLib.Contract.GetActiveServersCount(&_ServersLib.CallOpts, _storageContract)
}

// GetActiveServersCount is a free data retrieval call binding the contract method 0x3a5ed5fa.
//
// Solidity: function getActiveServersCount(_storageContract address) constant returns(activeServersCount uint256)
func (_ServersLib *ServersLibCallerSession) GetActiveServersCount(_storageContract common.Address) (*big.Int, error) {
	return _ServersLib.Contract.GetActiveServersCount(&_ServersLib.CallOpts, _storageContract)
}

// GetIDByAddress is a free data retrieval call binding the contract method 0xa838e451.
//
// Solidity: function getIDByAddress(_storageContract address, _address address) constant returns(found bool, id uint256)
func (_ServersLib *ServersLibCaller) GetIDByAddress(opts *bind.CallOpts, _storageContract common.Address, _address common.Address) (struct {
	Found bool
	Id    *big.Int
}, error) {
	ret := new(struct {
		Found bool
		Id    *big.Int
	})
	out := ret
	err := _ServersLib.contract.Call(opts, out, "getIDByAddress", _storageContract, _address)
	return *ret, err
}

// GetIDByAddress is a free data retrieval call binding the contract method 0xa838e451.
//
// Solidity: function getIDByAddress(_storageContract address, _address address) constant returns(found bool, id uint256)
func (_ServersLib *ServersLibSession) GetIDByAddress(_storageContract common.Address, _address common.Address) (struct {
	Found bool
	Id    *big.Int
}, error) {
	return _ServersLib.Contract.GetIDByAddress(&_ServersLib.CallOpts, _storageContract, _address)
}

// GetIDByAddress is a free data retrieval call binding the contract method 0xa838e451.
//
// Solidity: function getIDByAddress(_storageContract address, _address address) constant returns(found bool, id uint256)
func (_ServersLib *ServersLibCallerSession) GetIDByAddress(_storageContract common.Address, _address common.Address) (struct {
	Found bool
	Id    *big.Int
}, error) {
	return _ServersLib.Contract.GetIDByAddress(&_ServersLib.CallOpts, _storageContract, _address)
}

// GetServerDetailsByID is a free data retrieval call binding the contract method 0x239c0663.
//
// Solidity: function getServerDetailsByID(_storageContract address, _serverID uint256) constant returns(found bool, serverAddress address, serverIp bytes32, serverType bytes32)
func (_ServersLib *ServersLibCaller) GetServerDetailsByID(opts *bind.CallOpts, _storageContract common.Address, _serverID *big.Int) (struct {
	Found         bool
	ServerAddress common.Address
	ServerIp      [32]byte
	ServerType    [32]byte
}, error) {
	ret := new(struct {
		Found         bool
		ServerAddress common.Address
		ServerIp      [32]byte
		ServerType    [32]byte
	})
	out := ret
	err := _ServersLib.contract.Call(opts, out, "getServerDetailsByID", _storageContract, _serverID)
	return *ret, err
}

// GetServerDetailsByID is a free data retrieval call binding the contract method 0x239c0663.
//
// Solidity: function getServerDetailsByID(_storageContract address, _serverID uint256) constant returns(found bool, serverAddress address, serverIp bytes32, serverType bytes32)
func (_ServersLib *ServersLibSession) GetServerDetailsByID(_storageContract common.Address, _serverID *big.Int) (struct {
	Found         bool
	ServerAddress common.Address
	ServerIp      [32]byte
	ServerType    [32]byte
}, error) {
	return _ServersLib.Contract.GetServerDetailsByID(&_ServersLib.CallOpts, _storageContract, _serverID)
}

// GetServerDetailsByID is a free data retrieval call binding the contract method 0x239c0663.
//
// Solidity: function getServerDetailsByID(_storageContract address, _serverID uint256) constant returns(found bool, serverAddress address, serverIp bytes32, serverType bytes32)
func (_ServersLib *ServersLibCallerSession) GetServerDetailsByID(_storageContract common.Address, _serverID *big.Int) (struct {
	Found         bool
	ServerAddress common.Address
	ServerIp      [32]byte
	ServerType    [32]byte
}, error) {
	return _ServersLib.Contract.GetServerDetailsByID(&_ServersLib.CallOpts, _storageContract, _serverID)
}

// GetServersCount is a free data retrieval call binding the contract method 0xde44ef86.
//
// Solidity: function getServersCount(_storageContract address) constant returns(uint256)
func (_ServersLib *ServersLibCaller) GetServersCount(opts *bind.CallOpts, _storageContract common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ServersLib.contract.Call(opts, out, "getServersCount", _storageContract)
	return *ret0, err
}

// GetServersCount is a free data retrieval call binding the contract method 0xde44ef86.
//
// Solidity: function getServersCount(_storageContract address) constant returns(uint256)
func (_ServersLib *ServersLibSession) GetServersCount(_storageContract common.Address) (*big.Int, error) {
	return _ServersLib.Contract.GetServersCount(&_ServersLib.CallOpts, _storageContract)
}

// GetServersCount is a free data retrieval call binding the contract method 0xde44ef86.
//
// Solidity: function getServersCount(_storageContract address) constant returns(uint256)
func (_ServersLib *ServersLibCallerSession) GetServersCount(_storageContract common.Address) (*big.Int, error) {
	return _ServersLib.Contract.GetServersCount(&_ServersLib.CallOpts, _storageContract)
}

// IsActive is a free data retrieval call binding the contract method 0xd266e83b.
//
// Solidity: function isActive(_storageContract address, _id uint256) constant returns(bool)
func (_ServersLib *ServersLibCaller) IsActive(opts *bind.CallOpts, _storageContract common.Address, _id *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ServersLib.contract.Call(opts, out, "isActive", _storageContract, _id)
	return *ret0, err
}

// IsActive is a free data retrieval call binding the contract method 0xd266e83b.
//
// Solidity: function isActive(_storageContract address, _id uint256) constant returns(bool)
func (_ServersLib *ServersLibSession) IsActive(_storageContract common.Address, _id *big.Int) (bool, error) {
	return _ServersLib.Contract.IsActive(&_ServersLib.CallOpts, _storageContract, _id)
}

// IsActive is a free data retrieval call binding the contract method 0xd266e83b.
//
// Solidity: function isActive(_storageContract address, _id uint256) constant returns(bool)
func (_ServersLib *ServersLibCallerSession) IsActive(_storageContract common.Address, _id *big.Int) (bool, error) {
	return _ServersLib.Contract.IsActive(&_ServersLib.CallOpts, _storageContract, _id)
}

// AddServer is a paid mutator transaction binding the contract method 0x3331f7ea.
//
// Solidity: function addServer(_storageContract address, _address address, _ip bytes32, _type bytes32) returns(success bool)
func (_ServersLib *ServersLibTransactor) AddServer(opts *bind.TransactOpts, _storageContract common.Address, _address common.Address, _ip [32]byte, _type [32]byte) (*types.Transaction, error) {
	return _ServersLib.contract.Transact(opts, "addServer", _storageContract, _address, _ip, _type)
}

// AddServer is a paid mutator transaction binding the contract method 0x3331f7ea.
//
// Solidity: function addServer(_storageContract address, _address address, _ip bytes32, _type bytes32) returns(success bool)
func (_ServersLib *ServersLibSession) AddServer(_storageContract common.Address, _address common.Address, _ip [32]byte, _type [32]byte) (*types.Transaction, error) {
	return _ServersLib.Contract.AddServer(&_ServersLib.TransactOpts, _storageContract, _address, _ip, _type)
}

// AddServer is a paid mutator transaction binding the contract method 0x3331f7ea.
//
// Solidity: function addServer(_storageContract address, _address address, _ip bytes32, _type bytes32) returns(success bool)
func (_ServersLib *ServersLibTransactorSession) AddServer(_storageContract common.Address, _address common.Address, _ip [32]byte, _type [32]byte) (*types.Transaction, error) {
	return _ServersLib.Contract.AddServer(&_ServersLib.TransactOpts, _storageContract, _address, _ip, _type)
}

// RemoveServer is a paid mutator transaction binding the contract method 0x1d3d4f01.
//
// Solidity: function removeServer(_storageContract address, idx uint256) returns()
func (_ServersLib *ServersLibTransactor) RemoveServer(opts *bind.TransactOpts, _storageContract common.Address, idx *big.Int) (*types.Transaction, error) {
	return _ServersLib.contract.Transact(opts, "removeServer", _storageContract, idx)
}

// RemoveServer is a paid mutator transaction binding the contract method 0x1d3d4f01.
//
// Solidity: function removeServer(_storageContract address, idx uint256) returns()
func (_ServersLib *ServersLibSession) RemoveServer(_storageContract common.Address, idx *big.Int) (*types.Transaction, error) {
	return _ServersLib.Contract.RemoveServer(&_ServersLib.TransactOpts, _storageContract, idx)
}

// RemoveServer is a paid mutator transaction binding the contract method 0x1d3d4f01.
//
// Solidity: function removeServer(_storageContract address, idx uint256) returns()
func (_ServersLib *ServersLibTransactorSession) RemoveServer(_storageContract common.Address, idx *big.Int) (*types.Transaction, error) {
	return _ServersLib.Contract.RemoveServer(&_ServersLib.TransactOpts, _storageContract, idx)
}

// UpdateAddress is a paid mutator transaction binding the contract method 0x868562b7.
//
// Solidity: function updateAddress(_storageContract address, _address address, _newAddress address) returns(bool)
func (_ServersLib *ServersLibTransactor) UpdateAddress(opts *bind.TransactOpts, _storageContract common.Address, _address common.Address, _newAddress common.Address) (*types.Transaction, error) {
	return _ServersLib.contract.Transact(opts, "updateAddress", _storageContract, _address, _newAddress)
}

// UpdateAddress is a paid mutator transaction binding the contract method 0x868562b7.
//
// Solidity: function updateAddress(_storageContract address, _address address, _newAddress address) returns(bool)
func (_ServersLib *ServersLibSession) UpdateAddress(_storageContract common.Address, _address common.Address, _newAddress common.Address) (*types.Transaction, error) {
	return _ServersLib.Contract.UpdateAddress(&_ServersLib.TransactOpts, _storageContract, _address, _newAddress)
}

// UpdateAddress is a paid mutator transaction binding the contract method 0x868562b7.
//
// Solidity: function updateAddress(_storageContract address, _address address, _newAddress address) returns(bool)
func (_ServersLib *ServersLibTransactorSession) UpdateAddress(_storageContract common.Address, _address common.Address, _newAddress common.Address) (*types.Transaction, error) {
	return _ServersLib.Contract.UpdateAddress(&_ServersLib.TransactOpts, _storageContract, _address, _newAddress)
}

// UpdateIP is a paid mutator transaction binding the contract method 0x79b0a733.
//
// Solidity: function updateIP(_storageContract address, _address address, _ip bytes32) returns(bool)
func (_ServersLib *ServersLibTransactor) UpdateIP(opts *bind.TransactOpts, _storageContract common.Address, _address common.Address, _ip [32]byte) (*types.Transaction, error) {
	return _ServersLib.contract.Transact(opts, "updateIP", _storageContract, _address, _ip)
}

// UpdateIP is a paid mutator transaction binding the contract method 0x79b0a733.
//
// Solidity: function updateIP(_storageContract address, _address address, _ip bytes32) returns(bool)
func (_ServersLib *ServersLibSession) UpdateIP(_storageContract common.Address, _address common.Address, _ip [32]byte) (*types.Transaction, error) {
	return _ServersLib.Contract.UpdateIP(&_ServersLib.TransactOpts, _storageContract, _address, _ip)
}

// UpdateIP is a paid mutator transaction binding the contract method 0x79b0a733.
//
// Solidity: function updateIP(_storageContract address, _address address, _ip bytes32) returns(bool)
func (_ServersLib *ServersLibTransactorSession) UpdateIP(_storageContract common.Address, _address common.Address, _ip [32]byte) (*types.Transaction, error) {
	return _ServersLib.Contract.UpdateIP(&_ServersLib.TransactOpts, _storageContract, _address, _ip)
}

// UpdateType is a paid mutator transaction binding the contract method 0xde14c6bf.
//
// Solidity: function updateType(_storageContract address, _address address, _type bytes32) returns(bool)
func (_ServersLib *ServersLibTransactor) UpdateType(opts *bind.TransactOpts, _storageContract common.Address, _address common.Address, _type [32]byte) (*types.Transaction, error) {
	return _ServersLib.contract.Transact(opts, "updateType", _storageContract, _address, _type)
}

// UpdateType is a paid mutator transaction binding the contract method 0xde14c6bf.
//
// Solidity: function updateType(_storageContract address, _address address, _type bytes32) returns(bool)
func (_ServersLib *ServersLibSession) UpdateType(_storageContract common.Address, _address common.Address, _type [32]byte) (*types.Transaction, error) {
	return _ServersLib.Contract.UpdateType(&_ServersLib.TransactOpts, _storageContract, _address, _type)
}

// UpdateType is a paid mutator transaction binding the contract method 0xde14c6bf.
//
// Solidity: function updateType(_storageContract address, _address address, _type bytes32) returns(bool)
func (_ServersLib *ServersLibTransactorSession) UpdateType(_storageContract common.Address, _address common.Address, _type [32]byte) (*types.Transaction, error) {
	return _ServersLib.Contract.UpdateType(&_ServersLib.TransactOpts, _storageContract, _address, _type)
}

// XnoServersABI is the input ABI used to generate the binding from.
const XnoServersABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"removeXnoAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deregisterServer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"removeServerByAddress\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deleteStorage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"getServerIDByAddress\",\"outputs\":[{\"name\":\"found\",\"type\":\"bool\"},{\"name\":\"serverID\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deactivate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_currentAddress\",\"type\":\"address\"},{\"name\":\"_newAddress\",\"type\":\"address\"}],\"name\":\"changeServerAddress\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"passOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"createNewStorage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_currentAddress\",\"type\":\"address\"},{\"name\":\"_ip\",\"type\":\"bytes32\"}],\"name\":\"changeServerIP\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getStorageOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_currentAddress\",\"type\":\"address\"},{\"name\":\"_type\",\"type\":\"bytes32\"}],\"name\":\"changeServerType\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"addXnoAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"eternalStorage\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAllServerIDs\",\"outputs\":[{\"name\":\"list\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_ip\",\"type\":\"bytes32\"}],\"name\":\"updateServerIP\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_serverID\",\"type\":\"uint256\"}],\"name\":\"getServerByID\",\"outputs\":[{\"name\":\"found\",\"type\":\"bool\"},{\"name\":\"serverAddress\",\"type\":\"address\"},{\"name\":\"serverIP\",\"type\":\"bytes32\"},{\"name\":\"serverType\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_storageContract\",\"type\":\"address\"}],\"name\":\"attachExistingStorage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOnwer\",\"type\":\"address\"}],\"name\":\"detachStorage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_type\",\"type\":\"bytes32\"}],\"name\":\"updateServerType\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getServersCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_ip\",\"type\":\"bytes32\"},{\"name\":\"_type\",\"type\":\"bytes32\"}],\"name\":\"registerNewServer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_storageAddress\",\"type\":\"address\"}],\"name\":\"StorageCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_storageAddress\",\"type\":\"address\"}],\"name\":\"StorageAttached\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"StorageDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"StorageDetached\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"XnoServersCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"XnoServersDeactivated\",\"type\":\"event\"}]"

// XnoServersBin is the compiled bytecode used for deploying new contracts.
const XnoServersBin = `0x6060604052341561000f57600080fd5b60008054600160a060020a03191633600160a060020a03161790557f33dbf7bb72b1375808563e61bdb37fc26a37842c561a67d2638c813cdd32a7c56100606401000000006100808102610ea61704565b604051600160a060020a03909116815260200160405180910390a161008f565b600054600160a060020a031690565b61250e8061009e6000396000f3006060604052600436106101235763ffffffff60e060020a600035041662ef809e81146101285780630f1f22551461014957806313a6bca21461015c578063192150e61461018f57806331f2a5b4146101a257806341c0e1b5146101db57806351b42b00146101ee5780635b7074cf14610201578063620996df146102265780636dddf403146102455780636f94721f14610258578063893d20e81461027a5780638ae75977146102a95780638f8800b2146102bc578063928d9ef0146102de57806398ff9c54146102fd5780639c7c91c614610310578063b0f50aa514610376578063b66598a51461038c578063c99fe8d8146103d8578063df3db670146103f7578063f0951c2714610416578063f690a3ba1461042c578063fbc85d5914610451575b600080fd5b341561013357600080fd5b610147600160a060020a036004351661046a565b005b341561015457600080fd5b610147610523565b341561016757600080fd5b61017b600160a060020a0360043516610665565b604051901515815260200160405180910390f35b341561019a57600080fd5b610147610849565b34156101ad57600080fd5b6101c1600160a060020a036004351661091d565b604051911515825260208201526040908101905180910390f35b34156101e657600080fd5b6101476109db565b34156101f957600080fd5b610147610a04565b341561020c57600080fd5b61017b600160a060020a0360043581169060243516610a6f565b341561023157600080fd5b610147600160a060020a0360043516610bc9565b341561025057600080fd5b610147610c28565b341561026357600080fd5b61017b600160a060020a0360043516602435610d6d565b341561028557600080fd5b61028d610ea6565b604051600160a060020a03909116815260200160405180910390f35b34156102b457600080fd5b61028d610eb5565b34156102c757600080fd5b61017b600160a060020a0360043516602435610f3b565b34156102e957600080fd5b610147600160a060020a0360043516611074565b341561030857600080fd5b61028d611115565b341561031b57600080fd5b610323611124565b60405160208082528190810183818151815260200191508051906020019060200280838360005b8381101561036257808201518382015260200161034a565b505050509050019250505060405180910390f35b341561038157600080fd5b61017b60043561134c565b341561039757600080fd5b6103a260043561140b565b6040519315158452600160a060020a03909216602084015260408084019190915260608301919091526080909101905180910390f35b34156103e357600080fd5b610147600160a060020a03600435166114de565b341561040257600080fd5b610147600160a060020a036004351661166f565b341561042157600080fd5b61017b60043561170b565b341561043757600080fd5b61043f6117a9565b60405190815260200160405180910390f35b341561045c57600080fd5b61017b600435602435611845565b60018054600160a060020a03161515811461048457600080fd5b60005433600160a060020a0390811691161461049f57600080fd5b60015473__SecurityLib.sol:SecurityLib___________9063268959e590600160a060020a03168460405160e060020a63ffffffff8516028152600160a060020a0392831660048201529116602482015260440160006040518083038186803b151561050b57600080fd5b6102c65a03f4151561051c57600080fd5b5050505050565b600180546000918291600160a060020a03161515811461054257600080fd5b60015473__ServersLib.sol:ServersLib_____________9063a838e45190600160a060020a03163360006040516040015260405160e060020a63ffffffff8516028152600160a060020a03928316600482015291166024820152604401604080518083038186803b15156105b657600080fd5b6102c65a03f415156105c757600080fd5b505050604051805190602001805190509250925082156106605760015473__ServersLib.sol:ServersLib_____________90631d3d4f0190600160a060020a03168460405160e060020a63ffffffff8516028152600160a060020a039092166004830152602482015260440160006040518083038186803b151561064b57600080fd5b6102c65a03f4151561065c57600080fd5b5050505b505050565b600180546000918291829190600160a060020a03161515811461068757600080fd5b60015473__SecurityLib.sol:SecurityLib___________90634cbc5e9990600160a060020a03163360006040516020015260405160e060020a63ffffffff8516028152600160a060020a0392831660048201529116602482015260440160206040518083038186803b15156106fc57600080fd5b6102c65a03f4151561070d57600080fd5b50505060405180519050151561072257600080fd5b60015473__ServersLib.sol:ServersLib_____________9063a838e45190600160a060020a03168760006040516040015260405160e060020a63ffffffff8516028152600160a060020a03928316600482015291166024820152604401604080518083038186803b151561079657600080fd5b6102c65a03f415156107a757600080fd5b505050604051805190602001805190509250925082156108405760015473__ServersLib.sol:ServersLib_____________90631d3d4f0190600160a060020a03168460405160e060020a63ffffffff8516028152600160a060020a039092166004830152602482015260440160006040518083038186803b151561082b57600080fd5b6102c65a03f4151561083c57600080fd5b5050505b50909392505050565b60018054600160a060020a03161515811461086357600080fd5b60005433600160a060020a0390811691161461087e57600080fd5b600154600160a060020a03166341c0e1b56040518163ffffffff1660e060020a028152600401600060405180830381600087803b15156108bd57600080fd5b6102c65a03f115156108ce57600080fd5b50506001805473ffffffffffffffffffffffffffffffffffffffff19169055507f0f41df7e9c715020a9e0819fd4928247939cd858a13263fd926b1055992e81ae60405160405180910390a150565b600180546000918291600160a060020a03161515811461093c57600080fd5b60015473__ServersLib.sol:ServersLib_____________9063a838e45190600160a060020a03168660006040516040015260405160e060020a63ffffffff8516028152600160a060020a03928316600482015291166024820152604401604080518083038186803b15156109b057600080fd5b6102c65a03f415156109c157600080fd5b505050604051805190602001805190509250925050915091565b60005433600160a060020a039081169116146109f657600080fd5b600054600160a060020a0316ff5b600154600090600160a060020a031615610a1d57600080fd5b60005433600160a060020a03908116911614610a3857600080fd5b610a406109db565b7fbc9f82a0937a0c0ec0bc1d73bd19b17a36a8f9ec32ca3e88e1dbeb0d80c13c3760405160405180910390a150565b6001805460009190600160a060020a031615158114610a8d57600080fd5b60015473__SecurityLib.sol:SecurityLib___________90634cbc5e9990600160a060020a03163360006040516020015260405160e060020a63ffffffff8516028152600160a060020a0392831660048201529116602482015260440160206040518083038186803b1515610b0257600080fd5b6102c65a03f41515610b1357600080fd5b505050604051805190501515610b2857600080fd5b60015473__ServersLib.sol:ServersLib_____________9063868562b790600160a060020a0316868660006040516020015260405160e060020a63ffffffff8616028152600160a060020a0393841660048201529183166024830152909116604482015260640160206040518083038186803b1515610ba757600080fd5b6102c65a03f41515610bb857600080fd5b505050604051805195945050505050565b60005433600160a060020a03908116911614610be457600080fd5b600160a060020a0381161515610bf957600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600154600090600160a060020a031615610c4157600080fd5b60005433600160a060020a03908116911614610c5c57600080fd5b610c6461197f565b604051809103906000f0801515610c7a57600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03928316179081905573__SecurityLib.sol:SecurityLib___________91638a19b4d29116610cca610ea6565b60405160e060020a63ffffffff8516028152600160a060020a0392831660048201529116602482015260440160006040518083038186803b1515610d0d57600080fd5b6102c65a03f41515610d1e57600080fd5b50506001547f6c9d76c1677b0909ce573fabdda9f1bbbc891f981a7efa98415bdcd0ff9f54119150600160a060020a0316604051600160a060020a03909116815260200160405180910390a150565b6001805460009190600160a060020a031615158114610d8b57600080fd5b60015473__SecurityLib.sol:SecurityLib___________90634cbc5e9990600160a060020a03163360006040516020015260405160e060020a63ffffffff8516028152600160a060020a0392831660048201529116602482015260440160206040518083038186803b1515610e0057600080fd5b6102c65a03f41515610e1157600080fd5b505050604051805190501515610e2657600080fd5b60015473__ServersLib.sol:ServersLib_____________906379b0a73390600160a060020a0316868660006040516020015260405160e060020a63ffffffff8616028152600160a060020a039384166004820152919092166024820152604481019190915260640160206040518083038186803b1515610ba757600080fd5b600054600160a060020a031690565b6001805460009190600160a060020a031615158114610ed357600080fd5b600154600160a060020a031663893d20e86000604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b1515610f1b57600080fd5b6102c65a03f11515610f2c57600080fd5b50505060405180519250505090565b6001805460009190600160a060020a031615158114610f5957600080fd5b60015473__SecurityLib.sol:SecurityLib___________90634cbc5e9990600160a060020a03163360006040516020015260405160e060020a63ffffffff8516028152600160a060020a0392831660048201529116602482015260440160206040518083038186803b1515610fce57600080fd5b6102c65a03f41515610fdf57600080fd5b505050604051805190501515610ff457600080fd5b60015473__ServersLib.sol:ServersLib_____________9063de14c6bf90600160a060020a0316868660006040516020015260405160e060020a63ffffffff8616028152600160a060020a039384166004820152919092166024820152604481019190915260640160206040518083038186803b1515610ba757600080fd5b60018054600160a060020a03161515811461108e57600080fd5b60005433600160a060020a039081169116146110a957600080fd5b60015473__SecurityLib.sol:SecurityLib___________90638a19b4d290600160a060020a03168460405160e060020a63ffffffff8516028152600160a060020a0392831660048201529116602482015260440160006040518083038186803b151561050b57600080fd5b600154600160a060020a031681565b61112c61198f565b600180546000918291600160a060020a03161515811461114b57600080fd5b6001546000935073__ServersLib.sol:ServersLib_____________90633a5ed5fa90600160a060020a0316856040516020015260405160e060020a63ffffffff8416028152600160a060020a03909116600482015260240160206040518083038186803b15156111bb57600080fd5b6102c65a03f415156111cc57600080fd5b505050604051805190506040518059106111e35750595b90808252806020026020018201604052509350600091505b60015473__ServersLib.sol:ServersLib_____________9063de44ef8690600160a060020a031660006040516020015260405160e060020a63ffffffff8416028152600160a060020a03909116600482015260240160206040518083038186803b151561126857600080fd5b6102c65a03f4151561127957600080fd5b505050604051805190508210156113465760015473__ServersLib.sol:ServersLib_____________9063d266e83b90600160a060020a03168460006040516020015260405160e060020a63ffffffff8516028152600160a060020a039092166004830152602482015260440160206040518083038186803b15156112fd57600080fd5b6102c65a03f4151561130e57600080fd5b505050604051805190501561133b578184848060010195508151811061133057fe5b602090810290910101525b6001909101906111fb565b50505090565b6001805460009190600160a060020a03161515811461136a57600080fd5b60015473__ServersLib.sol:ServersLib_____________906379b0a73390600160a060020a0316338660006040516020015260405160e060020a63ffffffff8616028152600160a060020a039384166004820152919092166024820152604481019190915260640160206040518083038186803b15156113ea57600080fd5b6102c65a03f415156113fb57600080fd5b5050506040518051949350505050565b60018054600091829182918291600160a060020a03161515811461142e57600080fd5b60015473__ServersLib.sol:ServersLib_____________9063239c066390600160a060020a03168860006040516080015260405160e060020a63ffffffff8516028152600160a060020a039092166004830152602482015260440160806040518083038186803b15156114a157600080fd5b6102c65a03f415156114b257600080fd5b505050604051805190602001805190602001805190602001805190509450945094509450509193509193565b60005433600160a060020a039081169116146114f957600080fd5b30600160a060020a031681600160a060020a031663893d20e86000604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b151561154957600080fd5b6102c65a03f1151561155a57600080fd5b50505060405180519050600160a060020a031614151561157957600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03838116919091179182905573__SecurityLib.sol:SecurityLib___________91638a19b4d291166115cc610ea6565b60405160e060020a63ffffffff8516028152600160a060020a0392831660048201529116602482015260440160006040518083038186803b151561160f57600080fd5b6102c65a03f4151561162057600080fd5b50506001547f4fde8bbd254ef19fbcc2cc839169e2d44ffaf6a0091074d7e180c1484709e98b9150600160a060020a0316604051600160a060020a03909116815260200160405180910390a150565b60018054600160a060020a03161515811461168957600080fd5b60005433600160a060020a039081169116146116a457600080fd5b6116ad826118e9565b6001805473ffffffffffffffffffffffffffffffffffffffff191690557ff536614c5aed5495ea6a3a70d294685beeef2dc10128c2639c35d0228a8f545c82604051600160a060020a03909116815260200160405180910390a15050565b6001805460009190600160a060020a03161515811461172957600080fd5b60015473__ServersLib.sol:ServersLib_____________9063de14c6bf90600160a060020a0316338660006040516020015260405160e060020a63ffffffff8616028152600160a060020a039384166004820152919092166024820152604481019190915260640160206040518083038186803b15156113ea57600080fd5b6001805460009190600160a060020a0316151581146117c757600080fd5b60015473__ServersLib.sol:ServersLib_____________90633a5ed5fa90600160a060020a031660006040516020015260405160e060020a63ffffffff8416028152600160a060020a03909116600482015260240160206040518083038186803b151561183457600080fd5b6102c65a03f41515610f2c57600080fd5b6001805460009190600160a060020a03161515811461186357600080fd5b60015473__ServersLib.sol:ServersLib_____________90633331f7ea90600160a060020a031633878760006040516020015260405160e060020a63ffffffff8716028152600160a060020a0394851660048201529290931660248301526044820152606481019190915260840160206040518083038186803b1515610ba757600080fd5b60018054600160a060020a03161515811461190357600080fd5b60005433600160a060020a0390811691161461191e57600080fd5b600154600160a060020a031663620996df8360405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401600060405180830381600087803b151561196e57600080fd5b6102c65a03f1151561051c57600080fd5b604051610b41806119a283390190565b6020604051908101604052600081529056006060604052341561000f57600080fd5b60008054600160a060020a033316600160a060020a0319909116179055610b068061003b6000396000f30060606040526004361061013d5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663025ec81a8114610142578063043106c01461016a5780630c55d92514610182578063124f24181461019857806317e7dd22146101ae57806325cf512d146101d85780633562fd20146101f15780633cc1635c1461020a5780633eba9ed21461022057806341c0e1b51461023b57806344bfa56e1461024e5780634c77e5ba146102db5780635a2bf25a1461030d578063620996df1461032f5780638267a9ee1461034e578063893d20e8146103645780639007127b1461037757806393fe42481461038d578063a209a29c146103a3578063a77aa49e146103b9578063ba69fcaa146103d2578063bdc963d8146103e8578063c9a52d2c146103fe578063f586606614610454575b600080fd5b341561014d57600080fd5b6101586004356104aa565b60405190815260200160405180910390f35b341561017557600080fd5b6101806004356104bc565b005b341561018d57600080fd5b610180600435610502565b34156101a357600080fd5b610180600435610537565b34156101b957600080fd5b6101c4600435610563565b604051901515815260200160405180910390f35b34156101e357600080fd5b610180600435602435610578565b34156101fc57600080fd5b6101806004356024356105a5565b341561021557600080fd5b6101806004356105d2565b341561022b57600080fd5b6101806004356024351515610605565b341561024657600080fd5b610180610640565b341561025957600080fd5b610264600435610669565b60405160208082528190810183818151815260200191508051906020019080838360005b838110156102a0578082015183820152602001610288565b50505050905090810190601f1680156102cd5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156102e657600080fd5b6102f160043561072c565b604051600160a060020a03909116815260200160405180910390f35b341561031857600080fd5b610180600435600160a060020a0360243516610747565b341561033a57600080fd5b610180600160a060020a036004351661079d565b341561035957600080fd5b6101806004356107fc565b341561036f57600080fd5b6102f1610828565b341561038257600080fd5b610158600435610838565b341561039857600080fd5b61018060043561084a565b34156103ae57600080fd5b610264600435610876565b34156103c457600080fd5b610180600435602435610902565b34156103dd57600080fd5b61018060043561092f565b34156103f357600080fd5b610158600435610961565b341561040957600080fd5b610180600480359060446024803590810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965061097395505050505050565b341561045f57600080fd5b610180600480359060446024803590810190830135806020601f820181900481020160405190810160405281815292919060208401838380828437509496506109b295505050505050565b60009081526005602052604090205490565b60005433600160a060020a039081169116146104d757600080fd5b6000908152600360205260409020805473ffffffffffffffffffffffffffffffffffffffff19169055565b60005433600160a060020a0390811691161461051d57600080fd5b6000818152600460205260408120610534916109ec565b50565b60005433600160a060020a0390811691161461055257600080fd5b600090815260056020526040812055565b60009081526006602052604090205460ff1690565b60005433600160a060020a0390811691161461059357600080fd5b60009182526005602052604090912055565b60005433600160a060020a039081169116146105c057600080fd5b60009182526001602052604090912055565b60005433600160a060020a039081169116146105ed57600080fd5b6000908152600660205260409020805460ff19169055565b60005433600160a060020a0390811691161461062057600080fd5b600091825260066020526040909120805460ff1916911515919091179055565b60005433600160a060020a0390811691161461065b57600080fd5b600054600160a060020a0316ff5b610671610a30565b6004600083600019166000191681526020019081526020016000208054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156107205780601f106106f557610100808354040283529160200191610720565b820191906000526020600020905b81548152906001019060200180831161070357829003601f168201915b50505050509050919050565b600090815260036020526040902054600160a060020a031690565b60005433600160a060020a0390811691161461076257600080fd5b600091825260036020526040909120805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03909216919091179055565b60005433600160a060020a039081169116146107b857600080fd5b600160a060020a03811615156107cd57600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60005433600160a060020a0390811691161461081757600080fd5b600090815260076020526040812055565b600054600160a060020a03165b90565b60009081526007602052604090205490565b60005433600160a060020a0390811691161461086557600080fd5b600090815260016020526040812055565b61087e610a30565b6002600083600019166000191681526020019081526020016000208054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156107205780601f106106f557610100808354040283529160200191610720565b60005433600160a060020a0390811691161461091d57600080fd5b60009182526007602052604090912055565b60005433600160a060020a0390811691161461094a57600080fd5b6000818152600260205260408120610534916109ec565b60009081526001602052604090205490565b60005433600160a060020a0390811691161461098e57600080fd5b60008281526004602052604090208180516109ad929160200190610a42565b505050565b60005433600160a060020a039081169116146109cd57600080fd5b60008281526002602052604090208180516109ad929160200190610a42565b50805460018160011615610100020316600290046000825580601f10610a125750610534565b601f0160209004906000526020600020908101906105349190610ac0565b60206040519081016040526000815290565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610a8357805160ff1916838001178555610ab0565b82800160010185558215610ab0579182015b82811115610ab0578251825591602001919060010190610a95565b50610abc929150610ac0565b5090565b61083591905b80821115610abc5760008155600101610ac65600a165627a7a7230582089078e369701338df82d3be48ed51f1ec92dc24b78d1b6f8c6cd51c593647c300029a165627a7a72305820df5ff0a324223dd8ee0615f5ce09f2afd3c60f4a7d2ee9d4afe07174adc8a7e80029`

// DeployXnoServers deploys a new Ethereum contract, binding an instance of XnoServers to it.
func DeployXnoServers(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *XnoServers, error) {
	parsed, err := abi.JSON(strings.NewReader(XnoServersABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(XnoServersBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &XnoServers{XnoServersCaller: XnoServersCaller{contract: contract}, XnoServersTransactor: XnoServersTransactor{contract: contract}}, nil
}

// XnoServers is an auto generated Go binding around an Ethereum contract.
type XnoServers struct {
	XnoServersCaller     // Read-only binding to the contract
	XnoServersTransactor // Write-only binding to the contract
}

// XnoServersCaller is an auto generated read-only Go binding around an Ethereum contract.
type XnoServersCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// XnoServersTransactor is an auto generated write-only Go binding around an Ethereum contract.
type XnoServersTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// XnoServersSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type XnoServersSession struct {
	Contract     *XnoServers       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// XnoServersCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type XnoServersCallerSession struct {
	Contract *XnoServersCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// XnoServersTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type XnoServersTransactorSession struct {
	Contract     *XnoServersTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// XnoServersRaw is an auto generated low-level Go binding around an Ethereum contract.
type XnoServersRaw struct {
	Contract *XnoServers // Generic contract binding to access the raw methods on
}

// XnoServersCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type XnoServersCallerRaw struct {
	Contract *XnoServersCaller // Generic read-only contract binding to access the raw methods on
}

// XnoServersTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type XnoServersTransactorRaw struct {
	Contract *XnoServersTransactor // Generic write-only contract binding to access the raw methods on
}

// NewXnoServers creates a new instance of XnoServers, bound to a specific deployed contract.
func NewXnoServers(address common.Address, backend bind.ContractBackend) (*XnoServers, error) {
	contract, err := bindXnoServers(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &XnoServers{XnoServersCaller: XnoServersCaller{contract: contract}, XnoServersTransactor: XnoServersTransactor{contract: contract}}, nil
}

// NewXnoServersCaller creates a new read-only instance of XnoServers, bound to a specific deployed contract.
func NewXnoServersCaller(address common.Address, caller bind.ContractCaller) (*XnoServersCaller, error) {
	contract, err := bindXnoServers(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &XnoServersCaller{contract: contract}, nil
}

// NewXnoServersTransactor creates a new write-only instance of XnoServers, bound to a specific deployed contract.
func NewXnoServersTransactor(address common.Address, transactor bind.ContractTransactor) (*XnoServersTransactor, error) {
	contract, err := bindXnoServers(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &XnoServersTransactor{contract: contract}, nil
}

// bindXnoServers binds a generic wrapper to an already deployed contract.
func bindXnoServers(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(XnoServersABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_XnoServers *XnoServersRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _XnoServers.Contract.XnoServersCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_XnoServers *XnoServersRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _XnoServers.Contract.XnoServersTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_XnoServers *XnoServersRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _XnoServers.Contract.XnoServersTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_XnoServers *XnoServersCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _XnoServers.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_XnoServers *XnoServersTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _XnoServers.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_XnoServers *XnoServersTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _XnoServers.Contract.contract.Transact(opts, method, params...)
}

// EternalStorage is a free data retrieval call binding the contract method 0x98ff9c54.
//
// Solidity: function eternalStorage() constant returns(address)
func (_XnoServers *XnoServersCaller) EternalStorage(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _XnoServers.contract.Call(opts, out, "eternalStorage")
	return *ret0, err
}

// EternalStorage is a free data retrieval call binding the contract method 0x98ff9c54.
//
// Solidity: function eternalStorage() constant returns(address)
func (_XnoServers *XnoServersSession) EternalStorage() (common.Address, error) {
	return _XnoServers.Contract.EternalStorage(&_XnoServers.CallOpts)
}

// EternalStorage is a free data retrieval call binding the contract method 0x98ff9c54.
//
// Solidity: function eternalStorage() constant returns(address)
func (_XnoServers *XnoServersCallerSession) EternalStorage() (common.Address, error) {
	return _XnoServers.Contract.EternalStorage(&_XnoServers.CallOpts)
}

// GetAllServerIDs is a free data retrieval call binding the contract method 0x9c7c91c6.
//
// Solidity: function getAllServerIDs() constant returns(list uint256[])
func (_XnoServers *XnoServersCaller) GetAllServerIDs(opts *bind.CallOpts) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _XnoServers.contract.Call(opts, out, "getAllServerIDs")
	return *ret0, err
}

// GetAllServerIDs is a free data retrieval call binding the contract method 0x9c7c91c6.
//
// Solidity: function getAllServerIDs() constant returns(list uint256[])
func (_XnoServers *XnoServersSession) GetAllServerIDs() ([]*big.Int, error) {
	return _XnoServers.Contract.GetAllServerIDs(&_XnoServers.CallOpts)
}

// GetAllServerIDs is a free data retrieval call binding the contract method 0x9c7c91c6.
//
// Solidity: function getAllServerIDs() constant returns(list uint256[])
func (_XnoServers *XnoServersCallerSession) GetAllServerIDs() ([]*big.Int, error) {
	return _XnoServers.Contract.GetAllServerIDs(&_XnoServers.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_XnoServers *XnoServersCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _XnoServers.contract.Call(opts, out, "getOwner")
	return *ret0, err
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_XnoServers *XnoServersSession) GetOwner() (common.Address, error) {
	return _XnoServers.Contract.GetOwner(&_XnoServers.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_XnoServers *XnoServersCallerSession) GetOwner() (common.Address, error) {
	return _XnoServers.Contract.GetOwner(&_XnoServers.CallOpts)
}

// GetServerByID is a free data retrieval call binding the contract method 0xb66598a5.
//
// Solidity: function getServerByID(_serverID uint256) constant returns(found bool, serverAddress address, serverIP bytes32, serverType bytes32)
func (_XnoServers *XnoServersCaller) GetServerByID(opts *bind.CallOpts, _serverID *big.Int) (struct {
	Found         bool
	ServerAddress common.Address
	ServerIP      [32]byte
	ServerType    [32]byte
}, error) {
	ret := new(struct {
		Found         bool
		ServerAddress common.Address
		ServerIP      [32]byte
		ServerType    [32]byte
	})
	out := ret
	err := _XnoServers.contract.Call(opts, out, "getServerByID", _serverID)
	return *ret, err
}

// GetServerByID is a free data retrieval call binding the contract method 0xb66598a5.
//
// Solidity: function getServerByID(_serverID uint256) constant returns(found bool, serverAddress address, serverIP bytes32, serverType bytes32)
func (_XnoServers *XnoServersSession) GetServerByID(_serverID *big.Int) (struct {
	Found         bool
	ServerAddress common.Address
	ServerIP      [32]byte
	ServerType    [32]byte
}, error) {
	return _XnoServers.Contract.GetServerByID(&_XnoServers.CallOpts, _serverID)
}

// GetServerByID is a free data retrieval call binding the contract method 0xb66598a5.
//
// Solidity: function getServerByID(_serverID uint256) constant returns(found bool, serverAddress address, serverIP bytes32, serverType bytes32)
func (_XnoServers *XnoServersCallerSession) GetServerByID(_serverID *big.Int) (struct {
	Found         bool
	ServerAddress common.Address
	ServerIP      [32]byte
	ServerType    [32]byte
}, error) {
	return _XnoServers.Contract.GetServerByID(&_XnoServers.CallOpts, _serverID)
}

// GetServerIDByAddress is a free data retrieval call binding the contract method 0x31f2a5b4.
//
// Solidity: function getServerIDByAddress(_address address) constant returns(found bool, serverID uint256)
func (_XnoServers *XnoServersCaller) GetServerIDByAddress(opts *bind.CallOpts, _address common.Address) (struct {
	Found    bool
	ServerID *big.Int
}, error) {
	ret := new(struct {
		Found    bool
		ServerID *big.Int
	})
	out := ret
	err := _XnoServers.contract.Call(opts, out, "getServerIDByAddress", _address)
	return *ret, err
}

// GetServerIDByAddress is a free data retrieval call binding the contract method 0x31f2a5b4.
//
// Solidity: function getServerIDByAddress(_address address) constant returns(found bool, serverID uint256)
func (_XnoServers *XnoServersSession) GetServerIDByAddress(_address common.Address) (struct {
	Found    bool
	ServerID *big.Int
}, error) {
	return _XnoServers.Contract.GetServerIDByAddress(&_XnoServers.CallOpts, _address)
}

// GetServerIDByAddress is a free data retrieval call binding the contract method 0x31f2a5b4.
//
// Solidity: function getServerIDByAddress(_address address) constant returns(found bool, serverID uint256)
func (_XnoServers *XnoServersCallerSession) GetServerIDByAddress(_address common.Address) (struct {
	Found    bool
	ServerID *big.Int
}, error) {
	return _XnoServers.Contract.GetServerIDByAddress(&_XnoServers.CallOpts, _address)
}

// GetServersCount is a free data retrieval call binding the contract method 0xf690a3ba.
//
// Solidity: function getServersCount() constant returns(uint256)
func (_XnoServers *XnoServersCaller) GetServersCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _XnoServers.contract.Call(opts, out, "getServersCount")
	return *ret0, err
}

// GetServersCount is a free data retrieval call binding the contract method 0xf690a3ba.
//
// Solidity: function getServersCount() constant returns(uint256)
func (_XnoServers *XnoServersSession) GetServersCount() (*big.Int, error) {
	return _XnoServers.Contract.GetServersCount(&_XnoServers.CallOpts)
}

// GetServersCount is a free data retrieval call binding the contract method 0xf690a3ba.
//
// Solidity: function getServersCount() constant returns(uint256)
func (_XnoServers *XnoServersCallerSession) GetServersCount() (*big.Int, error) {
	return _XnoServers.Contract.GetServersCount(&_XnoServers.CallOpts)
}

// GetStorageOwner is a free data retrieval call binding the contract method 0x8ae75977.
//
// Solidity: function getStorageOwner() constant returns(address)
func (_XnoServers *XnoServersCaller) GetStorageOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _XnoServers.contract.Call(opts, out, "getStorageOwner")
	return *ret0, err
}

// GetStorageOwner is a free data retrieval call binding the contract method 0x8ae75977.
//
// Solidity: function getStorageOwner() constant returns(address)
func (_XnoServers *XnoServersSession) GetStorageOwner() (common.Address, error) {
	return _XnoServers.Contract.GetStorageOwner(&_XnoServers.CallOpts)
}

// GetStorageOwner is a free data retrieval call binding the contract method 0x8ae75977.
//
// Solidity: function getStorageOwner() constant returns(address)
func (_XnoServers *XnoServersCallerSession) GetStorageOwner() (common.Address, error) {
	return _XnoServers.Contract.GetStorageOwner(&_XnoServers.CallOpts)
}

// AddXnoAdmin is a paid mutator transaction binding the contract method 0x928d9ef0.
//
// Solidity: function addXnoAdmin(_admin address) returns()
func (_XnoServers *XnoServersTransactor) AddXnoAdmin(opts *bind.TransactOpts, _admin common.Address) (*types.Transaction, error) {
	return _XnoServers.contract.Transact(opts, "addXnoAdmin", _admin)
}

// AddXnoAdmin is a paid mutator transaction binding the contract method 0x928d9ef0.
//
// Solidity: function addXnoAdmin(_admin address) returns()
func (_XnoServers *XnoServersSession) AddXnoAdmin(_admin common.Address) (*types.Transaction, error) {
	return _XnoServers.Contract.AddXnoAdmin(&_XnoServers.TransactOpts, _admin)
}

// AddXnoAdmin is a paid mutator transaction binding the contract method 0x928d9ef0.
//
// Solidity: function addXnoAdmin(_admin address) returns()
func (_XnoServers *XnoServersTransactorSession) AddXnoAdmin(_admin common.Address) (*types.Transaction, error) {
	return _XnoServers.Contract.AddXnoAdmin(&_XnoServers.TransactOpts, _admin)
}

// AttachExistingStorage is a paid mutator transaction binding the contract method 0xc99fe8d8.
//
// Solidity: function attachExistingStorage(_storageContract address) returns()
func (_XnoServers *XnoServersTransactor) AttachExistingStorage(opts *bind.TransactOpts, _storageContract common.Address) (*types.Transaction, error) {
	return _XnoServers.contract.Transact(opts, "attachExistingStorage", _storageContract)
}

// AttachExistingStorage is a paid mutator transaction binding the contract method 0xc99fe8d8.
//
// Solidity: function attachExistingStorage(_storageContract address) returns()
func (_XnoServers *XnoServersSession) AttachExistingStorage(_storageContract common.Address) (*types.Transaction, error) {
	return _XnoServers.Contract.AttachExistingStorage(&_XnoServers.TransactOpts, _storageContract)
}

// AttachExistingStorage is a paid mutator transaction binding the contract method 0xc99fe8d8.
//
// Solidity: function attachExistingStorage(_storageContract address) returns()
func (_XnoServers *XnoServersTransactorSession) AttachExistingStorage(_storageContract common.Address) (*types.Transaction, error) {
	return _XnoServers.Contract.AttachExistingStorage(&_XnoServers.TransactOpts, _storageContract)
}

// ChangeServerAddress is a paid mutator transaction binding the contract method 0x5b7074cf.
//
// Solidity: function changeServerAddress(_currentAddress address, _newAddress address) returns(success bool)
func (_XnoServers *XnoServersTransactor) ChangeServerAddress(opts *bind.TransactOpts, _currentAddress common.Address, _newAddress common.Address) (*types.Transaction, error) {
	return _XnoServers.contract.Transact(opts, "changeServerAddress", _currentAddress, _newAddress)
}

// ChangeServerAddress is a paid mutator transaction binding the contract method 0x5b7074cf.
//
// Solidity: function changeServerAddress(_currentAddress address, _newAddress address) returns(success bool)
func (_XnoServers *XnoServersSession) ChangeServerAddress(_currentAddress common.Address, _newAddress common.Address) (*types.Transaction, error) {
	return _XnoServers.Contract.ChangeServerAddress(&_XnoServers.TransactOpts, _currentAddress, _newAddress)
}

// ChangeServerAddress is a paid mutator transaction binding the contract method 0x5b7074cf.
//
// Solidity: function changeServerAddress(_currentAddress address, _newAddress address) returns(success bool)
func (_XnoServers *XnoServersTransactorSession) ChangeServerAddress(_currentAddress common.Address, _newAddress common.Address) (*types.Transaction, error) {
	return _XnoServers.Contract.ChangeServerAddress(&_XnoServers.TransactOpts, _currentAddress, _newAddress)
}

// ChangeServerIP is a paid mutator transaction binding the contract method 0x6f94721f.
//
// Solidity: function changeServerIP(_currentAddress address, _ip bytes32) returns(success bool)
func (_XnoServers *XnoServersTransactor) ChangeServerIP(opts *bind.TransactOpts, _currentAddress common.Address, _ip [32]byte) (*types.Transaction, error) {
	return _XnoServers.contract.Transact(opts, "changeServerIP", _currentAddress, _ip)
}

// ChangeServerIP is a paid mutator transaction binding the contract method 0x6f94721f.
//
// Solidity: function changeServerIP(_currentAddress address, _ip bytes32) returns(success bool)
func (_XnoServers *XnoServersSession) ChangeServerIP(_currentAddress common.Address, _ip [32]byte) (*types.Transaction, error) {
	return _XnoServers.Contract.ChangeServerIP(&_XnoServers.TransactOpts, _currentAddress, _ip)
}

// ChangeServerIP is a paid mutator transaction binding the contract method 0x6f94721f.
//
// Solidity: function changeServerIP(_currentAddress address, _ip bytes32) returns(success bool)
func (_XnoServers *XnoServersTransactorSession) ChangeServerIP(_currentAddress common.Address, _ip [32]byte) (*types.Transaction, error) {
	return _XnoServers.Contract.ChangeServerIP(&_XnoServers.TransactOpts, _currentAddress, _ip)
}

// ChangeServerType is a paid mutator transaction binding the contract method 0x8f8800b2.
//
// Solidity: function changeServerType(_currentAddress address, _type bytes32) returns(success bool)
func (_XnoServers *XnoServersTransactor) ChangeServerType(opts *bind.TransactOpts, _currentAddress common.Address, _type [32]byte) (*types.Transaction, error) {
	return _XnoServers.contract.Transact(opts, "changeServerType", _currentAddress, _type)
}

// ChangeServerType is a paid mutator transaction binding the contract method 0x8f8800b2.
//
// Solidity: function changeServerType(_currentAddress address, _type bytes32) returns(success bool)
func (_XnoServers *XnoServersSession) ChangeServerType(_currentAddress common.Address, _type [32]byte) (*types.Transaction, error) {
	return _XnoServers.Contract.ChangeServerType(&_XnoServers.TransactOpts, _currentAddress, _type)
}

// ChangeServerType is a paid mutator transaction binding the contract method 0x8f8800b2.
//
// Solidity: function changeServerType(_currentAddress address, _type bytes32) returns(success bool)
func (_XnoServers *XnoServersTransactorSession) ChangeServerType(_currentAddress common.Address, _type [32]byte) (*types.Transaction, error) {
	return _XnoServers.Contract.ChangeServerType(&_XnoServers.TransactOpts, _currentAddress, _type)
}

// CreateNewStorage is a paid mutator transaction binding the contract method 0x6dddf403.
//
// Solidity: function createNewStorage() returns()
func (_XnoServers *XnoServersTransactor) CreateNewStorage(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _XnoServers.contract.Transact(opts, "createNewStorage")
}

// CreateNewStorage is a paid mutator transaction binding the contract method 0x6dddf403.
//
// Solidity: function createNewStorage() returns()
func (_XnoServers *XnoServersSession) CreateNewStorage() (*types.Transaction, error) {
	return _XnoServers.Contract.CreateNewStorage(&_XnoServers.TransactOpts)
}

// CreateNewStorage is a paid mutator transaction binding the contract method 0x6dddf403.
//
// Solidity: function createNewStorage() returns()
func (_XnoServers *XnoServersTransactorSession) CreateNewStorage() (*types.Transaction, error) {
	return _XnoServers.Contract.CreateNewStorage(&_XnoServers.TransactOpts)
}

// Deactivate is a paid mutator transaction binding the contract method 0x51b42b00.
//
// Solidity: function deactivate() returns()
func (_XnoServers *XnoServersTransactor) Deactivate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _XnoServers.contract.Transact(opts, "deactivate")
}

// Deactivate is a paid mutator transaction binding the contract method 0x51b42b00.
//
// Solidity: function deactivate() returns()
func (_XnoServers *XnoServersSession) Deactivate() (*types.Transaction, error) {
	return _XnoServers.Contract.Deactivate(&_XnoServers.TransactOpts)
}

// Deactivate is a paid mutator transaction binding the contract method 0x51b42b00.
//
// Solidity: function deactivate() returns()
func (_XnoServers *XnoServersTransactorSession) Deactivate() (*types.Transaction, error) {
	return _XnoServers.Contract.Deactivate(&_XnoServers.TransactOpts)
}

// DeleteStorage is a paid mutator transaction binding the contract method 0x192150e6.
//
// Solidity: function deleteStorage() returns()
func (_XnoServers *XnoServersTransactor) DeleteStorage(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _XnoServers.contract.Transact(opts, "deleteStorage")
}

// DeleteStorage is a paid mutator transaction binding the contract method 0x192150e6.
//
// Solidity: function deleteStorage() returns()
func (_XnoServers *XnoServersSession) DeleteStorage() (*types.Transaction, error) {
	return _XnoServers.Contract.DeleteStorage(&_XnoServers.TransactOpts)
}

// DeleteStorage is a paid mutator transaction binding the contract method 0x192150e6.
//
// Solidity: function deleteStorage() returns()
func (_XnoServers *XnoServersTransactorSession) DeleteStorage() (*types.Transaction, error) {
	return _XnoServers.Contract.DeleteStorage(&_XnoServers.TransactOpts)
}

// DeregisterServer is a paid mutator transaction binding the contract method 0x0f1f2255.
//
// Solidity: function deregisterServer() returns()
func (_XnoServers *XnoServersTransactor) DeregisterServer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _XnoServers.contract.Transact(opts, "deregisterServer")
}

// DeregisterServer is a paid mutator transaction binding the contract method 0x0f1f2255.
//
// Solidity: function deregisterServer() returns()
func (_XnoServers *XnoServersSession) DeregisterServer() (*types.Transaction, error) {
	return _XnoServers.Contract.DeregisterServer(&_XnoServers.TransactOpts)
}

// DeregisterServer is a paid mutator transaction binding the contract method 0x0f1f2255.
//
// Solidity: function deregisterServer() returns()
func (_XnoServers *XnoServersTransactorSession) DeregisterServer() (*types.Transaction, error) {
	return _XnoServers.Contract.DeregisterServer(&_XnoServers.TransactOpts)
}

// DetachStorage is a paid mutator transaction binding the contract method 0xdf3db670.
//
// Solidity: function detachStorage(_newOnwer address) returns()
func (_XnoServers *XnoServersTransactor) DetachStorage(opts *bind.TransactOpts, _newOnwer common.Address) (*types.Transaction, error) {
	return _XnoServers.contract.Transact(opts, "detachStorage", _newOnwer)
}

// DetachStorage is a paid mutator transaction binding the contract method 0xdf3db670.
//
// Solidity: function detachStorage(_newOnwer address) returns()
func (_XnoServers *XnoServersSession) DetachStorage(_newOnwer common.Address) (*types.Transaction, error) {
	return _XnoServers.Contract.DetachStorage(&_XnoServers.TransactOpts, _newOnwer)
}

// DetachStorage is a paid mutator transaction binding the contract method 0xdf3db670.
//
// Solidity: function detachStorage(_newOnwer address) returns()
func (_XnoServers *XnoServersTransactorSession) DetachStorage(_newOnwer common.Address) (*types.Transaction, error) {
	return _XnoServers.Contract.DetachStorage(&_XnoServers.TransactOpts, _newOnwer)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_XnoServers *XnoServersTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _XnoServers.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_XnoServers *XnoServersSession) Kill() (*types.Transaction, error) {
	return _XnoServers.Contract.Kill(&_XnoServers.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_XnoServers *XnoServersTransactorSession) Kill() (*types.Transaction, error) {
	return _XnoServers.Contract.Kill(&_XnoServers.TransactOpts)
}

// PassOwnership is a paid mutator transaction binding the contract method 0x620996df.
//
// Solidity: function passOwnership(_newOwner address) returns()
func (_XnoServers *XnoServersTransactor) PassOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _XnoServers.contract.Transact(opts, "passOwnership", _newOwner)
}

// PassOwnership is a paid mutator transaction binding the contract method 0x620996df.
//
// Solidity: function passOwnership(_newOwner address) returns()
func (_XnoServers *XnoServersSession) PassOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _XnoServers.Contract.PassOwnership(&_XnoServers.TransactOpts, _newOwner)
}

// PassOwnership is a paid mutator transaction binding the contract method 0x620996df.
//
// Solidity: function passOwnership(_newOwner address) returns()
func (_XnoServers *XnoServersTransactorSession) PassOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _XnoServers.Contract.PassOwnership(&_XnoServers.TransactOpts, _newOwner)
}

// RegisterNewServer is a paid mutator transaction binding the contract method 0xfbc85d59.
//
// Solidity: function registerNewServer(_ip bytes32, _type bytes32) returns(success bool)
func (_XnoServers *XnoServersTransactor) RegisterNewServer(opts *bind.TransactOpts, _ip [32]byte, _type [32]byte) (*types.Transaction, error) {
	return _XnoServers.contract.Transact(opts, "registerNewServer", _ip, _type)
}

// RegisterNewServer is a paid mutator transaction binding the contract method 0xfbc85d59.
//
// Solidity: function registerNewServer(_ip bytes32, _type bytes32) returns(success bool)
func (_XnoServers *XnoServersSession) RegisterNewServer(_ip [32]byte, _type [32]byte) (*types.Transaction, error) {
	return _XnoServers.Contract.RegisterNewServer(&_XnoServers.TransactOpts, _ip, _type)
}

// RegisterNewServer is a paid mutator transaction binding the contract method 0xfbc85d59.
//
// Solidity: function registerNewServer(_ip bytes32, _type bytes32) returns(success bool)
func (_XnoServers *XnoServersTransactorSession) RegisterNewServer(_ip [32]byte, _type [32]byte) (*types.Transaction, error) {
	return _XnoServers.Contract.RegisterNewServer(&_XnoServers.TransactOpts, _ip, _type)
}

// RemoveServerByAddress is a paid mutator transaction binding the contract method 0x13a6bca2.
//
// Solidity: function removeServerByAddress(_address address) returns(success bool)
func (_XnoServers *XnoServersTransactor) RemoveServerByAddress(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _XnoServers.contract.Transact(opts, "removeServerByAddress", _address)
}

// RemoveServerByAddress is a paid mutator transaction binding the contract method 0x13a6bca2.
//
// Solidity: function removeServerByAddress(_address address) returns(success bool)
func (_XnoServers *XnoServersSession) RemoveServerByAddress(_address common.Address) (*types.Transaction, error) {
	return _XnoServers.Contract.RemoveServerByAddress(&_XnoServers.TransactOpts, _address)
}

// RemoveServerByAddress is a paid mutator transaction binding the contract method 0x13a6bca2.
//
// Solidity: function removeServerByAddress(_address address) returns(success bool)
func (_XnoServers *XnoServersTransactorSession) RemoveServerByAddress(_address common.Address) (*types.Transaction, error) {
	return _XnoServers.Contract.RemoveServerByAddress(&_XnoServers.TransactOpts, _address)
}

// RemoveXnoAdmin is a paid mutator transaction binding the contract method 0x00ef809e.
//
// Solidity: function removeXnoAdmin(_admin address) returns()
func (_XnoServers *XnoServersTransactor) RemoveXnoAdmin(opts *bind.TransactOpts, _admin common.Address) (*types.Transaction, error) {
	return _XnoServers.contract.Transact(opts, "removeXnoAdmin", _admin)
}

// RemoveXnoAdmin is a paid mutator transaction binding the contract method 0x00ef809e.
//
// Solidity: function removeXnoAdmin(_admin address) returns()
func (_XnoServers *XnoServersSession) RemoveXnoAdmin(_admin common.Address) (*types.Transaction, error) {
	return _XnoServers.Contract.RemoveXnoAdmin(&_XnoServers.TransactOpts, _admin)
}

// RemoveXnoAdmin is a paid mutator transaction binding the contract method 0x00ef809e.
//
// Solidity: function removeXnoAdmin(_admin address) returns()
func (_XnoServers *XnoServersTransactorSession) RemoveXnoAdmin(_admin common.Address) (*types.Transaction, error) {
	return _XnoServers.Contract.RemoveXnoAdmin(&_XnoServers.TransactOpts, _admin)
}

// UpdateServerIP is a paid mutator transaction binding the contract method 0xb0f50aa5.
//
// Solidity: function updateServerIP(_ip bytes32) returns(success bool)
func (_XnoServers *XnoServersTransactor) UpdateServerIP(opts *bind.TransactOpts, _ip [32]byte) (*types.Transaction, error) {
	return _XnoServers.contract.Transact(opts, "updateServerIP", _ip)
}

// UpdateServerIP is a paid mutator transaction binding the contract method 0xb0f50aa5.
//
// Solidity: function updateServerIP(_ip bytes32) returns(success bool)
func (_XnoServers *XnoServersSession) UpdateServerIP(_ip [32]byte) (*types.Transaction, error) {
	return _XnoServers.Contract.UpdateServerIP(&_XnoServers.TransactOpts, _ip)
}

// UpdateServerIP is a paid mutator transaction binding the contract method 0xb0f50aa5.
//
// Solidity: function updateServerIP(_ip bytes32) returns(success bool)
func (_XnoServers *XnoServersTransactorSession) UpdateServerIP(_ip [32]byte) (*types.Transaction, error) {
	return _XnoServers.Contract.UpdateServerIP(&_XnoServers.TransactOpts, _ip)
}

// UpdateServerType is a paid mutator transaction binding the contract method 0xf0951c27.
//
// Solidity: function updateServerType(_type bytes32) returns(success bool)
func (_XnoServers *XnoServersTransactor) UpdateServerType(opts *bind.TransactOpts, _type [32]byte) (*types.Transaction, error) {
	return _XnoServers.contract.Transact(opts, "updateServerType", _type)
}

// UpdateServerType is a paid mutator transaction binding the contract method 0xf0951c27.
//
// Solidity: function updateServerType(_type bytes32) returns(success bool)
func (_XnoServers *XnoServersSession) UpdateServerType(_type [32]byte) (*types.Transaction, error) {
	return _XnoServers.Contract.UpdateServerType(&_XnoServers.TransactOpts, _type)
}

// UpdateServerType is a paid mutator transaction binding the contract method 0xf0951c27.
//
// Solidity: function updateServerType(_type bytes32) returns(success bool)
func (_XnoServers *XnoServersTransactorSession) UpdateServerType(_type [32]byte) (*types.Transaction, error) {
	return _XnoServers.Contract.UpdateServerType(&_XnoServers.TransactOpts, _type)
}
