pragma solidity ^0.4.18;


contract Validator {

  address public owner;
  // TODO(rbharath): A vote should be w.r.t a particular DataRegistry
  // and a particular datapoint. This probably requires a vote to be
  // a struct with this information.
  string[] public votes;

  function Validator() public {
    owner = msg.sender;
  }

  function addVote(string vote) public {
    require(msg.sender == owner);
    votes.push(vote);
  }

}
