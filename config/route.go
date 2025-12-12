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

	routes.AuthRouter(app, db)
	routes.ShopRouter(app, db)
	routes.ProductRouter(app, db)

	log.Fatalln(app.Listen(":" + os.Getenv("PORT")))
}