package repository

import "gorm.io/gorm"

type IconRepo interface {
	GetQuery(query string) (icon []Icon, err error)
	Create(icon Icon) (err error)
}

type iconRepo struct {
	db *gorm.DB
}

func NewIconRepo(db *gorm.DB) IconRepo {
	return iconRepo{db}
}

func (obj iconRepo) GetQuery(query string) (icon []Icon, err error) {
	err = obj.db.Raw(query).Scan(&icon).Error
	return icon, err
}

func (obj iconRepo) Create(icon Icon) (err error) {
	err = obj.db.Create(&icon).Error
	return err
}
