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

	diagByColumn := make(map[int]int, 0)  
	// gamma := make(map[int]int, 0)
	// epsilon := make(map[int]int)
	scanner := bufio.NewScanner(file)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
		i := scanner.Text()
		ints := stringToIntArr(i)
		// for each char add to corresponding map in gamma for each value
		for j := 0; j < len(ints); j ++ {
			currentInt := ints[j]
			diagByColumn[j] += currentInt  
		}
		fmt.Println("i -", i, "ints -", ints)
		// fmt.Println("gamma -", gamma, "epsilon -", epsilon)
		fmt.Println("lineCount -", lineCount, "diagByColumn -", diagByColumn)
	}
}
