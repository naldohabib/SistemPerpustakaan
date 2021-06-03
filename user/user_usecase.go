package user

import "Portofolio/SistemPerpustakaan/model"

// UserUsecase
type UserUsecase interface {
	Insert(data *model.User) (*model.User, error)
	GetByID(id int) (*model.User, error)
	GetAll() (*[]model.User, error)
	Update(id int, data *model.User) (*model.User, error)
	Delete(id int) error
	CheckMail(email string) bool
	GetUserByEmail(email string) (*model.User, string, error)
}
