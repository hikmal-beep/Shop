package routes

import (
	"Shop/module/product/http"
	"Shop/module/product/repository"
	"Shop/module/product/service"
	shopRepository "Shop/module/shop/repository"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ProductRouter(app *fiber.App, db *gorm.DB) {
	productRepo := repository.NewProductRepository(db)
	shopRepo := shopRepository.NewShopRepository(db)
	productService := service.NewProductService(productRepo, shopRepo)
	productHandler := http.NewProductHandler(productService)

	http.ProductRoutes(app, productHandler)
}

