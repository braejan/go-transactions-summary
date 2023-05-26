package file

import "errors"

var (
	// ErrFilePathIsEmpty is the error returned when the file path is empty.
	ErrFilePathIsEmpty = errors.New("file path is empty")
	// ErrTxFileIsEmpty is the error returned when the file is empty.
	ErrTxFileIsEmpty = errors.New("file is empty")
	// ErrFileCouldNotBeOpened is the error returned when the file could not be opened.
	ErrFileCouldNotBeOpened = errors.New("file could not be opened")
	// ErrFileReaderIsEmpty is the error returned when the file reader is empty.
	ErrFileReaderIsEmpty = errors.New("file reader is empty")
	// ErrFileCouldNotBeRead is the error returned when the file could not be read.
	ErrFileCouldNotBeRead = errors.New("file could not be read")
	// ErrFileLineIsInvalid is the error returned when the file line is invalid.
	ErrFileLineIsInvalid = errors.New("file line is invalid")
	// ErrAccountUseCasesIsEmpty
	ErrAccountUseCasesIsEmpty = errors.New("account use cases is empty")
	// ErrFileIsEmpty is the error returned when the file is empty.
	ErrFileIsEmpty = errors.New("file is empty")
)
