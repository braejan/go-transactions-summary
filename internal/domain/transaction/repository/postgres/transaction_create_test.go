package postgres_test

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/braejan/go-transactions-summary/internal/domain/transaction/entity"
	"github.com/braejan/go-transactions-summary/internal/domain/transaction/repository/postgres"
	voPostgres "github.com/braejan/go-transactions-summary/internal/valueobject/postgres"
	mockvoPostgres "github.com/braejan/go-transactions-summary/internal/valueobject/postgres/mock"
	"github.com/braejan/go-transactions-summary/internal/valueobject/transaction"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// TestCreateWithInvalidTransaction tests the error returned when the transaction is invalid.
func TestCreateWithInvalidTransaction(t *testing.T) {
	dbBase := mockvoPostgres.NewMockBasePostgresDatabase()
	// And a valid account repository.
	transactionRepo := postgres.NewPostgresTransactionRepository(dbBase)
	// When creating a account .
	err := transactionRepo.Create(nil)
	// Then the error returned is ErrBeginningTransaction.
	assert.NotNil(t, err)
	assert.Equal(t, transaction.ErrNilTransaction, err)
}

// TestCreateErrOpeningDatabase tests the error returned when the database cannot be opened.
func TestCreateErrOpeningDatabase(t *testing.T) {
	dbBaseMocked := mockvoPostgres.NewMockBasePostgresDatabase()
	// And a valid account repository.
	transactionRepo := postgres.NewPostgresTransactionRepository(dbBaseMocked)
	// And a mocked response calling Open.
	dbBaseMocked.On("Open").Return(nil, voPostgres.ErrOpeningDatabase)
	// When creating a account .
	err := transactionRepo.Create(&entity.Transaction{})
	// Then the error returned is ErrOpeningDatabase.
	assert.NotNil(t, err)
	assert.Equal(t, voPostgres.ErrOpeningDatabase, err)
}

// TestCreateErrBeginningTransaction tests the error returned when the transaction cannot be started.
func TestCreateErrBeginningTransaction(t *testing.T) {
	dbBaseMocked := mockvoPostgres.NewMockBasePostgresDatabase()
	// And a valid account repository.
	transactionRepo := postgres.NewPostgresTransactionRepository(dbBaseMocked)
	// And a mocked database.
	db, _, _ := sqlmock.New()
	// And a mocked response calling Open.
	dbBaseMocked.On("Open").Return(db, nil)
	// And a mocked response calling Close.
	dbBaseMocked.On("Close", db).Return(nil)
	// And a mocked response calling Begin.
	dbBaseMocked.On("BeginTx", db).Return(nil, voPostgres.ErrBeginningTransaction)
	// And a mocked response calling Rollback.
	dbBaseMocked.On("Rollback", mock.Anything).Return(nil)
	// When creating a account .
	err := transactionRepo.Create(&entity.Transaction{})
	// Then the error returned is ErrBeginningTransaction.
	assert.NotNil(t, err)
	assert.Equal(t, voPostgres.ErrBeginningTransaction, err)
}

// TestCreateErrExecutingQuery tests the error returned when the query cannot be executed.
func TestCreateErrExecutingQuery(t *testing.T) {
	dbBaseMocked := mockvoPostgres.NewMockBasePostgresDatabase()
	// And a valid account repository.
	transactionRepo := postgres.NewPostgresTransactionRepository(dbBaseMocked)
	// And a valid origin.
	origin := "txns.csv"
	// And a mocked database.
	db, _, _ := sqlmock.New()

	// And a mocked response calling Open.
	dbBaseMocked.On("Open").Return(db, nil)
	// And a mocked response calling Close.
	dbBaseMocked.On("Close", db).Return(nil)
	// And a mocked response calling BeginTx.
	dbTx, _ := db.Begin()
	dbBaseMocked.On("BeginTx", db).Return(dbTx, nil)
	// And a mocked response calling Rollback.
	dbBaseMocked.On("Rollback", dbTx).Return(nil)
	// And a valid entity.Transaction to create.
	tx, err := entity.NewTransaction(uuid.New(), 100.48, time.Now(), origin)
	assert.Nil(t, err)
	// And a mocked response calling Exec.
	dbBaseMocked.On(
		"Exec",
		dbTx,
		"INSERT INTO transactions (id, accountid, amount, date, origin, operation) VALUES ($1, $2, $3, $4, $5, $6)",
		[]interface{}{
			tx.ID,
			tx.AccountID,
			tx.Amount,
			tx.Date,
			tx.Origin,
			tx.Operation,
		}).Return(nil, voPostgres.ErrExec)
	// When creating a account .
	err = transactionRepo.Create(tx)
	// Then the error returned is ErrExecutingQuery.
	assert.NotNil(t, err)
	assert.Equal(t, transaction.ErrCreatingTransaction, err)
}

// TestCreateErrCommittingTransaction tests the error returned when the transaction cannot be committed.
func TestCreateErrCommittingTransaction(t *testing.T) {
	dbBaseMocked := mockvoPostgres.NewMockBasePostgresDatabase()
	// And a valid account repository.
	transactionRepo := postgres.NewPostgresTransactionRepository(dbBaseMocked)
	// And a valid origin.
	origin := "txns.csv"
	// And a mocked database.
	db, _, _ := sqlmock.New()

	// And a mocked response calling Open.
	dbBaseMocked.On("Open").Return(db, nil)
	// And a mocked response calling Close.
	dbBaseMocked.On("Close", db).Return(nil)
	// And a mocked response calling BeginTx.
	dbTx, _ := db.Begin()
	dbBaseMocked.On("BeginTx", db).Return(dbTx, nil)
	// And a mocked response calling Rollback.
	dbBaseMocked.On("Rollback", dbTx).Return(nil)
	// And a valid entity.Transaction to create.
	tx, err := entity.NewTransaction(uuid.New(), 100.48, time.Now(), origin)
	assert.Nil(t, err)
	// And a mocked response calling Exec.
	dbBaseMocked.On(
		"Exec",
		dbTx,
		"INSERT INTO transactions (id, accountid, amount, date, origin, operation) VALUES ($1, $2, $3, $4, $5, $6)",
		[]interface{}{
			tx.ID,
			tx.AccountID,
			tx.Amount,
			tx.Date,
			tx.Origin,
			tx.Operation,
		}).Return(nil, nil)
	// And a mocked response calling Commit.
	dbBaseMocked.On("Commit", dbTx).Return(voPostgres.ErrCommittingTransaction)
	// When creating a account .
	err = transactionRepo.Create(tx)
	// Then the error returned is ErrCommittingTransaction.
	assert.NotNil(t, err)
	assert.Equal(t, voPostgres.ErrCommittingTransaction, err)
}

// TestCreateSuccess tests the success when creating a transaction.
func TestCreateSuccess(t *testing.T) {
	dbBaseMocked := mockvoPostgres.NewMockBasePostgresDatabase()
	// And a valid account repository.
	transactionRepo := postgres.NewPostgresTransactionRepository(dbBaseMocked)
	// And a valid origin.
	origin := "txns.csv"
	// And a mocked database.
	db, _, _ := sqlmock.New()

	// And a mocked response calling Open.
	dbBaseMocked.On("Open").Return(db, nil)
	// And a mocked response calling Close.
	dbBaseMocked.On("Close", db).Return(nil)
	// And a mocked response calling BeginTx.
	dbTx, _ := db.Begin()
	dbBaseMocked.On("BeginTx", db).Return(dbTx, nil)
	// And a mocked response calling Rollback.
	dbBaseMocked.On("Rollback", dbTx).Return(nil)
	// And a mocked response calling Commit.
	dbBaseMocked.On("Commit", dbTx).Return(nil)
	// And a valid entity.Transaction to create.
	tx, err := entity.NewTransaction(uuid.New(), 100.48, time.Now(), origin)
	assert.Nil(t, err)
	// And a mocked response calling Exec.
	dbBaseMocked.On(
		"Exec",
		dbTx,
		"INSERT INTO transactions (id, accountid, amount, date, origin, operation) VALUES ($1, $2, $3, $4, $5, $6)",
		[]interface{}{
			tx.ID,
			tx.AccountID,
			tx.Amount,
			tx.Date,
			tx.Origin,
			tx.Operation,
		}).Return(nil, nil)
	// And a mocked response calling Commit.
	dbBaseMocked.On("Commit", dbTx).Return(nil)
	// When creating a account .
	err = transactionRepo.Create(tx)
	// Then the error returned is nil.
	assert.Nil(t, err)
}
