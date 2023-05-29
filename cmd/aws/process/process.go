package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	apRepo "github.com/braejan/go-transactions-summary/internal/domain/account/repository/postgres"
	ucAccount "github.com/braejan/go-transactions-summary/internal/domain/account/usecases"
	fileEntity "github.com/braejan/go-transactions-summary/internal/domain/file/entity"
	ucFile "github.com/braejan/go-transactions-summary/internal/domain/file/usecases"
	txRepo "github.com/braejan/go-transactions-summary/internal/domain/transaction/repository/postgres"
	ucTx "github.com/braejan/go-transactions-summary/internal/domain/transaction/usecases"
	upRepo "github.com/braejan/go-transactions-summary/internal/domain/user/repository/postgres"
	ucUser "github.com/braejan/go-transactions-summary/internal/domain/user/usecases"
	"github.com/braejan/go-transactions-summary/internal/valueobject/postgres"
)

func handler(ctx context.Context, s3Event events.S3Event) (err error) {
	// Create a new AWS session using environment variables
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")),
	})
	if err != nil {
		err = fmt.Errorf("failed to create AWS session: %v", err)
	}
	for _, record := range s3Event.Records {
		bucket := record.S3.Bucket.Name
		key := record.S3.Object.Key
		// Create a downloader with the session and default options
		downloader := s3manager.NewDownloader(sess)
		// Create a file to write the S3 Object contents to.
		f, err := os.Create("/tmp/" + key)
		if err != nil {
			return fmt.Errorf("failed to create file %q, %v", key, err)
		}
		defer f.Close()
		path := f.Name()
		// Write the contents of S3 Object to the file
		n, err := downloader.Download(f, &s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(key),
		})
		if err != nil {
			return fmt.Errorf("failed to download file, %v", err)
		}

		fmt.Printf("file downloaded, %d bytes\n", n)
		f.Seek(0, 0)
		errProccess := handleFile(f, key, path)
		if errProccess != nil {
			return fmt.Errorf("failed to process file, %v", errProccess)
		}

	}
	return
}

func main() {
	lambda.Start(handler)
}

func handleFile(file *os.File, fileName string, path string) (err error) {
	// Create a postgres configuration from environment variables
	postgresConfig := postgres.NewPostgresConfigurationFromEnv()
	// Create a db Based on the configuration
	postgresDatabase := postgres.NewBasePostgresDatabase(postgresConfig)
	// Create a user repository
	userRepository := upRepo.NewPostgresUserRepository(postgresDatabase)
	// Create a account repository
	accountRepository := apRepo.NewPostgresAccountRepository(postgresDatabase)
	// Create a transaction repository
	transactionRepository := txRepo.NewPostgresTransactionRepository(postgresDatabase)
	// Create a user usecase
	userUsecase, err := ucUser.NewUserUseCases(userRepository)
	if err != nil {
		return
	}
	// Create a account usecase
	accountUsecase, err := ucAccount.NewAccountUseCases(accountRepository, userRepository)
	if err != nil {
		return
	}
	// Create a transaction usecase
	transactionUsecase, err := ucTx.NewTransactionUseCases(transactionRepository)
	if err != nil {
		return
	}
	// Create a file usecase
	fileUsecases, err := ucFile.NewFileUseCases(userUsecase, accountUsecase, transactionUsecase)
	if err != nil {
		return
	}
	txFile := fileEntity.NewTxFile(fileName, path, "s3hash", 0)
	err = fileUsecases.ProcessFile(*txFile, file)
	return
}
