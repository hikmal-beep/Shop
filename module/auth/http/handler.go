package http

import (
	"Shop/helper"
	"Shop/module/auth/service"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	service  service.UserService
	validate *validator.Validate
}

func NewAuthHandler(service service.UserService) *AuthHandler {
	return &AuthHandler{
		service:  service,
		validate: validator.New(),
	}
}

// Register - POST /auth/register
func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var req RegisterRequest

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

	// Register user
	user, err := h.service.Register(c.Context(), service.RegisterUserData{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		// Check if it's a duplicate email error
		if err.Error() == "email already exists" {
			response := helper.APIResponse("Email already registered", fiber.StatusConflict, "error", "This email is already in use")
			return c.Status(fiber.StatusConflict).JSON(response)
		}
		response := helper.APIResponse("Failed to register user", fiber.StatusInternalServerError, "error", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	// Return response without token (user must login to get token)
	userData := UserData{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	response := helper.APIResponse("User registered successfully. Please login to continue.", fiber.StatusCreated, "success", userData)
	return c.Status(fiber.StatusCreated).JSON(response)
}

// Login - POST /auth/login
func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req LoginRequest

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

	// Login user
	user, err := h.service.Login(c.Context(), service.LoginUserData{
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		response := helper.APIResponse("Invalid email or password", fiber.StatusUnauthorized, "error", err.Error())
		return c.Status(fiber.StatusUnauthorized).JSON(response)
	}

	// Generate JWT token
	token, err := service.GenerateJWT(user.ID)
	if err != nil {
		response := helper.APIResponse("Failed to generate token", fiber.StatusInternalServerError, "error", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	// Return response with token
	authResponse := AuthResponse{
		Token: token,
		User: UserData{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
	}

	response := helper.APIResponse("Login successful", fiber.StatusOK, "success", authResponse)
	return c.Status(fiber.StatusOK).JSON(response)
}