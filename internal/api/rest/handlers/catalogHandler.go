package handlers

import (
	"github.com/gofiber/fiber/v2"
	"go-ecommerce-app/internal/api/rest"
	"go-ecommerce-app/internal/domain"
	"go-ecommerce-app/internal/dto"
	"go-ecommerce-app/internal/repository"
	"go-ecommerce-app/internal/service"
	"strconv"
)

type CatalogHandler struct {
	//svc UserService
	svc service.CatalogService
}

func SetupCatalogRoutes(rh *rest.RestHandler) {
	app := rh.App

	// create an instance of user service & inject to handler
	svc := service.CatalogService{
		Repo:   repository.NewCatalogRepository(rh.DB),
		Auth:   rh.Auth,
		Config: rh.Config,
	}
	handler := CatalogHandler{
		svc: svc,
	}

	// public
	// listing products and categories
	app.Get("/products", handler.GetProducts)
	app.Get("/products/:id", handler.GetProduct)
	app.Get("/categories", handler.GetCategories)
	app.Get("/categories/:id", handler.GetCategoryById)

	// private
	// manage products and categories
	selRoutes := app.Group("/seller", rh.Auth.AuthorizeSeller)
	// Categories
	selRoutes.Post("/categories", handler.CreateCategories)
	selRoutes.Patch("/categories/:id", handler.EditCategory)
	selRoutes.Delete("/categories/:id", handler.DeleteCategory)

	// Products
	selRoutes.Post("/products", handler.CreateProduct)
	selRoutes.Get("/products", handler.GetProducts)
	selRoutes.Get("/products/:id", handler.GetProduct)
	selRoutes.Put("/products/:id", handler.EditProduct)
	selRoutes.Patch("/products/:id", handler.UpdateStock) // update stock
	selRoutes.Delete("/products/:id", handler.DeleteProduct)
}

func (h CatalogHandler) CreateCategories(ctx *fiber.Ctx) error {

	req := dto.CreateCategoryRequest{}
	err := ctx.BodyParser(&req)
	if err != nil {
		return rest.BadRequestError(ctx, "create category request body parse error")
	}

	err = h.svc.CreateCategory(req)
	if err != nil {
		return rest.InternalServerError(ctx, err)
	}

	return rest.SuccessResponses(ctx, "category created successfully", nil)
}
func (h CatalogHandler) EditCategory(ctx *fiber.Ctx) error {

	id, _ := strconv.Atoi(ctx.Params("id"))

	req := dto.CreateCategoryRequest{}
	err := ctx.BodyParser(&req)
	if err != nil {
		return rest.BadRequestError(ctx, "create category request body parse error")
	}
	updatedCat, err := h.svc.EditCategory(id, req)
	if err != nil {
		return rest.InternalServerError(ctx, err)
	}
	return rest.SuccessResponses(ctx, "edit category endpoint", updatedCat)
}
func (h CatalogHandler) DeleteCategory(ctx *fiber.Ctx) error {

	id, _ := strconv.Atoi(ctx.Params("id"))

	err := h.svc.DeleteCategories(id)
	if err != nil {
		return rest.InternalServerError(ctx, err)
	}
	return rest.SuccessResponses(ctx, "category deleted successfully", nil)
}

func (h CatalogHandler) CreateProduct(ctx *fiber.Ctx) error {

	req := dto.CreateProductRequest{}
	err := ctx.BodyParser(&req)
	if err != nil {
		return rest.BadRequestError(ctx, "create product request body parse error")
	}
	user := h.svc.Auth.GetCurrentUser(ctx)
	err = h.svc.CreateProduct(req, user)
	if err != nil {
		return rest.InternalServerError(ctx, err)
	}

	return rest.SuccessResponses(ctx, "product create successfully", nil)
}
func (h CatalogHandler) GetProducts(ctx *fiber.Ctx) error {

	products, err := h.svc.GetProducts()
	if err != nil {
		return rest.ErrorMessage(ctx, 404, err)
	}

	return rest.SuccessResponses(ctx, "products", products)
}

func (h CatalogHandler) GetProduct(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	product, err := h.svc.GetProductById(id)
	if err != nil {
		return rest.BadRequestError(ctx, "get product by id error")
	}

	return rest.SuccessResponses(ctx, "product", product)

}
func (h CatalogHandler) EditProduct(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
	req := dto.CreateProductRequest{}
	err := ctx.BodyParser(&req)
	if err != nil {
		return rest.BadRequestError(ctx, "create product request body parse error")
	}
	user := h.svc.Auth.GetCurrentUser(ctx)
	product, err := h.svc.EditProduct(id, req, user)
	if err != nil {
		return rest.InternalServerError(ctx, err)
	}
	return rest.SuccessResponses(ctx, "product edit successfully", product)

}
func (h CatalogHandler) UpdateStock(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
	req := dto.UpdateStockRequest{}
	err := ctx.BodyParser(&req)
	if err != nil {
		return rest.BadRequestError(ctx, "update stock request body parse error")
	}
	user := h.svc.Auth.GetCurrentUser(ctx)

	product := domain.Product{
		ID:     uint(id),
		Stock:  uint(req.Stock),
		UserId: int(user.ID),
	}
	updatedProduct, err := h.svc.UpdateProductStock(product)
	if err != nil {
		return rest.InternalServerError(ctx, err)
	}

	return rest.SuccessResponses(ctx, "update stock product endpoint", updatedProduct)
}

func (h CatalogHandler) DeleteProduct(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	user := h.svc.Auth.GetCurrentUser(ctx)
	err := h.svc.DeleteProduct(id, user)

	return rest.SuccessResponses(ctx, "delete product endpoint", err)

}

func (h CatalogHandler) GetCategories(ctx *fiber.Ctx) error {

	cats, err := h.svc.GetCategories()
	if err != nil {
		return rest.ErrorMessage(ctx, fiber.StatusNotFound, err)
	}
	return rest.SuccessResponses(ctx, "categories", cats)
}

func (h CatalogHandler) GetCategoryById(ctx *fiber.Ctx) error {

	id, _ := strconv.Atoi(ctx.Params("id"))

	cat, err := h.svc.GetCategory(id)

	if err != nil {
		return rest.ErrorMessage(ctx, fiber.StatusNotFound, err)
	}

	return rest.SuccessResponses(ctx, "get category endpoint", cat)
}
