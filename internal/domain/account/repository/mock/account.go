package mock

import (
	"github.com/braejan/go-transactions-summary/internal/domain/account/entity"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

// mockAccountRepository is a mock of the AccountRepository interface implementation.
type mockAccountRepository struct {
	mock.Mock
}

// NewMockAccountRepository returns a new mock instance.
func NewMockAccountRepository() *mockAccountRepository {
	return &mockAccountRepository{}
}

// GetByID provides a mock function with given fields: ID uuid.UUID
func (_m *mockAccountRepository) GetByID(ID uuid.UUID) (acc *entity.Account, err error) {
	ret := _m.Called(ID)

	var r0 *entity.Account
	if rf, ok := ret.Get(0).(func(uuid.UUID) *entity.Account); ok {
		r0 = rf(ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Account)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uuid.UUID) error); ok {
		r1 = rf(ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByUserID provides a mock function with given fields: userID
func (_m *mockAccountRepository) GetByUserID(userID int64) (acc *entity.Account, err error) {
	ret := _m.Called(userID)

	var r0 *entity.Account
	if rf, ok := ret.Get(0).(func(int64) *entity.Account); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Account)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: entity.Account
func (_m *mockAccountRepository) Create(acc *entity.Account) (err error) {
	ret := _m.Called(acc)

	var r0 error
	if rf, ok := ret.Get(0).(func(*entity.Account) error); ok {
		r0 = rf(acc)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: acc *entity.Account
func (_m *mockAccountRepository) Update(acc *entity.Account) (err error) {
	ret := _m.Called(acc)

	var r0 error
	if rf, ok := ret.Get(0).(func(*entity.Account) error); ok {
		r0 = rf(acc)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
