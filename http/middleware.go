package main

import (
	"log"
	"net/http"
	"time"
)

// WithLogger 打印日志的中间件
func WithLogger(handler http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("path:%s process start...\n", r.URL.Path)
		defer func() {
			log.Printf("path:%s process end...\n", r.URL.Path)
		}()
		handler.ServeHTTP(w, r)
	}
}

// Metric 统计处理时间的中间件
func Metric(handler http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		defer func() {
			log.Printf("path:%s elapsed:%fs\n", r.URL.Path, time.Since(start).Seconds())
		}()
		time.Sleep(1 * time.Second)
		handler.ServeHTTP(w, r)
	}
}

const (
	ListenAddr = "10.0.0.101:8080"
)

func mid() {
	handler := http.NewServeMux()

	//handler.Handle("/", WithLogger(Metric(index())))
	handler.HandleFunc("/", WithLogger(Metric(index())))
	handler.HandleFunc("/ping/", WithLogger(Metric(ping())))

	s := &http.Server{
		Addr:    ListenAddr,
		Handler: handler,
	}
	log.Printf("server listen on %s", ListenAddr)
	s.ListenAndServe()
}
