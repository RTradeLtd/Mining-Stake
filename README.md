# Mining-Stake

Mining-Stake contains code related to the staking of RTC into smart contracts, which grants ethhash mining power at enabled locations.
The payouts are structured such that every 24 hours, a percentage of the usd value of the eth that was mined with the granted hash rate is given in RTC as a stake, directly to your wallet. Every 7 days, the total eth that was mined by the granted hash rate in the last 7 days is sent directly to your wallet. 

# Payout System

RTC stakes are *always* issued every 24 hours, but for now ETH payments are issued every 7 days. This is due to unpredicatable gas costs on the network, as well as depending on how much RTC you staked, the transaction fees for your payment may be larger than your payment. Thus, by sending every 7 days we limit the overall gas costs that will be incurred by a staker throughout their stake, but the gas costs will also be paid by RTrade instead of being deducted from your payment. This however is temporary, and we are exploring options to enable daily eth payments along with daily RTC paymetns in a cheap manner without burdening the transaction cost onto the staker.

Options being explored:
- Air drop channels
- Payment channels
- Precommitting large sum of funds to contracts, and update the balances every 24 hours
- state+payment channels
- sidechain payment routing



# Files

`go/src/stake_manager.go` - program used to administer a staking contract. Do not run when token_lockup_listener is running
`go/src/token_lockup_listener.go` - program used to listen for new stake deposits, and subsequently update our bolt db. Do not use while stake manager running
`solidity/src/TokenLockupV2.sol` - this is the actual contract used to stake in