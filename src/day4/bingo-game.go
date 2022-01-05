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
		if newErr != nil {
			if (err == nil){
				err = newErr
			} else {
				err = fmt.Errorf("Combined error: %v %v", err, newErr)
			}
		}

		newGameTemp.rounds = i
		newGame = newGameTemp
		newGame.answers, newErr = setLatestWinningNumber(newGame.answers, winningNumber)

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

	winningScore, err := findWinningScore(bingoGame)
	if err != nil {
		return err
	}

	fmt.Println(fmt.Sprintf("IT TOOK %v ROUNDS", bingoGame.rounds))
	fmt.Println(fmt.Sprintf("THERE WERE (%v) WINNERS:", len(bingoGame.winningBoards)))
	fmt.Println(printBingoBoardsStruct(bingoGame.winningBoards, false))
	fmt.Println("")
	fmt.Println(fmt.Sprintf("THE WINNING SCORE IS - %v", winningScore))

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

func findWinningScore(bingoGame BingoGame) (int, error) {
	winningBoardsLen := len(bingoGame.winningBoards)
	if winningBoardsLen > 1 || winningBoardsLen == 0 {
		return 0, fmt.Errorf("winningBoardsLen length invalid (expecting 1), winningBoardsLen length - %v", winningBoardsLen)
	}

	sumOfUnmarkedNumbers, err := sumUnmarkedNumbersGame(bingoGame)

	if err != nil {
		return 0, err
	}

	winningNumber := getLatestWinningNumber(bingoGame.answers)
	winningScore := sumOfUnmarkedNumbers * winningNumber

	fmt.Println(fmt.Sprintf("winningNumber - %v, sumOfUnmarkedNumbers - %v, winningScore - %v", winningNumber, sumOfUnmarkedNumbers, winningScore))

	return winningScore, nil
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

func sumUnmarkedNumbersGame(bingoGame BingoGame) (int, error) {
	if len(bingoGame.winningBoards) <= 0 {
		return 0, fmt.Errorf("bad number of winning boards - %v", len(bingoGame.winningBoards))
	}

	return sumUnmarkedNumbersBoard(bingoGame.winningBoards[0])
}
