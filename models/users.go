package models

import (
	"time"

	"github.com/jmoiron/sqlx/types"
	uuid "github.com/satori/go.uuid"
)

// User is the user model
type User struct {
	ID                 uuid.UUID      `json:"id" db:"id"`
	SchoolIDs          UUIDArray      `json:"school_ids" db:"school_ids"`
	RoleIDs            UUIDArray      `json:"role_ids" db:"role_ids"`
	Email              string         `json:"email" db:"email"`
	FirstName          string         `json:"first_name" db:"first_name"`
	LastName           string         `json:"last_name" db:"last_name"`
	Password           string         `json:"password"`
	PasswordHash       string         `json:"password_hash" db:"password_hash"`
	PasswordResetToken string         `json:"password_reset_token" db:"password_reset_token"`
	Metadata           types.JSONText `json:"metadata" db:"metadata"`
	Archived           bool           `json:"archived" db:"archived"`
	ArchivedOn         *time.Time     `json:"archived_on" db:"archived_on"`
	CreatedAt          time.Time      `json:"created_at" db:"created_at"`
}

// CreateParams implements the Cruddable interface
func (u *User) CreateParams() []interface{} {
	return []interface{}{
		&u.FirstName,
		&u.LastName,
		&u.Email,
		&u.PasswordHash,
	}
}

// CreateQuery implements the Cruddable interface
func (u *User) CreateQuery() string {
	return userQueries["create"]
}

// ReadParams implements the Cruddable interface
func (u *User) ReadParams() []interface{} {
	return []interface{}{}
}

// ReadQuery implements the Cruddable interface
func (u *User) ReadQuery() string {
	return userQueries["read"]
}

// UpdateManyParams implements the Cruddable interface
func (u *User) UpdateManyParams() []interface{} {
	return []interface{}{
		&u.FirstName,
		&u.LastName,
		&u.Email,
		&u.PasswordHash,
	}
}

// UpdateParams implements the Cruddable interface
func (u *User) UpdateParams() []interface{} {
	return []interface{}{
		&u.FirstName,
		&u.LastName,
		&u.Email,
		&u.PasswordHash,
	}
}

// UpdateQuery implements the Cruddable interface
func (u *User) UpdateQuery() string {
	return userQueries["update"]
}

// DeleteParams implements the Cruddable interface
func (u *User) DeleteParams() []interface{} {
	return []interface{}{}
}

// DeleteQuery implements the Cruddable interface
func (u *User) DeleteQuery() string {
	return userQueries["delete"]
}
