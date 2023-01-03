package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/google/uuid"
)

type User struct {
	Id        string `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

func handelUser(c *fiber.Ctx) error {
	user := User{
		FirstName: "juan",
		LastName:  "mange",
	}
	return c.Status(fiber.StatusOK).JSON(user)
}

func handleCreate(c *fiber.Ctx) error {
	user := User{}
	if err := c.BodyParser(&user); err != nil {
		return err
	}
	user.Id = uuid.NewString()
	return c.Status(fiber.StatusOK).JSON(user)
}

func handelCheck(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).SendString("running")
}

func main() {
	app := fiber.New()

	// Middleware

	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello")
	})

	app.Get("/about", func(c *fiber.Ctx) error {
		return c.SendString("about the api")
	})

	app.Use(requestid.New())

	userGroup := app.Group("/api/users")

	userGroup.Post("/create", handleCreate)
	app.Get("/check", handelCheck)
	userGroup.Get("", handelUser)

	app.Listen(":3000")
}
