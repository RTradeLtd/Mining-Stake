from web3 import Web3, IPCProvider
import web3
import sys



# python3.6 stakeUtil.py <command>

if len(sys.argv) > 2 or len(sys.argv) < 2:
    print('Improper invocation')
    print('python3.6 stakeUtil.py <command>')
    exit()

command = sys.argv[1]
commands = ['--lockup-tokens', '--deregister']

if command not in commands:
    print('Invalid command listing, see following valid commands')
    msg = ''
    for c in commands:
        msg += '%s\t' % c
    print(msg)
    exit()



def lockupTokens(contract):
    rtcToStake = int(input('how much RTC do you wish to stake: '))
    timeToStake = int(input('for how many weeks do you wish to stakeL: '))
    assert rtcToStake > 0 and timeToStake > 4
    contract.transact({'from': '123'}).lockupTokens(rtcToStake, timeToStake)

w3 = Web3(IPCProvider('ipcpath'))

# unlock account
web3.personal.unlockAccount(web3.eth.accounts[0], 'password', 0)

# load the contract
contract = w3.eth.contract('data')


if command == commands[0]:
    lockupTokens(contract)