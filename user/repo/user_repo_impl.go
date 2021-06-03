package repo

import (
	"Portofolio/SistemPerpustakaan/model"
	"Portofolio/SistemPerpustakaan/user"
	"fmt"

	"github.com/jinzhu/gorm"
)

//  UserRepoImpl ...
type UserRepoImpl struct {
	db *gorm.DB
}

// CreateUserRepoImpl ...
func CreateUserRepoImpl(db *gorm.DB) user.UserRepo {
	return &UserRepoImpl{db}
}

// Insert ...
func (c *UserRepoImpl) Insert(data *model.User) (*model.User, error) {
	err := c.db.Save(&data).Error
	if err != nil {
		return nil, fmt.Errorf("[UserRepoImpl.Insert] Error when query save data with : %w", err)
	}
	return data, nil
}

// GetByID ...
func (c *UserRepoImpl) GetByID(id int) (*model.User, error) {
	var data = model.User{}
	err := c.db.First(&data, id).Error
	if err != nil {
		return nil, fmt.Errorf("UserRepoImpl.GetByID Error when query get by id with error: %w", err)
	}
	return &data, nil
}

// GetAll ...
func (c *UserRepoImpl) GetAll() (*[]model.User, error) {
	var data []model.User
	err := c.db.Find(&data).Error
	if err != nil {
		return nil, fmt.Errorf("[UserRepoImpl.GetAll] Error when query get all data with error: %w", err)
	}
	return &data, nil
}

// Update ...
func (c *UserRepoImpl) Update(id int, data *model.User) (*model.User, error) {
	err := c.db.Model(&data).Where("id=?", id).Update(data).Error
	if err != nil {
		return nil, fmt.Errorf("UserRepoImpl.Update Error when query update data with error: %w", err)
	}
	return data, nil
}

// Delete ...
func (c *UserRepoImpl) Delete(id int) error {
	data := model.User{}
	err := c.db.Where("id=?", id).Delete(&data).Error
	if err != nil {
		return fmt.Errorf("[UserRepoImpl.Delete] Error when query delete data with error: %w", err)
	}
	return nil
}

// CheckMail ...
func (c *UserRepoImpl) CheckMail(email string) bool {
	var total int

	c.db.Debug().Table("tb_user").Where("email = ?", email).Count(&total)
	fmt.Println(total)
	if total > 0 {
		return true
	}
	return false
}

// GetUserByEmail ...
func (c *UserRepoImpl) GetUserByEmail(email string) (*model.User, error) {
	var data model.User
	err := c.db.Debug().Where("email = ?", email).Find(&data).Error
	fmt.Println(data.Email)
	if err != nil {
		return nil, fmt.Errorf("[UserRepoImpl.LoginUser] Error when query with error : %v", err)
	}

	return &data, nil
}
