// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package OneArgument

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

// OneArgumentMetaData contains all meta data concerning the OneArgument contract.
var OneArgumentMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"argumentOne\",\"type\":\"uint256\"}],\"name\":\"OneArgumentCalled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"argumentOne\",\"type\":\"uint256\"}],\"name\":\"oneArgument\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060008060006101000a81548160ff0219169083151502179055506104798061003a6000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80633f4ba83a146100515780635c975abb1461005b5780638456cb5914610079578063c95cf0d814610083575b600080fd5b61005961009f565b005b6100636100a9565b604051610070919061027b565b60405180910390f35b6100816100bf565b005b61009d600480360381019061009891906102d1565b6100c9565b005b6100a7610101565b565b60008060009054906101000a900460ff16905090565b6100c7610163565b565b6100d16101c5565b807f29ab08c845830c69b55a1fba5c95718f65dc24361a471e3da14cd5ff2b37315960405160405180910390a250565b61010961020f565b60008060006101000a81548160ff0219169083151502179055507f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa61014c610258565b604051610159919061033f565b60405180910390a1565b61016b6101c5565b60016000806101000a81548160ff0219169083151502179055507f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586101ae610258565b6040516101bb919061033f565b60405180910390a1565b6101cd6100a9565b1561020d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610204906103b7565b60405180910390fd5b565b6102176100a9565b610256576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161024d90610423565b60405180910390fd5b565b600033905090565b60008115159050919050565b61027581610260565b82525050565b6000602082019050610290600083018461026c565b92915050565b600080fd5b6000819050919050565b6102ae8161029b565b81146102b957600080fd5b50565b6000813590506102cb816102a5565b92915050565b6000602082840312156102e7576102e6610296565b5b60006102f5848285016102bc565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610329826102fe565b9050919050565b6103398161031e565b82525050565b60006020820190506103546000830184610330565b92915050565b600082825260208201905092915050565b7f5061757361626c653a2070617573656400000000000000000000000000000000600082015250565b60006103a160108361035a565b91506103ac8261036b565b602082019050919050565b600060208201905081810360008301526103d081610394565b9050919050565b7f5061757361626c653a206e6f7420706175736564000000000000000000000000600082015250565b600061040d60148361035a565b9150610418826103d7565b602082019050919050565b6000602082019050818103600083015261043c81610400565b905091905056fea26469706673582212208b7862d0679540580a7738388c07c9b5511993eccc7718f20d974261692b1d5564736f6c63430008130033",
}

// OneArgumentABI is the input ABI used to generate the binding from.
// Deprecated: Use OneArgumentMetaData.ABI instead.
var OneArgumentABI = OneArgumentMetaData.ABI

// OneArgumentBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OneArgumentMetaData.Bin instead.
var OneArgumentBin = OneArgumentMetaData.Bin

// DeployOneArgument deploys a new Ethereum contract, binding an instance of OneArgument to it.
func DeployOneArgument(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OneArgument, error) {
	parsed, err := OneArgumentMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OneArgumentBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OneArgument{OneArgumentCaller: OneArgumentCaller{contract: contract}, OneArgumentTransactor: OneArgumentTransactor{contract: contract}, OneArgumentFilterer: OneArgumentFilterer{contract: contract}}, nil
}

// OneArgument is an auto generated Go binding around an Ethereum contract.
type OneArgument struct {
	OneArgumentCaller     // Read-only binding to the contract
	OneArgumentTransactor // Write-only binding to the contract
	OneArgumentFilterer   // Log filterer for contract events
}

// OneArgumentCaller is an auto generated read-only Go binding around an Ethereum contract.
type OneArgumentCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneArgumentTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OneArgumentTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneArgumentFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OneArgumentFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneArgumentSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OneArgumentSession struct {
	Contract     *OneArgument      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OneArgumentCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OneArgumentCallerSession struct {
	Contract *OneArgumentCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// OneArgumentTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OneArgumentTransactorSession struct {
	Contract     *OneArgumentTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// OneArgumentRaw is an auto generated low-level Go binding around an Ethereum contract.
type OneArgumentRaw struct {
	Contract *OneArgument // Generic contract binding to access the raw methods on
}

// OneArgumentCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OneArgumentCallerRaw struct {
	Contract *OneArgumentCaller // Generic read-only contract binding to access the raw methods on
}

// OneArgumentTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OneArgumentTransactorRaw struct {
	Contract *OneArgumentTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOneArgument creates a new instance of OneArgument, bound to a specific deployed contract.
func NewOneArgument(address common.Address, backend bind.ContractBackend) (*OneArgument, error) {
	contract, err := bindOneArgument(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OneArgument{OneArgumentCaller: OneArgumentCaller{contract: contract}, OneArgumentTransactor: OneArgumentTransactor{contract: contract}, OneArgumentFilterer: OneArgumentFilterer{contract: contract}}, nil
}

// NewOneArgumentCaller creates a new read-only instance of OneArgument, bound to a specific deployed contract.
func NewOneArgumentCaller(address common.Address, caller bind.ContractCaller) (*OneArgumentCaller, error) {
	contract, err := bindOneArgument(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OneArgumentCaller{contract: contract}, nil
}

// NewOneArgumentTransactor creates a new write-only instance of OneArgument, bound to a specific deployed contract.
func NewOneArgumentTransactor(address common.Address, transactor bind.ContractTransactor) (*OneArgumentTransactor, error) {
	contract, err := bindOneArgument(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OneArgumentTransactor{contract: contract}, nil
}

// NewOneArgumentFilterer creates a new log filterer instance of OneArgument, bound to a specific deployed contract.
func NewOneArgumentFilterer(address common.Address, filterer bind.ContractFilterer) (*OneArgumentFilterer, error) {
	contract, err := bindOneArgument(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OneArgumentFilterer{contract: contract}, nil
}

// bindOneArgument binds a generic wrapper to an already deployed contract.
func bindOneArgument(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OneArgumentMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneArgument *OneArgumentRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OneArgument.Contract.OneArgumentCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneArgument *OneArgumentRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneArgument.Contract.OneArgumentTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneArgument *OneArgumentRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneArgument.Contract.OneArgumentTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneArgument *OneArgumentCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OneArgument.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneArgument *OneArgumentTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneArgument.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneArgument *OneArgumentTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneArgument.Contract.contract.Transact(opts, method, params...)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_OneArgument *OneArgumentCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _OneArgument.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_OneArgument *OneArgumentSession) Paused() (bool, error) {
	return _OneArgument.Contract.Paused(&_OneArgument.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_OneArgument *OneArgumentCallerSession) Paused() (bool, error) {
	return _OneArgument.Contract.Paused(&_OneArgument.CallOpts)
}

// OneArgument is a paid mutator transaction binding the contract method 0xc95cf0d8.
//
// Solidity: function oneArgument(uint256 argumentOne) returns()
func (_OneArgument *OneArgumentTransactor) OneArgument(opts *bind.TransactOpts, argumentOne *big.Int) (*types.Transaction, error) {
	return _OneArgument.contract.Transact(opts, "oneArgument", argumentOne)
}

// OneArgument is a paid mutator transaction binding the contract method 0xc95cf0d8.
//
// Solidity: function oneArgument(uint256 argumentOne) returns()
func (_OneArgument *OneArgumentSession) OneArgument(argumentOne *big.Int) (*types.Transaction, error) {
	return _OneArgument.Contract.OneArgument(&_OneArgument.TransactOpts, argumentOne)
}

// OneArgument is a paid mutator transaction binding the contract method 0xc95cf0d8.
//
// Solidity: function oneArgument(uint256 argumentOne) returns()
func (_OneArgument *OneArgumentTransactorSession) OneArgument(argumentOne *big.Int) (*types.Transaction, error) {
	return _OneArgument.Contract.OneArgument(&_OneArgument.TransactOpts, argumentOne)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_OneArgument *OneArgumentTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneArgument.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_OneArgument *OneArgumentSession) Pause() (*types.Transaction, error) {
	return _OneArgument.Contract.Pause(&_OneArgument.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_OneArgument *OneArgumentTransactorSession) Pause() (*types.Transaction, error) {
	return _OneArgument.Contract.Pause(&_OneArgument.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_OneArgument *OneArgumentTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneArgument.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_OneArgument *OneArgumentSession) Unpause() (*types.Transaction, error) {
	return _OneArgument.Contract.Unpause(&_OneArgument.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_OneArgument *OneArgumentTransactorSession) Unpause() (*types.Transaction, error) {
	return _OneArgument.Contract.Unpause(&_OneArgument.TransactOpts)
}

// OneArgumentOneArgumentCalledIterator is returned from FilterOneArgumentCalled and is used to iterate over the raw logs and unpacked data for OneArgumentCalled events raised by the OneArgument contract.
type OneArgumentOneArgumentCalledIterator struct {
	Event *OneArgumentOneArgumentCalled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OneArgumentOneArgumentCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OneArgumentOneArgumentCalled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OneArgumentOneArgumentCalled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OneArgumentOneArgumentCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OneArgumentOneArgumentCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OneArgumentOneArgumentCalled represents a OneArgumentCalled event raised by the OneArgument contract.
type OneArgumentOneArgumentCalled struct {
	ArgumentOne *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOneArgumentCalled is a free log retrieval operation binding the contract event 0x29ab08c845830c69b55a1fba5c95718f65dc24361a471e3da14cd5ff2b373159.
//
// Solidity: event OneArgumentCalled(uint256 indexed argumentOne)
func (_OneArgument *OneArgumentFilterer) FilterOneArgumentCalled(opts *bind.FilterOpts, argumentOne []*big.Int) (*OneArgumentOneArgumentCalledIterator, error) {

	var argumentOneRule []interface{}
	for _, argumentOneItem := range argumentOne {
		argumentOneRule = append(argumentOneRule, argumentOneItem)
	}

	logs, sub, err := _OneArgument.contract.FilterLogs(opts, "OneArgumentCalled", argumentOneRule)
	if err != nil {
		return nil, err
	}
	return &OneArgumentOneArgumentCalledIterator{contract: _OneArgument.contract, event: "OneArgumentCalled", logs: logs, sub: sub}, nil
}

// WatchOneArgumentCalled is a free log subscription operation binding the contract event 0x29ab08c845830c69b55a1fba5c95718f65dc24361a471e3da14cd5ff2b373159.
//
// Solidity: event OneArgumentCalled(uint256 indexed argumentOne)
func (_OneArgument *OneArgumentFilterer) WatchOneArgumentCalled(opts *bind.WatchOpts, sink chan<- *OneArgumentOneArgumentCalled, argumentOne []*big.Int) (event.Subscription, error) {

	var argumentOneRule []interface{}
	for _, argumentOneItem := range argumentOne {
		argumentOneRule = append(argumentOneRule, argumentOneItem)
	}

	logs, sub, err := _OneArgument.contract.WatchLogs(opts, "OneArgumentCalled", argumentOneRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OneArgumentOneArgumentCalled)
				if err := _OneArgument.contract.UnpackLog(event, "OneArgumentCalled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOneArgumentCalled is a log parse operation binding the contract event 0x29ab08c845830c69b55a1fba5c95718f65dc24361a471e3da14cd5ff2b373159.
//
// Solidity: event OneArgumentCalled(uint256 indexed argumentOne)
func (_OneArgument *OneArgumentFilterer) ParseOneArgumentCalled(log types.Log) (*OneArgumentOneArgumentCalled, error) {
	event := new(OneArgumentOneArgumentCalled)
	if err := _OneArgument.contract.UnpackLog(event, "OneArgumentCalled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OneArgumentPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the OneArgument contract.
type OneArgumentPausedIterator struct {
	Event *OneArgumentPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OneArgumentPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OneArgumentPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OneArgumentPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OneArgumentPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OneArgumentPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OneArgumentPaused represents a Paused event raised by the OneArgument contract.
type OneArgumentPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_OneArgument *OneArgumentFilterer) FilterPaused(opts *bind.FilterOpts) (*OneArgumentPausedIterator, error) {

	logs, sub, err := _OneArgument.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &OneArgumentPausedIterator{contract: _OneArgument.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_OneArgument *OneArgumentFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *OneArgumentPaused) (event.Subscription, error) {

	logs, sub, err := _OneArgument.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OneArgumentPaused)
				if err := _OneArgument.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_OneArgument *OneArgumentFilterer) ParsePaused(log types.Log) (*OneArgumentPaused, error) {
	event := new(OneArgumentPaused)
	if err := _OneArgument.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OneArgumentUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the OneArgument contract.
type OneArgumentUnpausedIterator struct {
	Event *OneArgumentUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OneArgumentUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OneArgumentUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OneArgumentUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OneArgumentUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OneArgumentUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OneArgumentUnpaused represents a Unpaused event raised by the OneArgument contract.
type OneArgumentUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_OneArgument *OneArgumentFilterer) FilterUnpaused(opts *bind.FilterOpts) (*OneArgumentUnpausedIterator, error) {

	logs, sub, err := _OneArgument.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &OneArgumentUnpausedIterator{contract: _OneArgument.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_OneArgument *OneArgumentFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *OneArgumentUnpaused) (event.Subscription, error) {

	logs, sub, err := _OneArgument.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OneArgumentUnpaused)
				if err := _OneArgument.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_OneArgument *OneArgumentFilterer) ParseUnpaused(log types.Log) (*OneArgumentUnpaused, error) {
	event := new(OneArgumentUnpaused)
	if err := _OneArgument.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
