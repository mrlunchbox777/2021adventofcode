package day4

import (
	"errors"
	"fmt"
	"strings"
)

type BingoBoard struct {
	boardLines []BingoBoardLine
	answerLines []BingoBoardLine
}

func checkForBingoBoardAnswerLinesWin(lines []BingoBoardLine) (bool, error) {
	var err error
	won := false

	for _, answerLine := range lines {
		wonTemp, newErr := checkForBingoBoardAnswerLineWin(answerLine)

		if newErr != nil {
			return won, newErr
		}

		if wonTemp {
			return wonTemp, err
		}
	}

	return won, err
}

func checkForBingoBoardWin(board BingoBoard) (bool, error) {
	won, err := checkForBingoBoardAnswerLinesWin(board.answerLines) 
	if won || err != nil {
		return won, err
	}

	newLines := []BingoBoardLine{}
	for i, _ := range board.answerLines[0].values {
		newLine := BingoBoardLine{}
		for _, currentLine := range board.answerLines {
			newLine.values = append(newLine.values, currentLine.values[i])
		}
		newLines = append(newLines, newLine)
	}

	won, err = checkForBingoBoardAnswerLinesWin(newLines) 
	return won, err
}

func getBingoBoardPrintString(board BingoBoard, getAnswersInstead bool) (string) {
	var boardValue strings.Builder
	var bingoBoardLen int
	if getAnswersInstead {
		bingoBoardLen = len(board.answerLines)
	} else {
		bingoBoardLen = len(board.boardLines)
	}

	for i := 0; i < bingoBoardLen; i++ {
		if getAnswersInstead {
			boardValue.WriteString(getBingoBoardLinePrintString(board.answerLines[i]))
		} else {
			boardValue.WriteString(getBingoBoardLinePrintString(board.boardLines[i]))
		}
		if i < (bingoBoardLen - 1) {
			boardValue.WriteString("\n")
		}
	}

	return boardValue.String()
}

func getBingoBoard(lineStrings []string) (BingoBoard, error) {
	var boardLines []BingoBoardLine
	var err error

	for _, value := range lineStrings {
		currentString := strings.TrimSpace(value)
		if currentString == ""{
			continue
		}
		valueStrings := strings.Split(currentString, " ")
		currentBoardLine, newErr := getBingoBoardLine(valueStrings)
		if newErr != nil {
			if (err == nil){
				err = newErr
			} else {
				err = fmt.Errorf("Combined error: %v %v", err, newErr)
			}
		}
		boardLines = append(boardLines, currentBoardLine)
	}

	return BingoBoard{ boardLines: boardLines }, err
}

func getBingoBoardAnswers(bingoBoard BingoBoard, winningNumber int) (BingoBoard, error) {
	var err error
	newBoard := BingoBoard{ boardLines: bingoBoard.boardLines }

	bingoBoardLinesLen := len(bingoBoard.boardLines)
	if bingoBoardLinesLen == 0 {
		return newBoard, errors.New("bingoBoard.boardLines count was 0")
	}

	if len(bingoBoard.answerLines) == 0 {
		bingoBoard.answerLines = make([]BingoBoardLine, bingoBoardLinesLen)
	}

	newBoard.answerLines = []BingoBoardLine{}
	for i, boardLine := range bingoBoard.boardLines {
		newAnswer, newErr := getBingoBoardLineAnswer(boardLine, bingoBoard.answerLines[i], winningNumber)
		newBoard.answerLines = append(newBoard.answerLines, newAnswer)
		if newErr != nil {
			if (err == nil){
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
			if (err == nil){
				err = newErr
			} else {
				err = fmt.Errorf("Combined error: %v %v", err, newErr)
			}
		}
	}

	return newBoards, err
}

func sumUnmarkedNumbersBoard(bingoBoard BingoBoard) (int, error) {
	bingoBoardLinesLen := len(bingoBoard.boardLines)
	bingoBoardAnswersLen := len(bingoBoard.answerLines)
	var err error
	currentSum := 0

	if bingoBoardLinesLen <= 0 {
		return 0, fmt.Errorf("bingoBoard.boardLines length was less than or equal to 0 - %v", bingoBoardLinesLen)
	}
	if bingoBoardLinesLen != bingoBoardAnswersLen {
		return 0, fmt.Errorf("bingoBoard.boardLines length didn't equal bingoBoard.answersLines length - %v - %v", bingoBoardLinesLen, bingoBoardAnswersLen)
	}

	for i, line := range bingoBoard.boardLines {
		currentSum, newErr = sumUnmarkedNumbersBoardLine(line, bingoBoard.answerLines[i])
		if newErr != nil {
			if (err == nil){
				err = newErr
			} else {
				err = fmt.Errorf("Combined error: %v %v", err, newErr)
			}
		}
	}

	return currentSum, err
}
