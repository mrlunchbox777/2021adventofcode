package main

import (
	"bufio"
	"fmt"
	// "math"
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

type BingoBoard struct {
	sourceLines []int
	answerLines []int
}

func getWinningNumbers(input string) ([]int, error) {
	var winningNumbers []int
	winningNumbersStringArr := strings.Split(input, ",")
	fmt.Println("winningNumbersStringArr -", winningNumbersStringArr)
	err := error(nil)
	for i := 0; i < len(winningNumbersStringArr); i++ {
		currentInt, err := strconv.Atoi(winningNumbersStringArr[i])
		if err != nil {
			panic(err)
		}
		winningNumbers = append(winningNumbers, currentInt)
	}
	return winningNumbers, err
}

func main() {
	file, err := os.Open("./data/input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// var bingoBoards []BingoBoard
	scanner := bufio.NewScanner(file)
	lineCount := 0

	for scanner.Scan() {
		lineCount++
		i := strings.TrimSpace(scanner.Text())
		if i == "" {
			fmt.Println("emptyString")
			continue
		}
		if (lineCount == 1) {
			winningNumbers, err := getWinningNumbers(i)
			if err != nil {
				panic(err)
			}
			fmt.Println("winningNumbers -", winningNumbers)
		} else{
			intsStrings := strings.Split(i, " ")
			// bingoBoardsCount := len(bingoBoards)
			for j := 0; j < len(intsStrings); j ++ {
				if intsStrings[j] == "" {
					fmt.Println("emptyString")
					continue
				}
				currentInt, err := strconv.Atoi(intsStrings[j])
				if err != nil {
					panic(err)
				}
				fmt.Println(currentInt);
			}
		}
	}

	fmt.Println("starting")

}

