package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql" // 没有导入会报错
)

// sample
// https://github.com/go-sql-driver/mysql/wiki/Examples

func main() {
	db, err := sql.Open("mysql", "root:12345678@tcp(172.28.104.225:32771)/router")
	if err != nil {
		log.Println(err)
		return
	}
	defer db.Close()

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		log.Println(err)
		return
	}
	rows, err := db.Query("SELECT * FROM member_state")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	log.Println(columns)
}
