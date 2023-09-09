package models

import (
	"gorm.io/gorm"
	"time"
)

// Note model corresponds to the "notes" table
type Note struct {
	Id           	 uint      		`gorm:"primaryKey" json:"id"`
	UserId           uint      		`json:"-"`
	NoteTitle        string    		`json:"note_title"`
	NoteDescription  string    		`json:"note_description"`
	CreatedAt        time.Time 		`json:"created_at"`
	UpdatedAt        time.Time 		`json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"-"`

	User 			 User 			`gorm:"foreignKey:UserId;references:Id" json:"-"`
}