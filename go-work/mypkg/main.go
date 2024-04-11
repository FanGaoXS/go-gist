package main

import (
	"github.com/fangaoxs/mypkg/internal/apis"
	apisv1 "github.com/fangaoxs/mypkg/internal/apis/v1"
)

func main() {
	apis.Print()
	apisv1.ListBooks()
}
