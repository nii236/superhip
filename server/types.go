package main

import uuid "github.com/satori/go.uuid"

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
	Filter map[string]interface{} `json:"filter,omitempty"`
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
	Filter map[string]interface{} `json:"filter,omitempty"`
}

// UpdateRequest is the JSON request that comes in for the userUpdate handler
type UpdateRequest struct {
	ID           uuid.UUID         `json:"id,omitempty"`
	Data         map[string]string `json:"data,omitempty"`
	PreviousData map[string]string `json:"previous_data,omitempty"`
}

// UpdateManyRequest is the JSON request that comes in for the userUpdateMany handler
type UpdateManyRequest struct {
	IDs  []uuid.UUID
	Data map[string]string
}

// CreateRequest is the JSON request that comes in for the userCreate handler
type CreateRequest struct {
	Data map[string]string `json:"data,omitempty"`
}

// DeleteRequest is the JSON request that comes in for the userDelete handler
type DeleteRequest struct {
	ID           uuid.UUID              `json:"id,omitempty"`
	PreviousData map[string]interface{} `json:"previous_data,omitempty"`
}

// DeleteManyRequest is the JSON request that comes in for the userDeleteMany handler
type DeleteManyRequest struct {
	IDs []uuid.UUID `json:"ids,omitempty"`
}
