
# External Contracts
contract Reserve:
    def getSupportPrice() -> uint256(wei): constant
    def support(offer: uint256(wei)): modifying
    def getWithdrawalProceeds(addr: address) -> uint256(wei): constant
    def withdraw(): modifying


