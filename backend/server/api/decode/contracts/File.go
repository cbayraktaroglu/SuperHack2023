// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package file

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

// FileFileParts is an auto generated low-level Go binding around an user-defined struct.
type FileFileParts struct {
	ChainID         string
	TransactionHash string
}

// BlockchainFileMetaData contains all meta data concerning the BlockchainFile contract.
var BlockchainFileMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_fileType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_fileName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_fileCheckSum\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"_transactionList\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"_chainIdList\",\"type\":\"string[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fileAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"result\",\"type\":\"bool\"}],\"name\":\"VerificationResult\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"contractIWorldID\",\"name\":\"_worldId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_appId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_actionId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"signal\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nullifierHash\",\"type\":\"uint256\"},{\"internalType\":\"uint256[8]\",\"name\":\"proof\",\"type\":\"uint256[8]\"}],\"name\":\"checkVerification\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factoryAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFileCheckSum\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFileName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFileType\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNumberofTransactions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getTransactionAtIndex\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"chainID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"transactionHash\",\"type\":\"string\"}],\"internalType\":\"structFile.FileParts\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVerification\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BlockchainFileABI is the input ABI used to generate the binding from.
// Deprecated: Use BlockchainFileMetaData.ABI instead.
var BlockchainFileABI = BlockchainFileMetaData.ABI

// BlockchainFile is an auto generated Go binding around an Ethereum contract.
type BlockchainFile struct {
	BlockchainFileCaller     // Read-only binding to the contract
	BlockchainFileTransactor // Write-only binding to the contract
	BlockchainFileFilterer   // Log filterer for contract events
}

// BlockchainFileCaller is an auto generated read-only Go binding around an Ethereum contract.
type BlockchainFileCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockchainFileTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BlockchainFileTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockchainFileFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BlockchainFileFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockchainFileSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BlockchainFileSession struct {
	Contract     *BlockchainFile   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BlockchainFileCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BlockchainFileCallerSession struct {
	Contract *BlockchainFileCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// BlockchainFileTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BlockchainFileTransactorSession struct {
	Contract     *BlockchainFileTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// BlockchainFileRaw is an auto generated low-level Go binding around an Ethereum contract.
type BlockchainFileRaw struct {
	Contract *BlockchainFile // Generic contract binding to access the raw methods on
}

// BlockchainFileCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BlockchainFileCallerRaw struct {
	Contract *BlockchainFileCaller // Generic read-only contract binding to access the raw methods on
}

// BlockchainFileTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BlockchainFileTransactorRaw struct {
	Contract *BlockchainFileTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBlockchainFile creates a new instance of BlockchainFile, bound to a specific deployed contract.
func NewBlockchainFile(address common.Address, backend bind.ContractBackend) (*BlockchainFile, error) {
	contract, err := bindBlockchainFile(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BlockchainFile{BlockchainFileCaller: BlockchainFileCaller{contract: contract}, BlockchainFileTransactor: BlockchainFileTransactor{contract: contract}, BlockchainFileFilterer: BlockchainFileFilterer{contract: contract}}, nil
}

// NewBlockchainFileCaller creates a new read-only instance of BlockchainFile, bound to a specific deployed contract.
func NewBlockchainFileCaller(address common.Address, caller bind.ContractCaller) (*BlockchainFileCaller, error) {
	contract, err := bindBlockchainFile(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BlockchainFileCaller{contract: contract}, nil
}

// NewBlockchainFileTransactor creates a new write-only instance of BlockchainFile, bound to a specific deployed contract.
func NewBlockchainFileTransactor(address common.Address, transactor bind.ContractTransactor) (*BlockchainFileTransactor, error) {
	contract, err := bindBlockchainFile(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BlockchainFileTransactor{contract: contract}, nil
}

// NewBlockchainFileFilterer creates a new log filterer instance of BlockchainFile, bound to a specific deployed contract.
func NewBlockchainFileFilterer(address common.Address, filterer bind.ContractFilterer) (*BlockchainFileFilterer, error) {
	contract, err := bindBlockchainFile(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BlockchainFileFilterer{contract: contract}, nil
}

// bindBlockchainFile binds a generic wrapper to an already deployed contract.
func bindBlockchainFile(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BlockchainFileMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BlockchainFile *BlockchainFileRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BlockchainFile.Contract.BlockchainFileCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BlockchainFile *BlockchainFileRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlockchainFile.Contract.BlockchainFileTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BlockchainFile *BlockchainFileRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BlockchainFile.Contract.BlockchainFileTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BlockchainFile *BlockchainFileCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BlockchainFile.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BlockchainFile *BlockchainFileTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlockchainFile.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BlockchainFile *BlockchainFileTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BlockchainFile.Contract.contract.Transact(opts, method, params...)
}

// FactoryAddress is a free data retrieval call binding the contract method 0x966dae0e.
//
// Solidity: function factoryAddress() view returns(address)
func (_BlockchainFile *BlockchainFileCaller) FactoryAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BlockchainFile.contract.Call(opts, &out, "factoryAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FactoryAddress is a free data retrieval call binding the contract method 0x966dae0e.
//
// Solidity: function factoryAddress() view returns(address)
func (_BlockchainFile *BlockchainFileSession) FactoryAddress() (common.Address, error) {
	return _BlockchainFile.Contract.FactoryAddress(&_BlockchainFile.CallOpts)
}

// FactoryAddress is a free data retrieval call binding the contract method 0x966dae0e.
//
// Solidity: function factoryAddress() view returns(address)
func (_BlockchainFile *BlockchainFileCallerSession) FactoryAddress() (common.Address, error) {
	return _BlockchainFile.Contract.FactoryAddress(&_BlockchainFile.CallOpts)
}

// GetFileCheckSum is a free data retrieval call binding the contract method 0x6a1b88a6.
//
// Solidity: function getFileCheckSum() view returns(string)
func (_BlockchainFile *BlockchainFileCaller) GetFileCheckSum(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BlockchainFile.contract.Call(opts, &out, "getFileCheckSum")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetFileCheckSum is a free data retrieval call binding the contract method 0x6a1b88a6.
//
// Solidity: function getFileCheckSum() view returns(string)
func (_BlockchainFile *BlockchainFileSession) GetFileCheckSum() (string, error) {
	return _BlockchainFile.Contract.GetFileCheckSum(&_BlockchainFile.CallOpts)
}

// GetFileCheckSum is a free data retrieval call binding the contract method 0x6a1b88a6.
//
// Solidity: function getFileCheckSum() view returns(string)
func (_BlockchainFile *BlockchainFileCallerSession) GetFileCheckSum() (string, error) {
	return _BlockchainFile.Contract.GetFileCheckSum(&_BlockchainFile.CallOpts)
}

// GetFileName is a free data retrieval call binding the contract method 0xfcf6c23c.
//
// Solidity: function getFileName() view returns(string)
func (_BlockchainFile *BlockchainFileCaller) GetFileName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BlockchainFile.contract.Call(opts, &out, "getFileName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetFileName is a free data retrieval call binding the contract method 0xfcf6c23c.
//
// Solidity: function getFileName() view returns(string)
func (_BlockchainFile *BlockchainFileSession) GetFileName() (string, error) {
	return _BlockchainFile.Contract.GetFileName(&_BlockchainFile.CallOpts)
}

// GetFileName is a free data retrieval call binding the contract method 0xfcf6c23c.
//
// Solidity: function getFileName() view returns(string)
func (_BlockchainFile *BlockchainFileCallerSession) GetFileName() (string, error) {
	return _BlockchainFile.Contract.GetFileName(&_BlockchainFile.CallOpts)
}

// GetFileType is a free data retrieval call binding the contract method 0x16b9e7ef.
//
// Solidity: function getFileType() view returns(string)
func (_BlockchainFile *BlockchainFileCaller) GetFileType(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BlockchainFile.contract.Call(opts, &out, "getFileType")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetFileType is a free data retrieval call binding the contract method 0x16b9e7ef.
//
// Solidity: function getFileType() view returns(string)
func (_BlockchainFile *BlockchainFileSession) GetFileType() (string, error) {
	return _BlockchainFile.Contract.GetFileType(&_BlockchainFile.CallOpts)
}

// GetFileType is a free data retrieval call binding the contract method 0x16b9e7ef.
//
// Solidity: function getFileType() view returns(string)
func (_BlockchainFile *BlockchainFileCallerSession) GetFileType() (string, error) {
	return _BlockchainFile.Contract.GetFileType(&_BlockchainFile.CallOpts)
}

// GetNumberofTransactions is a free data retrieval call binding the contract method 0xf2511d09.
//
// Solidity: function getNumberofTransactions() view returns(uint256)
func (_BlockchainFile *BlockchainFileCaller) GetNumberofTransactions(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BlockchainFile.contract.Call(opts, &out, "getNumberofTransactions")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumberofTransactions is a free data retrieval call binding the contract method 0xf2511d09.
//
// Solidity: function getNumberofTransactions() view returns(uint256)
func (_BlockchainFile *BlockchainFileSession) GetNumberofTransactions() (*big.Int, error) {
	return _BlockchainFile.Contract.GetNumberofTransactions(&_BlockchainFile.CallOpts)
}

// GetNumberofTransactions is a free data retrieval call binding the contract method 0xf2511d09.
//
// Solidity: function getNumberofTransactions() view returns(uint256)
func (_BlockchainFile *BlockchainFileCallerSession) GetNumberofTransactions() (*big.Int, error) {
	return _BlockchainFile.Contract.GetNumberofTransactions(&_BlockchainFile.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_BlockchainFile *BlockchainFileCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BlockchainFile.contract.Call(opts, &out, "getOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_BlockchainFile *BlockchainFileSession) GetOwner() (common.Address, error) {
	return _BlockchainFile.Contract.GetOwner(&_BlockchainFile.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_BlockchainFile *BlockchainFileCallerSession) GetOwner() (common.Address, error) {
	return _BlockchainFile.Contract.GetOwner(&_BlockchainFile.CallOpts)
}

// GetTransactionAtIndex is a free data retrieval call binding the contract method 0xff603b18.
//
// Solidity: function getTransactionAtIndex(uint256 index) view returns((string,string))
func (_BlockchainFile *BlockchainFileCaller) GetTransactionAtIndex(opts *bind.CallOpts, index *big.Int) (FileFileParts, error) {
	var out []interface{}
	err := _BlockchainFile.contract.Call(opts, &out, "getTransactionAtIndex", index)

	if err != nil {
		return *new(FileFileParts), err
	}

	out0 := *abi.ConvertType(out[0], new(FileFileParts)).(*FileFileParts)

	return out0, err

}

// GetTransactionAtIndex is a free data retrieval call binding the contract method 0xff603b18.
//
// Solidity: function getTransactionAtIndex(uint256 index) view returns((string,string))
func (_BlockchainFile *BlockchainFileSession) GetTransactionAtIndex(index *big.Int) (FileFileParts, error) {
	return _BlockchainFile.Contract.GetTransactionAtIndex(&_BlockchainFile.CallOpts, index)
}

// GetTransactionAtIndex is a free data retrieval call binding the contract method 0xff603b18.
//
// Solidity: function getTransactionAtIndex(uint256 index) view returns((string,string))
func (_BlockchainFile *BlockchainFileCallerSession) GetTransactionAtIndex(index *big.Int) (FileFileParts, error) {
	return _BlockchainFile.Contract.GetTransactionAtIndex(&_BlockchainFile.CallOpts, index)
}

// GetVerification is a free data retrieval call binding the contract method 0x1704f098.
//
// Solidity: function getVerification() view returns(bool)
func (_BlockchainFile *BlockchainFileCaller) GetVerification(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BlockchainFile.contract.Call(opts, &out, "getVerification")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetVerification is a free data retrieval call binding the contract method 0x1704f098.
//
// Solidity: function getVerification() view returns(bool)
func (_BlockchainFile *BlockchainFileSession) GetVerification() (bool, error) {
	return _BlockchainFile.Contract.GetVerification(&_BlockchainFile.CallOpts)
}

// GetVerification is a free data retrieval call binding the contract method 0x1704f098.
//
// Solidity: function getVerification() view returns(bool)
func (_BlockchainFile *BlockchainFileCallerSession) GetVerification() (bool, error) {
	return _BlockchainFile.Contract.GetVerification(&_BlockchainFile.CallOpts)
}

// CheckVerification is a paid mutator transaction binding the contract method 0xe02a577e.
//
// Solidity: function checkVerification(address _worldId, string _appId, string _actionId, address signal, uint256 root, uint256 nullifierHash, uint256[8] proof) returns()
func (_BlockchainFile *BlockchainFileTransactor) CheckVerification(opts *bind.TransactOpts, _worldId common.Address, _appId string, _actionId string, signal common.Address, root *big.Int, nullifierHash *big.Int, proof [8]*big.Int) (*types.Transaction, error) {
	return _BlockchainFile.contract.Transact(opts, "checkVerification", _worldId, _appId, _actionId, signal, root, nullifierHash, proof)
}

// CheckVerification is a paid mutator transaction binding the contract method 0xe02a577e.
//
// Solidity: function checkVerification(address _worldId, string _appId, string _actionId, address signal, uint256 root, uint256 nullifierHash, uint256[8] proof) returns()
func (_BlockchainFile *BlockchainFileSession) CheckVerification(_worldId common.Address, _appId string, _actionId string, signal common.Address, root *big.Int, nullifierHash *big.Int, proof [8]*big.Int) (*types.Transaction, error) {
	return _BlockchainFile.Contract.CheckVerification(&_BlockchainFile.TransactOpts, _worldId, _appId, _actionId, signal, root, nullifierHash, proof)
}

// CheckVerification is a paid mutator transaction binding the contract method 0xe02a577e.
//
// Solidity: function checkVerification(address _worldId, string _appId, string _actionId, address signal, uint256 root, uint256 nullifierHash, uint256[8] proof) returns()
func (_BlockchainFile *BlockchainFileTransactorSession) CheckVerification(_worldId common.Address, _appId string, _actionId string, signal common.Address, root *big.Int, nullifierHash *big.Int, proof [8]*big.Int) (*types.Transaction, error) {
	return _BlockchainFile.Contract.CheckVerification(&_BlockchainFile.TransactOpts, _worldId, _appId, _actionId, signal, root, nullifierHash, proof)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BlockchainFile *BlockchainFileTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BlockchainFile.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BlockchainFile *BlockchainFileSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BlockchainFile.Contract.TransferOwnership(&_BlockchainFile.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BlockchainFile *BlockchainFileTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BlockchainFile.Contract.TransferOwnership(&_BlockchainFile.TransactOpts, newOwner)
}

// BlockchainFileOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BlockchainFile contract.
type BlockchainFileOwnershipTransferredIterator struct {
	Event *BlockchainFileOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BlockchainFileOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockchainFileOwnershipTransferred)
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
		it.Event = new(BlockchainFileOwnershipTransferred)
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
func (it *BlockchainFileOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockchainFileOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockchainFileOwnershipTransferred represents a OwnershipTransferred event raised by the BlockchainFile contract.
type BlockchainFileOwnershipTransferred struct {
	FileAddress   common.Address
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0xc8894f26f396ce8c004245c8b7cd1b92103a6e4302fcbab883987149ac01b7ec.
//
// Solidity: event OwnershipTransferred(address indexed fileAddress, address indexed previousOwner, address indexed newOwner)
func (_BlockchainFile *BlockchainFileFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, fileAddress []common.Address, previousOwner []common.Address, newOwner []common.Address) (*BlockchainFileOwnershipTransferredIterator, error) {

	var fileAddressRule []interface{}
	for _, fileAddressItem := range fileAddress {
		fileAddressRule = append(fileAddressRule, fileAddressItem)
	}
	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BlockchainFile.contract.FilterLogs(opts, "OwnershipTransferred", fileAddressRule, previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BlockchainFileOwnershipTransferredIterator{contract: _BlockchainFile.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0xc8894f26f396ce8c004245c8b7cd1b92103a6e4302fcbab883987149ac01b7ec.
//
// Solidity: event OwnershipTransferred(address indexed fileAddress, address indexed previousOwner, address indexed newOwner)
func (_BlockchainFile *BlockchainFileFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BlockchainFileOwnershipTransferred, fileAddress []common.Address, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var fileAddressRule []interface{}
	for _, fileAddressItem := range fileAddress {
		fileAddressRule = append(fileAddressRule, fileAddressItem)
	}
	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BlockchainFile.contract.WatchLogs(opts, "OwnershipTransferred", fileAddressRule, previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockchainFileOwnershipTransferred)
				if err := _BlockchainFile.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0xc8894f26f396ce8c004245c8b7cd1b92103a6e4302fcbab883987149ac01b7ec.
//
// Solidity: event OwnershipTransferred(address indexed fileAddress, address indexed previousOwner, address indexed newOwner)
func (_BlockchainFile *BlockchainFileFilterer) ParseOwnershipTransferred(log types.Log) (*BlockchainFileOwnershipTransferred, error) {
	event := new(BlockchainFileOwnershipTransferred)
	if err := _BlockchainFile.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlockchainFileVerificationResultIterator is returned from FilterVerificationResult and is used to iterate over the raw logs and unpacked data for VerificationResult events raised by the BlockchainFile contract.
type BlockchainFileVerificationResultIterator struct {
	Event *BlockchainFileVerificationResult // Event containing the contract specifics and raw log

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
func (it *BlockchainFileVerificationResultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockchainFileVerificationResult)
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
		it.Event = new(BlockchainFileVerificationResult)
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
func (it *BlockchainFileVerificationResultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockchainFileVerificationResultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockchainFileVerificationResult represents a VerificationResult event raised by the BlockchainFile contract.
type BlockchainFileVerificationResult struct {
	Owner  common.Address
	Result bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterVerificationResult is a free log retrieval operation binding the contract event 0xe41b8ac779069cbbc13495ce67864b788f4ac7e77e4506f6ca7c165847b2d2a3.
//
// Solidity: event VerificationResult(address owner, bool result)
func (_BlockchainFile *BlockchainFileFilterer) FilterVerificationResult(opts *bind.FilterOpts) (*BlockchainFileVerificationResultIterator, error) {

	logs, sub, err := _BlockchainFile.contract.FilterLogs(opts, "VerificationResult")
	if err != nil {
		return nil, err
	}
	return &BlockchainFileVerificationResultIterator{contract: _BlockchainFile.contract, event: "VerificationResult", logs: logs, sub: sub}, nil
}

// WatchVerificationResult is a free log subscription operation binding the contract event 0xe41b8ac779069cbbc13495ce67864b788f4ac7e77e4506f6ca7c165847b2d2a3.
//
// Solidity: event VerificationResult(address owner, bool result)
func (_BlockchainFile *BlockchainFileFilterer) WatchVerificationResult(opts *bind.WatchOpts, sink chan<- *BlockchainFileVerificationResult) (event.Subscription, error) {

	logs, sub, err := _BlockchainFile.contract.WatchLogs(opts, "VerificationResult")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockchainFileVerificationResult)
				if err := _BlockchainFile.contract.UnpackLog(event, "VerificationResult", log); err != nil {
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

// ParseVerificationResult is a log parse operation binding the contract event 0xe41b8ac779069cbbc13495ce67864b788f4ac7e77e4506f6ca7c165847b2d2a3.
//
// Solidity: event VerificationResult(address owner, bool result)
func (_BlockchainFile *BlockchainFileFilterer) ParseVerificationResult(log types.Log) (*BlockchainFileVerificationResult, error) {
	event := new(BlockchainFileVerificationResult)
	if err := _BlockchainFile.contract.UnpackLog(event, "VerificationResult", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
