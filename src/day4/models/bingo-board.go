package models

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

func checkForBingoBoardWin(bingoBoard BingoBoard) (BingoBoard, error) {
	if bingoBoard.completed {
		return bingoBoard, nil
	}

	newBoard := BingoBoard{
		answerLines: bingoBoard.answerLines,
		boardLines:  bingoBoard.boardLines,
		id:          bingoBoard.id,
		completed:   bingoBoard.completed,
	}
	won, err := checkForBingoBoardAnswerLinesWin(bingoBoard.answerLines)
	if won && err != nil {
		newBoard.completed = won
		return newBoard, err
	}

	newLines := []bbl.BingoBoardLine{}
	for i := range bingoBoard.answerLines[0].Values() {
		newLine := bbl.BingoBoardLine{}
		for _, currentLine := range bingoBoard.answerLines {
			newLine.SetValues(append(newLine.Values(), currentLine.Values()[i]))
		}
		newLines = append(newLines, newLine)
	}

	won, err = checkForBingoBoardAnswerLinesWin(newLines)
	newBoard.completed = won
	return newBoard, err
}

func getBingoBoardPrintString(board BingoBoard, getAnswersInstead bool) string {
	var boardValue strings.Builder
	var bingoBoardLen int
	if getAnswersInstead {
		bingoBoardLen = len(board.answerLines)
	} else {
		bingoBoardLen = len(board.boardLines)
	}

	for i := 0; i < bingoBoardLen; i++ {
		if getAnswersInstead {
			boardValue.WriteString(board.answerLines[i].GetBingoBoardLinePrintString())
		} else {
			boardValue.WriteString(board.boardLines[i].GetBingoBoardLinePrintString())
		}
		if i < (bingoBoardLen - 1) {
			boardValue.WriteString("\n")
		}
	}

	return boardValue.String()
}

func getBingoBoard(lineStrings []string) (BingoBoard, error) {
	var boardLines []bbl.BingoBoardLine
	var err error

	for _, value := range lineStrings {
		currentString := strings.TrimSpace(value)
		if currentString == "" {
			continue
		}
		valueStrings := strings.Split(currentString, " ")
		currentBoardLine, newErr := bbl.GetBingoBoardLine(valueStrings)
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

func getBingoBoardAnswers(bingoBoard BingoBoard, winningNumber int) (BingoBoard, error) {
	var err error
	newBoard := BingoBoard{
		boardLines: bingoBoard.boardLines,
		id:         bingoBoard.id,
		completed:  bingoBoard.completed,
	}

	if newBoard.completed {
		return newBoard, nil
	}

	bingoBoardLinesLen := len(bingoBoard.boardLines)
	if bingoBoardLinesLen == 0 {
		return newBoard, errors.New("bingoBoard.boardLines count was 0")
	}

	if len(bingoBoard.answerLines) == 0 {
		bingoBoard.answerLines = make([]bbl.BingoBoardLine, bingoBoardLinesLen)
	}

	newBoard.answerLines = []bbl.BingoBoardLine{}
	for i, boardLine := range bingoBoard.boardLines {
		newAnswer, newErr := boardLine.GetBingoBoardLineAnswer(bingoBoard.answerLines[i], winningNumber)
		newBoard.answerLines = append(newBoard.answerLines, newAnswer)
		if newErr != nil {
			if err == nil {
				err = newErr
			} else {
				err = fmt.Errorf("Combined error: %v %v", err, newErr)
			}
		}
	}

	return newBoard, err
}

func getBingoBoardsAnswers(bingoBoards []BingoBoard, winningNumber int) ([]BingoBoard, error) {
	var err error
	newBoards := []BingoBoard{}

	if len(bingoBoards) == 0 {
		return newBoards, errors.New("bingoBoards count was 0")
	}

	for _, bingoBoard := range bingoBoards {
		newBoard, newErr := getBingoBoardAnswers(bingoBoard, winningNumber)
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

func sumUnmarkedNumbersBoard(bingoBoard BingoBoard, getLoser bool) (int, error) {
	bingoBoardLinesLen := len(bingoBoard.boardLines)
	bingoBoardAnswersLen := len(bingoBoard.answerLines)
	var err error
	var newErr error
	sum := 0
	currentSum := 0

	if bingoBoardLinesLen <= 0 {
		return 0, fmt.Errorf("bingoBoard.boardLines length was less than or equal to 0 - %v", bingoBoardLinesLen)
	}
	if bingoBoardLinesLen != bingoBoardAnswersLen {
		return 0, fmt.Errorf("bingoBoard.boardLines length (%v) didn't equal bingoBoard.answersLines length (%v)", bingoBoardLinesLen, bingoBoardAnswersLen)
	}

	for i, line := range bingoBoard.boardLines {
		currentSum, newErr = line.SumUnmarkedNumbersBoardLine(bingoBoard.answerLines[i], getLoser)
		if newErr != nil {
			if err == nil {
				err = newErr
			} else {
				err = fmt.Errorf("Combined error: %v %v", err, newErr)
			}
		}
		sum += currentSum
	}

	return sum, err
}
