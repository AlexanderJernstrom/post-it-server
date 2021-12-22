package routes

import "github.com/gofiber/fiber/v2"


type RegisterBody struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`

}

// /register
func Register(c *fiber.Ctx) error {
	newUser := new(RegisterBody)
	if err := c.BodyParser(newUser); err != nil {
		return c.Status(400).SendString("Invalid body")
	}
	return c.Status(200).SendString("OK")
}