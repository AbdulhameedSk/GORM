package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"golang.org/x/mod/sumdb/storage"
	"gorm.io/gorm"
)

type Book struct{
	Author string `json:"author"`
	Title  string `json:"title"`
	Publisher string `json:"publisher"`
}

type Repository struct {
	//DB is a pointer to gorm.DB
	DB *gorm.DB
}

func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.POST("/Create_books", r.CreateBooks)
	api.DELETE("/Delete_books/:id", r.DeleteBooks)
	api.GET("/Get_books/:id", r.GetBooks)
	api.GET("/books", r.GetBooks)
	api.PUT("/Update_books/:id", r.UpdateBooks)
	// Additional routes can be added here
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	
	db,err:=storage.newConnection(config)
	if err!=nil{
		log.Fatal("Could not connect to the database",err)
	}

	r := Repository{
		DB: db,
	}
	app := fiber.New()
	r.SetupRoutes(app)
	log.Fatal(app.Listen(":8080"))

}
