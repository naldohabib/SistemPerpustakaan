package model

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/jinzhu/gorm"
)

// Book ...
type Book struct {
	gorm.Model
	// CategoryID uint   `gorm:"not null" json:"category_id"`
	Title        string `gorm:"size:255" json:"title"`
	Author       string `gorm:"size:255" json:"author"`
	Publisher    string `gorm:"size:255" json:"publisher"`
	BookYear     int    `gorm:"size:255" json:"book_year"`
	Synopsis     string `gorm:"size:255" json:"synopsis"`
	Image        string `json:"image"`
	QtyBook      int    `json:"qty_book"`
	QtyAvailable int    `json:"qty_available"`
}

// TableName ...
func (c *Book) TableName() string {
	return "tb_book"
}

// Validate ...
func (c *Book) Validate() error {
	// if err := validation.Validate(c.CategoryID, validation.Required); err != nil {
	// 	return errors.New("category id cannot be blank")
	// }

	if err := validation.Validate(c.Title, validation.Required); err != nil {
		return errors.New("title cannot be blank")
	}

	if err := validation.Validate(c.Author, validation.Required); err != nil {
		return errors.New("author cannot be blank")
	}

	if err := validation.Validate(c.Publisher, validation.Required); err != nil {
		return errors.New("publisher cannot be blank")
	}

	if err := validation.Validate(c.BookYear, validation.Required); err != nil {
		return errors.New("book year cannot be blank")
	}

	// if err := validation.Validate(c.BookYear, validation.Length(4, 4)); err != nil {
	// 	return errors.New("book year cannot is not valid")
	// }

	if err := validation.Validate(c.Synopsis, validation.Required); err != nil {
		return errors.New("synopsis cannot be blank")
	}

	if err := validation.Validate(c.QtyBook, validation.Required); err != nil {
		return errors.New("Qtybook cannot be blank")
	}
	return nil

}
