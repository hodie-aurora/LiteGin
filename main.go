package main

import (
	"lg"
	"net/http"
)

func main() {
	r := lg.New()
	r.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("hello, World!"))
	})
	r.GET("/hi", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("hi, World!"))
	})
	r.POST("/BAZINGA", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("BAZINGA, World!"))
	})
	r.Run(":7777")
}
