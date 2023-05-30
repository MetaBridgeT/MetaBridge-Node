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

// BridgeCoreExternalCallData is an auto generated low-level Go binding around an user-defined struct.
type BridgeCoreExternalCallData struct {
	Chainid         *big.Int
	Caller          common.Address
	ContractAddress common.Address
	Data            []byte
	Nonce           *big.Int
	Height          *big.Int
	Value           *big.Int
}

// BridgeCoreMetaData contains all meta data concerning the BridgeCore contract.
var BridgeCoreMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"ExternalCall\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"bridgeCall\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"non\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"getExternalCall\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"chainid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structBridgeCore.ExternalCallData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// BridgeCoreABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeCoreMetaData.ABI instead.
var BridgeCoreABI = BridgeCoreMetaData.ABI

// BridgeCore is an auto generated Go binding around an Ethereum contract.
type BridgeCore struct {
	BridgeCoreCaller     // Read-only binding to the contract
	BridgeCoreTransactor // Write-only binding to the contract
	BridgeCoreFilterer   // Log filterer for contract events
}

// BridgeCoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeCoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeCoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeCoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeCoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeCoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeCoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeCoreSession struct {
	Contract     *BridgeCore       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeCoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeCoreCallerSession struct {
	Contract *BridgeCoreCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// BridgeCoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeCoreTransactorSession struct {
	Contract     *BridgeCoreTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BridgeCoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeCoreRaw struct {
	Contract *BridgeCore // Generic contract binding to access the raw methods on
}

// BridgeCoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeCoreCallerRaw struct {
	Contract *BridgeCoreCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeCoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeCoreTransactorRaw struct {
	Contract *BridgeCoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgeCore creates a new instance of BridgeCore, bound to a specific deployed contract.
func NewBridgeCore(address common.Address, backend bind.ContractBackend) (*BridgeCore, error) {
	contract, err := bindBridgeCore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BridgeCore{BridgeCoreCaller: BridgeCoreCaller{contract: contract}, BridgeCoreTransactor: BridgeCoreTransactor{contract: contract}, BridgeCoreFilterer: BridgeCoreFilterer{contract: contract}}, nil
}

// NewBridgeCoreCaller creates a new read-only instance of BridgeCore, bound to a specific deployed contract.
func NewBridgeCoreCaller(address common.Address, caller bind.ContractCaller) (*BridgeCoreCaller, error) {
	contract, err := bindBridgeCore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeCoreCaller{contract: contract}, nil
}

// NewBridgeCoreTransactor creates a new write-only instance of BridgeCore, bound to a specific deployed contract.
func NewBridgeCoreTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeCoreTransactor, error) {
	contract, err := bindBridgeCore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeCoreTransactor{contract: contract}, nil
}

// NewBridgeCoreFilterer creates a new log filterer instance of BridgeCore, bound to a specific deployed contract.
func NewBridgeCoreFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeCoreFilterer, error) {
	contract, err := bindBridgeCore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeCoreFilterer{contract: contract}, nil
}

// bindBridgeCore binds a generic wrapper to an already deployed contract.
func bindBridgeCore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BridgeCoreMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeCore *BridgeCoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeCore.Contract.BridgeCoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeCore *BridgeCoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeCore.Contract.BridgeCoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeCore *BridgeCoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeCore.Contract.BridgeCoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeCore *BridgeCoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeCore.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeCore *BridgeCoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeCore.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeCore *BridgeCoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeCore.Contract.contract.Transact(opts, method, params...)
}

// GetExternalCall is a free data retrieval call binding the contract method 0x23e38612.
//
// Solidity: function getExternalCall(uint256 nonce) view returns((uint256,address,address,bytes,uint256,uint256,uint256))
func (_BridgeCore *BridgeCoreCaller) GetExternalCall(opts *bind.CallOpts, nonce *big.Int) (BridgeCoreExternalCallData, error) {
	var out []interface{}
	err := _BridgeCore.contract.Call(opts, &out, "getExternalCall", nonce)

	if err != nil {
		return *new(BridgeCoreExternalCallData), err
	}

	out0 := *abi.ConvertType(out[0], new(BridgeCoreExternalCallData)).(*BridgeCoreExternalCallData)

	return out0, err

}

// GetExternalCall is a free data retrieval call binding the contract method 0x23e38612.
//
// Solidity: function getExternalCall(uint256 nonce) view returns((uint256,address,address,bytes,uint256,uint256,uint256))
func (_BridgeCore *BridgeCoreSession) GetExternalCall(nonce *big.Int) (BridgeCoreExternalCallData, error) {
	return _BridgeCore.Contract.GetExternalCall(&_BridgeCore.CallOpts, nonce)
}

// GetExternalCall is a free data retrieval call binding the contract method 0x23e38612.
//
// Solidity: function getExternalCall(uint256 nonce) view returns((uint256,address,address,bytes,uint256,uint256,uint256))
func (_BridgeCore *BridgeCoreCallerSession) GetExternalCall(nonce *big.Int) (BridgeCoreExternalCallData, error) {
	return _BridgeCore.Contract.GetExternalCall(&_BridgeCore.CallOpts, nonce)
}

// BridgeCall is a paid mutator transaction binding the contract method 0x133b61e0.
//
// Solidity: function bridgeCall(uint256 chainid, address contractAddress, bytes data) payable returns(uint256 non)
func (_BridgeCore *BridgeCoreTransactor) BridgeCall(opts *bind.TransactOpts, chainid *big.Int, contractAddress common.Address, data []byte) (*types.Transaction, error) {
	return _BridgeCore.contract.Transact(opts, "bridgeCall", chainid, contractAddress, data)
}

// BridgeCall is a paid mutator transaction binding the contract method 0x133b61e0.
//
// Solidity: function bridgeCall(uint256 chainid, address contractAddress, bytes data) payable returns(uint256 non)
func (_BridgeCore *BridgeCoreSession) BridgeCall(chainid *big.Int, contractAddress common.Address, data []byte) (*types.Transaction, error) {
	return _BridgeCore.Contract.BridgeCall(&_BridgeCore.TransactOpts, chainid, contractAddress, data)
}

// BridgeCall is a paid mutator transaction binding the contract method 0x133b61e0.
//
// Solidity: function bridgeCall(uint256 chainid, address contractAddress, bytes data) payable returns(uint256 non)
func (_BridgeCore *BridgeCoreTransactorSession) BridgeCall(chainid *big.Int, contractAddress common.Address, data []byte) (*types.Transaction, error) {
	return _BridgeCore.Contract.BridgeCall(&_BridgeCore.TransactOpts, chainid, contractAddress, data)
}

// BridgeCoreExternalCallIterator is returned from FilterExternalCall and is used to iterate over the raw logs and unpacked data for ExternalCall events raised by the BridgeCore contract.
type BridgeCoreExternalCallIterator struct {
	Event *BridgeCoreExternalCall // Event containing the contract specifics and raw log

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
func (it *BridgeCoreExternalCallIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeCoreExternalCall)
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
		it.Event = new(BridgeCoreExternalCall)
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
func (it *BridgeCoreExternalCallIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeCoreExternalCallIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeCoreExternalCall represents a ExternalCall event raised by the BridgeCore contract.
type BridgeCoreExternalCall struct {
	Chainid         *big.Int
	Caller          common.Address
	ContractAddress common.Address
	Data            []byte
	Nonce           *big.Int
	Height          *big.Int
	Value           *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterExternalCall is a free log retrieval operation binding the contract event 0x4010f393a3df9142bff016deb495cde5ff1d6b462d50ba09f942a711a7f5f55c.
//
// Solidity: event ExternalCall(uint256 chainid, address caller, address contractAddress, bytes data, uint256 nonce, uint256 height, uint256 value)
func (_BridgeCore *BridgeCoreFilterer) FilterExternalCall(opts *bind.FilterOpts) (*BridgeCoreExternalCallIterator, error) {

	logs, sub, err := _BridgeCore.contract.FilterLogs(opts, "ExternalCall")
	if err != nil {
		return nil, err
	}
	return &BridgeCoreExternalCallIterator{contract: _BridgeCore.contract, event: "ExternalCall", logs: logs, sub: sub}, nil
}

// WatchExternalCall is a free log subscription operation binding the contract event 0x4010f393a3df9142bff016deb495cde5ff1d6b462d50ba09f942a711a7f5f55c.
//
// Solidity: event ExternalCall(uint256 chainid, address caller, address contractAddress, bytes data, uint256 nonce, uint256 height, uint256 value)
func (_BridgeCore *BridgeCoreFilterer) WatchExternalCall(opts *bind.WatchOpts, sink chan<- *BridgeCoreExternalCall) (event.Subscription, error) {

	logs, sub, err := _BridgeCore.contract.WatchLogs(opts, "ExternalCall")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeCoreExternalCall)
				if err := _BridgeCore.contract.UnpackLog(event, "ExternalCall", log); err != nil {
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

// ParseExternalCall is a log parse operation binding the contract event 0x4010f393a3df9142bff016deb495cde5ff1d6b462d50ba09f942a711a7f5f55c.
//
// Solidity: event ExternalCall(uint256 chainid, address caller, address contractAddress, bytes data, uint256 nonce, uint256 height, uint256 value)
func (_BridgeCore *BridgeCoreFilterer) ParseExternalCall(log types.Log) (*BridgeCoreExternalCall, error) {
	event := new(BridgeCoreExternalCall)
	if err := _BridgeCore.contract.UnpackLog(event, "ExternalCall", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
