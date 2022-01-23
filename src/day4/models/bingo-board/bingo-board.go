package bingoboard

import (
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
	bbl "github.com/mrlunchbox777/2021adventofcode/src/day4/models/bingo-board-line"
)

type BingoBoard struct {
	answerLines []bbl.BingoBoardLine
	boardLines  []bbl.BingoBoardLine
	completed   bool
	id          string
}

func (board BingoBoard) AnswerLines() []bbl.BingoBoardLine {
	return board.answerLines
}

func (board BingoBoard) BoardLines() []bbl.BingoBoardLine {
	return board.boardLines
}

func (board BingoBoard) Completed() bool {
	return board.completed
}

func (board BingoBoard) Id() string {
	return board.id
}

//////////////////////////////////////////////////
// Original Extensions
//////////////////////////////////////////////////

func checkForBingoBoardAnswerLinesWin(lines []bbl.BingoBoardLine) (bool, error) {
	var err error
	won := false

	for _, answerLine := range lines {
		wonTemp, newErr := answerLine.CheckForBingoBoardAnswerLineWin()

		if newErr != nil {
			return won, newErr
		}

		if wonTemp {
			return wonTemp, err
		}
	}

	return won, err
}

func CreateBingoBoard(lineStrings []string) (BingoBoard, error) {
	var boardLines []bbl.BingoBoardLine
	var err error

	for _, value := range lineStrings {
		currentString := strings.TrimSpace(value)
		if currentString == "" {
			continue
		}
		valueStrings := strings.Split(currentString, " ")
		currentBoardLine, newErr := bbl.CreateBingoBoardLine(valueStrings)
		if newErr != nil {
			if err == nil {
				err = newErr
			} else {
				err = fmt.Errorf("Combined error: %v %v", err, newErr)
			}
		}
		boardLines = append(boardLines, currentBoardLine)
	}

	return BingoBoard{boardLines: boardLines, id: uuid.New().String(), completed: false}, err
}

func GetBingoBoardsAnswers(bingoBoards []BingoBoard, winningNumber int) ([]BingoBoard, error) {
	var err error
	newBoards := []BingoBoard{}

	if len(bingoBoards) == 0 {
		return newBoards, errors.New("bingoBoards count was 0")
	}

	for _, bingoBoard := range bingoBoards {
		newBoard, newErr := bingoBoard.getBingoBoardAnswers(winningNumber)
		newBoards = append(newBoards, newBoard)
		if newErr != nil {
			if err == nil {
				err = newErr
			} else {
				err = fmt.Errorf("Combined error: %v %v", err, newErr)
			}
		}
	}

	return newBoards, err
}
