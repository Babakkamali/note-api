package routes

import "github.com/gofiber/fiber/v2"
import "github.com/babakkamali/note-api/controllers"

func SetupRoutes(app *fiber.App){

	// User routes
	app.Post("/login", controllers.LoginUser)

	// Notes routes
	app.Post("/note", controllers.CreateNote)
	app.Get("/note/:id", controllers.GetNoteByID)
	app.Get("/notes", controllers.GetAllNotes)
	app.Put("/note/:id", controllers.UpdateNote)
	app.Delete("/note/:id", controllers.DeleteNote)
}