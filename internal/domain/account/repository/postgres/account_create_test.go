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

// TestCreateErrAccountNil tests the error returned when the account is nil.
func TestCreateErrAccountNil(t *testing.T) {
	dbBase := mock.NewMockBasePostgresDatabase()
	// And a valid account repository.
	accountRepo := postgres.NewPostgresAccountRepository(dbBase)
	// When creating a account.
	err := accountRepo.Create(nil)
	// Then the error returned is ErrNilAccount.
	assert.NotNil(t, err)
	assert.Equal(t, account.ErrNilAccount, err)

}

// TestCreateErrOpeningDatabase tests the error returned when opening the database.
func TestCreateErrOpeningDatabase(t *testing.T) {
	dbBase := mock.NewMockBasePostgresDatabase()
	// And a valid account repository.
	accountRepo := postgres.NewPostgresAccountRepository(dbBase)
	// And a valid account entity.
	account := entity.NewAccount(int64(1))
	// And a mocked response when calling Open.
	dbBase.On("Open").Return(nil, voPostgres.ErrOpeningDatabase)
	// When creating a account.
	err := accountRepo.Create(account)
	// Then the error returned is ErrOpeningDatabase.
	assert.NotNil(t, err)
	assert.Equal(t, voPostgres.ErrOpeningDatabase, err)
}

// TestCreateErrBeginningTransaction tests the error returned when beginning the transaction.
func TestCreateErrBeginningTransaction(t *testing.T) {
	dbBase := mock.NewMockBasePostgresDatabase()
	// And a valid account repository.
	accountRepo := postgres.NewPostgresAccountRepository(dbBase)
	// And a valid account entity.
	account := entity.NewAccount(int64(1))
	// And a mocked database.
	db, _, _ := sqlmock.New()
	// And a mocked response when calling Open.
	dbBase.On("Open").Return(db, nil)
	// And a mocked response when calling Close.
	dbBase.On("Close", db).Return(nil)
	// And a mocked response when calling Begin.
	dbBase.On("BeginTx", db).Return(nil, voPostgres.ErrBeginningTransaction)
	// When creating a account.
	err := accountRepo.Create(account)
	// Then the error returned is ErrBeginningTransaction.
	assert.NotNil(t, err)
	assert.Equal(t, voPostgres.ErrBeginningTransaction, err)
}

// TestCreateErrExec tests the error returned when executing the query.
func TestCreateErrExec(t *testing.T) {
	dbBase := mock.NewMockBasePostgresDatabase()
	// And a valid account repository.
	accountRepo := postgres.NewPostgresAccountRepository(dbBase)
	// And a valid account entity.
	acc := entity.NewAccount(int64(1))
	// And a mocked database.
	db, _, _ := sqlmock.New()
	// And a mocked response when calling Open.
	dbBase.On("Open").Return(db, nil)
	// And a mocked response when calling Close.
	dbBase.On("Close", db).Return(nil)
	// And a mocked response when calling Begin.
	tx, _ := db.Begin()
	dbBase.On("BeginTx", db).Return(tx, nil)
	// And a mocked response when calling Exec.
	dbBase.On("Exec", tx, "INSERT INTO accounts (id, balance, userid, active) VALUES ($1, $2, $3, $4)", []interface{}{acc.ID, acc.Balance, acc.UserID, acc.Active}).Return(nil, voPostgres.ErrExec)
	// And a mocked response when calling Rollback.
	dbBase.On("Rollback", tx).Return(nil)
	// When creating a account.
	err := accountRepo.Create(acc)
	// Then the error returned is ErrExec.
	assert.NotNil(t, err)
	assert.Equal(t, account.ErrCreatingAccount, err)
}

// TestCreateErrCommittingTransaction tests the error returned when committing the transaction.
func TestCreateErrCommittingTransaction(t *testing.T) {
	dbBase := mock.NewMockBasePostgresDatabase()
	// And a valid account repository.
	accountRepo := postgres.NewPostgresAccountRepository(dbBase)
	// And a valid account entity.
	acc := entity.NewAccount(int64(1))
	// And a mocked database.
	db, _, _ := sqlmock.New()
	// And a mocked response when calling Open.
	dbBase.On("Open").Return(db, nil)
	// And a mocked response when calling Close.
	dbBase.On("Close", db).Return(nil)
	// And a mocked response when calling Begin.
	tx, _ := db.Begin()
	dbBase.On("BeginTx", db).Return(tx, nil)
	// And a mocked response when calling Exec.
	dbBase.On("Exec", tx, "INSERT INTO accounts (id, balance, userid, active) VALUES ($1, $2, $3, $4)", []interface{}{acc.ID, acc.Balance, acc.UserID, acc.Active}).Return(nil, nil)
	// And a mocked response when calling Commit.
	dbBase.On("Commit", tx).Return(voPostgres.ErrCommittingTransaction)
	// When creating a account.
	err := accountRepo.Create(acc)
	// Then the error returned is ErrCommittingTransaction.
	assert.NotNil(t, err)
	assert.Equal(t, voPostgres.ErrCommittingTransaction, err)
}

// TestCreateSuccess tests the success when creating a account.
func TestCreateSuccess(t *testing.T) {
	dbBase := mock.NewMockBasePostgresDatabase()
	// And a valid account repository.
	accountRepo := postgres.NewPostgresAccountRepository(dbBase)
	// And a valid account entity.
	acc := entity.NewAccount(int64(1))
	// And a mocked database.
	db, _, _ := sqlmock.New()
	// And a mocked response when calling Open.
	dbBase.On("Open").Return(db, nil)
	// And a mocked response when calling Close.
	dbBase.On("Close", db).Return(nil)
	// And a mocked response when calling Begin.
	tx, _ := db.Begin()
	dbBase.On("BeginTx", db).Return(tx, nil)
	// And a mocked response when calling Exec.
	dbBase.On("Exec", tx, "INSERT INTO accounts (id, balance, userid, active) VALUES ($1, $2, $3, $4)", []interface{}{acc.ID, acc.Balance, acc.UserID, acc.Active}).Return(nil, nil)
	// And a mocked response when calling Commit.
	dbBase.On("Commit", tx).Return(nil)
	// When creating a account.
	err := accountRepo.Create(acc)
	// Then the error returned is nil.
	assert.Nil(t, err)
}
