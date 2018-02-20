pragma solidity ^0.4.18;


contract DataRegistry {

  // The name of this registry
  string public name;

  // The address of the owner of this contract
  address public owner;

  function DataRegistry(string _name) public {
    owner = msg.sender;
    name = _name;
  }
}
