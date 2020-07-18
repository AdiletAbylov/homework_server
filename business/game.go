package business

import (
	"fmt"
	"hw_server/model"
	"hw_server/repo"
)

// HandleGameEvent handles given game event.
func HandleGameEvent(event *model.Event) error {
	if repo.HasItemsWithTransactionID(event.TransactionID) {
		return fmt.Errorf("Transaction with ID:" + event.TransactionID + "is already exist.")
	}
	switch event.State {
	case model.WinState:
		repo.IncreaseUserBalance(event.Amount)
	case model.LostState:
		repo.DecreaseUserBalance(event.Amount)
	}

	repo.LogTransaction(event)
	return nil
}
