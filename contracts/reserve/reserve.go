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

// ReserveABI is the input ABI used to generate the binding from.
const ReserveABI = "[{\"name\":\"Withdrawn\",\"inputs\":[{\"type\":\"address\",\"name\":\"owner\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"transferred\",\"indexed\":false,\"unit\":\"wei\"}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Supported\",\"inputs\":[{\"type\":\"address\",\"name\":\"owner\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"offered\",\"indexed\":false,\"unit\":\"wei\"},{\"type\":\"uint256\",\"name\":\"minted\",\"indexed\":false,\"unit\":\"wei\"}],\"anonymous\":false,\"type\":\"event\"},{\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"ether_token_addr\"},{\"type\":\"address\",\"name\":\"market_token_addr\"},{\"type\":\"address\",\"name\":\"p11r_addr\"}],\"constant\":false,\"payable\":false,\"type\":\"constructor\"},{\"name\":\"getSupportPrice\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\",\"unit\":\"wei\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":10492},{\"name\":\"support\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256\",\"name\":\"offer\",\"unit\":\"wei\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":22186},{\"name\":\"getWithdrawalProceeds\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\",\"unit\":\"wei\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"addr\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":6038},{\"name\":\"withdraw\",\"outputs\":[],\"inputs\":[],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":13931}]"

// ReserveBin is the compiled bytecode used for deploying new contracts.
const ReserveBin = `0x740100000000000000000000000000000000000000006020526f7fffffffffffffffffffffffffffffff6040527fffffffffffffffffffffffffffffffff8000000000000000000000000000000060605274012a05f1fffffffffffffffffffffffffdabf41c006080527ffffffffffffffffffffffffed5fa0e000000000000000000000000000000000060a0526060610d1b6101403934156100a157600080fd5b6020610d1b60c03960c05160205181106100ba57600080fd5b5060206020610d1b0160c03960c05160205181106100d757600080fd5b5060206040610d1b0160c03960c05160205181106100f457600080fd5b50610140516000556101605160015561018051600255610d0356600035601c52740100000000000000000000000000000000000000006020526f7fffffffffffffffffffffffffffffff6040527fffffffffffffffffffffffffffffffff8000000000000000000000000000000060605274012a05f1fffffffffffffffffffffffffdabf41c006080527ffffffffffffffffffffffffed5fa0e000000000000000000000000000000000060a05263a056a5b9600051141561074e5734156100ac57600080fd5b6002543b6100b957600080fd5b60025430186100c757600080fd5b60206101c060046303bca7d36101605261017c6002545afa6100e857600080fd5b6000506101c051610140526002543b61010057600080fd5b600254301861010e57600080fd5b60206102606004637b3cee606102005261021c6002545afa61012f57600080fd5b600050610260516101e0526000543b61014757600080fd5b600054301861015557600080fd5b602061032060246370a082316102a052306102c0526102bc6000545afa61017b57600080fd5b60005061032051610280526001543b61019357600080fd5b60015430186101a157600080fd5b60206103c060046318160ddd6103605261037c6001545afa6101c257600080fd5b6000506103c05161034052670de0b6b3a764000061034051101561044e576101405168056bc75e2d631000006101f757600080fd5b68056bc75e2d631000006101e0511515610212576000610238565b610280516101e051610280516101e05102041461022e57600080fd5b610280516101e051025b1515610245576000610307565b633b9aca006101e051151561025b576000610281565b610280516101e051610280516101e05102041461027757600080fd5b610280516101e051025b633b9aca006101e05115156102975760006102bd565b610280516101e051610280516101e0510204146102b357600080fd5b610280516101e051025b0204146102c957600080fd5b633b9aca006101e05115156102df576000610305565b610280516101e051610280516101e0510204146102fb57600080fd5b610280516101e051025b025b046101405101101561031857600080fd5b68056bc75e2d6310000061032b57600080fd5b68056bc75e2d631000006101e051151561034657600061036c565b610280516101e051610280516101e05102041461036257600080fd5b610280516101e051025b151561037957600061043b565b633b9aca006101e051151561038f5760006103b5565b610280516101e051610280516101e0510204146103ab57600080fd5b610280516101e051025b633b9aca006101e05115156103cb5760006103f1565b610280516101e051610280516101e0510204146103e757600080fd5b610280516101e051025b0204146103fd57600080fd5b633b9aca006101e0511515610413576000610439565b610280516101e051610280516101e05102041461042f57600080fd5b610280516101e051025b025b04610140510160005260206000f361074c565b6101405160641515610461576000610481565b61034051606461034051606402041461047957600080fd5b610340516064025b61048a57600080fd5b606415156104995760006104b9565b6103405160646103405160640204146104b157600080fd5b610340516064025b6101e05115156104ca5760006104f0565b610280516101e051610280516101e0510204146104e657600080fd5b610280516101e051025b15156104fd5760006105bf565b633b9aca006101e0511515610513576000610539565b610280516101e051610280516101e05102041461052f57600080fd5b610280516101e051025b633b9aca006101e051151561054f576000610575565b610280516101e051610280516101e05102041461056b57600080fd5b610280516101e051025b02041461058157600080fd5b633b9aca006101e05115156105975760006105bd565b610280516101e051610280516101e0510204146105b357600080fd5b610280516101e051025b025b04610140510110156105d057600080fd5b606415156105df5760006105ff565b6103405160646103405160640204146105f757600080fd5b610340516064025b61060857600080fd5b60641515610617576000610637565b61034051606461034051606402041461062f57600080fd5b610340516064025b6101e051151561064857600061066e565b610280516101e051610280516101e05102041461066457600080fd5b610280516101e051025b151561067b57600061073d565b633b9aca006101e05115156106915760006106b7565b610280516101e051610280516101e0510204146106ad57600080fd5b610280516101e051025b633b9aca006101e05115156106cd5760006106f3565b610280516101e051610280516101e0510204146106e957600080fd5b610280516101e051025b0204146106ff57600080fd5b633b9aca006101e051151561071557600061073b565b610280516101e051610280516101e05102041461073157600080fd5b610280516101e051025b025b04610140510160005260206000f35b005b6356c9493f6000511415610959576020600461014037341561076f57600080fd5b60206101e0600463a056a5b96101805261019c6000305af161079057600080fd5b6101e05161016052610160516101405110156107ab57600080fd5b6000543b6107b857600080fd5b60005430186107c657600080fd5b60206102c060646323b872dd6102005233610220523061024052610140516102605261021c60006000545af16107fb57600080fd5b6000506102c0506101605161080f57600080fd5b6101605161014051041515610825576000610884565b633b9aca006101605161083757600080fd5b610160516101405104633b9aca006101605161085257600080fd5b61016051610140510402041461086757600080fd5b633b9aca006101605161087957600080fd5b610160516101405104025b6102e0526001543b61089557600080fd5b60015430186108a357600080fd5b60006000602463a0712d68610300526102e0516103205261031c60006001545af16108cd57600080fd5b6001543b6108da57600080fd5b60015430186108e857600080fd5b6020610420604463a9059cbb61038052336103a0526102e0516103c05261039c60006001545af161091857600080fd5b6000506104205061014051610440526102e05161046052337fb19981ad2efd084bc7ddcdc94541b38fd493a1c176ffe2cd1e512e6ee0c34fe96040610440a2005b634633020d6000511415610ad1576020600461014037341561097a57600080fd5b600435602051811061098b57600080fd5b506000610140511861099c57600080fd5b6001543b6109a957600080fd5b60015430186109b757600080fd5b602061020060246370a0823161018052610140516101a05261019c6001545afa6109e057600080fd5b60005061020051610160526000543b6109f857600080fd5b6000543018610a0657600080fd5b60206102c060246370a0823161024052306102605261025c6000545afa610a2c57600080fd5b6000506102c051610220526001543b610a4457600080fd5b6001543018610a5257600080fd5b602061036060046318160ddd6103005261031c6001545afa610a7357600080fd5b600050610360516102e0526102e051610a8b57600080fd5b6102e051610160511515610aa0576000610ac6565b61022051610160516102205161016051020414610abc57600080fd5b6102205161016051025b0460005260206000f3005b633ccfd60b6000511415610bee573415610aea57600080fd5b60206101e06024634633020d61016052336101805261017c6000305af1610b1057600080fd5b6101e0516101405260006101405111610b2857600080fd5b6001543b610b3557600080fd5b6001543018610b4357600080fd5b600060006024637e9d2ac161020052336102205261021c60006001545af1610b6a57600080fd5b6000543b610b7757600080fd5b6000543018610b8557600080fd5b6020610320604463a9059cbb61028052336102a052610140516102c05261029c60006000545af1610bb557600080fd5b600050610320506101405161034052337f7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d56020610340a2005b60006000fd5b61010f610d030361010f60003961010f610d03036000f3`

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
// Solidity: function getSupportPrice() constant returns(out uint256)
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
// Solidity: function getSupportPrice() constant returns(out uint256)
func (_Reserve *ReserveSession) GetSupportPrice() (*big.Int, error) {
	return _Reserve.Contract.GetSupportPrice(&_Reserve.CallOpts)
}

// GetSupportPrice is a free data retrieval call binding the contract method 0xa056a5b9.
//
// Solidity: function getSupportPrice() constant returns(out uint256)
func (_Reserve *ReserveCallerSession) GetSupportPrice() (*big.Int, error) {
	return _Reserve.Contract.GetSupportPrice(&_Reserve.CallOpts)
}

// GetWithdrawalProceeds is a free data retrieval call binding the contract method 0x4633020d.
//
// Solidity: function getWithdrawalProceeds(addr address) constant returns(out uint256)
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
// Solidity: function getWithdrawalProceeds(addr address) constant returns(out uint256)
func (_Reserve *ReserveSession) GetWithdrawalProceeds(addr common.Address) (*big.Int, error) {
	return _Reserve.Contract.GetWithdrawalProceeds(&_Reserve.CallOpts, addr)
}

// GetWithdrawalProceeds is a free data retrieval call binding the contract method 0x4633020d.
//
// Solidity: function getWithdrawalProceeds(addr address) constant returns(out uint256)
func (_Reserve *ReserveCallerSession) GetWithdrawalProceeds(addr common.Address) (*big.Int, error) {
	return _Reserve.Contract.GetWithdrawalProceeds(&_Reserve.CallOpts, addr)
}

// Support is a paid mutator transaction binding the contract method 0x56c9493f.
//
// Solidity: function support(offer uint256) returns()
func (_Reserve *ReserveTransactor) Support(opts *bind.TransactOpts, offer *big.Int) (*types.Transaction, error) {
	return _Reserve.contract.Transact(opts, "support", offer)
}

// Support is a paid mutator transaction binding the contract method 0x56c9493f.
//
// Solidity: function support(offer uint256) returns()
func (_Reserve *ReserveSession) Support(offer *big.Int) (*types.Transaction, error) {
	return _Reserve.Contract.Support(&_Reserve.TransactOpts, offer)
}

// Support is a paid mutator transaction binding the contract method 0x56c9493f.
//
// Solidity: function support(offer uint256) returns()
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
// Solidity: e Supported(owner indexed address, offered uint256, minted uint256)
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
// Solidity: e Supported(owner indexed address, offered uint256, minted uint256)
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
// Solidity: e Withdrawn(owner indexed address, transferred uint256)
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
// Solidity: e Withdrawn(owner indexed address, transferred uint256)
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
