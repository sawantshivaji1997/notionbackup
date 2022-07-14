// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	notionapi "github.com/jomei/notionapi"
	mock "github.com/stretchr/testify/mock"
)

// NewClient is an autogenerated mock type for the NewClient type
type NewClient struct {
	mock.Mock
}

// Execute provides a mock function with given fields: _a0, _a1
func (_m *NewClient) Execute(_a0 notionapi.Token, _a1 ...notionapi.ClientOption) *notionapi.Client {
	_va := make([]interface{}, len(_a1))
	for _i := range _a1 {
		_va[_i] = _a1[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *notionapi.Client
	if rf, ok := ret.Get(0).(func(notionapi.Token, ...notionapi.ClientOption) *notionapi.Client); ok {
		r0 = rf(_a0, _a1...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*notionapi.Client)
		}
	}

	return r0
}

type mockConstructorTestingTNewNewClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewNewClient creates a new instance of NewClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewNewClient(t mockConstructorTestingTNewNewClient) *NewClient {
	mock := &NewClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
