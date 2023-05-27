package repository

import (
	"github.com/braejan/go-transactions-summary/internal/domain/transaction/entity"
	"github.com/google/uuid"
)

// TransactionRepository interface defines the methods that the transaction repository must implement.
type TransactionRepository interface {
	// GetByID returns a transaction by its ID.
	GetByID(ID uuid.UUID) (tx *entity.Transaction, err error)
	// GetByAccountID returns a transaction by its account ID.
	GetByAccountID(accountID uuid.UUID) (txs []*entity.Transaction, err error)
	// GetCreditsByAccountID returns the credits of an account.
	GetCreditsByAccountID(accountID uuid.UUID) (txs []*entity.Transaction, err error)
	// GetDebitsByAccountID returns the debits of an account.
	GetDebitsByAccountID(accountID uuid.UUID) (txs []*entity.Transaction, err error)
	// GetTransactionsByOrigin returns the transactions of an account by origin.
	GetTransactionsByOrigin(origin string) (txs []*entity.Transaction, err error)
	// Create creates a new transaction.
	Create(tx *entity.Transaction) (err error)
}
