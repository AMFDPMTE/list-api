package main

import (
	"database/sql"
	"fmt"

	"github.com/AMFDPMTE/list"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	fmt.Println("Hello World")
	l := list.New()
	fmt.Println(l.Serialize())

	// sqlite time!
	db, err := sql.Open("sqlite3", "./db/db.sqlite3")
	if err != nil {
		panic(err)
	}

	// select
	stmt, err := db.Prepare("SELECT * FROM lists")
	if err != nil {
		panic(err)
	}
	rows, err := stmt.Query()
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		fmt.Println(rows.Columns())
	}
}
