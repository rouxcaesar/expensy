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
	_, err := os.Stat("../../internal/db/expenses.db")
	if os.IsNotExist(err) {
		_, err := os.Create("../../internal/db/expenses.db")
		if err != nil {
			fmt.Println(err)
			log.Fatal(err)
		}
	}

	//if info.IsDir() {
	//	fmt.Println(err)
	//	log.Fatal(err)
	//}

	// Open database file and create table if one doesn't exist yet.
	db, _ := sql.Open("sqlite3", "../../internal/db/expenses.db")
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS expenses (id INTEGER PRIMARY KEY, desc TEXT, cost NUMBER(10), date TEXT)")
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	statement.Exec()

	// Validate command.
	if len(os.Args) < 2 {
		log.Fatalln("Command needed to run program.")
	}

	// Handle command.
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
