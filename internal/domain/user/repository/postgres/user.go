package postgres

import (
	"context"
	"log"

	"github.com/braejan/go-transactions-summary/internal/domain/user/entity"
	"github.com/braejan/go-transactions-summary/internal/domain/user/repository"
	"github.com/braejan/go-transactions-summary/internal/valueobject/postgres"
	userErrors "github.com/braejan/go-transactions-summary/internal/valueobject/user"
	_ "github.com/lib/pq"
)

// postgresUserRepository struct implements the UserRepository interface using
// a PostgreSQL database.
type postgresUserRepository struct {
	configuration *postgres.PostgresConfiguration
	baseDB        postgres.PostgresDatabase
	repository.UserRepository
}

// NewPostgresUserRepository creates a new instance of postgresUserRepository.
func NewPostgresUserRepository(configuration *postgres.PostgresConfiguration, baseDB postgres.PostgresDatabase) (userRepo repository.UserRepository, err error) {
	if configuration == nil {
		err = postgres.ErrNilConfiguration
		return
	}
	userRepo = &postgresUserRepository{
		configuration: configuration,
		baseDB:        baseDB,
	}
	return
}

// GetByID returns a user by its ID.
const (
	getUserByID = `SELECT id, name, email FROM users WHERE id = $1`
)

func (postgresRepo *postgresUserRepository) GetByID(ID int64) (user *entity.User, err error) {
	db, err := postgresRepo.baseDB.Open(postgresRepo.configuration.GetDataSourceName())
	if err != nil {
		err = postgres.ErrOpeningDatabase
		return
	}
	defer postgresRepo.baseDB.Close(db)
	tx, err := postgresRepo.baseDB.BeginTx(db)
	if err != nil {
		err = postgres.ErrBeginningTransaction
		return
	}
	rows, err := postgresRepo.baseDB.Query(tx, getUserByID, ID)
	if err != nil {
		err = userErrors.ErrQueryingUserByID
		return
	}
	user = &entity.User{}
	if rows.Next() {
		err = rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			log.Printf("error scanning user row: %v", err)
			user = nil
			err = userErrors.ErrScanningUserByID
			return
		}
	}
	if user.ID == 0 && user.Name == "" && user.Email == "" {
		user = nil
		err = userErrors.ErrUserNotFound
		return
	}
	return
}

// GetByEmail returns a user by its email.
const (
	getUserByEmail = `SELECT id, name, email FROM users WHERE email = $1`
)

func (postgresRepo *postgresUserRepository) GetByEmail(email string) (user *entity.User, err error) {
	db, err := postgresRepo.baseDB.Open(postgresRepo.configuration.GetDataSourceName())
	if err != nil {
		err = postgres.ErrOpeningDatabase
		return
	}
	defer postgresRepo.baseDB.Close(db)
	tx, err := db.BeginTx(context.Background(), nil)
	if err != nil {
		err = postgres.ErrBeginningTransaction
		return
	}
	rows, err := postgresRepo.baseDB.Query(tx, getUserByEmail, email)
	if err != nil {
		err = userErrors.ErrScanningUserRow
		return
	}
	user = &entity.User{}
	if rows.Next() {
		err = rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			user = nil
			err = userErrors.ErrScanningUserRow
			return
		}
	}
	if user.ID == 0 && user.Name == "" && user.Email == "" {
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
	db, err := postgresRepo.baseDB.Open(postgresRepo.configuration.GetDataSourceName())
	if err != nil {
		err = postgres.ErrOpeningDatabase
		return
	}
	defer postgresRepo.baseDB.Close(db)
	tx, err := db.BeginTx(context.Background(), nil)
	if err != nil {
		err = postgres.ErrBeginningTransaction
		return
	}
	_, err = tx.Exec(createUser, user.ID, user.Name, user.Email)
	if err != nil {
		_ = postgresRepo.baseDB.Rollback(tx)
		err = userErrors.ErrCreatingUser
		return
	}
	_ = postgresRepo.baseDB.Commit(tx)
	return
}

// Update updates a user.
const (
	updateUser = `UPDATE users SET name = $1, email = $2 WHERE id = $3`
)

func (postgresRepo *postgresUserRepository) Update(user *entity.User) (err error) {
	db, err := postgresRepo.baseDB.Open(postgresRepo.configuration.GetDataSourceName())
	if err != nil {
		err = postgres.ErrOpeningDatabase
		return
	}
	defer postgresRepo.baseDB.Close(db)
	tx, err := db.BeginTx(context.Background(), nil)
	if err != nil {
		err = postgres.ErrBeginningTransaction
		return
	}
	_, err = postgresRepo.baseDB.Exec(tx, updateUser, user.Name, user.Email, user.ID)
	if err != nil {
		_ = postgresRepo.baseDB.Rollback(tx)
		err = userErrors.ErrUpdatingUser
		return
	}
	_ = postgresRepo.baseDB.Commit(tx)
	return
}
