package models

import (
	"fmt"
	"regexp"
)

// SchoolList is a list of schools
type SchoolList []*School

// ListQuery implements the Collection interface
func (tl *SchoolList) ListQuery() string {
	return schoolQueries["list"]
}

// ReferenceQuery implements the Collection interface
func (tl *SchoolList) ReferenceQuery(table string, column string) string {
	isAlpha := regexp.MustCompile(`^[A-Za-z]+$`).MatchString
	if !isAlpha(table) {
		panic("non alphanumeric table name provided")
	}

	if !isAlpha(column) {
		panic("non alphanumeric column name provided: " + column)
	}

	q := schoolQueries["reference"]
	return fmt.Sprintf(q, table, table, column, table)

}

// GetManyQuery implements the Collection interface
func (tl *SchoolList) GetManyQuery() string {
	return schoolQueries["getmany"]
}

// UpdateManyQuery implements the Collection interface
func (tl *SchoolList) UpdateManyQuery() string {
	return schoolQueries["updatemany"]
}

// DeleteManyQuery implements the Collection interface
func (tl *SchoolList) DeleteManyQuery() string {
	return schoolQueries["deletemany"]
}
