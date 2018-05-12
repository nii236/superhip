package main

import (
	"fmt"
	"log"
	"regexp"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

var _ Collectioner = &DB{}
var _ Crudder = &DB{}

// DB is the conn
type DB struct {
	conn *sqlx.DB
}

func newDB() (*DB, error) {
	db, err := sqlx.Connect("postgres", "user=dev dbname=superhip password=dev sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	return &DB{conn: db}, nil
}

// List implements the Collectioner interface
func (db *DB) List(collection Collection) error {
	return db.conn.Select(collection, collection.ListQuery())
}

// DropJoins drops all joins for a table and fk
func (db *DB) DropJoins(table string, col string, fk string) error {
	isAlpha := regexp.MustCompile(`^[A-Za-z_]+$`).MatchString
	if !isAlpha(table) {
		panic("non alphanumeric string provided: table: " + table)
	}
	if !isAlpha(col) {
		panic("non alphanumeric string provided: col1: " + col)
	}

	q := fmt.Sprintf(`DELETE FROM "%s" WHERE %s = $1`, table, col)
	_, err := db.conn.Exec(q, fk)
	if err != nil {
		return err
	}
	return nil
}

// MakeJoin creates a join
func (db *DB) MakeJoin(table string, col1 string, col2 string, fk1 string, fk2 string) error {
	isAlpha := regexp.MustCompile(`^[A-Za-z_]+$`).MatchString
	if !isAlpha(table) {
		panic("non alphanumeric string provided: table: " + table)
	}
	if !isAlpha(col1) {
		panic("non alphanumeric string provided: col1: " + col1)
	}
	if !isAlpha(col2) {
		panic("non alphanumeric string provided: col2: " + col2)
	}

	q := fmt.Sprintf(`INSERT INTO "%s" ("%s", "%s") VALUES ($1, $2)`, table, col1, col2)
	_, err := db.conn.Exec(q, fk1, fk2)
	if err != nil {
		return err
	}
	return nil
}

// UpdateMany implements the Collectioner interface
func (db *DB) UpdateMany(collection Collection, source Item, IDs []string) error {
	params := source.UpdateManyParams()
	params = append(params, pq.Array(IDs))
	return db.conn.Select(collection, collection.UpdateManyQuery(), params...)
}

// DeleteMany implements the Collectioner interface
func (db *DB) DeleteMany(collection Collection, IDs []string) error {
	return db.conn.Select(collection, collection.DeleteManyQuery(), pq.Array(IDs))
}

// GetMany implements the Collectioner interface
func (db *DB) GetMany(collection Collection, IDs []string) error {
	return db.conn.Select(collection, collection.GetManyQuery(), pq.Array(IDs))
}

// Reference implements the Collectioner interface
func (db *DB) Reference(collection Collection, table string, column string, ID string) error {
	return db.conn.Select(collection, collection.ReferenceQuery(table, column), ID)
}

// Create implements the Crudder interface
func (db *DB) Create(target Item, source Item) error {
	params := source.CreateParams()
	return db.conn.Get(target, source.CreateQuery(), params...)
}

// Read implements the Crudder interface
func (db *DB) Read(item Item, ID string) error {
	params := item.ReadParams()
	params = append(params, ID)
	return db.conn.Get(item, item.ReadQuery(), params...)
}

// Update implements the Crudder interface
func (db *DB) Update(target Item, source Item, ID string) error {
	params := source.UpdateParams()
	params = append(params, ID)
	return db.conn.Get(target, source.UpdateQuery(), params...)
}

// Delete implements the Crudder interface
func (db *DB) Delete(item Item, ID string) error {
	params := item.DeleteParams()
	params = append(params, ID)
	return db.conn.Get(item, item.DeleteQuery(), params...)
}
