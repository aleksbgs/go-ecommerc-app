package api

import (
	"go-ecommerce-app/config"
	"go-ecommerce-app/internal/api/rest"
	"go-ecommerce-app/internal/api/rest/handlers"
	"go-ecommerce-app/internal/domain"
	"go-ecommerce-app/internal/helper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"

	"github.com/gofiber/fiber/v2"
)

func StartServer(config config.AppConfig) {
	app := fiber.New()

	db, err := gorm.Open(postgres.Open(config.Dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("database connection error %v\n", err)
	}

	log.Printf("database connection established successfully")

	//run migration
	err = db.AutoMigrate(&domain.User{}, &domain.BankAccount{}, &domain.Category{}, &domain.Product{})

	if err != nil {
		log.Fatalf("database migration error %v\n", err)
	}
	log.Printf("database migration successfully")

	auth := helper.SetupAuth(config.AppSecret)

	rh := &rest.RestHandler{
		App:    app,
		DB:     db,
		Auth:   auth,
		Config: config,
	}
	setupRoutes(rh)

	//routes
	app.Listen(config.ServerPort)
}

func setupRoutes(rh *rest.RestHandler) {
	//user handler
	handlers.SetupUserRoutes(rh)
	handlers.SetupCatalogRoutes(rh)
	//transactions

	//

}
