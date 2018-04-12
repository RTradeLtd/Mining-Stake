pragma solidity 0.4.21;

import "Interfaces/RTCoinInterface.sol";
import "Math/SafeMath.sol";

contract Router {

    using SafeMath for uint256;

    address public admin;
    /**CONSTANTS*/
    uint256 public constant DEFAULTLOCKUPTIME = 4 weeks;
    uint256 public constant MINSTAKE = 100000000000000000000; // 100 RTC ($12.50 at $0.125/rtc)
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
    bool    public locked;


    RTCoinInterface public rtI = RTCoinInterface(0x0994f9595d28429584bfb5fcbfea75b9c9ea2c24);

    struct StakerStruct {
        address addr;
        uint256 rtcStaked;
        uint256 deposit;
        uint256 khSec;
        uint256 depositDate;
        uint256 releaseDate;
        uint256 id;
        bool    enabled;
    }

    struct RewardStruct {
        uint256 ethRewarded;
        uint256 rtcRewarded;
    }

    mapping (bytes32 => bool)         private validOraclizeIds; // keep to private, helps reduce gas costs
    mapping (address => StakerStruct[]) public stakers;
    mapping (address => RewardStruct) public rewards;
    mapping (address => uint256) public numStakes;

    event RtcReward(address _staker, uint256 _amount);
    event EthReward(address _staker, uint256 _amount);

    modifier onlyAdmin() {
        require(msg.sender == admin);
        _;
    }

    function Router() {
        admin = msg.sender;
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
}