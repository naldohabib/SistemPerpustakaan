package model

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/jinzhu/gorm"
)

// var IndonesianPhoneRegex = regexp.MustCompile(`^([\+]{1}628[1-35-9][0-9]{7,10})$`)

// User ...
type User struct {
	gorm.Model
	Username string `gorm:"not null;size:255" json:"username"`
	Email    string `gorm:"not null;size:255" json:"email"`
	Password string `gorm:"not null;size:500" json:"password"`
	Role     string `gorm:"size:15;DEFAULT:'user'" json:"role"`
}

// TableName ..
func (u User) TableName() string {
	return "tb_user"
}

// Validate ...
func (u *User) Validate() error {
	if err := validation.Validate(u.Username, validation.Required); err != nil {
		return errors.New("Username cannot be blank")
	}

	if err := validation.Validate(u.Email, validation.Required); err != nil {
		return errors.New("Email Cannot be blank and must be email")
	}

	if err := validation.Validate(u.Password, validation.Required); err != nil {
		return errors.New("Password Cannot be blank")
	}

	return nil
}
