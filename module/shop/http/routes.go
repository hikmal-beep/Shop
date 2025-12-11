package http

import (
	"Shop/middleware"

	"github.com/gofiber/fiber/v2"
)

func ShopRoutes(app *fiber.App, handler *ShopHandler) {
	shops := app.Group("/shops", middleware.JWTAuthMiddleware())
	shops.Get("/me", handler.GetMyShop)
	shops.Post("/", handler.CreateShop)
	shops.Put("/:id", handler.UpdateShop)
	shops.Delete("/:id", handler.DeleteShop)  
}