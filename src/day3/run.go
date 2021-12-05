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

	diagByColumn := make(map[int]int)  
	gamma := make(map[int]int)
	epsilon := make(map[int]int)
	scanner := bufio.NewScanner(file)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
		i := scanner.Text()
		ints := stringToIntArr(i)
		for j := 0; j < len(ints); j ++ {
			currentInt := ints[j]
			diagByColumn[j] += currentInt  
		}
		fmt.Println("i -", i, "ints -", ints)
	}
	halfLineCount := lineCount / 2
	for j, i := range diagByColumn {
		fmt.Println("diag by column i -", i)
		if i < halfLineCount {
			gamma[j] = 0
			epsilon[j] = 1
		} else {
			gamma[j] = 1
			epsilon[j] = 0
		}
		fmt.Println("gamma -", gamma, ", epsilon -", epsilon, ", j -", j)
	}
}
