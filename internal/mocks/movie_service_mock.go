// Code generated by MockGen. DO NOT EDIT.
// Source: internal/services/movie_service.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/stwrtrio/movie-festival-api/internal/models"
)

// MockMovieService is a mock of MovieService interface.
type MockMovieService struct {
	ctrl     *gomock.Controller
	recorder *MockMovieServiceMockRecorder
}

// MockMovieServiceMockRecorder is the mock recorder for MockMovieService.
type MockMovieServiceMockRecorder struct {
	mock *MockMovieService
}

// NewMockMovieService creates a new mock instance.
func NewMockMovieService(ctrl *gomock.Controller) *MockMovieService {
	mock := &MockMovieService{ctrl: ctrl}
	mock.recorder = &MockMovieServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMovieService) EXPECT() *MockMovieServiceMockRecorder {
	return m.recorder
}

// CreateMovie mocks base method.
func (m *MockMovieService) CreateMovie(ctx context.Context, movie *models.Movie) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMovie", ctx, movie)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateMovie indicates an expected call of CreateMovie.
func (mr *MockMovieServiceMockRecorder) CreateMovie(ctx, movie interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMovie", reflect.TypeOf((*MockMovieService)(nil).CreateMovie), ctx, movie)
}

// GetMovies mocks base method.
func (m *MockMovieService) GetMovies(ctx context.Context, pagination models.PaginationRequest, useCache bool) (*models.PaginationResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMovies", ctx, pagination, useCache)
	ret0, _ := ret[0].(*models.PaginationResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMovies indicates an expected call of GetMovies.
func (mr *MockMovieServiceMockRecorder) GetMovies(ctx, pagination, useCache interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMovies", reflect.TypeOf((*MockMovieService)(nil).GetMovies), ctx, pagination, useCache)
}

// UpdateMovie mocks base method.
func (m *MockMovieService) UpdateMovie(ctx context.Context, movie *models.Movie) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMovie", ctx, movie)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateMovie indicates an expected call of UpdateMovie.
func (mr *MockMovieServiceMockRecorder) UpdateMovie(ctx, movie interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMovie", reflect.TypeOf((*MockMovieService)(nil).UpdateMovie), ctx, movie)
}

// getMoviesFromCache mocks base method.
func (m *MockMovieService) getMoviesFromCache(ctx context.Context, key string) (*models.PaginationResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getMoviesFromCache", ctx, key)
	ret0, _ := ret[0].(*models.PaginationResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// getMoviesFromCache indicates an expected call of getMoviesFromCache.
func (mr *MockMovieServiceMockRecorder) getMoviesFromCache(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getMoviesFromCache", reflect.TypeOf((*MockMovieService)(nil).getMoviesFromCache), ctx, key)
}
