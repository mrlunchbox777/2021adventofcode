package bingoboardline

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type BingoBoardLine struct {
	isAnswerLine bool
	maxSize      int
	values       []int
}

func (boardLine BingoBoardLine) IsAnswerLine() bool {
	return boardLine.isAnswerLine
}

func (boardLine BingoBoardLine) MaxSize() int {
	return boardLine.maxSize
}

func (boardLine BingoBoardLine) Values() []int {
	return boardLine.values
}

func (boardLine BingoBoardLine) SetValues(values []int) {
	boardLine.values = values
}

//////////////////////////////////////////////////
// Original Extensions
//////////////////////////////////////////////////

func (answerLine BingoBoardLine) CheckForBingoBoardAnswerLineWin() (bool, error) {
	if !answerLine.isAnswerLine {
		return false, errors.New("This is not an answer line")
	}
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

func (boardLine BingoBoardLine) GetBingoBoardLinePrintString() string {
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

func GetBingoBoardLine(valueStrings []string) (BingoBoardLine, error) {
	var boardLine []int
	var err error

	for _, value := range valueStrings {
		currentString := strings.TrimSpace(value)
		if currentString == "" {
			continue
		}
		currentInt, newErr := strconv.Atoi(currentString)
		if newErr != nil {
			if err == nil {
				err = newErr
			} else {
				err = fmt.Errorf("Combined error: %v %v", err, newErr)
			}
		}
		boardLine = append(boardLine, currentInt)
	}

	return BingoBoardLine{values: boardLine, maxSize: len(boardLine), isAnswerLine: false}, err
}

func (boardLine BingoBoardLine) GetBingoBoardLineAnswer(answers BingoBoardLine, winningNumber int) (BingoBoardLine, error) {
	areExistingAnswers := len(answers.values) > 0
	newLine := BingoBoardLine{}
	newLine.values = []int{}

	if areExistingAnswers {
		if boardLine.MaxSize() != answers.MaxSize() {
			return newLine, errors.New("boardline answers.values length != boardline values.values length")
		}
	}

	for i := 0; i < boardLine.MaxSize(); i++ {
		if areExistingAnswers {
			if answers.values[i] == 1 {
				newLine.values = append(newLine.values, 1)
				continue
			}
		}

		if boardLine.values[i] == winningNumber {
			newLine.values = append(newLine.values, 1)
			continue
		}

		newLine.values = append(newLine.values, 0)
	}

	return newLine, nil
}

func (boardLine BingoBoardLine) SumUnmarkedNumbersBoardLine(answers BingoBoardLine, getLoser bool) (int, error) {
	if boardLine.MaxSize() == 0 {
		return 0, errors.New("boardLine length was 0")
	}
	if boardLine.MaxSize() != len(answers.values) {
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
