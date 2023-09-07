package models

import (
	"time"
)

// AuthToken model corresponds to the "auth_tokens" table
type AuthToken struct {
	Id     			uint      	`gorm:"primaryKey" json:"id"`
	UserId      	uint      	`gorm:"not null" json:"user_id"`
	Token       	string    	`gorm:"type:VARCHAR(7);not null" json:"token"`
	ExpiryDate  	*time.Time 	`json:"expiry_date"` // Pointer to allow for null values
	DateCreated 	time.Time 	`gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"date_created"`

	User 			User 		`gorm:"foreignKey:UserId;references:Id" json:"user"`
}