package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func getMysqlDB() *sql.DB {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", cfg.User, cfg.Password, cfg.MysqlAddr, cfg.MysqlPort, cfg.DbName))
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func getData(cid int) (result string, err error) {

	var rows *sql.Rows

	if cid == 0 {
		rows, err = db.Query("SELECT content FROM hitokoto_data")
	} else {
		rows, err = db.Query("SELECT content FROM hitokoto_data WHERE category = (SELECT category FROM hitokoto_category WHERE cid = ?)", cid)
	}
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}

	rand.Seed(time.Now().UnixNano())

	for i, num := 0, rand.Int()%count[cid]; i <= num; i++ {
		rows.Next()
		rows.Scan(&result)
	}
	return
}

func getCount() {

	rows, err := db.Query("SELECT * FROM hitokoto_category")
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}

	var (
		c   string
		cid int
		sum int
	)
	for rows.Next() {
		rows.Scan(&cid, &c)
		log.Println(cid, c)
		var temp int
		err = db.QueryRow("SELECT COUNT(*) FROM hitokoto_data WHERE category = ?", c).Scan(&temp)
		count[cid] = temp
		if err != nil {
			log.Fatal(err)
		}
		sum += temp
	}
	count[0] = sum
	log.Println(count)
}
