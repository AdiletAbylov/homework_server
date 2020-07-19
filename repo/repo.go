package repo

import (
	"hw_server/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Database is the ref to the GORM DB connection
var database *gorm.DB

// InitDB connects to database, makes automigrations and seeds default data
func InitDB(connectString string) {
	conn, err := gorm.Open("postgres", connectString)
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}
	database = conn

	database.AutoMigrate(&model.TransactionsModel{})
	database.AutoMigrate(&model.UserBalanceModel{})

	createDefautUser()
}

// Close closes database connection
func Close() {
	if database != nil {
		database.Close()
	}
}

// DefaultUser returns default user record
func DefaultUser() model.UserBalanceModel {
	user := model.UserBalanceModel{}

	database.First(&user)
	return user
}

// IncreaseUserBalance increases default user's balance on given value.
func IncreaseUserBalance(value float64) {
	user := DefaultUser()
	user.Balance += value
	database.Save(user)
}

// DecreaseUserBalance decreases default user's balance on given value.
// If new value is less than zero, puts zero value to balance.
func DecreaseUserBalance(value float64) {
	user := DefaultUser()
	balance := user.Balance
	if balance-value >= 0 {
		user.Balance -= value

	} else {
		user.Balance = 0
	}
	database.Save(user)
}

//HasItemsWithTransactionID checks if there are items with given transaction ID in database
func HasItemsWithTransactionID(transactionID string) bool {
	transaction := model.TransactionsModel{}
	database.Where("transaction_id=?", transactionID).First(&transaction)
	if transaction.TransactionID == "" {
		return false
	}
	return true
}

// LogTransaction saves game events in database
func LogTransaction(event *model.Event) {
	transaction := model.NewTransactionWithEvent(event)
	database.Create(&transaction)
}

// ResetDefaultUser resets default user's balance to zero
func ResetDefaultUser() {
	defaultUser := model.UserBalanceModel{}
	database.First(&defaultUser)
	defaultUser.Balance = 0
	database.Save(&defaultUser)
}
func createDefautUser() {
	defaultUser := model.UserBalanceModel{}
	database.First(&defaultUser)

	if defaultUser.Username != "Vasya Pupkin" {
		defaultUser = model.UserBalanceModel{Username: "Vasya Pupkin"}
		database.Create(&defaultUser)
	}
}
