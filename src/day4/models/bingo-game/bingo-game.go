package bingogame

import (
	"bufio"
	"fmt"
	"strings"

	bb "github.com/mrlunchbox777/2021adventofcode/src/day4/models/bingo-board"
	wn "github.com/mrlunchbox777/2021adventofcode/src/day4/models/winning-numbers"
)

type BingoGame struct {
	bingoBoards   []bb.BingoBoard
	answers       wn.WinningNumbers
	winningBoards []bb.BingoBoard
	losingBoards  []bb.BingoBoard
	rounds        int
}

func (game BingoGame) BingoBoards() []bb.BingoBoard {
	return game.bingoBoards
}

func (game BingoGame) Answers() wn.WinningNumbers {
	return game.answers
}

func (game BingoGame) WinningBoards() []bb.BingoBoard {
	return game.winningBoards
}

func (game BingoGame) LosingBoards() []bb.BingoBoard {
	return game.losingBoards
}

func (game BingoGame) Rounds() int {
	return game.rounds
}

//////////////////////////////////////////////////
// Original Extensions
//////////////////////////////////////////////////

func PrepGame(scanner *bufio.Scanner) (BingoGame, error) {
	gotWinningNumbers := false
	boardStrings := []string{}
	bingoBoards := []bb.BingoBoard{}
	var winningNumbers wn.WinningNumbers
	var err error

	for scanner.Scan() {
		i := strings.TrimSpace(scanner.Text())
		if !gotWinningNumbers {
			tempWinningNumbers, newErr := wn.GetWinningNumbers(i)
			winningNumbers = tempWinningNumbers
			if newErr != nil {
				if err == nil {
					err = newErr
				} else {
					err = fmt.Errorf("Combined error: %v %v", err, newErr)
				}
			}
			gotWinningNumbers = true
		} else {
			if i == "" {
				if len(boardStrings) > 0 {
					newBoard, newErr := bb.GetBingoBoard(boardStrings)
					if newErr != nil {
						if err == nil {
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
		newBoard, newErr := bb.GetBingoBoard(boardStrings)
		if newErr != nil {
			if err == nil {
				err = newErr
			} else {
				err = fmt.Errorf("Combined error: %v %v", err, newErr)
			}
		}
		bingoBoards = append(bingoBoards, newBoard)
	}

	return BingoGame{bingoBoards: bingoBoards, answers: winningNumbers}, err
}

func printBingoBoardsStruct(boards []bb.BingoBoard, getAnswersInstead bool) string {
	var gameValue strings.Builder

	for i, bingoBoard := range boards {
		if i > 0 {
			gameValue.WriteString("\n")
		}
		gameValue.WriteString(fmt.Sprintf("Bingo Board - %v\n", i))
		gameValue.WriteString(bingoBoard.GetBingoBoardPrintString(getAnswersInstead))
		gameValue.WriteString("\n")
	}

	return gameValue.String()
}

func reverseBingoBoards(boards []bb.BingoBoard) []bb.BingoBoard {
	for i := 0; i < len(boards)/2; i++ {
		j := len(boards) - i - 1
		boards[i], boards[j] = boards[j], boards[i]
	}
	return boards
}
