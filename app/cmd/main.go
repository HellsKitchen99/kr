package main

import (
	"fmt"

	"github.com/HellsKitchen99/kr/app/internal/build"
)

func main() {
	if err := build.Build(); err != nil {
		fmt.Println(err)
		return
	}
}
