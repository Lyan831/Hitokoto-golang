package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var db = getMysqlDB()
var count = make(map[HitokotoType]int)

func main() {
	router := httprouter.New()
	registerRoute(router)
	log.Println("Hitokoto server starts")
	log.Fatal(http.ListenAndServe(":8080", router))
}
