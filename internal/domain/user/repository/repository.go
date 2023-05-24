package repository

import "github.com/braejan/go-transactions-summary/internal/domain/user/entity"

// UserRepository interface defines the methods that the user repository must implement.

type UserRepository interface {
	// GetByID returns a user by its ID.
	GetByID(id int64) (user *entity.User, err error)
	// GetByEmail returns a user by its email.
	GetByEmail(email string) (user *entity.User, err error)
	// Create creates a new user.
	Create(user *entity.User) (err error)
	// Update updates a user.
	Update(user *entity.User) (err error)
}
