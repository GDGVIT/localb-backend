package handler

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/rithikjain/local-businesses-backend/api/middleware"
	"github.com/rithikjain/local-businesses-backend/api/view"
	"github.com/rithikjain/local-businesses-backend/pkg/admin"
	"github.com/rithikjain/local-businesses-backend/pkg/models"
	"os"
)

func Login(svc admin.Service) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		admin := &models.Admin{}
		if err := c.BodyParser(admin); err != nil {
			return view.Wrap(err, c)
		}

		ad, err := svc.Login(admin.Username, admin.Password)
		if err != nil {
			return view.Wrap(err, c)
		}
		ad.Password = ""

		// Handling JWT
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id":   ad.ID,
			"role": "admin",
		})
		tokenString, err := token.SignedString([]byte(os.Getenv("jwtSecret")))
		if err != nil {
			return view.Wrap(err, c)
		}

		return c.JSON(fiber.Map{
			"message": "Admin logged in",
			"token":   tokenString,
			"admin":   ad,
		})
	}
}

func ShowBusinessesToApprove(svc admin.Service) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		_, err := middleware.ValidateAndGetClaims(c, "admin")
		if err != nil {
			return view.Wrap(err, c)
		}

		bizs, err := svc.GetBusinessesToApprove()
		if err != nil {
			return view.Wrap(err, c)
		}

		return c.JSON(fiber.Map{
			"message":    "Businesses to approve fetched",
			"businesses": bizs,
		})
	}
}

func ApproveBusiness(svc admin.Service) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		_, err := middleware.ValidateAndGetClaims(c, "admin")
		if err != nil {
			return view.Wrap(err, c)
		}

		bizBody := &models.GetBusinessID{}
		if err := c.BodyParser(bizBody); err != nil {
			return view.Wrap(err, c)
		}

		err = svc.ApproveBusiness(bizBody.BusinessID)
		if err != nil {
			return view.Wrap(err, c)
		}

		return c.JSON(fiber.Map{
			"message": "Business approved",
		})
	}
}

func DeleteBusiness(svc admin.Service) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		_, err := middleware.ValidateAndGetClaims(c, "admin")
		if err != nil {
			return view.Wrap(err, c)
		}

		bizBody := &models.GetBusinessID{}
		if err := c.BodyParser(bizBody); err != nil {
			return view.Wrap(err, c)
		}

		err = svc.DeleteBusiness(bizBody.BusinessID)
		if err != nil {
			return view.Wrap(err, c)
		}

		return c.JSON(fiber.Map{
			"message": "Business deleted",
		})
	}
}

func MakeAdminHandler(app *fiber.App, svc admin.Service) {
	adminGroup := app.Group("/api/v1/admin")
	adminGroup.Post("/login", Login(svc))
	adminGroup.Get("/businessesToApprove", middleware.Protected(), ShowBusinessesToApprove(svc))
	adminGroup.Post("/approveBusiness", middleware.Protected(), ApproveBusiness(svc))
	adminGroup.Post("/deleteBusiness", middleware.Protected(), DeleteBusiness(svc))
}
