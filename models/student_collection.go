package models

import (
	"fmt"
	"regexp"
)

// StudentList is a list of students
type StudentList []*Student

// ListQuery implements the Collection interface
func (tl *StudentList) ListQuery() string {
	return studentQueries["list"]
}

// ReferenceQuery implements the Collection interface
func (tl *StudentList) ReferenceQuery(table string, column string) string {
	isAlpha := regexp.MustCompile(`^[A-Za-z]+$`).MatchString
	if !isAlpha(table) {
		panic("non alphanumeric table name provided")
	}

	if !isAlpha(column) {
		panic("non alphanumeric column name provided: " + column)
	}

	q := studentQueries["reference"]
	return fmt.Sprintf(q, table, table, column, table)

}

// GetManyQuery implements the Collection interface
func (tl *StudentList) GetManyQuery() string {
	return studentQueries["getmany"]
}

// UpdateManyQuery implements the Collection interface
func (tl *StudentList) UpdateManyQuery() string {
	return studentQueries["updatemany"]
}

// DeleteManyQuery implements the Collection interface
func (tl *StudentList) DeleteManyQuery() string {
	return studentQueries["deletemany"]
}
