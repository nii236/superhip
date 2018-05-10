package models

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx/types"
	uuid "github.com/satori/go.uuid"
)

// User is the user model
type User struct {
	ID                 uuid.UUID      `json:"id,omitempty" db:"id"`
	SchoolID           uuid.UUID      `json:"school_id,omitempty" db:"school_id"`
	Email              string         `json:"email,omitempty" db:"email"`
	FirstName          string         `json:"first_name,omitempty" db:"first_name"`
	LastName           string         `json:"last_name,omitempty" db:"last_name"`
	PasswordHash       string         `json:"password_hash,omitempty" db:"password_hash"`
	PasswordResetToken string         `json:"password_reset_token,omitempty" db:"password_reset_token"`
	Role               string         `json:"role,omitempty" db:"role"`
	Metadata           types.JSONText `json:"metadata,omitempty" db:"metadata"`
	Archived           bool           `json:"archived,omitempty" db:"archived"`
	ArchivedOn         *time.Time     `json:"archived_on,omitempty" db:"archived_on"`
	CreatedAt          time.Time      `json:"created_at,omitempty" db:"created_at"`
}

// CreateParams implements the Cruddable interface
func (u *User) CreateParams() []interface{} {
	return []interface{}{
		&u.SchoolID,
		&u.FirstName,
		&u.LastName,
		&u.Email,
		&u.PasswordHash,
		&u.Role,
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
		&u.SchoolID,
		&u.FirstName,
		&u.LastName,
		&u.Email,
		&u.PasswordHash,
		&u.Role,
	}
}

// UpdateParams implements the Cruddable interface
func (u *User) UpdateParams() []interface{} {
	fmt.Printf("%+v", u)
	return []interface{}{
		&u.SchoolID,
		&u.FirstName,
		&u.LastName,
		&u.Email,
		&u.PasswordHash,
		&u.Role,
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
