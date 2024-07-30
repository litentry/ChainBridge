// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package TwoArguments

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

// TwoArgumentsMetaData contains all meta data concerning the TwoArguments contract.
var TwoArgumentsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"argumentOne\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"argumentTwo\",\"type\":\"bytes4\"}],\"name\":\"TwoArgumentsCalled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"argumentOne\",\"type\":\"address[]\"},{\"internalType\":\"bytes4\",\"name\":\"argumentTwo\",\"type\":\"bytes4\"}],\"name\":\"twoArguments\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060008060006101000a81548160ff0219169083151502179055506106788061003a6000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80633f4ba83a146100515780635c975abb1461005b57806372e0745c146100795780638456cb5914610095575b600080fd5b61005961009f565b005b6100636100a9565b604051610070919061028b565b60405180910390f35b610093600480360381019061008e919061036d565b6100bf565b005b61009d610107565b005b6100a7610111565b565b60008060009054906101000a900460ff16905090565b6100c7610173565b7fc983106aca50fad459fb18ede1630e8ff8147ff28ad451a856427931fd7f15e38383836040516100fa939291906104fd565b60405180910390a1505050565b61010f6101bd565b565b61011961021f565b60008060006101000a81548160ff0219169083151502179055507f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa61015c610268565b604051610169919061053e565b60405180910390a1565b61017b6100a9565b156101bb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101b2906105b6565b60405180910390fd5b565b6101c5610173565b60016000806101000a81548160ff0219169083151502179055507f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258610208610268565b604051610215919061053e565b60405180910390a1565b6102276100a9565b610266576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161025d90610622565b60405180910390fd5b565b600033905090565b60008115159050919050565b61028581610270565b82525050565b60006020820190506102a0600083018461027c565b92915050565b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b60008083601f8401126102d5576102d46102b0565b5b8235905067ffffffffffffffff8111156102f2576102f16102b5565b5b60208301915083602082028301111561030e5761030d6102ba565b5b9250929050565b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b61034a81610315565b811461035557600080fd5b50565b60008135905061036781610341565b92915050565b600080600060408486031215610386576103856102a6565b5b600084013567ffffffffffffffff8111156103a4576103a36102ab565b5b6103b0868287016102bf565b935093505060206103c386828701610358565b9150509250925092565b600082825260208201905092915050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610413826103e8565b9050919050565b61042381610408565b82525050565b6000610435838361041a565b60208301905092915050565b61044a81610408565b811461045557600080fd5b50565b60008135905061046781610441565b92915050565b600061047c6020840184610458565b905092915050565b6000602082019050919050565b600061049d83856103cd565b93506104a8826103de565b8060005b858110156104e1576104be828461046d565b6104c88882610429565b97506104d383610484565b9250506001810190506104ac565b5085925050509392505050565b6104f781610315565b82525050565b60006040820190508181036000830152610518818587610491565b905061052760208301846104ee565b949350505050565b61053881610408565b82525050565b6000602082019050610553600083018461052f565b92915050565b600082825260208201905092915050565b7f5061757361626c653a2070617573656400000000000000000000000000000000600082015250565b60006105a0601083610559565b91506105ab8261056a565b602082019050919050565b600060208201905081810360008301526105cf81610593565b9050919050565b7f5061757361626c653a206e6f7420706175736564000000000000000000000000600082015250565b600061060c601483610559565b9150610617826105d6565b602082019050919050565b6000602082019050818103600083015261063b816105ff565b905091905056fea2646970667358221220a132c2e59b956c5feb380845a1b51fcb2e3e19a0799751d8fda60b2c722f927064736f6c63430008130033",
}

// TwoArgumentsABI is the input ABI used to generate the binding from.
// Deprecated: Use TwoArgumentsMetaData.ABI instead.
var TwoArgumentsABI = TwoArgumentsMetaData.ABI

// TwoArgumentsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TwoArgumentsMetaData.Bin instead.
var TwoArgumentsBin = TwoArgumentsMetaData.Bin

// DeployTwoArguments deploys a new Ethereum contract, binding an instance of TwoArguments to it.
func DeployTwoArguments(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TwoArguments, error) {
	parsed, err := TwoArgumentsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TwoArgumentsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TwoArguments{TwoArgumentsCaller: TwoArgumentsCaller{contract: contract}, TwoArgumentsTransactor: TwoArgumentsTransactor{contract: contract}, TwoArgumentsFilterer: TwoArgumentsFilterer{contract: contract}}, nil
}

// TwoArguments is an auto generated Go binding around an Ethereum contract.
type TwoArguments struct {
	TwoArgumentsCaller     // Read-only binding to the contract
	TwoArgumentsTransactor // Write-only binding to the contract
	TwoArgumentsFilterer   // Log filterer for contract events
}

// TwoArgumentsCaller is an auto generated read-only Go binding around an Ethereum contract.
type TwoArgumentsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TwoArgumentsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TwoArgumentsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TwoArgumentsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TwoArgumentsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TwoArgumentsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TwoArgumentsSession struct {
	Contract     *TwoArguments     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TwoArgumentsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TwoArgumentsCallerSession struct {
	Contract *TwoArgumentsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// TwoArgumentsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TwoArgumentsTransactorSession struct {
	Contract     *TwoArgumentsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// TwoArgumentsRaw is an auto generated low-level Go binding around an Ethereum contract.
type TwoArgumentsRaw struct {
	Contract *TwoArguments // Generic contract binding to access the raw methods on
}

// TwoArgumentsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TwoArgumentsCallerRaw struct {
	Contract *TwoArgumentsCaller // Generic read-only contract binding to access the raw methods on
}

// TwoArgumentsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TwoArgumentsTransactorRaw struct {
	Contract *TwoArgumentsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTwoArguments creates a new instance of TwoArguments, bound to a specific deployed contract.
func NewTwoArguments(address common.Address, backend bind.ContractBackend) (*TwoArguments, error) {
	contract, err := bindTwoArguments(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TwoArguments{TwoArgumentsCaller: TwoArgumentsCaller{contract: contract}, TwoArgumentsTransactor: TwoArgumentsTransactor{contract: contract}, TwoArgumentsFilterer: TwoArgumentsFilterer{contract: contract}}, nil
}

// NewTwoArgumentsCaller creates a new read-only instance of TwoArguments, bound to a specific deployed contract.
func NewTwoArgumentsCaller(address common.Address, caller bind.ContractCaller) (*TwoArgumentsCaller, error) {
	contract, err := bindTwoArguments(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TwoArgumentsCaller{contract: contract}, nil
}

// NewTwoArgumentsTransactor creates a new write-only instance of TwoArguments, bound to a specific deployed contract.
func NewTwoArgumentsTransactor(address common.Address, transactor bind.ContractTransactor) (*TwoArgumentsTransactor, error) {
	contract, err := bindTwoArguments(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TwoArgumentsTransactor{contract: contract}, nil
}

// NewTwoArgumentsFilterer creates a new log filterer instance of TwoArguments, bound to a specific deployed contract.
func NewTwoArgumentsFilterer(address common.Address, filterer bind.ContractFilterer) (*TwoArgumentsFilterer, error) {
	contract, err := bindTwoArguments(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TwoArgumentsFilterer{contract: contract}, nil
}

// bindTwoArguments binds a generic wrapper to an already deployed contract.
func bindTwoArguments(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TwoArgumentsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TwoArguments *TwoArgumentsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TwoArguments.Contract.TwoArgumentsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TwoArguments *TwoArgumentsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TwoArguments.Contract.TwoArgumentsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TwoArguments *TwoArgumentsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TwoArguments.Contract.TwoArgumentsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TwoArguments *TwoArgumentsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TwoArguments.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TwoArguments *TwoArgumentsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TwoArguments.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TwoArguments *TwoArgumentsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TwoArguments.Contract.contract.Transact(opts, method, params...)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_TwoArguments *TwoArgumentsCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _TwoArguments.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_TwoArguments *TwoArgumentsSession) Paused() (bool, error) {
	return _TwoArguments.Contract.Paused(&_TwoArguments.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_TwoArguments *TwoArgumentsCallerSession) Paused() (bool, error) {
	return _TwoArguments.Contract.Paused(&_TwoArguments.CallOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_TwoArguments *TwoArgumentsTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TwoArguments.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_TwoArguments *TwoArgumentsSession) Pause() (*types.Transaction, error) {
	return _TwoArguments.Contract.Pause(&_TwoArguments.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_TwoArguments *TwoArgumentsTransactorSession) Pause() (*types.Transaction, error) {
	return _TwoArguments.Contract.Pause(&_TwoArguments.TransactOpts)
}

// TwoArguments is a paid mutator transaction binding the contract method 0x72e0745c.
//
// Solidity: function twoArguments(address[] argumentOne, bytes4 argumentTwo) returns()
func (_TwoArguments *TwoArgumentsTransactor) TwoArguments(opts *bind.TransactOpts, argumentOne []common.Address, argumentTwo [4]byte) (*types.Transaction, error) {
	return _TwoArguments.contract.Transact(opts, "twoArguments", argumentOne, argumentTwo)
}

// TwoArguments is a paid mutator transaction binding the contract method 0x72e0745c.
//
// Solidity: function twoArguments(address[] argumentOne, bytes4 argumentTwo) returns()
func (_TwoArguments *TwoArgumentsSession) TwoArguments(argumentOne []common.Address, argumentTwo [4]byte) (*types.Transaction, error) {
	return _TwoArguments.Contract.TwoArguments(&_TwoArguments.TransactOpts, argumentOne, argumentTwo)
}

// TwoArguments is a paid mutator transaction binding the contract method 0x72e0745c.
//
// Solidity: function twoArguments(address[] argumentOne, bytes4 argumentTwo) returns()
func (_TwoArguments *TwoArgumentsTransactorSession) TwoArguments(argumentOne []common.Address, argumentTwo [4]byte) (*types.Transaction, error) {
	return _TwoArguments.Contract.TwoArguments(&_TwoArguments.TransactOpts, argumentOne, argumentTwo)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_TwoArguments *TwoArgumentsTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TwoArguments.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_TwoArguments *TwoArgumentsSession) Unpause() (*types.Transaction, error) {
	return _TwoArguments.Contract.Unpause(&_TwoArguments.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_TwoArguments *TwoArgumentsTransactorSession) Unpause() (*types.Transaction, error) {
	return _TwoArguments.Contract.Unpause(&_TwoArguments.TransactOpts)
}

// TwoArgumentsPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the TwoArguments contract.
type TwoArgumentsPausedIterator struct {
	Event *TwoArgumentsPaused // Event containing the contract specifics and raw log

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
func (it *TwoArgumentsPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TwoArgumentsPaused)
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
		it.Event = new(TwoArgumentsPaused)
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
func (it *TwoArgumentsPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TwoArgumentsPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TwoArgumentsPaused represents a Paused event raised by the TwoArguments contract.
type TwoArgumentsPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_TwoArguments *TwoArgumentsFilterer) FilterPaused(opts *bind.FilterOpts) (*TwoArgumentsPausedIterator, error) {

	logs, sub, err := _TwoArguments.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &TwoArgumentsPausedIterator{contract: _TwoArguments.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_TwoArguments *TwoArgumentsFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *TwoArgumentsPaused) (event.Subscription, error) {

	logs, sub, err := _TwoArguments.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TwoArgumentsPaused)
				if err := _TwoArguments.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_TwoArguments *TwoArgumentsFilterer) ParsePaused(log types.Log) (*TwoArgumentsPaused, error) {
	event := new(TwoArgumentsPaused)
	if err := _TwoArguments.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TwoArgumentsTwoArgumentsCalledIterator is returned from FilterTwoArgumentsCalled and is used to iterate over the raw logs and unpacked data for TwoArgumentsCalled events raised by the TwoArguments contract.
type TwoArgumentsTwoArgumentsCalledIterator struct {
	Event *TwoArgumentsTwoArgumentsCalled // Event containing the contract specifics and raw log

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
func (it *TwoArgumentsTwoArgumentsCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TwoArgumentsTwoArgumentsCalled)
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
		it.Event = new(TwoArgumentsTwoArgumentsCalled)
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
func (it *TwoArgumentsTwoArgumentsCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TwoArgumentsTwoArgumentsCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TwoArgumentsTwoArgumentsCalled represents a TwoArgumentsCalled event raised by the TwoArguments contract.
type TwoArgumentsTwoArgumentsCalled struct {
	ArgumentOne []common.Address
	ArgumentTwo [4]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTwoArgumentsCalled is a free log retrieval operation binding the contract event 0xc983106aca50fad459fb18ede1630e8ff8147ff28ad451a856427931fd7f15e3.
//
// Solidity: event TwoArgumentsCalled(address[] argumentOne, bytes4 argumentTwo)
func (_TwoArguments *TwoArgumentsFilterer) FilterTwoArgumentsCalled(opts *bind.FilterOpts) (*TwoArgumentsTwoArgumentsCalledIterator, error) {

	logs, sub, err := _TwoArguments.contract.FilterLogs(opts, "TwoArgumentsCalled")
	if err != nil {
		return nil, err
	}
	return &TwoArgumentsTwoArgumentsCalledIterator{contract: _TwoArguments.contract, event: "TwoArgumentsCalled", logs: logs, sub: sub}, nil
}

// WatchTwoArgumentsCalled is a free log subscription operation binding the contract event 0xc983106aca50fad459fb18ede1630e8ff8147ff28ad451a856427931fd7f15e3.
//
// Solidity: event TwoArgumentsCalled(address[] argumentOne, bytes4 argumentTwo)
func (_TwoArguments *TwoArgumentsFilterer) WatchTwoArgumentsCalled(opts *bind.WatchOpts, sink chan<- *TwoArgumentsTwoArgumentsCalled) (event.Subscription, error) {

	logs, sub, err := _TwoArguments.contract.WatchLogs(opts, "TwoArgumentsCalled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TwoArgumentsTwoArgumentsCalled)
				if err := _TwoArguments.contract.UnpackLog(event, "TwoArgumentsCalled", log); err != nil {
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

// ParseTwoArgumentsCalled is a log parse operation binding the contract event 0xc983106aca50fad459fb18ede1630e8ff8147ff28ad451a856427931fd7f15e3.
//
// Solidity: event TwoArgumentsCalled(address[] argumentOne, bytes4 argumentTwo)
func (_TwoArguments *TwoArgumentsFilterer) ParseTwoArgumentsCalled(log types.Log) (*TwoArgumentsTwoArgumentsCalled, error) {
	event := new(TwoArgumentsTwoArgumentsCalled)
	if err := _TwoArguments.contract.UnpackLog(event, "TwoArgumentsCalled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TwoArgumentsUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the TwoArguments contract.
type TwoArgumentsUnpausedIterator struct {
	Event *TwoArgumentsUnpaused // Event containing the contract specifics and raw log

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
func (it *TwoArgumentsUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TwoArgumentsUnpaused)
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
		it.Event = new(TwoArgumentsUnpaused)
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
func (it *TwoArgumentsUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TwoArgumentsUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TwoArgumentsUnpaused represents a Unpaused event raised by the TwoArguments contract.
type TwoArgumentsUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_TwoArguments *TwoArgumentsFilterer) FilterUnpaused(opts *bind.FilterOpts) (*TwoArgumentsUnpausedIterator, error) {

	logs, sub, err := _TwoArguments.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &TwoArgumentsUnpausedIterator{contract: _TwoArguments.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_TwoArguments *TwoArgumentsFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *TwoArgumentsUnpaused) (event.Subscription, error) {

	logs, sub, err := _TwoArguments.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TwoArgumentsUnpaused)
				if err := _TwoArguments.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_TwoArguments *TwoArgumentsFilterer) ParseUnpaused(log types.Log) (*TwoArgumentsUnpaused, error) {
	event := new(TwoArgumentsUnpaused)
	if err := _TwoArguments.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
