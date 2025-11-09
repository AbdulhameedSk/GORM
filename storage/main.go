package main

import (
	//For logging messages or errors to console
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Book struct {
	Author    string `json:"author"`
	Title     string `json:"title"`
	Publisher string `json:"publisher"`
}

type Repository struct {
	//DB is a pointer to gorm.DB
	DB *gorm.DB
}

func (r *Repository) CreateBooks(c *fiber.Ctx) error {
	// Implementation for creating a book
	var book Book
	if err := c.BodyParser(&book); err != nil {
		c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
		return err
	}
	err := r.DB.Create(&book).Error
	if err != nil {
		c.Status(500).JSON(fiber.Map{"error": "Could not create book"})
		return err
	}
	return c.Status(201).JSON(book)
}

func (r *Repository) GetBooks(c *fiber.Ctx) error{
	bookModels : =&[]models.Book{}
	err := r.DB.Find(bookModels).Error
	if err != nil {
		c.Status(500).JSON(fiber.Map{"error": "Could not retrieve books"})
		return err
	}
	return c.Status(200).JSON(bookModels)
}

func (r *Repository) GetBookByID(c *fiber.Ctx) error {
	id := c.Params("id")
}

func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.POST("/Create_books", r.CreateBooks)
	api.DELETE("/Delete_books/:id", r.DeleteBooks)
	api.GET("/Get_books/:id", r.GetBookByID)
	api.GET("/books", r.GetBooks)
	api.PUT("/Update_books/:id", r.UpdateBooks)
	// Additional routes can be added here
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	db, err := storage.newConnection(config)
	if err != nil {
		log.Fatal("Could not connect to the database", err)
	}

	r := Repository{
		DB: db,
	}
	app := fiber.New()
	r.SetupRoutes(app)
	log.Fatal(app.Listen(":8080"))

}
