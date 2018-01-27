pragma solidity ^0.4.18;

import "../contracts/DataCoin.sol";
import "../contracts/DataMining.sol";

contract TestDataMining {

  // Just tests that DataMining contract can be instantiated
  function testDataMiningCreation() public {
    address wallet = address(0x2bdd21761a483f71054e14f5b827213567971c676928d9a1808cbfa4b7501200);
    DataCoin datacoin = new DataCoin();
    DataMining datamining = new DataMining(wallet, datacoin);
  }
  
}
