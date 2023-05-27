package file

import (
	"log"
	"net/http"

	"github.com/braejan/go-transactions-summary/internal/domain/file/entity"
	"github.com/braejan/go-transactions-summary/internal/domain/file/usecases"
	voFile "github.com/braejan/go-transactions-summary/internal/valueobject/file"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type FileHandler struct {
	fileUsecases usecases.FileUseCases
}

func NewFileHandler(fileUsecases usecases.FileUseCases) (fileHandler *FileHandler, err error) {
	if fileUsecases == nil {
		err = voFile.ErrNilFileUseCases
		return
	}
	fileHandler = &FileHandler{
		fileUsecases: fileUsecases,
	}
	return
}

func (handler *FileHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/loadfile", handler.LoadFile).Methods("POST")
}

func (handler *FileHandler) LoadFile(writer http.ResponseWriter, request *http.Request) {

	// Get file from request
	file, _, err := request.FormFile("file")
	if err != nil {
		log.Printf("Error getting file from request: %v", err)
		http.Error(writer, "Error getting file from request", http.StatusBadRequest)
		return
	}
	fileName, err := request.FormValue("filename"), request.ParseMultipartForm(32<<20)
	if err != nil {
		log.Printf("Error getting file name from request: %v", err)
		http.Error(writer, "Error getting file from request", http.StatusBadRequest)
		return
	}
	log.Println("File name: ", fileName)
	txFile := entity.NewTxFile(fileName, "uploaded", uuid.New().String(), 0)
	if err != nil {
		log.Printf("Error converting multipart file to os file: %v", err)
		http.Error(writer, "Error converting multipart file to os file", http.StatusInternalServerError)
		return
	}
	err = handler.fileUsecases.ProcessMultipartFile(*txFile, file)
	if err != nil {
		log.Printf("Error processing file: %v", err)
		http.Error(writer, "Error processing file", http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusCreated)
}
