package models

import (
	"strings"

	"github.com/gobuffalo/packr"
	"github.com/lib/pq"
	"github.com/nleof/goyesql"
	uuid "github.com/satori/go.uuid"
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

// UUIDArray is a list of UUIDs
type UUIDArray []uuid.UUID

// Scan scans the result into an array of UUIDs
func (n *UUIDArray) Scan(value interface{}) error {
	array := string(value.([]byte))
	if array == "{NULL}" {
		*n = UUIDArray{}
		return nil
	}
	array = strings.Replace(array, "{", "", -1)
	array = strings.Replace(array, "}", "", -1)
	ids := strings.Split(array, ",")
	for _, v := range ids {
		id := uuid.FromStringOrNil(v)
		if id == uuid.Nil {
			continue
		}
		current := *n
		current = append(current, id)
		*n = current
	}
	return nil

	return pq.Array(&n).Scan(value)
}
