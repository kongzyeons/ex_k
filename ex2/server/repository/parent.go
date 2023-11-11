package repository

import "gorm.io/gorm"

type ParentRepo interface {
	GetQuery(query string) (parent []Parent, err error)
	Create(parent Parent) (err error)
}

type parentRepo struct {
	db *gorm.DB
}

func NewParentRepo(db *gorm.DB) ParentRepo {
	return parentRepo{db}
}

func (obj parentRepo) GetQuery(query string) (parent []Parent, err error) {
	err = obj.db.Raw(query).Scan(&parent).Error
	return parent, err
}

func (obj parentRepo) Create(parent Parent) (err error) {
	err = obj.db.Create(&parent).Error
	return err
}
