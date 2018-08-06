var Token = artifacts.require("Token");
var RCC = artifacts.require("RCC");

module.exports = function(deployer) {
	deployer.deploy(Token).then(() => { return deployer.deploy(RCC,Token.address)});
}

//module.exports = function(deployer) {
	
//}
//};