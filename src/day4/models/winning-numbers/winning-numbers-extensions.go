package winningnumbers

import (
	"errors"
	"fmt"
)

func (winningNumbers WinningNumbers) printWinningNumbersStruct() {
	fmt.Println("winningNumbersLen - ", len(winningNumbers.values))
	fmt.Println("winningNumbers - ", winningNumbers.values)
	fmt.Println("latestWinningNumber - ", winningNumbers.latestWinningNumber)
	fmt.Println("latestLosingNumber - ", winningNumbers.latestLosingNumber)
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
