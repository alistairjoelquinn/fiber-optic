package main

import (
	"log"

	"github.com/alistairjoelquinn/fiber-optic/cmd/books"
	"github.com/alistairjoelquinn/fiber-optic/cmd/database"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: routeErrorHandler,
	})
	openDB()

	app.Get("/", rootRouteResponse)
	app.Get("/v1/books", books.GetAllBooks)
	app.Get("/v1/book/:id", books.GetBook)
	app.Post("/v1/book/", books.AddBook)
	app.Put("/v1/book/:id", books.UpdateBook)
	app.Delete("/v1/book/:id", books.DeleteBook)

	log.Fatal(app.Listen(":3001"))
}

func routeErrorHandler(c *fiber.Ctx, err error) error {
	return c.Status(404).SendString("Error in route")
}

func openDB() *gorm.DB {
	var err error

	dsn := "postgres://postgres:postgres@localhost:5432/books?sslmode=disable"
	database.DB, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		panic("Error connecting to database")
	}

	log.Println("DB connected")
	database.DB.AutoMigrate(&books.Book{})
	log.Println("DB migrated")

	return database.DB
}
