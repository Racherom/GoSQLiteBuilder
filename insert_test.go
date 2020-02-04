package sqlite

import (
	"database/sql"
	"reflect"
	"testing"
)

func TestDB_InsertInto(t *testing.T) {
	type test struct {
		a string `db:"test"`
		b int    `db:"-"`
	}

	testTable := &Table{
		name: "test",
	}
	type fields struct {
		db *sql.DB
	}
	type args struct {
		t        *Table
		columnsi interface{}
		values   []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *InsertStmt
	}{
		{
			name: "Struct as Columns",
			fields: fields{
				db: nil,
			},
			args: args{
				t:        testTable,
				columnsi: &test{},
				values:   nil,
			},
			want: &InsertStmt{
				table:    testTable,
				colunmns: []string{"test"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &DB{
				db: tt.fields.db,
			}
			if got := db.InsertInto(tt.args.t, tt.args.columnsi, tt.args.values...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DB.InsertInto() = \ngot:  %#v,\nwant: %#v", got, tt.want)
			}
		})
	}
}
