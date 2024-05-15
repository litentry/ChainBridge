// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package CentrifugeAsset

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

// CentrifugeAssetMetaData contains all meta data concerning the CentrifugeAsset contract.
var CentrifugeAssetMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"asset\",\"type\":\"bytes32\"}],\"name\":\"AssetStored\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_assetsStored\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"asset\",\"type\":\"bytes32\"}],\"name\":\"store\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060008060006101000a81548160ff0219169083151502179055506105cc8061003a6000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c80633f4ba83a1461005c5780635c975abb14610066578063654cf88c146100845780638456cb59146100a057806396add600146100aa575b600080fd5b6100646100da565b005b61006e6100e4565b60405161007b9190610362565b60405180910390f35b61009e600480360381019061009991906103b8565b6100fa565b005b6100a86101be565b005b6100c460048036038101906100bf91906103b8565b6101c8565b6040516100d19190610362565b60405180910390f35b6100e26101e8565b565b60008060009054906101000a900460ff16905090565b61010261024a565b6001600082815260200190815260200160002060009054906101000a900460ff1615610163576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161015a90610442565b60405180910390fd5b600180600083815260200190815260200160002060006101000a81548160ff021916908315150217905550807f08ae553713effae7116be03743b167b8b803449ee8fb912c2ec43dc2c824f53560405160405180910390a250565b6101c6610294565b565b60016020528060005260406000206000915054906101000a900460ff1681565b6101f06102f6565b60008060006101000a81548160ff0219169083151502179055507f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa61023361033f565b60405161024091906104a3565b60405180910390a1565b6102526100e4565b15610292576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102899061050a565b60405180910390fd5b565b61029c61024a565b60016000806101000a81548160ff0219169083151502179055507f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586102df61033f565b6040516102ec91906104a3565b60405180910390a1565b6102fe6100e4565b61033d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161033490610576565b60405180910390fd5b565b600033905090565b60008115159050919050565b61035c81610347565b82525050565b60006020820190506103776000830184610353565b92915050565b600080fd5b6000819050919050565b61039581610382565b81146103a057600080fd5b50565b6000813590506103b28161038c565b92915050565b6000602082840312156103ce576103cd61037d565b5b60006103dc848285016103a3565b91505092915050565b600082825260208201905092915050565b7f617373657420697320616c72656164792073746f726564000000000000000000600082015250565b600061042c6017836103e5565b9150610437826103f6565b602082019050919050565b6000602082019050818103600083015261045b8161041f565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061048d82610462565b9050919050565b61049d81610482565b82525050565b60006020820190506104b86000830184610494565b92915050565b7f5061757361626c653a2070617573656400000000000000000000000000000000600082015250565b60006104f46010836103e5565b91506104ff826104be565b602082019050919050565b60006020820190508181036000830152610523816104e7565b9050919050565b7f5061757361626c653a206e6f7420706175736564000000000000000000000000600082015250565b60006105606014836103e5565b915061056b8261052a565b602082019050919050565b6000602082019050818103600083015261058f81610553565b905091905056fea2646970667358221220cdc1a390d1214d178c47b6bcf5e012d7dcef39213c8bd5e02b32e2267be8ecb264736f6c63430008130033",
}

// CentrifugeAssetABI is the input ABI used to generate the binding from.
// Deprecated: Use CentrifugeAssetMetaData.ABI instead.
var CentrifugeAssetABI = CentrifugeAssetMetaData.ABI

// CentrifugeAssetBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CentrifugeAssetMetaData.Bin instead.
var CentrifugeAssetBin = CentrifugeAssetMetaData.Bin

// DeployCentrifugeAsset deploys a new Ethereum contract, binding an instance of CentrifugeAsset to it.
func DeployCentrifugeAsset(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CentrifugeAsset, error) {
	parsed, err := CentrifugeAssetMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CentrifugeAssetBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CentrifugeAsset{CentrifugeAssetCaller: CentrifugeAssetCaller{contract: contract}, CentrifugeAssetTransactor: CentrifugeAssetTransactor{contract: contract}, CentrifugeAssetFilterer: CentrifugeAssetFilterer{contract: contract}}, nil
}

// CentrifugeAsset is an auto generated Go binding around an Ethereum contract.
type CentrifugeAsset struct {
	CentrifugeAssetCaller     // Read-only binding to the contract
	CentrifugeAssetTransactor // Write-only binding to the contract
	CentrifugeAssetFilterer   // Log filterer for contract events
}

// CentrifugeAssetCaller is an auto generated read-only Go binding around an Ethereum contract.
type CentrifugeAssetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CentrifugeAssetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CentrifugeAssetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CentrifugeAssetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CentrifugeAssetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CentrifugeAssetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CentrifugeAssetSession struct {
	Contract     *CentrifugeAsset  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CentrifugeAssetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CentrifugeAssetCallerSession struct {
	Contract *CentrifugeAssetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// CentrifugeAssetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CentrifugeAssetTransactorSession struct {
	Contract     *CentrifugeAssetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// CentrifugeAssetRaw is an auto generated low-level Go binding around an Ethereum contract.
type CentrifugeAssetRaw struct {
	Contract *CentrifugeAsset // Generic contract binding to access the raw methods on
}

// CentrifugeAssetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CentrifugeAssetCallerRaw struct {
	Contract *CentrifugeAssetCaller // Generic read-only contract binding to access the raw methods on
}

// CentrifugeAssetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CentrifugeAssetTransactorRaw struct {
	Contract *CentrifugeAssetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCentrifugeAsset creates a new instance of CentrifugeAsset, bound to a specific deployed contract.
func NewCentrifugeAsset(address common.Address, backend bind.ContractBackend) (*CentrifugeAsset, error) {
	contract, err := bindCentrifugeAsset(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CentrifugeAsset{CentrifugeAssetCaller: CentrifugeAssetCaller{contract: contract}, CentrifugeAssetTransactor: CentrifugeAssetTransactor{contract: contract}, CentrifugeAssetFilterer: CentrifugeAssetFilterer{contract: contract}}, nil
}

// NewCentrifugeAssetCaller creates a new read-only instance of CentrifugeAsset, bound to a specific deployed contract.
func NewCentrifugeAssetCaller(address common.Address, caller bind.ContractCaller) (*CentrifugeAssetCaller, error) {
	contract, err := bindCentrifugeAsset(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CentrifugeAssetCaller{contract: contract}, nil
}

// NewCentrifugeAssetTransactor creates a new write-only instance of CentrifugeAsset, bound to a specific deployed contract.
func NewCentrifugeAssetTransactor(address common.Address, transactor bind.ContractTransactor) (*CentrifugeAssetTransactor, error) {
	contract, err := bindCentrifugeAsset(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CentrifugeAssetTransactor{contract: contract}, nil
}

// NewCentrifugeAssetFilterer creates a new log filterer instance of CentrifugeAsset, bound to a specific deployed contract.
func NewCentrifugeAssetFilterer(address common.Address, filterer bind.ContractFilterer) (*CentrifugeAssetFilterer, error) {
	contract, err := bindCentrifugeAsset(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CentrifugeAssetFilterer{contract: contract}, nil
}

// bindCentrifugeAsset binds a generic wrapper to an already deployed contract.
func bindCentrifugeAsset(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CentrifugeAssetMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CentrifugeAsset *CentrifugeAssetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CentrifugeAsset.Contract.CentrifugeAssetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CentrifugeAsset *CentrifugeAssetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CentrifugeAsset.Contract.CentrifugeAssetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CentrifugeAsset *CentrifugeAssetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CentrifugeAsset.Contract.CentrifugeAssetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CentrifugeAsset *CentrifugeAssetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CentrifugeAsset.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CentrifugeAsset *CentrifugeAssetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CentrifugeAsset.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CentrifugeAsset *CentrifugeAssetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CentrifugeAsset.Contract.contract.Transact(opts, method, params...)
}

// AssetsStored is a free data retrieval call binding the contract method 0x96add600.
//
// Solidity: function _assetsStored(bytes32 ) view returns(bool)
func (_CentrifugeAsset *CentrifugeAssetCaller) AssetsStored(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _CentrifugeAsset.contract.Call(opts, &out, "_assetsStored", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AssetsStored is a free data retrieval call binding the contract method 0x96add600.
//
// Solidity: function _assetsStored(bytes32 ) view returns(bool)
func (_CentrifugeAsset *CentrifugeAssetSession) AssetsStored(arg0 [32]byte) (bool, error) {
	return _CentrifugeAsset.Contract.AssetsStored(&_CentrifugeAsset.CallOpts, arg0)
}

// AssetsStored is a free data retrieval call binding the contract method 0x96add600.
//
// Solidity: function _assetsStored(bytes32 ) view returns(bool)
func (_CentrifugeAsset *CentrifugeAssetCallerSession) AssetsStored(arg0 [32]byte) (bool, error) {
	return _CentrifugeAsset.Contract.AssetsStored(&_CentrifugeAsset.CallOpts, arg0)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CentrifugeAsset *CentrifugeAssetCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CentrifugeAsset.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CentrifugeAsset *CentrifugeAssetSession) Paused() (bool, error) {
	return _CentrifugeAsset.Contract.Paused(&_CentrifugeAsset.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CentrifugeAsset *CentrifugeAssetCallerSession) Paused() (bool, error) {
	return _CentrifugeAsset.Contract.Paused(&_CentrifugeAsset.CallOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CentrifugeAsset *CentrifugeAssetTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CentrifugeAsset.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CentrifugeAsset *CentrifugeAssetSession) Pause() (*types.Transaction, error) {
	return _CentrifugeAsset.Contract.Pause(&_CentrifugeAsset.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CentrifugeAsset *CentrifugeAssetTransactorSession) Pause() (*types.Transaction, error) {
	return _CentrifugeAsset.Contract.Pause(&_CentrifugeAsset.TransactOpts)
}

// Store is a paid mutator transaction binding the contract method 0x654cf88c.
//
// Solidity: function store(bytes32 asset) returns()
func (_CentrifugeAsset *CentrifugeAssetTransactor) Store(opts *bind.TransactOpts, asset [32]byte) (*types.Transaction, error) {
	return _CentrifugeAsset.contract.Transact(opts, "store", asset)
}

// Store is a paid mutator transaction binding the contract method 0x654cf88c.
//
// Solidity: function store(bytes32 asset) returns()
func (_CentrifugeAsset *CentrifugeAssetSession) Store(asset [32]byte) (*types.Transaction, error) {
	return _CentrifugeAsset.Contract.Store(&_CentrifugeAsset.TransactOpts, asset)
}

// Store is a paid mutator transaction binding the contract method 0x654cf88c.
//
// Solidity: function store(bytes32 asset) returns()
func (_CentrifugeAsset *CentrifugeAssetTransactorSession) Store(asset [32]byte) (*types.Transaction, error) {
	return _CentrifugeAsset.Contract.Store(&_CentrifugeAsset.TransactOpts, asset)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CentrifugeAsset *CentrifugeAssetTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CentrifugeAsset.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CentrifugeAsset *CentrifugeAssetSession) Unpause() (*types.Transaction, error) {
	return _CentrifugeAsset.Contract.Unpause(&_CentrifugeAsset.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CentrifugeAsset *CentrifugeAssetTransactorSession) Unpause() (*types.Transaction, error) {
	return _CentrifugeAsset.Contract.Unpause(&_CentrifugeAsset.TransactOpts)
}

// CentrifugeAssetAssetStoredIterator is returned from FilterAssetStored and is used to iterate over the raw logs and unpacked data for AssetStored events raised by the CentrifugeAsset contract.
type CentrifugeAssetAssetStoredIterator struct {
	Event *CentrifugeAssetAssetStored // Event containing the contract specifics and raw log

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
func (it *CentrifugeAssetAssetStoredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CentrifugeAssetAssetStored)
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
		it.Event = new(CentrifugeAssetAssetStored)
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
func (it *CentrifugeAssetAssetStoredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CentrifugeAssetAssetStoredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CentrifugeAssetAssetStored represents a AssetStored event raised by the CentrifugeAsset contract.
type CentrifugeAssetAssetStored struct {
	Asset [32]byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAssetStored is a free log retrieval operation binding the contract event 0x08ae553713effae7116be03743b167b8b803449ee8fb912c2ec43dc2c824f535.
//
// Solidity: event AssetStored(bytes32 indexed asset)
func (_CentrifugeAsset *CentrifugeAssetFilterer) FilterAssetStored(opts *bind.FilterOpts, asset [][32]byte) (*CentrifugeAssetAssetStoredIterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _CentrifugeAsset.contract.FilterLogs(opts, "AssetStored", assetRule)
	if err != nil {
		return nil, err
	}
	return &CentrifugeAssetAssetStoredIterator{contract: _CentrifugeAsset.contract, event: "AssetStored", logs: logs, sub: sub}, nil
}

// WatchAssetStored is a free log subscription operation binding the contract event 0x08ae553713effae7116be03743b167b8b803449ee8fb912c2ec43dc2c824f535.
//
// Solidity: event AssetStored(bytes32 indexed asset)
func (_CentrifugeAsset *CentrifugeAssetFilterer) WatchAssetStored(opts *bind.WatchOpts, sink chan<- *CentrifugeAssetAssetStored, asset [][32]byte) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _CentrifugeAsset.contract.WatchLogs(opts, "AssetStored", assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CentrifugeAssetAssetStored)
				if err := _CentrifugeAsset.contract.UnpackLog(event, "AssetStored", log); err != nil {
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

// ParseAssetStored is a log parse operation binding the contract event 0x08ae553713effae7116be03743b167b8b803449ee8fb912c2ec43dc2c824f535.
//
// Solidity: event AssetStored(bytes32 indexed asset)
func (_CentrifugeAsset *CentrifugeAssetFilterer) ParseAssetStored(log types.Log) (*CentrifugeAssetAssetStored, error) {
	event := new(CentrifugeAssetAssetStored)
	if err := _CentrifugeAsset.contract.UnpackLog(event, "AssetStored", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CentrifugeAssetPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the CentrifugeAsset contract.
type CentrifugeAssetPausedIterator struct {
	Event *CentrifugeAssetPaused // Event containing the contract specifics and raw log

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
func (it *CentrifugeAssetPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CentrifugeAssetPaused)
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
		it.Event = new(CentrifugeAssetPaused)
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
func (it *CentrifugeAssetPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CentrifugeAssetPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CentrifugeAssetPaused represents a Paused event raised by the CentrifugeAsset contract.
type CentrifugeAssetPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_CentrifugeAsset *CentrifugeAssetFilterer) FilterPaused(opts *bind.FilterOpts) (*CentrifugeAssetPausedIterator, error) {

	logs, sub, err := _CentrifugeAsset.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &CentrifugeAssetPausedIterator{contract: _CentrifugeAsset.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_CentrifugeAsset *CentrifugeAssetFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *CentrifugeAssetPaused) (event.Subscription, error) {

	logs, sub, err := _CentrifugeAsset.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CentrifugeAssetPaused)
				if err := _CentrifugeAsset.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_CentrifugeAsset *CentrifugeAssetFilterer) ParsePaused(log types.Log) (*CentrifugeAssetPaused, error) {
	event := new(CentrifugeAssetPaused)
	if err := _CentrifugeAsset.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CentrifugeAssetUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the CentrifugeAsset contract.
type CentrifugeAssetUnpausedIterator struct {
	Event *CentrifugeAssetUnpaused // Event containing the contract specifics and raw log

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
func (it *CentrifugeAssetUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CentrifugeAssetUnpaused)
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
		it.Event = new(CentrifugeAssetUnpaused)
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
func (it *CentrifugeAssetUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CentrifugeAssetUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CentrifugeAssetUnpaused represents a Unpaused event raised by the CentrifugeAsset contract.
type CentrifugeAssetUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_CentrifugeAsset *CentrifugeAssetFilterer) FilterUnpaused(opts *bind.FilterOpts) (*CentrifugeAssetUnpausedIterator, error) {

	logs, sub, err := _CentrifugeAsset.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &CentrifugeAssetUnpausedIterator{contract: _CentrifugeAsset.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_CentrifugeAsset *CentrifugeAssetFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *CentrifugeAssetUnpaused) (event.Subscription, error) {

	logs, sub, err := _CentrifugeAsset.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CentrifugeAssetUnpaused)
				if err := _CentrifugeAsset.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_CentrifugeAsset *CentrifugeAssetFilterer) ParseUnpaused(log types.Log) (*CentrifugeAssetUnpaused, error) {
	event := new(CentrifugeAssetUnpaused)
	if err := _CentrifugeAsset.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
