pragma solidity^0.4.24;
import "./Token.sol";

contract RCC {
	// to do: make it so that voters can't see each other until reveal stage
	struct Challenge {
		//mapping(address => uint) public secret_votes;
		mapping(address => bool) won_lottery;
		mapping(address => bytes32) concealed_votes;
		mapping(address => bool) open_votes;
		mapping(address => bool) has_voted;
		mapping(address => bool) has_revealed;
		mapping(address => uint) rewards;
		
		uint votes_counter;
		uint revealed_votes_counter;
		bool recruiting_votes;
		bool reveal_stage;
		bool finished;
		uint yes_vote_counter;
		bool result;
		bool can_claim_rewards;

	}
	uint counter;
	uint vote_deposit;
	uint votes_per_challenge;
	// we keep track of who staked how how much
	mapping(address => uint) public stakes;
	// we need to track how many challenges the agents participate in, to ensure they can't withdraw stake in the middle of vote
	mapping(address => uint) public num_active_challenges;
	// also the total number of active challenges
	uint public total_active_challenges;
	// we index challenges with ints
	mapping(uint => Challenge) public challenges;


	Token public token;
	uint public throughput;  // this is actually throughput times 1000, to avoid fixed point types. Hacky, need a better fix.

	constructor(address token_address) public {
		counter = 0;
		vote_deposit = 1;
		votes_per_challenge = 3;
		token = Token(token_address);
		throughput = 990;

	}

	function create_challenge() public returns(uint) {
		challenges[counter] = Challenge({votes_counter: 0, revealed_votes_counter: 0, recruiting_votes: true, reveal_stage: false, finished:false, yes_vote_counter:0, result:false, can_claim_rewards: false});
		counter +=1;
		total_active_challenges+=1;
		return(counter-1);
	}

	function claim_lottery_win(uint challenge_id) public {

		uint p = stakes[msg.sender] * throughput /total_active_challenges;
		uint256 random =  uint256(uint256(keccak256(abi.encodePacked(block.blockhash(block.number-1), 420))) % 1000);
		if (p>= random) {
			challenges[challenge_id].won_lottery[msg.sender] = true;
		}
	}

	function submit_concealed_vote(bytes32 concealed_vote,uint challenge_id) public {
		require(token.is_allowed(msg.sender,this) >= vote_deposit);
		require(challenges[challenge_id].recruiting_votes);
		require(challenges[challenge_id].has_voted[msg.sender] == false);
		num_active_challenges[msg.sender] +=1;
		token.send_from(msg.sender,this, vote_deposit);
		challenges[challenge_id].votes_counter +=1;
		challenges[challenge_id].has_voted[msg.sender] = true;
		challenges[challenge_id].concealed_votes[msg.sender] = concealed_vote;
		//if (vote_value) {
		//	challenge.yes_vote_counter +=1;
		//}
		if (challenges[challenge_id].votes_counter>=votes_per_challenge){
			challenges[challenge_id].reveal_stage = true;
			challenges[challenge_id].recruiting_votes=false;
			//tally_up(challenge_id);
		}

	}

	function reveal_vote(bool vote,uint challenge_id, uint salt) public {
		// vote is in the reveal stage
		require(challenges[challenge_id].reveal_stage);
		// voter has voted
		require(challenges[challenge_id].has_voted[msg.sender]);
		// voter hasn't revealed the vote yet
		require(!challenges[challenge_id].has_revealed[msg.sender]);
		// the vote is consistent with the hash
		uint vote_int = 0;
		if (vote) {
			vote_int = 1;
		}


		require(keccak256(vote_int, salt) == challenges[challenge_id].concealed_votes[msg.sender]);
		challenges[challenge_id].open_votes[msg.sender] = vote;
		if (vote) {
			challenges[challenge_id].yes_vote_counter +=1;
		}
		challenges[challenge_id].revealed_votes_counter +=1;
		challenges[challenge_id].has_revealed[msg.sender] = true;
		if (challenges[challenge_id].revealed_votes_counter>= votes_per_challenge) {
			challenges[challenge_id].reveal_stage = false;
			challenges[challenge_id].finished = true;
			tally_up(challenge_id);
		}

	}

	function tally_up(uint challenge_id) public {
		require(challenges[challenge_id].finished);
		if (2*challenges[challenge_id].yes_vote_counter>= votes_per_challenge) {
			challenges[challenge_id].result = true;
		}
		challenges[challenge_id].can_claim_rewards = true;
		total_active_challenges-=1;
		
		//challenge.tokens_distributed=true
	}

	function challenge_finished(uint challenge_id) public returns(bool) {
		return(challenges[challenge_id].finished);
	}
	function challenge_result(uint challenge_id) public returns(bool) {
		return(challenges[challenge_id].result);
	}

	function claim_reward(uint challenge_id) public {
		// can this fail due to rounding or something?
		require(challenges[challenge_id].can_claim_rewards);
		token.transfer(this,msg.sender,challenges[challenge_id].rewards[msg.sender]);
		challenges[challenge_id].rewards[msg.sender] = 0;
		num_active_challenges[msg.sender]-=1;

	}
}