// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package xnousers

import (
	"math/big"
	"strings"

	"github.com/xenioplatform/go-xenio/common"
	"github.com/xenioplatform/go-xenio/core/types"
	"github.com/xenioplatform/go-xenio/accounts/abi"
	"github.com/xenioplatform/go-xenio/accounts/abi/bind"
)

// XNOUsersABI is the input ABI used to generate the binding from.
const XNOUsersABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getUsers\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\"},{\"name\":\"\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getGamers\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\"},{\"name\":\"\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAllUsersAddresses\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getGamersAddresses\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getServersNumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"getUser\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getServers\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\"},{\"name\":\"\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getServersAddresses\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"gameAddress\",\"type\":\"address\"}],\"name\":\"removeUser\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getGamersNumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"source\",\"type\":\"string\"}],\"name\":\"stringToBytes32\",\"outputs\":[{\"name\":\"result\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_isServer\",\"type\":\"bool\"},{\"name\":\"_game\",\"type\":\"string\"}],\"name\":\"registerNewUser\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isServer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"}]"

// XNOUsersBin is the compiled bytecode used for deploying new contracts.
const XNOUsersBin = `0x606060405260008054600160a060020a033316600160a060020a03199091161790556113db806100306000396000f300606060405236156100c15763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041662ce8e3e81146100ce5780630d2e017b1461017a5780633a5aeec31461018d5780636603e8ac146101f35780636ecc46e0146102065780636f77926b1461022b57806375c1c7e7146103345780637d22f9ab14610347578063985751881461035a578063ae4191c81461038d578063cfb51928146103a0578063e5e8093a146103f1578063fda089581461048f575b34156100cc57600080fd5b005b34156100d957600080fd5b6100e16104ae565b604051808060200180602001838103835285818151815260200191508051906020019060200280838360005b8381101561012557808201518382015260200161010d565b50505050905001838103825284818151815260200191508051906020019060200280838360005b8381101561016457808201518382015260200161014c565b5050505090500194505050505060405180910390f35b341561018557600080fd5b6100e16108a0565b341561019857600080fd5b6101a0610ab0565b60405160208082528190810183818151815260200191508051906020019060200280838360005b838110156101df5780820151838201526020016101c7565b505050509050019250505060405180910390f35b34156101fe57600080fd5b6101a0610bca565b341561021157600080fd5b610219610c33565b60405190815260200160405180910390f35b341561023657600080fd5b61024a600160a060020a0360043516610c39565b6040518215156020820152606080825281906040820190820186818151815260200191508051906020019080838360005b8381101561029357808201518382015260200161027b565b50505050905090810190601f1680156102c05780820380516001836020036101000a031916815260200191505b50838103825284818151815260200191508051906020019080838360005b838110156102f65780820151838201526020016102de565b50505050905090810190601f1680156103235780820380516001836020036101000a031916815260200191505b509550505050505060405180910390f35b341561033f57600080fd5b6100e1610dc0565b341561035257600080fd5b6101a0610fc4565b341561036557600080fd5b610379600160a060020a036004351661102a565b604051901515815260200160405180910390f35b341561039857600080fd5b610219611097565b34156103ab57600080fd5b61021960046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965061109d95505050505050565b34156103fc57600080fd5b61037960046024813581810190830135806020601f82018190048102016040519081016040528181529291906020840183838082843782019150505050505091908035151590602001909190803590602001908201803590602001908080601f0160208091040260200160405190810160405281815292919060208401838380828437509496506110aa95505050505050565b341561049a57600080fd5b610379600160a060020a036004351661123e565b6104b6611260565b6104be611260565b60008060006104cb611260565b6104d3611260565b60025460035490955093506000808587016040518059106104f15750595b908082528060200260200182016040525093508587016040518059106105145750595b90808252806020026020018201604052509250600091505b868260ff1610156106f7576002805460ff841690811061054857fe5b6000918252602080832090910154600160a060020a03168083526001808352604093849020805492995061061894909360026000199385161561010002939093019093169190910491601f83018190048102019051908101604052809291908181526020018280546001816001161561010002031660029004801561060e5780601f106105e35761010080835404028352916020019161060e565b820191906000526020600020905b8154815290600101906020018083116105f157829003601f168201915b505050505061109d565b848360ff168151811061062757fe5b9060200190602002019060001916908160001916815250506106d36001600087600160a060020a0316600160a060020a031681526020019081526020016000206002018054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561060e5780601f106105e35761010080835404028352916020019161060e565b838360ff16815181106106e257fe5b6020908102909101015260019091019061052c565b5060005b858160ff161015610892576003805460ff831690811061071757fe5b6000918252602080832090910154600160a060020a0316808352600180835260409384902080549299506107b294909360026000199385161561010002939093019093169190910491601f83018190048102019051908101604052809291908181526020018280546001816001161561010002031660029004801561060e5780601f106105e35761010080835404028352916020019161060e565b848260ff168901815181106107c357fe5b90602001906020020190600019169081600019168152505061086f6001600087600160a060020a0316600160a060020a031681526020019081526020016000206002018054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561060e5780601f106105e35761010080835404028352916020019161060e565b838260ff1689018151811061088057fe5b602090810290910101526001016106fb565b509197909650945050505050565b6108a8611260565b6108b0611260565b6000806108bb611260565b6108c3611260565b60025493506000846040518059106108d85750595b90808252806020026020018201604052509250846040518059106108f95750595b90808252806020026020018201604052509150600090505b848160ff161015610aa4576002805460ff831690811061092d57fe5b6000918252602080832090910154600160a060020a0316808352600180835260409384902080549298506109c894909360026000199385161561010002939093019093169190910491601f83018190048102019051908101604052809291908181526020018280546001816001161561010002031660029004801561060e5780601f106105e35761010080835404028352916020019161060e565b838260ff16815181106109d757fe5b906020019060200201906000191690816000191681525050610a836001600086600160a060020a0316600160a060020a031681526020019081526020016000206002018054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561060e5780601f106105e35761010080835404028352916020019161060e565b828260ff1681518110610a9257fe5b60209081029091010152600101610911565b50909590945092505050565b610ab8611260565b600080610ac3611260565b6003546002549093509150600080838501604051805910610ae15750595b90808252806020026020018201604052509250600091505b848260ff161015610b5b576003805460ff8416908110610b1557fe5b600091825260209091200154600160a060020a03168360ff841681518110610b3957fe5b600160a060020a03909216602092830290910190910152600190910190610af9565b5060005b838160ff161015610bc0576002805460ff8316908110610b7b57fe5b600091825260209091200154600160a060020a03168360ff8316870181518110610ba157fe5b600160a060020a03909216602092830290910190910152600101610b5f565b5090949350505050565b610bd2611260565b6002805480602002602001604051908101604052809291908181526020018280548015610c2857602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610c0a575b505050505090505b90565b60035490565b610c41611260565b6000610c4b611260565b600160a060020a038416600090815260016020818152604092839020808301548154919460ff9091169360028087019487946000199381161561010002939093019092160491601f8301829004820290910190519081016040528092919081815260200182805460018160011615610100020316600290048015610d105780601f10610ce557610100808354040283529160200191610d10565b820191906000526020600020905b815481529060010190602001808311610cf357829003601f168201915b50505050509250808054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610dac5780601f10610d8157610100808354040283529160200191610dac565b820191906000526020600020905b815481529060010190602001808311610d8f57829003601f168201915b505050505090509250925092509193909250565b610dc8611260565b610dd0611260565b600080610ddb611260565b610de3611260565b6003549350600084604051805910610df85750595b9080825280602002602001820160405250925084604051805910610e195750595b90808252806020026020018201604052509150600090505b848160ff161015610aa4576003805460ff8316908110610e4d57fe5b6000918252602080832090910154600160a060020a031680835260018083526040938490208054929850610ee894909360026000199385161561010002939093019093169190910491601f83018190048102019051908101604052809291908181526020018280546001816001161561010002031660029004801561060e5780601f106105e35761010080835404028352916020019161060e565b838260ff1681518110610ef757fe5b906020019060200201906000191690816000191681525050610fa36001600086600160a060020a0316600160a060020a031681526020019081526020016000206002018054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561060e5780601f106105e35761010080835404028352916020019161060e565b828260ff1681518110610fb257fe5b60209081029091010152600101610e31565b610fcc611260565b6003805480602002602001604051908101604052809291908181526020018280548015610c2857602002820191906000526020600020908154600160a060020a03168152600190910190602001808311610c0a575050505050905090565b6000805433600160a060020a0390811691161461104657600080fd5b600160a060020a0382166000908152600160205260408120906110698282611272565b60018201805460ff19169055611083600283016000611272565b50600301805460ff19169055506001919050565b60025490565b6000602082015192915050565b6000806110b56112b9565b33600160a060020a03811660009081526001602081905260409091206003015491935060ff909116151514156110ee5760009250611235565b855115156110ff5760009250611235565b8581528415156020808301919091526040808301869052600160608401819052600160a060020a03851660009081529252902081908151819080516111489291602001906112ee565b50602082015160018201805460ff191691151591909117905560408201518160020190805161117b9291602001906112ee565b506060820151600391909101805460ff191691151591909117905550841515600114156111eb5760038054600181016111b4838261136c565b506000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038416179055611230565b60028054600181016111fd838261136c565b506000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0384161790555b600192505b50509392505050565b600160a060020a03166000908152600160208190526040909120015460ff1690565b60206040519081016040526000815290565b50805460018160011615610100020316600290046000825580601f1061129857506112b6565b601f0160209004906000526020600020908101906112b69190611395565b50565b6080604051908101604052806112cd611260565b8152600060208201526040016112e1611260565b8152600060209091015290565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061132f57805160ff191683800117855561135c565b8280016001018555821561135c579182015b8281111561135c578251825591602001919060010190611341565b50611368929150611395565b5090565b81548183558181151161139057600083815260209020611390918101908301611395565b505050565b610c3091905b80821115611368576000815560010161139b5600a165627a7a723058202ffd0aee33f489feba11165277ce1a09e0c524c925624bc0f2ad8508945c86c70029`

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

// GetAllUsersAddresses is a free data retrieval call binding the contract method 0x3a5aeec3.
//
// Solidity: function getAllUsersAddresses() constant returns(address[])
func (_XNOUsers *XNOUsersCaller) GetAllUsersAddresses(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _XNOUsers.contract.Call(opts, out, "getAllUsersAddresses")
	return *ret0, err
}

// GetAllUsersAddresses is a free data retrieval call binding the contract method 0x3a5aeec3.
//
// Solidity: function getAllUsersAddresses() constant returns(address[])
func (_XNOUsers *XNOUsersSession) GetAllUsersAddresses() ([]common.Address, error) {
	return _XNOUsers.Contract.GetAllUsersAddresses(&_XNOUsers.CallOpts)
}

// GetAllUsersAddresses is a free data retrieval call binding the contract method 0x3a5aeec3.
//
// Solidity: function getAllUsersAddresses() constant returns(address[])
func (_XNOUsers *XNOUsersCallerSession) GetAllUsersAddresses() ([]common.Address, error) {
	return _XNOUsers.Contract.GetAllUsersAddresses(&_XNOUsers.CallOpts)
}

// GetGamers is a free data retrieval call binding the contract method 0x0d2e017b.
//
// Solidity: function getGamers() constant returns(bytes32[], bytes32[])
func (_XNOUsers *XNOUsersCaller) GetGamers(opts *bind.CallOpts) ([][32]byte, [][32]byte, error) {
	var (
		ret0 = new([][32]byte)
		ret1 = new([][32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _XNOUsers.contract.Call(opts, out, "getGamers")
	return *ret0, *ret1, err
}

// GetGamers is a free data retrieval call binding the contract method 0x0d2e017b.
//
// Solidity: function getGamers() constant returns(bytes32[], bytes32[])
func (_XNOUsers *XNOUsersSession) GetGamers() ([][32]byte, [][32]byte, error) {
	return _XNOUsers.Contract.GetGamers(&_XNOUsers.CallOpts)
}

// GetGamers is a free data retrieval call binding the contract method 0x0d2e017b.
//
// Solidity: function getGamers() constant returns(bytes32[], bytes32[])
func (_XNOUsers *XNOUsersCallerSession) GetGamers() ([][32]byte, [][32]byte, error) {
	return _XNOUsers.Contract.GetGamers(&_XNOUsers.CallOpts)
}

// GetGamersAddresses is a free data retrieval call binding the contract method 0x6603e8ac.
//
// Solidity: function getGamersAddresses() constant returns(address[])
func (_XNOUsers *XNOUsersCaller) GetGamersAddresses(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _XNOUsers.contract.Call(opts, out, "getGamersAddresses")
	return *ret0, err
}

// GetGamersAddresses is a free data retrieval call binding the contract method 0x6603e8ac.
//
// Solidity: function getGamersAddresses() constant returns(address[])
func (_XNOUsers *XNOUsersSession) GetGamersAddresses() ([]common.Address, error) {
	return _XNOUsers.Contract.GetGamersAddresses(&_XNOUsers.CallOpts)
}

// GetGamersAddresses is a free data retrieval call binding the contract method 0x6603e8ac.
//
// Solidity: function getGamersAddresses() constant returns(address[])
func (_XNOUsers *XNOUsersCallerSession) GetGamersAddresses() ([]common.Address, error) {
	return _XNOUsers.Contract.GetGamersAddresses(&_XNOUsers.CallOpts)
}

// GetGamersNumber is a free data retrieval call binding the contract method 0xae4191c8.
//
// Solidity: function getGamersNumber() constant returns(uint256)
func (_XNOUsers *XNOUsersCaller) GetGamersNumber(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _XNOUsers.contract.Call(opts, out, "getGamersNumber")
	return *ret0, err
}

// GetGamersNumber is a free data retrieval call binding the contract method 0xae4191c8.
//
// Solidity: function getGamersNumber() constant returns(uint256)
func (_XNOUsers *XNOUsersSession) GetGamersNumber() (*big.Int, error) {
	return _XNOUsers.Contract.GetGamersNumber(&_XNOUsers.CallOpts)
}

// GetGamersNumber is a free data retrieval call binding the contract method 0xae4191c8.
//
// Solidity: function getGamersNumber() constant returns(uint256)
func (_XNOUsers *XNOUsersCallerSession) GetGamersNumber() (*big.Int, error) {
	return _XNOUsers.Contract.GetGamersNumber(&_XNOUsers.CallOpts)
}

// GetServers is a free data retrieval call binding the contract method 0x75c1c7e7.
//
// Solidity: function getServers() constant returns(bytes32[], bytes32[])
func (_XNOUsers *XNOUsersCaller) GetServers(opts *bind.CallOpts) ([][32]byte, [][32]byte, error) {
	var (
		ret0 = new([][32]byte)
		ret1 = new([][32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _XNOUsers.contract.Call(opts, out, "getServers")
	return *ret0, *ret1, err
}

// GetServers is a free data retrieval call binding the contract method 0x75c1c7e7.
//
// Solidity: function getServers() constant returns(bytes32[], bytes32[])
func (_XNOUsers *XNOUsersSession) GetServers() ([][32]byte, [][32]byte, error) {
	return _XNOUsers.Contract.GetServers(&_XNOUsers.CallOpts)
}

// GetServers is a free data retrieval call binding the contract method 0x75c1c7e7.
//
// Solidity: function getServers() constant returns(bytes32[], bytes32[])
func (_XNOUsers *XNOUsersCallerSession) GetServers() ([][32]byte, [][32]byte, error) {
	return _XNOUsers.Contract.GetServers(&_XNOUsers.CallOpts)
}

// GetServersAddresses is a free data retrieval call binding the contract method 0x7d22f9ab.
//
// Solidity: function getServersAddresses() constant returns(address[])
func (_XNOUsers *XNOUsersCaller) GetServersAddresses(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _XNOUsers.contract.Call(opts, out, "getServersAddresses")
	return *ret0, err
}

// GetServersAddresses is a free data retrieval call binding the contract method 0x7d22f9ab.
//
// Solidity: function getServersAddresses() constant returns(address[])
func (_XNOUsers *XNOUsersSession) GetServersAddresses() ([]common.Address, error) {
	return _XNOUsers.Contract.GetServersAddresses(&_XNOUsers.CallOpts)
}

// GetServersAddresses is a free data retrieval call binding the contract method 0x7d22f9ab.
//
// Solidity: function getServersAddresses() constant returns(address[])
func (_XNOUsers *XNOUsersCallerSession) GetServersAddresses() ([]common.Address, error) {
	return _XNOUsers.Contract.GetServersAddresses(&_XNOUsers.CallOpts)
}

// GetServersNumber is a free data retrieval call binding the contract method 0x6ecc46e0.
//
// Solidity: function getServersNumber() constant returns(uint256)
func (_XNOUsers *XNOUsersCaller) GetServersNumber(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _XNOUsers.contract.Call(opts, out, "getServersNumber")
	return *ret0, err
}

// GetServersNumber is a free data retrieval call binding the contract method 0x6ecc46e0.
//
// Solidity: function getServersNumber() constant returns(uint256)
func (_XNOUsers *XNOUsersSession) GetServersNumber() (*big.Int, error) {
	return _XNOUsers.Contract.GetServersNumber(&_XNOUsers.CallOpts)
}

// GetServersNumber is a free data retrieval call binding the contract method 0x6ecc46e0.
//
// Solidity: function getServersNumber() constant returns(uint256)
func (_XNOUsers *XNOUsersCallerSession) GetServersNumber() (*big.Int, error) {
	return _XNOUsers.Contract.GetServersNumber(&_XNOUsers.CallOpts)
}

// GetUser is a free data retrieval call binding the contract method 0x6f77926b.
//
// Solidity: function getUser(_address address) constant returns(string, bool, string)
func (_XNOUsers *XNOUsersCaller) GetUser(opts *bind.CallOpts, _address common.Address) (string, bool, string, error) {
	var (
		ret0 = new(string)
		ret1 = new(bool)
		ret2 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _XNOUsers.contract.Call(opts, out, "getUser", _address)
	return *ret0, *ret1, *ret2, err
}

// GetUser is a free data retrieval call binding the contract method 0x6f77926b.
//
// Solidity: function getUser(_address address) constant returns(string, bool, string)
func (_XNOUsers *XNOUsersSession) GetUser(_address common.Address) (string, bool, string, error) {
	return _XNOUsers.Contract.GetUser(&_XNOUsers.CallOpts, _address)
}

// GetUser is a free data retrieval call binding the contract method 0x6f77926b.
//
// Solidity: function getUser(_address address) constant returns(string, bool, string)
func (_XNOUsers *XNOUsersCallerSession) GetUser(_address common.Address) (string, bool, string, error) {
	return _XNOUsers.Contract.GetUser(&_XNOUsers.CallOpts, _address)
}

// GetUsers is a free data retrieval call binding the contract method 0x00ce8e3e.
//
// Solidity: function getUsers() constant returns(bytes32[], bytes32[])
func (_XNOUsers *XNOUsersCaller) GetUsers(opts *bind.CallOpts) ([][32]byte, [][32]byte, error) {
	var (
		ret0 = new([][32]byte)
		ret1 = new([][32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _XNOUsers.contract.Call(opts, out, "getUsers")
	return *ret0, *ret1, err
}

// GetUsers is a free data retrieval call binding the contract method 0x00ce8e3e.
//
// Solidity: function getUsers() constant returns(bytes32[], bytes32[])
func (_XNOUsers *XNOUsersSession) GetUsers() ([][32]byte, [][32]byte, error) {
	return _XNOUsers.Contract.GetUsers(&_XNOUsers.CallOpts)
}

// GetUsers is a free data retrieval call binding the contract method 0x00ce8e3e.
//
// Solidity: function getUsers() constant returns(bytes32[], bytes32[])
func (_XNOUsers *XNOUsersCallerSession) GetUsers() ([][32]byte, [][32]byte, error) {
	return _XNOUsers.Contract.GetUsers(&_XNOUsers.CallOpts)
}

// IsServer is a free data retrieval call binding the contract method 0xfda08958.
//
// Solidity: function isServer(_address address) constant returns(bool)
func (_XNOUsers *XNOUsersCaller) IsServer(opts *bind.CallOpts, _address common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _XNOUsers.contract.Call(opts, out, "isServer", _address)
	return *ret0, err
}

// IsServer is a free data retrieval call binding the contract method 0xfda08958.
//
// Solidity: function isServer(_address address) constant returns(bool)
func (_XNOUsers *XNOUsersSession) IsServer(_address common.Address) (bool, error) {
	return _XNOUsers.Contract.IsServer(&_XNOUsers.CallOpts, _address)
}

// IsServer is a free data retrieval call binding the contract method 0xfda08958.
//
// Solidity: function isServer(_address address) constant returns(bool)
func (_XNOUsers *XNOUsersCallerSession) IsServer(_address common.Address) (bool, error) {
	return _XNOUsers.Contract.IsServer(&_XNOUsers.CallOpts, _address)
}

// RegisterNewUser is a paid mutator transaction binding the contract method 0xe5e8093a.
//
// Solidity: function registerNewUser(_name string, _isServer bool, _game string) returns(success bool)
func (_XNOUsers *XNOUsersTransactor) RegisterNewUser(opts *bind.TransactOpts, _name string, _isServer bool, _game string) (*types.Transaction, error) {
	return _XNOUsers.contract.Transact(opts, "registerNewUser", _name, _isServer, _game)
}

// RegisterNewUser is a paid mutator transaction binding the contract method 0xe5e8093a.
//
// Solidity: function registerNewUser(_name string, _isServer bool, _game string) returns(success bool)
func (_XNOUsers *XNOUsersSession) RegisterNewUser(_name string, _isServer bool, _game string) (*types.Transaction, error) {
	return _XNOUsers.Contract.RegisterNewUser(&_XNOUsers.TransactOpts, _name, _isServer, _game)
}

// RegisterNewUser is a paid mutator transaction binding the contract method 0xe5e8093a.
//
// Solidity: function registerNewUser(_name string, _isServer bool, _game string) returns(success bool)
func (_XNOUsers *XNOUsersTransactorSession) RegisterNewUser(_name string, _isServer bool, _game string) (*types.Transaction, error) {
	return _XNOUsers.Contract.RegisterNewUser(&_XNOUsers.TransactOpts, _name, _isServer, _game)
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

// StringToBytes32 is a paid mutator transaction binding the contract method 0xcfb51928.
//
// Solidity: function stringToBytes32(source string) returns(result bytes32)
func (_XNOUsers *XNOUsersTransactor) StringToBytes32(opts *bind.TransactOpts, source string) (*types.Transaction, error) {
	return _XNOUsers.contract.Transact(opts, "stringToBytes32", source)
}

// StringToBytes32 is a paid mutator transaction binding the contract method 0xcfb51928.
//
// Solidity: function stringToBytes32(source string) returns(result bytes32)
func (_XNOUsers *XNOUsersSession) StringToBytes32(source string) (*types.Transaction, error) {
	return _XNOUsers.Contract.StringToBytes32(&_XNOUsers.TransactOpts, source)
}

// StringToBytes32 is a paid mutator transaction binding the contract method 0xcfb51928.
//
// Solidity: function stringToBytes32(source string) returns(result bytes32)
func (_XNOUsers *XNOUsersTransactorSession) StringToBytes32(source string) (*types.Transaction, error) {
	return _XNOUsers.Contract.StringToBytes32(&_XNOUsers.TransactOpts, source)
}
