pragma solidity 0.4.23;

import "./Modules/Administration.sol";
import "./Math/SafeMath.sol";
import "./Interfaces/RTCoinInterface.sol";
import "./Interfaces/ERC20Interface.sol";

contract RTCETH is Administration {
    using SafeMath for uint256;

    address public hotWallet;
    uint256 public ethUSD;
    uint256 public ethPerRtc;
    bool   public locked;

    RTCoinInterface public rtI;

    event EthUsdPriceUpdated(uint256 _ethUSD);
    event EthPerRtcUpdated(uint256 _ethPerRtc);
    event RtcPurchased(uint256 _rtcPurchased);

    constructor() public payable {
        rtI = RTCoinInterface(address(0));
    }

    function () public payable {}

    function updateEthPrice(
        uint256 _ethUSD)
        public
        onlyAdmin
        returns (bool)
    {
        ethUSD = _ethUSD;
        uint256 oneEth = 1 ether;
        uint256 oneUsdOfEth = oneEth.div(ethUSD);
        ethPerRtc = oneUsdOfEth.div(8);
        emit EthUsdPriceUpdated(ethUSD);
        emit EthPerRtcUpdated(ethPerRtc);
        return true;
    }

    function setHotWallet(
        address _hotWalletAddress)
        public
        onlyAdmin
        returns (bool)
    {
        hotWallet = _hotWalletAddress;
        return true;
    }

    function buyRtc()
        public
        payable
        returns (bool)
    {
        require(hotWallet != address(0));
        require(msg.value > 0);
        uint256 rtcPurchased = msg.value.div(ethPerRtc);
        emit RtcPurchased(rtcPurchased);
        hotWallet.transfer(msg.value);
        require(rtI.transfer(msg.sender, rtcPurchased));
        return true;
    }
}