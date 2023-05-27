package postgres_test

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/braejan/go-transactions-summary/internal/domain/account/repository/postgres"
	"github.com/braejan/go-transactions-summary/internal/valueobject/account"
	voPostgres "github.com/braejan/go-transactions-summary/internal/valueobject/postgres"
	mockvoPostgres "github.com/braejan/go-transactions-summary/internal/valueobject/postgres/mock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// TestGetByIDErrorOpeningDatabase tests the GetByID method when an error occurs while opening the database.
func TestGetByIDErrorOpeningDatabase(t *testing.T) {
	dbBase := mockvoPostgres.NewMockBasePostgresDatabase()
	// And a valid account repository.
	accountRepo := postgres.NewPostgresAccountRepository(dbBase)
	// And a valid account ID.
	ID := uuid.New()
	// And a mocked response when calling Open.
	dbBase.On("Open").Return(nil, errors.New("postgres: error opening database"))
	// When GetByID is called.
	_, err := accountRepo.GetByID(ID)
	// Then the error returned should be ErrOpeningDatabase.
	assert.NotNil(t, err)
	assert.Equal(t, voPostgres.ErrOpeningDatabase, err)
}

// TestGetByIDErrorBeginningTransaction tests the GetByID method when an error occurs while beginning a transaction.
func TestGetByIDErrorBeginningTransaction(t *testing.T) {
	dbBase := mockvoPostgres.NewMockBasePostgresDatabase()
	// And a valid account repository.
	accountRepo := postgres.NewPostgresAccountRepository(dbBase)
	// And a valid account ID.
	ID := uuid.New()
	// And a mocked database.
	db, _, _ := sqlmock.New()
	defer db.Close()
	// And a mocked response when calling Open.
	dbBase.On("Open").Return(db, nil)
	// And a mocked response when calling Close.
	dbBase.On("Close", db).Return(nil)
	// And a mocked response when calling BeginTx.
	dbBase.On("BeginTx", db).Return(nil, errors.New("postgres: error beginning transaction"))
	// And a mocked response when calling Rollback.
	dbBase.On("Rollback", mock.Anything).Return(nil)
	// When GetByID is called.
	_, err := accountRepo.GetByID(ID)
	// Then the error returned should be ErrBeginningTransaction.
	assert.NotNil(t, err)
	assert.Equal(t, voPostgres.ErrBeginningTransaction, err)
}

// TestGetByIDErrorQueryingAccountByID tests the GetByID method when an error occurs while querying the account by ID.
func TestGetByIDErrorQueryingAccountByID(t *testing.T) {
	dbBase := mockvoPostgres.NewMockBasePostgresDatabase()
	// And a valid account repository.
	accountRepo := postgres.NewPostgresAccountRepository(dbBase)
	// And a valid account ID.
	ID := uuid.New()
	// And a mocked database.
	db, _, _ := sqlmock.New()
	defer db.Close()
	// And a mocked response when calling Open.
	dbBase.On("Open").Return(db, nil)
	// And a mocked response when calling Close.
	dbBase.On("Close", db).Return(nil)
	// And a mocked response when calling BeginTx.
	tx, _ := db.BeginTx(context.Background(), nil)
	dbBase.On("BeginTx", db).Return(tx, nil)
	// And a mocked response when calling Rollback.
	dbBase.On("Rollback", mock.Anything).Return(nil)
	// And a mocked response when calling Query.
	dbBase.On("Query", tx, "SELECT id, balance, userid, active FROM accounts WHERE id = $1", []interface{}{ID}).Return(nil, errors.New("postgres: error querying account by ID"))
	// When GetByID is called.
	_, err := accountRepo.GetByID(ID)
	// Then the error returned should be ErrQueryingAccountByID.
	assert.NotNil(t, err)
	assert.Equal(t, account.ErrQueryingAccountByID, err)
}

// TestGetByIDErrorScanningAccountByID tests the GetByID method when an error occurs while scanning the account by ID.
func TestGetByIDErrorScanningAccountByID(t *testing.T) {
	// Given a valid configuration.
	configuration := voPostgres.NewPostgresConfigurationFromEnv()
	dbBase := voPostgres.NewBasePostgresDatabase(configuration)
	dbBaseMocked := mockvoPostgres.NewMockBasePostgresDatabase()
	// And a valid user ID.
	ID := uuid.New()
	// And a mocked database.
	db, dbMocked, _ := sqlmock.New()
	defer db.Close()
	// And a mocked transaction
	dbMocked.ExpectBegin()
	// And a mocked response when calling Open.
	dbBaseMocked.On("Open").Return(db, nil)
	// And a mocked response when calling BeginTx.
	tx, _ := db.BeginTx(context.Background(), nil)
	dbBaseMocked.On("BeginTx", db).Return(tx, nil)
	// And a mocked response when calling Rollback.
	dbBaseMocked.On("Rollback", mock.Anything).Return(nil)
	// And a mocked response when calling Close.
	dbBaseMocked.On("Close", db).Return(nil)
	// And a mocked response when calling Query.
	expected := sqlmock.NewRows([]string{"column1", "column2", "column3"}).AddRow(true, false, false)
	dbMocked.ExpectQuery("SELECT (.+) FROM accounts WHERE id = (.+)").WithArgs(ID).WillReturnRows(expected)
	rows, err := dbBase.Query(tx, "SELECT id, balance, userid, active FROM accounts WHERE id = $1", ID)
	assert.Nil(t, err)
	dbBaseMocked.On("Query", tx, "SELECT id, balance, userid, active FROM accounts WHERE id = $1", []interface{}{ID}).Return(rows, nil)
	// And a valid user repository.
	userRepo := postgres.NewPostgresAccountRepository(dbBaseMocked)
	// When GetByID is called.
	_, err = userRepo.GetByID(ID)
	// Then the error returned should be ErrScanningUser.
	assert.NotNil(t, err)
	assert.Equal(t, account.ErrScanningAccountByID, err)
}

// TestGetByIDSuccess tests the GetByID method when it succeeds.
func TestGetByIDSuccess(t *testing.T) {
	// Given a valid configuration.
	configuration := voPostgres.NewPostgresConfigurationFromEnv()
	dbBase := voPostgres.NewBasePostgresDatabase(configuration)
	dbBaseMocked := mockvoPostgres.NewMockBasePostgresDatabase()
	// And a valid account ID.
	ID := uuid.New()
	// And a mocked database.
	db, dbMocked, _ := sqlmock.New()
	defer db.Close()
	// And a mocked transaction
	dbMocked.ExpectBegin()
	// And a mocked response when calling Open.
	dbBaseMocked.On("Open").Return(db, nil)
	// And a mocked response when calling BeginTx.
	tx, _ := db.BeginTx(context.Background(), nil)
	dbBaseMocked.On("BeginTx", db).Return(tx, nil)
	// And a mocked response when calling Rollback.
	dbBaseMocked.On("Rollback", mock.Anything).Return(nil)
	// And a mocked response when calling Close.
	dbBaseMocked.On("Close", db).Return(nil)
	// And a mocked response when calling Query.
	expected := sqlmock.NewRows([]string{"id", "balance", "userid", "active"}).AddRow(ID, float64(1000), int64(1), true)
	dbMocked.ExpectQuery("SELECT (.+) FROM accounts WHERE id = (.+)").WithArgs(ID).WillReturnRows(expected)
	rows, err := dbBase.Query(tx, "SELECT id, balance, userid, active FROM accounts WHERE id = $1", ID)
	assert.Nil(t, err)
	dbBaseMocked.On("Query", tx, "SELECT id, balance, userid, active FROM accounts WHERE id = $1", []interface{}{ID}).Return(rows, nil)
	// And a valid user repository.
	accountRepo := postgres.NewPostgresAccountRepository(dbBaseMocked)
	// When GetByID is called.
	account, err := accountRepo.GetByID(ID)
	// Then the error returned should be nil.
	assert.Nil(t, err)
	// And the user returned should be the expected one.
	assert.Equal(t, ID, account.ID)
	assert.Equal(t, float64(1000), account.Balance)
	assert.Equal(t, int64(1), account.UserID)
	assert.Equal(t, true, account.Active)
}

// TestGetByIDErrEmptyResponse tests the GetByID method when the response is empty.
func TestGetByIDErrEmptyResponse(t *testing.T) {
	// Given a valid configuration.
	configuration := voPostgres.NewPostgresConfigurationFromEnv()
	dbBase := voPostgres.NewBasePostgresDatabase(configuration)
	dbBaseMocked := mockvoPostgres.NewMockBasePostgresDatabase()
	// And a valid account ID.
	ID := uuid.New()
	// And a mocked database.
	db, dbMocked, _ := sqlmock.New()
	defer db.Close()
	// And a mocked transaction
	dbMocked.ExpectBegin()
	// And a mocked response when calling Open.
	dbBaseMocked.On("Open").Return(db, nil)
	// And a mocked response when calling BeginTx.
	tx, _ := db.BeginTx(context.Background(), nil)
	dbBaseMocked.On("BeginTx", db).Return(tx, nil)
	// And a mocked response when calling Rollback.
	dbBaseMocked.On("Rollback", mock.Anything).Return(nil)
	// And a mocked response when calling Close.
	dbBaseMocked.On("Close", db).Return(nil)
	// And a mocked response when calling Query.
	expected := sqlmock.NewRows([]string{"id", "balance", "userid", "active"})
	dbMocked.ExpectQuery("SELECT (.+) FROM accounts WHERE id = (.+)").WithArgs(ID).WillReturnRows(expected)
	rows, err := dbBase.Query(tx, "SELECT id, balance, userid, active FROM accounts WHERE id = $1", ID)
	assert.Nil(t, err)
	dbBaseMocked.On("Query", tx, "SELECT id, balance, userid, active FROM accounts WHERE id = $1", []interface{}{ID}).Return(rows, nil)
	// And a valid user repository.
	accountRepo := postgres.NewPostgresAccountRepository(dbBaseMocked)
	// When GetByID is called.
	acc, err := accountRepo.GetByID(ID)
	// Then the error returned should be ErrAccountNotFound.
	assert.NotNil(t, err)
	assert.Equal(t, account.ErrAccountNotFound, err)
	// And the user returned should be the expected one.
	assert.Equal(t, uuid.Nil, acc.ID)
	assert.Equal(t, float64(0), acc.Balance)
	assert.Equal(t, int64(0), acc.UserID)
	assert.Equal(t, false, acc.Active)
}

// TestGetByUserIDErrorOpeningDatabase tests the GetByUserID method when the database cannot be opened.
func TestGetByUserIDErrorOpeningDatabase(t *testing.T) {
	dbBase := mockvoPostgres.NewMockBasePostgresDatabase()
	// And a valid account repository.
	accountRepo := postgres.NewPostgresAccountRepository(dbBase)
	// And a valid account ID.
	ID := int64(1)
	// And a mocked response when calling Open.
	dbBase.On("Open").Return(nil, errors.New("postgres: error opening database"))
	// When GetByUserID is called.
	_, err := accountRepo.GetByUserID(ID)
	// Then the error returned should be ErrOpeningDatabase.
	assert.NotNil(t, err)
	assert.Equal(t, voPostgres.ErrOpeningDatabase, err)
}

// TestGetByUserIDErrorBeginningTransaction tests the GetByUserID method when the transaction cannot be started.
func TestGetByUserIDErrorBeginningTransaction(t *testing.T) {
	dbBase := mockvoPostgres.NewMockBasePostgresDatabase()
	// And a valid account repository.
	accountRepo := postgres.NewPostgresAccountRepository(dbBase)
	// And a valid account ID.
	ID := int64(1)
	// And a mocked database.
	db, _, _ := sqlmock.New()
	defer db.Close()
	// And a mocked response when calling Open.
	dbBase.On("Open").Return(db, nil)
	// And a mocked response when calling Close.
	dbBase.On("Close", db).Return(nil)
	// And a mocked response when calling BeginTx.
	dbBase.On("BeginTx", db).Return(nil, errors.New("postgres: error beginning transaction"))
	// And a mocked response when calling Rollback.
	dbBase.On("Rollback", mock.Anything).Return(nil)
	// When GetByUserID is called.
	_, err := accountRepo.GetByUserID(ID)
	// Then the error returned should be ErrBeginningTransaction.
	assert.NotNil(t, err)
	assert.Equal(t, voPostgres.ErrBeginningTransaction, err)
}

// TestGetByUserIDErrorQueryingAccountByID tests the GetByUserID method when the account cannot be queried.
func TestGetByUserIDErrorQueryingAccountByID(t *testing.T) {
	dbBase := mockvoPostgres.NewMockBasePostgresDatabase()
	// And a valid account repository.
	accountRepo := postgres.NewPostgresAccountRepository(dbBase)
	// And a valid account ID.
	ID := int64(1)
	// And a mocked database.
	db, dbMocked, _ := sqlmock.New()
	defer db.Close()
	// And a mocked transaction
	dbMocked.ExpectBegin()
	// And a mocked response when calling Open.
	dbBase.On("Open").Return(db, nil)
	// And a mocked response when calling BeginTx.
	tx, _ := db.BeginTx(context.Background(), nil)
	dbBase.On("BeginTx", db).Return(tx, nil)
	// And a mocked response when calling Rollback.
	dbBase.On("Rollback", mock.Anything).Return(nil)
	// And a mocked response when calling Close.
	dbBase.On("Close", db).Return(nil)
	// And a mocked response when calling Query.
	dbBase.On("Query", tx, "SELECT id, balance, userid, active FROM accounts WHERE userid = $1", []interface{}{ID}).Return(nil, errors.New("postgres: error querying account by id"))
	// When GetByUserID is called.
	_, err := accountRepo.GetByUserID(ID)
	// Then the error returned should be ErrQueryingAccountByID.
	assert.NotNil(t, err)
	assert.Equal(t, account.ErrQueryingAccountByUserID, err)
}

// TestGetByUserIDErrorScanningAccountByID tests the GetByUserID method when the account cannot be scanned.
func TestGetByUserIDErrorScanningAccountByID(t *testing.T) {
	// Given a valid configuration.
	configuration := voPostgres.NewPostgresConfigurationFromEnv()
	dbBase := voPostgres.NewBasePostgresDatabase(configuration)
	dbBaseMocked := mockvoPostgres.NewMockBasePostgresDatabase()
	// And a valid user ID.
	ID := int64(1)
	// And a mocked database.
	db, dbMocked, _ := sqlmock.New()
	defer db.Close()
	// And a mocked transaction
	dbMocked.ExpectBegin()
	// And a mocked response when calling Open.
	dbBaseMocked.On("Open").Return(db, nil)
	// And a mocked response when calling BeginTx.
	tx, _ := db.BeginTx(context.Background(), nil)
	dbBaseMocked.On("BeginTx", db).Return(tx, nil)
	// And a mocked response when calling Rollback.
	dbBaseMocked.On("Rollback", mock.Anything).Return(nil)
	// And a mocked response when calling Close.
	dbBaseMocked.On("Close", db).Return(nil)
	// And a mocked response when calling Query.
	expected := sqlmock.NewRows([]string{"column1", "column2", "column3"}).AddRow(true, false, false)
	dbMocked.ExpectQuery("SELECT (.+) FROM accounts WHERE userid = (.+)").WithArgs(ID).WillReturnRows(expected)
	rows, err := dbBase.Query(tx, "SELECT id, balance, userid, active FROM accounts WHERE userid = $1", ID)
	assert.Nil(t, err)
	dbBaseMocked.On("Query", tx, "SELECT id, balance, userid, active FROM accounts WHERE userid = $1", []interface{}{ID}).Return(rows, nil)
	// And a valid user repository.
	userRepo := postgres.NewPostgresAccountRepository(dbBaseMocked)
	// When GetByUserID is called.
	_, err = userRepo.GetByUserID(ID)
	// Then the error returned should be ErrScanningUser.
	assert.NotNil(t, err)
	assert.Equal(t, account.ErrScanningAccountByUserID, err)
}

// TestGetByUserIDSuccess tests the GetByUserID method when it is successful.
func TestGetByUserIDSuccess(t *testing.T) {
	// Given a valid configuration.
	configuration := voPostgres.NewPostgresConfigurationFromEnv()
	dbBase := voPostgres.NewBasePostgresDatabase(configuration)
	dbBaseMocked := mockvoPostgres.NewMockBasePostgresDatabase()
	// And a valid account ID.
	ID := uuid.New()
	userID := int64(1)
	// And a mocked database.
	db, dbMocked, _ := sqlmock.New()
	defer db.Close()
	// And a mocked transaction
	dbMocked.ExpectBegin()
	// And a mocked response when calling Open.
	dbBaseMocked.On("Open").Return(db, nil)
	// And a mocked response when calling BeginTx.
	tx, _ := db.BeginTx(context.Background(), nil)
	dbBaseMocked.On("BeginTx", db).Return(tx, nil)
	// And a mocked response when calling Rollback.
	dbBaseMocked.On("Rollback", mock.Anything).Return(nil)
	// And a mocked response when calling Close.
	dbBaseMocked.On("Close", db).Return(nil)
	// And a mocked response when calling Query.
	expected := sqlmock.NewRows([]string{"id", "balance", "userid", "active"}).AddRow(ID, float64(1000), int64(1), true)
	dbMocked.ExpectQuery("SELECT (.+) FROM accounts WHERE userid = (.+)").WithArgs(userID).WillReturnRows(expected)
	rows, err := dbBase.Query(tx, "SELECT id, balance, userid, active FROM accounts WHERE userid = $1", userID)
	assert.Nil(t, err)
	dbBaseMocked.On("Query", tx, "SELECT id, balance, userid, active FROM accounts WHERE userid = $1", []interface{}{userID}).Return(rows, nil)
	// And a valid user repository.
	accountRepo := postgres.NewPostgresAccountRepository(dbBaseMocked)
	// When GetByUserID is called.
	account, err := accountRepo.GetByUserID(userID)
	// Then the error returned should be nil.
	assert.Nil(t, err)
	// And the user returned should be the expected one.
	assert.Equal(t, ID, account.ID)
	assert.Equal(t, float64(1000), account.Balance)
	assert.Equal(t, int64(1), account.UserID)
	assert.Equal(t, true, account.Active)
}

// TestGetByUserIDErrEmptyResponse tests the GetByUserID method when the response is empty.
func TestGetByUserIDErrEmptyResponse(t *testing.T) {
	// Given a valid configuration.
	configuration := voPostgres.NewPostgresConfigurationFromEnv()
	dbBase := voPostgres.NewBasePostgresDatabase(configuration)
	dbBaseMocked := mockvoPostgres.NewMockBasePostgresDatabase()
	// And a valid account ID.
	userID := int64(1)
	// And a mocked database.
	db, dbMocked, _ := sqlmock.New()
	defer db.Close()
	// And a mocked transaction
	dbMocked.ExpectBegin()
	// And a mocked response when calling Open.
	dbBaseMocked.On("Open").Return(db, nil)
	// And a mocked response when calling BeginTx.
	tx, _ := db.BeginTx(context.Background(), nil)
	dbBaseMocked.On("BeginTx", db).Return(tx, nil)
	// And a mocked response when calling Rollback.
	dbBaseMocked.On("Rollback", mock.Anything).Return(nil)
	// And a mocked response when calling Close.
	dbBaseMocked.On("Close", db).Return(nil)
	// And a mocked response when calling Query.
	expected := sqlmock.NewRows([]string{"id", "balance", "userid", "active"})
	dbMocked.ExpectQuery("SELECT (.+) FROM accounts WHERE userid = (.+)").WithArgs(userID).WillReturnRows(expected)
	rows, err := dbBase.Query(tx, "SELECT id, balance, userid, active FROM accounts WHERE userid = $1", userID)
	assert.Nil(t, err)
	dbBaseMocked.On("Query", tx, "SELECT id, balance, userid, active FROM accounts WHERE userid = $1", []interface{}{userID}).Return(rows, nil)
	// And a valid user repository.
	accountRepo := postgres.NewPostgresAccountRepository(dbBaseMocked)
	// When GetByUserID is called.
	_, err = accountRepo.GetByUserID(userID)
	// Then the error returned should be ErrAccountNotFound.
	assert.NotNil(t, err)
	assert.Equal(t, account.ErrAccountNotFound, err)
}
