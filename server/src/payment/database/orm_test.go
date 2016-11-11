package database

import (
	"fmt"
	"testing"
)

type dummyTable struct {
	Id     int `column:"id"`
	FooBar int `column:"foo_bar"`
}

func TestInsertSQL(t *testing.T) {
	dbHelper := &DBHelper{
		insertSQLCache: make(map[string]string),
	}
	fmt.Println(dbHelper.InsertSQL("dummy_table", dummyTable{}))
	fmt.Println(dbHelper.InsertSQL("dummy_table", dummyTable{}))
}
