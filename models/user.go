package models

import (
	"time"
)

type User struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Email     string    `gorm:"email" json:"email"`
	Password  string    `gorm:"hash" json:"-"`
	LastLogin time.Time `json:"lastLogin"`
	CreatedAt time.Time `json:"createdAt"`
}

type UserSession struct {
	SessionID int // uuid
	UserID    int
	Expiry    time.Time
}
