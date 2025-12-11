package http

import (
	"Shop/helper"
	"Shop/middleware"
	"Shop/module/shop/service"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ShopHandler struct {
	service  service.ShopService
	validate *validator.Validate
}

func NewShopHandler(service service.ShopService) *ShopHandler {
	return &ShopHandler{
		service:  service,
		validate: validator.New(),
	}
}

// GetMyShops - GET /shops/me (Get current user's shops)
func (h *ShopHandler) GetMyShop(c *fiber.Ctx) error {
	// Get user ID from JWT token
	userID, err := middleware.GetUserIDFromToken(c)
	if err != nil {
		response := helper.APIResponse("Unauthorized", fiber.StatusUnauthorized, "error", err.Error())
		return c.Status(fiber.StatusUnauthorized).JSON(response)
	}

	// Get user's shops
	shops, err := h.service.FindByUserID(c.Context(), userID)
	if err != nil {
		response := helper.APIResponse("Failed to get shops", fiber.StatusInternalServerError, "error", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	// Convert to response format
	var shopResponses []ShopResponse
	for _, shop := range shops {
		shopResponses = append(shopResponses, ShopResponse{
			ID:      shop.ID,
			UserID:  shop.UserID,
			Name:    shop.Name,
			Address: shop.Address,
		})
	}

	response := helper.APIResponse("Shops retrieved successfully", fiber.StatusOK, "success", shopResponses)
	return c.Status(fiber.StatusOK).JSON(response)
}

// CreateShop - POST /shops
func (h *ShopHandler) CreateShop(c *fiber.Ctx) error {
	var req CreateShopRequest

	// Get user ID from JWT token
	userID, err := middleware.GetUserIDFromToken(c)
	if err != nil {
		response := helper.APIResponse("Unauthorized", fiber.StatusUnauthorized, "error", err.Error())
		return c.Status(fiber.StatusUnauthorized).JSON(response)
	}

	// Parse request body
	if err := c.BodyParser(&req); err != nil {
		response := helper.APIResponse("Invalid request", fiber.StatusBadRequest, "error", nil)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	// Validate request
	if err := h.validate.Struct(req); err != nil {
		errors := helper.ErrorValidationFormat(err)
		response := helper.APIResponse("Validation failed", fiber.StatusBadRequest, "error", errors)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	// Create shop
	shop, err := h.service.Create(c.Context(), userID, service.CreateShopData{
		Name:    req.Name,
		Address: req.Address,
	})

	if err != nil {
		response := helper.APIResponse("Failed to create shop", fiber.StatusInternalServerError, "error", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	shopResponse := ShopResponse{
		ID:      shop.ID,
		UserID:  shop.UserID,
		Name:    shop.Name,
		Address: shop.Address,
	}

	response := helper.APIResponse("Shop created successfully", fiber.StatusCreated, "success", shopResponse)
	return c.Status(fiber.StatusCreated).JSON(response)
}

// UpdateShop - PUT /shops/:id
func (h *ShopHandler) UpdateShop(c *fiber.Ctx) error {
	var req UpdateShopRequest

	// Get user ID from JWT token
	userID, err := middleware.GetUserIDFromToken(c)
	if err != nil {
		response := helper.APIResponse("Unauthorized", fiber.StatusUnauthorized, "error", err.Error())
		return c.Status(fiber.StatusUnauthorized).JSON(response)
	}

	// Get shop ID from URL
	shopID, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		response := helper.APIResponse("Invalid shop ID", fiber.StatusBadRequest, "error", nil)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	// Parse request body
	if err := c.BodyParser(&req); err != nil {
		response := helper.APIResponse("Invalid request", fiber.StatusBadRequest, "error", nil)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	// Validate request
	if err := h.validate.Struct(req); err != nil {
		errors := helper.ErrorValidationFormat(err)
		response := helper.APIResponse("Validation failed", fiber.StatusBadRequest, "error", errors)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	// Update shop
	shop, err := h.service.Update(c.Context(), userID, service.UpdateShopData{
		ID:      shopID,
		Name:    req.Name,
		Address: req.Address,
	})

	if err != nil {
		response := helper.APIResponse("Failed to update shop", fiber.StatusInternalServerError, "error", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	shopResponse := ShopResponse{
		ID:      shop.ID,
		UserID:  shop.UserID,
		Name:    shop.Name,
		Address: shop.Address,
	}

	response := helper.APIResponse("Shop updated successfully", fiber.StatusOK, "success", shopResponse)
	return c.Status(fiber.StatusOK).JSON(response)
}

// DeleteShop - DELETE /shops/:id
func (h *ShopHandler) DeleteShop(c *fiber.Ctx) error {
	// Get user ID from JWT token
	userID, err := middleware.GetUserIDFromToken(c)
	if err != nil {
		response := helper.APIResponse("Unauthorized", fiber.StatusUnauthorized, "error", err.Error())
		return c.Status(fiber.StatusUnauthorized).JSON(response)
	}

	// Get shop ID from URL
	shopID, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		response := helper.APIResponse("Invalid shop ID", fiber.StatusBadRequest, "error", nil)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	// Delete shop
	err = h.service.Delete(c.Context(), userID, shopID)
	if err != nil {
		response := helper.APIResponse("Failed to delete shop", fiber.StatusInternalServerError, "error", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response := helper.APIResponse("Shop deleted successfully", fiber.StatusOK, "success", nil)
	return c.Status(fiber.StatusOK).JSON(response)
}