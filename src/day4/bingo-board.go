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

func getBingoBoard(lineStrings []string) (BingoBoard, error) {
	var boardLines []BingoBoardLine
	var err error

	for i := 0; i < len(lineStrings); i++ {
		currentString := strings.TrimSpace(lineStrings[i])
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
	for i := 0; i < bingoBoardLinesLen; i++ {
		newAnswer, newErr := getBingoBoardLineAnswer(bingoBoard.boardLines[i], bingoBoard.answerLines[i], winningNumber)
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
	bingoBoardsLen := len(bingoBoards)
	newBoards := []BingoBoard{}

	if bingoBoardsLen == 0 {
		return newBoards, errors.New("bingoBoards count was 0")
	}

	for i := 0; i < bingoBoardsLen; i++ {
		newBoard, newErr := getBingoBoardAnswers(bingoBoards[i], winningNumber)
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
