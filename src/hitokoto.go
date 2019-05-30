package main

import (
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// HitokotoType 一言类型
type HitokotoType int

// HitokotoTypeError 一言类型错误
type HitokotoTypeError struct {
	err string
}

// 定义一言类型常量
const (
	All HitokotoType = iota
	Anime
	Games
)

func (e HitokotoTypeError) Error() string {
	return "无效的一言类型"
}

func hitokoto(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := r.ParseForm()
	t, err := strconv.Atoi(r.FormValue("type"))
	if err != nil {
		t = int(All)
	}
	log.Println("Type:", t)

	ht, err := hitokotoType(HitokotoType(t))
	if err != nil {
		log.Println("Error:", err)
		http.NotFound(w, r)
		return
	}

	result := getData(ht)
	log.Println("Success:", result)
	io.WriteString(w, result)
}

func hitokotoType(t HitokotoType) (string, error) {
	switch t {
	case All:
		return "All", nil
	case Anime:
		return "Anime", nil
	case Games:
		return "Games", nil
	}
	return "", HitokotoTypeError{}
}
