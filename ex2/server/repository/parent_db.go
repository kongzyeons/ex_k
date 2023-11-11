package repository

type Parent struct {
	ID       int    `gorm:"primaryKey"`
	Name     string `gorm:"varchar"`
	Route    string `gorm:"varchar"`
	Children []Children
}
