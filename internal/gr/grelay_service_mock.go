package gr

import "github.com/stretchr/testify/mock"

type GrelayServiceMock struct {
	mock.Mock
}

func (m *GrelayServiceMock) exec(f func() (interface{}, error)) (interface{}, error) {
	args := m.Called(f)
	return args.Get(0), args.Error(1)
}
