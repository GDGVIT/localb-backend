package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/rithikjain/local-businesses-backend/api/view"
	"github.com/rithikjain/local-businesses-backend/pkg"
	"log"
	"os"
)

// Protected routes
func Protected() func(*fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte(os.Getenv("jwtSecret")),
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Missing or malformed JWT"})

	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid or expired JWT"})
	}
}

func ValidateAndGetClaims(c *fiber.Ctx, role string) (map[string]interface{}, error) {
	token, ok := c.Locals("user").(*jwt.Token)
	if !ok {
		log.Println(token)
		return nil, view.ErrInvalidToken
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		log.Println(claims)
		return nil, view.ErrInvalidToken
	}

	if claims.Valid() != nil {
		return nil, view.ErrInvalidToken
	}

	if claims["role"].(string) != role {
		log.Println(claims["role"])
		return nil, pkg.ErrUnauthorized
	}
	return claims, nil
}
