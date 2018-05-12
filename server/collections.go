package main

// Collections are model lists
type Collection interface {
	ListQuery() string
	ReferenceQuery(table string, column string) string
	GetManyQuery() string
	UpdateManyQuery() string
	DeleteManyQuery() string
}

// Collectioner are the methods needed by the DB to handle collections
type Collectioner interface {
	List(collection Collection, opts *ListOptions) error
	Reference(collection Collection, table string, column string, ID string) error
	GetMany(collection Collection, IDs []string) error
	UpdateMany(collection Collection, item Item, IDs []string) error
	DeleteMany(collection Collection, IDs []string) error
	MakeJoin(table string, col1 string, col2 string, fk1 string, fk2 string) error
}
