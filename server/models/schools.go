package models

import (
	"time"

	"github.com/jmoiron/sqlx/types"
	uuid "github.com/satori/go.uuid"
)

// School is the user model
type School struct {
	ID         uuid.UUID      `json:"id,omitempty" db:"id"`
	Name       string         `json:"name,omitempty" db:"name"`
	Metadata   types.JSONText `json:"metadata,omitempty" db:"metadata"`
	Archived   bool           `json:"archived,omitempty" db:"archived"`
	ArchivedOn *time.Time     `json:"archived_on,omitempty" db:"archived_on"`
	CreatedAt  time.Time      `json:"created_at,omitempty" db:"created_at"`
}

// CreateParams implements the Cruddable interface
func (t *School) CreateParams() []interface{} {
	return []interface{}{
		&t.Name,
	}
}

// CreateQuery implements the Cruddable interface
func (t *School) CreateQuery() string {
	return schoolQueries["create"]
}

// ReadParams implements the Cruddable interface
func (t *School) ReadParams() []interface{} {
	return []interface{}{}
}

// ReadQuery implements the Cruddable interface
func (t *School) ReadQuery() string {
	return schoolQueries["read"]
}

// UpdateManyParams implements the Cruddable interface
func (t *School) UpdateManyParams() []interface{} {
	return []interface{}{
		&t.Name,
	}
}

// UpdateParams implements the Cruddable interface
func (t *School) UpdateParams() []interface{} {
	return []interface{}{
		&t.Name,
	}
}

// UpdateQuery implements the Cruddable interface
func (t *School) UpdateQuery() string {
	return schoolQueries["update"]
}

// DeleteParams implements the Cruddable interface
func (t *School) DeleteParams() []interface{} {
	return []interface{}{}
}

// DeleteQuery implements the Cruddable interface
func (t *School) DeleteQuery() string {
	return schoolQueries["delete"]
}
