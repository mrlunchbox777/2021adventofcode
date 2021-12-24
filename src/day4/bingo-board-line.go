package day4

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type BingoBoardLine struct {
	values []int
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
