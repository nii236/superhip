package models

import (
	"fmt"
	"regexp"
)

// TeamList is a list of teams
type TeamList []*Team

// ListQuery implements the Collection interface
func (tl *TeamList) ListQuery() string {
	return teamQueries["list"]
}

// ReferenceQuery implements the Collection interface
func (tl *TeamList) ReferenceQuery(table string, column string) string {
	isAlpha := regexp.MustCompile(`^[A-Za-z]+$`).MatchString
	if !isAlpha(table) {
		panic("non alphanumeric table name provided")
	}

	if !isAlpha(column) {
		panic("non alphanumeric column name provided: " + column)
	}

	q := teamQueries["reference"]
	return fmt.Sprintf(q, table, table, column, table)

}

// GetManyQuery implements the Collection interface
func (tl *TeamList) GetManyQuery() string {
	return teamQueries["getmany"]
}

// UpdateManyQuery implements the Collection interface
func (tl *TeamList) UpdateManyQuery() string {
	return teamQueries["updatemany"]
}

// DeleteManyQuery implements the Collection interface
func (tl *TeamList) DeleteManyQuery() string {
	return teamQueries["deletemany"]
}
