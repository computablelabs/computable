pragma solidity ^0.4.18;

//import "zeppelin-solidity/contracts/ownership/Ownable.sol";

contract Migrations {
  address public owner;
  uint public last_completed_migration;

  modifier restricted() {
    if (msg.sender == owner) {
      _;
    }
  }

  function Migrations() public {
    owner = msg.sender;
  }

  function setCompleted(uint completed) public restricted {
    last_completed_migration = completed;
  }

  function upgrade(address newAddress) public restricted {
    Migrations upgraded = Migrations(newAddress);
    upgraded.setCompleted(last_completed_migration);
  }
}

///**
// * @title Migrations
// * @dev This is a truffle contract, needed for truffle integration, not meant for use by Zeppelin users.
// */
//contract Migrations is Ownable {
//  uint256 public lastCompletedMigration;
//
//  function setCompleted(uint256 completed) onlyOwner public {
//    lastCompletedMigration = completed;
//  }
//
//  function upgrade(address newAddress) onlyOwner public {
//    Migrations upgraded = Migrations(newAddress);
//    upgraded.setCompleted(lastCompletedMigration);
//  }
//}
