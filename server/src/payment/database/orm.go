package database

import (
	"fmt"
	"reflect"
	"strings"
	"sync"
)

var (
	g_Hepler *DBHelper
)

func init() {
	g_Hepler = &DBHelper{
		insertSQLCache: make(map[string]string),
	}
}

type DBHelper struct {
	sync.Mutex
	insertSQLCache map[string]string //map struct name to inser sql statement
}

func (helper *DBHelper) InsertSQL(tableName string, val interface{}) string {
	helper.Lock()
	defer helper.Unlock()
	if sql, ok := helper.insertSQLCache[tableName]; ok {
		return sql
	}
	sql := "insert into `%s` (%s) values (%s)"
	val = dereference(val)
	typ := reflect.TypeOf(val)
	fields := []string{}
	slots := []string{}
	for i := 0; i < typ.NumField(); i++ {
		columnName := typ.Field(i).Tag.Get("column")
		if len(columnName) > 0 && columnName != "-" {
			fields = append(fields, columnName)
			slots = append(slots, "?")
		}
	}
	sql = fmt.Sprintf(sql, tableName, strings.Join(fields, ","), strings.Join(slots, ","))
	helper.insertSQLCache[tableName] = sql
	return sql
}

func dereference(intptr interface{}) interface{} {
	var t reflect.Type
	var v reflect.Value
	t = reflect.TypeOf(intptr)
	v = reflect.ValueOf(intptr)
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}
	return v.Interface()
}
