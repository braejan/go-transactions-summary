package postgres

import (
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
	baseDB postgres.PostgresDatabase
	repository.UserRepository
}

// NewPostgresUserRepository creates a new instance of postgresUserRepository.
func NewPostgresUserRepository(baseDB postgres.PostgresDatabase) (userRepo repository.UserRepository) {
	userRepo = &postgresUserRepository{
		baseDB: baseDB,
	}
	return
}

// GetByID returns a user by its ID.
const (
	getUserByID = `SELECT id, name, email FROM users WHERE id = $1`
)

func (postgresRepo *postgresUserRepository) GetByID(ID int64) (user *entity.User, err error) {
	db, err := postgresRepo.baseDB.Open()
	if err != nil {
		err = postgres.ErrOpeningDatabase
		return
	}
	defer postgresRepo.baseDB.Close(db)
	tx, err := postgresRepo.baseDB.BeginTx(db)
	defer postgresRepo.baseDB.Rollback(tx)
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
	db, err := postgresRepo.baseDB.Open()
	if err != nil {
		err = postgres.ErrOpeningDatabase
		return
	}
	defer postgresRepo.baseDB.Close(db)
	tx, err := postgresRepo.baseDB.BeginTx(db)
	defer postgresRepo.baseDB.Rollback(tx)
	if err != nil {
		err = postgres.ErrBeginningTransaction
		return
	}
	rows, err := postgresRepo.baseDB.Query(tx, getUserByEmail, email)
	if err != nil {
		err = userErrors.ErrQueryingUserByEmail
		return
	}
	user = &entity.User{}
	if rows.Next() {
		err = rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			user = nil
			err = userErrors.ErrScanningUserByEmail
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
	if user == nil {
		err = userErrors.ErrNilUser
		return
	}
	db, err := postgresRepo.baseDB.Open()
	if err != nil {
		err = postgres.ErrOpeningDatabase
		return
	}
	defer postgresRepo.baseDB.Close(db)
	tx, err := postgresRepo.baseDB.BeginTx(db)
	defer postgresRepo.baseDB.Rollback(tx)
	if err != nil {
		err = postgres.ErrBeginningTransaction
		return
	}
	_, err = postgresRepo.baseDB.Exec(tx, createUser, user.ID, user.Name, user.Email)
	if err != nil {
		_ = postgresRepo.baseDB.Rollback(tx)
		err = userErrors.ErrCreatingUser
		return
	}
	err = postgresRepo.baseDB.Commit(tx)
	return
}

// Update updates a user.
const (
	updateUser = `UPDATE users SET name = $1, email = $2 WHERE id = $3`
)

func (postgresRepo *postgresUserRepository) Update(user *entity.User) (err error) {
	if user == nil {
		err = userErrors.ErrNilUser
		return
	}
	db, err := postgresRepo.baseDB.Open()
	if err != nil {
		err = postgres.ErrOpeningDatabase
		return
	}
	defer postgresRepo.baseDB.Close(db)
	tx, err := postgresRepo.baseDB.BeginTx(db)
	defer postgresRepo.baseDB.Rollback(tx)
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
	err = postgresRepo.baseDB.Commit(tx)
	return
}
