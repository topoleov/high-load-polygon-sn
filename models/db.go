package models

import (
	"database/sql"
	"log"

	// _ "github.com/lib/pq"
	_ "github.com/go-sql-driver/mysql"
)

var dbS *sql.DB
var dbM *sql.DB

func init() {
	var err error
	dbM, err = sql.Open("mysql", "go_app:go_pass@tcp(mysqlmaster:3306)/go_app_db")
	if err != nil {
		log.Fatal(err)
	}
	dbS, err = sql.Open("mysql", "go_app:go_pass@tcp(mysqlslave:3306)/go_app_db")
	if err != nil {
		log.Fatal(err)
	}

}
