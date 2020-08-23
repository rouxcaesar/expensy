package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Open database file and create table if one doesn't exist yet.
	db, _ := sql.Open("sqlite3", "../../internal/db/expenses.db")
	statement, _ := db.Prepare("CREATE TABLE IF NOT EXISTS expenses (id INTEGER PRIMARY KEY, desc TEXT, cost NUMBER(10), date TEXT)")
	statement.Exec()

	fmt.Println("Hello!")
}
