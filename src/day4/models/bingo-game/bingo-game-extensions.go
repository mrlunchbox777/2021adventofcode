package bingogame

import (
	"errors"
	"fmt"

	bb "github.com/mrlunchbox777/2021adventofcode/src/day4/models/bingo-board"
)

func (bingoGame BingoGame) CalcGame() (BingoGame, error) {
	var err error
	newGame := BingoGame{
		bingoBoards:   bingoGame.bingoBoards,
		answers:       bingoGame.answers,
		winningBoards: bingoGame.winningBoards,
		losingBoards:  bingoGame.losingBoards,
		rounds:        bingoGame.rounds,
	}

	if bingoGame.bingoBoards == nil {
		return newGame, errors.New("bingoGame.bingoBoards was nil")
	}
	if len(bingoGame.bingoBoards) == 0 {
		return newGame, errors.New("bingoGame.bingoBoards was empty")
	}
	if bingoGame.answers.Values() == nil {
		return newGame, errors.New("bingoGame.answers.values was nil")
	}
	winningNumbers := bingoGame.answers.Values()
	if len(winningNumbers) == 0 {
		return newGame, errors.New("winningNumbers was empty")
	}

	for i, winningNumber := range winningNumbers {
		newGameTemp, newErr := newGame.calcGameRound(winningNumber)
		if newErr != nil {
			if err == nil {
				err = newErr
			} else {
				err = fmt.Errorf("Combined error: %v %v", err, newErr)
			}
		}

		newGameTemp.rounds = i
		newGame = newGameTemp
		if len(newGame.winningBoards) == 0 {
			newGame.answers, newErr = newGame.answers.SetLatestNumber(winningNumber, false)
			if newErr != nil {
				if err == nil {
					err = newErr
				} else {
					err = fmt.Errorf("Combined error: %v %v", err, newErr)
				}
			}
		}

		newGame.answers, newErr = newGame.answers.SetLatestNumber(winningNumber, true)
		if newErr != nil {
			if err == nil {
				err = newErr
			} else {
				err = fmt.Errorf("Combined error: %v %v", err, newErr)
			}
		}
	}

	return newGame, err
}

func (bingoGame BingoGame) PrintBingoBoards(getAnswersInstead bool) {
	fmt.Println(printBingoBoardsStruct(bingoGame.bingoBoards, getAnswersInstead))
	fmt.Println("number of Boards -", len(bingoGame.bingoBoards))
}

func (bingoGame BingoGame) PrintResults(includeLoser bool) error {
	if len(bingoGame.winningBoards) == 0 {
		return errors.New("No Winning Boards")
	}

	// fmt.Println(printBingoBoardsStruct(bingoGame.bingoBoards, false))
	// fmt.Println(printBingoBoardsStruct(bingoGame.winningBoards, false))
	// fmt.Println(printBingoBoardsStruct(bingoGame.losingBoards, false))
	// getting errors here
	winningScore, err := bingoGame.findWinningScore(false)
	if err != nil {
		return err
	}
	fmt.Println(fmt.Sprintf("THE WINNING SCORE IS - %v", winningScore))

	// getting errors here
	losingScore, err := bingoGame.findWinningScore(true)
	if err != nil {
		return err
	}
	fmt.Println(fmt.Sprintf("THE LOSING SCORE IS - %v", losingScore))

	fmt.Println(fmt.Sprintf("IT TOOK %v ROUNDS", bingoGame.rounds))
	fmt.Println(fmt.Sprintf("THERE WERE (%v) WINNERS:", len(bingoGame.winningBoards)))
	fmt.Println("")
	fmt.Println(fmt.Sprintf("THE WINNING SCORE IS - %v", winningScore))
	if includeLoser {
		fmt.Println(fmt.Sprintf("THE LOSING SCORE IS - %v", losingScore))
	}

	return nil
}

func (bingoGame BingoGame) calcGameRound(winningNumber int) (BingoGame, error) {
	var err error
	newGame := BingoGame{
		bingoBoards:   bingoGame.bingoBoards,
		answers:       bingoGame.answers,
		winningBoards: bingoGame.winningBoards,
		losingBoards:  bingoGame.losingBoards,
		rounds:        bingoGame.rounds,
	}

	newBoards, newErr := bb.GetBingoBoardsAnswers(newGame.bingoBoards, winningNumber)
	newGame.bingoBoards = newBoards
	if newErr != nil {
		if err == nil {
			err = newErr
		} else {
			err = fmt.Errorf("Combined error: %v %v", err, newErr)
		}
	}

	newBoards = []bb.BingoBoard{}
	for _, board := range newGame.bingoBoards {
		potentialBoard, newErr := board.CheckForBingoBoardWin()
		newBoards = append(newBoards, potentialBoard)
		if potentialBoard.Completed() {
			newGame.winningBoards = append(bingoGame.winningBoards, potentialBoard)
		}
		if newErr != nil {
			if err == nil {
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

func (bingoGame BingoGame) findWinningScore(getLoser bool) (int, error) {
	winningBoardsLen := len(bingoGame.winningBoards)
	if winningBoardsLen == 0 {
		return 0, fmt.Errorf("winningBoardsLen length invalid (expecting >1), winningBoardsLen length - %v", winningBoardsLen)
	}

	sumOfUnmarkedNumbers, err := bingoGame.sumUnmarkedNumbersGame(getLoser)

	if err != nil {
		return 0, err
	}

	winningNumber := bingoGame.answers.GetLatestNumber(getLoser)
	winningScore := sumOfUnmarkedNumbers * winningNumber

	return winningScore, nil
}

func (bingoGame BingoGame) sumUnmarkedNumbersGame(getLoser bool) (int, error) {
	var boardsToUse []bb.BingoBoard
	if getLoser {
		boardsToUse = bingoGame.losingBoards
	} else {
		boardsToUse = bingoGame.winningBoards
	}
	if len(boardsToUse) <= 0 {
		return 0, fmt.Errorf("bad number of boardsToUse - %v", len(boardsToUse))
	}

	return boardsToUse[0].SumUnmarkedNumbersBoard(getLoser)
}
