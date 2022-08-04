package middleware

import "github.com/gofiber/fiber/v2"

func AuthLogin(c *fiber.Ctx) error {
	h := c.Get("Authorization")
	if h == "" {
		return fiber.ErrUnauthorized
	}
	user, err := Verify(h)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": 401, "message": "Session expired please login again"})
	}
	c.Locals("SUPPLIERID", user.SupplierId)
	c.Locals("PHONE", user.Phone)
	return c.Next()
}
