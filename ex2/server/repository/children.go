package repository

import "gorm.io/gorm"

type ChildrenRepo interface {
	GetQuery(query string) (children []Children, err error)
	Create(children Children) (err error)
}

type childrenRepo struct {
	db *gorm.DB
}

func NewChildrenRepo(db *gorm.DB) ChildrenRepo {
	return childrenRepo{db}
}

func (obj childrenRepo) GetQuery(query string) (children []Children, err error) {
	err = obj.db.Raw(query).Scan(&children).Error
	return children, err
}

func (obj childrenRepo) Create(children Children) (err error) {
	err = obj.db.Create(&children).Error
	return err
}
