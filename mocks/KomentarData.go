// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	komentar "project/features/komentar"

	mock "github.com/stretchr/testify/mock"
)

// KomentarData is an autogenerated mock type for the KomentarData type
type KomentarData struct {
	mock.Mock
}

// Add provides a mock function with given fields: userID, newComment
func (_m *KomentarData) Add(userID int, newComment komentar.Core) (komentar.Core, error) {
	ret := _m.Called(userID, newComment)

	var r0 komentar.Core
	if rf, ok := ret.Get(0).(func(int, komentar.Core) komentar.Core); ok {
		r0 = rf(userID, newComment)
	} else {
		r0 = ret.Get(0).(komentar.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, komentar.Core) error); ok {
		r1 = rf(userID, newComment)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: UserID, postID, commentID
func (_m *KomentarData) Delete(UserID int, postID int, commentID int) error {
	ret := _m.Called(UserID, postID, commentID)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, int, int) error); ok {
		r0 = rf(UserID, postID, commentID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetCommentsByPost provides a mock function with given fields: postID
func (_m *KomentarData) GetCommentsByPost(postID int) ([]komentar.Core, error) {
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

// Update provides a mock function with given fields: UserID, commentID, updatedComment
func (_m *KomentarData) Update(UserID int, commentID int, updatedComment komentar.Core) (komentar.Core, error) {
	ret := _m.Called(UserID, commentID, updatedComment)

	var r0 komentar.Core
	if rf, ok := ret.Get(0).(func(int, int, komentar.Core) komentar.Core); ok {
		r0 = rf(UserID, commentID, updatedComment)
	} else {
		r0 = ret.Get(0).(komentar.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int, komentar.Core) error); ok {
		r1 = rf(UserID, commentID, updatedComment)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewKomentarData interface {
	mock.TestingT
	Cleanup(func())
}

// NewKomentarData creates a new instance of KomentarData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewKomentarData(t mockConstructorTestingTNewKomentarData) *KomentarData {
	mock := &KomentarData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
