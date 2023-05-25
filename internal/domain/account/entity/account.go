package entity

import (
	"github.com/google/uuid"
)

// Account struct defines the account entity.
type Account struct {
	ID      uuid.UUID
	balance int64
	userID  int64
	active  bool
}

// NewAccount returns a new Account instance.
func NewAccount(userID int64) (account *Account) {
	account = &Account{
		ID:      uuid.New(),
		balance: 0,
		userID:  userID,
		active:  false,
	}
	return
}
