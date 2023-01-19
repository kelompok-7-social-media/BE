// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	komentar "project/features/komentar"

	mock "github.com/stretchr/testify/mock"
)

// KomentarService is an autogenerated mock type for the KomentarService type
type KomentarService struct {
	mock.Mock
}

// Add provides a mock function with given fields: token, newComment
func (_m *KomentarService) Add(token interface{}, newComment komentar.Core) (komentar.Core, error) {
	ret := _m.Called(token, newComment)

	var r0 komentar.Core
	if rf, ok := ret.Get(0).(func(interface{}, komentar.Core) komentar.Core); ok {
		r0 = rf(token, newComment)
	} else {
		r0 = ret.Get(0).(komentar.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}, komentar.Core) error); ok {
		r1 = rf(token, newComment)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: token, postID, commentID
func (_m *KomentarService) Delete(token interface{}, postID int, commentID int) error {
	ret := _m.Called(token, postID, commentID)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, int, int) error); ok {
		r0 = rf(token, postID, commentID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetCommentsByPost provides a mock function with given fields: postID
func (_m *KomentarService) GetCommentsByPost(postID int) ([]komentar.Core, error) {
	ret := _m.Called(postID)

	var r0 []komentar.Core
	if rf, ok := ret.Get(0).(func(int) []komentar.Core); ok {
		r0 = rf(postID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]komentar.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(postID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: token, commentID, updatedComment
func (_m *KomentarService) Update(token interface{}, commentID int, updatedComment komentar.Core) (komentar.Core, error) {
	ret := _m.Called(token, commentID, updatedComment)

	var r0 komentar.Core
	if rf, ok := ret.Get(0).(func(interface{}, int, komentar.Core) komentar.Core); ok {
		r0 = rf(token, commentID, updatedComment)
	} else {
		r0 = ret.Get(0).(komentar.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}, int, komentar.Core) error); ok {
		r1 = rf(token, commentID, updatedComment)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewKomentarService interface {
	mock.TestingT
	Cleanup(func())
}

// NewKomentarService creates a new instance of KomentarService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewKomentarService(t mockConstructorTestingTNewKomentarService) *KomentarService {
	mock := &KomentarService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
