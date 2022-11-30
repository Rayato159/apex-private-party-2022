package main

import (
	"fmt"

	"github.com/Rayato159/apex-cosplay-private-2022/src"
)

func main() {
	if err := src.FolderInit(); err != nil {
		panic(err)
	}
	fmt.Println("success, all folder have been created!!!")
}
