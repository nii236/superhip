package main

import (
	"database/sql"
	"fmt"
	"log"
	"regexp"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
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

type ListOptions struct {
	Limit          int
	Offset         int
	OrderBy        string
	OrderDirection string
}

func isAlpha(val string) bool {
	r := regexp.MustCompile(`^[A-Za-z_]+$`)
	return r.MatchString(val)
}

// Total returns the total number of records for a table
func (db *DB) Total(table string) (int, error) {
	if !isAlpha(table) {
		panic("non alphanumeric string provided: table: " + table)
	}
	var total int
	err := db.conn.Get(&total, fmt.Sprintf(`SELECT COUNT(id) FROM "%s"`, table))
	if err != nil {
		return 0, fmt.Errorf("could not get total number of records: %s", err)
	}

	return total, nil
}

// List implements the Collectioner interface
func (db *DB) List(collection Collection, opts *ListOptions) error {
	limit := &sql.NullInt64{}
	if opts.Limit != 0 {
		limit.Scan(int64(opts.Limit))
	}
	offset := &sql.NullInt64{}
	if opts.Offset != 0 {
		offset.Scan(int64(opts.Offset - 1))
	}

	// orderBy := "created_at"
	// if opts.OrderBy != "id" && opts.OrderBy != "" {
	// 	orderBy = opts.OrderBy
	// }
	// orderDirection := "desc"
	// if opts.OrderDirection != "id" && opts.OrderDirection != "" {
	// 	orderDirection = opts.OrderDirection
	// }
	return db.conn.Select(collection, collection.ListQuery(), limit, offset)
}

// UpdateJoins will drop then make joins
func (db *DB) UpdateJoins(table, col1, col2 string, fk1 uuid.UUID, fk2 []uuid.UUID) error {
	fmt.Println("DB: RUNNING UPDATEJOINS")
	err := db.DropJoins(table, col1, fk1.String())
	if err != nil {
		return fmt.Errorf("could not drop joins: %s", err)
	}

	for _, v := range fk2 {
		err = db.MakeJoin(table, col1, col2, fk1.String(), v.String())
		if err != nil {
			return fmt.Errorf("could not create joins: %s", err)
		}
	}
	return nil
}

// DropJoins drops all joins for a table and fk
func (db *DB) DropJoins(table string, col string, fk string) error {
	fmt.Println("DB: RUNNING DROPJOINS")
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
	fmt.Println("DB: RUNNING MAKEJOIN")
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
	fmt.Println("DB: RUNNING UPDATEMANY")
	params := source.UpdateManyParams()
	params = append(params, pq.Array(IDs))
	return db.conn.Select(collection, collection.UpdateManyQuery(), params...)
}

// DeleteMany implements the Collectioner interface
func (db *DB) DeleteMany(collection Collection, IDs []string) error {
	fmt.Println("DB: RUNNING DELETEMANY")
	return db.conn.Select(collection, collection.DeleteManyQuery(), pq.Array(IDs))
}

// GetMany implements the Collectioner interface
func (db *DB) GetMany(collection Collection, IDs []string) error {
	fmt.Println("DB: RUNNING GETMANY")
	return db.conn.Select(collection, collection.GetManyQuery(), pq.Array(IDs))
}

// Reference implements the Collectioner interface
func (db *DB) Reference(collection Collection, table string, column string, ID string) error {
	fmt.Println("DB: RUNNING REFERENCE")
	return db.conn.Select(collection, collection.ReferenceQuery(table, column), ID)
}

// Create implements the Crudder interface
func (db *DB) Create(target Item, source Item) error {
	fmt.Println("DB: RUNNING CREATE")
	params := source.CreateParams()
	return db.conn.Get(target, source.CreateQuery(), params...)
}

// Read implements the Crudder interface
func (db *DB) Read(item Item, ID string) error {
	fmt.Println("DB: RUNNING READ")
	params := item.ReadParams()
	params = append(params, ID)
	return db.conn.Get(item, item.ReadQuery(), params...)
}

// Update implements the Crudder interface
func (db *DB) Update(target Item, source Item, ID string) error {
	fmt.Println("DB: RUNNING UPDATE")
	params := source.UpdateParams()
	params = append(params, ID)
	return db.conn.Get(target, source.UpdateQuery(), params...)
}

// Delete implements the Crudder interface
func (db *DB) Delete(item Item, ID string) error {
	fmt.Println("DB: RUNNING DELETE")
	params := item.DeleteParams()
	params = append(params, ID)
	return db.conn.Get(item, item.DeleteQuery(), params...)
}
