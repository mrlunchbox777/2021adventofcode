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
	boardLines [5][5]int
	answerLines [5][5]int
}

func getWinningNumbers(input string) ([]int, error) {
	var winningNumbers []int
	winningNumbersStringArr := strings.Split(input, ",")
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

	bingoBoards := make(map[int]BingoBoard)
	scanner := bufio.NewScanner(file)
	lineCount := 0
	currentBingoBoard := 0
	currentBingoBoardRow := 0
	currentBingoBoardColumn := 0

	for scanner.Scan() {
		lineCount++
		i := strings.TrimSpace(scanner.Text())
		if i == "" {
			continue
		}
		if (lineCount == 1) {
			winningNumbers, err := getWinningNumbers(i)
			if err != nil {
				panic(err)
			}
			fmt.Println("winningNumbers - ", winningNumbers)
		} else {
			intsStrings := strings.Split(i, " ")
			currentBingoBoardStruct := BingoBoard{}
			for j := 0; j < len(intsStrings); j ++ {
				currentString := strings.TrimSpace(intsStrings[j])
				if currentString == ""{
					continue
				}
				currentInt, err := strconv.Atoi(currentString)
				if err != nil {
					panic(err)
				}
				fmt.Println("")
				fmt.Println("------------------------------------------------------")
				if currentBingoBoardColumn == 5 {
					currentBingoBoardRow++
					currentBingoBoardColumn = 0
				}
				if currentBingoBoardRow == 5 {
					oldBingoBoardStruct := currentBingoBoardStruct
					bingoBoards[currentBingoBoard] = oldBingoBoardStruct
					currentBingoBoardStruct = BingoBoard {}
					currentBingoBoard++
					currentBingoBoardRow = 0
				}
				fmt.Println("currentBingoBoard -", currentBingoBoard)
				fmt.Println("currentBingoBoardRow -", currentBingoBoardRow)
				fmt.Println("currentBingoBoardColumn -", currentBingoBoardColumn)
				fmt.Println("currentInt -", currentInt)
				fmt.Println("------------------------------------------------------")
				currentBingoBoardStruct.boardLines[currentBingoBoardRow][currentBingoBoardColumn] = currentInt
				currentBingoBoardColumn++
			}
		}

		for i := 0; i < len(bingoBoards); i++ {
			fmt.Println("Bingo Board - ", i)
			for j := 0; j < len(bingoBoards[i].boardLines); j++ {
				var lineValue strings.Builder
				for k := 0; k < len(bingoBoards[i].boardLines[j]); k++ {
					nextValue := strconv.Itoa(bingoBoards[i].boardLines[j][k])
					// fmt.Println("val -", nextValue, "raw -", bingoBoards[i].boardLines[j][k])
					lineValue.WriteString(nextValue)
				}
				fmt.Println(lineValue.String())
			}
			fmt.Println("")
		}
	}
}

