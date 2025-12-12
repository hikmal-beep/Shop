package http

import (
	"Shop/helper"
	"Shop/middleware"
	"Shop/module/product/service"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	service  service.ProductService
	validate *validator.Validate
}

func NewProductHandler(service service.ProductService) *ProductHandler {
	return &ProductHandler{
		service:  service,
		validate: validator.New(),
	}
}

func (h *ProductHandler) GetProduct(c *fiber.Ctx) error {
	productID, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		response := helper.APIResponse("Invalid product ID", fiber.StatusBadRequest, "error", nil)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	product, err := h.service.FindByID(c.Context(), productID)
	if err != nil {
		response := helper.APIResponse("Failed to get product", fiber.StatusInternalServerError, "error", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	if product == nil {
		response := helper.APIResponse("Product not found", fiber.StatusNotFound, "error", nil)
		return c.Status(fiber.StatusNotFound).JSON(response)
	}

	productResponse := ProductResponse{
		ID:          product.ID,
		ShopID:      product.ShopID,
		Product:     product.Product,
		Description: product.Description,
		Quantity:    product.Quantity,
	}

	response := helper.APIResponse("Product retrieved successfully", fiber.StatusOK, "success", productResponse)
	return c.Status(fiber.StatusOK).JSON(response)
}

func (h *ProductHandler) GetProductsByShop(c *fiber.Ctx) error {
	shopID, err := strconv.ParseInt(c.Params("shop_id"), 10, 64)
	if err != nil {
		response := helper.APIResponse("Invalid shop ID", fiber.StatusBadRequest, "error", nil)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	products, err := h.service.FindByShopID(c.Context(), shopID)
	if err != nil {
		response := helper.APIResponse("Failed to get products", fiber.StatusInternalServerError, "error", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	var productResponses []ProductResponse
	for _, product := range products {
		productResponses = append(productResponses, ProductResponse{
			ID:          product.ID,
			ShopID:      product.ShopID,
			Product:     product.Product,
			Description: product.Description,
			Quantity:    product.Quantity,
		})
	}

	response := helper.APIResponse("Products retrieved successfully", fiber.StatusOK, "success", productResponses)
	return c.Status(fiber.StatusOK).JSON(response)
}

func (h *ProductHandler) CreateProduct(c *fiber.Ctx) error {
	userID, err := middleware.GetUserIDFromToken(c)
	if err != nil {
		response := helper.APIResponse("Unauthorized", fiber.StatusUnauthorized, "error", err.Error())
		return c.Status(fiber.StatusUnauthorized).JSON(response)
	}

	var req CreateProductRequest
	if err := req.Bind(c); err != nil {
		response := helper.APIResponse("Invalid request", fiber.StatusBadRequest, "error", nil)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	if err := h.validate.Struct(req); err != nil {
		errors := helper.ErrorValidationFormat(err)
		response := helper.APIResponse("Validation failed", fiber.StatusBadRequest, "error", errors)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	product, err := h.service.Create(c.Context(), userID, service.CreateProductData{
		ShopID:      req.ShopID,
		Product:     req.Product,
		Description: req.Description,
		Quantity:    req.Quantity,
	})

	if err != nil {
		response := helper.APIResponse("Failed to create product", fiber.StatusInternalServerError, "error", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	productResponse := ProductResponse{
		ID:          product.ID,
		ShopID:      product.ShopID,
		Product:     product.Product,
		Description: product.Description,
		Quantity:    product.Quantity,
	}

	response := helper.APIResponse("Product created successfully", fiber.StatusCreated, "success", productResponse)
	return c.Status(fiber.StatusCreated).JSON(response)
}

func (h *ProductHandler) UpdateProduct(c *fiber.Ctx) error {
	userID, err := middleware.GetUserIDFromToken(c)
	if err != nil {
		response := helper.APIResponse("Unauthorized", fiber.StatusUnauthorized, "error", err.Error())
		return c.Status(fiber.StatusUnauthorized).JSON(response)
	}

	productID, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		response := helper.APIResponse("Invalid product ID", fiber.StatusBadRequest, "error", nil)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	var req UpdateProductRequest
	if err := req.Bind(c); err != nil {
		response := helper.APIResponse("Invalid request", fiber.StatusBadRequest, "error", nil)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	if err := h.validate.Struct(req); err != nil {
		errors := helper.ErrorValidationFormat(err)
		response := helper.APIResponse("Validation failed", fiber.StatusBadRequest, "error", errors)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	product, err := h.service.Update(c.Context(), userID, service.UpdateProductData{
		ID:          productID,
		Product:     req.Product,
		Description: req.Description,
		Quantity:    req.Quantity,
	})

	if err != nil {
		response := helper.APIResponse("Failed to update product", fiber.StatusInternalServerError, "error", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	productResponse := ProductResponse{
		ID:          product.ID,
		ShopID:      product.ShopID,
		Product:     product.Product,
		Description: product.Description,
		Quantity:    product.Quantity,
	}

	response := helper.APIResponse("Product updated successfully", fiber.StatusOK, "success", productResponse)
	return c.Status(fiber.StatusOK).JSON(response)
}

func (h *ProductHandler) DeleteProduct(c *fiber.Ctx) error {
	userID, err := middleware.GetUserIDFromToken(c)
	if err != nil {
		response := helper.APIResponse("Unauthorized", fiber.StatusUnauthorized, "error", err.Error())
		return c.Status(fiber.StatusUnauthorized).JSON(response)
	}

	productID, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		response := helper.APIResponse("Invalid product ID", fiber.StatusBadRequest, "error", nil)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	err = h.service.Delete(c.Context(), userID, productID)
	if err != nil {
		response := helper.APIResponse("Failed to delete product", fiber.StatusInternalServerError, "error", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response := helper.APIResponse("Product deleted successfully", fiber.StatusOK, "success", nil)
	return c.Status(fiber.StatusOK).JSON(response)
}