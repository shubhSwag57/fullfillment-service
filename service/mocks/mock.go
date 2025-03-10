package mocks

import (
	"github.com/stretchr/testify/mock"
)

type MockDB struct {
	mock.Mock
}

func (m *MockDB) Where(query interface{}, args ...interface{}) *MockDB {
	args = append([]interface{}{query}, args...)
	m.Called(args...)
	return m
}

func (m *MockDB) First(out interface{}, where ...interface{}) *MockDB {
	m.Called(out, where)
	return m
}

func (m *MockDB) Create(value interface{}) *MockDB {
	m.Called(value)
	return m
}

func (m *MockDB) Model(value interface{}) *MockDB {
	m.Called(value)
	return m
}

func (m *MockDB) Update(column string, value interface{}) *MockDB {
	m.Called(column, value)
	return m
}
