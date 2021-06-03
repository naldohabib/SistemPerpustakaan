package usecase

import (
	"Portofolio/SistemPerpustakaan/category"
	model2 "Portofolio/SistemPerpustakaan/model"
	"errors"
)

type CategoryUsecaseImpl struct {
	categoryRepo category.CategoryRepo
}

func CreateCategoryUsecaseImpl(categoryRepo category.CategoryRepo) category.CategoryUsecase {
	return &CategoryUsecaseImpl{categoryRepo}
}

func (c CategoryUsecaseImpl) Insert(data *model2.Category) (*model2.Category, error) {
	if err := data.Validate(); err != nil {
		return nil, err
	}

	category, err := c.categoryRepo.Insert(data)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (c CategoryUsecaseImpl) GetByID(id int) (*model2.Category, error) {
	category, err := c.categoryRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (c CategoryUsecaseImpl) GetAll() (*[]model2.Category, error) {
	category, err := c.categoryRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (c CategoryUsecaseImpl) Update(id int, data *model2.Category) (*model2.Category, error) {
	_, err := c.categoryRepo.GetByID(id)
	if err != nil {
		return nil, errors.New("category ID does not exist")
	}

	if err := data.Validate(); err != nil {
		return nil, err
	}

	category, err := c.categoryRepo.Update(id, data)
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (c CategoryUsecaseImpl) Delete(id int) error {
	_, err := c.categoryRepo.GetByID(id)
	if err != nil {
		return errors.New("category ID does not exist")
	}

	err = c.categoryRepo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
