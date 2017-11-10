// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package xnogames

import (
	"math/big"
	"strings"

	"github.com/xenioplatform/go-xenio/common"
	"github.com/xenioplatform/go-xenio/core/types"
	"github.com/xenioplatform/go-xenio/accounts/abi"
	"github.com/xenioplatform/go-xenio/accounts/abi/bind"
)

// XNOGamesABI is the input ABI used to generate the binding from.
const XNOGamesABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"imageURL\",\"type\":\"string\"},{\"name\":\"SHA256notaryHash\",\"type\":\"bytes32\"}],\"name\":\"addLogo\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getGamesAddresses\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"userAddress\",\"type\":\"address\"}],\"name\":\"getGameImages\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"SHA256notaryHash\",\"type\":\"bytes32\"}],\"name\":\"getImage\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAllImages\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"gameAddress\",\"type\":\"address\"}],\"name\":\"getGame\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"gameAddress\",\"type\":\"address\"}],\"name\":\"removeUser\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"imgHash\",\"type\":\"bytes32\"}],\"name\":\"removeImage\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_publisher\",\"type\":\"string\"},{\"name\":\"_developer\",\"type\":\"string\"},{\"name\":\"_country\",\"type\":\"bytes32\"},{\"name\":\"_state\",\"type\":\"bytes32\"}],\"name\":\"registerNewGame\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getGames\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\"},{\"name\":\"\",\"type\":\"bytes32[]\"},{\"name\":\"\",\"type\":\"bytes32[]\"},{\"name\":\"\",\"type\":\"bytes32[]\"},{\"name\":\"\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"source\",\"type\":\"string\"}],\"name\":\"stringToBytes32\",\"outputs\":[{\"name\":\"result\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"}]"

// XNOGamesBin is the compiled bytecode used for deploying new contracts.
const XNOGamesBin = `0x606060405260008054600160a060020a033316600160a060020a0319909116179055611387806100306000396000f300606060405236156100ac5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416632594677581146100b9578063281b9f771461012057806334c16694146101865780636ced1ae9146101b75780637b7284fc1461024c5780638d1a4b281461025f57806398575188146103e15780639f6c935814610400578063ad45644114610416578063c04c5947146104f2578063cfb519281461066d575b34156100b757600080fd5b005b34156100c457600080fd5b61010c60046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965050933593506106be92505050565b604051901515815260200160405180910390f35b341561012b57600080fd5b6101336107b7565b60405160208082528190810183818151815260200191508051906020019060200280838360005b8381101561017257808201518382015260200161015a565b505050509050019250505060405180910390f35b341561019157600080fd5b6101a5600160a060020a0360043516610820565b60405190815260200160405180910390f35b34156101c257600080fd5b6101cd60043561083e565b6040518080602001838152602001828103825284818151815260200191508051906020019080838360005b838110156102105780820151838201526020016101f8565b50505050905090810190601f16801561023d5780820380516001836020036101000a031916815260200191505b50935050505060405180910390f35b341561025757600080fd5b61013361090b565b341561026a57600080fd5b61027e600160a060020a0360043516610969565b604051606081018490526080810183905260a0810182905260c080825281906020820190604083019083018a818151815260200191508051906020019080838360005b838110156102d95780820151838201526020016102c1565b50505050905090810190601f1680156103065780820380516001836020036101000a031916815260200191505b50848103835289818151815260200191508051906020019080838360005b8381101561033c578082015183820152602001610324565b50505050905090810190601f1680156103695780820380516001836020036101000a031916815260200191505b50848103825288818151815260200191508051906020019080838360005b8381101561039f578082015183820152602001610387565b50505050905090810190601f1680156103cc5780820380516001836020036101000a031916815260200191505b50995050505050505050505060405180910390f35b34156103ec57600080fd5b61010c600160a060020a0360043516610ba3565b341561040b57600080fd5b61010c600435610c27565b341561042157600080fd5b61010c60046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284375094965050843594602001359350610c6b92505050565b34156104fd57600080fd5b610505610dd7565b60405180806020018060200180602001806020018060200186810386528b818151815260200191508051906020019060200280838360005b8381101561055557808201518382015260200161053d565b5050505090500186810385528a818151815260200191508051906020019060200280838360005b8381101561059457808201518382015260200161057c565b50505050905001868103845289818151815260200191508051906020019060200280838360005b838110156105d35780820151838201526020016105bb565b50505050905001868103835288818151815260200191508051906020019060200280838360005b838110156106125780820151838201526020016105fa565b50505050905001868103825287818151815260200191508051906020019060200280838360005b83811015610651578082015183820152602001610639565b505050509050019a505050505050505050505060405180910390f35b341561067857600080fd5b6101a560046024813581810190830135806020601f820181900481020160405190810160405281815292919060208401838380828437509496506111e495505050505050565b33600160a060020a038116600090815260016020819052604082205491929160029181161561010002600019011604156106fb57600091506107b0565b83511561070b57600091506107b0565b600083815260036020526040902054600261010060018316150260001901909116041561073b57600091506107b0565b600480546001810161074d83826111f1565b506000918252602080832091909101859055848252600390526040902084805161077b92916020019061121a565b50600083815260036020908152604080832042600191820155600160a060020a03851684529182905290912060050184905591505b5092915050565b6107bf611298565b600280548060200260200160405190810160405280929190818152602001828054801561081557602002820191906000526020600020905b8154600160a060020a031681526001909101906020018083116107f7575b505050505090505b90565b600160a060020a031660009081526001602052604090206005015490565b610846611298565b600082815260036020908152604080832060018082015482549294909385936002610100948216159490940260001901169290920491601f83018290048202909101905190810160405280929190818152602001828054600181600116156101000203166002900480156108fb5780601f106108d0576101008083540402835291602001916108fb565b820191906000526020600020905b8154815290600101906020018083116108de57829003601f168201915b5050505050915091509150915091565b610913611298565b600480548060200260200160405190810160405280929190818152602001828054801561081557602002820191906000526020600020905b8154815260019091019060200180831161094b575050505050905090565b610971611298565b610979611298565b610981611298565b600160a060020a03841660009081526001602081815260408084206003810154600482015460058301548354889788979596878201966002808a01979695948a948116156101000260001901160491601f8301829004820290910190519081016040528092919081815260200182805460018160011615610100020316600290048015610a4f5780601f10610a2457610100808354040283529160200191610a4f565b820191906000526020600020905b815481529060010190602001808311610a3257829003601f168201915b50505050509550848054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610aeb5780601f10610ac057610100808354040283529160200191610aeb565b820191906000526020600020905b815481529060010190602001808311610ace57829003601f168201915b50505050509450838054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610b875780601f10610b5c57610100808354040283529160200191610b87565b820191906000526020600020905b815481529060010190602001808311610b6a57829003601f168201915b5050505050935095509550955095509550955091939550919395565b6000805433600160a060020a03908116911614610bbf57600080fd5b600160a060020a038216600090815260016020526040812090610be282826112aa565b610bf06001830160006112aa565b610bfe6002830160006112aa565b50600060038201819055600482018190556005820155600601805460ff19169055506001919050565b6000805433600160a060020a03908116911614610c4357600080fd5b600082815260036020526040812090610c5c82826112aa565b50600060019182015592915050565b600080610c766112f1565b33600160a060020a03811660009081526001602081905260409091206006015491935060ff90911615151415610caf5760009250610dcc565b87511515610cc05760009250610dcc565b878152602080820188905260408083018890526060830187905260808301869052600160c08401819052600160a060020a0385166000908152925290208190815181908051610d1392916020019061121a565b50602082015181600101908051610d2e92916020019061121a565b50604082015181600201908051610d4992916020019061121a565b50606082015160038201556080820151600482015560a0820151600582015560c0820151600691909101805460ff1916911515919091179055506002805460018101610d9583826111f1565b506000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038416179055600192505b505095945050505050565b610ddf611298565b610de7611298565b610def611298565b610df7611298565b610dff611298565b600080610e0a611298565b610e12611298565b610e1a611298565b610e22611298565b610e2a611298565b6002549650600087604051805910610e3f5750595b9080825280602002602001820160405250955087604051805910610e605750595b9080825280602002602001820160405250945087604051805910610e815750595b9080825280602002602001820160405250935087604051805910610ea25750595b9080825280602002602001820160405250925087604051805910610ec35750595b90808252806020026020018201604052509150600090505b878160ff1610156111d0576002805460ff8316908110610ef757fe5b6000918252602080832090910154600160a060020a031680835260018083526040938490208054929b50610fc794909360026000199385161561010002939093019093169190910491601f830181900481020190519081016040528092919081815260200182805460018160011615610100020316600290048015610fbd5780601f10610f9257610100808354040283529160200191610fbd565b820191906000526020600020905b815481529060010190602001808311610fa057829003601f168201915b50505050506111e4565b868260ff1681518110610fd657fe5b9060200190602002019060001916908160001916815250506110826001600089600160a060020a0316600160a060020a031681526020019081526020016000206001018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610fbd5780601f10610f9257610100808354040283529160200191610fbd565b858260ff168151811061109157fe5b90602001906020020190600019169081600019168152505061113d6001600089600160a060020a0316600160a060020a031681526020019081526020016000206002018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610fbd5780601f10610f9257610100808354040283529160200191610fbd565b848260ff168151811061114c57fe5b6020908102909101810191909152600160a060020a0388166000908152600190915260409020600301548360ff83168151811061118557fe5b6020908102909101810191909152600160a060020a0388166000908152600190915260409020600401548260ff8316815181106111be57fe5b60209081029091010152600101610edb565b50939b929a50909850965090945092505050565b6000602082015192915050565b81548183558181151161121557600083815260209020611215918101908301611341565b505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061125b57805160ff1916838001178555611288565b82800160010185558215611288579182015b8281111561128857825182559160200191906001019061126d565b50611294929150611341565b5090565b60206040519081016040526000815290565b50805460018160011615610100020316600290046000825580601f106112d057506112ee565b601f0160209004906000526020600020908101906112ee9190611341565b50565b60e060405190810160405280611305611298565b8152602001611312611298565b815260200161131f611298565b8152600060208201819052604082018190526060820181905260809091015290565b61081d91905b8082111561129457600081556001016113475600a165627a7a723058209c1e37add880e334526080ef9d981656549a710186f7991f440d9ffb0bfa7b730029`

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

// GetAllImages is a free data retrieval call binding the contract method 0x7b7284fc.
//
// Solidity: function getAllImages() constant returns(bytes32[])
func (_XNOGames *XNOGamesCaller) GetAllImages(opts *bind.CallOpts) ([][]byte, error) {
	var (
		ret0 = new([][]byte)
	)
	out := ret0
	err := _XNOGames.contract.Call(opts, out, "getAllImages")
	return *ret0, err
}

// GetAllImages is a free data retrieval call binding the contract method 0x7b7284fc.
//
// Solidity: function getAllImages() constant returns(bytes32[])
func (_XNOGames *XNOGamesSession) GetAllImages() ([][]byte, error) {
	return _XNOGames.Contract.GetAllImages(&_XNOGames.CallOpts)
}

// GetAllImages is a free data retrieval call binding the contract method 0x7b7284fc.
//
// Solidity: function getAllImages() constant returns(bytes32[])
func (_XNOGames *XNOGamesCallerSession) GetAllImages() ([][]byte, error) {
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

// GetGames is a free data retrieval call binding the contract method 0xc04c5947.
//
// Solidity: function getGames() constant returns(bytes32[], bytes32[], bytes32[], bytes32[], bytes32[])
func (_XNOGames *XNOGamesCaller) GetGames(opts *bind.CallOpts) ([][32]byte, [][32]byte, [][32]byte, [][32]byte, [][32]byte, error) {
	var (
		ret0 = new([][32]byte)
		ret1 = new([][32]byte)
		ret2 = new([][32]byte)
		ret3 = new([][32]byte)
		ret4 = new([][32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
	}
	err := _XNOGames.contract.Call(opts, out, "getGames")
	return *ret0, *ret1, *ret2, *ret3, *ret4, err
}

// GetGames is a free data retrieval call binding the contract method 0xc04c5947.
//
// Solidity: function getGames() constant returns(bytes32[], bytes32[], bytes32[], bytes32[], bytes32[])
func (_XNOGames *XNOGamesSession) GetGames() ([][32]byte, [][32]byte, [][32]byte, [][32]byte, [][32]byte, error) {
	return _XNOGames.Contract.GetGames(&_XNOGames.CallOpts)
}

// GetGames is a free data retrieval call binding the contract method 0xc04c5947.
//
// Solidity: function getGames() constant returns(bytes32[], bytes32[], bytes32[], bytes32[], bytes32[])
func (_XNOGames *XNOGamesCallerSession) GetGames() ([][32]byte, [][32]byte, [][32]byte, [][32]byte, [][32]byte, error) {
	return _XNOGames.Contract.GetGames(&_XNOGames.CallOpts)
}

// GetGamesAddresses is a free data retrieval call binding the contract method 0x281b9f77.
//
// Solidity: function getGamesAddresses() constant returns(address[])
func (_XNOGames *XNOGamesCaller) GetGamesAddresses(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _XNOGames.contract.Call(opts, out, "getGamesAddresses")
	return *ret0, err
}

// GetGamesAddresses is a free data retrieval call binding the contract method 0x281b9f77.
//
// Solidity: function getGamesAddresses() constant returns(address[])
func (_XNOGames *XNOGamesSession) GetGamesAddresses() ([]common.Address, error) {
	return _XNOGames.Contract.GetGamesAddresses(&_XNOGames.CallOpts)
}

// GetGamesAddresses is a free data retrieval call binding the contract method 0x281b9f77.
//
// Solidity: function getGamesAddresses() constant returns(address[])
func (_XNOGames *XNOGamesCallerSession) GetGamesAddresses() ([]common.Address, error) {
	return _XNOGames.Contract.GetGamesAddresses(&_XNOGames.CallOpts)
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
// Solidity: function registerNewGame(_name string, _publisher string, _developer string, _country bytes32, _state bytes32) returns(success bool)
func (_XNOGames *XNOGamesTransactor) RegisterNewGame(opts *bind.TransactOpts, _name string, _publisher string, _developer string, _country [32]byte, _state [32]byte) (*types.Transaction, error) {
	return _XNOGames.contract.Transact(opts, "registerNewGame", _name, _publisher, _developer, _country, _state)
}

// RegisterNewGame is a paid mutator transaction binding the contract method 0xad456441.
//
// Solidity: function registerNewGame(_name string, _publisher string, _developer string, _country bytes32, _state bytes32) returns(success bool)
func (_XNOGames *XNOGamesSession) RegisterNewGame(_name string, _publisher string, _developer string, _country [32]byte, _state [32]byte) (*types.Transaction, error) {
	return _XNOGames.Contract.RegisterNewGame(&_XNOGames.TransactOpts, _name, _publisher, _developer, _country, _state)
}

// RegisterNewGame is a paid mutator transaction binding the contract method 0xad456441.
//
// Solidity: function registerNewGame(_name string, _publisher string, _developer string, _country bytes32, _state bytes32) returns(success bool)
func (_XNOGames *XNOGamesTransactorSession) RegisterNewGame(_name string, _publisher string, _developer string, _country [32]byte, _state [32]byte) (*types.Transaction, error) {
	return _XNOGames.Contract.RegisterNewGame(&_XNOGames.TransactOpts, _name, _publisher, _developer, _country, _state)
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

// StringToBytes32 is a paid mutator transaction binding the contract method 0xcfb51928.
//
// Solidity: function stringToBytes32(source string) returns(result bytes32)
func (_XNOGames *XNOGamesTransactor) StringToBytes32(opts *bind.TransactOpts, source string) (*types.Transaction, error) {
	return _XNOGames.contract.Transact(opts, "stringToBytes32", source)
}

// StringToBytes32 is a paid mutator transaction binding the contract method 0xcfb51928.
//
// Solidity: function stringToBytes32(source string) returns(result bytes32)
func (_XNOGames *XNOGamesSession) StringToBytes32(source string) (*types.Transaction, error) {
	return _XNOGames.Contract.StringToBytes32(&_XNOGames.TransactOpts, source)
}

// StringToBytes32 is a paid mutator transaction binding the contract method 0xcfb51928.
//
// Solidity: function stringToBytes32(source string) returns(result bytes32)
func (_XNOGames *XNOGamesTransactorSession) StringToBytes32(source string) (*types.Transaction, error) {
	return _XNOGames.Contract.StringToBytes32(&_XNOGames.TransactOpts, source)
}
