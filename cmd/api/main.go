package main

import "github.com/gofiber/fiber"

func main() {
	app := fiber.New()

	app.Get("/", rootRouteResponse)

	app.Listen(3001)
}
