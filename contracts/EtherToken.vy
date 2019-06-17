# @title Computable EtherToken
# @notice Implementation of an ERC20 compatible token which is purchased with ETH at a 1:1 ratio
# @author Computable
# @dev Implements ERC20 interface

# NOTE: We don't use our own past-tense style event names for Approval && Transfer per ERC20 Standard
Approval: event({owner: indexed(address), spender: indexed(address), amount: wei_value})
Deposited: event({source: indexed(address), amount: wei_value})
Transfer: event({source: indexed(address), to: indexed(address), amount: wei_value})
Withdrawn: event({to: indexed(address), amount: wei_value})

allowances: map(address, map(address, wei_value))
balances: map(address, wei_value)
decimals: public(uint256)
symbol: public(string[3])
supply: wei_value

@public
def __init__(initial_account: address, initial_balance: wei_value):
  self.balances[initial_account] = initial_balance
  self.decimals = 18
  self.symbol = "CET"
  self.supply = initial_balance


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
def decreaseApproval(spender: address, amount: wei_value):
  """
  @notice Decrement the amount allowed to a spender by the given amount
  @dev If the given amount is > the actual allowance we set it to 0
  @param spender The spender of the funds
  @param amount The amount to decrease a previous allowance by
  """
  self.allowances[msg.sender][spender] -= amount #vyper will throw if overrun here
  log.Approval(msg.sender, spender, self.allowances[msg.sender][spender])


@public
@payable
def deposit():
  """
  @notice Facilitate a user purchasing EtherToken with Eth at a 1:1 ratio
  """
  self.balances[msg.sender] += msg.value
  self.supply += msg.value
  log.Deposited(msg.sender, msg.value)


@public
def increaseApproval(spender: address, amount: wei_value):
  """
  @notice Increase the amount a spender has allotted to them, by the owner, by the given amount
  @param spender The address whose allowance to increase
  @param amount The amount to increase by
  """
  self.allowances[msg.sender][spender] += amount
  log.Approval(msg.sender, spender, self.allowances[msg.sender][spender])


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


@public
def withdraw(amount: wei_value):
  """
  @notice Allow msg.sender to withdraw an amount of ETH
  @dev The Vyper builtin `send` is used here
  @param amount An amount of ETH in wei to be withdrawn
  """
  self.balances[msg.sender] -= amount
  self.supply -= amount
  send(msg.sender, amount)
  log.Withdrawn(msg.sender, amount)
