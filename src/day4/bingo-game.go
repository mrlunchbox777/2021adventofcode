package day4

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type BingoGame struct {
	bingoBoards []BingoBoard
	answers []int
}

func GetBingoBoardsAnswers(bingoBoards BingoBoard[], winningNumbers []int) ([]BingoBoards, error) {
	var err error
	if winningNumbers == nil {
		return nil, errors.New("bingoBoards array was nil")
	}
	winningNumbersLen := len(winningNumbers)
	if  winningNumbersLen == 0 {
		return nil, errors.New("bingoBoards array was empty")
	}

	newBoards := []BingoBoard
	for i := 0; i < bingoBoardLen; i++ {
		newBoard, newErr := getBingoBoardAnswers(bingoBoards[i], winningNumber)
		if newErr != nil {
			if (err == nil){
				err = newErr
			} else {
				err = fmt.Errorf("Combined error: %v %v", err, newErr)
			}
		}

		newBoards := append(newBoards, newBoard)
	}

	return newBoards, err
}

func PrepGame(scanner *bufio.Scanner, printWinningNumbers bool) (bingoGame BingoGame) {
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
				if (err == nil){
					err = newErr
				} else {
					err = fmt.Errorf("Combined error: %v %v", err, newErr)
				}
			}
			gotWinningNumbers = true
			if (printWinningNumbers) {
				fmt.Println("winningNumbers - ", winningNumbers)
			}
		} else {
			if i == "" {
				if len(boardStrings) > 0 {
					newBoard, newErr := getBingoBoard(boardStrings)
					if newErr != nil {
						if (err == nil){
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
			if (err == nil){
				err = newErr
			} else {
				err = fmt.Errorf("Combined error: %v %v", err, newErr)
			}
		}
		bingoBoards = append(bingoBoards, newBoard)
	}

	return BingoGame{ bingoBoards: bingoBoards, answers: winningNumbers}, err
}

func PrintBingoBoards(bingoGame BingoGame) {
	for i := 0; i < len(bingoGame.bingoBoards); i++ {
		fmt.Println("Bingo Board -", i)
		for j := 0; j < len(bingoGame.bingoBoards[i].boardLines); j++ {
			var lineValue strings.Builder
			for k := 0; k < len(bingoGame.bingoBoards[i].boardLines[j].values); k++ {
				currentInt := bingoGame.bingoBoards[i].boardLines[j].values[k]
				nextValue := strconv.Itoa(currentInt)
				// fmt.Println("val -", nextValue, "raw -", bingoGame.bingoBoards[i].boardLines[j].values[k])
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
	fmt.Println("number of Boards -", len(bingoGame.bingoBoards))
}
