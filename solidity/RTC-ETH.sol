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
	bool	public pendingPriceUpdate;
	RTCoinInterface private rtI;

	mapping (address => bool) private bonus; 
	mapping (bytes32 => bool) private validOraclizeIds;

	event NewOraclizeQuery(string result);
	event EthUsdPriceUpdated(uint256 price);
	event EthPerRtcUpdated(uint256 price);
	event RtcPurchased(address _purchaser, uint256 _amountPurchased);
	event EthWithdrawn(address _recipient, uint256 _amountWithdrawn);
	event VztWithdrawn(address _recipient, uint256 _amountWithdrawn);

	modifier notLocked() {
		require(!locked);
		_;
	}


	function RTCETH() payable {
		// rinkeby address
		rtI = RTCoinInterface(0x89B44F01e1a363E46E175301603b141a00b3C884);
		bytes32 _id = oraclize_query("URL", "json(https://api.coinmarketcap.com/v1/ticker/ethereum/?convert=USD).0.price_usd");
		NewOraclizeQuery("Oraclize Query Was Sent. Standing By For Answer");
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
        ethPerRtc = oneUsdOfEth.div(8); // $1 / $0.125 = 8
        EthUsdPriceUpdated(ethUSD);
        EthPerRtcUpdated(ethPerRtc);
        delete validOraclizeIds[myid];
        locked = false;
        pendingPriceUpdate = false;
        update();
    }

    function update() payable {
    	require(!pendingPriceUpdate);
    	//require(msg.sender == owner || msg.sender == admin || msg.sender == oraclize_cbAddress() || msg.sender == address(this));
        require(this.balance >=oraclize_getPrice("URL"));
        if (!oracleUpdatesDisabled) {
        	pendingPriceUpdate = true;
        	NewOraclizeQuery("Oraclize query was sent, standing by for the answer..");
        	bytes32 _id = oraclize_query(120, "URL", "json(https://api.coinmarketcap.com/v1/ticker/ethereum/?convert=USD).0.price_usd");
        	validOraclizeIds[_id] = true;
        } else {
        	NewOraclizeQuery("Oracle Updates Are Disabled");
        }
    }

    // Used to manually force an update, requires user pay double the gas cost, but send less than 4x the gas cost
    // contract MUST have the required ether to pay the initial gas
    function forceUpdate() payable {
    	require(this.balance >= oraclize_getPrice("URL"));
    	require(msg.value >= oraclize_getPrice("URL").mul(2) && msg.value < oraclize_getPrice("URL").mul(4));
    	if (!oracleUpdatesDisabled) {
    		NewOraclizeQuery("Oraclize query was sent, standing by for the answer..");
        	bytes32 _id = oraclize_query("URL", "json(https://api.coinmarketcap.com/v1/ticker/ethereum/?convert=USD).0.price_usd");
        	validOraclizeIds[_id] = true;
    	} else {
    		NewOraclizeQuery("Oracle Updates Are Disabled, please enable before trying another force update");
    	}
    }

    function addBonusAddress(
    	address _bonusRecipient
    )
    	public
    	onlyAdmin
    	returns (bool)
    {
    	bonus[_bonusRecipient] = true;
    	return true;
    }

    function removeBonusAddress(
    	address _bonusRecipient
    )
    	public
    	onlyAdmin
    	returns (bool)
    {
    	bonus[_bonusRecipient] = false;
    	return true;
    }

    function setHotWallet(
    	address _hotWalletAddress
    )
    	public
    	onlyAdmin
    	returns (bool)
    {
    	require(_hotWalletAddress != hotWallet);
    	hotWallet = _hotWalletAddress;
    	return true;
    }


    function withdrawVzt(
    	address _recipient,
    	uint256 _amount
    )
    	public
    	onlyAdmin
    	notLocked
    	returns (bool)
    {
    	require(rtI.balanceOf(address(this)) >= _amount && _recipient != address(0x0));
    	VztWithdrawn(_recipient, _amount);
    	require(rtI.transfer(_recipient, _amount));
    	return true;
    }

    function withdrawEth(
    	address _recipient
    )
    	public
    	onlyAdmin
    	notLocked
    	returns (bool)
    {
    	oracleUpdatesDisabled = true;
    	uint256 fee = this.balance.sub(oraclize_getPrice("URL").mul(2));
    	require(_recipient != address(this));
    	EthWithdrawn(_recipient, fee);
    	_recipient.transfer(fee);
    }

    function buyRtc()
    	public
    	payable
    	notLocked
    	returns (bool)
    {
    	require(hotWallet != address(0x0));
    	require(msg.value > 0);
    	uint256 fee;
    	if (bonus[msg.sender]) { // sender is eligible for bonus, so fee reduction
    		fee = ethPerRtc.div(2);
    	} else { // sender is not eligible for bonus, so no fee reduction
    		fee = ethPerRtc;
    	}
    	uint256 rtcPurchased = (msg.value.div(fee)).mul(1 ether);
    	require(rtI.balanceOf(this) >= rtcPurchased);
    	require(rtI.transfer(msg.sender, rtcPurchased));
    	// lets make sure we have  enough for a future oracle call
    	uint256 amountMinusOracleFee = msg.value.sub(oraclize_getPrice("URL").mul(2));
    	RtcPurchased(msg.sender, rtcPurchased);
    	hotWallet.transfer(amountMinusOracleFee);
    	return true;
    }
}