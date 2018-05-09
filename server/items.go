package main

// Item are the methods needed for a single item struct
type Item interface {
	CreateParams() []interface{}
	CreateQuery() string

	ReadParams() []interface{}
	ReadQuery() string

	UpdateManyParams() []interface{}
	UpdateParams() []interface{}
	UpdateQuery() string

	DeleteParams() []interface{}
	DeleteQuery() string
}

// Crudder needs to be implemented by the DB, and handles single items
type Crudder interface {
	Create(target Item, source Item) error
	Read(target Item, id string) error
	Update(target Item, source Item, id string) error
	Delete(target Item, id string) error
}
