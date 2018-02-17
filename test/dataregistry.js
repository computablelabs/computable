// request abstraction for DataRegistry
var DataRegistry = artifacts.require("DataRegistry");

contract("DataRegistry", function(accounts) {
  it("Should instantiate DataRegistry correctly", function () {
    return DataRegistry.deployed(accounts[0]).then(function(instance) {
      return instance.owner.call();
    }).then(function(owner) {
      assert.equal(owner, accounts[0], "Owner isn't creator of DataRegistry");
    });
  });
});

