package valueobject

import "errors"

var (
	// ErrDBIsNil is the error returned when the database is nil.
	ErrDBIsNil = errors.New("database is nil")
	// ErrUserRepositoryIsNil is the error returned when the user repository is nil.
	ErrUserRepositoryIsNil = errors.New("user repository is nil")
	// ErrUserNotFound is the error returned when a user is not found.
	ErrUserNotFound = errors.New("user not found")
	// ErrUserAlreadyCreated is the error returned when a user is already created.
	ErrUserAlreadyCreated = errors.New("user already created")
)
