package config

import (
	"Shop/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/gorm"
)


func Route(db *gorm.DB) {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
	}))

	// Public routes (no JWT required)
	routes.AuthRouter(app, db)

	// Protected routes (JWT required)
	routes.ShopRouter(app, db)

	log.Fatalln(app.Listen(":" + os.Getenv("PORT")))
}