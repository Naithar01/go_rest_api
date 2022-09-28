package main

import (
	"log"

	"github.com/gofiber/fiber/v2/middleware/cors"

	"github/com/Naithar01/go_rest_api/database"
	"github/com/Naithar01/go_rest_api/router"
)

func main() {
	database.ConnectDB()

	app := router.InitApp()

	app.Use(cors.New())

	log.Fatal(app.Listen(":4000"))
}
