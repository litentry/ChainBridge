// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package GenericHandler

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

// GenericHandlerDepositRecord is an auto generated low-level Go binding around an user-defined struct.
type GenericHandlerDepositRecord struct {
	DestinationChainID uint8
	Depositer          common.Address
	ResourceID         [32]byte
	MetaData           []byte
}

// GenericHandlerABI is the input ABI used to generate the binding from.
const GenericHandlerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridgeAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"initialResourceIDs\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"initialContractAddresses\",\"type\":\"address[]\"},{\"internalType\":\"bytes4[]\",\"name\":\"initialDepositFunctionSignatures\",\"type\":\"bytes4[]\"},{\"internalType\":\"bytes4[]\",\"name\":\"initialExecuteFunctionSignatures\",\"type\":\"bytes4[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"len\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"name\":\"DebugDeposit\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"_bridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_contractAddressToDepositFunctionSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_contractAddressToExecuteFunctionSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_contractAddressToResourceID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_contractWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"_depositRecords\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"_destinationChainID\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_depositer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_metaData\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_resourceIDToContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"destId\",\"type\":\"uint8\"}],\"name\":\"getDepositRecord\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"_destinationChainID\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_depositer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_metaData\",\"type\":\"bytes\"}],\"internalType\":\"structGenericHandler.DepositRecord\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"depositFunctionSig\",\"type\":\"bytes4\"},{\"internalType\":\"bytes4\",\"name\":\"executeFunctionSig\",\"type\":\"bytes4\"}],\"name\":\"setResource\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"destinationChainID\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"depositer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// GenericHandlerBin is the compiled bytecode used for deploying new contracts.
var GenericHandlerBin = "0x60806040523480156200001157600080fd5b50604051620023283803806200232883398181016040528101906200003791906200056d565b82518451146200007e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040162000075906200078b565b60405180910390fd5b8151835114620000c5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620000bc90620007cf565b60405180910390fd5b80518251146200010c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200010390620007ad565b60405180910390fd5b846000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060005b8451811015620001cb57620001bd8582815181106200016a57fe5b60200260200101518583815181106200017f57fe5b60200260200101518584815181106200019457fe5b6020026020010151858581518110620001a957fe5b6020026020010151620001d760201b60201c565b80806001019150506200014f565b50505050505062000963565b826002600086815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555083600360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555081600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548163ffffffff021916908360e01c021790555080600560008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548163ffffffff021916908360e01c02179055506001600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050505050565b600081519050620003928162000915565b92915050565b600082601f830112620003aa57600080fd5b8151620003c1620003bb826200081f565b620007f1565b91508181835260208401935060208101905083856020840282011115620003e757600080fd5b60005b838110156200041b578162000400888262000381565b845260208401935060208301925050600181019050620003ea565b5050505092915050565b600082601f8301126200043757600080fd5b81516200044e620004488262000848565b620007f1565b915081818352602084019350602081019050838560208402820111156200047457600080fd5b60005b83811015620004a857816200048d88826200053f565b84526020840193506020830192505060018101905062000477565b5050505092915050565b600082601f830112620004c457600080fd5b8151620004db620004d58262000871565b620007f1565b915081818352602084019350602081019050838560208402820111156200050157600080fd5b60005b838110156200053557816200051a888262000556565b84526020840193506020830192505060018101905062000504565b5050505092915050565b60008151905062000550816200092f565b92915050565b600081519050620005678162000949565b92915050565b600080600080600060a086880312156200058657600080fd5b6000620005968882890162000381565b955050602086015167ffffffffffffffff811115620005b457600080fd5b620005c28882890162000425565b945050604086015167ffffffffffffffff811115620005e057600080fd5b620005ee8882890162000398565b935050606086015167ffffffffffffffff8111156200060c57600080fd5b6200061a88828901620004b2565b925050608086015167ffffffffffffffff8111156200063857600080fd5b6200064688828901620004b2565b9150509295509295909350565b600062000662603c836200089a565b91507f696e697469616c5265736f7572636549447320616e6420696e697469616c436f60008301527f6e7472616374416464726573736573206c656e206d69736d61746368000000006020830152604082019050919050565b6000620006ca603d836200089a565b91507f70726f7669646564206465706f73697420616e6420657865637574652066756e60008301527f6374696f6e207369676e617475726573206c656e206d69736d617463680000006020830152604082019050919050565b6000620007326040836200089a565b91507f70726f766964656420636f6e74726163742061646472657373657320616e642060008301527f66756e6374696f6e207369676e617475726573206c656e206d69736d617463686020830152604082019050919050565b60006020820190508181036000830152620007a68162000653565b9050919050565b60006020820190508181036000830152620007c881620006bb565b9050919050565b60006020820190508181036000830152620007ea8162000723565b9050919050565b6000604051905081810181811067ffffffffffffffff821117156200081557600080fd5b8060405250919050565b600067ffffffffffffffff8211156200083757600080fd5b602082029050602081019050919050565b600067ffffffffffffffff8211156200086057600080fd5b602082029050602081019050919050565b600067ffffffffffffffff8211156200088957600080fd5b602082029050602081019050919050565b600082825260208201905092915050565b6000620008b882620008f5565b9050919050565b6000819050919050565b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6200092081620008ab565b81146200092c57600080fd5b50565b6200093a81620008bf565b81146200094657600080fd5b50565b6200095481620008c9565b81146200096057600080fd5b50565b6119b580620009736000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c8063ba484c0911610071578063ba484c091461017b578063bba8185a146101ab578063c54c2a11146101c7578063cb624463146101f7578063e248cff214610227578063ec97d3b414610243576100a9565b8063318c136e146100ae57806338995da9146100cc5780634402027f146100e85780637f79bea81461011b578063a5c3a9851461014b575b600080fd5b6100b6610273565b6040516100c391906115f5565b60405180910390f35b6100e660048036038101906100e19190611212565b610297565b005b61010260048036038101906100fd9190611309565b6106e1565b604051610112949392919061174a565b60405180910390f35b61013560048036038101906101309190611105565b6107e3565b6040516101429190611610565b60405180910390f35b61016560048036038101906101609190611105565b610803565b6040516101729190611646565b60405180910390f35b610195600480360381019061019091906112cd565b610823565b6040516101a291906116e3565b60405180910390f35b6101c560048036038101906101c09190611157565b61099b565b005b6101e160048036038101906101dc919061112e565b6109b5565b6040516101ee91906115f5565b60405180910390f35b610211600480360381019061020c9190611105565b6109e8565b60405161021e9190611646565b60405180910390f35b610241600480360381019061023c91906111ba565b610a08565b005b61025d60048036038101906102589190611105565b610d0a565b60405161026a919061162b565b60405180910390f35b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b61029f610d22565b6000606083838101906102b291906112a4565b9150838360209084602001926102ca939291906117df565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505090506000600260008a815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600660008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff166103d3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103ca906116c3565b60405180910390fd5b6000600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460e01b9050600060e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916146105cb57606081846040516024016104819190611661565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505090507fbc9ef9c6f4013270ec48bca840ceff8550b59c280d0da01050b69fd97d59213085858360405161051393929190611705565b60405180910390a160008373ffffffffffffffffffffffffffffffffffffffff168260405161054291906115de565b6000604051808303816000865af19150503d806000811461057f576040519150601f19603f3d011682016040523d82523d6000602084013e610584565b606091505b50509050806105c8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105bf906116a3565b60405180910390fd5b50505b60405180608001604052808a60ff1681526020018873ffffffffffffffffffffffffffffffffffffffff1681526020018b815260200184815250600160008b60ff1660ff16815260200190815260200160002060008a67ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548160ff021916908360ff16021790555060208201518160000160016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506040820151816001015560608201518160020190805190602001906106d1929190610f5c565b5090505050505050505050505050565b6001602052816000526040600020602052806000526040600020600091509150508060000160009054906101000a900460ff16908060000160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806001015490806002018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156107d95780601f106107ae576101008083540402835291602001916107d9565b820191906000526020600020905b8154815290600101906020018083116107bc57829003601f168201915b5050505050905084565b60066020528060005260406000206000915054906101000a900460ff1681565b60056020528060005260406000206000915054906101000a900460e01b81565b61082b610fdc565b600160008360ff1660ff16815260200190815260200160002060008467ffffffffffffffff1667ffffffffffffffff1681526020019081526020016000206040518060800160405290816000820160009054906101000a900460ff1660ff1660ff1681526020016000820160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160018201548152602001600282018054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561098a5780601f1061095f5761010080835404028352916020019161098a565b820191906000526020600020905b81548152906001019060200180831161096d57829003601f168201915b505050505081525050905092915050565b6109a3610d22565b6109af84848484610db2565b50505050565b60026020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60046020528060005260406000206000915054906101000a900460e01b81565b610a10610d22565b600060608383810190610a2391906112a4565b915083836020908460200192610a3b939291906117df565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050905060006002600087815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600660008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610b44576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b3b906116c3565b60405180910390fd5b6000600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460e01b9050600060e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614610d015760608184604051602401610bf29190611661565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050905060008373ffffffffffffffffffffffffffffffffffffffff1682604051610c7891906115de565b6000604051808303816000865af19150503d8060008114610cb5576040519150601f19603f3d011682016040523d82523d6000602084013e610cba565b606091505b5050905080610cfe576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610cf5906116a3565b60405180910390fd5b50505b50505050505050565b60036020528060005260406000206000915090505481565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610db0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610da790611683565b60405180910390fd5b565b826002600086815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555083600360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555081600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548163ffffffff021916908360e01c021790555080600560008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548163ffffffff021916908360e01c02179055506001600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610f9d57805160ff1916838001178555610fcb565b82800160010185558215610fcb579182015b82811115610fca578251825591602001919060010190610faf565b5b509050610fd89190611020565b5090565b6040518060800160405280600060ff168152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160008019168152602001606081525090565b5b80821115611039576000816000905550600101611021565b5090565b60008135905061104c816118f5565b92915050565b6000813590506110618161190c565b92915050565b60008135905061107681611923565b92915050565b60008083601f84011261108e57600080fd5b8235905067ffffffffffffffff8111156110a757600080fd5b6020830191508360018202830111156110bf57600080fd5b9250929050565b6000813590506110d58161193a565b92915050565b6000813590506110ea81611951565b92915050565b6000813590506110ff81611968565b92915050565b60006020828403121561111757600080fd5b60006111258482850161103d565b91505092915050565b60006020828403121561114057600080fd5b600061114e84828501611052565b91505092915050565b6000806000806080858703121561116d57600080fd5b600061117b87828801611052565b945050602061118c8782880161103d565b935050604061119d87828801611067565b92505060606111ae87828801611067565b91505092959194509250565b6000806000604084860312156111cf57600080fd5b60006111dd86828701611052565b935050602084013567ffffffffffffffff8111156111fa57600080fd5b6112068682870161107c565b92509250509250925092565b60008060008060008060a0878903121561122b57600080fd5b600061123989828a01611052565b965050602061124a89828a016110f0565b955050604061125b89828a016110db565b945050606061126c89828a0161103d565b935050608087013567ffffffffffffffff81111561128957600080fd5b61129589828a0161107c565b92509250509295509295509295565b6000602082840312156112b657600080fd5b60006112c4848285016110c6565b91505092915050565b600080604083850312156112e057600080fd5b60006112ee858286016110db565b92505060206112ff858286016110f0565b9150509250929050565b6000806040838503121561131c57600080fd5b600061132a858286016110f0565b925050602061133b858286016110db565b9150509250929050565b61134e81611812565b82525050565b61135d81611812565b82525050565b61136c81611824565b82525050565b61137b81611830565b82525050565b61138a81611830565b82525050565b6113998161183a565b82525050565b60006113aa82611796565b6113b481856117a1565b93506113c48185602086016118b1565b6113cd816118e4565b840191505092915050565b60006113e382611796565b6113ed81856117b2565b93506113fd8185602086016118b1565b611406816118e4565b840191505092915050565b600061141c82611796565b61142681856117c3565b93506114368185602086016118b1565b80840191505092915050565b600061144f601e836117ce565b91507f73656e646572206d7573742062652062726964676520636f6e747261637400006000830152602082019050919050565b600061148f6026836117ce565b91507f64656c656761746563616c6c20746f20636f6e7472616374416464726573732060008301527f6661696c656400000000000000000000000000000000000000000000000000006020830152604082019050919050565b60006114f5602b836117ce565b91507f70726f766964656420636f6e747261637441646472657373206973206e6f742060008301527f77686974656c69737465640000000000000000000000000000000000000000006020830152604082019050919050565b600060808301600083015161156660008601826115c0565b5060208301516115796020860182611345565b50604083015161158c6040860182611372565b50606083015184820360608601526115a4828261139f565b9150508091505092915050565b6115ba81611886565b82525050565b6115c9816118a4565b82525050565b6115d8816118a4565b82525050565b60006115ea8284611411565b915081905092915050565b600060208201905061160a6000830184611354565b92915050565b60006020820190506116256000830184611363565b92915050565b60006020820190506116406000830184611381565b92915050565b600060208201905061165b6000830184611390565b92915050565b6000602082019050818103600083015261167b81846113d8565b905092915050565b6000602082019050818103600083015261169c81611442565b9050919050565b600060208201905081810360008301526116bc81611482565b9050919050565b600060208201905081810360008301526116dc816114e8565b9050919050565b600060208201905081810360008301526116fd818461154e565b905092915050565b600060608201905061171a60008301866115b1565b818103602083015261172c81856113d8565b9050818103604083015261174081846113d8565b9050949350505050565b600060808201905061175f60008301876115cf565b61176c6020830186611354565b6117796040830185611381565b818103606083015261178b81846113d8565b905095945050505050565b600081519050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b600080858511156117ef57600080fd5b838611156117fc57600080fd5b6001850283019150848603905094509492505050565b600061181d82611866565b9050919050565b60008115159050919050565b6000819050919050565b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600067ffffffffffffffff82169050919050565b600060ff82169050919050565b60005b838110156118cf5780820151818401526020810190506118b4565b838111156118de576000848401525b50505050565b6000601f19601f8301169050919050565b6118fe81611812565b811461190957600080fd5b50565b61191581611830565b811461192057600080fd5b50565b61192c8161183a565b811461193757600080fd5b50565b61194381611886565b811461194e57600080fd5b50565b61195a81611890565b811461196557600080fd5b50565b611971816118a4565b811461197c57600080fd5b5056fea264697066735822122028b1a9b602c034f01f4ba10ee38af1a50022eec88761c0a3685b5cbd67c5998264736f6c63430007000033"

// DeployGenericHandler deploys a new Ethereum contract, binding an instance of GenericHandler to it.
func DeployGenericHandler(auth *bind.TransactOpts, backend bind.ContractBackend, bridgeAddress common.Address, initialResourceIDs [][32]byte, initialContractAddresses []common.Address, initialDepositFunctionSignatures [][4]byte, initialExecuteFunctionSignatures [][4]byte) (common.Address, *types.Transaction, *GenericHandler, error) {
	parsed, err := abi.JSON(strings.NewReader(GenericHandlerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(GenericHandlerBin), backend, bridgeAddress, initialResourceIDs, initialContractAddresses, initialDepositFunctionSignatures, initialExecuteFunctionSignatures)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GenericHandler{GenericHandlerCaller: GenericHandlerCaller{contract: contract}, GenericHandlerTransactor: GenericHandlerTransactor{contract: contract}, GenericHandlerFilterer: GenericHandlerFilterer{contract: contract}}, nil
}

// GenericHandler is an auto generated Go binding around an Ethereum contract.
type GenericHandler struct {
	GenericHandlerCaller     // Read-only binding to the contract
	GenericHandlerTransactor // Write-only binding to the contract
	GenericHandlerFilterer   // Log filterer for contract events
}

// GenericHandlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type GenericHandlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GenericHandlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GenericHandlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GenericHandlerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GenericHandlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GenericHandlerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GenericHandlerSession struct {
	Contract     *GenericHandler   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GenericHandlerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GenericHandlerCallerSession struct {
	Contract *GenericHandlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// GenericHandlerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GenericHandlerTransactorSession struct {
	Contract     *GenericHandlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// GenericHandlerRaw is an auto generated low-level Go binding around an Ethereum contract.
type GenericHandlerRaw struct {
	Contract *GenericHandler // Generic contract binding to access the raw methods on
}

// GenericHandlerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GenericHandlerCallerRaw struct {
	Contract *GenericHandlerCaller // Generic read-only contract binding to access the raw methods on
}

// GenericHandlerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GenericHandlerTransactorRaw struct {
	Contract *GenericHandlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGenericHandler creates a new instance of GenericHandler, bound to a specific deployed contract.
func NewGenericHandler(address common.Address, backend bind.ContractBackend) (*GenericHandler, error) {
	contract, err := bindGenericHandler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GenericHandler{GenericHandlerCaller: GenericHandlerCaller{contract: contract}, GenericHandlerTransactor: GenericHandlerTransactor{contract: contract}, GenericHandlerFilterer: GenericHandlerFilterer{contract: contract}}, nil
}

// NewGenericHandlerCaller creates a new read-only instance of GenericHandler, bound to a specific deployed contract.
func NewGenericHandlerCaller(address common.Address, caller bind.ContractCaller) (*GenericHandlerCaller, error) {
	contract, err := bindGenericHandler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GenericHandlerCaller{contract: contract}, nil
}

// NewGenericHandlerTransactor creates a new write-only instance of GenericHandler, bound to a specific deployed contract.
func NewGenericHandlerTransactor(address common.Address, transactor bind.ContractTransactor) (*GenericHandlerTransactor, error) {
	contract, err := bindGenericHandler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GenericHandlerTransactor{contract: contract}, nil
}

// NewGenericHandlerFilterer creates a new log filterer instance of GenericHandler, bound to a specific deployed contract.
func NewGenericHandlerFilterer(address common.Address, filterer bind.ContractFilterer) (*GenericHandlerFilterer, error) {
	contract, err := bindGenericHandler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GenericHandlerFilterer{contract: contract}, nil
}

// bindGenericHandler binds a generic wrapper to an already deployed contract.
func bindGenericHandler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GenericHandlerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GenericHandler *GenericHandlerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GenericHandler.Contract.GenericHandlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GenericHandler *GenericHandlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GenericHandler.Contract.GenericHandlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GenericHandler *GenericHandlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GenericHandler.Contract.GenericHandlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GenericHandler *GenericHandlerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GenericHandler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GenericHandler *GenericHandlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GenericHandler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GenericHandler *GenericHandlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GenericHandler.Contract.contract.Transact(opts, method, params...)
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_GenericHandler *GenericHandlerCaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GenericHandler.contract.Call(opts, &out, "_bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_GenericHandler *GenericHandlerSession) BridgeAddress() (common.Address, error) {
	return _GenericHandler.Contract.BridgeAddress(&_GenericHandler.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_GenericHandler *GenericHandlerCallerSession) BridgeAddress() (common.Address, error) {
	return _GenericHandler.Contract.BridgeAddress(&_GenericHandler.CallOpts)
}

// ContractAddressToDepositFunctionSignature is a free data retrieval call binding the contract method 0xcb624463.
//
// Solidity: function _contractAddressToDepositFunctionSignature(address ) view returns(bytes4)
func (_GenericHandler *GenericHandlerCaller) ContractAddressToDepositFunctionSignature(opts *bind.CallOpts, arg0 common.Address) ([4]byte, error) {
	var out []interface{}
	err := _GenericHandler.contract.Call(opts, &out, "_contractAddressToDepositFunctionSignature", arg0)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// ContractAddressToDepositFunctionSignature is a free data retrieval call binding the contract method 0xcb624463.
//
// Solidity: function _contractAddressToDepositFunctionSignature(address ) view returns(bytes4)
func (_GenericHandler *GenericHandlerSession) ContractAddressToDepositFunctionSignature(arg0 common.Address) ([4]byte, error) {
	return _GenericHandler.Contract.ContractAddressToDepositFunctionSignature(&_GenericHandler.CallOpts, arg0)
}

// ContractAddressToDepositFunctionSignature is a free data retrieval call binding the contract method 0xcb624463.
//
// Solidity: function _contractAddressToDepositFunctionSignature(address ) view returns(bytes4)
func (_GenericHandler *GenericHandlerCallerSession) ContractAddressToDepositFunctionSignature(arg0 common.Address) ([4]byte, error) {
	return _GenericHandler.Contract.ContractAddressToDepositFunctionSignature(&_GenericHandler.CallOpts, arg0)
}

// ContractAddressToExecuteFunctionSignature is a free data retrieval call binding the contract method 0xa5c3a985.
//
// Solidity: function _contractAddressToExecuteFunctionSignature(address ) view returns(bytes4)
func (_GenericHandler *GenericHandlerCaller) ContractAddressToExecuteFunctionSignature(opts *bind.CallOpts, arg0 common.Address) ([4]byte, error) {
	var out []interface{}
	err := _GenericHandler.contract.Call(opts, &out, "_contractAddressToExecuteFunctionSignature", arg0)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// ContractAddressToExecuteFunctionSignature is a free data retrieval call binding the contract method 0xa5c3a985.
//
// Solidity: function _contractAddressToExecuteFunctionSignature(address ) view returns(bytes4)
func (_GenericHandler *GenericHandlerSession) ContractAddressToExecuteFunctionSignature(arg0 common.Address) ([4]byte, error) {
	return _GenericHandler.Contract.ContractAddressToExecuteFunctionSignature(&_GenericHandler.CallOpts, arg0)
}

// ContractAddressToExecuteFunctionSignature is a free data retrieval call binding the contract method 0xa5c3a985.
//
// Solidity: function _contractAddressToExecuteFunctionSignature(address ) view returns(bytes4)
func (_GenericHandler *GenericHandlerCallerSession) ContractAddressToExecuteFunctionSignature(arg0 common.Address) ([4]byte, error) {
	return _GenericHandler.Contract.ContractAddressToExecuteFunctionSignature(&_GenericHandler.CallOpts, arg0)
}

// ContractAddressToResourceID is a free data retrieval call binding the contract method 0xec97d3b4.
//
// Solidity: function _contractAddressToResourceID(address ) view returns(bytes32)
func (_GenericHandler *GenericHandlerCaller) ContractAddressToResourceID(opts *bind.CallOpts, arg0 common.Address) ([32]byte, error) {
	var out []interface{}
	err := _GenericHandler.contract.Call(opts, &out, "_contractAddressToResourceID", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ContractAddressToResourceID is a free data retrieval call binding the contract method 0xec97d3b4.
//
// Solidity: function _contractAddressToResourceID(address ) view returns(bytes32)
func (_GenericHandler *GenericHandlerSession) ContractAddressToResourceID(arg0 common.Address) ([32]byte, error) {
	return _GenericHandler.Contract.ContractAddressToResourceID(&_GenericHandler.CallOpts, arg0)
}

// ContractAddressToResourceID is a free data retrieval call binding the contract method 0xec97d3b4.
//
// Solidity: function _contractAddressToResourceID(address ) view returns(bytes32)
func (_GenericHandler *GenericHandlerCallerSession) ContractAddressToResourceID(arg0 common.Address) ([32]byte, error) {
	return _GenericHandler.Contract.ContractAddressToResourceID(&_GenericHandler.CallOpts, arg0)
}

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) view returns(bool)
func (_GenericHandler *GenericHandlerCaller) ContractWhitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _GenericHandler.contract.Call(opts, &out, "_contractWhitelist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) view returns(bool)
func (_GenericHandler *GenericHandlerSession) ContractWhitelist(arg0 common.Address) (bool, error) {
	return _GenericHandler.Contract.ContractWhitelist(&_GenericHandler.CallOpts, arg0)
}

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) view returns(bool)
func (_GenericHandler *GenericHandlerCallerSession) ContractWhitelist(arg0 common.Address) (bool, error) {
	return _GenericHandler.Contract.ContractWhitelist(&_GenericHandler.CallOpts, arg0)
}

// DepositRecords is a free data retrieval call binding the contract method 0x4402027f.
//
// Solidity: function _depositRecords(uint8 , uint64 ) view returns(uint8 _destinationChainID, address _depositer, bytes32 _resourceID, bytes _metaData)
func (_GenericHandler *GenericHandlerCaller) DepositRecords(opts *bind.CallOpts, arg0 uint8, arg1 uint64) (struct {
	DestinationChainID uint8
	Depositer          common.Address
	ResourceID         [32]byte
	MetaData           []byte
}, error) {
	var out []interface{}
	err := _GenericHandler.contract.Call(opts, &out, "_depositRecords", arg0, arg1)

	outstruct := new(struct {
		DestinationChainID uint8
		Depositer          common.Address
		ResourceID         [32]byte
		MetaData           []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.DestinationChainID = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.Depositer = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.ResourceID = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	outstruct.MetaData = *abi.ConvertType(out[3], new([]byte)).(*[]byte)

	return *outstruct, err

}

// DepositRecords is a free data retrieval call binding the contract method 0x4402027f.
//
// Solidity: function _depositRecords(uint8 , uint64 ) view returns(uint8 _destinationChainID, address _depositer, bytes32 _resourceID, bytes _metaData)
func (_GenericHandler *GenericHandlerSession) DepositRecords(arg0 uint8, arg1 uint64) (struct {
	DestinationChainID uint8
	Depositer          common.Address
	ResourceID         [32]byte
	MetaData           []byte
}, error) {
	return _GenericHandler.Contract.DepositRecords(&_GenericHandler.CallOpts, arg0, arg1)
}

// DepositRecords is a free data retrieval call binding the contract method 0x4402027f.
//
// Solidity: function _depositRecords(uint8 , uint64 ) view returns(uint8 _destinationChainID, address _depositer, bytes32 _resourceID, bytes _metaData)
func (_GenericHandler *GenericHandlerCallerSession) DepositRecords(arg0 uint8, arg1 uint64) (struct {
	DestinationChainID uint8
	Depositer          common.Address
	ResourceID         [32]byte
	MetaData           []byte
}, error) {
	return _GenericHandler.Contract.DepositRecords(&_GenericHandler.CallOpts, arg0, arg1)
}

// ResourceIDToContractAddress is a free data retrieval call binding the contract method 0xc54c2a11.
//
// Solidity: function _resourceIDToContractAddress(bytes32 ) view returns(address)
func (_GenericHandler *GenericHandlerCaller) ResourceIDToContractAddress(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _GenericHandler.contract.Call(opts, &out, "_resourceIDToContractAddress", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ResourceIDToContractAddress is a free data retrieval call binding the contract method 0xc54c2a11.
//
// Solidity: function _resourceIDToContractAddress(bytes32 ) view returns(address)
func (_GenericHandler *GenericHandlerSession) ResourceIDToContractAddress(arg0 [32]byte) (common.Address, error) {
	return _GenericHandler.Contract.ResourceIDToContractAddress(&_GenericHandler.CallOpts, arg0)
}

// ResourceIDToContractAddress is a free data retrieval call binding the contract method 0xc54c2a11.
//
// Solidity: function _resourceIDToContractAddress(bytes32 ) view returns(address)
func (_GenericHandler *GenericHandlerCallerSession) ResourceIDToContractAddress(arg0 [32]byte) (common.Address, error) {
	return _GenericHandler.Contract.ResourceIDToContractAddress(&_GenericHandler.CallOpts, arg0)
}

// GetDepositRecord is a free data retrieval call binding the contract method 0xba484c09.
//
// Solidity: function getDepositRecord(uint64 depositNonce, uint8 destId) view returns((uint8,address,bytes32,bytes))
func (_GenericHandler *GenericHandlerCaller) GetDepositRecord(opts *bind.CallOpts, depositNonce uint64, destId uint8) (GenericHandlerDepositRecord, error) {
	var out []interface{}
	err := _GenericHandler.contract.Call(opts, &out, "getDepositRecord", depositNonce, destId)

	if err != nil {
		return *new(GenericHandlerDepositRecord), err
	}

	out0 := *abi.ConvertType(out[0], new(GenericHandlerDepositRecord)).(*GenericHandlerDepositRecord)

	return out0, err

}

// GetDepositRecord is a free data retrieval call binding the contract method 0xba484c09.
//
// Solidity: function getDepositRecord(uint64 depositNonce, uint8 destId) view returns((uint8,address,bytes32,bytes))
func (_GenericHandler *GenericHandlerSession) GetDepositRecord(depositNonce uint64, destId uint8) (GenericHandlerDepositRecord, error) {
	return _GenericHandler.Contract.GetDepositRecord(&_GenericHandler.CallOpts, depositNonce, destId)
}

// GetDepositRecord is a free data retrieval call binding the contract method 0xba484c09.
//
// Solidity: function getDepositRecord(uint64 depositNonce, uint8 destId) view returns((uint8,address,bytes32,bytes))
func (_GenericHandler *GenericHandlerCallerSession) GetDepositRecord(depositNonce uint64, destId uint8) (GenericHandlerDepositRecord, error) {
	return _GenericHandler.Contract.GetDepositRecord(&_GenericHandler.CallOpts, depositNonce, destId)
}

// Deposit is a paid mutator transaction binding the contract method 0x38995da9.
//
// Solidity: function deposit(bytes32 resourceID, uint8 destinationChainID, uint64 depositNonce, address depositer, bytes data) returns()
func (_GenericHandler *GenericHandlerTransactor) Deposit(opts *bind.TransactOpts, resourceID [32]byte, destinationChainID uint8, depositNonce uint64, depositer common.Address, data []byte) (*types.Transaction, error) {
	return _GenericHandler.contract.Transact(opts, "deposit", resourceID, destinationChainID, depositNonce, depositer, data)
}

// Deposit is a paid mutator transaction binding the contract method 0x38995da9.
//
// Solidity: function deposit(bytes32 resourceID, uint8 destinationChainID, uint64 depositNonce, address depositer, bytes data) returns()
func (_GenericHandler *GenericHandlerSession) Deposit(resourceID [32]byte, destinationChainID uint8, depositNonce uint64, depositer common.Address, data []byte) (*types.Transaction, error) {
	return _GenericHandler.Contract.Deposit(&_GenericHandler.TransactOpts, resourceID, destinationChainID, depositNonce, depositer, data)
}

// Deposit is a paid mutator transaction binding the contract method 0x38995da9.
//
// Solidity: function deposit(bytes32 resourceID, uint8 destinationChainID, uint64 depositNonce, address depositer, bytes data) returns()
func (_GenericHandler *GenericHandlerTransactorSession) Deposit(resourceID [32]byte, destinationChainID uint8, depositNonce uint64, depositer common.Address, data []byte) (*types.Transaction, error) {
	return _GenericHandler.Contract.Deposit(&_GenericHandler.TransactOpts, resourceID, destinationChainID, depositNonce, depositer, data)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0xe248cff2.
//
// Solidity: function executeProposal(bytes32 resourceID, bytes data) returns()
func (_GenericHandler *GenericHandlerTransactor) ExecuteProposal(opts *bind.TransactOpts, resourceID [32]byte, data []byte) (*types.Transaction, error) {
	return _GenericHandler.contract.Transact(opts, "executeProposal", resourceID, data)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0xe248cff2.
//
// Solidity: function executeProposal(bytes32 resourceID, bytes data) returns()
func (_GenericHandler *GenericHandlerSession) ExecuteProposal(resourceID [32]byte, data []byte) (*types.Transaction, error) {
	return _GenericHandler.Contract.ExecuteProposal(&_GenericHandler.TransactOpts, resourceID, data)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0xe248cff2.
//
// Solidity: function executeProposal(bytes32 resourceID, bytes data) returns()
func (_GenericHandler *GenericHandlerTransactorSession) ExecuteProposal(resourceID [32]byte, data []byte) (*types.Transaction, error) {
	return _GenericHandler.Contract.ExecuteProposal(&_GenericHandler.TransactOpts, resourceID, data)
}

// SetResource is a paid mutator transaction binding the contract method 0xbba8185a.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress, bytes4 depositFunctionSig, bytes4 executeFunctionSig) returns()
func (_GenericHandler *GenericHandlerTransactor) SetResource(opts *bind.TransactOpts, resourceID [32]byte, contractAddress common.Address, depositFunctionSig [4]byte, executeFunctionSig [4]byte) (*types.Transaction, error) {
	return _GenericHandler.contract.Transact(opts, "setResource", resourceID, contractAddress, depositFunctionSig, executeFunctionSig)
}

// SetResource is a paid mutator transaction binding the contract method 0xbba8185a.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress, bytes4 depositFunctionSig, bytes4 executeFunctionSig) returns()
func (_GenericHandler *GenericHandlerSession) SetResource(resourceID [32]byte, contractAddress common.Address, depositFunctionSig [4]byte, executeFunctionSig [4]byte) (*types.Transaction, error) {
	return _GenericHandler.Contract.SetResource(&_GenericHandler.TransactOpts, resourceID, contractAddress, depositFunctionSig, executeFunctionSig)
}

// SetResource is a paid mutator transaction binding the contract method 0xbba8185a.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress, bytes4 depositFunctionSig, bytes4 executeFunctionSig) returns()
func (_GenericHandler *GenericHandlerTransactorSession) SetResource(resourceID [32]byte, contractAddress common.Address, depositFunctionSig [4]byte, executeFunctionSig [4]byte) (*types.Transaction, error) {
	return _GenericHandler.Contract.SetResource(&_GenericHandler.TransactOpts, resourceID, contractAddress, depositFunctionSig, executeFunctionSig)
}

// GenericHandlerDebugDepositIterator is returned from FilterDebugDeposit and is used to iterate over the raw logs and unpacked data for DebugDeposit events raised by the GenericHandler contract.
type GenericHandlerDebugDepositIterator struct {
	Event *GenericHandlerDebugDeposit // Event containing the contract specifics and raw log

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
func (it *GenericHandlerDebugDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GenericHandlerDebugDeposit)
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
		it.Event = new(GenericHandlerDebugDeposit)
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
func (it *GenericHandlerDebugDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GenericHandlerDebugDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GenericHandlerDebugDeposit represents a DebugDeposit event raised by the GenericHandler contract.
type GenericHandlerDebugDeposit struct {
	Len      *big.Int
	Metadata []byte
	CallData []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDebugDeposit is a free log retrieval operation binding the contract event 0xbc9ef9c6f4013270ec48bca840ceff8550b59c280d0da01050b69fd97d592130.
//
// Solidity: event DebugDeposit(uint256 len, bytes metadata, bytes callData)
func (_GenericHandler *GenericHandlerFilterer) FilterDebugDeposit(opts *bind.FilterOpts) (*GenericHandlerDebugDepositIterator, error) {

	logs, sub, err := _GenericHandler.contract.FilterLogs(opts, "DebugDeposit")
	if err != nil {
		return nil, err
	}
	return &GenericHandlerDebugDepositIterator{contract: _GenericHandler.contract, event: "DebugDeposit", logs: logs, sub: sub}, nil
}

// WatchDebugDeposit is a free log subscription operation binding the contract event 0xbc9ef9c6f4013270ec48bca840ceff8550b59c280d0da01050b69fd97d592130.
//
// Solidity: event DebugDeposit(uint256 len, bytes metadata, bytes callData)
func (_GenericHandler *GenericHandlerFilterer) WatchDebugDeposit(opts *bind.WatchOpts, sink chan<- *GenericHandlerDebugDeposit) (event.Subscription, error) {

	logs, sub, err := _GenericHandler.contract.WatchLogs(opts, "DebugDeposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GenericHandlerDebugDeposit)
				if err := _GenericHandler.contract.UnpackLog(event, "DebugDeposit", log); err != nil {
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

// ParseDebugDeposit is a log parse operation binding the contract event 0xbc9ef9c6f4013270ec48bca840ceff8550b59c280d0da01050b69fd97d592130.
//
// Solidity: event DebugDeposit(uint256 len, bytes metadata, bytes callData)
func (_GenericHandler *GenericHandlerFilterer) ParseDebugDeposit(log types.Log) (*GenericHandlerDebugDeposit, error) {
	event := new(GenericHandlerDebugDeposit)
	if err := _GenericHandler.contract.UnpackLog(event, "DebugDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
