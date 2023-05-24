package usecases_test

import (
	"errors"
	"testing"

	"github.com/braejan/go-transactions-summary/internal/domain/user/entity"
	"github.com/braejan/go-transactions-summary/internal/domain/user/repository/mock"
	"github.com/braejan/go-transactions-summary/internal/domain/user/usecases"
	"github.com/braejan/go-transactions-summary/internal/valueobject"
	"github.com/stretchr/testify/assert"
)

// TestNewUserUsecasesWithouUserRepository tests the NewUserUsecases function
// with a nil user repository.
func TestNewUserUsecasesWithouUserRepository(t *testing.T) {
	// Given a nil user repository
	// When call NewUserUsecases
	_, err := usecases.NewUserUsecases(nil)
	// Then get next errors
	assert.NotNil(t, err)
	assert.Equal(t, valueobject.ErrUserRepositoryIsNil, err)
}

// TestNewUserUsecasesWithUserRepository tests the NewUserUsecases function
// with a valid user repository.
func TestNewUserUsecasesWithUserRepository(t *testing.T) {
	// Given a valid user repository
	mockedUserRepo := mock.NewMockUserRepository()
	// When call NewUserUsecases
	userUsecases, err := usecases.NewUserUsecases(mockedUserRepo)
	// Then get no errors
	assert.Nil(t, err)
	assert.NotNil(t, userUsecases)
}

// TestGetByIDWithError tests the GetByID method with an error.
func TestGetByIDWithError(t *testing.T) {
	// Given a valid user repository
	mockedUserRepo := mock.NewMockUserRepository()
	// And a user usecases
	userUsecases, _ := usecases.NewUserUsecases(mockedUserRepo)
	// And a user ID
	ID := int64(1)
	// And a mocked error calling GetByID
	mockedUserRepo.On("GetByID", ID).Return(nil, valueobject.ErrUserNotFound)
	// When call GetByID with an invalid ID
	user, err := userUsecases.GetByID(ID)
	// Then get next errors
	assert.NotNil(t, err)
	assert.Equal(t, valueobject.ErrUserNotFound, err)
	assert.Nil(t, user)
}

// TestGetByIDSucess tests the GetByID method with success.
func TestGetByIDSucess(t *testing.T) {
	// Given a valid user repository
	mockedUserRepo := mock.NewMockUserRepository()
	// And a user usecases
	userUsecases, _ := usecases.NewUserUsecases(mockedUserRepo)
	// And a entity.User
	user := &entity.User{
		ID:    int64(1),
		Name:  "John Doe",
		Email: "john.doe@amazinemail.com",
	}
	// And a user ID
	ID := int64(1)
	// And a mocked entity.User response calling GetByID
	mockedUserRepo.On("GetByID", ID).Return(user, nil)
	// When call GetByID with a valid ID
	user, err := userUsecases.GetByID(ID)
	// Then get no errors
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, int64(1), user.ID)
	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, "john.doe@amazinemail.com", user.Email)
}

// TestGetByEmailWithError tests the GetByEmail method with an error.
func TestGetByEmailWithError(t *testing.T) {
	// Given a valid user repository
	mockedUserRepo := mock.NewMockUserRepository()
	// And a user usecases
	userUsecases, _ := usecases.NewUserUsecases(mockedUserRepo)
	// And a user email
	email := "john.doe@amazinemail.com"
	// And a mocked error calling GetByEmail
	mockedUserRepo.On("GetByEmail", email).Return(nil, valueobject.ErrUserNotFound)
	// When call GetByEmail with an invalid email
	user, err := userUsecases.GetByEmail(email)
	// Then get next errors
	assert.NotNil(t, err)
	assert.Equal(t, valueobject.ErrUserNotFound, err)
	assert.Nil(t, user)
}

// TestGetByEmailSucess tests the GetByEmail method with success.
func TestGetByEmailSucess(t *testing.T) {
	// Given a valid user repository
	mockedUserRepo := mock.NewMockUserRepository()
	// And a user usecases
	userUsecases, _ := usecases.NewUserUsecases(mockedUserRepo)
	// And a entity.User
	user := &entity.User{
		ID:    int64(1),
		Name:  "John Doe",
		Email: "john.doe@amazinemail.com",
	}
	// And a user email
	email := "john.doe@amazinemail.com"
	// And a mocked entity.User response calling GetByEmail
	mockedUserRepo.On("GetByEmail", email).Return(user, nil)
	// When call GetByEmail with a valid email
	user, err := userUsecases.GetByEmail(email)
	// Then get no errors
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, int64(1), user.ID)
	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, "john.doe@amazinemail.com", user.Email)
}

// TestCreateWithErrorGetByID tests the Create method with an error calling GetByID.
func TestCreateWithErrorGetByID(t *testing.T) {
	// Given a valid user repository
	mockedUserRepo := mock.NewMockUserRepository()
	// And a user usecases
	userUsecases, _ := usecases.NewUserUsecases(mockedUserRepo)
	// And a entity.User
	ID := int64(1)
	user := &entity.User{
		ID:    int64(1),
		Name:  "John Doe",
		Email: "john.doe@amazinemail.com",
	}
	// And a mocked entity.User response calling GetByID
	mockedUserRepo.On("GetByID", ID).Return(user, nil)
	// When call Create with an invalid user
	err := userUsecases.Create(ID, "John Doe", "john.doe@amazinemail.com")
	// Then get next errors
	assert.NotNil(t, err)
	assert.Equal(t, valueobject.ErrUserAlreadyCreated, err)
}

// TestCreateWithError tests the Create method with an error creating the user.
func TestCreateWithError(t *testing.T) {
	// Given a valid user repository
	mockedUserRepo := mock.NewMockUserRepository()
	// And a user usecases
	userUsecases, _ := usecases.NewUserUsecases(mockedUserRepo)
	// And a entity.User
	ID := int64(1)
	user := &entity.User{
		ID:    int64(1),
		Name:  "John Doe",
		Email: "john.doe@amazinemail.com",
	}
	// And a mocked entity.User response calling GetByID
	mockedUserRepo.On("GetByID", ID).Return(nil, valueobject.ErrUserNotFound)
	// And a mocked error calling Create
	mockedUserRepo.On("Create", user).Return(errors.New("new mocked error"))
	// When call Create with an invalid user
	err := userUsecases.Create(ID, "John Doe", "john.doe@amazinemail.com")
	// Then get next errors
	assert.NotNil(t, err)
	assert.Equal(t, errors.New("new mocked error"), err)
}

// TestCreateSucess tests the Create method with success.
func TestCreateSucess(t *testing.T) {
	// Given a valid user repository
	mockedUserRepo := mock.NewMockUserRepository()
	// And a user usecases
	userUsecases, _ := usecases.NewUserUsecases(mockedUserRepo)
	// And a entity.User
	ID := int64(1)
	user := &entity.User{
		ID:    int64(1),
		Name:  "John Doe",
		Email: "john.doe@amazinemail.com",
	}
	// And a mocked entity.User response calling GetByID
	mockedUserRepo.On("GetByID", ID).Return(nil, valueobject.ErrUserNotFound)
	// And a mocked entity.User response calling Create
	mockedUserRepo.On("Create", user).Return(nil)
	// When call Create with a valid user
	err := userUsecases.Create(ID, "John Doe", "john.doe@amazinemail.com")
	// Then get no errors
	assert.Nil(t, err)
}

// TestUpdateWithErrorGetByID tests the Update method with an error calling GetByID.
func TestUpdateWithErrorGetByID(t *testing.T) {
	// Given a valid user repository
	mockedUserRepo := mock.NewMockUserRepository()
	// And a user usecases
	userUsecases, _ := usecases.NewUserUsecases(mockedUserRepo)
	ID := int64(1)
	// And a mocked error response calling GetByID
	mockedUserRepo.On("GetByID", ID).Return(nil, valueobject.ErrUserNotFound)
	// When call Update with an invalid user
	err := userUsecases.Update(ID, "John Doe", "john.doe@amazinemail.com")
	// Then get next errors
	assert.NotNil(t, err)
	assert.Equal(t, valueobject.ErrUserNotFound, err)
}

// TestUpdateWithError tests the Update method with an error updating the user.
func TestUpdateWithError(t *testing.T) {
	// Given a valid user repository
	mockedUserRepo := mock.NewMockUserRepository()
	// And a user usecases
	userUsecases, _ := usecases.NewUserUsecases(mockedUserRepo)
	ID := int64(1)
	// And a mocked entity.User response calling GetByID
	user := &entity.User{
		ID:    int64(1),
		Name:  "John Doe",
		Email: "john.doe@amazinemail.com",
	}
	mockedUserRepo.On("GetByID", ID).Return(user, nil)
	// And a mocked error response calling Update
	mockedUserRepo.On("Update", user).Return(errors.New("new mocked error"))
	// When call Update with an invalid user
	err := userUsecases.Update(ID, "John Doe", "john.doe@amazinemail.com")
	// Then get next errors
	assert.NotNil(t, err)
	assert.Equal(t, errors.New("new mocked error"), err)
}

// TestUpdateSucess tests the Update method with success.
func TestUpdateSucess(t *testing.T) {
	// Given a valid user repository
	mockedUserRepo := mock.NewMockUserRepository()
	// And a user usecases
	userUsecases, _ := usecases.NewUserUsecases(mockedUserRepo)
	ID := int64(1)
	// And a mocked entity.User response calling GetByID
	user := &entity.User{
		ID:    int64(1),
		Name:  "John Doe",
		Email: "john.doe@amazinemail.com",
	}
	mockedUserRepo.On("GetByID", ID).Return(user, nil)
	// And a mocked entity.User response calling Update
	mockedUserRepo.On("Update", user).Return(nil)
	// When call Update with a valid user
	err := userUsecases.Update(ID, "John Doe", "john.doe@amazinemail.com")
	// Then get no errors
	assert.Nil(t, err)
}
