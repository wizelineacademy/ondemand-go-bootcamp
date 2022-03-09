package db

import (
	"github.com/stretchr/testify/mock"
)

type MockStore struct {
	mock.Mock
}

func (m *MockStore) ReadData() ([][]string, error) {
	ret := m.Called()

	r0 := ret.Get(0).([][]string)
	r1 := ret.Error(1)

	return r0, r1
}
