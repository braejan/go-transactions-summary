package mock

import (
	"github.com/braejan/go-transactions-summary/internal/domain/transaction/entity"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

// mockTransactionUseCases struct implements the TransactionUseCases interface.
type mockTransactionUseCases struct {
	mock.Mock
}

// NewMockTransactionUseCases returns a new mockTransactionUseCases instance.
func NewMockTransactionUseCases() (usecases *mockTransactionUseCases) {
	usecases = &mockTransactionUseCases{}
	return
}

// TransactionUseCases interface implementation:

// GetByID implements the TransactionUseCases interface method.
func (m *mockTransactionUseCases) GetByID(ID uuid.UUID) (tx entity.Transaction, err error) {
	args := m.Called(ID)
	tx = args.Get(0).(entity.Transaction)
	err = args.Error(1)
	return
}

// GetByAccountID implements the TransactionUseCases interface method.
func (m *mockTransactionUseCases) GetByAccountID(accountID uuid.UUID) (txs []entity.Transaction, err error) {
	args := m.Called(accountID)
	txs = args.Get(0).([]entity.Transaction)
	err = args.Error(1)
	return
}

// GetCreditsByAccountID implements the TransactionUseCases interface method.
func (m *mockTransactionUseCases) GetCreditsByAccountID(accountID uuid.UUID) (txs []entity.Transaction, err error) {
	args := m.Called(accountID)
	txs = args.Get(0).([]entity.Transaction)
	err = args.Error(1)
	return
}

// GetDebitsByAccountID implements the TransactionUseCases interface method.
func (m *mockTransactionUseCases) GetDebitsByAccountID(accountID uuid.UUID) (txs []entity.Transaction, err error) {
	args := m.Called(accountID)
	txs = args.Get(0).([]entity.Transaction)
	err = args.Error(1)
	return
}

// GetTransactionsByOrigin implements the TransactionUseCases interface method.
func (m *mockTransactionUseCases) GetTransactionsByOrigin(origin string) (txs []entity.Transaction, err error) {
	args := m.Called(origin)
	txs = args.Get(0).([]entity.Transaction)
	err = args.Error(1)
	return
}

// Create implements the TransactionUseCases interface method.
func (m *mockTransactionUseCases) Create(tx entity.Transaction) (err error) {
	args := m.Called(tx)
	err = args.Error(0)
	return
}
