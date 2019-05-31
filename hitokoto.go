package main

import (
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// HitokotoParamsError 一言类型错误
type HitokotoParamsError struct {
	err string
}

func (e HitokotoParamsError) Error() string {
	return "Invalid Param!"
}

func hitokoto(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := r.ParseForm()
	cid, err := strconv.Atoi(r.FormValue("category"))
	if err != nil {
		cid = 0
	}

	_, ok := count[cid]
	if !ok {
		log.Println("[Hitokoto by Lyan]Failed:", HitokotoParamsError{})
		io.WriteString(w, "参数不正确")
		return
	}

	result, err := getData(cid)
	log.Println("[Hitokoto by Lyan]Success:", result)
	io.WriteString(w, result)
}
