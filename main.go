package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
)

// Book struct to hold book data
type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Book // Slice to store books

func main() {

	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Initialize in-memory data
	books = append(books, Book{ID: 1, Title: "1984", Author: "George Orwell"})
	books = append(books, Book{ID: 2, Title: "The Great Gatsby", Author: "F. Scott Fitzgerald"})

	// CRUD routes
	app.Get("/book", getBooks)
	app.Get("/book/:id", getBook)
	app.Post("/book", createBook)
	app.Put("/book/:id", updateBook)
	app.Delete("/book/:id", deleteBook)

	app.Post("/upload", uploadFile)
	app.Get("/viewHTML", testHTML)
	app.Get("/config", getENV)

	app.Listen(":8080")
}

// Handlers
func uploadFile(c *fiber.Ctx) error {
	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	err = c.SaveFile(file, "./upload/"+file.Filename)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).SendString("File Uploaded")
}

func testHTML(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title": "Hello, World!",
		"Name":  "Safe",
	})
}

func getENV(c *fiber.Ctx) error {
	secret := os.Getenv("SECRET")
	if secret != "" {
		return c.JSON(fiber.Map{
			"secret": secret,
		})
	}
	return c.JSON(fiber.Map{
		"secret": "Not found",
	})
}
