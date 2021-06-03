package model

import "github.com/jinzhu/gorm"

// BookDetail ...
type BookDetail struct {
	gorm.Model
	StatusBook string `gorm:"not null;DEFAULT:'available'" json:"book_status"`
	BookID     uint   `gorm:"not null" json:"book_id"`
}

// TableName ...
func (c *BookDetail) TableName() string {
	return "tb_book_detail"
}
