package transaction

import "errors"

var (
	// ErrTransactionAmountIsZero is the error returned when the transaction amount is zero.
	ErrTransactionAmountIsZero = errors.New("transaction amount is zero")
	// ErrTransactionOriginIsEmpty is the error returned when the transaction origin is empty.
	ErrTransactionOriginIsEmpty = errors.New("transaction origin is empty")
	// ErrQueryingTransactionByID is the error returned when querying a transaction by ID.
	ErrQueryingTransactionByID = errors.New("error querying transaction by ID")
	// ErrScanningTransactionByID is the error returned when scanning a transaction by ID.
	ErrScanningTransactionByID = errors.New("error scanning transaction by ID")
	// ErrQueryingTransactionsByAccountID is the error returned when querying transactions by account ID.
	ErrQueryingTransactionsByAccountID = errors.New("error querying transactions by account ID")
	// ErrScanningTransactionsByAccountID is the error returned when scanning transactions by account ID.
	ErrScanningTransactionsByAccountID = errors.New("error scanning transactions by account ID")
	// ErrTransactionDateIsInvalid is the error returned when the transaction date is invalid.
	ErrTransactionDateIsInvalid = errors.New("transaction date is invalid")
	// ErrQueryingCreditsByAccountID is the error returned when querying credits by account ID.
	ErrQueryingCreditsByAccountID = errors.New("error querying credits by account ID")
	// ErrScanningCreditsByAccountID is the error returned when scanning credits by account ID.
	ErrScanningCreditsByAccountID = errors.New("error scanning credits by account ID")
	// ErrQueryingDebitsByAccountID is the error returned when querying debits by account ID.
	ErrQueryingDebitsByAccountID = errors.New("error querying debits by account ID")
	// ErrCreatingTransaction is the error returned when creating a transaction.
	ErrCreatingTransaction = errors.New("error creating transaction")
	// ErrNilTransaction is the error returned when the transaction is nil.
	ErrNilTransaction = errors.New("transaction is nil")
	// ErrEmptyOrigin is the error returned when the origin is empty.
	ErrEmptyOrigin = errors.New("origin is empty")
	// ErrQueryingTransactionsByOrigin is the error returned when querying transactions by origin.
	ErrQueryingTransactionsByOrigin = errors.New("error querying transactions by origin")
)
