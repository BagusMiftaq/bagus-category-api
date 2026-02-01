package services

import (
	"bagus-category-api/model"
	"bagus-category-api/model/repositories"
)

type CategoryService struct {
	repo *repositories.CategoryRepository
}

func NewCategoryService(repo *repositories.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (service *CategoryService) GetCategory() ([]model.Category, error) {
	return service.repo.GetCategory()
}

func (service *CategoryService) AddCategory(data *model.Category) error {
	return service.repo.AddCategory(data)
}

func (service *CategoryService) GetCategoryById(id int) (*model.Category, error) {
	return service.repo.GetCategoryById(id)
}

func (service *CategoryService) UpdateCategory(data *model.Category) error {
	return service.repo.UpdateCategory(data)
}

func (service *CategoryService) DeleteCategory(id int) error {
	return service.repo.DeleteCategory(id)
}
