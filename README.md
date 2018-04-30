# Minining-Stake, Big Stake Boi Edition

Welcome to RTrade's new revamped stake manager, with none of the lags, delays, missing stakes, or other bullshit that the previous versions had. This is also design with automation in mind, and you will notice there is no command line application. 

# Limitations

Staking can't be initiated from a contract, and must be initiated from an EOA. Attempts to initiate stakes from a contract account will result in a transaction failure

For email notifications, we will *ONLY* send an email to the email listed in the latest known stake id
# What Is This?

Mining Stake is RTrade's Mining/Staking platform. By using the RTC coin, you can lock them into a smart contract, the "Token Lockup" contract. You can lock your coins in multiples of 1 week periods, with a minimum duration of 1 week. At lockup, hash rate for the ethhash algorithm is calculated based on your stake. This hash rate is then used to mine ethereum at our mining farms. You receive the ETH you mine, and RTC as a staking bonus. 

# Payout Information

Currently there are two parts to the payout periods. Each day, your RTC stake is sent directly to the address you used to stake. This is equivalent to 10% of the USD evaluation of the eth that was mined with your hash rate in a 24 hour period.

In the ideal world, you would receive payouts everyday, but unfortunately due to gas costs this is not economically feasible at this time, so we have opted to issue eth once a week, and RTC everyday. We are exploring other avenues to get you your payouts immediately, without having to charge a boat load of fees.

# Staking Interface [BETA]

We currently have a beta of the staking interface online at https://stake.beta.rtradetecnologies.com 

Through this interface you can use metamask to interact with the smart contract, and view your rewards. However, myetherwallet, or mist can also be used.

*NOTE* Until launch, any stakes initiated with the web interface will populate the smart contract, but will not update our backend. So you will not receive any sort of compensation


# TODO

> Update RTC-ETH to reflect new RTC-CAD valuation
> Update interface for token lockup
> create interface for oracle contract
> create golang oracle backend
> Test new contract