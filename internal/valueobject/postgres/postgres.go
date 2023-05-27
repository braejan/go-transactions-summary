package postgres

import (
	"database/sql"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

type PostgresDatabase interface {
	Open() (db *sql.DB, err error)
	Close(db *sql.DB) (err error)
	BeginTx(db *sql.DB) (tx *sql.Tx, err error)
	Commit(tx *sql.Tx) (err error)
	Rollback(tx *sql.Tx) (err error)
	Exec(tx *sql.Tx, query string, args ...interface{}) (result sql.Result, err error)
	Query(tx *sql.Tx, query string, args ...interface{}) (rows *sql.Rows, err error)
}

type PostgresConfiguration struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

func NewPostgresConfiguration(host string, port int, user string, password string, database string) (configuration *PostgresConfiguration) {
	configuration = &PostgresConfiguration{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		Database: database,
	}
	return
}

func NewDefaultPostgresConfiguration() (configuration *PostgresConfiguration) {
	configuration = &PostgresConfiguration{
		Host:     "localhost",
		Port:     5432,
		User:     "postgres",
		Password: "postgres",
		Database: "stori-challenge-db",
	}
	return
}

func NewPostgresConfigurationFromEnv() (configuration *PostgresConfiguration) {
	host := os.Getenv("POSTGRES_HOST")
	if host == "" {
		// return default configuration
		return NewDefaultPostgresConfiguration()
	}
	port, err := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	if err != nil {
		// return default configuration
		return NewDefaultPostgresConfiguration()
	}
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	database := os.Getenv("POSTGRES_DATABASE")

	configuration = &PostgresConfiguration{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		Database: database,
	}
	return
}

func (configuration *PostgresConfiguration) GetDataSourceName() (dataSourceName string) {
	dataSourceName = "host=" + configuration.Host
	dataSourceName += " port=" + strconv.Itoa(configuration.Port)
	dataSourceName += " user=" + configuration.User
	dataSourceName += " password=" + configuration.Password
	dataSourceName += " dbname=" + configuration.Database
	dataSourceName += " sslmode=disable"
	return
}
