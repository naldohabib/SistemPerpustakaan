package driver

import (
	"Portofolio/SistemPerpustakaan/model"

	"github.com/jinzhu/gorm"
)

// InitTable ...
func InitTable(db *gorm.DB) {
	db.Debug().AutoMigrate(
		// &model.Category{},
		&model.Book{},
		&model.User{},
		&model.BookDetail{},
		&model.Transaksi{},
	)
	db.Model(&model.BookDetail{}).AddForeignKey("book_id", "tb_book(id)", "CASCADE", "CASCADE")
	// db.Model(&model.Transaksi{}).AddForeignKey("user_id", "tb_user(id)", "CASCADE", "CASCADE")
	// db.Model(&model.Transaksi{}).AddForeignKey("book_id", "tb_book(id)", "CASCADE", "CASCADE")
	// db.Model(&model.Transaksi{}).AddForeignKey("book_detail_id", "tb_book_detail(id)", "CASCADE", "CASCADE")
	// db.Model(&model.Book{}).AddForeignKey("category_id", "tb_category(id)", "CASCADE", "CASCADE")
}
