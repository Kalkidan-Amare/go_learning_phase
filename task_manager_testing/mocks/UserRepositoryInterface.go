// Code generated by mockery v2.44.1. DO NOT EDIT.

package mocks

import (
	domain "task_manager/domain"

	mock "github.com/stretchr/testify/mock"

	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

// UserRepositoryInterface is an autogenerated mock type for the UserRepositoryInterface type
type UserRepositoryInterface struct {
	mock.Mock
}

// AddUser provides a mock function with given fields: user
func (_m *UserRepositoryInterface) AddUser(user *domain.User) (interface{}, error) {
	ret := _m.Called(user)

	if len(ret) == 0 {
		panic("no return value specified for AddUser")
	}

	var r0 interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(*domain.User) (interface{}, error)); ok {
		return rf(user)
	}
	if rf, ok := ret.Get(0).(func(*domain.User) interface{}); ok {
		r0 = rf(user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(*domain.User) error); ok {
		r1 = rf(user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByID provides a mock function with given fields: objectID
func (_m *UserRepositoryInterface) GetUserByID(objectID primitive.ObjectID) (*domain.User, error) {
	ret := _m.Called(objectID)

	if len(ret) == 0 {
		panic("no return value specified for GetUserByID")
	}

	var r0 *domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(primitive.ObjectID) (*domain.User, error)); ok {
		return rf(objectID)
	}
	if rf, ok := ret.Get(0).(func(primitive.ObjectID) *domain.User); ok {
		r0 = rf(objectID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	if rf, ok := ret.Get(1).(func(primitive.ObjectID) error); ok {
		r1 = rf(objectID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByUsername provides a mock function with given fields: username
func (_m *UserRepositoryInterface) GetUserByUsername(username string) (*domain.User, error) {
	ret := _m.Called(username)

	if len(ret) == 0 {
		panic("no return value specified for GetUserByUsername")
	}

	var r0 *domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*domain.User, error)); ok {
		return rf(username)
	}
	if rf, ok := ret.Get(0).(func(string) *domain.User); ok {
		r0 = rf(username)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewUserRepositoryInterface creates a new instance of UserRepositoryInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserRepositoryInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserRepositoryInterface {
	mock := &UserRepositoryInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
