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
	// ErrAccountRepositoryIsNil is returned when an account repository is nil.
	ErrAccountRepositoryIsNil = errors.New("account repository is nil")
	// ErrProcessingAccountID is returned when an error occurs while processing an account ID.
	ErrProcessingAccountID = errors.New("error processing account ID")
	// ErrAccountAlreadyCreated is returned when an account is already created.
	ErrAccountAlreadyCreated = errors.New("account already created")
	// ErrNilAccountUseCases is returned when an account use cases is nil.
	ErrNilAccountUseCases = errors.New("account use cases is nil")
)
