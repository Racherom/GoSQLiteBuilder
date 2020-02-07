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
		t *Table
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
				t: testTable,
			},
			want: &InsertStmt{
				table: testTable,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &DB{
				db: tt.fields.db,
			}
			if got := db.InsertInto(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DB.InsertInto() = \ngot:  %#v,\nwant: %#v", got, tt.want)
			}
		})
	}
}

func TestInsertStmt_Exec(t *testing.T) {
	type test struct {
		a string `db:"test"`
		b int    `db:"-"`
	}
	type fields struct {
		table   *Table
		columns []string
		err     error
		query   string
	}
	type args struct {
		values interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Insert custom struct",
			args: args{
				values: test{
					a: "afds",
					b: 5,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			is := &InsertStmt{
				table:   tt.fields.table,
				columns: tt.fields.columns,
				err:     tt.fields.err,
				query:   tt.fields.query,
			}
			if err := is.Exec(tt.args.values); (err != nil) != tt.wantErr {
				t.Errorf("InsertStmt.Exec() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
