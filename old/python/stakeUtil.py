from web3 import Web3, IPCProvider
import web3
import sys
import json
from web3.middleware.pythonic import (
    pythonic_middleware,
    to_hexbytes,
)

size_extraData_for_poa = 200   # can change

web3 = Web3(IPCProvider('/home/solidity/.ethereum/geth.ipc'))
pythonic_middleware.__closure__[2].cell_contents['eth_getBlockByNumber'].args[1].args[0]['extraData'] = to_hexbytes(size_extraData_for_poa, variable_length=True)
pythonic_middleware.__closure__[2].cell_contents['eth_getBlockByHash'].args[1].args[0]['extraData'] = to_hexbytes(size_extraData_for_poa, variable_length=True)


# python3.6 stakeUtil.py <command>

if len(sys.argv) > 3 or len(sys.argv) < 3:
    print('Improper invocation')
    print('python3.6 stakeUtil.py <command> <contract-abi>')
    exit()

command = sys.argv[1]
contractAbi = sys.argv[2]
commands = ['--lockup-tokens', '--deregister']

with open(contractAbi, 'r') as fh:
    abi = json.load(fh)

if command not in commands:
    print('Invalid command listing, see following valid commands')
    msg = ''
    for c in commands:
        msg += '%s\t' % c
    print(msg)
    exit()



def lockupTokens(contract, w3):
    rtcToStake = int(input('how much RTC do you wish to stake: '))
    timeToStake = int(input('for how many weeks do you wish to stakeL: '))
    assert rtcToStake > 0 and timeToStake >= 4
    contract.transact({'from': '0xaFa0f5DbC3957d72Ba54787997D72E579fD419B9'}).lockupRtcTokens(rtcToStake, timeToStake)

w3 = Web3(IPCProvider('/home/solidity/.ethereum/geth.ipc'))

# unlock account
w3.personal.unlockAccount('0xaFa0f5DbC3957d72Ba54787997D72E579fD419B9', 'password123', 0)

# load the contract
contract = w3.eth.contract('0x1F60a5B6369179d7533Fc7f721EB8BB3dF0E95ea', abi=abi)


if command == commands[0]:
    lockupTokens(contract, w3)