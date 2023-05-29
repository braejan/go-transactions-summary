package usecases

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"regexp"
	"strconv"
	"time"

	acEntity "github.com/braejan/go-transactions-summary/internal/domain/account/entity"
	acUsecases "github.com/braejan/go-transactions-summary/internal/domain/account/usecases"
	fileEntity "github.com/braejan/go-transactions-summary/internal/domain/file/entity"
	txEntity "github.com/braejan/go-transactions-summary/internal/domain/transaction/entity"
	txUsecases "github.com/braejan/go-transactions-summary/internal/domain/transaction/usecases"
	txUtil "github.com/braejan/go-transactions-summary/internal/domain/transaction/util"
	userUsecases "github.com/braejan/go-transactions-summary/internal/domain/user/usecases"
	voAccount "github.com/braejan/go-transactions-summary/internal/valueobject/account"
	voFile "github.com/braejan/go-transactions-summary/internal/valueobject/file"
	voTransaction "github.com/braejan/go-transactions-summary/internal/valueobject/transaction"
	voUser "github.com/braejan/go-transactions-summary/internal/valueobject/user"
)

// localFileUseCases struct implements the FileUseCases interface.
type localFileUseCases struct {
	userUseCases        userUsecases.UserUseCases
	accountUseCases     acUsecases.AccountUseCases
	transactionUseCases txUsecases.TransactionUseCases
}

// NewFileUseCases returns a new localFileUseCases instance.
func NewFileUseCases(
	userUseCases userUsecases.UserUseCases,
	accountUseCases acUsecases.AccountUseCases,
	transactionUseCases txUsecases.TransactionUseCases,
) (useCases FileUseCases, err error) {
	if userUseCases == nil {
		err = voUser.ErrNilUserUseCases
		return
	}
	if accountUseCases == nil {
		err = voAccount.ErrNilAccountUseCases
		return
	}
	if transactionUseCases == nil {
		err = voTransaction.ErrNilTransactionUseCases
		return
	}
	useCases = &localFileUseCases{
		userUseCases:        userUseCases,
		accountUseCases:     accountUseCases,
		transactionUseCases: transactionUseCases,
	}
	return
}

// ReadFile reads the file from the given path.
func (useCases *localFileUseCases) ReadAndProcessFile(file fileEntity.TxFile, isS3 bool) (err error) {
	err = useCases.CheckFile(file, isS3)
	if err != nil {
		return
	}
	// Open the file. Omit validation previously done.
	osFile, _ := useCases.openOSFile(file.Path)
	defer osFile.Close()
	// Create a new reader.
	reader := csv.NewReader(osFile)
	// Read the file registers.
	txsAux, err := useCases.readFileRegisters(reader, file.Name)
	if err != nil {
		return
	}
	txs := txUtil.ArrayTxMemoryToArrayValue(txsAux)
	err = useCases.createTransactions(txs)
	return
}

// CheckFile checks if is a valid structured file.
func (useCases *localFileUseCases) CheckFile(file fileEntity.TxFile, isS3 bool) (err error) {
	if !isS3 {
		osFile, err := useCases.openOSFile(file.Path)
		if err != nil {
			return err
		}
		defer osFile.Close()
	}
	return

}

// ProcessFile processes the file.
func (useCases *localFileUseCases) ProcessFile(file fileEntity.TxFile, osFile *os.File) (err error) {
	log.Println("Processing file:", file.Name)
	// Create a new reader.
	reader := csv.NewReader(osFile)
	// Read the file registers.
	log.Println("Reading file:", file.Name)
	txsAux, err := useCases.readFileRegisters(reader, file.Name)
	if err != nil {
		return
	}
	txs := txUtil.ArrayTxMemoryToArrayValue(txsAux)
	err = useCases.createTransactions(txs)
	return
}

// ProcessMultipartFile processes the file.
func (useCases *localFileUseCases) ProcessMultipartFile(txFile fileEntity.TxFile, file multipart.File) (err error) {
	log.Println("Processing multipart file:", txFile.Name)
	// Create a new reader.
	reader := csv.NewReader(file)
	// Read the file registers.
	log.Println("Reading file:", txFile.Name)
	txsAux, err := useCases.readFileRegisters(reader, txFile.Name)
	if err != nil {
		return
	}
	txs := txUtil.ArrayTxMemoryToArrayValue(txsAux)
	err = useCases.createTransactions(txs)
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
	//Read the first line and ignore it.
	// TODO: Check if is a valid header.
	_, err = reader.Read()
	if err != nil {
		if err == io.EOF {
			err = voFile.ErrFileIsEmpty
			return
		}
		err = voFile.ErrFileCouldNotBeRead
		return
	}
	lineCounter := 1
	for {
		log.Println("Reading line:", lineCounter)
		lineCounter++
		record, errRead := reader.Read()
		if errRead != nil && errRead == io.EOF {
			break
		} else if err != nil {
			txs = nil
			err = errRead
			break
		}
		// Validate the line.
		userID, txDate, amount, errCheck := useCases.checkValidLine(record)
		if errCheck != nil {
			txs = nil
			err = errCheck
			break
		}
		errCheck = useCases.checkUser(userID)
		if errCheck != nil {
			txs = nil
			err = errCheck
			break
		}
		// Check if the account exists.
		acc, errAcc := useCases.checkAccountByUserID(userID)
		if errAcc != nil {
			txs = nil
			err = errAcc
			break
		}
		// Create the transaction entity and append it to the txs slice.
		tx, errTx := txEntity.NewTransaction(acc.ID, amount, txDate, fileName)
		if errTx != nil {
			err = errTx
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
	regex := `^[-|+]+[0-9]+(\.[0-9]*)?$`
	match, _ := regexp.MatchString(regex, record[2])
	if !match {
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
	log.Println("Checking account for user:", userID)
	// Check if the account exists.
	accAux, err := useCases.accountUseCases.GetByUserID(userID)
	if err != nil && err == voAccount.ErrAccountNotFound {
		// Create a new account.
		err = useCases.accountUseCases.Create(userID)
		if err != nil {
			account = nil
			return
		}
		accAux, err = useCases.accountUseCases.GetByUserID(userID)
	} else if err != nil {
		account = nil
		return
	}
	account = &accAux
	return
}

func (useCases *localFileUseCases) createTransactions(txs []txEntity.Transaction) (err error) {
	// TODO: If process fails at some point, the file should processed partially.
	// Insert the transactions.
	for _, tx := range txs {
		err = useCases.transactionUseCases.Create(tx)
		if err != nil {
			return
		}
	}
	return
}
