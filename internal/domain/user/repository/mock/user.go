package mock

import (
	"github.com/braejan/go-transactions-summary/internal/domain/user/entity"
	"github.com/stretchr/testify/mock"
)

// mockUserRepository is a mock of the UserRepository interface implementation.
type mockUserRepository struct {
	mock.Mock
}

// NewMockUserRepository returns a new mock instance.
func NewMockUserRepository() *mockUserRepository {
	return &mockUserRepository{}
}

// GetByID provides a mock function with given fields: ID
func (_m *mockUserRepository) GetByID(ID int64) (user *entity.User, err error) {
	ret := _m.Called(ID)

	var r0 *entity.User
	if rf, ok := ret.Get(0).(func(int64) *entity.User); ok {
		r0 = rf(ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByEmail provides a mock function with given fields: email
func (_m *mockUserRepository) GetByEmail(email string) (user *entity.User, err error) {
	ret := _m.Called(email)

	var r0 *entity.User
	if rf, ok := ret.Get(0).(func(string) *entity.User); ok {
		r0 = rf(email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: user
func (_m *mockUserRepository) Create(user *entity.User) (err error) {
	ret := _m.Called(user)

	var r0 error
	if rf, ok := ret.Get(0).(func(*entity.User) error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: user
func (_m *mockUserRepository) Update(user *entity.User) (err error) {
	ret := _m.Called(user)

	var r0 error
	if rf, ok := ret.Get(0).(func(*entity.User) error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
