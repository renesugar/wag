// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

package server

import (
	context "context"
	models "github.com/Clever/wag/samples/gen-go/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockController is a mock of Controller interface
type MockController struct {
	ctrl     *gomock.Controller
	recorder *MockControllerMockRecorder
}

// MockControllerMockRecorder is the mock recorder for MockController
type MockControllerMockRecorder struct {
	mock *MockController
}

// NewMockController creates a new mock instance
func NewMockController(ctrl *gomock.Controller) *MockController {
	mock := &MockController{ctrl: ctrl}
	mock.recorder = &MockControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockController) EXPECT() *MockControllerMockRecorder {
	return _m.recorder
}

// GetAuthors mocks base method
func (_m *MockController) GetAuthors(ctx context.Context, i *models.GetAuthorsInput) (*models.AuthorsResponse, string, error) {
	ret := _m.ctrl.Call(_m, "GetAuthors", ctx, i)
	ret0, _ := ret[0].(*models.AuthorsResponse)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetAuthors indicates an expected call of GetAuthors
func (_mr *MockControllerMockRecorder) GetAuthors(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetAuthors", reflect.TypeOf((*MockController)(nil).GetAuthors), arg0, arg1)
}

// GetAuthorsWithPut mocks base method
func (_m *MockController) GetAuthorsWithPut(ctx context.Context, i *models.GetAuthorsWithPutInput) (*models.AuthorsResponse, string, error) {
	ret := _m.ctrl.Call(_m, "GetAuthorsWithPut", ctx, i)
	ret0, _ := ret[0].(*models.AuthorsResponse)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetAuthorsWithPut indicates an expected call of GetAuthorsWithPut
func (_mr *MockControllerMockRecorder) GetAuthorsWithPut(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetAuthorsWithPut", reflect.TypeOf((*MockController)(nil).GetAuthorsWithPut), arg0, arg1)
}

// GetBooks mocks base method
func (_m *MockController) GetBooks(ctx context.Context, i *models.GetBooksInput) ([]models.Book, int64, error) {
	ret := _m.ctrl.Call(_m, "GetBooks", ctx, i)
	ret0, _ := ret[0].([]models.Book)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetBooks indicates an expected call of GetBooks
func (_mr *MockControllerMockRecorder) GetBooks(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetBooks", reflect.TypeOf((*MockController)(nil).GetBooks), arg0, arg1)
}

// CreateBook mocks base method
func (_m *MockController) CreateBook(ctx context.Context, i *models.Book) (*models.Book, error) {
	ret := _m.ctrl.Call(_m, "CreateBook", ctx, i)
	ret0, _ := ret[0].(*models.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBook indicates an expected call of CreateBook
func (_mr *MockControllerMockRecorder) CreateBook(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "CreateBook", reflect.TypeOf((*MockController)(nil).CreateBook), arg0, arg1)
}

// PutBook mocks base method
func (_m *MockController) PutBook(ctx context.Context, i *models.Book) (*models.Book, error) {
	ret := _m.ctrl.Call(_m, "PutBook", ctx, i)
	ret0, _ := ret[0].(*models.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutBook indicates an expected call of PutBook
func (_mr *MockControllerMockRecorder) PutBook(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "PutBook", reflect.TypeOf((*MockController)(nil).PutBook), arg0, arg1)
}

// GetBookByID mocks base method
func (_m *MockController) GetBookByID(ctx context.Context, i *models.GetBookByIDInput) (*models.Book, error) {
	ret := _m.ctrl.Call(_m, "GetBookByID", ctx, i)
	ret0, _ := ret[0].(*models.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBookByID indicates an expected call of GetBookByID
func (_mr *MockControllerMockRecorder) GetBookByID(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetBookByID", reflect.TypeOf((*MockController)(nil).GetBookByID), arg0, arg1)
}

// GetBookByID2 mocks base method
func (_m *MockController) GetBookByID2(ctx context.Context, id string) (*models.Book, error) {
	ret := _m.ctrl.Call(_m, "GetBookByID2", ctx, id)
	ret0, _ := ret[0].(*models.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBookByID2 indicates an expected call of GetBookByID2
func (_mr *MockControllerMockRecorder) GetBookByID2(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetBookByID2", reflect.TypeOf((*MockController)(nil).GetBookByID2), arg0, arg1)
}

// GetBookByIDCached mocks base method
func (_m *MockController) GetBookByIDCached(ctx context.Context, id string) (*models.Book, error) {
	ret := _m.ctrl.Call(_m, "GetBookByIDCached", ctx, id)
	ret0, _ := ret[0].(*models.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBookByIDCached indicates an expected call of GetBookByIDCached
func (_mr *MockControllerMockRecorder) GetBookByIDCached(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetBookByIDCached", reflect.TypeOf((*MockController)(nil).GetBookByIDCached), arg0, arg1)
}

// HealthCheck mocks base method
func (_m *MockController) HealthCheck(ctx context.Context) error {
	ret := _m.ctrl.Call(_m, "HealthCheck", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// HealthCheck indicates an expected call of HealthCheck
func (_mr *MockControllerMockRecorder) HealthCheck(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "HealthCheck", reflect.TypeOf((*MockController)(nil).HealthCheck), arg0)
}
