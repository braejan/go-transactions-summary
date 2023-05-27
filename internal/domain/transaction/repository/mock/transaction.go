package mock

import (
	"github.com/braejan/go-transactions-summary/internal/domain/transaction/entity"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

// mockTransactionRepository is a mock implementation of the transaction repository.
type mockTransactionRepository struct {
	mock.Mock
}

// NewMockTransactionRepository returns a new mock transaction repository.
func NewMockTransactionRepository() *mockTransactionRepository {
	return &mockTransactionRepository{}
}

// GetByID returns a transaction by its ID.
func (m *mockTransactionRepository) GetByID(ID uuid.UUID) (tx *entity.Transaction, err error) {
	args := m.Called(ID)

	var r0 *entity.Transaction
	if rf, ok := args.Get(0).(func(uuid.UUID) *entity.Transaction); ok {
		r0 = rf(ID)
	} else {
		if args.Get(0) != nil {
			r0 = args.Get(0).(*entity.Transaction)
		}
	}

	var r1 error
	if rf, ok := args.Get(1).(func(uuid.UUID) error); ok {
		r1 = rf(ID)
	} else {
		r1 = args.Error(1)
	}

	return r0, r1
}

// GetByAccountID returns a list of transactions by its account ID.
func (m *mockTransactionRepository) GetByAccountID(accountID uuid.UUID) (txs []*entity.Transaction, err error) {
	args := m.Called(accountID)

	var r0 []*entity.Transaction
	if rf, ok := args.Get(0).(func(uuid.UUID) []*entity.Transaction); ok {
		r0 = rf(accountID)
	} else {
		if args.Get(0) != nil {
			r0 = args.Get(0).([]*entity.Transaction)
		}
	}

	var r1 error
	if rf, ok := args.Get(1).(func(uuid.UUID) error); ok {
		r1 = rf(accountID)
	} else {
		r1 = args.Error(1)
	}

	return r0, r1
}

// GetCreditsByAccountID returns the credits of an account.
func (m *mockTransactionRepository) GetCreditsByAccountID(accountID uuid.UUID) (txs []*entity.Transaction, err error) {
	args := m.Called(accountID)

	var r0 []*entity.Transaction
	if rf, ok := args.Get(0).(func(uuid.UUID) []*entity.Transaction); ok {
		r0 = rf(accountID)
	} else {
		if args.Get(0) != nil {
			r0 = args.Get(0).([]*entity.Transaction)
		}
	}

	var r1 error
	if rf, ok := args.Get(1).(func(uuid.UUID) error); ok {
		r1 = rf(accountID)
	} else {
		r1 = args.Error(1)
	}

	return r0, r1
}

// GetDebitsByAccountID returns the debits of an account.
func (m *mockTransactionRepository) GetDebitsByAccountID(accountID uuid.UUID) (txs []*entity.Transaction, err error) {
	args := m.Called(accountID)

	var r0 []*entity.Transaction
	if rf, ok := args.Get(0).(func(uuid.UUID) []*entity.Transaction); ok {
		r0 = rf(accountID)
	} else {
		if args.Get(0) != nil {
			r0 = args.Get(0).([]*entity.Transaction)
		}
	}

	var r1 error
	if rf, ok := args.Get(1).(func(uuid.UUID) error); ok {
		r1 = rf(accountID)
	} else {
		r1 = args.Error(1)
	}

	return r0, r1
}

// GetTransactionsByOrigin returns the transactions of an account by origin.
func (m *mockTransactionRepository) GetTransactionsByOrigin(origin string) (txs []*entity.Transaction, err error) {
	args := m.Called(origin)

	var r0 []*entity.Transaction
	if rf, ok := args.Get(0).(func(string) []*entity.Transaction); ok {
		r0 = rf(origin)
	} else {
		if args.Get(0) != nil {
			r0 = args.Get(0).([]*entity.Transaction)
		}
	}

	var r1 error
	if rf, ok := args.Get(1).(func(string) error); ok {
		r1 = rf(origin)
	} else {
		r1 = args.Error(1)
	}

	return r0, r1
}

// Create creates a new transaction.
func (m *mockTransactionRepository) Create(tx *entity.Transaction) (err error) {
	args := m.Called(tx)

	var r0 error
	if rf, ok := args.Get(0).(func(*entity.Transaction) error); ok {
		r0 = rf(tx)
	} else {
		r0 = args.Error(0)
	}

	return r0
}
