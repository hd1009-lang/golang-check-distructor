package routes

import (
	"director/controllers"
	"github.com/gofiber/fiber/v2"
)

func Routes(app fiber.Router) {
	r := app.Group("/distributor")
	r.Post("/api/v1/list", controllers.GetList)
}
