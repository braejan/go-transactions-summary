package usecases

import "github.com/braejan/go-transactions-summary/internal/domain/account/entity"

// AccountUsecases interface defines the methods that the account usecases must implement.
type AccountUsecases interface {
	// GetByID returns an account by its ID.
	GetByID(ID int64) (account *entity.Account, err error)
	// GetByUserID returns an account by its user ID.
	GetByUserID(userID int64) (account *entity.Account, err error)
	// Create creates a new account.
	Create(userID int64) (err error)
	// Update updates an account.
	Update(ID int64, balance int64) (err error)
}
