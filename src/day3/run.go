package main

import (
	"fmt"
	"os"
	"bufio"
	// "strconv"
	// "strings"
)

func main() {
	file, err := os.Open("./data/input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// gamma := make([]map[int]int, 0)
	// epsilon := make(map[int]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i := scanner.Text()
		chars := []rune(i)
		ints := []int{}
		for _, currentChar := range chars {
			currentInt := int(currentChar)
			ints = append(ints, currentInt)
		}
			// find number of occurrences of each value in string and put into epsilon map
			// for each char add to corresponding map in gamma for each value
		fmt.Println("i -", i, "ints -", ints)
	}
}
