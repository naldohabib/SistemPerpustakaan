package repo

import (
	"Portofolio/SistemPerpustakaan/category"
	model2 "Portofolio/SistemPerpustakaan/model"
	"fmt"

	"github.com/jinzhu/gorm"
)

type CategoryRepoImpl struct {
	db *gorm.DB
}

// CreateCategoryRepoImpl ...
func CreateCategoryRepoImpl(db *gorm.DB) category.CategoryUsecase {
	return &CategoryRepoImpl{db}
}

func (c CategoryRepoImpl) Insert(data *model2.Category) (*model2.Category, error) {
	err := c.db.Save(&data).Error
	if err != nil {
		return nil, fmt.Errorf("[CategoryRepoImpl.Insert] Error when query save data with: %w\n", err)
	}
	return data, nil
}

func (c CategoryRepoImpl) GetByID(id int) (*model2.Category, error) {
	var data = model2.Category{}
	err := c.db.First(&data, id).Error
	if err != nil {
		return nil, fmt.Errorf("CategoryRepoImpl.GetByID Error when query get by id with error: %w\n", err)
	}
	return &data, nil
}

func (c CategoryRepoImpl) GetAll() (*[]model2.Category, error) {
	var data []model2.Category
	err := c.db.Find(&data).Error
	if err != nil {
		return nil, fmt.Errorf("[CategoryRepoImpl.GetAll] Error when query get all data with error: %w\n", err)
	}
	return &data, nil
}

func (c CategoryRepoImpl) Update(id int, data *model2.Category) (*model2.Category, error) {
	err := c.db.Model(&data).Where("id=?", id).Update(data).Error
	if err != nil {
		return nil, fmt.Errorf("CategoryRepoImpl.Update Error when query update data with error: %w\n", err)
	}
	return data, nil
}

func (c CategoryRepoImpl) Delete(id int) error {
	data := model2.Category{}
	err := c.db.Where("id=?", id).Delete(&data).Error
	if err != nil {
		return fmt.Errorf("[CategoryRepoImpl.Delete] Error when query delete data with error: %w\n", err)
	}
	return nil
}
