package main

import (
	"log"

	"github.com/alistairjoelquinn/fiber-optic/cmd/books"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: routeErrorHandler,
	})

	app.Get("/", rootRouteResponse)
	app.Get("/v1/books", books.GetAllBooks)
	app.Get("/v1/book/:id", books.GetBook)
	app.Post("/v1/book/", books.AddBook)
	app.Post("/v1/book/update/:id", books.UpdateBook)
	app.Post("/v1/book/delete/:id", books.DeleteBook)

	log.Fatal(app.Listen(":3001"))
}

func routeErrorHandler(c *fiber.Ctx, err error) error {
	return c.Status(404).SendString("Error in route")
}
