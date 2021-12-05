package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	// "strings"
)

func stringToIntArr(str string) ([]int) {
	chars := []rune(str)
	ints := []int{}
	for i := 0; i < len(chars); i++ {
		currentChar := string(chars[i])
		currentInt, err := strconv.Atoi(currentChar)
		if err != nil {
			panic(err)
		}
		ints = append(ints, currentInt)
	}
	return ints
}

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
		ints := stringToIntArr(i)
		// find number of occurrences of each value in string and put into epsilon map
		// for each char add to corresponding map in gamma for each value
		fmt.Println("i -", i, "ints -", ints)
	}
}
