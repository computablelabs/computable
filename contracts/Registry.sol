pragma solidity 0.4.24;

import "./ERC20.sol";
import "./Parameterizer.sol";
import "./PLCRVoting.sol";
import "./SafeMath.sol";


contract Registry {

  // ------
  // EVENTS
  // ------

  event _Application(bytes32 indexed listingHash, uint deposit, uint appEndDate, string data, address indexed applicant);
  event _Challenge(bytes32 indexed listingHash, uint challengeID, string data, uint commitEndDate, uint revealEndDate, address indexed challenger);
  event _Deposit(bytes32 indexed listingHash, uint added, uint newTotal, address indexed owner);
  event _Withdrawal(bytes32 indexed listingHash, uint withdrew, uint newTotal, address indexed owner);
  event _ApplicationWhitelisted(bytes32 indexed listingHash);
  event _ApplicationRemoved(bytes32 indexed listingHash);
  event _ListingRemoved(bytes32 indexed listingHash);
  event _ListingWithdrawn(bytes32 indexed listingHash);
  event _TouchAndRemoved(bytes32 indexed listingHash);
  event _ChallengeFailed(bytes32 indexed listingHash, uint indexed challengeID, uint rewardPool, uint totalTokens);
  event _ChallengeSucceeded(bytes32 indexed listingHash, uint indexed challengeID, uint rewardPool, uint totalTokens);
  event _RewardClaimed(uint indexed challengeID, uint reward, address indexed voter);

  using SafeMath for uint;

  struct Listing {
    uint applicationExpiry; // Expiration date of apply stage
    bool whitelisted;       // Indicates registry status
    address owner;          // Owner of Listing
    uint unstakedDeposit;   // Number of tokens in the listing not locked in a challenge
    uint challengeID;       // Corresponds to a PollID in PLCRVoting
  }

  struct Challenge {
    uint rewardPool;        // (remaining) Pool of tokens to be distributed to winning voters
    address challenger;     // Owner of Challenge
    bool resolved;          // Indication of if challenge is resolved
    uint stake;             // Number of tokens at stake for either party during challenge
    uint totalTokens;       // (remaining) Number of tokens used in voting by the winning side
    mapping(address => bool) tokenClaims; // Indicates whether a voter has claimed a reward yet
  }

  // Maps challengeIDs to associated challenge data
  mapping(uint => Challenge) public challenges;

  // Maps listingHashes to associated listingHash data
  mapping(bytes32 => Listing) public listings;

  // Global Variables
  ERC20 public token;
  PLCRVoting public voting;
  Parameterizer public parameterizer;
  string public name;

  // ------------
  // CONSTRUCTOR:
  // ------------

  /**
  @dev Contructor         Sets the addresses for token, voting, and parameterizer
  @param _tokenAddr       Address of the TCR's intrinsic ERC20 token
  @param _plcrAddr        Address of a PLCR voting contract for the provided token
  @param _paramsAddr      Address of a Parameterizer contract
  */
  constructor(
    address _tokenAddr,
    address _plcrAddr,
    address _paramsAddr,
    string _name
  ) public
  {
    token = ERC20(_tokenAddr);
    voting = PLCRVoting(_plcrAddr);
    parameterizer = Parameterizer(_paramsAddr);
    name = _name;
  }

  // --------------------
  // PUBLISHER INTERFACE:
  // --------------------

  /**
  @dev                Allows a user to start an application. Takes tokens from user and sets
                      apply stage end time.
  @param _listingHash The hash of a potential listing a user is applying to add to the registry
  @param _amount      The number of ERC20 tokens a user is willing to potentially stake
  @param _data        Extra data relevant to the application. Think IPFS hashes.
  */
  function apply(bytes32 _listingHash, uint _amount, string _data) external {
    require(!isWhitelisted(_listingHash));
    require(!appWasMade(_listingHash));
    require(_amount >= parameterizer.get("minDeposit"));

    // Sets owner
    Listing storage listing = listings[_listingHash];
    listing.owner = msg.sender;

    // Sets apply stage end time
    listing.applicationExpiry = block.timestamp.add(parameterizer.get("applyStageLen"));
    listing.unstakedDeposit = _amount;

    // Transfers tokens from user to Registry contract
    require(token.transferFrom(listing.owner, this, _amount));

    emit _Application(
      _listingHash,
      _amount,
      listing.applicationExpiry,
      _data,
      msg.sender
    );
  }

  /**
  @dev                Allows the owner of a listingHash to increase their unstaked deposit.
  @param _listingHash A listingHash msg.sender is the owner of
  @param _amount      The number of ERC20 tokens to increase a user's unstaked deposit
  */
  function deposit(bytes32 _listingHash, uint _amount) external {
    Listing storage listing = listings[_listingHash];

    require(listing.owner == msg.sender);

    listing.unstakedDeposit = listing.unstakedDeposit.add(_amount);
    require(token.transferFrom(msg.sender, this, _amount));

    emit _Deposit(
      _listingHash,
      _amount,
      listing.unstakedDeposit,
      msg.sender
    );
  }

  /**
  @dev                Allows the owner of a listingHash to decrease their unstaked deposit.
  @param _listingHash A listingHash msg.sender is the owner of.
  @param _amount      The number of ERC20 tokens to withdraw from the unstaked deposit.
  */
  function withdraw(bytes32 _listingHash, uint _amount) external {
    Listing storage listing = listings[_listingHash];

    require(listing.owner == msg.sender);
    require(_amount <= listing.unstakedDeposit);
    require(listing.unstakedDeposit.sub(_amount) >= parameterizer.get("minDeposit"));

    listing.unstakedDeposit = listing.unstakedDeposit.sub(_amount);
    require(token.transfer(msg.sender, _amount));

    emit _Withdrawal(
      _listingHash,
      _amount,
      listing.unstakedDeposit,
      msg.sender
    );
  }

  /**
  @dev                Allows the owner of a listingHash to remove the listingHash from the whitelist
                      Returns all tokens to the owner of the listingHash
  @param _listingHash A listingHash msg.sender is the owner of.
  */
  function exit(bytes32 _listingHash) external {
    Listing storage listing = listings[_listingHash];

    require(msg.sender == listing.owner);
    require(isWhitelisted(_listingHash));

    // Cannot exit during ongoing challenge
    require(listing.challengeID == 0 || challenges[listing.challengeID].resolved);

    // Remove listingHash & return tokens
    resetListing(_listingHash);
    emit _ListingWithdrawn(_listingHash);
  }

  // -----------------------
  // TOKEN HOLDER INTERFACE:
  // -----------------------

  /**
  @dev                Starts a poll for a listingHash which is either in the apply stage or
                      already in the whitelist. Tokens are taken from the challenger and the
                      applicant's deposits are locked.
  @param _listingHash The listingHash being challenged, whether listed or in application
  @param _data        Extra data relevant to the challenge. Think IPFS hashes.
  */
  function challenge(bytes32 _listingHash, string _data) external returns (uint challengeID) {
    Listing storage listing = listings[_listingHash];
    uint minDeposit = parameterizer.get("minDeposit");

    // Listing must be in apply stage or already on the whitelist
    require(appWasMade(_listingHash) || listing.whitelisted);
    // Prevent multiple challenges
    require(listing.challengeID == 0 || challenges[listing.challengeID].resolved);

    if (listing.unstakedDeposit < minDeposit) {
      // Not enough tokens, listingHash auto-delisted
      resetListing(_listingHash);
      emit _TouchAndRemoved(_listingHash);
      return 0;
    }

    // Starts poll
    uint pollID = voting.startPoll(
      parameterizer.get("voteQuorum"),
      parameterizer.get("commitStageLen"),
      parameterizer.get("revealStageLen")
    );

    challenges[pollID] = Challenge({
      challenger: msg.sender,
      rewardPool: uint(100).sub(parameterizer.get("dispensationPct")).mul(minDeposit).div(100),
      stake: minDeposit,
      resolved: false,
      totalTokens: 0
    });

    // Updates listingHash to store most recent challenge
    listing.challengeID = pollID;

    // Locks tokens for listingHash during challenge
    listing.unstakedDeposit = listing.unstakedDeposit.sub(minDeposit);

    // Takes tokens from challenger
    require(token.transferFrom(msg.sender, this, minDeposit));

    uint commitEndDate;
    uint revealEndDate;
    (commitEndDate, revealEndDate,) = voting.pollMap(pollID);

    emit _Challenge(
      _listingHash,
      pollID,
      _data,
      commitEndDate,
      revealEndDate,
      msg.sender
    );

    return pollID;
  }

  /**
  @dev                Updates a listingHash's status from 'application' to 'listing' or resolves
                      a challenge if one exists.
  @param _listingHash The listingHash whose status is being updated
  */
  function updateStatus(bytes32 _listingHash) public {
    if (canBeWhitelisted(_listingHash)) {
      whitelistApplication(_listingHash);
    } else if (challengeCanBeResolved(_listingHash)) {
      resolveChallenge(_listingHash);
    } else {
      revert();
    }
  }

  // ----------------
  // TOKEN FUNCTIONS:
  // ----------------

  /**
  @dev                Called by a voter to claim their reward for each completed vote. Someone
                      must call updateStatus() before this can be called.
  @param _challengeID The PLCR pollID of the challenge a reward is being claimed for
  @param _salt        The salt of a voter's commit hash in the given poll
  */
  function claimReward(uint _challengeID, uint _salt) public {
    // Ensures the voter has not already claimed tokens and challenge results have been processed
    require(challenges[_challengeID].tokenClaims[msg.sender] == false);
    require(challenges[_challengeID].resolved == true);

    uint voterTokens = voting.getNumPassingTokens(msg.sender, _challengeID, _salt);
    uint reward = voterReward(msg.sender, _challengeID, _salt);

    // Subtracts the voter's information to preserve the participation ratios
    // of other voters compared to the remaining pool of rewards
    challenges[_challengeID].totalTokens = challenges[_challengeID].totalTokens.sub(voterTokens);
    challenges[_challengeID].rewardPool = challenges[_challengeID].rewardPool.sub(reward);

    // Ensures a voter cannot claim tokens again
    challenges[_challengeID].tokenClaims[msg.sender] = true;

    require(token.transfer(msg.sender, reward));

    emit _RewardClaimed(_challengeID, reward, msg.sender);
  }

  // --------
  // GETTERS:
  // --------

  /**
  @dev                Calculates the provided voter's token reward for the given poll.
  @param _voter       The address of the voter whose reward balance is to be returned
  @param _challengeID The pollID of the challenge a reward balance is being queried for
  @param _salt        The salt of the voter's commit hash in the given poll
  @return             The uint indicating the voter's reward
  */
  function voterReward(address _voter, uint _challengeID, uint _salt)
  public view returns (uint)
  {
    uint totalTokens = challenges[_challengeID].totalTokens;
    uint rewardPool = challenges[_challengeID].rewardPool;
    uint voterTokens = voting.getNumPassingTokens(_voter, _challengeID, _salt);
    return voterTokens.mul(rewardPool).div(totalTokens);
  }

  /**
  @dev                Determines whether the given listingHash be whitelisted.
  @param _listingHash The listingHash whose status is to be examined
  */
  function canBeWhitelisted(bytes32 _listingHash) view public returns (bool) {
    uint challengeID = listings[_listingHash].challengeID;

    // Ensures that the application was made,
    // the application period has ended,
    // the listingHash can be whitelisted,
    // and either: the challengeID == 0, or the challenge has been resolved.
    if (
      /* solium-disable-next-line */
      appWasMade(_listingHash) &&
      /* solium-disable-next-line */
      listings[_listingHash].applicationExpiry < now &&
      /* solium-disable-next-line */
      !isWhitelisted(_listingHash) &&
      (challengeID == 0 || challenges[challengeID].resolved == true)
    ) {return true;}

    return false;
  }

  /**
  @dev                Returns true if the provided listingHash is whitelisted
  @param _listingHash The listingHash whose status is to be examined
  */
  function isWhitelisted(bytes32 _listingHash) view public returns (bool whitelisted) {
    return listings[_listingHash].whitelisted;
  }

  /**
  @dev                Returns true if apply was called for this listingHash
  @param _listingHash The listingHash whose status is to be examined
  */
  function appWasMade(bytes32 _listingHash) view public returns (bool exists) {
    return listings[_listingHash].applicationExpiry > 0;
  }

  /**
  @dev                Returns true if the application/listingHash has an unresolved challenge
  @param _listingHash The listingHash whose status is to be examined
  */
  function challengeExists(bytes32 _listingHash) view public returns (bool) {
    uint challengeID = listings[_listingHash].challengeID;

    return (listings[_listingHash].challengeID > 0 && !challenges[challengeID].resolved);
  }

  /**
  @dev                Determines whether voting has concluded in a challenge for a given
                      listingHash. Throws if no challenge exists.
  @param _listingHash A listingHash with an unresolved challenge
  */
  function challengeCanBeResolved(bytes32 _listingHash) view public returns (bool) {
    uint challengeID = listings[_listingHash].challengeID;

    require(challengeExists(_listingHash));

    return voting.pollEnded(challengeID);
  }

  /**
  @dev                Determines the number of tokens awarded to the winning party in a challenge.
  @param _challengeID The challengeID to determine a reward for
  */
  function determineReward(uint _challengeID) public view returns (uint) {
    require(!challenges[_challengeID].resolved && voting.pollEnded(_challengeID));

    // Edge case, nobody voted, give all tokens to the challenger.
    if (voting.getTotalNumberOfTokensForWinningOption(_challengeID) == 0) {
      return uint(2).mul(challenges[_challengeID].stake);
    }

    return uint(2).mul(challenges[_challengeID].stake).sub(challenges[_challengeID].rewardPool);
  }

  /**
  @dev                Getter for Challenge tokenClaims mappings
  @param _challengeID The challengeID to query
  @param _voter       The voter whose claim status to query for the provided challengeID
  */
  function tokenClaims(uint _challengeID, address _voter) public view returns (bool) {
    return challenges[_challengeID].tokenClaims[_voter];
  }

  // ----------------
  // PRIVATE FUNCTIONS:
  // ----------------

  /**
  @dev                Determines the winner in a challenge. Rewards the winner tokens and
                      either whitelists or de-whitelists the listingHash.
  @param _listingHash A listingHash with a challenge that is to be resolved
  */
  function resolveChallenge(bytes32 _listingHash) private {
    uint challengeID = listings[_listingHash].challengeID;

    // Calculates the winner's reward,
    // which is: (winner's full stake) + (dispensationPct * loser's stake)
    uint reward = determineReward(challengeID);

    // Sets flag on challenge being processed
    challenges[challengeID].resolved = true;

    // Stores the total tokens used for voting by the winning side for reward purposes
    challenges[challengeID].totalTokens = voting.getTotalNumberOfTokensForWinningOption(challengeID);

    // Case: challenge failed
    if (voting.isPassed(challengeID)) {
      whitelistApplication(_listingHash);
      // Unlock stake so that it can be retrieved by the applicant
      listings[_listingHash].unstakedDeposit = listings[_listingHash].unstakedDeposit.add(reward);

      emit _ChallengeFailed(
        _listingHash,
        challengeID,
        challenges[challengeID].rewardPool,
        challenges[challengeID].totalTokens
      );
    } else { // Case: challenge succeeded or nobody voted
      resetListing(_listingHash);
      // Transfer the reward to the challenger
      require(token.transfer(challenges[challengeID].challenger, reward));

      emit _ChallengeSucceeded(
        _listingHash,
        challengeID,
        challenges[challengeID].rewardPool,
        challenges[challengeID].totalTokens
      );
    }
  }

  /**
  @dev                Called by updateStatus() if the applicationExpiry date passed without a
                      challenge being made. Called by resolveChallenge() if an
                      application/listing beat a challenge.
  @param _listingHash The listingHash of an application/listingHash to be whitelisted
  */
  function whitelistApplication(bytes32 _listingHash) private {
    if (!listings[_listingHash].whitelisted) {emit _ApplicationWhitelisted(_listingHash);}
    listings[_listingHash].whitelisted = true;
  }

  /**
  @dev                Deletes a listingHash from the whitelist and transfers tokens back to owner
  @param _listingHash The listing hash to delete
  */
  function resetListing(bytes32 _listingHash) private {
    Listing storage listing = listings[_listingHash];

    // Emit events before deleting listing to check whether is whitelisted
    if (listing.whitelisted) {
      emit _ListingRemoved(_listingHash);
    } else {
      emit _ApplicationRemoved(_listingHash);
    }

    // Deleting listing to prevent reentry
    address owner = listing.owner;
    uint unstakedDeposit = listing.unstakedDeposit;
    delete listings[_listingHash];

    // Transfers any remaining balance back to the owner
    if (unstakedDeposit > 0) {
      require(token.transfer(owner, unstakedDeposit));
    }
  }
}
