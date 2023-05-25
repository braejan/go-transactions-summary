package postgres_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/braejan/go-transactions-summary/internal/domain/account/entity"
	"github.com/braejan/go-transactions-summary/internal/domain/account/repository/postgres"
	"github.com/braejan/go-transactions-summary/internal/valueobject/account"
	voPostgres "github.com/braejan/go-transactions-summary/internal/valueobject/postgres"
	"github.com/braejan/go-transactions-summary/internal/valueobject/postgres/mock"
	"github.com/stretchr/testify/assert"
)

// TestUpdateErrUserNil tests the error returned when the user is nil.
func TestUpdateErrUserNil(t *testing.T) {
	// Given a valid configuration.
	configuration := voPostgres.GetPostgresConfigurationFromEnv()
	dbBase := mock.NewMockBasePostgresDatabase()
	// And a valid account repository.
	accountRepo, _ := postgres.NewPostgresAccountRepository(configuration, dbBase)
	// When updating a account.
	err := accountRepo.Update(nil)
	// Then the error returned is ErrNilAccount.
	assert.NotNil(t, err)
	assert.Equal(t, account.ErrNilAccount, err)
}

// TestUpdateErrOpeningDatabase tests the error returned when opening the database.
func TestUpdateErrOpeningDatabase(t *testing.T) {
	// Given a valid configuration.
	configuration := voPostgres.GetPostgresConfigurationFromEnv()
	dbBase := mock.NewMockBasePostgresDatabase()
	// And a valid account repository.
	accountRepo, _ := postgres.NewPostgresAccountRepository(configuration, dbBase)
	// And a valid account entity.
	acc := entity.NewAccount(int64(1))
	// And a mocked response when calling Open.
	dbBase.On("Open", configuration.GetDataSourceName()).Return(nil, voPostgres.ErrOpeningDatabase)
	// When updating a account.
	err := accountRepo.Update(acc)
	// Then the error returned is ErrOpeningDatabase.
	assert.NotNil(t, err)
	assert.Equal(t, voPostgres.ErrOpeningDatabase, err)
}

// TestUpdateErrBeginningTransaction tests the error returned when beginning the transaction.
func TestUpdateErrBeginningTransaction(t *testing.T) {
	// Given a valid configuration.
	configuration := voPostgres.GetPostgresConfigurationFromEnv()
	dbBase := mock.NewMockBasePostgresDatabase()
	// And a valid account repository.
	accountRepo, _ := postgres.NewPostgresAccountRepository(configuration, dbBase)
	// And a valid account entity.
	acc := entity.NewAccount(int64(1))
	// And a mocked database.
	db, _, _ := sqlmock.New()
	// And a mocked response when calling Open.
	dbBase.On("Open", configuration.GetDataSourceName()).Return(db, nil)
	// And a mocked response when calling Close.
	dbBase.On("Close", db).Return(nil)
	// And a mocked response when calling Begin.
	dbBase.On("BeginTx", db).Return(nil, voPostgres.ErrBeginningTransaction)
	// When updating a account.
	err := accountRepo.Update(acc)
	// Then the error returned is ErrBeginningTransaction.
	assert.NotNil(t, err)
	assert.Equal(t, voPostgres.ErrBeginningTransaction, err)
}

// TestUpdateErrUpdatingUser tests the error returned when updating the user.
func TestUpdateErrUpdatingUser(t *testing.T) {
	// Given a valid configuration.
	configuration := voPostgres.GetPostgresConfigurationFromEnv()
	dbBase := mock.NewMockBasePostgresDatabase()
	// And a valid account repository.
	accountRepo, _ := postgres.NewPostgresAccountRepository(configuration, dbBase)
	// And a valid account entity.
	acc := entity.NewAccount(int64(1))
	// And a mocked database.
	db, _, _ := sqlmock.New()
	// And a mocked response when calling Open.
	dbBase.On("Open", configuration.GetDataSourceName()).Return(db, nil)
	// And a mocked response when calling Close.
	dbBase.On("Close", db).Return(nil)
	// And a mocked response when calling Begin.
	tx, _ := db.Begin()
	dbBase.On("BeginTx", db).Return(tx, nil)
	// And a mocked response when calling Rollback.
	dbBase.On("Rollback", tx).Return(nil)
	// And a mocked response when calling Exec.
	dbBase.On("Exec", tx, "UPDATE accounts SET balance = $1, active = $2 WHERE id = $3", []interface{}{acc.Balance, acc.Active, acc.ID}).Return(nil, account.ErrUpdatingAccount)
	// When updating a account.
	err := accountRepo.Update(acc)
	// Then the error returned is ErrUpdatingAccount.
	assert.NotNil(t, err)
	assert.Equal(t, account.ErrUpdatingAccount, err)
}

// TestUpdateErrCommittingTransaction tests the error returned when committing the transaction.
func TestUpdateErrCommittingTransaction(t *testing.T) {
	// Given a valid configuration.
	configuration := voPostgres.GetPostgresConfigurationFromEnv()
	dbBase := mock.NewMockBasePostgresDatabase()
	// And a valid account repository.
	accountRepo, _ := postgres.NewPostgresAccountRepository(configuration, dbBase)
	// And a valid account entity.
	acc := entity.NewAccount(int64(1))
	// And a mocked database.
	db, _, _ := sqlmock.New()
	// And a mocked response when calling Open.
	dbBase.On("Open", configuration.GetDataSourceName()).Return(db, nil)
	// And a mocked response when calling Close.
	dbBase.On("Close", db).Return(nil)
	// And a mocked response when calling Begin.
	tx, _ := db.Begin()
	dbBase.On("BeginTx", db).Return(tx, nil)
	// And a mocked response when calling Exec.
	dbBase.On("Exec", tx, "UPDATE accounts SET balance = $1, active = $2 WHERE id = $3", []interface{}{acc.Balance, acc.Active, acc.ID}).Return(nil, nil)
	// And a mocked response when calling Commit.
	dbBase.On("Commit", tx).Return(voPostgres.ErrCommittingTransaction)
	// When updating a account.
	err := accountRepo.Update(acc)
	// Then the error returned is ErrCommittingTransaction.
	assert.NotNil(t, err)
	assert.Equal(t, voPostgres.ErrCommittingTransaction, err)
}

// TestUpdateSuccess tests the success when updating the account.
func TestUpdateSuccess(t *testing.T) {
	// Given a valid configuration.
	configuration := voPostgres.GetPostgresConfigurationFromEnv()
	dbBase := mock.NewMockBasePostgresDatabase()
	// And a valid account repository.
	accountRepo, _ := postgres.NewPostgresAccountRepository(configuration, dbBase)
	// And a valid account entity.
	acc := entity.NewAccount(int64(1))
	// And a mocked database.
	db, _, _ := sqlmock.New()
	// And a mocked response when calling Open.
	dbBase.On("Open", configuration.GetDataSourceName()).Return(db, nil)
	// And a mocked response when calling Close.
	dbBase.On("Close", db).Return(nil)
	// And a mocked response when calling Begin.
	tx, _ := db.Begin()
	dbBase.On("BeginTx", db).Return(tx, nil)
	// And a mocked response when calling Exec.
	dbBase.On("Exec", tx, "UPDATE accounts SET balance = $1, active = $2 WHERE id = $3", []interface{}{acc.Balance, acc.Active, acc.ID}).Return(nil, nil)
	// And a mocked response when calling Commit.
	dbBase.On("Commit", tx).Return(nil)
	// When updating a account.
	err := accountRepo.Update(acc)
	// Then the error returned is nil.
	assert.Nil(t, err)
}
