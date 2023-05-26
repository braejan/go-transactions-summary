package usecases

import "github.com/braejan/go-transactions-summary/internal/domain/account/entity"

// AccountUseCases interface defines the methods that the account usecases must implement.
type AccountUseCases interface {
	// GetByID returns an account by its ID.
	GetByID(ID string) (acc *entity.Account, err error)
	// GetByUserID returns an account by its user ID.
	GetByUserID(userID int64) (acc *entity.Account, err error)
	// Create creates a new account.
	Create(userID int64) (err error)
	// Update updates an account.
	Update(ID string, balance float64, active bool) (err error)
}
