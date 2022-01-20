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

//////////////////////////////////////////////////
// Original Extensions
//////////////////////////////////////////////////

func (winningNumbers WinningNumbers) printWinningNumbersStruct() {
	fmt.Println("winningNumbersLen - ", len(winningNumbers.values))
	fmt.Println("winningNumbers - ", winningNumbers.values)
	fmt.Println("latestWinningNumber - ", winningNumbers.latestWinningNumber)
	fmt.Println("latestLosingNumber - ", winningNumbers.latestLosingNumber)
}

func GetWinningNumbers(input string) (WinningNumbers, error) {
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

func (winningNumbers WinningNumbers) SetLatestNumber(newLatest int, setLoser bool) (WinningNumbers, error) {
	if winningNumbers.values == nil || len(winningNumbers.values) == 0 {
		return winningNumbers, errors.New("winningNumbers was nil, or len(winningNumbers.values) == 0")
	}
	if setLoser {
		return WinningNumbers{values: winningNumbers.values, latestWinningNumber: winningNumbers.latestWinningNumber, latestLosingNumber: newLatest}, nil
	} else {
		return WinningNumbers{values: winningNumbers.values, latestWinningNumber: newLatest, latestLosingNumber: winningNumbers.latestLosingNumber}, nil
	}
}

func (winningNumbers WinningNumbers) GetLatestNumber(getLoser bool) int {
	if getLoser {
		return winningNumbers.latestLosingNumber
	} else {
		return winningNumbers.latestWinningNumber
	}
}
