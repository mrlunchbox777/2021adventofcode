package day4

import (
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
		if currentString == ""{
			continue
		}
		currentInt, newErr := strconv.Atoi(currentString)
		if newErr != nil {
			if (err != nil){
				err = newErr
			} else {
				err = fmt.Errorf("Combined error: %v %v", err, newErr)
			}
		}
		boardLine = append(boardLine, currentInt)
	}

	return BingoBoardLine{ values: boardLine }, err
}
