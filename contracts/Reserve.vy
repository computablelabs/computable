# @title Computable Reserve
# @notice Handle the details pertaining to the Reserve of a Computable Market
# @author Computable

# external contracts
contract EtherToken:
  def balanceOf(owner: address) -> uint256(wei): constant
  def transfer(to: address, amount: uint256(wei)) -> bool: modifying
  def transferFrom(source: address, to: address, amount: uint256(wei)) -> bool: modifying

contract MarketToken:
  def balanceOf(owner: address) -> uint256(wei): constant
  def burnAll(owner: address): modifying
  def mint(amount: uint256(wei)): modifying
  def totalSupply() -> uint256(wei): constant
  def transfer(to: address, amount: uint256(wei)) -> bool: modifying

contract Parameterizer:
  def getPriceFloor() -> uint256(wei): constant
  def getSpread() -> uint256: constant

# events
Withdrawn: event({owner: indexed(address), transferred: wei_value})
Supported: event({owner: indexed(address), offered: wei_value, minted: wei_value})

# state vars
ether_token: EtherToken
market_token: MarketToken
parameterizer: Parameterizer

@public
def __init__(ether_token_addr: address, market_token_addr: address, p11r_addr: address):
    self.ether_token = EtherToken(ether_token_addr)
    self.market_token = MarketToken(market_token_addr)
    self.parameterizer = Parameterizer(p11r_addr)


@public
@constant
def getSupportPrice() -> wei_value:
  """
  @notice Return the amount of Ether token (in wei) needed to purchase one billionth of a Market token
  """
  price_floor: wei_value = self.parameterizer.getPriceFloor()
  spread: uint256 = self.parameterizer.getSpread()
  reserve: wei_value = self.ether_token.balanceOf(self)
  total: wei_value = self.market_token.totalSupply()
  if total < 1000000000000000000: # that is, is total supply less than one token in wei
    return price_floor + ((spread * reserve * 1000000000) / (100 * 1000000000000000000))
  else:
    return price_floor + ((spread * reserve * 1000000000) / (100 * total)) # NOTE the multiplier ONE_GWEI


@public
def support(offer: wei_value):
  """
  @notice Allow the purchase MarketToken with EtherToken priced according to the "buy-curve"
  @param offer An amount of Ether Token in Wei
  """
  price: wei_value = self.getSupportPrice()
  assert offer >= price # you cannot buy less than one billionth of a market token
  self.ether_token.transferFrom(msg.sender, self, offer)
  minted: uint256 = (offer / price) * 1000000000 # NOTE the ONE_GWEI multiplier here as well
  self.market_token.mint(minted) # TODO maybe implement `mintFor()`
  self.market_token.transfer(msg.sender, minted)
  log.Supported(msg.sender, offer, minted)


@public
@constant
def getWithdrawalProceeds(addr: address) -> wei_value:
  """
  @notice Return the amount of Ether Token funds a supporter would recieve for withdrawal
  @return Amount in Ether Token Wei
  """
  assert addr != ZERO_ADDRESS
  bal: wei_value = self.market_token.balanceOf(addr)
  reserve: wei_value = self.ether_token.balanceOf(self)
  total: wei_value = self.market_token.totalSupply()
  return (bal * reserve) / total


@public
def withdraw():
  """
  @notice Allows a supporter to exit the market. Burning any market token owned and
  withdrawing their share of the reserve.
  @dev Supporter, if owning a challenge, may want to wait until that is over (in case they win)
  """
  withdrawn: wei_value = self.getWithdrawalProceeds(msg.sender)
  assert withdrawn > 0
  # before any transfer, burn their market tokens...
  self.market_token.burnAll(msg.sender)
  self.ether_token.transfer(msg.sender, withdrawn)
  log.Withdrawn(msg.sender, withdrawn)
