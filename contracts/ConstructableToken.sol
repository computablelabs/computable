pragma solidity 0.4.24;

import "./StandardToken.sol";


/**
 * A subclass of the Open Zeppelin StandardToken with a constructor.
 *
 * NOTE: This is essentially what O.Z. refers to as a "StandardTokenMock".
 * We feel that it is just a subclass that belongs with the other contracts
 */
contract ConstructableToken is StandardToken {

  constructor(address initialAccount, uint256 initialBalance) public {
    balances[initialAccount] = initialBalance;
    totalSupply_ = initialBalance;
  }

}


