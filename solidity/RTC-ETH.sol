pragma solidity 0.4.19;

import "./Modules/Administration.sol";
import "./Modules/Oraclize.sol";
import "./Math/SafeMath.sol";

contract RTCETH is Administration, usingOraclize {

	uint256 public ethUSD;
	uint256 public ethPerRtc;


	event NewOraclizeQuery(string result);
	event EthUsdPriceUpdated(uint256 price);
	event EthPerRtcUpdated(uint256 price);


	function __callback(bytes32 myid, string result) {
        locked = true;
        require(msg.sender == oraclize_cbAddress());
        require(validOraclizeIds[myid]);
        ethUSD = parseInt(result);
        uint256 oneEth = 1 ether;
        uint256 oneUsdOfEth = oneEth.div(ethUSD);
        ethPerRtc = oneUsdOfEth.div(8) // $1 / $0.125 = 8
        EthUsdPriceUpdated(ethUSD);
        EthPerRtcUpdated(ethPerRtc);
        delete validOraclizeIds[myid];
        locked = false;
    }

    function update() payable {
        require(this.balance >=oraclize_getPrice("URL"));
        NewOraclizeQuery("Oraclize query was sent, standing by for the answer..");
        bytes32 _id = oraclize_query(120, "URL", "json(https://api.coinmarketcap.com/v1/ticker/ethereum/?convert=USD).0.price_usd");
        validOraclizeIds[_id] = true;
    }


    function buyRtc()
    	public
    	returns (bool)
    {
    	require(msg.value > 0);
    	uint256 rtcPurchased = msg.value.div(ethPerRtc);
    }
}