package usecase

import (
	"Portofolio/SistemPerpustakaan/model"
	"Portofolio/SistemPerpustakaan/user"
	"errors"
	"fmt"
)

//  UserUsecaseImpl
type UserUsecaseImpl struct {
	userRepo user.UserRepo
}

// CreateUserUsecaseImpl ...
func CreateUserUsecaseImpl(userRepo user.UserRepo) user.UserUsecase {
	return &UserUsecaseImpl{userRepo}
}

func (u UserUsecaseImpl) Insert(data *model.User) (*model.User, error) {
	if err := data.Validate(); err != nil {
		return nil, err
	}

	user, err := u.userRepo.Insert(data)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u UserUsecaseImpl) GetByID(id int) (*model.User, error) {
	user, err := u.userRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u UserUsecaseImpl) GetAll() (*[]model.User, error) {
	user, err := u.userRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u UserUsecaseImpl) Update(id int, data *model.User) (*model.User, error) {
	_, err := u.userRepo.GetByID(id)
	if err != nil {
		return nil, errors.New("userID does not exist")
	}

	if err := data.Validate(); err != nil {
		return nil, err
	}

	user, err := u.userRepo.Update(id, data)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u UserUsecaseImpl) Delete(id int) error {
	_, err := u.userRepo.GetByID(id)
	if err != nil {
		return errors.New("userID does not exist")
	}

	err = u.userRepo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func (u UserUsecaseImpl) CheckMail(email string) bool {
	return u.userRepo.CheckMail(email)
}

func (u UserUsecaseImpl) GetUserByEmail(email string) (*model.User, string, error) {
	checkEmail := u.userRepo.CheckMail(email)

	if checkEmail == false {
		return nil, "Email Not Found", nil
	}
	fmt.Println(checkEmail)

	users, err := u.userRepo.GetUserByEmail(email)

	return users, "", err
}
