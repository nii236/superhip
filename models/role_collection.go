package models

import (
	"fmt"
	"regexp"
)

// RoleList is a list of roles
type RoleList []*Role

// ListQuery implements the Collection interface
func (tl *RoleList) ListQuery() string {
	return roleQueries["list"]
}

// ReferenceQuery implements the Collection interface
func (tl *RoleList) ReferenceQuery(table string, column string) string {
	isAlpha := regexp.MustCompile(`^[A-Za-z]+$`).MatchString
	if !isAlpha(table) {
		panic("non alphanumeric table name provided")
	}

	if !isAlpha(column) {
		panic("non alphanumeric column name provided: " + column)
	}

	q := roleQueries["reference"]
	return fmt.Sprintf(q, table, table, column, table)

}

// GetManyQuery implements the Collection interface
func (tl *RoleList) GetManyQuery() string {
	return roleQueries["getmany"]
}

// UpdateManyQuery implements the Collection interface
func (tl *RoleList) UpdateManyQuery() string {
	return roleQueries["updatemany"]
}

// DeleteManyQuery implements the Collection interface
func (tl *RoleList) DeleteManyQuery() string {
	return roleQueries["deletemany"]
}
