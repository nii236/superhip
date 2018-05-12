package models

import (
	"time"

	"github.com/jmoiron/sqlx/types"
	uuid "github.com/satori/go.uuid"
)

// Permission is the permission model
type Permission struct {
	ID         uuid.UUID      `json:"id,omitempty" db:"id"`
	Name       string         `json:"name,omitempty" db:"name"`
	Metadata   types.JSONText `json:"metadata,omitempty" db:"metadata"`
	Archived   bool           `json:"archived,omitempty" db:"archived"`
	ArchivedOn *time.Time     `json:"archived_on,omitempty" db:"archived_on"`
	CreatedAt  time.Time      `json:"created_at,omitempty" db:"created_at"`
}

// CreateParams implements the Cruddable interface
func (t *Permission) CreateParams() []interface{} {
	return []interface{}{
		&t.Name,
	}
}

// CreateQuery implements the Cruddable interface
func (t *Permission) CreateQuery() string {
	return permissionQueries["create"]
}

// ReadParams implements the Cruddable interface
func (t *Permission) ReadParams() []interface{} {
	return []interface{}{}
}

// ReadQuery implements the Cruddable interface
func (t *Permission) ReadQuery() string {
	return permissionQueries["read"]
}

// UpdateManyParams implements the Cruddable interface
func (t *Permission) UpdateManyParams() []interface{} {
	return []interface{}{
		&t.Name,
	}
}

// UpdateParams implements the Cruddable interface
func (t *Permission) UpdateParams() []interface{} {
	return []interface{}{
		&t.Name,
	}
}

// UpdateQuery implements the Cruddable interface
func (t *Permission) UpdateQuery() string {
	return permissionQueries["update"]
}

// DeleteParams implements the Cruddable interface
func (t *Permission) DeleteParams() []interface{} {
	return []interface{}{}
}

// DeleteQuery implements the Cruddable interface
func (t *Permission) DeleteQuery() string {
	return permissionQueries["delete"]
}
