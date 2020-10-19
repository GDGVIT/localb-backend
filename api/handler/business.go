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

func MakeBusinessHandler(app *fiber.App, svc business.Service) {
	businessGroup := app.Group("/api/v1/business")
	businessGroup.Post("/new", NewBusiness(svc))
}
