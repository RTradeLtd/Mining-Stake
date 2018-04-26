// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package TestRouter

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

// TestRouterABI is the input ABI used to generate the binding from.
const TestRouterABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_stakers\",\"type\":\"address[]\"},{\"name\":\"_payouts\",\"type\":\"uint256[]\"}],\"name\":\"testRouteRequire\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_stakers\",\"type\":\"address[]\"},{\"name\":\"_payouts\",\"type\":\"uint256[]\"}],\"name\":\"testRouteNoRequire\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rtI\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_recipient\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"TestSent\",\"type\":\"event\"}]"

// TestRouterBin is the compiled bytecode used for deploying new contracts.
const TestRouterBin = `606060405260008054600160a060020a031916730994f9595d28429584bfb5fcbfea75b9c9ea2c24179055341561003557600080fd5b610442806100446000396000f30060606040526004361061003d5763ffffffff60e060020a6000350416631e48b9768114610042578063245a1961146100e5578063be72ab5914610174575b600080fd5b341561004d57600080fd5b6100d16004602481358181019083013580602081810201604051908101604052809392919081815260200183836020028082843782019150505050505091908035906020019082018035906020019080806020026020016040519081016040528093929190818152602001838360200280828437509496506101a395505050505050565b604051901515815260200160405180910390f35b34156100f057600080fd5b6100d16004602481358181019083013580602081810201604051908101604052809392919081815260200183836020028082843782019150505050505091908035906020019082018035906020019080806020026020016040519081016040528093929190818152602001838360200280828437509496506102df95505050505050565b341561017f57600080fd5b610187610407565b604051600160a060020a03909116815260200160405180910390f35b6000805b83518110156102d557600054600160a060020a03166323b872dd338684815181106101ce57fe5b906020019060200201518685815181106101e457fe5b9060200190602002015160405160e060020a63ffffffff8616028152600160a060020a0393841660048201529190921660248201526044810191909152606401602060405180830381600087803b151561023d57600080fd5b5af1151561024a57600080fd5b50505060405180519050151561025f57600080fd5b7fcdc6b93edab43c5d838cac6d14cdb9dab880b6aad17e3dc1c0609a459c7dd5de84828151811061028c57fe5b906020019060200201518483815181106102a257fe5b90602001906020020151604051600160a060020a03909216825260208201526040908101905180910390a16001016101a7565b5060019392505050565b6000805b83518110156102d557600054600160a060020a03166323b872dd3386848151811061030a57fe5b9060200190602002015186858151811061032057fe5b9060200190602002015160405160e060020a63ffffffff8616028152600160a060020a0393841660048201529190921660248201526044810191909152606401602060405180830381600087803b151561037957600080fd5b5af1151561038657600080fd5b50505060405180519050507fcdc6b93edab43c5d838cac6d14cdb9dab880b6aad17e3dc1c0609a459c7dd5de8482815181106103be57fe5b906020019060200201518483815181106103d457fe5b90602001906020020151604051600160a060020a03909216825260208201526040908101905180910390a16001016102e3565b600054600160a060020a0316815600a165627a7a72305820572bee45fe8acc84e2fde1905dfb63f6535d0feed8e0e1333fb41923d2cd4e4d0029`

// DeployTestRouter deploys a new Ethereum contract, binding an instance of TestRouter to it.
func DeployTestRouter(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TestRouter, error) {
	parsed, err := abi.JSON(strings.NewReader(TestRouterABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TestRouterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TestRouter{TestRouterCaller: TestRouterCaller{contract: contract}, TestRouterTransactor: TestRouterTransactor{contract: contract}, TestRouterFilterer: TestRouterFilterer{contract: contract}}, nil
}

// TestRouter is an auto generated Go binding around an Ethereum contract.
type TestRouter struct {
	TestRouterCaller     // Read-only binding to the contract
	TestRouterTransactor // Write-only binding to the contract
	TestRouterFilterer   // Log filterer for contract events
}

// TestRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestRouterSession struct {
	Contract     *TestRouter       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestRouterCallerSession struct {
	Contract *TestRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// TestRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestRouterTransactorSession struct {
	Contract     *TestRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// TestRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestRouterRaw struct {
	Contract *TestRouter // Generic contract binding to access the raw methods on
}

// TestRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestRouterCallerRaw struct {
	Contract *TestRouterCaller // Generic read-only contract binding to access the raw methods on
}

// TestRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestRouterTransactorRaw struct {
	Contract *TestRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestRouter creates a new instance of TestRouter, bound to a specific deployed contract.
func NewTestRouter(address common.Address, backend bind.ContractBackend) (*TestRouter, error) {
	contract, err := bindTestRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestRouter{TestRouterCaller: TestRouterCaller{contract: contract}, TestRouterTransactor: TestRouterTransactor{contract: contract}, TestRouterFilterer: TestRouterFilterer{contract: contract}}, nil
}

// NewTestRouterCaller creates a new read-only instance of TestRouter, bound to a specific deployed contract.
func NewTestRouterCaller(address common.Address, caller bind.ContractCaller) (*TestRouterCaller, error) {
	contract, err := bindTestRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestRouterCaller{contract: contract}, nil
}

// NewTestRouterTransactor creates a new write-only instance of TestRouter, bound to a specific deployed contract.
func NewTestRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*TestRouterTransactor, error) {
	contract, err := bindTestRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestRouterTransactor{contract: contract}, nil
}

// NewTestRouterFilterer creates a new log filterer instance of TestRouter, bound to a specific deployed contract.
func NewTestRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*TestRouterFilterer, error) {
	contract, err := bindTestRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestRouterFilterer{contract: contract}, nil
}

// bindTestRouter binds a generic wrapper to an already deployed contract.
func bindTestRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TestRouterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestRouter *TestRouterRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TestRouter.Contract.TestRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestRouter *TestRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestRouter.Contract.TestRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestRouter *TestRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestRouter.Contract.TestRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestRouter *TestRouterCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TestRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestRouter *TestRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestRouter *TestRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestRouter.Contract.contract.Transact(opts, method, params...)
}

// RtI is a free data retrieval call binding the contract method 0xbe72ab59.
//
// Solidity: function rtI() constant returns(address)
func (_TestRouter *TestRouterCaller) RtI(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TestRouter.contract.Call(opts, out, "rtI")
	return *ret0, err
}

// RtI is a free data retrieval call binding the contract method 0xbe72ab59.
//
// Solidity: function rtI() constant returns(address)
func (_TestRouter *TestRouterSession) RtI() (common.Address, error) {
	return _TestRouter.Contract.RtI(&_TestRouter.CallOpts)
}

// RtI is a free data retrieval call binding the contract method 0xbe72ab59.
//
// Solidity: function rtI() constant returns(address)
func (_TestRouter *TestRouterCallerSession) RtI() (common.Address, error) {
	return _TestRouter.Contract.RtI(&_TestRouter.CallOpts)
}

// TestRouteNoRequire is a paid mutator transaction binding the contract method 0x245a1961.
//
// Solidity: function testRouteNoRequire(_stakers address[], _payouts uint256[]) returns(bool)
func (_TestRouter *TestRouterTransactor) TestRouteNoRequire(opts *bind.TransactOpts, _stakers []common.Address, _payouts []*big.Int) (*types.Transaction, error) {
	return _TestRouter.contract.Transact(opts, "testRouteNoRequire", _stakers, _payouts)
}

// TestRouteNoRequire is a paid mutator transaction binding the contract method 0x245a1961.
//
// Solidity: function testRouteNoRequire(_stakers address[], _payouts uint256[]) returns(bool)
func (_TestRouter *TestRouterSession) TestRouteNoRequire(_stakers []common.Address, _payouts []*big.Int) (*types.Transaction, error) {
	return _TestRouter.Contract.TestRouteNoRequire(&_TestRouter.TransactOpts, _stakers, _payouts)
}

// TestRouteNoRequire is a paid mutator transaction binding the contract method 0x245a1961.
//
// Solidity: function testRouteNoRequire(_stakers address[], _payouts uint256[]) returns(bool)
func (_TestRouter *TestRouterTransactorSession) TestRouteNoRequire(_stakers []common.Address, _payouts []*big.Int) (*types.Transaction, error) {
	return _TestRouter.Contract.TestRouteNoRequire(&_TestRouter.TransactOpts, _stakers, _payouts)
}

// TestRouteRequire is a paid mutator transaction binding the contract method 0x1e48b976.
//
// Solidity: function testRouteRequire(_stakers address[], _payouts uint256[]) returns(bool)
func (_TestRouter *TestRouterTransactor) TestRouteRequire(opts *bind.TransactOpts, _stakers []common.Address, _payouts []*big.Int) (*types.Transaction, error) {
	return _TestRouter.contract.Transact(opts, "testRouteRequire", _stakers, _payouts)
}

// TestRouteRequire is a paid mutator transaction binding the contract method 0x1e48b976.
//
// Solidity: function testRouteRequire(_stakers address[], _payouts uint256[]) returns(bool)
func (_TestRouter *TestRouterSession) TestRouteRequire(_stakers []common.Address, _payouts []*big.Int) (*types.Transaction, error) {
	return _TestRouter.Contract.TestRouteRequire(&_TestRouter.TransactOpts, _stakers, _payouts)
}

// TestRouteRequire is a paid mutator transaction binding the contract method 0x1e48b976.
//
// Solidity: function testRouteRequire(_stakers address[], _payouts uint256[]) returns(bool)
func (_TestRouter *TestRouterTransactorSession) TestRouteRequire(_stakers []common.Address, _payouts []*big.Int) (*types.Transaction, error) {
	return _TestRouter.Contract.TestRouteRequire(&_TestRouter.TransactOpts, _stakers, _payouts)
}

// TestRouterTestSentIterator is returned from FilterTestSent and is used to iterate over the raw logs and unpacked data for TestSent events raised by the TestRouter contract.
type TestRouterTestSentIterator struct {
	Event *TestRouterTestSent // Event containing the contract specifics and raw log

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
func (it *TestRouterTestSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestRouterTestSent)
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
		it.Event = new(TestRouterTestSent)
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
func (it *TestRouterTestSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestRouterTestSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestRouterTestSent represents a TestSent event raised by the TestRouter contract.
type TestRouterTestSent struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTestSent is a free log retrieval operation binding the contract event 0xcdc6b93edab43c5d838cac6d14cdb9dab880b6aad17e3dc1c0609a459c7dd5de.
//
// Solidity: event TestSent(_recipient address, _amount uint256)
func (_TestRouter *TestRouterFilterer) FilterTestSent(opts *bind.FilterOpts) (*TestRouterTestSentIterator, error) {

	logs, sub, err := _TestRouter.contract.FilterLogs(opts, "TestSent")
	if err != nil {
		return nil, err
	}
	return &TestRouterTestSentIterator{contract: _TestRouter.contract, event: "TestSent", logs: logs, sub: sub}, nil
}

// WatchTestSent is a free log subscription operation binding the contract event 0xcdc6b93edab43c5d838cac6d14cdb9dab880b6aad17e3dc1c0609a459c7dd5de.
//
// Solidity: event TestSent(_recipient address, _amount uint256)
func (_TestRouter *TestRouterFilterer) WatchTestSent(opts *bind.WatchOpts, sink chan<- *TestRouterTestSent) (event.Subscription, error) {

	logs, sub, err := _TestRouter.contract.WatchLogs(opts, "TestSent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestRouterTestSent)
				if err := _TestRouter.contract.UnpackLog(event, "TestSent", log); err != nil {
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
