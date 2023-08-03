package main

import (
	"log"

	"github.com/joho/godotenv"
)

type Env struct {
	Name string
	Age  string
}

func NewEnv() Env {
	envMap, err := godotenv.Read("./env/.env")
	if err != nil {
		log.Fatal(err)
	}

	name := "default"
	if envMap["name"] != "" {
		name = envMap["name"]
	}
	age := "1"
	if envMap["age"] != "" {
		age = envMap["age"]
	}

	return Env{
		Name: name,
		Age:  age,
	}
}
