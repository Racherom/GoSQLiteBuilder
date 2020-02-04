package sqlite

import (
	"fmt"
	"log"
	"reflect"
	"strings"
)

type InsertStmt struct {
	table   *Table
	columns []string
	err     error
	values  []interface{}
	query   string
}

func (db *DB) InsertInto(t *Table) *InsertStmt {
	return &InsertStmt{table: t}
}

func (is *InsertStmt) Columns(c interface{}) *InsertStmt {
	if is.err != nil {
		return is
	}

	if is.columns != nil {
		is.err = fmt.Errorf("Couldn't set columns twice. ")
		return is
	}

	switch v := c.(type) {
	case string:
		is.columns = trim(strings.Split(v, ", "))
	case []string:
		is.columns = trim(v)

	default:
		t := reflect.TypeOf(c)

		switch k := t.Kind(); k {
		case reflect.Ptr:
			t = t.Elem()
			k = t.Kind()
			fallthrough
		case reflect.Struct:
			for i := 0; i < t.NumField(); i++ {
				field := t.Field(i)

				tag := field.Tag.Get("db")
				if tag == "-" || tag == "" {
					continue
				}
				tagParts := strings.Split(tag, ",")
				if strings.Contains(tagParts[len(tagParts)-1], "Primary") {
					continue
				}
				is.columns = append(is.columns, tagParts[0])
				log.Printf("%#v\n", field)
			}
		default:
			log.Printf("%#v\n", k)
			return &InsertStmt{err: fmt.Errorf("Invalid columntype %s only sting, []string and struct are alowed. ", t.Name())}
		}

	}

	return is
}

func (i *InsertStmt) Exec(args ...interface{}) error {
	if i.err != nil {
		return i.err
	}

	if args == nil {
		return fmt.Errorf("Couldn't insert empty values. ")
	}

	if len(args) == 1 {

	} else {

	}

	return nil
}

func (i *InsertStmt) Prepare(columns interface{}) {
	if columns != nil {
		i.Columns(columns)
	}

}

func trim(in []string) (out []string) {
	for _, s := range in {
		out = append(out, strings.Trim(s, " \n\t\r"))
	}
	return
}

func (i *InsertStmt) buildQuery() error {

	values := []string{}

	i.query = fmt.Sprintf("INSERT INTO %s %s VALUES %s;", i.table.name, strings.Join(i.columns, ", "), strings.Join(values, ", "))

	return nil
}
