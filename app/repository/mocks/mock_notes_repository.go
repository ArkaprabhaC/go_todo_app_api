// Code generated by MockGen. DO NOT EDIT.
// Source: notes_repository.go

// Package repository_mock is a generated GoMock package.
package repository_mock

import (
	context "context"
	reflect "reflect"

	db_model "github.com/ArkaprabhaC/go_todo_app_api/app/model/db"
	gomock "github.com/golang/mock/gomock"
)

// MockNotesRepository is a mock of NotesRepository interface.
type MockNotesRepository struct {
	ctrl     *gomock.Controller
	recorder *MockNotesRepositoryMockRecorder
}

// MockNotesRepositoryMockRecorder is the mock recorder for MockNotesRepository.
type MockNotesRepositoryMockRecorder struct {
	mock *MockNotesRepository
}

// NewMockNotesRepository creates a new mock instance.
func NewMockNotesRepository(ctrl *gomock.Controller) *MockNotesRepository {
	mock := &MockNotesRepository{ctrl: ctrl}
	mock.recorder = &MockNotesRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNotesRepository) EXPECT() *MockNotesRepositoryMockRecorder {
	return m.recorder
}

// AddNote mocks base method.
func (m *MockNotesRepository) AddNote(ctx context.Context, note db_model.Note) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddNote", ctx, note)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddNote indicates an expected call of AddNote.
func (mr *MockNotesRepositoryMockRecorder) AddNote(ctx, note interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddNote", reflect.TypeOf((*MockNotesRepository)(nil).AddNote), ctx, note)
}

// GetNotes mocks base method.
func (m *MockNotesRepository) GetNotes(ctx context.Context) ([]db_model.Note, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNotes", ctx)
	ret0, _ := ret[0].([]db_model.Note)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNotes indicates an expected call of GetNotes.
func (mr *MockNotesRepositoryMockRecorder) GetNotes(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNotes", reflect.TypeOf((*MockNotesRepository)(nil).GetNotes), ctx)
}
