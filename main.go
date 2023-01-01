package main

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":3000"
	} else {
		port = ":" + port
	}

	return port
}
func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "https://exuberant-cowboy-hat-yak.cyclic.app",
		AllowCredentials: true,
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

	app.Listen(getPort())
}
