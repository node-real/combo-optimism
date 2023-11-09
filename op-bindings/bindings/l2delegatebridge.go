// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// L2DelegateBridgeMetaData contains all meta data concerning the L2DelegateBridge contract.
var L2DelegateBridgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"L2_STANDARD_BRIDGE\",\"outputs\":[{\"internalType\":\"contractL2StandardBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"L2_STANDARD_BRIDGE_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_minGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"withdrawTo\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x6080604052600080546001600160a01b03191673420000000000000000000000000000000000001017905534801561003657600080fd5b506102e2806100466000396000f3fe6080604052600436106100345760003560e01c806321d12763146100395780632cb7cb061461008f578063a3a79548146100b7575b600080fd5b34801561004557600080fd5b506000546100669073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b34801561009b57600080fd5b5061006673420000000000000000000000000000000000001081565b6100ca6100c5366004610191565b6100cc565b005b6000546040517fa3a7954800000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091169063a3a7954890349061012e908a908a908a908a908a908a90600401610249565b6000604051808303818588803b15801561014757600080fd5b505af115801561015b573d6000803e3d6000fd5b5050505050505050505050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461018c57600080fd5b919050565b60008060008060008060a087890312156101aa57600080fd5b6101b387610168565b95506101c160208801610168565b945060408701359350606087013563ffffffff811681146101e157600080fd5b9250608087013567ffffffffffffffff808211156101fe57600080fd5b818901915089601f83011261021257600080fd5b81358181111561022157600080fd5b8a602082850101111561023357600080fd5b6020830194508093505050509295509295509295565b600073ffffffffffffffffffffffffffffffffffffffff808916835280881660208401525085604083015263ffffffff8516606083015260a060808301528260a0830152828460c0840137600060c0848401015260c07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f850116830101905097965050505050505056fea164736f6c634300080f000a",
}

// L2DelegateBridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use L2DelegateBridgeMetaData.ABI instead.
var L2DelegateBridgeABI = L2DelegateBridgeMetaData.ABI

// L2DelegateBridgeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use L2DelegateBridgeMetaData.Bin instead.
var L2DelegateBridgeBin = L2DelegateBridgeMetaData.Bin

// DeployL2DelegateBridge deploys a new Ethereum contract, binding an instance of L2DelegateBridge to it.
func DeployL2DelegateBridge(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *L2DelegateBridge, error) {
	parsed, err := L2DelegateBridgeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(L2DelegateBridgeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &L2DelegateBridge{L2DelegateBridgeCaller: L2DelegateBridgeCaller{contract: contract}, L2DelegateBridgeTransactor: L2DelegateBridgeTransactor{contract: contract}, L2DelegateBridgeFilterer: L2DelegateBridgeFilterer{contract: contract}}, nil
}

// L2DelegateBridge is an auto generated Go binding around an Ethereum contract.
type L2DelegateBridge struct {
	L2DelegateBridgeCaller     // Read-only binding to the contract
	L2DelegateBridgeTransactor // Write-only binding to the contract
	L2DelegateBridgeFilterer   // Log filterer for contract events
}

// L2DelegateBridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type L2DelegateBridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2DelegateBridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L2DelegateBridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2DelegateBridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L2DelegateBridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2DelegateBridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L2DelegateBridgeSession struct {
	Contract     *L2DelegateBridge // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L2DelegateBridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L2DelegateBridgeCallerSession struct {
	Contract *L2DelegateBridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// L2DelegateBridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L2DelegateBridgeTransactorSession struct {
	Contract     *L2DelegateBridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// L2DelegateBridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type L2DelegateBridgeRaw struct {
	Contract *L2DelegateBridge // Generic contract binding to access the raw methods on
}

// L2DelegateBridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L2DelegateBridgeCallerRaw struct {
	Contract *L2DelegateBridgeCaller // Generic read-only contract binding to access the raw methods on
}

// L2DelegateBridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L2DelegateBridgeTransactorRaw struct {
	Contract *L2DelegateBridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL2DelegateBridge creates a new instance of L2DelegateBridge, bound to a specific deployed contract.
func NewL2DelegateBridge(address common.Address, backend bind.ContractBackend) (*L2DelegateBridge, error) {
	contract, err := bindL2DelegateBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L2DelegateBridge{L2DelegateBridgeCaller: L2DelegateBridgeCaller{contract: contract}, L2DelegateBridgeTransactor: L2DelegateBridgeTransactor{contract: contract}, L2DelegateBridgeFilterer: L2DelegateBridgeFilterer{contract: contract}}, nil
}

// NewL2DelegateBridgeCaller creates a new read-only instance of L2DelegateBridge, bound to a specific deployed contract.
func NewL2DelegateBridgeCaller(address common.Address, caller bind.ContractCaller) (*L2DelegateBridgeCaller, error) {
	contract, err := bindL2DelegateBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L2DelegateBridgeCaller{contract: contract}, nil
}

// NewL2DelegateBridgeTransactor creates a new write-only instance of L2DelegateBridge, bound to a specific deployed contract.
func NewL2DelegateBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*L2DelegateBridgeTransactor, error) {
	contract, err := bindL2DelegateBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L2DelegateBridgeTransactor{contract: contract}, nil
}

// NewL2DelegateBridgeFilterer creates a new log filterer instance of L2DelegateBridge, bound to a specific deployed contract.
func NewL2DelegateBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*L2DelegateBridgeFilterer, error) {
	contract, err := bindL2DelegateBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L2DelegateBridgeFilterer{contract: contract}, nil
}

// bindL2DelegateBridge binds a generic wrapper to an already deployed contract.
func bindL2DelegateBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := L2DelegateBridgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2DelegateBridge *L2DelegateBridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2DelegateBridge.Contract.L2DelegateBridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2DelegateBridge *L2DelegateBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2DelegateBridge.Contract.L2DelegateBridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2DelegateBridge *L2DelegateBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2DelegateBridge.Contract.L2DelegateBridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2DelegateBridge *L2DelegateBridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2DelegateBridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2DelegateBridge *L2DelegateBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2DelegateBridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2DelegateBridge *L2DelegateBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2DelegateBridge.Contract.contract.Transact(opts, method, params...)
}

// L2STANDARDBRIDGE is a free data retrieval call binding the contract method 0x21d12763.
//
// Solidity: function L2_STANDARD_BRIDGE() view returns(address)
func (_L2DelegateBridge *L2DelegateBridgeCaller) L2STANDARDBRIDGE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2DelegateBridge.contract.Call(opts, &out, "L2_STANDARD_BRIDGE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2STANDARDBRIDGE is a free data retrieval call binding the contract method 0x21d12763.
//
// Solidity: function L2_STANDARD_BRIDGE() view returns(address)
func (_L2DelegateBridge *L2DelegateBridgeSession) L2STANDARDBRIDGE() (common.Address, error) {
	return _L2DelegateBridge.Contract.L2STANDARDBRIDGE(&_L2DelegateBridge.CallOpts)
}

// L2STANDARDBRIDGE is a free data retrieval call binding the contract method 0x21d12763.
//
// Solidity: function L2_STANDARD_BRIDGE() view returns(address)
func (_L2DelegateBridge *L2DelegateBridgeCallerSession) L2STANDARDBRIDGE() (common.Address, error) {
	return _L2DelegateBridge.Contract.L2STANDARDBRIDGE(&_L2DelegateBridge.CallOpts)
}

// L2STANDARDBRIDGEADDRESS is a free data retrieval call binding the contract method 0x2cb7cb06.
//
// Solidity: function L2_STANDARD_BRIDGE_ADDRESS() view returns(address)
func (_L2DelegateBridge *L2DelegateBridgeCaller) L2STANDARDBRIDGEADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2DelegateBridge.contract.Call(opts, &out, "L2_STANDARD_BRIDGE_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2STANDARDBRIDGEADDRESS is a free data retrieval call binding the contract method 0x2cb7cb06.
//
// Solidity: function L2_STANDARD_BRIDGE_ADDRESS() view returns(address)
func (_L2DelegateBridge *L2DelegateBridgeSession) L2STANDARDBRIDGEADDRESS() (common.Address, error) {
	return _L2DelegateBridge.Contract.L2STANDARDBRIDGEADDRESS(&_L2DelegateBridge.CallOpts)
}

// L2STANDARDBRIDGEADDRESS is a free data retrieval call binding the contract method 0x2cb7cb06.
//
// Solidity: function L2_STANDARD_BRIDGE_ADDRESS() view returns(address)
func (_L2DelegateBridge *L2DelegateBridgeCallerSession) L2STANDARDBRIDGEADDRESS() (common.Address, error) {
	return _L2DelegateBridge.Contract.L2STANDARDBRIDGEADDRESS(&_L2DelegateBridge.CallOpts)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0xa3a79548.
//
// Solidity: function withdrawTo(address _l2Token, address _to, uint256 _amount, uint32 _minGasLimit, bytes _extraData) payable returns()
func (_L2DelegateBridge *L2DelegateBridgeTransactor) WithdrawTo(opts *bind.TransactOpts, _l2Token common.Address, _to common.Address, _amount *big.Int, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _L2DelegateBridge.contract.Transact(opts, "withdrawTo", _l2Token, _to, _amount, _minGasLimit, _extraData)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0xa3a79548.
//
// Solidity: function withdrawTo(address _l2Token, address _to, uint256 _amount, uint32 _minGasLimit, bytes _extraData) payable returns()
func (_L2DelegateBridge *L2DelegateBridgeSession) WithdrawTo(_l2Token common.Address, _to common.Address, _amount *big.Int, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _L2DelegateBridge.Contract.WithdrawTo(&_L2DelegateBridge.TransactOpts, _l2Token, _to, _amount, _minGasLimit, _extraData)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0xa3a79548.
//
// Solidity: function withdrawTo(address _l2Token, address _to, uint256 _amount, uint32 _minGasLimit, bytes _extraData) payable returns()
func (_L2DelegateBridge *L2DelegateBridgeTransactorSession) WithdrawTo(_l2Token common.Address, _to common.Address, _amount *big.Int, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _L2DelegateBridge.Contract.WithdrawTo(&_L2DelegateBridge.TransactOpts, _l2Token, _to, _amount, _minGasLimit, _extraData)
}
