package routes

import (
	"github.com/babakkamali/note-api/controllers"
	"github.com/babakkamali/note-api/repository"
	"github.com/babakkamali/note-api/services"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)


func SetupRoutes(app *fiber.App, dbConnection *gorm.DB) {
    userRepo := repository.NewUserRepository(dbConnection)
    userService := services.NewUserService(userRepo)
    userController := controllers.NewUserController(userService)
    
    noteRepo := repository.NewNoteRepository(dbConnection)
    noteService := services.NewNoteService(noteRepo)
    noteController := controllers.NewNoteController(noteService)

	api := app.Group("/api/v1")

	// User routes
	api.Post("/login", userController.AuthenticateOrRegister)
	api.Post("/verify", userController.ValidateSMSTokenAndLogin)

	// Notes routes
	notes := api.Group("/note")
	notes.Post("/note", noteController.CreateNote)
	notes.Get("/note/:id", noteController.GetNoteByID)
	notes.Get("/notes", noteController.GetAllNotes)
	notes.Put("/note/:id", noteController.UpdateNote)
	notes.Delete("/note/:id", noteController.DeleteNote)
}