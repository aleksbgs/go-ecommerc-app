package api

import (
	"go-ecommerce-app/config"
	"go-ecommerce-app/internal/api/rest"
	"go-ecommerce-app/internal/api/rest/handlers"
	"go-ecommerce-app/internal/domain"
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
	db.AutoMigrate(&domain.User{})

	rh := &rest.RestHandler{
		App: app,
		DB:  db,
	}
	setupRoutes(rh)

	//routes
	app.Listen(config.ServerPort)
}

func setupRoutes(rh *rest.RestHandler) {
	//user handler
	handlers.SetupUserRoutes(rh)

	//transactions

	//

}
