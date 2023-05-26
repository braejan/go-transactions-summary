package postgres_test

import (
	"testing"

	"github.com/braejan/go-transactions-summary/internal/domain/transaction/repository/postgres"
	voPostgres "github.com/braejan/go-transactions-summary/internal/valueobject/postgres"
	"github.com/stretchr/testify/assert"
)

// TestNewPostgresTransactionRepositoryWithNilConfig tests the NewPostgresTransactionRepository function
// when the configuration is nil.
func TestNewPostgresTransactionRepositoryWithNilConfig(t *testing.T) {
	// Given a nil configuration.
	// When NewPostgresTransactionRepository is called.
	_, err := postgres.NewPostgresTransactionRepository(nil, nil)
	// Then the error returned should be ErrNilConfiguration.
	assert.NotNil(t, err)
	assert.Equal(t, voPostgres.ErrNilConfiguration, err)
}
