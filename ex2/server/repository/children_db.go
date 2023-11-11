package repository

type Children struct {
	ID         int    `gorm:"primaryKey" json:"id"`
	ParentID   int    `gorm:"column:parent_id" json:"parent_id"`
	Parent     Parent `gorm:"foreignKey:ParentID"`
	IconID     int    `gorm:"column:icon_id" json:"icon_id"`
	Icon       Icon   `gorm:"foreignKey:IconID"`
	Name       string `gorm:"column:name" json:"name"`
	Route      string `gorm:"column:route" json:"route"`
	IsChildren bool   `gorm:"column:is_children" json:"is_children"`
}
