pragma solidity ^0.4.18;

import "../contracts/DataCoin.sol";

contract TestDataCoin {

  // Just tests that DataCoin can be instantiated
  function testDataCoinCreation() public {
    DataCoin datacoin = new DataCoin();
  }
}
