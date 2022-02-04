// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package AtomicBond

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
)

// AtomicBondMetaData contains all meta data concerning the AtomicBond contract.
var AtomicBondMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_MIM\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_TIME\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_MEMO\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_TIME_STAKING\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_JOE_ROUTER_02\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_MIM_TIME_BOND_DEPOSITORY\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"approveERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"approveFunder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"approveMinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"approveUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_funder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_minter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_bondIn\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_swapPath\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"_minSwapAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxBondPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"}],\"name\":\"atomicBondMemoMim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"finalize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"recoverLostNetworkToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"recoverLostTokenERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_minter\",\"type\":\"address\"}],\"name\":\"redeem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"revokeFunder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"revokeMinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"revokeUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6101606040523480156200001257600080fd5b50604051620022503803806200225083398101604081905262000035916200029f565b6001600160a01b0386166200004957600080fd5b6001600160a01b0385166200005d57600080fd5b6001600160a01b0384166200007157600080fd5b6001600160a01b0383166200008557600080fd5b6001600160a01b0382166200009957600080fd5b6001600160a01b038116620000ad57600080fd5b336080526001600160a01b0386811660a05285811660c05284811660e0819052848216610100819052848316610120529183166101405260405163095ea7b360e01b8152600481019290925260001960248301529063095ea7b3906044016020604051808303816000875af11580156200012b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062000151919062000320565b5060c0516101205160405163095ea7b360e01b81526001600160a01b039182166004820152600019602482015291169063095ea7b3906044016020604051808303816000875af1158015620001aa573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001d0919062000320565b5060a0516101405160405163095ea7b360e01b81526001600160a01b039182166004820152600019602482015291169063095ea7b3906044016020604051808303816000875af115801562000229573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200024f919062000320565b50506080516001600160a01b03166000908152602081905260409020805460ff19166001179055506200034b9350505050565b80516001600160a01b03811681146200029a57600080fd5b919050565b60008060008060008060c08789031215620002b957600080fd5b620002c48762000282565b9550620002d46020880162000282565b9450620002e46040880162000282565b9350620002f46060880162000282565b9250620003046080880162000282565b91506200031460a0880162000282565b90509295509295509295565b6000602082840312156200033357600080fd5b815180151581146200034457600080fd5b9392505050565b60805160a05160c05160e051610100516101205161014051611e2b62000425600039600081816105eb0152610d8e01526000610542015260006104920152600081816104340152818161099301528181610ba30152610e4101526000505060005050600081816106e5015281816107d00152818161086f015281816108a801528181610ab801528181610e6901528181610efb0152818161103901528181611124015281816111c3015281816112c801528181611382015281816113a9015281816114480152818161154a01526116010152611e2b6000f3fe608060405234801561001057600080fd5b50600436106100d45760003560e01c8063a8e5e4aa11610081578063e3e6cd181161005b578063e3e6cd181461017b578063ef097fc41461018e578063fa4ea2be146101a157600080fd5b8063a8e5e4aa14610142578063cfbd488514610155578063d624fdb61461016857600080fd5b80634d197a9b116100b25780634d197a9b1461010957806391aefcd81461011c57806395a2251f1461012f57600080fd5b80630dde5052146100d95780633cd80377146100ee5780634bb278f314610101575b600080fd5b6100ec6100e73660046119e3565b6101a9565b005b6100ec6100fc366004611a9d565b6106e3565b6100ec6107ce565b6100ec610117366004611a9d565b6108a6565b6100ec61012a366004611a9d565b610ab6565b6100ec61013d366004611a9d565b610cc9565b6100ec610150366004611ab8565b610ef9565b6100ec610163366004611a9d565b611037565b6100ec610176366004611a9d565b611122565b6100ec610189366004611af4565b6112c6565b6100ec61019c366004611a9d565b6113a7565b6100ec611548565b3360009081526020819052604090205460ff16610227576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f4e6f7420757365722e000000000000000000000000000000000000000000000060448201526064015b60405180910390fd5b8042811015610292576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600860248201527f457870697265642e000000000000000000000000000000000000000000000000604482015260640161021e565b600084116102fc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f496e76616c6964205f6d696e53776170416d6f756e742e000000000000000000604482015260640161021e565b73ffffffffffffffffffffffffffffffffffffffff891660009081526001602052604090205460ff1661038b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600b60248201527f4e6f742066756e6465722e000000000000000000000000000000000000000000604482015260640161021e565b73ffffffffffffffffffffffffffffffffffffffff881660009081526002602052604090205460ff1661041a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600b60248201527f4e6f74206d696e7465722e000000000000000000000000000000000000000000604482015260640161021e565b61045c73ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168a308a61164c565b6040517f9ebea88c00000000000000000000000000000000000000000000000000000000815260048101889052600160248201527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1690639ebea88c90604401600060405180830381600087803b1580156104eb57600080fd5b505af11580156104ff573d6000803e3d6000fd5b50506040517f38ed17390000000000000000000000000000000000000000000000000000000081526000925073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001691506338ed173990610582908b9089908c908c9030908b90600401611b1e565b6000604051808303816000875af11580156105a1573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01682016040526105e79190810190611bd8565b90507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16638dbdbe6d82600184516106349190611cb4565b8151811061064457610644611cf2565b6020026020010151868c6040518463ffffffff1660e01b815260040161069393929190928352602083019190915273ffffffffffffffffffffffffffffffffffffffff16604082015260600190565b6020604051808303816000875af11580156106b2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106d69190611d21565b5050505050505050505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163314610782576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600a60248201527f4e6f74206f776e65722e00000000000000000000000000000000000000000000604482015260640161021e565b73ffffffffffffffffffffffffffffffffffffffff16600090815260016020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16331461086d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600a60248201527f4e6f74206f776e65722e00000000000000000000000000000000000000000000604482015260640161021e565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16ff5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163314610945576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600a60248201527f4e6f74206f776e65722e00000000000000000000000000000000000000000000604482015260640161021e565b6040517fdd62ed3e00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82811660048301523060248301526000917f00000000000000000000000000000000000000000000000000000000000000009091169063dd62ed3e90604401602060405180830381865afa1580156109dc573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a009190611d21565b11610a67576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f4e6f7420617070726f7665642e00000000000000000000000000000000000000604482015260640161021e565b73ffffffffffffffffffffffffffffffffffffffff16600090815260026020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163314610b55576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600a60248201527f4e6f74206f776e65722e00000000000000000000000000000000000000000000604482015260640161021e565b6040517fdd62ed3e00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82811660048301523060248301526000917f00000000000000000000000000000000000000000000000000000000000000009091169063dd62ed3e90604401602060405180830381865afa158015610bec573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c109190611d21565b11610c77576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f4e6f7420617070726f7665642e00000000000000000000000000000000000000604482015260640161021e565b73ffffffffffffffffffffffffffffffffffffffff16600090815260016020819052604090912080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169091179055565b3360009081526020819052604090205460ff16610d42576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f4e6f7420757365722e0000000000000000000000000000000000000000000000604482015260640161021e565b6040517f1feed31f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8281166004830152600160248301527f00000000000000000000000000000000000000000000000000000000000000001690631feed31f906044016020604051808303816000875af1158015610dd7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610dfb9190611d21565b506040517f70a0823100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff80831660048301527f000000000000000000000000000000000000000000000000000000000000000091610ef59184917f0000000000000000000000000000000000000000000000000000000000000000918516906370a0823190602401602060405180830381865afa158015610eb2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ed69190611d21565b73ffffffffffffffffffffffffffffffffffffffff851692919061164c565b5050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163314610f98576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600a60248201527f4e6f74206f776e65722e00000000000000000000000000000000000000000000604482015260640161021e565b6040517f095ea7b300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83811660048301526024820183905284169063095ea7b3906044016020604051808303816000875af115801561100d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110319190611d3a565b50505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1633146110d6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600a60248201527f4e6f74206f776e65722e00000000000000000000000000000000000000000000604482015260640161021e565b73ffffffffffffffffffffffffffffffffffffffff16600090815260026020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1633146111c1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600a60248201527f4e6f74206f776e65722e00000000000000000000000000000000000000000000604482015260640161021e565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415611277576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f43616e6e6f7420617070726f7665206f776e65722e0000000000000000000000604482015260640161021e565b73ffffffffffffffffffffffffffffffffffffffff16600090815260208190526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163314611365576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600a60248201527f4e6f74206f776e65722e00000000000000000000000000000000000000000000604482015260640161021e565b610ef573ffffffffffffffffffffffffffffffffffffffff8316307f00000000000000000000000000000000000000000000000000000000000000008461164c565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163314611446576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600a60248201527f4e6f74206f776e65722e00000000000000000000000000000000000000000000604482015260640161021e565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156114fc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f43616e6e6f74207265766f6b65206f776e65722e000000000000000000000000604482015260640161021e565b73ffffffffffffffffffffffffffffffffffffffff16600090815260208190526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1633146115e7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600a60248201527f4e6f74206f776e65722e00000000000000000000000000000000000000000000604482015260640161021e565b60405173ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016904780156108fc02916000818181858888f19350505050158015611649573d6000803e3d6000fd5b50565b6040805173ffffffffffffffffffffffffffffffffffffffff8581166024830152848116604483015260648083018590528351808403909101815260849092018352602080830180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f23b872dd0000000000000000000000000000000000000000000000000000000017905283518085019094528084527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564908401526110319287929160009161171f9185169084906117ce565b8051909150156117c9578080602001905181019061173d9190611d3a565b6117c9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f74207375636365656400000000000000000000000000000000000000000000606482015260840161021e565b505050565b60606117dd84846000856117e7565b90505b9392505050565b606082471015611879576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c0000000000000000000000000000000000000000000000000000606482015260840161021e565b843b6118e1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161021e565b6000808673ffffffffffffffffffffffffffffffffffffffff16858760405161190a9190611d88565b60006040518083038185875af1925050503d8060008114611947576040519150601f19603f3d011682016040523d82523d6000602084013e61194c565b606091505b509150915061195c828286611967565b979650505050505050565b606083156119765750816117e0565b8251156119865782518084602001fd5b816040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161021e9190611da4565b803573ffffffffffffffffffffffffffffffffffffffff811681146119de57600080fd5b919050565b60008060008060008060008060e0898b0312156119ff57600080fd5b611a08896119ba565b9750611a1660208a016119ba565b965060408901359550606089013567ffffffffffffffff80821115611a3a57600080fd5b818b0191508b601f830112611a4e57600080fd5b813581811115611a5d57600080fd5b8c60208260051b8501011115611a7257600080fd5b999c989b50969960209190910198976080820135975060a0820135965060c090910135945092505050565b600060208284031215611aaf57600080fd5b6117e0826119ba565b600080600060608486031215611acd57600080fd5b611ad6846119ba565b9250611ae4602085016119ba565b9150604084013590509250925092565b60008060408385031215611b0757600080fd5b611b10836119ba565b946020939093013593505050565b868152602080820187905260a0604083018190528201859052600090869060c08401835b88811015611b7b5773ffffffffffffffffffffffffffffffffffffffff611b68856119ba565b1682529282019290820190600101611b42565b5073ffffffffffffffffffffffffffffffffffffffff96909616606085015250505060800152949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60006020808385031215611beb57600080fd5b825167ffffffffffffffff80821115611c0357600080fd5b818501915085601f830112611c1757600080fd5b815181811115611c2957611c29611ba9565b8060051b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f83011681018181108582111715611c6c57611c6c611ba9565b604052918252848201925083810185019188831115611c8a57600080fd5b938501935b82851015611ca857845184529385019392850192611c8f565b98975050505050505050565b600082821015611ced577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600060208284031215611d3357600080fd5b5051919050565b600060208284031215611d4c57600080fd5b815180151581146117e057600080fd5b60005b83811015611d77578181015183820152602001611d5f565b838111156110315750506000910152565b60008251611d9a818460208701611d5c565b9190910192915050565b6020815260008251806020840152611dc3816040850160208701611d5c565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016919091016040019291505056fea2646970667358221220c4d44bdc2c67be55f93e49cb30111cdbd50cf43e0ea31b04336959ffd7731ce664736f6c634300080b0033",
}

// AtomicBondABI is the input ABI used to generate the binding from.
// Deprecated: Use AtomicBondMetaData.ABI instead.
var AtomicBondABI = AtomicBondMetaData.ABI

// AtomicBondBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AtomicBondMetaData.Bin instead.
var AtomicBondBin = AtomicBondMetaData.Bin

// DeployAtomicBond deploys a new Ethereum contract, binding an instance of AtomicBond to it.
func DeployAtomicBond(auth *bind.TransactOpts, backend bind.ContractBackend, _MIM common.Address, _TIME common.Address, _MEMO common.Address, _TIME_STAKING common.Address, _JOE_ROUTER_02 common.Address, _MIM_TIME_BOND_DEPOSITORY common.Address) (common.Address, *types.Transaction, *AtomicBond, error) {
	parsed, err := AtomicBondMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AtomicBondBin), backend, _MIM, _TIME, _MEMO, _TIME_STAKING, _JOE_ROUTER_02, _MIM_TIME_BOND_DEPOSITORY)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AtomicBond{AtomicBondCaller: AtomicBondCaller{contract: contract}, AtomicBondTransactor: AtomicBondTransactor{contract: contract}, AtomicBondFilterer: AtomicBondFilterer{contract: contract}}, nil
}

// AtomicBond is an auto generated Go binding around an Ethereum contract.
type AtomicBond struct {
	AtomicBondCaller     // Read-only binding to the contract
	AtomicBondTransactor // Write-only binding to the contract
	AtomicBondFilterer   // Log filterer for contract events
}

// AtomicBondCaller is an auto generated read-only Go binding around an Ethereum contract.
type AtomicBondCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AtomicBondTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AtomicBondTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AtomicBondFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AtomicBondFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AtomicBondSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AtomicBondSession struct {
	Contract     *AtomicBond       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AtomicBondCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AtomicBondCallerSession struct {
	Contract *AtomicBondCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// AtomicBondTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AtomicBondTransactorSession struct {
	Contract     *AtomicBondTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// AtomicBondRaw is an auto generated low-level Go binding around an Ethereum contract.
type AtomicBondRaw struct {
	Contract *AtomicBond // Generic contract binding to access the raw methods on
}

// AtomicBondCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AtomicBondCallerRaw struct {
	Contract *AtomicBondCaller // Generic read-only contract binding to access the raw methods on
}

// AtomicBondTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AtomicBondTransactorRaw struct {
	Contract *AtomicBondTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAtomicBond creates a new instance of AtomicBond, bound to a specific deployed contract.
func NewAtomicBond(address common.Address, backend bind.ContractBackend) (*AtomicBond, error) {
	contract, err := bindAtomicBond(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AtomicBond{AtomicBondCaller: AtomicBondCaller{contract: contract}, AtomicBondTransactor: AtomicBondTransactor{contract: contract}, AtomicBondFilterer: AtomicBondFilterer{contract: contract}}, nil
}

// NewAtomicBondCaller creates a new read-only instance of AtomicBond, bound to a specific deployed contract.
func NewAtomicBondCaller(address common.Address, caller bind.ContractCaller) (*AtomicBondCaller, error) {
	contract, err := bindAtomicBond(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AtomicBondCaller{contract: contract}, nil
}

// NewAtomicBondTransactor creates a new write-only instance of AtomicBond, bound to a specific deployed contract.
func NewAtomicBondTransactor(address common.Address, transactor bind.ContractTransactor) (*AtomicBondTransactor, error) {
	contract, err := bindAtomicBond(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AtomicBondTransactor{contract: contract}, nil
}

// NewAtomicBondFilterer creates a new log filterer instance of AtomicBond, bound to a specific deployed contract.
func NewAtomicBondFilterer(address common.Address, filterer bind.ContractFilterer) (*AtomicBondFilterer, error) {
	contract, err := bindAtomicBond(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AtomicBondFilterer{contract: contract}, nil
}

// bindAtomicBond binds a generic wrapper to an already deployed contract.
func bindAtomicBond(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AtomicBondABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AtomicBond *AtomicBondRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AtomicBond.Contract.AtomicBondCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AtomicBond *AtomicBondRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AtomicBond.Contract.AtomicBondTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AtomicBond *AtomicBondRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AtomicBond.Contract.AtomicBondTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AtomicBond *AtomicBondCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AtomicBond.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AtomicBond *AtomicBondTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AtomicBond.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AtomicBond *AtomicBondTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AtomicBond.Contract.contract.Transact(opts, method, params...)
}

// ApproveERC20 is a paid mutator transaction binding the contract method 0xa8e5e4aa.
//
// Solidity: function approveERC20(address _token, address _address, uint256 _amount) returns()
func (_AtomicBond *AtomicBondTransactor) ApproveERC20(opts *bind.TransactOpts, _token common.Address, _address common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _AtomicBond.contract.Transact(opts, "approveERC20", _token, _address, _amount)
}

// ApproveERC20 is a paid mutator transaction binding the contract method 0xa8e5e4aa.
//
// Solidity: function approveERC20(address _token, address _address, uint256 _amount) returns()
func (_AtomicBond *AtomicBondSession) ApproveERC20(_token common.Address, _address common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _AtomicBond.Contract.ApproveERC20(&_AtomicBond.TransactOpts, _token, _address, _amount)
}

// ApproveERC20 is a paid mutator transaction binding the contract method 0xa8e5e4aa.
//
// Solidity: function approveERC20(address _token, address _address, uint256 _amount) returns()
func (_AtomicBond *AtomicBondTransactorSession) ApproveERC20(_token common.Address, _address common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _AtomicBond.Contract.ApproveERC20(&_AtomicBond.TransactOpts, _token, _address, _amount)
}

// ApproveFunder is a paid mutator transaction binding the contract method 0x91aefcd8.
//
// Solidity: function approveFunder(address _address) returns()
func (_AtomicBond *AtomicBondTransactor) ApproveFunder(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _AtomicBond.contract.Transact(opts, "approveFunder", _address)
}

// ApproveFunder is a paid mutator transaction binding the contract method 0x91aefcd8.
//
// Solidity: function approveFunder(address _address) returns()
func (_AtomicBond *AtomicBondSession) ApproveFunder(_address common.Address) (*types.Transaction, error) {
	return _AtomicBond.Contract.ApproveFunder(&_AtomicBond.TransactOpts, _address)
}

// ApproveFunder is a paid mutator transaction binding the contract method 0x91aefcd8.
//
// Solidity: function approveFunder(address _address) returns()
func (_AtomicBond *AtomicBondTransactorSession) ApproveFunder(_address common.Address) (*types.Transaction, error) {
	return _AtomicBond.Contract.ApproveFunder(&_AtomicBond.TransactOpts, _address)
}

// ApproveMinter is a paid mutator transaction binding the contract method 0x4d197a9b.
//
// Solidity: function approveMinter(address _address) returns()
func (_AtomicBond *AtomicBondTransactor) ApproveMinter(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _AtomicBond.contract.Transact(opts, "approveMinter", _address)
}

// ApproveMinter is a paid mutator transaction binding the contract method 0x4d197a9b.
//
// Solidity: function approveMinter(address _address) returns()
func (_AtomicBond *AtomicBondSession) ApproveMinter(_address common.Address) (*types.Transaction, error) {
	return _AtomicBond.Contract.ApproveMinter(&_AtomicBond.TransactOpts, _address)
}

// ApproveMinter is a paid mutator transaction binding the contract method 0x4d197a9b.
//
// Solidity: function approveMinter(address _address) returns()
func (_AtomicBond *AtomicBondTransactorSession) ApproveMinter(_address common.Address) (*types.Transaction, error) {
	return _AtomicBond.Contract.ApproveMinter(&_AtomicBond.TransactOpts, _address)
}

// ApproveUser is a paid mutator transaction binding the contract method 0xd624fdb6.
//
// Solidity: function approveUser(address _address) returns()
func (_AtomicBond *AtomicBondTransactor) ApproveUser(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _AtomicBond.contract.Transact(opts, "approveUser", _address)
}

// ApproveUser is a paid mutator transaction binding the contract method 0xd624fdb6.
//
// Solidity: function approveUser(address _address) returns()
func (_AtomicBond *AtomicBondSession) ApproveUser(_address common.Address) (*types.Transaction, error) {
	return _AtomicBond.Contract.ApproveUser(&_AtomicBond.TransactOpts, _address)
}

// ApproveUser is a paid mutator transaction binding the contract method 0xd624fdb6.
//
// Solidity: function approveUser(address _address) returns()
func (_AtomicBond *AtomicBondTransactorSession) ApproveUser(_address common.Address) (*types.Transaction, error) {
	return _AtomicBond.Contract.ApproveUser(&_AtomicBond.TransactOpts, _address)
}

// AtomicBondMemoMim is a paid mutator transaction binding the contract method 0x0dde5052.
//
// Solidity: function atomicBondMemoMim(address _funder, address _minter, uint256 _bondIn, address[] _swapPath, uint256 _minSwapAmount, uint256 _maxBondPrice, uint256 _deadline) returns()
func (_AtomicBond *AtomicBondTransactor) AtomicBondMemoMim(opts *bind.TransactOpts, _funder common.Address, _minter common.Address, _bondIn *big.Int, _swapPath []common.Address, _minSwapAmount *big.Int, _maxBondPrice *big.Int, _deadline *big.Int) (*types.Transaction, error) {
	return _AtomicBond.contract.Transact(opts, "atomicBondMemoMim", _funder, _minter, _bondIn, _swapPath, _minSwapAmount, _maxBondPrice, _deadline)
}

// AtomicBondMemoMim is a paid mutator transaction binding the contract method 0x0dde5052.
//
// Solidity: function atomicBondMemoMim(address _funder, address _minter, uint256 _bondIn, address[] _swapPath, uint256 _minSwapAmount, uint256 _maxBondPrice, uint256 _deadline) returns()
func (_AtomicBond *AtomicBondSession) AtomicBondMemoMim(_funder common.Address, _minter common.Address, _bondIn *big.Int, _swapPath []common.Address, _minSwapAmount *big.Int, _maxBondPrice *big.Int, _deadline *big.Int) (*types.Transaction, error) {
	return _AtomicBond.Contract.AtomicBondMemoMim(&_AtomicBond.TransactOpts, _funder, _minter, _bondIn, _swapPath, _minSwapAmount, _maxBondPrice, _deadline)
}

// AtomicBondMemoMim is a paid mutator transaction binding the contract method 0x0dde5052.
//
// Solidity: function atomicBondMemoMim(address _funder, address _minter, uint256 _bondIn, address[] _swapPath, uint256 _minSwapAmount, uint256 _maxBondPrice, uint256 _deadline) returns()
func (_AtomicBond *AtomicBondTransactorSession) AtomicBondMemoMim(_funder common.Address, _minter common.Address, _bondIn *big.Int, _swapPath []common.Address, _minSwapAmount *big.Int, _maxBondPrice *big.Int, _deadline *big.Int) (*types.Transaction, error) {
	return _AtomicBond.Contract.AtomicBondMemoMim(&_AtomicBond.TransactOpts, _funder, _minter, _bondIn, _swapPath, _minSwapAmount, _maxBondPrice, _deadline)
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_AtomicBond *AtomicBondTransactor) Finalize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AtomicBond.contract.Transact(opts, "finalize")
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_AtomicBond *AtomicBondSession) Finalize() (*types.Transaction, error) {
	return _AtomicBond.Contract.Finalize(&_AtomicBond.TransactOpts)
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_AtomicBond *AtomicBondTransactorSession) Finalize() (*types.Transaction, error) {
	return _AtomicBond.Contract.Finalize(&_AtomicBond.TransactOpts)
}

// RecoverLostNetworkToken is a paid mutator transaction binding the contract method 0xfa4ea2be.
//
// Solidity: function recoverLostNetworkToken() returns()
func (_AtomicBond *AtomicBondTransactor) RecoverLostNetworkToken(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AtomicBond.contract.Transact(opts, "recoverLostNetworkToken")
}

// RecoverLostNetworkToken is a paid mutator transaction binding the contract method 0xfa4ea2be.
//
// Solidity: function recoverLostNetworkToken() returns()
func (_AtomicBond *AtomicBondSession) RecoverLostNetworkToken() (*types.Transaction, error) {
	return _AtomicBond.Contract.RecoverLostNetworkToken(&_AtomicBond.TransactOpts)
}

// RecoverLostNetworkToken is a paid mutator transaction binding the contract method 0xfa4ea2be.
//
// Solidity: function recoverLostNetworkToken() returns()
func (_AtomicBond *AtomicBondTransactorSession) RecoverLostNetworkToken() (*types.Transaction, error) {
	return _AtomicBond.Contract.RecoverLostNetworkToken(&_AtomicBond.TransactOpts)
}

// RecoverLostTokenERC20 is a paid mutator transaction binding the contract method 0xe3e6cd18.
//
// Solidity: function recoverLostTokenERC20(address _token, uint256 _amount) returns()
func (_AtomicBond *AtomicBondTransactor) RecoverLostTokenERC20(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _AtomicBond.contract.Transact(opts, "recoverLostTokenERC20", _token, _amount)
}

// RecoverLostTokenERC20 is a paid mutator transaction binding the contract method 0xe3e6cd18.
//
// Solidity: function recoverLostTokenERC20(address _token, uint256 _amount) returns()
func (_AtomicBond *AtomicBondSession) RecoverLostTokenERC20(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _AtomicBond.Contract.RecoverLostTokenERC20(&_AtomicBond.TransactOpts, _token, _amount)
}

// RecoverLostTokenERC20 is a paid mutator transaction binding the contract method 0xe3e6cd18.
//
// Solidity: function recoverLostTokenERC20(address _token, uint256 _amount) returns()
func (_AtomicBond *AtomicBondTransactorSession) RecoverLostTokenERC20(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _AtomicBond.Contract.RecoverLostTokenERC20(&_AtomicBond.TransactOpts, _token, _amount)
}

// Redeem is a paid mutator transaction binding the contract method 0x95a2251f.
//
// Solidity: function redeem(address _minter) returns()
func (_AtomicBond *AtomicBondTransactor) Redeem(opts *bind.TransactOpts, _minter common.Address) (*types.Transaction, error) {
	return _AtomicBond.contract.Transact(opts, "redeem", _minter)
}

// Redeem is a paid mutator transaction binding the contract method 0x95a2251f.
//
// Solidity: function redeem(address _minter) returns()
func (_AtomicBond *AtomicBondSession) Redeem(_minter common.Address) (*types.Transaction, error) {
	return _AtomicBond.Contract.Redeem(&_AtomicBond.TransactOpts, _minter)
}

// Redeem is a paid mutator transaction binding the contract method 0x95a2251f.
//
// Solidity: function redeem(address _minter) returns()
func (_AtomicBond *AtomicBondTransactorSession) Redeem(_minter common.Address) (*types.Transaction, error) {
	return _AtomicBond.Contract.Redeem(&_AtomicBond.TransactOpts, _minter)
}

// RevokeFunder is a paid mutator transaction binding the contract method 0x3cd80377.
//
// Solidity: function revokeFunder(address _address) returns()
func (_AtomicBond *AtomicBondTransactor) RevokeFunder(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _AtomicBond.contract.Transact(opts, "revokeFunder", _address)
}

// RevokeFunder is a paid mutator transaction binding the contract method 0x3cd80377.
//
// Solidity: function revokeFunder(address _address) returns()
func (_AtomicBond *AtomicBondSession) RevokeFunder(_address common.Address) (*types.Transaction, error) {
	return _AtomicBond.Contract.RevokeFunder(&_AtomicBond.TransactOpts, _address)
}

// RevokeFunder is a paid mutator transaction binding the contract method 0x3cd80377.
//
// Solidity: function revokeFunder(address _address) returns()
func (_AtomicBond *AtomicBondTransactorSession) RevokeFunder(_address common.Address) (*types.Transaction, error) {
	return _AtomicBond.Contract.RevokeFunder(&_AtomicBond.TransactOpts, _address)
}

// RevokeMinter is a paid mutator transaction binding the contract method 0xcfbd4885.
//
// Solidity: function revokeMinter(address _address) returns()
func (_AtomicBond *AtomicBondTransactor) RevokeMinter(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _AtomicBond.contract.Transact(opts, "revokeMinter", _address)
}

// RevokeMinter is a paid mutator transaction binding the contract method 0xcfbd4885.
//
// Solidity: function revokeMinter(address _address) returns()
func (_AtomicBond *AtomicBondSession) RevokeMinter(_address common.Address) (*types.Transaction, error) {
	return _AtomicBond.Contract.RevokeMinter(&_AtomicBond.TransactOpts, _address)
}

// RevokeMinter is a paid mutator transaction binding the contract method 0xcfbd4885.
//
// Solidity: function revokeMinter(address _address) returns()
func (_AtomicBond *AtomicBondTransactorSession) RevokeMinter(_address common.Address) (*types.Transaction, error) {
	return _AtomicBond.Contract.RevokeMinter(&_AtomicBond.TransactOpts, _address)
}

// RevokeUser is a paid mutator transaction binding the contract method 0xef097fc4.
//
// Solidity: function revokeUser(address _address) returns()
func (_AtomicBond *AtomicBondTransactor) RevokeUser(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _AtomicBond.contract.Transact(opts, "revokeUser", _address)
}

// RevokeUser is a paid mutator transaction binding the contract method 0xef097fc4.
//
// Solidity: function revokeUser(address _address) returns()
func (_AtomicBond *AtomicBondSession) RevokeUser(_address common.Address) (*types.Transaction, error) {
	return _AtomicBond.Contract.RevokeUser(&_AtomicBond.TransactOpts, _address)
}

// RevokeUser is a paid mutator transaction binding the contract method 0xef097fc4.
//
// Solidity: function revokeUser(address _address) returns()
func (_AtomicBond *AtomicBondTransactorSession) RevokeUser(_address common.Address) (*types.Transaction, error) {
	return _AtomicBond.Contract.RevokeUser(&_AtomicBond.TransactOpts, _address)
}
