package mock

import (
	"mime/multipart"
	"os"

	fileEntity "github.com/braejan/go-transactions-summary/internal/domain/file/entity"
	"github.com/stretchr/testify/mock"
)

// mockFileUseCases is a mock of FileUseCases interface.
type mockFileUseCases struct {
	mock.Mock
}

// NewMockFileUseCases returns a new mock instance.
func NewMockFileUseCases() *mockFileUseCases {
	return &mockFileUseCases{}
}

// FileUseCases interface implementation.

// ReadAndProcessFile mocks base method.
func (m *mockFileUseCases) ReadAndProcessFile(txFile fileEntity.TxFile, isS3 bool) error {
	ret := m.Called(txFile, isS3)

	var r0 error
	if rf, ok := ret.Get(0).(func(fileEntity.TxFile, bool) error); ok {
		r0 = rf(txFile, isS3)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CheckFile mocks base method.
func (m *mockFileUseCases) CheckFile(txFile fileEntity.TxFile, isS3 bool) error {
	ret := m.Called(txFile, isS3)

	var r0 error
	if rf, ok := ret.Get(0).(func(fileEntity.TxFile, bool) error); ok {
		r0 = rf(txFile, isS3)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ProcessFile mocks base method.
func (m *mockFileUseCases) ProcessFile(txFile fileEntity.TxFile, file *os.File) error {
	ret := m.Called(txFile, file)

	var r0 error
	if rf, ok := ret.Get(0).(func(fileEntity.TxFile, *os.File) error); ok {
		r0 = rf(txFile, file)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ProcessMultipartFile mocks base method.
func (m *mockFileUseCases) ProcessMultipartFile(txFile fileEntity.TxFile, file multipart.File) error {
	ret := m.Called(txFile, file)

	var r0 error
	if rf, ok := ret.Get(0).(func(fileEntity.TxFile, multipart.File) error); ok {
		r0 = rf(txFile, file)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
