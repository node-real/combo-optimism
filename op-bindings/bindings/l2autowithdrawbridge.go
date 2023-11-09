// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// L2AutoWithdrawBridgeMetaData contains all meta data concerning the L2AutoWithdrawBridge contract.
var L2AutoWithdrawBridgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_minWithdrawalAmount\",\"type\":\"uint256\"},{\"internalType\":\"enumFeeVault.WithdrawalNetwork\",\"name\":\"_withdrawalNetwork\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"AutoWithdrawTo\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"Withdrawal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumFeeVault.WithdrawalNetwork\",\"name\":\"withdrawalNetwork\",\"type\":\"uint8\"}],\"name\":\"Withdrawal\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_DELEGATION_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"L2_STANDARD_BRIDGE\",\"outputs\":[{\"internalType\":\"contractL2StandardBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"L2_STANDARD_BRIDGE_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_WITHDRAWAL_AMOUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RECIPIENT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WITHDRAWAL_NETWORK\",\"outputs\":[{\"internalType\":\"enumFeeVault.WithdrawalNetwork\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_minGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"autoWithdrawTo\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delegation_fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDelegateFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_delegateFee\",\"type\":\"uint256\"}],\"name\":\"setDelegateFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalProcessed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x610140604052606680546001600160a01b03191673420000000000000000000000000000000000001017905566038d7ea4c680006067553480156200004357600080fd5b5060405162001a7f38038062001a7f83398101604081905262000066916200043e565b6001600160a01b03831660a0526080829052600160026000858585808681111562000095576200009562000497565b60c0816001811115620000ac57620000ac62000497565b905250505060e093909352610100919091526101205250620000ce84620000d8565b50505050620004ad565b600054610100900460ff1615808015620000f95750600054600160ff909116105b8062000129575062000116306200021660201b62000c541760201c565b15801562000129575060005460ff166001145b620001925760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b6000805460ff191660011790558015620001b6576000805461ff0019166101001790555b620001c062000225565b620001cb826200028d565b801562000212576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050565b6001600160a01b03163b151590565b600054610100900460ff16620002815760405162461bcd60e51b815260206004820152602b602482015260008051602062001a5f83398151915260448201526a6e697469616c697a696e6760a81b606482015260840162000189565b6200028b6200030c565b565b6200029762000373565b6001600160a01b038116620002fe5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840162000189565b6200030981620003cf565b50565b600054610100900460ff16620003685760405162461bcd60e51b815260206004820152602b602482015260008051602062001a5f83398151915260448201526a6e697469616c697a696e6760a81b606482015260840162000189565b6200028b33620003cf565b6033546001600160a01b031633146200028b5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640162000189565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b80516001600160a01b03811681146200043957600080fd5b919050565b600080600080608085870312156200045557600080fd5b620004608562000421565b9350620004706020860162000421565b9250604085015191506060850151600281106200048c57600080fd5b939692955090935050565b634e487b7160e01b600052602160045260246000fd5b60805160a05160c05160e05161010051610120516115316200052e60003960006107df015260006107b60152600061078d0152600081816103150152818161054901526105840152600081816101300152818161049801528181610527015281816105bd0152610724015260008181610356015261039a01526115316000f3fe6080604052600436106101125760003560e01c806384411d65116100a5578063a4d3c6f511610074578063d0e12f9011610059578063d0e12f9014610303578063d3e5792b14610344578063f2fde38b1461037857600080fd5b8063a4d3c6f5146102d0578063c4d66de8146102e357600080fd5b806384411d65146102595780638da5cb5b1461026f5780639931c34c1461029a57806399f979c5146102ba57600080fd5b806354fd4d50116100e157806354fd4d50146101e8578063715018a61461020a57806372b8a7161461021f578063840d60b21461023e57600080fd5b80630d9019e11461011e57806321d127631461017c5780632cb7cb06146101a95780633ccfd60b146101d157600080fd5b3661011957005b600080fd5b34801561012a57600080fd5b506101527f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b34801561018857600080fd5b506066546101529073ffffffffffffffffffffffffffffffffffffffff1681565b3480156101b557600080fd5b5061015273420000000000000000000000000000000000001081565b3480156101dd57600080fd5b506101e6610398565b005b3480156101f457600080fd5b506101fd610786565b6040516101739190611059565b34801561021657600080fd5b506101e6610829565b34801561022b57600080fd5b506067545b604051908152602001610173565b34801561024a57600080fd5b5061023066038d7ea4c6800081565b34801561026557600080fd5b5061023060655481565b34801561027b57600080fd5b5060335473ffffffffffffffffffffffffffffffffffffffff16610152565b3480156102a657600080fd5b506101e66102b5366004611073565b61083d565b3480156102c657600080fd5b5061023060675481565b6101e66102de3660046110b5565b61084a565b3480156102ef57600080fd5b506101e66102fe36600461116d565b610a05565b34801561030f57600080fd5b506103377f000000000000000000000000000000000000000000000000000000000000000081565b60405161017391906111f2565b34801561035057600080fd5b506102307f000000000000000000000000000000000000000000000000000000000000000081565b34801561038457600080fd5b506101e661039336600461116d565b610ba0565b7f0000000000000000000000000000000000000000000000000000000000000000471015610473576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604a60248201527f4665655661756c743a207769746864726177616c20616d6f756e74206d75737460448201527f2062652067726561746572207468616e206d696e696d756d207769746864726160648201527f77616c20616d6f756e7400000000000000000000000000000000000000000000608482015260a4015b60405180910390fd5b6000479050806065600082825461048a9190611235565b9091555050604080518281527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166020820152338183015290517fc8a211cc64b6ed1b50595a9fcb1932b6d1e5a6e8ef15b60e5b1f988ea9086bba9181900360600190a17f38e04cbeb8c10f8f568618aa75be0f10b6729b8b4237743b4de20cbcde2839ee817f0000000000000000000000000000000000000000000000000000000000000000337f0000000000000000000000000000000000000000000000000000000000000000604051610578949392919061124d565b60405180910390a160017f000000000000000000000000000000000000000000000000000000000000000060018111156105b4576105b4611188565b036106cd5760007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168260405160006040518083038185875af1925050503d8060008114610633576040519150601f19603f3d011682016040523d82523d6000602084013e610638565b606091505b50509050806106c9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603060248201527f4665655661756c743a206661696c656420746f2073656e642045544820746f2060448201527f4c322066656520726563697069656e7400000000000000000000000000000000606482015260840161046a565b5050565b604080516020810182526000815290517fe11013dd0000000000000000000000000000000000000000000000000000000081527342000000000000000000000000000000000000109163e11013dd918491610750917f0000000000000000000000000000000000000000000000000000000000000000916188b89160040161128e565b6000604051808303818588803b15801561076957600080fd5b505af115801561077d573d6000803e3d6000fd5b50505050505b50565b60606107b17f0000000000000000000000000000000000000000000000000000000000000000610c70565b6107da7f0000000000000000000000000000000000000000000000000000000000000000610c70565b6108037f0000000000000000000000000000000000000000000000000000000000000000610c70565b604051602001610815939291906112c9565b604051602081830303815290604052905090565b610831610dad565b61083b6000610e2e565b565b610845610dad565b606755565b836067546108589190611235565b34146108e6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603360248201527f6d73672e76616c756520646f6573206e6f7420657175616c20746f2064656c6560448201527f676174696f6e5f666565202b20616d6f756e7400000000000000000000000000606482015260840161046a565b3373ffffffffffffffffffffffffffffffffffffffff167ff3f58d6a31b0011b592ca174b425f161b372632c23c93791414c2a29ee3bdf07868685856040516109329493929190611388565b60405180910390a2606754604051309180156108fc02916000818181858888f19350505050158015610968573d6000803e3d6000fd5b506066546040517fa3a7954800000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091169063a3a795489086906109cb908a908a9084908a908a908a906004016113c8565b6000604051808303818588803b1580156109e457600080fd5b505af11580156109f8573d6000803e3d6000fd5b5050505050505050505050565b600054610100900460ff1615808015610a255750600054600160ff909116105b80610a3f5750303b158015610a3f575060005460ff166001145b610acb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a6564000000000000000000000000000000000000606482015260840161046a565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790558015610b2957600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b610b31610ea5565b610b3a82610ba0565b80156106c957600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15050565b610ba8610dad565b73ffffffffffffffffffffffffffffffffffffffff8116610c4b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f6464726573730000000000000000000000000000000000000000000000000000606482015260840161046a565b61078381610e2e565b73ffffffffffffffffffffffffffffffffffffffff163b151590565b606081600003610cb357505060408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015290565b8160005b8115610cdd5780610cc781611420565b9150610cd69050600a83611487565b9150610cb7565b60008167ffffffffffffffff811115610cf857610cf861149b565b6040519080825280601f01601f191660200182016040528015610d22576020820181803683370190505b5090505b8415610da557610d376001836114ca565b9150610d44600a866114e1565b610d4f906030611235565b60f81b818381518110610d6457610d646114f5565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350610d9e600a86611487565b9450610d26565b949350505050565b60335473ffffffffffffffffffffffffffffffffffffffff16331461083b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161046a565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600054610100900460ff16610f3c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161046a565b61083b600054610100900460ff16610fd6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161046a565b61083b33610e2e565b60005b83811015610ffa578181015183820152602001610fe2565b83811115611009576000848401525b50505050565b60008151808452611027816020860160208601610fdf565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b60208152600061106c602083018461100f565b9392505050565b60006020828403121561108557600080fd5b5035919050565b803573ffffffffffffffffffffffffffffffffffffffff811681146110b057600080fd5b919050565b60008060008060008060a087890312156110ce57600080fd5b6110d78761108c565b95506110e56020880161108c565b945060408701359350606087013563ffffffff8116811461110557600080fd5b9250608087013567ffffffffffffffff8082111561112257600080fd5b818901915089601f83011261113657600080fd5b81358181111561114557600080fd5b8a602082850101111561115757600080fd5b6020830194508093505050509295509295509295565b60006020828403121561117f57600080fd5b61106c8261108c565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b600281106111ee577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b9052565b6020810161120082846111b7565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000821982111561124857611248611206565b500190565b84815273ffffffffffffffffffffffffffffffffffffffff8481166020830152831660408201526080810161128560608301846111b7565b95945050505050565b73ffffffffffffffffffffffffffffffffffffffff8416815263ffffffff83166020820152606060408201526000611285606083018461100f565b600084516112db818460208901610fdf565b80830190507f2e000000000000000000000000000000000000000000000000000000000000008082528551611317816001850160208a01610fdf565b60019201918201528351611332816002840160208801610fdf565b0160020195945050505050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b73ffffffffffffffffffffffffffffffffffffffff851681528360208201526060604082015260006113be60608301848661133f565b9695505050505050565b600073ffffffffffffffffffffffffffffffffffffffff808916835280881660208401525085604083015263ffffffff8516606083015260a0608083015261141460a08301848661133f565b98975050505050505050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361145157611451611206565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b60008261149657611496611458565b500490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000828210156114dc576114dc611206565b500390565b6000826114f0576114f0611458565b500690565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fdfea164736f6c634300080f000a496e697469616c697a61626c653a20636f6e7472616374206973206e6f742069",
}

// L2AutoWithdrawBridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use L2AutoWithdrawBridgeMetaData.ABI instead.
var L2AutoWithdrawBridgeABI = L2AutoWithdrawBridgeMetaData.ABI

// L2AutoWithdrawBridgeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use L2AutoWithdrawBridgeMetaData.Bin instead.
var L2AutoWithdrawBridgeBin = L2AutoWithdrawBridgeMetaData.Bin

// DeployL2AutoWithdrawBridge deploys a new Ethereum contract, binding an instance of L2AutoWithdrawBridge to it.
func DeployL2AutoWithdrawBridge(auth *bind.TransactOpts, backend bind.ContractBackend, _owner common.Address, _recipient common.Address, _minWithdrawalAmount *big.Int, _withdrawalNetwork uint8) (common.Address, *types.Transaction, *L2AutoWithdrawBridge, error) {
	parsed, err := L2AutoWithdrawBridgeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(L2AutoWithdrawBridgeBin), backend, _owner, _recipient, _minWithdrawalAmount, _withdrawalNetwork)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &L2AutoWithdrawBridge{L2AutoWithdrawBridgeCaller: L2AutoWithdrawBridgeCaller{contract: contract}, L2AutoWithdrawBridgeTransactor: L2AutoWithdrawBridgeTransactor{contract: contract}, L2AutoWithdrawBridgeFilterer: L2AutoWithdrawBridgeFilterer{contract: contract}}, nil
}

// L2AutoWithdrawBridge is an auto generated Go binding around an Ethereum contract.
type L2AutoWithdrawBridge struct {
	L2AutoWithdrawBridgeCaller     // Read-only binding to the contract
	L2AutoWithdrawBridgeTransactor // Write-only binding to the contract
	L2AutoWithdrawBridgeFilterer   // Log filterer for contract events
}

// L2AutoWithdrawBridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type L2AutoWithdrawBridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2AutoWithdrawBridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L2AutoWithdrawBridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2AutoWithdrawBridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L2AutoWithdrawBridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2AutoWithdrawBridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L2AutoWithdrawBridgeSession struct {
	Contract     *L2AutoWithdrawBridge // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// L2AutoWithdrawBridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L2AutoWithdrawBridgeCallerSession struct {
	Contract *L2AutoWithdrawBridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// L2AutoWithdrawBridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L2AutoWithdrawBridgeTransactorSession struct {
	Contract     *L2AutoWithdrawBridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// L2AutoWithdrawBridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type L2AutoWithdrawBridgeRaw struct {
	Contract *L2AutoWithdrawBridge // Generic contract binding to access the raw methods on
}

// L2AutoWithdrawBridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L2AutoWithdrawBridgeCallerRaw struct {
	Contract *L2AutoWithdrawBridgeCaller // Generic read-only contract binding to access the raw methods on
}

// L2AutoWithdrawBridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L2AutoWithdrawBridgeTransactorRaw struct {
	Contract *L2AutoWithdrawBridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL2AutoWithdrawBridge creates a new instance of L2AutoWithdrawBridge, bound to a specific deployed contract.
func NewL2AutoWithdrawBridge(address common.Address, backend bind.ContractBackend) (*L2AutoWithdrawBridge, error) {
	contract, err := bindL2AutoWithdrawBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L2AutoWithdrawBridge{L2AutoWithdrawBridgeCaller: L2AutoWithdrawBridgeCaller{contract: contract}, L2AutoWithdrawBridgeTransactor: L2AutoWithdrawBridgeTransactor{contract: contract}, L2AutoWithdrawBridgeFilterer: L2AutoWithdrawBridgeFilterer{contract: contract}}, nil
}

// NewL2AutoWithdrawBridgeCaller creates a new read-only instance of L2AutoWithdrawBridge, bound to a specific deployed contract.
func NewL2AutoWithdrawBridgeCaller(address common.Address, caller bind.ContractCaller) (*L2AutoWithdrawBridgeCaller, error) {
	contract, err := bindL2AutoWithdrawBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L2AutoWithdrawBridgeCaller{contract: contract}, nil
}

// NewL2AutoWithdrawBridgeTransactor creates a new write-only instance of L2AutoWithdrawBridge, bound to a specific deployed contract.
func NewL2AutoWithdrawBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*L2AutoWithdrawBridgeTransactor, error) {
	contract, err := bindL2AutoWithdrawBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L2AutoWithdrawBridgeTransactor{contract: contract}, nil
}

// NewL2AutoWithdrawBridgeFilterer creates a new log filterer instance of L2AutoWithdrawBridge, bound to a specific deployed contract.
func NewL2AutoWithdrawBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*L2AutoWithdrawBridgeFilterer, error) {
	contract, err := bindL2AutoWithdrawBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L2AutoWithdrawBridgeFilterer{contract: contract}, nil
}

// bindL2AutoWithdrawBridge binds a generic wrapper to an already deployed contract.
func bindL2AutoWithdrawBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := L2AutoWithdrawBridgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2AutoWithdrawBridge *L2AutoWithdrawBridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2AutoWithdrawBridge.Contract.L2AutoWithdrawBridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2AutoWithdrawBridge *L2AutoWithdrawBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2AutoWithdrawBridge.Contract.L2AutoWithdrawBridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2AutoWithdrawBridge *L2AutoWithdrawBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2AutoWithdrawBridge.Contract.L2AutoWithdrawBridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2AutoWithdrawBridge *L2AutoWithdrawBridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2AutoWithdrawBridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2AutoWithdrawBridge *L2AutoWithdrawBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2AutoWithdrawBridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2AutoWithdrawBridge *L2AutoWithdrawBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2AutoWithdrawBridge.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTDELEGATIONFEE is a free data retrieval call binding the contract method 0x840d60b2.
//
// Solidity: function DEFAULT_DELEGATION_FEE() view returns(uint256)
func (_L2AutoWithdrawBridge *L2AutoWithdrawBridgeCaller) DEFAULTDELEGATIONFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L2AutoWithdrawBridge.contract.Call(opts, &out, "DEFAULT_DELEGATION_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DEFAULTDELEGATIONFEE is a free data retrieval call binding the contract method 0x840d60b2.
//
// Solidity: function DEFAULT_DELEGATION_FEE() view returns(uint256)
func (_L2AutoWithdrawBridge *L2AutoWithdrawBridgeSession) DEFAULTDELEGATIONFEE() (*big.Int, error) {
	return _L2AutoWithdrawBridge.Contract.DEFAULTDELEGATIONFEE(&_L2AutoWithdrawBridge.CallOpts)
}

// DEFAULTDELEGATIONFEE is a free data retrieval call binding the contract method 0x840d60b2.
//
// Solidity: function DEFAULT_DELEGATION_FEE() view returns(uint256)
func (_L2AutoWithdrawBridge *L2AutoWithdrawBridgeCallerSession) DEFAULTDELEGATIONFEE() (*big.Int, error) {
	return _L2AutoWithdrawBridge.Contract.DEFAULTDELEGATIONFEE(&_L2AutoWithdrawBridge.CallOpts)
}

// L2STANDARDBRIDGE is a free data retrieval call binding the contract method 0x21d12763.
//
// Solidity: function L2_STANDARD_BRIDGE() view returns(address)
func (_L2AutoWithdrawBridge *L2AutoWithdrawBridgeCaller) L2STANDARDBRIDGE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2AutoWithdrawBridge.contract.Call(opts, &out, "L2_STANDARD_BRIDGE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2STANDARDBRIDGE is a free data retrieval call binding the contract method 0x21d12763.
//
// Solidity: function L2_STANDARD_BRIDGE() view returns(address)
func (_L2AutoWithdrawBridge *L2AutoWithdrawBridgeSession) L2STANDARDBRIDGE() (common.Address, error) {
	return _L2AutoWithdrawBridge.Contract.L2STANDARDBRIDGE(&_L2AutoWithdrawBridge.CallOpts)
}

// L2STANDARDBRIDGE is a free data retrieval call binding the contract method 0x21d12763.
//
// Solidity: function L2_STANDARD_BRIDGE() view returns(address)
func (_L2AutoWithdrawBridge *L2AutoWithdrawBridgeCallerSession) L2STANDARDBRIDGE() (common.Address, error) {
	return _L2AutoWithdrawBridge.Contract.L2STANDARDBRIDGE(&_L2AutoWithdrawBridge.CallOpts)
}

// L2STANDARDBRIDGEADDRESS is a free data retrieval call binding the contract method 0x2cb7cb06.
//
// Solidity: function L2_STANDARD_BRIDGE_ADDRESS() view returns(address)
func (_L2AutoWithdrawBridge *L2AutoWithdrawBridgeCaller) L2STANDARDBRIDGEADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2AutoWithdrawBridge.contract.Call(opts, &out, "L2_STANDARD_BRIDGE_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2STANDARDBRIDGEADDRESS is a free data retrieval call binding the contract method 0x2cb7cb06.
//
// Solidity: function L2_STANDARD_BRIDGE_ADDRESS() view returns(address)
func (_L2AutoWithdrawBridge *L2AutoWithdrawBridgeSession) L2STANDARDBRIDGEADDRESS() (common.Address, error) {
	return _L2AutoWithdrawBridge.Contract.L2STANDARDBRIDGEADDRESS(&_L2AutoWithdrawBridge.CallOpts)
}

// L2STANDARDBRIDGEADDRESS is a free data retrieval call binding the contract method 0x2cb7cb06.
//
// Solidity: function L2_STANDARD_BRIDGE_ADDRESS() view returns(address)
func (_L2AutoWithdrawBridge *L2AutoWithdrawBridgeCallerSession) L2STANDARDBRIDGEADDRESS() (common.Address, error) {
	return _L2AutoWithdrawBridge.Contract.L2STANDARDBRIDGEADDRESS(&_L2AutoWithdrawBridge.CallOpts)
}

// MINWITHDRAWALAMOUNT is a free data retrieval call binding the contract method 0xd3e5792b.
//
// Solidity: function MIN_WITHDRAWAL_AMOUNT() view returns(uint256)
func (_L2AutoWithdrawBridge *L2AutoWithdrawBridgeCaller) MINWITHDRAWALAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L2AutoWithdrawBridge.contract.Call(opts, &out, "MIN_WITHDRAWAL_AMOUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINWITHDRAWALAMOUNT is a free data retrieval call binding the contract method 0xd3e5792b.
//
// Solidity: function MIN_WITHDRAWAL_AMOUNT() view returns(uint256)
func (_L2AutoWithdrawBridge *L2AutoWithdrawBridgeSession) MINWITHDRAWALAMOUNT() (*big.Int, error) {
	return _L2AutoWithdrawBridge.Contract.MINWITHDRAWALAMOUNT(&_L2AutoWithdrawBridge.CallOpts)
}

// MINWITHDRAWALAMOUNT is a free data retrieval call binding the contract method 0xd3e5792b.
//
// Solidity: function MIN_WITHDRAWAL_AMOUNT() view returns(uint256)
func (_L2AutoWithdrawBridge *L2AutoWithdrawBridgeCallerSession) MINWITHDRAWALAMOUNT() (*big.Int, error) {
	return _L2AutoWithdrawBridge.Contract.MINWITHDRAWALAMOUNT(&_L2AutoWithdrawBridge.CallOpts)
}

// RECIPIENT is a free data retrieval call binding the contract method 0x0d9019e1.
//
// Solidity: function RECIPIENT() view returns(address)
func (_L2AutoWithdrawBridge *L2AutoWithdrawBridgeCaller) RECIPIENT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2AutoWithdrawBridge.contract.Call(opts, &out, "RECIPIENT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RECIPIENT is a free data retrieval call binding the contract method 0x0d9019e1.
//
// Solidity: function RECIPIENT() view returns(address)
func (_L2AutoWithdrawBridge *L2AutoWithdrawBridgeSession) RECIPIENT() (common.Address, error) {
	return _L2AutoWithdrawBridge.Contract.RECIPIENT(&_L2AutoWithdrawBridge.CallOpts)
}

// RECIPIENT is a free data retrieval call binding the contract method 0x0d9019e1.
//
// Solidity: function RECIPIENT() view returns(address)
func (_L2AutoWithdrawBridge *L2AutoWithdrawBridgeCallerSession) RECIPIENT() (common.Address, error) {
	return _L2AutoWithdrawBridge.Contract.RECIPIENT(&_L2AutoWithdrawBridge.CallOpts)
}

// WITHDRAWALNETWORK is a free data retrieval call binding the contract method 0xd0e12f90.
//
// Solidity: function WITHDRAWAL_NETWORK() view returns(uint8)
func (_L2AutoWithdrawBridge *L2AutoWithdrawBridgeCaller) WITHDRAWALNETWORK(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _L2AutoWithdrawBridge.contract.Call(opts, &out, "WITHDRAWAL_NETWORK")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// WITHDRAWALNETWORK is a free data retrieval call binding the contract method 0xd0e12f90.
//
// Solidity: function WITHDRAWAL_NETWORK() view returns(uint8)
func (_L2AutoWithdrawBridge *L2AutoWithdrawBridgeSession) WITHDRAWALNETWORK() (uint8, error) {
	return _L2AutoWithdrawBridge.Contract.WITHDRAWALNETWORK(&_L2AutoWithdrawBridge.CallOpts)
}

// WITHDRAWALNETWORK is a free data retrieval call binding the contract method 0xd0e12f90.
//
// Solidity: function WITHDRAWAL_NETWORK() view returns(uint8)
func (_L2AutoWithdrawBridge *L2AutoWithdrawBridgeCallerSession) WITHDRAWALNETWORK() (uint8, error) {
	return _L2AutoWithdrawBridge.Contract.WITHDRAWALNETWORK(&_L2AutoWithdrawBridge.CallOpts)
}

// DelegationFee is a free data retrieval call binding the contract method 0x99f979c5.
//
// Solidity: function delegation_fee() view returns(uint256)
func (_L2AutoWithdrawBridge *L2AutoWithdrawBridgeCaller) DelegationFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L2AutoWithdrawBridge.contract.Call(opts, &out, "delegation_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelegationFee is a free data retrieval call binding the contract method 0x99f979c5.
//
// Solidity: function delegation_fee() view returns(uint256)
func (_L2AutoWithdrawBridge *L2AutoWithdrawBridgeSession) DelegationFee() (*big.Int, error) {
	return _L2AutoWithdrawBridge.Contract.DelegationFee(&_L2AutoWithdrawBridge.CallOpts)
}

// DelegationFee is a free data retrieval call binding the contract method 0x99f979c5.
//
// Solidity: function delegation_fee() view returns(uint256)
func (_L2AutoWithdrawBridge *L2AutoWithdrawBridgeCallerSession) DelegationFee() (*big.Int, error) {
	return _L2AutoWithdrawBridge.Contract.DelegationFee(&_L2AutoWithdrawBridge.CallOpts)
}

// GetDelegateFee is a free data retrieval call binding the contract method 0x72b8a716.
//
// Solidity: function getDelegateFee() view returns(uint256)
func (_L2AutoWithdrawBridge *L2AutoWithdrawBridgeCaller) GetDelegateFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L2AutoWithdrawBridge.contract.Call(opts, &out, "getDelegateFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDelegateFee is a free data retrieval call binding the contract method 0x72b8a716.
//
// Solidity: function getDelegateFee() view returns(uint256)
func (_L2AutoWithdrawBridge *L2AutoWithdrawBridgeSession) GetDelegateFee() (*big.Int, error) {
	return _L2AutoWithdrawBridge.Contract.GetDelegateFee(&_L2AutoWithdrawBridge.CallOpts)
}

// GetDelegateFee is a free data retrieval call binding the contract method 0x72b8a716.
//
// Solidity: function getDelegateFee() view returns(uint256)
func (_L2AutoWithdrawBridge *L2AutoWithdrawBridgeCallerSession) GetDelegateFee() (*big.Int, error) {
	return _L2AutoWithdrawBridge.Contract.GetDelegateFee(&_L2AutoWithdrawBridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L2AutoWithdrawBridge *L2AutoWithdrawBridgeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2AutoWithdrawBridge.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L2AutoWithdrawBridge *L2AutoWithdrawBridgeSession) Owner() (common.Address, error) {
	return _L2AutoWithdrawBridge.Contract.Owner(&_L2AutoWithdrawBridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L2AutoWithdrawBridge *L2AutoWithdrawBridgeCallerSession) Owner() (common.Address, error) {
	return _L2AutoWithdrawBridge.Contract.Owner(&_L2AutoWithdrawBridge.CallOpts)
}

// TotalProcessed is a free data retrieval call binding the contract method 0x84411d65.
//
// Solidity: function totalProcessed() view returns(uint256)
func (_L2AutoWithdrawBridge *L2AutoWithdrawBridgeCaller) TotalProcessed(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L2AutoWithdrawBridge.contract.Call(opts, &out, "totalProcessed")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalProcessed is a free data retrieval call binding the contract method 0x84411d65.
//
// Solidity: function totalProcessed() view returns(uint256)
func (_L2AutoWithdrawBridge *L2AutoWithdrawBridgeSession) TotalProcessed() (*big.Int, error) {
	return _L2AutoWithdrawBridge.Contract.TotalProcessed(&_L2AutoWithdrawBridge.CallOpts)
}

// TotalProcessed is a free data retrieval call binding the contract method 0x84411d65.
//
// Solidity: function totalProcessed() view returns(uint256)
func (_L2AutoWithdrawBridge *L2AutoWithdrawBridgeCallerSession) TotalProcessed() (*big.Int, error) {
	return _L2AutoWithdrawBridge.Contract.TotalProcessed(&_L2AutoWithdrawBridge.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_L2AutoWithdrawBridge *L2AutoWithdrawBridgeCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _L2AutoWithdrawBridge.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_L2AutoWithdrawBridge *L2AutoWithdrawBridgeSession) Version() (string, error) {
	return _L2AutoWithdrawBridge.Contract.Version(&_L2AutoWithdrawBridge.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_L2AutoWithdrawBridge *L2AutoWithdrawBridgeCallerSession) Version() (string, error) {
	return _L2AutoWithdrawBridge.Contract.Version(&_L2AutoWithdrawBridge.CallOpts)
}

// AutoWithdrawTo is a paid mutator transaction binding the contract method 0xa4d3c6f5.
//
// Solidity: function autoWithdrawTo(address _l2Token, address _to, uint256 _amount, uint32 _minGasLimit, bytes _extraData) payable returns()
func (_L2AutoWithdrawBridge *L2AutoWithdrawBridgeTransactor) AutoWithdrawTo(opts *bind.TransactOpts, _l2Token common.Address, _to common.Address, _amount *big.Int, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _L2AutoWithdrawBridge.contract.Transact(opts, "autoWithdrawTo", _l2Token, _to, _amount, _minGasLimit, _extraData)
}

// AutoWithdrawTo is a paid mutator transaction binding the contract method 0xa4d3c6f5.
//
// Solidity: function autoWithdrawTo(address _l2Token, address _to, uint256 _amount, uint32 _minGasLimit, bytes _extraData) payable returns()
func (_L2AutoWithdrawBridge *L2AutoWithdrawBridgeSession) AutoWithdrawTo(_l2Token common.Address, _to common.Address, _amount *big.Int, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _L2AutoWithdrawBridge.Contract.AutoWithdrawTo(&_L2AutoWithdrawBridge.TransactOpts, _l2Token, _to, _amount, _minGasLimit, _extraData)
}

// AutoWithdrawTo is a paid mutator transaction binding the contract method 0xa4d3c6f5.
//
// Solidity: function autoWithdrawTo(address _l2Token, address _to, uint256 _amount, uint32 _minGasLimit, bytes _extraData) payable returns()
func (_L2AutoWithdrawBridge *L2AutoWithdrawBridgeTransactorSession) AutoWithdrawTo(_l2Token common.Address, _to common.Address, _amount *big.Int, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _L2AutoWithdrawBridge.Contract.AutoWithdrawTo(&_L2AutoWithdrawBridge.TransactOpts, _l2Token, _to, _amount, _minGasLimit, _extraData)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_L2AutoWithdrawBridge *L2AutoWithdrawBridgeTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _L2AutoWithdrawBridge.contract.Transact(opts, "initialize", _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_L2AutoWithdrawBridge *L2AutoWithdrawBridgeSession) Initialize(_owner common.Address) (*types.Transaction, error) {
	return _L2AutoWithdrawBridge.Contract.Initialize(&_L2AutoWithdrawBridge.TransactOpts, _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_L2AutoWithdrawBridge *L2AutoWithdrawBridgeTransactorSession) Initialize(_owner common.Address) (*types.Transaction, error) {
	return _L2AutoWithdrawBridge.Contract.Initialize(&_L2AutoWithdrawBridge.TransactOpts, _owner)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L2AutoWithdrawBridge *L2AutoWithdrawBridgeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2AutoWithdrawBridge.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L2AutoWithdrawBridge *L2AutoWithdrawBridgeSession) RenounceOwnership() (*types.Transaction, error) {
	return _L2AutoWithdrawBridge.Contract.RenounceOwnership(&_L2AutoWithdrawBridge.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L2AutoWithdrawBridge *L2AutoWithdrawBridgeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _L2AutoWithdrawBridge.Contract.RenounceOwnership(&_L2AutoWithdrawBridge.TransactOpts)
}

// SetDelegateFee is a paid mutator transaction binding the contract method 0x9931c34c.
//
// Solidity: function setDelegateFee(uint256 _delegateFee) returns()
func (_L2AutoWithdrawBridge *L2AutoWithdrawBridgeTransactor) SetDelegateFee(opts *bind.TransactOpts, _delegateFee *big.Int) (*types.Transaction, error) {
	return _L2AutoWithdrawBridge.contract.Transact(opts, "setDelegateFee", _delegateFee)
}

// SetDelegateFee is a paid mutator transaction binding the contract method 0x9931c34c.
//
// Solidity: function setDelegateFee(uint256 _delegateFee) returns()
func (_L2AutoWithdrawBridge *L2AutoWithdrawBridgeSession) SetDelegateFee(_delegateFee *big.Int) (*types.Transaction, error) {
	return _L2AutoWithdrawBridge.Contract.SetDelegateFee(&_L2AutoWithdrawBridge.TransactOpts, _delegateFee)
}

// SetDelegateFee is a paid mutator transaction binding the contract method 0x9931c34c.
//
// Solidity: function setDelegateFee(uint256 _delegateFee) returns()
func (_L2AutoWithdrawBridge *L2AutoWithdrawBridgeTransactorSession) SetDelegateFee(_delegateFee *big.Int) (*types.Transaction, error) {
	return _L2AutoWithdrawBridge.Contract.SetDelegateFee(&_L2AutoWithdrawBridge.TransactOpts, _delegateFee)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L2AutoWithdrawBridge *L2AutoWithdrawBridgeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _L2AutoWithdrawBridge.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L2AutoWithdrawBridge *L2AutoWithdrawBridgeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L2AutoWithdrawBridge.Contract.TransferOwnership(&_L2AutoWithdrawBridge.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L2AutoWithdrawBridge *L2AutoWithdrawBridgeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L2AutoWithdrawBridge.Contract.TransferOwnership(&_L2AutoWithdrawBridge.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_L2AutoWithdrawBridge *L2AutoWithdrawBridgeTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2AutoWithdrawBridge.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_L2AutoWithdrawBridge *L2AutoWithdrawBridgeSession) Withdraw() (*types.Transaction, error) {
	return _L2AutoWithdrawBridge.Contract.Withdraw(&_L2AutoWithdrawBridge.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_L2AutoWithdrawBridge *L2AutoWithdrawBridgeTransactorSession) Withdraw() (*types.Transaction, error) {
	return _L2AutoWithdrawBridge.Contract.Withdraw(&_L2AutoWithdrawBridge.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_L2AutoWithdrawBridge *L2AutoWithdrawBridgeTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2AutoWithdrawBridge.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_L2AutoWithdrawBridge *L2AutoWithdrawBridgeSession) Receive() (*types.Transaction, error) {
	return _L2AutoWithdrawBridge.Contract.Receive(&_L2AutoWithdrawBridge.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_L2AutoWithdrawBridge *L2AutoWithdrawBridgeTransactorSession) Receive() (*types.Transaction, error) {
	return _L2AutoWithdrawBridge.Contract.Receive(&_L2AutoWithdrawBridge.TransactOpts)
}

// L2AutoWithdrawBridgeAutoWithdrawToIterator is returned from FilterAutoWithdrawTo and is used to iterate over the raw logs and unpacked data for AutoWithdrawTo events raised by the L2AutoWithdrawBridge contract.
type L2AutoWithdrawBridgeAutoWithdrawToIterator struct {
	Event *L2AutoWithdrawBridgeAutoWithdrawTo // Event containing the contract specifics and raw log

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
func (it *L2AutoWithdrawBridgeAutoWithdrawToIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2AutoWithdrawBridgeAutoWithdrawTo)
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
		it.Event = new(L2AutoWithdrawBridgeAutoWithdrawTo)
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
func (it *L2AutoWithdrawBridgeAutoWithdrawToIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2AutoWithdrawBridgeAutoWithdrawToIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2AutoWithdrawBridgeAutoWithdrawTo represents a AutoWithdrawTo event raised by the L2AutoWithdrawBridge contract.
type L2AutoWithdrawBridgeAutoWithdrawTo struct {
	From      common.Address
	To        common.Address
	Amount    *big.Int
	ExtraData []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAutoWithdrawTo is a free log retrieval operation binding the contract event 0xf3f58d6a31b0011b592ca174b425f161b372632c23c93791414c2a29ee3bdf07.
//
// Solidity: event AutoWithdrawTo(address indexed from, address to, uint256 amount, bytes extraData)
func (_L2AutoWithdrawBridge *L2AutoWithdrawBridgeFilterer) FilterAutoWithdrawTo(opts *bind.FilterOpts, from []common.Address) (*L2AutoWithdrawBridgeAutoWithdrawToIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _L2AutoWithdrawBridge.contract.FilterLogs(opts, "AutoWithdrawTo", fromRule)
	if err != nil {
		return nil, err
	}
	return &L2AutoWithdrawBridgeAutoWithdrawToIterator{contract: _L2AutoWithdrawBridge.contract, event: "AutoWithdrawTo", logs: logs, sub: sub}, nil
}

// WatchAutoWithdrawTo is a free log subscription operation binding the contract event 0xf3f58d6a31b0011b592ca174b425f161b372632c23c93791414c2a29ee3bdf07.
//
// Solidity: event AutoWithdrawTo(address indexed from, address to, uint256 amount, bytes extraData)
func (_L2AutoWithdrawBridge *L2AutoWithdrawBridgeFilterer) WatchAutoWithdrawTo(opts *bind.WatchOpts, sink chan<- *L2AutoWithdrawBridgeAutoWithdrawTo, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _L2AutoWithdrawBridge.contract.WatchLogs(opts, "AutoWithdrawTo", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2AutoWithdrawBridgeAutoWithdrawTo)
				if err := _L2AutoWithdrawBridge.contract.UnpackLog(event, "AutoWithdrawTo", log); err != nil {
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

// ParseAutoWithdrawTo is a log parse operation binding the contract event 0xf3f58d6a31b0011b592ca174b425f161b372632c23c93791414c2a29ee3bdf07.
//
// Solidity: event AutoWithdrawTo(address indexed from, address to, uint256 amount, bytes extraData)
func (_L2AutoWithdrawBridge *L2AutoWithdrawBridgeFilterer) ParseAutoWithdrawTo(log types.Log) (*L2AutoWithdrawBridgeAutoWithdrawTo, error) {
	event := new(L2AutoWithdrawBridgeAutoWithdrawTo)
	if err := _L2AutoWithdrawBridge.contract.UnpackLog(event, "AutoWithdrawTo", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2AutoWithdrawBridgeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the L2AutoWithdrawBridge contract.
type L2AutoWithdrawBridgeInitializedIterator struct {
	Event *L2AutoWithdrawBridgeInitialized // Event containing the contract specifics and raw log

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
func (it *L2AutoWithdrawBridgeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2AutoWithdrawBridgeInitialized)
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
		it.Event = new(L2AutoWithdrawBridgeInitialized)
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
func (it *L2AutoWithdrawBridgeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2AutoWithdrawBridgeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2AutoWithdrawBridgeInitialized represents a Initialized event raised by the L2AutoWithdrawBridge contract.
type L2AutoWithdrawBridgeInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L2AutoWithdrawBridge *L2AutoWithdrawBridgeFilterer) FilterInitialized(opts *bind.FilterOpts) (*L2AutoWithdrawBridgeInitializedIterator, error) {

	logs, sub, err := _L2AutoWithdrawBridge.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &L2AutoWithdrawBridgeInitializedIterator{contract: _L2AutoWithdrawBridge.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L2AutoWithdrawBridge *L2AutoWithdrawBridgeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *L2AutoWithdrawBridgeInitialized) (event.Subscription, error) {

	logs, sub, err := _L2AutoWithdrawBridge.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2AutoWithdrawBridgeInitialized)
				if err := _L2AutoWithdrawBridge.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L2AutoWithdrawBridge *L2AutoWithdrawBridgeFilterer) ParseInitialized(log types.Log) (*L2AutoWithdrawBridgeInitialized, error) {
	event := new(L2AutoWithdrawBridgeInitialized)
	if err := _L2AutoWithdrawBridge.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2AutoWithdrawBridgeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the L2AutoWithdrawBridge contract.
type L2AutoWithdrawBridgeOwnershipTransferredIterator struct {
	Event *L2AutoWithdrawBridgeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *L2AutoWithdrawBridgeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2AutoWithdrawBridgeOwnershipTransferred)
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
		it.Event = new(L2AutoWithdrawBridgeOwnershipTransferred)
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
func (it *L2AutoWithdrawBridgeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2AutoWithdrawBridgeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2AutoWithdrawBridgeOwnershipTransferred represents a OwnershipTransferred event raised by the L2AutoWithdrawBridge contract.
type L2AutoWithdrawBridgeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L2AutoWithdrawBridge *L2AutoWithdrawBridgeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*L2AutoWithdrawBridgeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L2AutoWithdrawBridge.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &L2AutoWithdrawBridgeOwnershipTransferredIterator{contract: _L2AutoWithdrawBridge.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L2AutoWithdrawBridge *L2AutoWithdrawBridgeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *L2AutoWithdrawBridgeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L2AutoWithdrawBridge.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2AutoWithdrawBridgeOwnershipTransferred)
				if err := _L2AutoWithdrawBridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L2AutoWithdrawBridge *L2AutoWithdrawBridgeFilterer) ParseOwnershipTransferred(log types.Log) (*L2AutoWithdrawBridgeOwnershipTransferred, error) {
	event := new(L2AutoWithdrawBridgeOwnershipTransferred)
	if err := _L2AutoWithdrawBridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2AutoWithdrawBridgeWithdrawalIterator is returned from FilterWithdrawal and is used to iterate over the raw logs and unpacked data for Withdrawal events raised by the L2AutoWithdrawBridge contract.
type L2AutoWithdrawBridgeWithdrawalIterator struct {
	Event *L2AutoWithdrawBridgeWithdrawal // Event containing the contract specifics and raw log

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
func (it *L2AutoWithdrawBridgeWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2AutoWithdrawBridgeWithdrawal)
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
		it.Event = new(L2AutoWithdrawBridgeWithdrawal)
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
func (it *L2AutoWithdrawBridgeWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2AutoWithdrawBridgeWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2AutoWithdrawBridgeWithdrawal represents a Withdrawal event raised by the L2AutoWithdrawBridge contract.
type L2AutoWithdrawBridgeWithdrawal struct {
	Value *big.Int
	To    common.Address
	From  common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0xc8a211cc64b6ed1b50595a9fcb1932b6d1e5a6e8ef15b60e5b1f988ea9086bba.
//
// Solidity: event Withdrawal(uint256 value, address to, address from)
func (_L2AutoWithdrawBridge *L2AutoWithdrawBridgeFilterer) FilterWithdrawal(opts *bind.FilterOpts) (*L2AutoWithdrawBridgeWithdrawalIterator, error) {

	logs, sub, err := _L2AutoWithdrawBridge.contract.FilterLogs(opts, "Withdrawal")
	if err != nil {
		return nil, err
	}
	return &L2AutoWithdrawBridgeWithdrawalIterator{contract: _L2AutoWithdrawBridge.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

// WatchWithdrawal is a free log subscription operation binding the contract event 0xc8a211cc64b6ed1b50595a9fcb1932b6d1e5a6e8ef15b60e5b1f988ea9086bba.
//
// Solidity: event Withdrawal(uint256 value, address to, address from)
func (_L2AutoWithdrawBridge *L2AutoWithdrawBridgeFilterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *L2AutoWithdrawBridgeWithdrawal) (event.Subscription, error) {

	logs, sub, err := _L2AutoWithdrawBridge.contract.WatchLogs(opts, "Withdrawal")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2AutoWithdrawBridgeWithdrawal)
				if err := _L2AutoWithdrawBridge.contract.UnpackLog(event, "Withdrawal", log); err != nil {
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

// ParseWithdrawal is a log parse operation binding the contract event 0xc8a211cc64b6ed1b50595a9fcb1932b6d1e5a6e8ef15b60e5b1f988ea9086bba.
//
// Solidity: event Withdrawal(uint256 value, address to, address from)
func (_L2AutoWithdrawBridge *L2AutoWithdrawBridgeFilterer) ParseWithdrawal(log types.Log) (*L2AutoWithdrawBridgeWithdrawal, error) {
	event := new(L2AutoWithdrawBridgeWithdrawal)
	if err := _L2AutoWithdrawBridge.contract.UnpackLog(event, "Withdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2AutoWithdrawBridgeWithdrawal0Iterator is returned from FilterWithdrawal0 and is used to iterate over the raw logs and unpacked data for Withdrawal0 events raised by the L2AutoWithdrawBridge contract.
type L2AutoWithdrawBridgeWithdrawal0Iterator struct {
	Event *L2AutoWithdrawBridgeWithdrawal0 // Event containing the contract specifics and raw log

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
func (it *L2AutoWithdrawBridgeWithdrawal0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2AutoWithdrawBridgeWithdrawal0)
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
		it.Event = new(L2AutoWithdrawBridgeWithdrawal0)
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
func (it *L2AutoWithdrawBridgeWithdrawal0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2AutoWithdrawBridgeWithdrawal0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2AutoWithdrawBridgeWithdrawal0 represents a Withdrawal0 event raised by the L2AutoWithdrawBridge contract.
type L2AutoWithdrawBridgeWithdrawal0 struct {
	Value             *big.Int
	To                common.Address
	From              common.Address
	WithdrawalNetwork uint8
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal0 is a free log retrieval operation binding the contract event 0x38e04cbeb8c10f8f568618aa75be0f10b6729b8b4237743b4de20cbcde2839ee.
//
// Solidity: event Withdrawal(uint256 value, address to, address from, uint8 withdrawalNetwork)
func (_L2AutoWithdrawBridge *L2AutoWithdrawBridgeFilterer) FilterWithdrawal0(opts *bind.FilterOpts) (*L2AutoWithdrawBridgeWithdrawal0Iterator, error) {

	logs, sub, err := _L2AutoWithdrawBridge.contract.FilterLogs(opts, "Withdrawal0")
	if err != nil {
		return nil, err
	}
	return &L2AutoWithdrawBridgeWithdrawal0Iterator{contract: _L2AutoWithdrawBridge.contract, event: "Withdrawal0", logs: logs, sub: sub}, nil
}

// WatchWithdrawal0 is a free log subscription operation binding the contract event 0x38e04cbeb8c10f8f568618aa75be0f10b6729b8b4237743b4de20cbcde2839ee.
//
// Solidity: event Withdrawal(uint256 value, address to, address from, uint8 withdrawalNetwork)
func (_L2AutoWithdrawBridge *L2AutoWithdrawBridgeFilterer) WatchWithdrawal0(opts *bind.WatchOpts, sink chan<- *L2AutoWithdrawBridgeWithdrawal0) (event.Subscription, error) {

	logs, sub, err := _L2AutoWithdrawBridge.contract.WatchLogs(opts, "Withdrawal0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2AutoWithdrawBridgeWithdrawal0)
				if err := _L2AutoWithdrawBridge.contract.UnpackLog(event, "Withdrawal0", log); err != nil {
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

// ParseWithdrawal0 is a log parse operation binding the contract event 0x38e04cbeb8c10f8f568618aa75be0f10b6729b8b4237743b4de20cbcde2839ee.
//
// Solidity: event Withdrawal(uint256 value, address to, address from, uint8 withdrawalNetwork)
func (_L2AutoWithdrawBridge *L2AutoWithdrawBridgeFilterer) ParseWithdrawal0(log types.Log) (*L2AutoWithdrawBridgeWithdrawal0, error) {
	event := new(L2AutoWithdrawBridgeWithdrawal0)
	if err := _L2AutoWithdrawBridge.contract.UnpackLog(event, "Withdrawal0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
