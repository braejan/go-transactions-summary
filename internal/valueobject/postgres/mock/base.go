package mock

import (
	"database/sql"

	"github.com/stretchr/testify/mock"
)

// mockBasePostgresDatabase is a mock of the basePostgresDatabase interface implementation.
type mockBasePostgresDatabase struct {
	mock.Mock
}

// NewMockBasePostgresDatabase returns a new mock instance.
func NewMockBasePostgresDatabase() *mockBasePostgresDatabase {
	return &mockBasePostgresDatabase{}
}

// Open provides a mock function with given fields: dataSourceName
func (_m *mockBasePostgresDatabase) Open() (db *sql.DB, err error) {
	ret := _m.Called()

	var r0 *sql.DB
	if rf, ok := ret.Get(0).(func() *sql.DB); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.DB)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Close provides a mock function with given fields: db
func (_m *mockBasePostgresDatabase) Close(db *sql.DB) (err error) {
	ret := _m.Called(db)

	var r0 error
	if rf, ok := ret.Get(0).(func(*sql.DB) error); ok {
		r0 = rf(db)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BeginTx provides a mock function with given fields: db
func (_m *mockBasePostgresDatabase) BeginTx(db *sql.DB) (tx *sql.Tx, err error) {
	ret := _m.Called(db)

	var r0 *sql.Tx
	if rf, ok := ret.Get(0).(func(*sql.DB) *sql.Tx); ok {
		r0 = rf(db)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.Tx)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*sql.DB) error); ok {
		r1 = rf(db)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Commit provides a mock function with given fields: tx
func (_m *mockBasePostgresDatabase) Commit(tx *sql.Tx) (err error) {
	ret := _m.Called(tx)

	var r0 error
	if rf, ok := ret.Get(0).(func(*sql.Tx) error); ok {
		r0 = rf(tx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Rollback provides a mock function with given fields: tx
func (_m *mockBasePostgresDatabase) Rollback(tx *sql.Tx) (err error) {
	ret := _m.Called(tx)

	var r0 error
	if rf, ok := ret.Get(0).(func(*sql.Tx) error); ok {
		r0 = rf(tx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Exec provides a mock function with given fields: tx, query, args
func (_m *mockBasePostgresDatabase) Exec(tx *sql.Tx, query string, args ...interface{}) (result sql.Result, err error) {
	ret := _m.Called(tx, query, args)
	var r0 sql.Result
	if rf, ok := ret.Get(0).(func(*sql.Tx, string, ...interface{}) sql.Result); ok {
		r0 = rf(tx, query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(sql.Result)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*sql.Tx, string, ...interface{}) error); ok {
		r1 = rf(tx, query, args...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Query provides a mock function with given fields: tx, query, args
func (_m *mockBasePostgresDatabase) Query(tx *sql.Tx, query string, args ...interface{}) (rows *sql.Rows, err error) {
	ret := _m.Called(tx, query, args)
	var r0 *sql.Rows
	if rf, ok := ret.Get(0).(func(*sql.Tx, string, ...interface{}) *sql.Rows); ok {
		r0 = rf(tx, query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.Rows)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*sql.Tx, string, ...interface{}) error); ok {
		r1 = rf(tx, query, args...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
