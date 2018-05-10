package models

import (
	"time"

	"github.com/jmoiron/sqlx/types"
	uuid "github.com/satori/go.uuid"
)

// Student is the user model
type Student struct {
	ID         uuid.UUID      `json:"id,omitempty" db:"id"`
	TeamID     uuid.UUID      `json:"team_id,omitempty" db:"team_id"`
	Name       string         `json:"name,omitempty" db:"name"`
	Metadata   types.JSONText `json:"metadata,omitempty" db:"metadata"`
	Archived   bool           `json:"archived,omitempty" db:"archived"`
	ArchivedOn *time.Time     `json:"archived_on,omitempty" db:"archived_on"`
	CreatedAt  time.Time      `json:"created_at,omitempty" db:"created_at"`
}

// CreateParams implements the Cruddable interface
func (t *Student) CreateParams() []interface{} {
	return []interface{}{
		&t.TeamID,
		&t.Name,
	}
}

// CreateQuery implements the Cruddable interface
func (t *Student) CreateQuery() string {
	return studentQueries["create"]
}

// ReadParams implements the Cruddable interface
func (t *Student) ReadParams() []interface{} {
	return []interface{}{}
}

// ReadQuery implements the Cruddable interface
func (t *Student) ReadQuery() string {
	return studentQueries["read"]
}

// UpdateManyParams implements the Cruddable interface
func (t *Student) UpdateManyParams() []interface{} {
	return []interface{}{
		&t.TeamID,
		&t.Name,
	}
}

// UpdateParams implements the Cruddable interface
func (t *Student) UpdateParams() []interface{} {
	return []interface{}{
		&t.TeamID,
		&t.Name,
	}
}

// UpdateQuery implements the Cruddable interface
func (t *Student) UpdateQuery() string {
	return studentQueries["update"]
}

// DeleteParams implements the Cruddable interface
func (t *Student) DeleteParams() []interface{} {
	return []interface{}{}
}

// DeleteQuery implements the Cruddable interface
func (t *Student) DeleteQuery() string {
	return studentQueries["delete"]
}
