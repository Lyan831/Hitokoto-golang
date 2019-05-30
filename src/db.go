package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func getMysqlDB() *sql.DB {
	db, err := sql.Open("mysql", "lyan:Lyan@831@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func getData(t string) (result string) {
	err := db.QueryRow("SELECT content FROM hitokoto WHERE type = ?", t).Scan(&result)
	if err != nil {
		log.Fatal(err)
	}
	return
}

// func getCount(HitokotoType) {
// 	err := db.Query("SELECT content FROM hitokoto WHERE type = ?")
// }
