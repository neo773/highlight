package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	H "github.com/highlight/highlight/sdk/highlight-go"
	highlightFiber "github.com/highlight/highlight/sdk/highlight-go/middleware/fiber"
	"log"
)

func main() {
	H.SetProjectID("1jdkoe52")
	H.SetGraphqlClientAddress("https://localhost:8082/public")
	H.SetOTLPEndpoint("http://localhost:4318")
	H.Start()
	defer H.Stop()

	app := fiber.New()
	app.Use(logger.New())
	app.Use(highlightFiber.Middleware())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	log.Fatal(app.Listen(":3456"))
}