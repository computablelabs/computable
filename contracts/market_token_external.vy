
# External Contracts
contract Markettoken:
    def allowance(owner: address, spender: address) -> uint256(wei): constant
    def approve(spender: address, amount: uint256(wei)) -> bool: modifying
    def balanceOf(owner: address) -> uint256(wei): constant
    def getPrivileged() -> (address, address): constant
    def hasPrivilege(sender: address) -> bool: constant
    def burn(amount: uint256(wei)): modifying
    def burnAll(owner: address): modifying
    def decreaseApproval(spender: address, amount: uint256(wei)): modifying
    def increaseApproval(spender: address, amount: uint256(wei)): modifying
    def mint(amount: uint256(wei)): modifying
    def setPrivileged(listing: address, reserve: address): modifying
    def totalSupply() -> uint256(wei): constant
    def transfer(to: address, amount: uint256(wei)) -> bool: modifying
    def transferFrom(source: address, to: address, amount: uint256(wei)) -> bool: modifying
    def decimals() -> uint256: constant


