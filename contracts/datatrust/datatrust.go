// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package datatrust

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

// DatatrustABI is the input ABI used to generate the binding from.
const DatatrustABI = "[{\"name\":\"Registered\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\",\"indexed\":true},{\"type\":\"address\",\"name\":\"registrant\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RegistrationSucceeded\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\",\"indexed\":true},{\"type\":\"address\",\"name\":\"registrant\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RegistrationFailed\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\",\"indexed\":true},{\"type\":\"address\",\"name\":\"registrant\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"DeliveryRequested\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\",\"indexed\":true},{\"type\":\"address\",\"name\":\"requester\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"amount\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Delivered\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\",\"indexed\":true},{\"type\":\"address\",\"name\":\"owner\",\"indexed\":true},{\"type\":\"bytes32\",\"name\":\"url\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"ether_token_addr\"},{\"type\":\"address\",\"name\":\"voting_addr\"},{\"type\":\"address\",\"name\":\"p11r_addr\"},{\"type\":\"address\",\"name\":\"res_addr\"}],\"constant\":false,\"payable\":false,\"type\":\"constructor\"},{\"name\":\"getPrivileged\",\"outputs\":[{\"type\":\"address\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":483},{\"name\":\"getReserve\",\"outputs\":[{\"type\":\"address\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":513},{\"name\":\"setPrivileged\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"listing\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":35978},{\"name\":\"getHash\",\"outputs\":[{\"type\":\"bytes32\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"string\",\"name\":\"url\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":615},{\"name\":\"getBackendAddress\",\"outputs\":[{\"type\":\"address\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":603},{\"name\":\"getBackendUrl\",\"outputs\":[{\"type\":\"string\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":21867},{\"name\":\"setBackendUrl\",\"outputs\":[],\"inputs\":[{\"type\":\"string\",\"name\":\"url\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":176967},{\"name\":\"getDataHash\",\"outputs\":[{\"type\":\"bytes32\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":826},{\"name\":\"setDataHash\",\"outputs\":[],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"listing\"},{\"type\":\"bytes32\",\"name\":\"data\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":35949},{\"name\":\"removeDataHash\",\"outputs\":[],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":20970},{\"name\":\"register\",\"outputs\":[],\"inputs\":[{\"type\":\"string\",\"name\":\"url\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":186313},{\"name\":\"resolveRegistration\",\"outputs\":[],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":48368},{\"name\":\"requestDelivery\",\"outputs\":[],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"},{\"type\":\"uint256\",\"name\":\"amount\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":116383},{\"name\":\"getBytesPurchased\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"addr\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":1045},{\"name\":\"getDelivery\",\"outputs\":[{\"type\":\"address\",\"name\":\"out\"},{\"type\":\"uint256\",\"name\":\"out\"},{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":1987},{\"name\":\"listingAccessed\",\"outputs\":[],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"listing\"},{\"type\":\"bytes32\",\"name\":\"delivery\"},{\"type\":\"uint256\",\"name\":\"amount\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":109361},{\"name\":\"getBytesAccessed\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":1096},{\"name\":\"bytesAccessedClaimed\",\"outputs\":[],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"},{\"type\":\"uint256\",\"name\":\"fee\",\"unit\":\"wei\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":23790},{\"name\":\"delivered\",\"outputs\":[],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"delivery\"},{\"type\":\"bytes32\",\"name\":\"url\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":77127}]"

// DatatrustBin is the compiled bytecode used for deploying new contracts.
const DatatrustBin = `0x740100000000000000000000000000000000000000006020526f7fffffffffffffffffffffffffffffff6040527fffffffffffffffffffffffffffffffff8000000000000000000000000000000060605274012a05f1fffffffffffffffffffffffffdabf41c006080527ffffffffffffffffffffffffed5fa0e000000000000000000000000000000000060a05260806113166101403934156100a157600080fd5b602061131660c03960c05160205181106100ba57600080fd5b50602060206113160160c03960c05160205181106100d757600080fd5b50602060406113160160c03960c05160205181106100f457600080fd5b50602060606113160160c03960c051602051811061011157600080fd5b50336009556101405160065561016051600755610180516008556101a051600a556112fe56600035601c52740100000000000000000000000000000000000000006020526f7fffffffffffffffffffffffffffffff6040527fffffffffffffffffffffffffffffffff8000000000000000000000000000000060605274012a05f1fffffffffffffffffffffffffdabf41c006080527ffffffffffffffffffffffffed5fa0e000000000000000000000000000000000060a05263d4c1753960005114156100b95734156100ac57600080fd5b600b5460005260206000f3005b6359bf5d3960005114156100df5734156100d257600080fd5b600a5460005260206000f3005b632ecace9c6000511415610136576020600461014037341561010057600080fd5b600435602051811061011157600080fd5b50600954331461012057600080fd5b600b541561012d57600080fd5b61014051600b55005b635b6beeb9600051141561018d576020600461014037341561015757600080fd5b60a060043560040161016037608060043560040135111561017757600080fd5b61016080516020820120905060005260206000f3005b63edb39a4060005114156101b35734156101a657600080fd5b60055460005260206000f3005b6376e1263560005114156102965734156101cc57600080fd5b60048060c052602060c020610180602082540161012060006005818352015b826101205160200211156101fe57610220565b61012051850154610120516020028501525b81516001018083528114156101eb575b5050505050506101805160206001820306601f8201039050610240610180516080818352015b8261024051111561025657610272565b6000610240516101a001535b8151600101808352811415610246575b5050506020610160526040610180510160206001820306601f8201039050610160f3005b6341d28f90600051141561034157602060046101403734156102b757600080fd5b60a06004356004016101603760806004356004013511156102d757600080fd5b60055433146102e557600080fd5b61016080600460c052602060c020602082510161012060006005818352015b8261012051602002111561031757610339565b61012051602002850151610120518501555b8151600101808352811415610304575b505050505050005b637a639f6e600051141561037e576020600461014037341561036257600080fd5b60006101405160e05260c052604060c0205460005260206000f3005b63b818bf0260005114156103c5576040600461014037341561039f57600080fd5b60055433146103ad57600080fd5b6101605160006101405160e05260c052604060c02055005b63bd82badc600051141561040a57602060046101403734156103e657600080fd5b600b5433146103f457600080fd5b600060006101405160e05260c052604060c02055005b63f2c298be6000511415610623576020600461014037341561042b57600080fd5b60a060043560040161016037608060043560040135111561044b57600080fd5b600554331861045957600080fd5b61016080600460c052602060c020602082510161012060006005818352015b8261012051602002111561048b576104ad565b61012051602002850151610120518501555b8151600101808352811415610478575b505050505050610160805160208201209050610220526007543b6104d057600080fd5b60075430186104de57600080fd5b60206102c0602463b89694c661024052610220516102605261025c6007545afa61050757600080fd5b6000506102c0511561051857600080fd5b6007543b61052557600080fd5b600754301861053357600080fd5b6000600060a463bb12c49e6103e052610220516104005260046104205233610440526008543b61056257600080fd5b600854301861057057600080fd5b6020610340600463fc0e3d906102e0526102fc6008545afa61059157600080fd5b60005061034051610460526008543b6105a957600080fd5b60085430186105b757600080fd5b60206103c06004632d0d2bc66103605261037c6008545afa6105d857600080fd5b6000506103c051610480526103fc60006007545af16105f657600080fd5b33610220517f7d917fcbc9a29a9705ff9936ffa599500e4fd902e4486bae317414fe967b307c60006000a3005b6384e1fe1560005114156108d6576020600461014037341561064457600080fd5b6007543b61065157600080fd5b600754301861065f57600080fd5b6020610200604463af61f76061016052610140516101805260046101a05261017c6007545afa61068e57600080fd5b6000506102005161069e57600080fd5b6007543b6106ab57600080fd5b60075430186106b957600080fd5b60206102a0602463327322c861022052610140516102405261023c6007545afa6106e257600080fd5b6000506102a0516106f257600080fd5b6007543b6106ff57600080fd5b600754301861070d57600080fd5b6020610360602463eb3014fd6102e05261014051610300526102fc6007545afa61073657600080fd5b600050610360516102c0526007543b61074e57600080fd5b600754301861075c57600080fd5b60206104a06044638f354b796104005261014051610420526008543b61078157600080fd5b600854301861078f57600080fd5b60206103e06004633d1e37d56103805261039c6008545afa6107b057600080fd5b6000506103e0516104405261041c6007545afa6107cc57600080fd5b6000506104a05115610812576102c0516005556102c051610140517ff9749f013eb1a881b147fd6da901e63089fadfb6fb375d6e56babcbcb5e0be4e60006000a361088f565b6102c051610140517ff83db24154eb020b1b0c94c9566e464732758c2c6bc070062458777d038baa3c60006000a3600080600460c052602060c020600161012060006001818352015b8261012051602002111561086e57610888565b6000610120518501555b815160010180835281141561085b575b5050505050505b6007543b61089c57600080fd5b60075430186108aa57600080fd5b6000600060246389bb617c6104c052610140516104e0526104dc60006007545af16108d457600080fd5b005b630ee206f56000511415610b4257604060046101403734156108f757600080fd5b60036101405160e05260c052604060c02060c052602060c020541561091b57600080fd5b6008543b61092857600080fd5b600854301861093657600080fd5b6020610200600463cd2d54386101a0526101bc6008545afa61095757600080fd5b6000506102005161022052610220511515610973576000610999565b6101605161022051610160516102205102041461098f57600080fd5b6101605161022051025b6101805260646109a857600080fd5b60646008543b6109b757600080fd5b60085430186109c557600080fd5b60206102c0600463ffe44da86102605261027c6008545afa6109e657600080fd5b6000506102c0516102e052610180511515610a02576000610a28565b6102e051610180516102e05161018051020414610a1e57600080fd5b6102e05161018051025b04610240526006543b610a3a57600080fd5b6006543018610a4857600080fd5b60206103c060646323b872dd6103005233610320523061034052610180516103605261031c60006006545af1610a7d57600080fd5b6000506103c0506006543b610a9157600080fd5b6006543018610a9f57600080fd5b6020610480604463a9059cbb6103e052600a546104005261024051610420526103fc60006006545af1610ad157600080fd5b6000506104805060013360e05260c052604060c0208054610160518254011015610afa57600080fd5b610160518154018155503360036101405160e05260c052604060c02060c052602060c0205561016051600160036101405160e05260c052604060c02060c052602060c0200155005b63f2cbc3fe6000511415610b915760206004610140373415610b6357600080fd5b6004356020518110610b7457600080fd5b5060016101405160e05260c052604060c0205460005260206000f3005b63edab743e6000511415610c265760206004610140373415610bb257600080fd5b60606101605261018060036101405160e05260c052604060c02060c052602060c020548152600160036101405160e05260c052604060c02060c052602060c02001548160200152600260036101405160e05260c052604060c02060c052602060c020015481604001525061016051610180f3005b63043b21666000511415610d205760606004610140373415610c4757600080fd5b6005543314610c5557600080fd5b600060006101405160e05260c052604060c0205418610c7357600080fd5b60026101405160e05260c052604060c0208054610180518254011015610c9857600080fd5b61018051815401815550600160036101605160e05260c052604060c02060c052602060c0205460e05260c052604060c0206101805181541015610cda57600080fd5b61018051815403815550600260036101605160e05260c052604060c02060c052602060c020018054610180518254011015610d1457600080fd5b61018051815401815550005b63e80239d56000511415610d5d5760206004610140373415610d4157600080fd5b60026101405160e05260c052604060c0205460005260206000f3005b639d38d6da6000511415610df65760406004610140373415610d7e57600080fd5b600b543314610d8c57600080fd5b600060026101405160e05260c052604060c020556006543b610dad57600080fd5b6006543018610dbb57600080fd5b6020610220604463a9059cbb61018052600a546101a052610160516101c05261019c60006006545af1610ded57600080fd5b60005061022050005b63d968041260005114156111c15760406004610140373415610e1757600080fd5b6005543314610e2557600080fd5b60036101405160e05260c052604060c02060c052602060c0205461018052600160036101405160e05260c052604060c02060c052602060c02001546101a0526101a051600260036101405160e05260c052604060c02060c052602060c02001541015610e9057600080fd5b60036101405160e05260c052604060c02060c052602060c020600081556000600182015560006002820155506064610ec757600080fd5b60646008543b610ed657600080fd5b6008543018610ee457600080fd5b60206102e06004634b8b7dce6102805261029c6008545afa610f0557600080fd5b6000506102e051610300526008543b610f1d57600080fd5b6008543018610f2b57600080fd5b6020610240600463cd2d54386101e0526101fc6008545afa610f4c57600080fd5b6000506102405161026052610260511515610f68576000610f8e565b6101a051610260516101a05161026051020414610f8457600080fd5b6101a05161026051025b1515610f9b57600061112f565b610300516008543b610fac57600080fd5b6008543018610fba57600080fd5b6020610240600463cd2d54386101e0526101fc6008545afa610fdb57600080fd5b6000506102405161026052610260511515610ff757600061101d565b6101a051610260516101a0516102605102041461101357600080fd5b6101a05161026051025b610300516008543b61102e57600080fd5b600854301861103c57600080fd5b6020610240600463cd2d54386101e0526101fc6008545afa61105d57600080fd5b600050610240516102605261026051151561107957600061109f565b6101a051610260516101a0516102605102041461109557600080fd5b6101a05161026051025b0204146110ab57600080fd5b610300516008543b6110bc57600080fd5b60085430186110ca57600080fd5b6020610240600463cd2d54386101e0526101fc6008545afa6110eb57600080fd5b600050610240516102605261026051151561110757600061112d565b6101a051610260516101a0516102605102041461112357600080fd5b6101a05161026051025b025b046101c0526006543b61114157600080fd5b600654301861114f57600080fd5b60206103c0604463a9059cbb61032052600554610340526101c0516103605261033c60006006545af161118157600080fd5b6000506103c050610160516103e05261018051610140517fa6bd58fcd4da90af9fea785d8cf919f762887f0f05b0656fb2ac5f7227183d5a60206103e0a3005b60006000fd5b6101376112fe036101376000396101376112fe036000f3`

// DeployDatatrust deploys a new Ethereum contract, binding an instance of Datatrust to it.
func DeployDatatrust(auth *bind.TransactOpts, backend bind.ContractBackend, ether_token_addr common.Address, voting_addr common.Address, p11r_addr common.Address, res_addr common.Address) (common.Address, *types.Transaction, *Datatrust, error) {
	parsed, err := abi.JSON(strings.NewReader(DatatrustABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DatatrustBin), backend, ether_token_addr, voting_addr, p11r_addr, res_addr)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Datatrust{DatatrustCaller: DatatrustCaller{contract: contract}, DatatrustTransactor: DatatrustTransactor{contract: contract}, DatatrustFilterer: DatatrustFilterer{contract: contract}}, nil
}

// Datatrust is an auto generated Go binding around an Ethereum contract.
type Datatrust struct {
	DatatrustCaller     // Read-only binding to the contract
	DatatrustTransactor // Write-only binding to the contract
	DatatrustFilterer   // Log filterer for contract events
}

// DatatrustCaller is an auto generated read-only Go binding around an Ethereum contract.
type DatatrustCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DatatrustTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DatatrustTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DatatrustFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DatatrustFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DatatrustSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DatatrustSession struct {
	Contract     *Datatrust        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DatatrustCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DatatrustCallerSession struct {
	Contract *DatatrustCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// DatatrustTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DatatrustTransactorSession struct {
	Contract     *DatatrustTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// DatatrustRaw is an auto generated low-level Go binding around an Ethereum contract.
type DatatrustRaw struct {
	Contract *Datatrust // Generic contract binding to access the raw methods on
}

// DatatrustCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DatatrustCallerRaw struct {
	Contract *DatatrustCaller // Generic read-only contract binding to access the raw methods on
}

// DatatrustTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DatatrustTransactorRaw struct {
	Contract *DatatrustTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDatatrust creates a new instance of Datatrust, bound to a specific deployed contract.
func NewDatatrust(address common.Address, backend bind.ContractBackend) (*Datatrust, error) {
	contract, err := bindDatatrust(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Datatrust{DatatrustCaller: DatatrustCaller{contract: contract}, DatatrustTransactor: DatatrustTransactor{contract: contract}, DatatrustFilterer: DatatrustFilterer{contract: contract}}, nil
}

// NewDatatrustCaller creates a new read-only instance of Datatrust, bound to a specific deployed contract.
func NewDatatrustCaller(address common.Address, caller bind.ContractCaller) (*DatatrustCaller, error) {
	contract, err := bindDatatrust(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DatatrustCaller{contract: contract}, nil
}

// NewDatatrustTransactor creates a new write-only instance of Datatrust, bound to a specific deployed contract.
func NewDatatrustTransactor(address common.Address, transactor bind.ContractTransactor) (*DatatrustTransactor, error) {
	contract, err := bindDatatrust(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DatatrustTransactor{contract: contract}, nil
}

// NewDatatrustFilterer creates a new log filterer instance of Datatrust, bound to a specific deployed contract.
func NewDatatrustFilterer(address common.Address, filterer bind.ContractFilterer) (*DatatrustFilterer, error) {
	contract, err := bindDatatrust(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DatatrustFilterer{contract: contract}, nil
}

// bindDatatrust binds a generic wrapper to an already deployed contract.
func bindDatatrust(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DatatrustABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Datatrust *DatatrustRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Datatrust.Contract.DatatrustCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Datatrust *DatatrustRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Datatrust.Contract.DatatrustTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Datatrust *DatatrustRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Datatrust.Contract.DatatrustTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Datatrust *DatatrustCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Datatrust.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Datatrust *DatatrustTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Datatrust.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Datatrust *DatatrustTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Datatrust.Contract.contract.Transact(opts, method, params...)
}

// GetBackendAddress is a free data retrieval call binding the contract method 0xedb39a40.
//
// Solidity: function getBackendAddress() constant returns(address out)
func (_Datatrust *DatatrustCaller) GetBackendAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Datatrust.contract.Call(opts, out, "getBackendAddress")
	return *ret0, err
}

// GetBackendAddress is a free data retrieval call binding the contract method 0xedb39a40.
//
// Solidity: function getBackendAddress() constant returns(address out)
func (_Datatrust *DatatrustSession) GetBackendAddress() (common.Address, error) {
	return _Datatrust.Contract.GetBackendAddress(&_Datatrust.CallOpts)
}

// GetBackendAddress is a free data retrieval call binding the contract method 0xedb39a40.
//
// Solidity: function getBackendAddress() constant returns(address out)
func (_Datatrust *DatatrustCallerSession) GetBackendAddress() (common.Address, error) {
	return _Datatrust.Contract.GetBackendAddress(&_Datatrust.CallOpts)
}

// GetBackendUrl is a free data retrieval call binding the contract method 0x76e12635.
//
// Solidity: function getBackendUrl() constant returns(string out)
func (_Datatrust *DatatrustCaller) GetBackendUrl(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Datatrust.contract.Call(opts, out, "getBackendUrl")
	return *ret0, err
}

// GetBackendUrl is a free data retrieval call binding the contract method 0x76e12635.
//
// Solidity: function getBackendUrl() constant returns(string out)
func (_Datatrust *DatatrustSession) GetBackendUrl() (string, error) {
	return _Datatrust.Contract.GetBackendUrl(&_Datatrust.CallOpts)
}

// GetBackendUrl is a free data retrieval call binding the contract method 0x76e12635.
//
// Solidity: function getBackendUrl() constant returns(string out)
func (_Datatrust *DatatrustCallerSession) GetBackendUrl() (string, error) {
	return _Datatrust.Contract.GetBackendUrl(&_Datatrust.CallOpts)
}

// GetBytesAccessed is a free data retrieval call binding the contract method 0xe80239d5.
//
// Solidity: function getBytesAccessed(bytes32 hash) constant returns(uint256 out)
func (_Datatrust *DatatrustCaller) GetBytesAccessed(opts *bind.CallOpts, hash [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Datatrust.contract.Call(opts, out, "getBytesAccessed", hash)
	return *ret0, err
}

// GetBytesAccessed is a free data retrieval call binding the contract method 0xe80239d5.
//
// Solidity: function getBytesAccessed(bytes32 hash) constant returns(uint256 out)
func (_Datatrust *DatatrustSession) GetBytesAccessed(hash [32]byte) (*big.Int, error) {
	return _Datatrust.Contract.GetBytesAccessed(&_Datatrust.CallOpts, hash)
}

// GetBytesAccessed is a free data retrieval call binding the contract method 0xe80239d5.
//
// Solidity: function getBytesAccessed(bytes32 hash) constant returns(uint256 out)
func (_Datatrust *DatatrustCallerSession) GetBytesAccessed(hash [32]byte) (*big.Int, error) {
	return _Datatrust.Contract.GetBytesAccessed(&_Datatrust.CallOpts, hash)
}

// GetBytesPurchased is a free data retrieval call binding the contract method 0xf2cbc3fe.
//
// Solidity: function getBytesPurchased(address addr) constant returns(uint256 out)
func (_Datatrust *DatatrustCaller) GetBytesPurchased(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Datatrust.contract.Call(opts, out, "getBytesPurchased", addr)
	return *ret0, err
}

// GetBytesPurchased is a free data retrieval call binding the contract method 0xf2cbc3fe.
//
// Solidity: function getBytesPurchased(address addr) constant returns(uint256 out)
func (_Datatrust *DatatrustSession) GetBytesPurchased(addr common.Address) (*big.Int, error) {
	return _Datatrust.Contract.GetBytesPurchased(&_Datatrust.CallOpts, addr)
}

// GetBytesPurchased is a free data retrieval call binding the contract method 0xf2cbc3fe.
//
// Solidity: function getBytesPurchased(address addr) constant returns(uint256 out)
func (_Datatrust *DatatrustCallerSession) GetBytesPurchased(addr common.Address) (*big.Int, error) {
	return _Datatrust.Contract.GetBytesPurchased(&_Datatrust.CallOpts, addr)
}

// GetDataHash is a free data retrieval call binding the contract method 0x7a639f6e.
//
// Solidity: function getDataHash(bytes32 hash) constant returns(bytes32 out)
func (_Datatrust *DatatrustCaller) GetDataHash(opts *bind.CallOpts, hash [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Datatrust.contract.Call(opts, out, "getDataHash", hash)
	return *ret0, err
}

// GetDataHash is a free data retrieval call binding the contract method 0x7a639f6e.
//
// Solidity: function getDataHash(bytes32 hash) constant returns(bytes32 out)
func (_Datatrust *DatatrustSession) GetDataHash(hash [32]byte) ([32]byte, error) {
	return _Datatrust.Contract.GetDataHash(&_Datatrust.CallOpts, hash)
}

// GetDataHash is a free data retrieval call binding the contract method 0x7a639f6e.
//
// Solidity: function getDataHash(bytes32 hash) constant returns(bytes32 out)
func (_Datatrust *DatatrustCallerSession) GetDataHash(hash [32]byte) ([32]byte, error) {
	return _Datatrust.Contract.GetDataHash(&_Datatrust.CallOpts, hash)
}

// GetDelivery is a free data retrieval call binding the contract method 0xedab743e.
//
// Solidity: function getDelivery(bytes32 hash) constant returns(address out, uint256 out, uint256 out)
func (_Datatrust *DatatrustCaller) GetDelivery(opts *bind.CallOpts, hash [32]byte) (common.Address, *big.Int, *big.Int, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _Datatrust.contract.Call(opts, out, "getDelivery", hash)
	return *ret0, *ret1, *ret2, err
}

// GetDelivery is a free data retrieval call binding the contract method 0xedab743e.
//
// Solidity: function getDelivery(bytes32 hash) constant returns(address out, uint256 out, uint256 out)
func (_Datatrust *DatatrustSession) GetDelivery(hash [32]byte) (common.Address, *big.Int, *big.Int, error) {
	return _Datatrust.Contract.GetDelivery(&_Datatrust.CallOpts, hash)
}

// GetDelivery is a free data retrieval call binding the contract method 0xedab743e.
//
// Solidity: function getDelivery(bytes32 hash) constant returns(address out, uint256 out, uint256 out)
func (_Datatrust *DatatrustCallerSession) GetDelivery(hash [32]byte) (common.Address, *big.Int, *big.Int, error) {
	return _Datatrust.Contract.GetDelivery(&_Datatrust.CallOpts, hash)
}

// GetHash is a free data retrieval call binding the contract method 0x5b6beeb9.
//
// Solidity: function getHash(string url) constant returns(bytes32 out)
func (_Datatrust *DatatrustCaller) GetHash(opts *bind.CallOpts, url string) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Datatrust.contract.Call(opts, out, "getHash", url)
	return *ret0, err
}

// GetHash is a free data retrieval call binding the contract method 0x5b6beeb9.
//
// Solidity: function getHash(string url) constant returns(bytes32 out)
func (_Datatrust *DatatrustSession) GetHash(url string) ([32]byte, error) {
	return _Datatrust.Contract.GetHash(&_Datatrust.CallOpts, url)
}

// GetHash is a free data retrieval call binding the contract method 0x5b6beeb9.
//
// Solidity: function getHash(string url) constant returns(bytes32 out)
func (_Datatrust *DatatrustCallerSession) GetHash(url string) ([32]byte, error) {
	return _Datatrust.Contract.GetHash(&_Datatrust.CallOpts, url)
}

// GetPrivileged is a free data retrieval call binding the contract method 0xd4c17539.
//
// Solidity: function getPrivileged() constant returns(address out)
func (_Datatrust *DatatrustCaller) GetPrivileged(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Datatrust.contract.Call(opts, out, "getPrivileged")
	return *ret0, err
}

// GetPrivileged is a free data retrieval call binding the contract method 0xd4c17539.
//
// Solidity: function getPrivileged() constant returns(address out)
func (_Datatrust *DatatrustSession) GetPrivileged() (common.Address, error) {
	return _Datatrust.Contract.GetPrivileged(&_Datatrust.CallOpts)
}

// GetPrivileged is a free data retrieval call binding the contract method 0xd4c17539.
//
// Solidity: function getPrivileged() constant returns(address out)
func (_Datatrust *DatatrustCallerSession) GetPrivileged() (common.Address, error) {
	return _Datatrust.Contract.GetPrivileged(&_Datatrust.CallOpts)
}

// GetReserve is a free data retrieval call binding the contract method 0x59bf5d39.
//
// Solidity: function getReserve() constant returns(address out)
func (_Datatrust *DatatrustCaller) GetReserve(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Datatrust.contract.Call(opts, out, "getReserve")
	return *ret0, err
}

// GetReserve is a free data retrieval call binding the contract method 0x59bf5d39.
//
// Solidity: function getReserve() constant returns(address out)
func (_Datatrust *DatatrustSession) GetReserve() (common.Address, error) {
	return _Datatrust.Contract.GetReserve(&_Datatrust.CallOpts)
}

// GetReserve is a free data retrieval call binding the contract method 0x59bf5d39.
//
// Solidity: function getReserve() constant returns(address out)
func (_Datatrust *DatatrustCallerSession) GetReserve() (common.Address, error) {
	return _Datatrust.Contract.GetReserve(&_Datatrust.CallOpts)
}

// BytesAccessedClaimed is a paid mutator transaction binding the contract method 0x9d38d6da.
//
// Solidity: function bytesAccessedClaimed(bytes32 hash, uint256 fee) returns()
func (_Datatrust *DatatrustTransactor) BytesAccessedClaimed(opts *bind.TransactOpts, hash [32]byte, fee *big.Int) (*types.Transaction, error) {
	return _Datatrust.contract.Transact(opts, "bytesAccessedClaimed", hash, fee)
}

// BytesAccessedClaimed is a paid mutator transaction binding the contract method 0x9d38d6da.
//
// Solidity: function bytesAccessedClaimed(bytes32 hash, uint256 fee) returns()
func (_Datatrust *DatatrustSession) BytesAccessedClaimed(hash [32]byte, fee *big.Int) (*types.Transaction, error) {
	return _Datatrust.Contract.BytesAccessedClaimed(&_Datatrust.TransactOpts, hash, fee)
}

// BytesAccessedClaimed is a paid mutator transaction binding the contract method 0x9d38d6da.
//
// Solidity: function bytesAccessedClaimed(bytes32 hash, uint256 fee) returns()
func (_Datatrust *DatatrustTransactorSession) BytesAccessedClaimed(hash [32]byte, fee *big.Int) (*types.Transaction, error) {
	return _Datatrust.Contract.BytesAccessedClaimed(&_Datatrust.TransactOpts, hash, fee)
}

// Delivered is a paid mutator transaction binding the contract method 0xd9680412.
//
// Solidity: function delivered(bytes32 delivery, bytes32 url) returns()
func (_Datatrust *DatatrustTransactor) Delivered(opts *bind.TransactOpts, delivery [32]byte, url [32]byte) (*types.Transaction, error) {
	return _Datatrust.contract.Transact(opts, "delivered", delivery, url)
}

// Delivered is a paid mutator transaction binding the contract method 0xd9680412.
//
// Solidity: function delivered(bytes32 delivery, bytes32 url) returns()
func (_Datatrust *DatatrustSession) Delivered(delivery [32]byte, url [32]byte) (*types.Transaction, error) {
	return _Datatrust.Contract.Delivered(&_Datatrust.TransactOpts, delivery, url)
}

// Delivered is a paid mutator transaction binding the contract method 0xd9680412.
//
// Solidity: function delivered(bytes32 delivery, bytes32 url) returns()
func (_Datatrust *DatatrustTransactorSession) Delivered(delivery [32]byte, url [32]byte) (*types.Transaction, error) {
	return _Datatrust.Contract.Delivered(&_Datatrust.TransactOpts, delivery, url)
}

// ListingAccessed is a paid mutator transaction binding the contract method 0x043b2166.
//
// Solidity: function listingAccessed(bytes32 listing, bytes32 delivery, uint256 amount) returns()
func (_Datatrust *DatatrustTransactor) ListingAccessed(opts *bind.TransactOpts, listing [32]byte, delivery [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Datatrust.contract.Transact(opts, "listingAccessed", listing, delivery, amount)
}

// ListingAccessed is a paid mutator transaction binding the contract method 0x043b2166.
//
// Solidity: function listingAccessed(bytes32 listing, bytes32 delivery, uint256 amount) returns()
func (_Datatrust *DatatrustSession) ListingAccessed(listing [32]byte, delivery [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Datatrust.Contract.ListingAccessed(&_Datatrust.TransactOpts, listing, delivery, amount)
}

// ListingAccessed is a paid mutator transaction binding the contract method 0x043b2166.
//
// Solidity: function listingAccessed(bytes32 listing, bytes32 delivery, uint256 amount) returns()
func (_Datatrust *DatatrustTransactorSession) ListingAccessed(listing [32]byte, delivery [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Datatrust.Contract.ListingAccessed(&_Datatrust.TransactOpts, listing, delivery, amount)
}

// Register is a paid mutator transaction binding the contract method 0xf2c298be.
//
// Solidity: function register(string url) returns()
func (_Datatrust *DatatrustTransactor) Register(opts *bind.TransactOpts, url string) (*types.Transaction, error) {
	return _Datatrust.contract.Transact(opts, "register", url)
}

// Register is a paid mutator transaction binding the contract method 0xf2c298be.
//
// Solidity: function register(string url) returns()
func (_Datatrust *DatatrustSession) Register(url string) (*types.Transaction, error) {
	return _Datatrust.Contract.Register(&_Datatrust.TransactOpts, url)
}

// Register is a paid mutator transaction binding the contract method 0xf2c298be.
//
// Solidity: function register(string url) returns()
func (_Datatrust *DatatrustTransactorSession) Register(url string) (*types.Transaction, error) {
	return _Datatrust.Contract.Register(&_Datatrust.TransactOpts, url)
}

// RemoveDataHash is a paid mutator transaction binding the contract method 0xbd82badc.
//
// Solidity: function removeDataHash(bytes32 hash) returns()
func (_Datatrust *DatatrustTransactor) RemoveDataHash(opts *bind.TransactOpts, hash [32]byte) (*types.Transaction, error) {
	return _Datatrust.contract.Transact(opts, "removeDataHash", hash)
}

// RemoveDataHash is a paid mutator transaction binding the contract method 0xbd82badc.
//
// Solidity: function removeDataHash(bytes32 hash) returns()
func (_Datatrust *DatatrustSession) RemoveDataHash(hash [32]byte) (*types.Transaction, error) {
	return _Datatrust.Contract.RemoveDataHash(&_Datatrust.TransactOpts, hash)
}

// RemoveDataHash is a paid mutator transaction binding the contract method 0xbd82badc.
//
// Solidity: function removeDataHash(bytes32 hash) returns()
func (_Datatrust *DatatrustTransactorSession) RemoveDataHash(hash [32]byte) (*types.Transaction, error) {
	return _Datatrust.Contract.RemoveDataHash(&_Datatrust.TransactOpts, hash)
}

// RequestDelivery is a paid mutator transaction binding the contract method 0x0ee206f5.
//
// Solidity: function requestDelivery(bytes32 hash, uint256 amount) returns()
func (_Datatrust *DatatrustTransactor) RequestDelivery(opts *bind.TransactOpts, hash [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Datatrust.contract.Transact(opts, "requestDelivery", hash, amount)
}

// RequestDelivery is a paid mutator transaction binding the contract method 0x0ee206f5.
//
// Solidity: function requestDelivery(bytes32 hash, uint256 amount) returns()
func (_Datatrust *DatatrustSession) RequestDelivery(hash [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Datatrust.Contract.RequestDelivery(&_Datatrust.TransactOpts, hash, amount)
}

// RequestDelivery is a paid mutator transaction binding the contract method 0x0ee206f5.
//
// Solidity: function requestDelivery(bytes32 hash, uint256 amount) returns()
func (_Datatrust *DatatrustTransactorSession) RequestDelivery(hash [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Datatrust.Contract.RequestDelivery(&_Datatrust.TransactOpts, hash, amount)
}

// ResolveRegistration is a paid mutator transaction binding the contract method 0x84e1fe15.
//
// Solidity: function resolveRegistration(bytes32 hash) returns()
func (_Datatrust *DatatrustTransactor) ResolveRegistration(opts *bind.TransactOpts, hash [32]byte) (*types.Transaction, error) {
	return _Datatrust.contract.Transact(opts, "resolveRegistration", hash)
}

// ResolveRegistration is a paid mutator transaction binding the contract method 0x84e1fe15.
//
// Solidity: function resolveRegistration(bytes32 hash) returns()
func (_Datatrust *DatatrustSession) ResolveRegistration(hash [32]byte) (*types.Transaction, error) {
	return _Datatrust.Contract.ResolveRegistration(&_Datatrust.TransactOpts, hash)
}

// ResolveRegistration is a paid mutator transaction binding the contract method 0x84e1fe15.
//
// Solidity: function resolveRegistration(bytes32 hash) returns()
func (_Datatrust *DatatrustTransactorSession) ResolveRegistration(hash [32]byte) (*types.Transaction, error) {
	return _Datatrust.Contract.ResolveRegistration(&_Datatrust.TransactOpts, hash)
}

// SetBackendUrl is a paid mutator transaction binding the contract method 0x41d28f90.
//
// Solidity: function setBackendUrl(string url) returns()
func (_Datatrust *DatatrustTransactor) SetBackendUrl(opts *bind.TransactOpts, url string) (*types.Transaction, error) {
	return _Datatrust.contract.Transact(opts, "setBackendUrl", url)
}

// SetBackendUrl is a paid mutator transaction binding the contract method 0x41d28f90.
//
// Solidity: function setBackendUrl(string url) returns()
func (_Datatrust *DatatrustSession) SetBackendUrl(url string) (*types.Transaction, error) {
	return _Datatrust.Contract.SetBackendUrl(&_Datatrust.TransactOpts, url)
}

// SetBackendUrl is a paid mutator transaction binding the contract method 0x41d28f90.
//
// Solidity: function setBackendUrl(string url) returns()
func (_Datatrust *DatatrustTransactorSession) SetBackendUrl(url string) (*types.Transaction, error) {
	return _Datatrust.Contract.SetBackendUrl(&_Datatrust.TransactOpts, url)
}

// SetDataHash is a paid mutator transaction binding the contract method 0xb818bf02.
//
// Solidity: function setDataHash(bytes32 listing, bytes32 data) returns()
func (_Datatrust *DatatrustTransactor) SetDataHash(opts *bind.TransactOpts, listing [32]byte, data [32]byte) (*types.Transaction, error) {
	return _Datatrust.contract.Transact(opts, "setDataHash", listing, data)
}

// SetDataHash is a paid mutator transaction binding the contract method 0xb818bf02.
//
// Solidity: function setDataHash(bytes32 listing, bytes32 data) returns()
func (_Datatrust *DatatrustSession) SetDataHash(listing [32]byte, data [32]byte) (*types.Transaction, error) {
	return _Datatrust.Contract.SetDataHash(&_Datatrust.TransactOpts, listing, data)
}

// SetDataHash is a paid mutator transaction binding the contract method 0xb818bf02.
//
// Solidity: function setDataHash(bytes32 listing, bytes32 data) returns()
func (_Datatrust *DatatrustTransactorSession) SetDataHash(listing [32]byte, data [32]byte) (*types.Transaction, error) {
	return _Datatrust.Contract.SetDataHash(&_Datatrust.TransactOpts, listing, data)
}

// SetPrivileged is a paid mutator transaction binding the contract method 0x2ecace9c.
//
// Solidity: function setPrivileged(address listing) returns()
func (_Datatrust *DatatrustTransactor) SetPrivileged(opts *bind.TransactOpts, listing common.Address) (*types.Transaction, error) {
	return _Datatrust.contract.Transact(opts, "setPrivileged", listing)
}

// SetPrivileged is a paid mutator transaction binding the contract method 0x2ecace9c.
//
// Solidity: function setPrivileged(address listing) returns()
func (_Datatrust *DatatrustSession) SetPrivileged(listing common.Address) (*types.Transaction, error) {
	return _Datatrust.Contract.SetPrivileged(&_Datatrust.TransactOpts, listing)
}

// SetPrivileged is a paid mutator transaction binding the contract method 0x2ecace9c.
//
// Solidity: function setPrivileged(address listing) returns()
func (_Datatrust *DatatrustTransactorSession) SetPrivileged(listing common.Address) (*types.Transaction, error) {
	return _Datatrust.Contract.SetPrivileged(&_Datatrust.TransactOpts, listing)
}

// DatatrustDeliveredIterator is returned from FilterDelivered and is used to iterate over the raw logs and unpacked data for Delivered events raised by the Datatrust contract.
type DatatrustDeliveredIterator struct {
	Event *DatatrustDelivered // Event containing the contract specifics and raw log

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
func (it *DatatrustDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DatatrustDelivered)
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
		it.Event = new(DatatrustDelivered)
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
func (it *DatatrustDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DatatrustDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DatatrustDelivered represents a Delivered event raised by the Datatrust contract.
type DatatrustDelivered struct {
	Hash  [32]byte
	Owner common.Address
	Url   [32]byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterDelivered is a free log retrieval operation binding the contract event 0xa6bd58fcd4da90af9fea785d8cf919f762887f0f05b0656fb2ac5f7227183d5a.
//
// Solidity: event Delivered(bytes32 indexed hash, address indexed owner, bytes32 url)
func (_Datatrust *DatatrustFilterer) FilterDelivered(opts *bind.FilterOpts, hash [][32]byte, owner []common.Address) (*DatatrustDeliveredIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Datatrust.contract.FilterLogs(opts, "Delivered", hashRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &DatatrustDeliveredIterator{contract: _Datatrust.contract, event: "Delivered", logs: logs, sub: sub}, nil
}

// WatchDelivered is a free log subscription operation binding the contract event 0xa6bd58fcd4da90af9fea785d8cf919f762887f0f05b0656fb2ac5f7227183d5a.
//
// Solidity: event Delivered(bytes32 indexed hash, address indexed owner, bytes32 url)
func (_Datatrust *DatatrustFilterer) WatchDelivered(opts *bind.WatchOpts, sink chan<- *DatatrustDelivered, hash [][32]byte, owner []common.Address) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Datatrust.contract.WatchLogs(opts, "Delivered", hashRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DatatrustDelivered)
				if err := _Datatrust.contract.UnpackLog(event, "Delivered", log); err != nil {
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

// DatatrustDeliveryRequestedIterator is returned from FilterDeliveryRequested and is used to iterate over the raw logs and unpacked data for DeliveryRequested events raised by the Datatrust contract.
type DatatrustDeliveryRequestedIterator struct {
	Event *DatatrustDeliveryRequested // Event containing the contract specifics and raw log

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
func (it *DatatrustDeliveryRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DatatrustDeliveryRequested)
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
		it.Event = new(DatatrustDeliveryRequested)
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
func (it *DatatrustDeliveryRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DatatrustDeliveryRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DatatrustDeliveryRequested represents a DeliveryRequested event raised by the Datatrust contract.
type DatatrustDeliveryRequested struct {
	Hash      [32]byte
	Requester common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeliveryRequested is a free log retrieval operation binding the contract event 0x898564ce29843afc6dc09bcabe85faaec3c18c45244d2c42e811a25bc0fd82c1.
//
// Solidity: event DeliveryRequested(bytes32 indexed hash, address indexed requester, uint256 amount)
func (_Datatrust *DatatrustFilterer) FilterDeliveryRequested(opts *bind.FilterOpts, hash [][32]byte, requester []common.Address) (*DatatrustDeliveryRequestedIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _Datatrust.contract.FilterLogs(opts, "DeliveryRequested", hashRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return &DatatrustDeliveryRequestedIterator{contract: _Datatrust.contract, event: "DeliveryRequested", logs: logs, sub: sub}, nil
}

// WatchDeliveryRequested is a free log subscription operation binding the contract event 0x898564ce29843afc6dc09bcabe85faaec3c18c45244d2c42e811a25bc0fd82c1.
//
// Solidity: event DeliveryRequested(bytes32 indexed hash, address indexed requester, uint256 amount)
func (_Datatrust *DatatrustFilterer) WatchDeliveryRequested(opts *bind.WatchOpts, sink chan<- *DatatrustDeliveryRequested, hash [][32]byte, requester []common.Address) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _Datatrust.contract.WatchLogs(opts, "DeliveryRequested", hashRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DatatrustDeliveryRequested)
				if err := _Datatrust.contract.UnpackLog(event, "DeliveryRequested", log); err != nil {
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

// DatatrustRegisteredIterator is returned from FilterRegistered and is used to iterate over the raw logs and unpacked data for Registered events raised by the Datatrust contract.
type DatatrustRegisteredIterator struct {
	Event *DatatrustRegistered // Event containing the contract specifics and raw log

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
func (it *DatatrustRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DatatrustRegistered)
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
		it.Event = new(DatatrustRegistered)
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
func (it *DatatrustRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DatatrustRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DatatrustRegistered represents a Registered event raised by the Datatrust contract.
type DatatrustRegistered struct {
	Hash       [32]byte
	Registrant common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRegistered is a free log retrieval operation binding the contract event 0x7d917fcbc9a29a9705ff9936ffa599500e4fd902e4486bae317414fe967b307c.
//
// Solidity: event Registered(bytes32 indexed hash, address indexed registrant)
func (_Datatrust *DatatrustFilterer) FilterRegistered(opts *bind.FilterOpts, hash [][32]byte, registrant []common.Address) (*DatatrustRegisteredIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var registrantRule []interface{}
	for _, registrantItem := range registrant {
		registrantRule = append(registrantRule, registrantItem)
	}

	logs, sub, err := _Datatrust.contract.FilterLogs(opts, "Registered", hashRule, registrantRule)
	if err != nil {
		return nil, err
	}
	return &DatatrustRegisteredIterator{contract: _Datatrust.contract, event: "Registered", logs: logs, sub: sub}, nil
}

// WatchRegistered is a free log subscription operation binding the contract event 0x7d917fcbc9a29a9705ff9936ffa599500e4fd902e4486bae317414fe967b307c.
//
// Solidity: event Registered(bytes32 indexed hash, address indexed registrant)
func (_Datatrust *DatatrustFilterer) WatchRegistered(opts *bind.WatchOpts, sink chan<- *DatatrustRegistered, hash [][32]byte, registrant []common.Address) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var registrantRule []interface{}
	for _, registrantItem := range registrant {
		registrantRule = append(registrantRule, registrantItem)
	}

	logs, sub, err := _Datatrust.contract.WatchLogs(opts, "Registered", hashRule, registrantRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DatatrustRegistered)
				if err := _Datatrust.contract.UnpackLog(event, "Registered", log); err != nil {
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

// DatatrustRegistrationFailedIterator is returned from FilterRegistrationFailed and is used to iterate over the raw logs and unpacked data for RegistrationFailed events raised by the Datatrust contract.
type DatatrustRegistrationFailedIterator struct {
	Event *DatatrustRegistrationFailed // Event containing the contract specifics and raw log

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
func (it *DatatrustRegistrationFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DatatrustRegistrationFailed)
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
		it.Event = new(DatatrustRegistrationFailed)
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
func (it *DatatrustRegistrationFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DatatrustRegistrationFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DatatrustRegistrationFailed represents a RegistrationFailed event raised by the Datatrust contract.
type DatatrustRegistrationFailed struct {
	Hash       [32]byte
	Registrant common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRegistrationFailed is a free log retrieval operation binding the contract event 0xf83db24154eb020b1b0c94c9566e464732758c2c6bc070062458777d038baa3c.
//
// Solidity: event RegistrationFailed(bytes32 indexed hash, address indexed registrant)
func (_Datatrust *DatatrustFilterer) FilterRegistrationFailed(opts *bind.FilterOpts, hash [][32]byte, registrant []common.Address) (*DatatrustRegistrationFailedIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var registrantRule []interface{}
	for _, registrantItem := range registrant {
		registrantRule = append(registrantRule, registrantItem)
	}

	logs, sub, err := _Datatrust.contract.FilterLogs(opts, "RegistrationFailed", hashRule, registrantRule)
	if err != nil {
		return nil, err
	}
	return &DatatrustRegistrationFailedIterator{contract: _Datatrust.contract, event: "RegistrationFailed", logs: logs, sub: sub}, nil
}

// WatchRegistrationFailed is a free log subscription operation binding the contract event 0xf83db24154eb020b1b0c94c9566e464732758c2c6bc070062458777d038baa3c.
//
// Solidity: event RegistrationFailed(bytes32 indexed hash, address indexed registrant)
func (_Datatrust *DatatrustFilterer) WatchRegistrationFailed(opts *bind.WatchOpts, sink chan<- *DatatrustRegistrationFailed, hash [][32]byte, registrant []common.Address) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var registrantRule []interface{}
	for _, registrantItem := range registrant {
		registrantRule = append(registrantRule, registrantItem)
	}

	logs, sub, err := _Datatrust.contract.WatchLogs(opts, "RegistrationFailed", hashRule, registrantRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DatatrustRegistrationFailed)
				if err := _Datatrust.contract.UnpackLog(event, "RegistrationFailed", log); err != nil {
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

// DatatrustRegistrationSucceededIterator is returned from FilterRegistrationSucceeded and is used to iterate over the raw logs and unpacked data for RegistrationSucceeded events raised by the Datatrust contract.
type DatatrustRegistrationSucceededIterator struct {
	Event *DatatrustRegistrationSucceeded // Event containing the contract specifics and raw log

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
func (it *DatatrustRegistrationSucceededIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DatatrustRegistrationSucceeded)
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
		it.Event = new(DatatrustRegistrationSucceeded)
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
func (it *DatatrustRegistrationSucceededIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DatatrustRegistrationSucceededIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DatatrustRegistrationSucceeded represents a RegistrationSucceeded event raised by the Datatrust contract.
type DatatrustRegistrationSucceeded struct {
	Hash       [32]byte
	Registrant common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRegistrationSucceeded is a free log retrieval operation binding the contract event 0xf9749f013eb1a881b147fd6da901e63089fadfb6fb375d6e56babcbcb5e0be4e.
//
// Solidity: event RegistrationSucceeded(bytes32 indexed hash, address indexed registrant)
func (_Datatrust *DatatrustFilterer) FilterRegistrationSucceeded(opts *bind.FilterOpts, hash [][32]byte, registrant []common.Address) (*DatatrustRegistrationSucceededIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var registrantRule []interface{}
	for _, registrantItem := range registrant {
		registrantRule = append(registrantRule, registrantItem)
	}

	logs, sub, err := _Datatrust.contract.FilterLogs(opts, "RegistrationSucceeded", hashRule, registrantRule)
	if err != nil {
		return nil, err
	}
	return &DatatrustRegistrationSucceededIterator{contract: _Datatrust.contract, event: "RegistrationSucceeded", logs: logs, sub: sub}, nil
}

// WatchRegistrationSucceeded is a free log subscription operation binding the contract event 0xf9749f013eb1a881b147fd6da901e63089fadfb6fb375d6e56babcbcb5e0be4e.
//
// Solidity: event RegistrationSucceeded(bytes32 indexed hash, address indexed registrant)
func (_Datatrust *DatatrustFilterer) WatchRegistrationSucceeded(opts *bind.WatchOpts, sink chan<- *DatatrustRegistrationSucceeded, hash [][32]byte, registrant []common.Address) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var registrantRule []interface{}
	for _, registrantItem := range registrant {
		registrantRule = append(registrantRule, registrantItem)
	}

	logs, sub, err := _Datatrust.contract.WatchLogs(opts, "RegistrationSucceeded", hashRule, registrantRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DatatrustRegistrationSucceeded)
				if err := _Datatrust.contract.UnpackLog(event, "RegistrationSucceeded", log); err != nil {
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
