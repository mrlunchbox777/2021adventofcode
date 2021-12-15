package day4

import (
	"fmt"
	"strconv"
	"strings"
)

func getWinningNumbers(input string) ([]int, error) {
	var winningNumbers []int
	winningNumbersStringArr := strings.Split(input, ",")
	err := error(nil)

	for i := 0; i < len(winningNumbersStringArr); i++ {
		currentInt, newErr := strconv.Atoi(winningNumbersStringArr[i])
		if newErr != nil {
			if (err == nil){
				err = newErr
			} else {
				err = fmt.Errorf("Combined error: %v %v", err, newErr)
			}
		}
		winningNumbers = append(winningNumbers, currentInt)
	}

	return winningNumbers, err
}
