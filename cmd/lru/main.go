package main

import (
	"fmt"

	"github.com/realroy/simple-lru-go/lru"
)

func main() {
	lru := lru.NewLRU(10)

	fmt.Println(lru)
}
