package main

import (
	"hw_server/config"
	"hw_server/repo"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestInitDB tests db initialization, table creation and default user creation
func TestInitDB(t *testing.T) {
	config.ReadConfigs()
	repo.InitDB(config.DBConnectString())
	user := repo.DefaultUser()
	assert.Equal(t, user.Username, "Vasya Pupkin", "Username of default user should be Vasya Pupkin")
}

func TestIncreaseUserBalance(t *testing.T) {
	config.ReadConfigs()
	repo.InitDB(config.DBConnectString())
	repo.ResetDefaultUser()
	repo.IncreaseUserBalance(100)
	user := repo.DefaultUser()
	assert.Equal(t, user.Balance, 100.0, "Balance value should be 100")
}

func TestDecreaseUserBalance(t *testing.T) {
	config.ReadConfigs()
	repo.InitDB(config.DBConnectString())
	repo.ResetDefaultUser()
	repo.IncreaseUserBalance(100)
	repo.DecreaseUserBalance(50)
	user := repo.DefaultUser()
	assert.Equal(t, user.Balance, 50.0, "Balance value should be 50")
	//
	// User balance can't be negative, minimal possible value is zero
	repo.DecreaseUserBalance(100)
	user = repo.DefaultUser()
	assert.Equal(t, user.Balance, 0.0, "Balance value should be 0")
}
