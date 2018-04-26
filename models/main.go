package models

import (
	"github.com/RTradeLtd/Mining-Stake/database"
	"github.com/RTradeLtd/Mining-Stake/token_lockup"
)

// Manager is a helper method to wrap around everything
type Manager struct {
	ContractHandler *TokenLockup.TokenLockup
	BoltDB          *database.BoltDB
}
