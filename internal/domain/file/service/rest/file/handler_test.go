package file_test

import (
	"bytes"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/braejan/go-transactions-summary/internal/domain/file/service/rest/file"
	fileMock "github.com/braejan/go-transactions-summary/internal/domain/file/usecases/mock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// TestNewFileHandler tests the NewFileHandler function.
func TestNewFileHandler(t *testing.T) {
	// Given a valid FileUseCases
	mockFileUseCases := fileMock.NewMockFileUseCases()
	// When NewFileHandler is called
	fileHandler, err := file.NewFileHandler(mockFileUseCases)
	assert.Nil(t, err)
	// Then the returned FileHandler is not nil
	assert.NotNil(t, fileHandler)
}

// TestLoadFile_Fail_FormFile tests the LoadFile function when FormFile fails.
func TestLoadFile_Fail_FormFile(t *testing.T) {
	// Given a valid FileHandler
	mockFileUseCases := fileMock.NewMockFileUseCases()
	fileHandler, err := file.NewFileHandler(mockFileUseCases)
	assert.Nil(t, err)
	// And a empty body file
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	// And a POST request with a file that fails to be read
	request, err := http.NewRequest("POST", "/loadfile", body)
	assert.Nil(t, err)
	request.Header.Add("Content-Type", writer.FormDataContentType())
	// And a HTTP response recorder
	responseRecorder := httptest.NewRecorder()
	// And a registered route
	router := mux.NewRouter()
	fileHandler.RegisterRoutes(router)
	// When send the request to /loadfile
	router.ServeHTTP(responseRecorder, request)
	// Then the returned status is BadRequest
	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
}

// TestLoadFile_Fail_ProcessFile tests the LoadFile function when ProcessFile fails.
func TestLoadFile_Fail_ProcessFile(t *testing.T) {
	// Given a valid FileHandler
	mockFileUseCases := fileMock.NewMockFileUseCases()
	mockFileUseCases.On("ProcessMultipartFile", mock.Anything, mock.Anything).Return(errors.New("error processing file"))
	fileHandler, err := file.NewFileHandler(mockFileUseCases)
	assert.Nil(t, err)
	// And a valid file
	currentDir, _ := os.Getwd()
	filePath := fmt.Sprintf("%s/%s", currentDir, "test/files/txns_invalid_last_record.csv")
	file, err := os.Open(filePath)
	assert.Nil(t, err)
	// And a empty body file
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", file.Name())
	assert.Nil(t, err)
	fileBytes, err := os.ReadFile(file.Name())
	assert.Nil(t, err)
	_, err = part.Write(fileBytes)
	assert.Nil(t, err)
	err = writer.Close()
	assert.Nil(t, err)
	// And a POST request with a file that fails to be read
	request, err := http.NewRequest("POST", "/loadfile", body)
	assert.Nil(t, err)
	request.Header.Add("Content-Type", writer.FormDataContentType())
	// And a HTTP response recorder
	responseRecorder := httptest.NewRecorder()
	// And a registered route
	router := mux.NewRouter()
	fileHandler.RegisterRoutes(router)
	// When send the request to /loadfile
	router.ServeHTTP(responseRecorder, request)
	// Then the returned status is BadRequest
	assert.Equal(t, http.StatusInternalServerError, responseRecorder.Code)
}

// TestLoadFile_Success tests the LoadFile function.
func TestLoadFile_Success(t *testing.T) {
	// Given a valid FileHandler
	mockFileUseCases := fileMock.NewMockFileUseCases()
	mockFileUseCases.On("ProcessMultipartFile", mock.Anything, mock.Anything).Return(nil)
	fileHandler, err := file.NewFileHandler(mockFileUseCases)
	assert.Nil(t, err)
	// And a valid file
	currentDir, _ := os.Getwd()
	filePath := fmt.Sprintf("%s/%s", currentDir, "test/files/txns_simple.csv")
	file, err := os.Open(filePath)
	assert.Nil(t, err)
	// And a empty body file
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", file.Name())
	assert.Nil(t, err)
	fileBytes, err := os.ReadFile(file.Name())
	assert.Nil(t, err)
	_, err = part.Write(fileBytes)
	assert.Nil(t, err)
	err = writer.Close()
	assert.Nil(t, err)
	// And a POST request with a file that fails to be read
	request, err := http.NewRequest("POST", "/loadfile", body)
	assert.Nil(t, err)
	request.Header.Add("Content-Type", writer.FormDataContentType())
	// And a HTTP response recorder
	responseRecorder := httptest.NewRecorder()
	// And a registered route
	router := mux.NewRouter()
	fileHandler.RegisterRoutes(router)
	// When send the request to /loadfile
	router.ServeHTTP(responseRecorder, request)
	// Then the returned status is BadRequest
	assert.Equal(t, http.StatusCreated, responseRecorder.Code)
}
