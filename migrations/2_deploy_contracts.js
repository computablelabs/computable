var DataCoin = artifacts.require('DataCoin');
var Validator = artifacts.require('Validator');
var DataRegistry = artifacts.require('DataRegistry');

module.exports = function (deployer) {
  deployer.deploy(DataCoin);
  deployer.deploy(Validator);
  deployer.deploy(DataRegistry, 'TestRegistry');
};
