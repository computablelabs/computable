pragma solidity ^0.4.18;


import "./DataCoin.sol";


// mock class using DataCoin 
contract DataCoinMock is DataCoin {

  function DataCoin(address initialAccount, uint256 initialBalance) public {
    balances[initialAccount] = initialBalance;
    totalSupply_ = initialBalance;
  }

}
