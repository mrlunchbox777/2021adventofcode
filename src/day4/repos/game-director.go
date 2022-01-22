package gamedirector

import (
	"bufio"
	"fmt"
	"strings"

	bb "github.com/mrlunchbox777/2021adventofcode/src/day4/models/bingo-board"
	bg "github.com/mrlunchbox777/2021adventofcode/src/day4/models/bingo-game"
	wn "github.com/mrlunchbox777/2021adventofcode/src/day4/models/winning-numbers"
)

func PrepGame(scanner *bufio.Scanner) (bg.BingoGame, error) {
	gotWinningNumbers := false
	boardStrings := []string{}
	bingoBoards := []bb.BingoBoard{}
	var winningNumbers wn.WinningNumbers
	var err error

	for scanner.Scan() {
		i := strings.TrimSpace(scanner.Text())
		if !gotWinningNumbers {
			tempWinningNumbers, newErr := wn.CreateWinningNumbers(i)
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

	return bg.CreateBingoGame(bingoBoards, winningNumbers), err
}
