package models

import (
	"time"
)

// User model corresponds to the "users" table
type User struct {
	Id         		uint      `gorm:"primaryKey" json:"id"`
	PhoneNumber    	string    `gorm:"type:VARCHAR(11);unique;not null" json:"phone_number"`
	FirstName      	string    `gorm:"type:VARCHAR(75)" json:"first_name"`
	LastName       	string    `gorm:"type:VARCHAR(75)" json:"last_name"`
	DateRegistered 	time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"date_registered"`
}