package main

import "github.com/gofiber/fiber"

func rootRouteResponse(a *fiber.Ctx) {
	a.Send("home page")
}
