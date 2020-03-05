package store

import (
	"reflect"
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func Test_getDB(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		db   *gorm.DB
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDB(tt.args.name); !reflect.DeepEqual(got, tt.db) {
				t.Errorf("getDB() = %v, want %v", got, tt.db)
			}
		})
	}
}
