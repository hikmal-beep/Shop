package http

import (
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App, handler *AuthHandler) {
	auth := app.Group("/auth")
	auth.Post("/register", handler.Register)
	auth.Post("/login", handler.Login)
}