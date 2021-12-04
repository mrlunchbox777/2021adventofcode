package main

import (
)

func main() {
	file, err := os.Open("./data/input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
}
