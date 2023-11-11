package repository

type Icon struct {
	ID       int    `gorm:"primaryKey" json:"id"`
	Name     string `gorm:"column:name" json:"name"`
	Children []Children
}
