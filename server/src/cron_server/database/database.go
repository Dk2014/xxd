package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	g_db *sql.DB
	o_db *sql.DB
)

func Init(connectString string) error {
	db, err := sql.Open("mysql", connectString)
	if err != nil {
		return err
	}
	g_db = db
	return nil
}

func OnlineInit(connectString string) error {
	if o_db != nil {
		panic("dup init")
	}
	onlinedb, err := sql.Open("mysql", connectString)
	if err != nil {
		return err
	}
	o_db = onlinedb
	return nil
}
