package account

import "errors"

var (
	// ErrNilAccount is returned when an account is nil.
	ErrNilAccount = errors.New("account is nil")
	// ErrQueryingAccountByID is returned when an error occurs while querying an account by its ID.
	ErrQueryingAccountByID = errors.New("error querying account by ID")
	// ErrScanningAccountByID is returned when an error occurs while scanning an account by its ID.
	ErrScanningAccountByID = errors.New("error scanning account by ID")
	// ErrAccountNotFound is returned when an account is not found.
	ErrAccountNotFound = errors.New("account not found")
	// ErrQueryingAccountByUserID is returned when an error occurs while querying an account by its user ID.
	ErrQueryingAccountByUserID = errors.New("error querying account by user ID")
	// ErrScanningAccountByUserID is returned when an error occurs while scanning an account by its user ID.
	ErrScanningAccountByUserID = errors.New("error scanning account by user ID")
	// ErrCreatingAccount is returned when an error occurs while creating an account.
	ErrCreatingAccount = errors.New("error creating account")
	// ErrUpdatingAccount is returned when an error occurs while updating an account.
	ErrUpdatingAccount = errors.New("error updating account")
	// ErrScanningAccountByID is returned when an error occurs while scanning an account by its ID.
	ErrScanningAccount = errors.New("error scanning account")
)
