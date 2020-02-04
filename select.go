package sqlite

import "fmt"

type SelectStmt struct {
	table    *Table
	colunmns interface{}
	err      error
	groupBy  string
	orderBy  string
	ad       AsendDecent
}

func (db *DB) Select(columns interface{}) *SelectStmt {
	return &SelectStmt{colunmns: columns}
}

func (s *SelectStmt) From(t *Table) *SelectStmt {
	if s.err != nil {
		return s
	}

	if s.table != nil {
		s.err = fmt.Errorf("Multiple from isn't allowed. ")
	} else {
		s.table = t
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
