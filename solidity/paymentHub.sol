pragma solidity 0.4.20;

import "./Interfaces/ERC20Interface.sol";
import "./Modules/Administration.sol";

/**
Used to route payments to stake holders

*/


contract PaymentRouter is Administration {

	function () payable {}

	function routePayments(
		address[] _recipients,
		address _tokenAddress,
		uint256 _tokensToSend,
		uint256 _ethToSend)
		public
		onlyAdmin
		returns (bool)
	{
		ERC20Interface eI = ERC20Interface(_tokenAddress);
		for (uint256 i = 0; i < _recipients.length; i++) {
			if (_recipients[i] != address(0)) {
				require(this.balance >= _ethToSend);
				require(eI.transfer(_recipients[i], _tokensToSend));
				_recipients[i].transfer(_ethToSend);
			}
		}
	}

}