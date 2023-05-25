package entity_test

import (
	"testing"

	"github.com/braejan/go-transactions-summary/internal/domain/account/entity"
	"github.com/stretchr/testify/assert"
)

// TestNewAccount tests the NewAccount function.
func TestNewAccount(t *testing.T) {
	// Given a valid user ID.
	userID := int64(1)
	// When call the NewAccount function.
	account := entity.NewAccount(userID)
	// Then the account must be created.
	assert.NotNil(t, account)
	assert.NotEmpty(t, account.ID)
	assert.Equal(t, userID, account.UserID)
	assert.Equal(t, float64(0), account.Balance)
	assert.False(t, account.Active)
}
