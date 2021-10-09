package grelay

import "github.com/stretchr/testify/mock"

type grelayServiceMock struct {
	mock.Mock
}

func (m *grelayServiceMock) exec(f func() (interface{}, error)) (interface{}, error) {
	args := m.Called(f)
	return args.Get(0), args.Error(1)
}
