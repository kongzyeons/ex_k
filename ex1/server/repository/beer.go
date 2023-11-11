package repository

import (
	"fmt"

	"gorm.io/gorm"
)

type BeerRepo interface {
	Create(beer Beer) (err error)
	// GetAll() (beers []Beer, err error)
	GetPagination(offset, limit int) (beers []Beer, err error)
	GetQuery(query string) (beer []Beer, err error)
	Delete(beer_id int) (err error)
	Update(beer Beer) (err error)
}

type beerRepo struct {
	db *gorm.DB
}

func NewBeerRepo(db *gorm.DB) BeerRepo {
	return beerRepo{db}
}

func (obj beerRepo) Create(beer Beer) (err error) {
	err = obj.db.Create(&beer).Error
	return err
}
func (obj beerRepo) GetPagination(offset, limit int) (beers []Beer, err error) {
	err = obj.db.Offset(offset).Limit(limit).Find(&beers).Error
	return beers, err
}
func (obj beerRepo) GetQuery(query string) (beer []Beer, err error) {
	err = obj.db.Raw(query).Scan(&beer).Error
	return beer, err
}
func (obj beerRepo) Delete(beer_id int) (err error) {
	queryDelete := "Delete From beers Where id = ?"
	err = obj.db.Exec(queryDelete, beer_id).Error
	return err
}
func (obj beerRepo) Update(beer Beer) (err error) {
	queryUpdate := fmt.Sprintf("Update beers Set name ='%v', category = '%v' , detail = '%v' Where id = '%v'",
		beer.Name, beer.Category, beer.Detail, beer.ID,
	)
	err = obj.db.Exec(queryUpdate).Error
	return err
}
