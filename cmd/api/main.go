package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: routeErrorHandler,
	})

	app.Get("/", rootRouteResponse)

	app.Listen(":3001")
}

func routeErrorHandler(c *fiber.Ctx, err error) error {
	return c.Status(404).SendString("Error in route")
}
