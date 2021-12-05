package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
)

func main() {
	file, err := os.Open("./data/input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i := scanner.Text()
		fmt.Println("i -", i)
	}
}
