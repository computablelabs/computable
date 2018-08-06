var Token = artifacts.require("Token")


contract("Simple token test", async(accounts) => {
	it("should give creator 10000 tokens", async() => {
		let instance = await Token.deployed();
		let balance = await instance.get_balance.call(accounts[0]);
		assert.equal(balance,10000);
	});

	it("should succesfully transfer tokens", async() =>{
		let coin = await Token.deployed();
		//var transfer_event = coin.transfer_event({_from:web3.eth.coinbase});
		//transfer_event.watch(function(err, result) {
    	//console.log(result);
		//});

		await coin.transfer(accounts[0],accounts[1],1000,{from:accounts[0]});
		let balance = await coin.get_balance.call(accounts[1]);
		assert.equal(balance,1000);

	});

	it("should transfer tokens on behalf of others", async() => {
		let coin = await Token.deployed();
		await coin.approve(accounts[1],1000,{from:accounts[0]});
		await coin.send_from(accounts[0],accounts[2],1000, {from: accounts[1]});
		let balance = await coin.get_balance.call(accounts[2]);
		assert.equal(balance,1000);

		try {await coin.send_from(accounts[0],accounts[2],1000, {from: accounts[1]});} catch (error) {}

		balance = await coin.get_balance.call(accounts[2]);
		//await console.log("asdfsdf");
		assert.equal(balance,1000);

	});
});