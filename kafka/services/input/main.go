package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/fangaoxs/go-gist/kafka/internal/writer"
)

const (
	ADDR  = "localhost:9092"
	TOPIC = "user-input"
)

func main() {
	w, err := writer.New(ADDR, TOPIC)
	if err != nil {
		log.Fatal(err)
	}
	defer w.Close()

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("请输入内容，退出输入exit：")
	for scanner.Scan() {
		if err = scanner.Err(); err != nil {
			fmt.Printf("读取用户输入失败：%t\n", err)
			break
		}
		input := scanner.Text()
		if input == "exit" {
			break
		}

		ctx := context.Background()
		if err = w.Write(ctx, input); err != nil {
			fmt.Println(err)
			break
		}
	}
}
