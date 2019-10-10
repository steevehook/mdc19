package controllers

import "github.com/stretchr/testify/mock"

type commentsServiceMock struct {
	mock.Mock
}

func (m *commentsServiceMock) GetComments() ([]string, error) {
	args := m.Called()
	return args.Get(0).([]string), args.Error(1)
}
