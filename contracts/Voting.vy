# @title Computable Voting
# @notice A simplified, weightless, public Voting contract
# @author Computable

# constants
CHALLENGE: constant(uint256) = 2 # candidate.kind that voting knows about

struct Candidate:
  kind: uint256 # one of [1,2,3,4] representing an application, challenge, reparam or registration respectively
  owner: address
  stake: wei_value
  vote_by: timestamp
  yea: uint256
  nay: uint256

# external contracts
contract MarketToken:
  def transfer(to: address, amount: uint256(wei)) -> bool: modifying
  def transferFrom(source: address, to: address, amount: uint256(wei)) -> bool: modifying

CandidateAdded: event({hash: indexed(bytes32), kind: indexed(uint256), owner: indexed(address), voteBy: timestamp})
CandidateRemoved: event({hash: indexed(bytes32)})
Voted: event({hash: indexed(bytes32), voter: indexed(address)})

candidates: map(bytes32, Candidate)
stakes: map(address, map(bytes32, wei_value)) # user -> candidate -> $
market_token: MarketToken
parameterizer_address: address
datatrust_address: address
listing_address: address
reserve_address: address
owner_address: address

@public
def __init__(market_token_addr: address):
  self.owner_address = msg.sender
  self.market_token = MarketToken(market_token_addr)


@public
@constant
def getPrivileged() -> (address, address, address, address):
  """
  @notice Fetch a list of each privileged address recognized by this contract
  @return privileged addresses
  """
  return (self.parameterizer_address, self.datatrust_address,
    self.listing_address, self.reserve_address)


@public
def setPrivileged(parameterizer: address, datatrust: address, listing: address, reserve: address):
  """
  @notice Allow the Market owner to set privileged contract addresses
  """
  assert msg.sender == self.owner_address
  self.parameterizer_address = parameterizer
  self.datatrust_address = datatrust
  self.listing_address = listing
  self.reserve_address = reserve


@public
@constant
def hasPrivilege(sender: address) -> bool:
  """
  @notice Return a bool indicating whether the given address is a member of this contracts privileged group
  @return bool
  """
  return (sender == self.parameterizer_address or sender == self.datatrust_address
    or sender == self.listing_address or sender == self.reserve_address)


@public
@constant
def candidateIs(hash: bytes32, kind: uint256) -> bool:
  """
  @notice Is a given candidate the given kind?
  @param hash The candidate to check
  @param kind The kind to compare
  @return bool
  """
  return self.candidates[hash].kind == kind


@public
@constant
def isCandidate(hash: bytes32) -> bool:
  """
  @notice Check if a given hash identifier is a current candidate
  @return bool
  """
  return self.candidates[hash].owner != ZERO_ADDRESS


@public
@constant
def getCandidate(hash: bytes32) -> (uint256, address, wei_value, timestamp, uint256, uint256):
  """
  @notice Return information about the given candidate identified by the given hash
  @dev Hash argument keys a candidate struct in the candidates mapping
  @return The type, vote_by timestamp and number of votes recieved
  """
  return (self.candidates[hash].kind, self.candidates[hash].owner, self.candidates[hash].stake,
    self.candidates[hash].vote_by, self.candidates[hash].yea, self.candidates[hash].nay)


@public
@constant
def getCandidateOwner(hash: bytes32) -> address:
  """
  @notice return the owner address of a given Candidate.
  """
  return self.candidates[hash].owner


@public
def addCandidate(hash: bytes32, kind: uint256, owner: address, stake: wei_value, vote_by: timedelta):
  """
  @notice Given a listing or parameter hash, create a new voting candidate
  @dev Only priveliged contracts may call this method
  @param hash The identifier for the listing or reparameterization candidate
  @param kind The type of candidate we are adding
  @param owner The adress which owns this created candidate
  @param stake How much, in wei, must be staked to vote or challenge
  @param vote_by How long into the future until polls for this candidate close
  """
  assert self.hasPrivilege(msg.sender)
  assert self.candidates[hash].owner == ZERO_ADDRESS
  if kind == CHALLENGE: # a challenger must successfully stake a challenge
    self.market_token.transferFrom(owner, self, stake)
    self.stakes[owner][hash] = stake
  end: timestamp = block.timestamp + vote_by
  self.candidates[hash].kind = kind
  self.candidates[hash].owner = owner
  self.candidates[hash].stake = stake
  self.candidates[hash].vote_by = end
  log.CandidateAdded(hash, kind, owner, end)


@public
def removeCandidate(hash: bytes32):
  """
  @notice Remove a candidate from the current list
  @dev Clears all members from a Candidate pointed to by a hash (enabling re-use)
  """
  assert self.hasPrivilege(msg.sender)
  assert self.candidates[hash].owner != ZERO_ADDRESS
  clear(self.candidates[hash]) # TODO assure this works vs individually setting to 0
  log.CandidateRemoved(hash)


@public
@constant
def didPass(hash: bytes32, plurality: uint256) -> bool:
  """
  @notice Return a bool indicating whether a given candidate recieved enough votes to exceed the plurality
  @dev The poll must be closed. Also we cover the corner case that no one voted.
  @return bool
  """
  assert self.candidates[hash].owner != ZERO_ADDRESS
  assert self.candidates[hash].vote_by < block.timestamp
  yea: uint256 = self.candidates[hash].yea
  total: uint256 = yea + self.candidates[hash].nay
  # edge case that no one voted
  if total == 0:
    # theoretically a market could have a 0 plurality
    return plurality == 0
  else:
    return ((yea * 100) / total) >= plurality


@public
@constant
def pollClosed(hash: bytes32) -> bool:
  """
  @notice Check to see if a given candidate's polling has closed
  @return bool
  """
  assert self.candidates[hash].owner != ZERO_ADDRESS
  return self.candidates[hash].vote_by < block.timestamp


@public
def vote(hash: bytes32, option: uint256):
  """
  @notice Cast a vote for a given candidate
  @dev User mush have approved market token to spend on their behalf
  @param hash The candidate identifier
  @param option Yea (1) or Nay (!1)
  """
  assert self.candidates[hash].owner != ZERO_ADDRESS
  assert self.candidates[hash].vote_by > block.timestamp
  stake: wei_value = self.candidates[hash].stake
  self.market_token.transferFrom(msg.sender, self, stake)
  self.stakes[msg.sender][hash] += stake
  if option == 1:
    self.candidates[hash].yea += 1
  else:
    self.candidates[hash].nay += 1
  log.Voted(hash, msg.sender)


@public
def transferStake(hash: bytes32, addr: address):
  """
  @notice The stakes belonging to one address are being credited to another due to a lost challenge.
  @param hash The Candidate identifier
  @param addr The Address recieving the credit
  """
  assert msg.sender == self.listing_address # only the listing contract will call this
  staked: wei_value = self.stakes[self.candidates[hash].owner][hash]
  clear(self.stakes[self.candidates[hash].owner][hash])
  self.stakes[addr][hash] = staked


@public
@constant
def getStake(hash: bytes32, addr: address) -> wei_value:
  """
  @notice Return the balance staked by a given user for a given candidate
  @param hash The candidate identifier
  @param addr Address that did the staking
  """
  return self.stakes[addr][hash]


@public
def unstake(hash: bytes32):
  """
  @notice Allow a user to (re)claim stakes that they have rights to
  @dev The candidate in question must have been resolved
  NOTE: There is an edge case where a candidate is resolved and another opens with the same name
  before unstaking. In that case a user simply waits until that poll is over.
  @param hash The Candidate identifier
  """
  assert self.candidates[hash].owner == ZERO_ADDRESS # must be resolved (have been removed)
  staked: wei_value = self.stakes[msg.sender][hash]
  clear(self.stakes[msg.sender][hash])
  self.market_token.transfer(msg.sender, staked)
