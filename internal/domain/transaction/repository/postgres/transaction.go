package postgres

import (
	"database/sql"

	"github.com/braejan/go-transactions-summary/internal/domain/transaction/entity"
	"github.com/braejan/go-transactions-summary/internal/domain/transaction/repository"
	"github.com/braejan/go-transactions-summary/internal/valueobject/postgres"
	"github.com/braejan/go-transactions-summary/internal/valueobject/transaction"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

// postgresTransactionRepository is the postgres implementation of the transaction repository.
type postgresTransactionRepository struct {
	configuration *postgres.PostgresConfiguration
	baseDB        postgres.PostgresDatabase
	repository.TransactionRepository
}

// NewPostgresTransactionRepository creates a new instance of repository.TransactionRepository.
func NewPostgresTransactionRepository(
	configuration *postgres.PostgresConfiguration,
	baseDB postgres.PostgresDatabase) (transactionRepo repository.TransactionRepository, err error) {
	if configuration == nil {
		err = postgres.ErrNilConfiguration
		return
	}
	transactionRepo = &postgresTransactionRepository{
		configuration: configuration,
		baseDB:        baseDB,
	}
	return
}

// repository.TransactionRepository implementation.

// GetByID returns a transaction by its ID.
const (
	getTransactionByID = `SELECT id, accountid, amount, date, origin FROM transactions WHERE id = $1`
)

func (postgresRepo *postgresTransactionRepository) GetByID(ID uuid.UUID) (tx *entity.Transaction, err error) {
	db, err := postgresRepo.baseDB.Open(postgresRepo.configuration.GetDataSourceName())
	if err != nil {
		err = postgres.ErrOpeningDatabase
		return
	}
	defer postgresRepo.baseDB.Close(db)
	dbTx, err := postgresRepo.baseDB.BeginTx(db)
	if err != nil {
		err = postgres.ErrBeginningTransaction
		return
	}
	rows, err := postgresRepo.baseDB.Query(dbTx, getTransactionByID, ID)
	if err != nil {
		err = transaction.ErrQueryingTransactionByID
		return
	}
	txs, err := rows2Transactions(rows)
	if err != nil {
		err = transaction.ErrScanningTransactionByID
		return
	}
	// There should be only one transaction.
	tx = txs[0]
	return
}

// GetByAccountID returns all transactions for an account.
const (
	getTransactionsByAccountID = `SELECT id, accountid, amount, date, origin FROM transactions WHERE accountid = $1`
)

func (postgresRepo *postgresTransactionRepository) GetByAccountID(accountID uuid.UUID) (txs []*entity.Transaction, err error) {
	db, err := postgresRepo.baseDB.Open(postgresRepo.configuration.GetDataSourceName())
	if err != nil {
		err = postgres.ErrOpeningDatabase
		return
	}
	defer postgresRepo.baseDB.Close(db)
	dbTx, err := postgresRepo.baseDB.BeginTx(db)
	if err != nil {
		err = postgres.ErrBeginningTransaction
		return
	}
	rows, err := postgresRepo.baseDB.Query(dbTx, getTransactionsByAccountID, accountID)
	if err != nil {
		err = transaction.ErrQueryingTransactionsByAccountID
		return
	}
	txs, err = rows2Transactions(rows)
	return
}

// GetCreditsByAccountID returns the credits of an account.
const (
	getCreditsByAccountID = `SELECT id, accountid, amount, date, origin FROM transactions WHERE accountid = $1 AND operation = 'credit'`
)

func (postgresRepo *postgresTransactionRepository) GetCreditsByAccountID(accountID uuid.UUID) (txs []*entity.Transaction, err error) {
	db, err := postgresRepo.baseDB.Open(postgresRepo.configuration.GetDataSourceName())
	if err != nil {
		err = postgres.ErrOpeningDatabase
		return
	}
	defer postgresRepo.baseDB.Close(db)
	dbTx, err := postgresRepo.baseDB.BeginTx(db)
	if err != nil {
		err = postgres.ErrBeginningTransaction
		return
	}
	rows, err := postgresRepo.baseDB.Query(dbTx, getCreditsByAccountID, accountID)
	if err != nil {
		err = transaction.ErrQueryingCreditsByAccountID
		return
	}
	txs, err = rows2Transactions(rows)
	if err != nil {
		err = transaction.ErrScanningCreditsByAccountID
	}
	return
}

// GetDebitsByAccountID returns the debits of an account.
const (
	getDebitsByAccountID = `SELECT id, accountid, amount, date, origin FROM transactions WHERE accountid = $1 AND operation = 'debit'`
)

func (postgresRepo *postgresTransactionRepository) GetDebitsByAccountID(accountID uuid.UUID) (txs []*entity.Transaction, err error) {
	db, err := postgresRepo.baseDB.Open(postgresRepo.configuration.GetDataSourceName())
	if err != nil {
		err = postgres.ErrOpeningDatabase
		return
	}
	defer postgresRepo.baseDB.Close(db)
	dbTx, err := postgresRepo.baseDB.BeginTx(db)
	if err != nil {
		err = postgres.ErrBeginningTransaction
		return
	}
	rows, err := postgresRepo.baseDB.Query(dbTx, getDebitsByAccountID, accountID)
	if err != nil {
		err = transaction.ErrQueryingDebitsByAccountID
		return
	}
	txs, err = rows2Transactions(rows)
	return
}

// GetTransactionsByOrigin returns all transactions for an origin.
const (
	getTransactionsByOrigin = `SELECT id, accountid, amount, date, origin FROM transactions WHERE origin = $1`
)

func (postgresRepo *postgresTransactionRepository) GetTransactionsByOrigin(origin string) (txs []*entity.Transaction, err error) {
	if origin == "" {
		err = transaction.ErrEmptyOrigin
		return
	}
	db, err := postgresRepo.baseDB.Open(postgresRepo.configuration.GetDataSourceName())
	if err != nil {
		err = postgres.ErrOpeningDatabase
		return
	}
	defer postgresRepo.baseDB.Close(db)
	dbTx, err := postgresRepo.baseDB.BeginTx(db)
	if err != nil {
		err = postgres.ErrBeginningTransaction
		return
	}
	rows, err := postgresRepo.baseDB.Query(dbTx, getTransactionsByOrigin, origin)
	if err != nil {
		err = transaction.ErrQueryingTransactionsByOrigin
		return
	}
	txs, err = rows2Transactions(rows)
	return
}

// Create creates a new transaction.
const (
	createTransaction = `INSERT INTO transactions (id, accountid, amount, date, origin) VALUES ($1, $2, $3, $4, $5)`
)

func (postgresRepo *postgresTransactionRepository) Create(tx *entity.Transaction) (err error) {
	if tx == nil {
		err = transaction.ErrNilTransaction
		return
	}
	db, err := postgresRepo.baseDB.Open(postgresRepo.configuration.GetDataSourceName())
	if err != nil {
		err = postgres.ErrOpeningDatabase
		return
	}
	defer postgresRepo.baseDB.Close(db)
	dbTx, err := postgresRepo.baseDB.BeginTx(db)
	if err != nil {
		err = postgres.ErrBeginningTransaction
		return
	}
	_, err = postgresRepo.baseDB.Exec(dbTx, createTransaction, tx.ID, tx.AccountID, tx.Amount, tx.Date, tx.Origin)
	if err != nil {
		_ = postgresRepo.baseDB.Rollback(dbTx)
		err = transaction.ErrCreatingTransaction
		return
	}
	err = postgresRepo.baseDB.Commit(dbTx)
	return
}

func rows2Transactions(rows *sql.Rows) (txs []*entity.Transaction, err error) {
	for rows.Next() {
		tx := &entity.Transaction{}
		err = rows.Scan(&tx.ID, &tx.AccountID, &tx.Amount, &tx.Date, &tx.Origin)
		if err != nil {
			txs = nil
			err = transaction.ErrScanningTransactionsByAccountID
			return
		}
		txs = append(txs, tx)
	}
	return
}
