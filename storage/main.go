package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()
	r.SetupRoutes(app)
	log.Fatal(app.Listen(":8080"))

}
