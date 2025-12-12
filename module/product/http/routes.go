package http

import (
	"Shop/middleware"

	"github.com/gofiber/fiber/v2"
)

func ProductRoutes(app *fiber.App, handler *ProductHandler) {
	products := app.Group("/products")

	products.Get("/:id", handler.GetProduct)
	products.Get("/shop/:shop_id", handler.GetProductsByShop)

	products.Use(middleware.JWTAuthMiddleware())
	products.Post("/", handler.CreateProduct)
	products.Put("/:id", handler.UpdateProduct)
	products.Delete("/:id", handler.DeleteProduct)
}