// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package accessors

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/thetatoken/thetasubchain/eth"
	"github.com/thetatoken/thetasubchain/eth/abi"
	"github.com/thetatoken/thetasubchain/eth/abi/bind"
	"github.com/thetatoken/theta/common"
	"github.com/thetatoken/thetasubchain/eth/core/types"
	"github.com/thetatoken/thetasubchain/eth/event"
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

// ProxyContractMetaData contains all meta data concerning the ProxyContract contract.
var ProxyContractMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Received\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"execute\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// ProxyContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ProxyContractMetaData.ABI instead.
var ProxyContractABI = ProxyContractMetaData.ABI

// ProxyContract is an auto generated Go binding around an Ethereum contract.
type ProxyContract struct {
	ProxyContractCaller     // Read-only binding to the contract
	ProxyContractTransactor // Write-only binding to the contract
	ProxyContractFilterer   // Log filterer for contract events
}

// ProxyContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProxyContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProxyContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProxyContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProxyContractSession struct {
	Contract     *ProxyContract    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProxyContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProxyContractCallerSession struct {
	Contract *ProxyContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ProxyContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProxyContractTransactorSession struct {
	Contract     *ProxyContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ProxyContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProxyContractRaw struct {
	Contract *ProxyContract // Generic contract binding to access the raw methods on
}

// ProxyContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProxyContractCallerRaw struct {
	Contract *ProxyContractCaller // Generic read-only contract binding to access the raw methods on
}

// ProxyContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProxyContractTransactorRaw struct {
	Contract *ProxyContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProxyContract creates a new instance of ProxyContract, bound to a specific deployed contract.
func NewProxyContract(address common.Address, backend bind.ContractBackend) (*ProxyContract, error) {
	contract, err := bindProxyContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ProxyContract{ProxyContractCaller: ProxyContractCaller{contract: contract}, ProxyContractTransactor: ProxyContractTransactor{contract: contract}, ProxyContractFilterer: ProxyContractFilterer{contract: contract}}, nil
}

// NewProxyContractCaller creates a new read-only instance of ProxyContract, bound to a specific deployed contract.
func NewProxyContractCaller(address common.Address, caller bind.ContractCaller) (*ProxyContractCaller, error) {
	contract, err := bindProxyContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProxyContractCaller{contract: contract}, nil
}

// NewProxyContractTransactor creates a new write-only instance of ProxyContract, bound to a specific deployed contract.
func NewProxyContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ProxyContractTransactor, error) {
	contract, err := bindProxyContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProxyContractTransactor{contract: contract}, nil
}

// NewProxyContractFilterer creates a new log filterer instance of ProxyContract, bound to a specific deployed contract.
func NewProxyContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ProxyContractFilterer, error) {
	contract, err := bindProxyContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProxyContractFilterer{contract: contract}, nil
}

// bindProxyContract binds a generic wrapper to an already deployed contract.
func bindProxyContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ProxyContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProxyContract *ProxyContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProxyContract.Contract.ProxyContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProxyContract *ProxyContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProxyContract.Contract.ProxyContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProxyContract *ProxyContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProxyContract.Contract.ProxyContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProxyContract *ProxyContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProxyContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProxyContract *ProxyContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProxyContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProxyContract *ProxyContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProxyContract.Contract.contract.Transact(opts, method, params...)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_ProxyContract *ProxyContractTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProxyContract.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_ProxyContract *ProxyContractSession) Deposit() (*types.Transaction, error) {
	return _ProxyContract.Contract.Deposit(&_ProxyContract.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_ProxyContract *ProxyContractTransactorSession) Deposit() (*types.Transaction, error) {
	return _ProxyContract.Contract.Deposit(&_ProxyContract.TransactOpts)
}

// Execute is a paid mutator transaction binding the contract method 0xf3541901.
//
// Solidity: function execute(address target, bytes data, uint256 value, uint256 nonce) payable returns(bool success, bytes returnData)
func (_ProxyContract *ProxyContractTransactor) Execute(opts *bind.TransactOpts, target common.Address, data []byte, value *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _ProxyContract.contract.Transact(opts, "execute", target, data, value, nonce)
}

// Execute is a paid mutator transaction binding the contract method 0xf3541901.
//
// Solidity: function execute(address target, bytes data, uint256 value, uint256 nonce) payable returns(bool success, bytes returnData)
func (_ProxyContract *ProxyContractSession) Execute(target common.Address, data []byte, value *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _ProxyContract.Contract.Execute(&_ProxyContract.TransactOpts, target, data, value, nonce)
}

// Execute is a paid mutator transaction binding the contract method 0xf3541901.
//
// Solidity: function execute(address target, bytes data, uint256 value, uint256 nonce) payable returns(bool success, bytes returnData)
func (_ProxyContract *ProxyContractTransactorSession) Execute(target common.Address, data []byte, value *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _ProxyContract.Contract.Execute(&_ProxyContract.TransactOpts, target, data, value, nonce)
}

// ProxyContractReceivedIterator is returned from FilterReceived and is used to iterate over the raw logs and unpacked data for Received events raised by the ProxyContract contract.
type ProxyContractReceivedIterator struct {
	Event *ProxyContractReceived // Event containing the contract specifics and raw log

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
func (it *ProxyContractReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProxyContractReceived)
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
		it.Event = new(ProxyContractReceived)
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
func (it *ProxyContractReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProxyContractReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProxyContractReceived represents a Received event raised by the ProxyContract contract.
type ProxyContractReceived struct {
	Arg0 common.Address
	Arg1 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterReceived is a free log retrieval operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address arg0, uint256 arg1)
func (_ProxyContract *ProxyContractFilterer) FilterReceived(opts *bind.FilterOpts) (*ProxyContractReceivedIterator, error) {

	logs, sub, err := _ProxyContract.contract.FilterLogs(opts, "Received")
	if err != nil {
		return nil, err
	}
	return &ProxyContractReceivedIterator{contract: _ProxyContract.contract, event: "Received", logs: logs, sub: sub}, nil
}

// WatchReceived is a free log subscription operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address arg0, uint256 arg1)
func (_ProxyContract *ProxyContractFilterer) WatchReceived(opts *bind.WatchOpts, sink chan<- *ProxyContractReceived) (event.Subscription, error) {

	logs, sub, err := _ProxyContract.contract.WatchLogs(opts, "Received")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProxyContractReceived)
				if err := _ProxyContract.contract.UnpackLog(event, "Received", log); err != nil {
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

// ParseReceived is a log parse operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address arg0, uint256 arg1)
func (_ProxyContract *ProxyContractFilterer) ParseReceived(log types.Log) (*ProxyContractReceived, error) {
	event := new(ProxyContractReceived)
	if err := _ProxyContract.contract.UnpackLog(event, "Received", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
