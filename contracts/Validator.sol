pragma solidity ^0.4.18;

contract Validator {

  address public owner;
  string[] public votes;

  function Validator() public {
    owner = msg.sender;
  }

  function addVote(string vote) public {
    require(msg.sender == owner);
    votes.push(vote);
  }

}
