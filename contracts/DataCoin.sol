pragma solidity ^0.4.18;

//import "zeppelin-solidity/contracts/token/ERC20/MintableToken.sol";
import "zeppelin-solidity/contracts/token/ERC20/BasicToken.sol";

//contract DataCoin is MintableToken {
contract DataCoin is BasicToken {

  string public name = "DataCoin";
  string public symbol = "DC";
  // TODO(rbharath): This is absolutely arbitrary and will likely
  // change.
  uint public INITIAL_SUPPLY = 10000;


  function DataCoin() public {
    balances[msg.sender] = INITIAL_SUPPLY;
  }

}
