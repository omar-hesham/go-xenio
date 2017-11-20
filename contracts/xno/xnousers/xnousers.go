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

// XnoUsersABI is the input ABI used to generate the binding from.
const XnoUsersABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"removeXnoAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_imgUrl\",\"type\":\"bytes32\"}],\"name\":\"changeUserProfilePicture\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deleteStorage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"removeUserByAddress\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deactivate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"passOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_userID\",\"type\":\"uint256\"}],\"name\":\"getUserByID\",\"outputs\":[{\"name\":\"found\",\"type\":\"bool\"},{\"name\":\"userAddress\",\"type\":\"address\"},{\"name\":\"userName\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"createNewStorage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deregisterUser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getStorageOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"addXnoAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"getUserIDByAddress\",\"outputs\":[{\"name\":\"found\",\"type\":\"bool\"},{\"name\":\"userID\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"eternalStorage\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getUsersCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_name\",\"type\":\"bytes32\"}],\"name\":\"registerNewUser\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_name\",\"type\":\"bytes32\"}],\"name\":\"changeUserName\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_name\",\"type\":\"bytes32\"}],\"name\":\"removeUserByName\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_storageContract\",\"type\":\"address\"}],\"name\":\"attachExistingStorage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_name\",\"type\":\"bytes32\"}],\"name\":\"getUserByName\",\"outputs\":[{\"name\":\"found\",\"type\":\"bool\"},{\"name\":\"userID\",\"type\":\"uint256\"},{\"name\":\"userAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_name\",\"type\":\"bytes32\"}],\"name\":\"getUserIDByName\",\"outputs\":[{\"name\":\"found\",\"type\":\"bool\"},{\"name\":\"userID\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_imgUrl\",\"type\":\"bytes32\"}],\"name\":\"updateUserProfilePicture\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_name\",\"type\":\"bytes32\"}],\"name\":\"updateUserName\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOnwer\",\"type\":\"address\"}],\"name\":\"detachStorage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_currentAddress\",\"type\":\"address\"},{\"name\":\"_newAddress\",\"type\":\"address\"}],\"name\":\"changeUserAddress\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAllUserIDs\",\"outputs\":[{\"name\":\"list\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_storageAddress\",\"type\":\"address\"}],\"name\":\"StorageCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_storageAddress\",\"type\":\"address\"}],\"name\":\"StorageAttached\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"StorageDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"StorageDetached\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"XnoUsersCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"XnoUsersDeactivated\",\"type\":\"event\"}]"

// XnoUsersBin is the compiled bytecode used for deploying new contracts.
const XnoUsersBin = `0x6060604052341561000f57600080fd5b60008054600160a060020a03191633600160a060020a03161790557f538cea439b4099a169f5388a42ff20adc3ec3cad60cff798c633d760687ad8356100606401000000006100808102610e001704565b604051600160a060020a03909116815260200160405180910390a161008f565b600054600160a060020a031690565b6127da8061009e6000396000f3006060604052600436106101445763ffffffff60e060020a600035041662ef809e811461014957806313ab582b1461016a578063192150e6146101a0578063349ce5d8146101b357806341c0e1b5146101d257806351b42b00146101e5578063620996df146101f85780636d3b5f18146102175780636dddf4031461025b5780637d13172a1461026e578063893d20e8146102815780638ae75977146102b0578063928d9ef0146102c35780639600f498146102e257806398ff9c541461031b578063a4a1e2631461032e578063b433361614610353578063beb377eb14610369578063bfc688661461038b578063c99fe8d8146103a1578063c9b3ac24146103c0578063cff5b31a14610405578063d26d09721461041b578063d9b05f5414610431578063df3db67014610447578063eef1b51514610466578063f6e3cd1b1461048b575b600080fd5b341561015457600080fd5b610168600160a060020a03600435166104f1565b005b341561017557600080fd5b61018c600160a060020a03600435166024356105aa565b604051901515815260200160405180910390f35b34156101ab57600080fd5b610168610705565b34156101be57600080fd5b61018c600160a060020a03600435166107d9565b34156101dd57600080fd5b6101686109bd565b34156101f057600080fd5b6101686109e6565b341561020357600080fd5b610168600160a060020a0360043516610a51565b341561022257600080fd5b61022d600435610ab0565b6040519215158352600160a060020a0390911660208301526040808301919091526060909101905180910390f35b341561026657600080fd5b610168610b79565b341561027957600080fd5b610168610cbe565b341561028c57600080fd5b610294610e00565b604051600160a060020a03909116815260200160405180910390f35b34156102bb57600080fd5b610294610e0f565b34156102ce57600080fd5b610168600160a060020a0360043516610e95565b34156102ed57600080fd5b610301600160a060020a0360043516610f36565b604051911515825260208201526040908101905180910390f35b341561032657600080fd5b610294610ff4565b341561033957600080fd5b610341611003565b60405190815260200160405180910390f35b341561035e57600080fd5b61018c60043561109f565b341561037457600080fd5b61018c600160a060020a036004351660243561115e565b341561039657600080fd5b61018c600435611297565b34156103ac57600080fd5b610168600160a060020a03600435166113c6565b34156103cb57600080fd5b6103d6600435611557565b60405192151583526020830191909152600160a060020a03166040808301919091526060909101905180910390f35b341561041057600080fd5b6103016004356115ec565b341561042657600080fd5b61018c60043561167d565b341561043c57600080fd5b61018c60043561171b565b341561045257600080fd5b610168600160a060020a03600435166117b9565b341561047157600080fd5b61018c600160a060020a0360043581169060243516611855565b341561049657600080fd5b61049e61198d565b60405160208082528190810183818151815260200191508051906020019060200280838360005b838110156104dd5780820151838201526020016104c5565b505050509050019250505060405180910390f35b60018054600160a060020a03161515811461050b57600080fd5b60005433600160a060020a0390811691161461052657600080fd5b60015473__SecurityLib.sol:SecurityLib___________9063268959e590600160a060020a03168460405160e060020a63ffffffff8516028152600160a060020a0392831660048201529116602482015260440160006040518083038186803b151561059257600080fd5b6102c65a03f415156105a357600080fd5b5050505050565b6001805460009190600160a060020a0316151581146105c857600080fd5b60015473__SecurityLib.sol:SecurityLib___________90634cbc5e9990600160a060020a03163360006040516020015260405160e060020a63ffffffff8516028152600160a060020a0392831660048201529116602482015260440160206040518083038186803b151561063d57600080fd5b6102c65a03f4151561064e57600080fd5b50505060405180519050151561066357600080fd5b60015473__UsersLib.sol:UsersLib_________________90630eb1f41190600160a060020a0316868660006040516020015260405160e060020a63ffffffff8616028152600160a060020a039384166004820152919092166024820152604481019190915260640160206040518083038186803b15156106e357600080fd5b6102c65a03f415156106f457600080fd5b505050604051805195945050505050565b60018054600160a060020a03161515811461071f57600080fd5b60005433600160a060020a0390811691161461073a57600080fd5b600154600160a060020a03166341c0e1b56040518163ffffffff1660e060020a028152600401600060405180830381600087803b151561077957600080fd5b6102c65a03f1151561078a57600080fd5b50506001805473ffffffffffffffffffffffffffffffffffffffff19169055507f0f41df7e9c715020a9e0819fd4928247939cd858a13263fd926b1055992e81ae60405160405180910390a150565b600180546000918291829190600160a060020a0316151581146107fb57600080fd5b60015473__SecurityLib.sol:SecurityLib___________90634cbc5e9990600160a060020a03163360006040516020015260405160e060020a63ffffffff8516028152600160a060020a0392831660048201529116602482015260440160206040518083038186803b151561087057600080fd5b6102c65a03f4151561088157600080fd5b50505060405180519050151561089657600080fd5b60015473__UsersLib.sol:UsersLib_________________9063a838e45190600160a060020a03168760006040516040015260405160e060020a63ffffffff8516028152600160a060020a03928316600482015291166024820152604401604080518083038186803b151561090a57600080fd5b6102c65a03f4151561091b57600080fd5b505050604051805190602001805190509250925082156109b45760015473__UsersLib.sol:UsersLib_________________906393d3f77490600160a060020a03168460405160e060020a63ffffffff8516028152600160a060020a039092166004830152602482015260440160006040518083038186803b151561099f57600080fd5b6102c65a03f415156109b057600080fd5b5050505b50909392505050565b60005433600160a060020a039081169116146109d857600080fd5b600054600160a060020a0316ff5b600154600090600160a060020a0316156109ff57600080fd5b60005433600160a060020a03908116911614610a1a57600080fd5b610a226109bd565b7fb757c26153d297b82ab7a0b5ede7ced1080090150244b22f235e9b8a8faba4b060405160405180910390a150565b60005433600160a060020a03908116911614610a6c57600080fd5b600160a060020a0381161515610a8157600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600180546000918291829190600160a060020a031615158114610ad257600080fd5b60015473__UsersLib.sol:UsersLib_________________9063656a196190600160a060020a03168760006040516060015260405160e060020a63ffffffff8516028152600160a060020a039092166004830152602482015260440160606040518083038186803b1515610b4557600080fd5b6102c65a03f41515610b5657600080fd5b505050604051805190602001805190602001805192989197509195509350505050565b600154600090600160a060020a031615610b9257600080fd5b60005433600160a060020a03908116911614610bad57600080fd5b610bb5611c4b565b604051809103906000f0801515610bcb57600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03928316179081905573__SecurityLib.sol:SecurityLib___________91638a19b4d29116610c1b610e00565b60405160e060020a63ffffffff8516028152600160a060020a0392831660048201529116602482015260440160006040518083038186803b1515610c5e57600080fd5b6102c65a03f41515610c6f57600080fd5b50506001547f6c9d76c1677b0909ce573fabdda9f1bbbc891f981a7efa98415bdcd0ff9f54119150600160a060020a0316604051600160a060020a03909116815260200160405180910390a150565b600180546000918291600160a060020a031615158114610cdd57600080fd5b60015473__UsersLib.sol:UsersLib_________________9063a838e45190600160a060020a03163360006040516040015260405160e060020a63ffffffff8516028152600160a060020a03928316600482015291166024820152604401604080518083038186803b1515610d5157600080fd5b6102c65a03f41515610d6257600080fd5b50505060405180519060200180519050925092508215610dfb5760015473__UsersLib.sol:UsersLib_________________906393d3f77490600160a060020a03168460405160e060020a63ffffffff8516028152600160a060020a039092166004830152602482015260440160006040518083038186803b1515610de657600080fd5b6102c65a03f41515610df757600080fd5b5050505b505050565b600054600160a060020a031690565b6001805460009190600160a060020a031615158114610e2d57600080fd5b600154600160a060020a031663893d20e86000604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b1515610e7557600080fd5b6102c65a03f11515610e8657600080fd5b50505060405180519250505090565b60018054600160a060020a031615158114610eaf57600080fd5b60005433600160a060020a03908116911614610eca57600080fd5b60015473__SecurityLib.sol:SecurityLib___________90638a19b4d290600160a060020a03168460405160e060020a63ffffffff8516028152600160a060020a0392831660048201529116602482015260440160006040518083038186803b151561059257600080fd5b600180546000918291600160a060020a031615158114610f5557600080fd5b60015473__UsersLib.sol:UsersLib_________________9063a838e45190600160a060020a03168660006040516040015260405160e060020a63ffffffff8516028152600160a060020a03928316600482015291166024820152604401604080518083038186803b1515610fc957600080fd5b6102c65a03f41515610fda57600080fd5b505050604051805190602001805190509250925050915091565b600154600160a060020a031681565b6001805460009190600160a060020a03161515811461102157600080fd5b60015473__UsersLib.sol:UsersLib_________________9063c817e2df90600160a060020a031660006040516020015260405160e060020a63ffffffff8416028152600160a060020a03909116600482015260240160206040518083038186803b151561108e57600080fd5b6102c65a03f41515610e8657600080fd5b6001805460009190600160a060020a0316151581146110bd57600080fd5b60015473__UsersLib.sol:UsersLib_________________9063364e4ddb90600160a060020a0316338660006040516020015260405160e060020a63ffffffff8616028152600160a060020a039384166004820152919092166024820152604481019190915260640160206040518083038186803b151561113d57600080fd5b6102c65a03f4151561114e57600080fd5b5050506040518051949350505050565b6001805460009190600160a060020a03161515811461117c57600080fd5b60015473__SecurityLib.sol:SecurityLib___________90634cbc5e9990600160a060020a03163360006040516020015260405160e060020a63ffffffff8516028152600160a060020a0392831660048201529116602482015260440160206040518083038186803b15156111f157600080fd5b6102c65a03f4151561120257600080fd5b50505060405180519050151561121757600080fd5b60015473__UsersLib.sol:UsersLib_________________90631fa768db90600160a060020a0316868660006040516020015260405160e060020a63ffffffff8616028152600160a060020a039384166004820152919092166024820152604481019190915260640160206040518083038186803b15156106e357600080fd5b600180546000918291829190600160a060020a0316151581146112b957600080fd5b60015473__SecurityLib.sol:SecurityLib___________90634cbc5e9990600160a060020a03163360006040516020015260405160e060020a63ffffffff8516028152600160a060020a0392831660048201529116602482015260440160206040518083038186803b151561132e57600080fd5b6102c65a03f4151561133f57600080fd5b50505060405180519050151561135457600080fd5b60015473__UsersLib.sol:UsersLib_________________90637c6e7d4590600160a060020a03168760006040516040015260405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401604080518083038186803b151561090a57600080fd5b60005433600160a060020a039081169116146113e157600080fd5b30600160a060020a031681600160a060020a031663893d20e86000604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b151561143157600080fd5b6102c65a03f1151561144257600080fd5b50505060405180519050600160a060020a031614151561146157600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03838116919091179182905573__SecurityLib.sol:SecurityLib___________91638a19b4d291166114b4610e00565b60405160e060020a63ffffffff8516028152600160a060020a0392831660048201529116602482015260440160006040518083038186803b15156114f757600080fd5b6102c65a03f4151561150857600080fd5b50506001547f4fde8bbd254ef19fbcc2cc839169e2d44ffaf6a0091074d7e180c1484709e98b9150600160a060020a0316604051600160a060020a03909116815260200160405180910390a150565b600180546000918291829190600160a060020a03161515811461157957600080fd5b60015473__UsersLib.sol:UsersLib_________________9063127de33390600160a060020a03168760006040516060015260405160e060020a63ffffffff8516028152600160a060020a039092166004830152602482015260440160606040518083038186803b1515610b4557600080fd5b600180546000918291600160a060020a03161515811461160b57600080fd5b60015473__UsersLib.sol:UsersLib_________________90637c6e7d4590600160a060020a03168660006040516040015260405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401604080518083038186803b1515610fc957600080fd5b6001805460009190600160a060020a03161515811461169b57600080fd5b60015473__UsersLib.sol:UsersLib_________________90630eb1f41190600160a060020a0316338660006040516020015260405160e060020a63ffffffff8616028152600160a060020a039384166004820152919092166024820152604481019190915260640160206040518083038186803b151561113d57600080fd5b6001805460009190600160a060020a03161515811461173957600080fd5b60015473__UsersLib.sol:UsersLib_________________90631fa768db90600160a060020a0316338660006040516020015260405160e060020a63ffffffff8616028152600160a060020a039384166004820152919092166024820152604481019190915260640160206040518083038186803b151561113d57600080fd5b60018054600160a060020a0316151581146117d357600080fd5b60005433600160a060020a039081169116146117ee57600080fd5b6117f782611bb5565b6001805473ffffffffffffffffffffffffffffffffffffffff191690557ff536614c5aed5495ea6a3a70d294685beeef2dc10128c2639c35d0228a8f545c82604051600160a060020a03909116815260200160405180910390a15050565b6001805460009190600160a060020a03161515811461187357600080fd5b60015473__SecurityLib.sol:SecurityLib___________90634cbc5e9990600160a060020a03163360006040516020015260405160e060020a63ffffffff8516028152600160a060020a0392831660048201529116602482015260440160206040518083038186803b15156118e857600080fd5b6102c65a03f415156118f957600080fd5b50505060405180519050151561190e57600080fd5b60015473__UsersLib.sol:UsersLib_________________9063868562b790600160a060020a0316868660006040516020015260405160e060020a63ffffffff8616028152600160a060020a0393841660048201529183166024830152909116604482015260640160206040518083038186803b15156106e357600080fd5b611995611c5b565b600180546000918291600160a060020a0316151581146119b457600080fd5b6001546000935073__UsersLib.sol:UsersLib_________________9063c817e2df90600160a060020a0316856040516020015260405160e060020a63ffffffff8416028152600160a060020a03909116600482015260240160206040518083038186803b1515611a2457600080fd5b6102c65a03f41515611a3557600080fd5b50505060405180519050604051805910611a4c5750595b90808252806020026020018201604052509350600091505b60015473__UsersLib.sol:UsersLib_________________906378f0826890600160a060020a031660006040516020015260405160e060020a63ffffffff8416028152600160a060020a03909116600482015260240160206040518083038186803b1515611ad157600080fd5b6102c65a03f41515611ae257600080fd5b50505060405180519050821015611baf5760015473__UsersLib.sol:UsersLib_________________9063d266e83b90600160a060020a03168460006040516020015260405160e060020a63ffffffff8516028152600160a060020a039092166004830152602482015260440160206040518083038186803b1515611b6657600080fd5b6102c65a03f41515611b7757600080fd5b5050506040518051905015611ba45781848480600101955081518110611b9957fe5b602090810290910101525b600190910190611a64565b50505090565b60018054600160a060020a031615158114611bcf57600080fd5b60005433600160a060020a03908116911614611bea57600080fd5b600154600160a060020a031663620996df8360405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401600060405180830381600087803b1515611c3a57600080fd5b6102c65a03f115156105a357600080fd5b604051610b4180611c6e83390190565b6020604051908101604052600081529056006060604052341561000f57600080fd5b60008054600160a060020a033316600160a060020a0319909116179055610b068061003b6000396000f30060606040526004361061013d5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663025ec81a8114610142578063043106c01461016a5780630c55d92514610182578063124f24181461019857806317e7dd22146101ae57806325cf512d146101d85780633562fd20146101f15780633cc1635c1461020a5780633eba9ed21461022057806341c0e1b51461023b57806344bfa56e1461024e5780634c77e5ba146102db5780635a2bf25a1461030d578063620996df1461032f5780638267a9ee1461034e578063893d20e8146103645780639007127b1461037757806393fe42481461038d578063a209a29c146103a3578063a77aa49e146103b9578063ba69fcaa146103d2578063bdc963d8146103e8578063c9a52d2c146103fe578063f586606614610454575b600080fd5b341561014d57600080fd5b6101586004356104aa565b60405190815260200160405180910390f35b341561017557600080fd5b6101806004356104bc565b005b341561018d57600080fd5b610180600435610502565b34156101a357600080fd5b610180600435610537565b34156101b957600080fd5b6101c4600435610563565b604051901515815260200160405180910390f35b34156101e357600080fd5b610180600435602435610578565b34156101fc57600080fd5b6101806004356024356105a5565b341561021557600080fd5b6101806004356105d2565b341561022b57600080fd5b6101806004356024351515610605565b341561024657600080fd5b610180610640565b341561025957600080fd5b610264600435610669565b60405160208082528190810183818151815260200191508051906020019080838360005b838110156102a0578082015183820152602001610288565b50505050905090810190601f1680156102cd5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156102e657600080fd5b6102f160043561072c565b604051600160a060020a03909116815260200160405180910390f35b341561031857600080fd5b610180600435600160a060020a0360243516610747565b341561033a57600080fd5b610180600160a060020a036004351661079d565b341561035957600080fd5b6101806004356107fc565b341561036f57600080fd5b6102f1610828565b341561038257600080fd5b610158600435610838565b341561039857600080fd5b61018060043561084a565b34156103ae57600080fd5b610264600435610876565b34156103c457600080fd5b610180600435602435610902565b34156103dd57600080fd5b61018060043561092f565b34156103f357600080fd5b610158600435610961565b341561040957600080fd5b610180600480359060446024803590810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965061097395505050505050565b341561045f57600080fd5b610180600480359060446024803590810190830135806020601f820181900481020160405190810160405281815292919060208401838380828437509496506109b295505050505050565b60009081526005602052604090205490565b60005433600160a060020a039081169116146104d757600080fd5b6000908152600360205260409020805473ffffffffffffffffffffffffffffffffffffffff19169055565b60005433600160a060020a0390811691161461051d57600080fd5b6000818152600460205260408120610534916109ec565b50565b60005433600160a060020a0390811691161461055257600080fd5b600090815260056020526040812055565b60009081526006602052604090205460ff1690565b60005433600160a060020a0390811691161461059357600080fd5b60009182526005602052604090912055565b60005433600160a060020a039081169116146105c057600080fd5b60009182526001602052604090912055565b60005433600160a060020a039081169116146105ed57600080fd5b6000908152600660205260409020805460ff19169055565b60005433600160a060020a0390811691161461062057600080fd5b600091825260066020526040909120805460ff1916911515919091179055565b60005433600160a060020a0390811691161461065b57600080fd5b600054600160a060020a0316ff5b610671610a30565b6004600083600019166000191681526020019081526020016000208054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156107205780601f106106f557610100808354040283529160200191610720565b820191906000526020600020905b81548152906001019060200180831161070357829003601f168201915b50505050509050919050565b600090815260036020526040902054600160a060020a031690565b60005433600160a060020a0390811691161461076257600080fd5b600091825260036020526040909120805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03909216919091179055565b60005433600160a060020a039081169116146107b857600080fd5b600160a060020a03811615156107cd57600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60005433600160a060020a0390811691161461081757600080fd5b600090815260076020526040812055565b600054600160a060020a03165b90565b60009081526007602052604090205490565b60005433600160a060020a0390811691161461086557600080fd5b600090815260016020526040812055565b61087e610a30565b6002600083600019166000191681526020019081526020016000208054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156107205780601f106106f557610100808354040283529160200191610720565b60005433600160a060020a0390811691161461091d57600080fd5b60009182526007602052604090912055565b60005433600160a060020a0390811691161461094a57600080fd5b6000818152600260205260408120610534916109ec565b60009081526001602052604090205490565b60005433600160a060020a0390811691161461098e57600080fd5b60008281526004602052604090208180516109ad929160200190610a42565b505050565b60005433600160a060020a039081169116146109cd57600080fd5b60008281526002602052604090208180516109ad929160200190610a42565b50805460018160011615610100020316600290046000825580601f10610a125750610534565b601f0160209004906000526020600020908101906105349190610ac0565b60206040519081016040526000815290565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610a8357805160ff1916838001178555610ab0565b82800160010185558215610ab0579182015b82811115610ab0578251825591602001919060010190610a95565b50610abc929150610ac0565b5090565b61083591905b80821115610abc5760008155600101610ac65600a165627a7a7230582089078e369701338df82d3be48ed51f1ec92dc24b78d1b6f8c6cd51c593647c300029a165627a7a7230582072715fa16db31c392b1894cb491e21e59802c98b68b1ad7df4450f3c5f104bd20029`

// DeployXnoUsers deploys a new Ethereum contract, binding an instance of XnoUsers to it.
func DeployXnoUsers(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *XnoUsers, error) {
	parsed, err := abi.JSON(strings.NewReader(XnoUsersABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(XnoUsersBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &XnoUsers{XnoUsersCaller: XnoUsersCaller{contract: contract}, XnoUsersTransactor: XnoUsersTransactor{contract: contract}}, nil
}

// XnoUsers is an auto generated Go binding around an Ethereum contract.
type XnoUsers struct {
	XnoUsersCaller     // Read-only binding to the contract
	XnoUsersTransactor // Write-only binding to the contract
}

// XnoUsersCaller is an auto generated read-only Go binding around an Ethereum contract.
type XnoUsersCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// XnoUsersTransactor is an auto generated write-only Go binding around an Ethereum contract.
type XnoUsersTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// XnoUsersSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type XnoUsersSession struct {
	Contract     *XnoUsers         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// XnoUsersCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type XnoUsersCallerSession struct {
	Contract *XnoUsersCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// XnoUsersTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type XnoUsersTransactorSession struct {
	Contract     *XnoUsersTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// XnoUsersRaw is an auto generated low-level Go binding around an Ethereum contract.
type XnoUsersRaw struct {
	Contract *XnoUsers // Generic contract binding to access the raw methods on
}

// XnoUsersCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type XnoUsersCallerRaw struct {
	Contract *XnoUsersCaller // Generic read-only contract binding to access the raw methods on
}

// XnoUsersTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type XnoUsersTransactorRaw struct {
	Contract *XnoUsersTransactor // Generic write-only contract binding to access the raw methods on
}

// NewXnoUsers creates a new instance of XnoUsers, bound to a specific deployed contract.
func NewXnoUsers(address common.Address, backend bind.ContractBackend) (*XnoUsers, error) {
	contract, err := bindXnoUsers(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &XnoUsers{XnoUsersCaller: XnoUsersCaller{contract: contract}, XnoUsersTransactor: XnoUsersTransactor{contract: contract}}, nil
}

// NewXnoUsersCaller creates a new read-only instance of XnoUsers, bound to a specific deployed contract.
func NewXnoUsersCaller(address common.Address, caller bind.ContractCaller) (*XnoUsersCaller, error) {
	contract, err := bindXnoUsers(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &XnoUsersCaller{contract: contract}, nil
}

// NewXnoUsersTransactor creates a new write-only instance of XnoUsers, bound to a specific deployed contract.
func NewXnoUsersTransactor(address common.Address, transactor bind.ContractTransactor) (*XnoUsersTransactor, error) {
	contract, err := bindXnoUsers(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &XnoUsersTransactor{contract: contract}, nil
}

// bindXnoUsers binds a generic wrapper to an already deployed contract.
func bindXnoUsers(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(XnoUsersABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_XnoUsers *XnoUsersRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _XnoUsers.Contract.XnoUsersCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_XnoUsers *XnoUsersRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _XnoUsers.Contract.XnoUsersTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_XnoUsers *XnoUsersRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _XnoUsers.Contract.XnoUsersTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_XnoUsers *XnoUsersCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _XnoUsers.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_XnoUsers *XnoUsersTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _XnoUsers.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_XnoUsers *XnoUsersTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _XnoUsers.Contract.contract.Transact(opts, method, params...)
}

// EternalStorage is a free data retrieval call binding the contract method 0x98ff9c54.
//
// Solidity: function eternalStorage() constant returns(address)
func (_XnoUsers *XnoUsersCaller) EternalStorage(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _XnoUsers.contract.Call(opts, out, "eternalStorage")
	return *ret0, err
}

// EternalStorage is a free data retrieval call binding the contract method 0x98ff9c54.
//
// Solidity: function eternalStorage() constant returns(address)
func (_XnoUsers *XnoUsersSession) EternalStorage() (common.Address, error) {
	return _XnoUsers.Contract.EternalStorage(&_XnoUsers.CallOpts)
}

// EternalStorage is a free data retrieval call binding the contract method 0x98ff9c54.
//
// Solidity: function eternalStorage() constant returns(address)
func (_XnoUsers *XnoUsersCallerSession) EternalStorage() (common.Address, error) {
	return _XnoUsers.Contract.EternalStorage(&_XnoUsers.CallOpts)
}

// GetAllUserIDs is a free data retrieval call binding the contract method 0xf6e3cd1b.
//
// Solidity: function getAllUserIDs() constant returns(list uint256[])
func (_XnoUsers *XnoUsersCaller) GetAllUserIDs(opts *bind.CallOpts) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _XnoUsers.contract.Call(opts, out, "getAllUserIDs")
	return *ret0, err
}

// GetAllUserIDs is a free data retrieval call binding the contract method 0xf6e3cd1b.
//
// Solidity: function getAllUserIDs() constant returns(list uint256[])
func (_XnoUsers *XnoUsersSession) GetAllUserIDs() ([]*big.Int, error) {
	return _XnoUsers.Contract.GetAllUserIDs(&_XnoUsers.CallOpts)
}

// GetAllUserIDs is a free data retrieval call binding the contract method 0xf6e3cd1b.
//
// Solidity: function getAllUserIDs() constant returns(list uint256[])
func (_XnoUsers *XnoUsersCallerSession) GetAllUserIDs() ([]*big.Int, error) {
	return _XnoUsers.Contract.GetAllUserIDs(&_XnoUsers.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_XnoUsers *XnoUsersCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _XnoUsers.contract.Call(opts, out, "getOwner")
	return *ret0, err
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_XnoUsers *XnoUsersSession) GetOwner() (common.Address, error) {
	return _XnoUsers.Contract.GetOwner(&_XnoUsers.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_XnoUsers *XnoUsersCallerSession) GetOwner() (common.Address, error) {
	return _XnoUsers.Contract.GetOwner(&_XnoUsers.CallOpts)
}

// GetStorageOwner is a free data retrieval call binding the contract method 0x8ae75977.
//
// Solidity: function getStorageOwner() constant returns(address)
func (_XnoUsers *XnoUsersCaller) GetStorageOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _XnoUsers.contract.Call(opts, out, "getStorageOwner")
	return *ret0, err
}

// GetStorageOwner is a free data retrieval call binding the contract method 0x8ae75977.
//
// Solidity: function getStorageOwner() constant returns(address)
func (_XnoUsers *XnoUsersSession) GetStorageOwner() (common.Address, error) {
	return _XnoUsers.Contract.GetStorageOwner(&_XnoUsers.CallOpts)
}

// GetStorageOwner is a free data retrieval call binding the contract method 0x8ae75977.
//
// Solidity: function getStorageOwner() constant returns(address)
func (_XnoUsers *XnoUsersCallerSession) GetStorageOwner() (common.Address, error) {
	return _XnoUsers.Contract.GetStorageOwner(&_XnoUsers.CallOpts)
}

// GetUserByID is a free data retrieval call binding the contract method 0x6d3b5f18.
//
// Solidity: function getUserByID(_userID uint256) constant returns(found bool, userAddress address, userName bytes32)
func (_XnoUsers *XnoUsersCaller) GetUserByID(opts *bind.CallOpts, _userID *big.Int) (struct {
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
	err := _XnoUsers.contract.Call(opts, out, "getUserByID", _userID)
	return *ret, err
}

// GetUserByID is a free data retrieval call binding the contract method 0x6d3b5f18.
//
// Solidity: function getUserByID(_userID uint256) constant returns(found bool, userAddress address, userName bytes32)
func (_XnoUsers *XnoUsersSession) GetUserByID(_userID *big.Int) (struct {
	Found       bool
	UserAddress common.Address
	UserName    [32]byte
}, error) {
	return _XnoUsers.Contract.GetUserByID(&_XnoUsers.CallOpts, _userID)
}

// GetUserByID is a free data retrieval call binding the contract method 0x6d3b5f18.
//
// Solidity: function getUserByID(_userID uint256) constant returns(found bool, userAddress address, userName bytes32)
func (_XnoUsers *XnoUsersCallerSession) GetUserByID(_userID *big.Int) (struct {
	Found       bool
	UserAddress common.Address
	UserName    [32]byte
}, error) {
	return _XnoUsers.Contract.GetUserByID(&_XnoUsers.CallOpts, _userID)
}

// GetUserByName is a free data retrieval call binding the contract method 0xc9b3ac24.
//
// Solidity: function getUserByName(_name bytes32) constant returns(found bool, userID uint256, userAddress address)
func (_XnoUsers *XnoUsersCaller) GetUserByName(opts *bind.CallOpts, _name [32]byte) (struct {
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
	err := _XnoUsers.contract.Call(opts, out, "getUserByName", _name)
	return *ret, err
}

// GetUserByName is a free data retrieval call binding the contract method 0xc9b3ac24.
//
// Solidity: function getUserByName(_name bytes32) constant returns(found bool, userID uint256, userAddress address)
func (_XnoUsers *XnoUsersSession) GetUserByName(_name [32]byte) (struct {
	Found       bool
	UserID      *big.Int
	UserAddress common.Address
}, error) {
	return _XnoUsers.Contract.GetUserByName(&_XnoUsers.CallOpts, _name)
}

// GetUserByName is a free data retrieval call binding the contract method 0xc9b3ac24.
//
// Solidity: function getUserByName(_name bytes32) constant returns(found bool, userID uint256, userAddress address)
func (_XnoUsers *XnoUsersCallerSession) GetUserByName(_name [32]byte) (struct {
	Found       bool
	UserID      *big.Int
	UserAddress common.Address
}, error) {
	return _XnoUsers.Contract.GetUserByName(&_XnoUsers.CallOpts, _name)
}

// GetUserIDByAddress is a free data retrieval call binding the contract method 0x9600f498.
//
// Solidity: function getUserIDByAddress(_address address) constant returns(found bool, userID uint256)
func (_XnoUsers *XnoUsersCaller) GetUserIDByAddress(opts *bind.CallOpts, _address common.Address) (struct {
	Found  bool
	UserID *big.Int
}, error) {
	ret := new(struct {
		Found  bool
		UserID *big.Int
	})
	out := ret
	err := _XnoUsers.contract.Call(opts, out, "getUserIDByAddress", _address)
	return *ret, err
}

// GetUserIDByAddress is a free data retrieval call binding the contract method 0x9600f498.
//
// Solidity: function getUserIDByAddress(_address address) constant returns(found bool, userID uint256)
func (_XnoUsers *XnoUsersSession) GetUserIDByAddress(_address common.Address) (struct {
	Found  bool
	UserID *big.Int
}, error) {
	return _XnoUsers.Contract.GetUserIDByAddress(&_XnoUsers.CallOpts, _address)
}

// GetUserIDByAddress is a free data retrieval call binding the contract method 0x9600f498.
//
// Solidity: function getUserIDByAddress(_address address) constant returns(found bool, userID uint256)
func (_XnoUsers *XnoUsersCallerSession) GetUserIDByAddress(_address common.Address) (struct {
	Found  bool
	UserID *big.Int
}, error) {
	return _XnoUsers.Contract.GetUserIDByAddress(&_XnoUsers.CallOpts, _address)
}

// GetUserIDByName is a free data retrieval call binding the contract method 0xcff5b31a.
//
// Solidity: function getUserIDByName(_name bytes32) constant returns(found bool, userID uint256)
func (_XnoUsers *XnoUsersCaller) GetUserIDByName(opts *bind.CallOpts, _name [32]byte) (struct {
	Found  bool
	UserID *big.Int
}, error) {
	ret := new(struct {
		Found  bool
		UserID *big.Int
	})
	out := ret
	err := _XnoUsers.contract.Call(opts, out, "getUserIDByName", _name)
	return *ret, err
}

// GetUserIDByName is a free data retrieval call binding the contract method 0xcff5b31a.
//
// Solidity: function getUserIDByName(_name bytes32) constant returns(found bool, userID uint256)
func (_XnoUsers *XnoUsersSession) GetUserIDByName(_name [32]byte) (struct {
	Found  bool
	UserID *big.Int
}, error) {
	return _XnoUsers.Contract.GetUserIDByName(&_XnoUsers.CallOpts, _name)
}

// GetUserIDByName is a free data retrieval call binding the contract method 0xcff5b31a.
//
// Solidity: function getUserIDByName(_name bytes32) constant returns(found bool, userID uint256)
func (_XnoUsers *XnoUsersCallerSession) GetUserIDByName(_name [32]byte) (struct {
	Found  bool
	UserID *big.Int
}, error) {
	return _XnoUsers.Contract.GetUserIDByName(&_XnoUsers.CallOpts, _name)
}

// GetUsersCount is a free data retrieval call binding the contract method 0xa4a1e263.
//
// Solidity: function getUsersCount() constant returns(uint256)
func (_XnoUsers *XnoUsersCaller) GetUsersCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _XnoUsers.contract.Call(opts, out, "getUsersCount")
	return *ret0, err
}

// GetUsersCount is a free data retrieval call binding the contract method 0xa4a1e263.
//
// Solidity: function getUsersCount() constant returns(uint256)
func (_XnoUsers *XnoUsersSession) GetUsersCount() (*big.Int, error) {
	return _XnoUsers.Contract.GetUsersCount(&_XnoUsers.CallOpts)
}

// GetUsersCount is a free data retrieval call binding the contract method 0xa4a1e263.
//
// Solidity: function getUsersCount() constant returns(uint256)
func (_XnoUsers *XnoUsersCallerSession) GetUsersCount() (*big.Int, error) {
	return _XnoUsers.Contract.GetUsersCount(&_XnoUsers.CallOpts)
}

// AddXnoAdmin is a paid mutator transaction binding the contract method 0x928d9ef0.
//
// Solidity: function addXnoAdmin(_admin address) returns()
func (_XnoUsers *XnoUsersTransactor) AddXnoAdmin(opts *bind.TransactOpts, _admin common.Address) (*types.Transaction, error) {
	return _XnoUsers.contract.Transact(opts, "addXnoAdmin", _admin)
}

// AddXnoAdmin is a paid mutator transaction binding the contract method 0x928d9ef0.
//
// Solidity: function addXnoAdmin(_admin address) returns()
func (_XnoUsers *XnoUsersSession) AddXnoAdmin(_admin common.Address) (*types.Transaction, error) {
	return _XnoUsers.Contract.AddXnoAdmin(&_XnoUsers.TransactOpts, _admin)
}

// AddXnoAdmin is a paid mutator transaction binding the contract method 0x928d9ef0.
//
// Solidity: function addXnoAdmin(_admin address) returns()
func (_XnoUsers *XnoUsersTransactorSession) AddXnoAdmin(_admin common.Address) (*types.Transaction, error) {
	return _XnoUsers.Contract.AddXnoAdmin(&_XnoUsers.TransactOpts, _admin)
}

// AttachExistingStorage is a paid mutator transaction binding the contract method 0xc99fe8d8.
//
// Solidity: function attachExistingStorage(_storageContract address) returns()
func (_XnoUsers *XnoUsersTransactor) AttachExistingStorage(opts *bind.TransactOpts, _storageContract common.Address) (*types.Transaction, error) {
	return _XnoUsers.contract.Transact(opts, "attachExistingStorage", _storageContract)
}

// AttachExistingStorage is a paid mutator transaction binding the contract method 0xc99fe8d8.
//
// Solidity: function attachExistingStorage(_storageContract address) returns()
func (_XnoUsers *XnoUsersSession) AttachExistingStorage(_storageContract common.Address) (*types.Transaction, error) {
	return _XnoUsers.Contract.AttachExistingStorage(&_XnoUsers.TransactOpts, _storageContract)
}

// AttachExistingStorage is a paid mutator transaction binding the contract method 0xc99fe8d8.
//
// Solidity: function attachExistingStorage(_storageContract address) returns()
func (_XnoUsers *XnoUsersTransactorSession) AttachExistingStorage(_storageContract common.Address) (*types.Transaction, error) {
	return _XnoUsers.Contract.AttachExistingStorage(&_XnoUsers.TransactOpts, _storageContract)
}

// ChangeUserAddress is a paid mutator transaction binding the contract method 0xeef1b515.
//
// Solidity: function changeUserAddress(_currentAddress address, _newAddress address) returns(success bool)
func (_XnoUsers *XnoUsersTransactor) ChangeUserAddress(opts *bind.TransactOpts, _currentAddress common.Address, _newAddress common.Address) (*types.Transaction, error) {
	return _XnoUsers.contract.Transact(opts, "changeUserAddress", _currentAddress, _newAddress)
}

// ChangeUserAddress is a paid mutator transaction binding the contract method 0xeef1b515.
//
// Solidity: function changeUserAddress(_currentAddress address, _newAddress address) returns(success bool)
func (_XnoUsers *XnoUsersSession) ChangeUserAddress(_currentAddress common.Address, _newAddress common.Address) (*types.Transaction, error) {
	return _XnoUsers.Contract.ChangeUserAddress(&_XnoUsers.TransactOpts, _currentAddress, _newAddress)
}

// ChangeUserAddress is a paid mutator transaction binding the contract method 0xeef1b515.
//
// Solidity: function changeUserAddress(_currentAddress address, _newAddress address) returns(success bool)
func (_XnoUsers *XnoUsersTransactorSession) ChangeUserAddress(_currentAddress common.Address, _newAddress common.Address) (*types.Transaction, error) {
	return _XnoUsers.Contract.ChangeUserAddress(&_XnoUsers.TransactOpts, _currentAddress, _newAddress)
}

// ChangeUserName is a paid mutator transaction binding the contract method 0xbeb377eb.
//
// Solidity: function changeUserName(_address address, _name bytes32) returns(success bool)
func (_XnoUsers *XnoUsersTransactor) ChangeUserName(opts *bind.TransactOpts, _address common.Address, _name [32]byte) (*types.Transaction, error) {
	return _XnoUsers.contract.Transact(opts, "changeUserName", _address, _name)
}

// ChangeUserName is a paid mutator transaction binding the contract method 0xbeb377eb.
//
// Solidity: function changeUserName(_address address, _name bytes32) returns(success bool)
func (_XnoUsers *XnoUsersSession) ChangeUserName(_address common.Address, _name [32]byte) (*types.Transaction, error) {
	return _XnoUsers.Contract.ChangeUserName(&_XnoUsers.TransactOpts, _address, _name)
}

// ChangeUserName is a paid mutator transaction binding the contract method 0xbeb377eb.
//
// Solidity: function changeUserName(_address address, _name bytes32) returns(success bool)
func (_XnoUsers *XnoUsersTransactorSession) ChangeUserName(_address common.Address, _name [32]byte) (*types.Transaction, error) {
	return _XnoUsers.Contract.ChangeUserName(&_XnoUsers.TransactOpts, _address, _name)
}

// ChangeUserProfilePicture is a paid mutator transaction binding the contract method 0x13ab582b.
//
// Solidity: function changeUserProfilePicture(_address address, _imgUrl bytes32) returns(success bool)
func (_XnoUsers *XnoUsersTransactor) ChangeUserProfilePicture(opts *bind.TransactOpts, _address common.Address, _imgUrl [32]byte) (*types.Transaction, error) {
	return _XnoUsers.contract.Transact(opts, "changeUserProfilePicture", _address, _imgUrl)
}

// ChangeUserProfilePicture is a paid mutator transaction binding the contract method 0x13ab582b.
//
// Solidity: function changeUserProfilePicture(_address address, _imgUrl bytes32) returns(success bool)
func (_XnoUsers *XnoUsersSession) ChangeUserProfilePicture(_address common.Address, _imgUrl [32]byte) (*types.Transaction, error) {
	return _XnoUsers.Contract.ChangeUserProfilePicture(&_XnoUsers.TransactOpts, _address, _imgUrl)
}

// ChangeUserProfilePicture is a paid mutator transaction binding the contract method 0x13ab582b.
//
// Solidity: function changeUserProfilePicture(_address address, _imgUrl bytes32) returns(success bool)
func (_XnoUsers *XnoUsersTransactorSession) ChangeUserProfilePicture(_address common.Address, _imgUrl [32]byte) (*types.Transaction, error) {
	return _XnoUsers.Contract.ChangeUserProfilePicture(&_XnoUsers.TransactOpts, _address, _imgUrl)
}

// CreateNewStorage is a paid mutator transaction binding the contract method 0x6dddf403.
//
// Solidity: function createNewStorage() returns()
func (_XnoUsers *XnoUsersTransactor) CreateNewStorage(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _XnoUsers.contract.Transact(opts, "createNewStorage")
}

// CreateNewStorage is a paid mutator transaction binding the contract method 0x6dddf403.
//
// Solidity: function createNewStorage() returns()
func (_XnoUsers *XnoUsersSession) CreateNewStorage() (*types.Transaction, error) {
	return _XnoUsers.Contract.CreateNewStorage(&_XnoUsers.TransactOpts)
}

// CreateNewStorage is a paid mutator transaction binding the contract method 0x6dddf403.
//
// Solidity: function createNewStorage() returns()
func (_XnoUsers *XnoUsersTransactorSession) CreateNewStorage() (*types.Transaction, error) {
	return _XnoUsers.Contract.CreateNewStorage(&_XnoUsers.TransactOpts)
}

// Deactivate is a paid mutator transaction binding the contract method 0x51b42b00.
//
// Solidity: function deactivate() returns()
func (_XnoUsers *XnoUsersTransactor) Deactivate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _XnoUsers.contract.Transact(opts, "deactivate")
}

// Deactivate is a paid mutator transaction binding the contract method 0x51b42b00.
//
// Solidity: function deactivate() returns()
func (_XnoUsers *XnoUsersSession) Deactivate() (*types.Transaction, error) {
	return _XnoUsers.Contract.Deactivate(&_XnoUsers.TransactOpts)
}

// Deactivate is a paid mutator transaction binding the contract method 0x51b42b00.
//
// Solidity: function deactivate() returns()
func (_XnoUsers *XnoUsersTransactorSession) Deactivate() (*types.Transaction, error) {
	return _XnoUsers.Contract.Deactivate(&_XnoUsers.TransactOpts)
}

// DeleteStorage is a paid mutator transaction binding the contract method 0x192150e6.
//
// Solidity: function deleteStorage() returns()
func (_XnoUsers *XnoUsersTransactor) DeleteStorage(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _XnoUsers.contract.Transact(opts, "deleteStorage")
}

// DeleteStorage is a paid mutator transaction binding the contract method 0x192150e6.
//
// Solidity: function deleteStorage() returns()
func (_XnoUsers *XnoUsersSession) DeleteStorage() (*types.Transaction, error) {
	return _XnoUsers.Contract.DeleteStorage(&_XnoUsers.TransactOpts)
}

// DeleteStorage is a paid mutator transaction binding the contract method 0x192150e6.
//
// Solidity: function deleteStorage() returns()
func (_XnoUsers *XnoUsersTransactorSession) DeleteStorage() (*types.Transaction, error) {
	return _XnoUsers.Contract.DeleteStorage(&_XnoUsers.TransactOpts)
}

// DeregisterUser is a paid mutator transaction binding the contract method 0x7d13172a.
//
// Solidity: function deregisterUser() returns()
func (_XnoUsers *XnoUsersTransactor) DeregisterUser(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _XnoUsers.contract.Transact(opts, "deregisterUser")
}

// DeregisterUser is a paid mutator transaction binding the contract method 0x7d13172a.
//
// Solidity: function deregisterUser() returns()
func (_XnoUsers *XnoUsersSession) DeregisterUser() (*types.Transaction, error) {
	return _XnoUsers.Contract.DeregisterUser(&_XnoUsers.TransactOpts)
}

// DeregisterUser is a paid mutator transaction binding the contract method 0x7d13172a.
//
// Solidity: function deregisterUser() returns()
func (_XnoUsers *XnoUsersTransactorSession) DeregisterUser() (*types.Transaction, error) {
	return _XnoUsers.Contract.DeregisterUser(&_XnoUsers.TransactOpts)
}

// DetachStorage is a paid mutator transaction binding the contract method 0xdf3db670.
//
// Solidity: function detachStorage(_newOnwer address) returns()
func (_XnoUsers *XnoUsersTransactor) DetachStorage(opts *bind.TransactOpts, _newOnwer common.Address) (*types.Transaction, error) {
	return _XnoUsers.contract.Transact(opts, "detachStorage", _newOnwer)
}

// DetachStorage is a paid mutator transaction binding the contract method 0xdf3db670.
//
// Solidity: function detachStorage(_newOnwer address) returns()
func (_XnoUsers *XnoUsersSession) DetachStorage(_newOnwer common.Address) (*types.Transaction, error) {
	return _XnoUsers.Contract.DetachStorage(&_XnoUsers.TransactOpts, _newOnwer)
}

// DetachStorage is a paid mutator transaction binding the contract method 0xdf3db670.
//
// Solidity: function detachStorage(_newOnwer address) returns()
func (_XnoUsers *XnoUsersTransactorSession) DetachStorage(_newOnwer common.Address) (*types.Transaction, error) {
	return _XnoUsers.Contract.DetachStorage(&_XnoUsers.TransactOpts, _newOnwer)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_XnoUsers *XnoUsersTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _XnoUsers.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_XnoUsers *XnoUsersSession) Kill() (*types.Transaction, error) {
	return _XnoUsers.Contract.Kill(&_XnoUsers.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_XnoUsers *XnoUsersTransactorSession) Kill() (*types.Transaction, error) {
	return _XnoUsers.Contract.Kill(&_XnoUsers.TransactOpts)
}

// PassOwnership is a paid mutator transaction binding the contract method 0x620996df.
//
// Solidity: function passOwnership(_newOwner address) returns()
func (_XnoUsers *XnoUsersTransactor) PassOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _XnoUsers.contract.Transact(opts, "passOwnership", _newOwner)
}

// PassOwnership is a paid mutator transaction binding the contract method 0x620996df.
//
// Solidity: function passOwnership(_newOwner address) returns()
func (_XnoUsers *XnoUsersSession) PassOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _XnoUsers.Contract.PassOwnership(&_XnoUsers.TransactOpts, _newOwner)
}

// PassOwnership is a paid mutator transaction binding the contract method 0x620996df.
//
// Solidity: function passOwnership(_newOwner address) returns()
func (_XnoUsers *XnoUsersTransactorSession) PassOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _XnoUsers.Contract.PassOwnership(&_XnoUsers.TransactOpts, _newOwner)
}

// RegisterNewUser is a paid mutator transaction binding the contract method 0xb4333616.
//
// Solidity: function registerNewUser(_name bytes32) returns(success bool)
func (_XnoUsers *XnoUsersTransactor) RegisterNewUser(opts *bind.TransactOpts, _name [32]byte) (*types.Transaction, error) {
	return _XnoUsers.contract.Transact(opts, "registerNewUser", _name)
}

// RegisterNewUser is a paid mutator transaction binding the contract method 0xb4333616.
//
// Solidity: function registerNewUser(_name bytes32) returns(success bool)
func (_XnoUsers *XnoUsersSession) RegisterNewUser(_name [32]byte) (*types.Transaction, error) {
	return _XnoUsers.Contract.RegisterNewUser(&_XnoUsers.TransactOpts, _name)
}

// RegisterNewUser is a paid mutator transaction binding the contract method 0xb4333616.
//
// Solidity: function registerNewUser(_name bytes32) returns(success bool)
func (_XnoUsers *XnoUsersTransactorSession) RegisterNewUser(_name [32]byte) (*types.Transaction, error) {
	return _XnoUsers.Contract.RegisterNewUser(&_XnoUsers.TransactOpts, _name)
}

// RemoveUserByAddress is a paid mutator transaction binding the contract method 0x349ce5d8.
//
// Solidity: function removeUserByAddress(_address address) returns(success bool)
func (_XnoUsers *XnoUsersTransactor) RemoveUserByAddress(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _XnoUsers.contract.Transact(opts, "removeUserByAddress", _address)
}

// RemoveUserByAddress is a paid mutator transaction binding the contract method 0x349ce5d8.
//
// Solidity: function removeUserByAddress(_address address) returns(success bool)
func (_XnoUsers *XnoUsersSession) RemoveUserByAddress(_address common.Address) (*types.Transaction, error) {
	return _XnoUsers.Contract.RemoveUserByAddress(&_XnoUsers.TransactOpts, _address)
}

// RemoveUserByAddress is a paid mutator transaction binding the contract method 0x349ce5d8.
//
// Solidity: function removeUserByAddress(_address address) returns(success bool)
func (_XnoUsers *XnoUsersTransactorSession) RemoveUserByAddress(_address common.Address) (*types.Transaction, error) {
	return _XnoUsers.Contract.RemoveUserByAddress(&_XnoUsers.TransactOpts, _address)
}

// RemoveUserByName is a paid mutator transaction binding the contract method 0xbfc68866.
//
// Solidity: function removeUserByName(_name bytes32) returns(success bool)
func (_XnoUsers *XnoUsersTransactor) RemoveUserByName(opts *bind.TransactOpts, _name [32]byte) (*types.Transaction, error) {
	return _XnoUsers.contract.Transact(opts, "removeUserByName", _name)
}

// RemoveUserByName is a paid mutator transaction binding the contract method 0xbfc68866.
//
// Solidity: function removeUserByName(_name bytes32) returns(success bool)
func (_XnoUsers *XnoUsersSession) RemoveUserByName(_name [32]byte) (*types.Transaction, error) {
	return _XnoUsers.Contract.RemoveUserByName(&_XnoUsers.TransactOpts, _name)
}

// RemoveUserByName is a paid mutator transaction binding the contract method 0xbfc68866.
//
// Solidity: function removeUserByName(_name bytes32) returns(success bool)
func (_XnoUsers *XnoUsersTransactorSession) RemoveUserByName(_name [32]byte) (*types.Transaction, error) {
	return _XnoUsers.Contract.RemoveUserByName(&_XnoUsers.TransactOpts, _name)
}

// RemoveXnoAdmin is a paid mutator transaction binding the contract method 0x00ef809e.
//
// Solidity: function removeXnoAdmin(_admin address) returns()
func (_XnoUsers *XnoUsersTransactor) RemoveXnoAdmin(opts *bind.TransactOpts, _admin common.Address) (*types.Transaction, error) {
	return _XnoUsers.contract.Transact(opts, "removeXnoAdmin", _admin)
}

// RemoveXnoAdmin is a paid mutator transaction binding the contract method 0x00ef809e.
//
// Solidity: function removeXnoAdmin(_admin address) returns()
func (_XnoUsers *XnoUsersSession) RemoveXnoAdmin(_admin common.Address) (*types.Transaction, error) {
	return _XnoUsers.Contract.RemoveXnoAdmin(&_XnoUsers.TransactOpts, _admin)
}

// RemoveXnoAdmin is a paid mutator transaction binding the contract method 0x00ef809e.
//
// Solidity: function removeXnoAdmin(_admin address) returns()
func (_XnoUsers *XnoUsersTransactorSession) RemoveXnoAdmin(_admin common.Address) (*types.Transaction, error) {
	return _XnoUsers.Contract.RemoveXnoAdmin(&_XnoUsers.TransactOpts, _admin)
}

// UpdateUserName is a paid mutator transaction binding the contract method 0xd9b05f54.
//
// Solidity: function updateUserName(_name bytes32) returns(success bool)
func (_XnoUsers *XnoUsersTransactor) UpdateUserName(opts *bind.TransactOpts, _name [32]byte) (*types.Transaction, error) {
	return _XnoUsers.contract.Transact(opts, "updateUserName", _name)
}

// UpdateUserName is a paid mutator transaction binding the contract method 0xd9b05f54.
//
// Solidity: function updateUserName(_name bytes32) returns(success bool)
func (_XnoUsers *XnoUsersSession) UpdateUserName(_name [32]byte) (*types.Transaction, error) {
	return _XnoUsers.Contract.UpdateUserName(&_XnoUsers.TransactOpts, _name)
}

// UpdateUserName is a paid mutator transaction binding the contract method 0xd9b05f54.
//
// Solidity: function updateUserName(_name bytes32) returns(success bool)
func (_XnoUsers *XnoUsersTransactorSession) UpdateUserName(_name [32]byte) (*types.Transaction, error) {
	return _XnoUsers.Contract.UpdateUserName(&_XnoUsers.TransactOpts, _name)
}

// UpdateUserProfilePicture is a paid mutator transaction binding the contract method 0xd26d0972.
//
// Solidity: function updateUserProfilePicture(_imgUrl bytes32) returns(success bool)
func (_XnoUsers *XnoUsersTransactor) UpdateUserProfilePicture(opts *bind.TransactOpts, _imgUrl [32]byte) (*types.Transaction, error) {
	return _XnoUsers.contract.Transact(opts, "updateUserProfilePicture", _imgUrl)
}

// UpdateUserProfilePicture is a paid mutator transaction binding the contract method 0xd26d0972.
//
// Solidity: function updateUserProfilePicture(_imgUrl bytes32) returns(success bool)
func (_XnoUsers *XnoUsersSession) UpdateUserProfilePicture(_imgUrl [32]byte) (*types.Transaction, error) {
	return _XnoUsers.Contract.UpdateUserProfilePicture(&_XnoUsers.TransactOpts, _imgUrl)
}

// UpdateUserProfilePicture is a paid mutator transaction binding the contract method 0xd26d0972.
//
// Solidity: function updateUserProfilePicture(_imgUrl bytes32) returns(success bool)
func (_XnoUsers *XnoUsersTransactorSession) UpdateUserProfilePicture(_imgUrl [32]byte) (*types.Transaction, error) {
	return _XnoUsers.Contract.UpdateUserProfilePicture(&_XnoUsers.TransactOpts, _imgUrl)
}
