// Specifically request an abstraction for Validator 
var Validator = artifacts.require("Validator");

contract('Validator', function(accounts) {
  it("Should instantiate validator correctly", function () {
    // TODO(rbharath): Seems to always deploy at account[0] even if
    // we change to Validator.deployed(accounts[1]). Figure out
    // what's going on.
    return Validator.deployed(accounts[0]).then(function(instance) {
      return instance.owner.call();
    }).then(function(owner) {
      assert.equal(owner, accounts[0], "Owner isn't creator of validator");
    });
  });

  it("Owners should be able to record votes", function () {
    return Validator.deployed(accounts[0]).then(function(instance) {
      var meta = instance;
      meta.addVote("hello");
      return meta;
    }).then(function(instance) {
      var vote = instance.votes.call(0);
      return vote; 
    }).then(function(vote) {
      assert.equal(vote, "hello");
    });
  });

}); 
