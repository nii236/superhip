package models

import (
	"time"

	"github.com/jmoiron/sqlx/types"
	uuid "github.com/satori/go.uuid"
)

// Team is the user model
type Team struct {
	ID         uuid.UUID      `json:"id,omitempty" db:"id"`
	SchoolID   uuid.UUID      `json:"school_id,omitempty" db:"school_id"`
	Name       string         `json:"name,omitempty" db:"name"`
	Metadata   types.JSONText `json:"metadata,omitempty" db:"metadata"`
	Archived   bool           `json:"archived,omitempty" db:"archived"`
	ArchivedOn *time.Time     `json:"archived_on,omitempty" db:"archived_on"`
	CreatedAt  time.Time      `json:"created_at,omitempty" db:"created_at"`
}

// CreateParams implements the Cruddable interface
func (t *Team) CreateParams() []interface{} {
	return []interface{}{
		&t.SchoolID,
		&t.Name,
	}
}

// CreateQuery implements the Cruddable interface
func (t *Team) CreateQuery() string {
	return teamQueries["create"]
}

// ReadParams implements the Cruddable interface
func (t *Team) ReadParams() []interface{} {
	return []interface{}{}
}

// ReadQuery implements the Cruddable interface
func (t *Team) ReadQuery() string {
	return teamQueries["read"]
}

// UpdateManyParams implements the Cruddable interface
func (t *Team) UpdateManyParams() []interface{} {
	return []interface{}{
		&t.SchoolID,
		&t.Name,
	}
}

// UpdateParams implements the Cruddable interface
func (t *Team) UpdateParams() []interface{} {
	return []interface{}{
		&t.SchoolID,
		&t.Name,
	}
}

// UpdateQuery implements the Cruddable interface
func (t *Team) UpdateQuery() string {
	return teamQueries["update"]
}

// DeleteParams implements the Cruddable interface
func (t *Team) DeleteParams() []interface{} {
	return []interface{}{}
}

// DeleteQuery implements the Cruddable interface
func (t *Team) DeleteQuery() string {
	return teamQueries["delete"]
}
