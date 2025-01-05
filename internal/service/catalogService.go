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
func (s CatalogService) CreateProduct(input dto.CreateProductRequest, user domain.User) error {
	err := s.Repo.CreateProduct(&domain.Product{
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
		CategoryId:  input.CategoryId,
		ImageUrl:    input.ImageUrl,
		UserId:      int(user.ID),
		Stock:       uint(input.Stock),
	})
	return err
}
func (s CatalogService) EditProduct(id int, input dto.CreateProductRequest, user domain.User) (*domain.Product, error) {
	existingProduct, err := s.Repo.FindProductById(id)

	if err != nil {
		return nil, errors.New("Product does not found")
	}

	if existingProduct.UserId != int(user.ID) {
		return nil, errors.New("User does not match")
	}

	if len(input.Name) > 0 {
		existingProduct.Name = input.Name
	}
	if input.Description != "" {
		existingProduct.Description = input.Description
	}
	if input.Price > 0 {
		existingProduct.Price = input.Price
	}
	updateProduct, err := s.Repo.EditProduct(existingProduct)
	if err != nil {
		return nil, errors.New("failed to update product")
	}
	return updateProduct, nil
}
func (s CatalogService) DeleteProduct(id int, user domain.User) error {
	existingProduct, err := s.Repo.FindProductById(id)
	if err != nil {
		return errors.New("Product does not found")
	}
	if existingProduct.UserId != int(user.ID) {
		return errors.New("User don't have rights of the product")
	}

	err = s.Repo.DeleteProduct(existingProduct)
	if err != nil {
		return errors.New("failed to delete product")
	}
	return nil
}

func (s CatalogService) GetProducts() ([]*domain.Product, error) {
	products, err := s.Repo.FindProducts()
	if err != nil {
		return nil, errors.New("failed to find products")
	}
	return products, err
}
func (s CatalogService) GetProductById(id int) (*domain.Product, error) {
	product, err := s.Repo.FindProductById(id)
	if err != nil {
		return nil, errors.New("Product does not found")
	}
	return product, err
}
func (s CatalogService) GetSellerProducts(id int) ([]*domain.Product, error) {
	products, err := s.Repo.FindSellerProducts(id)
	if err != nil {
		return nil, errors.New("failed to find seller products")
	}
	return products, err
}

func (s CatalogService) UpdateProductStock(e domain.Product) (*domain.Product, error) {
	product, err := s.Repo.FindProductById(int(e.ID))
	if err != nil {
		return nil, errors.New("Product not found")
	}

	if product.UserId != e.UserId {
		return nil, errors.New("you don't have permission to change stock of this product")
	}

	product.Stock = e.Stock

	editProduct, err := s.Repo.EditProduct(product)
	if err != nil {
		return nil, err
	}
	return editProduct, nil

}
