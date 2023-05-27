package entity_test

import (
	"testing"
	"time"

	"github.com/braejan/go-transactions-summary/internal/domain/transaction/entity"
	"github.com/braejan/go-transactions-summary/internal/valueobject/transaction"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// TestNewTransactionWithZeroAmount tests the NewTransaction function with zero amount.
func TestNewTransactionWithZeroAmount(t *testing.T) {
	// Given a valid account ID, zero amount, date and origin.
	accountID := uuid.New()
	amount := float64(0)
	origin := "txns.csv"
	// When call the NewTransaction function.
	tx, err := entity.NewTransaction(accountID, amount, time.Now(), origin)
	// Then the transaction must not be created.
	assert.Nil(t, tx)
	assert.Equal(t, transaction.ErrTransactionAmountIsZero, err)
}

// TestNewTransactionWithEmptyOrigin tests the NewTransaction function with empty origin.
func TestNewTransactionWithEmptyOrigin(t *testing.T) {
	// Given a valid account ID, amount, date and empty origin.
	accountID := uuid.New()
	amount := float64(100)
	origin := ""
	// When call the NewTransaction function.
	tx, err := entity.NewTransaction(accountID, amount, time.Now(), origin)
	// Then the transaction must not be created.
	assert.Nil(t, tx)
	assert.Equal(t, transaction.ErrTransactionOriginIsEmpty, err)
}

// TestNewTransactionWithCredit tests the NewTransaction function with credit operation.
func TestNewTransactionWithCredit(t *testing.T) {
	// Given a valid account ID, amount, date and origin.
	accountID := uuid.New()
	amount := float64(100.48)
	// A valid date parsed for the string "7/28".
	date, err := time.Parse("1/2", "7/28")
	assert.Nil(t, err)
	origin := "txns.csv"
	// When call the NewTransaction function.
	tx, err := entity.NewTransaction(accountID, amount, date, origin)
	// Then the transaction must be created.
	assert.Nil(t, err)
	assert.NotNil(t, tx)
	assert.Equal(t, accountID, tx.AccountID)
	assert.Equal(t, amount, tx.Amount)
	assert.Equal(t, "credit", tx.Operation)
	assert.Equal(t, date, tx.Date)
	assert.NotEmpty(t, tx.CreatedAt)
	assert.Equal(t, origin, tx.Origin)
}

// TestNewTransactionWithDebit tests the NewTransaction function with debit operation.
func TestNewTransactionWithDebit(t *testing.T) {
	// Given a valid account ID, amount, date and origin.
	accountID := uuid.New()
	amount := float64(-100.48)
	// A valid date parsed for the string "7/28".
	date, err := time.Parse("1/2", "7/28")
	assert.Nil(t, err)
	origin := "txns.csv"
	// When call the NewTransaction function.
	tx, err := entity.NewTransaction(accountID, amount, date, origin)
	// Then the transaction must be created.
	assert.Nil(t, err)
	assert.NotNil(t, tx)
	assert.Equal(t, accountID, tx.AccountID)
	assert.Equal(t, amount, tx.Amount)
	assert.Equal(t, "debit", tx.Operation)
	assert.Equal(t, date, tx.Date)
	assert.NotEmpty(t, tx.CreatedAt)
	assert.Equal(t, origin, tx.Origin)
}

// TestNewTransactionWithInvalidDate tests the NewTransaction function with invalid date.
func TestNewTransactionWithInvalidDate(t *testing.T) {
	// Given a valid account ID, amount, invalid date and origin.
	accountID := uuid.New()
	amount := float64(100.48)
	date := time.Time{}
	origin := "txns.csv"
	// When call the NewTransaction function.
	tx, err := entity.NewTransaction(accountID, amount, date, origin)
	// Then the transaction must not be created.
	assert.Nil(t, tx)
	assert.Equal(t, transaction.ErrTransactionDateIsInvalid, err)
}
