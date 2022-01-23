package bingoboard

import (
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
