package postgres

import (
	"context"
	"fmt"

	"github.com/braejan/go-transactions-summary/internal/domain/user/entity"
	"github.com/braejan/go-transactions-summary/internal/domain/user/repository"
	"github.com/braejan/go-transactions-summary/internal/valueobject/postgres"
	userErrors "github.com/braejan/go-transactions-summary/internal/valueobject/user"
	_ "github.com/lib/pq"
)

/*

 */
// postgresUserRepository struct implements the UserRepository interface using
// a PostgreSQL database.
type postgresUserRepository struct {
	configuration *postgres.PostgresConfiguration
	dbBase        postgres.PostgresDatabase
	repository.UserRepository
}

// NewPostgresUserRepository creates a new instance of postgresUserRepository.
func NewPostgresUserRepository(configuration *postgres.PostgresConfiguration) repository.UserRepository {
	return &postgresUserRepository{
		configuration: configuration,
		dbBase:        postgres.NewBasePostgresDatabase(),
	}
}

// GetByID returns a user by its ID.
const (
	getUserByID = `SELECT id, name, email FROM users WHERE id = $1`
)

func (postgresRepo *postgresUserRepository) GetByID(ID int64) (user *entity.User, err error) {
	db, err := postgresRepo.dbBase.Open(postgresRepo.configuration.GetDataSourceName())
	if err != nil {
		err = fmt.Errorf("%w:\n%w", postgres.ErrOpeningDatabase, err)
		return
	}
	defer postgresRepo.dbBase.Close(db)
	tx, err := db.BeginTx(context.Background(), nil)
	if err != nil {
		err = fmt.Errorf("%w:\n%w", postgres.ErrBeginningTransaction, err)
		return
	}
	row := tx.QueryRow(getUserByID, ID)
	user = &entity.User{}
	err = row.Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		user = nil
		err = fmt.Errorf("%w:\n%w", userErrors.ErrScanningUserRow, err)
		return
	}
	return
}

// GetByEmail returns a user by its email.
const (
	getUserByEmail = `SELECT id, name, email FROM users WHERE email = $1`
)

func (postgresRepo *postgresUserRepository) GetByEmail(email string) (user *entity.User, err error) {
	db, err := postgresRepo.dbBase.Open(postgresRepo.configuration.GetDataSourceName())
	if err != nil {
		err = fmt.Errorf("%w:\n%w", postgres.ErrOpeningDatabase, err)
		return
	}
	defer postgresRepo.dbBase.Close(db)
	tx, err := db.BeginTx(context.Background(), nil)
	if err != nil {
		err = fmt.Errorf("%w:\n%w", postgres.ErrBeginningTransaction, err)
		return
	}
	row, err := postgresRepo.dbBase.Query(tx, getUserByEmail, email)
	if err != nil {
		err = fmt.Errorf("%w:\n%w", userErrors.ErrScanningUserRow, err)
		return
	}
	user = &entity.User{}
	err = row.Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		user = nil
		err = fmt.Errorf("%w:\n%w", userErrors.ErrScanningUserRow, err)
		return
	}
	if user.ID == 0 {
		user = nil
		err = userErrors.ErrUserNotFound
	}
	return
}

// Create creates a new user.
const (
	createUser = `INSERT INTO users (id, name, email) VALUES ($1, $2, $3)`
)

func (postgresRepo *postgresUserRepository) Create(user *entity.User) (err error) {
	db, err := postgresRepo.dbBase.Open(postgresRepo.configuration.GetDataSourceName())
	if err != nil {
		err = fmt.Errorf("%w:\n%w", postgres.ErrOpeningDatabase, err)
		return
	}
	defer postgresRepo.dbBase.Close(db)
	tx, err := db.BeginTx(context.Background(), nil)
	if err != nil {
		err = fmt.Errorf("%w:\n%w", postgres.ErrBeginningTransaction, err)
		return
	}
	_, err = tx.Exec(createUser, user.ID, user.Name, user.Email)
	if err != nil {
		_ = postgresRepo.dbBase.Rollback(tx)
		err = fmt.Errorf("%w:\n%w", userErrors.ErrCreatingUser, err)
		return
	}
	_ = postgresRepo.dbBase.Commit(tx)
	return
}

// Update updates a user.
const (
	updateUser = `UPDATE users SET name = $1, email = $2 WHERE id = $3`
)

func (postgresRepo *postgresUserRepository) Update(user *entity.User) (err error) {
	db, err := postgresRepo.dbBase.Open(postgresRepo.configuration.GetDataSourceName())
	if err != nil {
		err = fmt.Errorf("%w:\n%w", postgres.ErrOpeningDatabase, err)
		return
	}
	defer postgresRepo.dbBase.Close(db)
	tx, err := db.BeginTx(context.Background(), nil)
	if err != nil {
		err = fmt.Errorf("%w:\n%w", postgres.ErrBeginningTransaction, err)
		return
	}
	_, err = postgresRepo.dbBase.Exec(tx, updateUser, user.Name, user.Email, user.ID)
	if err != nil {
		_ = postgresRepo.dbBase.Rollback(tx)
		err = fmt.Errorf("%w:\n%w", userErrors.ErrUpdatingUser, err)
		return
	}
	_ = postgresRepo.dbBase.Commit(tx)
	return
}
