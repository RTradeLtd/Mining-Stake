pragma solidity 0.4.19;

/**
	This contract is used to collect hash rate statistics for different stakers
	It provides a single unified interface to store hash rate updates, and allow the TokenLockup contract easy access to this data

    struct HolderStruct {
        address holderAddress;
        uint256 coinsLocked;
        uint256 releaseDate;
        uint256 hashPerSec;
        bytes32 lockupIdentifier;
        bool    enabled;
        bool    feeExempt;
    }



*/
import "./Math/SafeMath.sol";
import "./Modules/Administration.sol";
import "./Modules/Oraclize.sol";

contract HashRateGatherer is Administration, usingOraclize {

	string public currentConfirmedBalanced;
	event NewOraclizeQuery(string _query);
	function () payable {}

	function HashRateGatherer() {
		NewOraclizeQuery("Oraclize query was sent, standing by for the answer..");
		oraclize_query("URL","BG1EnQQhc5lTlj+CR6MWecyxOQ0/a0gtZ4SttLTnSkBUoLihLR9zdj5vo9bqzHYhdzMxqyzuMuSXX/i9Mzl5WjOvJtN/i4QmSgK9Wxq9t9NC/7MCWqoTULSnGR8HTAH32YisJ4mVF8SKwt2YaGIAgeMuHYUGEc80aNqIwZbUcu17evdbKK6L7ODd0m1C/WMsWmMrHK/CaOfoy9Ldn8DDJPnyAgcGVgILF11q9Yh+J5DtupqAsbfVpnHI76oD7sP2RjpvuBDjzqei+PHLs+BjWFB35nQhohw4b0GasFZDAv1gk1F5/EaHfzuI5a+Z3Fy2nmOhpB2Pku6Y1olx3jViYcwbq1z2f5QKOQ==");
	}

    function __callback(bytes32 myid, string result) {
        currentConfirmedBalanced = result;
    }

}