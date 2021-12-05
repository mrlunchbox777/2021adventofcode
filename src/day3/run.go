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

	x := 0
	y := 0
	aim := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i := scanner.Text()
	}
}
