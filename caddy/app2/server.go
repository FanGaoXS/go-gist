package main

import (
	"log"
	"net/http"
)

const (
	AppName    = "app2"
	ListenAddr = "10.0.0.2:8080"
)

func serve() {
	handler := http.NewServeMux()
	handler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		msg := AppName + " Welcome!"
		w.Write([]byte(msg))
	})
	handler.HandleFunc("/ping/", func(w http.ResponseWriter, r *http.Request) {
		msg := AppName + " pong"
		w.Write([]byte(msg))
	})

	s := &http.Server{
		Addr:    ListenAddr,
		Handler: handler,
	}

	log.Printf("%s listen on %s", AppName, ListenAddr)
	s.ListenAndServe()
}
