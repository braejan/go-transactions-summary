package postgres_test

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/braejan/go-transactions-summary/internal/domain/user/repository/postgres"
	voPostgres "github.com/braejan/go-transactions-summary/internal/valueobject/postgres"
	mockvoPostgres "github.com/braejan/go-transactions-summary/internal/valueobject/postgres/mock"
	"github.com/braejan/go-transactions-summary/internal/valueobject/user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// TestGetByIDErrorOpeningDatabase tests the GetByID function
// when the database cannot be opened.
func TestGetByIDErrorOpeningDatabase(t *testing.T) {
	dbBase := mockvoPostgres.NewMockBasePostgresDatabase()
	// And a valid user repository.
	userRepo := postgres.NewPostgresUserRepository(dbBase)
	// And a valid user ID.
	ID := int64(1)
	// And a mocked response when calling Open.
	dbBase.On("Open").Return(nil, errors.New("postgres: error opening database"))
	// When GetByID is called.
	_, err := userRepo.GetByID(ID)
	// Then the error returned should be ErrOpeningDatabase.
	assert.NotNil(t, err)
	assert.Equal(t, voPostgres.ErrOpeningDatabase, err)
}

// TestGetByIDErrorBeginningTransaction tests the GetByID function
// when the transaction cannot be started.
func TestGetByIDErrorBeginningTransaction(t *testing.T) {
	dbBase := mockvoPostgres.NewMockBasePostgresDatabase()
	// And a valid user repository.
	userRepo := postgres.NewPostgresUserRepository(dbBase)
	// And a valid user ID.
	ID := int64(1)
	// And a mocked database.
	db, _, _ := sqlmock.New()
	defer db.Close()
	// And a mocked response when calling Open.
	dbBase.On("Open").Return(db, nil)
	// And a mocked response when calling BeginTx.
	dbBase.On("BeginTx", db).Return(nil, errors.New("postgres: error beginning transaction"))
	// And a mocked response when calling Rollback.
	dbBase.On("Rollback", mock.Anything).Return(nil)
	// And a mocked response when calling Close.
	dbBase.On("Close", db).Return(nil)
	// When GetByID is called.
	_, err := userRepo.GetByID(ID)
	// Then the error returned should be ErrBeginningTransaction.
	assert.NotNil(t, err)
	assert.Equal(t, voPostgres.ErrBeginningTransaction, err)
	assert.True(t, dbBase.AssertExpectations(t))
}

// TestGetByIDErrorQuerying tests the GetByID function
// when the query cannot be executed.
func TestGetByIDErrorQuerying(t *testing.T) {
	dbBase := mockvoPostgres.NewMockBasePostgresDatabase()
	// And a valid user repository.
	userRepo := postgres.NewPostgresUserRepository(dbBase)
	// And a valid user ID.
	ID := int64(1)
	// And a mocked database.
	db, _, _ := sqlmock.New()
	defer db.Close()
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
	dbBase.On("Query", tx, "SELECT id, name, email FROM users WHERE id = $1", []interface{}{ID}).Return(nil, errors.New("postgres: error querying"))
	_, err := userRepo.GetByID(ID)
	// Then the error returned should be ErrQuerying.
	assert.NotNil(t, err)
	assert.Equal(t, user.ErrQueryingUserByID, err)
	assert.True(t, dbBase.AssertExpectations(t))
}

// TestGetByIDErrorScanning tests the GetByID function
// when the query cannot be scanned.
func TestGetByIDErrorScanning(t *testing.T) {
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
	dbMocked.ExpectQuery("SELECT (.+) FROM users WHERE id = (.+)").WithArgs(ID).WillReturnRows(expected)
	rows, err := dbBase.Query(tx, "SELECT id, name, email FROM users WHERE id = $1", ID)
	assert.Nil(t, err)
	dbBaseMocked.On("Query", tx, "SELECT id, name, email FROM users WHERE id = $1", []interface{}{ID}).Return(rows, nil)
	// And a valid user repository.
	userRepo := postgres.NewPostgresUserRepository(dbBaseMocked)
	// When GetByID is called.
	_, err = userRepo.GetByID(ID)
	// Then the error returned should be ErrScanningUser.
	assert.NotNil(t, err)
	assert.Equal(t, user.ErrScanningUserByID, err)
}

// TestGetByIDSucess tests the GetByID function
// when the query is successful.
func TestGetByIDSuccess(t *testing.T) {
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
	expected := sqlmock.NewRows([]string{"id", "name", "email"}).AddRow(ID, "John Doe", "john.doe@amazingemail.com")
	dbMocked.ExpectQuery("SELECT (.+) FROM users WHERE id = (.+)").WithArgs(ID).WillReturnRows(expected)
	rows, err := dbBase.Query(tx, "SELECT id, name, email FROM users WHERE id = $1", ID)
	assert.Nil(t, err)
	dbBaseMocked.On("Query", tx, "SELECT id, name, email FROM users WHERE id = $1", []interface{}{ID}).Return(rows, nil)
	// And a valid user repository.
	userRepo := postgres.NewPostgresUserRepository(dbBaseMocked)
	// When GetByID is called.
	user, err := userRepo.GetByID(ID)
	// Then the error returned should be nil.
	assert.Nil(t, err)
	// And the user returned should be the expected one.
	assert.Equal(t, ID, user.ID)
	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, "john.doe@amazingemail.com", user.Email)
}

// TestGetByIDErrEmptyResponse results in an error when the user is not found.
func TestGetByIDErrEmptyResponse(t *testing.T) {
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
	expected := sqlmock.NewRows([]string{"id", "name", "email"})
	dbMocked.ExpectQuery("SELECT (.+) FROM users WHERE id = (.+)").WithArgs(ID).WillReturnRows(expected)
	rows, err := dbBase.Query(tx, "SELECT id, name, email FROM users WHERE id = $1", ID)
	assert.Nil(t, err)
	dbBaseMocked.On("Query", tx, "SELECT id, name, email FROM users WHERE id = $1", []interface{}{ID}).Return(rows, nil)
	// And a valid user repository.
	userRepo := postgres.NewPostgresUserRepository(dbBaseMocked)
	// When GetByID is called.
	_, err = userRepo.GetByID(ID)
	// Then the error returned should be ErrEmpty.
	assert.NotNil(t, err)
	assert.Equal(t, user.ErrUserNotFound, err)
}

// TestGetByEmailErrorOpeningDatabase results in an error when the database cannot be opened.
func TestGetByEmailErrorOpeningDatabase(t *testing.T) {
	dbBaseMocked := mockvoPostgres.NewMockBasePostgresDatabase()
	// And a valid email.
	email := "john.doe@amazingemail.com"
	// And a mocked response when calling Open.
	dbBaseMocked.On("Open").Return(nil, errors.New("error opening database"))
	// And a valid user repository.
	userRepo := postgres.NewPostgresUserRepository(dbBaseMocked)
	// When GetByEmail is called.
	_, err := userRepo.GetByEmail(email)
	// Then the error returned should be ErrOpeningDatabase.
	assert.NotNil(t, err)
	assert.Equal(t, voPostgres.ErrOpeningDatabase, err)
}

// TestGetByEmailErrorBeginningTransaction results in an error when the transaction cannot be started.
func TestGetByEmailErrorBeginningTransaction(t *testing.T) {
	dbBaseMocked := mockvoPostgres.NewMockBasePostgresDatabase()
	// And a valid email.
	email := "john.doe@amazingemail.com"
	// And a mocked database.
	db, _, _ := sqlmock.New()
	defer db.Close()
	// And a mocked response when calling Open.
	dbBaseMocked.On("Open").Return(db, nil)
	// And a mocked response when calling BeginTx.
	dbBaseMocked.On("BeginTx", db).Return(nil, errors.New("error beginning transaction"))
	// And a mocked response when calling Rollback.
	dbBaseMocked.On("Rollback", mock.Anything).Return(nil)
	// And a mocked response when calling Close.
	dbBaseMocked.On("Close", db).Return(nil)
	// And a valid user repository.
	userRepo := postgres.NewPostgresUserRepository(dbBaseMocked)
	// When GetByEmail is called.
	_, err := userRepo.GetByEmail(email)
	// Then the error returned should be ErrBeginningTransaction.
	assert.NotNil(t, err)
	assert.Equal(t, voPostgres.ErrBeginningTransaction, err)
}

// TestGetByEmailErrorQuerying results in an error when the query cannot be executed.
func TestGetByEmailErrorQuerying(t *testing.T) {
	dbBase := mockvoPostgres.NewMockBasePostgresDatabase()
	// And a valid user repository.
	userRepo := postgres.NewPostgresUserRepository(dbBase)
	// And a valid email.
	email := "john.doe@amazingemail.com"
	// And a mocked database.
	db, _, _ := sqlmock.New()
	defer db.Close()
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
	dbBase.On("Query", tx, "SELECT id, name, email FROM users WHERE email = $1", []interface{}{email}).Return(nil, errors.New("postgres: error querying"))
	_, err := userRepo.GetByEmail(email)
	// Then the error returned should be ErrQuerying.
	assert.NotNil(t, err)
	assert.Equal(t, user.ErrQueryingUserByEmail, err)
	assert.True(t, dbBase.AssertExpectations(t))
}

// TestGetByEmailErrorScanning results in an error when the query cannot be scanned.
func TestGetByEmailErrorScanning(t *testing.T) {
	// Given a valid configuration.
	configuration := voPostgres.NewPostgresConfigurationFromEnv()
	dbBase := voPostgres.NewBasePostgresDatabase(configuration)
	dbBaseMocked := mockvoPostgres.NewMockBasePostgresDatabase()
	// And a valid email.
	email := "john.doe@amazingemail.com"
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
	dbMocked.ExpectQuery("SELECT (.+) FROM users WHERE email = (.+)").WithArgs(email).WillReturnRows(expected)
	rows, err := dbBase.Query(tx, "SELECT id, name, email FROM users WHERE email = $1", email)
	assert.Nil(t, err)
	dbBaseMocked.On("Query", tx, "SELECT id, name, email FROM users WHERE email = $1", []interface{}{email}).Return(rows, nil)
	// And a valid user repository.
	userRepo := postgres.NewPostgresUserRepository(dbBaseMocked)
	// When GetByEmail is called.
	_, err = userRepo.GetByEmail(email)
	// Then the error returned should be ErrScanningUser.
	assert.NotNil(t, err)
	assert.Equal(t, user.ErrScanningUserByEmail, err)
}

// TestGetByEmailSuccess results in a user when the query is successful.
func TestGetByEmailSuccess(t *testing.T) {
	// Given a valid configuration.
	configuration := voPostgres.NewPostgresConfigurationFromEnv()
	dbBase := voPostgres.NewBasePostgresDatabase(configuration)
	dbBaseMocked := mockvoPostgres.NewMockBasePostgresDatabase()
	// And a valid email.
	email := "john.doe@amazingemail.com"
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
	expected := sqlmock.NewRows([]string{"id", "name", "email"}).AddRow(1, "John Doe", email)
	dbMocked.ExpectQuery("SELECT (.+) FROM users WHERE email = (.+)").WithArgs(email).WillReturnRows(expected)
	rows, err := dbBase.Query(tx, "SELECT id, name, email FROM users WHERE email = $1", email)
	assert.Nil(t, err)
	dbBaseMocked.On("Query", tx, "SELECT id, name, email FROM users WHERE email = $1", []interface{}{email}).Return(rows, nil)
	// And a valid user repository.
	userRepo := postgres.NewPostgresUserRepository(dbBaseMocked)
	// When GetByEmail is called.
	user, err := userRepo.GetByEmail(email)
	// Then the error returned should be nil.
	assert.Nil(t, err)
	// And the user returned should be valid.
	assert.NotNil(t, user)
	assert.Equal(t, int64(1), user.ID)
	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, email, user.Email)
}

// TestGetByEmailErrEmptyResponse results in an error when the email is empty.
func TestGetByEmailErrEmptyResponse(t *testing.T) {
	// Given a valid configuration.
	configuration := voPostgres.NewPostgresConfigurationFromEnv()
	dbBase := voPostgres.NewBasePostgresDatabase(configuration)
	dbBaseMocked := mockvoPostgres.NewMockBasePostgresDatabase()
	// And a valid email.
	email := "john.deo@amazingemail.com"
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
	expected := sqlmock.NewRows([]string{"id", "name", "email"})
	dbMocked.ExpectQuery("SELECT (.+) FROM users WHERE email = (.+)").WithArgs(email).WillReturnRows(expected)
	rows, err := dbBase.Query(tx, "SELECT id, name, email FROM users WHERE email = $1", email)
	assert.Nil(t, err)
	dbBaseMocked.On("Query", tx, "SELECT id, name, email FROM users WHERE email = $1", []interface{}{email}).Return(rows, nil)
	// And a valid user repository.
	userRepo := postgres.NewPostgresUserRepository(dbBaseMocked)
	// When GetByEmail is called.
	_, err = userRepo.GetByEmail(email)
	// Then the error returned should be ErrEmpty.
	assert.NotNil(t, err)
	assert.Equal(t, user.ErrUserNotFound, err)
}
