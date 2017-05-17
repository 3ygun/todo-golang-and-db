package main

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dbName  = "todoApp"
	dbTable = "todos"
)

func CreateDatabase() *sql.DB {
	var password = os.Getenv("MYSQL_ROOT_PASSWORD")

	// Create the database handle, confirm driver is present
	// Use ?parseTime=true to get correct DATETIME to time.Time parsing
	db, err := sql.Open("mysql", "root:"+password+"@/?parseTime=true")
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	initializeDatabase(db)

	return db
}

func initializeDatabase(db *sql.DB) {
	var err error
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + dbName)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("USE " + dbName)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS " + dbTable + " ( " +
		"id INTEGER NOT NULL AUTO_INCREMENT, " +
		"name VARCHAR(255), " +
		"completed BOOLEAN, " +
		"due DATETIME, " +
		"PRIMARY KEY (id) )")
	if err != nil {
		panic(err)
	}
}
