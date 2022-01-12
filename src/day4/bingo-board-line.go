package day4

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// TODO refactor this out
type BingoBoardLine struct {
	values []int
}

func checkForBingoBoardAnswerLineWin(answerLine BingoBoardLine) (bool, error) {
	if len(answerLine.values) == 0 {
		return false, errors.New("answerLine was empty")
	}

	winningLine := true
	for _, val := range answerLine.values {
		if val != 1 {
			winningLine = false
		}
	}

	return winningLine, nil
}

func getBingoBoardLinePrintString(boardLine BingoBoardLine) (string) {
	var lineValue strings.Builder
	bingoBoardLinesLen := len(boardLine.values)

	for i, currentInt := range boardLine.values {
		nextValue := strconv.Itoa(currentInt)

		if currentInt < 10 {
			lineValue.WriteString(" ")
		}

		lineValue.WriteString(nextValue)

		if i < (bingoBoardLinesLen - 1) {
			lineValue.WriteString(" ")
		}
	}

	return lineValue.String()
}

func getBingoBoardLine(valueStrings []string) (BingoBoardLine, error) {
	var boardLine []int
	var err error

	for _, value := range valueStrings {
		currentString := strings.TrimSpace(value)
		if currentString == "" {
			continue
		}
		currentInt, newErr := strconv.Atoi(currentString)
		if newErr != nil {
			if (err == nil){
				err = newErr
			} else {
				err = fmt.Errorf("Combined error: %v %v", err, newErr)
			}
		}
		boardLine = append(boardLine, currentInt)
	}

	return BingoBoardLine{ values: boardLine }, err
}

func getBingoBoardLineAnswer(values BingoBoardLine, answers BingoBoardLine, winningNumber int) (BingoBoardLine, error) {
	areExistingAnswers := len(answers.values) > 0
	newLine := BingoBoardLine{}
	newLine.values = []int{}
	boardLength := len(values.values)

	if areExistingAnswers {
		if boardLength != len(values.values) {
			return newLine, errors.New("boardline answers.values length != boardline values.values length")
		}
	}

	for i := 0; i < boardLength; i++ {
		if areExistingAnswers {
			if answers.values[i] == 1 {
				newLine.values = append(newLine.values, 1)
				continue
			}
		}

		if values.values[i] == winningNumber {
			newLine.values = append(newLine.values, 1)
			continue
		}

		newLine.values = append(newLine.values, 0)
	}

	return newLine, nil
}

func sumUnmarkedNumbersBoardLine(boardLine BingoBoardLine, answers BingoBoardLine, getLoser bool) (int, error) {
	boardLineLen := len(boardLine.values)
	if boardLineLen == 0 {
		return 0, errors.New("boardLine length was 0")
	}
	if boardLineLen != len(answers.values) {
		return 0, errors.New("boardLine length didn't equal answers length")
	}

	newSum := 0
	for i, val := range boardLine.values {
		if getLoser {
			if answers.values[i] != 0 {
				newSum += val
			}
		} else {
			if answers.values[i] == 0 {
				newSum += val
			}
		}
	}

	return newSum, nil
}
