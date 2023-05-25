package postgres

import "errors"

var (
	// ErrOpeningDatabase is returned when the database cannot be opened.
	ErrOpeningDatabase = errors.New("error opening database")
	// ErrBeginningTransaction is returned when the transaction cannot be started.
	ErrBeginningTransaction = errors.New("error beginning transaction")
	// ErrNilConfiguration is returned when the configuration is nil.
	ErrNilConfiguration = errors.New("error nil configuration")
)
