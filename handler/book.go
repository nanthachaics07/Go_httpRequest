package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// Book struct to hold book data
type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

// BooksData struct encapsulates book-related data
type BooksData struct {
	books []Book
}

func (bd *BooksData) InitializeBooks() {
	bd.books = append(bd.books, Book{ID: 1, Title: "1984", Author: "George Orwell"})
	bd.books = append(bd.books, Book{ID: 2, Title: "The Great Gatsby", Author: "F. Scott Fitzgerald"})
	bd.books = append(bd.books, Book{ID: 3, Title: "To Kill a Mockingbird", Author: "Harper Lee"})
	bd.books = append(bd.books, Book{ID: 4, Title: "Pride and Prejudice", Author: "Jane Austen"})
	bd.books = append(bd.books, Book{ID: 5, Title: "The Catcher in the Rye", Author: "J.D. Salinger"})
	bd.books = append(bd.books, Book{ID: 6, Title: "Animal Farm", Author: "George Orwell"})
	bd.books = append(bd.books, Book{ID: 7, Title: "Wuthering Heights", Author: "Emily Bronte"})
	bd.books = append(bd.books, Book{ID: 8, Title: "Lord of the Flies", Author: "William Golding"})
	bd.books = append(bd.books, Book{ID: 9, Title: "The Grapes of Wrath", Author: "John Steinbeck"})
	bd.books = append(bd.books, Book{ID: 10, Title: "The Picture of Dorian Gray", Author: "Oscar Wilde"})
}

func (bd *BooksData) GetBooks(c *fiber.Ctx) error {
	return c.JSON(bd.books)
}

func (bd *BooksData) GetBook(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	for _, book := range bd.books {
		if book.ID == id {
			return c.JSON(book)
		}
	}

	return c.SendStatus(fiber.StatusNotFound)
}

func (bd *BooksData) CreateBook(c *fiber.Ctx) error {
	book := new(Book)

	if err := c.BodyParser(book); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	book.ID = len(bd.books) + 1
	bd.books = append(bd.books, *book)

	return c.JSON(book)
}

func (bd *BooksData) UpdateBook(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	bookUpdate := new(Book)
	if err := c.BodyParser(bookUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	for i, book := range bd.books {
		if book.ID == id {
			book.Title = bookUpdate.Title
			book.Author = bookUpdate.Author
			bd.books[i] = book
			return c.JSON(book)
		}
	}

	return c.SendStatus(fiber.StatusNotFound)
}

func (bd *BooksData) DeleteBook(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	for i, book := range bd.books {
		if book.ID == id {
			bd.books = append(bd.books[:i], bd.books[i+1:]...)
			return c.SendStatus(fiber.StatusNoContent)
		}
	}

	return c.SendStatus(fiber.StatusNotFound)
}
