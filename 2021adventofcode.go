package main

import (
	"math/rand"
	"time"
	"fmt"
	"os"
	"log"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	content, err := os.ReadFile("./data/data-day1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(content))
}