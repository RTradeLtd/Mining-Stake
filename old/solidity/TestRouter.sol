pragma solidity 0.4.21;


import "Interfaces/RTCoinInterface.sol";
import "Math/SafeMath.sol";

contract TestRouter {
	using SafeMath for uint256;

	RTCoinInterface public rtI = RTCoinInterface(0x0994f9595d28429584bfb5fcbfea75b9c9ea2c24);

	event TestSent(address _recipient, uint256 _amount);

	function testRouteNoRequire(
		address[] _stakers,
		uint256[] _payouts)
		public
		returns (bool)
	{
		for (uint256 i = 0; i < _stakers.length; i++) {
			rtI.transferFrom(msg.sender, _stakers[i], _payouts[i]);
			emit TestSent(_stakers[i], _payouts[i]);
		}
		return true;
	}

	function testRouteRequire(
		address[] _stakers,
		uint256[] _payouts)
		public
		returns (bool)
	{
		for (uint256 i = 0; i < _stakers.length; i++) {
			require(rtI.transferFrom(msg.sender, _stakers[i], _payouts[i]));
			emit TestSent(_stakers[i], _payouts[i]);
		}
		return true;
	}

}