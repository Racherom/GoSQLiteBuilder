package sqlite

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"strings"
)

type InsertStmt struct {
	table   *Table
	columns []string
	err     error
	query   string
}

type PreparedInsertStmt struct {
	columns []string
	err     error
	stmt    sql.Stmt
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
			return &InsertStmt{err: fmt.Errorf("Invalid columntype %s only sting, []string and struct are alowed. ", t.Name())}
		}

	}

	return is
}

func (is *InsertStmt) Exec(values ...interface{}) error {
	if is.err != nil {
		return is.err
	}

	if values == nil {
		return fmt.Errorf("Couldn't insert empty values. ")
	}

	var lastType reflect.Type
	allSameStruct := true

	for i := range values {
		valueType := reflect.TypeOf(values[i])

		if (valueType.Kind() != reflect.Struct && valueType.Kind() != reflect.Map) || (lastType != nil && lastType == valueType) {
			allSameStruct = false
			break
		}

		lastType = valueType
	}

	rows := 1
	var querryValues []interface{}
	if allSameStruct {
		rows = len(values)
		switch v := values[0].(type) {
		case map[string]interface{}:
			log.Printf("map[string]interface{}: %#v\n", v)
		case struct{}:
			log.Printf("struct{}: %#v\n", v)
		default:
			log.Printf("default: %#v\n", v)
		}

	} else {

	}

	res, err := is.table.db.db.Exec(is.buildQuery(rows), querryValues...)
	if err != nil {
		return fmt.Errorf("Couldnt exec insert: %v", err)
	}
	res.LastInsertId()

	return nil
}

func (is *InsertStmt) Prepare(columns interface{}) error {
	if columns != nil {
		is.Columns(columns)
	} else if is.columns == nil {
		return fmt.Errorf("Couldn't prepare without columns. ")
	}
	if is.err != nil {
		return is.err
	}

	return nil

}

func trim(in []string) (out []string) {
	for _, s := range in {
		out = append(out, strings.Trim(s, " \n\t\r"))
	}
	return
}

func (is *InsertStmt) buildQuery(lenRows int) string {

	rows := []string{}

	values := []string{}
	for i := 0; i < len(is.columns); i++ {
		values = append(values, "?")
	}

	row := fmt.Sprintf("(%s)", strings.Join(values, ", "))

	for i := 0; i < lenRows; i++ {
		rows = append(rows, row)
	}

	return fmt.Sprintf("INSERT INTO %s %s VALUES %s;", is.table.name, strings.Join(is.columns, ", "), strings.Join(rows, ", "))

}

// Error returns error
func (is *InsertStmt) Error() error {
	return is.err
}

func (p *PreparedInsertStmt) Exec(values ...interface{}) error {
	p.stmt.Exec()
	return nil
}
