package repository

import "github.com/stretchr/testify/mock"

type beerRepoMock struct {
	mock.Mock
}

func NewBeerRepoMock() *beerRepoMock {
	return &beerRepoMock{}
}

func (obj *beerRepoMock) Create(beer Beer) (err error) {
	args := obj.Called(beer)
	return args.Error(0)
}

// func (obj *beerRepoMock) GetAll() (beers []Beer, err error) {
// 	args := obj.Called()
// 	return args.Get(0).([]Beer), args.Error(1)
// }

func (obj *beerRepoMock) GetPagination(offset, limit int) (beers []Beer, err error) {
	args := obj.Called(offset, limit)
	return args.Get(0).([]Beer), args.Error(1)
}

func (obj *beerRepoMock) GetQuery(query string) (beer []Beer, err error) {
	args := obj.Called(query)
	return args.Get(0).([]Beer), args.Error(1)
}

func (obj *beerRepoMock) Delete(beer_id int) (err error) {
	args := obj.Called(beer_id)
	return args.Error(0)
}

func (obj *beerRepoMock) Update(beer Beer) (err error) {
	args := obj.Called(beer)
	return args.Error(0)
}
