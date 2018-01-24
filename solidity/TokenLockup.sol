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
    uint256 public constant MINIMUMLOCKUPAMOUNT = 100;


    uint256 public ethUSD;
    // hot wallet used to collect sign up fees,
    address public rtcHotWallet;
    // 0.1 ETH
    uint256 public signUpFee = 100000000000000000;
    // floating value, how many RTC for a single hash rate a second, starts out at 4.5RTC = 1 hash/sec
    uint256 public rtcPerHashSecond = 4500000000000000000;
    uint256 public stakerCount;

    bool    private locked;


    RTCoinInterface private rtI;

    struct HolderStruct {
        address holderAddress;
        uint256 coinsLocked;
        uint256 releaseDate;
        uint256 hashPerSec;
        bool    enabled;
        bool    feeExempt;
    }

    mapping (address => HolderStruct) public holders;
    mapping (address => bool)   public registeredHolders;
    mapping (address => uint256) public holderBalances; // for lockupTokens
    mapping (address => uint256) public holderRewards;
    mapping (address => uint256) public memberNumber;
    mapping (bytes32 => bool)   private validOraclizeIds; // keep to private, helps reduce gas costs

    event RTCoinInterfaceSet(address indexed _rtcContractAddress, bool indexed _rtcInteraceSet);
    event MiningRewardDeposited(address indexed _miningPayoutRewardee, uint256 _amountInRtcPaidOut, bool indexed _miningRewardPayout);
    event LockupWithdrawn(address indexed _withdrawee, uint256 _amountWithdrawn, bool indexed _lockupWithdrawn);
    event LockupDeposited(address indexed _lockee, uint256 _amountLocked, uint256 indexed _lockupDuration, uint256 indexed _hashesPerSecond, bool _tokensLockedUp);

    event NewOraclizeQuery(string result);
    event EthUsdPriceUpdated(uint256 price);
    event SignUpFeeUpdated(uint256 fee);

    modifier registeredUser(address _user) {
        require(_user != address(0x0) && registeredHolders[_user] == true);
        _;
    }

    modifier nonRegisteredUser(address _user) {
        require(_user != address(0x0) && registeredHolders[_user] != true);
        _;
    }

    modifier isEnabledUser(address _user) {
        require(_user != address(0x0) && holders[_user].enabled == true);
        _;
    }

    modifier validReleaseDate(address _user) {
        require(_user != address(0x0) && now > holders[_user].releaseDate);
        _;
    }

    modifier validBalance(address _holder) {
        require(holderBalances[_holder] > 0);
        _;
    }

    modifier hasValidBalance() {
        require(holderBalances[msg.sender] > 0 && holders[msg.sender].coinsLocked > 0);
        _;
    }

    modifier notLocked() {
        require(!locked);
        _;
    }

    function TokenLockup() payable {
        bytes32 id = oraclize_query("URL", "json(https://api.coinmarketcap.com/v1/ticker/ethereum/?convert=USD).0.price_usd");
        validOraclizeIds[id] = true;
    }

    function () payable {}


    function __callback(bytes32 myid, string result) {
        locked = true;
        require(msg.sender == oraclize_cbAddress());
        require(validOraclizeIds[myid]);
        ethUSD = parseInt(result);
        uint256 oneEth = 1 ether;
        signUpFee = oneEth.div(ethUSD);
        signUpFee = signUpFee.div(1 ether);
        signUpFee = signUpFee.mul(10);
        EthUsdPriceUpdated(ethUSD);
        SignUpFeeUpdated(signUpFee);
        delete validOraclizeIds[myid];
        locked = false;
        update();
    }

    function update() payable {
        require(this.balance >=oraclize_getPrice("URL"));
        NewOraclizeQuery("Oraclize query was sent, standing by for the answer..");
        bytes32 _id = oraclize_query(600, "URL", "json(https://api.coinmarketcap.com/v1/ticker/ethereum/?convert=USD).0.price_usd");
        validOraclizeIds[_id] = true;
    }

    function lockupTokens(
            uint256 _amountToLockup,
            uint256 _lockupDurationInWeeks
    )
        public
        payable
        nonRegisteredUser(msg.sender)
        notLocked
        returns (bool)
    {
        require(_amountToLockup >= MINIMUMLOCKUPAMOUNT);
        require(_lockupDurationInWeeks >= 4);
        uint256 lockupDuration = _lockupDurationInWeeks * 1 weeks;
        // check if they are fee exempt
        if (stakerCount > 100) {
            require(msg.value == signUpFee);
            rtcHotWallet.transfer(msg.value);
        }
        uint256 _hashSecond = _amountToLockup.div(rtcPerHashSecond);
        holders[msg.sender].holderAddress = msg.sender;
        holders[msg.sender].coinsLocked = _amountToLockup;
        holders[msg.sender].releaseDate = lockupDuration;
        holders[msg.sender].hashPerSec = _hashSecond;
        holders[msg.sender].enabled = true;
        holderBalances[msg.sender] = holderBalances[msg.sender].add(_amountToLockup);
        registeredHolders[msg.sender] = true;
        require(rtI.transferFrom(msg.sender, this, _amountToLockup));
        LockupDeposited(msg.sender, _amountToLockup, lockupDuration, _hashSecond, true);
        return true;
    }

    function setRtI(
        address _rtcAddress
    )
        public
        onlyOwner
        returns (bool)
    {
        rtI = RTCoinInterface(_rtcAddress);
        RTCoinInterfaceSet(_rtcAddress, true);
        return true;
    }


    function depositReward(
        address _holder,
        uint256 _amountToDeposit
    )
        public
        registeredUser(_holder)
        isEnabledUser(_holder)
        validBalance(_holder)
        returns (bool)
    {
        require(_amountToDeposit > 0);
        holderRewards[_holder] = holderRewards[_holder].add(_amountToDeposit);
        MiningRewardDeposited(_holder, _amountToDeposit, true);
        require(rtI.transferFrom(msg.sender, this, _amountToDeposit));
        return true;
    }

    function retrieveTokens()
        public
        registeredUser(msg.sender)
        isEnabledUser(msg.sender)
        validReleaseDate(msg.sender)
        hasValidBalance
        returns (bool)
    {

        uint256 amountToWithdraw = holderBalances[msg.sender];
        require(amountToWithdraw == holders[msg.sender].coinsLocked);
        holders[msg.sender].coinsLocked = 0;
        holderBalances[msg.sender] = 0; // reset balance, prevent reentrancy
        holders[msg.sender].enabled = false;
        require(rtI.transfer(msg.sender, amountToWithdraw));
        LockupWithdrawn(msg.sender, amountToWithdraw, true);
        return true;
    }


    /**INTERNALS*/


    /**GETTERS*/
    function getContractRtcBalance()
        public
        view
        returns (uint256)
    {
        return rtI.balanceOf(this);
    }

    function getHolderStruct(
        address _holderAddress
    )
        public
        view
        returns (address, uint256, uint256, bool)
    {
        HolderStruct memory s = holders[_holderAddress];
        return (s.holderAddress, s.coinsLocked, s.releaseDate, s.enabled);
    }

}
