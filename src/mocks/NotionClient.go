// Code generated by mockery v2.12.1. DO NOT EDIT.

package mocks

import (
	context "context"

	notionapi "github.com/jomei/notionapi"
	mock "github.com/stretchr/testify/mock"

	notionclient "github.com/sawantshivaji1997/notionbackup/src/notionclient"

	testing "testing"
)

// NotionClient is an autogenerated mock type for the NotionClient type
type NotionClient struct {
	mock.Mock
}

// GetAllDatabases provides a mock function with given fields: _a0, _a1
func (_m *NotionClient) GetAllDatabases(_a0 context.Context, _a1 notionapi.Cursor) ([]notionapi.Database, notionapi.Cursor, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []notionapi.Database
	if rf, ok := ret.Get(0).(func(context.Context, notionapi.Cursor) []notionapi.Database); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]notionapi.Database)
		}
	}

	var r1 notionapi.Cursor
	if rf, ok := ret.Get(1).(func(context.Context, notionapi.Cursor) notionapi.Cursor); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Get(1).(notionapi.Cursor)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, notionapi.Cursor) error); ok {
		r2 = rf(_a0, _a1)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetAllPages provides a mock function with given fields: _a0, _a1
func (_m *NotionClient) GetAllPages(_a0 context.Context, _a1 notionapi.Cursor) ([]notionapi.Page, notionapi.Cursor, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []notionapi.Page
	if rf, ok := ret.Get(0).(func(context.Context, notionapi.Cursor) []notionapi.Page); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]notionapi.Page)
		}
	}

	var r1 notionapi.Cursor
	if rf, ok := ret.Get(1).(func(context.Context, notionapi.Cursor) notionapi.Cursor); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Get(1).(notionapi.Cursor)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, notionapi.Cursor) error); ok {
		r2 = rf(_a0, _a1)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetBlockByID provides a mock function with given fields: _a0, _a1
func (_m *NotionClient) GetBlockByID(_a0 context.Context, _a1 notionclient.BlockID) (notionapi.Block, error) {
	ret := _m.Called(_a0, _a1)

	var r0 notionapi.Block
	if rf, ok := ret.Get(0).(func(context.Context, notionclient.BlockID) notionapi.Block); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(notionapi.Block)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, notionclient.BlockID) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBlocksOfPages provides a mock function with given fields: _a0, _a1, _a2
func (_m *NotionClient) GetBlocksOfPages(_a0 context.Context, _a1 notionclient.PageID, _a2 notionapi.Cursor) ([]notionapi.Block, notionapi.Cursor, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 []notionapi.Block
	if rf, ok := ret.Get(0).(func(context.Context, notionclient.PageID, notionapi.Cursor) []notionapi.Block); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]notionapi.Block)
		}
	}

	var r1 notionapi.Cursor
	if rf, ok := ret.Get(1).(func(context.Context, notionclient.PageID, notionapi.Cursor) notionapi.Cursor); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Get(1).(notionapi.Cursor)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, notionclient.PageID, notionapi.Cursor) error); ok {
		r2 = rf(_a0, _a1, _a2)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetChildBlocksOfBlock provides a mock function with given fields: _a0, _a1, _a2
func (_m *NotionClient) GetChildBlocksOfBlock(_a0 context.Context, _a1 notionclient.BlockID, _a2 notionapi.Cursor) ([]notionapi.Block, notionapi.Cursor, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 []notionapi.Block
	if rf, ok := ret.Get(0).(func(context.Context, notionclient.BlockID, notionapi.Cursor) []notionapi.Block); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]notionapi.Block)
		}
	}

	var r1 notionapi.Cursor
	if rf, ok := ret.Get(1).(func(context.Context, notionclient.BlockID, notionapi.Cursor) notionapi.Cursor); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Get(1).(notionapi.Cursor)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, notionclient.BlockID, notionapi.Cursor) error); ok {
		r2 = rf(_a0, _a1, _a2)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetDatabaseByID provides a mock function with given fields: _a0, _a1
func (_m *NotionClient) GetDatabaseByID(_a0 context.Context, _a1 notionclient.DatabaseID) (*notionapi.Database, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *notionapi.Database
	if rf, ok := ret.Get(0).(func(context.Context, notionclient.DatabaseID) *notionapi.Database); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*notionapi.Database)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, notionclient.DatabaseID) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDatabasesByName provides a mock function with given fields: _a0, _a1, _a2
func (_m *NotionClient) GetDatabasesByName(_a0 context.Context, _a1 notionclient.DatabaseName, _a2 notionapi.Cursor) ([]notionapi.Database, notionapi.Cursor, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 []notionapi.Database
	if rf, ok := ret.Get(0).(func(context.Context, notionclient.DatabaseName, notionapi.Cursor) []notionapi.Database); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]notionapi.Database)
		}
	}

	var r1 notionapi.Cursor
	if rf, ok := ret.Get(1).(func(context.Context, notionclient.DatabaseName, notionapi.Cursor) notionapi.Cursor); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Get(1).(notionapi.Cursor)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, notionclient.DatabaseName, notionapi.Cursor) error); ok {
		r2 = rf(_a0, _a1, _a2)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetPageByID provides a mock function with given fields: _a0, _a1
func (_m *NotionClient) GetPageByID(_a0 context.Context, _a1 notionclient.PageID) (*notionapi.Page, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *notionapi.Page
	if rf, ok := ret.Get(0).(func(context.Context, notionclient.PageID) *notionapi.Page); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*notionapi.Page)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, notionclient.PageID) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPagesByName provides a mock function with given fields: _a0, _a1, _a2
func (_m *NotionClient) GetPagesByName(_a0 context.Context, _a1 notionclient.PageName, _a2 notionapi.Cursor) ([]notionapi.Page, notionapi.Cursor, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 []notionapi.Page
	if rf, ok := ret.Get(0).(func(context.Context, notionclient.PageName, notionapi.Cursor) []notionapi.Page); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]notionapi.Page)
		}
	}

	var r1 notionapi.Cursor
	if rf, ok := ret.Get(1).(func(context.Context, notionclient.PageName, notionapi.Cursor) notionapi.Cursor); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Get(1).(notionapi.Cursor)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, notionclient.PageName, notionapi.Cursor) error); ok {
		r2 = rf(_a0, _a1, _a2)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetPagesOfDatabase provides a mock function with given fields: _a0, _a1, _a2
func (_m *NotionClient) GetPagesOfDatabase(_a0 context.Context, _a1 notionclient.DatabaseID, _a2 notionapi.Cursor) ([]notionapi.Page, notionapi.Cursor, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 []notionapi.Page
	if rf, ok := ret.Get(0).(func(context.Context, notionclient.DatabaseID, notionapi.Cursor) []notionapi.Page); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]notionapi.Page)
		}
	}

	var r1 notionapi.Cursor
	if rf, ok := ret.Get(1).(func(context.Context, notionclient.DatabaseID, notionapi.Cursor) notionapi.Cursor); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Get(1).(notionapi.Cursor)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, notionclient.DatabaseID, notionapi.Cursor) error); ok {
		r2 = rf(_a0, _a1, _a2)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// NewNotionClient creates a new instance of NotionClient. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewNotionClient(t testing.TB) *NotionClient {
	mock := &NotionClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
