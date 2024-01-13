package main

import (
	"GOhttpServer/handler"
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
)

func checkMiddleware(c *fiber.Ctx) error {
	start := time.Now().UTC()

	fmt.Printf("URL_Request: %s , Method: %s , Time: %s\n",
		c.OriginalURL(), c.Method(), start)
	return c.Next()
}

func main() {

	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	booksData := handler.BooksData{}
	booksData.InitializeBooks()

	// Routes
	app.Post("/login", handler.Login)

	app.Use(checkMiddleware)

	// CRUD routes
	app.Get("/book", booksData.GetBooks)
	app.Get("/book/:id", booksData.GetBook)
	app.Post("/book", booksData.CreateBook)
	app.Put("/book/:id", booksData.UpdateBook)
	app.Delete("/book/:id", booksData.DeleteBook)

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
		secret = "Not Found"
	}
	return c.JSON(fiber.Map{
		"secret": secret,
	})
}
