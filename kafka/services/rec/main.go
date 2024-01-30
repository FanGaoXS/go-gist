package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fangaoxs/go-gist/kafka/internal/reader"
)

const (
	ADDR  = "localhost:9092"
	GROUP = "rec_team"
	TOPIC = "user-input"
)

func main() {
	r, err := reader.New(ADDR, GROUP, TOPIC)
	if err != nil {
		log.Fatal(err)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-c
		fmt.Printf("接收到程序信号: %s", sig.String())
		r.Close()
		os.Exit(0)
	}()

	ctx := context.Background()
	for {
		message, err := r.Read(ctx)
		if err != nil {
			fmt.Println(err)
			fmt.Println("wait for 10 seconds")
			time.Sleep(10 * time.Second)
			continue
		}
		fmt.Println("rec_team: message: ", message)
		// do something as recommend team.
	}
}
