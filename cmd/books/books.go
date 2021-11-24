package books

import "github.com/gofiber/fiber/v2"

func GetAllBooks(c *fiber.Ctx) error {
	return c.SendString("All Books!")
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
