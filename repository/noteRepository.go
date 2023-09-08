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

func (repo *NoteRepository) GetNoteByID(note models.Note) (*models.Note, error) {
    err := repo.db.Where("id = ? AND user_id = ?", note.Id, note.UserId).First(&note).Error
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
    return repo.db.Model(&models.Note{}).Where("id = ? AND user_id = ?", note.Id, note.UserId).Updates(note).Error
}

func (repo *NoteRepository) DeleteNote(note *models.Note) error {
    return repo.db.Where("id = ? AND user_id = ?", note.Id, note.UserId).Delete(&models.Note{}).Error
}