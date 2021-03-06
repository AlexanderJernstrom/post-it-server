package routes

import (
	"server/models"
	service "server/services"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)



type RegisterBody struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type LoginBody struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

// /register
func Register(c *fiber.Ctx) error {
	requestBody := new(RegisterBody)
	if err := c.BodyParser(requestBody); err != nil {
		return c.Status(400).SendString("Invalid body")
	}
	
	newUser := new(models.User)
	newUser.Name = requestBody.Name
	newUser.Email = requestBody.Email
	newUser.Password = requestBody.Password
	
	userWithEmailExists, err := service.UserWithEmailExists(newUser.Email)

	if err != nil {
		return c.Status(500).SendString("Internal server error")
	}

	if userWithEmailExists {
		return c.Status(500).SendString("User with email already exists")
	}

	user, err := service.CreateUser(*newUser)

	if err != nil {
		return c.Status(500).SendString("Could not create user")
	}

	return c.Status(200).JSON(user)
}



// /login
func Login(c *fiber.Ctx) error {
	requestBody := new(LoginBody)
	if err := c.BodyParser(requestBody); err != nil {
		return c.Status(400).SendString("Invalid body")
	}

	if _, err := service.UserWithEmailExists(requestBody.Email); err != nil {
		return c.Status(500).SendString("User with that email does not exist")
	}

	user, err := service.GetUserByEmail(requestBody.Email)

	if err != nil {
		return c.Status(500).SendString("Could not get user")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(requestBody.Password)); err != nil {
		return c.Status(500).SendString("Invalid password")
	}

	token, err := user.CreateToken()

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}


	return c.SendString(token)
}