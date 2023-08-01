package main

import "net/http"

func index() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		msg := "Hello, world!"
		w.Write([]byte(msg))
	}
}

func ping() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		msg := "pong"
		w.Write([]byte(msg))
	}
}

func serve1() {
	http.HandleFunc("/", index())
	http.ListenAndServe(":8080", nil)
}

func serve2() {
	handler := http.NewServeMux()
	handler.HandleFunc("/", index())
	s := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}
	s.ListenAndServe()
}
