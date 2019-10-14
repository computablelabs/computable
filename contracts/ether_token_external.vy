
# External Contracts
contract Ethertoken:
    def allowance(owner: address, spender: address) -> uint256(wei): constant
    def approve(spender: address, amount: uint256(wei)) -> bool: modifying
    def balanceOf(owner: address) -> uint256(wei): constant
    def decreaseAllowance(spender: address, amount: uint256(wei)) -> bool: modifying
    def deposit(): modifying
    def increaseAllowance(spender: address, amount: uint256(wei)) -> bool: modifying
    def totalSupply() -> uint256(wei): constant
    def transfer(to: address, amount: uint256(wei)) -> bool: modifying
    def transferFrom(source: address, to: address, amount: uint256(wei)) -> bool: modifying
    def withdraw(amount: uint256(wei)): modifying
    def decimals() -> uint256: constant
    def symbol() -> string[3]: constant


