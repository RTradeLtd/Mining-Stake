// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Oracle

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

// OracleABI is the input ABI used to generate the binding from.
const OracleABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"frozen\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"moderators\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_destinationContract\",\"type\":\"address\"},{\"name\":\"_ethUSD\",\"type\":\"uint256\"}],\"name\":\"updateEthPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_destinationContract\",\"type\":\"address\"},{\"name\":\"_rtcUSD\",\"type\":\"uint256\"}],\"name\":\"updateRtcPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"contracts\",\"outputs\":[{\"name\":\"contractAddress\",\"type\":\"address\"},{\"name\":\"updateFrequencyInHours\",\"type\":\"uint256\"},{\"name\":\"nextUpdate\",\"type\":\"uint256\"},{\"name\":\"enabled\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newAdmin\",\"type\":\"address\"}],\"name\":\"setAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_contractAddress\",\"type\":\"address\"},{\"name\":\"_updateFrequencyInHours\",\"type\":\"uint256\"},{\"name\":\"_enabledFunctions\",\"type\":\"bytes4[]\"}],\"name\":\"addAuthorizedContract\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_contractAddress\",\"type\":\"address\"}],\"name\":\"AuthorizedContractAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_admin\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_adminSet\",\"type\":\"bool\"}],\"name\":\"AdminSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_newOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_ownershipTransferred\",\"type\":\"bool\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// OracleBin is the compiled bytecode used for deploying new contracts.
const OracleBin = `608060405260008054600160a060020a033316600160a060020a031991821681179092556001805490911690911790556109068061003e6000396000f30060806040526004361061008a5763ffffffff60e060020a600035041663054f7d9c811461008f57806314d0f1ba146100b85780632f24b6f8146100d957806366aeb856146100fd57806369dc9ff314610121578063704b6c02146101745780638da5cb5b14610195578063a981320b146101c6578063f2fde38b1461022d578063f851a4401461024e575b600080fd5b34801561009b57600080fd5b506100a4610263565b604080519115158252519081900360200190f35b3480156100c457600080fd5b506100a4600160a060020a0360043516610284565b3480156100e557600080fd5b506100a4600160a060020a0360043516602435610299565b34801561010957600080fd5b506100a4600160a060020a0360043516602435610411565b34801561012d57600080fd5b50610142600160a060020a0360043516610546565b60408051600160a060020a03909516855260208501939093528383019190915215156060830152519081900360800190f35b34801561018057600080fd5b506100a4600160a060020a036004351661057c565b3480156101a157600080fd5b506101aa61060b565b60408051600160a060020a039092168252519081900360200190f35b3480156101d257600080fd5b5060408051602060046044358181013583810280860185019096528085526100a4958335600160a060020a031695602480359636969560649593949201929182918501908490808284375094975061061a9650505050505050565b34801561023957600080fd5b506100a4600160a060020a03600435166107cf565b34801561025a57600080fd5b506101aa610864565b60015474010000000000000000000000000000000000000000900460ff1681565b60026020526000908152604090205460ff1681565b6000805433600160a060020a03908116911614806102c5575060015433600160a060020a039081169116145b15156102d057600080fd5b600160a060020a03831660009081526003602081905260409091200154839060ff1615156102fd57600080fd5b604080517f75706461746545746850726963652875696e743235362900000000000000000081528151908190036017019020600160a060020a0386166000908152600360209081528382207bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19841683526004019052919091205485919060ff16151561038557600080fd5b85600160a060020a03166321370942866040518263ffffffff1660e060020a02815260040180828152602001915050602060405180830381600087803b1580156103ce57600080fd5b505af11580156103e2573d6000803e3d6000fd5b505050506040513d60208110156103f857600080fd5b5051151561040557600080fd5b50600195945050505050565b6000805433600160a060020a039081169116148061043d575060015433600160a060020a039081169116145b151561044857600080fd5b600160a060020a03831660009081526003602081905260409091200154839060ff16151561047557600080fd5b604080517f75706461746552746350726963652875696e743235362900000000000000000081528151908190036017019020600160a060020a0386166000908152600360209081528382207bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19841683526004019052919091205485919060ff1615156104fd57600080fd5b85600160a060020a03166358243471866040518263ffffffff1660e060020a02815260040180828152602001915050602060405180830381600087803b1580156103ce57600080fd5b60036020819052600091825260409091208054600182015460028301549290930154600160a060020a0390911692919060ff1684565b6000805433600160a060020a0390811691161461059857600080fd5b600154600160a060020a03838116911614156105b357600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03841690811782556040517fe68d2c359a771606c400cf8b87000cf5864010363d6a736e98f5047b7bbe18e990600090a3919050565b600054600160a060020a031690565b60006106246108b3565b6000805433600160a060020a0390811691161480610650575060015433600160a060020a039081169116145b151561065b57600080fd5b600160a060020a03861682526020820185905261069061068386610e1063ffffffff61087316565b429063ffffffff6108a116565b6040838101918252600160a060020a038881166000908152600360208181529382208751815473ffffffffffffffffffffffffffffffffffffffff191694169390931783559286015160018301559251600282015560608501519101805460ff191691151591909117905590505b835181101561078757600160a060020a038616600090815260036020526040812085516001926004909201919087908590811061073757fe5b6020908102919091018101517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168252810191909152604001600020805460ff19169115159190911790556001016106fe565b60408051600160a060020a038816815290517f351e40fe333cede9dd79a1ef4d16228b1702f272eb2c2166ec3c7ee6a9d4ce2c9181900360200190a150600195945050505050565b6000805433600160a060020a039081169116146107eb57600080fd5b600054600160a060020a038381169116141561080657600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03848116918217835560405160019333909216917f7fdc2a4b6eb39ec3363d710d188620bd1e97b3c434161f187b4d0dc0544faa5891a4919050565b600154600160a060020a031690565b600082820283158061088f575082848281151561088c57fe5b04145b151561089a57600080fd5b9392505050565b60008282018381101561089a57600080fd5b604080516080810182526000808252602082018190529181018290526060810191909152905600a165627a7a723058200c0bf1c5396159f9c03886e4927127159870314b95f778b7957724bb155015530029`

// DeployOracle deploys a new Ethereum contract, binding an instance of Oracle to it.
func DeployOracle(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Oracle, error) {
	parsed, err := abi.JSON(strings.NewReader(OracleABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OracleBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Oracle{OracleCaller: OracleCaller{contract: contract}, OracleTransactor: OracleTransactor{contract: contract}, OracleFilterer: OracleFilterer{contract: contract}}, nil
}

// Oracle is an auto generated Go binding around an Ethereum contract.
type Oracle struct {
	OracleCaller     // Read-only binding to the contract
	OracleTransactor // Write-only binding to the contract
	OracleFilterer   // Log filterer for contract events
}

// OracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type OracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OracleSession struct {
	Contract     *Oracle           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OracleCallerSession struct {
	Contract *OracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OracleTransactorSession struct {
	Contract     *OracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type OracleRaw struct {
	Contract *Oracle // Generic contract binding to access the raw methods on
}

// OracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OracleCallerRaw struct {
	Contract *OracleCaller // Generic read-only contract binding to access the raw methods on
}

// OracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OracleTransactorRaw struct {
	Contract *OracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOracle creates a new instance of Oracle, bound to a specific deployed contract.
func NewOracle(address common.Address, backend bind.ContractBackend) (*Oracle, error) {
	contract, err := bindOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Oracle{OracleCaller: OracleCaller{contract: contract}, OracleTransactor: OracleTransactor{contract: contract}, OracleFilterer: OracleFilterer{contract: contract}}, nil
}

// NewOracleCaller creates a new read-only instance of Oracle, bound to a specific deployed contract.
func NewOracleCaller(address common.Address, caller bind.ContractCaller) (*OracleCaller, error) {
	contract, err := bindOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OracleCaller{contract: contract}, nil
}

// NewOracleTransactor creates a new write-only instance of Oracle, bound to a specific deployed contract.
func NewOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*OracleTransactor, error) {
	contract, err := bindOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OracleTransactor{contract: contract}, nil
}

// NewOracleFilterer creates a new log filterer instance of Oracle, bound to a specific deployed contract.
func NewOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*OracleFilterer, error) {
	contract, err := bindOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OracleFilterer{contract: contract}, nil
}

// bindOracle binds a generic wrapper to an already deployed contract.
func bindOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OracleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oracle *OracleRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Oracle.Contract.OracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oracle *OracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oracle.Contract.OracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oracle *OracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Oracle.Contract.OracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oracle *OracleCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Oracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oracle *OracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oracle *OracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Oracle.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_Oracle *OracleCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Oracle.contract.Call(opts, out, "admin")
	return *ret0, err
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_Oracle *OracleSession) Admin() (common.Address, error) {
	return _Oracle.Contract.Admin(&_Oracle.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_Oracle *OracleCallerSession) Admin() (common.Address, error) {
	return _Oracle.Contract.Admin(&_Oracle.CallOpts)
}

// Contracts is a free data retrieval call binding the contract method 0x69dc9ff3.
//
// Solidity: function contracts( address) constant returns(contractAddress address, updateFrequencyInHours uint256, nextUpdate uint256, enabled bool)
func (_Oracle *OracleCaller) Contracts(opts *bind.CallOpts, arg0 common.Address) (struct {
	ContractAddress        common.Address
	UpdateFrequencyInHours *big.Int
	NextUpdate             *big.Int
	Enabled                bool
}, error) {
	ret := new(struct {
		ContractAddress        common.Address
		UpdateFrequencyInHours *big.Int
		NextUpdate             *big.Int
		Enabled                bool
	})
	out := ret
	err := _Oracle.contract.Call(opts, out, "contracts", arg0)
	return *ret, err
}

// Contracts is a free data retrieval call binding the contract method 0x69dc9ff3.
//
// Solidity: function contracts( address) constant returns(contractAddress address, updateFrequencyInHours uint256, nextUpdate uint256, enabled bool)
func (_Oracle *OracleSession) Contracts(arg0 common.Address) (struct {
	ContractAddress        common.Address
	UpdateFrequencyInHours *big.Int
	NextUpdate             *big.Int
	Enabled                bool
}, error) {
	return _Oracle.Contract.Contracts(&_Oracle.CallOpts, arg0)
}

// Contracts is a free data retrieval call binding the contract method 0x69dc9ff3.
//
// Solidity: function contracts( address) constant returns(contractAddress address, updateFrequencyInHours uint256, nextUpdate uint256, enabled bool)
func (_Oracle *OracleCallerSession) Contracts(arg0 common.Address) (struct {
	ContractAddress        common.Address
	UpdateFrequencyInHours *big.Int
	NextUpdate             *big.Int
	Enabled                bool
}, error) {
	return _Oracle.Contract.Contracts(&_Oracle.CallOpts, arg0)
}

// Frozen is a free data retrieval call binding the contract method 0x054f7d9c.
//
// Solidity: function frozen() constant returns(bool)
func (_Oracle *OracleCaller) Frozen(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Oracle.contract.Call(opts, out, "frozen")
	return *ret0, err
}

// Frozen is a free data retrieval call binding the contract method 0x054f7d9c.
//
// Solidity: function frozen() constant returns(bool)
func (_Oracle *OracleSession) Frozen() (bool, error) {
	return _Oracle.Contract.Frozen(&_Oracle.CallOpts)
}

// Frozen is a free data retrieval call binding the contract method 0x054f7d9c.
//
// Solidity: function frozen() constant returns(bool)
func (_Oracle *OracleCallerSession) Frozen() (bool, error) {
	return _Oracle.Contract.Frozen(&_Oracle.CallOpts)
}

// Moderators is a free data retrieval call binding the contract method 0x14d0f1ba.
//
// Solidity: function moderators( address) constant returns(bool)
func (_Oracle *OracleCaller) Moderators(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Oracle.contract.Call(opts, out, "moderators", arg0)
	return *ret0, err
}

// Moderators is a free data retrieval call binding the contract method 0x14d0f1ba.
//
// Solidity: function moderators( address) constant returns(bool)
func (_Oracle *OracleSession) Moderators(arg0 common.Address) (bool, error) {
	return _Oracle.Contract.Moderators(&_Oracle.CallOpts, arg0)
}

// Moderators is a free data retrieval call binding the contract method 0x14d0f1ba.
//
// Solidity: function moderators( address) constant returns(bool)
func (_Oracle *OracleCallerSession) Moderators(arg0 common.Address) (bool, error) {
	return _Oracle.Contract.Moderators(&_Oracle.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Oracle *OracleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Oracle.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Oracle *OracleSession) Owner() (common.Address, error) {
	return _Oracle.Contract.Owner(&_Oracle.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Oracle *OracleCallerSession) Owner() (common.Address, error) {
	return _Oracle.Contract.Owner(&_Oracle.CallOpts)
}

// AddAuthorizedContract is a paid mutator transaction binding the contract method 0xa981320b.
//
// Solidity: function addAuthorizedContract(_contractAddress address, _updateFrequencyInHours uint256, _enabledFunctions bytes4[]) returns(bool)
func (_Oracle *OracleTransactor) AddAuthorizedContract(opts *bind.TransactOpts, _contractAddress common.Address, _updateFrequencyInHours *big.Int, _enabledFunctions [][4]byte) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "addAuthorizedContract", _contractAddress, _updateFrequencyInHours, _enabledFunctions)
}

// AddAuthorizedContract is a paid mutator transaction binding the contract method 0xa981320b.
//
// Solidity: function addAuthorizedContract(_contractAddress address, _updateFrequencyInHours uint256, _enabledFunctions bytes4[]) returns(bool)
func (_Oracle *OracleSession) AddAuthorizedContract(_contractAddress common.Address, _updateFrequencyInHours *big.Int, _enabledFunctions [][4]byte) (*types.Transaction, error) {
	return _Oracle.Contract.AddAuthorizedContract(&_Oracle.TransactOpts, _contractAddress, _updateFrequencyInHours, _enabledFunctions)
}

// AddAuthorizedContract is a paid mutator transaction binding the contract method 0xa981320b.
//
// Solidity: function addAuthorizedContract(_contractAddress address, _updateFrequencyInHours uint256, _enabledFunctions bytes4[]) returns(bool)
func (_Oracle *OracleTransactorSession) AddAuthorizedContract(_contractAddress common.Address, _updateFrequencyInHours *big.Int, _enabledFunctions [][4]byte) (*types.Transaction, error) {
	return _Oracle.Contract.AddAuthorizedContract(&_Oracle.TransactOpts, _contractAddress, _updateFrequencyInHours, _enabledFunctions)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(_newAdmin address) returns(bool)
func (_Oracle *OracleTransactor) SetAdmin(opts *bind.TransactOpts, _newAdmin common.Address) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "setAdmin", _newAdmin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(_newAdmin address) returns(bool)
func (_Oracle *OracleSession) SetAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.SetAdmin(&_Oracle.TransactOpts, _newAdmin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(_newAdmin address) returns(bool)
func (_Oracle *OracleTransactorSession) SetAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.SetAdmin(&_Oracle.TransactOpts, _newAdmin)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns(bool)
func (_Oracle *OracleTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns(bool)
func (_Oracle *OracleSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.TransferOwnership(&_Oracle.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns(bool)
func (_Oracle *OracleTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.TransferOwnership(&_Oracle.TransactOpts, _newOwner)
}

// UpdateEthPrice is a paid mutator transaction binding the contract method 0x2f24b6f8.
//
// Solidity: function updateEthPrice(_destinationContract address, _ethUSD uint256) returns(bool)
func (_Oracle *OracleTransactor) UpdateEthPrice(opts *bind.TransactOpts, _destinationContract common.Address, _ethUSD *big.Int) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "updateEthPrice", _destinationContract, _ethUSD)
}

// UpdateEthPrice is a paid mutator transaction binding the contract method 0x2f24b6f8.
//
// Solidity: function updateEthPrice(_destinationContract address, _ethUSD uint256) returns(bool)
func (_Oracle *OracleSession) UpdateEthPrice(_destinationContract common.Address, _ethUSD *big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.UpdateEthPrice(&_Oracle.TransactOpts, _destinationContract, _ethUSD)
}

// UpdateEthPrice is a paid mutator transaction binding the contract method 0x2f24b6f8.
//
// Solidity: function updateEthPrice(_destinationContract address, _ethUSD uint256) returns(bool)
func (_Oracle *OracleTransactorSession) UpdateEthPrice(_destinationContract common.Address, _ethUSD *big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.UpdateEthPrice(&_Oracle.TransactOpts, _destinationContract, _ethUSD)
}

// UpdateRtcPrice is a paid mutator transaction binding the contract method 0x66aeb856.
//
// Solidity: function updateRtcPrice(_destinationContract address, _rtcUSD uint256) returns(bool)
func (_Oracle *OracleTransactor) UpdateRtcPrice(opts *bind.TransactOpts, _destinationContract common.Address, _rtcUSD *big.Int) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "updateRtcPrice", _destinationContract, _rtcUSD)
}

// UpdateRtcPrice is a paid mutator transaction binding the contract method 0x66aeb856.
//
// Solidity: function updateRtcPrice(_destinationContract address, _rtcUSD uint256) returns(bool)
func (_Oracle *OracleSession) UpdateRtcPrice(_destinationContract common.Address, _rtcUSD *big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.UpdateRtcPrice(&_Oracle.TransactOpts, _destinationContract, _rtcUSD)
}

// UpdateRtcPrice is a paid mutator transaction binding the contract method 0x66aeb856.
//
// Solidity: function updateRtcPrice(_destinationContract address, _rtcUSD uint256) returns(bool)
func (_Oracle *OracleTransactorSession) UpdateRtcPrice(_destinationContract common.Address, _rtcUSD *big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.UpdateRtcPrice(&_Oracle.TransactOpts, _destinationContract, _rtcUSD)
}

// OracleAdminSetIterator is returned from FilterAdminSet and is used to iterate over the raw logs and unpacked data for AdminSet events raised by the Oracle contract.
type OracleAdminSetIterator struct {
	Event *OracleAdminSet // Event containing the contract specifics and raw log

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
func (it *OracleAdminSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleAdminSet)
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
		it.Event = new(OracleAdminSet)
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
func (it *OracleAdminSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleAdminSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleAdminSet represents a AdminSet event raised by the Oracle contract.
type OracleAdminSet struct {
	Admin    common.Address
	AdminSet bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAdminSet is a free log retrieval operation binding the contract event 0xe68d2c359a771606c400cf8b87000cf5864010363d6a736e98f5047b7bbe18e9.
//
// Solidity: event AdminSet(_admin indexed address, _adminSet indexed bool)
func (_Oracle *OracleFilterer) FilterAdminSet(opts *bind.FilterOpts, _admin []common.Address, _adminSet []bool) (*OracleAdminSetIterator, error) {

	var _adminRule []interface{}
	for _, _adminItem := range _admin {
		_adminRule = append(_adminRule, _adminItem)
	}
	var _adminSetRule []interface{}
	for _, _adminSetItem := range _adminSet {
		_adminSetRule = append(_adminSetRule, _adminSetItem)
	}

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "AdminSet", _adminRule, _adminSetRule)
	if err != nil {
		return nil, err
	}
	return &OracleAdminSetIterator{contract: _Oracle.contract, event: "AdminSet", logs: logs, sub: sub}, nil
}

// WatchAdminSet is a free log subscription operation binding the contract event 0xe68d2c359a771606c400cf8b87000cf5864010363d6a736e98f5047b7bbe18e9.
//
// Solidity: event AdminSet(_admin indexed address, _adminSet indexed bool)
func (_Oracle *OracleFilterer) WatchAdminSet(opts *bind.WatchOpts, sink chan<- *OracleAdminSet, _admin []common.Address, _adminSet []bool) (event.Subscription, error) {

	var _adminRule []interface{}
	for _, _adminItem := range _admin {
		_adminRule = append(_adminRule, _adminItem)
	}
	var _adminSetRule []interface{}
	for _, _adminSetItem := range _adminSet {
		_adminSetRule = append(_adminSetRule, _adminSetItem)
	}

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "AdminSet", _adminRule, _adminSetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleAdminSet)
				if err := _Oracle.contract.UnpackLog(event, "AdminSet", log); err != nil {
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

// OracleAuthorizedContractAddedIterator is returned from FilterAuthorizedContractAdded and is used to iterate over the raw logs and unpacked data for AuthorizedContractAdded events raised by the Oracle contract.
type OracleAuthorizedContractAddedIterator struct {
	Event *OracleAuthorizedContractAdded // Event containing the contract specifics and raw log

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
func (it *OracleAuthorizedContractAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleAuthorizedContractAdded)
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
		it.Event = new(OracleAuthorizedContractAdded)
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
func (it *OracleAuthorizedContractAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleAuthorizedContractAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleAuthorizedContractAdded represents a AuthorizedContractAdded event raised by the Oracle contract.
type OracleAuthorizedContractAdded struct {
	ContractAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterAuthorizedContractAdded is a free log retrieval operation binding the contract event 0x351e40fe333cede9dd79a1ef4d16228b1702f272eb2c2166ec3c7ee6a9d4ce2c.
//
// Solidity: event AuthorizedContractAdded(_contractAddress address)
func (_Oracle *OracleFilterer) FilterAuthorizedContractAdded(opts *bind.FilterOpts) (*OracleAuthorizedContractAddedIterator, error) {

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "AuthorizedContractAdded")
	if err != nil {
		return nil, err
	}
	return &OracleAuthorizedContractAddedIterator{contract: _Oracle.contract, event: "AuthorizedContractAdded", logs: logs, sub: sub}, nil
}

// WatchAuthorizedContractAdded is a free log subscription operation binding the contract event 0x351e40fe333cede9dd79a1ef4d16228b1702f272eb2c2166ec3c7ee6a9d4ce2c.
//
// Solidity: event AuthorizedContractAdded(_contractAddress address)
func (_Oracle *OracleFilterer) WatchAuthorizedContractAdded(opts *bind.WatchOpts, sink chan<- *OracleAuthorizedContractAdded) (event.Subscription, error) {

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "AuthorizedContractAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleAuthorizedContractAdded)
				if err := _Oracle.contract.UnpackLog(event, "AuthorizedContractAdded", log); err != nil {
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

// OracleOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Oracle contract.
type OracleOwnershipTransferredIterator struct {
	Event *OracleOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OracleOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleOwnershipTransferred)
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
		it.Event = new(OracleOwnershipTransferred)
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
func (it *OracleOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleOwnershipTransferred represents a OwnershipTransferred event raised by the Oracle contract.
type OracleOwnershipTransferred struct {
	PreviousOwner        common.Address
	NewOwner             common.Address
	OwnershipTransferred bool
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x7fdc2a4b6eb39ec3363d710d188620bd1e97b3c434161f187b4d0dc0544faa58.
//
// Solidity: event OwnershipTransferred(_previousOwner indexed address, _newOwner indexed address, _ownershipTransferred indexed bool)
func (_Oracle *OracleFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, _previousOwner []common.Address, _newOwner []common.Address, _ownershipTransferred []bool) (*OracleOwnershipTransferredIterator, error) {

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

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "OwnershipTransferred", _previousOwnerRule, _newOwnerRule, _ownershipTransferredRule)
	if err != nil {
		return nil, err
	}
	return &OracleOwnershipTransferredIterator{contract: _Oracle.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x7fdc2a4b6eb39ec3363d710d188620bd1e97b3c434161f187b4d0dc0544faa58.
//
// Solidity: event OwnershipTransferred(_previousOwner indexed address, _newOwner indexed address, _ownershipTransferred indexed bool)
func (_Oracle *OracleFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OracleOwnershipTransferred, _previousOwner []common.Address, _newOwner []common.Address, _ownershipTransferred []bool) (event.Subscription, error) {

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

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "OwnershipTransferred", _previousOwnerRule, _newOwnerRule, _ownershipTransferredRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleOwnershipTransferred)
				if err := _Oracle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
