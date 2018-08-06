const abi = require('ethereumjs-abi');
var RCC = artifacts.require("RCC")
var Token = artifacts.require("Token")

function getVoteSaltHash(vote, salt) {
    //return `0x${abi.soliditySHA3(['uint','uint'], [vote, salt]).toString('hex')}`;
    return `0x${abi.soliditySHA3(['uint', 'uint'], [vote, salt]).toString('hex')}`;
  }


contract("RCC initial test", async(accounts) =>{
	it("should create a challenge, place votes, pass a challenge", async() =>{
		// deploy the contracts
		let token = await Token.deployed();
		
		let rcc = await RCC.deployed();
		
		// make allowances on behalf of users
		await token.transfer(accounts[0],accounts[1],1000,{from: accounts[0]})
		await token.transfer(accounts[0],accounts[2],1000,{from: accounts[0]})

		await token.approve(rcc.address,1000,{from:accounts[0]});
		await token.approve(rcc.address,1000,{from:accounts[1]});
		await token.approve(rcc.address,1000,{from:accounts[2]});

		//create a challenge
		hash = getVoteSaltHash(1,420);
		hash2 = getVoteSaltHash(0,420);
		//hash = `0x${abi.soliditySHA3(['uint'],[1]).toString('hex')}`;
		//hash="a57cbc4172acb2b0136b6ab0a21bdf543d2afd272124ee0970b565406e7e975a";
		await rcc.create_challenge({from: accounts[0]});
		//vote

		await rcc.submit_concealed_vote(hash,0,{from:accounts[0]});
		await rcc.submit_concealed_vote(hash,0,{from:accounts[1]});
		await rcc.submit_concealed_vote(hash2,0,{from:accounts[2]});
		//console.log("encrypted:");
		//console.log(keccak256(1,420))
		await rcc.reveal_vote(true,0,420,{from:accounts[0]});

		await rcc.reveal_vote(true,0,420,{from:accounts[1]});
		await rcc.reveal_vote(false,0,420,{from:accounts[2]});


		//console.log(id.toNumber());
		assert.equal(await rcc.challenge_finished.call(0),true);
		assert.equal(await rcc.challenge_result.call(0),true);

	});
});