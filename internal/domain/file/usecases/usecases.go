package usecases

import (
	"mime/multipart"
	"os"

	fileEntity "github.com/braejan/go-transactions-summary/internal/domain/file/entity"
)

// FileUseCases interface defines the file use cases.
type FileUseCases interface {
	// ReadFile reads the file from the given path or S3 bucket.
	ReadAndProcessFile(txFile fileEntity.TxFile, isS3 bool) (err error)
	// CheckFile checks if is a valid structured file.
	CheckFile(txFile fileEntity.TxFile, isS3 bool) (err error)
	// ProcessFile processes the file.
	ProcessFile(txFile fileEntity.TxFile, file *os.File) (err error)
	// ProcessMultipartFile processes the file.
	ProcessMultipartFile(txFile fileEntity.TxFile, file multipart.File) (err error)
}
