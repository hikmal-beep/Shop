package http

import (
	"github.com/gofiber/fiber/v2"
)

// The base User data structure (analogous to ProductDetail)
type UserData struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// --- REGISTER Handlers ---
type (
	RegisterRequest struct {
		Name     string `json:"name" validate:"required"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=6"`
	}
	RegisterResponse struct {
		Token string   `json:"token"`
		User  UserData `json:"user"`
	}
)

func (r *RegisterRequest) Bind(c *fiber.Ctx) error {
	// Assuming BodyParser for standard POST request body
	return c.BodyParser(r)
}

// --- LOGIN Handlers ---
type (
	LoginRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}
	LoginResponse struct {
		Token string   `json:"token"`
		User  UserData `json:"user"`
	}
)

func (r *LoginRequest) Bind(c *fiber.Ctx) error {
	// Assuming BodyParser for standard POST request body
	return c.BodyParser(r)
}