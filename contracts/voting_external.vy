
# External Contracts
contract Voting:
    def getPrivileged() -> (address, address, address): constant
    def setPrivileged(parameterizer: address, datatrust: address, listing: address): modifying
    def hasPrivilege(sender: address) -> bool: constant
    def candidateIs(hash: bytes32, kind: uint256) -> bool: constant
    def isCandidate(hash: bytes32) -> bool: constant
    def getCandidate(hash: bytes32) -> (uint256, address, uint256(wei), uint256(sec, positional), uint256, uint256): constant
    def getCandidateOwner(hash: bytes32) -> address: constant
    def addCandidate(hash: bytes32, kind: uint256, owner: address, stake: uint256(wei), vote_by: uint256(sec)): modifying
    def removeCandidate(hash: bytes32): modifying
    def didPass(hash: bytes32, plurality: uint256) -> bool: constant
    def pollClosed(hash: bytes32) -> bool: constant
    def vote(hash: bytes32, option: uint256): modifying
    def transferStake(hash: bytes32, addr: address): modifying
    def getStake(hash: bytes32, addr: address) -> uint256(wei): constant
    def unstake(hash: bytes32): modifying


