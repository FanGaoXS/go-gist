package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	load()
	read()
	env := NewEnv()
	fmt.Printf("%v\n", env)
}

func load() {
	if err := godotenv.Load("./env/.env"); err != nil {
		log.Fatal(err)
	}

	fmt.Println("name: ", os.Getenv("name"))
	fmt.Println("age: ", os.Getenv("age"))
}

func read() {
	envMap, err := godotenv.Read("./env/.env")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("name: ", envMap["name"])
	fmt.Println("age: ", envMap["age"])
}
