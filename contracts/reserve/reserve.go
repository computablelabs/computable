// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package reserve

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

// ReserveABI is the input ABI used to generate the binding from.
const ReserveABI = "[{\"name\":\"Withdrawn\",\"inputs\":[{\"type\":\"address\",\"name\":\"owner\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"transferred\",\"indexed\":false,\"unit\":\"wei\"}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Supported\",\"inputs\":[{\"type\":\"address\",\"name\":\"owner\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"offered\",\"indexed\":false,\"unit\":\"wei\"},{\"type\":\"uint256\",\"name\":\"minted\",\"indexed\":false,\"unit\":\"wei\"}],\"anonymous\":false,\"type\":\"event\"},{\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"ether_token_addr\"},{\"type\":\"address\",\"name\":\"market_token_addr\"},{\"type\":\"address\",\"name\":\"p11r_addr\"}],\"constant\":false,\"payable\":false,\"type\":\"constructor\"},{\"name\":\"getSupportPrice\",\"outputs\":[{\"type\":\"uint256\",\"unit\":\"wei\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":11142},{\"name\":\"support\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256\",\"unit\":\"wei\",\"name\":\"offer\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":21428},{\"name\":\"getWithdrawalProceeds\",\"outputs\":[{\"type\":\"uint256\",\"unit\":\"wei\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"addr\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":6765},{\"name\":\"withdraw\",\"outputs\":[],\"inputs\":[],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":13277}]"

// ReserveBin is the compiled bytecode used for deploying new contracts.
const ReserveBin = `0x740100000000000000000000000000000000000000006020526f7fffffffffffffffffffffffffffffff6040527fffffffffffffffffffffffffffffffff8000000000000000000000000000000060605274012a05f1fffffffffffffffffffffffffdabf41c006080527ffffffffffffffffffffffffed5fa0e000000000000000000000000000000000060a0526060610d7a6101403934156100a157600080fd5b6020610d7a60c03960c05160205181106100ba57600080fd5b5060206020610d7a0160c03960c05160205181106100d757600080fd5b5060206040610d7a0160c03960c05160205181106100f457600080fd5b50610140516000556101605160015561018051600255610d6256600436101561000d57610c4d565b600035601c52740100000000000000000000000000000000000000006020526f7fffffffffffffffffffffffffffffff6040527fffffffffffffffffffffffffffffffff8000000000000000000000000000000060605274012a05f1fffffffffffffffffffffffffdabf41c006080527ffffffffffffffffffffffffed5fa0e000000000000000000000000000000000060a052600015610755575b610140526002543b6100ba57600080fd5b60025430186100c857600080fd5b60206101e060046303bca7d36101805261019c6002545afa6100e957600080fd5b6000506101e051610160526002543b61010157600080fd5b600254301861010f57600080fd5b60206102806004637b3cee606102205261023c6002545afa61013057600080fd5b60005061028051610200526000543b61014857600080fd5b600054301861015657600080fd5b602061034060246370a082316102c052306102e0526102dc6000545afa61017c57600080fd5b600050610340516102a0526001543b61019457600080fd5b60015430186101a257600080fd5b60206103e060046318160ddd6103805261039c6001545afa6101c357600080fd5b6000506103e05161036052670de0b6b3a7640000610360511015610452576101605168056bc75e2d631000006101f857600080fd5b68056bc75e2d63100000610200511515610213576000610239565b6102a051610200516102a0516102005102041461022f57600080fd5b6102a05161020051025b1515610246576000610308565b633b9aca0061020051151561025c576000610282565b6102a051610200516102a0516102005102041461027857600080fd5b6102a05161020051025b633b9aca006102005115156102985760006102be565b6102a051610200516102a051610200510204146102b457600080fd5b6102a05161020051025b0204146102ca57600080fd5b633b9aca006102005115156102e0576000610306565b6102a051610200516102a051610200510204146102fc57600080fd5b6102a05161020051025b025b046101605101101561031957600080fd5b68056bc75e2d6310000061032c57600080fd5b68056bc75e2d6310000061020051151561034757600061036d565b6102a051610200516102a0516102005102041461036357600080fd5b6102a05161020051025b151561037a57600061043c565b633b9aca006102005115156103905760006103b6565b6102a051610200516102a051610200510204146103ac57600080fd5b6102a05161020051025b633b9aca006102005115156103cc5760006103f2565b6102a051610200516102a051610200510204146103e857600080fd5b6102a05161020051025b0204146103fe57600080fd5b633b9aca0061020051151561041457600061043a565b6102a051610200516102a0516102005102041461043057600080fd5b6102a05161020051025b025b0461016051016000526000516101405156610753565b6101605160641515610465576000610485565b61036051606461036051606402041461047d57600080fd5b610360516064025b61048e57600080fd5b6064151561049d5760006104bd565b6103605160646103605160640204146104b557600080fd5b610360516064025b6102005115156104ce5760006104f4565b6102a051610200516102a051610200510204146104ea57600080fd5b6102a05161020051025b15156105015760006105c3565b633b9aca0061020051151561051757600061053d565b6102a051610200516102a0516102005102041461053357600080fd5b6102a05161020051025b633b9aca00610200511515610553576000610579565b6102a051610200516102a0516102005102041461056f57600080fd5b6102a05161020051025b02041461058557600080fd5b633b9aca0061020051151561059b5760006105c1565b6102a051610200516102a051610200510204146105b757600080fd5b6102a05161020051025b025b04610160510110156105d457600080fd5b606415156105e3576000610603565b6103605160646103605160640204146105fb57600080fd5b610360516064025b61060c57600080fd5b6064151561061b57600061063b565b61036051606461036051606402041461063357600080fd5b610360516064025b61020051151561064c576000610672565b6102a051610200516102a0516102005102041461066857600080fd5b6102a05161020051025b151561067f576000610741565b633b9aca006102005115156106955760006106bb565b6102a051610200516102a051610200510204146106b157600080fd5b6102a05161020051025b633b9aca006102005115156106d15760006106f7565b6102a051610200516102a051610200510204146106ed57600080fd5b6102a05161020051025b02041461070357600080fd5b633b9aca0061020051151561071957600061073f565b6102a051610200516102a0516102005102041461073557600080fd5b6102a05161020051025b025b04610160510160005260005161014051565b005b63a056a5b9600051141561078a57341561076e57600080fd5b600658016100a9565b610140526101405160005260206000f350005b6356c9493f600051141561097a5734156107a357600080fd5b61014051600658016100a9565b610180526101405261018051610140526101405160043510156107d257600080fd5b6000543b6107df57600080fd5b60005430186107ed57600080fd5b602061026060646323b872dd6101a052336101c052306101e052600435610200526101bc60006000545af161082157600080fd5b600050610260506101405161083557600080fd5b6101405160043504151561084a5760006108a6565b633b9aca006101405161085c57600080fd5b6101405160043504633b9aca006101405161087657600080fd5b610140516004350402041461088a57600080fd5b633b9aca006101405161089c57600080fd5b6101405160043504025b610280526001543b6108b757600080fd5b60015430186108c557600080fd5b60006000602463a0712d686102a052610280516102c0526102bc60006001545af16108ef57600080fd5b6001543b6108fc57600080fd5b600154301861090a57600080fd5b60206103c0604463a9059cbb610320523361034052610280516103605261033c60006001545af161093a57600080fd5b6000506103c0506004356103e0526102805161040052337fb19981ad2efd084bc7ddcdc94541b38fd493a1c176ffe2cd1e512e6ee0c34fe960406103e0a2005b600015610ad3575b61016052610140526000610140511861099a57600080fd5b6001543b6109a757600080fd5b60015430186109b557600080fd5b602061022060246370a082316101a052610140516101c0526101bc6001545afa6109de57600080fd5b60005061022051610180526000543b6109f657600080fd5b6000543018610a0457600080fd5b60206102e060246370a0823161026052306102805261027c6000545afa610a2a57600080fd5b6000506102e051610240526001543b610a4257600080fd5b6001543018610a5057600080fd5b602061038060046318160ddd6103205261033c6001545afa610a7157600080fd5b600050610380516103005261030051610a8957600080fd5b61030051610180511515610a9e576000610ac4565b61024051610180516102405161018051020414610aba57600080fd5b6102405161018051025b04600052600051610160515650005b634633020d6000511415610b2e573415610aec57600080fd5b6004356020518110610afd57600080fd5b506321c1de4661014052600435610160526101605160065801610982565b6101c0526101c05160005260206000f350005b633ccfd60b6000511415610c4c573415610b4757600080fd5b610140516321c1de4661018052336101a0526101a05160065801610982565b6102005261014052610200516101405260006101405111610b8657600080fd5b6001543b610b9357600080fd5b6001543018610ba157600080fd5b600060006024637e9d2ac161022052336102405261023c60006001545af1610bc857600080fd5b6000543b610bd557600080fd5b6000543018610be357600080fd5b6020610340604463a9059cbb6102a052336102c052610140516102e0526102bc60006000545af1610c1357600080fd5b600050610340506101405161036052337f7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d56020610360a2005b5b60006000fd5b61010f610d620361010f60003961010f610d62036000f3`

// DeployReserve deploys a new Ethereum contract, binding an instance of Reserve to it.
func DeployReserve(auth *bind.TransactOpts, backend bind.ContractBackend, ether_token_addr common.Address, market_token_addr common.Address, p11r_addr common.Address) (common.Address, *types.Transaction, *Reserve, error) {
	parsed, err := abi.JSON(strings.NewReader(ReserveABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ReserveBin), backend, ether_token_addr, market_token_addr, p11r_addr)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Reserve{ReserveCaller: ReserveCaller{contract: contract}, ReserveTransactor: ReserveTransactor{contract: contract}, ReserveFilterer: ReserveFilterer{contract: contract}}, nil
}

// Reserve is an auto generated Go binding around an Ethereum contract.
type Reserve struct {
	ReserveCaller     // Read-only binding to the contract
	ReserveTransactor // Write-only binding to the contract
	ReserveFilterer   // Log filterer for contract events
}

// ReserveCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReserveCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReserveTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReserveTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReserveFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReserveFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReserveSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReserveSession struct {
	Contract     *Reserve          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReserveCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReserveCallerSession struct {
	Contract *ReserveCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ReserveTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReserveTransactorSession struct {
	Contract     *ReserveTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ReserveRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReserveRaw struct {
	Contract *Reserve // Generic contract binding to access the raw methods on
}

// ReserveCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReserveCallerRaw struct {
	Contract *ReserveCaller // Generic read-only contract binding to access the raw methods on
}

// ReserveTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReserveTransactorRaw struct {
	Contract *ReserveTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReserve creates a new instance of Reserve, bound to a specific deployed contract.
func NewReserve(address common.Address, backend bind.ContractBackend) (*Reserve, error) {
	contract, err := bindReserve(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Reserve{ReserveCaller: ReserveCaller{contract: contract}, ReserveTransactor: ReserveTransactor{contract: contract}, ReserveFilterer: ReserveFilterer{contract: contract}}, nil
}

// NewReserveCaller creates a new read-only instance of Reserve, bound to a specific deployed contract.
func NewReserveCaller(address common.Address, caller bind.ContractCaller) (*ReserveCaller, error) {
	contract, err := bindReserve(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReserveCaller{contract: contract}, nil
}

// NewReserveTransactor creates a new write-only instance of Reserve, bound to a specific deployed contract.
func NewReserveTransactor(address common.Address, transactor bind.ContractTransactor) (*ReserveTransactor, error) {
	contract, err := bindReserve(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReserveTransactor{contract: contract}, nil
}

// NewReserveFilterer creates a new log filterer instance of Reserve, bound to a specific deployed contract.
func NewReserveFilterer(address common.Address, filterer bind.ContractFilterer) (*ReserveFilterer, error) {
	contract, err := bindReserve(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReserveFilterer{contract: contract}, nil
}

// bindReserve binds a generic wrapper to an already deployed contract.
func bindReserve(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ReserveABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Reserve *ReserveRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Reserve.Contract.ReserveCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Reserve *ReserveRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Reserve.Contract.ReserveTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Reserve *ReserveRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Reserve.Contract.ReserveTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Reserve *ReserveCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Reserve.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Reserve *ReserveTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Reserve.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Reserve *ReserveTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Reserve.Contract.contract.Transact(opts, method, params...)
}

// GetSupportPrice is a free data retrieval call binding the contract method 0xa056a5b9.
//
// Solidity: function getSupportPrice() constant returns(uint256 out)
func (_Reserve *ReserveCaller) GetSupportPrice(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Reserve.contract.Call(opts, out, "getSupportPrice")
	return *ret0, err
}

// GetSupportPrice is a free data retrieval call binding the contract method 0xa056a5b9.
//
// Solidity: function getSupportPrice() constant returns(uint256 out)
func (_Reserve *ReserveSession) GetSupportPrice() (*big.Int, error) {
	return _Reserve.Contract.GetSupportPrice(&_Reserve.CallOpts)
}

// GetSupportPrice is a free data retrieval call binding the contract method 0xa056a5b9.
//
// Solidity: function getSupportPrice() constant returns(uint256 out)
func (_Reserve *ReserveCallerSession) GetSupportPrice() (*big.Int, error) {
	return _Reserve.Contract.GetSupportPrice(&_Reserve.CallOpts)
}

// GetWithdrawalProceeds is a free data retrieval call binding the contract method 0x4633020d.
//
// Solidity: function getWithdrawalProceeds(address addr) constant returns(uint256 out)
func (_Reserve *ReserveCaller) GetWithdrawalProceeds(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Reserve.contract.Call(opts, out, "getWithdrawalProceeds", addr)
	return *ret0, err
}

// GetWithdrawalProceeds is a free data retrieval call binding the contract method 0x4633020d.
//
// Solidity: function getWithdrawalProceeds(address addr) constant returns(uint256 out)
func (_Reserve *ReserveSession) GetWithdrawalProceeds(addr common.Address) (*big.Int, error) {
	return _Reserve.Contract.GetWithdrawalProceeds(&_Reserve.CallOpts, addr)
}

// GetWithdrawalProceeds is a free data retrieval call binding the contract method 0x4633020d.
//
// Solidity: function getWithdrawalProceeds(address addr) constant returns(uint256 out)
func (_Reserve *ReserveCallerSession) GetWithdrawalProceeds(addr common.Address) (*big.Int, error) {
	return _Reserve.Contract.GetWithdrawalProceeds(&_Reserve.CallOpts, addr)
}

// Support is a paid mutator transaction binding the contract method 0x56c9493f.
//
// Solidity: function support(uint256 offer) returns()
func (_Reserve *ReserveTransactor) Support(opts *bind.TransactOpts, offer *big.Int) (*types.Transaction, error) {
	return _Reserve.contract.Transact(opts, "support", offer)
}

// Support is a paid mutator transaction binding the contract method 0x56c9493f.
//
// Solidity: function support(uint256 offer) returns()
func (_Reserve *ReserveSession) Support(offer *big.Int) (*types.Transaction, error) {
	return _Reserve.Contract.Support(&_Reserve.TransactOpts, offer)
}

// Support is a paid mutator transaction binding the contract method 0x56c9493f.
//
// Solidity: function support(uint256 offer) returns()
func (_Reserve *ReserveTransactorSession) Support(offer *big.Int) (*types.Transaction, error) {
	return _Reserve.Contract.Support(&_Reserve.TransactOpts, offer)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Reserve *ReserveTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Reserve.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Reserve *ReserveSession) Withdraw() (*types.Transaction, error) {
	return _Reserve.Contract.Withdraw(&_Reserve.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Reserve *ReserveTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Reserve.Contract.Withdraw(&_Reserve.TransactOpts)
}

// ReserveSupportedIterator is returned from FilterSupported and is used to iterate over the raw logs and unpacked data for Supported events raised by the Reserve contract.
type ReserveSupportedIterator struct {
	Event *ReserveSupported // Event containing the contract specifics and raw log

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
func (it *ReserveSupportedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReserveSupported)
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
		it.Event = new(ReserveSupported)
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
func (it *ReserveSupportedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReserveSupportedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReserveSupported represents a Supported event raised by the Reserve contract.
type ReserveSupported struct {
	Owner   common.Address
	Offered *big.Int
	Minted  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSupported is a free log retrieval operation binding the contract event 0xb19981ad2efd084bc7ddcdc94541b38fd493a1c176ffe2cd1e512e6ee0c34fe9.
//
// Solidity: event Supported(address indexed owner, uint256 offered, uint256 minted)
func (_Reserve *ReserveFilterer) FilterSupported(opts *bind.FilterOpts, owner []common.Address) (*ReserveSupportedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Reserve.contract.FilterLogs(opts, "Supported", ownerRule)
	if err != nil {
		return nil, err
	}
	return &ReserveSupportedIterator{contract: _Reserve.contract, event: "Supported", logs: logs, sub: sub}, nil
}

// WatchSupported is a free log subscription operation binding the contract event 0xb19981ad2efd084bc7ddcdc94541b38fd493a1c176ffe2cd1e512e6ee0c34fe9.
//
// Solidity: event Supported(address indexed owner, uint256 offered, uint256 minted)
func (_Reserve *ReserveFilterer) WatchSupported(opts *bind.WatchOpts, sink chan<- *ReserveSupported, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Reserve.contract.WatchLogs(opts, "Supported", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReserveSupported)
				if err := _Reserve.contract.UnpackLog(event, "Supported", log); err != nil {
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

// ReserveWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the Reserve contract.
type ReserveWithdrawnIterator struct {
	Event *ReserveWithdrawn // Event containing the contract specifics and raw log

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
func (it *ReserveWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReserveWithdrawn)
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
		it.Event = new(ReserveWithdrawn)
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
func (it *ReserveWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReserveWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReserveWithdrawn represents a Withdrawn event raised by the Reserve contract.
type ReserveWithdrawn struct {
	Owner       common.Address
	Transferred *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed owner, uint256 transferred)
func (_Reserve *ReserveFilterer) FilterWithdrawn(opts *bind.FilterOpts, owner []common.Address) (*ReserveWithdrawnIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Reserve.contract.FilterLogs(opts, "Withdrawn", ownerRule)
	if err != nil {
		return nil, err
	}
	return &ReserveWithdrawnIterator{contract: _Reserve.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed owner, uint256 transferred)
func (_Reserve *ReserveFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *ReserveWithdrawn, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Reserve.contract.WatchLogs(opts, "Withdrawn", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReserveWithdrawn)
				if err := _Reserve.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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
