pragma solidity 0.4.19;

import "./Modules/Oraclize.sol";


contract DieselPrice is usingOraclize {
    
    uint public DieselPriceUSD;

    event newOraclizeQuery(string description);
    event newDieselPrice(string price);

    function DieselPrice() payable {}
    function () payable {}

    function __callback(bytes32 myid, string result) {
        if (msg.sender != oraclize_cbAddress()) throw;
        newDieselPrice(result);
        DieselPriceUSD = parseInt(result); // let's save it as $ cents
        // do something with the USD Diesel price
    }
    
    function update() payable {
        newOraclizeQuery("Oraclize query was sent, standing by for the answer..");
        oraclize_query("URL", "json(https://api.coinmarketcap.com/v1/ticker/ethereum/?convert=USD).0.price_usd");
    }
    
}

