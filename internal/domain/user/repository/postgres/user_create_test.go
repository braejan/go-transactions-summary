package postgres_test

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/braejan/go-transactions-summary/internal/domain/user/entity"
	"github.com/braejan/go-transactions-summary/internal/domain/user/repository/postgres"
	voPostgres "github.com/braejan/go-transactions-summary/internal/valueobject/postgres"
	mockvoPostgres "github.com/braejan/go-transactions-summary/internal/valueobject/postgres/mock"
	"github.com/braejan/go-transactions-summary/internal/valueobject/user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// TestCreateErrUserNil tests the error returned when the user is nil.
func TestCreateErrUserNil(t *testing.T) {
	dbBase := mockvoPostgres.NewMockBasePostgresDatabase()
	// And a valid user repository.
	userRepo := postgres.NewPostgresUserRepository(dbBase)
	// When creating a user.
	err := userRepo.Create(nil)
	// Then the error returned is ErrNilUser.
	assert.NotNil(t, err)
	assert.Equal(t, user.ErrNilUser, err)
}

// TestCreateErrOpeningDatabase tests the error returned when opening the database.
func TestCreateErrOpeningDatabase(t *testing.T) {
	dbBase := mockvoPostgres.NewMockBasePostgresDatabase()
	// And a valid user repository.
	userRepo := postgres.NewPostgresUserRepository(dbBase)
	// And a valid user entity.
	user := entity.NewUser(int64(1), "John Doe", "john.doe@amazingemail.com")
	// And a mocked response when calling Open.
	dbBase.On("Open").Return(nil, voPostgres.ErrOpeningDatabase)
	// When creating a user.
	err := userRepo.Create(user)
	// Then the error returned is ErrOpeningDatabase.
	assert.NotNil(t, err)
	assert.Equal(t, voPostgres.ErrOpeningDatabase, err)

}

// TestCreateErrBeginningTransaction tests the error returned when beginning the transaction.
func TestCreateErrBeginningTransaction(t *testing.T) {
	dbBase := mockvoPostgres.NewMockBasePostgresDatabase()
	// And a valid user repository.
	userRepo := postgres.NewPostgresUserRepository(dbBase)
	// And a valid user entity.
	user := entity.NewUser(int64(1), "John Doe", "john.doe@amazingemail.com")
	// And a mocked database.
	db, _, _ := sqlmock.New()
	// And a mocked response when calling Open.
	dbBase.On("Open").Return(db, nil)
	// And a mocked response when calling BeginTx.
	dbBase.On("BeginTx", db).Return(nil, voPostgres.ErrBeginningTransaction)
	// And a mocked response when calling Rollback.
	dbBase.On("Rollback", mock.Anything).Return(nil)
	// And a mocked response when calling Close.
	dbBase.On("Close", db).Return(nil)
	// When creating a user.
	err := userRepo.Create(user)
	// Then the error returned is ErrBeginningTransaction.
	assert.NotNil(t, err)
	assert.Equal(t, voPostgres.ErrBeginningTransaction, err)
}

// TestCreateErrExec tests the error returned when executing the query.
func TestCreateErrExec(t *testing.T) {
	dbBase := mockvoPostgres.NewMockBasePostgresDatabase()
	// And a valid user repository.
	userRepo := postgres.NewPostgresUserRepository(dbBase)
	// And a valid user entity.
	userToTest := entity.NewUser(int64(1), "John Doe", "john.doe@amazingemail.com")
	// And a mocked database.
	db, mockedDB, _ := sqlmock.New()
	// And a mocked response when calling Open.
	dbBase.On("Open").Return(db, nil)
	// And a mocked Tx.
	tx, _ := db.BeginTx(context.Background(), nil)
	// And a mocked response when calling BeginTx.
	dbBase.On("BeginTx", db).Return(tx, nil)
	// And a mocked response when calling Rollback.
	dbBase.On("Rollback", mock.Anything).Return(nil)
	mockedDB.ExpectBegin()
	// And a mocked response when calling Exec.
	dbBase.On("Exec", tx, "INSERT INTO users (id, name, email) VALUES ($1, $2, $3)", []interface{}{int64(1), "John Doe", "john.doe@amazingemail.com"}).Return(nil, voPostgres.ErrExec)
	mockedDB.ExpectRollback()
	// And a mocked response when calling Close.
	dbBase.On("Close", db).Return(nil)
	// When creating a user.
	err := userRepo.Create(userToTest)
	// Then the error returned is ErrExec.
	assert.NotNil(t, err)
	assert.Equal(t, user.ErrCreatingUser, err)
}

// TestCreateSuccess tests the success of creating a user.
func TestCreateSuccess(t *testing.T) {
	dbBase := mockvoPostgres.NewMockBasePostgresDatabase()
	// And a valid user repository.
	userRepo := postgres.NewPostgresUserRepository(dbBase)
	// And a valid user entity.
	userToTest := entity.NewUser(int64(1), "John Doe", "john.doe@amazingemail.com")
	// And a mocked database.
	db, mockedDB, _ := sqlmock.New()
	// And a mocked response when calling Open.
	dbBase.On("Open").Return(db, nil)
	// And a mocked Tx.
	tx, _ := db.BeginTx(context.Background(), nil)
	// And a mocked response when calling BeginTx.
	dbBase.On("BeginTx", db).Return(tx, nil)
	// And a mocked response when calling Rollback.
	dbBase.On("Rollback", mock.Anything).Return(nil)
	mockedDB.ExpectBegin()
	// And a mocked response when calling Exec.
	dbBase.On("Exec", tx, "INSERT INTO users (id, name, email) VALUES ($1, $2, $3)", []interface{}{int64(1), "John Doe", "john.doe@amazingemail.com"}).Return(nil, nil)
	// And a mocked response when calling Commit.
	dbBase.On("Commit", tx).Return(nil)
	mockedDB.ExpectCommit()
	// And a mocked response when calling Close.
	dbBase.On("Close", db).Return(nil)
	// When creating a user.
	err := userRepo.Create(userToTest)
	// Then the error returned is nil.
	assert.Nil(t, err)
}
