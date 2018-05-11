package models

import (
	"strings"
	"time"

	"github.com/lib/pq"

	"github.com/jmoiron/sqlx/types"
	uuid "github.com/satori/go.uuid"
)

// User is the user model
type User struct {
	ID                 uuid.UUID      `json:"id" db:"id"`
	SchoolIDs          UUIDArray      `json:"school_ids" db:"school_ids"`
	Email              string         `json:"email" db:"email"`
	FirstName          string         `json:"first_name" db:"first_name"`
	LastName           string         `json:"last_name" db:"last_name"`
	PasswordHash       string         `json:"password_hash" db:"password_hash"`
	PasswordResetToken string         `json:"password_reset_token" db:"password_reset_token"`
	Role               string         `json:"role" db:"role"`
	Metadata           types.JSONText `json:"metadata" db:"metadata"`
	Archived           bool           `json:"archived" db:"archived"`
	ArchivedOn         *time.Time     `json:"archived_on" db:"archived_on"`
	CreatedAt          time.Time      `json:"created_at" db:"created_at"`
}

// UUIDArray is a list of UUIDs
type UUIDArray []uuid.UUID

// Scan scans the result into an array of UUIDs
func (n *UUIDArray) Scan(value interface{}) error {
	array := string(value.([]byte))
	if array == "{NULL}" {
		*n = UUIDArray{}
		return nil
	}
	array = strings.Replace(array, "{", "", -1)
	array = strings.Replace(array, "}", "", -1)
	ids := strings.Split(array, ",")
	for _, v := range ids {
		id := uuid.FromStringOrNil(v)
		if id == uuid.Nil {
			continue
		}
		current := *n
		current = append(current, id)
		*n = current
	}
	return nil

	return pq.Array(&n).Scan(value)
}

// CreateParams implements the Cruddable interface
func (u *User) CreateParams() []interface{} {
	return []interface{}{
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
		&u.FirstName,
		&u.LastName,
		&u.Email,
		&u.PasswordHash,
		&u.Role,
	}
}

// UpdateParams implements the Cruddable interface
func (u *User) UpdateParams() []interface{} {
	return []interface{}{
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
