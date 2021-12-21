package day4

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type BingoGame struct {
	bingoBoards []BingoBoard
	answers WinningNumbers
}

func CalcGame(bingoGame BingoGame) (BingoGame, error) {
	var err error
	if bingoGame == nil {
		return nil, errors.New("bingoGame was nil")
	}
	if bingoGame.bingoBoards == nil {
		return nil, errors.New("bingoGame.bingoBoards was nil")
	}
	if bingoGame.answers == nil {
		return nil, errors.New("bingoGame.answers was nil")
	}
	if len(bingoGame.bingoBoards) == 0 {
		return nil, errors.New("bingoGame.bingoBoards was empty")
	}
	if len(bingoGame.answers) == 0 {
		return nil, errors.New("bingoGame.answers was empty")
	}
	if bingoGame.answers.values == nil {
		return nil, errors.New("bingoGame.answers.values was nil")
	}
	winningNumbers := bingoGame.answers.values
	if len(winningNumbers) == 0 {
		return nil, errors.New("winningNumbers was empty")
	}

	newBoards := []BingoBoard
	for i := 0; i < len(winningNumbers); i++ {
		winningNumber := winningNumber[i]
		newBoard, newErr := getBingoBoardAnswers(bingoGame.bingoBoards, winningNumber)
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
	var err error

	for scanner.Scan() {
		i := strings.TrimSpace(scanner.Text())
		if (!gotWinningNumbers) {
			winningNumbers, newErr := getWinningNumbers(i)
			if newErr != nil {
				if (err == nil){
					err = newErr
				} else {
					err = fmt.Errorf("Combined error: %v %v", err, newErr)
				}
			}
			gotWinningNumbers = true
			if (printWinningNumbers) {
				fmt.Println("winningNumbers - ", winningNumbers.values)
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
