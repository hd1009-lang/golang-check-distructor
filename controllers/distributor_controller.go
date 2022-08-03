package controllers

import (
	"director/models"
	"director/types"
	"github.com/gofiber/fiber/v2"
)

func GetList(c *fiber.Ctx) error {
	data := new(types.DistributorNewsList)
	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": 404, "message": "Error on set permission request", "data": err})
	}
	result := models.GetDistributorNews(data.Page, data.PageSize)
	return c.JSON(fiber.Map{"status": 200, "message": "Success", "data": result, "currentPage": data.Page})
}
