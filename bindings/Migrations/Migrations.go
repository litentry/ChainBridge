// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Migrations

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

// MigrationsMetaData contains all meta data concerning the Migrations contract.
var MigrationsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"last_completed_migration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"completed\",\"type\":\"uint256\"}],\"name\":\"setCompleted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"new_address\",\"type\":\"address\"}],\"name\":\"upgrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610394806100606000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80630900f01014610051578063445df0ac1461006d5780638da5cb5b1461008b578063fdacd576146100a9575b600080fd5b61006b6004803603810190610066919061027a565b6100c5565b005b61007561018f565b60405161008291906102c0565b60405180910390f35b610093610195565b6040516100a091906102ea565b60405180910390f35b6100c360048036038101906100be9190610331565b6101b9565b005b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff160361018c5760008190508073ffffffffffffffffffffffffffffffffffffffff1663fdacd5766001546040518263ffffffff1660e01b815260040161015891906102c0565b600060405180830381600087803b15801561017257600080fd5b505af1158015610186573d6000803e3d6000fd5b50505050505b50565b60015481565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff160361021457806001819055505b50565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006102478261021c565b9050919050565b6102578161023c565b811461026257600080fd5b50565b6000813590506102748161024e565b92915050565b6000602082840312156102905761028f610217565b5b600061029e84828501610265565b91505092915050565b6000819050919050565b6102ba816102a7565b82525050565b60006020820190506102d560008301846102b1565b92915050565b6102e48161023c565b82525050565b60006020820190506102ff60008301846102db565b92915050565b61030e816102a7565b811461031957600080fd5b50565b60008135905061032b81610305565b92915050565b60006020828403121561034757610346610217565b5b60006103558482850161031c565b9150509291505056fea2646970667358221220a889040bd43db0958b465f89a9d4b6a88f96ec8b9afa1fea680e3e26aedb3dc464736f6c63430008130033",
}

// MigrationsABI is the input ABI used to generate the binding from.
// Deprecated: Use MigrationsMetaData.ABI instead.
var MigrationsABI = MigrationsMetaData.ABI

// MigrationsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MigrationsMetaData.Bin instead.
var MigrationsBin = MigrationsMetaData.Bin

// DeployMigrations deploys a new Ethereum contract, binding an instance of Migrations to it.
func DeployMigrations(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Migrations, error) {
	parsed, err := MigrationsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MigrationsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Migrations{MigrationsCaller: MigrationsCaller{contract: contract}, MigrationsTransactor: MigrationsTransactor{contract: contract}, MigrationsFilterer: MigrationsFilterer{contract: contract}}, nil
}

// Migrations is an auto generated Go binding around an Ethereum contract.
type Migrations struct {
	MigrationsCaller     // Read-only binding to the contract
	MigrationsTransactor // Write-only binding to the contract
	MigrationsFilterer   // Log filterer for contract events
}

// MigrationsCaller is an auto generated read-only Go binding around an Ethereum contract.
type MigrationsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MigrationsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MigrationsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MigrationsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MigrationsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MigrationsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MigrationsSession struct {
	Contract     *Migrations       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MigrationsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MigrationsCallerSession struct {
	Contract *MigrationsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// MigrationsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MigrationsTransactorSession struct {
	Contract     *MigrationsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MigrationsRaw is an auto generated low-level Go binding around an Ethereum contract.
type MigrationsRaw struct {
	Contract *Migrations // Generic contract binding to access the raw methods on
}

// MigrationsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MigrationsCallerRaw struct {
	Contract *MigrationsCaller // Generic read-only contract binding to access the raw methods on
}

// MigrationsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MigrationsTransactorRaw struct {
	Contract *MigrationsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMigrations creates a new instance of Migrations, bound to a specific deployed contract.
func NewMigrations(address common.Address, backend bind.ContractBackend) (*Migrations, error) {
	contract, err := bindMigrations(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Migrations{MigrationsCaller: MigrationsCaller{contract: contract}, MigrationsTransactor: MigrationsTransactor{contract: contract}, MigrationsFilterer: MigrationsFilterer{contract: contract}}, nil
}

// NewMigrationsCaller creates a new read-only instance of Migrations, bound to a specific deployed contract.
func NewMigrationsCaller(address common.Address, caller bind.ContractCaller) (*MigrationsCaller, error) {
	contract, err := bindMigrations(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MigrationsCaller{contract: contract}, nil
}

// NewMigrationsTransactor creates a new write-only instance of Migrations, bound to a specific deployed contract.
func NewMigrationsTransactor(address common.Address, transactor bind.ContractTransactor) (*MigrationsTransactor, error) {
	contract, err := bindMigrations(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MigrationsTransactor{contract: contract}, nil
}

// NewMigrationsFilterer creates a new log filterer instance of Migrations, bound to a specific deployed contract.
func NewMigrationsFilterer(address common.Address, filterer bind.ContractFilterer) (*MigrationsFilterer, error) {
	contract, err := bindMigrations(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MigrationsFilterer{contract: contract}, nil
}

// bindMigrations binds a generic wrapper to an already deployed contract.
func bindMigrations(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MigrationsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Migrations *MigrationsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Migrations.Contract.MigrationsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Migrations *MigrationsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Migrations.Contract.MigrationsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Migrations *MigrationsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Migrations.Contract.MigrationsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Migrations *MigrationsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Migrations.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Migrations *MigrationsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Migrations.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Migrations *MigrationsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Migrations.Contract.contract.Transact(opts, method, params...)
}

// LastCompletedMigration is a free data retrieval call binding the contract method 0x445df0ac.
//
// Solidity: function last_completed_migration() view returns(uint256)
func (_Migrations *MigrationsCaller) LastCompletedMigration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Migrations.contract.Call(opts, &out, "last_completed_migration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastCompletedMigration is a free data retrieval call binding the contract method 0x445df0ac.
//
// Solidity: function last_completed_migration() view returns(uint256)
func (_Migrations *MigrationsSession) LastCompletedMigration() (*big.Int, error) {
	return _Migrations.Contract.LastCompletedMigration(&_Migrations.CallOpts)
}

// LastCompletedMigration is a free data retrieval call binding the contract method 0x445df0ac.
//
// Solidity: function last_completed_migration() view returns(uint256)
func (_Migrations *MigrationsCallerSession) LastCompletedMigration() (*big.Int, error) {
	return _Migrations.Contract.LastCompletedMigration(&_Migrations.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Migrations *MigrationsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Migrations.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Migrations *MigrationsSession) Owner() (common.Address, error) {
	return _Migrations.Contract.Owner(&_Migrations.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Migrations *MigrationsCallerSession) Owner() (common.Address, error) {
	return _Migrations.Contract.Owner(&_Migrations.CallOpts)
}

// SetCompleted is a paid mutator transaction binding the contract method 0xfdacd576.
//
// Solidity: function setCompleted(uint256 completed) returns()
func (_Migrations *MigrationsTransactor) SetCompleted(opts *bind.TransactOpts, completed *big.Int) (*types.Transaction, error) {
	return _Migrations.contract.Transact(opts, "setCompleted", completed)
}

// SetCompleted is a paid mutator transaction binding the contract method 0xfdacd576.
//
// Solidity: function setCompleted(uint256 completed) returns()
func (_Migrations *MigrationsSession) SetCompleted(completed *big.Int) (*types.Transaction, error) {
	return _Migrations.Contract.SetCompleted(&_Migrations.TransactOpts, completed)
}

// SetCompleted is a paid mutator transaction binding the contract method 0xfdacd576.
//
// Solidity: function setCompleted(uint256 completed) returns()
func (_Migrations *MigrationsTransactorSession) SetCompleted(completed *big.Int) (*types.Transaction, error) {
	return _Migrations.Contract.SetCompleted(&_Migrations.TransactOpts, completed)
}

// Upgrade is a paid mutator transaction binding the contract method 0x0900f010.
//
// Solidity: function upgrade(address new_address) returns()
func (_Migrations *MigrationsTransactor) Upgrade(opts *bind.TransactOpts, new_address common.Address) (*types.Transaction, error) {
	return _Migrations.contract.Transact(opts, "upgrade", new_address)
}

// Upgrade is a paid mutator transaction binding the contract method 0x0900f010.
//
// Solidity: function upgrade(address new_address) returns()
func (_Migrations *MigrationsSession) Upgrade(new_address common.Address) (*types.Transaction, error) {
	return _Migrations.Contract.Upgrade(&_Migrations.TransactOpts, new_address)
}

// Upgrade is a paid mutator transaction binding the contract method 0x0900f010.
//
// Solidity: function upgrade(address new_address) returns()
func (_Migrations *MigrationsTransactorSession) Upgrade(new_address common.Address) (*types.Transaction, error) {
	return _Migrations.Contract.Upgrade(&_Migrations.TransactOpts, new_address)
}
