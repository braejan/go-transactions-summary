package usecases

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	acEntity "github.com/braejan/go-transactions-summary/internal/domain/account/entity"
	acUsecases "github.com/braejan/go-transactions-summary/internal/domain/account/usecases"
	fileEntity "github.com/braejan/go-transactions-summary/internal/domain/file/entity"
	txEntity "github.com/braejan/go-transactions-summary/internal/domain/transaction/entity"
	userUsecases "github.com/braejan/go-transactions-summary/internal/domain/user/usecases"
	voFile "github.com/braejan/go-transactions-summary/internal/valueobject/file"
	voUser "github.com/braejan/go-transactions-summary/internal/valueobject/user"
)

// localFileUseCases struct implements the FileUseCases interface.
type localFileUseCases struct {
	userUseCases    userUsecases.UserUseCases
	accountUseCases acUsecases.AccountUseCases
}

// NewLocalFileUseCases returns a new localFileUseCases instance.
func NewLocalFileUseCases(accountUsecases acUsecases.AccountUseCases) (useCases *localFileUseCases, err error) {
	if accountUsecases == nil {
		err = voFile.ErrAccountUseCasesIsEmpty
		return
	}
	useCases = &localFileUseCases{
		accountUseCases: accountUsecases,
	}
	return
}

// ReadFile reads the file from the given path.
func (useCases *localFileUseCases) ReadFile(file *fileEntity.TxFile, isS3 bool) (txs []*txEntity.Transaction, err error) {
	err = useCases.CheckFile(file, isS3)
	if err != nil {
		txs = nil
		return
	}
	// Open the file. Omit validation previously done.
	osFile, _ := useCases.openOSFile(file.Path)
	defer osFile.Close()
	// Create a new reader.
	reader := csv.NewReader(osFile)
	// Read the file registers.
	txs, err = useCases.readFileRegisters(reader, file.Name)
	return
}

// CheckFile checks if is a valid structured file.
func (useCases *localFileUseCases) CheckFile(file *fileEntity.TxFile, isS3 bool) (err error) {
	if file == nil {
		err = voFile.ErrTxFileIsEmpty
		return
	}
	if !isS3 {
		osFile, err := useCases.openOSFile(file.Path)
		if err != nil {
			return err
		}
		defer osFile.Close()
	}
	return

}

func (useCases *localFileUseCases) openOSFile(path string) (file *os.File, err error) {
	// Validate the path.
	if path == "" {
		err = voFile.ErrFilePathIsEmpty
		return
	}
	// Open the file.
	file, err = os.Open(path)
	if err != nil {
		err = voFile.ErrFileCouldNotBeOpened
		return
	}
	return
}

func (useCases *localFileUseCases) readFileRegisters(reader *csv.Reader, fileName string) (txs []*txEntity.Transaction, err error) {
	if reader == nil {
		err = voFile.ErrFileReaderIsEmpty
		return
	}
	//Read the first line and ignore it.
	// TODO: Check if is a valid header.
	_, err = reader.Read()
	if err != nil {
		err = voFile.ErrFileCouldNotBeRead
		return
	}
	lineCounter := 1
	for {
		lineCounter++
		record, err := reader.Read()
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			txs = nil
			break
		}
		// Validate the line.
		userID, txDate, amount, err := useCases.checkValidLine(record)
		if err != nil {
			txs = nil
			break
		}
		err = useCases.checkUser(userID)
		if err != nil {
			txs = nil
			break
		}
		// Check if the account exists.
		acc, err := useCases.checkAccountByUserID(userID)
		if err != nil {
			txs = nil
			break
		}
		// Create the transaction entity and append it to the txs slice.
		tx, err := txEntity.NewTransaction(acc.ID, amount, txDate, fileName)
		if err != nil {
			txs = nil
			break
		}
		txs = append(txs, tx)
	}
	if txs == nil {
		log.Printf("Error reading the file %s at line %d", fileName, lineCounter)
	} else {
		log.Printf("File %s readed successfully", fileName)
	}
	return
}

func (useCases *localFileUseCases) checkValidLine(record []string) (id int64, txDate time.Time, amount float64, err error) {
	if len(record) != 3 {
		err = voFile.ErrFileLineIsInvalid
		return
	}
	// Validate the position 0 as a valid int64.
	id, err = strconv.ParseInt(record[0], 10, 64)
	if err != nil {
		err = voFile.ErrFileLineIsInvalid
		return
	}
	// Validate the position 1 as a valid date format "1/2".
	txDate, err = time.Parse("1/2", record[1])
	if err != nil {
		err = voFile.ErrFileLineIsInvalid
		return
	}
	// Validate the position 2 as a valid float64.
	amount, err = strconv.ParseFloat(record[2], 64)
	if err != nil {
		err = voFile.ErrFileLineIsInvalid
	}
	return
}

func (useCases *localFileUseCases) checkUser(ID int64) (err error) {
	// Check if the user exists.
	_, err = useCases.userUseCases.GetByID(ID)
	if err != nil && err == voUser.ErrUserNotFound {
		// Create a new user.
		err = useCases.userUseCases.Create(ID, fmt.Sprintf("User Name %d", ID), fmt.Sprintf("user.email%d@amazingemail.com", ID))
		if err != nil {
			return
		}
		_, err = useCases.userUseCases.GetByID(ID)
	} else {
		return
	}
	return

}

func (useCases *localFileUseCases) checkAccountByUserID(userID int64) (account *acEntity.Account, err error) {
	// Check if the account exists.
	account, err = useCases.accountUseCases.GetByUserID(userID)
	if err != nil && err == voUser.ErrUserNotFound {
		// Create a new account.
		err = useCases.accountUseCases.Create(userID)
		if err != nil {
			account = nil
			return
		}
		account, err = useCases.accountUseCases.GetByUserID(userID)
	} else if err != nil {
		account = nil
		return
	}
	return
}
