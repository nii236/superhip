package models

import (
	"fmt"
	"regexp"
)

// PermissionList is a list of permissions
type PermissionList []*Permission

// ListQuery implements the Collection interface
func (tl *PermissionList) ListQuery() string {
	return permissionQueries["list"]
}

// ReferenceQuery implements the Collection interface
func (tl *PermissionList) ReferenceQuery(table string, column string) string {
	isAlpha := regexp.MustCompile(`^[A-Za-z]+$`).MatchString
	if !isAlpha(table) {
		panic("non alphanumeric table name provided")
	}

	if !isAlpha(column) {
		panic("non alphanumeric column name provided: " + column)
	}

	q := permissionQueries["reference"]
	return fmt.Sprintf(q, table, table, column, table)

}

// GetManyQuery implements the Collection interface
func (tl *PermissionList) GetManyQuery() string {
	return permissionQueries["getmany"]
}

// UpdateManyQuery implements the Collection interface
func (tl *PermissionList) UpdateManyQuery() string {
	return permissionQueries["updatemany"]
}

// DeleteManyQuery implements the Collection interface
func (tl *PermissionList) DeleteManyQuery() string {
	return permissionQueries["deletemany"]
}
