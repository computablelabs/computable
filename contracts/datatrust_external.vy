
# External Contracts
contract Datatrust:
    def getPrivileged() -> address: constant
    def getReserve() -> address: constant
    def setPrivileged(listing: address): modifying
    def getHash(url: string[128]) -> bytes32: constant
    def getBackendAddress() -> address: constant
    def getBackendUrl() -> string[128]: constant
    def setBackendUrl(url: string[128]): modifying
    def getDataHash(hash: bytes32) -> bytes32: constant
    def setDataHash(listing: bytes32, data: bytes32): modifying
    def removeDataHash(hash: bytes32): modifying
    def register(url: string[128]): modifying
    def resolveRegistration(hash: bytes32): modifying
    def requestDelivery(hash: bytes32, amount: uint256): modifying
    def getBytesPurchased(addr: address) -> uint256: constant
    def getDelivery(hash: bytes32) -> (address, uint256, uint256): constant
    def listingAccessed(listing: bytes32, delivery: bytes32, amount: uint256): modifying
    def getAccessRewardEarned(hash: bytes32) -> uint256(wei): constant
    def accessRewardClaimed(hash: bytes32): modifying
    def delivered(delivery: bytes32, url: bytes32): modifying


