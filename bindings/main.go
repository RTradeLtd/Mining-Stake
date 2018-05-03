// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// TokenLockupABI is the input ABI used to generate the binding from.
const TokenLockupABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"frozen\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_stakers\",\"type\":\"address[]\"},{\"name\":\"_payments\",\"type\":\"uint256[]\"}],\"name\":\"routeEthReward\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewards\",\"outputs\":[{\"name\":\"ethRewarded\",\"type\":\"uint256\"},{\"name\":\"rtcRewarded\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rtcHotWallet\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unlockStaking\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_staker\",\"type\":\"address\"}],\"name\":\"getNumStakes\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rtcCAD\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"moderators\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MINSTAKE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_stakeId\",\"type\":\"uint256\"}],\"name\":\"withdrawStake\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"kiloHashSecondPerOneCentCad\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"numStakes\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_contractAddress\",\"type\":\"address\"}],\"name\":\"setOracleContract\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"oracleContractAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_rtcCAD\",\"type\":\"uint256\"}],\"name\":\"updateRtcPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_stakers\",\"type\":\"address[]\"},{\"name\":\"_payments\",\"type\":\"uint256[]\"}],\"name\":\"routeRtcRewards\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_rtHotWallet\",\"type\":\"address\"}],\"name\":\"setRtHotWallet\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newAdmin\",\"type\":\"address\"}],\"name\":\"setAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_rtcToStake\",\"type\":\"uint256\"}],\"name\":\"calculateKhSecForNumRtc\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_rtcAddress\",\"type\":\"address\"}],\"name\":\"setRtI\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"signUpFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_staker\",\"type\":\"address\"}],\"name\":\"getRewardStruct\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rtI\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakers\",\"outputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"rtcStaked\",\"type\":\"uint256\"},{\"name\":\"deposit\",\"type\":\"uint256\"},{\"name\":\"khSec\",\"type\":\"uint256\"},{\"name\":\"depositDate\",\"type\":\"uint256\"},{\"name\":\"releaseDate\",\"type\":\"uint256\"},{\"name\":\"id\",\"type\":\"uint256\"},{\"name\":\"encryptedEmail\",\"type\":\"string\"},{\"name\":\"enabled\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"locked\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_rtcToStake\",\"type\":\"uint256\"},{\"name\":\"_durationInWeeksToStake\",\"type\":\"uint256\"},{\"name\":\"_encryptedEmail\",\"type\":\"string\"}],\"name\":\"depositStake\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_staker\",\"type\":\"address\"},{\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getStakerStruct\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"DEFAULTLOCKUPTIME\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stakerCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_staker\",\"type\":\"address\"},{\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getStakerEmailForStakeId\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"lockStaking\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_depositer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_weeksStaked\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_khSec\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"StakeDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_staker\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_stakeId\",\"type\":\"uint256\"}],\"name\":\"DepositWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_withdrawer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"EthWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_staker\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"RtcReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_staker\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"EthReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"result\",\"type\":\"string\"}],\"name\":\"NewOraclizeQuery\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"EthUsdPriceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"SignUpFeeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_admin\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_adminSet\",\"type\":\"bool\"}],\"name\":\"AdminSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_newOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_ownershipTransferred\",\"type\":\"bool\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// TokenLockupBin is the compiled bytecode used for deploying new contracts.
const TokenLockupBin = `60806040526008805461010060a860020a031916740994f9595d28429584bfb5fcbfea75b9c9ea2c240017905560008054600160a060020a0319908116600160a060020a033316908117835560018054909216179055611ae090819061006590396000f3006080604052600436106101b65763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663054f7d9c81146101b8578063059d2bf4146101e15780630700037d146102625780630bc0eadd1461029c5780630d376c80146102cd5780630de9b519146102e2578063120442121461031557806314d0f1ba1461032a5780631d0a0e8d1461034b57806325d5971f146103605780632996c25f146103785780633467e9e11461038d578063496ccd9b146103ae5780634bc60aca146103cf57806358243471146103e45780636198ce14146103fc5780636b41f4111461048a578063704b6c02146104ab5780637694a03d146104cc57806379867480146104e45780638da5cb5b146105055780639278b5871461051a578063b62682b61461052f578063be72ab5914610550578063c8b6cbf714610565578063cf3090121461064f578063d914c9d214610664578063de20e7e9146106c4578063df9ec0f41461071d578063dff6978714610732578063e4b7555114610747578063e527e56f146107e0578063f2fde38b146107f5578063f851a44014610816578063ffa1ad741461082b575b005b3480156101c457600080fd5b506101cd610840565b604080519115158252519081900360200190f35b604080516020600480358082013583810280860185019096528085526101cd95369593946024949385019291829185019084908082843750506040805187358901803560208181028481018201909552818452989b9a9989019892975090820195509350839250850190849080828437509497506108619650505050505050565b34801561026e57600080fd5b50610283600160a060020a0360043516610a16565b6040805192835260208301919091528051918290030190f35b3480156102a857600080fd5b506102b1610a2f565b60408051600160a060020a039092168252519081900360200190f35b3480156102d957600080fd5b506101cd610a3e565b3480156102ee57600080fd5b50610303600160a060020a0360043516610a97565b60408051918252519081900360200190f35b34801561032157600080fd5b50610303610ab2565b34801561033657600080fd5b506101cd600160a060020a0360043516610ab8565b34801561035757600080fd5b50610303610acd565b34801561036c57600080fd5b506101cd600435610ad9565b34801561038457600080fd5b50610303610d6f565b34801561039957600080fd5b50610303600160a060020a0360043516610d74565b3480156103ba57600080fd5b506101cd600160a060020a0360043516610d86565b3480156103db57600080fd5b506102b1610def565b3480156103f057600080fd5b506101cd600435610dfe565b34801561040857600080fd5b50604080516020600480358082013583810280860185019096528085526101cd95369593946024949385019291829185019084908082843750506040805187358901803560208181028481018201909552818452989b9a998901989297509082019550935083925085019084908082843750949750610e279650505050505050565b34801561049657600080fd5b506101cd600160a060020a036004351661105e565b3480156104b757600080fd5b506101cd600160a060020a03600435166110c7565b3480156104d857600080fd5b50610303600435611156565b3480156104f057600080fd5b506101cd600160a060020a0360043516611188565b34801561051157600080fd5b506102b16111f6565b34801561052657600080fd5b50610303611205565b34801561053b57600080fd5b50610283600160a060020a036004351661120b565b34801561055c57600080fd5b506102b161122e565b34801561057157600080fd5b50610589600160a060020a0360043516602435611242565b604051808a600160a060020a0316600160a060020a031681526020018981526020018881526020018781526020018681526020018581526020018481526020018060200183151515158152602001828103825284818151815260200191508051906020019080838360005b8381101561060c5781810151838201526020016105f4565b50505050905090810190601f1680156106395780820380516001836020036101000a031916815260200191505b509a505050505050505050505060405180910390f35b34801561065b57600080fd5b506101cd611344565b34801561067057600080fd5b50604080516020600460443581810135601f81018490048402850184019095528484526101cd94823594602480359536959460649492019190819084018382808284375094975061134d9650505050505050565b3480156106d057600080fd5b506106e8600160a060020a036004351660243561164e565b6040805196875260208701959095528585019390935260608501919091526080840152151560a0830152519081900360c00190f35b34801561072957600080fd5b506103036117be565b34801561073e57600080fd5b506103036117c5565b34801561075357600080fd5b5061076b600160a060020a03600435166024356117cb565b6040805160208082528351818301528351919283929083019185019080838360005b838110156107a557818101518382015260200161078d565b50505050905090810190601f1680156107d25780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156107ec57600080fd5b506101cd611896565b34801561080157600080fd5b506101cd600160a060020a03600435166118f1565b34801561082257600080fd5b506102b1611986565b34801561083757600080fd5b5061076b611995565b60015474010000000000000000000000000000000000000000900460ff1681565b600080548190819033600160a060020a0390811691161480610891575060015433600160a060020a039081169116145b151561089c57600080fd5b600034116108a957600080fd5b83518551146108b757600080fd5b600091505b8451821015610a0b5783828151811015156108d357fe5b90602001906020020151905061092481600a600088868151811015156108f557fe5b6020908102909101810151600160a060020a03168252810191909152604001600020549063ffffffff6119cc16565b600a6000878581518110151561093657fe5b6020908102909101810151600160a060020a031682528101919091526040016000205584517f196f95be2bd8e6aa49ac118195dec22c01e6adf2c34eaae9fefbb9efc8b78a799086908490811061098957fe5b602090810290910181015160408051600160a060020a03909216825291810184905281519081900390910190a184828151811015156109c457fe5b90602001906020020151600160a060020a03166108fc829081150290604051600060405180830381858888f193505050501515610a0057600080fd5b6001909101906108bc565b506001949350505050565b600a602052600090815260409020805460019091015482565b600554600160a060020a031681565b6000805433600160a060020a0390811691161480610a6a575060015433600160a060020a039081169116145b1515610a7557600080fd5b60085460ff161515610a8657600080fd5b506008805460ff1916905560015b90565b600160a060020a03166000908152600b602052604090205490565b60045481565b60026020526000908152604090205460ff1681565b6704fefa17b724000081565b60008033836009600083600160a060020a0316600160a060020a0316815260200190815260200160002081815481101515610b1057fe5b90600052602060002090600902016005015442111515610b2f57600080fd5b33600160a060020a0381166000908152600960205260409020805487919082908110610b5757fe5b600091825260209091206008600990920201015460ff161515610b7957600080fd5b600160a060020a0333166000908152600960205260408120805489908110610b9d57fe5b906000526020600020906009020160020154111515610bb857fe5b600160a060020a0333166000908152600960205260409020805488908110610bdc57fe5b600091825260208083206009928302016002015433600160a060020a031684529190526040822080549197509089908110610c1357fe5b600091825260208083206009928302016002019390935533600160a060020a0316825290915260408120805489908110610c4957fe5b6000918252602091829020600860099092020101805492151560ff199093169290921790915560408051600160a060020a0333168152918201879052818101899052517f7719804546c0185709e60c90d164447ff251a5ba29af0216faa921350f6bebf79181900360600190a1600854604080517fa9059cbb000000000000000000000000000000000000000000000000000000008152600160a060020a0333811660048301526024820189905291516101009093049091169163a9059cbb916044808201926020929091908290030181600087803b158015610d2b57600080fd5b505af1158015610d3f573d6000803e3d6000fd5b505050506040513d6020811015610d5557600080fd5b50511515610d6257600080fd5b5060019695505050505050565b606481565b600b6020526000908152604090205481565b6000805433600160a060020a0390811691161480610db2575060015433600160a060020a039081169116145b1515610dbd57600080fd5b5060038054600160a060020a03831673ffffffffffffffffffffffffffffffffffffffff199091161790556001919050565b600354600160a060020a031681565b6003546000903390600160a060020a03808316911614610e1d57600080fd5b5050600455600190565b600080548190819033600160a060020a0390811691161480610e57575060015433600160a060020a039081169116145b1515610e6257600080fd5b8351855114610e7057600080fd5b600091505b8451821015610a0b578382815181101515610e8c57fe5b906020019060200201519050610ee081600a60008886815181101515610eae57fe5b6020908102909101810151600160a060020a03168252810191909152604001600020600101549063ffffffff6119cc16565b600a60008785815181101515610ef257fe5b6020908102909101810151600160a060020a031682528101919091526040016000206001015584517fe0ba89edeae157ec385468cf95ff7ea61497f95bf3e0fe9637fa358aefdf7e2890869084908110610f4857fe5b602090810290910181015160408051600160a060020a03909216825291810184905281519081900390910190a1600860019054906101000a9004600160a060020a0316600160a060020a03166323b872dd338785815181101515610fa857fe5b6020908102909101810151604080517c010000000000000000000000000000000000000000000000000000000063ffffffff8716028152600160a060020a03948516600482015293909116602484015260448301869052516064808401938290030181600087803b15801561101c57600080fd5b505af1158015611030573d6000803e3d6000fd5b505050506040513d602081101561104657600080fd5b5051151561105357600080fd5b600190910190610e75565b6000805433600160a060020a039081169116148061108a575060015433600160a060020a039081169116145b151561109557600080fd5b5060058054600160a060020a03831673ffffffffffffffffffffffffffffffffffffffff199091161790556001919050565b6000805433600160a060020a039081169116146110e357600080fd5b600154600160a060020a03838116911614156110fe57600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03841690811782556040517fe68d2c359a771606c400cf8b87000cf5864010363d6a736e98f5047b7bbe18e990600090a3919050565b60008061116f60646004546119de90919063ffffffff16565b9050611181838263ffffffff6119f516565b9392505050565b6000805433600160a060020a03908116911614806111b4575060015433600160a060020a039081169116145b15156111bf57600080fd5b5060088054600160a060020a0383166101000274ffffffffffffffffffffffffffffffffffffffff00199091161790556001919050565b600054600160a060020a031690565b60065481565b600160a060020a03166000908152600a6020526040902080546001909101549091565b6008546101009004600160a060020a031681565b60096020528160005260406000208181548110151561125d57fe5b60009182526020918290206009909102018054600180830154600280850154600386015460048701546005880154600689015460078a01805460408051601f6000199c841615610100029c909c0190921698909804998a018d90048d0281018d01909752888752600160a060020a039099169c50959a509298919790969295909392918301828280156113315780601f1061130657610100808354040283529160200191611331565b820191906000526020600020905b81548152906001019060200180831161131457829003601f168201915b5050506008909301549192505060ff1689565b60085460ff1681565b60085460009081908190819060ff161561136657600080fd5b33600160a060020a031632600160a060020a031614151561138357fe5b6704fefa17b7240000871015801561139c575060048610155b15156113a757600080fd5b6004546113bb90606463ffffffff6119de16565b600160a060020a0333166000908152600b602052604090205490935091506113ea82600163ffffffff6119cc16565b600160a060020a0333166000908152600b6020526040902055611413878463ffffffff6119f516565b905061142d81670de0b6b3a764000063ffffffff6119de16565b90506009600033600160a060020a0316600160a060020a031681526020019081526020016000206101206040519081016040528033600160a060020a031681526020018981526020018981526020018381526020014281526020018862093a800242018152602001848152602001878152602001600115158152509080600181540180825580915050906001820390600052602060002090600902016000909192909190915060008201518160000160006101000a815481600160a060020a030219169083600160a060020a031602179055506020820151816001015560408201518160020155606082015181600301556080820151816004015560a0820151816005015560c0820151816006015560e0820151816007019080519060200190611558929190611a1c565b5061010091909101516008909101805460ff1916911515919091179055506040805133600160a060020a0316815260208101899052808201889052606081018390526080810184905290517f1a325385f16807e99fb688b597db78b00faee313dcf02e882dd16daab6fc3e1f9160a0908290030190a1600854600554604080517f23b872dd000000000000000000000000000000000000000000000000000000008152600160a060020a0333811660048301529283166024820152604481018b90529051610100909304909116916323b872dd916064808201926020929091908290030181600087803b158015610d2b57600080fd5b600160a060020a0382166000908152600960205260408120805482918291829182918291908890811061167d57fe5b6000918252602080832060099283020160010154600160a060020a038c16845291905260409091208054899081106116b157fe5b6000918252602080832060099283020160030154600160a060020a038d168452919052604090912080548a9081106116e557fe5b6000918252602080832060099283020160040154600160a060020a038e168452919052604090912080548b90811061171957fe5b6000918252602080832060099283020160050154600160a060020a038f168452919052604090912080548c90811061174d57fe5b906000526020600020906009020160060154600960008e600160a060020a0316600160a060020a031681526020019081526020016000208c81548110151561179157fe5b6000918252602090912060099091020160080154949d939c50919a509850965060ff909116945092505050565b6224ea0081565b60075481565b600160a060020a0382166000908152600960205260409020805460609190839081106117f357fe5b6000918252602091829020600760099092020101805460408051601f60026000196101006001871615020190941693909304928301859004850281018501909152818152928301828280156118895780601f1061185e57610100808354040283529160200191611889565b820191906000526020600020905b81548152906001019060200180831161186c57829003601f168201915b5050505050905092915050565b6000805433600160a060020a03908116911614806118c2575060015433600160a060020a039081169116145b15156118cd57600080fd5b60085460ff16156118dd57600080fd5b506008805460ff1916600190811790915590565b6000805433600160a060020a0390811691161461190d57600080fd5b600054600160a060020a038381169116141561192857600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03848116918217835560405160019333909216917f7fdc2a4b6eb39ec3363d710d188620bd1e97b3c434161f187b4d0dc0544faa5891a4919050565b600154600160a060020a031690565b60408051808201909152600981527f312e302e30626574610000000000000000000000000000000000000000000000602082015281565b60008282018381101561118157600080fd5b60008082848115156119ec57fe5b04949350505050565b6000828202831580611a115750828482811515611a0e57fe5b04145b151561118157600080fd5b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10611a5d57805160ff1916838001178555611a8a565b82800160010185558215611a8a579182015b82811115611a8a578251825591602001919060010190611a6f565b50611a96929150611a9a565b5090565b610a9491905b80821115611a965760008155600101611aa05600a165627a7a7230582041c0c6a610fcb3c794cb5b1dd970f86f8032b34bea28c06d96efc4279d8122a90029`

// DeployTokenLockup deploys a new Ethereum contract, binding an instance of TokenLockup to it.
func DeployTokenLockup(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TokenLockup, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenLockupABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TokenLockupBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TokenLockup{TokenLockupCaller: TokenLockupCaller{contract: contract}, TokenLockupTransactor: TokenLockupTransactor{contract: contract}, TokenLockupFilterer: TokenLockupFilterer{contract: contract}}, nil
}

// TokenLockup is an auto generated Go binding around an Ethereum contract.
type TokenLockup struct {
	TokenLockupCaller     // Read-only binding to the contract
	TokenLockupTransactor // Write-only binding to the contract
	TokenLockupFilterer   // Log filterer for contract events
}

// TokenLockupCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenLockupCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenLockupTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenLockupTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenLockupFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenLockupFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenLockupSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenLockupSession struct {
	Contract     *TokenLockup      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenLockupCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenLockupCallerSession struct {
	Contract *TokenLockupCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// TokenLockupTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenLockupTransactorSession struct {
	Contract     *TokenLockupTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// TokenLockupRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenLockupRaw struct {
	Contract *TokenLockup // Generic contract binding to access the raw methods on
}

// TokenLockupCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenLockupCallerRaw struct {
	Contract *TokenLockupCaller // Generic read-only contract binding to access the raw methods on
}

// TokenLockupTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenLockupTransactorRaw struct {
	Contract *TokenLockupTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenLockup creates a new instance of TokenLockup, bound to a specific deployed contract.
func NewTokenLockup(address common.Address, backend bind.ContractBackend) (*TokenLockup, error) {
	contract, err := bindTokenLockup(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenLockup{TokenLockupCaller: TokenLockupCaller{contract: contract}, TokenLockupTransactor: TokenLockupTransactor{contract: contract}, TokenLockupFilterer: TokenLockupFilterer{contract: contract}}, nil
}

// NewTokenLockupCaller creates a new read-only instance of TokenLockup, bound to a specific deployed contract.
func NewTokenLockupCaller(address common.Address, caller bind.ContractCaller) (*TokenLockupCaller, error) {
	contract, err := bindTokenLockup(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenLockupCaller{contract: contract}, nil
}

// NewTokenLockupTransactor creates a new write-only instance of TokenLockup, bound to a specific deployed contract.
func NewTokenLockupTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenLockupTransactor, error) {
	contract, err := bindTokenLockup(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenLockupTransactor{contract: contract}, nil
}

// NewTokenLockupFilterer creates a new log filterer instance of TokenLockup, bound to a specific deployed contract.
func NewTokenLockupFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenLockupFilterer, error) {
	contract, err := bindTokenLockup(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenLockupFilterer{contract: contract}, nil
}

// bindTokenLockup binds a generic wrapper to an already deployed contract.
func bindTokenLockup(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenLockupABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenLockup *TokenLockupRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TokenLockup.Contract.TokenLockupCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenLockup *TokenLockupRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenLockup.Contract.TokenLockupTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenLockup *TokenLockupRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenLockup.Contract.TokenLockupTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenLockup *TokenLockupCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TokenLockup.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenLockup *TokenLockupTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenLockup.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenLockup *TokenLockupTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenLockup.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTLOCKUPTIME is a free data retrieval call binding the contract method 0xdf9ec0f4.
//
// Solidity: function DEFAULTLOCKUPTIME() constant returns(uint256)
func (_TokenLockup *TokenLockupCaller) DEFAULTLOCKUPTIME(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TokenLockup.contract.Call(opts, out, "DEFAULTLOCKUPTIME")
	return *ret0, err
}

// DEFAULTLOCKUPTIME is a free data retrieval call binding the contract method 0xdf9ec0f4.
//
// Solidity: function DEFAULTLOCKUPTIME() constant returns(uint256)
func (_TokenLockup *TokenLockupSession) DEFAULTLOCKUPTIME() (*big.Int, error) {
	return _TokenLockup.Contract.DEFAULTLOCKUPTIME(&_TokenLockup.CallOpts)
}

// DEFAULTLOCKUPTIME is a free data retrieval call binding the contract method 0xdf9ec0f4.
//
// Solidity: function DEFAULTLOCKUPTIME() constant returns(uint256)
func (_TokenLockup *TokenLockupCallerSession) DEFAULTLOCKUPTIME() (*big.Int, error) {
	return _TokenLockup.Contract.DEFAULTLOCKUPTIME(&_TokenLockup.CallOpts)
}

// MINSTAKE is a free data retrieval call binding the contract method 0x1d0a0e8d.
//
// Solidity: function MINSTAKE() constant returns(uint256)
func (_TokenLockup *TokenLockupCaller) MINSTAKE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TokenLockup.contract.Call(opts, out, "MINSTAKE")
	return *ret0, err
}

// MINSTAKE is a free data retrieval call binding the contract method 0x1d0a0e8d.
//
// Solidity: function MINSTAKE() constant returns(uint256)
func (_TokenLockup *TokenLockupSession) MINSTAKE() (*big.Int, error) {
	return _TokenLockup.Contract.MINSTAKE(&_TokenLockup.CallOpts)
}

// MINSTAKE is a free data retrieval call binding the contract method 0x1d0a0e8d.
//
// Solidity: function MINSTAKE() constant returns(uint256)
func (_TokenLockup *TokenLockupCallerSession) MINSTAKE() (*big.Int, error) {
	return _TokenLockup.Contract.MINSTAKE(&_TokenLockup.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() constant returns(string)
func (_TokenLockup *TokenLockupCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _TokenLockup.contract.Call(opts, out, "VERSION")
	return *ret0, err
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() constant returns(string)
func (_TokenLockup *TokenLockupSession) VERSION() (string, error) {
	return _TokenLockup.Contract.VERSION(&_TokenLockup.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() constant returns(string)
func (_TokenLockup *TokenLockupCallerSession) VERSION() (string, error) {
	return _TokenLockup.Contract.VERSION(&_TokenLockup.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_TokenLockup *TokenLockupCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TokenLockup.contract.Call(opts, out, "admin")
	return *ret0, err
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_TokenLockup *TokenLockupSession) Admin() (common.Address, error) {
	return _TokenLockup.Contract.Admin(&_TokenLockup.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_TokenLockup *TokenLockupCallerSession) Admin() (common.Address, error) {
	return _TokenLockup.Contract.Admin(&_TokenLockup.CallOpts)
}

// CalculateKhSecForNumRtc is a free data retrieval call binding the contract method 0x7694a03d.
//
// Solidity: function calculateKhSecForNumRtc(_rtcToStake uint256) constant returns(uint256)
func (_TokenLockup *TokenLockupCaller) CalculateKhSecForNumRtc(opts *bind.CallOpts, _rtcToStake *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TokenLockup.contract.Call(opts, out, "calculateKhSecForNumRtc", _rtcToStake)
	return *ret0, err
}

// CalculateKhSecForNumRtc is a free data retrieval call binding the contract method 0x7694a03d.
//
// Solidity: function calculateKhSecForNumRtc(_rtcToStake uint256) constant returns(uint256)
func (_TokenLockup *TokenLockupSession) CalculateKhSecForNumRtc(_rtcToStake *big.Int) (*big.Int, error) {
	return _TokenLockup.Contract.CalculateKhSecForNumRtc(&_TokenLockup.CallOpts, _rtcToStake)
}

// CalculateKhSecForNumRtc is a free data retrieval call binding the contract method 0x7694a03d.
//
// Solidity: function calculateKhSecForNumRtc(_rtcToStake uint256) constant returns(uint256)
func (_TokenLockup *TokenLockupCallerSession) CalculateKhSecForNumRtc(_rtcToStake *big.Int) (*big.Int, error) {
	return _TokenLockup.Contract.CalculateKhSecForNumRtc(&_TokenLockup.CallOpts, _rtcToStake)
}

// Frozen is a free data retrieval call binding the contract method 0x054f7d9c.
//
// Solidity: function frozen() constant returns(bool)
func (_TokenLockup *TokenLockupCaller) Frozen(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TokenLockup.contract.Call(opts, out, "frozen")
	return *ret0, err
}

// Frozen is a free data retrieval call binding the contract method 0x054f7d9c.
//
// Solidity: function frozen() constant returns(bool)
func (_TokenLockup *TokenLockupSession) Frozen() (bool, error) {
	return _TokenLockup.Contract.Frozen(&_TokenLockup.CallOpts)
}

// Frozen is a free data retrieval call binding the contract method 0x054f7d9c.
//
// Solidity: function frozen() constant returns(bool)
func (_TokenLockup *TokenLockupCallerSession) Frozen() (bool, error) {
	return _TokenLockup.Contract.Frozen(&_TokenLockup.CallOpts)
}

// GetNumStakes is a free data retrieval call binding the contract method 0x0de9b519.
//
// Solidity: function getNumStakes(_staker address) constant returns(uint256)
func (_TokenLockup *TokenLockupCaller) GetNumStakes(opts *bind.CallOpts, _staker common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TokenLockup.contract.Call(opts, out, "getNumStakes", _staker)
	return *ret0, err
}

// GetNumStakes is a free data retrieval call binding the contract method 0x0de9b519.
//
// Solidity: function getNumStakes(_staker address) constant returns(uint256)
func (_TokenLockup *TokenLockupSession) GetNumStakes(_staker common.Address) (*big.Int, error) {
	return _TokenLockup.Contract.GetNumStakes(&_TokenLockup.CallOpts, _staker)
}

// GetNumStakes is a free data retrieval call binding the contract method 0x0de9b519.
//
// Solidity: function getNumStakes(_staker address) constant returns(uint256)
func (_TokenLockup *TokenLockupCallerSession) GetNumStakes(_staker common.Address) (*big.Int, error) {
	return _TokenLockup.Contract.GetNumStakes(&_TokenLockup.CallOpts, _staker)
}

// GetRewardStruct is a free data retrieval call binding the contract method 0xb62682b6.
//
// Solidity: function getRewardStruct(_staker address) constant returns(uint256, uint256)
func (_TokenLockup *TokenLockupCaller) GetRewardStruct(opts *bind.CallOpts, _staker common.Address) (*big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _TokenLockup.contract.Call(opts, out, "getRewardStruct", _staker)
	return *ret0, *ret1, err
}

// GetRewardStruct is a free data retrieval call binding the contract method 0xb62682b6.
//
// Solidity: function getRewardStruct(_staker address) constant returns(uint256, uint256)
func (_TokenLockup *TokenLockupSession) GetRewardStruct(_staker common.Address) (*big.Int, *big.Int, error) {
	return _TokenLockup.Contract.GetRewardStruct(&_TokenLockup.CallOpts, _staker)
}

// GetRewardStruct is a free data retrieval call binding the contract method 0xb62682b6.
//
// Solidity: function getRewardStruct(_staker address) constant returns(uint256, uint256)
func (_TokenLockup *TokenLockupCallerSession) GetRewardStruct(_staker common.Address) (*big.Int, *big.Int, error) {
	return _TokenLockup.Contract.GetRewardStruct(&_TokenLockup.CallOpts, _staker)
}

// GetStakerEmailForStakeId is a free data retrieval call binding the contract method 0xe4b75551.
//
// Solidity: function getStakerEmailForStakeId(_staker address, _id uint256) constant returns(string)
func (_TokenLockup *TokenLockupCaller) GetStakerEmailForStakeId(opts *bind.CallOpts, _staker common.Address, _id *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _TokenLockup.contract.Call(opts, out, "getStakerEmailForStakeId", _staker, _id)
	return *ret0, err
}

// GetStakerEmailForStakeId is a free data retrieval call binding the contract method 0xe4b75551.
//
// Solidity: function getStakerEmailForStakeId(_staker address, _id uint256) constant returns(string)
func (_TokenLockup *TokenLockupSession) GetStakerEmailForStakeId(_staker common.Address, _id *big.Int) (string, error) {
	return _TokenLockup.Contract.GetStakerEmailForStakeId(&_TokenLockup.CallOpts, _staker, _id)
}

// GetStakerEmailForStakeId is a free data retrieval call binding the contract method 0xe4b75551.
//
// Solidity: function getStakerEmailForStakeId(_staker address, _id uint256) constant returns(string)
func (_TokenLockup *TokenLockupCallerSession) GetStakerEmailForStakeId(_staker common.Address, _id *big.Int) (string, error) {
	return _TokenLockup.Contract.GetStakerEmailForStakeId(&_TokenLockup.CallOpts, _staker, _id)
}

// GetStakerStruct is a free data retrieval call binding the contract method 0xde20e7e9.
//
// Solidity: function getStakerStruct(_staker address, _id uint256) constant returns(uint256, uint256, uint256, uint256, uint256, bool)
func (_TokenLockup *TokenLockupCaller) GetStakerStruct(opts *bind.CallOpts, _staker common.Address, _id *big.Int) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, bool, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
		ret3 = new(*big.Int)
		ret4 = new(*big.Int)
		ret5 = new(bool)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
	}
	err := _TokenLockup.contract.Call(opts, out, "getStakerStruct", _staker, _id)
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, err
}

// GetStakerStruct is a free data retrieval call binding the contract method 0xde20e7e9.
//
// Solidity: function getStakerStruct(_staker address, _id uint256) constant returns(uint256, uint256, uint256, uint256, uint256, bool)
func (_TokenLockup *TokenLockupSession) GetStakerStruct(_staker common.Address, _id *big.Int) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, bool, error) {
	return _TokenLockup.Contract.GetStakerStruct(&_TokenLockup.CallOpts, _staker, _id)
}

// GetStakerStruct is a free data retrieval call binding the contract method 0xde20e7e9.
//
// Solidity: function getStakerStruct(_staker address, _id uint256) constant returns(uint256, uint256, uint256, uint256, uint256, bool)
func (_TokenLockup *TokenLockupCallerSession) GetStakerStruct(_staker common.Address, _id *big.Int) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, bool, error) {
	return _TokenLockup.Contract.GetStakerStruct(&_TokenLockup.CallOpts, _staker, _id)
}

// KiloHashSecondPerOneCentCad is a free data retrieval call binding the contract method 0x2996c25f.
//
// Solidity: function kiloHashSecondPerOneCentCad() constant returns(uint256)
func (_TokenLockup *TokenLockupCaller) KiloHashSecondPerOneCentCad(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TokenLockup.contract.Call(opts, out, "kiloHashSecondPerOneCentCad")
	return *ret0, err
}

// KiloHashSecondPerOneCentCad is a free data retrieval call binding the contract method 0x2996c25f.
//
// Solidity: function kiloHashSecondPerOneCentCad() constant returns(uint256)
func (_TokenLockup *TokenLockupSession) KiloHashSecondPerOneCentCad() (*big.Int, error) {
	return _TokenLockup.Contract.KiloHashSecondPerOneCentCad(&_TokenLockup.CallOpts)
}

// KiloHashSecondPerOneCentCad is a free data retrieval call binding the contract method 0x2996c25f.
//
// Solidity: function kiloHashSecondPerOneCentCad() constant returns(uint256)
func (_TokenLockup *TokenLockupCallerSession) KiloHashSecondPerOneCentCad() (*big.Int, error) {
	return _TokenLockup.Contract.KiloHashSecondPerOneCentCad(&_TokenLockup.CallOpts)
}

// Locked is a free data retrieval call binding the contract method 0xcf309012.
//
// Solidity: function locked() constant returns(bool)
func (_TokenLockup *TokenLockupCaller) Locked(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TokenLockup.contract.Call(opts, out, "locked")
	return *ret0, err
}

// Locked is a free data retrieval call binding the contract method 0xcf309012.
//
// Solidity: function locked() constant returns(bool)
func (_TokenLockup *TokenLockupSession) Locked() (bool, error) {
	return _TokenLockup.Contract.Locked(&_TokenLockup.CallOpts)
}

// Locked is a free data retrieval call binding the contract method 0xcf309012.
//
// Solidity: function locked() constant returns(bool)
func (_TokenLockup *TokenLockupCallerSession) Locked() (bool, error) {
	return _TokenLockup.Contract.Locked(&_TokenLockup.CallOpts)
}

// Moderators is a free data retrieval call binding the contract method 0x14d0f1ba.
//
// Solidity: function moderators( address) constant returns(bool)
func (_TokenLockup *TokenLockupCaller) Moderators(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TokenLockup.contract.Call(opts, out, "moderators", arg0)
	return *ret0, err
}

// Moderators is a free data retrieval call binding the contract method 0x14d0f1ba.
//
// Solidity: function moderators( address) constant returns(bool)
func (_TokenLockup *TokenLockupSession) Moderators(arg0 common.Address) (bool, error) {
	return _TokenLockup.Contract.Moderators(&_TokenLockup.CallOpts, arg0)
}

// Moderators is a free data retrieval call binding the contract method 0x14d0f1ba.
//
// Solidity: function moderators( address) constant returns(bool)
func (_TokenLockup *TokenLockupCallerSession) Moderators(arg0 common.Address) (bool, error) {
	return _TokenLockup.Contract.Moderators(&_TokenLockup.CallOpts, arg0)
}

// NumStakes is a free data retrieval call binding the contract method 0x3467e9e1.
//
// Solidity: function numStakes( address) constant returns(uint256)
func (_TokenLockup *TokenLockupCaller) NumStakes(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TokenLockup.contract.Call(opts, out, "numStakes", arg0)
	return *ret0, err
}

// NumStakes is a free data retrieval call binding the contract method 0x3467e9e1.
//
// Solidity: function numStakes( address) constant returns(uint256)
func (_TokenLockup *TokenLockupSession) NumStakes(arg0 common.Address) (*big.Int, error) {
	return _TokenLockup.Contract.NumStakes(&_TokenLockup.CallOpts, arg0)
}

// NumStakes is a free data retrieval call binding the contract method 0x3467e9e1.
//
// Solidity: function numStakes( address) constant returns(uint256)
func (_TokenLockup *TokenLockupCallerSession) NumStakes(arg0 common.Address) (*big.Int, error) {
	return _TokenLockup.Contract.NumStakes(&_TokenLockup.CallOpts, arg0)
}

// OracleContractAddress is a free data retrieval call binding the contract method 0x4bc60aca.
//
// Solidity: function oracleContractAddress() constant returns(address)
func (_TokenLockup *TokenLockupCaller) OracleContractAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TokenLockup.contract.Call(opts, out, "oracleContractAddress")
	return *ret0, err
}

// OracleContractAddress is a free data retrieval call binding the contract method 0x4bc60aca.
//
// Solidity: function oracleContractAddress() constant returns(address)
func (_TokenLockup *TokenLockupSession) OracleContractAddress() (common.Address, error) {
	return _TokenLockup.Contract.OracleContractAddress(&_TokenLockup.CallOpts)
}

// OracleContractAddress is a free data retrieval call binding the contract method 0x4bc60aca.
//
// Solidity: function oracleContractAddress() constant returns(address)
func (_TokenLockup *TokenLockupCallerSession) OracleContractAddress() (common.Address, error) {
	return _TokenLockup.Contract.OracleContractAddress(&_TokenLockup.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_TokenLockup *TokenLockupCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TokenLockup.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_TokenLockup *TokenLockupSession) Owner() (common.Address, error) {
	return _TokenLockup.Contract.Owner(&_TokenLockup.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_TokenLockup *TokenLockupCallerSession) Owner() (common.Address, error) {
	return _TokenLockup.Contract.Owner(&_TokenLockup.CallOpts)
}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards( address) constant returns(ethRewarded uint256, rtcRewarded uint256)
func (_TokenLockup *TokenLockupCaller) Rewards(opts *bind.CallOpts, arg0 common.Address) (struct {
	EthRewarded *big.Int
	RtcRewarded *big.Int
}, error) {
	ret := new(struct {
		EthRewarded *big.Int
		RtcRewarded *big.Int
	})
	out := ret
	err := _TokenLockup.contract.Call(opts, out, "rewards", arg0)
	return *ret, err
}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards( address) constant returns(ethRewarded uint256, rtcRewarded uint256)
func (_TokenLockup *TokenLockupSession) Rewards(arg0 common.Address) (struct {
	EthRewarded *big.Int
	RtcRewarded *big.Int
}, error) {
	return _TokenLockup.Contract.Rewards(&_TokenLockup.CallOpts, arg0)
}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards( address) constant returns(ethRewarded uint256, rtcRewarded uint256)
func (_TokenLockup *TokenLockupCallerSession) Rewards(arg0 common.Address) (struct {
	EthRewarded *big.Int
	RtcRewarded *big.Int
}, error) {
	return _TokenLockup.Contract.Rewards(&_TokenLockup.CallOpts, arg0)
}

// RtI is a free data retrieval call binding the contract method 0xbe72ab59.
//
// Solidity: function rtI() constant returns(address)
func (_TokenLockup *TokenLockupCaller) RtI(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TokenLockup.contract.Call(opts, out, "rtI")
	return *ret0, err
}

// RtI is a free data retrieval call binding the contract method 0xbe72ab59.
//
// Solidity: function rtI() constant returns(address)
func (_TokenLockup *TokenLockupSession) RtI() (common.Address, error) {
	return _TokenLockup.Contract.RtI(&_TokenLockup.CallOpts)
}

// RtI is a free data retrieval call binding the contract method 0xbe72ab59.
//
// Solidity: function rtI() constant returns(address)
func (_TokenLockup *TokenLockupCallerSession) RtI() (common.Address, error) {
	return _TokenLockup.Contract.RtI(&_TokenLockup.CallOpts)
}

// RtcCAD is a free data retrieval call binding the contract method 0x12044212.
//
// Solidity: function rtcCAD() constant returns(uint256)
func (_TokenLockup *TokenLockupCaller) RtcCAD(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TokenLockup.contract.Call(opts, out, "rtcCAD")
	return *ret0, err
}

// RtcCAD is a free data retrieval call binding the contract method 0x12044212.
//
// Solidity: function rtcCAD() constant returns(uint256)
func (_TokenLockup *TokenLockupSession) RtcCAD() (*big.Int, error) {
	return _TokenLockup.Contract.RtcCAD(&_TokenLockup.CallOpts)
}

// RtcCAD is a free data retrieval call binding the contract method 0x12044212.
//
// Solidity: function rtcCAD() constant returns(uint256)
func (_TokenLockup *TokenLockupCallerSession) RtcCAD() (*big.Int, error) {
	return _TokenLockup.Contract.RtcCAD(&_TokenLockup.CallOpts)
}

// RtcHotWallet is a free data retrieval call binding the contract method 0x0bc0eadd.
//
// Solidity: function rtcHotWallet() constant returns(address)
func (_TokenLockup *TokenLockupCaller) RtcHotWallet(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TokenLockup.contract.Call(opts, out, "rtcHotWallet")
	return *ret0, err
}

// RtcHotWallet is a free data retrieval call binding the contract method 0x0bc0eadd.
//
// Solidity: function rtcHotWallet() constant returns(address)
func (_TokenLockup *TokenLockupSession) RtcHotWallet() (common.Address, error) {
	return _TokenLockup.Contract.RtcHotWallet(&_TokenLockup.CallOpts)
}

// RtcHotWallet is a free data retrieval call binding the contract method 0x0bc0eadd.
//
// Solidity: function rtcHotWallet() constant returns(address)
func (_TokenLockup *TokenLockupCallerSession) RtcHotWallet() (common.Address, error) {
	return _TokenLockup.Contract.RtcHotWallet(&_TokenLockup.CallOpts)
}

// SignUpFee is a free data retrieval call binding the contract method 0x9278b587.
//
// Solidity: function signUpFee() constant returns(uint256)
func (_TokenLockup *TokenLockupCaller) SignUpFee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TokenLockup.contract.Call(opts, out, "signUpFee")
	return *ret0, err
}

// SignUpFee is a free data retrieval call binding the contract method 0x9278b587.
//
// Solidity: function signUpFee() constant returns(uint256)
func (_TokenLockup *TokenLockupSession) SignUpFee() (*big.Int, error) {
	return _TokenLockup.Contract.SignUpFee(&_TokenLockup.CallOpts)
}

// SignUpFee is a free data retrieval call binding the contract method 0x9278b587.
//
// Solidity: function signUpFee() constant returns(uint256)
func (_TokenLockup *TokenLockupCallerSession) SignUpFee() (*big.Int, error) {
	return _TokenLockup.Contract.SignUpFee(&_TokenLockup.CallOpts)
}

// StakerCount is a free data retrieval call binding the contract method 0xdff69787.
//
// Solidity: function stakerCount() constant returns(uint256)
func (_TokenLockup *TokenLockupCaller) StakerCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TokenLockup.contract.Call(opts, out, "stakerCount")
	return *ret0, err
}

// StakerCount is a free data retrieval call binding the contract method 0xdff69787.
//
// Solidity: function stakerCount() constant returns(uint256)
func (_TokenLockup *TokenLockupSession) StakerCount() (*big.Int, error) {
	return _TokenLockup.Contract.StakerCount(&_TokenLockup.CallOpts)
}

// StakerCount is a free data retrieval call binding the contract method 0xdff69787.
//
// Solidity: function stakerCount() constant returns(uint256)
func (_TokenLockup *TokenLockupCallerSession) StakerCount() (*big.Int, error) {
	return _TokenLockup.Contract.StakerCount(&_TokenLockup.CallOpts)
}

// Stakers is a free data retrieval call binding the contract method 0xc8b6cbf7.
//
// Solidity: function stakers( address,  uint256) constant returns(addr address, rtcStaked uint256, deposit uint256, khSec uint256, depositDate uint256, releaseDate uint256, id uint256, encryptedEmail string, enabled bool)
func (_TokenLockup *TokenLockupCaller) Stakers(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	Addr           common.Address
	RtcStaked      *big.Int
	Deposit        *big.Int
	KhSec          *big.Int
	DepositDate    *big.Int
	ReleaseDate    *big.Int
	Id             *big.Int
	EncryptedEmail string
	Enabled        bool
}, error) {
	ret := new(struct {
		Addr           common.Address
		RtcStaked      *big.Int
		Deposit        *big.Int
		KhSec          *big.Int
		DepositDate    *big.Int
		ReleaseDate    *big.Int
		Id             *big.Int
		EncryptedEmail string
		Enabled        bool
	})
	out := ret
	err := _TokenLockup.contract.Call(opts, out, "stakers", arg0, arg1)
	return *ret, err
}

// Stakers is a free data retrieval call binding the contract method 0xc8b6cbf7.
//
// Solidity: function stakers( address,  uint256) constant returns(addr address, rtcStaked uint256, deposit uint256, khSec uint256, depositDate uint256, releaseDate uint256, id uint256, encryptedEmail string, enabled bool)
func (_TokenLockup *TokenLockupSession) Stakers(arg0 common.Address, arg1 *big.Int) (struct {
	Addr           common.Address
	RtcStaked      *big.Int
	Deposit        *big.Int
	KhSec          *big.Int
	DepositDate    *big.Int
	ReleaseDate    *big.Int
	Id             *big.Int
	EncryptedEmail string
	Enabled        bool
}, error) {
	return _TokenLockup.Contract.Stakers(&_TokenLockup.CallOpts, arg0, arg1)
}

// Stakers is a free data retrieval call binding the contract method 0xc8b6cbf7.
//
// Solidity: function stakers( address,  uint256) constant returns(addr address, rtcStaked uint256, deposit uint256, khSec uint256, depositDate uint256, releaseDate uint256, id uint256, encryptedEmail string, enabled bool)
func (_TokenLockup *TokenLockupCallerSession) Stakers(arg0 common.Address, arg1 *big.Int) (struct {
	Addr           common.Address
	RtcStaked      *big.Int
	Deposit        *big.Int
	KhSec          *big.Int
	DepositDate    *big.Int
	ReleaseDate    *big.Int
	Id             *big.Int
	EncryptedEmail string
	Enabled        bool
}, error) {
	return _TokenLockup.Contract.Stakers(&_TokenLockup.CallOpts, arg0, arg1)
}

// DepositStake is a paid mutator transaction binding the contract method 0xd914c9d2.
//
// Solidity: function depositStake(_rtcToStake uint256, _durationInWeeksToStake uint256, _encryptedEmail string) returns(bool)
func (_TokenLockup *TokenLockupTransactor) DepositStake(opts *bind.TransactOpts, _rtcToStake *big.Int, _durationInWeeksToStake *big.Int, _encryptedEmail string) (*types.Transaction, error) {
	return _TokenLockup.contract.Transact(opts, "depositStake", _rtcToStake, _durationInWeeksToStake, _encryptedEmail)
}

// DepositStake is a paid mutator transaction binding the contract method 0xd914c9d2.
//
// Solidity: function depositStake(_rtcToStake uint256, _durationInWeeksToStake uint256, _encryptedEmail string) returns(bool)
func (_TokenLockup *TokenLockupSession) DepositStake(_rtcToStake *big.Int, _durationInWeeksToStake *big.Int, _encryptedEmail string) (*types.Transaction, error) {
	return _TokenLockup.Contract.DepositStake(&_TokenLockup.TransactOpts, _rtcToStake, _durationInWeeksToStake, _encryptedEmail)
}

// DepositStake is a paid mutator transaction binding the contract method 0xd914c9d2.
//
// Solidity: function depositStake(_rtcToStake uint256, _durationInWeeksToStake uint256, _encryptedEmail string) returns(bool)
func (_TokenLockup *TokenLockupTransactorSession) DepositStake(_rtcToStake *big.Int, _durationInWeeksToStake *big.Int, _encryptedEmail string) (*types.Transaction, error) {
	return _TokenLockup.Contract.DepositStake(&_TokenLockup.TransactOpts, _rtcToStake, _durationInWeeksToStake, _encryptedEmail)
}

// LockStaking is a paid mutator transaction binding the contract method 0xe527e56f.
//
// Solidity: function lockStaking() returns(bool)
func (_TokenLockup *TokenLockupTransactor) LockStaking(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenLockup.contract.Transact(opts, "lockStaking")
}

// LockStaking is a paid mutator transaction binding the contract method 0xe527e56f.
//
// Solidity: function lockStaking() returns(bool)
func (_TokenLockup *TokenLockupSession) LockStaking() (*types.Transaction, error) {
	return _TokenLockup.Contract.LockStaking(&_TokenLockup.TransactOpts)
}

// LockStaking is a paid mutator transaction binding the contract method 0xe527e56f.
//
// Solidity: function lockStaking() returns(bool)
func (_TokenLockup *TokenLockupTransactorSession) LockStaking() (*types.Transaction, error) {
	return _TokenLockup.Contract.LockStaking(&_TokenLockup.TransactOpts)
}

// RouteEthReward is a paid mutator transaction binding the contract method 0x059d2bf4.
//
// Solidity: function routeEthReward(_stakers address[], _payments uint256[]) returns(bool)
func (_TokenLockup *TokenLockupTransactor) RouteEthReward(opts *bind.TransactOpts, _stakers []common.Address, _payments []*big.Int) (*types.Transaction, error) {
	return _TokenLockup.contract.Transact(opts, "routeEthReward", _stakers, _payments)
}

// RouteEthReward is a paid mutator transaction binding the contract method 0x059d2bf4.
//
// Solidity: function routeEthReward(_stakers address[], _payments uint256[]) returns(bool)
func (_TokenLockup *TokenLockupSession) RouteEthReward(_stakers []common.Address, _payments []*big.Int) (*types.Transaction, error) {
	return _TokenLockup.Contract.RouteEthReward(&_TokenLockup.TransactOpts, _stakers, _payments)
}

// RouteEthReward is a paid mutator transaction binding the contract method 0x059d2bf4.
//
// Solidity: function routeEthReward(_stakers address[], _payments uint256[]) returns(bool)
func (_TokenLockup *TokenLockupTransactorSession) RouteEthReward(_stakers []common.Address, _payments []*big.Int) (*types.Transaction, error) {
	return _TokenLockup.Contract.RouteEthReward(&_TokenLockup.TransactOpts, _stakers, _payments)
}

// RouteRtcRewards is a paid mutator transaction binding the contract method 0x6198ce14.
//
// Solidity: function routeRtcRewards(_stakers address[], _payments uint256[]) returns(bool)
func (_TokenLockup *TokenLockupTransactor) RouteRtcRewards(opts *bind.TransactOpts, _stakers []common.Address, _payments []*big.Int) (*types.Transaction, error) {
	return _TokenLockup.contract.Transact(opts, "routeRtcRewards", _stakers, _payments)
}

// RouteRtcRewards is a paid mutator transaction binding the contract method 0x6198ce14.
//
// Solidity: function routeRtcRewards(_stakers address[], _payments uint256[]) returns(bool)
func (_TokenLockup *TokenLockupSession) RouteRtcRewards(_stakers []common.Address, _payments []*big.Int) (*types.Transaction, error) {
	return _TokenLockup.Contract.RouteRtcRewards(&_TokenLockup.TransactOpts, _stakers, _payments)
}

// RouteRtcRewards is a paid mutator transaction binding the contract method 0x6198ce14.
//
// Solidity: function routeRtcRewards(_stakers address[], _payments uint256[]) returns(bool)
func (_TokenLockup *TokenLockupTransactorSession) RouteRtcRewards(_stakers []common.Address, _payments []*big.Int) (*types.Transaction, error) {
	return _TokenLockup.Contract.RouteRtcRewards(&_TokenLockup.TransactOpts, _stakers, _payments)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(_newAdmin address) returns(bool)
func (_TokenLockup *TokenLockupTransactor) SetAdmin(opts *bind.TransactOpts, _newAdmin common.Address) (*types.Transaction, error) {
	return _TokenLockup.contract.Transact(opts, "setAdmin", _newAdmin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(_newAdmin address) returns(bool)
func (_TokenLockup *TokenLockupSession) SetAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _TokenLockup.Contract.SetAdmin(&_TokenLockup.TransactOpts, _newAdmin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(_newAdmin address) returns(bool)
func (_TokenLockup *TokenLockupTransactorSession) SetAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _TokenLockup.Contract.SetAdmin(&_TokenLockup.TransactOpts, _newAdmin)
}

// SetOracleContract is a paid mutator transaction binding the contract method 0x496ccd9b.
//
// Solidity: function setOracleContract(_contractAddress address) returns(bool)
func (_TokenLockup *TokenLockupTransactor) SetOracleContract(opts *bind.TransactOpts, _contractAddress common.Address) (*types.Transaction, error) {
	return _TokenLockup.contract.Transact(opts, "setOracleContract", _contractAddress)
}

// SetOracleContract is a paid mutator transaction binding the contract method 0x496ccd9b.
//
// Solidity: function setOracleContract(_contractAddress address) returns(bool)
func (_TokenLockup *TokenLockupSession) SetOracleContract(_contractAddress common.Address) (*types.Transaction, error) {
	return _TokenLockup.Contract.SetOracleContract(&_TokenLockup.TransactOpts, _contractAddress)
}

// SetOracleContract is a paid mutator transaction binding the contract method 0x496ccd9b.
//
// Solidity: function setOracleContract(_contractAddress address) returns(bool)
func (_TokenLockup *TokenLockupTransactorSession) SetOracleContract(_contractAddress common.Address) (*types.Transaction, error) {
	return _TokenLockup.Contract.SetOracleContract(&_TokenLockup.TransactOpts, _contractAddress)
}

// SetRtHotWallet is a paid mutator transaction binding the contract method 0x6b41f411.
//
// Solidity: function setRtHotWallet(_rtHotWallet address) returns(bool)
func (_TokenLockup *TokenLockupTransactor) SetRtHotWallet(opts *bind.TransactOpts, _rtHotWallet common.Address) (*types.Transaction, error) {
	return _TokenLockup.contract.Transact(opts, "setRtHotWallet", _rtHotWallet)
}

// SetRtHotWallet is a paid mutator transaction binding the contract method 0x6b41f411.
//
// Solidity: function setRtHotWallet(_rtHotWallet address) returns(bool)
func (_TokenLockup *TokenLockupSession) SetRtHotWallet(_rtHotWallet common.Address) (*types.Transaction, error) {
	return _TokenLockup.Contract.SetRtHotWallet(&_TokenLockup.TransactOpts, _rtHotWallet)
}

// SetRtHotWallet is a paid mutator transaction binding the contract method 0x6b41f411.
//
// Solidity: function setRtHotWallet(_rtHotWallet address) returns(bool)
func (_TokenLockup *TokenLockupTransactorSession) SetRtHotWallet(_rtHotWallet common.Address) (*types.Transaction, error) {
	return _TokenLockup.Contract.SetRtHotWallet(&_TokenLockup.TransactOpts, _rtHotWallet)
}

// SetRtI is a paid mutator transaction binding the contract method 0x79867480.
//
// Solidity: function setRtI(_rtcAddress address) returns(bool)
func (_TokenLockup *TokenLockupTransactor) SetRtI(opts *bind.TransactOpts, _rtcAddress common.Address) (*types.Transaction, error) {
	return _TokenLockup.contract.Transact(opts, "setRtI", _rtcAddress)
}

// SetRtI is a paid mutator transaction binding the contract method 0x79867480.
//
// Solidity: function setRtI(_rtcAddress address) returns(bool)
func (_TokenLockup *TokenLockupSession) SetRtI(_rtcAddress common.Address) (*types.Transaction, error) {
	return _TokenLockup.Contract.SetRtI(&_TokenLockup.TransactOpts, _rtcAddress)
}

// SetRtI is a paid mutator transaction binding the contract method 0x79867480.
//
// Solidity: function setRtI(_rtcAddress address) returns(bool)
func (_TokenLockup *TokenLockupTransactorSession) SetRtI(_rtcAddress common.Address) (*types.Transaction, error) {
	return _TokenLockup.Contract.SetRtI(&_TokenLockup.TransactOpts, _rtcAddress)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns(bool)
func (_TokenLockup *TokenLockupTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _TokenLockup.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns(bool)
func (_TokenLockup *TokenLockupSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _TokenLockup.Contract.TransferOwnership(&_TokenLockup.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns(bool)
func (_TokenLockup *TokenLockupTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _TokenLockup.Contract.TransferOwnership(&_TokenLockup.TransactOpts, _newOwner)
}

// UnlockStaking is a paid mutator transaction binding the contract method 0x0d376c80.
//
// Solidity: function unlockStaking() returns(bool)
func (_TokenLockup *TokenLockupTransactor) UnlockStaking(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenLockup.contract.Transact(opts, "unlockStaking")
}

// UnlockStaking is a paid mutator transaction binding the contract method 0x0d376c80.
//
// Solidity: function unlockStaking() returns(bool)
func (_TokenLockup *TokenLockupSession) UnlockStaking() (*types.Transaction, error) {
	return _TokenLockup.Contract.UnlockStaking(&_TokenLockup.TransactOpts)
}

// UnlockStaking is a paid mutator transaction binding the contract method 0x0d376c80.
//
// Solidity: function unlockStaking() returns(bool)
func (_TokenLockup *TokenLockupTransactorSession) UnlockStaking() (*types.Transaction, error) {
	return _TokenLockup.Contract.UnlockStaking(&_TokenLockup.TransactOpts)
}

// UpdateRtcPrice is a paid mutator transaction binding the contract method 0x58243471.
//
// Solidity: function updateRtcPrice(_rtcCAD uint256) returns(bool)
func (_TokenLockup *TokenLockupTransactor) UpdateRtcPrice(opts *bind.TransactOpts, _rtcCAD *big.Int) (*types.Transaction, error) {
	return _TokenLockup.contract.Transact(opts, "updateRtcPrice", _rtcCAD)
}

// UpdateRtcPrice is a paid mutator transaction binding the contract method 0x58243471.
//
// Solidity: function updateRtcPrice(_rtcCAD uint256) returns(bool)
func (_TokenLockup *TokenLockupSession) UpdateRtcPrice(_rtcCAD *big.Int) (*types.Transaction, error) {
	return _TokenLockup.Contract.UpdateRtcPrice(&_TokenLockup.TransactOpts, _rtcCAD)
}

// UpdateRtcPrice is a paid mutator transaction binding the contract method 0x58243471.
//
// Solidity: function updateRtcPrice(_rtcCAD uint256) returns(bool)
func (_TokenLockup *TokenLockupTransactorSession) UpdateRtcPrice(_rtcCAD *big.Int) (*types.Transaction, error) {
	return _TokenLockup.Contract.UpdateRtcPrice(&_TokenLockup.TransactOpts, _rtcCAD)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0x25d5971f.
//
// Solidity: function withdrawStake(_stakeId uint256) returns(bool)
func (_TokenLockup *TokenLockupTransactor) WithdrawStake(opts *bind.TransactOpts, _stakeId *big.Int) (*types.Transaction, error) {
	return _TokenLockup.contract.Transact(opts, "withdrawStake", _stakeId)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0x25d5971f.
//
// Solidity: function withdrawStake(_stakeId uint256) returns(bool)
func (_TokenLockup *TokenLockupSession) WithdrawStake(_stakeId *big.Int) (*types.Transaction, error) {
	return _TokenLockup.Contract.WithdrawStake(&_TokenLockup.TransactOpts, _stakeId)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0x25d5971f.
//
// Solidity: function withdrawStake(_stakeId uint256) returns(bool)
func (_TokenLockup *TokenLockupTransactorSession) WithdrawStake(_stakeId *big.Int) (*types.Transaction, error) {
	return _TokenLockup.Contract.WithdrawStake(&_TokenLockup.TransactOpts, _stakeId)
}

// TokenLockupAdminSetIterator is returned from FilterAdminSet and is used to iterate over the raw logs and unpacked data for AdminSet events raised by the TokenLockup contract.
type TokenLockupAdminSetIterator struct {
	Event *TokenLockupAdminSet // Event containing the contract specifics and raw log

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
func (it *TokenLockupAdminSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenLockupAdminSet)
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
		it.Event = new(TokenLockupAdminSet)
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
func (it *TokenLockupAdminSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenLockupAdminSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenLockupAdminSet represents a AdminSet event raised by the TokenLockup contract.
type TokenLockupAdminSet struct {
	Admin    common.Address
	AdminSet bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAdminSet is a free log retrieval operation binding the contract event 0xe68d2c359a771606c400cf8b87000cf5864010363d6a736e98f5047b7bbe18e9.
//
// Solidity: event AdminSet(_admin indexed address, _adminSet indexed bool)
func (_TokenLockup *TokenLockupFilterer) FilterAdminSet(opts *bind.FilterOpts, _admin []common.Address, _adminSet []bool) (*TokenLockupAdminSetIterator, error) {

	var _adminRule []interface{}
	for _, _adminItem := range _admin {
		_adminRule = append(_adminRule, _adminItem)
	}
	var _adminSetRule []interface{}
	for _, _adminSetItem := range _adminSet {
		_adminSetRule = append(_adminSetRule, _adminSetItem)
	}

	logs, sub, err := _TokenLockup.contract.FilterLogs(opts, "AdminSet", _adminRule, _adminSetRule)
	if err != nil {
		return nil, err
	}
	return &TokenLockupAdminSetIterator{contract: _TokenLockup.contract, event: "AdminSet", logs: logs, sub: sub}, nil
}

// WatchAdminSet is a free log subscription operation binding the contract event 0xe68d2c359a771606c400cf8b87000cf5864010363d6a736e98f5047b7bbe18e9.
//
// Solidity: event AdminSet(_admin indexed address, _adminSet indexed bool)
func (_TokenLockup *TokenLockupFilterer) WatchAdminSet(opts *bind.WatchOpts, sink chan<- *TokenLockupAdminSet, _admin []common.Address, _adminSet []bool) (event.Subscription, error) {

	var _adminRule []interface{}
	for _, _adminItem := range _admin {
		_adminRule = append(_adminRule, _adminItem)
	}
	var _adminSetRule []interface{}
	for _, _adminSetItem := range _adminSet {
		_adminSetRule = append(_adminSetRule, _adminSetItem)
	}

	logs, sub, err := _TokenLockup.contract.WatchLogs(opts, "AdminSet", _adminRule, _adminSetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenLockupAdminSet)
				if err := _TokenLockup.contract.UnpackLog(event, "AdminSet", log); err != nil {
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

// TokenLockupDepositWithdrawnIterator is returned from FilterDepositWithdrawn and is used to iterate over the raw logs and unpacked data for DepositWithdrawn events raised by the TokenLockup contract.
type TokenLockupDepositWithdrawnIterator struct {
	Event *TokenLockupDepositWithdrawn // Event containing the contract specifics and raw log

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
func (it *TokenLockupDepositWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenLockupDepositWithdrawn)
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
		it.Event = new(TokenLockupDepositWithdrawn)
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
func (it *TokenLockupDepositWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenLockupDepositWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenLockupDepositWithdrawn represents a DepositWithdrawn event raised by the TokenLockup contract.
type TokenLockupDepositWithdrawn struct {
	Staker  common.Address
	Amount  *big.Int
	StakeId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDepositWithdrawn is a free log retrieval operation binding the contract event 0x7719804546c0185709e60c90d164447ff251a5ba29af0216faa921350f6bebf7.
//
// Solidity: event DepositWithdrawn(_staker address, _amount uint256, _stakeId uint256)
func (_TokenLockup *TokenLockupFilterer) FilterDepositWithdrawn(opts *bind.FilterOpts) (*TokenLockupDepositWithdrawnIterator, error) {

	logs, sub, err := _TokenLockup.contract.FilterLogs(opts, "DepositWithdrawn")
	if err != nil {
		return nil, err
	}
	return &TokenLockupDepositWithdrawnIterator{contract: _TokenLockup.contract, event: "DepositWithdrawn", logs: logs, sub: sub}, nil
}

// WatchDepositWithdrawn is a free log subscription operation binding the contract event 0x7719804546c0185709e60c90d164447ff251a5ba29af0216faa921350f6bebf7.
//
// Solidity: event DepositWithdrawn(_staker address, _amount uint256, _stakeId uint256)
func (_TokenLockup *TokenLockupFilterer) WatchDepositWithdrawn(opts *bind.WatchOpts, sink chan<- *TokenLockupDepositWithdrawn) (event.Subscription, error) {

	logs, sub, err := _TokenLockup.contract.WatchLogs(opts, "DepositWithdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenLockupDepositWithdrawn)
				if err := _TokenLockup.contract.UnpackLog(event, "DepositWithdrawn", log); err != nil {
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

// TokenLockupEthRewardIterator is returned from FilterEthReward and is used to iterate over the raw logs and unpacked data for EthReward events raised by the TokenLockup contract.
type TokenLockupEthRewardIterator struct {
	Event *TokenLockupEthReward // Event containing the contract specifics and raw log

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
func (it *TokenLockupEthRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenLockupEthReward)
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
		it.Event = new(TokenLockupEthReward)
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
func (it *TokenLockupEthRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenLockupEthRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenLockupEthReward represents a EthReward event raised by the TokenLockup contract.
type TokenLockupEthReward struct {
	Staker common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEthReward is a free log retrieval operation binding the contract event 0x196f95be2bd8e6aa49ac118195dec22c01e6adf2c34eaae9fefbb9efc8b78a79.
//
// Solidity: event EthReward(_staker address, _amount uint256)
func (_TokenLockup *TokenLockupFilterer) FilterEthReward(opts *bind.FilterOpts) (*TokenLockupEthRewardIterator, error) {

	logs, sub, err := _TokenLockup.contract.FilterLogs(opts, "EthReward")
	if err != nil {
		return nil, err
	}
	return &TokenLockupEthRewardIterator{contract: _TokenLockup.contract, event: "EthReward", logs: logs, sub: sub}, nil
}

// WatchEthReward is a free log subscription operation binding the contract event 0x196f95be2bd8e6aa49ac118195dec22c01e6adf2c34eaae9fefbb9efc8b78a79.
//
// Solidity: event EthReward(_staker address, _amount uint256)
func (_TokenLockup *TokenLockupFilterer) WatchEthReward(opts *bind.WatchOpts, sink chan<- *TokenLockupEthReward) (event.Subscription, error) {

	logs, sub, err := _TokenLockup.contract.WatchLogs(opts, "EthReward")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenLockupEthReward)
				if err := _TokenLockup.contract.UnpackLog(event, "EthReward", log); err != nil {
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

// TokenLockupEthUsdPriceUpdatedIterator is returned from FilterEthUsdPriceUpdated and is used to iterate over the raw logs and unpacked data for EthUsdPriceUpdated events raised by the TokenLockup contract.
type TokenLockupEthUsdPriceUpdatedIterator struct {
	Event *TokenLockupEthUsdPriceUpdated // Event containing the contract specifics and raw log

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
func (it *TokenLockupEthUsdPriceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenLockupEthUsdPriceUpdated)
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
		it.Event = new(TokenLockupEthUsdPriceUpdated)
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
func (it *TokenLockupEthUsdPriceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenLockupEthUsdPriceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenLockupEthUsdPriceUpdated represents a EthUsdPriceUpdated event raised by the TokenLockup contract.
type TokenLockupEthUsdPriceUpdated struct {
	Price *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterEthUsdPriceUpdated is a free log retrieval operation binding the contract event 0xb8a7d16d8966ae3f48e95e49ed078690c23bd91afb16363bbaaaac00ff99b03f.
//
// Solidity: event EthUsdPriceUpdated(price uint256)
func (_TokenLockup *TokenLockupFilterer) FilterEthUsdPriceUpdated(opts *bind.FilterOpts) (*TokenLockupEthUsdPriceUpdatedIterator, error) {

	logs, sub, err := _TokenLockup.contract.FilterLogs(opts, "EthUsdPriceUpdated")
	if err != nil {
		return nil, err
	}
	return &TokenLockupEthUsdPriceUpdatedIterator{contract: _TokenLockup.contract, event: "EthUsdPriceUpdated", logs: logs, sub: sub}, nil
}

// WatchEthUsdPriceUpdated is a free log subscription operation binding the contract event 0xb8a7d16d8966ae3f48e95e49ed078690c23bd91afb16363bbaaaac00ff99b03f.
//
// Solidity: event EthUsdPriceUpdated(price uint256)
func (_TokenLockup *TokenLockupFilterer) WatchEthUsdPriceUpdated(opts *bind.WatchOpts, sink chan<- *TokenLockupEthUsdPriceUpdated) (event.Subscription, error) {

	logs, sub, err := _TokenLockup.contract.WatchLogs(opts, "EthUsdPriceUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenLockupEthUsdPriceUpdated)
				if err := _TokenLockup.contract.UnpackLog(event, "EthUsdPriceUpdated", log); err != nil {
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

// TokenLockupEthWithdrawnIterator is returned from FilterEthWithdrawn and is used to iterate over the raw logs and unpacked data for EthWithdrawn events raised by the TokenLockup contract.
type TokenLockupEthWithdrawnIterator struct {
	Event *TokenLockupEthWithdrawn // Event containing the contract specifics and raw log

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
func (it *TokenLockupEthWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenLockupEthWithdrawn)
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
		it.Event = new(TokenLockupEthWithdrawn)
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
func (it *TokenLockupEthWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenLockupEthWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenLockupEthWithdrawn represents a EthWithdrawn event raised by the TokenLockup contract.
type TokenLockupEthWithdrawn struct {
	Withdrawer common.Address
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterEthWithdrawn is a free log retrieval operation binding the contract event 0x8455ae6be5d92f1df1c3c1484388e247a36c7e60d72055ae216dbc258f257d4b.
//
// Solidity: event EthWithdrawn(_withdrawer address, _amount uint256)
func (_TokenLockup *TokenLockupFilterer) FilterEthWithdrawn(opts *bind.FilterOpts) (*TokenLockupEthWithdrawnIterator, error) {

	logs, sub, err := _TokenLockup.contract.FilterLogs(opts, "EthWithdrawn")
	if err != nil {
		return nil, err
	}
	return &TokenLockupEthWithdrawnIterator{contract: _TokenLockup.contract, event: "EthWithdrawn", logs: logs, sub: sub}, nil
}

// WatchEthWithdrawn is a free log subscription operation binding the contract event 0x8455ae6be5d92f1df1c3c1484388e247a36c7e60d72055ae216dbc258f257d4b.
//
// Solidity: event EthWithdrawn(_withdrawer address, _amount uint256)
func (_TokenLockup *TokenLockupFilterer) WatchEthWithdrawn(opts *bind.WatchOpts, sink chan<- *TokenLockupEthWithdrawn) (event.Subscription, error) {

	logs, sub, err := _TokenLockup.contract.WatchLogs(opts, "EthWithdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenLockupEthWithdrawn)
				if err := _TokenLockup.contract.UnpackLog(event, "EthWithdrawn", log); err != nil {
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

// TokenLockupNewOraclizeQueryIterator is returned from FilterNewOraclizeQuery and is used to iterate over the raw logs and unpacked data for NewOraclizeQuery events raised by the TokenLockup contract.
type TokenLockupNewOraclizeQueryIterator struct {
	Event *TokenLockupNewOraclizeQuery // Event containing the contract specifics and raw log

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
func (it *TokenLockupNewOraclizeQueryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenLockupNewOraclizeQuery)
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
		it.Event = new(TokenLockupNewOraclizeQuery)
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
func (it *TokenLockupNewOraclizeQueryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenLockupNewOraclizeQueryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenLockupNewOraclizeQuery represents a NewOraclizeQuery event raised by the TokenLockup contract.
type TokenLockupNewOraclizeQuery struct {
	Result string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNewOraclizeQuery is a free log retrieval operation binding the contract event 0x096835e36c2ccea88ff2b3aca87dfc938b977e52ea656873ff76a8dba50d4d34.
//
// Solidity: event NewOraclizeQuery(result string)
func (_TokenLockup *TokenLockupFilterer) FilterNewOraclizeQuery(opts *bind.FilterOpts) (*TokenLockupNewOraclizeQueryIterator, error) {

	logs, sub, err := _TokenLockup.contract.FilterLogs(opts, "NewOraclizeQuery")
	if err != nil {
		return nil, err
	}
	return &TokenLockupNewOraclizeQueryIterator{contract: _TokenLockup.contract, event: "NewOraclizeQuery", logs: logs, sub: sub}, nil
}

// WatchNewOraclizeQuery is a free log subscription operation binding the contract event 0x096835e36c2ccea88ff2b3aca87dfc938b977e52ea656873ff76a8dba50d4d34.
//
// Solidity: event NewOraclizeQuery(result string)
func (_TokenLockup *TokenLockupFilterer) WatchNewOraclizeQuery(opts *bind.WatchOpts, sink chan<- *TokenLockupNewOraclizeQuery) (event.Subscription, error) {

	logs, sub, err := _TokenLockup.contract.WatchLogs(opts, "NewOraclizeQuery")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenLockupNewOraclizeQuery)
				if err := _TokenLockup.contract.UnpackLog(event, "NewOraclizeQuery", log); err != nil {
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

// TokenLockupOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TokenLockup contract.
type TokenLockupOwnershipTransferredIterator struct {
	Event *TokenLockupOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TokenLockupOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenLockupOwnershipTransferred)
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
		it.Event = new(TokenLockupOwnershipTransferred)
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
func (it *TokenLockupOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenLockupOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenLockupOwnershipTransferred represents a OwnershipTransferred event raised by the TokenLockup contract.
type TokenLockupOwnershipTransferred struct {
	PreviousOwner        common.Address
	NewOwner             common.Address
	OwnershipTransferred bool
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x7fdc2a4b6eb39ec3363d710d188620bd1e97b3c434161f187b4d0dc0544faa58.
//
// Solidity: event OwnershipTransferred(_previousOwner indexed address, _newOwner indexed address, _ownershipTransferred indexed bool)
func (_TokenLockup *TokenLockupFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, _previousOwner []common.Address, _newOwner []common.Address, _ownershipTransferred []bool) (*TokenLockupOwnershipTransferredIterator, error) {

	var _previousOwnerRule []interface{}
	for _, _previousOwnerItem := range _previousOwner {
		_previousOwnerRule = append(_previousOwnerRule, _previousOwnerItem)
	}
	var _newOwnerRule []interface{}
	for _, _newOwnerItem := range _newOwner {
		_newOwnerRule = append(_newOwnerRule, _newOwnerItem)
	}
	var _ownershipTransferredRule []interface{}
	for _, _ownershipTransferredItem := range _ownershipTransferred {
		_ownershipTransferredRule = append(_ownershipTransferredRule, _ownershipTransferredItem)
	}

	logs, sub, err := _TokenLockup.contract.FilterLogs(opts, "OwnershipTransferred", _previousOwnerRule, _newOwnerRule, _ownershipTransferredRule)
	if err != nil {
		return nil, err
	}
	return &TokenLockupOwnershipTransferredIterator{contract: _TokenLockup.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x7fdc2a4b6eb39ec3363d710d188620bd1e97b3c434161f187b4d0dc0544faa58.
//
// Solidity: event OwnershipTransferred(_previousOwner indexed address, _newOwner indexed address, _ownershipTransferred indexed bool)
func (_TokenLockup *TokenLockupFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TokenLockupOwnershipTransferred, _previousOwner []common.Address, _newOwner []common.Address, _ownershipTransferred []bool) (event.Subscription, error) {

	var _previousOwnerRule []interface{}
	for _, _previousOwnerItem := range _previousOwner {
		_previousOwnerRule = append(_previousOwnerRule, _previousOwnerItem)
	}
	var _newOwnerRule []interface{}
	for _, _newOwnerItem := range _newOwner {
		_newOwnerRule = append(_newOwnerRule, _newOwnerItem)
	}
	var _ownershipTransferredRule []interface{}
	for _, _ownershipTransferredItem := range _ownershipTransferred {
		_ownershipTransferredRule = append(_ownershipTransferredRule, _ownershipTransferredItem)
	}

	logs, sub, err := _TokenLockup.contract.WatchLogs(opts, "OwnershipTransferred", _previousOwnerRule, _newOwnerRule, _ownershipTransferredRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenLockupOwnershipTransferred)
				if err := _TokenLockup.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// TokenLockupRtcRewardIterator is returned from FilterRtcReward and is used to iterate over the raw logs and unpacked data for RtcReward events raised by the TokenLockup contract.
type TokenLockupRtcRewardIterator struct {
	Event *TokenLockupRtcReward // Event containing the contract specifics and raw log

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
func (it *TokenLockupRtcRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenLockupRtcReward)
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
		it.Event = new(TokenLockupRtcReward)
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
func (it *TokenLockupRtcRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenLockupRtcRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenLockupRtcReward represents a RtcReward event raised by the TokenLockup contract.
type TokenLockupRtcReward struct {
	Staker common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRtcReward is a free log retrieval operation binding the contract event 0xe0ba89edeae157ec385468cf95ff7ea61497f95bf3e0fe9637fa358aefdf7e28.
//
// Solidity: event RtcReward(_staker address, _amount uint256)
func (_TokenLockup *TokenLockupFilterer) FilterRtcReward(opts *bind.FilterOpts) (*TokenLockupRtcRewardIterator, error) {

	logs, sub, err := _TokenLockup.contract.FilterLogs(opts, "RtcReward")
	if err != nil {
		return nil, err
	}
	return &TokenLockupRtcRewardIterator{contract: _TokenLockup.contract, event: "RtcReward", logs: logs, sub: sub}, nil
}

// WatchRtcReward is a free log subscription operation binding the contract event 0xe0ba89edeae157ec385468cf95ff7ea61497f95bf3e0fe9637fa358aefdf7e28.
//
// Solidity: event RtcReward(_staker address, _amount uint256)
func (_TokenLockup *TokenLockupFilterer) WatchRtcReward(opts *bind.WatchOpts, sink chan<- *TokenLockupRtcReward) (event.Subscription, error) {

	logs, sub, err := _TokenLockup.contract.WatchLogs(opts, "RtcReward")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenLockupRtcReward)
				if err := _TokenLockup.contract.UnpackLog(event, "RtcReward", log); err != nil {
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

// TokenLockupSignUpFeeUpdatedIterator is returned from FilterSignUpFeeUpdated and is used to iterate over the raw logs and unpacked data for SignUpFeeUpdated events raised by the TokenLockup contract.
type TokenLockupSignUpFeeUpdatedIterator struct {
	Event *TokenLockupSignUpFeeUpdated // Event containing the contract specifics and raw log

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
func (it *TokenLockupSignUpFeeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenLockupSignUpFeeUpdated)
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
		it.Event = new(TokenLockupSignUpFeeUpdated)
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
func (it *TokenLockupSignUpFeeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenLockupSignUpFeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenLockupSignUpFeeUpdated represents a SignUpFeeUpdated event raised by the TokenLockup contract.
type TokenLockupSignUpFeeUpdated struct {
	Fee *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSignUpFeeUpdated is a free log retrieval operation binding the contract event 0x72cfec5345d9622cfd9fdeadefa51edec8e3d432a8f9fcb12cd65d484b47b204.
//
// Solidity: event SignUpFeeUpdated(fee uint256)
func (_TokenLockup *TokenLockupFilterer) FilterSignUpFeeUpdated(opts *bind.FilterOpts) (*TokenLockupSignUpFeeUpdatedIterator, error) {

	logs, sub, err := _TokenLockup.contract.FilterLogs(opts, "SignUpFeeUpdated")
	if err != nil {
		return nil, err
	}
	return &TokenLockupSignUpFeeUpdatedIterator{contract: _TokenLockup.contract, event: "SignUpFeeUpdated", logs: logs, sub: sub}, nil
}

// WatchSignUpFeeUpdated is a free log subscription operation binding the contract event 0x72cfec5345d9622cfd9fdeadefa51edec8e3d432a8f9fcb12cd65d484b47b204.
//
// Solidity: event SignUpFeeUpdated(fee uint256)
func (_TokenLockup *TokenLockupFilterer) WatchSignUpFeeUpdated(opts *bind.WatchOpts, sink chan<- *TokenLockupSignUpFeeUpdated) (event.Subscription, error) {

	logs, sub, err := _TokenLockup.contract.WatchLogs(opts, "SignUpFeeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenLockupSignUpFeeUpdated)
				if err := _TokenLockup.contract.UnpackLog(event, "SignUpFeeUpdated", log); err != nil {
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

// TokenLockupStakeDepositedIterator is returned from FilterStakeDeposited and is used to iterate over the raw logs and unpacked data for StakeDeposited events raised by the TokenLockup contract.
type TokenLockupStakeDepositedIterator struct {
	Event *TokenLockupStakeDeposited // Event containing the contract specifics and raw log

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
func (it *TokenLockupStakeDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenLockupStakeDeposited)
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
		it.Event = new(TokenLockupStakeDeposited)
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
func (it *TokenLockupStakeDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenLockupStakeDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenLockupStakeDeposited represents a StakeDeposited event raised by the TokenLockup contract.
type TokenLockupStakeDeposited struct {
	Depositer   common.Address
	Amount      *big.Int
	WeeksStaked *big.Int
	KhSec       *big.Int
	Id          *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterStakeDeposited is a free log retrieval operation binding the contract event 0x1a325385f16807e99fb688b597db78b00faee313dcf02e882dd16daab6fc3e1f.
//
// Solidity: event StakeDeposited(_depositer address, _amount uint256, _weeksStaked uint256, _khSec uint256, _id uint256)
func (_TokenLockup *TokenLockupFilterer) FilterStakeDeposited(opts *bind.FilterOpts) (*TokenLockupStakeDepositedIterator, error) {

	logs, sub, err := _TokenLockup.contract.FilterLogs(opts, "StakeDeposited")
	if err != nil {
		return nil, err
	}
	return &TokenLockupStakeDepositedIterator{contract: _TokenLockup.contract, event: "StakeDeposited", logs: logs, sub: sub}, nil
}

// WatchStakeDeposited is a free log subscription operation binding the contract event 0x1a325385f16807e99fb688b597db78b00faee313dcf02e882dd16daab6fc3e1f.
//
// Solidity: event StakeDeposited(_depositer address, _amount uint256, _weeksStaked uint256, _khSec uint256, _id uint256)
func (_TokenLockup *TokenLockupFilterer) WatchStakeDeposited(opts *bind.WatchOpts, sink chan<- *TokenLockupStakeDeposited) (event.Subscription, error) {

	logs, sub, err := _TokenLockup.contract.WatchLogs(opts, "StakeDeposited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenLockupStakeDeposited)
				if err := _TokenLockup.contract.UnpackLog(event, "StakeDeposited", log); err != nil {
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
