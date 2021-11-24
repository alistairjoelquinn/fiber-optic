package main

import "github.com/gofiber/fiber/v2"

func rootRouteResponse(a *fiber.Ctx) error {
	return a.SendString("home page")
}
