// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ERC20Safe

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

// // ERC20SafeMetaData contains all meta data concerning the ERC20Safe contract.
// var ERC20SafeMetaData = &bind.MetaData{
// 	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"fundERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
// 	Bin: "0x608060405234801561001057600080fd5b50610383806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c806395601f0914610030575b600080fd5b61009c6004803603606081101561004657600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061009e565b005b60008390506100af818430856100b5565b50505050565b610170846323b872dd60e01b858585604051602401808473ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050610176565b50505050565b600060608373ffffffffffffffffffffffffffffffffffffffff16836040518082805190602001908083835b602083106101c557805182526020820191506020810190506020830392506101a2565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114610227576040519150601f19603f3d011682016040523d82523d6000602084013e61022c565b606091505b5091509150816102a4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f45524332303a2063616c6c206661696c6564000000000000000000000000000081525060200191505060405180910390fd5b600081511115610347578080602001905160208110156102c357600080fd5b8101908080519060200190929190505050610346576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f45524332303a206f7065726174696f6e20646964206e6f74207375636365656481525060200191505060405180910390fd5b5b5050505056fea26469706673582212207e4550a9040043b13c9082b998fdc223be0039b02641b76ecf66f04a0fd969de64736f6c63430007000033",
// }

// ERC20SafeABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC20SafeMetaData.ABI instead.
var ERC20SafeABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"fundERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ERC20SafeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ERC20SafeMetaData.Bin instead.
var ERC20SafeBin = "0x608060405234801561001057600080fd5b50610383806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c806395601f0914610030575b600080fd5b61009c6004803603606081101561004657600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061009e565b005b60008390506100af818430856100b5565b50505050565b610170846323b872dd60e01b858585604051602401808473ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050610176565b50505050565b600060608373ffffffffffffffffffffffffffffffffffffffff16836040518082805190602001908083835b602083106101c557805182526020820191506020810190506020830392506101a2565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114610227576040519150601f19603f3d011682016040523d82523d6000602084013e61022c565b606091505b5091509150816102a4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f45524332303a2063616c6c206661696c6564000000000000000000000000000081525060200191505060405180910390fd5b600081511115610347578080602001905160208110156102c357600080fd5b8101908080519060200190929190505050610346576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f45524332303a206f7065726174696f6e20646964206e6f74207375636365656481525060200191505060405180910390fd5b5b5050505056fea26469706673582212207e4550a9040043b13c9082b998fdc223be0039b02641b76ecf66f04a0fd969de64736f6c63430007000033"

// DeployERC20Safe deploys a new Ethereum contract, binding an instance of ERC20Safe to it.
func DeployERC20Safe(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC20Safe, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20SafeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ERC20SafeBin), backend)
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
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Safe *ERC20SafeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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
func (_ERC20Safe *ERC20SafeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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
