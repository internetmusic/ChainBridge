// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ERC20Safe

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ERC20SafeABI is the input ABI used to generate the binding from.
const ERC20SafeABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"fundERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ERC20SafeBin is the compiled bytecode used for deploying new contracts.
var ERC20SafeBin = "0x608060405234801561001057600080fd5b5061036c806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80636ebcf6071461003b57806395601f0914610093575b600080fd5b61007d6004803603602081101561005157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610101565b6040518082815260200191505060405180910390f35b6100ff600480360360608110156100a957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610119565b005b60006020528060005260406000206000915090505481565b60008390508073ffffffffffffffffffffffffffffffffffffffff166323b872dd8430856040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050602060405180830381600087803b1580156101d957600080fd5b505af11580156101ed573d6000803e3d6000fd5b505050506040513d602081101561020357600080fd5b810190808051906020019092919050505050610266826000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546102ae90919063ffffffff16565b6000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555050505050565b60008082840190508381101561032c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f536166654d6174683a206164646974696f6e206f766572666c6f77000000000081525060200191505060405180910390fd5b809150509291505056fea264697066735822122017abc7ed841ab1a3a46bde3bec1e5cf1c2e9d27f59d9d86e35fecb6cc7bc33f064736f6c63430006040033"

// DeployERC20Safe deploys a new Ethereum contract, binding an instance of ERC20Safe to it.
func DeployERC20Safe(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC20Safe, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20SafeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20SafeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20Safe{ERC20SafeCaller: ERC20SafeCaller{contract: contract}, ERC20SafeTransactor: ERC20SafeTransactor{contract: contract}, ERC20SafeFilterer: ERC20SafeFilterer{contract: contract}}, nil
}

// ERC20Safe is an auto generated Go binding around an Ethereum contract.
type ERC20Safe struct {
	ERC20SafeCaller     // Read-only binding to the contract
	ERC20SafeTransactor // Write-only binding to the contract
	ERC20SafeFilterer   // Log filterer for contract events
}

// ERC20SafeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20SafeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20SafeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20SafeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20SafeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20SafeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20SafeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20SafeSession struct {
	Contract     *ERC20Safe        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20SafeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20SafeCallerSession struct {
	Contract *ERC20SafeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ERC20SafeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20SafeTransactorSession struct {
	Contract     *ERC20SafeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ERC20SafeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20SafeRaw struct {
	Contract *ERC20Safe // Generic contract binding to access the raw methods on
}

// ERC20SafeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20SafeCallerRaw struct {
	Contract *ERC20SafeCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20SafeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20SafeTransactorRaw struct {
	Contract *ERC20SafeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Safe creates a new instance of ERC20Safe, bound to a specific deployed contract.
func NewERC20Safe(address common.Address, backend bind.ContractBackend) (*ERC20Safe, error) {
	contract, err := bindERC20Safe(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Safe{ERC20SafeCaller: ERC20SafeCaller{contract: contract}, ERC20SafeTransactor: ERC20SafeTransactor{contract: contract}, ERC20SafeFilterer: ERC20SafeFilterer{contract: contract}}, nil
}

// NewERC20SafeCaller creates a new read-only instance of ERC20Safe, bound to a specific deployed contract.
func NewERC20SafeCaller(address common.Address, caller bind.ContractCaller) (*ERC20SafeCaller, error) {
	contract, err := bindERC20Safe(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20SafeCaller{contract: contract}, nil
}

// NewERC20SafeTransactor creates a new write-only instance of ERC20Safe, bound to a specific deployed contract.
func NewERC20SafeTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20SafeTransactor, error) {
	contract, err := bindERC20Safe(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20SafeTransactor{contract: contract}, nil
}

// NewERC20SafeFilterer creates a new log filterer instance of ERC20Safe, bound to a specific deployed contract.
func NewERC20SafeFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20SafeFilterer, error) {
	contract, err := bindERC20Safe(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20SafeFilterer{contract: contract}, nil
}

// bindERC20Safe binds a generic wrapper to an already deployed contract.
func bindERC20Safe(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20SafeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Safe *ERC20SafeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Safe.Contract.ERC20SafeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Safe *ERC20SafeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Safe.Contract.ERC20SafeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Safe *ERC20SafeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Safe.Contract.ERC20SafeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Safe *ERC20SafeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Safe.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Safe *ERC20SafeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Safe.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Safe *ERC20SafeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Safe.Contract.contract.Transact(opts, method, params...)
}

// Balances is a free data retrieval call binding the contract method 0x6ebcf607.
//
// Solidity: function _balances(address ) constant returns(uint256)
func (_ERC20Safe *ERC20SafeCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Safe.contract.Call(opts, out, "_balances", arg0)
	return *ret0, err
}

// Balances is a free data retrieval call binding the contract method 0x6ebcf607.
//
// Solidity: function _balances(address ) constant returns(uint256)
func (_ERC20Safe *ERC20SafeSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _ERC20Safe.Contract.Balances(&_ERC20Safe.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x6ebcf607.
//
// Solidity: function _balances(address ) constant returns(uint256)
func (_ERC20Safe *ERC20SafeCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _ERC20Safe.Contract.Balances(&_ERC20Safe.CallOpts, arg0)
}

// FundERC20 is a paid mutator transaction binding the contract method 0x95601f09.
//
// Solidity: function fundERC20(address tokenAddress, address owner, uint256 amount) returns()
func (_ERC20Safe *ERC20SafeTransactor) FundERC20(opts *bind.TransactOpts, tokenAddress common.Address, owner common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Safe.contract.Transact(opts, "fundERC20", tokenAddress, owner, amount)
}

// FundERC20 is a paid mutator transaction binding the contract method 0x95601f09.
//
// Solidity: function fundERC20(address tokenAddress, address owner, uint256 amount) returns()
func (_ERC20Safe *ERC20SafeSession) FundERC20(tokenAddress common.Address, owner common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Safe.Contract.FundERC20(&_ERC20Safe.TransactOpts, tokenAddress, owner, amount)
}

// FundERC20 is a paid mutator transaction binding the contract method 0x95601f09.
//
// Solidity: function fundERC20(address tokenAddress, address owner, uint256 amount) returns()
func (_ERC20Safe *ERC20SafeTransactorSession) FundERC20(tokenAddress common.Address, owner common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Safe.Contract.FundERC20(&_ERC20Safe.TransactOpts, tokenAddress, owner, amount)
}

var RuntimeBytecode = "0x608060405234801561001057600080fd5b50600436106100365760003560e01c80636ebcf6071461003b57806395601f0914610093575b600080fd5b61007d6004803603602081101561005157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610101565b6040518082815260200191505060405180910390f35b6100ff600480360360608110156100a957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610119565b005b60006020528060005260406000206000915090505481565b60008390508073ffffffffffffffffffffffffffffffffffffffff166323b872dd8430856040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050602060405180830381600087803b1580156101d957600080fd5b505af11580156101ed573d6000803e3d6000fd5b505050506040513d602081101561020357600080fd5b810190808051906020019092919050505050610266826000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546102ae90919063ffffffff16565b6000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555050505050565b60008082840190508381101561032c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f536166654d6174683a206164646974696f6e206f766572666c6f77000000000081525060200191505060405180910390fd5b809150509291505056fea264697066735822122017abc7ed841ab1a3a46bde3bec1e5cf1c2e9d27f59d9d86e35fecb6cc7bc33f064736f6c63430006040033"