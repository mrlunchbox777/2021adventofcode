package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
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

func reverse(numbers map[int]int) map[int]int {
	for i := 0; i < len(numbers)/2; i++ {
		j := len(numbers) - i - 1
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	return numbers
}

func getDecimalFromBinary(numbers map[int]int) float64 {
	retVal := float64(0)
	reverseNumbers := reverse(numbers)
	for i, j := range reverseNumbers {
		retVal += float64(j) * math.Pow(float64(2), float64(i))
	}
	return retVal
}

type BingoBoard struct {
	lines map[int]int
}

func getWinningNumbers(input string) (map[int]int, error) {
	winningNumbers := make(map[int]int)
	winningNumbersStringArr := strings.Split(input, ",")
	err := error(nil)
	for i := 0; i < len(winningNumbersStringArr); i++ {
		winningNumbers[i], err = strconv.Atoi(winningNumbersStringArr[i])
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("winning Numbers -", winningNumbers)
	return winningNumbers, err
}

func main() {
	file, err := os.Open("./data/input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// bingoBoards := make(map[int]BingoBoard)
	scanner := bufio.NewScanner(file)
	lineCount := 0

	for scanner.Scan() {
		lineCount++
		i := strings.TrimSpace(scanner.Text())
		if (lineCount == 1) {
			winningNumbers, err := getWinningNumbers(i)
			if err != nil {
				panic(err)
			}
			fmt.Println(winningNumbers)
		} else{
			if i != "" {
				ints := stringToIntArr(i)
				for j := 0; j < len(ints); j ++ {
					currentInt := ints[j]
					fmt.Println(currentInt);
					// diagByColumn[j] += currentInt  
				}
			} else {
				fmt.Println("emptyString")
			}
		}
	}

	fmt.Println("starting")

}

