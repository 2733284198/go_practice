package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	var h Handle
	h.dst = "/tmp/test.txt"
	router.POST("/upload", h.uploadFile)

	s := &http.Server{Addr: "localhost:8088", Handler: router}
	if err := s.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
