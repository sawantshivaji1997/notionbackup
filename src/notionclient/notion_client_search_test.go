package notionclient_test

import (
	"context"
	"errors"
	"testing"

	"github.com/jomei/notionapi"
	"github.com/sawantshivaji1997/notionbackup/src/notionclient"
	"github.com/sawantshivaji1997/notionbackup/src/utils"
	"github.com/stretchr/testify/assert"
)

const (
	ERROR_STR                             = "error occurred"
	TEST_DATA_PATH                        = "./../../testdata/notionclient/"
	SEARCH_PATH                           = TEST_DATA_PATH + "search/"
	SEARCH_ALL_PAGES_JSON                 = SEARCH_PATH + "search_all_pages.json"
	SEARCH_PAGES_WITH_PAGINATION_JSON     = SEARCH_PATH + "search_pages_with_pagination.json"
	SEARCH_PAGES_WITH_NAME_JSON           = SEARCH_PATH + "search_pages_with_name.json"
	SEARCH_ALL_DATABASES_JSON             = SEARCH_PATH + "search_all_databases.json"
	SEARCH_DATABASES_WITH_PAGINATION_JSON = SEARCH_PATH + "search_databases_with_pagination.json"
	SEARCH_DATABASES_WITH_NAME_JSON       = SEARCH_PATH + "search_databases_with_name.json"
	EMPTY_SEARCH_RESULT                   = SEARCH_PATH + "empty_search_result.json"

	PAGE_PATH                = TEST_DATA_PATH + "page/"
	PAGE_JSON                = PAGE_PATH + "page.json"
	PAGE_NOT_EXIST_ERROR_STR = "Page does not exist"

	DATABASE_PATH                = TEST_DATA_PATH + "database/"
	DATABASE_JSON                = DATABASE_PATH + "database.json"
	DATABASE_NOT_EXIST_ERROR_STR = "Database does not exist"
)

// Mocking the NewClient from github.com/jomei/notionapi
func newMockedClient(token notionapi.Token, opt ...notionapi.ClientOption) *notionapi.Client {
	return &notionapi.Client{}
}

func TestGetNotionClient(t *testing.T) {
	t.Run("Get Client with valid parameters", func(t *testing.T) {
		client := notionclient.GetNotionApiClient(context.Background(), "asdasd", newMockedClient)
		assert.NotNil(t, client)
	})
}

// Mocking Search service from github.com/jomei/notionapi
type MockedSearchService struct {
	response  *notionapi.SearchResponse
	response2 *notionapi.SearchResponse
	err       error
}

func GetMockedSearchService(t *testing.T, mockFilePath string, err error) *notionclient.NotionApiClient {
	if err != nil {
		return &notionclient.NotionApiClient{
			Client: &notionapi.Client{
				Search: &MockedSearchService{
					response:  nil,
					response2: nil,
					err:       err,
				},
			},
		}
	}

	jsonBytes, err := utils.ReadContentsOfFile(mockFilePath)
	if err != nil {
		t.Fatal(err)
	}

	searchResponse, err := utils.ParseSearchResponseJsonString(jsonBytes)

	if err != nil {
		t.Fatal(err)
	}

	var searchResponse2 notionapi.SearchResponse
	if searchResponse.HasMore {
		searchResponse2, _ := utils.ParseSearchResponseJsonString(jsonBytes)
		searchResponse2.HasMore = false
		searchResponse2.NextCursor = notionapi.Cursor("")
	}

	return &notionclient.NotionApiClient{
		Client: &notionapi.Client{
			Search: &MockedSearchService{
				response:  searchResponse,
				response2: &searchResponse2,
				err:       nil,
			},
		},
	}
}

func (srv *MockedSearchService) Do(ctx context.Context, req *notionapi.SearchRequest) (*notionapi.SearchResponse, error) {
	if req.StartCursor != "" {
		return srv.response2, srv.err
	}

	return srv.response, srv.err
}

func TestGetAllPages(t *testing.T) {

	tests := []struct {
		name        string
		filePath    string
		wantErr     bool
		emptyResult bool
		err         error
		cursorEmpty bool
	}{
		{
			name:        "Get all Pages",
			filePath:    SEARCH_ALL_PAGES_JSON,
			wantErr:     false,
			emptyResult: false,
			err:         nil,
			cursorEmpty: true,
		},
		{
			name:        "Get pages with pagination",
			filePath:    SEARCH_PAGES_WITH_PAGINATION_JSON,
			wantErr:     false,
			emptyResult: false,
			err:         nil,
			cursorEmpty: false,
		},
		{
			name:        "Error in fetching pages",
			filePath:    "",
			wantErr:     true,
			emptyResult: true,
			err:         errors.New(ERROR_STR),
			cursorEmpty: true,
		},
		{
			name:        "Get empty Page list",
			filePath:    EMPTY_SEARCH_RESULT,
			wantErr:     false,
			emptyResult: true,
			err:         nil,
			cursorEmpty: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client := GetMockedSearchService(t, test.filePath, test.err)
			pages, cursor, err := client.GetAllPages(context.Background(), notionapi.Cursor(""))
			if test.wantErr {
				assert.Nil(t, pages)
				assert.NotNil(t, err)
				assert.Empty(t, cursor)
			} else {
				if test.emptyResult {
					assert.Empty(t, pages)
				} else {
					assert.NotEmpty(t, pages)
					if test.cursorEmpty {
						assert.Empty(t, cursor)
					} else {
						assert.NotEmpty(t, cursor)
					}
				}
				assert.Nil(t, err)
			}
		})
	}
}

func TestGetPagesByName(t *testing.T) {
	tests := []struct {
		name        string
		pageName    string
		filePath    string
		wantErr     bool
		emptyResult bool
		err         error
	}{
		{
			name:        "Get existing pages",
			pageName:    "Page-1",
			filePath:    SEARCH_PAGES_WITH_NAME_JSON,
			wantErr:     false,
			emptyResult: false,
			err:         nil,
		},
		{
			name:        "Get non-existing pages",
			pageName:    "Page-2",
			filePath:    EMPTY_SEARCH_RESULT,
			wantErr:     false,
			emptyResult: true,
			err:         nil,
		},
		{
			name:        "Error while fetching pages",
			pageName:    "Page-1",
			filePath:    SEARCH_PAGES_WITH_NAME_JSON,
			wantErr:     true,
			emptyResult: true,
			err:         errors.New(ERROR_STR),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client := GetMockedSearchService(t, test.filePath, test.err)
			pages, _, err := client.GetPagesByName(context.Background(), notionclient.PageName(test.pageName), "")
			if test.wantErr {
				assert.Nil(t, pages)
				assert.NotNil(t, err)
			} else {
				if test.emptyResult {
					assert.Empty(t, pages)
				} else {
					assert.NotEmpty(t, pages)
				}
				assert.Nil(t, err)
			}
		})
	}
}

func TestGetAllDatabases(t *testing.T) {
	tests := []struct {
		name        string
		filePath    string
		wantErr     bool
		emptyResult bool
		err         error
	}{
		{
			name:        "Get all Databases",
			filePath:    SEARCH_ALL_DATABASES_JSON,
			wantErr:     false,
			emptyResult: false,
			err:         nil,
		},
		{
			name:        "Get databases with pagination",
			filePath:    SEARCH_DATABASES_WITH_PAGINATION_JSON,
			wantErr:     false,
			emptyResult: false,
			err:         nil,
		},
		{
			name:        "Error in fetching databases",
			filePath:    "",
			wantErr:     true,
			emptyResult: true,
			err:         errors.New(ERROR_STR),
		},
		{
			name:        "Get empty Database list",
			filePath:    EMPTY_SEARCH_RESULT,
			wantErr:     false,
			emptyResult: true,
			err:         nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client := GetMockedSearchService(t, test.filePath, test.err)
			databases, _, err := client.GetAllDatabases(context.Background(), "")
			if test.wantErr {
				assert.Nil(t, databases)
				assert.NotNil(t, err)
			} else {
				if test.emptyResult {
					assert.Empty(t, databases)
				} else {
					assert.NotEmpty(t, databases)
				}
				assert.Nil(t, err)
			}
		})
	}
}

func TestGetDatabasesByName(t *testing.T) {
	tests := []struct {
		name         string
		databaseName string
		filePath     string
		wantErr      bool
		emptyResult  bool
		err          error
	}{
		{
			name:         "Get existing database",
			databaseName: "Database-1",
			filePath:     SEARCH_DATABASES_WITH_NAME_JSON,
			wantErr:      false,
			emptyResult:  false,
			err:          nil,
		},
		{
			name:         "Get non-existing database",
			databaseName: "Database-2",
			filePath:     EMPTY_SEARCH_RESULT,
			wantErr:      false,
			emptyResult:  true,
			err:          nil,
		},
		{
			name:         "Error while fetching database",
			databaseName: "Database-1",
			filePath:     SEARCH_DATABASES_WITH_NAME_JSON,
			wantErr:      true,
			emptyResult:  true,
			err:          errors.New(ERROR_STR),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client := GetMockedSearchService(t, test.filePath, test.err)
			databases, _, err := client.GetDatabasesByName(context.Background(), notionclient.DatabaseName(test.databaseName), "")
			if test.wantErr {
				assert.Nil(t, databases)
				assert.NotNil(t, err)
			} else {
				if test.emptyResult {
					assert.Empty(t, databases)
				} else {
					assert.NotEmpty(t, databases)
				}
				assert.Nil(t, err)
			}
		})
	}
}

type MockedPageService struct {
	Page *notionapi.Page
	err  error
}

func GetMockedPageService(t *testing.T, mockFilePath string, err error) *notionclient.NotionApiClient {
	if err != nil {
		return &notionclient.NotionApiClient{
			Client: &notionapi.Client{
				Page: &MockedPageService{
					Page: nil,
					err:  err,
				},
			},
		}
	}

	jsonBytes, err := utils.ReadContentsOfFile(mockFilePath)
	if err != nil {
		t.Fatal(err)
	}

	page, err := utils.ParsePageJsonString(jsonBytes)

	if err != nil {
		t.Fatal(err)
	}

	return &notionclient.NotionApiClient{
		Client: &notionapi.Client{
			Page: &MockedPageService{
				Page: page,
				err:  nil,
			},
		},
	}
}

func (srv *MockedPageService) Get(ctx context.Context, id notionapi.PageID) (*notionapi.Page, error) {
	return srv.Page, srv.err
}

func (srv *MockedPageService) Create(ctx context.Context, req *notionapi.PageCreateRequest) (*notionapi.Page, error) {
	// TODO
	return nil, nil
}

func (srv *MockedPageService) Update(ctx context.Context, id notionapi.PageID, req *notionapi.PageUpdateRequest) (*notionapi.Page, error) {
	// TODO
	return nil, nil
}

func TestGetPageByID(t *testing.T) {
	tests := []struct {
		name        string
		pageid      string
		filePath    string
		wantErr     bool
		emptyResult bool
		err         error
	}{
		{
			name:        "Get existing Page",
			pageid:      "05034203-2870-4bc8-b1f9-22c0ae6e56ba",
			filePath:    PAGE_JSON,
			wantErr:     false,
			emptyResult: false,
			err:         nil,
		},
		{
			name:        "Get non-existing Page",
			pageid:      "05034203-2870-4bc8-b1f9-22c0ae6e56aa",
			filePath:    PAGE_JSON,
			wantErr:     true,
			emptyResult: true,
			err:         errors.New(PAGE_NOT_EXIST_ERROR_STR),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client := GetMockedPageService(t, test.filePath, test.err)
			page, err := client.GetPageByID(context.Background(), notionclient.PageID(test.pageid))
			if test.wantErr {
				assert.Nil(t, page)
				assert.NotNil(t, err)
			} else {
				if test.emptyResult {
					assert.Empty(t, page)
				} else {
					assert.NotEmpty(t, page)
				}
				assert.Nil(t, err)
			}
		})

	}
}

type MockedDatabaseService struct {
	Database *notionapi.Database
	err      error
}

func GetMockedDatabaseService(t *testing.T, mockFilePath string, err error) *notionclient.NotionApiClient {
	if err != nil {
		return &notionclient.NotionApiClient{
			Client: &notionapi.Client{
				Database: &MockedDatabaseService{
					Database: nil,
					err:      err,
				},
			},
		}
	}

	jsonBytes, err := utils.ReadContentsOfFile(mockFilePath)
	if err != nil {
		t.Fatal(err)
	}

	database, err := utils.ParseDatabaseJsonString(jsonBytes)

	if err != nil {
		t.Fatal(err)
	}

	return &notionclient.NotionApiClient{
		Client: &notionapi.Client{
			Database: &MockedDatabaseService{
				Database: database,
				err:      nil,
			},
		},
	}
}

func (srv *MockedDatabaseService) Get(ctx context.Context, id notionapi.DatabaseID) (*notionapi.Database, error) {
	return srv.Database, srv.err
}

func (srv *MockedDatabaseService) List(ctx context.Context, pagination *notionapi.Pagination) (*notionapi.DatabaseListResponse, error) {
	// TODO
	return nil, nil
}

func (srv *MockedDatabaseService) Query(ctx context.Context, id notionapi.DatabaseID, req *notionapi.DatabaseQueryRequest) (*notionapi.DatabaseQueryResponse, error) {
	// TODO
	return nil, nil
}

func (srv *MockedDatabaseService) Update(ctx context.Context, id notionapi.DatabaseID, req *notionapi.DatabaseUpdateRequest) (*notionapi.Database, error) {
	// TODO
	return nil, nil
}

func (srv *MockedDatabaseService) Create(ctx context.Context, req *notionapi.DatabaseCreateRequest) (*notionapi.Database, error) {
	// TODO
	return nil, nil
}

func TestGetDatabaseByID(t *testing.T) {
	tests := []struct {
		name        string
		databaseid  string
		filePath    string
		wantErr     bool
		emptyResult bool
		err         error
	}{
		{
			name:        "Get existing Database",
			databaseid:  "3caeee7e-2884-4d17-a911-a17b5d1b92d4",
			filePath:    DATABASE_JSON,
			wantErr:     false,
			emptyResult: false,
			err:         nil,
		},
		{
			name:        "Get non-existing Database",
			databaseid:  "3caeee7e-2774-4d17-a911-a17b5d1b92da",
			filePath:    DATABASE_JSON,
			wantErr:     true,
			emptyResult: true,
			err:         errors.New(DATABASE_NOT_EXIST_ERROR_STR),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client := GetMockedDatabaseService(t, test.filePath, test.err)
			database, err := client.GetDatabaseByID(context.Background(), notionclient.DatabaseID(test.databaseid))
			if test.wantErr {
				assert.Nil(t, database)
				assert.NotNil(t, err)
			} else {
				if test.emptyResult {
					assert.Empty(t, database)
				} else {
					assert.NotEmpty(t, database)
				}
				assert.Nil(t, err)
			}
		})

	}
}
