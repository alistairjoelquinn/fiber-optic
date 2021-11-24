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
	book := new(Book)
	if err := c.BodyParser(book); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	database.DB.Create(&book)

	return c.JSON(book)
}

func UpdateBook(c *fiber.Ctx) error {
	return c.SendString("Update a book")
}

func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")

	var book Book
	dbResponse := database.DB.First(&book, id)
	if dbResponse.Error != nil {
		return c.Status(500).SendString("No book matches this ID")
	}

	database.DB.Delete(&book)

	return c.Status(200).SendString("book deleted")
}
