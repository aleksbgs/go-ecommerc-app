package handlers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"go-ecommerce-app/internal/api/rest"
	"go-ecommerce-app/internal/dto"
	"go-ecommerce-app/internal/repository"
	"go-ecommerce-app/internal/service"
	"log"
	"net/http"
)

type UserHandler struct {
	//svc UserService
	svc service.UserService
}

func SetupUserRoutes(rh *rest.RestHandler) {
	app := rh.App

	//create an instance of user service & incejt to
	svc := service.UserService{
		Repo:   repository.NewUserRepository(rh.DB),
		CRepo:  repository.NewCatalogRepository(rh.DB),
		Auth:   rh.Auth,
		Config: rh.Config,
	}

	handler := UserHandler{
		svc: svc,
	}

	pubRoutes := app.Group("/users")
	// Public endpoints
	pubRoutes.Post("/register", handler.Register)
	pubRoutes.Post("/login", handler.Login)

	pvtRoutes := pubRoutes.Group("/", rh.Auth.Authorize)

	// Private endpoint
	pvtRoutes.Get("/verify", handler.GetVerificationCode)
	pvtRoutes.Post("/verify", handler.Verify)

	pvtRoutes.Post("/profile", handler.CreateProfile)
	pvtRoutes.Get("/profile", handler.GetProfile)

	pvtRoutes.Patch("/profile", handler.UpdateProfile)

	pvtRoutes.Post("/cart", handler.AddToCart)
	pvtRoutes.Get("/cart", handler.GetCart)

	pvtRoutes.Get("/order", handler.GetOrders)
	pvtRoutes.Get("/order/:id", handler.GetOrder)

	pvtRoutes.Post("/order", handler.CreateOrder)

	pvtRoutes.Post("/become-seller", handler.BecomeSeller)

}

func (h *UserHandler) Register(ctx *fiber.Ctx) error {

	user := dto.UserSignup{}

	err := ctx.BodyParser(&user)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "please provide a valid input",
		})
	}

	token, err := h.svc.Signup(user)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "register",
		"token":   token,
	})
}

func (h *UserHandler) Login(ctx *fiber.Ctx) error {
	loginInput := dto.UserLogin{}

	err := ctx.BodyParser(&loginInput)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "please provide a valid input",
		})
	}

	token, err := h.svc.Login(loginInput.Email, loginInput.Password)

	if err != nil {
		return ctx.Status(http.StatusUnauthorized).JSON(&fiber.Map{
			"message": "error on Login",
		})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "login success",
		"token":   token,
	})
}

func (u *UserHandler) GetVerificationCode(ctx *fiber.Ctx) error {

	user := u.svc.Auth.GetCurrentUser(ctx)

	//create verification code and update to user profile IN Db

	err := u.svc.GetVerificationCode(user)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "get verification code",
	})
}
func (u *UserHandler) Verify(ctx *fiber.Ctx) error {

	user := u.svc.Auth.GetCurrentUser(ctx)

	var req dto.VerificationCodeInput

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "please provide a valid input",
		})
	}

	err := u.svc.VerifyCode(user.ID, req.Code)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "verify success",
	})
}

func (u *UserHandler) CreateProfile(ctx *fiber.Ctx) error {

	user := u.svc.Auth.GetCurrentUser(ctx)

	req := dto.ProfileInput{}
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{})
	}

	log.Printf("User %v", user)

	err := u.svc.CreateProfile(user.ID, req)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "create profile",
	})

}

func (u *UserHandler) GetProfile(ctx *fiber.Ctx) error {

	user := u.svc.Auth.GetCurrentUser(ctx)
	log.Printf("User %v", user)

	//call user service and perform get profile
	profile, err := u.svc.GetProfile(user.ID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "get profile",
		"profile": profile,
	})
}

func (u *UserHandler) AddToCart(ctx *fiber.Ctx) error {

	req := dto.CreateCartRequest{}
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "please provide a valid product and qty",
		})
	}
	user := u.svc.Auth.GetCurrentUser(ctx)

	log.Println(user)

	//call user service and perform create cart
	cartItems, err := u.svc.CreateCart(req, user)

	if err != nil {
		return rest.InternalServerError(ctx, err)

	}
	return rest.SuccessResponses(ctx, "cart created successfully", cartItems)

}
func (u *UserHandler) GetCart(ctx *fiber.Ctx) error {

	user := u.svc.Auth.GetCurrentUser(ctx)

	cart, err := u.svc.FindCart(user.ID)

	if err != nil {
		return rest.InternalServerError(ctx, errors.New("cart not found"))
	}
	return rest.SuccessResponses(ctx, "cart found successfully", cart)

}

func (u *UserHandler) CreateOrder(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "create Order",
	})
}

func (u *UserHandler) GetOrders(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "get orders",
	})
}
func (u *UserHandler) GetOrder(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "get order by id",
	})
}

func (u *UserHandler) BecomeSeller(ctx *fiber.Ctx) error {

	user := u.svc.Auth.GetCurrentUser(ctx)

	req := dto.SellerInput{}
	err := ctx.BodyParser(&req)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "please provide a valid input",
		})
	}

	token, err := u.svc.BecomeSeller(user.ID, req)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "failed to become seller",
		})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "become a seller",
		"token":   token,
	})
}
func (h *UserHandler) UpdateProfile(ctx *fiber.Ctx) error {

	user := h.svc.Auth.GetCurrentUser(ctx)
	req := dto.ProfileInput{}
	err := ctx.BodyParser(&req)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{})
	}

	err = h.svc.UpdateProfile(user.ID, req)

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "update profile success",
	})
}
