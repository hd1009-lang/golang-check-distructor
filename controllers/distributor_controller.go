package controllers

import (
	"director/models"
	"director/types"
	"fmt"
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

func CreateNews(c *fiber.Ctx) error {
	data := new(types.DistributorNewsForm)
	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": 400, "message": "Data request invalid"})
	}
	fmt.Println(data)
	var supplier_id = c.Locals("SUPPLIERID").(uint64)
	fmt.Println(supplier_id)
	err := models.CreateDistributorNews(data, supplier_id)
	fmt.Println(err)
	return c.JSON(fiber.Map{"status": 200, "message": "Success!"})

}
func DeleteNews(c *fiber.Ctx) error {
	var id = c.Params("id")
	err := models.DeleteNews(id)
	if err != nil {
		fmt.Println(err)
	}
	return c.JSON(fiber.Map{"status": 200, "message": "Success!"})
}
