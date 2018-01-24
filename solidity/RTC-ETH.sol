pragma solidity 0.4.19;

import "./Modules/Administration.sol";
import "./Modules/Oraclize.sol";
import "./Math/SafeMath.sol";
import "./Interfaces/RTCoinInterface.sol";


contract RTCETH is Administration, usingOraclize {

	using SafeMath for uint256;

	address public hotWallet;
	uint256 public ethUSD;
	uint256 public ethPerRtc;
	bool	public oracleUpdatesDisabled;
	bool	public locked;

	RTCoinInterface private rtI;

	mapping (address => bool) private bonus; 


	event NewOraclizeQuery(string result);
	event EthUsdPriceUpdated(uint256 price);
	event EthPerRtcUpdated(uint256 price);

	modifier notLocked() {
		require(!locked);
		_;
	}


	function RTCETH() payable {
		rtI = RTCoinInterface(address(0x0));
		bytes32 _id = oraclize_query("URL", "json(https://api.coinmarketcap.com/v1/ticker/ethereum/?convert=USD).0.price_usd")
		NewOraclizeQuery("Oraclize Query Was Sent. Standing By For Answer")
		validOraclizeIds[_id] = true;
	}

	function () payable {}

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
        update();
        locked = false;
    }

    function update() payable {
        require(this.balance >=oraclize_getPrice("URL"));
        if (!oracleUpdatesDisabled) {
        	NewOraclizeQuery("Oraclize query was sent, standing by for the answer..");
        	bytes32 _id = oraclize_query(120, "URL", "json(https://api.coinmarketcap.com/v1/ticker/ethereum/?convert=USD).0.price_usd");
        	validOraclizeIds[_id] = true;
        } else {
        	NewOraclizeQuery("Oracle Updates Are Disabled");
        }
    }


    function buyRtc()
    	public
    	payable
    	notLocked
    	returns (bool)
    {
    	require(msg.value > 0);
    	uint256 fee;
    	if (bonus[msg.sender]) { // sender is eligible for bonus, so fee reduction
    		fee = ethPerRtc.div(2);
    	} else { // sender is not eligible for bonus, so no fee reduction
    		fee = ethPerRtc;
    	}
    	uint256 rtcPurchased = msg.value.div(fee);
    	requiire(rtI.balanceOf(this) >= rtcPurchased);
    	rtI.transfer(msg.sender, rtcPurchased);
    	uint256 amountMinusOracleFee = msg.value.sub(oraclize_getPrice("URL").mul(2));
    	hotWallet.transfer(amountMinusOracleFee);
    	return true;
    }
}