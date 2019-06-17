// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package markettoken

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

// MarketTokenABI is the input ABI used to generate the binding from.
const MarketTokenABI = "[{\"name\":\"Approval\",\"inputs\":[{\"type\":\"address\",\"name\":\"owner\",\"indexed\":true},{\"type\":\"address\",\"name\":\"spender\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"amount\",\"indexed\":false,\"unit\":\"wei\"}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Burned\",\"inputs\":[{\"type\":\"address\",\"name\":\"burner\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"amount\",\"indexed\":false,\"unit\":\"wei\"}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Minted\",\"inputs\":[{\"type\":\"address\",\"name\":\"to\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"amount\",\"indexed\":false,\"unit\":\"wei\"}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Transfer\",\"inputs\":[{\"type\":\"address\",\"name\":\"source\",\"indexed\":true},{\"type\":\"address\",\"name\":\"to\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"amount\",\"indexed\":false,\"unit\":\"wei\"}],\"anonymous\":false,\"type\":\"event\"},{\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"initial_account\"},{\"type\":\"uint256\",\"name\":\"initial_balance\",\"unit\":\"wei\"}],\"constant\":false,\"payable\":false,\"type\":\"constructor\"},{\"name\":\"allowance\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\",\"unit\":\"wei\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"owner\"},{\"type\":\"address\",\"name\":\"spender\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":815},{\"name\":\"approve\",\"outputs\":[{\"type\":\"bool\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"spender\"},{\"type\":\"uint256\",\"name\":\"amount\",\"unit\":\"wei\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":37719},{\"name\":\"balanceOf\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\",\"unit\":\"wei\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"owner\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":715},{\"name\":\"getPrivileged\",\"outputs\":[{\"type\":\"address\",\"name\":\"out\"},{\"type\":\"address\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":851},{\"name\":\"hasPrivilege\",\"outputs\":[{\"type\":\"bool\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"sender\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":884},{\"name\":\"burn\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256\",\"name\":\"amount\",\"unit\":\"wei\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":75309},{\"name\":\"burnAll\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"owner\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":80387},{\"name\":\"decreaseApproval\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"spender\"},{\"type\":\"uint256\",\"name\":\"amount\",\"unit\":\"wei\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":38862},{\"name\":\"increaseApproval\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"spender\"},{\"type\":\"uint256\",\"name\":\"amount\",\"unit\":\"wei\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":39098},{\"name\":\"mint\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256\",\"name\":\"amount\",\"unit\":\"wei\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":75841},{\"name\":\"setPrivileged\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"reserve\"},{\"type\":\"address\",\"name\":\"listing\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":71563},{\"name\":\"totalSupply\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\",\"unit\":\"wei\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":813},{\"name\":\"transfer\",\"outputs\":[{\"type\":\"bool\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"to\"},{\"type\":\"uint256\",\"name\":\"amount\",\"unit\":\"wei\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":74471},{\"name\":\"transferFrom\",\"outputs\":[{\"type\":\"bool\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"source\"},{\"type\":\"address\",\"name\":\"to\"},{\"type\":\"uint256\",\"name\":\"amount\",\"unit\":\"wei\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":110441},{\"name\":\"decimals\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":903},{\"name\":\"symbol\",\"outputs\":[{\"type\":\"string\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2411}]"

// MarketTokenBin is the compiled bytecode used for deploying new contracts.
const MarketTokenBin = `0x740100000000000000000000000000000000000000006020526f7fffffffffffffffffffffffffffffff6040527fffffffffffffffffffffffffffffffff8000000000000000000000000000000060605274012a05f1fffffffffffffffffffffffffdabf41c006080527ffffffffffffffffffffffffed5fa0e000000000000000000000000000000000060a0526040610bd36101403934156100a157600080fd5b6020610bd360c03960c05160205181106100ba57600080fd5b50336004556101605160016101405160e05260c052604060c0205560126002556003610180527f434d5400000000000000000000000000000000000000000000000000000000006101a05261018080600360c052602060c020602082510161012060006002818352015b8261012051602002111561013757610159565b61012051602002850151610120518501555b8151600101808352811415610124575b50505050505061016051600755610bbb56600035601c52740100000000000000000000000000000000000000006020526f7fffffffffffffffffffffffffffffff6040527fffffffffffffffffffffffffffffffff8000000000000000000000000000000060605274012a05f1fffffffffffffffffffffffffdabf41c006080527ffffffffffffffffffffffffed5fa0e000000000000000000000000000000000060a05263dd62ed3e600051141561010357604060046101403734156100b457600080fd5b60043560205181106100c557600080fd5b5060243560205181106100d757600080fd5b5060006101405160e05260c052604060c0206101605160e05260c052604060c0205460005260206000f3005b63095ea7b36000511415610198576040600461014037341561012457600080fd5b600435602051811061013557600080fd5b506101605160003360e05260c052604060c0206101405160e05260c052604060c02055610160516101805261014051337f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9256020610180a3600160005260206000f3005b6370a0823160005114156101e757602060046101403734156101b957600080fd5b60043560205181106101ca57600080fd5b5060016101405160e05260c052604060c0205460005260206000f3005b63d4c17539600051141561022157341561020057600080fd5b604061014052610160600554815260065481602001525061014051610160f3005b63b7bf210c600051141561026f576020600461014037341561024257600080fd5b600435602051811061025357600080fd5b50600654610140511460055461014051141760005260206000f3005b6342966c68600051141561033b576020600461014037341561029057600080fd5b60206101e0602463b7bf210c61016052336101805261017c6000305af16102b657600080fd5b6101e0516102c357600080fd5b60013360e05260c052604060c02061014051815410156102e257600080fd5b61014051815403815550600761014051815410156102ff57600080fd5b610140518154038155506101405161020052337f696de425f79f4a40bc6d2122ca50507f0efbeabbff86a84871b7196ab8ea8df76020610200a2005b637e9d2ac1600051141561043d576020600461014037341561035c57600080fd5b600435602051811061036d57600080fd5b5060206101e0602463b7bf210c61016052336101805261017c6000305af161039457600080fd5b6101e0516103a157600080fd5b60016101405160e05260c052604060c0205461020052600761020051815410156103ca57600080fd5b61020051815403815550600060016101405160e05260c052604060c02055600060006101405160e05260c052604060c0203360e05260c052604060c020556102005161022052610140517f696de425f79f4a40bc6d2122ca50507f0efbeabbff86a84871b7196ab8ea8df76020610220a2005b636618846360005114156104f8576040600461014037341561045e57600080fd5b600435602051811061046f57600080fd5b5060003360e05260c052604060c0206101405160e05260c052604060c020610160518154101561049e57600080fd5b6101605181540381555060003360e05260c052604060c0206101405160e05260c052604060c020546101805261014051337f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9256020610180a3005b63d73dd62360005114156105b6576040600461014037341561051957600080fd5b600435602051811061052a57600080fd5b5060003360e05260c052604060c0206101405160e05260c052604060c020805461016051825401101561055c57600080fd5b6101605181540181555060003360e05260c052604060c0206101405160e05260c052604060c020546101805261014051337f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9256020610180a3005b63a0712d68600051141561068857602060046101403734156105d757600080fd5b60206101e0602463b7bf210c61016052336101805261017c6000305af16105fd57600080fd5b6101e05161060a57600080fd5b6007805461014051825401101561062057600080fd5b6101405181540181555060013360e05260c052604060c020805461014051825401101561064c57600080fd5b610140518154018155506101405161020052337f30385c845b448a36257a6a1716e6ad2e1bc2cbe333cde1e69fe849ad6511adfe6020610200a2005b63b01e369f600051141561070557604060046101403734156106a957600080fd5b60043560205181106106ba57600080fd5b5060243560205181106106cc57600080fd5b5060045433146106db57600080fd5b600654156106e857600080fd5b600554156106f557600080fd5b6101405160055561016051600655005b6318160ddd600051141561072b57341561071e57600080fd5b60075460005260206000f3005b63a9059cbb6000511415610806576040600461014037341561074c57600080fd5b600435602051811061075d57600080fd5b506000610140511861076e57600080fd5b60013360e05260c052604060c020610160518154101561078d57600080fd5b6101605181540381555060016101405160e05260c052604060c02080546101605182540110156107bc57600080fd5b61016051815401815550610160516101805261014051337fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef6020610180a3600160005260206000f3005b6323b872dd6000511415610941576060600461014037341561082757600080fd5b600435602051811061083857600080fd5b50602435602051811061084a57600080fd5b506000610140511861085b57600080fd5b6000610160511861086b57600080fd5b60016101405160e05260c052604060c020610180518154101561088d57600080fd5b6101805181540381555060016101605160e05260c052604060c02080546101805182540110156108bc57600080fd5b6101805181540181555060006101405160e05260c052604060c0203360e05260c052604060c02061018051815410156108f457600080fd5b61018051815403815550610180516101a05261016051610140517fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60206101a0a3600160005260206000f3005b63313ce567600051141561096757341561095a57600080fd5b60025460005260206000f3005b6395d89b416000511415610a4a57341561098057600080fd5b60038060c052602060c020610180602082540161012060006002818352015b826101205160200211156109b2576109d4565b61012051850154610120516020028501525b815160010180835281141561099f575b5050505050506101805160206001820306601f82010390506101e0610180516003818352015b826101e0511115610a0a57610a26565b60006101e0516101a001535b81516001018083528114156109fa575b5050506020610160526040610180510160206001820306601f8201039050610160f3005b60006000fd5b61016b610bbb0361016b60003961016b610bbb036000f3`

// DeployMarketToken deploys a new Ethereum contract, binding an instance of MarketToken to it.
func DeployMarketToken(auth *bind.TransactOpts, backend bind.ContractBackend, initial_account common.Address, initial_balance *big.Int) (common.Address, *types.Transaction, *MarketToken, error) {
	parsed, err := abi.JSON(strings.NewReader(MarketTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MarketTokenBin), backend, initial_account, initial_balance)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MarketToken{MarketTokenCaller: MarketTokenCaller{contract: contract}, MarketTokenTransactor: MarketTokenTransactor{contract: contract}, MarketTokenFilterer: MarketTokenFilterer{contract: contract}}, nil
}

// MarketToken is an auto generated Go binding around an Ethereum contract.
type MarketToken struct {
	MarketTokenCaller     // Read-only binding to the contract
	MarketTokenTransactor // Write-only binding to the contract
	MarketTokenFilterer   // Log filterer for contract events
}

// MarketTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type MarketTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MarketTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MarketTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MarketTokenSession struct {
	Contract     *MarketToken      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MarketTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MarketTokenCallerSession struct {
	Contract *MarketTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// MarketTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MarketTokenTransactorSession struct {
	Contract     *MarketTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// MarketTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type MarketTokenRaw struct {
	Contract *MarketToken // Generic contract binding to access the raw methods on
}

// MarketTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MarketTokenCallerRaw struct {
	Contract *MarketTokenCaller // Generic read-only contract binding to access the raw methods on
}

// MarketTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MarketTokenTransactorRaw struct {
	Contract *MarketTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMarketToken creates a new instance of MarketToken, bound to a specific deployed contract.
func NewMarketToken(address common.Address, backend bind.ContractBackend) (*MarketToken, error) {
	contract, err := bindMarketToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MarketToken{MarketTokenCaller: MarketTokenCaller{contract: contract}, MarketTokenTransactor: MarketTokenTransactor{contract: contract}, MarketTokenFilterer: MarketTokenFilterer{contract: contract}}, nil
}

// NewMarketTokenCaller creates a new read-only instance of MarketToken, bound to a specific deployed contract.
func NewMarketTokenCaller(address common.Address, caller bind.ContractCaller) (*MarketTokenCaller, error) {
	contract, err := bindMarketToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MarketTokenCaller{contract: contract}, nil
}

// NewMarketTokenTransactor creates a new write-only instance of MarketToken, bound to a specific deployed contract.
func NewMarketTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*MarketTokenTransactor, error) {
	contract, err := bindMarketToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MarketTokenTransactor{contract: contract}, nil
}

// NewMarketTokenFilterer creates a new log filterer instance of MarketToken, bound to a specific deployed contract.
func NewMarketTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*MarketTokenFilterer, error) {
	contract, err := bindMarketToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MarketTokenFilterer{contract: contract}, nil
}

// bindMarketToken binds a generic wrapper to an already deployed contract.
func bindMarketToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MarketTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MarketToken *MarketTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MarketToken.Contract.MarketTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MarketToken *MarketTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MarketToken.Contract.MarketTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MarketToken *MarketTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MarketToken.Contract.MarketTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MarketToken *MarketTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MarketToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MarketToken *MarketTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MarketToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MarketToken *MarketTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MarketToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(out uint256)
func (_MarketToken *MarketTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MarketToken.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(out uint256)
func (_MarketToken *MarketTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _MarketToken.Contract.Allowance(&_MarketToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(out uint256)
func (_MarketToken *MarketTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _MarketToken.Contract.Allowance(&_MarketToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(owner address) constant returns(out uint256)
func (_MarketToken *MarketTokenCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MarketToken.contract.Call(opts, out, "balanceOf", owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(owner address) constant returns(out uint256)
func (_MarketToken *MarketTokenSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _MarketToken.Contract.BalanceOf(&_MarketToken.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(owner address) constant returns(out uint256)
func (_MarketToken *MarketTokenCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _MarketToken.Contract.BalanceOf(&_MarketToken.CallOpts, owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(out uint256)
func (_MarketToken *MarketTokenCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MarketToken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(out uint256)
func (_MarketToken *MarketTokenSession) Decimals() (*big.Int, error) {
	return _MarketToken.Contract.Decimals(&_MarketToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(out uint256)
func (_MarketToken *MarketTokenCallerSession) Decimals() (*big.Int, error) {
	return _MarketToken.Contract.Decimals(&_MarketToken.CallOpts)
}

// GetPrivileged is a free data retrieval call binding the contract method 0xd4c17539.
//
// Solidity: function getPrivileged() constant returns(out address, out address)
func (_MarketToken *MarketTokenCaller) GetPrivileged(opts *bind.CallOpts) (common.Address, common.Address, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(common.Address)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _MarketToken.contract.Call(opts, out, "getPrivileged")
	return *ret0, *ret1, err
}

// GetPrivileged is a free data retrieval call binding the contract method 0xd4c17539.
//
// Solidity: function getPrivileged() constant returns(out address, out address)
func (_MarketToken *MarketTokenSession) GetPrivileged() (common.Address, common.Address, error) {
	return _MarketToken.Contract.GetPrivileged(&_MarketToken.CallOpts)
}

// GetPrivileged is a free data retrieval call binding the contract method 0xd4c17539.
//
// Solidity: function getPrivileged() constant returns(out address, out address)
func (_MarketToken *MarketTokenCallerSession) GetPrivileged() (common.Address, common.Address, error) {
	return _MarketToken.Contract.GetPrivileged(&_MarketToken.CallOpts)
}

// HasPrivilege is a free data retrieval call binding the contract method 0xb7bf210c.
//
// Solidity: function hasPrivilege(sender address) constant returns(out bool)
func (_MarketToken *MarketTokenCaller) HasPrivilege(opts *bind.CallOpts, sender common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MarketToken.contract.Call(opts, out, "hasPrivilege", sender)
	return *ret0, err
}

// HasPrivilege is a free data retrieval call binding the contract method 0xb7bf210c.
//
// Solidity: function hasPrivilege(sender address) constant returns(out bool)
func (_MarketToken *MarketTokenSession) HasPrivilege(sender common.Address) (bool, error) {
	return _MarketToken.Contract.HasPrivilege(&_MarketToken.CallOpts, sender)
}

// HasPrivilege is a free data retrieval call binding the contract method 0xb7bf210c.
//
// Solidity: function hasPrivilege(sender address) constant returns(out bool)
func (_MarketToken *MarketTokenCallerSession) HasPrivilege(sender common.Address) (bool, error) {
	return _MarketToken.Contract.HasPrivilege(&_MarketToken.CallOpts, sender)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(out string)
func (_MarketToken *MarketTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _MarketToken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(out string)
func (_MarketToken *MarketTokenSession) Symbol() (string, error) {
	return _MarketToken.Contract.Symbol(&_MarketToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(out string)
func (_MarketToken *MarketTokenCallerSession) Symbol() (string, error) {
	return _MarketToken.Contract.Symbol(&_MarketToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(out uint256)
func (_MarketToken *MarketTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MarketToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(out uint256)
func (_MarketToken *MarketTokenSession) TotalSupply() (*big.Int, error) {
	return _MarketToken.Contract.TotalSupply(&_MarketToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(out uint256)
func (_MarketToken *MarketTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _MarketToken.Contract.TotalSupply(&_MarketToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, amount uint256) returns(out bool)
func (_MarketToken *MarketTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MarketToken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, amount uint256) returns(out bool)
func (_MarketToken *MarketTokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MarketToken.Contract.Approve(&_MarketToken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, amount uint256) returns(out bool)
func (_MarketToken *MarketTokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MarketToken.Contract.Approve(&_MarketToken.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(amount uint256) returns()
func (_MarketToken *MarketTokenTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _MarketToken.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(amount uint256) returns()
func (_MarketToken *MarketTokenSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _MarketToken.Contract.Burn(&_MarketToken.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(amount uint256) returns()
func (_MarketToken *MarketTokenTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _MarketToken.Contract.Burn(&_MarketToken.TransactOpts, amount)
}

// BurnAll is a paid mutator transaction binding the contract method 0x7e9d2ac1.
//
// Solidity: function burnAll(owner address) returns()
func (_MarketToken *MarketTokenTransactor) BurnAll(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _MarketToken.contract.Transact(opts, "burnAll", owner)
}

// BurnAll is a paid mutator transaction binding the contract method 0x7e9d2ac1.
//
// Solidity: function burnAll(owner address) returns()
func (_MarketToken *MarketTokenSession) BurnAll(owner common.Address) (*types.Transaction, error) {
	return _MarketToken.Contract.BurnAll(&_MarketToken.TransactOpts, owner)
}

// BurnAll is a paid mutator transaction binding the contract method 0x7e9d2ac1.
//
// Solidity: function burnAll(owner address) returns()
func (_MarketToken *MarketTokenTransactorSession) BurnAll(owner common.Address) (*types.Transaction, error) {
	return _MarketToken.Contract.BurnAll(&_MarketToken.TransactOpts, owner)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(spender address, amount uint256) returns()
func (_MarketToken *MarketTokenTransactor) DecreaseApproval(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MarketToken.contract.Transact(opts, "decreaseApproval", spender, amount)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(spender address, amount uint256) returns()
func (_MarketToken *MarketTokenSession) DecreaseApproval(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MarketToken.Contract.DecreaseApproval(&_MarketToken.TransactOpts, spender, amount)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(spender address, amount uint256) returns()
func (_MarketToken *MarketTokenTransactorSession) DecreaseApproval(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MarketToken.Contract.DecreaseApproval(&_MarketToken.TransactOpts, spender, amount)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(spender address, amount uint256) returns()
func (_MarketToken *MarketTokenTransactor) IncreaseApproval(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MarketToken.contract.Transact(opts, "increaseApproval", spender, amount)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(spender address, amount uint256) returns()
func (_MarketToken *MarketTokenSession) IncreaseApproval(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MarketToken.Contract.IncreaseApproval(&_MarketToken.TransactOpts, spender, amount)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(spender address, amount uint256) returns()
func (_MarketToken *MarketTokenTransactorSession) IncreaseApproval(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MarketToken.Contract.IncreaseApproval(&_MarketToken.TransactOpts, spender, amount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(amount uint256) returns()
func (_MarketToken *MarketTokenTransactor) Mint(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _MarketToken.contract.Transact(opts, "mint", amount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(amount uint256) returns()
func (_MarketToken *MarketTokenSession) Mint(amount *big.Int) (*types.Transaction, error) {
	return _MarketToken.Contract.Mint(&_MarketToken.TransactOpts, amount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(amount uint256) returns()
func (_MarketToken *MarketTokenTransactorSession) Mint(amount *big.Int) (*types.Transaction, error) {
	return _MarketToken.Contract.Mint(&_MarketToken.TransactOpts, amount)
}

// SetPrivileged is a paid mutator transaction binding the contract method 0xb01e369f.
//
// Solidity: function setPrivileged(reserve address, listing address) returns()
func (_MarketToken *MarketTokenTransactor) SetPrivileged(opts *bind.TransactOpts, reserve common.Address, listing common.Address) (*types.Transaction, error) {
	return _MarketToken.contract.Transact(opts, "setPrivileged", reserve, listing)
}

// SetPrivileged is a paid mutator transaction binding the contract method 0xb01e369f.
//
// Solidity: function setPrivileged(reserve address, listing address) returns()
func (_MarketToken *MarketTokenSession) SetPrivileged(reserve common.Address, listing common.Address) (*types.Transaction, error) {
	return _MarketToken.Contract.SetPrivileged(&_MarketToken.TransactOpts, reserve, listing)
}

// SetPrivileged is a paid mutator transaction binding the contract method 0xb01e369f.
//
// Solidity: function setPrivileged(reserve address, listing address) returns()
func (_MarketToken *MarketTokenTransactorSession) SetPrivileged(reserve common.Address, listing common.Address) (*types.Transaction, error) {
	return _MarketToken.Contract.SetPrivileged(&_MarketToken.TransactOpts, reserve, listing)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, amount uint256) returns(out bool)
func (_MarketToken *MarketTokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MarketToken.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, amount uint256) returns(out bool)
func (_MarketToken *MarketTokenSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MarketToken.Contract.Transfer(&_MarketToken.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, amount uint256) returns(out bool)
func (_MarketToken *MarketTokenTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MarketToken.Contract.Transfer(&_MarketToken.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(source address, to address, amount uint256) returns(out bool)
func (_MarketToken *MarketTokenTransactor) TransferFrom(opts *bind.TransactOpts, source common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MarketToken.contract.Transact(opts, "transferFrom", source, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(source address, to address, amount uint256) returns(out bool)
func (_MarketToken *MarketTokenSession) TransferFrom(source common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MarketToken.Contract.TransferFrom(&_MarketToken.TransactOpts, source, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(source address, to address, amount uint256) returns(out bool)
func (_MarketToken *MarketTokenTransactorSession) TransferFrom(source common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MarketToken.Contract.TransferFrom(&_MarketToken.TransactOpts, source, to, amount)
}

// MarketTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the MarketToken contract.
type MarketTokenApprovalIterator struct {
	Event *MarketTokenApproval // Event containing the contract specifics and raw log

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
func (it *MarketTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketTokenApproval)
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
		it.Event = new(MarketTokenApproval)
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
func (it *MarketTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketTokenApproval represents a Approval event raised by the MarketToken contract.
type MarketTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, amount uint256)
func (_MarketToken *MarketTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*MarketTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MarketToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &MarketTokenApprovalIterator{contract: _MarketToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, amount uint256)
func (_MarketToken *MarketTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MarketTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MarketToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketTokenApproval)
				if err := _MarketToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// MarketTokenBurnedIterator is returned from FilterBurned and is used to iterate over the raw logs and unpacked data for Burned events raised by the MarketToken contract.
type MarketTokenBurnedIterator struct {
	Event *MarketTokenBurned // Event containing the contract specifics and raw log

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
func (it *MarketTokenBurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketTokenBurned)
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
		it.Event = new(MarketTokenBurned)
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
func (it *MarketTokenBurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketTokenBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketTokenBurned represents a Burned event raised by the MarketToken contract.
type MarketTokenBurned struct {
	Burner common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBurned is a free log retrieval operation binding the contract event 0x696de425f79f4a40bc6d2122ca50507f0efbeabbff86a84871b7196ab8ea8df7.
//
// Solidity: e Burned(burner indexed address, amount uint256)
func (_MarketToken *MarketTokenFilterer) FilterBurned(opts *bind.FilterOpts, burner []common.Address) (*MarketTokenBurnedIterator, error) {

	var burnerRule []interface{}
	for _, burnerItem := range burner {
		burnerRule = append(burnerRule, burnerItem)
	}

	logs, sub, err := _MarketToken.contract.FilterLogs(opts, "Burned", burnerRule)
	if err != nil {
		return nil, err
	}
	return &MarketTokenBurnedIterator{contract: _MarketToken.contract, event: "Burned", logs: logs, sub: sub}, nil
}

// WatchBurned is a free log subscription operation binding the contract event 0x696de425f79f4a40bc6d2122ca50507f0efbeabbff86a84871b7196ab8ea8df7.
//
// Solidity: e Burned(burner indexed address, amount uint256)
func (_MarketToken *MarketTokenFilterer) WatchBurned(opts *bind.WatchOpts, sink chan<- *MarketTokenBurned, burner []common.Address) (event.Subscription, error) {

	var burnerRule []interface{}
	for _, burnerItem := range burner {
		burnerRule = append(burnerRule, burnerItem)
	}

	logs, sub, err := _MarketToken.contract.WatchLogs(opts, "Burned", burnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketTokenBurned)
				if err := _MarketToken.contract.UnpackLog(event, "Burned", log); err != nil {
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

// MarketTokenMintedIterator is returned from FilterMinted and is used to iterate over the raw logs and unpacked data for Minted events raised by the MarketToken contract.
type MarketTokenMintedIterator struct {
	Event *MarketTokenMinted // Event containing the contract specifics and raw log

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
func (it *MarketTokenMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketTokenMinted)
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
		it.Event = new(MarketTokenMinted)
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
func (it *MarketTokenMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketTokenMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketTokenMinted represents a Minted event raised by the MarketToken contract.
type MarketTokenMinted struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMinted is a free log retrieval operation binding the contract event 0x30385c845b448a36257a6a1716e6ad2e1bc2cbe333cde1e69fe849ad6511adfe.
//
// Solidity: e Minted(to indexed address, amount uint256)
func (_MarketToken *MarketTokenFilterer) FilterMinted(opts *bind.FilterOpts, to []common.Address) (*MarketTokenMintedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MarketToken.contract.FilterLogs(opts, "Minted", toRule)
	if err != nil {
		return nil, err
	}
	return &MarketTokenMintedIterator{contract: _MarketToken.contract, event: "Minted", logs: logs, sub: sub}, nil
}

// WatchMinted is a free log subscription operation binding the contract event 0x30385c845b448a36257a6a1716e6ad2e1bc2cbe333cde1e69fe849ad6511adfe.
//
// Solidity: e Minted(to indexed address, amount uint256)
func (_MarketToken *MarketTokenFilterer) WatchMinted(opts *bind.WatchOpts, sink chan<- *MarketTokenMinted, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MarketToken.contract.WatchLogs(opts, "Minted", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketTokenMinted)
				if err := _MarketToken.contract.UnpackLog(event, "Minted", log); err != nil {
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

// MarketTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the MarketToken contract.
type MarketTokenTransferIterator struct {
	Event *MarketTokenTransfer // Event containing the contract specifics and raw log

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
func (it *MarketTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketTokenTransfer)
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
		it.Event = new(MarketTokenTransfer)
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
func (it *MarketTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketTokenTransfer represents a Transfer event raised by the MarketToken contract.
type MarketTokenTransfer struct {
	Source common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(source indexed address, to indexed address, amount uint256)
func (_MarketToken *MarketTokenFilterer) FilterTransfer(opts *bind.FilterOpts, source []common.Address, to []common.Address) (*MarketTokenTransferIterator, error) {

	var sourceRule []interface{}
	for _, sourceItem := range source {
		sourceRule = append(sourceRule, sourceItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MarketToken.contract.FilterLogs(opts, "Transfer", sourceRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MarketTokenTransferIterator{contract: _MarketToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(source indexed address, to indexed address, amount uint256)
func (_MarketToken *MarketTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MarketTokenTransfer, source []common.Address, to []common.Address) (event.Subscription, error) {

	var sourceRule []interface{}
	for _, sourceItem := range source {
		sourceRule = append(sourceRule, sourceItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MarketToken.contract.WatchLogs(opts, "Transfer", sourceRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketTokenTransfer)
				if err := _MarketToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
