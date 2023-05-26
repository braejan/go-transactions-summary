package postgres_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/braejan/go-transactions-summary/internal/valueobject/postgres"
	"github.com/stretchr/testify/assert"
)

// TestNewBasePostgresDatabase tests the NewBasePostgresDatabase function.
func TestNewBasePostgresDatabase(t *testing.T) {
	//When call NewBasePostgresDatabase
	postgresRepo := postgres.NewBasePostgresDatabase()
	// Then return a new instance of PostgresDatabase interface implementation.
	assert.NotNil(t, postgresRepo)
}

// TestCloseFail tests the Close function fails.
func TestCloseFail(t *testing.T) {
	// Given a mock database
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	// And a mock database that fails
	mock.ExpectClose().WillReturnError(assert.AnError)
	// When call Close
	err = postgres.NewBasePostgresDatabase().Close(db)
	// Then return an error
	assert.Error(t, err)
}

// TestCloseSuccess tests the Close function succeeds.
func TestCloseSuccess(t *testing.T) {
	// Given a mock database
	db, dbMock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	// And a mock database that succeeds
	dbMock.ExpectClose().WillReturnError(nil)
	// When call Close
	err = postgres.NewBasePostgresDatabase().Close(db)
	// Then return no error
	assert.NoError(t, err)
}

// TestBeginTxFail tests the BeginTx function fails.
func TestBeginTxFail(t *testing.T) {
	// Given a mock database
	db, dbMock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	// And a mock database that fails
	dbMock.ExpectBegin().WillReturnError(assert.AnError)
	// When call BeginTx
	tx, err := postgres.NewBasePostgresDatabase().BeginTx(db)
	// Then return an error
	assert.Error(t, err)
	assert.Nil(t, tx)
}

// TestBeginTxSuccess tests the BeginTx function succeeds.
func TestBeginTxSuccess(t *testing.T) {
	// Given a mock database
	db, dbMock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	// And a mock database that succeeds
	dbMock.ExpectBegin().WillReturnError(nil)
	// When call BeginTx
	tx, err := postgres.NewBasePostgresDatabase().BeginTx(db)
	// Then return no error
	assert.NoError(t, err)
	assert.NotNil(t, tx)
}

// TestCommitFail tests the Commit function fails.
func TestCommitFail(t *testing.T) {
	// Given a mock database
	db, dbMock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	// And a mock database that fails
	dbMock.ExpectBegin().WillReturnError(nil)
	dbMock.ExpectCommit().WillReturnError(assert.AnError)
	// And a mocked dbTx
	tx, err := db.Begin()
	assert.NoError(t, err)
	// When call Commit
	err = postgres.NewBasePostgresDatabase().Commit(tx)
	// Then return an error
	assert.Error(t, err)
}

// TestCommitSuccess tests the Commit function succeeds.
func TestCommitSuccess(t *testing.T) {
	// Given a mock database
	db, dbMock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	// And a mock database that succeeds
	dbMock.ExpectBegin().WillReturnError(nil)
	dbMock.ExpectCommit().WillReturnError(nil)
	// And a mocked dbTx
	tx, err := db.Begin()
	assert.NoError(t, err)
	// When call Commit
	err = postgres.NewBasePostgresDatabase().Commit(tx)
	// Then return no error
	assert.NoError(t, err)
}

// TestRollbackFail tests the Rollback function fails.
func TestRollbackFail(t *testing.T) {
	// Given a mock database
	db, dbMock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	// And a mock database that fails
	dbMock.ExpectBegin().WillReturnError(nil)
	dbMock.ExpectRollback().WillReturnError(assert.AnError)
	// And a mocked dbTx
	tx, err := db.Begin()
	assert.NoError(t, err)
	// When call Rollback
	err = postgres.NewBasePostgresDatabase().Rollback(tx)
	// Then return an error
	assert.Error(t, err)
}

// TestRollbackSuccess tests the Rollback function succeeds.
func TestRollbackSuccess(t *testing.T) {
	// Given a mock database
	db, dbMock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	// And a mock database that succeeds
	dbMock.ExpectBegin().WillReturnError(nil)
	dbMock.ExpectRollback().WillReturnError(nil)
	// And a mocked dbTx
	tx, err := db.Begin()
	assert.NoError(t, err)
	// When call Rollback
	err = postgres.NewBasePostgresDatabase().Rollback(tx)
	// Then return no error
	assert.NoError(t, err)
}

// TestExecFail tests the Exec function fails.
func TestExecFail(t *testing.T) {
	// Given a mock database
	db, dbMock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	// And a mock database that fails
	dbMock.ExpectBegin().WillReturnError(nil)
	dbMock.ExpectExec("test").WillReturnError(assert.AnError)
	// And a mocked dbTx
	tx, err := db.Begin()
	assert.NoError(t, err)
	// When call Exec
	_, err = postgres.NewBasePostgresDatabase().Exec(tx, "test")
	// Then return an error
	assert.Error(t, err)
}

// TestExecSuccess tests the Exec function succeeds.
func TestExecSuccess(t *testing.T) {
	// Given a mock database
	db, dbMock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	// And a mock database that succeeds
	dbMock.ExpectBegin().WillReturnError(nil)
	dbMock.ExpectExec("test").WillReturnResult(sqlmock.NewResult(1, 1))
	// And a mocked dbTx
	tx, err := db.Begin()
	assert.NoError(t, err)
	// When call Exec
	_, err = postgres.NewBasePostgresDatabase().Exec(tx, "test")
	// Then return no error
	assert.NoError(t, err)
}

// TestQueryFail tests the Query function fails.
func TestQueryFail(t *testing.T) {
	// Given a mock database
	db, dbMock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	// And a mock database that fails
	dbMock.ExpectBegin().WillReturnError(nil)
	dbMock.ExpectQuery("test").WillReturnError(assert.AnError)
	// And a mocked dbTx
	tx, err := db.Begin()
	assert.NoError(t, err)
	// When call Query
	_, err = postgres.NewBasePostgresDatabase().Query(tx, "test")
	// Then return an error
	assert.Error(t, err)
}

// TestQuerySuccess tests the Query function succeeds.
func TestQuerySuccess(t *testing.T) {
	// Given a mock database
	db, dbMock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	// And a mock database that succeeds
	dbMock.ExpectBegin().WillReturnError(nil)
	dbMock.ExpectQuery("test").WillReturnRows(sqlmock.NewRows([]string{"test"}).AddRow("test"))
	// And a mocked dbTx
	tx, err := db.Begin()
	assert.NoError(t, err)
	// When call Query
	_, err = postgres.NewBasePostgresDatabase().Query(tx, "test")
	// Then return no error
	assert.NoError(t, err)
}
