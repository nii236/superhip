package models

import (
	"time"

	"github.com/jmoiron/sqlx/types"
	uuid "github.com/satori/go.uuid"
)

// Role is the user model
type Role struct {
	ID         uuid.UUID      `json:"id,omitempty" db:"id"`
	Name       string         `json:"name,omitempty" db:"name"`
	Metadata   types.JSONText `json:"metadata,omitempty" db:"metadata"`
	Archived   bool           `json:"archived,omitempty" db:"archived"`
	ArchivedOn *time.Time     `json:"archived_on,omitempty" db:"archived_on"`
	CreatedAt  time.Time      `json:"created_at,omitempty" db:"created_at"`
}

// CreateParams implements the Cruddable interface
func (t *Role) CreateParams() []interface{} {
	return []interface{}{
		&t.Name,
	}
}

// CreateQuery implements the Cruddable interface
func (t *Role) CreateQuery() string {
	return roleQueries["create"]
}

// ReadParams implements the Cruddable interface
func (t *Role) ReadParams() []interface{} {
	return []interface{}{}
}

// ReadQuery implements the Cruddable interface
func (t *Role) ReadQuery() string {
	return roleQueries["read"]
}

// UpdateManyParams implements the Cruddable interface
func (t *Role) UpdateManyParams() []interface{} {
	return []interface{}{
		&t.Name,
	}
}

// UpdateParams implements the Cruddable interface
func (t *Role) UpdateParams() []interface{} {
	return []interface{}{
		&t.Name,
	}
}

// UpdateQuery implements the Cruddable interface
func (t *Role) UpdateQuery() string {
	return roleQueries["update"]
}

// DeleteParams implements the Cruddable interface
func (t *Role) DeleteParams() []interface{} {
	return []interface{}{}
}

// DeleteQuery implements the Cruddable interface
func (t *Role) DeleteQuery() string {
	return roleQueries["delete"]
}
