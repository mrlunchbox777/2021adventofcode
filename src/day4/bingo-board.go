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

	bingoBoardLinesLen := len(bingoBoard.boardLines)
	if bingoBoardLinesLen == 0 {
		return BingoBoard{}, errors.New("bingoBoard.boardLines count was 0")
	}

	var newAnswers = []BingoBoardLine{}
	for i := 0; i < bingoBoardLinesLen; i++ {
		// TODO: this is going error because we haven't set up blank answers (should probably do that at this point)
		newAnswer, newErr := getBingoBoardLineAnswer(bingoBoard.boardLines[i], bingoBoard.answerLines[i], winningNumber)
		if newErr != nil {
			if (err == nil){
				err = newErr
			} else {
				err = fmt.Errorf("Combined error: %v %v", err, newErr)
			}
		}
		newAnswers := append(newAnswers, newAnswer)
	}

	return BingoBoard{ boardLines: bingoBoard.boardLines, answerLines: newAnswers }, err
}

func getBingoBoardsAnswersForWinningNumber(bingoBoards []BingoBoard, winningNumber int) ([]BingoBoard, error) {
	var err error
	if bingoBoards == nil {
		return nil, errors.New("bingoBoards array was nil")
	}
	bingoBoardsLen := len(bingoBoards)
	if  bingoBoardsLen == 0 {
		return nil, errors.New("bingoBoards array was empty")
	}

	newBoards := []BingoBoard
	for i := 0; i < bingoBoardLen; i++ {
		newBoard, newErr := getBingoBoardAnswers(bingoBoards[i], winningNumber)
		if newErr != nil {
			if (err == nil){
				err = newErr
			} else {
				err = fmt.Errorf("Combined error: %v %v", err, newErr)
			}
		}

		newBoards := append(newBoards, newBoard)
	}

	return newBoards, err
}
