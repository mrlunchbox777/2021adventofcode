package bingoboardline

import (
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
