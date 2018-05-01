package manager_test

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"

	"github.com/RTradeLtd/Mining-Stake/database"
	"github.com/RTradeLtd/Mining-Stake/manager"
)

const dbPath = "stakersTest.db"
const stakeBucket = "stakers"
const emailBucket = "emails"
const tokenLockupAddress = "0x5ae6c285eeb2e5a9234956cbcf9dea2c97c3a773"
const rpcURL = "http://127.0.0.1:8501"
const ipcPath = "/home/soliidty/DevNet/node1/geth.ipc"
const dev = true

func TestAuthenticateWithNetworkOnline(t *testing.T) {
	m := setupManager()
	err := m.AuthenticateWithNetwork()
	if err != nil {
		t.Error(err)
	}
}

func TestSetupDatabase(t *testing.T) {
	m := setupManager()
	err := m.Bolt.Setup(dbPath, stakeBucket, emailBucket)
	if err != nil {
		t.Fatal(err)
	}
	err = m.Bolt.UpdateStakeIDBucket(common.HexToAddress("0x5ae6c285eeb2e5a9234956cbcf9dea2c97c3a773"), big.NewInt(100))
	if err != nil {
		t.Error(err)
	}
	id := m.Bolt.RetrieveStakeIDInformationForAddress(common.HexToAddress("0x5ae6c285eeb2e5a9234956cbcf9dea2c97c3a773"))
	if id.Cmp(big.NewInt(100)) != 0 {
		t.Error("unexpected response for stake id")
	}
	_, err = m.Bolt.TFetchStakeIDs()
	if err != nil {
		t.Error(err)
	}
}

/* this wont complete for whatever reason
func TestUpdateStakeIDBucket(t *testing.T) {
	m := setupManager()
	err := m.Bolt.Setup(dbPath, stakeBucket, emailBucket)
	if err != nil {
		t.Fatal(err)
	}
	err = m.Bolt.UpdateStakeIDBucket(common.HexToAddress("0x5ae6c285eeb2e5a9234956cbcf9dea2c97c3a773"), big.NewInt(100))
	if err != nil {
		t.Error(err)
	}
}*/
/*
func TestRetrieveStakeIDInformationForAddress(t *testing.T) {
	m := setupManager()
	m.Bolt.Setup(dbPath, stakeBucket, emailBucket)
	id := m.Bolt.RetrieveStakeIDInformationForAddress(common.HexToAddress("0x5ae6c285eeb2e5a9234956cbcf9dea2c97c3a773"))
	if id.Cmp(big.NewInt(100)) != 0 {
		t.Error("unexpected response for stake id")
	}
}

func TestFetchStakeIDs(t *testing.T) {
	m := setupManager()
	m.Bolt.Setup(dbPath, stakeBucket, emailBucket)
	_, err := m.Bolt.TFetchStakeIDs()
	if err != nil {
		t.Error(err)
	}
}*/
func setupManager() *manager.Manager {

	manager := &manager.Manager{
		Key:            "",
		Password:       "",
		SendGridAPIKey: "",
		IpcPath:        ipcPath,
		RPCURL:         rpcURL,
		Bolt: &database.BoltDB{
			StakeIDBucketName: stakeBucket,
			EmailBucketName:   emailBucket,
		},
	}

	return manager
}
