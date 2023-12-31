package main

import (
	"log"

	db "github.com/babakkamali/note-api/config"
	routes "github.com/babakkamali/note-api/routes"
	"github.com/babakkamali/note-api/utils"

	"github.com/gofiber/fiber/v2"
)


func main() {
    dbConnection, err := db.Connect()
    if err != nil {
        log.Fatalf("Failed to connect to the database: %v", err)
    }

    app := fiber.New()
    utils.HealthCheck(app)

    routes.SetupRoutes(app, dbConnection)

    app.Listen(":80")
}

