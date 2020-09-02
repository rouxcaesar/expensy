package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	ctl "expensy/cmd/expensyctl"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// First check whether a database file exists.
	// If it does not, then create one.

	// Open database file and create table if one doesn't exist yet.
	db, _ := sql.Open("sqlite3", "../../internal/db/expenses.db")
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS expenses (id INTEGER PRIMARY KEY, desc TEXT, cost NUMBER(10), date TEXT)")
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	statement.Exec()

	fmt.Println("Hello!")

	// Handle command
	cmd := os.Args[1]
	args := os.Args[2:]

	switch cmd {
	case "add":
		fmt.Println(cmd)
		ctl.AddCommand(args)
	default:
		fmt.Println("Default case")
	}
}
