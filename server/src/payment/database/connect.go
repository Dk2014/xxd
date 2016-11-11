package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	g_db *sql.DB
)

func Init(connectString string) error {
	if g_db != nil {
		panic("dup init")
	}
	db, err := sql.Open("mysql", connectString)
	if err != nil {
		return err
	}
	g_db = db
	return nil
}
