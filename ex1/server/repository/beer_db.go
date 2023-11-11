package repository

type Beer struct {
	ID       int    `gorm:"primaryKey" json:"id"`
	Name     string `gorm:"column:name" json:"name"`
	Category string `gorm:"column:category" json:"category"`
	Detail   string `gorm:"column:detail" json:"detail"`
}
