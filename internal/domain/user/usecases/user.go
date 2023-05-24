package usecases

import (
	"github.com/braejan/go-transactions-summary/internal/domain/user/entity"
	"github.com/braejan/go-transactions-summary/internal/domain/user/repository"
	"github.com/braejan/go-transactions-summary/internal/valueobject"
)

// userUsecases struct implements the UserUsecases interface.

type userUsecases struct {
	userRepo repository.UserRepository
}

// NewUserUsecases returns a new userUsecases instance.
func NewUserUsecases(userRepo repository.UserRepository) (usecases UserUsecases, err error) {

	if userRepo == nil {
		err = valueobject.ErrUserRepositoryIsNil
		return
	}
	usecases = &userUsecases{
		userRepo: userRepo,
	}
	return
}

// GetByID implements the UserUsecases interface method.
func (u *userUsecases) GetByID(ID int64) (user *entity.User, err error) {
	user, err = u.userRepo.GetByID(ID)
	return
}

// GetByEmail implements the UserUsecases interface method.
func (u *userUsecases) GetByEmail(email string) (user *entity.User, err error) {
	user, err = u.userRepo.GetByEmail(email)
	return
}

// Create implements the UserUsecases interface method.
func (u *userUsecases) Create(ID int64, name string, email string) (err error) {
	_, err = u.userRepo.GetByID(ID)
	if err == valueobject.ErrUserNotFound {
		// The user is not created.
		err = u.userRepo.Create(entity.NewUser(ID, name, email))
	} else if err == nil {
		// The user is already created.
		err = valueobject.ErrUserAlreadyCreated
		return
	}
	return
}

// Update implements the UserUsecases interface method.
func (u *userUsecases) Update(ID int64, name string, email string) (err error) {
	_, err = u.userRepo.GetByID(ID)
	if err != nil {
		// The user is not created.
		err = valueobject.ErrUserNotFound
		return
	}
	err = u.userRepo.Update(entity.NewUser(ID, name, email))
	return
}
