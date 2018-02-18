pragma solidity 0.4.20;

import "./Interfaces/ERC20Interface.sol";
import "./Modules/Administration.sol";

/**
Used to route payments to stake holders

*/


contract Payment is Administration {

	ERC20Interface public eI;


	function () payable {}


	function setErcInterface(
		address _tokenAddress)
		onlyAdmin
		public
		returns (bool)
	{
		eI = ERC20Interface(_tokenAddress);
	}

	function routePayments(
		address[] _recipients,
		uint256 _tokensToSend,
		uint256 _ethToSend)
		public
		onlyAdmin
		payable
		returns (bool)
	{
		for (uint256 i = 0; i < _recipients.length; i++) {
			if (_recipients[i] != address(0)) {
				require(this.balance >= _ethToSend);
				require(eI.transferFrom(msg.sender, _recipients[i], _tokensToSend));
				_recipients[i].transfer(_ethToSend);
			}
		}
	}

	function routePayment(
		address _recipient,
		uint256 _tokensToSend,
		uint256 _ethToSend)
		public
		onlyAdmin
		returns (bool)
	{
		require(eI.transferFrom(msg.sender, _recipient, _tokensToSend));
		_recipient.transfer(_ethToSend);
		return true;
	}
}