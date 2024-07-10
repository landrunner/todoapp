package models

import (
	"database/sql"
	"reflect"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestDataSource_FetchTodos(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	tests := []struct {
		name    string
		fields  fields
		want    []Todo
		wantErr bool
	}{
		{"test", fields{}, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DataSource{
				db: tt.fields.db,
			}
			d.InitDB("testdb.sqlite3")
			got, err := d.FetchTodos()
			if (err != nil) != tt.wantErr {
				t.Errorf("DataSource.FetchTodos() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DataSource.FetchTodos() = %v, want %v", got, tt.want)
			}
		})
	}
}
