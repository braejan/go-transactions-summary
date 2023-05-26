package usecases

import (
	fileEntity "github.com/braejan/go-transactions-summary/internal/domain/file/entity"
	txEntity "github.com/braejan/go-transactions-summary/internal/domain/transaction/entity"
)

// FileUseCases interface defines the file use cases.
type FileUseCases interface {
	// ReadFile reads the file from the given path or S3 bucket.
	ReadFile(file *fileEntity.TxFile) (txs []*txEntity.Transaction, err error)
	// CheckFile checks if is a valid structured file.
	CheckFile(file *fileEntity.TxFile) (err error)
}
