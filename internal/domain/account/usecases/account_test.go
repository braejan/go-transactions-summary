package usecases_test

import (
	"errors"
	"testing"

	"github.com/braejan/go-transactions-summary/internal/domain/account/entity"
	accMock "github.com/braejan/go-transactions-summary/internal/domain/account/repository/mock"
	"github.com/braejan/go-transactions-summary/internal/domain/account/usecases"
	userEntity "github.com/braejan/go-transactions-summary/internal/domain/user/entity"
	userMock "github.com/braejan/go-transactions-summary/internal/domain/user/repository/mock"
	"github.com/braejan/go-transactions-summary/internal/valueobject/account"
	"github.com/braejan/go-transactions-summary/internal/valueobject/user"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// TestNewAccountUseCasesWithUserRepoNil tests the NewAccountUseCases function with a nil user repository.
func TestNewAccountUseCasesWithUserRepoNil(t *testing.T) {
	// Given a valid account repository.
	accRepo := accMock.NewMockAccountRepository()
	// When NewAccountUseCases is called with a nil user repository.
	_, err := usecases.NewAccountUseCases(accRepo, nil)
	// Then the error ErrUserRepositoryIsNil is returned.
	assert.EqualError(t, err, user.ErrUserRepositoryIsNil.Error())
}

// TestNewAccountUseCasesWithAccountRepoNil tests the NewAccountUseCases function with a nil account repository.
func TestNewAccountUseCasesWithAccountRepoNil(t *testing.T) {
	// Given a valid user repository.
	userRepo := userMock.NewMockUserRepository()
	// When NewAccountUseCases is called with a nil account repository.
	_, err := usecases.NewAccountUseCases(nil, userRepo)
	// Then the error ErrAccountRepositoryIsNil is returned.
	assert.EqualError(t, err, account.ErrAccountRepositoryIsNil.Error())
}

// TestNewAccountUseCasesWithValidRepos tests the NewAccountUseCases function with valid repositories.
func TestNewAccountUseCasesWithValidRepos(t *testing.T) {
	// Given valid repositories.
	accRepo := accMock.NewMockAccountRepository()
	userRepo := userMock.NewMockUserRepository()
	// When NewAccountUseCases is called with valid repositories.
	usecases, err := usecases.NewAccountUseCases(accRepo, userRepo)
	// Then no error is returned.
	assert.NoError(t, err)
	// And the usecases is not nil.
	assert.NotNil(t, usecases)
}

// TestGetByIDWithInvalidID tests the GetByID method with an invalid ID.
func TestGetByIDWithInvalidID(t *testing.T) {
	// Given valid repositories.
	accRepo := accMock.NewMockAccountRepository()
	userRepo := userMock.NewMockUserRepository()
	// And a valid usecases.
	usecases, err := usecases.NewAccountUseCases(accRepo, userRepo)
	assert.NoError(t, err)
	assert.NotNil(t, usecases)
	// When GetByID is called with an invalid ID.
	_, err = usecases.GetByID("invalid")
	// Then the error ErrProcessingAccountID is returned.
	assert.EqualError(t, err, account.ErrProcessingAccountID.Error())
}

// TestGetByIDSucess tests the GetByID method with a valid ID.
func TestGetByIDSucess(t *testing.T) {
	// Given valid repositories.
	accRepo := accMock.NewMockAccountRepository()
	userRepo := userMock.NewMockUserRepository()
	// And a valid usecases.
	usecases, err := usecases.NewAccountUseCases(accRepo, userRepo)
	assert.NoError(t, err)
	assert.NotNil(t, usecases)
	// And a valid account
	accExpected := entity.NewAccount(int64(1))
	// And a mocked response when GetByID is called.
	accRepo.On("GetByID", accExpected.ID).Return(accExpected, nil)
	// When GetByID is called with a valid ID.
	acc, err := usecases.GetByID(accExpected.ID.String())
	// Then no error is returned.
	assert.Nil(t, err)
	// And the account returned is the expected.
	assert.Equal(t, accExpected, acc)
}

// TestGetByUserIDWithInvalidUserID tests the GetByUserID method with an invalid user ID.
func TestGetByUserIDWithInvalidUserID(t *testing.T) {
	// Given valid repositories.
	accRepo := accMock.NewMockAccountRepository()
	userRepo := userMock.NewMockUserRepository()
	// And a valid usecases.
	usecases, err := usecases.NewAccountUseCases(accRepo, userRepo)
	assert.NoError(t, err)
	assert.NotNil(t, usecases)
	// And a mocked response when GetByUserID is called.
	accRepo.On("GetByUserID", int64(0)).Return(nil, user.ErrUserNotFound)
	// When GetByUserID is called with an invalid user ID.
	_, err = usecases.GetByUserID(int64(0))
	// Then the error ErrProcessingUserID is returned.
	assert.EqualError(t, err, user.ErrUserNotFound.Error())
}

// TestCreateWithInvalidUserID tests the Create method with an invalid user ID.
func TestCreateWithInvalidUserID(t *testing.T) {
	// Given valid repositories.
	accRepo := accMock.NewMockAccountRepository()
	userRepo := userMock.NewMockUserRepository()
	// And a valid usecases.
	usecases, err := usecases.NewAccountUseCases(accRepo, userRepo)
	assert.NoError(t, err)
	assert.NotNil(t, usecases)
	// And a mocked response when userRepo.GetByID is called.
	userRepo.On("GetByID", int64(0)).Return(nil, user.ErrUserNotFound)
	// When Create is called with an invalid user ID.
	err = usecases.Create(int64(0))
	// Then the error ErrProcessingUserID is returned.
	assert.EqualError(t, err, user.ErrUserNotFound.Error())
}

// TestCreateWithExistingAccount tests the Create method with an existing account.
func TestCreateWithExistingAccount(t *testing.T) {
	// Given valid repositories.
	accRepo := accMock.NewMockAccountRepository()
	userRepo := userMock.NewMockUserRepository()
	// And a valid usecases.
	usecases, err := usecases.NewAccountUseCases(accRepo, userRepo)
	assert.NoError(t, err)
	assert.NotNil(t, usecases)
	// And a mocked response when userRepo.GetByID is called.
	user := userEntity.NewUser(int64(1), "John Doe", "john.doe@amazingemail.com")
	userRepo.On("GetByID", user.ID).Return(user, nil)
	// And a mocked response when accRepo.GetByUserID is called.
	acc := entity.NewAccount(user.ID)
	accRepo.On("GetByUserID", user.ID).Return(acc, nil)
	// When Create is called with an existing account.
	err = usecases.Create(user.ID)
	// Then the error is not nil
	assert.NotNil(t, err)
	// Then the error ErrAccountAlreadyExists is returned.
	assert.Equal(t, err, account.ErrAccountAlreadyCreated)
}

// TestCreateWithErrorGettingAccount tests the Create method with an error getting the account.
func TestCreateWithErrorGettingAccount(t *testing.T) {
	// Given valid repositories.
	accRepo := accMock.NewMockAccountRepository()
	userRepo := userMock.NewMockUserRepository()
	// And a valid usecases.
	usecases, err := usecases.NewAccountUseCases(accRepo, userRepo)
	assert.NoError(t, err)
	assert.NotNil(t, usecases)
	// And a mocked response when userRepo.GetByID is called.
	user := userEntity.NewUser(int64(1), "John Doe", "john.doe@amazinemail.com")
	userRepo.On("GetByID", user.ID).Return(user, nil)
	// And a mocked response when accRepo.GetByUserID is called.
	accRepo.On("GetByUserID", user.ID).Return(nil, errors.New("error"))
	// When Create is called with an error getting the account.
	err = usecases.Create(user.ID)
	// Then the error is not nil
	assert.NotNil(t, err)
}

// TestCreateWithRepoError tests the Create method with an error saving the account.
func TestCreateWithRepoError(t *testing.T) {
	// Given valid repositories.
	accRepo := accMock.NewMockAccountRepository()
	userRepo := userMock.NewMockUserRepository()
	// And a valid usecases.
	usecases, err := usecases.NewAccountUseCases(accRepo, userRepo)
	assert.NoError(t, err)
	assert.NotNil(t, usecases)
	// And a mocked response when userRepo.GetByID is called.
	user := userEntity.NewUser(int64(1), "John Doe", "john.doe@amazinemail.com")
	userRepo.On("GetByID", user.ID).Return(user, nil)
	// And a mocked response when accRepo.GetByUserID is called.
	accRepo.On("GetByUserID", user.ID).Return(nil, account.ErrAccountNotFound)
	// And a mocked response when accRepo.Create is called.
	accRepo.On("Create", mock.Anything).Return(errors.New("error"))
	// When Create is called with an error saving the account.
	err = usecases.Create(user.ID)
	// Then the error is not nil
	assert.NotNil(t, err)
}

// TestCreateWithSuccess tests the Create method with success.
func TestCreateWithSuccess(t *testing.T) {
	// Given valid repositories.
	accRepo := accMock.NewMockAccountRepository()
	userRepo := userMock.NewMockUserRepository()
	// And a valid usecases.
	usecases, err := usecases.NewAccountUseCases(accRepo, userRepo)
	assert.NoError(t, err)
	assert.NotNil(t, usecases)
	// And a mocked response when userRepo.GetByID is called.
	user := userEntity.NewUser(int64(1), "John Doe", "john.doe@amazingemail.com")
	userRepo.On("GetByID", user.ID).Return(user, nil)
	// And a mocked response when accRepo.GetByUserID is called.
	accRepo.On("GetByUserID", user.ID).Return(nil, account.ErrAccountNotFound)
	// And a mocked response when accRepo.Create is called.
	accRepo.On("Create", mock.Anything).Return(nil)
	// When Create is called with success.
	err = usecases.Create(user.ID)
	// Then the error is nil
	assert.Nil(t, err)
}

// TestUpdateWithInvalidAccountID tests the Update method with an invalid account ID.
func TestUpdateWithInvalidAccountID(t *testing.T) {
	// Given valid repositories.
	accRepo := accMock.NewMockAccountRepository()
	userRepo := userMock.NewMockUserRepository()
	// And a valid usecases.
	usecases, err := usecases.NewAccountUseCases(accRepo, userRepo)
	assert.NoError(t, err)
	assert.NotNil(t, usecases)
	// When Update is called with an invalid account ID.
	err = usecases.Update("invalid", 1.00, true)
	// Then the error ErrProcessingAccountID is returned.
	assert.EqualError(t, err, account.ErrProcessingAccountID.Error())
}

// TestUpdateWithAccountNotFound tests the Update method with an account not found.
func TestUpdateWithAccountNotFound(t *testing.T) {
	// Given valid repositories.
	accRepo := accMock.NewMockAccountRepository()
	userRepo := userMock.NewMockUserRepository()
	// And a valid usecases.
	usecases, err := usecases.NewAccountUseCases(accRepo, userRepo)
	assert.NoError(t, err)
	assert.NotNil(t, usecases)
	// And a valid uuid.UUID accID
	accID := uuid.New()
	// And a mocked response when accRepo.GetByID is called.
	accRepo.On("GetByID", accID).Return(nil, account.ErrAccountNotFound)
	// When Update is called with an account not found.
	err = usecases.Update(accID.String(), 1.00, true)
	// Then the error ErrAccountNotFound is returned.
	assert.EqualError(t, err, account.ErrAccountNotFound.Error())
}

// TestUpdateWithRepoError tests the Update method with an error saving the account.
func TestUpdateWithRepoError(t *testing.T) {
	// Given valid repositories.
	accRepo := accMock.NewMockAccountRepository()
	userRepo := userMock.NewMockUserRepository()
	// And a valid usecases.
	usecases, err := usecases.NewAccountUseCases(accRepo, userRepo)
	assert.NoError(t, err)
	assert.NotNil(t, usecases)
	// And a valid uuid.UUID accID
	accID := uuid.New()
	// And a mocked response when accRepo.GetByID is called.
	acc := entity.NewAccount(int64(1))
	accRepo.On("GetByID", accID).Return(acc, nil)
	// And a mocked response when accRepo.Update is called.
	accRepo.On("Update", mock.Anything).Return(errors.New("error"))
	// When Update is called with an error saving the account.
	err = usecases.Update(accID.String(), 1.00, true)
	// Then the error is not nil
	assert.NotNil(t, err)
}

// TestUpdateWithSuccess tests the Update method with success.
func TestUpdateWithSuccess(t *testing.T) {
	// Given valid repositories.
	accRepo := accMock.NewMockAccountRepository()
	userRepo := userMock.NewMockUserRepository()
	// And a valid usecases.
	usecases, err := usecases.NewAccountUseCases(accRepo, userRepo)
	assert.NoError(t, err)
	assert.NotNil(t, usecases)
	// And a valid uuid.UUID accID
	accID := uuid.New()
	// And a mocked response when accRepo.GetByID is called.
	acc := entity.NewAccount(int64(1))
	accRepo.On("GetByID", accID).Return(acc, nil)
	// And a mocked response when accRepo.Update is called.
	accRepo.On("Update", mock.Anything).Return(nil)
	// When Update is called with success.
	err = usecases.Update(accID.String(), 1.00, true)
	// Then the error is nil
	assert.Nil(t, err)
}
