// Code generated by mockery v1.0.0. DO NOT EDIT.

package transaction

import mock "github.com/stretchr/testify/mock"

// MockDatastore is an autogenerated mock type for the MockDatastore type
type MockDatastore struct {
	mock.Mock
}

// Create provides a mock function with given fields: txn
func (_m *MockDatastore) Create(txn Transaction) error {
	ret := _m.Called(txn)

	var r0 error
	if rf, ok := ret.Get(0).(func(Transaction) error); ok {
		r0 = rf(txn)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
