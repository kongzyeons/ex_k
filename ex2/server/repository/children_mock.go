package repository

import "github.com/stretchr/testify/mock"

type childrenRepoMock struct {
	mock.Mock
}

func NewChildrenRepoMock() *childrenRepoMock {
	return &childrenRepoMock{}
}

func (obj *childrenRepoMock) GetQuery(query string) (children []Children, err error) {
	args := obj.Called(query)
	return args.Get(0).([]Children), args.Error(1)
}

func (obj *childrenRepoMock) Create(children Children) (err error) {
	args := obj.Called(children)
	return args.Error(0)
}
