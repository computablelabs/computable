# @title Computable Parameterizer
# @notice A simplified parameterizer, hard-coded to our protocol.
# @author Computable

# We refer to any 'params' as integers to avoid conversions to and from strings
PRICE_FLOOR: constant(uint256) = 2
SPREAD: constant(uint256) = 4
LIST_REWARD: constant(uint256) = 5
STAKE: constant(uint256) = 1
VOTE_BY: constant(uint256) = 7
PLURALITY: constant(uint256) = 6
# _PAYMENTs are % based so the sum of both should be < 100 as there is an implicit reserve_payment
BACKEND_PAYMENT: constant(uint256) = 8
MAKER_PAYMENT: constant(uint256) = 9
COST_PER_BYTE: constant(uint256) = 11

# The sole Candidate 'kind' known to the Parameterizer
REPARAM: constant(uint256) = 3

# The number of seconds in a day
SECONDS_IN_DAY: constant(uint256) = 86400 

struct Reparam:
  param: uint256
  value: uint256

# Parameterizer has access to the MarketToken contract in order to query
# for the total number of market tokens
contract MarketToken:
  def totalSupply() -> uint256(wei): constant

# Parameterizer has access to the Voting contract, being recognized by it as privileged
contract Voting:
  def candidateIs(hash: bytes32, kind: uint256) -> bool: constant
  def isCandidate(hash: bytes32) -> bool: constant
  def addCandidate(hash: bytes32, kind: uint256, owner: address, stake: uint256(wei), vote_by: uint256(sec)): modifying
  def removeCandidate(hash: bytes32): modifying
  def didPass(hash: bytes32, plurality: uint256) -> bool: constant
  def pollClosed(hash: bytes32) -> bool: constant

ReparamProposed: event({owner: indexed(address), hash: indexed(bytes32), param: indexed(uint256), value: uint256})
ReparamFailed: event({hash: indexed(bytes32), param: indexed(uint256), value: uint256})
ReparamSucceeded: event({hash: indexed(bytes32), param: indexed(uint256), value: uint256})

reparams: map(bytes32, Reparam)
price_floor: wei_value
spread: uint256
list_reward: wei_value
stake: wei_value
vote_by: timedelta
plurality: uint256
backend_payment: uint256
maker_payment: uint256
cost_per_byte: wei_value
market_token: MarketToken 
voting: Voting

@public
def __init__(mkt_addr: address, v_addr: address, pr_fl: wei_value, spd: uint256, list_re: wei_value, stk: wei_value,
  vote_by_d: timedelta, pl: uint256, back_p: uint256, maker_p: uint256, cost: wei_value):
    self.market_token = MarketToken(mkt_addr)
    self.voting = Voting(v_addr)
    self.price_floor = pr_fl
    self.spread = spd
    self.list_reward = list_re
    self.stake = stk
    self.vote_by = vote_by_d
    self.plurality = pl
    self.backend_payment = back_p
    self.maker_payment = maker_p
    self.cost_per_byte = cost


@public
@constant
def getBackendPayment() -> uint256:
  """
  @notice Returns the percentage of a delivery cost given to the backend upon delivering
  """
  return self.backend_payment


@public
@constant
def getMakerPayment() -> uint256:
  """
  @notice Returns the percentage of a delivery cost given to a maker per their list usage
  """
  return self.maker_payment


@public
@constant
def getReservePayment() -> uint256:
  """
  @notice Returns the percentage of a delivery cost given to the reserve upon purchase
  @dev As you can see we dont require an explicit Reserve amount as we can simply use this delta.
  """
  return 100 - (self.backend_payment + self.maker_payment)


@public
@constant
def getCostPerByte() -> wei_value:
  """
  @notice Return the amount, in wei, to be paid per output byte of listing data
  """
  return self.cost_per_byte


@public
@constant
def getStake() -> wei_value:
  """
  @notice Return the current voting/challenge stake value in wei
  """
  return self.stake


@public
@constant
def getPriceFloor() -> wei_value:
  """
  @notice Return the current Ethertoken to MarketToken price floor in wei
  """
  return self.price_floor


@public
@constant
def getHash(param: uint256, value: uint256) -> bytes32:
  """
  @notice Given the param: value keypair - generate a hash
  @dev Returns the same value we use internally for constructing mapping keys
  @param param the integer representing a parameterizer member
  @param value what to set it to
  @return The generated hash
  """
  return keccak256(concat(convert(param, bytes32), convert(value, bytes32)))


@public
@constant
def getSpread() -> uint256:
  """
  @notice Return the current spread scaling factor
  """
  return self.spread


@public
@constant
def getListReward() -> wei_value:
  """
  @notice Return the current amount of market token, in wei, awarded to a new listing
  """
  return self.list_reward


@public
@constant
def getPlurality() -> uint256:
  """
  @notice Return the percent of 100 needed by a candidate to 'pass' in a poll
  """
  return self.plurality


@public
@constant
def getReparam(hash: bytes32) -> (uint256, uint256):
  """
  @notice Return the data about the given Reparam
  @param hash The Reparam identifier
  """
  return (self.reparams[hash].param, self.reparams[hash].value)


@public
@constant
def getVoteBy() -> timedelta:
  """
  @notice Return the number of seconds, from when begun, that a candidate poll should close
  """
  return self.vote_by


@public
def reparameterize(param: uint256, value: uint256):
  """
  @notice Suggest a change to a Parameterizer attribute, creating a candidate for it
  @dev Sender must not be a matching candidate already open
  @param param The attribute to change
  @param value What to change it to
  """
  # hashed identifier made up of the prop and its proposed value
  hash: bytes32 = keccak256(concat(convert(param, bytes32), convert(value, bytes32)))
  assert not self.voting.isCandidate(hash)
  self.reparams[hash] = Reparam({param: param, value:value})
  self.voting.addCandidate(hash, REPARAM, msg.sender, self.stake, self.vote_by)
  log.ReparamProposed(msg.sender, hash, param, value)


@public
def resolveReparam(hash: bytes32):
  """
  @notice Determine if a Reparam Candidate collected enough votes to pass, setting it if so
  @dev This method enforces that the candidate is of the correct type and its poll is closed
  @param hash The Reparam identifier
  """
  assert self.voting.candidateIs(hash, REPARAM)
  assert self.voting.pollClosed(hash)
  # ascertain which param,value we are looking at
  param: uint256 = self.reparams[hash].param
  value: uint256 = self.reparams[hash].value
  if self.voting.didPass(hash, self.plurality):
    #TODO in time we can likely tell an optimal order for these...
    if param == PRICE_FLOOR:
      self.price_floor = value
    elif param == SPREAD:
      # The spread is a percentage >= 100%
      assert value >= 100 
      self.spread = value
    elif param == LIST_REWARD:
      self.list_reward = value
    elif param == STAKE:
      # The stake can't be zero
      assert value > 0
      # The stake can't exceed 1/3 of MarketToken supply
      assert value <= (self.market_token.totalSupply()/3) 
      self.stake = value
    elif param == VOTE_BY:
      assert value >= SECONDS_IN_DAY
      assert value <= 2 * (7 * SECONDS_IN_DAY)
      self.vote_by = value
    elif param == PLURALITY:
      # Plurality can be at most 100
      assert value <= 100
      self.plurality = value
    # Backend percent + Maker percent should be <= 100
    elif param == MAKER_PAYMENT:
      assert (value + self.backend_payment) <= 100
      self.maker_payment = value
    elif param == BACKEND_PAYMENT:
      assert (value + self.maker_payment) <= 100
      self.backend_payment = value
    elif param == COST_PER_BYTE:
      self.cost_per_byte = value
    log.ReparamSucceeded(hash, param, value)
  else: # did not get enough votes...
    log.ReparamFailed(hash, param, value)
  # regardless, cleanup the reparam and candidate
  self.voting.removeCandidate(hash)
  # TODO make sure this works as expected
  clear(self.reparams[hash])
