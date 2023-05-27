package mock

import (
	"github.com/braejan/go-transactions-summary/internal/domain/account/entity"
	"github.com/stretchr/testify/mock"
)

// mockAccountUseCases is a mock of the AccountUsecases interface implementation.
type mockAccountUseCases struct {
	mock.Mock
}

// NewMockAccountUseCases returns a new mock instance.
func NewMockAccountUseCases() *mockAccountUseCases {
	return &mockAccountUseCases{}
}

// GetByID provides a mock function with given fields: ID
func (_m *mockAccountUseCases) GetByID(ID string) (acc entity.Account, err error) {
	ret := _m.Called(ID)

	var r0 entity.Account
	if rf, ok := ret.Get(0).(func(string) entity.Account); ok {
		r0 = rf(ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(entity.Account)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByUserID provides a mock function with given fields: userID
func (_m *mockAccountUseCases) GetByUserID(userID int64) (acc entity.Account, err error) {
	ret := _m.Called(userID)

	var r0 entity.Account
	if rf, ok := ret.Get(0).(func(int64) entity.Account); ok {
		r0 = rf(userID)
	} else {
		r0 = ret.Get(0).(entity.Account)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: userID
func (_m *mockAccountUseCases) Create(userID int64) (err error) {
	ret := _m.Called(userID)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64) error); ok {
		r0 = rf(userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: ID, balance, active
func (_m *mockAccountUseCases) Update(ID string, balance float64, active bool) (err error) {
	ret := _m.Called(ID, balance, active)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, float64, bool) error); ok {
		r0 = rf(ID, balance, active)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
