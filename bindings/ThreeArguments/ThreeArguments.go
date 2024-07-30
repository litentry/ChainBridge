// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ThreeArguments

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

// ThreeArgumentsMetaData contains all meta data concerning the ThreeArguments contract.
var ThreeArgumentsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"argumentOne\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"int8\",\"name\":\"argumentTwo\",\"type\":\"int8\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"argumentThree\",\"type\":\"bool\"}],\"name\":\"ThreeArgumentsCalled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"argumentOne\",\"type\":\"string\"},{\"internalType\":\"int8\",\"name\":\"argumentTwo\",\"type\":\"int8\"},{\"internalType\":\"bool\",\"name\":\"argumentThree\",\"type\":\"bool\"}],\"name\":\"threeArguments\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060008060006101000a81548160ff0219169083151502179055506106088061003a6000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80633f4ba83a146100515780635c975abb1461005b5780638456cb59146100795780639280b90514610083575b600080fd5b61005961009f565b005b6100636100a9565b604051610070919061028e565b60405180910390f35b6100816100bf565b005b61009d6004803603810190610098919061037d565b6100c9565b005b6100a7610114565b565b60008060009054906101000a900460ff16905090565b6100c7610176565b565b6100d16101d8565b7fd589183661fa75f94e2db32f4eb7ebb50f4154c160e15eb43f772a46f360a3a884848484604051610106949392919061045e565b60405180910390a150505050565b61011c610222565b60008060006101000a81548160ff0219169083151502179055507f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa61015f61026b565b60405161016c91906104df565b60405180910390a1565b61017e6101d8565b60016000806101000a81548160ff0219169083151502179055507f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586101c161026b565b6040516101ce91906104df565b60405180910390a1565b6101e06100a9565b15610220576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161021790610546565b60405180910390fd5b565b61022a6100a9565b610269576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610260906105b2565b60405180910390fd5b565b600033905090565b60008115159050919050565b61028881610273565b82525050565b60006020820190506102a3600083018461027f565b92915050565b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b60008083601f8401126102d8576102d76102b3565b5b8235905067ffffffffffffffff8111156102f5576102f46102b8565b5b602083019150836001820283011115610311576103106102bd565b5b9250929050565b60008160000b9050919050565b61032e81610318565b811461033957600080fd5b50565b60008135905061034b81610325565b92915050565b61035a81610273565b811461036557600080fd5b50565b60008135905061037781610351565b92915050565b60008060008060608587031215610397576103966102a9565b5b600085013567ffffffffffffffff8111156103b5576103b46102ae565b5b6103c1878288016102c2565b945094505060206103d48782880161033c565b92505060406103e587828801610368565b91505092959194509250565b600082825260208201905092915050565b82818337600083830152505050565b6000601f19601f8301169050919050565b600061042e83856103f1565b935061043b838584610402565b61044483610411565b840190509392505050565b61045881610318565b82525050565b60006060820190508181036000830152610479818688610422565b9050610488602083018561044f565b610495604083018461027f565b95945050505050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006104c98261049e565b9050919050565b6104d9816104be565b82525050565b60006020820190506104f460008301846104d0565b92915050565b7f5061757361626c653a2070617573656400000000000000000000000000000000600082015250565b60006105306010836103f1565b915061053b826104fa565b602082019050919050565b6000602082019050818103600083015261055f81610523565b9050919050565b7f5061757361626c653a206e6f7420706175736564000000000000000000000000600082015250565b600061059c6014836103f1565b91506105a782610566565b602082019050919050565b600060208201905081810360008301526105cb8161058f565b905091905056fea26469706673582212204caaab0a388b49fda864e816f3fb774f57c8dc90a2f455b81837bc303fabd5ca64736f6c63430008130033",
}

// ThreeArgumentsABI is the input ABI used to generate the binding from.
// Deprecated: Use ThreeArgumentsMetaData.ABI instead.
var ThreeArgumentsABI = ThreeArgumentsMetaData.ABI

// ThreeArgumentsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ThreeArgumentsMetaData.Bin instead.
var ThreeArgumentsBin = ThreeArgumentsMetaData.Bin

// DeployThreeArguments deploys a new Ethereum contract, binding an instance of ThreeArguments to it.
func DeployThreeArguments(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ThreeArguments, error) {
	parsed, err := ThreeArgumentsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ThreeArgumentsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ThreeArguments{ThreeArgumentsCaller: ThreeArgumentsCaller{contract: contract}, ThreeArgumentsTransactor: ThreeArgumentsTransactor{contract: contract}, ThreeArgumentsFilterer: ThreeArgumentsFilterer{contract: contract}}, nil
}

// ThreeArguments is an auto generated Go binding around an Ethereum contract.
type ThreeArguments struct {
	ThreeArgumentsCaller     // Read-only binding to the contract
	ThreeArgumentsTransactor // Write-only binding to the contract
	ThreeArgumentsFilterer   // Log filterer for contract events
}

// ThreeArgumentsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ThreeArgumentsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ThreeArgumentsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ThreeArgumentsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ThreeArgumentsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ThreeArgumentsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ThreeArgumentsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ThreeArgumentsSession struct {
	Contract     *ThreeArguments   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ThreeArgumentsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ThreeArgumentsCallerSession struct {
	Contract *ThreeArgumentsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ThreeArgumentsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ThreeArgumentsTransactorSession struct {
	Contract     *ThreeArgumentsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ThreeArgumentsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ThreeArgumentsRaw struct {
	Contract *ThreeArguments // Generic contract binding to access the raw methods on
}

// ThreeArgumentsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ThreeArgumentsCallerRaw struct {
	Contract *ThreeArgumentsCaller // Generic read-only contract binding to access the raw methods on
}

// ThreeArgumentsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ThreeArgumentsTransactorRaw struct {
	Contract *ThreeArgumentsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewThreeArguments creates a new instance of ThreeArguments, bound to a specific deployed contract.
func NewThreeArguments(address common.Address, backend bind.ContractBackend) (*ThreeArguments, error) {
	contract, err := bindThreeArguments(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ThreeArguments{ThreeArgumentsCaller: ThreeArgumentsCaller{contract: contract}, ThreeArgumentsTransactor: ThreeArgumentsTransactor{contract: contract}, ThreeArgumentsFilterer: ThreeArgumentsFilterer{contract: contract}}, nil
}

// NewThreeArgumentsCaller creates a new read-only instance of ThreeArguments, bound to a specific deployed contract.
func NewThreeArgumentsCaller(address common.Address, caller bind.ContractCaller) (*ThreeArgumentsCaller, error) {
	contract, err := bindThreeArguments(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ThreeArgumentsCaller{contract: contract}, nil
}

// NewThreeArgumentsTransactor creates a new write-only instance of ThreeArguments, bound to a specific deployed contract.
func NewThreeArgumentsTransactor(address common.Address, transactor bind.ContractTransactor) (*ThreeArgumentsTransactor, error) {
	contract, err := bindThreeArguments(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ThreeArgumentsTransactor{contract: contract}, nil
}

// NewThreeArgumentsFilterer creates a new log filterer instance of ThreeArguments, bound to a specific deployed contract.
func NewThreeArgumentsFilterer(address common.Address, filterer bind.ContractFilterer) (*ThreeArgumentsFilterer, error) {
	contract, err := bindThreeArguments(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ThreeArgumentsFilterer{contract: contract}, nil
}

// bindThreeArguments binds a generic wrapper to an already deployed contract.
func bindThreeArguments(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ThreeArgumentsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ThreeArguments *ThreeArgumentsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ThreeArguments.Contract.ThreeArgumentsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ThreeArguments *ThreeArgumentsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ThreeArguments.Contract.ThreeArgumentsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ThreeArguments *ThreeArgumentsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ThreeArguments.Contract.ThreeArgumentsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ThreeArguments *ThreeArgumentsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ThreeArguments.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ThreeArguments *ThreeArgumentsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ThreeArguments.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ThreeArguments *ThreeArgumentsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ThreeArguments.Contract.contract.Transact(opts, method, params...)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ThreeArguments *ThreeArgumentsCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ThreeArguments.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ThreeArguments *ThreeArgumentsSession) Paused() (bool, error) {
	return _ThreeArguments.Contract.Paused(&_ThreeArguments.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ThreeArguments *ThreeArgumentsCallerSession) Paused() (bool, error) {
	return _ThreeArguments.Contract.Paused(&_ThreeArguments.CallOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ThreeArguments *ThreeArgumentsTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ThreeArguments.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ThreeArguments *ThreeArgumentsSession) Pause() (*types.Transaction, error) {
	return _ThreeArguments.Contract.Pause(&_ThreeArguments.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ThreeArguments *ThreeArgumentsTransactorSession) Pause() (*types.Transaction, error) {
	return _ThreeArguments.Contract.Pause(&_ThreeArguments.TransactOpts)
}

// ThreeArguments is a paid mutator transaction binding the contract method 0x9280b905.
//
// Solidity: function threeArguments(string argumentOne, int8 argumentTwo, bool argumentThree) returns()
func (_ThreeArguments *ThreeArgumentsTransactor) ThreeArguments(opts *bind.TransactOpts, argumentOne string, argumentTwo int8, argumentThree bool) (*types.Transaction, error) {
	return _ThreeArguments.contract.Transact(opts, "threeArguments", argumentOne, argumentTwo, argumentThree)
}

// ThreeArguments is a paid mutator transaction binding the contract method 0x9280b905.
//
// Solidity: function threeArguments(string argumentOne, int8 argumentTwo, bool argumentThree) returns()
func (_ThreeArguments *ThreeArgumentsSession) ThreeArguments(argumentOne string, argumentTwo int8, argumentThree bool) (*types.Transaction, error) {
	return _ThreeArguments.Contract.ThreeArguments(&_ThreeArguments.TransactOpts, argumentOne, argumentTwo, argumentThree)
}

// ThreeArguments is a paid mutator transaction binding the contract method 0x9280b905.
//
// Solidity: function threeArguments(string argumentOne, int8 argumentTwo, bool argumentThree) returns()
func (_ThreeArguments *ThreeArgumentsTransactorSession) ThreeArguments(argumentOne string, argumentTwo int8, argumentThree bool) (*types.Transaction, error) {
	return _ThreeArguments.Contract.ThreeArguments(&_ThreeArguments.TransactOpts, argumentOne, argumentTwo, argumentThree)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ThreeArguments *ThreeArgumentsTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ThreeArguments.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ThreeArguments *ThreeArgumentsSession) Unpause() (*types.Transaction, error) {
	return _ThreeArguments.Contract.Unpause(&_ThreeArguments.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ThreeArguments *ThreeArgumentsTransactorSession) Unpause() (*types.Transaction, error) {
	return _ThreeArguments.Contract.Unpause(&_ThreeArguments.TransactOpts)
}

// ThreeArgumentsPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ThreeArguments contract.
type ThreeArgumentsPausedIterator struct {
	Event *ThreeArgumentsPaused // Event containing the contract specifics and raw log

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
func (it *ThreeArgumentsPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ThreeArgumentsPaused)
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
		it.Event = new(ThreeArgumentsPaused)
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
func (it *ThreeArgumentsPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ThreeArgumentsPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ThreeArgumentsPaused represents a Paused event raised by the ThreeArguments contract.
type ThreeArgumentsPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ThreeArguments *ThreeArgumentsFilterer) FilterPaused(opts *bind.FilterOpts) (*ThreeArgumentsPausedIterator, error) {

	logs, sub, err := _ThreeArguments.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &ThreeArgumentsPausedIterator{contract: _ThreeArguments.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ThreeArguments *ThreeArgumentsFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ThreeArgumentsPaused) (event.Subscription, error) {

	logs, sub, err := _ThreeArguments.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ThreeArgumentsPaused)
				if err := _ThreeArguments.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_ThreeArguments *ThreeArgumentsFilterer) ParsePaused(log types.Log) (*ThreeArgumentsPaused, error) {
	event := new(ThreeArgumentsPaused)
	if err := _ThreeArguments.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ThreeArgumentsThreeArgumentsCalledIterator is returned from FilterThreeArgumentsCalled and is used to iterate over the raw logs and unpacked data for ThreeArgumentsCalled events raised by the ThreeArguments contract.
type ThreeArgumentsThreeArgumentsCalledIterator struct {
	Event *ThreeArgumentsThreeArgumentsCalled // Event containing the contract specifics and raw log

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
func (it *ThreeArgumentsThreeArgumentsCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ThreeArgumentsThreeArgumentsCalled)
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
		it.Event = new(ThreeArgumentsThreeArgumentsCalled)
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
func (it *ThreeArgumentsThreeArgumentsCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ThreeArgumentsThreeArgumentsCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ThreeArgumentsThreeArgumentsCalled represents a ThreeArgumentsCalled event raised by the ThreeArguments contract.
type ThreeArgumentsThreeArgumentsCalled struct {
	ArgumentOne   string
	ArgumentTwo   int8
	ArgumentThree bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterThreeArgumentsCalled is a free log retrieval operation binding the contract event 0xd589183661fa75f94e2db32f4eb7ebb50f4154c160e15eb43f772a46f360a3a8.
//
// Solidity: event ThreeArgumentsCalled(string argumentOne, int8 argumentTwo, bool argumentThree)
func (_ThreeArguments *ThreeArgumentsFilterer) FilterThreeArgumentsCalled(opts *bind.FilterOpts) (*ThreeArgumentsThreeArgumentsCalledIterator, error) {

	logs, sub, err := _ThreeArguments.contract.FilterLogs(opts, "ThreeArgumentsCalled")
	if err != nil {
		return nil, err
	}
	return &ThreeArgumentsThreeArgumentsCalledIterator{contract: _ThreeArguments.contract, event: "ThreeArgumentsCalled", logs: logs, sub: sub}, nil
}

// WatchThreeArgumentsCalled is a free log subscription operation binding the contract event 0xd589183661fa75f94e2db32f4eb7ebb50f4154c160e15eb43f772a46f360a3a8.
//
// Solidity: event ThreeArgumentsCalled(string argumentOne, int8 argumentTwo, bool argumentThree)
func (_ThreeArguments *ThreeArgumentsFilterer) WatchThreeArgumentsCalled(opts *bind.WatchOpts, sink chan<- *ThreeArgumentsThreeArgumentsCalled) (event.Subscription, error) {

	logs, sub, err := _ThreeArguments.contract.WatchLogs(opts, "ThreeArgumentsCalled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ThreeArgumentsThreeArgumentsCalled)
				if err := _ThreeArguments.contract.UnpackLog(event, "ThreeArgumentsCalled", log); err != nil {
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

// ParseThreeArgumentsCalled is a log parse operation binding the contract event 0xd589183661fa75f94e2db32f4eb7ebb50f4154c160e15eb43f772a46f360a3a8.
//
// Solidity: event ThreeArgumentsCalled(string argumentOne, int8 argumentTwo, bool argumentThree)
func (_ThreeArguments *ThreeArgumentsFilterer) ParseThreeArgumentsCalled(log types.Log) (*ThreeArgumentsThreeArgumentsCalled, error) {
	event := new(ThreeArgumentsThreeArgumentsCalled)
	if err := _ThreeArguments.contract.UnpackLog(event, "ThreeArgumentsCalled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ThreeArgumentsUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ThreeArguments contract.
type ThreeArgumentsUnpausedIterator struct {
	Event *ThreeArgumentsUnpaused // Event containing the contract specifics and raw log

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
func (it *ThreeArgumentsUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ThreeArgumentsUnpaused)
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
		it.Event = new(ThreeArgumentsUnpaused)
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
func (it *ThreeArgumentsUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ThreeArgumentsUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ThreeArgumentsUnpaused represents a Unpaused event raised by the ThreeArguments contract.
type ThreeArgumentsUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ThreeArguments *ThreeArgumentsFilterer) FilterUnpaused(opts *bind.FilterOpts) (*ThreeArgumentsUnpausedIterator, error) {

	logs, sub, err := _ThreeArguments.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &ThreeArgumentsUnpausedIterator{contract: _ThreeArguments.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ThreeArguments *ThreeArgumentsFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ThreeArgumentsUnpaused) (event.Subscription, error) {

	logs, sub, err := _ThreeArguments.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ThreeArgumentsUnpaused)
				if err := _ThreeArguments.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_ThreeArguments *ThreeArgumentsFilterer) ParseUnpaused(log types.Log) (*ThreeArgumentsUnpaused, error) {
	event := new(ThreeArgumentsUnpaused)
	if err := _ThreeArguments.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
