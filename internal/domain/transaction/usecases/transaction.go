package usecases

import (
	txEntity "github.com/braejan/go-transactions-summary/internal/domain/transaction/entity"
	"github.com/braejan/go-transactions-summary/internal/domain/transaction/repository"
	"github.com/braejan/go-transactions-summary/internal/domain/transaction/util"
	voTransaction "github.com/braejan/go-transactions-summary/internal/valueobject/transaction"
	"github.com/google/uuid"
)

// transactionUseCases implements the transaction use cases.
type transactionUseCases struct {
	// transactionRepo is the transaction repository.
	transactionRepo repository.TransactionRepository
}

// NewTransactionUseCases returns a new transaction use cases.
func NewTransactionUseCases(transactionRepo repository.TransactionRepository) (usecases TransactionUseCases, err error) {
	if transactionRepo == nil {
		err = voTransaction.ErrNilTransactionRepo
		return
	}
	usecases = &transactionUseCases{
		transactionRepo: transactionRepo,
	}
	return
}

// TransactionUseCases interface implementation

// GetByID returns a transaction by its ID.
func (uc *transactionUseCases) GetByID(txID uuid.UUID) (tx txEntity.Transaction, err error) {
	txAux, err := uc.transactionRepo.GetByID(txID)
	if err != nil {
		return
	}
	tx = *txAux
	return
}

// GetByAccountID returns a transaction by its account ID.
func (uc *transactionUseCases) GetByAccountID(accountID uuid.UUID) (txs []txEntity.Transaction, err error) {
	accAux, err := uc.transactionRepo.GetByAccountID(accountID)
	if err != nil {
		txs = nil
		return
	}
	txs = util.ArrayTxMemoryToArrayValue(accAux)
	return
}

// GetCreditsByAccountID returns the credits of an account.
func (uc *transactionUseCases) GetCreditsByAccountID(accountID uuid.UUID) (txs []txEntity.Transaction, err error) {
	accAux, err := uc.transactionRepo.GetCreditsByAccountID(accountID)
	if err != nil {
		txs = nil
		return
	}
	txs = util.ArrayTxMemoryToArrayValue(accAux)
	return
}

// GetDebitsByAccountID returns the debits of an account.
func (uc *transactionUseCases) GetDebitsByAccountID(accountID uuid.UUID) (txs []txEntity.Transaction, err error) {
	accAux, err := uc.transactionRepo.GetDebitsByAccountID(accountID)
	if err != nil {
		txs = nil
		return
	}
	txs = util.ArrayTxMemoryToArrayValue(accAux)
	return
}

// GetTransactionsByOrigin returns the transactions of an account by origin.
func (uc *transactionUseCases) GetTransactionsByOrigin(origin string) (txs []txEntity.Transaction, err error) {
	accAux, err := uc.transactionRepo.GetTransactionsByOrigin(origin)
	if err != nil {
		txs = nil
		return
	}
	txs = util.ArrayTxMemoryToArrayValue(accAux)
	return
}

// Create creates a new transaction.
func (uc *transactionUseCases) Create(tx txEntity.Transaction) (err error) {
	err = uc.transactionRepo.Create(&tx)
	return
}
