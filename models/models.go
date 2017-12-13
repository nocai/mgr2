package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

const (
	username = "root"
	password = "root"
)

var db *sql.DB

func init() {
	db = getDB()
}

func getDB() *sql.DB {
	d, err := sql.Open("mysql", username + ":" + password + "@tcp(localhost:3306)/mgr?charset=utf8")
	if err != nil {
		panic(err.Error())
	}

	d.SetMaxOpenConns(2000)
	d.SetMaxIdleConns(1000)
	d.Ping()

	return d
}

func GetDB() *sql.DB {
	if db == nil {
		db = getDB()
	}
	return db
}

