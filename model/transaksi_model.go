package model

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/jinzhu/gorm"
)

// Transaksi ...
type Transaksi struct {
	gorm.Model
	UserID        uint   `gorm:"not null" json:"user_id"`
	BookID        uint   `gorm:"not null" json:"book_id"`
	BookDetailID  uint   `json:"book_detail_id"`
	NameUser      string `json:"name_user,omitempty"`
	NameBook      string `json:"name_book,omitempty"`
	TimeStamp     int64  `json:"time_st'amp"`
	DateFineStamp int64  `json:"date_fine_stamp"`
	JumlahHari    int64  `gorm:"not null" json:"jumlah_hari"`
	Status        string `gorm:"DEFAULT:'waiting'" json:"status"`
	TotalDenda    int    `json:"total_denda"`
}

// TableName ...
func (c *Transaksi) TableName() string {
	return "tb_transaksi"
}

// Validate ...
func (u *Transaksi) Validate() error {
	if err := validation.Validate(u.BookID, validation.Required); err != nil {
		return errors.New("BookID cannot be blank")
	}

	if err := validation.Validate(u.UserID, validation.Required); err != nil {
		return errors.New("UserID cannot be blank")
	}

	if err := validation.Validate(u.JumlahHari, validation.Required); err != nil {
		return errors.New("JumlahHari cannot be blank")
	}

	return nil
}
