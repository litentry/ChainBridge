// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package NoArgument

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

// NoArgumentMetaData contains all meta data concerning the NoArgument contract.
var NoArgumentMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[],\"name\":\"NoArgumentCalled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"noArgument\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060008060006101000a81548160ff0219169083151502179055506103fd8061003a6000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80633f4ba83a14610051578063568959ca1461005b5780635c975abb146100655780638456cb5914610083575b600080fd5b61005961008d565b005b610063610097565b005b61006d6100cd565b60405161007a9190610267565b60405180910390f35b61008b6100e3565b005b6100956100ed565b565b61009f61014f565b7fc582abe1670c5a7f7cad8f171e4af03c793dd9f59fee6714179f56b6e9aea26f60405160405180910390a1565b60008060009054906101000a900460ff16905090565b6100eb610199565b565b6100f56101fb565b60008060006101000a81548160ff0219169083151502179055507f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa610138610244565b60405161014591906102c3565b60405180910390a1565b6101576100cd565b15610197576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161018e9061033b565b60405180910390fd5b565b6101a161014f565b60016000806101000a81548160ff0219169083151502179055507f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586101e4610244565b6040516101f191906102c3565b60405180910390a1565b6102036100cd565b610242576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610239906103a7565b60405180910390fd5b565b600033905090565b60008115159050919050565b6102618161024c565b82525050565b600060208201905061027c6000830184610258565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006102ad82610282565b9050919050565b6102bd816102a2565b82525050565b60006020820190506102d860008301846102b4565b92915050565b600082825260208201905092915050565b7f5061757361626c653a2070617573656400000000000000000000000000000000600082015250565b60006103256010836102de565b9150610330826102ef565b602082019050919050565b6000602082019050818103600083015261035481610318565b9050919050565b7f5061757361626c653a206e6f7420706175736564000000000000000000000000600082015250565b60006103916014836102de565b915061039c8261035b565b602082019050919050565b600060208201905081810360008301526103c081610384565b905091905056fea2646970667358221220f04876a8f421cfcc54a0a94bf9db550f2bd26849cf8495788339bbed158a226e64736f6c63430008130033",
}

// NoArgumentABI is the input ABI used to generate the binding from.
// Deprecated: Use NoArgumentMetaData.ABI instead.
var NoArgumentABI = NoArgumentMetaData.ABI

// NoArgumentBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use NoArgumentMetaData.Bin instead.
var NoArgumentBin = NoArgumentMetaData.Bin

// DeployNoArgument deploys a new Ethereum contract, binding an instance of NoArgument to it.
func DeployNoArgument(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *NoArgument, error) {
	parsed, err := NoArgumentMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(NoArgumentBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NoArgument{NoArgumentCaller: NoArgumentCaller{contract: contract}, NoArgumentTransactor: NoArgumentTransactor{contract: contract}, NoArgumentFilterer: NoArgumentFilterer{contract: contract}}, nil
}

// NoArgument is an auto generated Go binding around an Ethereum contract.
type NoArgument struct {
	NoArgumentCaller     // Read-only binding to the contract
	NoArgumentTransactor // Write-only binding to the contract
	NoArgumentFilterer   // Log filterer for contract events
}

// NoArgumentCaller is an auto generated read-only Go binding around an Ethereum contract.
type NoArgumentCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NoArgumentTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NoArgumentTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NoArgumentFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NoArgumentFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NoArgumentSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NoArgumentSession struct {
	Contract     *NoArgument       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NoArgumentCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NoArgumentCallerSession struct {
	Contract *NoArgumentCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// NoArgumentTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NoArgumentTransactorSession struct {
	Contract     *NoArgumentTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// NoArgumentRaw is an auto generated low-level Go binding around an Ethereum contract.
type NoArgumentRaw struct {
	Contract *NoArgument // Generic contract binding to access the raw methods on
}

// NoArgumentCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NoArgumentCallerRaw struct {
	Contract *NoArgumentCaller // Generic read-only contract binding to access the raw methods on
}

// NoArgumentTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NoArgumentTransactorRaw struct {
	Contract *NoArgumentTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNoArgument creates a new instance of NoArgument, bound to a specific deployed contract.
func NewNoArgument(address common.Address, backend bind.ContractBackend) (*NoArgument, error) {
	contract, err := bindNoArgument(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NoArgument{NoArgumentCaller: NoArgumentCaller{contract: contract}, NoArgumentTransactor: NoArgumentTransactor{contract: contract}, NoArgumentFilterer: NoArgumentFilterer{contract: contract}}, nil
}

// NewNoArgumentCaller creates a new read-only instance of NoArgument, bound to a specific deployed contract.
func NewNoArgumentCaller(address common.Address, caller bind.ContractCaller) (*NoArgumentCaller, error) {
	contract, err := bindNoArgument(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NoArgumentCaller{contract: contract}, nil
}

// NewNoArgumentTransactor creates a new write-only instance of NoArgument, bound to a specific deployed contract.
func NewNoArgumentTransactor(address common.Address, transactor bind.ContractTransactor) (*NoArgumentTransactor, error) {
	contract, err := bindNoArgument(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NoArgumentTransactor{contract: contract}, nil
}

// NewNoArgumentFilterer creates a new log filterer instance of NoArgument, bound to a specific deployed contract.
func NewNoArgumentFilterer(address common.Address, filterer bind.ContractFilterer) (*NoArgumentFilterer, error) {
	contract, err := bindNoArgument(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NoArgumentFilterer{contract: contract}, nil
}

// bindNoArgument binds a generic wrapper to an already deployed contract.
func bindNoArgument(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NoArgumentMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NoArgument *NoArgumentRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NoArgument.Contract.NoArgumentCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NoArgument *NoArgumentRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NoArgument.Contract.NoArgumentTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NoArgument *NoArgumentRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NoArgument.Contract.NoArgumentTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NoArgument *NoArgumentCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NoArgument.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NoArgument *NoArgumentTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NoArgument.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NoArgument *NoArgumentTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NoArgument.Contract.contract.Transact(opts, method, params...)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_NoArgument *NoArgumentCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _NoArgument.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_NoArgument *NoArgumentSession) Paused() (bool, error) {
	return _NoArgument.Contract.Paused(&_NoArgument.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_NoArgument *NoArgumentCallerSession) Paused() (bool, error) {
	return _NoArgument.Contract.Paused(&_NoArgument.CallOpts)
}

// NoArgument is a paid mutator transaction binding the contract method 0x568959ca.
//
// Solidity: function noArgument() returns()
func (_NoArgument *NoArgumentTransactor) NoArgument(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NoArgument.contract.Transact(opts, "noArgument")
}

// NoArgument is a paid mutator transaction binding the contract method 0x568959ca.
//
// Solidity: function noArgument() returns()
func (_NoArgument *NoArgumentSession) NoArgument() (*types.Transaction, error) {
	return _NoArgument.Contract.NoArgument(&_NoArgument.TransactOpts)
}

// NoArgument is a paid mutator transaction binding the contract method 0x568959ca.
//
// Solidity: function noArgument() returns()
func (_NoArgument *NoArgumentTransactorSession) NoArgument() (*types.Transaction, error) {
	return _NoArgument.Contract.NoArgument(&_NoArgument.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_NoArgument *NoArgumentTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NoArgument.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_NoArgument *NoArgumentSession) Pause() (*types.Transaction, error) {
	return _NoArgument.Contract.Pause(&_NoArgument.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_NoArgument *NoArgumentTransactorSession) Pause() (*types.Transaction, error) {
	return _NoArgument.Contract.Pause(&_NoArgument.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_NoArgument *NoArgumentTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NoArgument.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_NoArgument *NoArgumentSession) Unpause() (*types.Transaction, error) {
	return _NoArgument.Contract.Unpause(&_NoArgument.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_NoArgument *NoArgumentTransactorSession) Unpause() (*types.Transaction, error) {
	return _NoArgument.Contract.Unpause(&_NoArgument.TransactOpts)
}

// NoArgumentNoArgumentCalledIterator is returned from FilterNoArgumentCalled and is used to iterate over the raw logs and unpacked data for NoArgumentCalled events raised by the NoArgument contract.
type NoArgumentNoArgumentCalledIterator struct {
	Event *NoArgumentNoArgumentCalled // Event containing the contract specifics and raw log

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
func (it *NoArgumentNoArgumentCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NoArgumentNoArgumentCalled)
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
		it.Event = new(NoArgumentNoArgumentCalled)
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
func (it *NoArgumentNoArgumentCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NoArgumentNoArgumentCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NoArgumentNoArgumentCalled represents a NoArgumentCalled event raised by the NoArgument contract.
type NoArgumentNoArgumentCalled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterNoArgumentCalled is a free log retrieval operation binding the contract event 0xc582abe1670c5a7f7cad8f171e4af03c793dd9f59fee6714179f56b6e9aea26f.
//
// Solidity: event NoArgumentCalled()
func (_NoArgument *NoArgumentFilterer) FilterNoArgumentCalled(opts *bind.FilterOpts) (*NoArgumentNoArgumentCalledIterator, error) {

	logs, sub, err := _NoArgument.contract.FilterLogs(opts, "NoArgumentCalled")
	if err != nil {
		return nil, err
	}
	return &NoArgumentNoArgumentCalledIterator{contract: _NoArgument.contract, event: "NoArgumentCalled", logs: logs, sub: sub}, nil
}

// WatchNoArgumentCalled is a free log subscription operation binding the contract event 0xc582abe1670c5a7f7cad8f171e4af03c793dd9f59fee6714179f56b6e9aea26f.
//
// Solidity: event NoArgumentCalled()
func (_NoArgument *NoArgumentFilterer) WatchNoArgumentCalled(opts *bind.WatchOpts, sink chan<- *NoArgumentNoArgumentCalled) (event.Subscription, error) {

	logs, sub, err := _NoArgument.contract.WatchLogs(opts, "NoArgumentCalled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NoArgumentNoArgumentCalled)
				if err := _NoArgument.contract.UnpackLog(event, "NoArgumentCalled", log); err != nil {
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

// ParseNoArgumentCalled is a log parse operation binding the contract event 0xc582abe1670c5a7f7cad8f171e4af03c793dd9f59fee6714179f56b6e9aea26f.
//
// Solidity: event NoArgumentCalled()
func (_NoArgument *NoArgumentFilterer) ParseNoArgumentCalled(log types.Log) (*NoArgumentNoArgumentCalled, error) {
	event := new(NoArgumentNoArgumentCalled)
	if err := _NoArgument.contract.UnpackLog(event, "NoArgumentCalled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NoArgumentPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the NoArgument contract.
type NoArgumentPausedIterator struct {
	Event *NoArgumentPaused // Event containing the contract specifics and raw log

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
func (it *NoArgumentPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NoArgumentPaused)
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
		it.Event = new(NoArgumentPaused)
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
func (it *NoArgumentPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NoArgumentPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NoArgumentPaused represents a Paused event raised by the NoArgument contract.
type NoArgumentPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_NoArgument *NoArgumentFilterer) FilterPaused(opts *bind.FilterOpts) (*NoArgumentPausedIterator, error) {

	logs, sub, err := _NoArgument.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &NoArgumentPausedIterator{contract: _NoArgument.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_NoArgument *NoArgumentFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *NoArgumentPaused) (event.Subscription, error) {

	logs, sub, err := _NoArgument.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NoArgumentPaused)
				if err := _NoArgument.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_NoArgument *NoArgumentFilterer) ParsePaused(log types.Log) (*NoArgumentPaused, error) {
	event := new(NoArgumentPaused)
	if err := _NoArgument.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NoArgumentUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the NoArgument contract.
type NoArgumentUnpausedIterator struct {
	Event *NoArgumentUnpaused // Event containing the contract specifics and raw log

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
func (it *NoArgumentUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NoArgumentUnpaused)
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
		it.Event = new(NoArgumentUnpaused)
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
func (it *NoArgumentUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NoArgumentUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NoArgumentUnpaused represents a Unpaused event raised by the NoArgument contract.
type NoArgumentUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_NoArgument *NoArgumentFilterer) FilterUnpaused(opts *bind.FilterOpts) (*NoArgumentUnpausedIterator, error) {

	logs, sub, err := _NoArgument.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &NoArgumentUnpausedIterator{contract: _NoArgument.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_NoArgument *NoArgumentFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *NoArgumentUnpaused) (event.Subscription, error) {

	logs, sub, err := _NoArgument.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NoArgumentUnpaused)
				if err := _NoArgument.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_NoArgument *NoArgumentFilterer) ParseUnpaused(log types.Log) (*NoArgumentUnpaused, error) {
	event := new(NoArgumentUnpaused)
	if err := _NoArgument.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
