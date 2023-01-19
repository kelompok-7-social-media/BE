// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	posting "project/features/posting"

	mock "github.com/stretchr/testify/mock"
)

// PostingService is an autogenerated mock type for the PostingService type
type PostingService struct {
	mock.Mock
}

// Add provides a mock function with given fields: token, newBook
func (_m *PostingService) Add(token interface{}, newBook posting.Core) (posting.Core, error) {
	ret := _m.Called(token, newBook)

	var r0 posting.Core
	if rf, ok := ret.Get(0).(func(interface{}, posting.Core) posting.Core); ok {
		r0 = rf(token, newBook)
	} else {
		r0 = ret.Get(0).(posting.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}, posting.Core) error); ok {
		r1 = rf(token, newBook)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: token, bookID
func (_m *PostingService) Delete(token interface{}, bookID int) error {
	ret := _m.Called(token, bookID)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, int) error); ok {
		r0 = rf(token, bookID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllPost provides a mock function with given fields:
func (_m *PostingService) GetAllPost() ([]posting.Core, error) {
	ret := _m.Called()

	var r0 []posting.Core
	if rf, ok := ret.Get(0).(func() []posting.Core); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]posting.Core)
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

// MyPost provides a mock function with given fields: token
func (_m *PostingService) MyPost(token interface{}) ([]posting.Core, error) {
	ret := _m.Called(token)

	var r0 []posting.Core
	if rf, ok := ret.Get(0).(func(interface{}) []posting.Core); ok {
		r0 = rf(token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]posting.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}) error); ok {
		r1 = rf(token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: token, bookID, updatedData
func (_m *PostingService) Update(token interface{}, bookID int, updatedData posting.Core) (posting.Core, error) {
	ret := _m.Called(token, bookID, updatedData)

	var r0 posting.Core
	if rf, ok := ret.Get(0).(func(interface{}, int, posting.Core) posting.Core); ok {
		r0 = rf(token, bookID, updatedData)
	} else {
		r0 = ret.Get(0).(posting.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}, int, posting.Core) error); ok {
		r1 = rf(token, bookID, updatedData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewPostingService interface {
	mock.TestingT
	Cleanup(func())
}

// NewPostingService creates a new instance of PostingService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPostingService(t mockConstructorTestingTNewPostingService) *PostingService {
	mock := &PostingService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
