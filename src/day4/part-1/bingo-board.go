package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type BingoBoard struct {
	boardLines []BingoBoardLine
	answerLines []BingoBoardLine
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
			} else {
				err = fmt.Errorf("Combined error: %v %v", err, newErr)
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
				} else {
					err = fmt.Errorf("Combined error: %v %v", err, newErr)
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
						} else {
							err = fmt.Errorf("Combined error: %v %v", err, newErr)
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
			} else {
				err = fmt.Errorf("Combined error: %v %v", err, newErr)
			}
		}
		bingoBoards = append(bingoBoards, newBoard)
	}

	return winningNumbers, bingoBoards, err
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
