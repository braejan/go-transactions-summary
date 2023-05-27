package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	apRepo "github.com/braejan/go-transactions-summary/internal/domain/account/repository/postgres"
	ucAccount "github.com/braejan/go-transactions-summary/internal/domain/account/usecases"
	"github.com/braejan/go-transactions-summary/internal/domain/file/service/rest/file"
	ucFile "github.com/braejan/go-transactions-summary/internal/domain/file/usecases"
	txRepo "github.com/braejan/go-transactions-summary/internal/domain/transaction/repository/postgres"
	ucTx "github.com/braejan/go-transactions-summary/internal/domain/transaction/usecases"
	upRepo "github.com/braejan/go-transactions-summary/internal/domain/user/repository/postgres"
	ucUser "github.com/braejan/go-transactions-summary/internal/domain/user/usecases"
	"github.com/braejan/go-transactions-summary/internal/valueobject/postgres"
	"github.com/gorilla/mux"
)

var fileUsecases ucFile.FileUseCases

func init() {
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
	fataAnyErr(err)
	// Create a account usecase
	accountUsecase, err := ucAccount.NewAccountUseCases(accountRepository, userRepository)
	fataAnyErr(err)
	// Create a transaction usecase
	transactionUsecase, err := ucTx.NewTransactionUseCases(transactionRepository)
	fataAnyErr(err)
	// Create a file usecase
	fileUsecases, err = ucFile.NewFileUseCases(userUsecase, accountUsecase, transactionUsecase)
	fataAnyErr(err)

}

func main() {
	// Create context and register handlers
	ctx := context.Background()
	router := mux.NewRouter()
	fileHandler, err := file.NewFileHandler(fileUsecases)
	fataAnyErr(err)
	fileHandler.RegisterRoutes(router)
	// Create the server
	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	// Start server
	// start server
	go func() {
		log.Printf("🆙 starting server 🙎 file on port %s", server.Addr[1:])
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()
	// wait for SIGINT or SIGTERM signal
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)
	<-stopChan
	// shutdown server
	log.Println("🛑 shutting down server...")
	err = server.Shutdown(ctx)
	if err != nil {
		log.Fatal(err, "shutting down server")
	}

}

func fataAnyErr(err error) {
	if err != nil {
		panic(err)
	}
}
