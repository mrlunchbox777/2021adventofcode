package main

import (
	"bufio"
	"fmt"
	// "math"
	"os"
	"strconv"
	"strings"
)

type BingoBoardLine struct {
	values []int
}

type BingoBoard struct {
	boardLines []BingoBoardLine
	answerLines []BingoBoardLine
}

func getBingoBoardLine(valueStrings []string) (BingoBoardLine) {
	var boardLine []int
	for i := 0; i < len(valueStrings); i++ {
		currentString := strings.TrimSpace(valueStrings[i])
		if currentString == ""{
			continue
		}
		currentInt, err := strconv.Atoi(currentString)
		if err != nil {
			panic(err)
		}
		boardLine = append(boardLine, currentInt)
	}
	return BingoBoardLine{ values: boardLine }
}

func getBingoBoard(lineStrings []string) (BingoBoard) {
	var boardLines []BingoBoardLine
	for i := 0; i < len(lineStrings); i++ {
		currentString := strings.TrimSpace(lineStrings[i])
		if currentString == ""{
			continue
		}
		valueStrings := strings.Split(currentString, " ")
		currentBoardLine := getBingoBoardLine(valueStrings)
		boardLines = append(boardLines, currentBoardLine)
	}
	return BingoBoard{ boardLines: boardLines }
}

func getBingoBoards(scanner *bufio.Scanner) ([]int, []BingoBoard) {
	gotWinningNumbers := false
	boardStrings := []string{}
	bingoBoards := []BingoBoard{}
	var winningNumbers []int

	for scanner.Scan() {
		i := strings.TrimSpace(scanner.Text())
		if (!gotWinningNumbers) {
			tempWinningNumbers, err := getWinningNumbers(i)
			winningNumbers = tempWinningNumbers
			if err != nil {
				panic(err)
			}
			gotWinningNumbers = true
			fmt.Println("winningNumbers - ", winningNumbers)
		} else {
			if i == "" {
				if len(boardStrings) > 0 {
					bingoBoards = append(bingoBoards, getBingoBoard(boardStrings))
					boardStrings = []string{}
				}
			} else {
				boardStrings = append(boardStrings, i)
			}
		}
	}
	if len(boardStrings) > 0 {
		bingoBoards = append(bingoBoards, getBingoBoard(boardStrings))
	}

	return winningNumbers, bingoBoards
}

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

func printAllBingoBoards(bingoBoards []BingoBoard) {
	for i := 0; i < len(bingoBoards); i++ {
		fmt.Println("Bingo Board -", i)
		for j := 0; j < len(bingoBoards[i].boardLines); j++ {
			var lineValue strings.Builder
			for k := 0; k < len(bingoBoards[i].boardLines[j].values); k++ {
				currentInt := bingoBoards[i].boardLines[j].values[k]
				nextValue := strconv.Itoa(currentInt)
				// fmt.Println("val -", nextValue, "raw -", bingoBoards[i].boardLines[j].values[k])
				if currentInt < 10 {
					lineValue.WriteString(" ")
				}
				lineValue.WriteString(nextValue)
				lineValue.WriteString(" ")
			}
			fmt.Println(lineValue.String())
		}
		fmt.Println("")
	}
	fmt.Println("number of Boards -", len(bingoBoards))
}

func main() {
	file, err := os.Open("./data/input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	winningNumbers, bingoBoards := getBingoBoards(scanner)

	printAllBingoBoards(bingoBoards)
	fmt.Println("winningNumbers - ", winningNumbers)

}

