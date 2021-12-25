package day4

import (
	"bufio"
	"errors"
	"fmt"
	"strings"
)

type BingoGame struct {
	bingoBoards []BingoBoard
	answers WinningNumbers
	winningBoards []BingoBoard
	rounds int
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

	for i, winningNumber := range winningNumbers {
		if (i > 10) {
			// continue
		}
		newGameTemp, newErr := calcGameRound(newGame, winningNumber)
		newGameTemp.rounds = i
		newGame = newGameTemp

		if newErr != nil {
			if (err == nil){
				err = newErr
			} else {
				err = fmt.Errorf("Combined error: %v %v", err, newErr)
			}
		}

		if len(newGame.winningBoards) > 0 {
			return newGame, err
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

func PrintBingoBoards(bingoGame BingoGame, getAnswersInstead bool) {
	fmt.Println(printBingoBoardsStruct(bingoGame.bingoBoards, getAnswersInstead))
	fmt.Println("number of Boards -", len(bingoGame.bingoBoards))
}

func PrintResults(bingoGame BingoGame) (error) {
	if len(bingoGame.winningBoards) == 0 {
		return errors.New("No Winning Boards")
	}

	fmt.Println(fmt.Sprintf("IT TOOK %v ROUNDS", bingoGame.rounds))
	fmt.Println(fmt.Sprintf("THERE WERE (%v) WINNERS:", len(bingoGame.winningBoards)))
	fmt.Println(printBingoBoardsStruct(bingoGame.winningBoards, false))
	return nil
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

	for _, board := range newGame.bingoBoards {
		gotWinner, newErr := checkForBingoBoardWin(board)
		if gotWinner {
			newGame.winningBoards = append(bingoGame.winningBoards, board)
		}
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

func printBingoBoardsStruct(boards []BingoBoard, getAnswersInstead bool) (string) {
	var gameValue strings.Builder

	for i, bingoBoard := range boards {
		if i > 0 {
			gameValue.WriteString("\n")
		}
		gameValue.WriteString(fmt.Sprintf("Bingo Board - %v\n", i))
		gameValue.WriteString(getBingoBoardPrintString(bingoBoard, getAnswersInstead))
		gameValue.WriteString("\n")
	}

	return gameValue.String()
}
