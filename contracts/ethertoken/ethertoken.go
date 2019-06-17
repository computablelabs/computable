// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethertoken

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

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// EtherTokenABI is the input ABI used to generate the binding from.
const EtherTokenABI = "[{\"name\":\"Approval\",\"inputs\":[{\"type\":\"address\",\"name\":\"owner\",\"indexed\":true},{\"type\":\"address\",\"name\":\"spender\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"amount\",\"indexed\":false,\"unit\":\"wei\"}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Deposited\",\"inputs\":[{\"type\":\"address\",\"name\":\"source\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"amount\",\"indexed\":false,\"unit\":\"wei\"}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Transfer\",\"inputs\":[{\"type\":\"address\",\"name\":\"source\",\"indexed\":true},{\"type\":\"address\",\"name\":\"to\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"amount\",\"indexed\":false,\"unit\":\"wei\"}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Withdrawn\",\"inputs\":[{\"type\":\"address\",\"name\":\"to\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"amount\",\"indexed\":false,\"unit\":\"wei\"}],\"anonymous\":false,\"type\":\"event\"},{\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"initial_account\"},{\"type\":\"uint256\",\"name\":\"initial_balance\",\"unit\":\"wei\"}],\"constant\":false,\"payable\":false,\"type\":\"constructor\"},{\"name\":\"allowance\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\",\"unit\":\"wei\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"owner\"},{\"type\":\"address\",\"name\":\"spender\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":815},{\"name\":\"approve\",\"outputs\":[{\"type\":\"bool\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"spender\"},{\"type\":\"uint256\",\"name\":\"amount\",\"unit\":\"wei\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":37719},{\"name\":\"balanceOf\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\",\"unit\":\"wei\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"owner\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":715},{\"name\":\"decreaseApproval\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"spender\"},{\"type\":\"uint256\",\"name\":\"amount\",\"unit\":\"wei\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":38742},{\"name\":\"deposit\",\"outputs\":[],\"inputs\":[],\"constant\":false,\"payable\":true,\"type\":\"function\",\"gas\":73656},{\"name\":\"increaseApproval\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"spender\"},{\"type\":\"uint256\",\"name\":\"amount\",\"unit\":\"wei\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":39008},{\"name\":\"totalSupply\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\",\"unit\":\"wei\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":663},{\"name\":\"transfer\",\"outputs\":[{\"type\":\"bool\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"to\"},{\"type\":\"uint256\",\"name\":\"amount\",\"unit\":\"wei\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":74321},{\"name\":\"transferFrom\",\"outputs\":[{\"type\":\"bool\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"source\"},{\"type\":\"address\",\"name\":\"to\"},{\"type\":\"uint256\",\"name\":\"amount\",\"unit\":\"wei\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":110291},{\"name\":\"withdraw\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256\",\"name\":\"amount\",\"unit\":\"wei\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":108330},{\"name\":\"decimals\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":783},{\"name\":\"symbol\",\"outputs\":[{\"type\":\"string\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2291}]"

// EtherTokenBin is the compiled bytecode used for deploying new contracts.
const EtherTokenBin = `0x740100000000000000000000000000000000000000006020526f7fffffffffffffffffffffffffffffff6040527fffffffffffffffffffffffffffffffff8000000000000000000000000000000060605274012a05f1fffffffffffffffffffffffffdabf41c006080527ffffffffffffffffffffffffed5fa0e000000000000000000000000000000000060a05260406109596101403934156100a157600080fd5b602061095960c03960c05160205181106100ba57600080fd5b506101605160016101405160e05260c052604060c0205560126002556003610180527f43455400000000000000000000000000000000000000000000000000000000006101a05261018080600360c052602060c020602082510161012060006002818352015b8261012051602002111561013357610155565b61012051602002850151610120518501555b8151600101808352811415610120575b5050505050506101605160045561094156600035601c52740100000000000000000000000000000000000000006020526f7fffffffffffffffffffffffffffffff6040527fffffffffffffffffffffffffffffffff8000000000000000000000000000000060605274012a05f1fffffffffffffffffffffffffdabf41c006080527ffffffffffffffffffffffffed5fa0e000000000000000000000000000000000060a05263dd62ed3e600051141561010357604060046101403734156100b457600080fd5b60043560205181106100c557600080fd5b5060243560205181106100d757600080fd5b5060006101405160e05260c052604060c0206101605160e05260c052604060c0205460005260206000f3005b63095ea7b36000511415610198576040600461014037341561012457600080fd5b600435602051811061013557600080fd5b506101605160003360e05260c052604060c0206101405160e05260c052604060c02055610160516101805261014051337f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9256020610180a3600160005260206000f3005b6370a0823160005114156101e757602060046101403734156101b957600080fd5b60043560205181106101ca57600080fd5b5060016101405160e05260c052604060c0205460005260206000f3005b636618846360005114156102a2576040600461014037341561020857600080fd5b600435602051811061021957600080fd5b5060003360e05260c052604060c0206101405160e05260c052604060c020610160518154101561024857600080fd5b6101605181540381555060003360e05260c052604060c0206101405160e05260c052604060c020546101805261014051337f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9256020610180a3005b63d0e30db0600051141561031f5760013360e05260c052604060c02080543482540110156102cf57600080fd5b34815401815550600480543482540110156102e957600080fd5b348154018155503461014052337f2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c46020610140a2005b63d73dd62360005114156103dd576040600461014037341561034057600080fd5b600435602051811061035157600080fd5b5060003360e05260c052604060c0206101405160e05260c052604060c020805461016051825401101561038357600080fd5b6101605181540181555060003360e05260c052604060c0206101405160e05260c052604060c020546101805261014051337f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9256020610180a3005b6318160ddd60005114156104035734156103f657600080fd5b60045460005260206000f3005b63a9059cbb60005114156104de576040600461014037341561042457600080fd5b600435602051811061043557600080fd5b506000610140511861044657600080fd5b60013360e05260c052604060c020610160518154101561046557600080fd5b6101605181540381555060016101405160e05260c052604060c020805461016051825401101561049457600080fd5b61016051815401815550610160516101805261014051337fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef6020610180a3600160005260206000f3005b6323b872dd600051141561061957606060046101403734156104ff57600080fd5b600435602051811061051057600080fd5b50602435602051811061052257600080fd5b506000610140511861053357600080fd5b6000610160511861054357600080fd5b60016101405160e05260c052604060c020610180518154101561056557600080fd5b6101805181540381555060016101605160e05260c052604060c020805461018051825401101561059457600080fd5b6101805181540181555060006101405160e05260c052604060c0203360e05260c052604060c02061018051815410156105cc57600080fd5b61018051815403815550610180516101a05261016051610140517fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60206101a0a3600160005260206000f3005b632e1a7d4d60005114156106cb576020600461014037341561063a57600080fd5b60013360e05260c052604060c020610140518154101561065957600080fd5b610140518154038155506004610140518154101561067657600080fd5b61014051815403815550600060006000600061014051336000f161069957600080fd5b6101405161016052337f7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d56020610160a2005b63313ce56760005114156106f15734156106e457600080fd5b60025460005260206000f3005b6395d89b4160005114156107d457341561070a57600080fd5b60038060c052602060c020610180602082540161012060006002818352015b8261012051602002111561073c5761075e565b61012051850154610120516020028501525b8151600101808352811415610729575b5050505050506101805160206001820306601f82010390506101e0610180516003818352015b826101e0511115610794576107b0565b60006101e0516101a001535b8151600101808352811415610784575b5050506020610160526040610180510160206001820306601f8201039050610160f3005b60006000fd5b61016761094103610167600039610167610941036000f3`

// DeployEtherToken deploys a new Ethereum contract, binding an instance of EtherToken to it.
func DeployEtherToken(auth *bind.TransactOpts, backend bind.ContractBackend, initial_account common.Address, initial_balance *big.Int) (common.Address, *types.Transaction, *EtherToken, error) {
	parsed, err := abi.JSON(strings.NewReader(EtherTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(EtherTokenBin), backend, initial_account, initial_balance)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EtherToken{EtherTokenCaller: EtherTokenCaller{contract: contract}, EtherTokenTransactor: EtherTokenTransactor{contract: contract}, EtherTokenFilterer: EtherTokenFilterer{contract: contract}}, nil
}

// EtherToken is an auto generated Go binding around an Ethereum contract.
type EtherToken struct {
	EtherTokenCaller     // Read-only binding to the contract
	EtherTokenTransactor // Write-only binding to the contract
	EtherTokenFilterer   // Log filterer for contract events
}

// EtherTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type EtherTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtherTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EtherTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtherTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EtherTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtherTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EtherTokenSession struct {
	Contract     *EtherToken       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EtherTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EtherTokenCallerSession struct {
	Contract *EtherTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// EtherTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EtherTokenTransactorSession struct {
	Contract     *EtherTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// EtherTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type EtherTokenRaw struct {
	Contract *EtherToken // Generic contract binding to access the raw methods on
}

// EtherTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EtherTokenCallerRaw struct {
	Contract *EtherTokenCaller // Generic read-only contract binding to access the raw methods on
}

// EtherTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EtherTokenTransactorRaw struct {
	Contract *EtherTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEtherToken creates a new instance of EtherToken, bound to a specific deployed contract.
func NewEtherToken(address common.Address, backend bind.ContractBackend) (*EtherToken, error) {
	contract, err := bindEtherToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EtherToken{EtherTokenCaller: EtherTokenCaller{contract: contract}, EtherTokenTransactor: EtherTokenTransactor{contract: contract}, EtherTokenFilterer: EtherTokenFilterer{contract: contract}}, nil
}

// NewEtherTokenCaller creates a new read-only instance of EtherToken, bound to a specific deployed contract.
func NewEtherTokenCaller(address common.Address, caller bind.ContractCaller) (*EtherTokenCaller, error) {
	contract, err := bindEtherToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EtherTokenCaller{contract: contract}, nil
}

// NewEtherTokenTransactor creates a new write-only instance of EtherToken, bound to a specific deployed contract.
func NewEtherTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*EtherTokenTransactor, error) {
	contract, err := bindEtherToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EtherTokenTransactor{contract: contract}, nil
}

// NewEtherTokenFilterer creates a new log filterer instance of EtherToken, bound to a specific deployed contract.
func NewEtherTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*EtherTokenFilterer, error) {
	contract, err := bindEtherToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EtherTokenFilterer{contract: contract}, nil
}

// bindEtherToken binds a generic wrapper to an already deployed contract.
func bindEtherToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EtherTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EtherToken *EtherTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EtherToken.Contract.EtherTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EtherToken *EtherTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherToken.Contract.EtherTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EtherToken *EtherTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EtherToken.Contract.EtherTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EtherToken *EtherTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EtherToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EtherToken *EtherTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EtherToken *EtherTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EtherToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256 out)
func (_EtherToken *EtherTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EtherToken.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256 out)
func (_EtherToken *EtherTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _EtherToken.Contract.Allowance(&_EtherToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256 out)
func (_EtherToken *EtherTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _EtherToken.Contract.Allowance(&_EtherToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256 out)
func (_EtherToken *EtherTokenCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EtherToken.contract.Call(opts, out, "balanceOf", owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256 out)
func (_EtherToken *EtherTokenSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _EtherToken.Contract.BalanceOf(&_EtherToken.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256 out)
func (_EtherToken *EtherTokenCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _EtherToken.Contract.BalanceOf(&_EtherToken.CallOpts, owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256 out)
func (_EtherToken *EtherTokenCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EtherToken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256 out)
func (_EtherToken *EtherTokenSession) Decimals() (*big.Int, error) {
	return _EtherToken.Contract.Decimals(&_EtherToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256 out)
func (_EtherToken *EtherTokenCallerSession) Decimals() (*big.Int, error) {
	return _EtherToken.Contract.Decimals(&_EtherToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string out)
func (_EtherToken *EtherTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _EtherToken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string out)
func (_EtherToken *EtherTokenSession) Symbol() (string, error) {
	return _EtherToken.Contract.Symbol(&_EtherToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string out)
func (_EtherToken *EtherTokenCallerSession) Symbol() (string, error) {
	return _EtherToken.Contract.Symbol(&_EtherToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256 out)
func (_EtherToken *EtherTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EtherToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256 out)
func (_EtherToken *EtherTokenSession) TotalSupply() (*big.Int, error) {
	return _EtherToken.Contract.TotalSupply(&_EtherToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256 out)
func (_EtherToken *EtherTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _EtherToken.Contract.TotalSupply(&_EtherToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool out)
func (_EtherToken *EtherTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EtherToken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool out)
func (_EtherToken *EtherTokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EtherToken.Contract.Approve(&_EtherToken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool out)
func (_EtherToken *EtherTokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EtherToken.Contract.Approve(&_EtherToken.TransactOpts, spender, amount)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address spender, uint256 amount) returns()
func (_EtherToken *EtherTokenTransactor) DecreaseApproval(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EtherToken.contract.Transact(opts, "decreaseApproval", spender, amount)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address spender, uint256 amount) returns()
func (_EtherToken *EtherTokenSession) DecreaseApproval(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EtherToken.Contract.DecreaseApproval(&_EtherToken.TransactOpts, spender, amount)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address spender, uint256 amount) returns()
func (_EtherToken *EtherTokenTransactorSession) DecreaseApproval(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EtherToken.Contract.DecreaseApproval(&_EtherToken.TransactOpts, spender, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_EtherToken *EtherTokenTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherToken.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_EtherToken *EtherTokenSession) Deposit() (*types.Transaction, error) {
	return _EtherToken.Contract.Deposit(&_EtherToken.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_EtherToken *EtherTokenTransactorSession) Deposit() (*types.Transaction, error) {
	return _EtherToken.Contract.Deposit(&_EtherToken.TransactOpts)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address spender, uint256 amount) returns()
func (_EtherToken *EtherTokenTransactor) IncreaseApproval(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EtherToken.contract.Transact(opts, "increaseApproval", spender, amount)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address spender, uint256 amount) returns()
func (_EtherToken *EtherTokenSession) IncreaseApproval(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EtherToken.Contract.IncreaseApproval(&_EtherToken.TransactOpts, spender, amount)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address spender, uint256 amount) returns()
func (_EtherToken *EtherTokenTransactorSession) IncreaseApproval(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EtherToken.Contract.IncreaseApproval(&_EtherToken.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool out)
func (_EtherToken *EtherTokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EtherToken.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool out)
func (_EtherToken *EtherTokenSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EtherToken.Contract.Transfer(&_EtherToken.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool out)
func (_EtherToken *EtherTokenTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EtherToken.Contract.Transfer(&_EtherToken.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address source, address to, uint256 amount) returns(bool out)
func (_EtherToken *EtherTokenTransactor) TransferFrom(opts *bind.TransactOpts, source common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EtherToken.contract.Transact(opts, "transferFrom", source, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address source, address to, uint256 amount) returns(bool out)
func (_EtherToken *EtherTokenSession) TransferFrom(source common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EtherToken.Contract.TransferFrom(&_EtherToken.TransactOpts, source, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address source, address to, uint256 amount) returns(bool out)
func (_EtherToken *EtherTokenTransactorSession) TransferFrom(source common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EtherToken.Contract.TransferFrom(&_EtherToken.TransactOpts, source, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_EtherToken *EtherTokenTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _EtherToken.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_EtherToken *EtherTokenSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _EtherToken.Contract.Withdraw(&_EtherToken.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_EtherToken *EtherTokenTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _EtherToken.Contract.Withdraw(&_EtherToken.TransactOpts, amount)
}

// EtherTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the EtherToken contract.
type EtherTokenApprovalIterator struct {
	Event *EtherTokenApproval // Event containing the contract specifics and raw log

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
func (it *EtherTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherTokenApproval)
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
		it.Event = new(EtherTokenApproval)
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
func (it *EtherTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherTokenApproval represents a Approval event raised by the EtherToken contract.
type EtherTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_EtherToken *EtherTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*EtherTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _EtherToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &EtherTokenApprovalIterator{contract: _EtherToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_EtherToken *EtherTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *EtherTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _EtherToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherTokenApproval)
				if err := _EtherToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// EtherTokenDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the EtherToken contract.
type EtherTokenDepositedIterator struct {
	Event *EtherTokenDeposited // Event containing the contract specifics and raw log

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
func (it *EtherTokenDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherTokenDeposited)
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
		it.Event = new(EtherTokenDeposited)
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
func (it *EtherTokenDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherTokenDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherTokenDeposited represents a Deposited event raised by the EtherToken contract.
type EtherTokenDeposited struct {
	Source common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address indexed source, uint256 amount)
func (_EtherToken *EtherTokenFilterer) FilterDeposited(opts *bind.FilterOpts, source []common.Address) (*EtherTokenDepositedIterator, error) {

	var sourceRule []interface{}
	for _, sourceItem := range source {
		sourceRule = append(sourceRule, sourceItem)
	}

	logs, sub, err := _EtherToken.contract.FilterLogs(opts, "Deposited", sourceRule)
	if err != nil {
		return nil, err
	}
	return &EtherTokenDepositedIterator{contract: _EtherToken.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address indexed source, uint256 amount)
func (_EtherToken *EtherTokenFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *EtherTokenDeposited, source []common.Address) (event.Subscription, error) {

	var sourceRule []interface{}
	for _, sourceItem := range source {
		sourceRule = append(sourceRule, sourceItem)
	}

	logs, sub, err := _EtherToken.contract.WatchLogs(opts, "Deposited", sourceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherTokenDeposited)
				if err := _EtherToken.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// EtherTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the EtherToken contract.
type EtherTokenTransferIterator struct {
	Event *EtherTokenTransfer // Event containing the contract specifics and raw log

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
func (it *EtherTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherTokenTransfer)
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
		it.Event = new(EtherTokenTransfer)
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
func (it *EtherTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherTokenTransfer represents a Transfer event raised by the EtherToken contract.
type EtherTokenTransfer struct {
	Source common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed source, address indexed to, uint256 amount)
func (_EtherToken *EtherTokenFilterer) FilterTransfer(opts *bind.FilterOpts, source []common.Address, to []common.Address) (*EtherTokenTransferIterator, error) {

	var sourceRule []interface{}
	for _, sourceItem := range source {
		sourceRule = append(sourceRule, sourceItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EtherToken.contract.FilterLogs(opts, "Transfer", sourceRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EtherTokenTransferIterator{contract: _EtherToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed source, address indexed to, uint256 amount)
func (_EtherToken *EtherTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *EtherTokenTransfer, source []common.Address, to []common.Address) (event.Subscription, error) {

	var sourceRule []interface{}
	for _, sourceItem := range source {
		sourceRule = append(sourceRule, sourceItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EtherToken.contract.WatchLogs(opts, "Transfer", sourceRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherTokenTransfer)
				if err := _EtherToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// EtherTokenWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the EtherToken contract.
type EtherTokenWithdrawnIterator struct {
	Event *EtherTokenWithdrawn // Event containing the contract specifics and raw log

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
func (it *EtherTokenWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherTokenWithdrawn)
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
		it.Event = new(EtherTokenWithdrawn)
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
func (it *EtherTokenWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherTokenWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherTokenWithdrawn represents a Withdrawn event raised by the EtherToken contract.
type EtherTokenWithdrawn struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed to, uint256 amount)
func (_EtherToken *EtherTokenFilterer) FilterWithdrawn(opts *bind.FilterOpts, to []common.Address) (*EtherTokenWithdrawnIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EtherToken.contract.FilterLogs(opts, "Withdrawn", toRule)
	if err != nil {
		return nil, err
	}
	return &EtherTokenWithdrawnIterator{contract: _EtherToken.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed to, uint256 amount)
func (_EtherToken *EtherTokenFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *EtherTokenWithdrawn, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EtherToken.contract.WatchLogs(opts, "Withdrawn", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherTokenWithdrawn)
				if err := _EtherToken.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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
