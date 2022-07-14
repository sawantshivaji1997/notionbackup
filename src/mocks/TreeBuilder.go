// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	tree "github.com/sawantshivaji1997/notionbackup/src/tree"
	mock "github.com/stretchr/testify/mock"
)

// TreeBuilder is an autogenerated mock type for the TreeBuilder type
type TreeBuilder struct {
	mock.Mock
}

// BuildTree provides a mock function with given fields: _a0
func (_m *TreeBuilder) BuildTree(_a0 context.Context) (*tree.Tree, error) {
	ret := _m.Called(_a0)

	var r0 *tree.Tree
	if rf, ok := ret.Get(0).(func(context.Context) *tree.Tree); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*tree.Tree)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewTreeBuilder interface {
	mock.TestingT
	Cleanup(func())
}

// NewTreeBuilder creates a new instance of TreeBuilder. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTreeBuilder(t mockConstructorTestingTNewTreeBuilder) *TreeBuilder {
	mock := &TreeBuilder{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
