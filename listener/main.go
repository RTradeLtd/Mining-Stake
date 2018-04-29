package listener

import (
	"fmt"
	"log"

	"github.com/RTradeLtd/Mining-Stake/TokenLockup"
	"github.com/RTradeLtd/Mining-Stake/manager"
)

// EventParser is used to parse evm events to notify us
// when we have a new stake
func EventParser(m *manager.Manager) {
	var ch = make(chan *TokenLockup.TokenLockupStakeDeposited)
	sub, err := m.ContractHandler.WatchStakeDeposited(nil, ch)
	if err != nil {
		log.Fatal("error creating subscription for stake deposited")
	}

	for {
		select {
		case err := <-sub.Err():
			fmt.Println("error parsing event ", err)
		case evLog := <-ch:
			m.SendNotificationEmail(evLog.Depositer, evLog.Amount, evLog.WeeksStaked, evLog.KhSec, evLog.Id)
			m.Bolt.UpdateStakeIDBucket(evLog.Depositer, evLog.Id)
		}
	}
}
