package postgres_test

import (
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/braejan/go-transactions-summary/internal/domain/transaction/repository/postgres"
	voPostgres "github.com/braejan/go-transactions-summary/internal/valueobject/postgres"
	"github.com/braejan/go-transactions-summary/internal/valueobject/postgres/mock"
	"github.com/braejan/go-transactions-summary/internal/valueobject/transaction"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// TestGetByIDErrOpeningDatabase tests the error returned when opening the database.
func TestGetByIDErrOpeningDatabase(t *testing.T) {
	// Given a valid configuration.
	configuration := voPostgres.GetPostgresConfigurationFromEnv()
	dbBase := mock.NewMockBasePostgresDatabase()
	// And a valid transaction repository.
	txRepo, err := postgres.NewPostgresTransactionRepository(configuration, dbBase)
	assert.Nil(t, err)
	// And a valid uuid.UUID txID.
	accID := uuid.New()
	// And a mocked response when calling Open.
	dbBase.On("Open", configuration.GetDataSourceName()).Return(nil, voPostgres.ErrOpeningDatabase)
	// When getting a account by ID.
	_, err = txRepo.GetByID(accID)
	// Then the error returned is ErrOpeningDatabase.
	assert.NotNil(t, err)
	assert.Equal(t, voPostgres.ErrOpeningDatabase, err)
}

// TestGetByIDErrOpeningTransaction tests the error returned when opening the transaction.
func TestGetByIDErrOpeningTransaction(t *testing.T) {
	// Given a valid configuration.
	configuration := voPostgres.GetPostgresConfigurationFromEnv()
	dbBase := mock.NewMockBasePostgresDatabase()
	// And a valid transaction repository.
	txRepo, err := postgres.NewPostgresTransactionRepository(configuration, dbBase)
	assert.Nil(t, err)
	// And a valid uuid.UUID txID.
	accID := uuid.New()
	// And a mocked database.
	db, _, _ := sqlmock.New()
	// And a mocked response when calling Open.
	dbBase.On("Open", configuration.GetDataSourceName()).Return(db, nil)
	// And a mocked response when calling Close.
	dbBase.On("Close", db).Return(nil)
	// And a mocked response when calling Begin.
	dbBase.On("BeginTx", db).Return(nil, voPostgres.ErrBeginningTransaction)
	// When getting a account by ID.
	_, err = txRepo.GetByID(accID)
	// Then the error returned is ErrBeginningTransaction.
	assert.NotNil(t, err)
	assert.Equal(t, voPostgres.ErrBeginningTransaction, err)
}

// TestGetByIDErrQueryingDatabase tests the error returned when querying the database.
func TestGetByIDErrQueryingDatabase(t *testing.T) {
	// Given a valid configuration.
	configuration := voPostgres.GetPostgresConfigurationFromEnv()
	dbBase := mock.NewMockBasePostgresDatabase()
	// And a valid transaction repository.
	txRepo, err := postgres.NewPostgresTransactionRepository(configuration, dbBase)
	assert.Nil(t, err)
	// And a valid uuid.UUID txID.
	txID := uuid.New()
	// And a mocked database.
	db, _, _ := sqlmock.New()
	// And a mocked response when calling Open.
	dbBase.On("Open", configuration.GetDataSourceName()).Return(db, nil)
	// And a mocked response when calling Close.
	dbBase.On("Close", db).Return(nil)
	// And a mocked response when calling Begin.
	tx, _ := db.Begin()
	dbBase.On("BeginTx", db).Return(tx, nil)
	// And a mocked response when querying the database.
	dbBase.On("Query", tx, "SELECT id, accountid, amount, date, origin FROM transactions WHERE id = $1", []interface{}{txID}).Return(nil, voPostgres.ErrQueryingDatabase)
	// When getting a account by ID.
	_, err = txRepo.GetByID(txID)
	// Then the error returned is ErrQueryingDatabase.
	assert.NotNil(t, err)
	assert.Equal(t, transaction.ErrQueryingTransactionByID, err)
}

// TestGetByIDErrScanningRow tests the error returned when scanning the row.
func TestGetByIDErrScanningRow(t *testing.T) {
	// Given a valid configuration.
	configuration := voPostgres.GetPostgresConfigurationFromEnv()
	dbBase := voPostgres.NewBasePostgresDatabase()
	dbBaseMocked := mock.NewMockBasePostgresDatabase()
	// And a mocked database.
	db, dbMocked, _ := sqlmock.New()
	defer db.Close()
	// And a mocked transaction
	dbMocked.ExpectBegin()
	// And a valid uuid.UUID txID.
	txID := uuid.New()
	// And a mocked response when calling Open.
	dbBaseMocked.On("Open", configuration.GetDataSourceName()).Return(db, nil)
	// And a mocked response when calling BeginTx.
	tx, _ := db.BeginTx(context.Background(), nil)
	dbBaseMocked.On("BeginTx", db).Return(tx, nil)
	// And a mocked response when calling Close.
	dbBaseMocked.On("Close", db).Return(nil)
	// And a mocked response when calling Query.
	expected := sqlmock.NewRows([]string{"column1", "column2", "column3"}).AddRow(true, false, false)
	dbMocked.ExpectQuery("SELECT (.+) FROM transactions WHERE id = (.+)").WithArgs(txID).WillReturnRows(expected)
	rows, err := dbBase.Query(tx, "SELECT id, accountid, amount, date, origin FROM transactions WHERE id = $1", txID)
	assert.Nil(t, err)
	dbBaseMocked.On("Query", tx, "SELECT id, accountid, amount, date, origin FROM transactions WHERE id = $1", []interface{}{txID}).Return(rows, nil)
	// And a valid transaction repository
	transactionRepo, err := postgres.NewPostgresTransactionRepository(configuration, dbBaseMocked)
	assert.Nil(t, err)
	// When GetByID is called.
	_, err = transactionRepo.GetByID(txID)
	// Then the error returned should be ErrScanningUser.
	assert.NotNil(t, err)
	assert.Equal(t, transaction.ErrScanningTransactionByID, err)
}

// TestGetByIDSucess tests the success when getting a transaction by ID.
func TestGetByIDSucess(t *testing.T) {
	// Given a valid configuration.
	configuration := voPostgres.GetPostgresConfigurationFromEnv()
	dbBase := voPostgres.NewBasePostgresDatabase()
	dbBaseMocked := mock.NewMockBasePostgresDatabase()
	// And a mocked database.
	db, dbMocked, _ := sqlmock.New()
	defer db.Close()
	// And a mocked transaction
	dbMocked.ExpectBegin()
	// And a valid uuid.UUID txID.
	txID := uuid.New()
	// And a mocked response when calling Open.
	dbBaseMocked.On("Open", configuration.GetDataSourceName()).Return(db, nil)
	// And a mocked response when calling BeginTx.
	tx, _ := db.BeginTx(context.Background(), nil)
	dbBaseMocked.On("BeginTx", db).Return(tx, nil)
	// And a mocked response when calling Close.
	dbBaseMocked.On("Close", db).Return(nil)
	// And a mocked response when calling Query.
	expected := sqlmock.NewRows([]string{"id", "accountid", "amount", "date", "origin"}).AddRow(txID, uuid.New(), 100.0, time.Now(), "txns.csv")
	dbMocked.ExpectQuery("SELECT (.+) FROM transactions WHERE id = (.+)").WithArgs(txID).WillReturnRows(expected)
	rows, err := dbBase.Query(tx, "SELECT id, accountid, amount, date, origin FROM transactions WHERE id = $1", txID)
	assert.Nil(t, err)
	dbBaseMocked.On("Query", tx, "SELECT id, accountid, amount, date, origin FROM transactions WHERE id = $1", []interface{}{txID}).Return(rows, nil)
	// And a valid transaction repository
	transactionRepo, err := postgres.NewPostgresTransactionRepository(configuration, dbBaseMocked)
	assert.Nil(t, err)
	// When GetByID is called.
	transaction, err := transactionRepo.GetByID(txID)
	// Then the error returned should be ErrScanningUser.
	assert.Nil(t, err)
	assert.NotNil(t, transaction)
	assert.Equal(t, txID, transaction.ID)
	assert.Equal(t, 100.0, transaction.Amount)
	assert.Equal(t, "txns.csv", transaction.Origin)
}

// TestGetByAccountIDErrOpeningDatabase tests the error returned when opening the database.
func TestGetByAccountIDErrOpeningDatabase(t *testing.T) {
	// Given a valid configuration.
	configuration := voPostgres.GetPostgresConfigurationFromEnv()
	dbBase := mock.NewMockBasePostgresDatabase()
	// And a valid transaction repository.
	txRepo, err := postgres.NewPostgresTransactionRepository(configuration, dbBase)
	assert.Nil(t, err)
	// And a valid uuid.UUID accountID.
	accountID := uuid.New()
	// And a mocked response when calling Open.
	dbBase.On("Open", configuration.GetDataSourceName()).Return(nil, voPostgres.ErrOpeningDatabase)
	// When getting a account by ID.
	_, err = txRepo.GetByAccountID(accountID)
	// Then the error returned is ErrOpeningDatabase.
	assert.NotNil(t, err)
	assert.Equal(t, voPostgres.ErrOpeningDatabase, err)
}

// TestGetByAccountIDErrBeginningTransaction tests the error returned when beginning the transaction.
func TestGetByAccountIDErrBeginningTransaction(t *testing.T) {
	// Given a valid configuration.
	configuration := voPostgres.GetPostgresConfigurationFromEnv()
	dbBase := mock.NewMockBasePostgresDatabase()
	// And a valid transaction repository.
	txRepo, err := postgres.NewPostgresTransactionRepository(configuration, dbBase)
	assert.Nil(t, err)
	// And a valid uuid.UUID accountID.
	accountID := uuid.New()
	// And a sqlmock database.
	db, _, _ := sqlmock.New()
	// And a mocked response when calling Open.
	dbBase.On("Open", configuration.GetDataSourceName()).Return(db, nil)
	// And a mocked response when calling Close.
	dbBase.On("Close", db).Return(nil)
	// And a mocked response when calling BeginTx.
	dbBase.On("BeginTx", db).Return(nil, voPostgres.ErrBeginningTransaction)
	// When getting a account by ID.
	_, err = txRepo.GetByAccountID(accountID)
	// Then the error returned is ErrBeginningTransaction.
	assert.NotNil(t, err)
	assert.Equal(t, voPostgres.ErrBeginningTransaction, err)
}

// TestGetByAccountIDErrQuerying tests the error returned when querying the database.
func TestGetByAccountIDErrQuerying(t *testing.T) {
	// Given a valid configuration.
	configuration := voPostgres.GetPostgresConfigurationFromEnv()
	dbBaseMocked := mock.NewMockBasePostgresDatabase()
	// And a valid transaction repository.
	txRepo, err := postgres.NewPostgresTransactionRepository(configuration, dbBaseMocked)
	assert.Nil(t, err)
	// And a valid uuid.UUID accountID.
	accountID := uuid.New()
	// And a sqlmock database.
	db, _, _ := sqlmock.New()
	// And a mocked response when calling Open.
	dbBaseMocked.On("Open", configuration.GetDataSourceName()).Return(db, nil)
	// And a mocked response when calling Close.
	dbBaseMocked.On("Close", db).Return(nil)
	// And a mocked response when calling BeginTx.
	tx, _ := db.BeginTx(context.Background(), nil)
	dbBaseMocked.On("BeginTx", db).Return(tx, nil)
	// And a mocked response when calling Query.
	dbBaseMocked.On("Query", tx, "SELECT id, accountid, amount, date, origin FROM transactions WHERE accountid = $1", []interface{}{accountID}).Return(nil, voPostgres.ErrQueryingDatabase)
	// When getting a account by ID.
	_, err = txRepo.GetByAccountID(accountID)
	// Then the error returned is ErrQuerying.
	assert.NotNil(t, err)
	assert.Equal(t, transaction.ErrQueryingTransactionsByAccountID, err)
}

// TestGetByAccountIDErrScanning tests the error returned when scanning the database.
func TestGetByAccountIDErrScanning(t *testing.T) {
	// Given a valid configuration.
	configuration := voPostgres.GetPostgresConfigurationFromEnv()
	dbBaseMocked := mock.NewMockBasePostgresDatabase()
	dbBase := voPostgres.NewBasePostgresDatabase()
	// And a valid transaction repository.
	txRepo, err := postgres.NewPostgresTransactionRepository(configuration, dbBaseMocked)
	assert.Nil(t, err)
	// And a valid uuid.UUID accountID.
	accountID := uuid.New()
	// And a sqlmock database.
	db, dbMocked, _ := sqlmock.New()
	defer db.Close()
	dbMocked.ExpectBegin()
	// And a mocked response when calling Open.
	dbBaseMocked.On("Open", configuration.GetDataSourceName()).Return(db, nil)
	// And a mocked response when calling Close.
	dbBaseMocked.On("Close", db).Return(nil)
	// And a mocked response when calling BeginTx.
	tx, _ := db.BeginTx(context.Background(), nil)
	dbBaseMocked.On("BeginTx", db).Return(tx, nil)
	// And a mocked response when calling Query.
	expected := sqlmock.NewRows([]string{"column1", "column2"}).AddRow("100.0", "txns.csv")
	dbMocked.ExpectQuery("SELECT (.+) FROM transactions WHERE accountid = (.+)").WithArgs(accountID).WillReturnRows(expected)
	rows, err := dbBase.Query(tx, "SELECT id, accountid, amount, date, origin FROM transactions WHERE accountid = $1", accountID)
	assert.Nil(t, err)
	dbBaseMocked.On("Query", tx, "SELECT id, accountid, amount, date, origin FROM transactions WHERE accountid = $1", []interface{}{accountID}).Return(rows, nil)
	// When getting a account by ID.
	_, err = txRepo.GetByAccountID(accountID)
	// Then the error returned is ErrScanning.
	assert.NotNil(t, err)
	assert.Equal(t, transaction.ErrScanningTransactionsByAccountID, err)
}

// TestGetByAccountIDSucess tests the success when getting a transaction by account ID.
func TestGetByAccountIDSucess(t *testing.T) {
	// Given a valid configuration.
	configuration := voPostgres.GetPostgresConfigurationFromEnv()
	dbBaseMocked := mock.NewMockBasePostgresDatabase()
	dbBase := voPostgres.NewBasePostgresDatabase()
	// And a valid transaction repository.
	txRepo, err := postgres.NewPostgresTransactionRepository(configuration, dbBaseMocked)
	assert.Nil(t, err)
	// And a valid uuid.UUID accountID.
	accountID := uuid.New()
	// And a sqlmock database.
	db, dbMocked, _ := sqlmock.New()
	defer db.Close()
	dbMocked.ExpectBegin()
	// And a mocked response when calling Open.
	dbBaseMocked.On("Open", configuration.GetDataSourceName()).Return(db, nil)
	// And a mocked response when calling Close.
	dbBaseMocked.On("Close", db).Return(nil)
	// And a mocked response when calling BeginTx.
	tx, _ := db.BeginTx(context.Background(), nil)
	dbBaseMocked.On("BeginTx", db).Return(tx, nil)
	// And a mocked response when calling Query.
	expected := sqlmock.NewRows([]string{"id", "accountid", "amount", "date", "origin"})
	expected.AddRow(uuid.New(), accountID, 100.0, time.Now(), "txns.csv")
	expected.AddRow(uuid.New(), accountID, -200.0, time.Now(), "txns.csv")
	expected.AddRow(uuid.New(), accountID, 300.0, time.Now(), "txns.csv")
	dbMocked.ExpectQuery("SELECT (.+) FROM transactions WHERE accountid = (.+)").WithArgs(accountID).WillReturnRows(expected)
	rows, err := dbBase.Query(tx, "SELECT id, accountid, amount, date, origin FROM transactions WHERE accountid = $1", accountID)
	assert.Nil(t, err)
	dbBaseMocked.On("Query", tx, "SELECT id, accountid, amount, date, origin FROM transactions WHERE accountid = $1", []interface{}{accountID}).Return(rows, nil)
	// When getting a account by ID.
	transactions, err := txRepo.GetByAccountID(accountID)
	// Then the error returned is nil.
	assert.Nil(t, err)
	// And the transactions returned are not nil.
	assert.NotNil(t, transactions)
	// And the transactions returned are 3.
	assert.Equal(t, 3, len(transactions))
}

// TestGetCreditsByAccountIDErrOpening tests the error returned when opening the database.
func TestGetCreditsByAccountIDErrOpening(t *testing.T) {
	// Given a valid configuration.
	configuration := voPostgres.GetPostgresConfigurationFromEnv()
	dbBaseMocked := mock.NewMockBasePostgresDatabase()
	// And a valid transaction repository.
	txRepo, err := postgres.NewPostgresTransactionRepository(configuration, dbBaseMocked)
	assert.Nil(t, err)
	// And a valid uuid.UUID accountID.
	accountID := uuid.New()
	// And a mocked response when calling Open.
	dbBaseMocked.On("Open", configuration.GetDataSourceName()).Return(nil, voPostgres.ErrOpeningDatabase)
	// When getting a account by ID.
	_, err = txRepo.GetCreditsByAccountID(accountID)
	// Then the error returned is ErrOpening.
	assert.NotNil(t, err)
	assert.Equal(t, voPostgres.ErrOpeningDatabase, err)
}

// TestGetCreditsByAccountIDErrBeginningTx tests the error returned when beginning the transaction.
func TestGetCreditsByAccountIDErrBeginningTx(t *testing.T) {
	// Given a valid configuration.
	configuration := voPostgres.GetPostgresConfigurationFromEnv()
	dbBaseMocked := mock.NewMockBasePostgresDatabase()
	// And a valid transaction repository.
	txRepo, err := postgres.NewPostgresTransactionRepository(configuration, dbBaseMocked)
	assert.Nil(t, err)
	// And a valid uuid.UUID accountID.
	accountID := uuid.New()
	// And a sqlmock database.
	db, _, _ := sqlmock.New()
	// And a mocked response when calling Open.
	dbBaseMocked.On("Open", configuration.GetDataSourceName()).Return(db, nil)
	// And a mocked response when calling Close.
	dbBaseMocked.On("Close", db).Return(nil)
	// And a mocked response when calling BeginTx.
	dbBaseMocked.On("BeginTx", db).Return(nil, voPostgres.ErrBeginningTransaction)
	// When getting a account by ID.
	_, err = txRepo.GetCreditsByAccountID(accountID)
	// Then the error returned is ErrBeginningTx.
	assert.NotNil(t, err)
	assert.Equal(t, voPostgres.ErrBeginningTransaction, err)
}

// TestGetCreditsByAccountIDErrQuerying tests the error returned when querying the database.
func TestGetCreditsByAccountIDErrQuerying(t *testing.T) {
	// Given a valid configuration.
	configuration := voPostgres.GetPostgresConfigurationFromEnv()
	dbBaseMocked := mock.NewMockBasePostgresDatabase()
	// And a valid transaction repository.
	txRepo, err := postgres.NewPostgresTransactionRepository(configuration, dbBaseMocked)
	assert.Nil(t, err)
	// And a valid uuid.UUID accountID.
	accountID := uuid.New()
	// And a sqlmock database.
	db, dbMocked, _ := sqlmock.New()
	defer db.Close()
	dbMocked.ExpectBegin()
	// And a mocked response when calling Open.
	dbBaseMocked.On("Open", configuration.GetDataSourceName()).Return(db, nil)
	// And a mocked response when calling Close.
	dbBaseMocked.On("Close", db).Return(nil)
	// And a mocked response when calling BeginTx.
	tx, _ := db.BeginTx(context.Background(), nil)
	dbBaseMocked.On("BeginTx", db).Return(tx, nil)
	dbBaseMocked.On(
		"Query",
		tx,
		"SELECT id, accountid, amount, date, origin FROM transactions WHERE accountid = $1 AND operation = 'credit'",
		[]interface{}{accountID}).Return(nil, voPostgres.ErrQueryingDatabase)
	// When getting a account by ID.
	_, err = txRepo.GetCreditsByAccountID(accountID)
	// Then the error returned is ErrQuerying.
	assert.NotNil(t, err)
	assert.Equal(t, transaction.ErrQueryingCreditsByAccountID, err)
}

// TestGetCreditsByAccountIDErrScanning tests the error returned when scanning the database.
func TestGetCreditsByAccountIDErrScanning(t *testing.T) {
	// Given a valid configuration.
	configuration := voPostgres.GetPostgresConfigurationFromEnv()
	dbBaseMocked := mock.NewMockBasePostgresDatabase()
	dbBase := voPostgres.NewBasePostgresDatabase()
	// And a valid transaction repository.
	txRepo, err := postgres.NewPostgresTransactionRepository(configuration, dbBaseMocked)
	assert.Nil(t, err)
	// And a valid uuid.UUID txID.
	txID := uuid.New()
	// And a sqlmock database.
	db, dbMocked, _ := sqlmock.New()
	dbMocked.ExpectBegin()
	defer db.Close()
	// And a mocked response when calling Open.
	dbBaseMocked.On("Open", configuration.GetDataSourceName()).Return(db, nil)
	// And a mocked response when calling Close.
	dbBaseMocked.On("Close", db).Return(nil)
	// And a mocked response when calling BeginTx.
	tx, _ := db.BeginTx(context.Background(), nil)
	dbBaseMocked.On("BeginTx", db).Return(tx, nil)
	expected := sqlmock.NewRows([]string{"column1", "colum2"})
	expected.AddRow(100.0, "txns.csv")
	expected.AddRow(200.0, "txns.csv")
	expected.AddRow(-300.0, "txns.csv")
	dbMocked.ExpectQuery("SELECT (.+) FROM transactions WHERE accountid = (.+) AND operation = 'credit'").WithArgs(txID).WillReturnRows(expected)
	rows, err := dbBase.Query(tx, "SELECT id, accountid, amount, date, origin FROM transactions WHERE accountid = $1 AND operation = 'credit'", txID)
	assert.Nil(t, err)
	dbBaseMocked.On(
		"Query",
		tx,
		"SELECT id, accountid, amount, date, origin FROM transactions WHERE accountid = $1 AND operation = 'credit'",
		[]interface{}{txID}).Return(rows, nil)
	// When getting a account by ID.
	_, err = txRepo.GetCreditsByAccountID(txID)
	// Then the error returned is ErrScanning.
	assert.NotNil(t, err)
	assert.Equal(t, transaction.ErrScanningCreditsByAccountID, err)
}

// TestGetCreditsByAccountIDSucess tests the success when getting the credits by account ID.
func TestGetCreditsByAccountIDSucess(t *testing.T) {
	// Given a valid configuration.
	configuration := voPostgres.GetPostgresConfigurationFromEnv()
	dbBaseMocked := mock.NewMockBasePostgresDatabase()
	dbBase := voPostgres.NewBasePostgresDatabase()
	// And a valid transaction repository.
	txRepo, err := postgres.NewPostgresTransactionRepository(configuration, dbBaseMocked)
	assert.Nil(t, err)
	// And a valid uuid.UUID txID.
	txID := uuid.New()
	// And a sqlmock database.
	db, dbMocked, _ := sqlmock.New()
	dbMocked.ExpectBegin()
	defer db.Close()
	// And a mocked response when calling Open.
	dbBaseMocked.On("Open", configuration.GetDataSourceName()).Return(db, nil)
	// And a mocked response when calling Close.
	dbBaseMocked.On("Close", db).Return(nil)
	// And a mocked response when calling BeginTx.
	tx, _ := db.BeginTx(context.Background(), nil)
	dbBaseMocked.On("BeginTx", db).Return(tx, nil)
	expected := sqlmock.NewRows([]string{"id", "accountid", "amount", "date", "origin"})
	expected.AddRow(uuid.New(), txID, 100.0, time.Now(), "txns.csv")
	expected.AddRow(uuid.New(), txID, 200.0, time.Now(), "txns.csv")
	expected.AddRow(uuid.New(), txID, -300.0, time.Now(), "txns.csv")
	dbMocked.ExpectQuery("SELECT (.+) FROM transactions WHERE accountid = (.+) AND operation = 'credit'").WithArgs(txID).WillReturnRows(expected)
	rows, err := dbBase.Query(tx, "SELECT id, accountid, amount, date, origin FROM transactions WHERE accountid = $1 AND operation = 'credit'", txID)
	assert.Nil(t, err)
	dbBaseMocked.On(
		"Query",
		tx,
		"SELECT id, accountid, amount, date, origin FROM transactions WHERE accountid = $1 AND operation = 'credit'",
		[]interface{}{txID}).Return(rows, nil)
	// When getting a account by ID.
	txs, err := txRepo.GetCreditsByAccountID(txID)
	// Then the error returned is Nil
	assert.Nil(t, err)
	assert.NotNil(t, txs)
	assert.Equal(t, 3, len(txs))
	assert.Equal(t, -300.0, txs[2].Amount)
}

// TestGetDebitsByAccountIDErrOpening tests the error returned when opening the database.
func TestGetDebitsByAccountIDErrOpening(t *testing.T) {
	// Given a valid configuration.
	configuration := voPostgres.GetPostgresConfigurationFromEnv()
	dbBaseMocked := mock.NewMockBasePostgresDatabase()
	// And a valid transaction repository.
	txRepo, err := postgres.NewPostgresTransactionRepository(configuration, dbBaseMocked)
	assert.Nil(t, err)
	// And a valid uuid.UUID accountID.
	accountID := uuid.New()
	// And a mocked response when calling Open.
	dbBaseMocked.On("Open", configuration.GetDataSourceName()).Return(nil, voPostgres.ErrOpeningDatabase)
	// When getting a account by ID.
	_, err = txRepo.GetDebitsByAccountID(accountID)
	// Then the error returned is ErrOpening.
	assert.NotNil(t, err)
	assert.Equal(t, voPostgres.ErrOpeningDatabase, err)
}

// TestGetDebitsByAccountErrBeginTx tests the error returned when beginning the transaction.
func TestGetDebitsByAccountErrBeginTx(t *testing.T) {
	// Given a valid configuration.
	configuration := voPostgres.GetPostgresConfigurationFromEnv()
	dbBaseMocked := mock.NewMockBasePostgresDatabase()
	// And a valid transaction repository.
	txRepo, err := postgres.NewPostgresTransactionRepository(configuration, dbBaseMocked)
	assert.Nil(t, err)
	// And a valid uuid.UUID accountID.
	accountID := uuid.New()
	// And a sqlmock database.
	db, _, _ := sqlmock.New()
	defer db.Close()
	// And a mocked response when calling Open.
	dbBaseMocked.On("Open", configuration.GetDataSourceName()).Return(db, nil)
	// And a mocked response when calling Close.
	dbBaseMocked.On("Close", db).Return(nil)
	// And a mocked response when calling BeginTx.
	dbBaseMocked.On("BeginTx", db).Return(nil, voPostgres.ErrBeginningTransaction)
	// When getting a account by ID.
	_, err = txRepo.GetDebitsByAccountID(accountID)
	// Then the error returned is ErrBeginningTransaction.
	assert.NotNil(t, err)
	assert.Equal(t, voPostgres.ErrBeginningTransaction, err)
}

// TestGetDebitsByAccountIDErrQuerying tests the error returned when querying the database.
func TestGetDebitsByAccountIDErrQuerying(t *testing.T) {
	// Given a valid configuration.
	configuration := voPostgres.GetPostgresConfigurationFromEnv()
	dbBaseMocked := mock.NewMockBasePostgresDatabase()
	// And a valid transaction repository.
	txRepo, err := postgres.NewPostgresTransactionRepository(configuration, dbBaseMocked)
	assert.Nil(t, err)
	// And a valid uuid.UUID accountID.
	accountID := uuid.New()
	// And a sqlmock database.
	db, dbMocked, _ := sqlmock.New()
	dbMocked.ExpectBegin()
	defer db.Close()
	// And a mocked response when calling Open.
	dbBaseMocked.On("Open", configuration.GetDataSourceName()).Return(db, nil)
	// And a mocked response when calling Close.
	dbBaseMocked.On("Close", db).Return(nil)
	// And a mocked response when calling BeginTx.
	tx, _ := db.BeginTx(context.Background(), nil)
	dbBaseMocked.On("BeginTx", db).Return(tx, nil)
	dbBaseMocked.On(
		"Query",
		tx,
		"SELECT id, accountid, amount, date, origin FROM transactions WHERE accountid = $1 AND operation = 'debit'",
		[]interface{}{accountID}).Return(nil, voPostgres.ErrQueryingDatabase)
	// When getting a account by ID.
	_, err = txRepo.GetDebitsByAccountID(accountID)
	// Then the error returned is ErrQueryingDatabase.
	assert.NotNil(t, err)
	assert.Equal(t, transaction.ErrQueryingDebitsByAccountID, err)
}

// TestGetDebitsByAccountIDErrScanning tests the error returned when scanning the database response.
func TestGetDebitsByAccountIDErrScanning(t *testing.T) {
	// Given a valid configuration.
	configuration := voPostgres.GetPostgresConfigurationFromEnv()
	dbBaseMocked := mock.NewMockBasePostgresDatabase()
	dbBase := voPostgres.NewBasePostgresDatabase()
	// And a valid transaction repository.
	txRepo, err := postgres.NewPostgresTransactionRepository(configuration, dbBaseMocked)
	assert.Nil(t, err)
	// And a valid uuid.UUID accountID.
	accountID := uuid.New()
	// And a sqlmock database.
	db, dbMocked, _ := sqlmock.New()
	dbMocked.ExpectBegin()
	defer db.Close()
	// And a mocked response when calling Open.
	dbBaseMocked.On("Open", configuration.GetDataSourceName()).Return(db, nil)
	// And a mocked response when calling Close.
	dbBaseMocked.On("Close", db).Return(nil)
	// And a mocked response when calling BeginTx.
	tx, _ := db.BeginTx(context.Background(), nil)
	dbBaseMocked.On("BeginTx", db).Return(tx, nil)
	expected := sqlmock.NewRows([]string{"column1", "column2"})
	expected.AddRow("invalid", "txns.csv")
	dbMocked.ExpectQuery("SELECT (.+) FROM transactions WHERE accountid = (.+) AND operation = 'debit'").WithArgs(accountID).WillReturnRows(expected)
	rows, err := dbBase.Query(tx, "SELECT id, accountid, amount, date, origin FROM transactions WHERE accountid = $1 AND operation = 'debit'", accountID)
	assert.Nil(t, err)
	dbBaseMocked.On(
		"Query",
		tx,
		"SELECT id, accountid, amount, date, origin FROM transactions WHERE accountid = $1 AND operation = 'debit'",
		[]interface{}{accountID}).Return(rows, nil)
	// When getting a account by ID.
	_, err = txRepo.GetDebitsByAccountID(accountID)
	// Then the error returned is ErrScanning.
	assert.NotNil(t, err)
	assert.Equal(t, transaction.ErrScanningDebitsByAccountID, err)
}

// TestGetDebitsByAccountIDSucess tests the success in getting the debits by account ID.
func TestGetDebitsByAccountIDSucess(t *testing.T) {
	// Given a valid configuration.
	configuration := voPostgres.GetPostgresConfigurationFromEnv()
	dbBaseMocked := mock.NewMockBasePostgresDatabase()
	dbBase := voPostgres.NewBasePostgresDatabase()
	// And a valid transaction repository.
	txRepo, err := postgres.NewPostgresTransactionRepository(configuration, dbBaseMocked)
	assert.Nil(t, err)
	// And a valid uuid.UUID accountID.
	accountID := uuid.New()
	// And a sqlmock database.
	db, dbMocked, _ := sqlmock.New()
	dbMocked.ExpectBegin()
	defer db.Close()
	// And a mocked response when calling Open.
	dbBaseMocked.On("Open", configuration.GetDataSourceName()).Return(db, nil)
	// And a mocked response when calling Close.
	dbBaseMocked.On("Close", db).Return(nil)
	// And a mocked response when calling BeginTx.
	tx, _ := db.BeginTx(context.Background(), nil)
	dbBaseMocked.On("BeginTx", db).Return(tx, nil)
	expected := sqlmock.NewRows([]string{"id", "accountid", "amount", "date", "origin"})
	expected.AddRow(uuid.New(), accountID, 100.00, time.Now(), "txns.csv")
	expected.AddRow(uuid.New(), accountID, 200.00, time.Now(), "txns.csv")
	expected.AddRow(uuid.New(), accountID, -300.00, time.Now(), "txns.csv")
	dbMocked.ExpectQuery("SELECT (.+) FROM transactions WHERE accountid = (.+) AND operation = 'debit'").WithArgs(accountID).WillReturnRows(expected)
	rows, err := dbBase.Query(tx, "SELECT id, accountid, amount, date, origin FROM transactions WHERE accountid = $1 AND operation = 'debit'", accountID)
	assert.Nil(t, err)
	dbBaseMocked.On(
		"Query",
		tx,
		"SELECT id, accountid, amount, date, origin FROM transactions WHERE accountid = $1 AND operation = 'debit'",
		[]interface{}{accountID}).Return(rows, nil)
	// When getting a account by ID.
	txs, err := txRepo.GetDebitsByAccountID(accountID)
	// Then the error returned is Nil.
	assert.Nil(t, err)
	assert.NotNil(t, txs)
	assert.Equal(t, 3, len(txs))
	assert.Equal(t, -300.00, txs[2].Amount)

}

// TestGetTransactionsByOriginInvalidOrigin tests the error returned when the origin is invalid.
func TestGetTransactionsByOriginInvalidOrigin(t *testing.T) {
	// Given a valid configuration.
	configuration := voPostgres.GetPostgresConfigurationFromEnv()
	dbBaseMocked := mock.NewMockBasePostgresDatabase()
	// And a valid transaction repository.
	txRepo, err := postgres.NewPostgresTransactionRepository(configuration, dbBaseMocked)
	assert.Nil(t, err)
	// And a invalid origin.
	origin := ""
	// When getting a account by ID.
	_, err = txRepo.GetTransactionsByOrigin(origin)
	// Then the error returned is ErrEmptyOrigin.
	assert.NotNil(t, err)
	assert.Equal(t, transaction.ErrEmptyOrigin, err)
}

// TestGetTransactionsByOriginErrOpening tests the error returned when opening the database.
func TestGetTransactionsByOriginErrOpening(t *testing.T) {
	// Given a valid configuration.
	configuration := voPostgres.GetPostgresConfigurationFromEnv()
	dbBaseMocked := mock.NewMockBasePostgresDatabase()
	// And a valid transaction repository.
	txRepo, err := postgres.NewPostgresTransactionRepository(configuration, dbBaseMocked)
	assert.Nil(t, err)
	// And a valid origin.
	origin := "txns.csv"
	// And a mocked response when calling Open.
	dbBaseMocked.On("Open", configuration.GetDataSourceName()).Return(nil, voPostgres.ErrOpeningDatabase)
	// When getting a account by ID.
	_, err = txRepo.GetTransactionsByOrigin(origin)
	// Then the error returned is ErrOpeningDatabase.
	assert.NotNil(t, err)
	assert.Equal(t, voPostgres.ErrOpeningDatabase, err)
}

// TestGetTransactionsByOriginErrBegginingTx tests the error returned when beginning the transaction.
func TestGetTransactionsByOriginErrBegginingTx(t *testing.T) {
	// Given a valid configuration.
	configuration := voPostgres.GetPostgresConfigurationFromEnv()
	dbBaseMocked := mock.NewMockBasePostgresDatabase()
	// And a valid transaction repository.
	txRepo, err := postgres.NewPostgresTransactionRepository(configuration, dbBaseMocked)
	assert.Nil(t, err)
	// And a valid origin.
	origin := "txns.csv"
	// And a sqlmock database.
	db, _, _ := sqlmock.New()
	db.Close()
	// And a mocked response when calling Open.
	dbBaseMocked.On("Open", configuration.GetDataSourceName()).Return(db, nil)
	// And a mocked response when calling Close.
	dbBaseMocked.On("Close", db).Return(nil)
	// And a mocked response when calling BeginTx.
	dbBaseMocked.On("BeginTx", db).Return(nil, voPostgres.ErrBeginningTransaction)
	// When getting a account by ID.
	_, err = txRepo.GetTransactionsByOrigin(origin)
	// Then the error returned is ErrBeginningTx.
	assert.NotNil(t, err)
	assert.Equal(t, voPostgres.ErrBeginningTransaction, err)
}

// TestGetTransactionsByOriginErrQuerying tests the error returned when querying the database.
func TestGetTransactionsByOriginErrQuerying(t *testing.T) {
	// Given a valid configuration.
	configuration := voPostgres.GetPostgresConfigurationFromEnv()
	dbBaseMocked := mock.NewMockBasePostgresDatabase()
	// And a valid transaction repository.
	txRepo, err := postgres.NewPostgresTransactionRepository(configuration, dbBaseMocked)
	assert.Nil(t, err)
	// And a valid origin.
	origin := "txns.csv"
	// And a sqlmock database.
	db, dbMocked, _ := sqlmock.New()
	dbMocked.ExpectBegin()
	defer db.Close()
	// And a mocked response when calling Open.
	dbBaseMocked.On("Open", configuration.GetDataSourceName()).Return(db, nil)
	// And a mocked response when calling Close.
	dbBaseMocked.On("Close", db).Return(nil)
	// And a mocked response when calling BeginTx.
	tx, _ := db.BeginTx(context.Background(), nil)
	dbBaseMocked.On("BeginTx", db).Return(tx, nil)
	dbBaseMocked.On("Query", tx, "SELECT id, accountid, amount, date, origin FROM transactions WHERE origin = $1", []interface{}{origin}).Return(nil, voPostgres.ErrQueryingDatabase)
	// When getting a account by ID.
	_, err = txRepo.GetTransactionsByOrigin(origin)
	// Then the error returned is ErrQueryingDatabase.
	assert.NotNil(t, err)
	assert.Equal(t, transaction.ErrQueryingTransactionsByOrigin, err)
}

// TestGetTransactionsByOriginSuccess tests the success when querying the database.
func TestGetTransactionsByOriginSuccess(t *testing.T) {
	// Given a valid configuration.
	configuration := voPostgres.GetPostgresConfigurationFromEnv()
	dbBaseMocked := mock.NewMockBasePostgresDatabase()
	dbBase := voPostgres.NewBasePostgresDatabase()
	// And a valid transaction repository.
	txRepo, err := postgres.NewPostgresTransactionRepository(configuration, dbBaseMocked)
	assert.Nil(t, err)
	// And a valid origin.
	origin := "txns.csv"
	// And a sqlmock database.
	db, dbMocked, _ := sqlmock.New()
	dbMocked.ExpectBegin()
	defer db.Close()
	// And a mocked response when calling Open.
	dbBaseMocked.On("Open", configuration.GetDataSourceName()).Return(db, nil)
	// And a mocked response when calling Close.
	dbBaseMocked.On("Close", db).Return(nil)
	// And a mocked response when calling BeginTx.
	tx, _ := db.BeginTx(context.Background(), nil)
	dbBaseMocked.On("BeginTx", db).Return(tx, nil)
	// And a mocked response when calling Query.
	expected := sqlmock.NewRows([]string{"id", "accountid", "amount", "date", "origin"})
	accountID := uuid.New()
	expected.AddRow(uuid.New(), accountID, 100.00, time.Now(), "txns.csv")
	expected.AddRow(uuid.New(), accountID, 200.00, time.Now(), "txns.csv")
	expected.AddRow(uuid.New(), accountID, -300.00, time.Now(), "txns.csv")
	dbMocked.ExpectQuery("SELECT (.+) FROM transactions WHERE origin = (.+)").WithArgs(origin).WillReturnRows(expected)
	rows, err := dbBase.Query(tx, "SELECT id, accountid, amount, date, origin FROM transactions WHERE origin = $1", origin)
	assert.Nil(t, err)
	dbBaseMocked.On(
		"Query",
		tx,
		"SELECT id, accountid, amount, date, origin FROM transactions WHERE origin = $1",
		[]interface{}{origin}).Return(rows, nil)
	// When getting a account by ID.
	txs, err := txRepo.GetTransactionsByOrigin(origin)
	// Then the error returned is Nil.
	assert.Nil(t, err)
	assert.NotNil(t, txs)
	assert.Equal(t, 3, len(txs))
	assert.Equal(t, -300.00, txs[2].Amount)
}
