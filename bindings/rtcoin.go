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

// RTCoinABI is the input ABI used to generate the binding from.
const RTCoinABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"freezeTransfers\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"frozen\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"approved\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"moderators\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_recipient\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"transferredFrom\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowed\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newAdmin\",\"type\":\"address\"}],\"name\":\"setAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_holder\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_recipient\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"transferred\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"thawTransfers\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"transfersFrozen\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_recipient\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_transfersFrozen\",\"type\":\"bool\"}],\"name\":\"TransfersFrozen\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_transfersThawed\",\"type\":\"bool\"}],\"name\":\"TransfersThawed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_admin\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_adminSet\",\"type\":\"bool\"}],\"name\":\"AdminSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_newOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_ownershipTransferred\",\"type\":\"bool\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// RTCoinBin is the compiled bytecode used for deploying new contracts.
const RTCoinBin = `608060405234801561001057600080fd5b5060008054600160a060020a033316600160a060020a031991821681179092556001805490911690911790556040805180820190915260068082527f5254436f696e0000000000000000000000000000000000000000000000000000602090920191825261008091600391610103565b506040805180820190915260038082527f525443000000000000000000000000000000000000000000000000000000000060209092019182526100c591600491610103565b506006805460ff191660121790556b02df458b2c635dcf55e000006005819055600160a060020a03331660009081526007602052604090205561019e565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061014457805160ff1916838001178555610171565b82800160010185558215610171579182015b82811115610171578251825591602001919060010190610156565b5061017d929150610181565b5090565b61019b91905b8082111561017d5760008155600101610187565b90565b610cd6806101ad6000396000f3006080604052600436106101115763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663015024608114610116578063054f7d9c1461013f57806306fdde0314610154578063095ea7b3146101de57806314d0f1ba1461020257806318160ddd1461022357806323b872dd1461024a57806327e235e314610274578063313ce567146102955780635c658165146102c0578063704b6c02146102e757806370a08231146103085780638da5cb5b1461032957806395d89b411461035a578063a9059cbb1461036f578063ce8e120a14610393578063dd62ed3e146103a8578063e45b8134146103cf578063f2fde38b146103e4578063f851a44014610405575b600080fd5b34801561012257600080fd5b5061012b61041a565b604080519115158252519081900360200190f35b34801561014b57600080fd5b5061012b610492565b34801561016057600080fd5b506101696104b3565b6040805160208082528351818301528351919283929083019185019080838360005b838110156101a357818101518382015260200161018b565b50505050905090810190601f1680156101d05780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156101ea57600080fd5b5061012b600160a060020a0360043516602435610541565b34801561020e57600080fd5b5061012b600160a060020a0360043516610641565b34801561022f57600080fd5b50610238610656565b60408051918252519081900360200190f35b34801561025657600080fd5b5061012b600160a060020a036004358116906024351660443561065c565b34801561028057600080fd5b50610238600160a060020a0360043516610815565b3480156102a157600080fd5b506102aa610827565b6040805160ff9092168252519081900360200190f35b3480156102cc57600080fd5b50610238600160a060020a0360043581169060243516610830565b3480156102f357600080fd5b5061012b600160a060020a036004351661084d565b34801561031457600080fd5b50610238600160a060020a03600435166108dc565b34801561033557600080fd5b5061033e6108f7565b60408051600160a060020a039092168252519081900360200190f35b34801561036657600080fd5b50610169610906565b34801561037b57600080fd5b5061012b600160a060020a0360043516602435610961565b34801561039f57600080fd5b5061012b610a4e565b3480156103b457600080fd5b50610238600160a060020a0360043581169060243516610ac2565b3480156103db57600080fd5b5061012b610aed565b3480156103f057600080fd5b5061012b600160a060020a0360043516610afb565b34801561041157600080fd5b5061033e610b90565b6000805433600160a060020a0390811691161480610446575060015433600160a060020a039081169116145b151561045157600080fd5b6006805461ff0019166101001790556040516001907fff7ea91c52ebd8c0d8018fdba50cb801e862f6795b1e17eeac882d4288b0934090600090a250600190565b60015474010000000000000000000000000000000000000000900460ff1681565b6003805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156105395780601f1061050e57610100808354040283529160200191610539565b820191906000526020600020905b81548152906001019060200180831161051c57829003601f168201915b505050505081565b6000600160a060020a038316151561055857600080fd5b6000821161056557600080fd5b600160a060020a033381166000908152600860209081526040808320938716835292905220546105958184610b9f565b1161059f57600080fd5b600160a060020a033381166000908152600860209081526040808320938716835292905220546105d5908363ffffffff610b9f16565b600160a060020a033381166000818152600860209081526040808320948916808452948252918290209490945580518681529051929391927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a350600192915050565b60026020526000908152604090205460ff1681565b60055490565b600654600090610100900460ff161561067457600080fd5b61067f848484610bb8565b151561068a57600080fd5b600160a060020a03808516600090815260086020908152604080832033909416835292905220548211156106bd57600080fd5b600160a060020a038085166000908152600860209081526040808320339094168352929052908120546106f6908463ffffffff610c9516565b101561070157600080fd5b600160a060020a0380851660009081526008602090815260408083203390941683529290522054610738908363ffffffff610c9516565b600160a060020a03808616600081815260086020908152604080832033909516835293815283822094909455908152600790925290205461077f908363ffffffff610c9516565b600160a060020a0380861660009081526007602052604080822093909355908516815220546107b4908363ffffffff610b9f16565b600160a060020a0380851660008181526007602090815260409182902094909455805186815290519193928816927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a35060019392505050565b60076020526000908152604090205481565b60065460ff1681565b600860209081526000928352604080842090915290825290205481565b6000805433600160a060020a0390811691161461086957600080fd5b600154600160a060020a038381169116141561088457600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03841690811782556040517fe68d2c359a771606c400cf8b87000cf5864010363d6a736e98f5047b7bbe18e990600090a3919050565b600160a060020a031660009081526007602052604090205490565b600054600160a060020a031690565b6004805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156105395780601f1061050e57610100808354040283529160200191610539565b600654600090610100900460ff161561097957600080fd5b610984338484610bb8565b151561098f57600080fd5b600160a060020a0333166000908152600760205260409020546109b8908363ffffffff610c9516565b600160a060020a0333811660009081526007602052604080822093909355908516815220546109ed908363ffffffff610b9f16565b600160a060020a038085166000818152600760209081526040918290209490945580518681529051919333909316927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a350600192915050565b6000805433600160a060020a0390811691161480610a7a575060015433600160a060020a039081169116145b1515610a8557600080fd5b6006805461ff00191690556040516001907fb36ea4d45a6246e5ea6da988f57a5bf9a9022c85940cc6fe92dd9e45bf632cf690600090a250600190565b600160a060020a03918216600090815260086020908152604080832093909416825291909152205490565b600654610100900460ff1681565b6000805433600160a060020a03908116911614610b1757600080fd5b600054600160a060020a0383811691161415610b3257600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03848116918217835560405160019333909216917f7fdc2a4b6eb39ec3363d710d188620bd1e97b3c434161f187b4d0dc0544faa5891a4919050565b600154600160a060020a031690565b600082820183811015610bb157600080fd5b9392505050565b6000600160a060020a03841615801590610bda5750600160a060020a03831615155b8015610be65750600082115b1515610bf157600080fd5b600160a060020a038416600090815260076020526040812054610c1a908463ffffffff610c9516565b1015610c2557600080fd5b600160a060020a038316600090815260076020526040812054610c4e908463ffffffff610b9f16565b11610c5857600080fd5b600160a060020a038316600090815260076020526040902054610c81818463ffffffff610b9f16565b11610c8b57600080fd5b5060019392505050565b600082821115610ca457600080fd5b509003905600a165627a7a7230582085aabeaa1e97149ab9b8bfd50b3f90a6c2cbff5f4766971521ba1477cf44c28e0029`

// DeployRTCoin deploys a new Ethereum contract, binding an instance of RTCoin to it.
func DeployRTCoin(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RTCoin, error) {
	parsed, err := abi.JSON(strings.NewReader(RTCoinABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RTCoinBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RTCoin{RTCoinCaller: RTCoinCaller{contract: contract}, RTCoinTransactor: RTCoinTransactor{contract: contract}, RTCoinFilterer: RTCoinFilterer{contract: contract}}, nil
}

// RTCoin is an auto generated Go binding around an Ethereum contract.
type RTCoin struct {
	RTCoinCaller     // Read-only binding to the contract
	RTCoinTransactor // Write-only binding to the contract
	RTCoinFilterer   // Log filterer for contract events
}

// RTCoinCaller is an auto generated read-only Go binding around an Ethereum contract.
type RTCoinCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RTCoinTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RTCoinTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RTCoinFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RTCoinFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RTCoinSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RTCoinSession struct {
	Contract     *RTCoin           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RTCoinCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RTCoinCallerSession struct {
	Contract *RTCoinCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RTCoinTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RTCoinTransactorSession struct {
	Contract     *RTCoinTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RTCoinRaw is an auto generated low-level Go binding around an Ethereum contract.
type RTCoinRaw struct {
	Contract *RTCoin // Generic contract binding to access the raw methods on
}

// RTCoinCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RTCoinCallerRaw struct {
	Contract *RTCoinCaller // Generic read-only contract binding to access the raw methods on
}

// RTCoinTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RTCoinTransactorRaw struct {
	Contract *RTCoinTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRTCoin creates a new instance of RTCoin, bound to a specific deployed contract.
func NewRTCoin(address common.Address, backend bind.ContractBackend) (*RTCoin, error) {
	contract, err := bindRTCoin(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RTCoin{RTCoinCaller: RTCoinCaller{contract: contract}, RTCoinTransactor: RTCoinTransactor{contract: contract}, RTCoinFilterer: RTCoinFilterer{contract: contract}}, nil
}

// NewRTCoinCaller creates a new read-only instance of RTCoin, bound to a specific deployed contract.
func NewRTCoinCaller(address common.Address, caller bind.ContractCaller) (*RTCoinCaller, error) {
	contract, err := bindRTCoin(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RTCoinCaller{contract: contract}, nil
}

// NewRTCoinTransactor creates a new write-only instance of RTCoin, bound to a specific deployed contract.
func NewRTCoinTransactor(address common.Address, transactor bind.ContractTransactor) (*RTCoinTransactor, error) {
	contract, err := bindRTCoin(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RTCoinTransactor{contract: contract}, nil
}

// NewRTCoinFilterer creates a new log filterer instance of RTCoin, bound to a specific deployed contract.
func NewRTCoinFilterer(address common.Address, filterer bind.ContractFilterer) (*RTCoinFilterer, error) {
	contract, err := bindRTCoin(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RTCoinFilterer{contract: contract}, nil
}

// bindRTCoin binds a generic wrapper to an already deployed contract.
func bindRTCoin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RTCoinABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RTCoin *RTCoinRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RTCoin.Contract.RTCoinCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RTCoin *RTCoinRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RTCoin.Contract.RTCoinTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RTCoin *RTCoinRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RTCoin.Contract.RTCoinTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RTCoin *RTCoinCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RTCoin.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RTCoin *RTCoinTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RTCoin.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RTCoin *RTCoinTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RTCoin.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_RTCoin *RTCoinCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RTCoin.contract.Call(opts, out, "admin")
	return *ret0, err
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_RTCoin *RTCoinSession) Admin() (common.Address, error) {
	return _RTCoin.Contract.Admin(&_RTCoin.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_RTCoin *RTCoinCallerSession) Admin() (common.Address, error) {
	return _RTCoin.Contract.Admin(&_RTCoin.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_RTCoin *RTCoinCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RTCoin.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_RTCoin *RTCoinSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _RTCoin.Contract.Allowance(&_RTCoin.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_RTCoin *RTCoinCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _RTCoin.Contract.Allowance(&_RTCoin.CallOpts, _owner, _spender)
}

// Allowed is a free data retrieval call binding the contract method 0x5c658165.
//
// Solidity: function allowed( address,  address) constant returns(uint256)
func (_RTCoin *RTCoinCaller) Allowed(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RTCoin.contract.Call(opts, out, "allowed", arg0, arg1)
	return *ret0, err
}

// Allowed is a free data retrieval call binding the contract method 0x5c658165.
//
// Solidity: function allowed( address,  address) constant returns(uint256)
func (_RTCoin *RTCoinSession) Allowed(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _RTCoin.Contract.Allowed(&_RTCoin.CallOpts, arg0, arg1)
}

// Allowed is a free data retrieval call binding the contract method 0x5c658165.
//
// Solidity: function allowed( address,  address) constant returns(uint256)
func (_RTCoin *RTCoinCallerSession) Allowed(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _RTCoin.Contract.Allowed(&_RTCoin.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_holder address) constant returns(uint256)
func (_RTCoin *RTCoinCaller) BalanceOf(opts *bind.CallOpts, _holder common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RTCoin.contract.Call(opts, out, "balanceOf", _holder)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_holder address) constant returns(uint256)
func (_RTCoin *RTCoinSession) BalanceOf(_holder common.Address) (*big.Int, error) {
	return _RTCoin.Contract.BalanceOf(&_RTCoin.CallOpts, _holder)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_holder address) constant returns(uint256)
func (_RTCoin *RTCoinCallerSession) BalanceOf(_holder common.Address) (*big.Int, error) {
	return _RTCoin.Contract.BalanceOf(&_RTCoin.CallOpts, _holder)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances( address) constant returns(uint256)
func (_RTCoin *RTCoinCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RTCoin.contract.Call(opts, out, "balances", arg0)
	return *ret0, err
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances( address) constant returns(uint256)
func (_RTCoin *RTCoinSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _RTCoin.Contract.Balances(&_RTCoin.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances( address) constant returns(uint256)
func (_RTCoin *RTCoinCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _RTCoin.Contract.Balances(&_RTCoin.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_RTCoin *RTCoinCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _RTCoin.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_RTCoin *RTCoinSession) Decimals() (uint8, error) {
	return _RTCoin.Contract.Decimals(&_RTCoin.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_RTCoin *RTCoinCallerSession) Decimals() (uint8, error) {
	return _RTCoin.Contract.Decimals(&_RTCoin.CallOpts)
}

// Frozen is a free data retrieval call binding the contract method 0x054f7d9c.
//
// Solidity: function frozen() constant returns(bool)
func (_RTCoin *RTCoinCaller) Frozen(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RTCoin.contract.Call(opts, out, "frozen")
	return *ret0, err
}

// Frozen is a free data retrieval call binding the contract method 0x054f7d9c.
//
// Solidity: function frozen() constant returns(bool)
func (_RTCoin *RTCoinSession) Frozen() (bool, error) {
	return _RTCoin.Contract.Frozen(&_RTCoin.CallOpts)
}

// Frozen is a free data retrieval call binding the contract method 0x054f7d9c.
//
// Solidity: function frozen() constant returns(bool)
func (_RTCoin *RTCoinCallerSession) Frozen() (bool, error) {
	return _RTCoin.Contract.Frozen(&_RTCoin.CallOpts)
}

// Moderators is a free data retrieval call binding the contract method 0x14d0f1ba.
//
// Solidity: function moderators( address) constant returns(bool)
func (_RTCoin *RTCoinCaller) Moderators(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RTCoin.contract.Call(opts, out, "moderators", arg0)
	return *ret0, err
}

// Moderators is a free data retrieval call binding the contract method 0x14d0f1ba.
//
// Solidity: function moderators( address) constant returns(bool)
func (_RTCoin *RTCoinSession) Moderators(arg0 common.Address) (bool, error) {
	return _RTCoin.Contract.Moderators(&_RTCoin.CallOpts, arg0)
}

// Moderators is a free data retrieval call binding the contract method 0x14d0f1ba.
//
// Solidity: function moderators( address) constant returns(bool)
func (_RTCoin *RTCoinCallerSession) Moderators(arg0 common.Address) (bool, error) {
	return _RTCoin.Contract.Moderators(&_RTCoin.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_RTCoin *RTCoinCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _RTCoin.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_RTCoin *RTCoinSession) Name() (string, error) {
	return _RTCoin.Contract.Name(&_RTCoin.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_RTCoin *RTCoinCallerSession) Name() (string, error) {
	return _RTCoin.Contract.Name(&_RTCoin.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RTCoin *RTCoinCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RTCoin.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RTCoin *RTCoinSession) Owner() (common.Address, error) {
	return _RTCoin.Contract.Owner(&_RTCoin.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RTCoin *RTCoinCallerSession) Owner() (common.Address, error) {
	return _RTCoin.Contract.Owner(&_RTCoin.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_RTCoin *RTCoinCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _RTCoin.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_RTCoin *RTCoinSession) Symbol() (string, error) {
	return _RTCoin.Contract.Symbol(&_RTCoin.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_RTCoin *RTCoinCallerSession) Symbol() (string, error) {
	return _RTCoin.Contract.Symbol(&_RTCoin.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_RTCoin *RTCoinCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RTCoin.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_RTCoin *RTCoinSession) TotalSupply() (*big.Int, error) {
	return _RTCoin.Contract.TotalSupply(&_RTCoin.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_RTCoin *RTCoinCallerSession) TotalSupply() (*big.Int, error) {
	return _RTCoin.Contract.TotalSupply(&_RTCoin.CallOpts)
}

// TransfersFrozen is a free data retrieval call binding the contract method 0xe45b8134.
//
// Solidity: function transfersFrozen() constant returns(bool)
func (_RTCoin *RTCoinCaller) TransfersFrozen(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RTCoin.contract.Call(opts, out, "transfersFrozen")
	return *ret0, err
}

// TransfersFrozen is a free data retrieval call binding the contract method 0xe45b8134.
//
// Solidity: function transfersFrozen() constant returns(bool)
func (_RTCoin *RTCoinSession) TransfersFrozen() (bool, error) {
	return _RTCoin.Contract.TransfersFrozen(&_RTCoin.CallOpts)
}

// TransfersFrozen is a free data retrieval call binding the contract method 0xe45b8134.
//
// Solidity: function transfersFrozen() constant returns(bool)
func (_RTCoin *RTCoinCallerSession) TransfersFrozen() (bool, error) {
	return _RTCoin.Contract.TransfersFrozen(&_RTCoin.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _amount uint256) returns(approved bool)
func (_RTCoin *RTCoinTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RTCoin.contract.Transact(opts, "approve", _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _amount uint256) returns(approved bool)
func (_RTCoin *RTCoinSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RTCoin.Contract.Approve(&_RTCoin.TransactOpts, _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _amount uint256) returns(approved bool)
func (_RTCoin *RTCoinTransactorSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RTCoin.Contract.Approve(&_RTCoin.TransactOpts, _spender, _amount)
}

// FreezeTransfers is a paid mutator transaction binding the contract method 0x01502460.
//
// Solidity: function freezeTransfers() returns(bool)
func (_RTCoin *RTCoinTransactor) FreezeTransfers(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RTCoin.contract.Transact(opts, "freezeTransfers")
}

// FreezeTransfers is a paid mutator transaction binding the contract method 0x01502460.
//
// Solidity: function freezeTransfers() returns(bool)
func (_RTCoin *RTCoinSession) FreezeTransfers() (*types.Transaction, error) {
	return _RTCoin.Contract.FreezeTransfers(&_RTCoin.TransactOpts)
}

// FreezeTransfers is a paid mutator transaction binding the contract method 0x01502460.
//
// Solidity: function freezeTransfers() returns(bool)
func (_RTCoin *RTCoinTransactorSession) FreezeTransfers() (*types.Transaction, error) {
	return _RTCoin.Contract.FreezeTransfers(&_RTCoin.TransactOpts)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(_newAdmin address) returns(bool)
func (_RTCoin *RTCoinTransactor) SetAdmin(opts *bind.TransactOpts, _newAdmin common.Address) (*types.Transaction, error) {
	return _RTCoin.contract.Transact(opts, "setAdmin", _newAdmin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(_newAdmin address) returns(bool)
func (_RTCoin *RTCoinSession) SetAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _RTCoin.Contract.SetAdmin(&_RTCoin.TransactOpts, _newAdmin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(_newAdmin address) returns(bool)
func (_RTCoin *RTCoinTransactorSession) SetAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _RTCoin.Contract.SetAdmin(&_RTCoin.TransactOpts, _newAdmin)
}

// ThawTransfers is a paid mutator transaction binding the contract method 0xce8e120a.
//
// Solidity: function thawTransfers() returns(bool)
func (_RTCoin *RTCoinTransactor) ThawTransfers(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RTCoin.contract.Transact(opts, "thawTransfers")
}

// ThawTransfers is a paid mutator transaction binding the contract method 0xce8e120a.
//
// Solidity: function thawTransfers() returns(bool)
func (_RTCoin *RTCoinSession) ThawTransfers() (*types.Transaction, error) {
	return _RTCoin.Contract.ThawTransfers(&_RTCoin.TransactOpts)
}

// ThawTransfers is a paid mutator transaction binding the contract method 0xce8e120a.
//
// Solidity: function thawTransfers() returns(bool)
func (_RTCoin *RTCoinTransactorSession) ThawTransfers() (*types.Transaction, error) {
	return _RTCoin.Contract.ThawTransfers(&_RTCoin.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_recipient address, _amount uint256) returns(transferred bool)
func (_RTCoin *RTCoinTransactor) Transfer(opts *bind.TransactOpts, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RTCoin.contract.Transact(opts, "transfer", _recipient, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_recipient address, _amount uint256) returns(transferred bool)
func (_RTCoin *RTCoinSession) Transfer(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RTCoin.Contract.Transfer(&_RTCoin.TransactOpts, _recipient, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_recipient address, _amount uint256) returns(transferred bool)
func (_RTCoin *RTCoinTransactorSession) Transfer(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RTCoin.Contract.Transfer(&_RTCoin.TransactOpts, _recipient, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_owner address, _recipient address, _amount uint256) returns(transferredFrom bool)
func (_RTCoin *RTCoinTransactor) TransferFrom(opts *bind.TransactOpts, _owner common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RTCoin.contract.Transact(opts, "transferFrom", _owner, _recipient, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_owner address, _recipient address, _amount uint256) returns(transferredFrom bool)
func (_RTCoin *RTCoinSession) TransferFrom(_owner common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RTCoin.Contract.TransferFrom(&_RTCoin.TransactOpts, _owner, _recipient, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_owner address, _recipient address, _amount uint256) returns(transferredFrom bool)
func (_RTCoin *RTCoinTransactorSession) TransferFrom(_owner common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RTCoin.Contract.TransferFrom(&_RTCoin.TransactOpts, _owner, _recipient, _amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns(bool)
func (_RTCoin *RTCoinTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _RTCoin.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns(bool)
func (_RTCoin *RTCoinSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _RTCoin.Contract.TransferOwnership(&_RTCoin.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns(bool)
func (_RTCoin *RTCoinTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _RTCoin.Contract.TransferOwnership(&_RTCoin.TransactOpts, _newOwner)
}

// RTCoinAdminSetIterator is returned from FilterAdminSet and is used to iterate over the raw logs and unpacked data for AdminSet events raised by the RTCoin contract.
type RTCoinAdminSetIterator struct {
	Event *RTCoinAdminSet // Event containing the contract specifics and raw log

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
func (it *RTCoinAdminSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RTCoinAdminSet)
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
		it.Event = new(RTCoinAdminSet)
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
func (it *RTCoinAdminSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RTCoinAdminSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RTCoinAdminSet represents a AdminSet event raised by the RTCoin contract.
type RTCoinAdminSet struct {
	Admin    common.Address
	AdminSet bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAdminSet is a free log retrieval operation binding the contract event 0xe68d2c359a771606c400cf8b87000cf5864010363d6a736e98f5047b7bbe18e9.
//
// Solidity: e AdminSet(_admin indexed address, _adminSet indexed bool)
func (_RTCoin *RTCoinFilterer) FilterAdminSet(opts *bind.FilterOpts, _admin []common.Address, _adminSet []bool) (*RTCoinAdminSetIterator, error) {

	var _adminRule []interface{}
	for _, _adminItem := range _admin {
		_adminRule = append(_adminRule, _adminItem)
	}
	var _adminSetRule []interface{}
	for _, _adminSetItem := range _adminSet {
		_adminSetRule = append(_adminSetRule, _adminSetItem)
	}

	logs, sub, err := _RTCoin.contract.FilterLogs(opts, "AdminSet", _adminRule, _adminSetRule)
	if err != nil {
		return nil, err
	}
	return &RTCoinAdminSetIterator{contract: _RTCoin.contract, event: "AdminSet", logs: logs, sub: sub}, nil
}

// WatchAdminSet is a free log subscription operation binding the contract event 0xe68d2c359a771606c400cf8b87000cf5864010363d6a736e98f5047b7bbe18e9.
//
// Solidity: e AdminSet(_admin indexed address, _adminSet indexed bool)
func (_RTCoin *RTCoinFilterer) WatchAdminSet(opts *bind.WatchOpts, sink chan<- *RTCoinAdminSet, _admin []common.Address, _adminSet []bool) (event.Subscription, error) {

	var _adminRule []interface{}
	for _, _adminItem := range _admin {
		_adminRule = append(_adminRule, _adminItem)
	}
	var _adminSetRule []interface{}
	for _, _adminSetItem := range _adminSet {
		_adminSetRule = append(_adminSetRule, _adminSetItem)
	}

	logs, sub, err := _RTCoin.contract.WatchLogs(opts, "AdminSet", _adminRule, _adminSetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RTCoinAdminSet)
				if err := _RTCoin.contract.UnpackLog(event, "AdminSet", log); err != nil {
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

// RTCoinApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the RTCoin contract.
type RTCoinApprovalIterator struct {
	Event *RTCoinApproval // Event containing the contract specifics and raw log

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
func (it *RTCoinApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RTCoinApproval)
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
		it.Event = new(RTCoinApproval)
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
func (it *RTCoinApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RTCoinApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RTCoinApproval represents a Approval event raised by the RTCoin contract.
type RTCoinApproval struct {
	Owner   common.Address
	Spender common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, _spender indexed address, _amount uint256)
func (_RTCoin *RTCoinFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _spender []common.Address) (*RTCoinApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _RTCoin.contract.FilterLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return &RTCoinApprovalIterator{contract: _RTCoin.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, _spender indexed address, _amount uint256)
func (_RTCoin *RTCoinFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *RTCoinApproval, _owner []common.Address, _spender []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _RTCoin.contract.WatchLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RTCoinApproval)
				if err := _RTCoin.contract.UnpackLog(event, "Approval", log); err != nil {
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

// RTCoinOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the RTCoin contract.
type RTCoinOwnershipTransferredIterator struct {
	Event *RTCoinOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RTCoinOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RTCoinOwnershipTransferred)
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
		it.Event = new(RTCoinOwnershipTransferred)
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
func (it *RTCoinOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RTCoinOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RTCoinOwnershipTransferred represents a OwnershipTransferred event raised by the RTCoin contract.
type RTCoinOwnershipTransferred struct {
	PreviousOwner        common.Address
	NewOwner             common.Address
	OwnershipTransferred bool
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x7fdc2a4b6eb39ec3363d710d188620bd1e97b3c434161f187b4d0dc0544faa58.
//
// Solidity: e OwnershipTransferred(_previousOwner indexed address, _newOwner indexed address, _ownershipTransferred indexed bool)
func (_RTCoin *RTCoinFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, _previousOwner []common.Address, _newOwner []common.Address, _ownershipTransferred []bool) (*RTCoinOwnershipTransferredIterator, error) {

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

	logs, sub, err := _RTCoin.contract.FilterLogs(opts, "OwnershipTransferred", _previousOwnerRule, _newOwnerRule, _ownershipTransferredRule)
	if err != nil {
		return nil, err
	}
	return &RTCoinOwnershipTransferredIterator{contract: _RTCoin.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x7fdc2a4b6eb39ec3363d710d188620bd1e97b3c434161f187b4d0dc0544faa58.
//
// Solidity: e OwnershipTransferred(_previousOwner indexed address, _newOwner indexed address, _ownershipTransferred indexed bool)
func (_RTCoin *RTCoinFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RTCoinOwnershipTransferred, _previousOwner []common.Address, _newOwner []common.Address, _ownershipTransferred []bool) (event.Subscription, error) {

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

	logs, sub, err := _RTCoin.contract.WatchLogs(opts, "OwnershipTransferred", _previousOwnerRule, _newOwnerRule, _ownershipTransferredRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RTCoinOwnershipTransferred)
				if err := _RTCoin.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// RTCoinTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the RTCoin contract.
type RTCoinTransferIterator struct {
	Event *RTCoinTransfer // Event containing the contract specifics and raw log

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
func (it *RTCoinTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RTCoinTransfer)
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
		it.Event = new(RTCoinTransfer)
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
func (it *RTCoinTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RTCoinTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RTCoinTransfer represents a Transfer event raised by the RTCoin contract.
type RTCoinTransfer struct {
	Sender    common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_sender indexed address, _recipient indexed address, _amount uint256)
func (_RTCoin *RTCoinFilterer) FilterTransfer(opts *bind.FilterOpts, _sender []common.Address, _recipient []common.Address) (*RTCoinTransferIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}
	var _recipientRule []interface{}
	for _, _recipientItem := range _recipient {
		_recipientRule = append(_recipientRule, _recipientItem)
	}

	logs, sub, err := _RTCoin.contract.FilterLogs(opts, "Transfer", _senderRule, _recipientRule)
	if err != nil {
		return nil, err
	}
	return &RTCoinTransferIterator{contract: _RTCoin.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_sender indexed address, _recipient indexed address, _amount uint256)
func (_RTCoin *RTCoinFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *RTCoinTransfer, _sender []common.Address, _recipient []common.Address) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}
	var _recipientRule []interface{}
	for _, _recipientItem := range _recipient {
		_recipientRule = append(_recipientRule, _recipientItem)
	}

	logs, sub, err := _RTCoin.contract.WatchLogs(opts, "Transfer", _senderRule, _recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RTCoinTransfer)
				if err := _RTCoin.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// RTCoinTransfersFrozenIterator is returned from FilterTransfersFrozen and is used to iterate over the raw logs and unpacked data for TransfersFrozen events raised by the RTCoin contract.
type RTCoinTransfersFrozenIterator struct {
	Event *RTCoinTransfersFrozen // Event containing the contract specifics and raw log

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
func (it *RTCoinTransfersFrozenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RTCoinTransfersFrozen)
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
		it.Event = new(RTCoinTransfersFrozen)
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
func (it *RTCoinTransfersFrozenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RTCoinTransfersFrozenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RTCoinTransfersFrozen represents a TransfersFrozen event raised by the RTCoin contract.
type RTCoinTransfersFrozen struct {
	TransfersFrozen bool
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTransfersFrozen is a free log retrieval operation binding the contract event 0xff7ea91c52ebd8c0d8018fdba50cb801e862f6795b1e17eeac882d4288b09340.
//
// Solidity: e TransfersFrozen(_transfersFrozen indexed bool)
func (_RTCoin *RTCoinFilterer) FilterTransfersFrozen(opts *bind.FilterOpts, _transfersFrozen []bool) (*RTCoinTransfersFrozenIterator, error) {

	var _transfersFrozenRule []interface{}
	for _, _transfersFrozenItem := range _transfersFrozen {
		_transfersFrozenRule = append(_transfersFrozenRule, _transfersFrozenItem)
	}

	logs, sub, err := _RTCoin.contract.FilterLogs(opts, "TransfersFrozen", _transfersFrozenRule)
	if err != nil {
		return nil, err
	}
	return &RTCoinTransfersFrozenIterator{contract: _RTCoin.contract, event: "TransfersFrozen", logs: logs, sub: sub}, nil
}

// WatchTransfersFrozen is a free log subscription operation binding the contract event 0xff7ea91c52ebd8c0d8018fdba50cb801e862f6795b1e17eeac882d4288b09340.
//
// Solidity: e TransfersFrozen(_transfersFrozen indexed bool)
func (_RTCoin *RTCoinFilterer) WatchTransfersFrozen(opts *bind.WatchOpts, sink chan<- *RTCoinTransfersFrozen, _transfersFrozen []bool) (event.Subscription, error) {

	var _transfersFrozenRule []interface{}
	for _, _transfersFrozenItem := range _transfersFrozen {
		_transfersFrozenRule = append(_transfersFrozenRule, _transfersFrozenItem)
	}

	logs, sub, err := _RTCoin.contract.WatchLogs(opts, "TransfersFrozen", _transfersFrozenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RTCoinTransfersFrozen)
				if err := _RTCoin.contract.UnpackLog(event, "TransfersFrozen", log); err != nil {
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

// RTCoinTransfersThawedIterator is returned from FilterTransfersThawed and is used to iterate over the raw logs and unpacked data for TransfersThawed events raised by the RTCoin contract.
type RTCoinTransfersThawedIterator struct {
	Event *RTCoinTransfersThawed // Event containing the contract specifics and raw log

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
func (it *RTCoinTransfersThawedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RTCoinTransfersThawed)
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
		it.Event = new(RTCoinTransfersThawed)
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
func (it *RTCoinTransfersThawedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RTCoinTransfersThawedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RTCoinTransfersThawed represents a TransfersThawed event raised by the RTCoin contract.
type RTCoinTransfersThawed struct {
	TransfersThawed bool
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTransfersThawed is a free log retrieval operation binding the contract event 0xb36ea4d45a6246e5ea6da988f57a5bf9a9022c85940cc6fe92dd9e45bf632cf6.
//
// Solidity: e TransfersThawed(_transfersThawed indexed bool)
func (_RTCoin *RTCoinFilterer) FilterTransfersThawed(opts *bind.FilterOpts, _transfersThawed []bool) (*RTCoinTransfersThawedIterator, error) {

	var _transfersThawedRule []interface{}
	for _, _transfersThawedItem := range _transfersThawed {
		_transfersThawedRule = append(_transfersThawedRule, _transfersThawedItem)
	}

	logs, sub, err := _RTCoin.contract.FilterLogs(opts, "TransfersThawed", _transfersThawedRule)
	if err != nil {
		return nil, err
	}
	return &RTCoinTransfersThawedIterator{contract: _RTCoin.contract, event: "TransfersThawed", logs: logs, sub: sub}, nil
}

// WatchTransfersThawed is a free log subscription operation binding the contract event 0xb36ea4d45a6246e5ea6da988f57a5bf9a9022c85940cc6fe92dd9e45bf632cf6.
//
// Solidity: e TransfersThawed(_transfersThawed indexed bool)
func (_RTCoin *RTCoinFilterer) WatchTransfersThawed(opts *bind.WatchOpts, sink chan<- *RTCoinTransfersThawed, _transfersThawed []bool) (event.Subscription, error) {

	var _transfersThawedRule []interface{}
	for _, _transfersThawedItem := range _transfersThawed {
		_transfersThawedRule = append(_transfersThawedRule, _transfersThawedItem)
	}

	logs, sub, err := _RTCoin.contract.WatchLogs(opts, "TransfersThawed", _transfersThawedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RTCoinTransfersThawed)
				if err := _RTCoin.contract.UnpackLog(event, "TransfersThawed", log); err != nil {
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
