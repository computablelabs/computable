pragma solidity ^0.4.18;

import "truffle/Assert.sol";
import "truffle/DeployedAddresses.sol";
import "../contracts/Validator.sol";

contract TestValidator {

  // Just tests that validator can be instantiated
  function testValidatorConstruction() public {
    Validator validator = new Validator();
  }

  function testValidatorAddVote() public {
    Validator validator = new Validator();

    string memory vote = "vote";
    validator.addVote(vote);

  }
}
