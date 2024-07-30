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

// ERC20SafeMetaData contains all meta data concerning the ERC20Safe contract.
var ERC20SafeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"fundERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061054c806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c806395601f0914610030575b600080fd5b61004a60048036038101906100459190610298565b61004c565b005b600083905061005d81843085610063565b50505050565b6100e6846323b872dd60e01b85858560405160240161008493929190610309565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506100ec565b50505050565b6000808373ffffffffffffffffffffffffffffffffffffffff168360405161011491906103b1565b6000604051808303816000865af19150503d8060008114610151576040519150601f19603f3d011682016040523d82523d6000602084013e610156565b606091505b50915091508161019b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161019290610425565b60405180910390fd5b6000815111156101f957808060200190518101906101b9919061047d565b6101f8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101ef906104f6565b60405180910390fd5b5b50505050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061022f82610204565b9050919050565b61023f81610224565b811461024a57600080fd5b50565b60008135905061025c81610236565b92915050565b6000819050919050565b61027581610262565b811461028057600080fd5b50565b6000813590506102928161026c565b92915050565b6000806000606084860312156102b1576102b06101ff565b5b60006102bf8682870161024d565b93505060206102d08682870161024d565b92505060406102e186828701610283565b9150509250925092565b6102f481610224565b82525050565b61030381610262565b82525050565b600060608201905061031e60008301866102eb565b61032b60208301856102eb565b61033860408301846102fa565b949350505050565b600081519050919050565b600081905092915050565b60005b83811015610374578082015181840152602081019050610359565b60008484015250505050565b600061038b82610340565b610395818561034b565b93506103a5818560208601610356565b80840191505092915050565b60006103bd8284610380565b915081905092915050565b600082825260208201905092915050565b7f45524332303a2063616c6c206661696c65640000000000000000000000000000600082015250565b600061040f6012836103c8565b915061041a826103d9565b602082019050919050565b6000602082019050818103600083015261043e81610402565b9050919050565b60008115159050919050565b61045a81610445565b811461046557600080fd5b50565b60008151905061047781610451565b92915050565b600060208284031215610493576104926101ff565b5b60006104a184828501610468565b91505092915050565b7f45524332303a206f7065726174696f6e20646964206e6f742073756363656564600082015250565b60006104e06020836103c8565b91506104eb826104aa565b602082019050919050565b6000602082019050818103600083015261050f816104d3565b905091905056fea264697066735822122029d9f19e7a34164455138d68d9252a27e579bcd72fe642c99fb368b6df46703164736f6c63430008130033",
}

// ERC20SafeABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC20SafeMetaData.ABI instead.
var ERC20SafeABI = ERC20SafeMetaData.ABI

// ERC20SafeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ERC20SafeMetaData.Bin instead.
var ERC20SafeBin = ERC20SafeMetaData.Bin

// DeployERC20Safe deploys a new Ethereum contract, binding an instance of ERC20Safe to it.
func DeployERC20Safe(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC20Safe, error) {
	parsed, err := ERC20SafeMetaData.GetAbi()
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
	parsed, err := ERC20SafeMetaData.GetAbi()
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
