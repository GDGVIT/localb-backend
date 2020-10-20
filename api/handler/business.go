package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rithikjain/local-businesses-backend/api/view"
	"github.com/rithikjain/local-businesses-backend/pkg/business"
	"github.com/rithikjain/local-businesses-backend/pkg/models"
)

func NewBusiness(svc business.Service) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		biz := &models.Business{}
		if err := c.BodyParser(biz); err != nil {
			return view.Wrap(err, c)
		}

		err := svc.AddBusiness(biz)
		if err != nil {
			return view.Wrap(err, c)
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "Added new business",
		})
	}
}

func ShowBusinesses(svc business.Service) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		bizs, err := svc.GetApprovedBusinesses()
		if err != nil {
			return view.Wrap(err, c)
		}

		return c.JSON(fiber.Map{
			"message":    "Businesses fetched",
			"businesses": bizs,
		})
	}
}

func ShowBusinessesByCity(svc business.Service) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		city := c.Query("city")
		bizs, err := svc.GetBusinessesByCity(city)
		if err != nil {
			return view.Wrap(err, c)
		}

		return c.JSON(fiber.Map{
			"message":    "Businesses fetched",
			"businesses": bizs,
		})
	}
}

func MakeBusinessHandler(app *fiber.App, svc business.Service) {
	businessGroup := app.Group("/api/v1/business")
	businessGroup.Post("/new", NewBusiness(svc))
	businessGroup.Get("/listAll", ShowBusinesses(svc))
	businessGroup.Get("/list", ShowBusinessesByCity(svc))
}
