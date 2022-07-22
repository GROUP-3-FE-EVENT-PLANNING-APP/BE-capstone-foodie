// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	comments "capstone/group3/features/comments"

	mock "github.com/stretchr/testify/mock"
)

// Data is an autogenerated mock type for the Data type
type Data struct {
	mock.Mock
}

// InsertComment provides a mock function with given fields: input
func (_m *Data) InsertComment(input comments.Core) (int, error) {
	ret := _m.Called(input)

	var r0 int
	if rf, ok := ret.Get(0).(func(comments.Core) int); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(comments.Core) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectCommentByIdResto provides a mock function with given fields: idResto, limitint, offsetint
func (_m *Data) SelectCommentByIdResto(idResto int, limitint int, offsetint int) ([]comments.Core, error) {
	ret := _m.Called(idResto, limitint, offsetint)

	var r0 []comments.Core
	if rf, ok := ret.Get(0).(func(int, int, int) []comments.Core); ok {
		r0 = rf(idResto, limitint, offsetint)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]comments.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int, int) error); ok {
		r1 = rf(idResto, limitint, offsetint)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectRatingByIdResto provides a mock function with given fields: idResto
func (_m *Data) SelectRatingByIdResto(idResto int) (float64, error) {
	ret := _m.Called(idResto)

	var r0 float64
	if rf, ok := ret.Get(0).(func(int) float64); ok {
		r0 = rf(idResto)
	} else {
		r0 = ret.Get(0).(float64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(idResto)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewData interface {
	mock.TestingT
	Cleanup(func())
}

// NewData creates a new instance of Data. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewData(t mockConstructorTestingTNewData) *Data {
	mock := &Data{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
