// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package parameterizer

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

// ParameterizerABI is the input ABI used to generate the binding from.
const ParameterizerABI = "[{\"name\":\"ReparamProposed\",\"inputs\":[{\"type\":\"address\",\"name\":\"owner\",\"indexed\":true},{\"type\":\"bytes32\",\"name\":\"hash\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"param\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"value\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"ReparamFailed\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"param\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"value\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"ReparamSucceeded\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"param\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"value\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"mkt_addr\"},{\"type\":\"address\",\"name\":\"v_addr\"},{\"type\":\"uint256\",\"unit\":\"wei\",\"name\":\"pr_fl\"},{\"type\":\"uint256\",\"name\":\"spd\"},{\"type\":\"uint256\",\"unit\":\"wei\",\"name\":\"list_re\"},{\"type\":\"uint256\",\"unit\":\"wei\",\"name\":\"stk\"},{\"type\":\"uint256\",\"unit\":\"sec\",\"name\":\"vote_by_d\"},{\"type\":\"uint256\",\"name\":\"pl\"},{\"type\":\"uint256\",\"name\":\"back_p\"},{\"type\":\"uint256\",\"name\":\"maker_p\"},{\"type\":\"uint256\",\"unit\":\"wei\",\"name\":\"cost\"}],\"constant\":false,\"payable\":false,\"type\":\"constructor\"},{\"name\":\"getBackendPayment\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":551},{\"name\":\"getMakerPayment\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":581},{\"name\":\"getReservePayment\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2834},{\"name\":\"getCostPerByte\",\"outputs\":[{\"type\":\"uint256\",\"unit\":\"wei\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":641},{\"name\":\"getStake\",\"outputs\":[{\"type\":\"uint256\",\"unit\":\"wei\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":671},{\"name\":\"getPriceFloor\",\"outputs\":[{\"type\":\"uint256\",\"unit\":\"wei\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":701},{\"name\":\"getHash\",\"outputs\":[{\"type\":\"bytes32\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"param\"},{\"type\":\"uint256\",\"name\":\"value\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":766},{\"name\":\"getSpread\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":761},{\"name\":\"getListReward\",\"outputs\":[{\"type\":\"uint256\",\"unit\":\"wei\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":791},{\"name\":\"getPlurality\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":821},{\"name\":\"getReparam\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"},{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":1509},{\"name\":\"getVoteBy\",\"outputs\":[{\"type\":\"uint256\",\"unit\":\"sec\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":881},{\"name\":\"reparameterize\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256\",\"name\":\"param\"},{\"type\":\"uint256\",\"name\":\"value\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":78273},{\"name\":\"resolveReparam\",\"outputs\":[],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":89087}]"

// ParameterizerBin is the compiled bytecode used for deploying new contracts.
const ParameterizerBin = `0x740100000000000000000000000000000000000000006020526f7fffffffffffffffffffffffffffffff6040527fffffffffffffffffffffffffffffffff8000000000000000000000000000000060605274012a05f1fffffffffffffffffffffffffdabf41c006080527ffffffffffffffffffffffffed5fa0e000000000000000000000000000000000060a052610160610a0d6101403934156100a257600080fd5b6020610a0d60c03960c05160205181106100bb57600080fd5b5060206020610a0d0160c03960c05160205181106100d857600080fd5b5061014051600a5561016051600b55610180516001556101a0516002556101c0516003556101e05160045561020051600555610220516006556102405160075561026051600855610280516009556109f556600436101561000d576108c4565b600035601c52740100000000000000000000000000000000000000006020526f7fffffffffffffffffffffffffffffff6040527fffffffffffffffffffffffffffffffff8000000000000000000000000000000060605274012a05f1fffffffffffffffffffffffffdabf41c006080527ffffffffffffffffffffffffed5fa0e000000000000000000000000000000000060a052634b8b7dce60005114156100c85734156100ba57600080fd5b60075460005260206000f350005b63ccc3412f60005114156100ef5734156100e157600080fd5b60085460005260206000f350005b63ffe44da8600051141561015b57341561010857600080fd5b60075460085460075401101561011d57600080fd5b600854600754016064101561013157600080fd5b60075460085460075401101561014657600080fd5b6008546007540160640360005260206000f350005b63cd2d5438600051141561018257341561017457600080fd5b60095460005260206000f350005b63fc0e3d9060005114156101a957341561019b57600080fd5b60045460005260206000f350005b6303bca7d360005114156101d05734156101c257600080fd5b60015460005260206000f350005b631b855044600051141561022d5734156101e957600080fd5b60006004356020826101400101526020810190506024356020826101400101526020810190508061014052610140905080516020820120905060005260206000f350005b637b3cee60600051141561025457341561024657600080fd5b60025460005260206000f350005b63dc2b2853600051141561027b57341561026d57600080fd5b60035460005260206000f350005b633d1e37d560005114156102a257341561029457600080fd5b60065460005260206000f350005b63d90f8f0d600051141561030c5734156102bb57600080fd5b604061014052610160600060043560e05260c052604060c02060c052602060c0205481526001600060043560e05260c052604060c02060c052602060c020015481602001525061014051610160f350005b632d0d2bc6600051141561033357341561032557600080fd5b60055460005260206000f350005b63c1836218600051141561049b57341561034c57600080fd5b60006004356020826101600101526020810190506024356020826101600101526020810190508061016052610160905080516020820120905061014052600b543b61039657600080fd5b600b5430186103a457600080fd5b6020610260602463b89694c66101e05261014051610200526101fc600b545afa6103cd57600080fd5b60005061026051156103de57600080fd5b60006101405160e05260c052604060c02060c052602060c0206004358155602435600182015550600b543b61041257600080fd5b600b54301861042057600080fd5b6000600060a463bb12c49e61028052610140516102a05260036102c052336102e052600454610300526005546103205261029c6000600b545af161046357600080fd5b6024356103805260043561014051337fb38eb1d14189ba2de4ca92c3ead8d4c2a61d566d266d46b487c4fa82ae7328426020610380a4005b63435c709a60005114156108c35734156104b457600080fd5b600b543b6104c157600080fd5b600b5430186104cf57600080fd5b60206101e0604463af61f760610140526004356101605260036101805261015c600b545afa6104fd57600080fd5b6000506101e05161050d57600080fd5b600b543b61051a57600080fd5b600b54301861052857600080fd5b6020610280602463327322c8610200526004356102205261021c600b545afa61055057600080fd5b6000506102805161056057600080fd5b600060043560e05260c052604060c02060c052602060c020546102a0526001600060043560e05260c052604060c02060c052602060c02001546102c052600b543b6105aa57600080fd5b600b5430186105b857600080fd5b60206103a06044638f354b7961030052600435610320526006546103405261031c600b545afa6105e757600080fd5b6000506103a051156108225760026102a051141561060b576102c0516001556107e7565b60046102a05114156106345760646102c051101561062857600080fd5b6102c0516002556107e6565b60056102a051141561064c576102c0516003556107e5565b60016102a05114156106dc5760006102c0511161066857600080fd5b600a543b61067557600080fd5b600a54301861068357600080fd5b602061042060046318160ddd6103c0526103dc600a545afa6106a457600080fd5b600050610420516104405260036106ba57600080fd5b600361044051046102c05111156106d057600080fd5b6102c0516004556107e4565b60076102a051141561071a57620151806102c05110156106fb57600080fd5b621275006102c051111561070e57600080fd5b6102c0516005556107e3565b60066102a05114156107435760646102c051111561073757600080fd5b6102c0516006556107e2565b60096102a05114156107875760646102c0516007546102c05101101561076857600080fd5b6007546102c05101111561077b57600080fd5b6102c0516008556107e1565b60086102a05114156107cb5760646102c0516008546102c0510110156107ac57600080fd5b6008546102c0510111156107bf57600080fd5b6102c0516007556107e0565b600b6102a05114156107df576102c0516009555b5b5b5b5b5b5b5b5b6102c051610460526102a0516004357fe6f29ce1d705c668b223a35e77096d95bbd994c121503fba4b29e65be92bc7e86020610460a3610859565b6102c0516102e0526102a0516004357f1f050cdbc74210070c4a499a73459732a158933d789331de2757eb4d3563025360206102e0a35b600b543b61086657600080fd5b600b54301861087457600080fd5b6000600060246389bb617c610480526004356104a05261049c6000600b545af161089d57600080fd5b600060043560e05260c052604060c02060c052602060c020600081556000600182015550005b5b60006000fd5b61012b6109f50361012b60003961012b6109f5036000f3`

// DeployParameterizer deploys a new Ethereum contract, binding an instance of Parameterizer to it.
func DeployParameterizer(auth *bind.TransactOpts, backend bind.ContractBackend, mkt_addr common.Address, v_addr common.Address, pr_fl *big.Int, spd *big.Int, list_re *big.Int, stk *big.Int, vote_by_d *big.Int, pl *big.Int, back_p *big.Int, maker_p *big.Int, cost *big.Int) (common.Address, *types.Transaction, *Parameterizer, error) {
	parsed, err := abi.JSON(strings.NewReader(ParameterizerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ParameterizerBin), backend, mkt_addr, v_addr, pr_fl, spd, list_re, stk, vote_by_d, pl, back_p, maker_p, cost)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Parameterizer{ParameterizerCaller: ParameterizerCaller{contract: contract}, ParameterizerTransactor: ParameterizerTransactor{contract: contract}, ParameterizerFilterer: ParameterizerFilterer{contract: contract}}, nil
}

// Parameterizer is an auto generated Go binding around an Ethereum contract.
type Parameterizer struct {
	ParameterizerCaller     // Read-only binding to the contract
	ParameterizerTransactor // Write-only binding to the contract
	ParameterizerFilterer   // Log filterer for contract events
}

// ParameterizerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ParameterizerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ParameterizerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ParameterizerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ParameterizerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ParameterizerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ParameterizerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ParameterizerSession struct {
	Contract     *Parameterizer    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ParameterizerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ParameterizerCallerSession struct {
	Contract *ParameterizerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ParameterizerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ParameterizerTransactorSession struct {
	Contract     *ParameterizerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ParameterizerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ParameterizerRaw struct {
	Contract *Parameterizer // Generic contract binding to access the raw methods on
}

// ParameterizerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ParameterizerCallerRaw struct {
	Contract *ParameterizerCaller // Generic read-only contract binding to access the raw methods on
}

// ParameterizerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ParameterizerTransactorRaw struct {
	Contract *ParameterizerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewParameterizer creates a new instance of Parameterizer, bound to a specific deployed contract.
func NewParameterizer(address common.Address, backend bind.ContractBackend) (*Parameterizer, error) {
	contract, err := bindParameterizer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Parameterizer{ParameterizerCaller: ParameterizerCaller{contract: contract}, ParameterizerTransactor: ParameterizerTransactor{contract: contract}, ParameterizerFilterer: ParameterizerFilterer{contract: contract}}, nil
}

// NewParameterizerCaller creates a new read-only instance of Parameterizer, bound to a specific deployed contract.
func NewParameterizerCaller(address common.Address, caller bind.ContractCaller) (*ParameterizerCaller, error) {
	contract, err := bindParameterizer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ParameterizerCaller{contract: contract}, nil
}

// NewParameterizerTransactor creates a new write-only instance of Parameterizer, bound to a specific deployed contract.
func NewParameterizerTransactor(address common.Address, transactor bind.ContractTransactor) (*ParameterizerTransactor, error) {
	contract, err := bindParameterizer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ParameterizerTransactor{contract: contract}, nil
}

// NewParameterizerFilterer creates a new log filterer instance of Parameterizer, bound to a specific deployed contract.
func NewParameterizerFilterer(address common.Address, filterer bind.ContractFilterer) (*ParameterizerFilterer, error) {
	contract, err := bindParameterizer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ParameterizerFilterer{contract: contract}, nil
}

// bindParameterizer binds a generic wrapper to an already deployed contract.
func bindParameterizer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ParameterizerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Parameterizer *ParameterizerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Parameterizer.Contract.ParameterizerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Parameterizer *ParameterizerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Parameterizer.Contract.ParameterizerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Parameterizer *ParameterizerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Parameterizer.Contract.ParameterizerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Parameterizer *ParameterizerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Parameterizer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Parameterizer *ParameterizerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Parameterizer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Parameterizer *ParameterizerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Parameterizer.Contract.contract.Transact(opts, method, params...)
}

// GetBackendPayment is a free data retrieval call binding the contract method 0x4b8b7dce.
//
// Solidity: function getBackendPayment() constant returns(uint256 out)
func (_Parameterizer *ParameterizerCaller) GetBackendPayment(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Parameterizer.contract.Call(opts, out, "getBackendPayment")
	return *ret0, err
}

// GetBackendPayment is a free data retrieval call binding the contract method 0x4b8b7dce.
//
// Solidity: function getBackendPayment() constant returns(uint256 out)
func (_Parameterizer *ParameterizerSession) GetBackendPayment() (*big.Int, error) {
	return _Parameterizer.Contract.GetBackendPayment(&_Parameterizer.CallOpts)
}

// GetBackendPayment is a free data retrieval call binding the contract method 0x4b8b7dce.
//
// Solidity: function getBackendPayment() constant returns(uint256 out)
func (_Parameterizer *ParameterizerCallerSession) GetBackendPayment() (*big.Int, error) {
	return _Parameterizer.Contract.GetBackendPayment(&_Parameterizer.CallOpts)
}

// GetCostPerByte is a free data retrieval call binding the contract method 0xcd2d5438.
//
// Solidity: function getCostPerByte() constant returns(uint256 out)
func (_Parameterizer *ParameterizerCaller) GetCostPerByte(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Parameterizer.contract.Call(opts, out, "getCostPerByte")
	return *ret0, err
}

// GetCostPerByte is a free data retrieval call binding the contract method 0xcd2d5438.
//
// Solidity: function getCostPerByte() constant returns(uint256 out)
func (_Parameterizer *ParameterizerSession) GetCostPerByte() (*big.Int, error) {
	return _Parameterizer.Contract.GetCostPerByte(&_Parameterizer.CallOpts)
}

// GetCostPerByte is a free data retrieval call binding the contract method 0xcd2d5438.
//
// Solidity: function getCostPerByte() constant returns(uint256 out)
func (_Parameterizer *ParameterizerCallerSession) GetCostPerByte() (*big.Int, error) {
	return _Parameterizer.Contract.GetCostPerByte(&_Parameterizer.CallOpts)
}

// GetHash is a free data retrieval call binding the contract method 0x1b855044.
//
// Solidity: function getHash(uint256 param, uint256 value) constant returns(bytes32 out)
func (_Parameterizer *ParameterizerCaller) GetHash(opts *bind.CallOpts, param *big.Int, value *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Parameterizer.contract.Call(opts, out, "getHash", param, value)
	return *ret0, err
}

// GetHash is a free data retrieval call binding the contract method 0x1b855044.
//
// Solidity: function getHash(uint256 param, uint256 value) constant returns(bytes32 out)
func (_Parameterizer *ParameterizerSession) GetHash(param *big.Int, value *big.Int) ([32]byte, error) {
	return _Parameterizer.Contract.GetHash(&_Parameterizer.CallOpts, param, value)
}

// GetHash is a free data retrieval call binding the contract method 0x1b855044.
//
// Solidity: function getHash(uint256 param, uint256 value) constant returns(bytes32 out)
func (_Parameterizer *ParameterizerCallerSession) GetHash(param *big.Int, value *big.Int) ([32]byte, error) {
	return _Parameterizer.Contract.GetHash(&_Parameterizer.CallOpts, param, value)
}

// GetListReward is a free data retrieval call binding the contract method 0xdc2b2853.
//
// Solidity: function getListReward() constant returns(uint256 out)
func (_Parameterizer *ParameterizerCaller) GetListReward(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Parameterizer.contract.Call(opts, out, "getListReward")
	return *ret0, err
}

// GetListReward is a free data retrieval call binding the contract method 0xdc2b2853.
//
// Solidity: function getListReward() constant returns(uint256 out)
func (_Parameterizer *ParameterizerSession) GetListReward() (*big.Int, error) {
	return _Parameterizer.Contract.GetListReward(&_Parameterizer.CallOpts)
}

// GetListReward is a free data retrieval call binding the contract method 0xdc2b2853.
//
// Solidity: function getListReward() constant returns(uint256 out)
func (_Parameterizer *ParameterizerCallerSession) GetListReward() (*big.Int, error) {
	return _Parameterizer.Contract.GetListReward(&_Parameterizer.CallOpts)
}

// GetMakerPayment is a free data retrieval call binding the contract method 0xccc3412f.
//
// Solidity: function getMakerPayment() constant returns(uint256 out)
func (_Parameterizer *ParameterizerCaller) GetMakerPayment(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Parameterizer.contract.Call(opts, out, "getMakerPayment")
	return *ret0, err
}

// GetMakerPayment is a free data retrieval call binding the contract method 0xccc3412f.
//
// Solidity: function getMakerPayment() constant returns(uint256 out)
func (_Parameterizer *ParameterizerSession) GetMakerPayment() (*big.Int, error) {
	return _Parameterizer.Contract.GetMakerPayment(&_Parameterizer.CallOpts)
}

// GetMakerPayment is a free data retrieval call binding the contract method 0xccc3412f.
//
// Solidity: function getMakerPayment() constant returns(uint256 out)
func (_Parameterizer *ParameterizerCallerSession) GetMakerPayment() (*big.Int, error) {
	return _Parameterizer.Contract.GetMakerPayment(&_Parameterizer.CallOpts)
}

// GetPlurality is a free data retrieval call binding the contract method 0x3d1e37d5.
//
// Solidity: function getPlurality() constant returns(uint256 out)
func (_Parameterizer *ParameterizerCaller) GetPlurality(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Parameterizer.contract.Call(opts, out, "getPlurality")
	return *ret0, err
}

// GetPlurality is a free data retrieval call binding the contract method 0x3d1e37d5.
//
// Solidity: function getPlurality() constant returns(uint256 out)
func (_Parameterizer *ParameterizerSession) GetPlurality() (*big.Int, error) {
	return _Parameterizer.Contract.GetPlurality(&_Parameterizer.CallOpts)
}

// GetPlurality is a free data retrieval call binding the contract method 0x3d1e37d5.
//
// Solidity: function getPlurality() constant returns(uint256 out)
func (_Parameterizer *ParameterizerCallerSession) GetPlurality() (*big.Int, error) {
	return _Parameterizer.Contract.GetPlurality(&_Parameterizer.CallOpts)
}

// GetPriceFloor is a free data retrieval call binding the contract method 0x03bca7d3.
//
// Solidity: function getPriceFloor() constant returns(uint256 out)
func (_Parameterizer *ParameterizerCaller) GetPriceFloor(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Parameterizer.contract.Call(opts, out, "getPriceFloor")
	return *ret0, err
}

// GetPriceFloor is a free data retrieval call binding the contract method 0x03bca7d3.
//
// Solidity: function getPriceFloor() constant returns(uint256 out)
func (_Parameterizer *ParameterizerSession) GetPriceFloor() (*big.Int, error) {
	return _Parameterizer.Contract.GetPriceFloor(&_Parameterizer.CallOpts)
}

// GetPriceFloor is a free data retrieval call binding the contract method 0x03bca7d3.
//
// Solidity: function getPriceFloor() constant returns(uint256 out)
func (_Parameterizer *ParameterizerCallerSession) GetPriceFloor() (*big.Int, error) {
	return _Parameterizer.Contract.GetPriceFloor(&_Parameterizer.CallOpts)
}

// GetReparam is a free data retrieval call binding the contract method 0xd90f8f0d.
//
// Solidity: function getReparam(bytes32 hash) constant returns(uint256 out, uint256 out)
func (_Parameterizer *ParameterizerCaller) GetReparam(opts *bind.CallOpts, hash [32]byte) (*big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Parameterizer.contract.Call(opts, out, "getReparam", hash)
	return *ret0, *ret1, err
}

// GetReparam is a free data retrieval call binding the contract method 0xd90f8f0d.
//
// Solidity: function getReparam(bytes32 hash) constant returns(uint256 out, uint256 out)
func (_Parameterizer *ParameterizerSession) GetReparam(hash [32]byte) (*big.Int, *big.Int, error) {
	return _Parameterizer.Contract.GetReparam(&_Parameterizer.CallOpts, hash)
}

// GetReparam is a free data retrieval call binding the contract method 0xd90f8f0d.
//
// Solidity: function getReparam(bytes32 hash) constant returns(uint256 out, uint256 out)
func (_Parameterizer *ParameterizerCallerSession) GetReparam(hash [32]byte) (*big.Int, *big.Int, error) {
	return _Parameterizer.Contract.GetReparam(&_Parameterizer.CallOpts, hash)
}

// GetReservePayment is a free data retrieval call binding the contract method 0xffe44da8.
//
// Solidity: function getReservePayment() constant returns(uint256 out)
func (_Parameterizer *ParameterizerCaller) GetReservePayment(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Parameterizer.contract.Call(opts, out, "getReservePayment")
	return *ret0, err
}

// GetReservePayment is a free data retrieval call binding the contract method 0xffe44da8.
//
// Solidity: function getReservePayment() constant returns(uint256 out)
func (_Parameterizer *ParameterizerSession) GetReservePayment() (*big.Int, error) {
	return _Parameterizer.Contract.GetReservePayment(&_Parameterizer.CallOpts)
}

// GetReservePayment is a free data retrieval call binding the contract method 0xffe44da8.
//
// Solidity: function getReservePayment() constant returns(uint256 out)
func (_Parameterizer *ParameterizerCallerSession) GetReservePayment() (*big.Int, error) {
	return _Parameterizer.Contract.GetReservePayment(&_Parameterizer.CallOpts)
}

// GetSpread is a free data retrieval call binding the contract method 0x7b3cee60.
//
// Solidity: function getSpread() constant returns(uint256 out)
func (_Parameterizer *ParameterizerCaller) GetSpread(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Parameterizer.contract.Call(opts, out, "getSpread")
	return *ret0, err
}

// GetSpread is a free data retrieval call binding the contract method 0x7b3cee60.
//
// Solidity: function getSpread() constant returns(uint256 out)
func (_Parameterizer *ParameterizerSession) GetSpread() (*big.Int, error) {
	return _Parameterizer.Contract.GetSpread(&_Parameterizer.CallOpts)
}

// GetSpread is a free data retrieval call binding the contract method 0x7b3cee60.
//
// Solidity: function getSpread() constant returns(uint256 out)
func (_Parameterizer *ParameterizerCallerSession) GetSpread() (*big.Int, error) {
	return _Parameterizer.Contract.GetSpread(&_Parameterizer.CallOpts)
}

// GetStake is a free data retrieval call binding the contract method 0xfc0e3d90.
//
// Solidity: function getStake() constant returns(uint256 out)
func (_Parameterizer *ParameterizerCaller) GetStake(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Parameterizer.contract.Call(opts, out, "getStake")
	return *ret0, err
}

// GetStake is a free data retrieval call binding the contract method 0xfc0e3d90.
//
// Solidity: function getStake() constant returns(uint256 out)
func (_Parameterizer *ParameterizerSession) GetStake() (*big.Int, error) {
	return _Parameterizer.Contract.GetStake(&_Parameterizer.CallOpts)
}

// GetStake is a free data retrieval call binding the contract method 0xfc0e3d90.
//
// Solidity: function getStake() constant returns(uint256 out)
func (_Parameterizer *ParameterizerCallerSession) GetStake() (*big.Int, error) {
	return _Parameterizer.Contract.GetStake(&_Parameterizer.CallOpts)
}

// GetVoteBy is a free data retrieval call binding the contract method 0x2d0d2bc6.
//
// Solidity: function getVoteBy() constant returns(uint256 out)
func (_Parameterizer *ParameterizerCaller) GetVoteBy(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Parameterizer.contract.Call(opts, out, "getVoteBy")
	return *ret0, err
}

// GetVoteBy is a free data retrieval call binding the contract method 0x2d0d2bc6.
//
// Solidity: function getVoteBy() constant returns(uint256 out)
func (_Parameterizer *ParameterizerSession) GetVoteBy() (*big.Int, error) {
	return _Parameterizer.Contract.GetVoteBy(&_Parameterizer.CallOpts)
}

// GetVoteBy is a free data retrieval call binding the contract method 0x2d0d2bc6.
//
// Solidity: function getVoteBy() constant returns(uint256 out)
func (_Parameterizer *ParameterizerCallerSession) GetVoteBy() (*big.Int, error) {
	return _Parameterizer.Contract.GetVoteBy(&_Parameterizer.CallOpts)
}

// Reparameterize is a paid mutator transaction binding the contract method 0xc1836218.
//
// Solidity: function reparameterize(uint256 param, uint256 value) returns()
func (_Parameterizer *ParameterizerTransactor) Reparameterize(opts *bind.TransactOpts, param *big.Int, value *big.Int) (*types.Transaction, error) {
	return _Parameterizer.contract.Transact(opts, "reparameterize", param, value)
}

// Reparameterize is a paid mutator transaction binding the contract method 0xc1836218.
//
// Solidity: function reparameterize(uint256 param, uint256 value) returns()
func (_Parameterizer *ParameterizerSession) Reparameterize(param *big.Int, value *big.Int) (*types.Transaction, error) {
	return _Parameterizer.Contract.Reparameterize(&_Parameterizer.TransactOpts, param, value)
}

// Reparameterize is a paid mutator transaction binding the contract method 0xc1836218.
//
// Solidity: function reparameterize(uint256 param, uint256 value) returns()
func (_Parameterizer *ParameterizerTransactorSession) Reparameterize(param *big.Int, value *big.Int) (*types.Transaction, error) {
	return _Parameterizer.Contract.Reparameterize(&_Parameterizer.TransactOpts, param, value)
}

// ResolveReparam is a paid mutator transaction binding the contract method 0x435c709a.
//
// Solidity: function resolveReparam(bytes32 hash) returns()
func (_Parameterizer *ParameterizerTransactor) ResolveReparam(opts *bind.TransactOpts, hash [32]byte) (*types.Transaction, error) {
	return _Parameterizer.contract.Transact(opts, "resolveReparam", hash)
}

// ResolveReparam is a paid mutator transaction binding the contract method 0x435c709a.
//
// Solidity: function resolveReparam(bytes32 hash) returns()
func (_Parameterizer *ParameterizerSession) ResolveReparam(hash [32]byte) (*types.Transaction, error) {
	return _Parameterizer.Contract.ResolveReparam(&_Parameterizer.TransactOpts, hash)
}

// ResolveReparam is a paid mutator transaction binding the contract method 0x435c709a.
//
// Solidity: function resolveReparam(bytes32 hash) returns()
func (_Parameterizer *ParameterizerTransactorSession) ResolveReparam(hash [32]byte) (*types.Transaction, error) {
	return _Parameterizer.Contract.ResolveReparam(&_Parameterizer.TransactOpts, hash)
}

// ParameterizerReparamFailedIterator is returned from FilterReparamFailed and is used to iterate over the raw logs and unpacked data for ReparamFailed events raised by the Parameterizer contract.
type ParameterizerReparamFailedIterator struct {
	Event *ParameterizerReparamFailed // Event containing the contract specifics and raw log

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
func (it *ParameterizerReparamFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ParameterizerReparamFailed)
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
		it.Event = new(ParameterizerReparamFailed)
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
func (it *ParameterizerReparamFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ParameterizerReparamFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ParameterizerReparamFailed represents a ReparamFailed event raised by the Parameterizer contract.
type ParameterizerReparamFailed struct {
	Hash  [32]byte
	Param *big.Int
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterReparamFailed is a free log retrieval operation binding the contract event 0x1f050cdbc74210070c4a499a73459732a158933d789331de2757eb4d35630253.
//
// Solidity: event ReparamFailed(bytes32 indexed hash, uint256 indexed param, uint256 value)
func (_Parameterizer *ParameterizerFilterer) FilterReparamFailed(opts *bind.FilterOpts, hash [][32]byte, param []*big.Int) (*ParameterizerReparamFailedIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var paramRule []interface{}
	for _, paramItem := range param {
		paramRule = append(paramRule, paramItem)
	}

	logs, sub, err := _Parameterizer.contract.FilterLogs(opts, "ReparamFailed", hashRule, paramRule)
	if err != nil {
		return nil, err
	}
	return &ParameterizerReparamFailedIterator{contract: _Parameterizer.contract, event: "ReparamFailed", logs: logs, sub: sub}, nil
}

// WatchReparamFailed is a free log subscription operation binding the contract event 0x1f050cdbc74210070c4a499a73459732a158933d789331de2757eb4d35630253.
//
// Solidity: event ReparamFailed(bytes32 indexed hash, uint256 indexed param, uint256 value)
func (_Parameterizer *ParameterizerFilterer) WatchReparamFailed(opts *bind.WatchOpts, sink chan<- *ParameterizerReparamFailed, hash [][32]byte, param []*big.Int) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var paramRule []interface{}
	for _, paramItem := range param {
		paramRule = append(paramRule, paramItem)
	}

	logs, sub, err := _Parameterizer.contract.WatchLogs(opts, "ReparamFailed", hashRule, paramRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ParameterizerReparamFailed)
				if err := _Parameterizer.contract.UnpackLog(event, "ReparamFailed", log); err != nil {
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

// ParameterizerReparamProposedIterator is returned from FilterReparamProposed and is used to iterate over the raw logs and unpacked data for ReparamProposed events raised by the Parameterizer contract.
type ParameterizerReparamProposedIterator struct {
	Event *ParameterizerReparamProposed // Event containing the contract specifics and raw log

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
func (it *ParameterizerReparamProposedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ParameterizerReparamProposed)
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
		it.Event = new(ParameterizerReparamProposed)
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
func (it *ParameterizerReparamProposedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ParameterizerReparamProposedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ParameterizerReparamProposed represents a ReparamProposed event raised by the Parameterizer contract.
type ParameterizerReparamProposed struct {
	Owner common.Address
	Hash  [32]byte
	Param *big.Int
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterReparamProposed is a free log retrieval operation binding the contract event 0xb38eb1d14189ba2de4ca92c3ead8d4c2a61d566d266d46b487c4fa82ae732842.
//
// Solidity: event ReparamProposed(address indexed owner, bytes32 indexed hash, uint256 indexed param, uint256 value)
func (_Parameterizer *ParameterizerFilterer) FilterReparamProposed(opts *bind.FilterOpts, owner []common.Address, hash [][32]byte, param []*big.Int) (*ParameterizerReparamProposedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var paramRule []interface{}
	for _, paramItem := range param {
		paramRule = append(paramRule, paramItem)
	}

	logs, sub, err := _Parameterizer.contract.FilterLogs(opts, "ReparamProposed", ownerRule, hashRule, paramRule)
	if err != nil {
		return nil, err
	}
	return &ParameterizerReparamProposedIterator{contract: _Parameterizer.contract, event: "ReparamProposed", logs: logs, sub: sub}, nil
}

// WatchReparamProposed is a free log subscription operation binding the contract event 0xb38eb1d14189ba2de4ca92c3ead8d4c2a61d566d266d46b487c4fa82ae732842.
//
// Solidity: event ReparamProposed(address indexed owner, bytes32 indexed hash, uint256 indexed param, uint256 value)
func (_Parameterizer *ParameterizerFilterer) WatchReparamProposed(opts *bind.WatchOpts, sink chan<- *ParameterizerReparamProposed, owner []common.Address, hash [][32]byte, param []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var paramRule []interface{}
	for _, paramItem := range param {
		paramRule = append(paramRule, paramItem)
	}

	logs, sub, err := _Parameterizer.contract.WatchLogs(opts, "ReparamProposed", ownerRule, hashRule, paramRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ParameterizerReparamProposed)
				if err := _Parameterizer.contract.UnpackLog(event, "ReparamProposed", log); err != nil {
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

// ParameterizerReparamSucceededIterator is returned from FilterReparamSucceeded and is used to iterate over the raw logs and unpacked data for ReparamSucceeded events raised by the Parameterizer contract.
type ParameterizerReparamSucceededIterator struct {
	Event *ParameterizerReparamSucceeded // Event containing the contract specifics and raw log

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
func (it *ParameterizerReparamSucceededIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ParameterizerReparamSucceeded)
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
		it.Event = new(ParameterizerReparamSucceeded)
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
func (it *ParameterizerReparamSucceededIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ParameterizerReparamSucceededIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ParameterizerReparamSucceeded represents a ReparamSucceeded event raised by the Parameterizer contract.
type ParameterizerReparamSucceeded struct {
	Hash  [32]byte
	Param *big.Int
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterReparamSucceeded is a free log retrieval operation binding the contract event 0xe6f29ce1d705c668b223a35e77096d95bbd994c121503fba4b29e65be92bc7e8.
//
// Solidity: event ReparamSucceeded(bytes32 indexed hash, uint256 indexed param, uint256 value)
func (_Parameterizer *ParameterizerFilterer) FilterReparamSucceeded(opts *bind.FilterOpts, hash [][32]byte, param []*big.Int) (*ParameterizerReparamSucceededIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var paramRule []interface{}
	for _, paramItem := range param {
		paramRule = append(paramRule, paramItem)
	}

	logs, sub, err := _Parameterizer.contract.FilterLogs(opts, "ReparamSucceeded", hashRule, paramRule)
	if err != nil {
		return nil, err
	}
	return &ParameterizerReparamSucceededIterator{contract: _Parameterizer.contract, event: "ReparamSucceeded", logs: logs, sub: sub}, nil
}

// WatchReparamSucceeded is a free log subscription operation binding the contract event 0xe6f29ce1d705c668b223a35e77096d95bbd994c121503fba4b29e65be92bc7e8.
//
// Solidity: event ReparamSucceeded(bytes32 indexed hash, uint256 indexed param, uint256 value)
func (_Parameterizer *ParameterizerFilterer) WatchReparamSucceeded(opts *bind.WatchOpts, sink chan<- *ParameterizerReparamSucceeded, hash [][32]byte, param []*big.Int) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var paramRule []interface{}
	for _, paramItem := range param {
		paramRule = append(paramRule, paramItem)
	}

	logs, sub, err := _Parameterizer.contract.WatchLogs(opts, "ReparamSucceeded", hashRule, paramRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ParameterizerReparamSucceeded)
				if err := _Parameterizer.contract.UnpackLog(event, "ReparamSucceeded", log); err != nil {
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
