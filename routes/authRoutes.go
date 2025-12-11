package routes

import (
	"Shop/module/auth/http"
	"Shop/module/auth/repository"
	"Shop/module/auth/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AuthRouter(app *fiber.App, db *gorm.DB) {
	AuthRepo := repository.NewUserRepository(db)
	AuthService := service.NewAuthService(AuthRepo)
	AuthHandler := http.NewAuthHandler(AuthService)

	http.AuthRoutes(app, AuthHandler)
}
