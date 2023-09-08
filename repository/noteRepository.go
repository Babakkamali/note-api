package repository

import (
    "github.com/babakkamali/note-api/models"
    "gorm.io/gorm"
)

type NoteRepository struct {
    db *gorm.DB
}

func NewNoteRepository(db *gorm.DB) *NoteRepository {
    return &NoteRepository{db}
}

func (repo *NoteRepository) CreateNote(note *models.Note) error {
    return repo.db.Create(note).Error
}

func (repo *NoteRepository) GetNoteByID(noteID uint) (*models.Note, error) {
    var note models.Note
    err := repo.db.First(&note, noteID).Error
    if err != nil {
        return nil, err
    }
    return &note, nil
}

func (repo *NoteRepository) GetAllNotes(userID uint) ([]models.Note, error) {
    var notes []models.Note
    err := repo.db.Where("user_id = ?", userID).Find(&notes).Error
    if err != nil {
        return nil, err
    }
    return notes, nil
}

func (repo *NoteRepository) UpdateNote(note *models.Note) error {
    return repo.db.Save(note).Error
}

func (repo *NoteRepository) DeleteNote(noteID uint) error {
    return repo.db.Delete(&models.Note{}, noteID).Error
}