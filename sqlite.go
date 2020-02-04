package sqlite

import (
	"database/sql"
	"fmt"
	"reflect"

	// sqlite
	_ "github.com/mattn/go-sqlite3"
)

const (
	ASEND AsendDecent = iota
	DESEND
)

type AsendDecent int8

type DB struct {
	db *sql.DB
}

type Table struct {
	name string
	db   *DB
}

type QuerryStmt interface {
}

type ExecStmt interface {
	Exec(args ...interface{})
}

type Where struct {
}

func New(dataSourceName string) (*DB, error) {
	db, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("Couldn't open db: %v", err)
	}

	return &DB{db: db}, nil
}

func (db *DB) PrepareTable(name string, t reflect.Type) (*Table, error) {
	return nil, nil
}

func (t *Table) PrepareGet() (*sql.Stmt, error) {

	return nil, nil
}
