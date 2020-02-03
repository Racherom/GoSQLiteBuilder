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

type SelectStmt struct {
	from     *Table
	colunmns interface{}
	err      error
	groupBy  string
	orderBy  string
	ad       AsendDecent
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

func (db *DB) Select(columns interface{}) *SelectStmt {
	return &SelectStmt{colunmns: columns}
}

func (db *DB) PrepareTable(name string, t reflect.Type) (*Table, error) {
	return nil, nil
}

func (t *Table) PrepareGet() (*sql.Stmt, error) {

	return nil, nil
}

func (s *SelectStmt) From(t *Table) *SelectStmt {
	if s.err != nil {
		return s
	}

	if s.from != nil {
		s.err = fmt.Errorf("Multiple from isn't allowed. ")
	} else {
		s.from = t
	}
	return s
}

func (s *SelectStmt) GroupBy(group string) *SelectStmt {
	s.groupBy = group
	return s
}

func (s *SelectStmt) OrderBy(group string) *SelectStmt {
	s.orderBy = group
	return s
}

func (s *SelectStmt) Error() error {
	return s.err
}
