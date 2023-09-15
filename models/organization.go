package models

import (
	"time"
)

type Organization struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}
