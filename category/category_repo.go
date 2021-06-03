package category

import (
	model2 "Portofolio/SistemPerpustakaan/model"
)

type CategoryRepo interface {
	Insert(data *model2.Category) (*model2.Category, error)
	GetByID(id int) (*model2.Category, error)
	GetAll() (*[]model2.Category, error)
	Update(id int, data *model2.Category) (*model2.Category, error)
	Delete(id int) error
}
