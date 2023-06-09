package entity

import (
	"time"

	"github.com/braejan/go-transactions-summary/internal/valueobject/transaction"
	"github.com/google/uuid"
)

// Transaction struct defines the transaction entity.
type Transaction struct {
	ID uuid.UUID
	// AccountID is the ID of the account that the transaction belongs to.
	AccountID uuid.UUID
	// Amount is the amount of the transaction.
	Amount float64
	// Operation is the operation of the transaction.
	Operation string
	// Date is the date of the transaction.
	Date time.Time
	// CreatedAt is the date and time when the transaction was created.
	CreatedAt time.Time
	// Origin is the origin of the transaction.
	Origin string
}

// NewTransaction returns a new Transaction instance.
func NewTransaction(accountID uuid.UUID, amount float64, dateTx time.Time, origin string) (tx *Transaction, err error) {
	if amount == 0 {
		err = transaction.ErrTransactionAmountIsZero
		return
	}
	if origin == "" {
		err = transaction.ErrTransactionOriginIsEmpty
		return
	}
	if dateTx.IsZero() {
		err = transaction.ErrTransactionDateIsInvalid
		return
	}
	operation := "credit"
	if amount < 0 {
		operation = "debit"
	}

	tx = &Transaction{
		ID:        uuid.New(),
		AccountID: accountID,
		Amount:    amount,
		Operation: operation,
		Date:      dateTx,
		CreatedAt: time.Now(),
		Origin:    origin,
	}
	return
}
