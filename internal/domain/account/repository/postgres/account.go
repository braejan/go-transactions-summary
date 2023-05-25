package postgres

import (
	"github.com/braejan/go-transactions-summary/internal/domain/account/entity"
	"github.com/braejan/go-transactions-summary/internal/domain/account/repository"
	"github.com/braejan/go-transactions-summary/internal/valueobject/account"
	"github.com/braejan/go-transactions-summary/internal/valueobject/postgres"
	"github.com/google/uuid"
)

// postgresAccountRepository struct implements the AccountRepository interface using
// a PostgreSQL database.
type postgresAccountRepository struct {
	configuration *postgres.PostgresConfiguration
	baseDB        postgres.PostgresDatabase
	repository.AccountRepository
}

// NewPostgresAccountRepository creates a new instance of postgresAccountRepository.
func NewPostgresAccountRepository(
	configuration *postgres.PostgresConfiguration,
	baseDB postgres.PostgresDatabase) (accountRepo repository.AccountRepository, err error) {
	if configuration == nil {
		err = postgres.ErrNilConfiguration
		return
	}
	accountRepo = &postgresAccountRepository{
		configuration: configuration,
		baseDB:        baseDB,
	}
	return
}

// GetByID returns an account by its ID.
const (
	getAccountByID = `SELECT id, balance, userid, active FROM accounts WHERE id = $1`
)

func (postgresRepo *postgresAccountRepository) GetByID(ID uuid.UUID) (acc *entity.Account, err error) {
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
	rows, err := postgresRepo.baseDB.Query(tx, getAccountByID, ID)
	if err != nil {
		err = account.ErrQueryingAccountByID
		return
	}
	acc = &entity.Account{}
	if rows.Next() {
		err = rows.Scan(&acc.ID, &acc.Balance, &acc.UserID, &acc.Active)
		if err != nil {
			err = account.ErrScanningAccountByID
			return
		}
	}
	if acc.UserID == 0 {
		err = account.ErrAccountNotFound
	}
	return
}

// GetByUserID returns an account by its user ID.
const (
	getAccountByUserID = `SELECT id, balance, userid, active FROM accounts WHERE userid = $1`
)

func (postgresRepo *postgresAccountRepository) GetByUserID(userID int64) (acc *entity.Account, err error) {
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
	rows, err := postgresRepo.baseDB.Query(tx, getAccountByUserID, userID)
	if err != nil {
		err = account.ErrQueryingAccountByUserID
		return
	}
	acc = &entity.Account{}
	if rows.Next() {
		err = rows.Scan(&acc.ID, &acc.Balance, &acc.UserID, &acc.Active)
		if err != nil {
			err = account.ErrScanningAccountByUserID
			return
		}
	}
	if acc.UserID == 0 {
		err = account.ErrAccountNotFound
	}
	return
}

// Create creates a new account.
const (
	createAccount = `INSERT INTO accounts (id, balance, userid, active) VALUES ($1, $2, $3, $4)`
)

func (postgresRepo *postgresAccountRepository) Create(acc *entity.Account) (err error) {
	if acc == nil {
		err = account.ErrNilAccount
		return
	}
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
	_, err = postgresRepo.baseDB.Exec(tx, createAccount, acc.ID, acc.Balance, acc.UserID, acc.Active)
	if err != nil {
		_ = postgresRepo.baseDB.Rollback(tx)
		err = account.ErrCreatingAccount
		return
	}
	err = postgresRepo.baseDB.Commit(tx)
	return
}

// Update updates an account.
const (
	updateAccount = `UPDATE accounts SET balance = $1, active = $2 WHERE id = $3`
)

func (postgresRepo *postgresAccountRepository) Update(acc *entity.Account) (err error) {
	if acc == nil {
		err = account.ErrNilAccount
		return
	}
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
	_, err = postgresRepo.baseDB.Exec(tx, updateAccount, acc.Balance, acc.Active, acc.ID)
	if err != nil {
		_ = postgresRepo.baseDB.Rollback(tx)
		err = account.ErrUpdatingAccount
	}
	err = postgresRepo.baseDB.Commit(tx)
	return
}
