package main

import (
	"URLShortener/internal/generator"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	fmt.Println(generator.Encode(6341))
}
