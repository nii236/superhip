package models

import "github.com/nleof/goyesql"

var userQueries goyesql.Queries
var schoolQueries goyesql.Queries
var teamQueries goyesql.Queries
var studentQueries goyesql.Queries

func init() {
	userQueries = goyesql.MustParseFile("./queries/users.sql")
	schoolQueries = goyesql.MustParseFile("./queries/schools.sql")
	teamQueries = goyesql.MustParseFile("./queries/teams.sql")
	studentQueries = goyesql.MustParseFile("./queries/students.sql")
}
