package repository

import "github.com/stretchr/testify/mock"

type iconRepoMock struct {
	mock.Mock
}

func NewIconRepoMock() *iconRepoMock {
	return &iconRepoMock{}
}

func (obj *iconRepoMock) GetQuery(query string) (icon []Icon, err error) {
	args := obj.Called(query)
	return args.Get(0).([]Icon), args.Error(1)
}

func (obj *iconRepoMock) Create(icon Icon) (err error) {
	args := obj.Called(icon)
	return args.Error(0)
}
