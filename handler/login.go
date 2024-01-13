package handler

import "github.com/gofiber/fiber/v2"

type Users struct {
	Email    string
	Password string
}

var emberUserTest = Users{
	Email:    "user@example.com",
	Password: "password123456",
}

func Login(c *fiber.Ctx) error {
	user := new(Users)
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	return c.Render("login", fiber.Map{
		"Title": "Login",
		"Name":  "Safe",
	})
}
