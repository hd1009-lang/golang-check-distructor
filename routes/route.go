package routes

import (
	"director/controllers"
	"director/middleware"
	"github.com/gofiber/fiber/v2"
)

func Routes(app fiber.Router) {
	r := app.Group("/distributor")
	r.Post("/api/v1/list", middleware.AuthLogin, controllers.GetList)
	r.Post("/api/v1/create", middleware.AuthLogin, controllers.CreateNews)
	r.Delete("/api/v1/delete/:id", middleware.AuthLogin, controllers.DeleteNews)
}
