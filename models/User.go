package models

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type User struct {
	ID uuid.UUID `json:"id"`

	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	DeletedAt time.Time `json:"deleted_at,omitempty"`

	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Code     string `json:"code"`

	Active bool `json:"active"`
}
