pragma solidity ^0.4.18;

import "./DataCoin.sol";


contract DataRegistry {

  // ------
  // DATA STRUCTURES
  // ------

  struct Listing {
    uint applicationExpiry; // Expiration date of apply stage
    bool whitelisted;       // Indicates registry status
    address owner;          // Owner of Listing
    uint unstakedDeposit;   // Number of tokens in the listing not locked in a challenge
    uint challengeID;       // Corresponds to a PollID in PLCRVoting
  }

  // The name of this registry
  string public name;

  // The address of the owner of this contract
  address public owner;

  // The minimum required deposit
  uint public minDeposit = 10;

  // Maps listingHashes to associated listingHash data
  mapping(bytes32 => Listing) public listings;

  // Global Variables
  DataCoin public token;


  function DataRegistry(string _name) public {
    owner = msg.sender;
    name = _name;
  }

  // --------------------
  // PUBLISHER INTERFACE:
  // --------------------

  /**
  @dev                Allows a user to start an application. Takes tokens from user and sets
                      apply stage end time.
  @param _listingHash The hash of a potential listing a user is applying to add to the registry
  @param _amount      The number of DataCoins a user is willing to potentially stake
  @param _data        Extra data relevant to the application. Think IPFS hashes.
  */
  function apply(bytes32 _listingHash, uint _amount, string _data) external {
    //require(!isWhitelisted(_listingHash));
    //require(!appWasMade(_listingHash));
    //require(_amount >= parameterizer.get("minDeposit"));
    require(_amount >= minDeposit);

    // Sets owner
    Listing storage listing = listings[_listingHash];
    listing.owner = msg.sender;

    // Transfers tokens from user to Registry contract
    require(token.transferFrom(listing.owner, this, _amount));

    // Sets apply stage end time
    //listing.applicationExpiry = block.timestamp + parameterizer.get("applyStageLen");
    listing.unstakedDeposit = _amount;

    //_Application(_listingHash, _amount, _data);
  }
}
