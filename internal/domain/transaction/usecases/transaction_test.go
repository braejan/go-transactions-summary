package usecases_test

import (
	"testing"
	"time"

	txEntity "github.com/braejan/go-transactions-summary/internal/domain/transaction/entity"
	txMock "github.com/braejan/go-transactions-summary/internal/domain/transaction/repository/mock"
	"github.com/braejan/go-transactions-summary/internal/domain/transaction/usecases"
	voTransaction "github.com/braejan/go-transactions-summary/internal/valueobject/transaction"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func getTestTransactions() (transactions []*txEntity.Transaction) {
	// Append 5 more transactions with negatives amounts
	for i := 0; i < 5; i++ {
		transactions = append(transactions, &txEntity.Transaction{
			ID:        uuid.New(),
			AccountID: uuid.New(),
			Amount:    100.51,
			Operation: "creadit",
			Date:      time.Now(),
			CreatedAt: time.Now(),
			Origin:    "txns.csv",
		})
	}
	// Append 5 more transactions with negatives amounts
	for i := 0; i < 5; i++ {
		transactions = append(transactions, &txEntity.Transaction{
			ID:        uuid.New(),
			AccountID: uuid.New(),
			Amount:    -100.52,
			Operation: "debit",
			Date:      time.Now(),
			CreatedAt: time.Now(),
			Origin:    "txns.csv",
		})
	}
	return
}

// TestNewTransactionUseCases_NilTransactionRepo tests the NewTransactionUseCases function with a nil transaction repository.
func TestNewTransactionUseCases_NilTransactionRepo(t *testing.T) {
	// Given a nil transaction repository
	// When calling NewTransactionUseCases
	// Then it should return an error.
	_, err := usecases.NewTransactionUseCases(nil)
	assert.Error(t, voTransaction.ErrNilTransactionRepo, err)
}

// TestNewTransactionUseCases tests the NewTransactionUseCases function.
func TestNewTransactionUseCases(t *testing.T) {
	// Given a valid transaction repository
	mockTransactionRepo := txMock.NewMockTransactionRepository()
	// When calling NewTransactionUseCases
	// Then it should return a new TransactionUseCases instance.
	uc, err := usecases.NewTransactionUseCases(mockTransactionRepo)
	assert.Nil(t, err)
	assert.NotNil(t, uc)
}

// TestTransactionUseCases_GetByID tests the GetByID function.
func TestTransactionUseCases_GetByID(t *testing.T) {
	// Given a valid transaction repository
	mockTransactionRepo := txMock.NewMockTransactionRepository()
	// And a valid transaction entity
	txToTest, err := txEntity.NewTransaction(uuid.New(), 100.50, time.Now(), "txns.csv")
	assert.Nil(t, err)
	mockTransactionRepo.On("GetByID", txToTest.ID).Return(nil, voTransaction.ErrQueryingTransactionByID)
	// And a valid transaction use cases
	uc, err := usecases.NewTransactionUseCases(mockTransactionRepo)
	assert.Nil(t, err)
	assert.NotNil(t, uc)
	// When calling GetByID with a valid transaction ID
	// Then it should return a transaction.
	_, err = uc.GetByID(txToTest.ID)
	assert.NotNil(t, err)
	assert.Equal(t, voTransaction.ErrQueryingTransactionByID, err)
}

// TestTransactionUseCasesSuccess_GetByID tests the GetByID function.
func TestTransactionUseCasesSuccess_GetByID(t *testing.T) {
	// Given a valid transaction repository
	mockTransactionRepo := txMock.NewMockTransactionRepository()
	// And a valid transaction entity
	txToTest, err := txEntity.NewTransaction(uuid.New(), 100.50, time.Now(), "txns.csv")
	assert.Nil(t, err)
	mockTransactionRepo.On("GetByID", txToTest.ID).Return(txToTest, nil)
	// And a valid transaction use cases
	uc, err := usecases.NewTransactionUseCases(mockTransactionRepo)
	assert.Nil(t, err)
	assert.NotNil(t, uc)
	// When calling GetByID with a valid transaction ID
	// Then it should return a transaction.
	tx, err := uc.GetByID(txToTest.ID)
	assert.Nil(t, err)
	assert.Equal(t, *txToTest, tx)
}

// TestTransactionUseCases_GetByAccountID_Err
func TestTransactionUseCases_GetByAccountID_Err(t *testing.T) {
	// Given a valid transaction repository
	mockTransactionRepo := txMock.NewMockTransactionRepository()
	// And a valid transaction entity
	txToTest, err := txEntity.NewTransaction(uuid.New(), 100.50, time.Now(), "txns.csv")
	assert.Nil(t, err)
	mockTransactionRepo.On("GetByAccountID", txToTest.AccountID).Return(nil, voTransaction.ErrQueryingTransactionsByAccountID)
	// And a valid transaction use cases
	uc, err := usecases.NewTransactionUseCases(mockTransactionRepo)
	assert.Nil(t, err)
	assert.NotNil(t, uc)
	// When calling GetByAccountID with a valid transaction ID
	// Then it should return a transaction.
	_, err = uc.GetByAccountID(txToTest.AccountID)
	assert.NotNil(t, err)
	assert.Equal(t, voTransaction.ErrQueryingTransactionsByAccountID, err)
}

// TestTransactionUseCases_GetByAccountID_Success
func TestTransactionUseCases_GetByAccountID_Success(t *testing.T) {
	// Given a valid transaction repository
	mockTransactionRepo := txMock.NewMockTransactionRepository()
	txToTest := getTestTransactions()
	for _, transaction := range txToTest {
		mockTransactionRepo.On("GetByAccountID", transaction.AccountID).Return(getTestTransactions(), nil)
	}
	// And a valid transaction use cases
	uc, err := usecases.NewTransactionUseCases(mockTransactionRepo)
	assert.Nil(t, err)
	assert.NotNil(t, uc)
	// When calling GetByAccountID with a valid transaction ID
	// Then it should return a transaction.
	txs, err := uc.GetByAccountID(txToTest[0].AccountID)
	assert.Nil(t, err)
	assert.Equal(t, 10, len(txs))
}

// TestGetCreditsByAccountID_GetCreditsByAccountID
func TestGetCreditsByAccountID_GetCreditsByAccountID(t *testing.T) {
	// Given a valid transaction repository
	mockTransactionRepo := txMock.NewMockTransactionRepository()
	txToTest := getTestTransactions()
	for _, transaction := range txToTest {
		mockTransactionRepo.On("GetCreditsByAccountID", transaction.AccountID).Return(nil, voTransaction.ErrQueryingCreditsByAccountID)
	}
	// And a valid transaction use cases
	uc, err := usecases.NewTransactionUseCases(mockTransactionRepo)
	assert.Nil(t, err)
	assert.NotNil(t, uc)
	// When calling GetByAccountID with a valid transaction ID
	txs, err := uc.GetCreditsByAccountID(txToTest[0].AccountID)
	// Then it should return a error ErrQueryingCreditsByAccountID
	assert.NotNil(t, err)
	assert.Nil(t, txs)
	assert.Equal(t, voTransaction.ErrQueryingCreditsByAccountID, err)

}

// TestGetCreditsByAccountID_Success
func TestGetCreditsByAccountID_Success(t *testing.T) {
	// Given a valid transaction repository
	mockTransactionRepo := txMock.NewMockTransactionRepository()
	txToTest := getTestTransactions()
	for _, transaction := range txToTest {
		mockTransactionRepo.On("GetCreditsByAccountID", transaction.AccountID).Return(getTestTransactions(), nil)
	}
	// And a valid transaction use cases
	uc, err := usecases.NewTransactionUseCases(mockTransactionRepo)
	assert.Nil(t, err)
	assert.NotNil(t, uc)
	// When calling GetByAccountID with a valid transaction ID
	txs, err := uc.GetCreditsByAccountID(txToTest[0].AccountID)
	// Then it should return a error ErrQueryingCreditsByAccountID
	assert.Nil(t, err)
	assert.Equal(t, 10, len(txs))
}

// TestGetDebitsByAccountID_GetDebitsByAccountID
func TestGetDebitsByAccountID_GetDebitsByAccountID(t *testing.T) {
	// Given a valid transaction repository
	mockTransactionRepo := txMock.NewMockTransactionRepository()
	txToTest := getTestTransactions()
	for _, transaction := range txToTest {
		mockTransactionRepo.On("GetDebitsByAccountID", transaction.AccountID).Return(nil, voTransaction.ErrQueryingDebitsByAccountID)
	}
	// And a valid transaction use cases
	uc, err := usecases.NewTransactionUseCases(mockTransactionRepo)
	assert.Nil(t, err)
	assert.NotNil(t, uc)
	// When calling GetByAccountID with a valid transaction ID
	txs, err := uc.GetDebitsByAccountID(txToTest[0].AccountID)
	// Then it should return a error ErrQueryingCreditsByAccountID
	assert.NotNil(t, err)
	assert.Nil(t, txs)
	assert.Equal(t, voTransaction.ErrQueryingDebitsByAccountID, err)

}

// TestGetDebitsByAccountID_Success
func TestGetDebitsByAccountID_Success(t *testing.T) {
	// Given a valid transaction repository
	mockTransactionRepo := txMock.NewMockTransactionRepository()
	txToTest := getTestTransactions()
	for _, transaction := range txToTest {
		mockTransactionRepo.On("GetDebitsByAccountID", transaction.AccountID).Return(getTestTransactions(), nil)
	}
	// And a valid transaction use cases
	uc, err := usecases.NewTransactionUseCases(mockTransactionRepo)
	assert.Nil(t, err)
	assert.NotNil(t, uc)
	// When calling GetByAccountID with a valid transaction ID
	txs, err := uc.GetDebitsByAccountID(txToTest[0].AccountID)
	// Then it should return a error ErrQueryingCreditsByAccountID
	assert.Nil(t, err)
	assert.Equal(t, 10, len(txs))
}

// TestGetTransactionsByOrigin_GetTransactionsByOrigin
func TestGetTransactionsByOrigin_GetTransactionsByOrigin(t *testing.T) {
	// Given a valid transaction repository
	mockTransactionRepo := txMock.NewMockTransactionRepository()
	txToTest := getTestTransactions()
	for _, transaction := range txToTest {
		mockTransactionRepo.On("GetTransactionsByOrigin", transaction.Origin).Return(nil, voTransaction.ErrQueryingTransactionsByOrigin)
	}
	// And a valid transaction use cases
	uc, err := usecases.NewTransactionUseCases(mockTransactionRepo)
	assert.Nil(t, err)
	assert.NotNil(t, uc)
	// When calling GetByAccountID with a valid transaction ID
	txs, err := uc.GetTransactionsByOrigin(txToTest[0].Origin)
	// Then it should return a error ErrQueryingCreditsByAccountID
	assert.NotNil(t, err)
	assert.Nil(t, txs)
	assert.Equal(t, voTransaction.ErrQueryingTransactionsByOrigin, err)
}

// TestGetTransactionsByOrigin_Success
func TestGetTransactionsByOrigin_Success(t *testing.T) {
	// Given a valid transaction repository
	mockTransactionRepo := txMock.NewMockTransactionRepository()
	txToTest := getTestTransactions()
	for _, transaction := range txToTest {
		mockTransactionRepo.On("GetTransactionsByOrigin", transaction.Origin).Return(getTestTransactions(), nil)
	}
	// And a valid transaction use cases
	uc, err := usecases.NewTransactionUseCases(mockTransactionRepo)
	assert.Nil(t, err)
	assert.NotNil(t, uc)
	// When calling GetByAccountID with a valid transaction ID
	txs, err := uc.GetTransactionsByOrigin(txToTest[0].Origin)
	// Then it should return a error ErrQueryingCreditsByAccountID
	assert.Nil(t, err)
	assert.Equal(t, 10, len(txs))
}

// TestCreate_Create
func TestCreate_Create(t *testing.T) {
	// Given a valid transaction repository
	mockTransactionRepo := txMock.NewMockTransactionRepository()
	transactionToTest := getTestTransactions()[0]
	mockTransactionRepo.On("Create", transactionToTest).Return(voTransaction.ErrCreatingTransaction)
	// And a valid transaction use cases
	uc, err := usecases.NewTransactionUseCases(mockTransactionRepo)
	assert.Nil(t, err)
	assert.NotNil(t, uc)
	err = uc.Create(*transactionToTest)
	// Then it should return a error ErrCreatingTransaction
	assert.NotNil(t, err)
	assert.Equal(t, voTransaction.ErrCreatingTransaction, err)
}

// TestCreate_Success
func TestCreate_Success(t *testing.T) {
	// Given a valid transaction repository
	mockTransactionRepo := txMock.NewMockTransactionRepository()
	transactionToTest := getTestTransactions()[0]
	mockTransactionRepo.On("Create", transactionToTest).Return(nil)
	// And a valid transaction use cases
	uc, err := usecases.NewTransactionUseCases(mockTransactionRepo)
	assert.Nil(t, err)
	assert.NotNil(t, uc)
	err = uc.Create(*transactionToTest)
	// Then it should return a error ErrCreatingTransaction
	assert.Nil(t, err)
}
