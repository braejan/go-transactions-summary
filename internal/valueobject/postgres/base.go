package postgres

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
)

// basePostgresDatabase is the base implementation of PostgresDatabase interface.
type basePostgresDatabase struct {
	postgresConfig *PostgresConfiguration
}

// NewBasePostgresDatabase creates a new instance of PostgresDatabase interface implementation.
func NewBasePostgresDatabase(postgresConfig *PostgresConfiguration) PostgresDatabase {
	return &basePostgresDatabase{
		postgresConfig: postgresConfig,
	}
}

// PostgresDatabase interface implementation.

// Open opens a new database connection.
func (postgresRepo *basePostgresDatabase) Open() (db *sql.DB, err error) {
	db, err = sql.Open("postgres", postgresRepo.postgresConfig.GetDataSourceName())
	return
}

// Close closes a database connection.
func (postgresRepo *basePostgresDatabase) Close(db *sql.DB) (err error) {
	err = db.Close()
	return
}

// Begin begins a transaction.
func (postgresRepo *basePostgresDatabase) BeginTx(db *sql.DB) (tx *sql.Tx, err error) {
	tx, err = db.BeginTx(context.Background(), nil)
	return
}

// Commit commits a transaction.
func (postgresRepo *basePostgresDatabase) Commit(tx *sql.Tx) (err error) {
	err = tx.Commit()
	return
}

// Rollback rollbacks a transaction.
func (postgresRepo *basePostgresDatabase) Rollback(tx *sql.Tx) (err error) {
	if tx == nil {
		return
	}
	err = tx.Rollback()
	return
}

// Exec executes a sql instruction.
func (postgresRepo *basePostgresDatabase) Exec(tx *sql.Tx, dml string, args ...interface{}) (result sql.Result, err error) {
	result, err = tx.ExecContext(context.Background(), dml, args...)
	return
}

// Query executes a sql query.
func (postgresRepo *basePostgresDatabase) Query(tx *sql.Tx, query string, args ...interface{}) (rows *sql.Rows, err error) {
	rows, err = tx.QueryContext(context.Background(), query, args...)
	return
}
