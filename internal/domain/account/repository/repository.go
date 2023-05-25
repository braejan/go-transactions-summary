package repository

import "github.com/braejan/go-transactions-summary/internal/domain/account/entity"

// AccountRepository interface defines the methods that the account repository must implement.
type AccountRepository interface {
	// GetByID returns an account by its ID.
	GetByID(id int64) (account *entity.Account, err error)
	// GetByUserID returns an account by its user ID.
	GetByUserID(userID int64) (account *entity.Account, err error)
	// Create creates a new account.
	Create(account *entity.Account) (err error)
	// Update updates an account.
	Update(account *entity.Account) (err error)
}
