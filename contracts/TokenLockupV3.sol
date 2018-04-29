pragma solidity 0.4.23;


/**

    This contract is used to lockup RTCoins to collect a payout in RTCoin.

*/

import "./Math/SafeMath.sol";
import "./Modules/Administration.sol";
import "./Interfaces/RTCoinInterface.sol";

contract TokenLockup is Administration {

    using SafeMath for uint256;


    /**CONSTANTS*/
    uint256 public constant DEFAULTLOCKUPTIME = 4 weeks;
    uint256 public constant MINSTAKE = 100000000000000000000; // 100 RTC ($12.50 at $0.125/rtc)
    string  public constant VERSION = "1.0.0beta";


    // keeps track of hte oraclize contract address from which all price updates will come from
    address public oracleContractAddress;
    // keeps track of the latest rtc-usd ratio, with no decimals
    uint256 public rtcUSD;
    
    // hot wallet used to collect sign up fees,
    address public rtcHotWallet;
    // will always be equivalent to $10 USD of ethereum +/- a few cents
    uint256 public signUpFee;
    // how many RTC for a single kilo hash rate a second, starts out at 1rtc = 6.25KH/s
    uint256 public kiloHashSecondPerRtc = 6250000000000000000;
    // keeps track of how many KH/s for a single USD cent (0.02 kh/sec)
    uint256 public kiloHashSecondPerOneCentUsd = 20000000000000000;
    uint256 public stakerCount;
    bool    public locked;

    function updatePrices(
        uint256 _rtcUSD)
        public
        onlyAdmin
        onlyOracleContract(msg.sender)
        returns (bool)
    {
        rtcUSD = _rtcUSD;
        return true;
    }

    RTCoinInterface public rtI = RTCoinInterface(0x0994f9595d28429584bfb5fcbfea75b9c9ea2c24);

    struct StakerStruct {
        address addr;
        uint256 rtcStaked;
        uint256 deposit;
        uint256 khSec;
        uint256 depositDate;
        uint256 releaseDate;
        uint256 id;
        string encryptedEmail;
        bool    enabled;
    }

    struct RewardStruct {
        uint256 ethRewarded;
        uint256 rtcRewarded;
    }

    mapping (address => StakerStruct[]) public stakers;
    mapping (address => RewardStruct) public rewards;
    mapping (address => uint256) public numStakes;

    event StakeDeposited(address _depositer, uint256 _amount, uint256 _weeksStaked, uint256 _khSec, uint256 _id);
    event DepositWithdrawn(address _staker, uint256 _amount, uint256 _stakeId);
    event EthWithdrawn(address _withdrawer, uint256 _amount);
    event RtcReward(address _staker, uint256 _amount);
    event EthReward(address _staker, uint256 _amount);
    event NewOraclizeQuery(string result);
    event EthUsdPriceUpdated(uint256 price);
    event SignUpFeeUpdated(uint256 fee);

    modifier onlyOracleContract(address _sender) {
        require(_sender == oracleContractAddress);
        _;
    }

    modifier registeredStaker(address _staker, uint256 _id) {
        require(stakers[_staker][_id].enabled);
        _;
    }

    modifier pastReleaseDate(address _staker, uint256 _id) {
        require(now > stakers[_staker][_id].releaseDate);
        _;
    }

    modifier stakeEnabled(address _staker, uint256 _id) {
        require(stakers[_staker][_id].enabled);
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
        @dev Fallback, allows depositing of ether into the contract
    */
    function ()public  payable {}

    function setRtI(
        address _rtcAddress)
        public
        onlyAdmin
        returns (bool)
    {
        rtI = RTCoinInterface(_rtcAddress);
        return true;
    }

    function setRtHotWallet(
        address _rtHotWallet)
        public
        onlyAdmin
        returns (bool)
    {
        rtcHotWallet = _rtHotWallet;
        return true;
    }

    function depositStake(
        uint256 _rtcToStake,
        uint256 _durationInWeeksToStake,
        string _encryptedEmail)
        public
        notLocked
        returns (bool)
    {
        require(_rtcToStake >= MINSTAKE && _durationInWeeksToStake >= 4);
        kiloHashSecondPerRtc = rtcUSD.div(kiloHashSecondPerOneCentUsd);
        uint256 id = numStakes[msg.sender];
        numStakes[msg.sender] = numStakes[msg.sender].add(1);
        uint256 khSec = _rtcToStake.mul(kiloHashSecondPerRtc);
        khSec = khSec.div(1 ether);
        stakers[msg.sender].push(StakerStruct(
            msg.sender,
            _rtcToStake,
            _rtcToStake,
            khSec,
            now,
            (now + (_durationInWeeksToStake * 1 weeks)),
            id,
            _encryptedEmail,
            true));
        emit StakeDeposited(msg.sender, _rtcToStake, _durationInWeeksToStake, khSec, id);
        require(rtI.transferFrom(msg.sender, rtcHotWallet, _rtcToStake));
        return true;
    }

    function withdrawStake(
        uint256 _stakeId)
        public
        pastReleaseDate(msg.sender, _stakeId)
        stakeEnabled(msg.sender, _stakeId)
        returns (bool)
    {
        assert(stakers[msg.sender][_stakeId].deposit > 0);
        uint256 deposit = stakers[msg.sender][_stakeId].deposit;
        stakers[msg.sender][_stakeId].deposit = 0;
        stakers[msg.sender][_stakeId].enabled = false;
        emit DepositWithdrawn(msg.sender, deposit, _stakeId);
        require(rtI.transfer(msg.sender, deposit));
        return true;
    }

    function routeRtcRewards(
        address[] _stakers,
        uint256[] _payments)
        public
        onlyAdmin
        returns (bool)
    {
        require(_stakers.length == _payments.length);
        for (uint256 i = 0; i < _stakers.length; i++) {
            uint256 rtc = _payments[i];
            rewards[_stakers[i]].rtcRewarded = rewards[_stakers[i]].rtcRewarded.add(rtc);
            emit RtcReward(_stakers[i], rtc);
            require(rtI.transferFrom(msg.sender, _stakers[i], rtc));
        }
        return true;
    }

    function routeEthReward(
        address[] _stakers,
        uint256[] _payments)  
        public
        onlyAdmin
        payable
        returns (bool)
    {
        require(msg.value > 0);
        require(_stakers.length == _payments.length);
        for (uint256 i = 0; i < _stakers.length; i++) {
            uint256 eth = _payments[i];
            rewards[_stakers[i]].ethRewarded = rewards[_stakers[i]].ethRewarded.add(eth);
            emit EthReward(_stakers[i], eth);
            require(_stakers[i].send(eth));
        }
        return true;
    }

    function getRewardStruct(
        address _staker)
        public
        view
        returns (uint256, uint256)
    {
        return (rewards[_staker].ethRewarded, rewards[_staker].rtcRewarded);
    }

    function getStakerStruct(
        address _staker,
        uint256 _id)
        public
        view
        returns (uint256, uint256, uint256, uint256, uint256, bool)
    {
        return (
            stakers[_staker][_id].rtcStaked,
            stakers[_staker][_id].khSec,
            stakers[_staker][_id].depositDate,
            stakers[_staker][_id].releaseDate,
            stakers[_staker][_id].id,
            stakers[_staker][_id].enabled);
    }

    function getNumStakes(
        address _staker)
        public
        view
        returns (uint256)
    {
        return numStakes[_staker];
    }

}