package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rithikjain/local-businesses-backend/api/view"
	"github.com/rithikjain/local-businesses-backend/pkg/business"
	"github.com/rithikjain/local-businesses-backend/pkg/models"
	"strconv"
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
		page, _ := strconv.Atoi(c.Query("page", "1"))
		pageSize, _ := strconv.Atoi(c.Query("pagesize", "10"))

		bizs, err := svc.GetApprovedBusinesses(page, pageSize)
		if err != nil {
			return view.Wrap(err, c)
		}

		return c.JSON(fiber.Map{
			"message":    "Businesses fetched",
			"page":       page,
			"page_size":  pageSize,
			"businesses": bizs,
		})
	}
}

func ShowBusinessesByCity(svc business.Service) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		page, _ := strconv.Atoi(c.Query("page", "1"))
		pageSize, _ := strconv.Atoi(c.Query("pagesize", "10"))

		city := c.Query("city")
		bizs, err := svc.GetBusinessesByCity(city, page, pageSize)
		if err != nil {
			return view.Wrap(err, c)
		}

		return c.JSON(fiber.Map{
			"message":    "Businesses fetched",
			"page":       page,
			"page_size":  pageSize,
			"businesses": bizs,
		})
	}
}

func ShowBusinessesByCityAndType(svc business.Service) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		page, _ := strconv.Atoi(c.Query("page", "1"))
		pageSize, _ := strconv.Atoi(c.Query("pagesize", "10"))
		typ := c.Query("type", "Small")

		city := c.Query("city")
		bizs, err := svc.GetBusinessesByCityAndType(city, typ, page, pageSize)
		if err != nil {
			return view.Wrap(err, c)
		}

		return c.JSON(fiber.Map{
			"message":    "Businesses fetched",
			"page":       page,
			"page_size":  pageSize,
			"businesses": bizs,
		})
	}
}

func MakeBusinessHandler(app *fiber.App, svc business.Service) {
	businessGroup := app.Group("/api/v1/business")
	businessGroup.Post("/new", NewBusiness(svc))
	businessGroup.Get("/listAll", ShowBusinesses(svc))
	businessGroup.Get("/list", ShowBusinessesByCity(svc))
	businessGroup.Get("/listByType", ShowBusinessesByCityAndType(svc))
}
