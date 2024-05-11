// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package GenericHandler

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

// GenericHandlerDepositRecord is an auto generated low-level Go binding around an user-defined struct.
type GenericHandlerDepositRecord struct {
	DestinationChainID uint8
	Depositer          common.Address
	ResourceID         [32]byte
	MetaData           []byte
}

// GenericHandlerMetaData contains all meta data concerning the GenericHandler contract.
var GenericHandlerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridgeAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"initialResourceIDs\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"initialContractAddresses\",\"type\":\"address[]\"},{\"internalType\":\"bytes4[]\",\"name\":\"initialDepositFunctionSignatures\",\"type\":\"bytes4[]\"},{\"internalType\":\"bytes4[]\",\"name\":\"initialExecuteFunctionSignatures\",\"type\":\"bytes4[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"_bridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_contractAddressToDepositFunctionSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_contractAddressToExecuteFunctionSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_contractAddressToResourceID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_contractWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"_depositRecords\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"_destinationChainID\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_depositer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_metaData\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_resourceIDToContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"destId\",\"type\":\"uint8\"}],\"name\":\"getDepositRecord\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"_destinationChainID\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_depositer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_metaData\",\"type\":\"bytes\"}],\"internalType\":\"structGenericHandler.DepositRecord\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"depositFunctionSig\",\"type\":\"bytes4\"},{\"internalType\":\"bytes4\",\"name\":\"executeFunctionSig\",\"type\":\"bytes4\"}],\"name\":\"setResource\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"destinationChainID\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"depositer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040516200276c3803806200276c8339818101604052810190620000379190620007df565b82518451146200007e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040162000075906200096a565b60405180910390fd5b8151835114620000c5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620000bc9062000a02565b60405180910390fd5b80518251146200010c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620001039062000a9a565b60405180910390fd5b846000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060005b8451811015620001f757620001e185828151811062000173576200017262000abc565b5b602002602001015185838151811062000191576200019062000abc565b5b6020026020010151858481518110620001af57620001ae62000abc565b5b6020026020010151858581518110620001cd57620001cc62000abc565b5b60200260200101516200020360201b60201c565b8080620001ee9062000b24565b9150506200014f565b50505050505062000b71565b826002600086815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555083600360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555081600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548163ffffffff021916908360e01c021790555080600560008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548163ffffffff021916908360e01c02179055506001600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050505050565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620003ee82620003c1565b9050919050565b6200040081620003e1565b81146200040c57600080fd5b50565b6000815190506200042081620003f5565b92915050565b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b62000476826200042b565b810181811067ffffffffffffffff821117156200049857620004976200043c565b5b80604052505050565b6000620004ad620003ad565b9050620004bb82826200046b565b919050565b600067ffffffffffffffff821115620004de57620004dd6200043c565b5b602082029050602081019050919050565b600080fd5b6000819050919050565b6200050981620004f4565b81146200051557600080fd5b50565b6000815190506200052981620004fe565b92915050565b6000620005466200054084620004c0565b620004a1565b905080838252602082019050602084028301858111156200056c576200056b620004ef565b5b835b8181101562000599578062000584888262000518565b8452602084019350506020810190506200056e565b5050509392505050565b600082601f830112620005bb57620005ba62000426565b5b8151620005cd8482602086016200052f565b91505092915050565b600067ffffffffffffffff821115620005f457620005f36200043c565b5b602082029050602081019050919050565b60006200061c6200061684620005d6565b620004a1565b90508083825260208201905060208402830185811115620006425762000641620004ef565b5b835b818110156200066f57806200065a88826200040f565b84526020840193505060208101905062000644565b5050509392505050565b600082601f83011262000691576200069062000426565b5b8151620006a384826020860162000605565b91505092915050565b600067ffffffffffffffff821115620006ca57620006c96200043c565b5b602082029050602081019050919050565b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b6200071281620006db565b81146200071e57600080fd5b50565b600081519050620007328162000707565b92915050565b60006200074f6200074984620006ac565b620004a1565b90508083825260208201905060208402830185811115620007755762000774620004ef565b5b835b81811015620007a257806200078d888262000721565b84526020840193505060208101905062000777565b5050509392505050565b600082601f830112620007c457620007c362000426565b5b8151620007d684826020860162000738565b91505092915050565b600080600080600060a08688031215620007fe57620007fd620003b7565b5b60006200080e888289016200040f565b955050602086015167ffffffffffffffff811115620008325762000831620003bc565b5b6200084088828901620005a3565b945050604086015167ffffffffffffffff811115620008645762000863620003bc565b5b620008728882890162000679565b935050606086015167ffffffffffffffff811115620008965762000895620003bc565b5b620008a488828901620007ac565b925050608086015167ffffffffffffffff811115620008c857620008c7620003bc565b5b620008d688828901620007ac565b9150509295509295909350565b600082825260208201905092915050565b7f696e697469616c5265736f7572636549447320616e6420696e697469616c436f60008201527f6e7472616374416464726573736573206c656e206d69736d6174636800000000602082015250565b600062000952603c83620008e3565b91506200095f82620008f4565b604082019050919050565b60006020820190508181036000830152620009858162000943565b9050919050565b7f70726f766964656420636f6e74726163742061646472657373657320616e642060008201527f66756e6374696f6e207369676e617475726573206c656e206d69736d61746368602082015250565b6000620009ea604083620008e3565b9150620009f7826200098c565b604082019050919050565b6000602082019050818103600083015262000a1d81620009db565b9050919050565b7f70726f7669646564206465706f73697420616e6420657865637574652066756e60008201527f6374696f6e207369676e617475726573206c656e206d69736d61746368000000602082015250565b600062000a82603d83620008e3565b915062000a8f8262000a24565b604082019050919050565b6000602082019050818103600083015262000ab58162000a73565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000819050919050565b600062000b318262000b1a565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820362000b665762000b6562000aeb565b5b600182019050919050565b611beb8062000b816000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c8063ba484c0911610071578063ba484c091461017b578063bba8185a146101ab578063c54c2a11146101c7578063cb624463146101f7578063e248cff214610227578063ec97d3b414610243576100a9565b8063318c136e146100ae57806338995da9146100cc5780634402027f146100e85780637f79bea81461011b578063a5c3a9851461014b575b600080fd5b6100b6610273565b6040516100c39190610ef7565b60405180910390f35b6100e660048036038101906100e1919061105c565b610297565b005b61010260048036038101906100fd91906110f6565b61065b565b60405161011294939291906111e4565b60405180910390f35b61013560048036038101906101309190611230565b61074d565b6040516101429190611278565b60405180910390f35b61016560048036038101906101609190611230565b61076d565b60405161017291906112ce565b60405180910390f35b610195600480360381019061019091906112e9565b61078d565b6040516101a29190611403565b60405180910390f35b6101c560048036038101906101c09190611451565b6108f5565b005b6101e160048036038101906101dc91906114b8565b61090f565b6040516101ee9190610ef7565b60405180910390f35b610211600480360381019061020c9190611230565b610942565b60405161021e91906112ce565b60405180910390f35b610241600480360381019061023c91906114e5565b610962565b005b61025d60048036038101906102589190611230565b610c20565b60405161026a9190611545565b60405180910390f35b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b61029f610c38565b6000606083838101906102b29190611596565b915083836020908460206102c691906115f2565b926102d393929190611630565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505090506000600260008a815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600660008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff166103dc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103d3906116ee565b60405180910390fd5b6000600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460e01b9050600060e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161461054c576000818460405160200161048b92919061176b565b604051602081830303815290604052905060008373ffffffffffffffffffffffffffffffffffffffff16826040516104c39190611793565b6000604051808303816000865af19150503d8060008114610500576040519150601f19603f3d011682016040523d82523d6000602084013e610505565b606091505b5050905080610549576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105409061181c565b60405180910390fd5b50505b60405180608001604052808a60ff1681526020018873ffffffffffffffffffffffffffffffffffffffff1681526020018b815260200184815250600160008b60ff1660ff16815260200190815260200160002060008a67ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548160ff021916908360ff16021790555060208201518160000160016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160010155606082015181600201908161064b9190611a77565b5090505050505050505050505050565b6001602052816000526040600020602052806000526040600020600091509150508060000160009054906101000a900460ff16908060000160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154908060020180546106ca9061189a565b80601f01602080910402602001604051908101604052809291908181526020018280546106f69061189a565b80156107435780601f1061071857610100808354040283529160200191610743565b820191906000526020600020905b81548152906001019060200180831161072657829003601f168201915b5050505050905084565b60066020528060005260406000206000915054906101000a900460ff1681565b60056020528060005260406000206000915054906101000a900460e01b81565b610795610e72565b600160008360ff1660ff16815260200190815260200160002060008467ffffffffffffffff1667ffffffffffffffff1681526020019081526020016000206040518060800160405290816000820160009054906101000a900460ff1660ff1660ff1681526020016000820160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820154815260200160028201805461086b9061189a565b80601f01602080910402602001604051908101604052809291908181526020018280546108979061189a565b80156108e45780601f106108b9576101008083540402835291602001916108e4565b820191906000526020600020905b8154815290600101906020018083116108c757829003601f168201915b505050505081525050905092915050565b6108fd610c38565b61090984848484610cc8565b50505050565b60026020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60046020528060005260406000206000915054906101000a900460e01b81565b61096a610c38565b60006060838381019061097d9190611596565b9150838360209084602061099191906115f2565b9261099e93929190611630565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050905060006002600087815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600660008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610aa7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a9e906116ee565b60405180910390fd5b6000600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460e01b9050600060e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614610c175760008184604051602001610b5692919061176b565b604051602081830303815290604052905060008373ffffffffffffffffffffffffffffffffffffffff1682604051610b8e9190611793565b6000604051808303816000865af19150503d8060008114610bcb576040519150601f19603f3d011682016040523d82523d6000602084013e610bd0565b606091505b5050905080610c14576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c0b9061181c565b60405180910390fd5b50505b50505050505050565b60036020528060005260406000206000915090505481565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610cc6576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610cbd90611b95565b60405180910390fd5b565b826002600086815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555083600360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555081600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548163ffffffff021916908360e01c021790555080600560008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548163ffffffff021916908360e01c02179055506001600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050505050565b6040518060800160405280600060ff168152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160008019168152602001606081525090565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610ee182610eb6565b9050919050565b610ef181610ed6565b82525050565b6000602082019050610f0c6000830184610ee8565b92915050565b600080fd5b600080fd5b6000819050919050565b610f2f81610f1c565b8114610f3a57600080fd5b50565b600081359050610f4c81610f26565b92915050565b600060ff82169050919050565b610f6881610f52565b8114610f7357600080fd5b50565b600081359050610f8581610f5f565b92915050565b600067ffffffffffffffff82169050919050565b610fa881610f8b565b8114610fb357600080fd5b50565b600081359050610fc581610f9f565b92915050565b610fd481610ed6565b8114610fdf57600080fd5b50565b600081359050610ff181610fcb565b92915050565b600080fd5b600080fd5b600080fd5b60008083601f84011261101c5761101b610ff7565b5b8235905067ffffffffffffffff81111561103957611038610ffc565b5b60208301915083600182028301111561105557611054611001565b5b9250929050565b60008060008060008060a0878903121561107957611078610f12565b5b600061108789828a01610f3d565b965050602061109889828a01610f76565b95505060406110a989828a01610fb6565b94505060606110ba89828a01610fe2565b935050608087013567ffffffffffffffff8111156110db576110da610f17565b5b6110e789828a01611006565b92509250509295509295509295565b6000806040838503121561110d5761110c610f12565b5b600061111b85828601610f76565b925050602061112c85828601610fb6565b9150509250929050565b61113f81610f52565b82525050565b61114e81610f1c565b82525050565b600081519050919050565b600082825260208201905092915050565b60005b8381101561118e578082015181840152602081019050611173565b60008484015250505050565b6000601f19601f8301169050919050565b60006111b682611154565b6111c0818561115f565b93506111d0818560208601611170565b6111d98161119a565b840191505092915050565b60006080820190506111f96000830187611136565b6112066020830186610ee8565b6112136040830185611145565b818103606083015261122581846111ab565b905095945050505050565b60006020828403121561124657611245610f12565b5b600061125484828501610fe2565b91505092915050565b60008115159050919050565b6112728161125d565b82525050565b600060208201905061128d6000830184611269565b92915050565b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b6112c881611293565b82525050565b60006020820190506112e360008301846112bf565b92915050565b60008060408385031215611300576112ff610f12565b5b600061130e85828601610fb6565b925050602061131f85828601610f76565b9150509250929050565b61133281610f52565b82525050565b61134181610ed6565b82525050565b61135081610f1c565b82525050565b600082825260208201905092915050565b600061137282611154565b61137c8185611356565b935061138c818560208601611170565b6113958161119a565b840191505092915050565b60006080830160008301516113b86000860182611329565b5060208301516113cb6020860182611338565b5060408301516113de6040860182611347565b50606083015184820360608601526113f68282611367565b9150508091505092915050565b6000602082019050818103600083015261141d81846113a0565b905092915050565b61142e81611293565b811461143957600080fd5b50565b60008135905061144b81611425565b92915050565b6000806000806080858703121561146b5761146a610f12565b5b600061147987828801610f3d565b945050602061148a87828801610fe2565b935050604061149b8782880161143c565b92505060606114ac8782880161143c565b91505092959194509250565b6000602082840312156114ce576114cd610f12565b5b60006114dc84828501610f3d565b91505092915050565b6000806000604084860312156114fe576114fd610f12565b5b600061150c86828701610f3d565b935050602084013567ffffffffffffffff81111561152d5761152c610f17565b5b61153986828701611006565b92509250509250925092565b600060208201905061155a6000830184611145565b92915050565b6000819050919050565b61157381611560565b811461157e57600080fd5b50565b6000813590506115908161156a565b92915050565b6000602082840312156115ac576115ab610f12565b5b60006115ba84828501611581565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006115fd82611560565b915061160883611560565b92508282019050808211156116205761161f6115c3565b5b92915050565b600080fd5b600080fd5b6000808585111561164457611643611626565b5b838611156116555761165461162b565b5b6001850283019150848603905094509492505050565b600082825260208201905092915050565b7f70726f766964656420636f6e747261637441646472657373206973206e6f742060008201527f77686974656c6973746564000000000000000000000000000000000000000000602082015250565b60006116d8602b8361166b565b91506116e38261167c565b604082019050919050565b60006020820190508181036000830152611707816116cb565b9050919050565b6000819050919050565b61172961172482611293565b61170e565b82525050565b600081905092915050565b600061174582611154565b61174f818561172f565b935061175f818560208601611170565b80840191505092915050565b60006117778285611718565b600482019150611787828461173a565b91508190509392505050565b600061179f828461173a565b915081905092915050565b7f64656c656761746563616c6c20746f20636f6e7472616374416464726573732060008201527f6661696c65640000000000000000000000000000000000000000000000000000602082015250565b600061180660268361166b565b9150611811826117aa565b604082019050919050565b60006020820190508181036000830152611835816117f9565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806118b257607f821691505b6020821081036118c5576118c461186b565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b60006008830261192d7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826118f0565b61193786836118f0565b95508019841693508086168417925050509392505050565b6000819050919050565b600061197461196f61196a84611560565b61194f565b611560565b9050919050565b6000819050919050565b61198e83611959565b6119a261199a8261197b565b8484546118fd565b825550505050565b600090565b6119b76119aa565b6119c2818484611985565b505050565b5b818110156119e6576119db6000826119af565b6001810190506119c8565b5050565b601f821115611a2b576119fc816118cb565b611a05846118e0565b81016020851015611a14578190505b611a28611a20856118e0565b8301826119c7565b50505b505050565b600082821c905092915050565b6000611a4e60001984600802611a30565b1980831691505092915050565b6000611a678383611a3d565b9150826002028217905092915050565b611a8082611154565b67ffffffffffffffff811115611a9957611a9861183c565b5b611aa3825461189a565b611aae8282856119ea565b600060209050601f831160018114611ae15760008415611acf578287015190505b611ad98582611a5b565b865550611b41565b601f198416611aef866118cb565b60005b82811015611b1757848901518255600182019150602085019450602081019050611af2565b86831015611b345784890151611b30601f891682611a3d565b8355505b6001600288020188555050505b505050505050565b7f73656e646572206d7573742062652062726964676520636f6e74726163740000600082015250565b6000611b7f601e8361166b565b9150611b8a82611b49565b602082019050919050565b60006020820190508181036000830152611bae81611b72565b905091905056fea26469706673582212203519853128bc339fa1a4cd054a886eefb4c2d2ec17909685ae890e6d0c19cb5964736f6c63430008130033",
}

// GenericHandlerABI is the input ABI used to generate the binding from.
// Deprecated: Use GenericHandlerMetaData.ABI instead.
var GenericHandlerABI = GenericHandlerMetaData.ABI

// GenericHandlerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GenericHandlerMetaData.Bin instead.
var GenericHandlerBin = GenericHandlerMetaData.Bin

// DeployGenericHandler deploys a new Ethereum contract, binding an instance of GenericHandler to it.
func DeployGenericHandler(auth *bind.TransactOpts, backend bind.ContractBackend, bridgeAddress common.Address, initialResourceIDs [][32]byte, initialContractAddresses []common.Address, initialDepositFunctionSignatures [][4]byte, initialExecuteFunctionSignatures [][4]byte) (common.Address, *types.Transaction, *GenericHandler, error) {
	parsed, err := GenericHandlerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GenericHandlerBin), backend, bridgeAddress, initialResourceIDs, initialContractAddresses, initialDepositFunctionSignatures, initialExecuteFunctionSignatures)
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
	parsed, err := GenericHandlerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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
