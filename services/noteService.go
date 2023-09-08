package services

import (
	"github.com/babakkamali/note-api/models"
	"github.com/babakkamali/note-api/repository"
)

type NoteService struct {
	repo *repository.NoteRepository
}

func NewNoteService(repo *repository.NoteRepository) *NoteService {
	return &NoteService{repo}
}

func (s *NoteService) CreateNote(note *models.Note) error {
	if err := s.repo.CreateNote(note); err != nil {
		return err
	}
	return nil
}

func (s *NoteService) GetNoteByID(note models.Note) (*models.Note, error) {
	return s.repo.GetNoteByID(note)
}

func (s *NoteService) GetAllNotes(userID uint) ([]models.Note, error) {
	notes, err := s.repo.GetAllNotes(userID)
	if err != nil {
		return nil, err
	}
	return notes, nil
}

func (s *NoteService) UpdateNote(note *models.Note) error {
	return s.repo.UpdateNote(note)
}

func (s *NoteService) DeleteNote(note *models.Note) error {
	return s.repo.DeleteNote(note)
}