package winningnumbers

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type WinningNumbers struct {
	values              []int
	latestWinningNumber int
	latestLosingNumber  int
}

func (wNumbers WinningNumbers) Values() []int {
	return wNumbers.values
}

func (wNumbers WinningNumbers) LatestWinningNumber() int {
	return wNumbers.latestWinningNumber
}

func (wNumbers WinningNumbers) LatestLosingNumber() int {
	return wNumbers.latestLosingNumber
}

func CreateWinningNumbers(input string) (WinningNumbers, error) {
	if len(input) == 0 {
		var emptyWinners WinningNumbers
		return emptyWinners, errors.New("input for getWinningNumbers was nil or empty")
	}

	var winningNumbers []int
	winningNumbersStringArr := strings.Split(input, ",")
	err := error(nil)

	for _, currentIntStr := range winningNumbersStringArr {
		currentInt, newErr := strconv.Atoi(currentIntStr)
		if newErr != nil {
			if err == nil {
				err = newErr
			} else {
				err = fmt.Errorf("Combined error: %v %v", err, newErr)
			}
		}
		winningNumbers = append(winningNumbers, currentInt)
	}

	return WinningNumbers{values: winningNumbers}, err
}
