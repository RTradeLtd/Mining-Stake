// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package PaymentRouter

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

// PaymentRouterABI is the input ABI used to generate the binding from.
const PaymentRouterABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_stakers\",\"type\":\"address[]\"},{\"name\":\"_payments\",\"type\":\"uint256[]\"}],\"name\":\"routeEthReward\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewards\",\"outputs\":[{\"name\":\"ethRewarded\",\"type\":\"uint256\"},{\"name\":\"rtcRewarded\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rtcHotWallet\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MINSTAKE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"numStakes\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_stakers\",\"type\":\"address[]\"},{\"name\":\"_payments\",\"type\":\"uint256[]\"}],\"name\":\"routeRtcRewards\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"signUpFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ethUSD\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"kiloHashSecondPerRtc\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rtI\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakers\",\"outputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"rtcStaked\",\"type\":\"uint256\"},{\"name\":\"deposit\",\"type\":\"uint256\"},{\"name\":\"khSec\",\"type\":\"uint256\"},{\"name\":\"depositDate\",\"type\":\"uint256\"},{\"name\":\"releaseDate\",\"type\":\"uint256\"},{\"name\":\"id\",\"type\":\"uint256\"},{\"name\":\"enabled\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"locked\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"DEFAULTLOCKUPTIME\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stakerCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_staker\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"RtcReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_staker\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"EthReward\",\"type\":\"event\"}]"

// PaymentRouterBin is the compiled bytecode used for deploying new contracts.
const PaymentRouterBin = `60606040526756bc75e2d6310000600455730994f9595d28429584bfb5fcbfea75b9c9ea2c24600660016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550341561007057600080fd5b336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610dde806100bf6000396000f3006060604052600436106100e6576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063059d2bf4146100eb5780630700037d146101925780630bc0eadd146101e65780631d0a0e8d1461023b5780633467e9e1146102645780636198ce14146102b15780639278b58714610363578063ac48bd5a1461038c578063acd98bbf146103b5578063be72ab59146103de578063c8b6cbf714610433578063cf309012146104ea578063df9ec0f414610517578063dff6978714610540578063f851a44014610569578063ffa1ad74146105be575b600080fd5b6101786004808035906020019082018035906020019080806020026020016040519081016040528093929190818152602001838360200280828437820191505050505050919080359060200190820180359060200190808060200260200160405190810160405280939291908181526020018383602002808284378201915050505050509190505061064c565b604051808215151515815260200191505060405180910390f35b341561019d57600080fd5b6101c9600480803573ffffffffffffffffffffffffffffffffffffffff169060200190919050506108ae565b604051808381526020018281526020019250505060405180910390f35b34156101f157600080fd5b6101f96108d2565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b341561024657600080fd5b61024e6108f8565b6040518082815260200191505060405180910390f35b341561026f57600080fd5b61029b600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610905565b6040518082815260200191505060405180910390f35b34156102bc57600080fd5b6103496004808035906020019082018035906020019080806020026020016040519081016040528093929190818152602001838360200280828437820191505050505050919080359060200190820180359060200190808060200260200160405190810160405280939291908181526020018383602002808284378201915050505050509190505061091d565b604051808215151515815260200191505060405180910390f35b341561036e57600080fd5b610376610c4a565b6040518082815260200191505060405180910390f35b341561039757600080fd5b61039f610c50565b6040518082815260200191505060405180910390f35b34156103c057600080fd5b6103c8610c56565b6040518082815260200191505060405180910390f35b34156103e957600080fd5b6103f1610c5c565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b341561043e57600080fd5b610473600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091908035906020019091905050610c82565b604051808973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001888152602001878152602001868152602001858152602001848152602001838152602001821515151581526020019850505050505050505060405180910390f35b34156104f557600080fd5b6104fd610d13565b604051808215151515815260200191505060405180910390f35b341561052257600080fd5b61052a610d26565b6040518082815260200191505060405180910390f35b341561054b57600080fd5b610553610d2d565b6040518082815260200191505060405180910390f35b341561057457600080fd5b61057c610d33565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34156105c957600080fd5b6105d1610d58565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156106115780820151818401526020810190506105f6565b50505050905090810190601f16801561063e5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b60008060008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156106ac57600080fd5b6000341115156106bb57600080fd5b835185511415156106cb57600080fd5b600091505b84518210156108a25783828151811015156106e757fe5b90602001906020020151905061075f8160096000888681518110151561070957fe5b9060200190602002015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000154610d9190919063ffffffff16565b60096000878581518110151561077157fe5b9060200190602002015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600001819055507f196f95be2bd8e6aa49ac118195dec22c01e6adf2c34eaae9fefbb9efc8b78a7985838151811015156107eb57fe5b9060200190602002015182604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390a1848281518110151561084c57fe5b9060200190602002015173ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f19350505050151561089557600080fd5b81806001019250506106d0565b60019250505092915050565b60096020528060005260406000206000915090508060000154908060010154905082565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b68056bc75e2d6310000081565b600a6020528060005260406000206000915090505481565b60008060008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561097d57600080fd5b8351855114151561098d57600080fd5b600091505b8451821015610c3e5783828151811015156109a957fe5b906020019060200201519050610a21816009600088868151811015156109cb57fe5b9060200190602002015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010154610d9190919063ffffffff16565b600960008785815181101515610a3357fe5b9060200190602002015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101819055507fe0ba89edeae157ec385468cf95ff7ea61497f95bf3e0fe9637fa358aefdf7e288583815181101515610aad57fe5b9060200190602002015182604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390a1600660019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd338785815181101515610b4d57fe5b90602001906020020151846040518463ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050602060405180830381600087803b1515610c0f57600080fd5b5af11515610c1c57600080fd5b505050604051805190501515610c3157600080fd5b8180600101925050610992565b60019250505092915050565b60035481565b60015481565b60045481565b600660019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600860205281600052604060002081815481101515610c9d57fe5b9060005260206000209060080201600091509150508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154908060020154908060030154908060040154908060050154908060060154908060070160009054906101000a900460ff16905088565b600660009054906101000a900460ff1681565b6224ea0081565b60055481565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6040805190810160405280600981526020017f302e302e3162657461000000000000000000000000000000000000000000000081525081565b6000808284019050838110151515610da857600080fd5b80915050929150505600a165627a7a7230582065c94e89e08eefdc891ad64861c10bf27c24e335d4819ae2a4a4b2f2e0c6452b0029`

// DeployPaymentRouter deploys a new Ethereum contract, binding an instance of PaymentRouter to it.
func DeployPaymentRouter(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PaymentRouter, error) {
	parsed, err := abi.JSON(strings.NewReader(PaymentRouterABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PaymentRouterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PaymentRouter{PaymentRouterCaller: PaymentRouterCaller{contract: contract}, PaymentRouterTransactor: PaymentRouterTransactor{contract: contract}, PaymentRouterFilterer: PaymentRouterFilterer{contract: contract}}, nil
}

// PaymentRouter is an auto generated Go binding around an Ethereum contract.
type PaymentRouter struct {
	PaymentRouterCaller     // Read-only binding to the contract
	PaymentRouterTransactor // Write-only binding to the contract
	PaymentRouterFilterer   // Log filterer for contract events
}

// PaymentRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type PaymentRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymentRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PaymentRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymentRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PaymentRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymentRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PaymentRouterSession struct {
	Contract     *PaymentRouter    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PaymentRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PaymentRouterCallerSession struct {
	Contract *PaymentRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// PaymentRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PaymentRouterTransactorSession struct {
	Contract     *PaymentRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// PaymentRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type PaymentRouterRaw struct {
	Contract *PaymentRouter // Generic contract binding to access the raw methods on
}

// PaymentRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PaymentRouterCallerRaw struct {
	Contract *PaymentRouterCaller // Generic read-only contract binding to access the raw methods on
}

// PaymentRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PaymentRouterTransactorRaw struct {
	Contract *PaymentRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPaymentRouter creates a new instance of PaymentRouter, bound to a specific deployed contract.
func NewPaymentRouter(address common.Address, backend bind.ContractBackend) (*PaymentRouter, error) {
	contract, err := bindPaymentRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PaymentRouter{PaymentRouterCaller: PaymentRouterCaller{contract: contract}, PaymentRouterTransactor: PaymentRouterTransactor{contract: contract}, PaymentRouterFilterer: PaymentRouterFilterer{contract: contract}}, nil
}

// NewPaymentRouterCaller creates a new read-only instance of PaymentRouter, bound to a specific deployed contract.
func NewPaymentRouterCaller(address common.Address, caller bind.ContractCaller) (*PaymentRouterCaller, error) {
	contract, err := bindPaymentRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PaymentRouterCaller{contract: contract}, nil
}

// NewPaymentRouterTransactor creates a new write-only instance of PaymentRouter, bound to a specific deployed contract.
func NewPaymentRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*PaymentRouterTransactor, error) {
	contract, err := bindPaymentRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PaymentRouterTransactor{contract: contract}, nil
}

// NewPaymentRouterFilterer creates a new log filterer instance of PaymentRouter, bound to a specific deployed contract.
func NewPaymentRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*PaymentRouterFilterer, error) {
	contract, err := bindPaymentRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PaymentRouterFilterer{contract: contract}, nil
}

// bindPaymentRouter binds a generic wrapper to an already deployed contract.
func bindPaymentRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PaymentRouterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PaymentRouter *PaymentRouterRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PaymentRouter.Contract.PaymentRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PaymentRouter *PaymentRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PaymentRouter.Contract.PaymentRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PaymentRouter *PaymentRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PaymentRouter.Contract.PaymentRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PaymentRouter *PaymentRouterCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PaymentRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PaymentRouter *PaymentRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PaymentRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PaymentRouter *PaymentRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PaymentRouter.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTLOCKUPTIME is a free data retrieval call binding the contract method 0xdf9ec0f4.
//
// Solidity: function DEFAULTLOCKUPTIME() constant returns(uint256)
func (_PaymentRouter *PaymentRouterCaller) DEFAULTLOCKUPTIME(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PaymentRouter.contract.Call(opts, out, "DEFAULTLOCKUPTIME")
	return *ret0, err
}

// DEFAULTLOCKUPTIME is a free data retrieval call binding the contract method 0xdf9ec0f4.
//
// Solidity: function DEFAULTLOCKUPTIME() constant returns(uint256)
func (_PaymentRouter *PaymentRouterSession) DEFAULTLOCKUPTIME() (*big.Int, error) {
	return _PaymentRouter.Contract.DEFAULTLOCKUPTIME(&_PaymentRouter.CallOpts)
}

// DEFAULTLOCKUPTIME is a free data retrieval call binding the contract method 0xdf9ec0f4.
//
// Solidity: function DEFAULTLOCKUPTIME() constant returns(uint256)
func (_PaymentRouter *PaymentRouterCallerSession) DEFAULTLOCKUPTIME() (*big.Int, error) {
	return _PaymentRouter.Contract.DEFAULTLOCKUPTIME(&_PaymentRouter.CallOpts)
}

// MINSTAKE is a free data retrieval call binding the contract method 0x1d0a0e8d.
//
// Solidity: function MINSTAKE() constant returns(uint256)
func (_PaymentRouter *PaymentRouterCaller) MINSTAKE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PaymentRouter.contract.Call(opts, out, "MINSTAKE")
	return *ret0, err
}

// MINSTAKE is a free data retrieval call binding the contract method 0x1d0a0e8d.
//
// Solidity: function MINSTAKE() constant returns(uint256)
func (_PaymentRouter *PaymentRouterSession) MINSTAKE() (*big.Int, error) {
	return _PaymentRouter.Contract.MINSTAKE(&_PaymentRouter.CallOpts)
}

// MINSTAKE is a free data retrieval call binding the contract method 0x1d0a0e8d.
//
// Solidity: function MINSTAKE() constant returns(uint256)
func (_PaymentRouter *PaymentRouterCallerSession) MINSTAKE() (*big.Int, error) {
	return _PaymentRouter.Contract.MINSTAKE(&_PaymentRouter.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() constant returns(string)
func (_PaymentRouter *PaymentRouterCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _PaymentRouter.contract.Call(opts, out, "VERSION")
	return *ret0, err
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() constant returns(string)
func (_PaymentRouter *PaymentRouterSession) VERSION() (string, error) {
	return _PaymentRouter.Contract.VERSION(&_PaymentRouter.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() constant returns(string)
func (_PaymentRouter *PaymentRouterCallerSession) VERSION() (string, error) {
	return _PaymentRouter.Contract.VERSION(&_PaymentRouter.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_PaymentRouter *PaymentRouterCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PaymentRouter.contract.Call(opts, out, "admin")
	return *ret0, err
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_PaymentRouter *PaymentRouterSession) Admin() (common.Address, error) {
	return _PaymentRouter.Contract.Admin(&_PaymentRouter.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_PaymentRouter *PaymentRouterCallerSession) Admin() (common.Address, error) {
	return _PaymentRouter.Contract.Admin(&_PaymentRouter.CallOpts)
}

// EthUSD is a free data retrieval call binding the contract method 0xac48bd5a.
//
// Solidity: function ethUSD() constant returns(uint256)
func (_PaymentRouter *PaymentRouterCaller) EthUSD(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PaymentRouter.contract.Call(opts, out, "ethUSD")
	return *ret0, err
}

// EthUSD is a free data retrieval call binding the contract method 0xac48bd5a.
//
// Solidity: function ethUSD() constant returns(uint256)
func (_PaymentRouter *PaymentRouterSession) EthUSD() (*big.Int, error) {
	return _PaymentRouter.Contract.EthUSD(&_PaymentRouter.CallOpts)
}

// EthUSD is a free data retrieval call binding the contract method 0xac48bd5a.
//
// Solidity: function ethUSD() constant returns(uint256)
func (_PaymentRouter *PaymentRouterCallerSession) EthUSD() (*big.Int, error) {
	return _PaymentRouter.Contract.EthUSD(&_PaymentRouter.CallOpts)
}

// KiloHashSecondPerRtc is a free data retrieval call binding the contract method 0xacd98bbf.
//
// Solidity: function kiloHashSecondPerRtc() constant returns(uint256)
func (_PaymentRouter *PaymentRouterCaller) KiloHashSecondPerRtc(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PaymentRouter.contract.Call(opts, out, "kiloHashSecondPerRtc")
	return *ret0, err
}

// KiloHashSecondPerRtc is a free data retrieval call binding the contract method 0xacd98bbf.
//
// Solidity: function kiloHashSecondPerRtc() constant returns(uint256)
func (_PaymentRouter *PaymentRouterSession) KiloHashSecondPerRtc() (*big.Int, error) {
	return _PaymentRouter.Contract.KiloHashSecondPerRtc(&_PaymentRouter.CallOpts)
}

// KiloHashSecondPerRtc is a free data retrieval call binding the contract method 0xacd98bbf.
//
// Solidity: function kiloHashSecondPerRtc() constant returns(uint256)
func (_PaymentRouter *PaymentRouterCallerSession) KiloHashSecondPerRtc() (*big.Int, error) {
	return _PaymentRouter.Contract.KiloHashSecondPerRtc(&_PaymentRouter.CallOpts)
}

// Locked is a free data retrieval call binding the contract method 0xcf309012.
//
// Solidity: function locked() constant returns(bool)
func (_PaymentRouter *PaymentRouterCaller) Locked(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PaymentRouter.contract.Call(opts, out, "locked")
	return *ret0, err
}

// Locked is a free data retrieval call binding the contract method 0xcf309012.
//
// Solidity: function locked() constant returns(bool)
func (_PaymentRouter *PaymentRouterSession) Locked() (bool, error) {
	return _PaymentRouter.Contract.Locked(&_PaymentRouter.CallOpts)
}

// Locked is a free data retrieval call binding the contract method 0xcf309012.
//
// Solidity: function locked() constant returns(bool)
func (_PaymentRouter *PaymentRouterCallerSession) Locked() (bool, error) {
	return _PaymentRouter.Contract.Locked(&_PaymentRouter.CallOpts)
}

// NumStakes is a free data retrieval call binding the contract method 0x3467e9e1.
//
// Solidity: function numStakes( address) constant returns(uint256)
func (_PaymentRouter *PaymentRouterCaller) NumStakes(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PaymentRouter.contract.Call(opts, out, "numStakes", arg0)
	return *ret0, err
}

// NumStakes is a free data retrieval call binding the contract method 0x3467e9e1.
//
// Solidity: function numStakes( address) constant returns(uint256)
func (_PaymentRouter *PaymentRouterSession) NumStakes(arg0 common.Address) (*big.Int, error) {
	return _PaymentRouter.Contract.NumStakes(&_PaymentRouter.CallOpts, arg0)
}

// NumStakes is a free data retrieval call binding the contract method 0x3467e9e1.
//
// Solidity: function numStakes( address) constant returns(uint256)
func (_PaymentRouter *PaymentRouterCallerSession) NumStakes(arg0 common.Address) (*big.Int, error) {
	return _PaymentRouter.Contract.NumStakes(&_PaymentRouter.CallOpts, arg0)
}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards( address) constant returns(ethRewarded uint256, rtcRewarded uint256)
func (_PaymentRouter *PaymentRouterCaller) Rewards(opts *bind.CallOpts, arg0 common.Address) (struct {
	EthRewarded *big.Int
	RtcRewarded *big.Int
}, error) {
	ret := new(struct {
		EthRewarded *big.Int
		RtcRewarded *big.Int
	})
	out := ret
	err := _PaymentRouter.contract.Call(opts, out, "rewards", arg0)
	return *ret, err
}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards( address) constant returns(ethRewarded uint256, rtcRewarded uint256)
func (_PaymentRouter *PaymentRouterSession) Rewards(arg0 common.Address) (struct {
	EthRewarded *big.Int
	RtcRewarded *big.Int
}, error) {
	return _PaymentRouter.Contract.Rewards(&_PaymentRouter.CallOpts, arg0)
}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards( address) constant returns(ethRewarded uint256, rtcRewarded uint256)
func (_PaymentRouter *PaymentRouterCallerSession) Rewards(arg0 common.Address) (struct {
	EthRewarded *big.Int
	RtcRewarded *big.Int
}, error) {
	return _PaymentRouter.Contract.Rewards(&_PaymentRouter.CallOpts, arg0)
}

// RtI is a free data retrieval call binding the contract method 0xbe72ab59.
//
// Solidity: function rtI() constant returns(address)
func (_PaymentRouter *PaymentRouterCaller) RtI(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PaymentRouter.contract.Call(opts, out, "rtI")
	return *ret0, err
}

// RtI is a free data retrieval call binding the contract method 0xbe72ab59.
//
// Solidity: function rtI() constant returns(address)
func (_PaymentRouter *PaymentRouterSession) RtI() (common.Address, error) {
	return _PaymentRouter.Contract.RtI(&_PaymentRouter.CallOpts)
}

// RtI is a free data retrieval call binding the contract method 0xbe72ab59.
//
// Solidity: function rtI() constant returns(address)
func (_PaymentRouter *PaymentRouterCallerSession) RtI() (common.Address, error) {
	return _PaymentRouter.Contract.RtI(&_PaymentRouter.CallOpts)
}

// RtcHotWallet is a free data retrieval call binding the contract method 0x0bc0eadd.
//
// Solidity: function rtcHotWallet() constant returns(address)
func (_PaymentRouter *PaymentRouterCaller) RtcHotWallet(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PaymentRouter.contract.Call(opts, out, "rtcHotWallet")
	return *ret0, err
}

// RtcHotWallet is a free data retrieval call binding the contract method 0x0bc0eadd.
//
// Solidity: function rtcHotWallet() constant returns(address)
func (_PaymentRouter *PaymentRouterSession) RtcHotWallet() (common.Address, error) {
	return _PaymentRouter.Contract.RtcHotWallet(&_PaymentRouter.CallOpts)
}

// RtcHotWallet is a free data retrieval call binding the contract method 0x0bc0eadd.
//
// Solidity: function rtcHotWallet() constant returns(address)
func (_PaymentRouter *PaymentRouterCallerSession) RtcHotWallet() (common.Address, error) {
	return _PaymentRouter.Contract.RtcHotWallet(&_PaymentRouter.CallOpts)
}

// SignUpFee is a free data retrieval call binding the contract method 0x9278b587.
//
// Solidity: function signUpFee() constant returns(uint256)
func (_PaymentRouter *PaymentRouterCaller) SignUpFee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PaymentRouter.contract.Call(opts, out, "signUpFee")
	return *ret0, err
}

// SignUpFee is a free data retrieval call binding the contract method 0x9278b587.
//
// Solidity: function signUpFee() constant returns(uint256)
func (_PaymentRouter *PaymentRouterSession) SignUpFee() (*big.Int, error) {
	return _PaymentRouter.Contract.SignUpFee(&_PaymentRouter.CallOpts)
}

// SignUpFee is a free data retrieval call binding the contract method 0x9278b587.
//
// Solidity: function signUpFee() constant returns(uint256)
func (_PaymentRouter *PaymentRouterCallerSession) SignUpFee() (*big.Int, error) {
	return _PaymentRouter.Contract.SignUpFee(&_PaymentRouter.CallOpts)
}

// StakerCount is a free data retrieval call binding the contract method 0xdff69787.
//
// Solidity: function stakerCount() constant returns(uint256)
func (_PaymentRouter *PaymentRouterCaller) StakerCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PaymentRouter.contract.Call(opts, out, "stakerCount")
	return *ret0, err
}

// StakerCount is a free data retrieval call binding the contract method 0xdff69787.
//
// Solidity: function stakerCount() constant returns(uint256)
func (_PaymentRouter *PaymentRouterSession) StakerCount() (*big.Int, error) {
	return _PaymentRouter.Contract.StakerCount(&_PaymentRouter.CallOpts)
}

// StakerCount is a free data retrieval call binding the contract method 0xdff69787.
//
// Solidity: function stakerCount() constant returns(uint256)
func (_PaymentRouter *PaymentRouterCallerSession) StakerCount() (*big.Int, error) {
	return _PaymentRouter.Contract.StakerCount(&_PaymentRouter.CallOpts)
}

// Stakers is a free data retrieval call binding the contract method 0xc8b6cbf7.
//
// Solidity: function stakers( address,  uint256) constant returns(addr address, rtcStaked uint256, deposit uint256, khSec uint256, depositDate uint256, releaseDate uint256, id uint256, enabled bool)
func (_PaymentRouter *PaymentRouterCaller) Stakers(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	Addr        common.Address
	RtcStaked   *big.Int
	Deposit     *big.Int
	KhSec       *big.Int
	DepositDate *big.Int
	ReleaseDate *big.Int
	Id          *big.Int
	Enabled     bool
}, error) {
	ret := new(struct {
		Addr        common.Address
		RtcStaked   *big.Int
		Deposit     *big.Int
		KhSec       *big.Int
		DepositDate *big.Int
		ReleaseDate *big.Int
		Id          *big.Int
		Enabled     bool
	})
	out := ret
	err := _PaymentRouter.contract.Call(opts, out, "stakers", arg0, arg1)
	return *ret, err
}

// Stakers is a free data retrieval call binding the contract method 0xc8b6cbf7.
//
// Solidity: function stakers( address,  uint256) constant returns(addr address, rtcStaked uint256, deposit uint256, khSec uint256, depositDate uint256, releaseDate uint256, id uint256, enabled bool)
func (_PaymentRouter *PaymentRouterSession) Stakers(arg0 common.Address, arg1 *big.Int) (struct {
	Addr        common.Address
	RtcStaked   *big.Int
	Deposit     *big.Int
	KhSec       *big.Int
	DepositDate *big.Int
	ReleaseDate *big.Int
	Id          *big.Int
	Enabled     bool
}, error) {
	return _PaymentRouter.Contract.Stakers(&_PaymentRouter.CallOpts, arg0, arg1)
}

// Stakers is a free data retrieval call binding the contract method 0xc8b6cbf7.
//
// Solidity: function stakers( address,  uint256) constant returns(addr address, rtcStaked uint256, deposit uint256, khSec uint256, depositDate uint256, releaseDate uint256, id uint256, enabled bool)
func (_PaymentRouter *PaymentRouterCallerSession) Stakers(arg0 common.Address, arg1 *big.Int) (struct {
	Addr        common.Address
	RtcStaked   *big.Int
	Deposit     *big.Int
	KhSec       *big.Int
	DepositDate *big.Int
	ReleaseDate *big.Int
	Id          *big.Int
	Enabled     bool
}, error) {
	return _PaymentRouter.Contract.Stakers(&_PaymentRouter.CallOpts, arg0, arg1)
}

// RouteEthReward is a paid mutator transaction binding the contract method 0x059d2bf4.
//
// Solidity: function routeEthReward(_stakers address[], _payments uint256[]) returns(bool)
func (_PaymentRouter *PaymentRouterTransactor) RouteEthReward(opts *bind.TransactOpts, _stakers []common.Address, _payments []*big.Int) (*types.Transaction, error) {
	return _PaymentRouter.contract.Transact(opts, "routeEthReward", _stakers, _payments)
}

// RouteEthReward is a paid mutator transaction binding the contract method 0x059d2bf4.
//
// Solidity: function routeEthReward(_stakers address[], _payments uint256[]) returns(bool)
func (_PaymentRouter *PaymentRouterSession) RouteEthReward(_stakers []common.Address, _payments []*big.Int) (*types.Transaction, error) {
	return _PaymentRouter.Contract.RouteEthReward(&_PaymentRouter.TransactOpts, _stakers, _payments)
}

// RouteEthReward is a paid mutator transaction binding the contract method 0x059d2bf4.
//
// Solidity: function routeEthReward(_stakers address[], _payments uint256[]) returns(bool)
func (_PaymentRouter *PaymentRouterTransactorSession) RouteEthReward(_stakers []common.Address, _payments []*big.Int) (*types.Transaction, error) {
	return _PaymentRouter.Contract.RouteEthReward(&_PaymentRouter.TransactOpts, _stakers, _payments)
}

// RouteRtcRewards is a paid mutator transaction binding the contract method 0x6198ce14.
//
// Solidity: function routeRtcRewards(_stakers address[], _payments uint256[]) returns(bool)
func (_PaymentRouter *PaymentRouterTransactor) RouteRtcRewards(opts *bind.TransactOpts, _stakers []common.Address, _payments []*big.Int) (*types.Transaction, error) {
	return _PaymentRouter.contract.Transact(opts, "routeRtcRewards", _stakers, _payments)
}

// RouteRtcRewards is a paid mutator transaction binding the contract method 0x6198ce14.
//
// Solidity: function routeRtcRewards(_stakers address[], _payments uint256[]) returns(bool)
func (_PaymentRouter *PaymentRouterSession) RouteRtcRewards(_stakers []common.Address, _payments []*big.Int) (*types.Transaction, error) {
	return _PaymentRouter.Contract.RouteRtcRewards(&_PaymentRouter.TransactOpts, _stakers, _payments)
}

// RouteRtcRewards is a paid mutator transaction binding the contract method 0x6198ce14.
//
// Solidity: function routeRtcRewards(_stakers address[], _payments uint256[]) returns(bool)
func (_PaymentRouter *PaymentRouterTransactorSession) RouteRtcRewards(_stakers []common.Address, _payments []*big.Int) (*types.Transaction, error) {
	return _PaymentRouter.Contract.RouteRtcRewards(&_PaymentRouter.TransactOpts, _stakers, _payments)
}

// PaymentRouterEthRewardIterator is returned from FilterEthReward and is used to iterate over the raw logs and unpacked data for EthReward events raised by the PaymentRouter contract.
type PaymentRouterEthRewardIterator struct {
	Event *PaymentRouterEthReward // Event containing the contract specifics and raw log

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
func (it *PaymentRouterEthRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentRouterEthReward)
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
		it.Event = new(PaymentRouterEthReward)
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
func (it *PaymentRouterEthRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentRouterEthRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentRouterEthReward represents a EthReward event raised by the PaymentRouter contract.
type PaymentRouterEthReward struct {
	Staker common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEthReward is a free log retrieval operation binding the contract event 0x196f95be2bd8e6aa49ac118195dec22c01e6adf2c34eaae9fefbb9efc8b78a79.
//
// Solidity: event EthReward(_staker address, _amount uint256)
func (_PaymentRouter *PaymentRouterFilterer) FilterEthReward(opts *bind.FilterOpts) (*PaymentRouterEthRewardIterator, error) {

	logs, sub, err := _PaymentRouter.contract.FilterLogs(opts, "EthReward")
	if err != nil {
		return nil, err
	}
	return &PaymentRouterEthRewardIterator{contract: _PaymentRouter.contract, event: "EthReward", logs: logs, sub: sub}, nil
}

// WatchEthReward is a free log subscription operation binding the contract event 0x196f95be2bd8e6aa49ac118195dec22c01e6adf2c34eaae9fefbb9efc8b78a79.
//
// Solidity: event EthReward(_staker address, _amount uint256)
func (_PaymentRouter *PaymentRouterFilterer) WatchEthReward(opts *bind.WatchOpts, sink chan<- *PaymentRouterEthReward) (event.Subscription, error) {

	logs, sub, err := _PaymentRouter.contract.WatchLogs(opts, "EthReward")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentRouterEthReward)
				if err := _PaymentRouter.contract.UnpackLog(event, "EthReward", log); err != nil {
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

// PaymentRouterRtcRewardIterator is returned from FilterRtcReward and is used to iterate over the raw logs and unpacked data for RtcReward events raised by the PaymentRouter contract.
type PaymentRouterRtcRewardIterator struct {
	Event *PaymentRouterRtcReward // Event containing the contract specifics and raw log

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
func (it *PaymentRouterRtcRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentRouterRtcReward)
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
		it.Event = new(PaymentRouterRtcReward)
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
func (it *PaymentRouterRtcRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentRouterRtcRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentRouterRtcReward represents a RtcReward event raised by the PaymentRouter contract.
type PaymentRouterRtcReward struct {
	Staker common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRtcReward is a free log retrieval operation binding the contract event 0xe0ba89edeae157ec385468cf95ff7ea61497f95bf3e0fe9637fa358aefdf7e28.
//
// Solidity: event RtcReward(_staker address, _amount uint256)
func (_PaymentRouter *PaymentRouterFilterer) FilterRtcReward(opts *bind.FilterOpts) (*PaymentRouterRtcRewardIterator, error) {

	logs, sub, err := _PaymentRouter.contract.FilterLogs(opts, "RtcReward")
	if err != nil {
		return nil, err
	}
	return &PaymentRouterRtcRewardIterator{contract: _PaymentRouter.contract, event: "RtcReward", logs: logs, sub: sub}, nil
}

// WatchRtcReward is a free log subscription operation binding the contract event 0xe0ba89edeae157ec385468cf95ff7ea61497f95bf3e0fe9637fa358aefdf7e28.
//
// Solidity: event RtcReward(_staker address, _amount uint256)
func (_PaymentRouter *PaymentRouterFilterer) WatchRtcReward(opts *bind.WatchOpts, sink chan<- *PaymentRouterRtcReward) (event.Subscription, error) {

	logs, sub, err := _PaymentRouter.contract.WatchLogs(opts, "RtcReward")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentRouterRtcReward)
				if err := _PaymentRouter.contract.UnpackLog(event, "RtcReward", log); err != nil {
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
