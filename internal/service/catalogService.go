package service

import (
	"errors"
	"go-ecommerce-app/config"
	"go-ecommerce-app/internal/domain"
	"go-ecommerce-app/internal/dto"
	"go-ecommerce-app/internal/helper"
	"go-ecommerce-app/internal/repository"
)

type CatalogService struct {
	Repo   repository.CatalogRepository
	Auth   helper.Auth
	Config config.AppConfig
}

func (s CatalogService) CreateCategory(input dto.CreateCategoryRequest) error {

	err := s.Repo.CreateCategory(&domain.Category{
		Name:         input.Name,
		ImageUrl:     input.ImageUrl,
		DisplayOrder: input.DisplayOrder,
	})

	return err
}
func (s CatalogService) EditCategory(id int, input dto.CreateCategoryRequest) (*domain.Category, error) {
	existingCat, err := s.Repo.FindCategoryById(id)
	if err != nil {
		return nil, errors.New("Category does not found")
	}

	if len(input.Name) > 0 {
		existingCat.Name = input.Name
	}
	if input.ParentId > 0 {
		existingCat.ParentId = input.ParentId
	}
	if len(input.ImageUrl) > 0 {
		existingCat.ImageUrl = input.ImageUrl
	}
	if input.DisplayOrder > 0 {
		existingCat.DisplayOrder = input.DisplayOrder
	}

	updateCat, err := s.Repo.EditCategory(existingCat)

	return updateCat, nil
}
func (s CatalogService) DeleteCategories(id int) error {
	err := s.Repo.DeleteCategory(id)
	if err != nil {
		return errors.New("Category does not found")
	}
	return nil
}
func (s CatalogService) GetCategories() ([]*domain.Category, error) {
	categories, err := s.Repo.FindCategories()
	if err != nil {
		return nil, errors.New("failed to find categories")
	}
	return categories, err

}

func (s CatalogService) GetCategory(id int) (*domain.Category, error) {
	cat, err := s.Repo.FindCategoryById(id)
	if err != nil {
		return nil, errors.New("failed to find category")
	}

	return cat, nil
}
