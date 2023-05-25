package postgres_test

import (
	"testing"

	"github.com/braejan/go-transactions-summary/internal/domain/account/repository/postgres"
	voPostgres "github.com/braejan/go-transactions-summary/internal/valueobject/postgres"
	"github.com/braejan/go-transactions-summary/internal/valueobject/postgres/mock"
	"github.com/stretchr/testify/assert"
)

// TestNewPostgresAccountRepositoryConfigNil tests the NewPostgresAccountRepository function
// when the configuration is nil.
func TestNewPostgresAccountRepositoryConfigNil(t *testing.T) {
	// Given a nil configuration.
	// When NewPostgresAccountRepository is called.
	_, err := postgres.NewPostgresAccountRepository(nil, nil)
	// Then the error returned should be ErrNilConfiguration.
	assert.NotNil(t, err)
	assert.Equal(t, voPostgres.ErrNilConfiguration, err)
}

// TestNewPostgresAccountRepositoryWithConfig tests the NewPostgresAccountRepository function
// when the configuration is not nil.
func TestNewPostgresAccountRepositoryWithConfig(t *testing.T) {
	// Given a valid configuration.
	configuration := voPostgres.GetPostgresConfigurationFromEnv()
	dbBase := mock.NewMockBasePostgresDatabase()
	// When NewPostgresUserRepository is called.
	_, err := postgres.NewPostgresAccountRepository(configuration, dbBase)
	// Then the error returned should be nil.
	assert.Nil(t, err)
}
