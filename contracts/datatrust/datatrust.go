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
const DatatrustABI = "[{\"name\":\"Registered\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\",\"indexed\":true},{\"type\":\"address\",\"name\":\"registrant\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RegistrationSucceeded\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\",\"indexed\":true},{\"type\":\"address\",\"name\":\"registrant\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RegistrationFailed\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\",\"indexed\":true},{\"type\":\"address\",\"name\":\"registrant\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"DeliveryRequested\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\",\"indexed\":true},{\"type\":\"address\",\"name\":\"requester\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"amount\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Delivered\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\",\"indexed\":true},{\"type\":\"address\",\"name\":\"owner\",\"indexed\":true},{\"type\":\"bytes32\",\"name\":\"url\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"ListingAccessed\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\",\"indexed\":true},{\"type\":\"bytes32\",\"name\":\"delivery\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"amount\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"ether_token_addr\"},{\"type\":\"address\",\"name\":\"voting_addr\"},{\"type\":\"address\",\"name\":\"p11r_addr\"},{\"type\":\"address\",\"name\":\"res_addr\"}],\"constant\":false,\"payable\":false,\"type\":\"constructor\"},{\"name\":\"getPrivileged\",\"outputs\":[{\"type\":\"address\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":551},{\"name\":\"getReserve\",\"outputs\":[{\"type\":\"address\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":581},{\"name\":\"setPrivileged\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"listing\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":35998},{\"name\":\"getHash\",\"outputs\":[{\"type\":\"bytes32\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"string\",\"name\":\"url\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":665},{\"name\":\"getBackendAddress\",\"outputs\":[{\"type\":\"address\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":671},{\"name\":\"getBackendUrl\",\"outputs\":[{\"type\":\"string\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":22831},{\"name\":\"setBackendUrl\",\"outputs\":[],\"inputs\":[{\"type\":\"string\",\"name\":\"url\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":176987},{\"name\":\"getDataHash\",\"outputs\":[{\"type\":\"bytes32\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":876},{\"name\":\"setDataHash\",\"outputs\":[],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"listing\"},{\"type\":\"bytes32\",\"name\":\"data\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":35963},{\"name\":\"removeDataHash\",\"outputs\":[],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":20990},{\"name\":\"register\",\"outputs\":[],\"inputs\":[{\"type\":\"string\",\"name\":\"url\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":10292},{\"name\":\"resolveRegistration\",\"outputs\":[],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":68704},{\"name\":\"requestDelivery\",\"outputs\":[],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"},{\"type\":\"uint256\",\"name\":\"amount\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":225422},{\"name\":\"getBytesPurchased\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"addr\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":1095},{\"name\":\"getDelivery\",\"outputs\":[{\"type\":\"address\",\"name\":\"out\"},{\"type\":\"uint256\",\"name\":\"out\"},{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2037},{\"name\":\"listingAccessed\",\"outputs\":[],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"listing\"},{\"type\":\"bytes32\",\"name\":\"delivery\"},{\"type\":\"uint256\",\"name\":\"amount\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":120067},{\"name\":\"getAccessRewardEarned\",\"outputs\":[{\"type\":\"uint256\",\"unit\":\"wei\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":1146},{\"name\":\"accessRewardClaimed\",\"outputs\":[],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":24131},{\"name\":\"delivered\",\"outputs\":[],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"delivery\"},{\"type\":\"bytes32\",\"name\":\"url\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":136010}]"

// DatatrustBin is the compiled bytecode used for deploying new contracts.
const DatatrustBin = `0x740100000000000000000000000000000000000000006020526f7fffffffffffffffffffffffffffffff6040527fffffffffffffffffffffffffffffffff8000000000000000000000000000000060605274012a05f1fffffffffffffffffffffffffdabf41c006080527ffffffffffffffffffffffffed5fa0e000000000000000000000000000000000060a05260806116a66101403934156100a157600080fd5b60206116a660c03960c05160205181106100ba57600080fd5b50602060206116a60160c03960c05160205181106100d757600080fd5b50602060406116a60160c03960c05160205181106100f457600080fd5b50602060606116a60160c03960c051602051811061011157600080fd5b50336009556101405160065561016051600755610180516008556101a051600a5561168e56600436101561000d57611551565b600035601c52740100000000000000000000000000000000000000006020526f7fffffffffffffffffffffffffffffff6040527fffffffffffffffffffffffffffffffff8000000000000000000000000000000060605274012a05f1fffffffffffffffffffffffffdabf41c006080527ffffffffffffffffffffffffed5fa0e000000000000000000000000000000000060a05263d4c1753960005114156100c85734156100ba57600080fd5b600b5460005260206000f350005b6359bf5d3960005114156100ef5734156100e157600080fd5b600a5460005260206000f350005b632ecace9c600051141561013d57341561010857600080fd5b600435602051811061011957600080fd5b50600954331461012857600080fd5b600b541561013557600080fd5b600435600b55005b635b6beeb9600051141561018d57341561015657600080fd5b60a060043560040161014037608060043560040135111561017657600080fd5b61014080516020820120905060005260206000f350005b63edb39a4060005114156101b45734156101a657600080fd5b60055460005260206000f350005b6376e1263560005114156102995734156101cd57600080fd5b60048060c052602060c020610180602082540161012060006005818352015b826101205160200211156101ff57610221565b61012051850154610120516020028501525b81516001018083528114156101ec575b5050505050506101805160206001820306601f8201039050610240610180516080818352015b826102405110151561025857610274565b6000610240516101a001535b8151600101808352811415610247575b5050506020610160526040610180510160206001820306601f8201039050610160f350005b6341d28f90600051141561033c5734156102b257600080fd5b60a06004356004016101403760806004356004013511156102d257600080fd5b60055433146102e057600080fd5b61014080600460c052602060c020602082510161012060006005818352015b8261012051602002111561031257610334565b61012051602002850151610120518501555b81516001018083528114156102ff575b505050505050005b637a639f6e600051141561037157341561035557600080fd5b600060043560e05260c052604060c0205460005260206000f350005b63b818bf0260005114156103ae57341561038a57600080fd5b600554331461039857600080fd5b602435600060043560e05260c052604060c02055005b63bd82badc60005114156103ea5734156103c757600080fd5b600b5433146103d557600080fd5b6000600060043560e05260c052604060c02055005b63f2c298be60005114156105a157341561040357600080fd5b60a060043560040161014037608060043560040135111561042357600080fd5b600554331861043157600080fd5b610140805160208201209050610200526007543b61044e57600080fd5b600754301861045c57600080fd5b60206102a0602463b89694c661022052610200516102405261023c6007545afa61048557600080fd5b6000506102a0511561049657600080fd5b6007543b6104a357600080fd5b60075430186104b157600080fd5b6000600060a463bb12c49e6103c052610200516103e05260046104005233610420526008543b6104e057600080fd5b60085430186104ee57600080fd5b6020610320600463fc0e3d906102c0526102dc6008545afa61050f57600080fd5b60005061032051610440526008543b61052757600080fd5b600854301861053557600080fd5b60206103a06004632d0d2bc66103405261035c6008545afa61055657600080fd5b6000506103a051610460526103dc60006007545af161057457600080fd5b33610200517f7d917fcbc9a29a9705ff9936ffa599500e4fd902e4486bae317414fe967b307c60006000a3005b6384e1fe1560005114156108455734156105ba57600080fd5b6007543b6105c757600080fd5b60075430186105d557600080fd5b60206101e0604463af61f760610140526004356101605260046101805261015c6007545afa61060357600080fd5b6000506101e05161061357600080fd5b6007543b61062057600080fd5b600754301861062e57600080fd5b6020610280602463327322c8610200526004356102205261021c6007545afa61065657600080fd5b6000506102805161066657600080fd5b6007543b61067357600080fd5b600754301861068157600080fd5b6020610340602463eb3014fd6102c0526004356102e0526102dc6007545afa6106a957600080fd5b600050610340516102a0526007543b6106c157600080fd5b60075430186106cf57600080fd5b60206104806044638f354b796103e052600435610400526008543b6106f357600080fd5b600854301861070157600080fd5b60206103c06004633d1e37d56103605261037c6008545afa61072257600080fd5b6000506103c051610420526103fc6007545afa61073e57600080fd5b60005061048051156107d1576102a0516005556102a0516004357ff9749f013eb1a881b147fd6da901e63089fadfb6fb375d6e56babcbcb5e0be4e60006000a3600080600460c052602060c020600161012060006001818352015b826101205160200211156107ac576107c6565b6000610120518501555b8151600101808352811415610799575b5050505050506107ff565b6102a0516004357ff83db24154eb020b1b0c94c9566e464732758c2c6bc070062458777d038baa3c60006000a35b6007543b61080c57600080fd5b600754301861081a57600080fd5b6000600060246389bb617c6104a0526004356104c0526104bc60006007545af161084357600080fd5b005b630ee206f56000511415610b7e57341561085e57600080fd5b600260043560e05260c052604060c02060c052602060c020541561088157600080fd5b6008543b61088e57600080fd5b600854301861089c57600080fd5b60206101c0600463cd2d54386101605261017c6008545afa6108bd57600080fd5b6000506101c051610140526101405115156108d95760006108fc565b60243561014051602435610140510204146108f357600080fd5b60243561014051025b6101e052606461090b57600080fd5b60646008543b61091a57600080fd5b600854301861092857600080fd5b6020610280600463ffe44da86102205261023c6008545afa61094957600080fd5b600050610280516102a0526101e051151561096557600061098b565b6102a0516101e0516102a0516101e05102041461098157600080fd5b6102a0516101e051025b04610200526006543b61099d57600080fd5b60065430186109ab57600080fd5b602061038060646323b872dd6102c052336102e05230610300526101e051610320526102dc60006006545af16109e057600080fd5b600050610380506006543b6109f457600080fd5b6006543018610a0257600080fd5b6020610440604463a9059cbb6103a052600a546103c052610200516103e0526103bc60006006545af1610a3457600080fd5b6000506104405060013360e05260c052604060c02080546024358254011015610a5c57600080fd5b60243581540181555033600260043560e05260c052604060c02060c052602060c020556024356001600260043560e05260c052604060c02060c052602060c0200155610140516003600260043560e05260c052604060c02060c052602060c02001556008543b610acb57600080fd5b6008543018610ad957600080fd5b60206104c06004634b8b7dce6104605261047c6008545afa610afa57600080fd5b6000506104c0516004600260043560e05260c052604060c02060c052602060c02001556008543b610b2a57600080fd5b6008543018610b3857600080fd5b6020610540600463ccc3412f6104e0526104fc6008545afa610b5957600080fd5b600050610540516005600260043560e05260c052604060c02060c052602060c0200155005b63f2cbc3fe6000511415610bc5573415610b9757600080fd5b6004356020518110610ba857600080fd5b50600160043560e05260c052604060c0205460005260206000f350005b63edab743e6000511415610c50573415610bde57600080fd5b606061014052610160600260043560e05260c052604060c02060c052602060c0205481526001600260043560e05260c052604060c02060c052602060c020015481602001526002600260043560e05260c052604060c02060c052602060c020015481604001525061014051610160f350005b63043b2166600051141561103f573415610c6957600080fd5b6005543314610c7757600080fd5b6000600060043560e05260c052604060c0205418610c9457600080fd5b6001600260243560e05260c052604060c02060c052602060c0205460e05260c052604060c02060443581541015610cca57600080fd5b6044358154038155506002600260243560e05260c052604060c02060c052602060c0200180546044358254011015610d0157600080fd5b6044358154018155506064610d1557600080fd5b60646003600260243560e05260c052604060c02060c052602060c02001541515610d40576000610dab565b6044356003600260243560e05260c052604060c02060c052602060c02001546044356003600260243560e05260c052604060c02060c052602060c0200154020414610d8a57600080fd5b6044356003600260243560e05260c052604060c02060c052602060c0200154025b1515610db8576000610fd6565b6005600260243560e05260c052604060c02060c052602060c02001546003600260243560e05260c052604060c02060c052602060c02001541515610dfd576000610e68565b6044356003600260243560e05260c052604060c02060c052602060c02001546044356003600260243560e05260c052604060c02060c052602060c0200154020414610e4757600080fd5b6044356003600260243560e05260c052604060c02060c052602060c0200154025b6005600260243560e05260c052604060c02060c052602060c02001546003600260243560e05260c052604060c02060c052602060c02001541515610ead576000610f18565b6044356003600260243560e05260c052604060c02060c052602060c02001546044356003600260243560e05260c052604060c02060c052602060c0200154020414610ef757600080fd5b6044356003600260243560e05260c052604060c02060c052602060c0200154025b020414610f2457600080fd5b6005600260243560e05260c052604060c02060c052602060c02001546003600260243560e05260c052604060c02060c052602060c02001541515610f69576000610fd4565b6044356003600260243560e05260c052604060c02060c052602060c02001546044356003600260243560e05260c052604060c02060c052602060c0200154020414610fb357600080fd5b6044356003600260243560e05260c052604060c02060c052602060c0200154025b025b0461014052600360043560e05260c052604060c0208054610140518254011015610fff57600080fd5b61014051815401815550604435610160526024356004357f43bab07295538e531559fd23ac7cd44b6f236590ce18919de563b56d474f4aa56020610160a3005b63ea5450d0600051141561107457341561105857600080fd5b600360043560e05260c052604060c0205460005260206000f350005b63fcefc458600051141561111957341561108d57600080fd5b600b54331461109b57600080fd5b600360043560e05260c052604060c02054610140526000600360043560e05260c052604060c020556006543b6110d057600080fd5b60065430186110de57600080fd5b6020610200604463a9059cbb61016052600a5461018052610140516101a05261017c60006006545af161111057600080fd5b60005061020050005b63d9680412600051141561155057341561113257600080fd5b600554331461114057600080fd5b600260043560e05260c052604060c02060c052602060c02054610140526001600260043560e05260c052604060c02060c052602060c020015461016052610160516002600260043560e05260c052604060c02060c052602060c020015410156111a857600080fd5b60646111b357600080fd5b60646003600260043560e05260c052604060c02060c052602060c020015415156111de57600061124c565b610160516003600260043560e05260c052604060c02060c052602060c0200154610160516003600260043560e05260c052604060c02060c052602060c020015402041461122a57600080fd5b610160516003600260043560e05260c052604060c02060c052602060c0200154025b1515611259576000611480565b6004600260043560e05260c052604060c02060c052602060c02001546003600260043560e05260c052604060c02060c052602060c0200154151561129e57600061130c565b610160516003600260043560e05260c052604060c02060c052602060c0200154610160516003600260043560e05260c052604060c02060c052602060c02001540204146112ea57600080fd5b610160516003600260043560e05260c052604060c02060c052602060c0200154025b6004600260043560e05260c052604060c02060c052602060c02001546003600260043560e05260c052604060c02060c052602060c020015415156113515760006113bf565b610160516003600260043560e05260c052604060c02060c052602060c0200154610160516003600260043560e05260c052604060c02060c052602060c020015402041461139d57600080fd5b610160516003600260043560e05260c052604060c02060c052602060c0200154025b0204146113cb57600080fd5b6004600260043560e05260c052604060c02060c052602060c02001546003600260043560e05260c052604060c02060c052602060c0200154151561141057600061147e565b610160516003600260043560e05260c052604060c02060c052602060c0200154610160516003600260043560e05260c052604060c02060c052602060c020015402041461145c57600080fd5b610160516003600260043560e05260c052604060c02060c052602060c0200154025b025b0461018052600260043560e05260c052604060c02060c052602060c020600081556000600182015560006002820155600060038201556000600482015560006005820155506006543b6114d257600080fd5b60065430186114e057600080fd5b6020610240604463a9059cbb6101a0526005546101c052610180516101e0526101bc60006006545af161151257600080fd5b6000506102405060243561026052610140516004357fa6bd58fcd4da90af9fea785d8cf919f762887f0f05b0656fb2ac5f7227183d5a6020610260a3005b5b60006000fd5b61013761168e0361013760003961013761168e036000f3`

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

// GetAccessRewardEarned is a free data retrieval call binding the contract method 0xea5450d0.
//
// Solidity: function getAccessRewardEarned(bytes32 hash) constant returns(uint256 out)
func (_Datatrust *DatatrustCaller) GetAccessRewardEarned(opts *bind.CallOpts, hash [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Datatrust.contract.Call(opts, out, "getAccessRewardEarned", hash)
	return *ret0, err
}

// GetAccessRewardEarned is a free data retrieval call binding the contract method 0xea5450d0.
//
// Solidity: function getAccessRewardEarned(bytes32 hash) constant returns(uint256 out)
func (_Datatrust *DatatrustSession) GetAccessRewardEarned(hash [32]byte) (*big.Int, error) {
	return _Datatrust.Contract.GetAccessRewardEarned(&_Datatrust.CallOpts, hash)
}

// GetAccessRewardEarned is a free data retrieval call binding the contract method 0xea5450d0.
//
// Solidity: function getAccessRewardEarned(bytes32 hash) constant returns(uint256 out)
func (_Datatrust *DatatrustCallerSession) GetAccessRewardEarned(hash [32]byte) (*big.Int, error) {
	return _Datatrust.Contract.GetAccessRewardEarned(&_Datatrust.CallOpts, hash)
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

// AccessRewardClaimed is a paid mutator transaction binding the contract method 0xfcefc458.
//
// Solidity: function accessRewardClaimed(bytes32 hash) returns()
func (_Datatrust *DatatrustTransactor) AccessRewardClaimed(opts *bind.TransactOpts, hash [32]byte) (*types.Transaction, error) {
	return _Datatrust.contract.Transact(opts, "accessRewardClaimed", hash)
}

// AccessRewardClaimed is a paid mutator transaction binding the contract method 0xfcefc458.
//
// Solidity: function accessRewardClaimed(bytes32 hash) returns()
func (_Datatrust *DatatrustSession) AccessRewardClaimed(hash [32]byte) (*types.Transaction, error) {
	return _Datatrust.Contract.AccessRewardClaimed(&_Datatrust.TransactOpts, hash)
}

// AccessRewardClaimed is a paid mutator transaction binding the contract method 0xfcefc458.
//
// Solidity: function accessRewardClaimed(bytes32 hash) returns()
func (_Datatrust *DatatrustTransactorSession) AccessRewardClaimed(hash [32]byte) (*types.Transaction, error) {
	return _Datatrust.Contract.AccessRewardClaimed(&_Datatrust.TransactOpts, hash)
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

// DatatrustListingAccessedIterator is returned from FilterListingAccessed and is used to iterate over the raw logs and unpacked data for ListingAccessed events raised by the Datatrust contract.
type DatatrustListingAccessedIterator struct {
	Event *DatatrustListingAccessed // Event containing the contract specifics and raw log

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
func (it *DatatrustListingAccessedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DatatrustListingAccessed)
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
		it.Event = new(DatatrustListingAccessed)
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
func (it *DatatrustListingAccessedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DatatrustListingAccessedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DatatrustListingAccessed represents a ListingAccessed event raised by the Datatrust contract.
type DatatrustListingAccessed struct {
	Hash     [32]byte
	Delivery [32]byte
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterListingAccessed is a free log retrieval operation binding the contract event 0x43bab07295538e531559fd23ac7cd44b6f236590ce18919de563b56d474f4aa5.
//
// Solidity: event ListingAccessed(bytes32 indexed hash, bytes32 indexed delivery, uint256 amount)
func (_Datatrust *DatatrustFilterer) FilterListingAccessed(opts *bind.FilterOpts, hash [][32]byte, delivery [][32]byte) (*DatatrustListingAccessedIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var deliveryRule []interface{}
	for _, deliveryItem := range delivery {
		deliveryRule = append(deliveryRule, deliveryItem)
	}

	logs, sub, err := _Datatrust.contract.FilterLogs(opts, "ListingAccessed", hashRule, deliveryRule)
	if err != nil {
		return nil, err
	}
	return &DatatrustListingAccessedIterator{contract: _Datatrust.contract, event: "ListingAccessed", logs: logs, sub: sub}, nil
}

// WatchListingAccessed is a free log subscription operation binding the contract event 0x43bab07295538e531559fd23ac7cd44b6f236590ce18919de563b56d474f4aa5.
//
// Solidity: event ListingAccessed(bytes32 indexed hash, bytes32 indexed delivery, uint256 amount)
func (_Datatrust *DatatrustFilterer) WatchListingAccessed(opts *bind.WatchOpts, sink chan<- *DatatrustListingAccessed, hash [][32]byte, delivery [][32]byte) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var deliveryRule []interface{}
	for _, deliveryItem := range delivery {
		deliveryRule = append(deliveryRule, deliveryItem)
	}

	logs, sub, err := _Datatrust.contract.WatchLogs(opts, "ListingAccessed", hashRule, deliveryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DatatrustListingAccessed)
				if err := _Datatrust.contract.UnpackLog(event, "ListingAccessed", log); err != nil {
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
