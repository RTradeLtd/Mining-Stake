pragma solidity 0.4.19;


/**

    This contract is used to lockup RTCoins to collect a payout in RTCoin.

*/

import "./Math/SafeMath.sol";
import "./Modules/Administration.sol";
import "./Modules/Oraclize.sol";
import "./Interfaces/RTCoinInterface.sol";

contract TokenLockup is Administration, usingOraclize {

    using SafeMath for uint256;

    /**CONSTANTS*/
    uint256 public constant DEFAULTLOCKUPTIME = 4 weeks;
    uint256 public constant MINIMUMLOCKUPAMOUNT = 100000000000000000000; // 100 RTC ($12.50 at $0.125/rtc)
    string  public constant VERSION = "0.0.1beta";

    // keeps track of the latest eth-usd ratio, with no decimals
    uint256 public ethUSD;
    // hot wallet used to collect sign up fees,
    address public rtcHotWallet;
    // will always be equivalent to $10 USD of ethereum +/- a few cents
    uint256 public signUpFee;
    // how many RTC for a single kilo hash rate a second, starts out at 1rtc = 6.25KH/s
    // at starting evaluation, it costs $0.1255USD to get 1 hash a second.
    uint256 public kiloHashSecondPerRtc = 6250000000000000000;
    uint256 public stakerCount;
    bool    private locked;


    RTCoinInterface private rtI;

    struct StakerStruct {
        address addr;
        uint256 rtcStaked;
        uint256 khSec;
        uint256 depositDate;
        uint256 releaseDate;
        bytes32 id;
        bool    enabled;
    }

    struct RewardStruct {
        address addr;
        uint256 eth;
        uint256 rtc;
    }

    mapping (bytes32 => bool)   private validOraclizeIds; // keep to private, helps reduce gas costs
    mapping (address => StakerStruct) private stakers;
    mapping (address => RewardStruct) private rewards;
    mapping (address => uint256)      private ethBalances;

    event StakeDeposited(address _depositer, uint256 _amount, uint256 _weeksStaked, uint256 _khSec, bytes32 _id);
    event RewardDeposited(address _staker, uint256 _rtcStaked, uint256 _ethMined);
    event EthWithdrawn(address _withdrawer, uint256 _amount);
    event NewOraclizeQuery(string result);
    event EthUsdPriceUpdated(uint256 price);
    event SignUpFeeUpdated(uint256 fee);

    modifier registeredStaker(address _addr) {
        require(stakers[_addr].enabled);
        _;
    }

    modifier notRegisteredStaker(address _addr) {
        require(!stakers[_addr].enabled);
        _;
    }

    modifier notLocked() {
        require(!locked);
        _;
    }

    modifier isLocked() {
        require(locked);
        _;
    }

    /**
        CONSTRUCTOR
    */
    function TokenLockup() payable {
        bytes32 id = oraclize_query("URL", "json(https://api.coinmarketcap.com/v1/ticker/ethereum/?convert=USD).0.price_usd");
        validOraclizeIds[id] = true;
    }

    /**
        @dev Fallback, allows depositing of ether into the contract
    */
    function () payable {}

    function depositStake(
        uint256 _rtcToStake,
        uint256 _durationInWeeksToStake)
        public
        notRegisteredStaker(msg.sender)
        notLocked
        returns (bool)
    {
        bytes32 id = keccak256(msg.sender, _rtcToStake, _durationInWeeksToStake, now);
        uint256 khSec = _rtcToStake.mul(kiloHashSecondPerRtc);
        khSec = khSec.div(1 ether);
        stakers[msg.sender].addr = msg.sender;
        stakers[msg.sender].rtcStaked = _rtcToStake;
        stakers[msg.sender].khSec = khSec;
        stakers[msg.sender].depositDate = now;
        stakers[msg.sender].releaseDate = (now + (_durationInWeeksToStake * 1 weeks));
        stakers[msg.sender].id = id;
        stakers[msg.sender].enabled = true;
        StakeDeposited(msg.sender, _rtcToStake, _durationInWeeksToStake, khSec, id);
        require(rtI.transferFrom(msg.sender, address(this), _rtcToStake));
        return true;
    }

    function depositReward(
        address _staker,
        uint256 _ethMined,
        uint256 _rtcStaked)
        public
        onlyAdmin
        registeredStaker(_staker)
        notLocked
        payable
        returns (bool)
    {
        require(msg.value == _ethMined &&  _ethMined > 0 && _rtcStaked > 0);
        rewards[_staker].eth =  rewards[_staker].eth.add(_ethMined);
        rewards[_staker].rtc = rewards[_staker].rtc.add(_rtcStaked);
        ethBalances[_staker] = ethBalances[_staker].add(_ethMined);
        RewardDeposited(_staker, _rtcStaked, _ethMined);
        // we can transfer tokens right award since it doesn't trigger code execution
        require(rtI.transferFrom(msg.sender, _staker, _rtcStaked));
        return true;
    }

    function withdrawEth()
        public
        registeredStaker(msg.sender)
        notLocked
        returns (bool)
    {
        uint256 eth = ethBalances[msg.sender];
        // reset, prevent re-entrancy
        ethBalances[msg.sender] = 0;
        EthWithdrawn(msg.sender, eth);
        msg.sender.transfer(eth);
        return true;
    }

    /**
        @dev Callback function, used by Oraclize to update the eth-usd conversion rate
    */
    function __callback(bytes32 myid, string result) {
        locked = true;
        require(msg.sender == oraclize_cbAddress());
        require(validOraclizeIds[myid]);
        ethUSD = parseInt(result);
        uint256 oneEth = 1 ether;
        signUpFee = oneEth.div(ethUSD);
        signUpFee = signUpFee.mul(10);
        EthUsdPriceUpdated(ethUSD);
        SignUpFeeUpdated(signUpFee);
        delete validOraclizeIds[myid];
        locked = false;
        update();
    }

    /**
        @dev Used to trigger an ETH-USD state update
        @notice Marked private to prevent anyone from forcing an udpate and wasting our ethereum
    */
    function update() private {
        require(this.balance >=oraclize_getPrice("URL"));
        NewOraclizeQuery("Oraclize query was sent, standing by for the answer..");
        bytes32 _id = oraclize_query(600, "URL", "json(https://api.coinmarketcap.com/v1/ticker/ethereum/?convert=USD).0.price_usd");
        validOraclizeIds[_id] = true;
    }

}
