package postgres

import "errors"

var (
	// ErrOpeningDatabase is returned when the database cannot be opened.
	ErrOpeningDatabase = errors.New("error opening database")
	// ErrBeginningTransaction is returned when the transaction cannot be started.
	ErrBeginningTransaction = errors.New("error beginning transaction")
	// ErrNilConfiguration is returned when the configuration is nil.
	ErrNilConfiguration = errors.New("error nil configuration")
	// ErrDBIsNil is the error returned when the database is nil.
	ErrDBIsNil = errors.New("database is nil")
	// ErrExec is the error returned when the query cannot be executed.
	ErrExec = errors.New("error executing dml")
	// ErrCommittingTransaction is the error returned when the transaction cannot be committed.
	ErrCommittingTransaction = errors.New("error committing transaction")
	// ErrRollingBackTransaction is the error returned when the transaction cannot be rolled back.
	ErrRollingBackTransaction = errors.New("error rolling back transaction")
	// ErrQueryingDatabase is the error returned when the database cannot be queried.
	ErrQueryingDatabase = errors.New("error querying database")
)
