package repository

import (
	"github.com/stretchr/testify/mock"
)

type parentRepoMock struct {
	mock.Mock
}

func NewParentRepoMock() *parentRepoMock {
	return &parentRepoMock{}
}

func (obj *parentRepoMock) GetQuery(query string) (parent []Parent, err error) {
	args := obj.Called(query)
	return args.Get(0).([]Parent), args.Error(1)
}

func (obj *parentRepoMock) Create(parent Parent) (err error) {
	args := obj.Called(parent)
	return args.Error(0)
}
