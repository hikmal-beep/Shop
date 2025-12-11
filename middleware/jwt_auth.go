package middleware

import (
	"os"
	"strconv"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Missing or malformed JWT"})
	}
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid or expired JWT"})
}

func JWTAuthMiddleware() fiber.Handler {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "A_VERY_STRONG_JWT_SECRET_KEY_12345_SHOP" // Fallback secret
	}
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(jwtSecret)},
		ContextKey: "user",
		ErrorHandler: ErrorHandler,
	})
}

func GetUserIDFromToken(c *fiber.Ctx) (int64, error) {
	token, ok := c.Locals("user").(*jwt.Token)
	if !ok {
		return 0, fiber.NewError(fiber.StatusUnauthorized, "Token not found in context")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, fiber.NewError(fiber.StatusInternalServerError, "Invalid token Claims")
	}
	userID, ok := claims["user_id"].(float64)
	if !ok {
		return 0, fiber.NewError(fiber.StatusInternalServerError, "Invalid user_id in token")
	}
	return int64(userID), nil
}

func GetUserIDStringFromToken(c *fiber.Ctx) (string, error) {
	userID, err := GetUserIDFromToken(c)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(userID, 10), nil
}