package books

import (
	"github.com/alistairjoelquinn/fiber-optic/cmd/database"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

func GetAllBooks(c *fiber.Ctx) error {
	var books []Book
	database.DB.Find(&books)

	return c.JSON(books)
}

func GetBook(c *fiber.Ctx) error {
	return c.SendString("single Book")
}

func AddBook(c *fiber.Ctx) error {
	return c.SendString("Add new book")
}

func UpdateBook(c *fiber.Ctx) error {
	return c.SendString("Update a book")
}

func DeleteBook(c *fiber.Ctx) error {
	return c.SendString("All Books!")
}
