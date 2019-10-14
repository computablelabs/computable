# @title Computable MarketToken
# @notice Implementation of an ERC20 compatible token which is both burnable and mintable
# @author Computable
# @dev Implements the ERC20 interface

Approval: event({owner: indexed(address), spender: indexed(address), amount: wei_value})
Burned: event({burner: indexed(address), amount: wei_value})
Minted: event({to: indexed(address), amount: wei_value})
Transfer: event({source: indexed(address), to: indexed(address), amount: wei_value})

allowances: map(address, map(address, wei_value))
balances: map(address, wei_value)
decimals: public(uint256)
symbol: public(string[3])
owner_address: address
reserve_address: address
listing_address: address
supply: wei_value

@public
def __init__(initial_account: address, initial_balance: wei_value):
  self.owner_address = msg.sender
  self.balances[initial_account] = initial_balance
  self.decimals = 18
  self.symbol = "CMT"
  self.supply = initial_balance
  log.Transfer(ZERO_ADDRESS, initial_account, initial_balance)


@public
@constant
def allowance(owner: address, spender: address) -> wei_value:
  """
  @notice Fetch the amount a given spender is allowed to use on the owner's behalf
  @param owner Who owns the funds
  @param spender The appointed spender of said funds
  @return Said amount
  """
  return self.allowances[owner][spender]


@public
def approve(spender: address, amount: wei_value) -> bool:
  """
  @notice Approve a given spender for a given amount of funds
  @dev Note that there is an attack surface area here that someone may use both the
  old and new amounts via unfortunate TX ordering. A possible solution is to first
  reduce the spender's amount to 0 - then set the desired amt afterward. See
  https://github.com/ethereum/EIPs/issues/20#issuecomment-263524729
  @return true if successful
  """
  self.allowances[msg.sender][spender] = amount
  log.Approval(msg.sender, spender, amount)
  return True


@public
@constant
def balanceOf(owner: address) -> wei_value:
  """
  @notice Fetch the balance of the given address
  @param owner The address to fetch the balance of
  @return Balance of said address
  """
  return self.balances[owner]


@public
@constant
def getPrivileged() -> (address, address):
  """
  @notice return the address(es) of contracts that are recognized as being privileged
  @return The address(es)
  """
  return (self.reserve_address, self.listing_address)


@public
@constant
def hasPrivilege(sender: address) -> bool:
  """
  @notice Return a bool indicating whether the given address is a member of this contracts privileged group
  @return bool
  """
  return (sender == self.reserve_address or sender == self.listing_address)


@public
def burn(amount: wei_value):
  """
  @notice Burns the given amount of token wei
  @dev We only allow the market contract to call burn
  @param amount The amount to burn
  """
  # only the listing may call burn
  assert msg.sender == self.listing_address
  self.balances[msg.sender] -= amount
  self.supply -= amount
  log.Burned(msg.sender, amount)


@public
def burnAll(owner: address):
  """
  @notice Burns all market token associated with an address
  @dev We only allow the market contract to call burnAll
  @param address The owner of the tokens being burnt
  """
  # only the reserve may call burn all
  assert msg.sender == self.reserve_address
  bal: wei_value = self.balances[owner]
  self.supply -= bal
  clear(self.balances[owner])
  clear(self.allowances[owner][msg.sender])
  log.Burned(owner, bal)


@public
def decreaseAllowance(spender: address, amount: wei_value) -> bool:
  """
  @notice Decrement the amount allowed to a spender by the given amount
  @dev If the given amount is > the actual allowance we set it to 0
  @param spender The spender of the funds
  @param amount The amount to decrease a previous allowance by
  @return True if successful
  """
  self.allowances[msg.sender][spender] -= amount
  log.Approval(msg.sender, spender, self.allowances[msg.sender][spender])
  return True


@public
def increaseAllowance(spender: address, amount: wei_value) -> bool:
  """
  @notice Increase the amount a spender has allotted to them, by the owner, by the given amount
  @param spender The address whose allowance to increase
  @param amount The amount to increase by
  return True if successful
  """
  self.allowances[msg.sender][spender] += amount
  log.Approval(msg.sender, spender, self.allowances[msg.sender][spender])
  return True


@public
def mint(amount: wei_value):
  """
  @notice Create new Market Token funds and add them to the Market Contract balance
  @dev We only allow the Market Contract to call for minting
  """
  # either may call mint
  assert msg.sender == self.reserve_address or msg.sender == self.listing_address
  self.supply += amount
  self.balances[msg.sender] += amount
  log.Minted(msg.sender, amount)
  # TODO look at implementing a mintFor method to avoid mint followed by immediate transfer


# TODO look into the @noreentrant decorator
@public
def setPrivileged(reserve: address, listing: address):
  """
  @notice We restrict some activities to only privileged contracts. Can only be called once.
  @dev We only allow the owner to set the privileged address(es)
  @param reserve The deployed address of the reserve Contract
  @param listing The deployed address of the Listing Contract
  """
  assert msg.sender == self.owner_address
  assert self.reserve_address == ZERO_ADDRESS
  assert self.listing_address == ZERO_ADDRESS
  self.reserve_address = reserve
  self.listing_address = listing


@public
@constant
def totalSupply() -> wei_value:
  """
  @notice Total number of tokens, in wei, that exist
  @return The current total supply
  """
  return self.supply


@public
def transfer(to: address, amount: wei_value) -> bool:
  """
  @notice Move funds from the sender to a given address
  @param to The reciever of the funds
  @param amount How much to transfer
  @return true if successful
  """
  assert to != ZERO_ADDRESS
  self.balances[msg.sender] -= amount
  self.balances[to] += amount
  log.Transfer(msg.sender, to, amount)
  return True


@public
def transferFrom(source: address, to: address, amount: wei_value) -> bool:
  """
  @notice Move funds from one address to another
  @dev The allowance to `source` by `sender` is enforced (and decremented) here
  @param source Where the funds originate
  @param to Where the funds are going
  @param amount How much
  @return true if successful
  """
  assert source != ZERO_ADDRESS
  assert to != ZERO_ADDRESS
  self.balances[source] -= amount
  self.balances[to] += amount
  self.allowances[source][msg.sender] -= amount
  log.Transfer(source, to, amount)
  return True
