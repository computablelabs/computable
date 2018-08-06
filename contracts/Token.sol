pragma solidity^0.4.24;

contract Token {
	mapping(address => uint) public balances;
	mapping(address => mapping(address=> uint)) public allowed;
	event transfer_event(address from, address to, uint value);

	constructor() public {
		balances[msg.sender] = 10000;

	}
	function transfer(address from, address to, uint value) public {
		require(msg.sender == from);
		require(balances[msg.sender]>=value);
		balances[msg.sender]-= value;
		balances[to]+= value;
		emit transfer_event(from,to,value);
	}


	function approve(address spender, uint value) public {
		require(get_balance(msg.sender) >= value);
		allowed[msg.sender][spender] = value; 
	}

	function send_from(address sender, address receiver, uint value) public {
		require( allowed[sender][msg.sender] >= value);
		allowed[sender][msg.sender] -= value;
		balances[sender] -= value;
		balances[receiver] +=value;
		emit transfer_event(sender, receiver,value);
	}

	function get_balance(address user) public view returns(uint){
		return(balances[user]);
	}

	function is_allowed(address from, address sender) public view returns(uint) {
		return(allowed[from][sender]);
	}




}