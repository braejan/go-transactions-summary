package postgres_test

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/braejan/go-transactions-summary/internal/domain/user/entity"
	"github.com/braejan/go-transactions-summary/internal/domain/user/repository/postgres"
	voPostgres "github.com/braejan/go-transactions-summary/internal/valueobject/postgres"
	"github.com/braejan/go-transactions-summary/internal/valueobject/postgres/mock"
	"github.com/braejan/go-transactions-summary/internal/valueobject/user"
	"github.com/stretchr/testify/assert"
)

// TestUpdateErrUserNil tests the error returned when the user is nil.
func TestUpdateErrUserNil(t *testing.T) {
	// Given a valid configuration.
	configuration := voPostgres.GetPostgresConfigurationFromEnv()
	dbBase := mock.NewMockBasePostgresDatabase()
	// And a valid user repository.
	userRepo, _ := postgres.NewPostgresUserRepository(configuration, dbBase)
	// When updating a user.
	err := userRepo.Update(nil)
	// Then the error returned is ErrNilUser.
	assert.NotNil(t, err)
	assert.Equal(t, user.ErrNilUser, err)
}

// TestUpdateErrOpeningDatabase tests the error returned when opening the database.
func TestUpdateErrOpeningDatabase(t *testing.T) {
	// Given a valid configuration.
	configuration := voPostgres.GetPostgresConfigurationFromEnv()
	dbBase := mock.NewMockBasePostgresDatabase()
	// And a valid user repository.
	userRepo, _ := postgres.NewPostgresUserRepository(configuration, dbBase)
	// And a valid user entity.
	user := entity.NewUser(int64(1), "John Doe", "john.doe@amazingemail.com")
	// And a mocked response when calling Open.
	dbBase.On("Open", configuration.GetDataSourceName()).Return(nil, voPostgres.ErrOpeningDatabase)
	// When updating a user.
	err := userRepo.Update(user)
	// Then the error returned is ErrOpeningDatabase.
	assert.NotNil(t, err)
	assert.Equal(t, voPostgres.ErrOpeningDatabase, err)
}

// TestUpdateErrBeginningTransaction tests the error returned when beginning the transaction.
func TestUpdateErrBeginningTransaction(t *testing.T) {
	// Given a valid configuration.
	configuration := voPostgres.GetPostgresConfigurationFromEnv()
	dbBase := mock.NewMockBasePostgresDatabase()
	// And a valid user repository.
	userRepo, _ := postgres.NewPostgresUserRepository(configuration, dbBase)
	// And a valid user entity.
	userToTest := entity.NewUser(int64(1), "John Doe", "john.doe@amazingemail.com")
	// And a mocked database.
	db, _, _ := sqlmock.New()
	defer db.Close()
	// And a mocked response when calling Open.
	dbBase.On("Open", configuration.GetDataSourceName()).Return(db, nil)
	// And a mocked response when calling Close.
	dbBase.On("Close", db).Return(nil)
	// And a mocked response when calling BeginTx.
	dbBase.On("BeginTx", db).Return(nil, voPostgres.ErrBeginningTransaction)
	// When updating a user.
	err := userRepo.Update(userToTest)
	// Then the error returned is ErrBeginningTransaction.
	assert.NotNil(t, err)
	assert.Equal(t, voPostgres.ErrBeginningTransaction, err)
}

// TestUpdateErrUpdatingUser tests the error returned when updating the user.
func TestUpdateErrUpdatingUser(t *testing.T) {
	// Given a valid configuration.
	configuration := voPostgres.GetPostgresConfigurationFromEnv()
	dbBase := mock.NewMockBasePostgresDatabase()
	// And a valid user repository.
	userRepo, _ := postgres.NewPostgresUserRepository(configuration, dbBase)
	// And a valid user entity.
	userToTest := entity.NewUser(int64(1), "John Doe", "john.doe@amazingemail.com")
	// And a mocked database.
	db, dbMocked, _ := sqlmock.New()
	defer db.Close()
	// And a mocked response when calling Open.
	dbBase.On("Open", configuration.GetDataSourceName()).Return(db, nil)
	// And a mocked response when calling Close.
	dbBase.On("Close", db).Return(nil)
	// And a mocked tx.
	tx, _ := db.BeginTx(context.Background(), nil)
	// And a mocked response when calling BeginTx.
	dbBase.On("BeginTx", db).Return(tx, nil)
	dbMocked.ExpectBegin()
	// And a mocked response when calling Exec.
	dbBase.On("Exec", tx, "UPDATE users SET name = $1, email = $2 WHERE id = $3", []interface{}{"John Doe", "john.doe@amazingemail.com", int64(1)}).Return(nil, user.ErrUpdatingUser)
	// And a mocked response when calling Rollback.
	dbBase.On("Rollback", tx).Return(nil)
	dbMocked.ExpectRollback()
	// When updating a user.
	err := userRepo.Update(userToTest)
	// Then the error returned is ErrUpdatingUser.
	assert.NotNil(t, err)
	assert.Equal(t, user.ErrUpdatingUser, err)
}

// TestUpdateSuccess tests the success of updating a user.
func TestUpdateSuccess(t *testing.T) {
	// Given a valid configuration.
	configuration := voPostgres.GetPostgresConfigurationFromEnv()
	dbBase := mock.NewMockBasePostgresDatabase()
	// And a valid user repository.
	userRepo, _ := postgres.NewPostgresUserRepository(configuration, dbBase)
	// And a valid user entity.
	userToTest := entity.NewUser(int64(1), "John Doe", "john.doe@amazingemail.com")
	// And a mocked database.
	db, dbMocked, _ := sqlmock.New()
	defer db.Close()
	// And a mocked response when calling Open.
	dbBase.On("Open", configuration.GetDataSourceName()).Return(db, nil)
	// And a mocked response when calling Close.
	dbBase.On("Close", db).Return(nil)
	// And a mocked tx.
	tx, _ := db.BeginTx(context.Background(), nil)
	// And a mocked response when calling BeginTx.
	dbBase.On("BeginTx", db).Return(tx, nil)
	dbMocked.ExpectBegin()
	// And a mocked response when calling Update.
	dbBase.On("Exec", tx, "UPDATE users SET name = $1, email = $2 WHERE id = $3", []interface{}{"John Doe", "john.doe@amazingemail.com", int64(1)}).Return(nil, nil)
	// And a mocked response when calling Commit.
	dbBase.On("Commit", tx).Return(nil)
	dbMocked.ExpectCommit()
	// When updating a user.
	err := userRepo.Update(userToTest)
	// Then the error returned is nil.
	assert.Nil(t, err)
}
