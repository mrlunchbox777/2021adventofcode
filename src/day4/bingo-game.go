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
	losingBoards []BingoBoard
	rounds int
}

func CalcGame(bingoGame BingoGame) (BingoGame, error) {
	var err error
	newGame := BingoGame {
		bingoBoards: bingoGame.bingoBoards,
		answers: bingoGame.answers,
		winningBoards: bingoGame.winningBoards,
		losingBoards: bingoGame.losingBoards,
		rounds: bingoGame.rounds,
	}

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
		if len(newGame.winningBoards) > 0 {
			newGame.answers, newErr = setLatestNumber(newGame.answers, winningNumber, false)
			if newErr != nil {
				if (err == nil){
					err = newErr
				} else {
					err = fmt.Errorf("Combined error: %v %v", err, newErr)
				}
			}
		}

		newGame.answers, newErr = setLatestNumber(newGame.answers, winningNumber, true)
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

func PrintBingoBoards(bingoGame BingoGame, getAnswersInstead bool) {
	fmt.Println(printBingoBoardsStruct(bingoGame.bingoBoards, getAnswersInstead))
	fmt.Println("number of Boards -", len(bingoGame.bingoBoards))
}

func PrintResults(bingoGame BingoGame, includeLoser bool) (error) {
	if len(bingoGame.winningBoards) == 0 {
		return errors.New("No Winning Boards")
	}

	winningScore, err := findWinningScore(bingoGame, false)
	if err != nil {
		return err
	}

	losingScore, err := findWinningScore(bingoGame, true)
	if err != nil {
		return err
	}

	fmt.Println(fmt.Sprintf("IT TOOK %v ROUNDS", bingoGame.rounds))
	fmt.Println(fmt.Sprintf("THERE WERE (%v) WINNERS:", len(bingoGame.winningBoards)))
	// fmt.Println(printBingoBoardsStruct(bingoGame.winningBoards, false))
	fmt.Println("")
	fmt.Println(fmt.Sprintf("THE WINNING SCORE IS - %v", winningScore))
	if includeLoser {
		fmt.Println(fmt.Sprintf("THE LOSING SCORE IS - %v", losingScore))
	}

	return nil
}

func calcGameRound(bingoGame BingoGame, winningNumber int) (BingoGame, error) {
	var err error
	newGame := BingoGame {
		bingoBoards: bingoGame.bingoBoards,
		answers: bingoGame.answers,
		winningBoards: bingoGame.winningBoards,
		losingBoards: bingoGame.losingBoards,
		rounds: bingoGame.rounds,
	}

	newBoards, newErr := getBingoBoardsAnswers(newGame.bingoBoards, winningNumber)
	newGame.bingoBoards = newBoards
	if newErr != nil {
		if (err == nil){
			err = newErr
		} else {
			err = fmt.Errorf("Combined error: %v %v", err, newErr)
		}
	}

	newBoards = []BingoBoard{}
	for _, board := range newGame.bingoBoards {
		potentialBoard, newErr := checkForBingoBoardWin(board)
		newBoards = append(newBoards, potentialBoard)
		if potentialBoard.completed {
			newGame.winningBoards = append(bingoGame.winningBoards, potentialBoard)
		}
		if newErr != nil {
			if (err == nil){
				err = newErr
			} else {
				err = fmt.Errorf("Combined error: %v %v", err, newErr)
			}
		}
	}

	newGame.bingoBoards = newBoards
	newGame.losingBoards = reverseBingoBoards(bingoGame.winningBoards)

	return newGame, err
}

func findWinningScore(bingoGame BingoGame, getLoser bool) (int, error) {
	winningBoardsLen := len(bingoGame.winningBoards)
	if winningBoardsLen == 0 {
		return 0, fmt.Errorf("winningBoardsLen length invalid (expecting >1), winningBoardsLen length - %v", winningBoardsLen)
	}

	sumOfUnmarkedNumbers, err := sumUnmarkedNumbersGame(bingoGame, getLoser)

	if err != nil {
		return 0, err
	}

	winningNumber := getLatestNumber(bingoGame.answers, getLoser)
	winningScore := sumOfUnmarkedNumbers * winningNumber

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

func reverseBingoBoards(boards []BingoBoard) []BingoBoard {
	for i := 0; i < len(boards)/2; i++ {
		j := len(boards) - i - 1
		boards[i], boards[j] = boards[j], boards[i]
	}
	return boards
}

func sumUnmarkedNumbersGame(bingoGame BingoGame, getLoser bool) (int, error) {
	var boardsToUse []BingoBoard
	if getLoser {
		boardsToUse = bingoGame.losingBoards
	} else {
		boardsToUse = bingoGame.winningBoards
	}
	if len(boardsToUse) <= 0 {
		return 0, fmt.Errorf("bad number of boardsToUse - %v", len(boardsToUse))
	}

	return sumUnmarkedNumbersBoard(boardsToUse[0], getLoser)
}
