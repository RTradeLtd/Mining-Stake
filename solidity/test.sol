pragma solidity 0.4.19;

import "./Modules/Oraclize.sol";


contract test is usingOraclize {


	uint256 public ethUSD;

	mapping (bytes32 => bool) public validOraclizeIds;

	event newOraclizeQuery(string result);
	event newEthUsdPrice(string result);

	function test() payable {}

	function () payable {}

	function __callback(bytes32 myid, string result) {
        require(msg.sender == oraclize_cbAddress());
        require(validOraclizeIds[myid]);
        ethUSD = parseInt(result);
        delete validOraclizeIds[myid];
        //updateEthUsd();
    }

    // WIP, not totally, set to update price every 600 seconds
    function update() payable {
        require(this.balance >=oraclize_getPrice("URL"));
        newOraclizeQuery("Oraclize query was sent, standing by for the answer..");
        bytes32 _id = oraclize_query("60", "URL", "json(https://api.coinmarketcap.com/v1/ticker/ethereum/?convert=USD).0.price_usd");
        validOraclizeIds[_id] = true;
    }
}