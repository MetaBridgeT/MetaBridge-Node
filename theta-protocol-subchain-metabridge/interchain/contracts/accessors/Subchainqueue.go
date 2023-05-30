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

// SubchainqueueMapAddr is an auto generated low-level Go binding around an user-defined struct.
type SubchainqueueMapAddr struct {
	Ownner common.Address
	Target common.Address
	Height *big.Int
	Value  *big.Int
}

// MainMetaData contains all meta data concerning the Main contract.
var MainMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ChainAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"signedTxData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"rpcUrl\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proxyContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"TransactionAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newValidator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"TransactionClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"TransactionCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"TransactionRejectReq\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"TransactionRejected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"TransactionSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"yesVotes\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalVotes\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"TransactionVotes\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"GetMaxProcessedTransactionNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nonces\",\"type\":\"uint256\"}],\"name\":\"GetTransactionEventHeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_rpcUrl\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_proxyContract\",\"type\":\"address\"}],\"name\":\"addChain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newValidator\",\"type\":\"address\"}],\"name\":\"addValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"chains\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"rpcUrl\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"proxyContract\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentValidatorIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"getChain\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_signedTxData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonces\",\"type\":\"uint256\"}],\"name\":\"queueTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"reject\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_txHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"_yesVote\",\"type\":\"bool\"}],\"name\":\"rejectvote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_txHash\",\"type\":\"bytes32\"}],\"name\":\"submit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"transactions\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signedTxData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"completed\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"submitted\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"rejectsubmit\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"rejected\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"yesVotes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"noVotes\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"ownner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structSubchainqueue.MapAddr\",\"name\":\"mapaddr\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"validators\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_yesVote\",\"type\":\"bool\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MainABI is the input ABI used to generate the binding from.
// Deprecated: Use MainMetaData.ABI instead.
var MainABI = MainMetaData.ABI

// Main is an auto generated Go binding around an Ethereum contract.
type Main struct {
	MainCaller     // Read-only binding to the contract
	MainTransactor // Write-only binding to the contract
	MainFilterer   // Log filterer for contract events
}

// MainCaller is an auto generated read-only Go binding around an Ethereum contract.
type MainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MainSession struct {
	Contract     *Main             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MainCallerSession struct {
	Contract *MainCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MainTransactorSession struct {
	Contract     *MainTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MainRaw is an auto generated low-level Go binding around an Ethereum contract.
type MainRaw struct {
	Contract *Main // Generic contract binding to access the raw methods on
}

// MainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MainCallerRaw struct {
	Contract *MainCaller // Generic read-only contract binding to access the raw methods on
}

// MainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MainTransactorRaw struct {
	Contract *MainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMain creates a new instance of Main, bound to a specific deployed contract.
func NewMain(address common.Address, backend bind.ContractBackend) (*Main, error) {
	contract, err := bindMain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Main{MainCaller: MainCaller{contract: contract}, MainTransactor: MainTransactor{contract: contract}, MainFilterer: MainFilterer{contract: contract}}, nil
}

// NewMainCaller creates a new read-only instance of Main, bound to a specific deployed contract.
func NewMainCaller(address common.Address, caller bind.ContractCaller) (*MainCaller, error) {
	contract, err := bindMain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MainCaller{contract: contract}, nil
}

// NewMainTransactor creates a new write-only instance of Main, bound to a specific deployed contract.
func NewMainTransactor(address common.Address, transactor bind.ContractTransactor) (*MainTransactor, error) {
	contract, err := bindMain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MainTransactor{contract: contract}, nil
}

// NewMainFilterer creates a new log filterer instance of Main, bound to a specific deployed contract.
func NewMainFilterer(address common.Address, filterer bind.ContractFilterer) (*MainFilterer, error) {
	contract, err := bindMain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MainFilterer{contract: contract}, nil
}

// bindMain binds a generic wrapper to an already deployed contract.
func bindMain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MainMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main *MainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Main.Contract.MainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main *MainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.Contract.MainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main *MainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main.Contract.MainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main *MainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Main.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main *MainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main *MainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main.Contract.contract.Transact(opts, method, params...)
}

// GetMaxProcessedTransactionNonce is a free data retrieval call binding the contract method 0x9266bfe7.
//
// Solidity: function GetMaxProcessedTransactionNonce() view returns(uint256)
func (_Main *MainCaller) GetMaxProcessedTransactionNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "GetMaxProcessedTransactionNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxProcessedTransactionNonce is a free data retrieval call binding the contract method 0x9266bfe7.
//
// Solidity: function GetMaxProcessedTransactionNonce() view returns(uint256)
func (_Main *MainSession) GetMaxProcessedTransactionNonce() (*big.Int, error) {
	return _Main.Contract.GetMaxProcessedTransactionNonce(&_Main.CallOpts)
}

// GetMaxProcessedTransactionNonce is a free data retrieval call binding the contract method 0x9266bfe7.
//
// Solidity: function GetMaxProcessedTransactionNonce() view returns(uint256)
func (_Main *MainCallerSession) GetMaxProcessedTransactionNonce() (*big.Int, error) {
	return _Main.Contract.GetMaxProcessedTransactionNonce(&_Main.CallOpts)
}

// GetTransactionEventHeight is a free data retrieval call binding the contract method 0x08affa94.
//
// Solidity: function GetTransactionEventHeight(uint256 nonces) view returns(uint256)
func (_Main *MainCaller) GetTransactionEventHeight(opts *bind.CallOpts, nonces *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "GetTransactionEventHeight", nonces)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTransactionEventHeight is a free data retrieval call binding the contract method 0x08affa94.
//
// Solidity: function GetTransactionEventHeight(uint256 nonces) view returns(uint256)
func (_Main *MainSession) GetTransactionEventHeight(nonces *big.Int) (*big.Int, error) {
	return _Main.Contract.GetTransactionEventHeight(&_Main.CallOpts, nonces)
}

// GetTransactionEventHeight is a free data retrieval call binding the contract method 0x08affa94.
//
// Solidity: function GetTransactionEventHeight(uint256 nonces) view returns(uint256)
func (_Main *MainCallerSession) GetTransactionEventHeight(nonces *big.Int) (*big.Int, error) {
	return _Main.Contract.GetTransactionEventHeight(&_Main.CallOpts, nonces)
}

// Chains is a free data retrieval call binding the contract method 0x550325b5.
//
// Solidity: function chains(uint256 ) view returns(uint256 chainId, string rpcUrl, address proxyContract)
func (_Main *MainCaller) Chains(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ChainId       *big.Int
	RpcUrl        string
	ProxyContract common.Address
}, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "chains", arg0)

	outstruct := new(struct {
		ChainId       *big.Int
		RpcUrl        string
		ProxyContract common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ChainId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.RpcUrl = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.ProxyContract = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Chains is a free data retrieval call binding the contract method 0x550325b5.
//
// Solidity: function chains(uint256 ) view returns(uint256 chainId, string rpcUrl, address proxyContract)
func (_Main *MainSession) Chains(arg0 *big.Int) (struct {
	ChainId       *big.Int
	RpcUrl        string
	ProxyContract common.Address
}, error) {
	return _Main.Contract.Chains(&_Main.CallOpts, arg0)
}

// Chains is a free data retrieval call binding the contract method 0x550325b5.
//
// Solidity: function chains(uint256 ) view returns(uint256 chainId, string rpcUrl, address proxyContract)
func (_Main *MainCallerSession) Chains(arg0 *big.Int) (struct {
	ChainId       *big.Int
	RpcUrl        string
	ProxyContract common.Address
}, error) {
	return _Main.Contract.Chains(&_Main.CallOpts, arg0)
}

// CurrentValidatorIndex is a free data retrieval call binding the contract method 0x50e4491f.
//
// Solidity: function currentValidatorIndex() view returns(uint256)
func (_Main *MainCaller) CurrentValidatorIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "currentValidatorIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentValidatorIndex is a free data retrieval call binding the contract method 0x50e4491f.
//
// Solidity: function currentValidatorIndex() view returns(uint256)
func (_Main *MainSession) CurrentValidatorIndex() (*big.Int, error) {
	return _Main.Contract.CurrentValidatorIndex(&_Main.CallOpts)
}

// CurrentValidatorIndex is a free data retrieval call binding the contract method 0x50e4491f.
//
// Solidity: function currentValidatorIndex() view returns(uint256)
func (_Main *MainCallerSession) CurrentValidatorIndex() (*big.Int, error) {
	return _Main.Contract.CurrentValidatorIndex(&_Main.CallOpts)
}

// GetChain is a free data retrieval call binding the contract method 0xb6791ad4.
//
// Solidity: function getChain(uint256 num) view returns(string)
func (_Main *MainCaller) GetChain(opts *bind.CallOpts, num *big.Int) (string, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getChain", num)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetChain is a free data retrieval call binding the contract method 0xb6791ad4.
//
// Solidity: function getChain(uint256 num) view returns(string)
func (_Main *MainSession) GetChain(num *big.Int) (string, error) {
	return _Main.Contract.GetChain(&_Main.CallOpts, num)
}

// GetChain is a free data retrieval call binding the contract method 0xb6791ad4.
//
// Solidity: function getChain(uint256 num) view returns(string)
func (_Main *MainCallerSession) GetChain(num *big.Int) (string, error) {
	return _Main.Contract.GetChain(&_Main.CallOpts, num)
}

// Transactions is a free data retrieval call binding the contract method 0x9ace38c2.
//
// Solidity: function transactions(uint256 ) view returns(bytes32 txHash, uint256 chainId, bytes signedTxData, bool completed, bool submitted, bool rejectsubmit, bool rejected, uint256 yesVotes, uint256 noVotes, address validator, (address,address,uint256,uint256) mapaddr)
func (_Main *MainCaller) Transactions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	TxHash       [32]byte
	ChainId      *big.Int
	SignedTxData []byte
	Completed    bool
	Submitted    bool
	Rejectsubmit bool
	Rejected     bool
	YesVotes     *big.Int
	NoVotes      *big.Int
	Validator    common.Address
	Mapaddr      SubchainqueueMapAddr
}, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "transactions", arg0)

	outstruct := new(struct {
		TxHash       [32]byte
		ChainId      *big.Int
		SignedTxData []byte
		Completed    bool
		Submitted    bool
		Rejectsubmit bool
		Rejected     bool
		YesVotes     *big.Int
		NoVotes      *big.Int
		Validator    common.Address
		Mapaddr      SubchainqueueMapAddr
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TxHash = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.ChainId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.SignedTxData = *abi.ConvertType(out[2], new([]byte)).(*[]byte)
	outstruct.Completed = *abi.ConvertType(out[3], new(bool)).(*bool)
	outstruct.Submitted = *abi.ConvertType(out[4], new(bool)).(*bool)
	outstruct.Rejectsubmit = *abi.ConvertType(out[5], new(bool)).(*bool)
	outstruct.Rejected = *abi.ConvertType(out[6], new(bool)).(*bool)
	outstruct.YesVotes = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.NoVotes = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.Validator = *abi.ConvertType(out[9], new(common.Address)).(*common.Address)
	outstruct.Mapaddr = *abi.ConvertType(out[10], new(SubchainqueueMapAddr)).(*SubchainqueueMapAddr)

	return *outstruct, err

}

// Transactions is a free data retrieval call binding the contract method 0x9ace38c2.
//
// Solidity: function transactions(uint256 ) view returns(bytes32 txHash, uint256 chainId, bytes signedTxData, bool completed, bool submitted, bool rejectsubmit, bool rejected, uint256 yesVotes, uint256 noVotes, address validator, (address,address,uint256,uint256) mapaddr)
func (_Main *MainSession) Transactions(arg0 *big.Int) (struct {
	TxHash       [32]byte
	ChainId      *big.Int
	SignedTxData []byte
	Completed    bool
	Submitted    bool
	Rejectsubmit bool
	Rejected     bool
	YesVotes     *big.Int
	NoVotes      *big.Int
	Validator    common.Address
	Mapaddr      SubchainqueueMapAddr
}, error) {
	return _Main.Contract.Transactions(&_Main.CallOpts, arg0)
}

// Transactions is a free data retrieval call binding the contract method 0x9ace38c2.
//
// Solidity: function transactions(uint256 ) view returns(bytes32 txHash, uint256 chainId, bytes signedTxData, bool completed, bool submitted, bool rejectsubmit, bool rejected, uint256 yesVotes, uint256 noVotes, address validator, (address,address,uint256,uint256) mapaddr)
func (_Main *MainCallerSession) Transactions(arg0 *big.Int) (struct {
	TxHash       [32]byte
	ChainId      *big.Int
	SignedTxData []byte
	Completed    bool
	Submitted    bool
	Rejectsubmit bool
	Rejected     bool
	YesVotes     *big.Int
	NoVotes      *big.Int
	Validator    common.Address
	Mapaddr      SubchainqueueMapAddr
}, error) {
	return _Main.Contract.Transactions(&_Main.CallOpts, arg0)
}

// Validators is a free data retrieval call binding the contract method 0xfa52c7d8.
//
// Solidity: function validators(address ) view returns(bool)
func (_Main *MainCaller) Validators(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "validators", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Validators is a free data retrieval call binding the contract method 0xfa52c7d8.
//
// Solidity: function validators(address ) view returns(bool)
func (_Main *MainSession) Validators(arg0 common.Address) (bool, error) {
	return _Main.Contract.Validators(&_Main.CallOpts, arg0)
}

// Validators is a free data retrieval call binding the contract method 0xfa52c7d8.
//
// Solidity: function validators(address ) view returns(bool)
func (_Main *MainCallerSession) Validators(arg0 common.Address) (bool, error) {
	return _Main.Contract.Validators(&_Main.CallOpts, arg0)
}

// AddChain is a paid mutator transaction binding the contract method 0xa6788809.
//
// Solidity: function addChain(uint256 _chainId, string _rpcUrl, address _proxyContract) returns()
func (_Main *MainTransactor) AddChain(opts *bind.TransactOpts, _chainId *big.Int, _rpcUrl string, _proxyContract common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "addChain", _chainId, _rpcUrl, _proxyContract)
}

// AddChain is a paid mutator transaction binding the contract method 0xa6788809.
//
// Solidity: function addChain(uint256 _chainId, string _rpcUrl, address _proxyContract) returns()
func (_Main *MainSession) AddChain(_chainId *big.Int, _rpcUrl string, _proxyContract common.Address) (*types.Transaction, error) {
	return _Main.Contract.AddChain(&_Main.TransactOpts, _chainId, _rpcUrl, _proxyContract)
}

// AddChain is a paid mutator transaction binding the contract method 0xa6788809.
//
// Solidity: function addChain(uint256 _chainId, string _rpcUrl, address _proxyContract) returns()
func (_Main *MainTransactorSession) AddChain(_chainId *big.Int, _rpcUrl string, _proxyContract common.Address) (*types.Transaction, error) {
	return _Main.Contract.AddChain(&_Main.TransactOpts, _chainId, _rpcUrl, _proxyContract)
}

// AddValidator is a paid mutator transaction binding the contract method 0x4d238c8e.
//
// Solidity: function addValidator(address newValidator) returns()
func (_Main *MainTransactor) AddValidator(opts *bind.TransactOpts, newValidator common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "addValidator", newValidator)
}

// AddValidator is a paid mutator transaction binding the contract method 0x4d238c8e.
//
// Solidity: function addValidator(address newValidator) returns()
func (_Main *MainSession) AddValidator(newValidator common.Address) (*types.Transaction, error) {
	return _Main.Contract.AddValidator(&_Main.TransactOpts, newValidator)
}

// AddValidator is a paid mutator transaction binding the contract method 0x4d238c8e.
//
// Solidity: function addValidator(address newValidator) returns()
func (_Main *MainTransactorSession) AddValidator(newValidator common.Address) (*types.Transaction, error) {
	return _Main.Contract.AddValidator(&_Main.TransactOpts, newValidator)
}

// Claim is a paid mutator transaction binding the contract method 0x379607f5.
//
// Solidity: function claim(uint256 index) returns()
func (_Main *MainTransactor) Claim(opts *bind.TransactOpts, index *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "claim", index)
}

// Claim is a paid mutator transaction binding the contract method 0x379607f5.
//
// Solidity: function claim(uint256 index) returns()
func (_Main *MainSession) Claim(index *big.Int) (*types.Transaction, error) {
	return _Main.Contract.Claim(&_Main.TransactOpts, index)
}

// Claim is a paid mutator transaction binding the contract method 0x379607f5.
//
// Solidity: function claim(uint256 index) returns()
func (_Main *MainTransactorSession) Claim(index *big.Int) (*types.Transaction, error) {
	return _Main.Contract.Claim(&_Main.TransactOpts, index)
}

// QueueTransaction is a paid mutator transaction binding the contract method 0xf2b772aa.
//
// Solidity: function queueTransaction(uint256 _chainId, bytes _signedTxData, address target, uint256 value, uint256 nonces) returns()
func (_Main *MainTransactor) QueueTransaction(opts *bind.TransactOpts, _chainId *big.Int, _signedTxData []byte, target common.Address, value *big.Int, nonces *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "queueTransaction", _chainId, _signedTxData, target, value, nonces)
}

// QueueTransaction is a paid mutator transaction binding the contract method 0xf2b772aa.
//
// Solidity: function queueTransaction(uint256 _chainId, bytes _signedTxData, address target, uint256 value, uint256 nonces) returns()
func (_Main *MainSession) QueueTransaction(_chainId *big.Int, _signedTxData []byte, target common.Address, value *big.Int, nonces *big.Int) (*types.Transaction, error) {
	return _Main.Contract.QueueTransaction(&_Main.TransactOpts, _chainId, _signedTxData, target, value, nonces)
}

// QueueTransaction is a paid mutator transaction binding the contract method 0xf2b772aa.
//
// Solidity: function queueTransaction(uint256 _chainId, bytes _signedTxData, address target, uint256 value, uint256 nonces) returns()
func (_Main *MainTransactorSession) QueueTransaction(_chainId *big.Int, _signedTxData []byte, target common.Address, value *big.Int, nonces *big.Int) (*types.Transaction, error) {
	return _Main.Contract.QueueTransaction(&_Main.TransactOpts, _chainId, _signedTxData, target, value, nonces)
}

// Reject is a paid mutator transaction binding the contract method 0xb8adaa11.
//
// Solidity: function reject(uint256 _index) returns()
func (_Main *MainTransactor) Reject(opts *bind.TransactOpts, _index *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "reject", _index)
}

// Reject is a paid mutator transaction binding the contract method 0xb8adaa11.
//
// Solidity: function reject(uint256 _index) returns()
func (_Main *MainSession) Reject(_index *big.Int) (*types.Transaction, error) {
	return _Main.Contract.Reject(&_Main.TransactOpts, _index)
}

// Reject is a paid mutator transaction binding the contract method 0xb8adaa11.
//
// Solidity: function reject(uint256 _index) returns()
func (_Main *MainTransactorSession) Reject(_index *big.Int) (*types.Transaction, error) {
	return _Main.Contract.Reject(&_Main.TransactOpts, _index)
}

// Rejectvote is a paid mutator transaction binding the contract method 0xcd597ac4.
//
// Solidity: function rejectvote(uint256 _index, bytes32 _txHash, bool _yesVote) returns()
func (_Main *MainTransactor) Rejectvote(opts *bind.TransactOpts, _index *big.Int, _txHash [32]byte, _yesVote bool) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "rejectvote", _index, _txHash, _yesVote)
}

// Rejectvote is a paid mutator transaction binding the contract method 0xcd597ac4.
//
// Solidity: function rejectvote(uint256 _index, bytes32 _txHash, bool _yesVote) returns()
func (_Main *MainSession) Rejectvote(_index *big.Int, _txHash [32]byte, _yesVote bool) (*types.Transaction, error) {
	return _Main.Contract.Rejectvote(&_Main.TransactOpts, _index, _txHash, _yesVote)
}

// Rejectvote is a paid mutator transaction binding the contract method 0xcd597ac4.
//
// Solidity: function rejectvote(uint256 _index, bytes32 _txHash, bool _yesVote) returns()
func (_Main *MainTransactorSession) Rejectvote(_index *big.Int, _txHash [32]byte, _yesVote bool) (*types.Transaction, error) {
	return _Main.Contract.Rejectvote(&_Main.TransactOpts, _index, _txHash, _yesVote)
}

// Submit is a paid mutator transaction binding the contract method 0x2ecea788.
//
// Solidity: function submit(uint256 _index, bytes32 _txHash) returns()
func (_Main *MainTransactor) Submit(opts *bind.TransactOpts, _index *big.Int, _txHash [32]byte) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "submit", _index, _txHash)
}

// Submit is a paid mutator transaction binding the contract method 0x2ecea788.
//
// Solidity: function submit(uint256 _index, bytes32 _txHash) returns()
func (_Main *MainSession) Submit(_index *big.Int, _txHash [32]byte) (*types.Transaction, error) {
	return _Main.Contract.Submit(&_Main.TransactOpts, _index, _txHash)
}

// Submit is a paid mutator transaction binding the contract method 0x2ecea788.
//
// Solidity: function submit(uint256 _index, bytes32 _txHash) returns()
func (_Main *MainTransactorSession) Submit(_index *big.Int, _txHash [32]byte) (*types.Transaction, error) {
	return _Main.Contract.Submit(&_Main.TransactOpts, _index, _txHash)
}

// Vote is a paid mutator transaction binding the contract method 0xc9d27afe.
//
// Solidity: function vote(uint256 _index, bool _yesVote) returns()
func (_Main *MainTransactor) Vote(opts *bind.TransactOpts, _index *big.Int, _yesVote bool) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "vote", _index, _yesVote)
}

// Vote is a paid mutator transaction binding the contract method 0xc9d27afe.
//
// Solidity: function vote(uint256 _index, bool _yesVote) returns()
func (_Main *MainSession) Vote(_index *big.Int, _yesVote bool) (*types.Transaction, error) {
	return _Main.Contract.Vote(&_Main.TransactOpts, _index, _yesVote)
}

// Vote is a paid mutator transaction binding the contract method 0xc9d27afe.
//
// Solidity: function vote(uint256 _index, bool _yesVote) returns()
func (_Main *MainTransactorSession) Vote(_index *big.Int, _yesVote bool) (*types.Transaction, error) {
	return _Main.Contract.Vote(&_Main.TransactOpts, _index, _yesVote)
}

// MainChainAddedIterator is returned from FilterChainAdded and is used to iterate over the raw logs and unpacked data for ChainAdded events raised by the Main contract.
type MainChainAddedIterator struct {
	Event *MainChainAdded // Event containing the contract specifics and raw log

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
func (it *MainChainAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainChainAdded)
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
		it.Event = new(MainChainAdded)
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
func (it *MainChainAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainChainAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainChainAdded represents a ChainAdded event raised by the Main contract.
type MainChainAdded struct {
	Arg0   *big.Int
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterChainAdded is a free log retrieval operation binding the contract event 0xd34413f9daf08c09e55b3ad5e2f7f18249f3ff26c061cb904f99597dd14ad53e.
//
// Solidity: event ChainAdded(uint256 arg0, address sender)
func (_Main *MainFilterer) FilterChainAdded(opts *bind.FilterOpts) (*MainChainAddedIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "ChainAdded")
	if err != nil {
		return nil, err
	}
	return &MainChainAddedIterator{contract: _Main.contract, event: "ChainAdded", logs: logs, sub: sub}, nil
}

// WatchChainAdded is a free log subscription operation binding the contract event 0xd34413f9daf08c09e55b3ad5e2f7f18249f3ff26c061cb904f99597dd14ad53e.
//
// Solidity: event ChainAdded(uint256 arg0, address sender)
func (_Main *MainFilterer) WatchChainAdded(opts *bind.WatchOpts, sink chan<- *MainChainAdded) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "ChainAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainChainAdded)
				if err := _Main.contract.UnpackLog(event, "ChainAdded", log); err != nil {
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

// ParseChainAdded is a log parse operation binding the contract event 0xd34413f9daf08c09e55b3ad5e2f7f18249f3ff26c061cb904f99597dd14ad53e.
//
// Solidity: event ChainAdded(uint256 arg0, address sender)
func (_Main *MainFilterer) ParseChainAdded(log types.Log) (*MainChainAdded, error) {
	event := new(MainChainAdded)
	if err := _Main.contract.UnpackLog(event, "ChainAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainTransactionAddedIterator is returned from FilterTransactionAdded and is used to iterate over the raw logs and unpacked data for TransactionAdded events raised by the Main contract.
type MainTransactionAddedIterator struct {
	Event *MainTransactionAdded // Event containing the contract specifics and raw log

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
func (it *MainTransactionAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainTransactionAdded)
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
		it.Event = new(MainTransactionAdded)
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
func (it *MainTransactionAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainTransactionAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainTransactionAdded represents a TransactionAdded event raised by the Main contract.
type MainTransactionAdded struct {
	Index         *big.Int
	TxHash        [32]byte
	ChainId       *big.Int
	SignedTxData  []byte
	Height        *big.Int
	Validator     common.Address
	RpcUrl        string
	ProxyContract common.Address
	Value         *big.Int
	Sender        common.Address
	Target        common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTransactionAdded is a free log retrieval operation binding the contract event 0x851a265a65a06334b90f7394629088b8bfccab83e098587ce68e5fa68b27543d.
//
// Solidity: event TransactionAdded(uint256 index, bytes32 txHash, uint256 chainId, bytes signedTxData, uint256 height, address validator, string rpcUrl, address proxyContract, uint256 value, address sender, address target)
func (_Main *MainFilterer) FilterTransactionAdded(opts *bind.FilterOpts) (*MainTransactionAddedIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "TransactionAdded")
	if err != nil {
		return nil, err
	}
	return &MainTransactionAddedIterator{contract: _Main.contract, event: "TransactionAdded", logs: logs, sub: sub}, nil
}

// WatchTransactionAdded is a free log subscription operation binding the contract event 0x851a265a65a06334b90f7394629088b8bfccab83e098587ce68e5fa68b27543d.
//
// Solidity: event TransactionAdded(uint256 index, bytes32 txHash, uint256 chainId, bytes signedTxData, uint256 height, address validator, string rpcUrl, address proxyContract, uint256 value, address sender, address target)
func (_Main *MainFilterer) WatchTransactionAdded(opts *bind.WatchOpts, sink chan<- *MainTransactionAdded) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "TransactionAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainTransactionAdded)
				if err := _Main.contract.UnpackLog(event, "TransactionAdded", log); err != nil {
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

// ParseTransactionAdded is a log parse operation binding the contract event 0x851a265a65a06334b90f7394629088b8bfccab83e098587ce68e5fa68b27543d.
//
// Solidity: event TransactionAdded(uint256 index, bytes32 txHash, uint256 chainId, bytes signedTxData, uint256 height, address validator, string rpcUrl, address proxyContract, uint256 value, address sender, address target)
func (_Main *MainFilterer) ParseTransactionAdded(log types.Log) (*MainTransactionAdded, error) {
	event := new(MainTransactionAdded)
	if err := _Main.contract.UnpackLog(event, "TransactionAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainTransactionClaimedIterator is returned from FilterTransactionClaimed and is used to iterate over the raw logs and unpacked data for TransactionClaimed events raised by the Main contract.
type MainTransactionClaimedIterator struct {
	Event *MainTransactionClaimed // Event containing the contract specifics and raw log

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
func (it *MainTransactionClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainTransactionClaimed)
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
		it.Event = new(MainTransactionClaimed)
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
func (it *MainTransactionClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainTransactionClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainTransactionClaimed represents a TransactionClaimed event raised by the Main contract.
type MainTransactionClaimed struct {
	Index        *big.Int
	NewValidator common.Address
	Sender       common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTransactionClaimed is a free log retrieval operation binding the contract event 0x30a807b8ab28eebb93ad225ce76f327ae85af49589c51adcdd95da0aba88674a.
//
// Solidity: event TransactionClaimed(uint256 index, address newValidator, address sender)
func (_Main *MainFilterer) FilterTransactionClaimed(opts *bind.FilterOpts) (*MainTransactionClaimedIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "TransactionClaimed")
	if err != nil {
		return nil, err
	}
	return &MainTransactionClaimedIterator{contract: _Main.contract, event: "TransactionClaimed", logs: logs, sub: sub}, nil
}

// WatchTransactionClaimed is a free log subscription operation binding the contract event 0x30a807b8ab28eebb93ad225ce76f327ae85af49589c51adcdd95da0aba88674a.
//
// Solidity: event TransactionClaimed(uint256 index, address newValidator, address sender)
func (_Main *MainFilterer) WatchTransactionClaimed(opts *bind.WatchOpts, sink chan<- *MainTransactionClaimed) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "TransactionClaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainTransactionClaimed)
				if err := _Main.contract.UnpackLog(event, "TransactionClaimed", log); err != nil {
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

// ParseTransactionClaimed is a log parse operation binding the contract event 0x30a807b8ab28eebb93ad225ce76f327ae85af49589c51adcdd95da0aba88674a.
//
// Solidity: event TransactionClaimed(uint256 index, address newValidator, address sender)
func (_Main *MainFilterer) ParseTransactionClaimed(log types.Log) (*MainTransactionClaimed, error) {
	event := new(MainTransactionClaimed)
	if err := _Main.contract.UnpackLog(event, "TransactionClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainTransactionCompletedIterator is returned from FilterTransactionCompleted and is used to iterate over the raw logs and unpacked data for TransactionCompleted events raised by the Main contract.
type MainTransactionCompletedIterator struct {
	Event *MainTransactionCompleted // Event containing the contract specifics and raw log

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
func (it *MainTransactionCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainTransactionCompleted)
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
		it.Event = new(MainTransactionCompleted)
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
func (it *MainTransactionCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainTransactionCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainTransactionCompleted represents a TransactionCompleted event raised by the Main contract.
type MainTransactionCompleted struct {
	Index   *big.Int
	TxHash  [32]byte
	Success bool
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransactionCompleted is a free log retrieval operation binding the contract event 0xf23d2b2aa31db7dbfa1f15516d23df0b11101336c1b0e33aba0d0c6bf863c4e6.
//
// Solidity: event TransactionCompleted(uint256 index, bytes32 txHash, bool success, address sender)
func (_Main *MainFilterer) FilterTransactionCompleted(opts *bind.FilterOpts) (*MainTransactionCompletedIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "TransactionCompleted")
	if err != nil {
		return nil, err
	}
	return &MainTransactionCompletedIterator{contract: _Main.contract, event: "TransactionCompleted", logs: logs, sub: sub}, nil
}

// WatchTransactionCompleted is a free log subscription operation binding the contract event 0xf23d2b2aa31db7dbfa1f15516d23df0b11101336c1b0e33aba0d0c6bf863c4e6.
//
// Solidity: event TransactionCompleted(uint256 index, bytes32 txHash, bool success, address sender)
func (_Main *MainFilterer) WatchTransactionCompleted(opts *bind.WatchOpts, sink chan<- *MainTransactionCompleted) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "TransactionCompleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainTransactionCompleted)
				if err := _Main.contract.UnpackLog(event, "TransactionCompleted", log); err != nil {
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

// ParseTransactionCompleted is a log parse operation binding the contract event 0xf23d2b2aa31db7dbfa1f15516d23df0b11101336c1b0e33aba0d0c6bf863c4e6.
//
// Solidity: event TransactionCompleted(uint256 index, bytes32 txHash, bool success, address sender)
func (_Main *MainFilterer) ParseTransactionCompleted(log types.Log) (*MainTransactionCompleted, error) {
	event := new(MainTransactionCompleted)
	if err := _Main.contract.UnpackLog(event, "TransactionCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainTransactionRejectReqIterator is returned from FilterTransactionRejectReq and is used to iterate over the raw logs and unpacked data for TransactionRejectReq events raised by the Main contract.
type MainTransactionRejectReqIterator struct {
	Event *MainTransactionRejectReq // Event containing the contract specifics and raw log

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
func (it *MainTransactionRejectReqIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainTransactionRejectReq)
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
		it.Event = new(MainTransactionRejectReq)
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
func (it *MainTransactionRejectReqIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainTransactionRejectReqIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainTransactionRejectReq represents a TransactionRejectReq event raised by the Main contract.
type MainTransactionRejectReq struct {
	Index  *big.Int
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransactionRejectReq is a free log retrieval operation binding the contract event 0x3cfcffbf945df6c5613d5a6849d6ace06d9c5680d9080ca3cfab7e49c3a723bd.
//
// Solidity: event TransactionRejectReq(uint256 index, address sender)
func (_Main *MainFilterer) FilterTransactionRejectReq(opts *bind.FilterOpts) (*MainTransactionRejectReqIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "TransactionRejectReq")
	if err != nil {
		return nil, err
	}
	return &MainTransactionRejectReqIterator{contract: _Main.contract, event: "TransactionRejectReq", logs: logs, sub: sub}, nil
}

// WatchTransactionRejectReq is a free log subscription operation binding the contract event 0x3cfcffbf945df6c5613d5a6849d6ace06d9c5680d9080ca3cfab7e49c3a723bd.
//
// Solidity: event TransactionRejectReq(uint256 index, address sender)
func (_Main *MainFilterer) WatchTransactionRejectReq(opts *bind.WatchOpts, sink chan<- *MainTransactionRejectReq) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "TransactionRejectReq")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainTransactionRejectReq)
				if err := _Main.contract.UnpackLog(event, "TransactionRejectReq", log); err != nil {
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

// ParseTransactionRejectReq is a log parse operation binding the contract event 0x3cfcffbf945df6c5613d5a6849d6ace06d9c5680d9080ca3cfab7e49c3a723bd.
//
// Solidity: event TransactionRejectReq(uint256 index, address sender)
func (_Main *MainFilterer) ParseTransactionRejectReq(log types.Log) (*MainTransactionRejectReq, error) {
	event := new(MainTransactionRejectReq)
	if err := _Main.contract.UnpackLog(event, "TransactionRejectReq", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainTransactionRejectedIterator is returned from FilterTransactionRejected and is used to iterate over the raw logs and unpacked data for TransactionRejected events raised by the Main contract.
type MainTransactionRejectedIterator struct {
	Event *MainTransactionRejected // Event containing the contract specifics and raw log

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
func (it *MainTransactionRejectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainTransactionRejected)
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
		it.Event = new(MainTransactionRejected)
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
func (it *MainTransactionRejectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainTransactionRejectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainTransactionRejected represents a TransactionRejected event raised by the Main contract.
type MainTransactionRejected struct {
	Index  *big.Int
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransactionRejected is a free log retrieval operation binding the contract event 0x7777ac68274c98eb80c448bc97e699ea01d00cd3a5a304bde947b35c82e1dbd1.
//
// Solidity: event TransactionRejected(uint256 index, address sender)
func (_Main *MainFilterer) FilterTransactionRejected(opts *bind.FilterOpts) (*MainTransactionRejectedIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "TransactionRejected")
	if err != nil {
		return nil, err
	}
	return &MainTransactionRejectedIterator{contract: _Main.contract, event: "TransactionRejected", logs: logs, sub: sub}, nil
}

// WatchTransactionRejected is a free log subscription operation binding the contract event 0x7777ac68274c98eb80c448bc97e699ea01d00cd3a5a304bde947b35c82e1dbd1.
//
// Solidity: event TransactionRejected(uint256 index, address sender)
func (_Main *MainFilterer) WatchTransactionRejected(opts *bind.WatchOpts, sink chan<- *MainTransactionRejected) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "TransactionRejected")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainTransactionRejected)
				if err := _Main.contract.UnpackLog(event, "TransactionRejected", log); err != nil {
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

// ParseTransactionRejected is a log parse operation binding the contract event 0x7777ac68274c98eb80c448bc97e699ea01d00cd3a5a304bde947b35c82e1dbd1.
//
// Solidity: event TransactionRejected(uint256 index, address sender)
func (_Main *MainFilterer) ParseTransactionRejected(log types.Log) (*MainTransactionRejected, error) {
	event := new(MainTransactionRejected)
	if err := _Main.contract.UnpackLog(event, "TransactionRejected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainTransactionSubmittedIterator is returned from FilterTransactionSubmitted and is used to iterate over the raw logs and unpacked data for TransactionSubmitted events raised by the Main contract.
type MainTransactionSubmittedIterator struct {
	Event *MainTransactionSubmitted // Event containing the contract specifics and raw log

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
func (it *MainTransactionSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainTransactionSubmitted)
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
		it.Event = new(MainTransactionSubmitted)
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
func (it *MainTransactionSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainTransactionSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainTransactionSubmitted represents a TransactionSubmitted event raised by the Main contract.
type MainTransactionSubmitted struct {
	Index   *big.Int
	TxHash  [32]byte
	Success bool
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransactionSubmitted is a free log retrieval operation binding the contract event 0x6db377314210eb14c94abafe230add59d50a3d1d40d8dfb86044d81b7f8350fb.
//
// Solidity: event TransactionSubmitted(uint256 index, bytes32 txHash, bool success, address sender)
func (_Main *MainFilterer) FilterTransactionSubmitted(opts *bind.FilterOpts) (*MainTransactionSubmittedIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "TransactionSubmitted")
	if err != nil {
		return nil, err
	}
	return &MainTransactionSubmittedIterator{contract: _Main.contract, event: "TransactionSubmitted", logs: logs, sub: sub}, nil
}

// WatchTransactionSubmitted is a free log subscription operation binding the contract event 0x6db377314210eb14c94abafe230add59d50a3d1d40d8dfb86044d81b7f8350fb.
//
// Solidity: event TransactionSubmitted(uint256 index, bytes32 txHash, bool success, address sender)
func (_Main *MainFilterer) WatchTransactionSubmitted(opts *bind.WatchOpts, sink chan<- *MainTransactionSubmitted) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "TransactionSubmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainTransactionSubmitted)
				if err := _Main.contract.UnpackLog(event, "TransactionSubmitted", log); err != nil {
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

// ParseTransactionSubmitted is a log parse operation binding the contract event 0x6db377314210eb14c94abafe230add59d50a3d1d40d8dfb86044d81b7f8350fb.
//
// Solidity: event TransactionSubmitted(uint256 index, bytes32 txHash, bool success, address sender)
func (_Main *MainFilterer) ParseTransactionSubmitted(log types.Log) (*MainTransactionSubmitted, error) {
	event := new(MainTransactionSubmitted)
	if err := _Main.contract.UnpackLog(event, "TransactionSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainTransactionVotesIterator is returned from FilterTransactionVotes and is used to iterate over the raw logs and unpacked data for TransactionVotes events raised by the Main contract.
type MainTransactionVotesIterator struct {
	Event *MainTransactionVotes // Event containing the contract specifics and raw log

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
func (it *MainTransactionVotesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainTransactionVotes)
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
		it.Event = new(MainTransactionVotes)
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
func (it *MainTransactionVotesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainTransactionVotesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainTransactionVotes represents a TransactionVotes event raised by the Main contract.
type MainTransactionVotes struct {
	Index      *big.Int
	YesVotes   *big.Int
	TotalVotes *big.Int
	Sender     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTransactionVotes is a free log retrieval operation binding the contract event 0x1e969d1826189dc4cc1012da76c2eefaeae2c7d5e6ff82004745c2b8a61c8619.
//
// Solidity: event TransactionVotes(uint256 index, uint256 yesVotes, uint256 totalVotes, address sender)
func (_Main *MainFilterer) FilterTransactionVotes(opts *bind.FilterOpts) (*MainTransactionVotesIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "TransactionVotes")
	if err != nil {
		return nil, err
	}
	return &MainTransactionVotesIterator{contract: _Main.contract, event: "TransactionVotes", logs: logs, sub: sub}, nil
}

// WatchTransactionVotes is a free log subscription operation binding the contract event 0x1e969d1826189dc4cc1012da76c2eefaeae2c7d5e6ff82004745c2b8a61c8619.
//
// Solidity: event TransactionVotes(uint256 index, uint256 yesVotes, uint256 totalVotes, address sender)
func (_Main *MainFilterer) WatchTransactionVotes(opts *bind.WatchOpts, sink chan<- *MainTransactionVotes) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "TransactionVotes")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainTransactionVotes)
				if err := _Main.contract.UnpackLog(event, "TransactionVotes", log); err != nil {
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

// ParseTransactionVotes is a log parse operation binding the contract event 0x1e969d1826189dc4cc1012da76c2eefaeae2c7d5e6ff82004745c2b8a61c8619.
//
// Solidity: event TransactionVotes(uint256 index, uint256 yesVotes, uint256 totalVotes, address sender)
func (_Main *MainFilterer) ParseTransactionVotes(log types.Log) (*MainTransactionVotes, error) {
	event := new(MainTransactionVotes)
	if err := _Main.contract.UnpackLog(event, "TransactionVotes", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
