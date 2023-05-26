package usecases_test

import (
	"errors"
	"fmt"
	"os"
	"testing"

	acEntity "github.com/braejan/go-transactions-summary/internal/domain/account/entity"
	accMockUseCases "github.com/braejan/go-transactions-summary/internal/domain/account/usecases/mock"
	"github.com/braejan/go-transactions-summary/internal/domain/file/entity"
	"github.com/braejan/go-transactions-summary/internal/domain/file/usecases"
	txMockUseCases "github.com/braejan/go-transactions-summary/internal/domain/transaction/usecases/mock"
	userEntity "github.com/braejan/go-transactions-summary/internal/domain/user/entity"
	userMockUseCases "github.com/braejan/go-transactions-summary/internal/domain/user/usecases/mock"
	voAccount "github.com/braejan/go-transactions-summary/internal/valueobject/account"
	voFile "github.com/braejan/go-transactions-summary/internal/valueobject/file"
	voTransaction "github.com/braejan/go-transactions-summary/internal/valueobject/transaction"
	voUser "github.com/braejan/go-transactions-summary/internal/valueobject/user"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func getTestUsers() (users []*userEntity.User) {
	user := userEntity.NewUser(int64(0), "User Name 0", "user.email0@amazingemail.com")
	users = append(users, user)
	user = userEntity.NewUser(int64(1), "User Name 1", "user.email1@amazingemail.com")
	users = append(users, user)
	user = userEntity.NewUser(int64(2), "User Name 2", "user.email2@amazingemail.com")
	users = append(users, user)
	user = userEntity.NewUser(int64(3), "User Name 3", "user.email3@amazingemail.com")
	users = append(users, user)
	return
}

// TestNewLocalFileUseCasesWithNilUserUseCases tests the NewLocalFileUseCases function with a nil userUseCases parameter.
func TestNewLocalFileUseCasesWithNilUserUseCases(t *testing.T) {
	// Given a valid accountUseCases
	accountUseCases := accMockUseCases.NewMockAccountUseCases()
	// And a valid transactionUseCases
	transactionUseCases := txMockUseCases.NewMockTransactionUseCases()
	// When NewLocalFileUseCases is called with a nil userUseCases
	useCases, err := usecases.NewLocalFileUseCases(nil, accountUseCases, transactionUseCases)
	// Then the returned useCases should be nil
	assert.Nil(t, useCases)
	// And the returned error should be ErrNilUserUseCases
	assert.Equal(t, voUser.ErrNilUserUseCases, err)
}

// TestNewLocalFileUseCasesWithNilAccountUseCases tests the NewLocalFileUseCases function with a nil accountUseCases parameter.
func TestNewLocalFileUseCasesWithNilAccountUseCases(t *testing.T) {
	// Given a valid userUseCases
	userUseCases := userMockUseCases.NewMockUserUseCases()
	// And a valid transactionUseCases
	transactionUseCases := txMockUseCases.NewMockTransactionUseCases()
	// When NewLocalFileUseCases is called with a nil accountUseCases
	useCases, err := usecases.NewLocalFileUseCases(userUseCases, nil, transactionUseCases)
	// Then the returned useCases should be nil
	assert.Nil(t, useCases)
	// And the returned error should be ErrNilAccountUseCases
	assert.Equal(t, voAccount.ErrNilAccountUseCases, err)
}

// TestNewLocalFileUseCasesWithNilTransactionUseCases tests the NewLocalFileUseCases function with a nil transactionUseCases parameter.
func TestNewLocalFileUseCasesWithNilTransactionUseCases(t *testing.T) {
	// Given a valid userUseCases
	userUseCases := userMockUseCases.NewMockUserUseCases()
	// And a valid accountUseCases
	accountUseCases := accMockUseCases.NewMockAccountUseCases()
	// When NewLocalFileUseCases is called with a nil transactionUseCases
	useCases, err := usecases.NewLocalFileUseCases(userUseCases, accountUseCases, nil)
	// Then the returned useCases should be nil
	assert.Nil(t, useCases)
	// And the returned error should be ErrNilTransactionUseCases
	assert.Equal(t, voTransaction.ErrNilTransactionUseCases, err)
}

// TestNewLocalUseCasesSuccess tests the NewLocalFileUseCases function with valid parameters.
func TestNewLocalUseCasesSuccess(t *testing.T) {
	// Given a valid userUseCases
	userUseCases := userMockUseCases.NewMockUserUseCases()
	// And a valid accountUseCases
	accountUseCases := accMockUseCases.NewMockAccountUseCases()
	// And a valid transactionUseCases
	transactionUseCases := txMockUseCases.NewMockTransactionUseCases()
	// When NewLocalFileUseCases is called with valid parameters
	useCases, err := usecases.NewLocalFileUseCases(userUseCases, accountUseCases, transactionUseCases)
	// Then the returned useCases should not be nil
	assert.NotNil(t, useCases)
	// And the returned error should be nil
	assert.Nil(t, err)
}

// TestReadAndProcessFileWithEmtyFile tests the ReadAndProcessFile function with an empty file entity.
func TestReadAndProcessFileWithEmtyFile(t *testing.T) {
	// Given a valid userUseCases
	userUseCases := userMockUseCases.NewMockUserUseCases()
	// And a valid accountUseCases
	accountUseCases := accMockUseCases.NewMockAccountUseCases()
	// And a valid transactionUseCases
	transactionUseCases := txMockUseCases.NewMockTransactionUseCases()
	// And a valid useCases
	useCases, _ := usecases.NewLocalFileUseCases(userUseCases, accountUseCases, transactionUseCases)
	// When ReadAndProcessFile is called with an empty file entity
	err := useCases.ReadAndProcessFile(entity.TxFile{}, false)
	// Then the returned error should be ErrFilePathIsEmpty
	assert.Equal(t, voFile.ErrFilePathIsEmpty, err)
}

// TestReadAndProcessFileWithNonExistingFile tests the ReadAndProcessFile function with a non existing file entity.
func TestReadAndProcessFileWithNonExistingFile(t *testing.T) {
	// Given a valid userUseCases
	userUseCases := userMockUseCases.NewMockUserUseCases()
	// And a valid accountUseCases
	accountUseCases := accMockUseCases.NewMockAccountUseCases()
	// And a valid transactionUseCases
	transactionUseCases := txMockUseCases.NewMockTransactionUseCases()
	// And a valid useCases
	useCases, _ := usecases.NewLocalFileUseCases(userUseCases, accountUseCases, transactionUseCases)
	// When ReadAndProcessFile is called with a non existing file entity
	err := useCases.ReadAndProcessFile(entity.TxFile{Path: "non-existing-file"}, false)
	// Then the returned error should be ErrFileCouldNotBeOpened
	assert.Equal(t, voFile.ErrFileCouldNotBeOpened, err)
}

// TestReadAndProcessFileWithEmptyFile tests the ReadAndProcessFile function with an empty file entity.
func TestReadAndProcessFileWithEmptyFile(t *testing.T) {
	// Given a valid userUseCases
	userUseCases := userMockUseCases.NewMockUserUseCases()
	// And a valid accountUseCases
	accountUseCases := accMockUseCases.NewMockAccountUseCases()
	// And a valid transactionUseCases
	transactionUseCases := txMockUseCases.NewMockTransactionUseCases()
	// And a valid useCases
	useCases, _ := usecases.NewLocalFileUseCases(userUseCases, accountUseCases, transactionUseCases)
	// And a valid file entity
	currentDir, _ := os.Getwd()
	filePath := fmt.Sprintf("%s/%s", currentDir, "test/files/txns_empty.csv")
	fmt.Printf("filePath: %s\n", filePath)
	fileEntity := entity.NewTxFile("txns_empty.csv", filePath, uuid.New().String(), 0)
	// When ReadAndProcessFile is called with an empty file entity
	err := useCases.ReadAndProcessFile(*fileEntity, false)
	// Then the returned error should be ErrFileIsEmpty
	assert.Equal(t, voFile.ErrFileIsEmpty, err)
}

// TestReadAndProcessFileWithInvalidColumns tests the ReadAndProcessFile function with an invalid file entry columns.
func TestReadAndProcessFileWithInvalidColumns(t *testing.T) {
	// Given a valid userUseCases
	userUseCases := userMockUseCases.NewMockUserUseCases()
	// And a valid accountUseCases
	accountUseCases := accMockUseCases.NewMockAccountUseCases()
	// And a valid transactionUseCases
	transactionUseCases := txMockUseCases.NewMockTransactionUseCases()
	// And a valid useCases
	useCases, _ := usecases.NewLocalFileUseCases(userUseCases, accountUseCases, transactionUseCases)
	// And a valid file entity
	currentDir, _ := os.Getwd()
	filePath := fmt.Sprintf("%s/%s", currentDir, "test/files/txns_invalid_columns.csv")
	fmt.Printf("filePath: %s\n", filePath)
	fileEntity := entity.NewTxFile("txns_invalid.csv", filePath, uuid.New().String(), 0)
	// When ReadAndProcessFile is called with an invalid file entity
	err := useCases.ReadAndProcessFile(*fileEntity, false)
	// Then the returned error should be ErrFileLineIsInvalid
	assert.Equal(t, voFile.ErrFileLineIsInvalid, err)
}

// TestReadAndProcessFileWithInvalidID tests the ReadAndProcessFile function with an invalid file entry id.
func TestReadAndProcessFileWithInvalidID(t *testing.T) {
	// Given a valid userUseCases
	userUseCases := userMockUseCases.NewMockUserUseCases()
	// And a valid accountUseCases
	accountUseCases := accMockUseCases.NewMockAccountUseCases()
	// And a valid transactionUseCases
	transactionUseCases := txMockUseCases.NewMockTransactionUseCases()
	// And a valid useCases
	useCases, _ := usecases.NewLocalFileUseCases(userUseCases, accountUseCases, transactionUseCases)
	// And a valid file entity
	currentDir, _ := os.Getwd()
	filePath := fmt.Sprintf("%s/%s", currentDir, "test/files/txns_invalid_id.csv")
	fmt.Printf("filePath: %s\n", filePath)
	fileEntity := entity.NewTxFile("txns_invalid.csv", filePath, uuid.New().String(), 0)
	// When ReadAndProcessFile is called with an invalid file entity
	err := useCases.ReadAndProcessFile(*fileEntity, false)
	// Then the returned error should be ErrFileLineIsInvalid
	assert.Equal(t, voFile.ErrFileLineIsInvalid, err)
}

// TestReadAndProcessFileWithInvalidDate tests the ReadAndProcessFile function with an invalid file entry date.
func TestReadAndProcessFileWithInvalidDate(t *testing.T) {
	// Given a valid userUseCases
	userUseCases := userMockUseCases.NewMockUserUseCases()
	// And a valid accountUseCases
	accountUseCases := accMockUseCases.NewMockAccountUseCases()
	// And a valid transactionUseCases
	transactionUseCases := txMockUseCases.NewMockTransactionUseCases()
	// And a valid useCases
	useCases, _ := usecases.NewLocalFileUseCases(userUseCases, accountUseCases, transactionUseCases)
	// And a valid file entity
	currentDir, _ := os.Getwd()
	filePath := fmt.Sprintf("%s/%s", currentDir, "test/files/txns_invalid_date.csv")
	fmt.Printf("filePath: %s\n", filePath)
	fileEntity := entity.NewTxFile("txns_invalid.csv", filePath, uuid.New().String(), 0)
	// When ReadAndProcessFile is called with an invalid file entity
	err := useCases.ReadAndProcessFile(*fileEntity, false)
	// Then the returned error should be ErrFileLineIsInvalid
	assert.Equal(t, voFile.ErrFileLineIsInvalid, err)
}

// TestReadAndProcessFileWithInvalidAmount tests the ReadAndProcessFile function with an invalid file entry amount.
func TestReadAndProcessFileWithInvalidAmount(t *testing.T) {
	// Given a valid userUseCases
	userUseCases := userMockUseCases.NewMockUserUseCases()
	// And a valid accountUseCases
	accountUseCases := accMockUseCases.NewMockAccountUseCases()
	// And a valid transactionUseCases
	transactionUseCases := txMockUseCases.NewMockTransactionUseCases()
	// And a valid useCases
	useCases, _ := usecases.NewLocalFileUseCases(userUseCases, accountUseCases, transactionUseCases)
	// And a valid file entity
	currentDir, _ := os.Getwd()
	filePath := fmt.Sprintf("%s/%s", currentDir, "test/files/txns_invalid_amount.csv")
	fmt.Printf("filePath: %s\n", filePath)
	fileEntity := entity.NewTxFile("txns_invalid.csv", filePath, uuid.New().String(), 0)
	// When ReadAndProcessFile is called with an invalid file entity
	err := useCases.ReadAndProcessFile(*fileEntity, false)
	// Then the returned error should be ErrFileLineIsInvalid
	assert.Equal(t, voFile.ErrFileLineIsInvalid, err)
}

// TestReadAndProcessErrGettingUserByID tests the ReadAndProcessFile function with an error getting the user by id.
func TestReadAndProcessErrGettingUserByID(t *testing.T) {
	// Given a valid userUseCases
	userUseCases := userMockUseCases.NewMockUserUseCases()
	// And a valid accountUseCases
	accountUseCases := accMockUseCases.NewMockAccountUseCases()
	// And a valid transactionUseCases
	transactionUseCases := txMockUseCases.NewMockTransactionUseCases()
	userUseCases.On("GetByID", mock.Anything).Return(nil, errors.New("error getting user by id"))
	// And a valid useCases
	useCases, _ := usecases.NewLocalFileUseCases(userUseCases, accountUseCases, transactionUseCases)
	// And a valid file entity
	currentDir, _ := os.Getwd()
	filePath := fmt.Sprintf("%s/%s", currentDir, "test/files/txns_simple.csv")
	fmt.Printf("filePath: %s\n", filePath)
	fileEntity := entity.NewTxFile("txns.csv", filePath, uuid.New().String(), 0)
	// When ReadAndProcessFile is called with an invalid file entity
	err := useCases.ReadAndProcessFile(*fileEntity, false)
	// Then the returned error should be not nil
	assert.NotNil(t, err)
}

// TestReadAndProcessErrCreatingUser tests the ReadAndProcessFile function with an error creating the user.
func TestReadAndProcessErrCreatingUser(t *testing.T) {
	// Given a valid userUseCases
	userUseCases := userMockUseCases.NewMockUserUseCases()
	userUseCases.On("GetByID", mock.Anything).Return(nil, voUser.ErrUserNotFound)
	userUseCases.On("Create", int64(0), "User Name 0", "user.email0@amazingemail.com").Return(nil, errors.New("error creating user"))
	// And a valid accountUseCases
	accountUseCases := accMockUseCases.NewMockAccountUseCases()
	// And a valid transactionUseCases
	transactionUseCases := txMockUseCases.NewMockTransactionUseCases()

	// And a valid useCases
	useCases, _ := usecases.NewLocalFileUseCases(userUseCases, accountUseCases, transactionUseCases)

	// And a valid file entity
	currentDir, _ := os.Getwd()
	filePath := fmt.Sprintf("%s/%s", currentDir, "test/files/txns_simple.csv")
	fmt.Printf("filePath: %s\n", filePath)
	fileEntity := entity.NewTxFile("txns.csv", filePath, uuid.New().String(), 0)

	// When ReadAndProcessFile is called with an invalid file entity
	err := useCases.ReadAndProcessFile(*fileEntity, false)

	// Then the returned error should be not nil
	assert.NotNil(t, err)
}

// TestReadAndProcessErrCheckAccountByUserID_GetByUserID tests the ReadAndProcessFile function with an error checking the account by user id.
func TestReadAndProcessErrCheckAccountByUserID_GetByUserID(t *testing.T) {
	// Given a valid userUseCases
	userUseCases := userMockUseCases.NewMockUserUseCases()
	userUseCases.On("GetByID", mock.Anything).Return(nil, voUser.ErrUserNotFound).Once()
	userUseCases.On("Create", int64(0), "User Name 0", "user.email0@amazingemail.com").Return(nil)
	// And a valid user with id 0
	user := userEntity.NewUser(0, "User Name 0", "user.email0@amazingemail.com")
	userUseCases.On("GetByID", mock.Anything).Return(*user, nil)
	// And a valid accountUseCases
	accountUseCases := accMockUseCases.NewMockAccountUseCases()
	accountUseCases.On("GetByUserID", int64(0)).Return(acEntity.Account{}, errors.New("error checking account by user id"))
	// And a valid transactionUseCases
	transactionUseCases := txMockUseCases.NewMockTransactionUseCases()
	// And a valid useCases
	useCases, _ := usecases.NewLocalFileUseCases(userUseCases, accountUseCases, transactionUseCases)
	// And a valid file entity
	currentDir, _ := os.Getwd()
	filePath := fmt.Sprintf("%s/%s", currentDir, "test/files/txns_simple.csv")
	fmt.Printf("filePath: %s\n", filePath)
	fileEntity := entity.NewTxFile("txns.csv", filePath, uuid.New().String(), 0)
	// When ReadAndProcessFile is called with an invalid file entity
	err := useCases.ReadAndProcessFile(*fileEntity, false)

	// Then the returned error should be not nil
	assert.NotNil(t, err)
}

// TestReadAndProcessErrCheckAccountByUserID_Create tests the ReadAndProcessFile function with an error checking the account by user id.
func TestReadAndProcessErrCheckAccountByUserID_Create(t *testing.T) {
	// Given a valid userUseCases
	userUseCases := userMockUseCases.NewMockUserUseCases()
	userUseCases.On("GetByID", mock.Anything).Return(nil, voUser.ErrUserNotFound).Once()
	userUseCases.On("Create", int64(0), "User Name 0", "user.email0@amazingemail.com").Return(nil)
	// And a valid user with id 0
	user := userEntity.NewUser(0, "User Name 0", "user.email0@amazingemail.com")
	userUseCases.On("GetByID", mock.Anything).Return(*user, nil)
	// And a valid accountUseCases
	accountUseCases := accMockUseCases.NewMockAccountUseCases()
	accountUseCases.On("GetByUserID", int64(0)).Return(acEntity.Account{}, voAccount.ErrAccountNotFound).Once()
	accountUseCases.On("Create", int64(0)).Return(errors.New("error checking account by user id"))
	// And a valid transactionUseCases
	transactionUseCases := txMockUseCases.NewMockTransactionUseCases()
	// And a valid useCases
	useCases, _ := usecases.NewLocalFileUseCases(userUseCases, accountUseCases, transactionUseCases)
	// And a valid file entity
	currentDir, _ := os.Getwd()
	filePath := fmt.Sprintf("%s/%s", currentDir, "test/files/txns_simple.csv")
	fmt.Printf("filePath: %s\n", filePath)
	fileEntity := entity.NewTxFile("txns.csv", filePath, uuid.New().String(), 0)
	// When ReadAndProcessFile is called with an invalid file entity
	err := useCases.ReadAndProcessFile(*fileEntity, false)
	// Then the returned error should be not nil
	assert.NotNil(t, err)
}

// TestReadAndProcessErrCheckAccountByUserID_2GetByUserID tests the ReadAndProcessFile function with an error checking the account by user id.
func TestReadAndProcessErrCheckAccountByUserID_2GetByUserID(t *testing.T) {
	// Given a valid userUseCases
	userUseCases := userMockUseCases.NewMockUserUseCases()
	userUseCases.On("GetByID", mock.Anything).Return(nil, voUser.ErrUserNotFound).Once()
	userUseCases.On("Create", int64(0), "User Name 0", "user.email0@amazingemail.com").Return(nil)
	// And a valid user with id 0
	user := userEntity.NewUser(0, "User Name 0", "user.email0@amazingemail.com")
	userUseCases.On("GetByID", mock.Anything).Return(*user, nil)
	// And a valid accountUseCases
	accountUseCases := accMockUseCases.NewMockAccountUseCases()
	accountUseCases.On("GetByUserID", int64(0)).Return(acEntity.Account{}, voAccount.ErrAccountNotFound).Once()
	accountUseCases.On("Create", int64(0)).Return(nil)
	accountUseCases.On("GetByUserID", int64(0)).Return(acEntity.Account{}, voAccount.ErrQueryingAccountByUserID).Once()
	// And a valid transactionUseCases
	transactionUseCases := txMockUseCases.NewMockTransactionUseCases()
	// And a valid useCases
	useCases, _ := usecases.NewLocalFileUseCases(userUseCases, accountUseCases, transactionUseCases)
	// And a valid file entity
	currentDir, _ := os.Getwd()
	filePath := fmt.Sprintf("%s/%s", currentDir, "test/files/txns_simple.csv")
	fileEntity := entity.NewTxFile("txns.csv", filePath, uuid.New().String(), 0)
	// When ReadAndProcessFile is called with an invalid file entity
	err := useCases.ReadAndProcessFile(*fileEntity, false)
	// Then the returned error should be ErrQueryingAccountByUserID
	assert.NotNil(t, err)
	assert.Equal(t, voAccount.ErrQueryingAccountByUserID, err)
}

// TestReadAndProcessErrCreatingTransactions tests the ReadAndProcessFile function with an error creating transactions.
func TestReadAndProcessErrCreatingTransactions(t *testing.T) {
	// Given a valid user array
	users := getTestUsers()
	// Given a valid userUseCases
	userUseCases := userMockUseCases.NewMockUserUseCases()
	// And a valid accountUseCases
	accountUseCases := accMockUseCases.NewMockAccountUseCases()
	for _, user := range users {
		userUseCases.On("Create", user.ID, user.Name, user.Email).Return(nil)
		userUseCases.On("GetByID", user.ID).Return(*user, nil)
		// And a valid user account
		account := acEntity.NewAccount(user.ID)
		accountUseCases.On("GetByUserID", user.ID).Return(*account, nil)
	}
	// And a valid transactionUseCases
	transactionUseCases := txMockUseCases.NewMockTransactionUseCases()
	transactionUseCases.On("Create", mock.Anything).Return(errors.New("error creating transaction"))
	// And a valid useCases
	useCases, _ := usecases.NewLocalFileUseCases(userUseCases, accountUseCases, transactionUseCases)
	// And a valid file entity
	currentDir, _ := os.Getwd()
	filePath := fmt.Sprintf("%s/%s", currentDir, "test/files/txns_simple.csv")
	fileEntity := entity.NewTxFile("txns.csv", filePath, uuid.New().String(), 0)
	// When ReadAndProcessFile is called with an invalid file entity
	err := useCases.ReadAndProcessFile(*fileEntity, false)
	// Then the returned error should be ErrQueryingAccountByUserID
	assert.NotNil(t, err)
	assert.Equal(t, voTransaction.ErrCreatingTransaction, err)
}

// TestReadAndProcessSucess tests the ReadAndProcessFile function with success.
func TestReadAndProcessSucess(t *testing.T) {
	// Given a valid user array
	users := getTestUsers()
	// Given a valid userUseCases
	userUseCases := userMockUseCases.NewMockUserUseCases()
	// And a valid accountUseCases
	accountUseCases := accMockUseCases.NewMockAccountUseCases()
	for _, user := range users {
		userUseCases.On("Create", user.ID, user.Name, user.Email).Return(nil)
		userUseCases.On("GetByID", user.ID).Return(*user, nil)
		// And a valid user account
		account := acEntity.NewAccount(user.ID)
		accountUseCases.On("GetByUserID", user.ID).Return(*account, nil)
	}
	// And a valid transactionUseCases
	transactionUseCases := txMockUseCases.NewMockTransactionUseCases()
	transactionUseCases.On("Create", mock.Anything).Return(nil)
	// And a valid useCases
	useCases, _ := usecases.NewLocalFileUseCases(userUseCases, accountUseCases, transactionUseCases)
	// And a valid file entity
	currentDir, _ := os.Getwd()
	filePath := fmt.Sprintf("%s/%s", currentDir, "test/files/txns_simple.csv")
	fileEntity := entity.NewTxFile("txns.csv", filePath, uuid.New().String(), 0)
	// When ReadAndProcessFile is called with an invalid file entity
	err := useCases.ReadAndProcessFile(*fileEntity, false)
	// Then the returned error should be nil
	assert.Nil(t, err)
}

// TestReadAndProcessErrFileLastRecord tests the ReadAndProcessFile function with an error getting the last record of the file.
func TestReadAndProcessErrFileLastRecord(t *testing.T) {
	// Given a valid user array
	users := getTestUsers()
	// Given a valid userUseCases
	userUseCases := userMockUseCases.NewMockUserUseCases()
	// And a valid accountUseCases
	accountUseCases := accMockUseCases.NewMockAccountUseCases()
	for _, user := range users {
		userUseCases.On("Create", user.ID, user.Name, user.Email).Return(nil)
		userUseCases.On("GetByID", user.ID).Return(*user, nil)
		// And a valid user account
		account := acEntity.NewAccount(user.ID)
		accountUseCases.On("GetByUserID", user.ID).Return(*account, nil)
	}
	// And a valid transactionUseCases
	transactionUseCases := txMockUseCases.NewMockTransactionUseCases()
	transactionUseCases.On("Create", mock.Anything).Return(nil)
	// And a valid useCases
	useCases, _ := usecases.NewLocalFileUseCases(userUseCases, accountUseCases, transactionUseCases)
	// And a valid file entity
	currentDir, _ := os.Getwd()
	filePath := fmt.Sprintf("%s/%s", currentDir, "test/files/txns_invalid_last_record.csv")
	fileEntity := entity.NewTxFile("txns.csv", filePath, uuid.New().String(), 0)
	// When ReadAndProcessFile is called with an invalid file entity
	err := useCases.ReadAndProcessFile(*fileEntity, false)
	// Then the returned error should be not nil
	assert.NotNil(t, err)
}
