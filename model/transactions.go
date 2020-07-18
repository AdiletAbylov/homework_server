package model

import (
	"github.com/jinzhu/gorm"
)

// TransactionsModel is an ORM model for transactions table
type TransactionsModel struct {
	gorm.Model
	TransactionID string
	State         string
	Amount        float64
	SourceType    string
}

// NewTransactionWithEvent creates Transaction model with data from model.GameEvent
func NewTransactionWithEvent(event *Event) TransactionsModel {
	transaction := TransactionsModel{}
	transaction.Amount = event.Amount
	transaction.TransactionID = event.TransactionID
	transaction.SourceType = event.SourceType
	transaction.State = event.State
	return transaction
}
