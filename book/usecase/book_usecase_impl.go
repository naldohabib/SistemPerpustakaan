package usecase

import (
	"Portofolio/SistemPerpustakaan/book"
	model "Portofolio/SistemPerpustakaan/model"
	"errors"

	validation "github.com/go-ozzo/ozzo-validation"
)

// BookUsecaseImpl ...
type BookUsecaseImpl struct {
	bookRepo book.BookRepo
}

// CreateBookUsecaseImpl ...
func CreateBookUsecaseImpl(bookRepo book.BookRepo) book.BookUsecase {
	return &BookUsecaseImpl{bookRepo}
}

// Insert ...
func (b *BookUsecaseImpl) Insert(data *model.Book) (*model.Book, error) {
	if err := data.Validate(); err != nil {
		return nil, err
	}

	data, err := b.bookRepo.Insert(data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// GetAll ...
func (b *BookUsecaseImpl) GetAll() (*[]model.Book, error) {
	data, err := b.bookRepo.GetAll()
	if err != nil {
		return nil, err
	}

	return data, nil
}

// GetBookByID ...
func (b *BookUsecaseImpl) GetBookByID(id int) (*model.Book, error) {
	data, err := b.bookRepo.GetBookByID(id)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// GetAllBookDetail ...
func (b *BookUsecaseImpl) GetAllBookDetail(id int) (*[]model.BookDetail, error) {
	_, err := b.bookRepo.GetBookByID(id)
	if err != nil {
		return nil, errors.New("category ID does not exist")
	}

	data, err := b.bookRepo.GetAllBookDetail(id)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// Delete ...
func (b *BookUsecaseImpl) Delete(id int) error {
	_, err := b.bookRepo.GetBookByID(id)
	if err != nil {
		return errors.New("bookID does not exist")
	}

	err = b.bookRepo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

// Update ...
func (b *BookUsecaseImpl) Update(id int, data *model.Book) (*model.Book, error) {
	firstData, err := b.bookRepo.GetBookByID(id)
	if err != nil {
		return nil, errors.New("bookID does not exist")
	}

	data.QtyBook = firstData.QtyBook
	data.QtyAvailable = firstData.QtyAvailable

	if data.Image == "" {
		data.Image = firstData.Image
	}
	if data.Title == "" {
		data.Title = firstData.Title
	}
	if data.Author == "" {
		data.Author = firstData.Author
	}
	if data.Publisher == "" {
		data.Publisher = firstData.Publisher
	}
	if data.Synopsis == "" {
		data.Synopsis = firstData.Synopsis
	}
	if err := validation.Validate(data.BookYear, validation.Required); err != nil {
		return nil, errors.New("book year cannot be blank")
	}

	dataFix, err := b.bookRepo.Update(id, data)
	if err != nil {
		return nil, err
	}

	return dataFix, nil
}
