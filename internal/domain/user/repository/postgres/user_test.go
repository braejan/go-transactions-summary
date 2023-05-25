package postgres_test

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/braejan/go-transactions-summary/internal/domain/user/repository/postgres"
	voPostgres "github.com/braejan/go-transactions-summary/internal/valueobject/postgres"
	"github.com/braejan/go-transactions-summary/internal/valueobject/postgres/mock"
	"github.com/braejan/go-transactions-summary/internal/valueobject/user"

	"github.com/stretchr/testify/assert"
)

func getTestRows() (rows *sql.Rows) {
	rows = new(sql.Rows)
	return
}

// TestNewPostgresUserRepositoryConfigNil tests the NewPostgresUserRepository function
// when the configuration is nil.
func TestNewPostgresUserRepositoryConfigNil(t *testing.T) {
	// Given a nil configuration.
	// When NewPostgresUserRepository is called.
	_, err := postgres.NewPostgresUserRepository(nil, nil)
	// Then the error returned should be ErrNilConfiguration.
	assert.NotNil(t, err)
	assert.Equal(t, voPostgres.ErrNilConfiguration, err)
}

// TestNewPostgresUserRepositoryWithConfig tests the NewPostgresUserRepository function
// when the configuration is not nil.
func TestNewPostgresUserRepositoryWithConfig(t *testing.T) {
	// Given a valid configuration.
	configuration := voPostgres.GetPostgresConfigurationFromEnv()
	dbBase := mock.NewMockBasePostgresDatabase()
	// When NewPostgresUserRepository is called.
	_, err := postgres.NewPostgresUserRepository(configuration, dbBase)
	// Then the error returned should be nil.
	assert.Nil(t, err)
}

// TestGetByIDErrorOpeningDatabase tests the GetByID function
// when the database cannot be opened.
func TestGetByIDErrorOpeningDatabase(t *testing.T) {
	// Given a valid configuration.
	configuration := voPostgres.GetPostgresConfigurationFromEnv()
	dbBase := mock.NewMockBasePostgresDatabase()
	// And a valid user repository.
	userRepo, _ := postgres.NewPostgresUserRepository(configuration, dbBase)
	// And a valid user ID.
	ID := int64(1)
	// And a mocked response when calling Open.
	dbBase.On("Open", configuration.GetDataSourceName()).Return(nil, errors.New("postgres: error opening database"))
	// When GetByID is called.
	_, err := userRepo.GetByID(ID)
	// Then the error returned should be ErrOpeningDatabase.
	assert.NotNil(t, err)
	assert.Equal(t, voPostgres.ErrOpeningDatabase, err)
}

// TestGetByIDErrorBeginningTransaction tests the GetByID function
// when the transaction cannot be started.
func TestGetByIDErrorBeginningTransaction(t *testing.T) {
	// Given a valid configuration.
	configuration := voPostgres.GetPostgresConfigurationFromEnv()
	dbBase := mock.NewMockBasePostgresDatabase()
	// And a valid user repository.
	userRepo, _ := postgres.NewPostgresUserRepository(configuration, dbBase)
	// And a valid user ID.
	ID := int64(1)
	// And a mocked database.
	db, _, _ := sqlmock.New()
	defer db.Close()
	// And a mocked response when calling Open.
	dbBase.On("Open", configuration.GetDataSourceName()).Return(db, nil)
	// And a mocked response when calling BeginTx.
	dbBase.On("BeginTx", db).Return(nil, errors.New("postgres: error beginning transaction"))
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
	// Given a valid configuration.
	configuration := voPostgres.GetPostgresConfigurationFromEnv()
	dbBase := mock.NewMockBasePostgresDatabase()
	// And a valid user repository.
	userRepo, _ := postgres.NewPostgresUserRepository(configuration, dbBase)
	// And a valid user ID.
	ID := int64(1)
	// And a mocked database.
	db, _, _ := sqlmock.New()
	defer db.Close()
	// And a mocked response when calling Open.
	dbBase.On("Open", configuration.GetDataSourceName()).Return(db, nil)
	// And a mocked response when calling BeginTx.
	tx, _ := db.BeginTx(context.Background(), nil)
	dbBase.On("BeginTx", db).Return(tx, nil)
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
	configuration := voPostgres.GetPostgresConfigurationFromEnv()
	dbBase := voPostgres.NewBasePostgresDatabase()
	dbBaseMocked := mock.NewMockBasePostgresDatabase()
	// And a valid user ID.
	ID := int64(1)
	// And a mocked database.
	db, dbMocked, _ := sqlmock.New()
	defer db.Close()
	// And a mocked transaction
	dbMocked.ExpectBegin()
	// And a mocked response when calling Open.
	dbBaseMocked.On("Open", configuration.GetDataSourceName()).Return(db, nil)
	// And a mocked response when calling BeginTx.
	tx, _ := db.BeginTx(context.Background(), nil)
	dbBaseMocked.On("BeginTx", db).Return(tx, nil)
	// And a mocked response when calling Close.
	dbBaseMocked.On("Close", db).Return(nil)
	// And a mocked response when calling Query.
	expected := sqlmock.NewRows([]string{"column1", "column2", "column3"}).AddRow(true, false, false)
	dbMocked.ExpectQuery("SELECT (.+) FROM users WHERE id = (.+)").WithArgs(ID).WillReturnRows(expected)
	rows, err := dbBase.Query(tx, "SELECT id, name, email FROM users WHERE id = $1", ID)
	assert.Nil(t, err)
	dbBaseMocked.On("Query", tx, "SELECT id, name, email FROM users WHERE id = $1", []interface{}{ID}).Return(rows, nil)
	// And a valid user repository.
	userRepo, _ := postgres.NewPostgresUserRepository(configuration, dbBaseMocked)
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
	configuration := voPostgres.GetPostgresConfigurationFromEnv()
	dbBase := voPostgres.NewBasePostgresDatabase()
	dbBaseMocked := mock.NewMockBasePostgresDatabase()
	// And a valid user ID.
	ID := int64(1)
	// And a mocked database.
	db, dbMocked, _ := sqlmock.New()
	defer db.Close()
	// And a mocked transaction
	dbMocked.ExpectBegin()
	// And a mocked response when calling Open.
	dbBaseMocked.On("Open", configuration.GetDataSourceName()).Return(db, nil)
	// And a mocked response when calling BeginTx.
	tx, _ := db.BeginTx(context.Background(), nil)
	dbBaseMocked.On("BeginTx", db).Return(tx, nil)
	// And a mocked response when calling Close.
	dbBaseMocked.On("Close", db).Return(nil)
	// And a mocked response when calling Query.
	expected := sqlmock.NewRows([]string{"id", "name", "email"}).AddRow(ID, "John Doe", "john.doe@amazingemail.com")
	dbMocked.ExpectQuery("SELECT (.+) FROM users WHERE id = (.+)").WithArgs(ID).WillReturnRows(expected)
	rows, err := dbBase.Query(tx, "SELECT id, name, email FROM users WHERE id = $1", ID)
	assert.Nil(t, err)
	dbBaseMocked.On("Query", tx, "SELECT id, name, email FROM users WHERE id = $1", []interface{}{ID}).Return(rows, nil)
	// And a valid user repository.
	userRepo, _ := postgres.NewPostgresUserRepository(configuration, dbBaseMocked)
	// When GetByID is called.
	user, err := userRepo.GetByID(ID)
	// Then the error returned should be nil.
	assert.Nil(t, err)
	// And the user returned should be the expected one.
	assert.Equal(t, ID, user.ID)
	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, "john.doe@amazingemail.com", user.Email)
}
