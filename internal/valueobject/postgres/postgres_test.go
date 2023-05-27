package postgres_test

import (
	"log"
	"os"
	"testing"

	"github.com/braejan/go-transactions-summary/internal/valueobject/postgres"
	"github.com/stretchr/testify/assert"
)

func resetEnvironmentPostgresVariables() {
	os.Unsetenv("POSTGRES_HOST")
	os.Unsetenv("POSTGRES_PORT")
	os.Unsetenv("POSTGRES_USER")
	os.Unsetenv("POSTGRES_PASSWORD")
	os.Unsetenv("POSTGRES_DATABASE")
}

// TestNewPostgresConfigurationSuccess tests the NewPostgresConfiguration function succeeds.
func TestNewPostgresConfigurationSuccess(t *testing.T) {
	// Given a host
	host := "test"
	// And a port
	port := 1234
	// And a user
	user := "test"
	// And a password
	password := "test"
	// And a database
	database := "test"
	// When call NewPostgresConfiguration
	configuration := postgres.NewPostgresConfiguration(host, port, user, password, database)
	// Then return a PostgresConfiguration
	assert.Equal(t, host, configuration.Host)
	assert.Equal(t, port, configuration.Port)
	assert.Equal(t, user, configuration.User)
	assert.Equal(t, password, configuration.Password)
	assert.Equal(t, database, configuration.Database)
}

// TestGetDefaultPostgresConfigurationSuccess tests the GetDefaultPostgresConfiguration function succeeds.
func TestGetDefaultPostgresConfigurationSuccess(t *testing.T) {
	// When call GetDefaultPostgresConfiguration
	configuration := postgres.NewDefaultPostgresConfiguration()
	log.Printf("datasource: %s\n", configuration.GetDataSourceName())
	// Then return a PostgresConfiguration
	assert.Equal(t, "localhost", configuration.Host)
	assert.Equal(t, 5432, configuration.Port)
	assert.Equal(t, "postgres", configuration.User)
	assert.Equal(t, "postgres", configuration.Password)
	assert.Equal(t, "stori-challenge-db", configuration.Database)
}

// TestGetPostgresConfigurationFromEnvSuccess tests the GetPostgresConfigurationFromEnv function succeeds.
func TestGetPostgresConfigurationFromEnvSuccess(t *testing.T) {
	// reset environment variables
	resetEnvironmentPostgresVariables()
	// Given a POSTGRES_HOST environment variable
	err := os.Setenv("POSTGRES_HOST", "localhost")
	assert.NoError(t, err)
	// And a POSTGRES_PORT environment variable
	err = os.Setenv("POSTGRES_PORT", "1234")
	assert.NoError(t, err)
	// And a POSTGRES_USER environment variable
	err = os.Setenv("POSTGRES_USER", "postgres")
	assert.NoError(t, err)
	// And a POSTGRES_PASSWORD environment variable
	err = os.Setenv("POSTGRES_PASSWORD", "postgres")
	assert.NoError(t, err)
	// And a POSTGRES_DATABASE environment variable
	err = os.Setenv("POSTGRES_DATABASE", "stori-challenge-db")
	assert.NoError(t, err)
	// When call GetPostgresConfigurationFromEnv
	configuration := postgres.NewPostgresConfigurationFromEnv()
	// Then return a PostgresConfiguration
	assert.Equal(t, "localhost", configuration.Host)
	assert.Equal(t, 1234, configuration.Port)
	assert.Equal(t, "postgres", configuration.User)
	assert.Equal(t, "postgres", configuration.Password)
	assert.Equal(t, "stori-challenge-db", configuration.Database)
}

// TestGetPostgresConfigurationFromEnvWithEmptyHostSuccess tests the GetPostgresConfigurationFromEnv function succeeds.
func TestGetPostgresConfigurationFromEnvWithEmptyHostSuccess(t *testing.T) {
	// reset environment variables
	resetEnvironmentPostgresVariables()
	// Given a POSTGRES_PORT environment variable
	err := os.Setenv("POSTGRES_PORT", "1234")
	assert.NoError(t, err)
	// And a POSTGRES_USER environment variable
	err = os.Setenv("POSTGRES_USER", "postgres")
	assert.NoError(t, err)
	// And a POSTGRES_PASSWORD environment variable
	err = os.Setenv("POSTGRES_PASSWORD", "postgres")
	assert.NoError(t, err)
	// And a POSTGRES_DATABASE environment variable
	err = os.Setenv("POSTGRES_DATABASE", "other-db")
	assert.NoError(t, err)
	// When call GetPostgresConfigurationFromEnv
	configuration := postgres.NewPostgresConfigurationFromEnv()
	// Then return a PostgresConfiguration
	assert.Equal(t, "localhost", configuration.Host)
	assert.Equal(t, 5432, configuration.Port)
	assert.Equal(t, "postgres", configuration.User)
	assert.Equal(t, "postgres", configuration.Password)
	assert.Equal(t, "stori-challenge-db", configuration.Database)
}

// TestGetPostgresConfigurationFromEnvWithEmptyPortSuccess tests the GetPostgresConfigurationFromEnv function succeeds.
func TestGetPostgresConfigurationFromEnvWithEmptyPortSuccess(t *testing.T) {
	// reset environment variables
	resetEnvironmentPostgresVariables()
	// Given a POSTGRES_HOST environment variable
	err := os.Setenv("POSTGRES_HOST", "localhost")
	assert.NoError(t, err)
	// And a POSTGRES_USER environment variable
	err = os.Setenv("POSTGRES_USER", "postgres")
	assert.NoError(t, err)
	// And a POSTGRES_PASSWORD environment variable
	err = os.Setenv("POSTGRES_PASSWORD", "postgres")
	assert.NoError(t, err)
	// And a POSTGRES_DATABASE environment variable
	err = os.Setenv("POSTGRES_DATABASE", "other-db")
	assert.NoError(t, err)
	// When call GetPostgresConfigurationFromEnv
	configuration := postgres.NewPostgresConfigurationFromEnv()
	// Then return a PostgresConfiguration
	assert.Equal(t, "localhost", configuration.Host)
	assert.Equal(t, 5432, configuration.Port)
	assert.Equal(t, "postgres", configuration.User)
	assert.Equal(t, "postgres", configuration.Password)
	assert.Equal(t, "stori-challenge-db", configuration.Database)
}

// TestGetDataSoureNameSuccess tests the GetDataSourceName function succeeds.
func TestGetDataSoureNameSuccess(t *testing.T) {
	// Given a PostgresConfiguration
	configuration := postgres.NewDefaultPostgresConfiguration()
	// When call GetDataSourceName
	dataSourceName := configuration.GetDataSourceName()
	// Then return a data source name
	assert.Equal(t, "host=localhost port=5432 user=postgres password=postgres dbname=stori-challenge-db sslmode=disable", dataSourceName)
}
