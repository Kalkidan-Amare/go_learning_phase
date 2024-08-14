// Code generated by mockery v2.44.1. DO NOT EDIT.

package mocks

import (
	domain "task_manager/domain"

	mock "github.com/stretchr/testify/mock"

	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

// TaskRepositoryInterface is an autogenerated mock type for the TaskRepositoryInterface type
type TaskRepositoryInterface struct {
	mock.Mock
}

// AddTask provides a mock function with given fields: task
func (_m *TaskRepositoryInterface) AddTask(task *domain.Task) (interface{}, error) {
	ret := _m.Called(task)

	if len(ret) == 0 {
		panic("no return value specified for AddTask")
	}

	var r0 interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(*domain.Task) (interface{}, error)); ok {
		return rf(task)
	}
	if rf, ok := ret.Get(0).(func(*domain.Task) interface{}); ok {
		r0 = rf(task)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(*domain.Task) error); ok {
		r1 = rf(task)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteTask provides a mock function with given fields: id
func (_m *TaskRepositoryInterface) DeleteTask(id primitive.ObjectID) error {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteTask")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(primitive.ObjectID) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllTasks provides a mock function with given fields:
func (_m *TaskRepositoryInterface) GetAllTasks() ([]domain.Task, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAllTasks")
	}

	var r0 []domain.Task
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]domain.Task, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []domain.Task); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Task)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTaskByID provides a mock function with given fields: id
func (_m *TaskRepositoryInterface) GetTaskByID(id primitive.ObjectID) (*domain.Task, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetTaskByID")
	}

	var r0 *domain.Task
	var r1 error
	if rf, ok := ret.Get(0).(func(primitive.ObjectID) (*domain.Task, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(primitive.ObjectID) *domain.Task); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Task)
		}
	}

	if rf, ok := ret.Get(1).(func(primitive.ObjectID) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateTask provides a mock function with given fields: id, taskData
func (_m *TaskRepositoryInterface) UpdateTask(id primitive.ObjectID, taskData *domain.Task) (*domain.Task, error) {
	ret := _m.Called(id, taskData)

	if len(ret) == 0 {
		panic("no return value specified for UpdateTask")
	}

	var r0 *domain.Task
	var r1 error
	if rf, ok := ret.Get(0).(func(primitive.ObjectID, *domain.Task) (*domain.Task, error)); ok {
		return rf(id, taskData)
	}
	if rf, ok := ret.Get(0).(func(primitive.ObjectID, *domain.Task) *domain.Task); ok {
		r0 = rf(id, taskData)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Task)
		}
	}

	if rf, ok := ret.Get(1).(func(primitive.ObjectID, *domain.Task) error); ok {
		r1 = rf(id, taskData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewTaskRepositoryInterface creates a new instance of TaskRepositoryInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTaskRepositoryInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *TaskRepositoryInterface {
	mock := &TaskRepositoryInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
