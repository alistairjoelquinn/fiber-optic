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
	id := c.Params("id")
	var book Book

	database.DB.Find(&book, id)

	return c.JSON(book)
}

func AddBook(c *fiber.Ctx) error {
	book := Book{
		Title:  "The Brothers Karamazov",
		Author: "Fyodor Dostoyevsky",
		Rating: 5,
	}

	database.DB.Create(&book)

	return c.JSON(book)
}

func UpdateBook(c *fiber.Ctx) error {
	return c.SendString("Update a book")
}

func DeleteBook(c *fiber.Ctx) error {
	return c.SendString("All Books!")
}
