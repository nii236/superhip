package models

import (
	"fmt"
	"regexp"
)

// UserList is a list of users
type UserList []*User

// ListQuery implements the Collection interface
func (u *UserList) ListQuery() string {
	return userQueries["list"]
}

// ReferenceQuery implements the Collection interface
func (u *UserList) ReferenceQuery(table string, column string) string {
	isAlpha := regexp.MustCompile(`^[A-Za-z]+$`).MatchString
	if !isAlpha(table) {
		panic("non alphanumeric table name provided")
	}

	if !isAlpha(column) {
		panic("non alphanumeric column name provided: " + column)
	}

	q := userQueries["reference"]
	return fmt.Sprintf(q, table, table, column, table)

}

// GetManyQuery implements the Collection interface
func (u *UserList) GetManyQuery() string {
	return userQueries["getmany"]
}

// UpdateManyQuery implements the Collection interface
func (u *UserList) UpdateManyQuery() string {
	return userQueries["updatemany"]
}

// DeleteManyQuery implements the Collection interface
func (u *UserList) DeleteManyQuery() string {
	return userQueries["deletemany"]
}
