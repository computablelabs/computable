pragma solidity ^0.4.18;

import "truffle/Assert.sol";
import "../contracts/DataCoin.sol";


contract TestDataCoin {

  // Just tests that DataCoin can be instantiated
  function testDataCoinCreation() public {
    // This address is created in scripts/test.sh
    DataCoin datacoin = new DataCoin();
    uint256 totalSupply = datacoin.totalSupply();
    Assert.equal(totalSupply, 10000, "Total supply should be 10000");
  }
}
