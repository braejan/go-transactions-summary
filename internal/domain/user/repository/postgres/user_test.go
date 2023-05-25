package postgres_test

import (
	"testing"

	"github.com/braejan/go-transactions-summary/internal/domain/user/repository/postgres"
	voPostgres "github.com/braejan/go-transactions-summary/internal/valueobject/postgres"
	"github.com/braejan/go-transactions-summary/internal/valueobject/postgres/mock"

	"github.com/stretchr/testify/assert"
)

// TestNewPostgresUserRepositoryConfigNil tests the NewPostgresUserRepository function
// when the configuration is nil.
func TestNewPostgresUserRepositoryConfigNil(t *testing.T) {
	// Given a nil configuration.
	// When NewPostgresUserRepository is called.
	_, err := postgres.NewPostgresUserRepository(nil, nil)
	// Then the error returned should be ErrNilConfiguration.
	assert.NotNil(t, err)
	assert.Equal(t, voPostgres.ErrNilConfiguration, err)
}

// TestNewPostgresUserRepositoryWithConfig tests the NewPostgresUserRepository function
// when the configuration is not nil.
func TestNewPostgresUserRepositoryWithConfig(t *testing.T) {
	// Given a valid configuration.
	configuration := voPostgres.GetPostgresConfigurationFromEnv()
	dbBase := mock.NewMockBasePostgresDatabase()
	// When NewPostgresUserRepository is called.
	_, err := postgres.NewPostgresUserRepository(configuration, dbBase)
	// Then the error returned should be nil.
	assert.Nil(t, err)
}
