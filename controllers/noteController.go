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
    // Parse the note from the request body
    var newNote models.Note
    if err := c.BodyParser(&newNote); err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, "Failed to parse request body")
    }

    // Fetch the user ID from the context (provided by the Authentication middleware)
    userID, ok := c.Locals("userID").(uint)
    if !ok {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, "Invalid or missing user ID")
    }

    newNote.UserId = userID

    // Use the service to create a new note
    if err := nc.noteService.CreateNote(&newNote); err != nil {
        return utils.SendErrorResponse(c, fiber.StatusInternalServerError, "Failed to create note")
    }

    return utils.SendResponse(c, fiber.StatusCreated, "Note created", newNote)
}

// GetNoteByID handles GET requests to retrieve a note by its ID.
func (nc *NoteController) GetNoteByID(c *fiber.Ctx) error {
    var note models.Note
    noteIDStr := c.Params("id")
    noteID, err := strconv.Atoi(noteIDStr)
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, "Invalid note ID")
    }

    userID, ok := c.Locals("userID").(uint)
    if !ok {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, "Invalid or missing user ID")
    }

    note.Id = uint(noteID)
    note.UserId = userID

    retrievedNote, err := nc.noteService.GetNoteByID(note)
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusNotFound, "Note not found")
    }

    return utils.SendResponse(c, fiber.StatusOK, "Note fetched successfully", retrievedNote)
}

// GetAllNotes handles GET requests to retrieve all notes of a user.
func (nc *NoteController) GetAllNotes(c *fiber.Ctx) error {
    userID, ok := c.Locals("userID").(uint) // Two-variable type assertion
    if !ok {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, "Invalid or missing user ID")
    }

    notes, err := nc.noteService.GetAllNotes(userID)
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusInternalServerError, "Failed to retrieve notes")
    }

    return utils.SendResponse(c, fiber.StatusOK, "Notes fetched successfully", notes)
}

// UpdateNote handles PUT requests to update a note by ID.
func (nc *NoteController) UpdateNote(c *fiber.Ctx) error {
    noteIDStr := c.Params("id")
    noteID, err := strconv.Atoi(noteIDStr)
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, "Invalid note ID")
    }

    userID, ok := c.Locals("userID").(uint)
    if !ok {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, "Invalid or missing user ID")
    }

    // Parsing the body to get the updated fields
    var updatedNote models.Note
    if err := c.BodyParser(&updatedNote); err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, "Failed to parse request body")
    }

    updatedNote.Id = uint(noteID)
    updatedNote.UserId = userID // Setting the UserId from the authenticated session, not the request body
    err = nc.noteService.UpdateNote(&updatedNote)
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusInternalServerError, "Failed to update note")
    }

    return utils.SendResponse(c, fiber.StatusOK, "Note updated successfully", nil)
}


// DeleteNote handles DELETE requests to delete a note by ID.
func (nc *NoteController) DeleteNote(c *fiber.Ctx) error {
    noteIDStr := c.Params("id")
    noteID, err := strconv.Atoi(noteIDStr)
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, "Invalid note ID")
    }

    userID, ok := c.Locals("userID").(uint)
    if !ok {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, "Invalid or missing user ID")
    }

    noteToDelete := models.Note{
        Id: uint(noteID),
        UserId: userID, // Setting the UserId from the authenticated session
    }
    err = nc.noteService.DeleteNote(&noteToDelete)
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusInternalServerError, "Failed to delete note")
    }

    return utils.SendResponse(c, fiber.StatusOK, "Note deleted successfully", nil)
}
