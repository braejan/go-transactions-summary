package usecases

import (
	"github.com/braejan/go-transactions-summary/internal/domain/user/entity"
	"github.com/braejan/go-transactions-summary/internal/domain/user/repository"
	"github.com/braejan/go-transactions-summary/internal/valueobject/user"
)

// userUsecases struct implements the UserUseCases interface.

type userUsecases struct {
	userRepo repository.UserRepository
}

// NewUserUseCases returns a new userUsecases instance.
func NewUserUseCases(userRepo repository.UserRepository) (usecases UserUseCases, err error) {

	if userRepo == nil {
		err = user.ErrUserRepositoryIsNil
		return
	}
	usecases = &userUsecases{
		userRepo: userRepo,
	}
	return
}

// GetByID implements the UserUseCases interface method.
func (u *userUsecases) GetByID(ID int64) (user *entity.User, err error) {
	user, err = u.userRepo.GetByID(ID)
	return
}

// GetByEmail implements the UserUseCases interface method.
func (u *userUsecases) GetByEmail(email string) (user *entity.User, err error) {
	user, err = u.userRepo.GetByEmail(email)
	return
}

// Create implements the UserUseCases interface method.
func (u *userUsecases) Create(ID int64, name string, email string) (err error) {
	_, err = u.userRepo.GetByID(ID)
	if err == user.ErrUserNotFound {
		// The user is not created.
		err = u.userRepo.Create(entity.NewUser(ID, name, email))
	} else if err == nil {
		// The user is already created.
		err = user.ErrUserAlreadyCreated
		return
	}
	return
}

// Update implements the UserUseCases interface method.
func (u *userUsecases) Update(ID int64, name string, email string) (err error) {
	_, err = u.userRepo.GetByID(ID)
	if err != nil {
		// The user is not created.
		err = user.ErrUserNotFound
		return
	}
	err = u.userRepo.Update(entity.NewUser(ID, name, email))
	return
}
