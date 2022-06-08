package models

import (
	"time"
)

type Ping struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	CreatedBy string    `json:"createdBy"`
	CreatedAt time.Time `json:"createdAt"`
}
