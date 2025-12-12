package http

import "github.com/gofiber/fiber/v2"

type ProductResponse struct {
	ID          int64  `json:"id"`
	ShopID      int64  `json:"shop_id"`
	Product     string `json:"product"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}

type CreateProductRequest struct {
	ShopID      int64  `json:"shop_id" validate:"required"`
	Product     string `json:"product" validate:"required"`
	Description string `json:"description" validate:"required"`
	Quantity    int    `json:"quantity" validate:"required,min=0"`
}

func (r *CreateProductRequest) Bind(c *fiber.Ctx) error {
	return c.BodyParser(r)
}

type UpdateProductRequest struct {
	Product     string `json:"product" validate:"required"`
	Description string `json:"description" validate:"required"`
	Quantity    int    `json:"quantity" validate:"required,min=0"`
}

func (r *UpdateProductRequest) Bind(c *fiber.Ctx) error {
	return c.BodyParser(r)
}