package valueobject

import "errors"

var (
	// ErrDBIsNil is the error returned when the database is nil.
	ErrDBIsNil = errors.New("database is nil")
)
