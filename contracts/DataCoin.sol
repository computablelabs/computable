pragma solidity ^0.4.18;

import "zeppelin-solidity/contracts/token/ERC20/MintableToken.sol";

contract DataCoin is MintableToken {

  string public name = "DataCoin";
  string public symbol = "DC";
  uint8 public decimals = 2;
  // TODO(rbharath): This is absolutely arbitrary and will likely
  // change.
  uint public INITIAL_SUPPLY = 10000;


  function DataCoin() public {
    balances[msg.sender] = INITIAL_SUPPLY;
  }

	function sendCoin(address receiver, uint amount) returns(bool sufficient) {
		if (balances[msg.sender] < amount) return false;
		balances[msg.sender] -= amount;
		balances[receiver] += amount;
		return true;
	}

	function getBalance(address addr) returns(uint) {
  	return balances[addr];
	}

}
