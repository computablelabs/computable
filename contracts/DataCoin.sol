pragma solidity ^0.4.18;

import "zeppelin-solidity/contracts/token/ERC20/MintableToken.sol";

contract DataCoin is MintableToken {

  string public name = "DataCoin";
  string public symbol = "DC";
  uint8 public decimals = 2;
  // TODO(rbharath): This is absolutely arbitrary and will likely
  // change.
  uint public INITIAL_SUPPLY = 12000;


  function DataCoin() public {
    totalSupply_ = INITIAL_SUPPLY;
    balances[msg.sender] = INITIAL_SUPPLY;
  }

}
