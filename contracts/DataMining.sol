pragma solidity ^0.4.18;

import "./DataCoin.sol";
import "./Validator.sol";

/**
 * @title DataMining 
 * @dev DataMining is the contract that controls how new data is mined. 
 */
contract DataMining {
  using SafeMath for uint256;

  // The token being sold
  DataCoin public datacoin;

  // address where validator stakes are collected
  address public wallet;

  // TODO(rbharath): This is absolutely arbitrary and will likely
  // change.
  // Stake required for validators to be approved.
  uint public VALIDATOR_STAKE = 100;

  // Contains the list of validators
  Validator[] public validators;

  /**
   * event for Validator approval. 
   * @param validator address of validator approved
   * @param stake DataCoin placed as stake 
   */
  event ValidatorApproval(address indexed validator, uint256 stake);

  function DataMining(address _wallet, DataCoin _datacoin) public {
    require(_wallet != address(0));
    require(_datacoin != address(0));
    wallet = _wallet;
    datacoin = _datacoin;
  }

}
