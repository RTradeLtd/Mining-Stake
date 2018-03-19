pragma solidity 0.4.21;

/**
RTCoin Interface
*/

interface RTCoinInterface {
    
    function freezeTransfers() public returns (bool);

    function thawTransfers() public returns (bool);

    function transfer(address _recipient, uint256 _amount) public returns (bool transferred);

    function transferFrom(address _owner, address _recipient, uint256 _amount) public returns (bool transferredFrom);

    function approve(address _spender, uint256 _amount) public returns (bool approved);
    /**GETTERS */

    function totalSupply() public view returns (uint256);

    function balanceOf(address _holder) public view returns (uint256);

    function allowance(address _owner, address _spender) public view returns (uint256);
}
