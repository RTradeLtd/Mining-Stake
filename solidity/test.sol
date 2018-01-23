pragma solidity 0.4.19;

import "./Modules/Oraclize.sol";


contract test is usingOraclize {


	uint256 public ethUSD;

	mapping (bytes32 => bool) public validOraclizeIds;

	event newOraclizeQuery(string result);

	function test() {}

	function () payable {}

	function __callback(bytes32 id, string result) {
        require(validOraclizeIds[id]);
        require(msg.sender == oraclize_cbAddress());
        ethUSD = parseInt(result);
        delete validOraclizeIds[id];
        //updateEthUsd();
    }

    // WIP, not totally, set to update price every 600 seconds
    function updateEthUsd()
        payable
        returns (bool)
    {
        require(this.balance >=oraclize_getPrice("URL"));
        newOraclizeQuery("Oraclize query was sent, standing by for the answer..");
        bytes32 id = oraclize_query("URL", "json(https://api.coinmarketcap.com/v1/ticker/ethereum/?convert=USD).0.price_usd", 30000000000);
        validOraclizeIds[id] = true;
    }
}