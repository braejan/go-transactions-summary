package usecases

import "github.com/braejan/go-transactions-summary/internal/domain/user/entity"

// UserUsecases interface defines the methods that the user usecases must implement.

type UserUseCases interface {
	// GetByID returns a user by its ID.
	GetByID(ID int64) (user *entity.User, err error)
	// GetByEmail returns a user by its email.
	GetByEmail(email string) (user *entity.User, err error)
	// Create creates a new user.
	Create(ID int64, name string, email string) (err error)
	// Update updates a user.
	Update(ID int64, name string, email string) (err error)
}
