
# External Contracts
contract Parameterizer:
    def getBackendPayment() -> uint256: constant
    def getMakerPayment() -> uint256: constant
    def getReservePayment() -> uint256: constant
    def getCostPerByte() -> uint256(wei): constant
    def getStake() -> uint256(wei): constant
    def getPriceFloor() -> uint256(wei): constant
    def getHash(param: uint256, value: uint256) -> bytes32: constant
    def getSpread() -> uint256: constant
    def getListReward() -> uint256(wei): constant
    def getPlurality() -> uint256: constant
    def getReparam(hash: bytes32) -> (uint256, uint256): constant
    def getVoteBy() -> uint256(sec): constant
    def reparameterize(param: uint256, value: uint256): modifying
    def resolveReparam(hash: bytes32): modifying


