package models

import (
	"time"
)

type User struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Username     string    `gorm:"not null;unique" json:"username" validate:"required"`
	Email        string    `gorm:"not null;unique" json:"email" validate:"required,email"`
	Password     string    `gorm:"not null" json:"password" validate:"required,min=6"`
	Age          int       `gorm:"not null" json:"age" validate:"required,gt=8"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Photos       []Photo
	Comments     []Comment
	SocialMedias []SocialMedia
}
