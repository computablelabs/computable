pragma solidity ^0.4.18;


//import "./DataCoin.sol";
import "zeppelin-solidity/contracts/token/ERC20/BasicToken.sol";
import "zeppelin-solidity/contracts/mocks/BasicTokenMock.sol";


// mock class using DataCoin 
contract DataCoinMock is BasicTokenMock {


}

//contract DataCoinMock is BasicToken {
//
//  function BasicTokenMock(address initialAccount, uint256 initialBalance) public {
//    balances[initialAccount] = initialBalance;
//    totalSupply_ = initialBalance;
//  }
//
//}