package routes

import (
	"Shop/module/shop/http"
	"Shop/module/shop/repository"
	"Shop/module/shop/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ShopRouter(app *fiber.App, db *gorm.DB) {
	// Initialize layers: Repository → Service → Handler
	shopRepo := repository.NewShopRepository(db)
	shopService := service.NewShopService(shopRepo)
	shopHandler := http.NewShopHandler(shopService)

	// Register routes
	http.ShopRoutes(app, shopHandler)
}

