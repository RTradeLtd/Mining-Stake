pragma solidity 0.4.23;

import "./Math/SafeMath.sol";
import "./Modules/Administration.sol";

/*
    Used to facilitate oracle style updates of our smart contracts without having to rely on third-party products
    Currently only one update is supported, ETH-USD prices
*/

contract Oracle is Administration {
    using SafeMath for uint256;

    struct AuthorizedContractStruct {
        address contractAddress;
        uint256 updateFrequencyInHours;
        uint256 nextUpdate;
        bool    enabled;
    }

    mapping (address => AuthorizedContractStruct) public contracts;

    event AuthorizedContractAdded(address _contractAddress);

    function addAuthorizedContract(
        address _contractAddress,
        uint256 _updateFrequencyInHours,
        uint256 _nextUpdate)
        public
        onlyAdmin
        returns (bool)
    {
        AuthorizedContractStruct memory a;
        a.contractAddress = _contractAddress;
        a.updateFrequencyInHours = _updateFrequencyInHours;
        a.nextUpdate = now.add(_updateFrequencyInHours.mul(1 hours));
        contracts[_contractAddress] = a;
        emit AuthorizedContractAdded(_contractAddress);
        return true;
    }
}