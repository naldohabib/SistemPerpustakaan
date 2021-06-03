package model

import (
	"errors"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/jinzhu/gorm"
)

// Category ...
type Category struct {
	gorm.Model
	Name string `gorm:"size:255" json:"name"`
}

// TableName ...
func (c Category)TableName() string {
	return "tb_category"
}

func (c Category)Validate() error {
	if err := validation.Validate(c.Name, validation.Required); err != nil {
		return errors.New("category name cannot be blank")
	}

	if err:= validation.Validate(c.Name, validation.Length(1,12)); err!= nil{
		return errors.New("category name minimum length and maximum 12")
	}

	return nil
}
