package model

type ItemBook struct {
	ID      int    `gorm:"primaryKey" json:"id"`
	Name    string `gorm:"not null;type:varchar(225)" json:"name"`
	Genre   string `gorm:"not null;type:varchar(225)" json:"genre"`
	Author  string `gorm:"not null;type:varchar(225)" json:"author"`
	ShelfID int    `json:"shelf_id"`
}
