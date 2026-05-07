package main

import (
	"fmt"

	"github.com/HellsKitchen99/kr/writer/internal/external"
)

func main() {
	writer := external.NewWriter("http://app:1111/system", 60)
	fmt.Println("writer has been started")
	if err := writer.Ask(); err != nil {
		fmt.Println("ОШИБКА", err)
		return
	}
}
