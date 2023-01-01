package main

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000, https://gofiber.net",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	app.Get("/s", func(c *fiber.Ctx) error {
		c.Cookie(&fiber.Cookie{
			Name:     "SSIDCP",
			Value:    "raj",
			Secure:   true,
			HTTPOnly: true,
			Expires:  time.Now().Add(time.Hour * 72),
			SameSite: "None",
		})
		return c.SendStatus(200)
	})
	app.Get("/r", func(c *fiber.Ctx) error {
		c.Cookie(&fiber.Cookie{
			Name:     "SSIDCP",
			Secure:   true,
			HTTPOnly: true,
			Expires:  time.Now().Add(-(time.Hour * 5)),
			SameSite: "None",
		})
		return c.SendStatus(200)
	})
	app.Get("/c", func(c *fiber.Ctx) error {
		return c.JSON(c.Cookies("SSIDCP"))
	})
	PORT := ":8000"
	if p := os.Getenv("PORT"); p != "" {
		PORT = ":" + PORT
	}

	app.Listen(PORT)
}
