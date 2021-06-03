package book

import (
	"Portofolio/SistemPerpustakaan/model"
)

type BookUsecase interface {
	Insert(data *model.Book) (*model.Book, error)
	GetAll() (*[]model.Book, error)
	GetBookByID(id int) (*model.Book, error)
	GetAllBookDetail(id int) (*[]model.BookDetail, error)
	Delete(id int) error
	Update(id int, data *model.Book) (*model.Book, error)
}
