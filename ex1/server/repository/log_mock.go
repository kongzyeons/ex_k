package repository

import "github.com/stretchr/testify/mock"

type logRepoMock struct {
	mock.Mock
}

func NewLogRepoMock() *logRepoMock {
	return &logRepoMock{}
}

func (obj *logRepoMock) CreateLog(bearLog BearLog) (errr error) {
	args := obj.Called(bearLog)
	return args.Error(0)
}
