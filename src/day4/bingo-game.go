package day4

import (
	"bufio"
	"errors"
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
	newGame := BingoGame{bingoBoards: bingoGame.bingoBoards, answers: bingoGame.answers}

	if bingoGame.bingoBoards == nil {
		return newGame, errors.New("bingoGame.bingoBoards was nil")
	}
	if len(bingoGame.bingoBoards) == 0 {
		return newGame, errors.New("bingoGame.bingoBoards was empty")
	}
	if bingoGame.answers.values == nil {
		return newGame, errors.New("bingoGame.answers.values was nil")
	}
	winningNumbers := bingoGame.answers.values
	if len(winningNumbers) == 0 {
		return newGame, errors.New("winningNumbers was empty")
	}

	for i := 0; i < len(winningNumbers); i++ {
		if (i > 10) {
			// continue
		}
		newGameTemp, newErr := calcGameRound(newGame, winningNumbers[i])
		newGame = newGameTemp

		if newErr != nil {
			if (err == nil){
				err = newErr
			} else {
				err = fmt.Errorf("Combined error: %v %v", err, newErr)
			}
		}
	}

	return newGame, err
}

func PrepGame(scanner *bufio.Scanner) (BingoGame, error) {
	gotWinningNumbers := false
	boardStrings := []string{}
	bingoBoards := []BingoBoard{}
	var winningNumbers WinningNumbers
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

func PrintBingoBoardsAnswers(bingoGame BingoGame) {
	for i := 0; i < len(bingoGame.bingoBoards); i++ {
		fmt.Println("Bingo Board -", i)
		for j := 0; j < len(bingoGame.bingoBoards[i].answerLines); j++ {
			var lineValue strings.Builder
			for k := 0; k < len(bingoGame.bingoBoards[i].answerLines[j].values); k++ {
				currentInt := bingoGame.bingoBoards[i].answerLines[j].values[k]
				nextValue := strconv.Itoa(currentInt)
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

func calcGameRound(bingoGame BingoGame, winningNumber int) (BingoGame, error) {
	var err error
	newGame := BingoGame{bingoBoards: bingoGame.bingoBoards, answers: bingoGame.answers}

	newBoards, newErr := getBingoBoardsAnswers(newGame.bingoBoards, winningNumber)
	newGame.bingoBoards = newBoards
	if newErr != nil {
		if (err == nil){
			err = newErr
		} else {
			err = fmt.Errorf("Combined error: %v %v", err, newErr)
		}
	}
	// check for win

	return newGame, err
}
