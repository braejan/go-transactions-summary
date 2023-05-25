package entity

import (
	"github.com/google/uuid"
)

// Account struct defines the account entity.
type Account struct {
	ID      uuid.UUID
	Balance int64
	UserID  int64
	Active  bool
}

// NewAccount returns a new Account instance.
func NewAccount(userID int64) (account *Account) {
	account = &Account{
		ID:      uuid.New(),
		Balance: 0,
		UserID:  userID,
		Active:  false,
	}
	return
}
