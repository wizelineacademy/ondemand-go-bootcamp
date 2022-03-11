package db

import (
	"github.com/stretchr/testify/mock"
)

// MockStore is a mock of Store interface
type MockStore struct {
	mock.Mock
}

// ReadData mocks ReadData function
func (m MockStore) ReadData() ([][]string, error) {
	ret := m.Called()

	r0 := ret.Get(0).([][]string)
	r1 := ret.Error(1)

	return r0, r1
}
