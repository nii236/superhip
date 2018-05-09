package main

import (
	"encoding/json"

	uuid "github.com/satori/go.uuid"
)

// GetListRequest is the JSON request that comes in for the userGetList handler
type GetListRequest struct {
	Pagination struct {
		Page    int `json:"page,omitempty"`
		PerPage int `json:"per_page,omitempty"`
	} `json:"pagination,omitempty"`
	Sort struct {
		Field string `json:"field,omitempty"`
		Order string `json:"order,omitempty"`
	} `json:"sort,omitempty"`
	Filter json.RawMessage `json:"filter,omitempty"`
}

// GetOneRequest is the JSON request that comes in for the userGetOne handler
type GetOneRequest struct {
	ID uuid.UUID `json:"id,omitempty"`
}

// GetManyRequest is the JSON request that comes in for the userGetMany handler
type GetManyRequest struct {
	IDs []uuid.UUID `json:"ids,omitempty"`
}

// GetManyReferenceRequest is the JSON request that comes in for the userGetManyReference handler
type GetManyReferenceRequest struct {
	Target     string    `json:"target,omitempty"`
	ID         uuid.UUID `json:"id,omitempty"`
	Pagination struct {
		Page    int `json:"page,omitempty"`
		PerPage int `json:"per_page,omitempty"`
	} `json:"pagination,omitempty"`
	Sort struct {
		Field string `json:"field,omitempty"`
		Order string `json:"order,omitempty"`
	} `json:"sort,omitempty"`
	Filter json.RawMessage `json:"filter,omitempty"`
	Column string          `json:"column,omitempty`
}

// UpdateRequest is the JSON request that comes in for the userUpdate handler
type UpdateRequest struct {
	ID           uuid.UUID       `json:"id,omitempty"`
	Data         json.RawMessage `json:"data,omitempty"`
	PreviousData json.RawMessage `json:"previous_data,omitempty"`
}

// UpdateManyRequest is the JSON request that comes in for the userUpdateMany handler
type UpdateManyRequest struct {
	IDs  []*uuid.UUID    `json:"ids,omitempty"`
	Data json.RawMessage `json:"data,omitempty"`
}

// CreateRequest is the JSON request that comes in for the userCreate handler
type CreateRequest struct {
	Data json.RawMessage `json:"data,omitempty"`
}

// DeleteRequest is the JSON request that comes in for the userDelete handler
type DeleteRequest struct {
	ID           uuid.UUID       `json:"id,omitempty"`
	PreviousData json.RawMessage `json:"previous_data,omitempty"`
}

// DeleteManyRequest is the JSON request that comes in for the userDeleteMany handler
type DeleteManyRequest struct {
	IDs []uuid.UUID `json:"ids,omitempty"`
}

// Response is the generic response for react admin
type Response struct {
	Total   int             `json:"total"`
	Data    json.RawMessage `json:"data"`
	Message string          `json:"message,omitempty"`
}
