var DataCoin = artifacts.require("DataCoin");
var Validator = artifacts.require("Validator");

module.exports = function(deployer) {
  deployer.deploy(DataCoin);
  deployer.deploy(Validator);
}
