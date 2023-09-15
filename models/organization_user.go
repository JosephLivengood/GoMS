package models

import (
	"time"
)

type OrganizationUser struct {
	UserID         int       `json:"-"`
	OrganizationID int       `json:"-"`
	Role           string    `json:"-"`
	CreatedAt      time.Time `json:"-"`
}
