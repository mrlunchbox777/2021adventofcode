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

func getBingoBoardLine(valueStrings []string) (BingoBoardLine, error) {
	var boardLine []int
	var err error
	for i := 0; i < len(valueStrings); i++ {
		currentString := strings.TrimSpace(valueStrings[i])
		if currentString == ""{
			continue
		}
		currentInt, newErr := strconv.Atoi(currentString)
		if newErr != nil {
			if (err != nil){
				err = newErr
				fmt.Println("7", err)
			} else {
				err = fmt.Errorf("Combined error: %v %v", err, newErr)
				fmt.Println("8", err)
			}
		}
		boardLine = append(boardLine, currentInt)
	}

	return BingoBoardLine{ values: boardLine }, err
}

func getBingoBoard(lineStrings []string) (BingoBoard, error) {
	var boardLines []BingoBoardLine
	var err error

	for i := 0; i < len(lineStrings); i++ {
		currentString := strings.TrimSpace(lineStrings[i])
		if currentString == ""{
			continue
		}
		valueStrings := strings.Split(currentString, " ")
		currentBoardLine, newErr := getBingoBoardLine(valueStrings)
		if newErr != nil {
			if (err != nil){
				err = newErr
				fmt.Println("9", err)
			} else {
				err = fmt.Errorf("Combined error: %v %v", err, newErr)
				fmt.Println("10", err)
			}
		}
		boardLines = append(boardLines, currentBoardLine)
	}

	return BingoBoard{ boardLines: boardLines }, err
}

func getBingoBoards(scanner *bufio.Scanner) ([]int, []BingoBoard, error) {
	gotWinningNumbers := false
	boardStrings := []string{}
	bingoBoards := []BingoBoard{}
	var winningNumbers []int
	var err error

	for scanner.Scan() {
		i := strings.TrimSpace(scanner.Text())
		if (!gotWinningNumbers) {
			tempWinningNumbers, newErr := getWinningNumbers(i)
			winningNumbers = tempWinningNumbers
			if newErr != nil {
				if (err != nil){
					err = newErr
					fmt.Println("11", err)
				} else {
					err = fmt.Errorf("Combined error: %v %v", err, newErr)
					fmt.Println("12", err)
				}
			}
			gotWinningNumbers = true
			fmt.Println("winningNumbers - ", winningNumbers)
		} else {
			if i == "" {
				if len(boardStrings) > 0 {
					newBoard, newErr := getBingoBoard(boardStrings)
					if newErr != nil {
						if (err != nil){
							err = newErr
							fmt.Println("13", err)
						} else {
							err = fmt.Errorf("Combined error: %v %v", err, newErr)
							fmt.Println("14", err)
						}
					}
					bingoBoards = append(bingoBoards, newBoard)
					boardStrings = []string{}
				}
			} else {
				boardStrings = append(boardStrings, i)
			}
		}
	}
	if len(boardStrings) > 0 {
		newBoard, newErr := getBingoBoard(boardStrings)
		if newErr != nil {
			if (err != nil){
				err = newErr
				fmt.Println("1", err)
			} else {
				err = fmt.Errorf("Combined error: %v %v", err, newErr)
				fmt.Println("2", err)
			}
		}
		bingoBoards = append(bingoBoards, newBoard)
	}

	return winningNumbers, bingoBoards, err
}

func stringToIntArr(str string) ([]int, error) {
	chars := []rune(str)
	ints := []int{}
	var err error

	for i := 0; i < len(chars); i++ {
		currentChar := string(chars[i])
		currentInt, newErr := strconv.Atoi(currentChar)
		if newErr != nil {
			if (err != nil){
				err = newErr
				fmt.Println("3", err)
			} else {
				err = fmt.Errorf("Combined error: %v %v", err, newErr)
				fmt.Println("4", err)
			}
		}
		ints = append(ints, currentInt)
	}

	return ints, err
}

func getWinningNumbers(input string) ([]int, error) {
	var winningNumbers []int
	winningNumbersStringArr := strings.Split(input, ",")
	err := error(nil)
	for i := 0; i < len(winningNumbersStringArr); i++ {
		currentInt, newErr := strconv.Atoi(winningNumbersStringArr[i])
		if newErr != nil {
			if (err != nil){
				err = newErr
				fmt.Println("5", err)
			} else {
				err = fmt.Errorf("Combined error: %v %v", err, newErr)
				fmt.Println("6", err)
			}
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
	winningNumbers, bingoBoards, err := getBingoBoards(scanner)
	if (err != nil){
		panic(err)
	}

	printAllBingoBoards(bingoBoards)
	fmt.Println("winningNumbers - ", winningNumbers)

}

