// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package voting

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

// VotingABI is the input ABI used to generate the binding from.
const VotingABI = "[{\"name\":\"CandidateAdded\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"kind\",\"indexed\":true},{\"type\":\"address\",\"name\":\"owner\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"voteBy\",\"indexed\":false,\"unit\":\"sec\"}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"CandidateRemoved\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Voted\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\",\"indexed\":true},{\"type\":\"address\",\"name\":\"voter\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"market_token_addr\"}],\"constant\":false,\"payable\":false,\"type\":\"constructor\"},{\"name\":\"getPrivileged\",\"outputs\":[{\"type\":\"address\",\"name\":\"out\"},{\"type\":\"address\",\"name\":\"out\"},{\"type\":\"address\",\"name\":\"out\"},{\"type\":\"address\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":1191},{\"name\":\"setPrivileged\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"parameterizer\"},{\"type\":\"address\",\"name\":\"reserve\"},{\"type\":\"address\",\"name\":\"datatrust\"},{\"type\":\"address\",\"name\":\"listing\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":141983},{\"name\":\"hasPrivilege\",\"outputs\":[{\"type\":\"bool\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"sender\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":1284},{\"name\":\"candidateIs\",\"outputs\":[{\"type\":\"bool\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"},{\"type\":\"uint256\",\"name\":\"kind\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":793},{\"name\":\"isCandidate\",\"outputs\":[{\"type\":\"bool\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":823},{\"name\":\"getCandidate\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"},{\"type\":\"address\",\"name\":\"out\"},{\"type\":\"uint256\",\"name\":\"out\",\"unit\":\"wei\"},{\"type\":\"uint256\",\"name\":\"out\",\"unit\":\"sec\"},{\"type\":\"uint256\",\"name\":\"out\"},{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2941},{\"name\":\"getCandidateOwner\",\"outputs\":[{\"type\":\"address\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":874},{\"name\":\"addCandidate\",\"outputs\":[],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"},{\"type\":\"uint256\",\"name\":\"kind\"},{\"type\":\"address\",\"name\":\"owner\"},{\"type\":\"uint256\",\"name\":\"stake\",\"unit\":\"wei\"},{\"type\":\"uint256\",\"name\":\"vote_by\",\"unit\":\"sec\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":184492},{\"name\":\"removeCandidate\",\"outputs\":[],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":124758},{\"name\":\"didPass\",\"outputs\":[{\"type\":\"bool\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"},{\"type\":\"uint256\",\"name\":\"plurality\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":3298},{\"name\":\"pollClosed\",\"outputs\":[{\"type\":\"bool\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":1486},{\"name\":\"vote\",\"outputs\":[],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"},{\"type\":\"uint256\",\"name\":\"option\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":77901},{\"name\":\"transferStake\",\"outputs\":[],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"},{\"type\":\"address\",\"name\":\"addr\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":57681},{\"name\":\"getStake\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\",\"unit\":\"wei\"}],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"},{\"type\":\"address\",\"name\":\"addr\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":1166},{\"name\":\"unstake\",\"outputs\":[],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":24233}]"

// VotingBin is the compiled bytecode used for deploying new contracts.
const VotingBin = `0x740100000000000000000000000000000000000000006020526f7fffffffffffffffffffffffffffffff6040527fffffffffffffffffffffffffffffffff8000000000000000000000000000000060605274012a05f1fffffffffffffffffffffffffdabf41c006080527ffffffffffffffffffffffffed5fa0e000000000000000000000000000000000060a0526020610d606101403934156100a157600080fd5b6020610d6060c03960c05160205181106100ba57600080fd5b503360035561014051600255610d4856600035601c52740100000000000000000000000000000000000000006020526f7fffffffffffffffffffffffffffffff6040527fffffffffffffffffffffffffffffffff8000000000000000000000000000000060605274012a05f1fffffffffffffffffffffffffdabf41c006080527ffffffffffffffffffffffffed5fa0e000000000000000000000000000000000060a05263d4c1753960005114156100dd5734156100ac57600080fd5b60806101405261016060045481526005548160200152600654816040015260075481606001525061014051610160f3005b6339db04e660005114156101a657608060046101403734156100fe57600080fd5b600435602051811061010f57600080fd5b50602435602051811061012157600080fd5b50604435602051811061013357600080fd5b50606435602051811061014557600080fd5b50600354331461015457600080fd5b6004541561016157600080fd5b6005541561016e57600080fd5b6006541561017b57600080fd5b6007541561018857600080fd5b6101405160045561016051600555610180516006556101a051600755005b63b7bf210c600051141561020657602060046101403734156101c757600080fd5b60043560205181106101d857600080fd5b50600554610140511460045461014051141760065461014051141760075461014051141760005260206000f3005b63af61f7606000511415610250576040600461014037341561022757600080fd5b6101605160006101405160e05260c052604060c02060c052602060c020541460005260206000f3005b63b89694c6600051141561029c576020600461014037341561027157600080fd5b6000600160006101405160e05260c052604060c02060c052602060c0200154141560005260206000f3005b63dfb6419f600051141561039757602060046101403734156102bd57600080fd5b60c06101605261018060006101405160e05260c052604060c02060c052602060c020548152600160006101405160e05260c052604060c02060c052602060c02001548160200152600260006101405160e05260c052604060c02060c052602060c02001548160400152600360006101405160e05260c052604060c02060c052602060c02001548160600152600460006101405160e05260c052604060c02060c052602060c02001548160800152600560006101405160e05260c052604060c02060c052602060c02001548160a001525061016051610180f3005b63eb3014fd60005114156103df57602060046101403734156103b857600080fd5b600160006101405160e05260c052604060c02060c052602060c020015460005260206000f3005b63bb12c49e60005114156105d25760a0600461014037341561040057600080fd5b604435602051811061041157600080fd5b506020610260602463b7bf210c6101e05233610200526101fc6000305af161043857600080fd5b6102605161044557600080fd5b600160006101405160e05260c052604060c02060c052602060c02001541561046c57600080fd5b60026101605114156104f8576002543b61048557600080fd5b600254301861049357600080fd5b602061034060646323b872dd61028052610180516102a052306102c0526101a0516102e05261029c60006002545af16104cb57600080fd5b600050610340506101a05160016101805160e05260c052604060c0206101405160e05260c052604060c020555b426101c0514201101561050a57600080fd5b6101c0514201610360526101605160006101405160e05260c052604060c02060c052602060c0205561018051600160006101405160e05260c052604060c02060c052602060c02001556101a051600260006101405160e05260c052604060c02060c052602060c020015561036051600360006101405160e05260c052604060c02060c052602060c020015561036051610380526101805161016051610140517f51baeaa00f5969dc416d8e5299022ac3886d317d13685d1c6906ced9bac71b026020610380a4005b6389bb617c60005114156106bc57602060046101403734156105f357600080fd5b60206101e0602463b7bf210c61016052336101805261017c6000305af161061957600080fd5b6101e05161062657600080fd5b6000600160006101405160e05260c052604060c02060c052602060c02001541861064f57600080fd5b60006101405160e05260c052604060c02060c052602060c02060008155600060018201556000600282015560006003820155600060048201556000600582015550610140517fb314642ce4aa3156c1090458cfccabff50f32bc85d110d7799c369e1a660bcf460006000a2005b638f354b79600051141561081657604060046101403734156106dd57600080fd5b6000600160006101405160e05260c052604060c02060c052602060c02001541861070657600080fd5b42600360006101405160e05260c052604060c02060c052602060c02001541061072e57600080fd5b600460006101405160e05260c052604060c02060c052602060c02001546101805261018051600560006101405160e05260c052604060c02060c052602060c02001546101805101101561078057600080fd5b600560006101405160e05260c052604060c02060c052602060c020015461018051016101a0526101a05115156107c257610160511560005260206000f3610814565b610160516101a0516107d357600080fd5b6101a0516101805115156107e8576000610808565b60646101805160646101805102041461080057600080fd5b606461018051025b04101560005260206000f35b005b63327322c86000511415610889576020600461014037341561083757600080fd5b6000600160006101405160e05260c052604060c02060c052602060c02001541861086057600080fd5b42600360006101405160e05260c052604060c02060c052602060c02001541060005260206000f3005b639ef1204c6000511415610a5957604060046101403734156108aa57600080fd5b6000600160006101405160e05260c052604060c02060c052602060c0200154186108d357600080fd5b42600360006101405160e05260c052604060c02060c052602060c0200154116108fb57600080fd5b600260006101405160e05260c052604060c02060c052602060c0200154610180526002543b61092957600080fd5b600254301861093757600080fd5b602061026060646323b872dd6101a052336101c052306101e05261018051610200526101bc60006002545af161096c57600080fd5b6000506102605060013360e05260c052604060c0206101405160e05260c052604060c02080546101805182540110156109a457600080fd5b6101805181540181555060016101605114156109f557600460006101405160e05260c052604060c02060c052602060c020018054600182540110156109e857600080fd5b6001815401815550610a2c565b600560006101405160e05260c052604060c02060c052602060c02001805460018254011015610a2357600080fd5b60018154018155505b33610140517f0b2f654b7e608ce51a82ce8157e79c350ed670605e8985266ad89fc85060e74960006000a3005b63b6b692066000511415610b3b5760406004610140373415610a7a57600080fd5b6024356020518110610a8b57600080fd5b506007543314610a9a57600080fd5b6001600160006101405160e05260c052604060c02060c052602060c020015460e05260c052604060c0206101405160e05260c052604060c020546101805260006001600160006101405160e05260c052604060c02060c052602060c020015460e05260c052604060c0206101405160e05260c052604060c020556101805160016101605160e05260c052604060c0206101405160e05260c052604060c02055005b635399ab4c6000511415610b995760406004610140373415610b5c57600080fd5b6024356020518110610b6d57600080fd5b5060016101605160e05260c052604060c0206101405160e05260c052604060c0205460005260206000f3005b6371ed5d1a6000511415610c775760206004610140373415610bba57600080fd5b600160006101405160e05260c052604060c02060c052602060c020015415610be157600080fd5b60013360e05260c052604060c0206101405160e05260c052604060c0205461016052600060013360e05260c052604060c0206101405160e05260c052604060c020556002543b610c3057600080fd5b6002543018610c3e57600080fd5b6020610220604463a9059cbb61018052336101a052610160516101c05261019c60006002545af1610c6e57600080fd5b60005061022050005b60006000fd5b6100cb610d48036100cb6000396100cb610d48036000f3`

// DeployVoting deploys a new Ethereum contract, binding an instance of Voting to it.
func DeployVoting(auth *bind.TransactOpts, backend bind.ContractBackend, market_token_addr common.Address) (common.Address, *types.Transaction, *Voting, error) {
	parsed, err := abi.JSON(strings.NewReader(VotingABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(VotingBin), backend, market_token_addr)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Voting{VotingCaller: VotingCaller{contract: contract}, VotingTransactor: VotingTransactor{contract: contract}, VotingFilterer: VotingFilterer{contract: contract}}, nil
}

// Voting is an auto generated Go binding around an Ethereum contract.
type Voting struct {
	VotingCaller     // Read-only binding to the contract
	VotingTransactor // Write-only binding to the contract
	VotingFilterer   // Log filterer for contract events
}

// VotingCaller is an auto generated read-only Go binding around an Ethereum contract.
type VotingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VotingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VotingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VotingSession struct {
	Contract     *Voting           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VotingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VotingCallerSession struct {
	Contract *VotingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VotingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VotingTransactorSession struct {
	Contract     *VotingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VotingRaw is an auto generated low-level Go binding around an Ethereum contract.
type VotingRaw struct {
	Contract *Voting // Generic contract binding to access the raw methods on
}

// VotingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VotingCallerRaw struct {
	Contract *VotingCaller // Generic read-only contract binding to access the raw methods on
}

// VotingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VotingTransactorRaw struct {
	Contract *VotingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVoting creates a new instance of Voting, bound to a specific deployed contract.
func NewVoting(address common.Address, backend bind.ContractBackend) (*Voting, error) {
	contract, err := bindVoting(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Voting{VotingCaller: VotingCaller{contract: contract}, VotingTransactor: VotingTransactor{contract: contract}, VotingFilterer: VotingFilterer{contract: contract}}, nil
}

// NewVotingCaller creates a new read-only instance of Voting, bound to a specific deployed contract.
func NewVotingCaller(address common.Address, caller bind.ContractCaller) (*VotingCaller, error) {
	contract, err := bindVoting(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VotingCaller{contract: contract}, nil
}

// NewVotingTransactor creates a new write-only instance of Voting, bound to a specific deployed contract.
func NewVotingTransactor(address common.Address, transactor bind.ContractTransactor) (*VotingTransactor, error) {
	contract, err := bindVoting(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VotingTransactor{contract: contract}, nil
}

// NewVotingFilterer creates a new log filterer instance of Voting, bound to a specific deployed contract.
func NewVotingFilterer(address common.Address, filterer bind.ContractFilterer) (*VotingFilterer, error) {
	contract, err := bindVoting(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VotingFilterer{contract: contract}, nil
}

// bindVoting binds a generic wrapper to an already deployed contract.
func bindVoting(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VotingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Voting *VotingRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Voting.Contract.VotingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Voting *VotingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Voting.Contract.VotingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Voting *VotingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Voting.Contract.VotingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Voting *VotingCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Voting.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Voting *VotingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Voting.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Voting *VotingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Voting.Contract.contract.Transact(opts, method, params...)
}

// CandidateIs is a free data retrieval call binding the contract method 0xaf61f760.
//
// Solidity: function candidateIs(hash bytes32, kind uint256) constant returns(out bool)
func (_Voting *VotingCaller) CandidateIs(opts *bind.CallOpts, hash [32]byte, kind *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Voting.contract.Call(opts, out, "candidateIs", hash, kind)
	return *ret0, err
}

// CandidateIs is a free data retrieval call binding the contract method 0xaf61f760.
//
// Solidity: function candidateIs(hash bytes32, kind uint256) constant returns(out bool)
func (_Voting *VotingSession) CandidateIs(hash [32]byte, kind *big.Int) (bool, error) {
	return _Voting.Contract.CandidateIs(&_Voting.CallOpts, hash, kind)
}

// CandidateIs is a free data retrieval call binding the contract method 0xaf61f760.
//
// Solidity: function candidateIs(hash bytes32, kind uint256) constant returns(out bool)
func (_Voting *VotingCallerSession) CandidateIs(hash [32]byte, kind *big.Int) (bool, error) {
	return _Voting.Contract.CandidateIs(&_Voting.CallOpts, hash, kind)
}

// DidPass is a free data retrieval call binding the contract method 0x8f354b79.
//
// Solidity: function didPass(hash bytes32, plurality uint256) constant returns(out bool)
func (_Voting *VotingCaller) DidPass(opts *bind.CallOpts, hash [32]byte, plurality *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Voting.contract.Call(opts, out, "didPass", hash, plurality)
	return *ret0, err
}

// DidPass is a free data retrieval call binding the contract method 0x8f354b79.
//
// Solidity: function didPass(hash bytes32, plurality uint256) constant returns(out bool)
func (_Voting *VotingSession) DidPass(hash [32]byte, plurality *big.Int) (bool, error) {
	return _Voting.Contract.DidPass(&_Voting.CallOpts, hash, plurality)
}

// DidPass is a free data retrieval call binding the contract method 0x8f354b79.
//
// Solidity: function didPass(hash bytes32, plurality uint256) constant returns(out bool)
func (_Voting *VotingCallerSession) DidPass(hash [32]byte, plurality *big.Int) (bool, error) {
	return _Voting.Contract.DidPass(&_Voting.CallOpts, hash, plurality)
}

// GetCandidate is a free data retrieval call binding the contract method 0xdfb6419f.
//
// Solidity: function getCandidate(hash bytes32) constant returns(out uint256, out address, out uint256, out uint256, out uint256, out uint256)
func (_Voting *VotingCaller) GetCandidate(opts *bind.CallOpts, hash [32]byte) (*big.Int, common.Address, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(common.Address)
		ret2 = new(*big.Int)
		ret3 = new(*big.Int)
		ret4 = new(*big.Int)
		ret5 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
	}
	err := _Voting.contract.Call(opts, out, "getCandidate", hash)
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, err
}

// GetCandidate is a free data retrieval call binding the contract method 0xdfb6419f.
//
// Solidity: function getCandidate(hash bytes32) constant returns(out uint256, out address, out uint256, out uint256, out uint256, out uint256)
func (_Voting *VotingSession) GetCandidate(hash [32]byte) (*big.Int, common.Address, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Voting.Contract.GetCandidate(&_Voting.CallOpts, hash)
}

// GetCandidate is a free data retrieval call binding the contract method 0xdfb6419f.
//
// Solidity: function getCandidate(hash bytes32) constant returns(out uint256, out address, out uint256, out uint256, out uint256, out uint256)
func (_Voting *VotingCallerSession) GetCandidate(hash [32]byte) (*big.Int, common.Address, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Voting.Contract.GetCandidate(&_Voting.CallOpts, hash)
}

// GetCandidateOwner is a free data retrieval call binding the contract method 0xeb3014fd.
//
// Solidity: function getCandidateOwner(hash bytes32) constant returns(out address)
func (_Voting *VotingCaller) GetCandidateOwner(opts *bind.CallOpts, hash [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Voting.contract.Call(opts, out, "getCandidateOwner", hash)
	return *ret0, err
}

// GetCandidateOwner is a free data retrieval call binding the contract method 0xeb3014fd.
//
// Solidity: function getCandidateOwner(hash bytes32) constant returns(out address)
func (_Voting *VotingSession) GetCandidateOwner(hash [32]byte) (common.Address, error) {
	return _Voting.Contract.GetCandidateOwner(&_Voting.CallOpts, hash)
}

// GetCandidateOwner is a free data retrieval call binding the contract method 0xeb3014fd.
//
// Solidity: function getCandidateOwner(hash bytes32) constant returns(out address)
func (_Voting *VotingCallerSession) GetCandidateOwner(hash [32]byte) (common.Address, error) {
	return _Voting.Contract.GetCandidateOwner(&_Voting.CallOpts, hash)
}

// GetPrivileged is a free data retrieval call binding the contract method 0xd4c17539.
//
// Solidity: function getPrivileged() constant returns(out address, out address, out address, out address)
func (_Voting *VotingCaller) GetPrivileged(opts *bind.CallOpts) (common.Address, common.Address, common.Address, common.Address, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(common.Address)
		ret2 = new(common.Address)
		ret3 = new(common.Address)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
	}
	err := _Voting.contract.Call(opts, out, "getPrivileged")
	return *ret0, *ret1, *ret2, *ret3, err
}

// GetPrivileged is a free data retrieval call binding the contract method 0xd4c17539.
//
// Solidity: function getPrivileged() constant returns(out address, out address, out address, out address)
func (_Voting *VotingSession) GetPrivileged() (common.Address, common.Address, common.Address, common.Address, error) {
	return _Voting.Contract.GetPrivileged(&_Voting.CallOpts)
}

// GetPrivileged is a free data retrieval call binding the contract method 0xd4c17539.
//
// Solidity: function getPrivileged() constant returns(out address, out address, out address, out address)
func (_Voting *VotingCallerSession) GetPrivileged() (common.Address, common.Address, common.Address, common.Address, error) {
	return _Voting.Contract.GetPrivileged(&_Voting.CallOpts)
}

// GetStake is a free data retrieval call binding the contract method 0x5399ab4c.
//
// Solidity: function getStake(hash bytes32, addr address) constant returns(out uint256)
func (_Voting *VotingCaller) GetStake(opts *bind.CallOpts, hash [32]byte, addr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Voting.contract.Call(opts, out, "getStake", hash, addr)
	return *ret0, err
}

// GetStake is a free data retrieval call binding the contract method 0x5399ab4c.
//
// Solidity: function getStake(hash bytes32, addr address) constant returns(out uint256)
func (_Voting *VotingSession) GetStake(hash [32]byte, addr common.Address) (*big.Int, error) {
	return _Voting.Contract.GetStake(&_Voting.CallOpts, hash, addr)
}

// GetStake is a free data retrieval call binding the contract method 0x5399ab4c.
//
// Solidity: function getStake(hash bytes32, addr address) constant returns(out uint256)
func (_Voting *VotingCallerSession) GetStake(hash [32]byte, addr common.Address) (*big.Int, error) {
	return _Voting.Contract.GetStake(&_Voting.CallOpts, hash, addr)
}

// HasPrivilege is a free data retrieval call binding the contract method 0xb7bf210c.
//
// Solidity: function hasPrivilege(sender address) constant returns(out bool)
func (_Voting *VotingCaller) HasPrivilege(opts *bind.CallOpts, sender common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Voting.contract.Call(opts, out, "hasPrivilege", sender)
	return *ret0, err
}

// HasPrivilege is a free data retrieval call binding the contract method 0xb7bf210c.
//
// Solidity: function hasPrivilege(sender address) constant returns(out bool)
func (_Voting *VotingSession) HasPrivilege(sender common.Address) (bool, error) {
	return _Voting.Contract.HasPrivilege(&_Voting.CallOpts, sender)
}

// HasPrivilege is a free data retrieval call binding the contract method 0xb7bf210c.
//
// Solidity: function hasPrivilege(sender address) constant returns(out bool)
func (_Voting *VotingCallerSession) HasPrivilege(sender common.Address) (bool, error) {
	return _Voting.Contract.HasPrivilege(&_Voting.CallOpts, sender)
}

// IsCandidate is a free data retrieval call binding the contract method 0xb89694c6.
//
// Solidity: function isCandidate(hash bytes32) constant returns(out bool)
func (_Voting *VotingCaller) IsCandidate(opts *bind.CallOpts, hash [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Voting.contract.Call(opts, out, "isCandidate", hash)
	return *ret0, err
}

// IsCandidate is a free data retrieval call binding the contract method 0xb89694c6.
//
// Solidity: function isCandidate(hash bytes32) constant returns(out bool)
func (_Voting *VotingSession) IsCandidate(hash [32]byte) (bool, error) {
	return _Voting.Contract.IsCandidate(&_Voting.CallOpts, hash)
}

// IsCandidate is a free data retrieval call binding the contract method 0xb89694c6.
//
// Solidity: function isCandidate(hash bytes32) constant returns(out bool)
func (_Voting *VotingCallerSession) IsCandidate(hash [32]byte) (bool, error) {
	return _Voting.Contract.IsCandidate(&_Voting.CallOpts, hash)
}

// PollClosed is a free data retrieval call binding the contract method 0x327322c8.
//
// Solidity: function pollClosed(hash bytes32) constant returns(out bool)
func (_Voting *VotingCaller) PollClosed(opts *bind.CallOpts, hash [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Voting.contract.Call(opts, out, "pollClosed", hash)
	return *ret0, err
}

// PollClosed is a free data retrieval call binding the contract method 0x327322c8.
//
// Solidity: function pollClosed(hash bytes32) constant returns(out bool)
func (_Voting *VotingSession) PollClosed(hash [32]byte) (bool, error) {
	return _Voting.Contract.PollClosed(&_Voting.CallOpts, hash)
}

// PollClosed is a free data retrieval call binding the contract method 0x327322c8.
//
// Solidity: function pollClosed(hash bytes32) constant returns(out bool)
func (_Voting *VotingCallerSession) PollClosed(hash [32]byte) (bool, error) {
	return _Voting.Contract.PollClosed(&_Voting.CallOpts, hash)
}

// AddCandidate is a paid mutator transaction binding the contract method 0xbb12c49e.
//
// Solidity: function addCandidate(hash bytes32, kind uint256, owner address, stake uint256, vote_by uint256) returns()
func (_Voting *VotingTransactor) AddCandidate(opts *bind.TransactOpts, hash [32]byte, kind *big.Int, owner common.Address, stake *big.Int, vote_by *big.Int) (*types.Transaction, error) {
	return _Voting.contract.Transact(opts, "addCandidate", hash, kind, owner, stake, vote_by)
}

// AddCandidate is a paid mutator transaction binding the contract method 0xbb12c49e.
//
// Solidity: function addCandidate(hash bytes32, kind uint256, owner address, stake uint256, vote_by uint256) returns()
func (_Voting *VotingSession) AddCandidate(hash [32]byte, kind *big.Int, owner common.Address, stake *big.Int, vote_by *big.Int) (*types.Transaction, error) {
	return _Voting.Contract.AddCandidate(&_Voting.TransactOpts, hash, kind, owner, stake, vote_by)
}

// AddCandidate is a paid mutator transaction binding the contract method 0xbb12c49e.
//
// Solidity: function addCandidate(hash bytes32, kind uint256, owner address, stake uint256, vote_by uint256) returns()
func (_Voting *VotingTransactorSession) AddCandidate(hash [32]byte, kind *big.Int, owner common.Address, stake *big.Int, vote_by *big.Int) (*types.Transaction, error) {
	return _Voting.Contract.AddCandidate(&_Voting.TransactOpts, hash, kind, owner, stake, vote_by)
}

// RemoveCandidate is a paid mutator transaction binding the contract method 0x89bb617c.
//
// Solidity: function removeCandidate(hash bytes32) returns()
func (_Voting *VotingTransactor) RemoveCandidate(opts *bind.TransactOpts, hash [32]byte) (*types.Transaction, error) {
	return _Voting.contract.Transact(opts, "removeCandidate", hash)
}

// RemoveCandidate is a paid mutator transaction binding the contract method 0x89bb617c.
//
// Solidity: function removeCandidate(hash bytes32) returns()
func (_Voting *VotingSession) RemoveCandidate(hash [32]byte) (*types.Transaction, error) {
	return _Voting.Contract.RemoveCandidate(&_Voting.TransactOpts, hash)
}

// RemoveCandidate is a paid mutator transaction binding the contract method 0x89bb617c.
//
// Solidity: function removeCandidate(hash bytes32) returns()
func (_Voting *VotingTransactorSession) RemoveCandidate(hash [32]byte) (*types.Transaction, error) {
	return _Voting.Contract.RemoveCandidate(&_Voting.TransactOpts, hash)
}

// SetPrivileged is a paid mutator transaction binding the contract method 0x39db04e6.
//
// Solidity: function setPrivileged(parameterizer address, reserve address, datatrust address, listing address) returns()
func (_Voting *VotingTransactor) SetPrivileged(opts *bind.TransactOpts, parameterizer common.Address, reserve common.Address, datatrust common.Address, listing common.Address) (*types.Transaction, error) {
	return _Voting.contract.Transact(opts, "setPrivileged", parameterizer, reserve, datatrust, listing)
}

// SetPrivileged is a paid mutator transaction binding the contract method 0x39db04e6.
//
// Solidity: function setPrivileged(parameterizer address, reserve address, datatrust address, listing address) returns()
func (_Voting *VotingSession) SetPrivileged(parameterizer common.Address, reserve common.Address, datatrust common.Address, listing common.Address) (*types.Transaction, error) {
	return _Voting.Contract.SetPrivileged(&_Voting.TransactOpts, parameterizer, reserve, datatrust, listing)
}

// SetPrivileged is a paid mutator transaction binding the contract method 0x39db04e6.
//
// Solidity: function setPrivileged(parameterizer address, reserve address, datatrust address, listing address) returns()
func (_Voting *VotingTransactorSession) SetPrivileged(parameterizer common.Address, reserve common.Address, datatrust common.Address, listing common.Address) (*types.Transaction, error) {
	return _Voting.Contract.SetPrivileged(&_Voting.TransactOpts, parameterizer, reserve, datatrust, listing)
}

// TransferStake is a paid mutator transaction binding the contract method 0xb6b69206.
//
// Solidity: function transferStake(hash bytes32, addr address) returns()
func (_Voting *VotingTransactor) TransferStake(opts *bind.TransactOpts, hash [32]byte, addr common.Address) (*types.Transaction, error) {
	return _Voting.contract.Transact(opts, "transferStake", hash, addr)
}

// TransferStake is a paid mutator transaction binding the contract method 0xb6b69206.
//
// Solidity: function transferStake(hash bytes32, addr address) returns()
func (_Voting *VotingSession) TransferStake(hash [32]byte, addr common.Address) (*types.Transaction, error) {
	return _Voting.Contract.TransferStake(&_Voting.TransactOpts, hash, addr)
}

// TransferStake is a paid mutator transaction binding the contract method 0xb6b69206.
//
// Solidity: function transferStake(hash bytes32, addr address) returns()
func (_Voting *VotingTransactorSession) TransferStake(hash [32]byte, addr common.Address) (*types.Transaction, error) {
	return _Voting.Contract.TransferStake(&_Voting.TransactOpts, hash, addr)
}

// Unstake is a paid mutator transaction binding the contract method 0x71ed5d1a.
//
// Solidity: function unstake(hash bytes32) returns()
func (_Voting *VotingTransactor) Unstake(opts *bind.TransactOpts, hash [32]byte) (*types.Transaction, error) {
	return _Voting.contract.Transact(opts, "unstake", hash)
}

// Unstake is a paid mutator transaction binding the contract method 0x71ed5d1a.
//
// Solidity: function unstake(hash bytes32) returns()
func (_Voting *VotingSession) Unstake(hash [32]byte) (*types.Transaction, error) {
	return _Voting.Contract.Unstake(&_Voting.TransactOpts, hash)
}

// Unstake is a paid mutator transaction binding the contract method 0x71ed5d1a.
//
// Solidity: function unstake(hash bytes32) returns()
func (_Voting *VotingTransactorSession) Unstake(hash [32]byte) (*types.Transaction, error) {
	return _Voting.Contract.Unstake(&_Voting.TransactOpts, hash)
}

// Vote is a paid mutator transaction binding the contract method 0x9ef1204c.
//
// Solidity: function vote(hash bytes32, option uint256) returns()
func (_Voting *VotingTransactor) Vote(opts *bind.TransactOpts, hash [32]byte, option *big.Int) (*types.Transaction, error) {
	return _Voting.contract.Transact(opts, "vote", hash, option)
}

// Vote is a paid mutator transaction binding the contract method 0x9ef1204c.
//
// Solidity: function vote(hash bytes32, option uint256) returns()
func (_Voting *VotingSession) Vote(hash [32]byte, option *big.Int) (*types.Transaction, error) {
	return _Voting.Contract.Vote(&_Voting.TransactOpts, hash, option)
}

// Vote is a paid mutator transaction binding the contract method 0x9ef1204c.
//
// Solidity: function vote(hash bytes32, option uint256) returns()
func (_Voting *VotingTransactorSession) Vote(hash [32]byte, option *big.Int) (*types.Transaction, error) {
	return _Voting.Contract.Vote(&_Voting.TransactOpts, hash, option)
}

// VotingCandidateAddedIterator is returned from FilterCandidateAdded and is used to iterate over the raw logs and unpacked data for CandidateAdded events raised by the Voting contract.
type VotingCandidateAddedIterator struct {
	Event *VotingCandidateAdded // Event containing the contract specifics and raw log

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
func (it *VotingCandidateAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotingCandidateAdded)
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
		it.Event = new(VotingCandidateAdded)
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
func (it *VotingCandidateAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotingCandidateAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotingCandidateAdded represents a CandidateAdded event raised by the Voting contract.
type VotingCandidateAdded struct {
	Hash   [32]byte
	Kind   *big.Int
	Owner  common.Address
	VoteBy *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCandidateAdded is a free log retrieval operation binding the contract event 0x51baeaa00f5969dc416d8e5299022ac3886d317d13685d1c6906ced9bac71b02.
//
// Solidity: e CandidateAdded(hash indexed bytes32, kind indexed uint256, owner indexed address, voteBy uint256)
func (_Voting *VotingFilterer) FilterCandidateAdded(opts *bind.FilterOpts, hash [][32]byte, kind []*big.Int, owner []common.Address) (*VotingCandidateAddedIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var kindRule []interface{}
	for _, kindItem := range kind {
		kindRule = append(kindRule, kindItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Voting.contract.FilterLogs(opts, "CandidateAdded", hashRule, kindRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &VotingCandidateAddedIterator{contract: _Voting.contract, event: "CandidateAdded", logs: logs, sub: sub}, nil
}

// WatchCandidateAdded is a free log subscription operation binding the contract event 0x51baeaa00f5969dc416d8e5299022ac3886d317d13685d1c6906ced9bac71b02.
//
// Solidity: e CandidateAdded(hash indexed bytes32, kind indexed uint256, owner indexed address, voteBy uint256)
func (_Voting *VotingFilterer) WatchCandidateAdded(opts *bind.WatchOpts, sink chan<- *VotingCandidateAdded, hash [][32]byte, kind []*big.Int, owner []common.Address) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var kindRule []interface{}
	for _, kindItem := range kind {
		kindRule = append(kindRule, kindItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Voting.contract.WatchLogs(opts, "CandidateAdded", hashRule, kindRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotingCandidateAdded)
				if err := _Voting.contract.UnpackLog(event, "CandidateAdded", log); err != nil {
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

// VotingCandidateRemovedIterator is returned from FilterCandidateRemoved and is used to iterate over the raw logs and unpacked data for CandidateRemoved events raised by the Voting contract.
type VotingCandidateRemovedIterator struct {
	Event *VotingCandidateRemoved // Event containing the contract specifics and raw log

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
func (it *VotingCandidateRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotingCandidateRemoved)
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
		it.Event = new(VotingCandidateRemoved)
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
func (it *VotingCandidateRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotingCandidateRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotingCandidateRemoved represents a CandidateRemoved event raised by the Voting contract.
type VotingCandidateRemoved struct {
	Hash [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterCandidateRemoved is a free log retrieval operation binding the contract event 0xb314642ce4aa3156c1090458cfccabff50f32bc85d110d7799c369e1a660bcf4.
//
// Solidity: e CandidateRemoved(hash indexed bytes32)
func (_Voting *VotingFilterer) FilterCandidateRemoved(opts *bind.FilterOpts, hash [][32]byte) (*VotingCandidateRemovedIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _Voting.contract.FilterLogs(opts, "CandidateRemoved", hashRule)
	if err != nil {
		return nil, err
	}
	return &VotingCandidateRemovedIterator{contract: _Voting.contract, event: "CandidateRemoved", logs: logs, sub: sub}, nil
}

// WatchCandidateRemoved is a free log subscription operation binding the contract event 0xb314642ce4aa3156c1090458cfccabff50f32bc85d110d7799c369e1a660bcf4.
//
// Solidity: e CandidateRemoved(hash indexed bytes32)
func (_Voting *VotingFilterer) WatchCandidateRemoved(opts *bind.WatchOpts, sink chan<- *VotingCandidateRemoved, hash [][32]byte) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _Voting.contract.WatchLogs(opts, "CandidateRemoved", hashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotingCandidateRemoved)
				if err := _Voting.contract.UnpackLog(event, "CandidateRemoved", log); err != nil {
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

// VotingVotedIterator is returned from FilterVoted and is used to iterate over the raw logs and unpacked data for Voted events raised by the Voting contract.
type VotingVotedIterator struct {
	Event *VotingVoted // Event containing the contract specifics and raw log

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
func (it *VotingVotedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotingVoted)
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
		it.Event = new(VotingVoted)
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
func (it *VotingVotedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotingVotedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotingVoted represents a Voted event raised by the Voting contract.
type VotingVoted struct {
	Hash  [32]byte
	Voter common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterVoted is a free log retrieval operation binding the contract event 0x0b2f654b7e608ce51a82ce8157e79c350ed670605e8985266ad89fc85060e749.
//
// Solidity: e Voted(hash indexed bytes32, voter indexed address)
func (_Voting *VotingFilterer) FilterVoted(opts *bind.FilterOpts, hash [][32]byte, voter []common.Address) (*VotingVotedIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _Voting.contract.FilterLogs(opts, "Voted", hashRule, voterRule)
	if err != nil {
		return nil, err
	}
	return &VotingVotedIterator{contract: _Voting.contract, event: "Voted", logs: logs, sub: sub}, nil
}

// WatchVoted is a free log subscription operation binding the contract event 0x0b2f654b7e608ce51a82ce8157e79c350ed670605e8985266ad89fc85060e749.
//
// Solidity: e Voted(hash indexed bytes32, voter indexed address)
func (_Voting *VotingFilterer) WatchVoted(opts *bind.WatchOpts, sink chan<- *VotingVoted, hash [][32]byte, voter []common.Address) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _Voting.contract.WatchLogs(opts, "Voted", hashRule, voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotingVoted)
				if err := _Voting.contract.UnpackLog(event, "Voted", log); err != nil {
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
