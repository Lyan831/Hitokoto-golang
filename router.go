package main

import (
	"github.com/julienschmidt/httprouter"
)

func registerRoute(r *httprouter.Router) {
	r.GET("/", hitokoto)
}
