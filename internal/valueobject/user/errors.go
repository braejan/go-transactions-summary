package user

import "errors"

var (
	// ErrScanningUserRow is returned when the user row cannot be scanned.
	ErrScanningUserRow = errors.New("error scanning user row")
	// ErrCreatingUser is returned when the user cannot be created.
	ErrCreatingUser = errors.New("error creating user")
	// ErrUpdatingUser is returned when the user cannot be updated.
	ErrUpdatingUser = errors.New("error updating user")
	// ErrUserRepositoryIsNil is the error returned when the user repository is nil.
	ErrUserRepositoryIsNil = errors.New("user repository is nil")
	// ErrUserNotFound is the error returned when a user is not found.
	ErrUserNotFound = errors.New("user not found")
	// ErrUserAlreadyCreated is the error returned when a user is already created.
	ErrUserAlreadyCreated = errors.New("user already created")
	// ErrQueryingUserByID is the error returned when querying a user by ID.
	ErrQueryingUserByID = errors.New("error querying user by ID")
	// ErrScanningUserByID is the error returned when scanning a user by ID.
	ErrScanningUserByID = errors.New("error scanning user by ID")
)
