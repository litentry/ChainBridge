// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package PhalaBTCLottery

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
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// PhalaBTCLotteryABI is the input ABI used to generate the binding from.
const PhalaBTCLotteryABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_bridge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_genericHandler\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"roundId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"totalCount\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"winnerCount\",\"type\":\"uint32\"}],\"name\":\"NewRound\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"roundId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"tokenId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"btcAddr\",\"type\":\"bytes\"}],\"name\":\"OpenLottery\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"roundId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"tokenId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"signedTx\",\"type\":\"bytes\"}],\"name\":\"SignedTxStored\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"openedBox\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"storedTx\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"txStorage\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"depositHandler\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"roundId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"tokenId\",\"type\":\"uint32\"}],\"name\":\"isLotteryOpened\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"recordSignedBTCTx\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// PhalaBTCLotteryBin is the compiled bytecode used for deploying new contracts.
var PhalaBTCLotteryBin = "0x60806040523480156200001157600080fd5b506040516200136e3803806200136e8339818101604052810190620000379190620000d8565b826000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050506200017c565b600081519050620000d28162000162565b92915050565b600080600060608486031215620000ee57600080fd5b6000620000fe86828701620000c1565b93505060206200011186828701620000c1565b92505060406200012486828701620000c1565b9150509250925092565b60006200013b8262000142565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6200016d816200012e565b81146200017957600080fd5b50565b6111e2806200018c6000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c806312aad3d4146100675780639a4ae77714610097578063a2b58507146100c7578063c066b75f146100e3578063f2db379714610113578063f4190f8314610143575b600080fd5b610081600480360381019061007c9190610bc3565b61015f565b60405161008e9190610eee565b60405180910390f35b6100b160048036038101906100ac9190610bc3565b61018e565b6040516100be9190610f09565b60405180910390f35b6100e160048036038101906100dc9190610b82565b61024b565b005b6100fd60048036038101906100f89190610bc3565b610355565b60405161010a9190610eee565b60405180910390f35b61012d60048036038101906101289190610bc3565b610384565b60405161013a9190610eee565b60405180910390f35b61015d60048036038101906101589190610b82565b6103d8565b005b60026020528160005260406000206020528060005260406000206000915091509054906101000a900460ff1681565b6004602052816000526040600020602052806000526040600020600091509150508054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156102435780601f1061021857610100808354040283529160200191610243565b820191906000526020600020905b81548152906001019060200180831161022657829003601f168201915b505050505081565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146102db576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102d290610f2b565b60405180910390fd5b60006102f16000836104c790919063ffffffff16565b905060006103096004846104c790919063ffffffff16565b905060006103216008856104c790919063ffffffff16565b90506060610341600c8363ffffffff16876105259092919063ffffffff16565b905061034e848483610631565b5050505050565b60036020528160005260406000206020528060005260406000206000915091509054906101000a900460ff1681565b6000600260008463ffffffff1663ffffffff16815260200190815260200160002060008363ffffffff1663ffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b60006103ee6000836107a190919063ffffffff16565b905060008160ff1614156104575760006104126001846104c790919063ffffffff16565b9050600061042a6005856104c790919063ffffffff16565b905060006104426009866104c790919063ffffffff16565b905061044f8383836107ff565b5050506104c3565b60018160ff1614156104c15760006104796001846104c790919063ffffffff16565b905060006104916005856104c790919063ffffffff16565b905060606104ac60096020876105259092919063ffffffff16565b90506104b983838361095d565b5050506104c2565b5b5b5050565b60006004820183511015610510576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105079061100b565b60405180910390fd5b60008260048501015190508091505092915050565b606081601f8301101561056d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161056490610f6b565b60405180910390fd5b818301845110156105b3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105aa90610fcb565b60405180910390fd5b60608215600081146105d45760405191506000825260208201604052610625565b6040519150601f8416801560200281840101858101878315602002848b0101015b8183101561061257805183526020830192506020810190506105f5565b50868552601f19601f8301166040525050505b50809150509392505050565b600360008463ffffffff1663ffffffff16815260200190815260200160002060008363ffffffff1663ffffffff16815260200190815260200160002060009054906101000a900460ff16156106bb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106b290610fab565b60405180910390fd5b6001600360008563ffffffff1663ffffffff16815260200190815260200160002060008463ffffffff1663ffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555080600460008563ffffffff1663ffffffff16815260200190815260200160002060008463ffffffff1663ffffffff1681526020019081526020016000209080519060200190610760929190610a7c565b507fe7b2b0726b9d9cd1592e2621a382872738fc5bf81fd1d81f9c1f4f962615b8af8383836040516107949392919061102b565b60405180910390a1505050565b600060018201835110156107ea576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107e190610feb565b60405180910390fd5b60008260018501015190508091505092915050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461088f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161088690610f2b565b60405180910390fd5b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff161461091d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161091490610f4b565b60405180910390fd5b7fdb03f7d591b8e310b058b3c3c80e9415acee180707443368d8bae2d55383124083838360405161095093929190611069565b60405180910390a1505050565b600260008463ffffffff1663ffffffff16815260200190815260200160002060008363ffffffff1663ffffffff16815260200190815260200160002060009054906101000a900460ff16156109e7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109de90610f8b565b60405180910390fd5b6001600260008563ffffffff1663ffffffff16815260200190815260200160002060008463ffffffff1663ffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055507f66b4e6a794fe8bbff66bf6eeb41af302d4f7e2bde331cb703b4ded081cac3d77838383604051610a6f9392919061102b565b60405180910390a1505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610abd57805160ff1916838001178555610aeb565b82800160010185558215610aeb579182015b82811115610aea578251825591602001919060010190610acf565b5b509050610af89190610afc565b5090565b5b80821115610b15576000816000905550600101610afd565b5090565b600082601f830112610b2a57600080fd5b8135610b3d610b38826110cd565b6110a0565b91508082526020830160208301858383011115610b5957600080fd5b610b64838284611142565b50505092915050565b600081359050610b7c81611195565b92915050565b600060208284031215610b9457600080fd5b600082013567ffffffffffffffff811115610bae57600080fd5b610bba84828501610b19565b91505092915050565b60008060408385031215610bd657600080fd5b6000610be485828601610b6d565b9250506020610bf585828601610b6d565b9150509250929050565b610c0881611126565b82525050565b6000610c19826110f9565b610c238185611104565b9350610c33818560208601611151565b610c3c81611184565b840191505092915050565b6000610c54604383611115565b91507f5065726d697373696f6e2044656e6965643a204d6573736167652053656e646560008301527f722073686f756c642062652047656e6572696348616e646c657220636f6e747260208301527f61637400000000000000000000000000000000000000000000000000000000006040830152606082019050919050565b6000610ce0603383611115565b91507f5065726d697373696f6e2044656e6965643a205478204f726967696e2073686f60008301527f756c64206265204e465420636f6e7472616374000000000000000000000000006020830152604082019050919050565b6000610d46600e83611115565b91507f736c6963655f6f766572666c6f770000000000000000000000000000000000006000830152602082019050919050565b6000610d86602483611115565b91507f496e76616c69642043616c6c3a2042544320626f7820616c7265616479206f7060008301527f656e6564000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000610dec601f83611115565b91507f496e76616c69642043616c6c3a20547820616c72656164792073746f726564006000830152602082019050919050565b6000610e2c601183611115565b91507f736c6963655f6f75744f66426f756e64730000000000000000000000000000006000830152602082019050919050565b6000610e6c601383611115565b91507f746f55696e74385f6f75744f66426f756e6473000000000000000000000000006000830152602082019050919050565b6000610eac601483611115565b91507f746f55696e7433325f6f75744f66426f756e64730000000000000000000000006000830152602082019050919050565b610ee881611132565b82525050565b6000602082019050610f036000830184610bff565b92915050565b60006020820190508181036000830152610f238184610c0e565b905092915050565b60006020820190508181036000830152610f4481610c47565b9050919050565b60006020820190508181036000830152610f6481610cd3565b9050919050565b60006020820190508181036000830152610f8481610d39565b9050919050565b60006020820190508181036000830152610fa481610d79565b9050919050565b60006020820190508181036000830152610fc481610ddf565b9050919050565b60006020820190508181036000830152610fe481610e1f565b9050919050565b6000602082019050818103600083015261100481610e5f565b9050919050565b6000602082019050818103600083015261102481610e9f565b9050919050565b60006060820190506110406000830186610edf565b61104d6020830185610edf565b818103604083015261105f8184610c0e565b9050949350505050565b600060608201905061107e6000830186610edf565b61108b6020830185610edf565b6110986040830184610edf565b949350505050565b6000604051905081810181811067ffffffffffffffff821117156110c357600080fd5b8060405250919050565b600067ffffffffffffffff8211156110e457600080fd5b601f19601f8301169050602081019050919050565b600081519050919050565b600082825260208201905092915050565b600082825260208201905092915050565b60008115159050919050565b600063ffffffff82169050919050565b82818337600083830152505050565b60005b8381101561116f578082015181840152602081019050611154565b8381111561117e576000848401525b50505050565b6000601f19601f8301169050919050565b61119e81611132565b81146111a957600080fd5b5056fea264697066735822122047eb05eb685d91496d5dfb4cddd1553ea21e9920247674eebd09ec5555e8952764736f6c63430007000033"

// DeployPhalaBTCLottery deploys a new Ethereum contract, binding an instance of PhalaBTCLottery to it.
func DeployPhalaBTCLottery(auth *bind.TransactOpts, backend bind.ContractBackend, _nftContract common.Address, _bridge common.Address, _genericHandler common.Address) (common.Address, *types.Transaction, *PhalaBTCLottery, error) {
	parsed, err := abi.JSON(strings.NewReader(PhalaBTCLotteryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PhalaBTCLotteryBin), backend, _nftContract, _bridge, _genericHandler)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PhalaBTCLottery{PhalaBTCLotteryCaller: PhalaBTCLotteryCaller{contract: contract}, PhalaBTCLotteryTransactor: PhalaBTCLotteryTransactor{contract: contract}, PhalaBTCLotteryFilterer: PhalaBTCLotteryFilterer{contract: contract}}, nil
}

// PhalaBTCLottery is an auto generated Go binding around an Ethereum contract.
type PhalaBTCLottery struct {
	PhalaBTCLotteryCaller     // Read-only binding to the contract
	PhalaBTCLotteryTransactor // Write-only binding to the contract
	PhalaBTCLotteryFilterer   // Log filterer for contract events
}

// PhalaBTCLotteryCaller is an auto generated read-only Go binding around an Ethereum contract.
type PhalaBTCLotteryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PhalaBTCLotteryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PhalaBTCLotteryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PhalaBTCLotteryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PhalaBTCLotteryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PhalaBTCLotterySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PhalaBTCLotterySession struct {
	Contract     *PhalaBTCLottery  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PhalaBTCLotteryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PhalaBTCLotteryCallerSession struct {
	Contract *PhalaBTCLotteryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// PhalaBTCLotteryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PhalaBTCLotteryTransactorSession struct {
	Contract     *PhalaBTCLotteryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// PhalaBTCLotteryRaw is an auto generated low-level Go binding around an Ethereum contract.
type PhalaBTCLotteryRaw struct {
	Contract *PhalaBTCLottery // Generic contract binding to access the raw methods on
}

// PhalaBTCLotteryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PhalaBTCLotteryCallerRaw struct {
	Contract *PhalaBTCLotteryCaller // Generic read-only contract binding to access the raw methods on
}

// PhalaBTCLotteryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PhalaBTCLotteryTransactorRaw struct {
	Contract *PhalaBTCLotteryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPhalaBTCLottery creates a new instance of PhalaBTCLottery, bound to a specific deployed contract.
func NewPhalaBTCLottery(address common.Address, backend bind.ContractBackend) (*PhalaBTCLottery, error) {
	contract, err := bindPhalaBTCLottery(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PhalaBTCLottery{PhalaBTCLotteryCaller: PhalaBTCLotteryCaller{contract: contract}, PhalaBTCLotteryTransactor: PhalaBTCLotteryTransactor{contract: contract}, PhalaBTCLotteryFilterer: PhalaBTCLotteryFilterer{contract: contract}}, nil
}

// NewPhalaBTCLotteryCaller creates a new read-only instance of PhalaBTCLottery, bound to a specific deployed contract.
func NewPhalaBTCLotteryCaller(address common.Address, caller bind.ContractCaller) (*PhalaBTCLotteryCaller, error) {
	contract, err := bindPhalaBTCLottery(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PhalaBTCLotteryCaller{contract: contract}, nil
}

// NewPhalaBTCLotteryTransactor creates a new write-only instance of PhalaBTCLottery, bound to a specific deployed contract.
func NewPhalaBTCLotteryTransactor(address common.Address, transactor bind.ContractTransactor) (*PhalaBTCLotteryTransactor, error) {
	contract, err := bindPhalaBTCLottery(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PhalaBTCLotteryTransactor{contract: contract}, nil
}

// NewPhalaBTCLotteryFilterer creates a new log filterer instance of PhalaBTCLottery, bound to a specific deployed contract.
func NewPhalaBTCLotteryFilterer(address common.Address, filterer bind.ContractFilterer) (*PhalaBTCLotteryFilterer, error) {
	contract, err := bindPhalaBTCLottery(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PhalaBTCLotteryFilterer{contract: contract}, nil
}

// bindPhalaBTCLottery binds a generic wrapper to an already deployed contract.
func bindPhalaBTCLottery(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PhalaBTCLotteryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PhalaBTCLottery *PhalaBTCLotteryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PhalaBTCLottery.Contract.PhalaBTCLotteryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PhalaBTCLottery *PhalaBTCLotteryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PhalaBTCLottery.Contract.PhalaBTCLotteryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PhalaBTCLottery *PhalaBTCLotteryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PhalaBTCLottery.Contract.PhalaBTCLotteryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PhalaBTCLottery *PhalaBTCLotteryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PhalaBTCLottery.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PhalaBTCLottery *PhalaBTCLotteryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PhalaBTCLottery.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PhalaBTCLottery *PhalaBTCLotteryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PhalaBTCLottery.Contract.contract.Transact(opts, method, params...)
}

// IsLotteryOpened is a free data retrieval call binding the contract method 0xf2db3797.
//
// Solidity: function isLotteryOpened(uint32 roundId, uint32 tokenId) view returns(bool)
func (_PhalaBTCLottery *PhalaBTCLotteryCaller) IsLotteryOpened(opts *bind.CallOpts, roundId uint32, tokenId uint32) (bool, error) {
	var out []interface{}
	err := _PhalaBTCLottery.contract.Call(opts, &out, "isLotteryOpened", roundId, tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLotteryOpened is a free data retrieval call binding the contract method 0xf2db3797.
//
// Solidity: function isLotteryOpened(uint32 roundId, uint32 tokenId) view returns(bool)
func (_PhalaBTCLottery *PhalaBTCLotterySession) IsLotteryOpened(roundId uint32, tokenId uint32) (bool, error) {
	return _PhalaBTCLottery.Contract.IsLotteryOpened(&_PhalaBTCLottery.CallOpts, roundId, tokenId)
}

// IsLotteryOpened is a free data retrieval call binding the contract method 0xf2db3797.
//
// Solidity: function isLotteryOpened(uint32 roundId, uint32 tokenId) view returns(bool)
func (_PhalaBTCLottery *PhalaBTCLotteryCallerSession) IsLotteryOpened(roundId uint32, tokenId uint32) (bool, error) {
	return _PhalaBTCLottery.Contract.IsLotteryOpened(&_PhalaBTCLottery.CallOpts, roundId, tokenId)
}

// OpenedBox is a free data retrieval call binding the contract method 0x12aad3d4.
//
// Solidity: function openedBox(uint32 , uint32 ) view returns(bool)
func (_PhalaBTCLottery *PhalaBTCLotteryCaller) OpenedBox(opts *bind.CallOpts, arg0 uint32, arg1 uint32) (bool, error) {
	var out []interface{}
	err := _PhalaBTCLottery.contract.Call(opts, &out, "openedBox", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OpenedBox is a free data retrieval call binding the contract method 0x12aad3d4.
//
// Solidity: function openedBox(uint32 , uint32 ) view returns(bool)
func (_PhalaBTCLottery *PhalaBTCLotterySession) OpenedBox(arg0 uint32, arg1 uint32) (bool, error) {
	return _PhalaBTCLottery.Contract.OpenedBox(&_PhalaBTCLottery.CallOpts, arg0, arg1)
}

// OpenedBox is a free data retrieval call binding the contract method 0x12aad3d4.
//
// Solidity: function openedBox(uint32 , uint32 ) view returns(bool)
func (_PhalaBTCLottery *PhalaBTCLotteryCallerSession) OpenedBox(arg0 uint32, arg1 uint32) (bool, error) {
	return _PhalaBTCLottery.Contract.OpenedBox(&_PhalaBTCLottery.CallOpts, arg0, arg1)
}

// StoredTx is a free data retrieval call binding the contract method 0xc066b75f.
//
// Solidity: function storedTx(uint32 , uint32 ) view returns(bool)
func (_PhalaBTCLottery *PhalaBTCLotteryCaller) StoredTx(opts *bind.CallOpts, arg0 uint32, arg1 uint32) (bool, error) {
	var out []interface{}
	err := _PhalaBTCLottery.contract.Call(opts, &out, "storedTx", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// StoredTx is a free data retrieval call binding the contract method 0xc066b75f.
//
// Solidity: function storedTx(uint32 , uint32 ) view returns(bool)
func (_PhalaBTCLottery *PhalaBTCLotterySession) StoredTx(arg0 uint32, arg1 uint32) (bool, error) {
	return _PhalaBTCLottery.Contract.StoredTx(&_PhalaBTCLottery.CallOpts, arg0, arg1)
}

// StoredTx is a free data retrieval call binding the contract method 0xc066b75f.
//
// Solidity: function storedTx(uint32 , uint32 ) view returns(bool)
func (_PhalaBTCLottery *PhalaBTCLotteryCallerSession) StoredTx(arg0 uint32, arg1 uint32) (bool, error) {
	return _PhalaBTCLottery.Contract.StoredTx(&_PhalaBTCLottery.CallOpts, arg0, arg1)
}

// TxStorage is a free data retrieval call binding the contract method 0x9a4ae777.
//
// Solidity: function txStorage(uint32 , uint32 ) view returns(bytes)
func (_PhalaBTCLottery *PhalaBTCLotteryCaller) TxStorage(opts *bind.CallOpts, arg0 uint32, arg1 uint32) ([]byte, error) {
	var out []interface{}
	err := _PhalaBTCLottery.contract.Call(opts, &out, "txStorage", arg0, arg1)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// TxStorage is a free data retrieval call binding the contract method 0x9a4ae777.
//
// Solidity: function txStorage(uint32 , uint32 ) view returns(bytes)
func (_PhalaBTCLottery *PhalaBTCLotterySession) TxStorage(arg0 uint32, arg1 uint32) ([]byte, error) {
	return _PhalaBTCLottery.Contract.TxStorage(&_PhalaBTCLottery.CallOpts, arg0, arg1)
}

// TxStorage is a free data retrieval call binding the contract method 0x9a4ae777.
//
// Solidity: function txStorage(uint32 , uint32 ) view returns(bytes)
func (_PhalaBTCLottery *PhalaBTCLotteryCallerSession) TxStorage(arg0 uint32, arg1 uint32) ([]byte, error) {
	return _PhalaBTCLottery.Contract.TxStorage(&_PhalaBTCLottery.CallOpts, arg0, arg1)
}

// DepositHandler is a paid mutator transaction binding the contract method 0xf4190f83.
//
// Solidity: function depositHandler(bytes data) returns()
func (_PhalaBTCLottery *PhalaBTCLotteryTransactor) DepositHandler(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _PhalaBTCLottery.contract.Transact(opts, "depositHandler", data)
}

// DepositHandler is a paid mutator transaction binding the contract method 0xf4190f83.
//
// Solidity: function depositHandler(bytes data) returns()
func (_PhalaBTCLottery *PhalaBTCLotterySession) DepositHandler(data []byte) (*types.Transaction, error) {
	return _PhalaBTCLottery.Contract.DepositHandler(&_PhalaBTCLottery.TransactOpts, data)
}

// DepositHandler is a paid mutator transaction binding the contract method 0xf4190f83.
//
// Solidity: function depositHandler(bytes data) returns()
func (_PhalaBTCLottery *PhalaBTCLotteryTransactorSession) DepositHandler(data []byte) (*types.Transaction, error) {
	return _PhalaBTCLottery.Contract.DepositHandler(&_PhalaBTCLottery.TransactOpts, data)
}

// RecordSignedBTCTx is a paid mutator transaction binding the contract method 0xa2b58507.
//
// Solidity: function recordSignedBTCTx(bytes data) returns()
func (_PhalaBTCLottery *PhalaBTCLotteryTransactor) RecordSignedBTCTx(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _PhalaBTCLottery.contract.Transact(opts, "recordSignedBTCTx", data)
}

// RecordSignedBTCTx is a paid mutator transaction binding the contract method 0xa2b58507.
//
// Solidity: function recordSignedBTCTx(bytes data) returns()
func (_PhalaBTCLottery *PhalaBTCLotterySession) RecordSignedBTCTx(data []byte) (*types.Transaction, error) {
	return _PhalaBTCLottery.Contract.RecordSignedBTCTx(&_PhalaBTCLottery.TransactOpts, data)
}

// RecordSignedBTCTx is a paid mutator transaction binding the contract method 0xa2b58507.
//
// Solidity: function recordSignedBTCTx(bytes data) returns()
func (_PhalaBTCLottery *PhalaBTCLotteryTransactorSession) RecordSignedBTCTx(data []byte) (*types.Transaction, error) {
	return _PhalaBTCLottery.Contract.RecordSignedBTCTx(&_PhalaBTCLottery.TransactOpts, data)
}

// PhalaBTCLotteryNewRoundIterator is returned from FilterNewRound and is used to iterate over the raw logs and unpacked data for NewRound events raised by the PhalaBTCLottery contract.
type PhalaBTCLotteryNewRoundIterator struct {
	Event *PhalaBTCLotteryNewRound // Event containing the contract specifics and raw log

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
func (it *PhalaBTCLotteryNewRoundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PhalaBTCLotteryNewRound)
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
		it.Event = new(PhalaBTCLotteryNewRound)
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
func (it *PhalaBTCLotteryNewRoundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PhalaBTCLotteryNewRoundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PhalaBTCLotteryNewRound represents a NewRound event raised by the PhalaBTCLottery contract.
type PhalaBTCLotteryNewRound struct {
	RoundId     uint32
	TotalCount  uint32
	WinnerCount uint32
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewRound is a free log retrieval operation binding the contract event 0xdb03f7d591b8e310b058b3c3c80e9415acee180707443368d8bae2d553831240.
//
// Solidity: event NewRound(uint32 roundId, uint32 totalCount, uint32 winnerCount)
func (_PhalaBTCLottery *PhalaBTCLotteryFilterer) FilterNewRound(opts *bind.FilterOpts) (*PhalaBTCLotteryNewRoundIterator, error) {

	logs, sub, err := _PhalaBTCLottery.contract.FilterLogs(opts, "NewRound")
	if err != nil {
		return nil, err
	}
	return &PhalaBTCLotteryNewRoundIterator{contract: _PhalaBTCLottery.contract, event: "NewRound", logs: logs, sub: sub}, nil
}

// WatchNewRound is a free log subscription operation binding the contract event 0xdb03f7d591b8e310b058b3c3c80e9415acee180707443368d8bae2d553831240.
//
// Solidity: event NewRound(uint32 roundId, uint32 totalCount, uint32 winnerCount)
func (_PhalaBTCLottery *PhalaBTCLotteryFilterer) WatchNewRound(opts *bind.WatchOpts, sink chan<- *PhalaBTCLotteryNewRound) (event.Subscription, error) {

	logs, sub, err := _PhalaBTCLottery.contract.WatchLogs(opts, "NewRound")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PhalaBTCLotteryNewRound)
				if err := _PhalaBTCLottery.contract.UnpackLog(event, "NewRound", log); err != nil {
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

// ParseNewRound is a log parse operation binding the contract event 0xdb03f7d591b8e310b058b3c3c80e9415acee180707443368d8bae2d553831240.
//
// Solidity: event NewRound(uint32 roundId, uint32 totalCount, uint32 winnerCount)
func (_PhalaBTCLottery *PhalaBTCLotteryFilterer) ParseNewRound(log types.Log) (*PhalaBTCLotteryNewRound, error) {
	event := new(PhalaBTCLotteryNewRound)
	if err := _PhalaBTCLottery.contract.UnpackLog(event, "NewRound", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PhalaBTCLotteryOpenLotteryIterator is returned from FilterOpenLottery and is used to iterate over the raw logs and unpacked data for OpenLottery events raised by the PhalaBTCLottery contract.
type PhalaBTCLotteryOpenLotteryIterator struct {
	Event *PhalaBTCLotteryOpenLottery // Event containing the contract specifics and raw log

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
func (it *PhalaBTCLotteryOpenLotteryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PhalaBTCLotteryOpenLottery)
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
		it.Event = new(PhalaBTCLotteryOpenLottery)
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
func (it *PhalaBTCLotteryOpenLotteryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PhalaBTCLotteryOpenLotteryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PhalaBTCLotteryOpenLottery represents a OpenLottery event raised by the PhalaBTCLottery contract.
type PhalaBTCLotteryOpenLottery struct {
	RoundId uint32
	TokenId uint32
	BtcAddr []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOpenLottery is a free log retrieval operation binding the contract event 0x66b4e6a794fe8bbff66bf6eeb41af302d4f7e2bde331cb703b4ded081cac3d77.
//
// Solidity: event OpenLottery(uint32 roundId, uint32 tokenId, bytes btcAddr)
func (_PhalaBTCLottery *PhalaBTCLotteryFilterer) FilterOpenLottery(opts *bind.FilterOpts) (*PhalaBTCLotteryOpenLotteryIterator, error) {

	logs, sub, err := _PhalaBTCLottery.contract.FilterLogs(opts, "OpenLottery")
	if err != nil {
		return nil, err
	}
	return &PhalaBTCLotteryOpenLotteryIterator{contract: _PhalaBTCLottery.contract, event: "OpenLottery", logs: logs, sub: sub}, nil
}

// WatchOpenLottery is a free log subscription operation binding the contract event 0x66b4e6a794fe8bbff66bf6eeb41af302d4f7e2bde331cb703b4ded081cac3d77.
//
// Solidity: event OpenLottery(uint32 roundId, uint32 tokenId, bytes btcAddr)
func (_PhalaBTCLottery *PhalaBTCLotteryFilterer) WatchOpenLottery(opts *bind.WatchOpts, sink chan<- *PhalaBTCLotteryOpenLottery) (event.Subscription, error) {

	logs, sub, err := _PhalaBTCLottery.contract.WatchLogs(opts, "OpenLottery")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PhalaBTCLotteryOpenLottery)
				if err := _PhalaBTCLottery.contract.UnpackLog(event, "OpenLottery", log); err != nil {
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

// ParseOpenLottery is a log parse operation binding the contract event 0x66b4e6a794fe8bbff66bf6eeb41af302d4f7e2bde331cb703b4ded081cac3d77.
//
// Solidity: event OpenLottery(uint32 roundId, uint32 tokenId, bytes btcAddr)
func (_PhalaBTCLottery *PhalaBTCLotteryFilterer) ParseOpenLottery(log types.Log) (*PhalaBTCLotteryOpenLottery, error) {
	event := new(PhalaBTCLotteryOpenLottery)
	if err := _PhalaBTCLottery.contract.UnpackLog(event, "OpenLottery", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PhalaBTCLotterySignedTxStoredIterator is returned from FilterSignedTxStored and is used to iterate over the raw logs and unpacked data for SignedTxStored events raised by the PhalaBTCLottery contract.
type PhalaBTCLotterySignedTxStoredIterator struct {
	Event *PhalaBTCLotterySignedTxStored // Event containing the contract specifics and raw log

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
func (it *PhalaBTCLotterySignedTxStoredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PhalaBTCLotterySignedTxStored)
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
		it.Event = new(PhalaBTCLotterySignedTxStored)
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
func (it *PhalaBTCLotterySignedTxStoredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PhalaBTCLotterySignedTxStoredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PhalaBTCLotterySignedTxStored represents a SignedTxStored event raised by the PhalaBTCLottery contract.
type PhalaBTCLotterySignedTxStored struct {
	RoundId  uint32
	TokenId  uint32
	SignedTx []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSignedTxStored is a free log retrieval operation binding the contract event 0xe7b2b0726b9d9cd1592e2621a382872738fc5bf81fd1d81f9c1f4f962615b8af.
//
// Solidity: event SignedTxStored(uint32 roundId, uint32 tokenId, bytes signedTx)
func (_PhalaBTCLottery *PhalaBTCLotteryFilterer) FilterSignedTxStored(opts *bind.FilterOpts) (*PhalaBTCLotterySignedTxStoredIterator, error) {

	logs, sub, err := _PhalaBTCLottery.contract.FilterLogs(opts, "SignedTxStored")
	if err != nil {
		return nil, err
	}
	return &PhalaBTCLotterySignedTxStoredIterator{contract: _PhalaBTCLottery.contract, event: "SignedTxStored", logs: logs, sub: sub}, nil
}

// WatchSignedTxStored is a free log subscription operation binding the contract event 0xe7b2b0726b9d9cd1592e2621a382872738fc5bf81fd1d81f9c1f4f962615b8af.
//
// Solidity: event SignedTxStored(uint32 roundId, uint32 tokenId, bytes signedTx)
func (_PhalaBTCLottery *PhalaBTCLotteryFilterer) WatchSignedTxStored(opts *bind.WatchOpts, sink chan<- *PhalaBTCLotterySignedTxStored) (event.Subscription, error) {

	logs, sub, err := _PhalaBTCLottery.contract.WatchLogs(opts, "SignedTxStored")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PhalaBTCLotterySignedTxStored)
				if err := _PhalaBTCLottery.contract.UnpackLog(event, "SignedTxStored", log); err != nil {
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

// ParseSignedTxStored is a log parse operation binding the contract event 0xe7b2b0726b9d9cd1592e2621a382872738fc5bf81fd1d81f9c1f4f962615b8af.
//
// Solidity: event SignedTxStored(uint32 roundId, uint32 tokenId, bytes signedTx)
func (_PhalaBTCLottery *PhalaBTCLotteryFilterer) ParseSignedTxStored(log types.Log) (*PhalaBTCLotterySignedTxStored, error) {
	event := new(PhalaBTCLotterySignedTxStored)
	if err := _PhalaBTCLottery.contract.UnpackLog(event, "SignedTxStored", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
