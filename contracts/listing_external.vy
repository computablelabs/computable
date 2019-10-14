
# External Contracts
contract Listing:
    def isListed(hash: bytes32) -> bool: constant
    def withdrawFromListing(hash: bytes32, amount: uint256(wei)): modifying
    def list(hash: bytes32): modifying
    def getListing(hash: bytes32) -> (address, uint256(wei)): constant
    def resolveApplication(hash: bytes32): modifying
    def claimAccessReward(hash: bytes32): modifying
    def challenge(hash: bytes32): modifying
    def resolveChallenge(hash: bytes32): modifying
    def exit(hash: bytes32): modifying


