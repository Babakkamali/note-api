package routes

import (
	"github.com/babakkamali/note-api/controllers"
	"github.com/babakkamali/note-api/middleware"
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

	// api v1 group for the whole api 
	api := app.Group("/api/v1")

	// User routes
	api.Post("/login", userController.AuthenticateOrRegister)
	api.Post("/verify", userController.ValidateSMSTokenAndLogin)

	// Notes group
	notes := api.Group("/note")

	// Notes middleware
	notes.Use(middleware.Authentication)

	// Notes routes
	notes.Post("/", noteController.CreateNote)
	notes.Get("/:id", noteController.GetNoteByID)
	notes.Get("/", noteController.GetAllNotes)
	notes.Patch("/:id", noteController.UpdateNote)
	notes.Delete("/:id", noteController.DeleteNote)
}