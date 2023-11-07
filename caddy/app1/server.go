package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	AppName    = "app1"
	ListenAddr = "10.0.0.1:8080"
)

func serve() {
	handler := gin.Default()
	handler.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome %s!", AppName)
	})
	handler.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "%s pong!", AppName)
	})

	s := &http.Server{
		Addr:    ListenAddr,
		Handler: handler,
	}

	log.Printf("%s listen on %s", AppName, ListenAddr)
	s.ListenAndServe()
}
