package models

import (
	"github.com/gobuffalo/packr"
	"github.com/nleof/goyesql"
)

var roleQueries goyesql.Queries
var permissionQueries goyesql.Queries
var userQueries goyesql.Queries
var schoolQueries goyesql.Queries
var teamQueries goyesql.Queries
var studentQueries goyesql.Queries

func init() {
	b := packr.NewBox("./queries/")

	roleQueries = goyesql.MustParseBytes(b.Bytes("roles.sql"))
	permissionQueries = goyesql.MustParseBytes(b.Bytes("permissions.sql"))
	userQueries = goyesql.MustParseBytes(b.Bytes("users.sql"))
	schoolQueries = goyesql.MustParseBytes(b.Bytes("schools.sql"))
	teamQueries = goyesql.MustParseBytes(b.Bytes("teams.sql"))
	studentQueries = goyesql.MustParseBytes(b.Bytes("students.sql"))
}
