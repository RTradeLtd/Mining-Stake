pragma solidity 0.4.20;

import "./Interfaces/ERC20Interface.sol";
import "./Modules/Administration.sol";

/**
Used to route payments to stake holders

*/


contract PaymentRouter is Administration {

	event RTCSent(address _recipient, uint256 _amount);
	event ETHSent(address _recipient, uint256 _amount);

	function () payable {}

	function routeTokenPayments(
		address[] _recipients,
		address _tokenAddress,
		uint256 _tokensToSend)
		public
		onlyAdmin
		returns (bool)
	{
		ERC20Interface eI = ERC20Interface(_tokenAddress);
		for (uint256 i = 0; i < _recipients.length; i++) {
			if (_recipients[i] != address(0)) {
				RTCSent(_recipients[i], _tokensToSend);
				require(eI.transfer(_recipients[i], _tokensToSend));
			}
		}
	}

	function routeEthPayments(
		address[] 	_recipients,
		uint256 	_ethToSend)
		public
		onlyAdmin
		returns (bool)
	{
		for (uint256 i = 0; i < _recipients.length; i++) {
			if (_recipients[i] != address(0)) {
				require(this.balance >= _ethToSend);
				ETHSent(_recipients[i], _ethToSend);
				require(_recipients[i].send(_ethToSend));
			}
		}
	}

}