package repo

import (
	"Portofolio/SistemPerpustakaan/book"
	"Portofolio/SistemPerpustakaan/model"
	"fmt"

	"github.com/jinzhu/gorm"
)

// BookRepoImpl ...
type BookRepoImpl struct {
	db *gorm.DB
}

// CreateBookRepoImpl ...
func CreateBookRepoImpl(db *gorm.DB) book.BookRepo {
	return &BookRepoImpl{db}
}

//BeginTrans ...
func (b *BookRepoImpl) BeginTrans() *gorm.DB {
	return b.db.Begin()
}

// Insert ...
func (b *BookRepoImpl) Insert(data *model.Book) (*model.Book, error) {

	tx := b.db.Begin()

	err := tx.Save(&data).Error
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("[BookRepoImpl.Insert] Error occured while inserting Book data to database : %w", err)
	}

	for i := 0; i < data.QtyBook; i++ {
		bookdetail := model.BookDetail{
			StatusBook: "available",
			BookID:     data.ID,
		}
		err := tx.Save(&bookdetail).Error
		if err != nil {
			tx.Rollback()
			return nil, fmt.Errorf("[BookRepoImpl.Insert] Error occured while inserting Book_Detail data to database : %w", err)
		}
	}

	tx.Commit()

	return data, nil
}

// GetAll ...
func (b *BookRepoImpl) GetAll() (*[]model.Book, error) {
	var data []model.Book
	err := b.db.Find(&data).Error
	if err != nil {
		return nil, fmt.Errorf("[BookRepoImpl.GetAll] Error when query get all data with error: %w", err)
	}
	return &data, nil
}

// GetBookByID ...
func (b *BookRepoImpl) GetBookByID(id int) (*model.Book, error) {
	var data = model.Book{}
	err := b.db.First(&data, id).Error
	if err != nil {
		return nil, fmt.Errorf("BookRepoImpl.GetByID Error when query get by id with error: %w", err)
	}
	return &data, nil
}

// GetAllBookDetail ...
func (b *BookRepoImpl) GetAllBookDetail(id int) (*[]model.BookDetail, error) {
	var data []model.BookDetail
	err := b.db.Where("book_id = ?", id).Find(&data).Error
	if err != nil {
		return nil, fmt.Errorf("[BookRepoImpl.GetAllBookDetail] Error when query get all data with error: %w", err)
	}
	return &data, nil
}

// Delete ...
func (b *BookRepoImpl) Delete(id int) error {
	data := model.Book{}
	dataDetailBook := model.BookDetail{}

	err := b.db.Where("book_id=?", id).Delete(&dataDetailBook).Error
	if err != nil {
		return fmt.Errorf("[BookRepoImpl.Delete] Error when query delete data with error: %w", err)
	}

	err = b.db.Where("id=?", id).Delete(&data).Error
	if err != nil {
		return fmt.Errorf("[BookRepoImpl.Delete] Error when query delete data with error: %w", err)
	}
	return nil
}

// Update ...
func (b *BookRepoImpl) Update(id int, data *model.Book) (*model.Book, error) {
	err := b.db.Model(&data).Where("id=?", id).Update(data).Error
	if err != nil {
		return nil, fmt.Errorf("BookRepoImpl.Update Error when query update data with error: %w", err)
	}
	return data, nil
}
