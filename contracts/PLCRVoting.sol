pragma solidity 0.4.24;

import "./ERC20.sol";
import "./DLL.sol";
import "./AttributeStore.sol";
import "./SafeMath.sol";


/**
@title Partial-Lock-Commit-Reveal Voting scheme with ERC20 tokens
@author Team: Aspyn Palatnick, Cem Ozer, Yorke Rhodes
*/
contract PLCRVoting {

  // ============
  // EVENTS:
  // ============

  event _VoteCommitted(uint indexed id, uint numTokens, address indexed voter);
  event _VoteRevealed(uint indexed id, uint numTokens, uint votesFor, uint votesAgainst, uint indexed choice, address indexed voter);
  event _PollCreated(uint voteQuorum, uint commitEndDate, uint revealEndDate, uint indexed id, address indexed creator);
  event _VotingRightsGranted(uint numTokens, address indexed voter);
  event _VotingRightsWithdrawn(uint numTokens, address indexed voter);
  event _TokensRescued(uint indexed id, address indexed voter);

  // ============
  // DATA STRUCTURES:
  // ============

  using AttributeStore for AttributeStore.Data;
  using DLL for DLL.Data;
  using SafeMath for uint;

  struct Poll {
    uint commitEndDate;     /// expiration date of commit period for poll
    uint revealEndDate;     /// expiration date of reveal period for poll
    uint voteQuorum;    /// number of votes required for a proposal to pass
    uint votesFor;    /// tally of votes supporting proposal
    uint votesAgainst;      /// tally of votes countering proposal
    mapping(address => bool) didCommit;  /// indicates whether an address committed a vote for this poll
    mapping(address => bool) didReveal;   /// indicates whether an address revealed a vote for this poll
  }

  // ============
  // STATE VARIABLES:
  // ============

  uint constant public INITIAL_POLL_NONCE = 0;
  uint public pollNonce;

  mapping(uint => Poll) public pollMap; // maps id to Poll struct
  mapping(address => uint) public voteTokenBalance; // maps user's address to voteToken balance

  mapping(address => DLL.Data) dllMap;
  AttributeStore.Data store;

  ERC20 public token;

  // ============
  // CONSTRUCTOR:
  // ============

  /**
  @dev Initializes voteQuorum, commitDuration, revealDuration, and pollNonce in addition to token contract and trusted mapping
  @param _tokenAddr The address where the ERC20 token contract is deployed
  */
  constructor(address _tokenAddr) public {
    token = ERC20(_tokenAddr);
    pollNonce = INITIAL_POLL_NONCE;
  }

  // ================
  // TOKEN INTERFACE:
  // ================

  /**
  @notice Loads _numTokens ERC20 tokens into the voting contract for one-to-one voting rights
  @dev Assumes that msg.sender has approved voting contract to spend on their behalf
  @param _numTokens The number of votingTokens desired in exchange for ERC20 tokens
  */
  function requestVotingRights(uint _numTokens) external {
    require(token.balanceOf(msg.sender) >= _numTokens);
    voteTokenBalance[msg.sender] = voteTokenBalance[msg.sender].add(_numTokens);
    require(token.transferFrom(msg.sender, this, _numTokens));
    emit _VotingRightsGranted(_numTokens, msg.sender);
  }

  /**
  @notice Withdraw _numTokens ERC20 tokens from the voting contract, revoking these voting rights
  @param _numTokens The number of ERC20 tokens desired in exchange for voting rights
  */
  function withdrawVotingRights(uint _numTokens) external {
    uint availableTokens = voteTokenBalance[msg.sender].sub(getLockedTokens(msg.sender));
    require(availableTokens >= _numTokens);
    voteTokenBalance[msg.sender] = voteTokenBalance[msg.sender].sub(_numTokens);
    require(token.transfer(msg.sender, _numTokens));
    emit _VotingRightsWithdrawn(_numTokens, msg.sender);
  }

  /**
  @dev Unlocks tokens locked in unrevealed vote where poll has ended
  @param _id Integer identifier associated with the target poll
  */
  function rescueTokens(uint _id) external {
    require(isExpired(pollMap[_id].revealEndDate));
    require(dllMap[msg.sender].contains(_id));

    dllMap[msg.sender].remove(_id);
    emit _TokensRescued(_id, msg.sender);
  }

  // =================
  // VOTING INTERFACE:
  // =================

  /**
  @notice Commits vote using hash of choice and secret salt to conceal vote until reveal
  @param _id Integer identifier associated with target poll
  @param _secretHash Commit keccak256 hash of voter's choice and salt (tightly packed in this order)
  @param _numTokens The number of tokens to be committed towards the target poll
  @param _prevID The ID of the poll that the user has voted the maximum number of tokens in which is still less than or equal to numTokens
  */
  function commitVote(
    uint _id,
    bytes32 _secretHash,
    uint _numTokens,
    uint _prevID
  ) external
  {
    require(commitPeriodActive(_id));
    require(voteTokenBalance[msg.sender] >= _numTokens); // prevent user from overspending
    require(_id != 0);                // prevent user from committing to zero node placeholder

    // Check if _prevID exists in the user's DLL or if _prevID is 0
    require(_prevID == 0 || dllMap[msg.sender].contains(_prevID));

    uint nextID = dllMap[msg.sender].getNext(_prevID);

    // if nextID is equal to _id, _id is being updated,
    nextID = (nextID == _id) ? dllMap[msg.sender].getNext(_id) : nextID;

    require(
      validPosition(
        _prevID,
        nextID,
        msg.sender,
        _numTokens
      )
    );

    dllMap[msg.sender].insert(_prevID, _id, nextID);

    bytes32 UUID = attrUUID(msg.sender, _id);

    store.setAttribute(UUID, "numTokens", _numTokens);
    store.setAttribute(UUID, "commitHash", uint(_secretHash));

    pollMap[_id].didCommit[msg.sender] = true;
    emit _VoteCommitted(_id, _numTokens, msg.sender);
  }

  /**
  @dev Compares previous and next poll's committed tokens for sorting purposes
  @param _prevID Integer identifier associated with previous poll in sorted order
  @param _nextID Integer identifier associated with next poll in sorted order
  @param _voter Address of user to check DLL position for
  @param _numTokens The number of tokens to be committed towards the poll (used for sorting)
  @return valid Boolean indication of if the specified position maintains the sort
  */
  function validPosition(
    uint _prevID,
    uint _nextID,
    address _voter,
    uint _numTokens
  ) public view returns (bool valid)
  {
    bool prevValid = (_numTokens >= getNumTokens(_voter, _prevID));
    // if next is zero node, _numTokens does not need to be greater
    bool nextValid = (_numTokens <= getNumTokens(_voter, _nextID) || _nextID == 0);
    return prevValid && nextValid;
  }

  /**
  @notice Reveals vote with choice and secret salt used in generating commitHash to attribute committed tokens
  @param _id Integer identifier associated with target poll
  @param _voteOption Vote choice used to generate commitHash for associated poll
  @param _salt Secret number used to generate commitHash for associated poll
  */
  function revealVote(uint _id, uint _voteOption, uint _salt) external {
    // Make sure the reveal period is active
    require(revealPeriodActive(_id));
    require(pollMap[_id].didCommit[msg.sender]);                         // make sure user has committed a vote for this poll
    require(!pollMap[_id].didReveal[msg.sender]);                        // prevent user from revealing multiple times
    require(keccak256(abi.encodePacked(_voteOption, _salt)) == getCommitHash(msg.sender, _id)); // compare resultant hash from inputs to original commitHash

    uint numTokens = getNumTokens(msg.sender, _id);

    if (_voteOption == 1) {// apply numTokens to appropriate poll choice
      pollMap[_id].votesFor = pollMap[_id].votesFor.add(numTokens);
    } else {
      pollMap[_id].votesAgainst = pollMap[_id].votesAgainst.add(numTokens);
    }

    dllMap[msg.sender].remove(_id); // remove the node referring to this vote upon reveal
    pollMap[_id].didReveal[msg.sender] = true;

    emit _VoteRevealed(
      _id,
      numTokens,
      pollMap[_id].votesFor,
      pollMap[_id].votesAgainst,
      _voteOption, msg.sender
    );
  }

  /**
  @param _id Integer identifier associated with target poll
  @param _salt Arbitrarily chosen integer used to generate secretHash
  @return correctVotes Number of tokens voted for winning option
  */
  function getNumPassingTokens(address _voter, uint _id, uint _salt) public view returns (uint correctVotes) {
    require(pollEnded(_id));
    require(pollMap[_id].didReveal[_voter]);

    uint winningChoice = isPassed(_id) ? 1 : 0;
    bytes32 winnerHash = keccak256(abi.encodePacked(winningChoice, _salt));
    bytes32 commitHash = getCommitHash(_voter, _id);

    require(winnerHash == commitHash);

    return getNumTokens(_voter, _id);
  }

  // ==================
  // POLLING INTERFACE:
  // ==================

  /**
  @dev Initiates a poll with canonical configured parameters at id emitted by PollCreated event
  @param _voteQuorum Type of majority (out of 100) that is necessary for poll to be successful
  @param _commitDuration Length of desired commit period in seconds
  @param _revealDuration Length of desired reveal period in seconds
  */
  function startPoll(uint _voteQuorum, uint _commitDuration, uint _revealDuration) public returns (uint id) {
    pollNonce = pollNonce.add(1);

    uint commitEndDate = block.timestamp.add(_commitDuration);
    uint revealEndDate = commitEndDate.add(_revealDuration);

    pollMap[pollNonce] = Poll({
      voteQuorum: _voteQuorum,
      commitEndDate: commitEndDate,
      revealEndDate: revealEndDate,
      votesFor: 0,
      votesAgainst: 0
    });

    emit _PollCreated(
      _voteQuorum,
      commitEndDate,
      revealEndDate,
      pollNonce,
      msg.sender
    );

    return pollNonce;
  }

  /**
  @notice Determines if proposal has passed
  @dev Check if votesFor out of totalVotes exceeds votesQuorum (requires pollEnded)
  @param _id Integer identifier associated with target poll
  */
  function isPassed(uint _id) view public returns (bool passed) {
    require(pollEnded(_id));

    Poll memory poll = pollMap[_id];
    return (poll.votesFor.mul(100)) > (poll.voteQuorum.mul(poll.votesFor.add(poll.votesAgainst)));
  }

  // ----------------
  // POLLING HELPERS:
  // ----------------

  /**
  @dev Gets the total winning votes for reward distribution purposes
  @param _id Integer identifier associated with target poll
  @return Total number of votes committed to the winning option for specified poll
  */
  function getTotalNumberOfTokensForWinningOption(uint _id) view public returns (uint numTokens) {
    require(pollEnded(_id));

    if (isPassed(_id))
      return pollMap[_id].votesFor;
    else
      return pollMap[_id].votesAgainst;
  }

  /**
  @notice Determines if poll is over
  @dev Checks isExpired for specified poll's revealEndDate
  @return Boolean indication of whether polling period is over
  */
  function pollEnded(uint _id) view public returns (bool ended) {
    require(pollExists(_id));

    return isExpired(pollMap[_id].revealEndDate);
  }

  /**
  @notice Checks if the commit period is still active for the specified poll
  @dev Checks isExpired for the specified poll's commitEndDate
  @param _id Integer identifier associated with target poll
  @return Boolean indication of isCommitPeriodActive for target poll
  */
  function commitPeriodActive(uint _id) view public returns (bool active) {
    require(pollExists(_id));

    return !isExpired(pollMap[_id].commitEndDate);
  }

  /**
  @notice Checks if the reveal period is still active for the specified poll
  @dev Checks isExpired for the specified poll's revealEndDate
  @param _id Integer identifier associated with target poll
  */
  function revealPeriodActive(uint _id) view public returns (bool active) {
    require(pollExists(_id));

    return !isExpired(pollMap[_id].revealEndDate) && !commitPeriodActive(_id);
  }

  /**
  @dev Checks if user has committed for specified poll
  @param _voter Address of user to check against
  @param _id Integer identifier associated with target poll
  @return Boolean indication of whether user has committed
  */
  function didCommit(address _voter, uint _id) view public returns (bool committed) {
    require(pollExists(_id));

    return pollMap[_id].didCommit[_voter];
  }

  /**
  @dev Checks if user has revealed for specified poll
  @param _voter Address of user to check against
  @param _id Integer identifier associated with target poll
  @return Boolean indication of whether user has revealed
  */
  function didReveal(address _voter, uint _id) view public returns (bool revealed) {
    require(pollExists(_id));

    return pollMap[_id].didReveal[_voter];
  }

  /**
  @dev Checks if a poll exists
  @param _id The id whose existance is to be evaluated.
  @return Boolean Indicates whether a poll exists for the provided id
  */
  function pollExists(uint _id) view public returns (bool exists) {
    return (_id != 0 && _id <= pollNonce);
  }

  // ---------------------------
  // DOUBLE-LINKED-LIST HELPERS:
  // ---------------------------

  /**
  @dev Gets the bytes32 commitHash property of target poll
  @param _voter Address of user to check against
  @param _id Integer identifier associated with target poll
  @return Bytes32 hash property attached to target poll
  */
  function getCommitHash(address _voter, uint _id) view public returns (bytes32 commitHash) {
    return bytes32(store.getAttribute(attrUUID(_voter, _id), "commitHash"));
  }

  /**
  @dev Wrapper for getAttribute with attrName="numTokens"
  @param _voter Address of user to check against
  @param _id Integer identifier associated with target poll
  @return Number of tokens committed to poll in sorted poll-linked-list
  */
  function getNumTokens(address _voter, uint _id) view public returns (uint numTokens) {
    return store.getAttribute(attrUUID(_voter, _id), "numTokens");
  }

  /**
  @dev Gets top element of sorted poll-linked-list
  @param _voter Address of user to check against
  @return Integer identifier to poll with maximum number of tokens committed to it
  */
  function getLastNode(address _voter) view public returns (uint id) {
    return dllMap[_voter].getPrev(0);
  }

  /**
  @dev Gets the numTokens property of getLastNode
  @param _voter Address of user to check against
  @return Maximum number of tokens committed in poll specified
  */
  function getLockedTokens(address _voter) view public returns (uint numTokens) {
    return getNumTokens(_voter, getLastNode(_voter));
  }

  /*
  @dev Takes the last node in the user's DLL and iterates backwards through the list searching
  for a node with a value less than or equal to the provided _numTokens value. When such a node
  is found, if the provided _id matches the found nodeID, this operation is an in-place
  update. In that case, return the previous node of the node being updated. Otherwise return the
  first node that was found with a value less than or equal to the provided _numTokens.
  @param _voter The voter whose DLL will be searched
  @param _numTokens The value for the numTokens attribute in the node to be inserted
  @return the node which the propoded node should be inserted after
  */
  function getInsertPointForNumTokens(address _voter, uint _numTokens, uint _id)
  view public returns (uint prevNode)
  {
    // Get the last node in the list and the number of tokens in that node
    uint nodeID = getLastNode(_voter);
    uint tokensInNode = getNumTokens(_voter, nodeID);

    // Iterate backwards through the list until reaching the root node
    while (nodeID != 0) {
      // Get the number of tokens in the current node
      tokensInNode = getNumTokens(_voter, nodeID);
      if (tokensInNode <= _numTokens) { // We found the insert point!
        if (nodeID == _id) {
          // This is an in-place update. Return the prev node of the node being updated
          nodeID = dllMap[_voter].getPrev(nodeID);
        }
        // Return the insert point
        return nodeID;
      }
      // We did not find the insert point. Continue iterating backwards through the list
      nodeID = dllMap[_voter].getPrev(nodeID);
    }

    // The list is empty, or a smaller value than anything else in the list is being inserted
    return nodeID;
  }

  // ----------------
  // GENERAL HELPERS:
  // ----------------

  /**
  @dev Checks if an expiration date has been reached
  @param _terminationDate Integer timestamp of date to compare current timestamp with
  @return expired Boolean indication of whether the terminationDate has passed
  */
  function isExpired(uint _terminationDate) view public returns (bool expired) {
    return (block.timestamp > _terminationDate);
  }

  /**
  @dev Generates an identifier which associates a user and a poll together
  @param _id Integer identifier associated with target poll
  @return UUID Hash which is deterministic from _user and _id
  */
  function attrUUID(address _user, uint _id) public pure returns (bytes32 UUID) {
    return keccak256(abi.encodePacked(_user, _id));
  }
}
