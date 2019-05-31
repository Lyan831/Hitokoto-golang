package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/julienschmidt/httprouter"
)

var cfg = readConfig()
var db = getMysqlDB()
var count = make(map[int]int)

// Config 包含了 MySQL 配置等相关字段
type Config struct {
	ListenPort int
	MysqlAddr  string
	MysqlPort  int
	DbName     string
	User       string
	Password   string
}

func readConfig() Config {
	jsonData, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatal(err)
	}
	var c Config
	err = json.Unmarshal(jsonData, &c)
	if err != nil {
		log.Fatal(err)
	}
	return c
}

func main() {

	defer db.Close()

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-ch
		log.Printf("[Hitokoto by Lyan]Hitokoto server stops now...")
		db.Close()
		os.Exit(0)
	}()

	getCount()
	router := httprouter.New()
	registerRoute(router)
	log.Printf("[Hitokoto by Lyan]Hitokoto server is running on port %d...", cfg.ListenPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", cfg.ListenPort), router))
}
