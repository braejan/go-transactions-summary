package usecases

import (
	"log"

	"github.com/braejan/go-transactions-summary/internal/domain/account/entity"
	accRepo "github.com/braejan/go-transactions-summary/internal/domain/account/repository"
	userRepo "github.com/braejan/go-transactions-summary/internal/domain/user/repository"
	"github.com/braejan/go-transactions-summary/internal/valueobject/account"
	"github.com/braejan/go-transactions-summary/internal/valueobject/user"
	"github.com/google/uuid"
)

// accountUsecases struct implements the AccountUsecases interface.
type accountUsecases struct {
	accountRepo accRepo.AccountRepository
	userRepo    userRepo.UserRepository
}

// NewAccountUseCases returns a new accountUsecases instance.
func NewAccountUseCases(accountRepo accRepo.AccountRepository, usrRepo userRepo.UserRepository) (usecases AccountUseCases, err error) {
	if accountRepo == nil {
		err = account.ErrAccountRepositoryIsNil
		return
	}
	if usrRepo == nil {
		err = user.ErrUserRepositoryIsNil
		return
	}
	usecases = &accountUsecases{
		accountRepo: accountRepo,
		userRepo:    usrRepo,
	}
	return
}

// Usecases interface implementation:

// GetByID implements the AccountUsecases interface method.
func (u *accountUsecases) GetByID(ID string) (acc entity.Account, err error) {
	accID, err := uuid.Parse(ID)
	if err != nil {
		err = account.ErrProcessingAccountID
		return
	}
	accAux, err := u.accountRepo.GetByID(accID)
	if err != nil {
		return
	}
	acc = *accAux
	return
}

// GetByUserID implements the AccountUsecases interface method.
func (u *accountUsecases) GetByUserID(userID int64) (acc entity.Account, err error) {
	accAux, err := u.accountRepo.GetByUserID(userID)
	if err != nil {
		return
	}
	acc = *accAux
	return
}

// Create implements the AccountUsecases interface method.
func (u *accountUsecases) Create(userID int64) (err error) {
	log.Println("Creating account for user", userID)
	// Check if the user exists.
	_, err = u.userRepo.GetByID(userID)
	if err != nil {
		// The user is not created.
		err = user.ErrUserNotFound
		return
	}
	// Check if the user already has an account.
	acc, err := u.accountRepo.GetByUserID(userID)
	if err != nil && err != account.ErrAccountNotFound {
		return
	}
	if acc != nil {
		// The user already has an account.
		err = account.ErrAccountAlreadyCreated
		return
	}
	// Create the account.
	acc = entity.NewAccount(userID)
	err = u.accountRepo.Create(acc)
	return
}

// Update implements the AccountUsecases interface method.
func (u *accountUsecases) Update(ID string, balance float64, active bool) (err error) {
	accID, err := uuid.Parse(ID)
	if err != nil {
		err = account.ErrProcessingAccountID
		return
	}
	// Check if the account exists.
	acc, err := u.accountRepo.GetByID(accID)
	if err != nil {
		return
	}
	// Update the account.
	acc.Balance = balance
	acc.Active = active
	err = u.accountRepo.Update(acc)
	return
}
