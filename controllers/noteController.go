package controllers

import (
	"strconv"

	"github.com/babakkamali/note-api/models"
	"github.com/babakkamali/note-api/services"
	"github.com/babakkamali/note-api/utils"
	"github.com/gofiber/fiber/v2"
)

type NoteController struct {
	noteService *services.NoteService
}

func NewNoteController(service *services.NoteService) *NoteController {
	return &NoteController{
		noteService: service,
	}
}

// CreateNote handles POST requests to create a new note.
func (nc *NoteController) CreateNote(c *fiber.Ctx) error {
	// Logic for creating a new note...

	return c.JSON(fiber.Map{"status": "success", "message": "Note created"})
}

// GetNoteByID handles GET requests to retrieve a note by its ID.
func (nc *NoteController) GetNoteByID(c *fiber.Ctx) error {
	noteIDStr := c.Params("id")
	noteID, err := strconv.Atoi(noteIDStr)
	if err != nil {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, "Invalid note ID")
	}

	note, err := nc.noteService.GetNoteByID(uint(noteID))
	if err != nil {
		return utils.SendErrorResponse(c, fiber.StatusNotFound, "Note not found")
	}

	return c.JSON(note)
}

// GetAllNotes handles GET requests to retrieve all notes of a user.
func (nc *NoteController) GetAllNotes(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint) // Assuming you've stored the user ID in the context using some middleware
	notes, err := nc.noteService.GetAllNotes(userID)
	if err != nil {
		return utils.SendErrorResponse(c, fiber.StatusInternalServerError, "Failed to retrieve notes")
	}

	return c.JSON(fiber.Map{"status": "success", "data": notes})
}

// UpdateNote handles PUT requests to update a note by ID.
func (nc *NoteController) UpdateNote(c *fiber.Ctx) error {
	noteIDStr := c.Params("id")
	noteID, err := strconv.Atoi(noteIDStr)
	if err != nil {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, "Invalid note ID")
	}

	// Parsing the body to get the updated fields
	var updatedNote models.Note
	if err := c.BodyParser(&updatedNote); err != nil {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, "Failed to parse request body")
	}

	updatedNote.Id = uint(noteID)
	err = nc.noteService.UpdateNote(&updatedNote)
	if err != nil {
		return utils.SendErrorResponse(c, fiber.StatusInternalServerError, "Failed to update note")
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Note updated successfully"})
}

// DeleteNote handles DELETE requests to delete a note by ID.
func (nc *NoteController) DeleteNote(c *fiber.Ctx) error {
	noteIDStr := c.Params("id")
	noteID, err := strconv.Atoi(noteIDStr)
	if err != nil {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, "Invalid note ID")
	}

	err = nc.noteService.DeleteNote(uint(noteID))
	if err != nil {
		return utils.SendErrorResponse(c, fiber.StatusInternalServerError, "Failed to delete note")
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Note deleted successfully"})
}
