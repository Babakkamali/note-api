package main

import (
	"github.com/gofiber/fiber/v2"
	db "github.com/babakkamali/note-api/config"
	routes "github.com/babakkamali/note-api/routes"
)


func main()	{

	db.Connect()

	app := fiber.New()

	routes.SetupRoutes(app)

	app.Listen(":8000")

}
