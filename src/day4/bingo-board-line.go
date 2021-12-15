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

func getBingoBoardLine(valueStrings []string) (BingoBoardLine, error) {
	var boardLine []int
	var err error

	for i := 0; i < len(valueStrings); i++ {
		currentString := strings.TrimSpace(valueStrings[i])
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
	var areExistingAnswers bool
	if values == nil {
		return nil, errors.New("boardline values poco was nil")
	}
	if values.values == nil {
		return nil, errors.New("boardline values.values poco was nil")
	}
	if answers == nil {
		areExistingAnswers = false
	} else {
		areExistingAnswers := len(answers.values) > 0
	}

	var newAnswers []int
	boardLength := len(values.values)

	if areExistingAnswers {
		if boardLength != len(values.values) {
			return nil, errors.New("boardline answers.values length != boardline values.values length")
		}
	}

	for i := 0; i < boardLength; i++ {
		if areExistingAnswers {
			if answers.values[i] == 1 {
				append(newAnswers, 1)
				continue
			}
		}

		if values.values[i] == winningNumber {
			append(newAnswers, 1)
			continue
		}

		append(newAnswers, 0)
	}

	return newAnswers
}
