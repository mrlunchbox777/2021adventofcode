package day4

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type WinningNumbers struct {
	values []int
	latestWinningNumber int
}

func PrintWinningNumbers(bingoGame BingoGame) () {
	printWinningNumbersStruct(bingoGame.answers)
}

func printWinningNumbersStruct(winningNumbers WinningNumbers) () {
	fmt.Println("winningNumbersLen - ", len(winningNumbers.values))
	fmt.Println("winningNumbers - ", winningNumbers.values)
	fmt.Println("latestWinningNumber - ", winningNumbers.latestWinningNumber)
}

func getWinningNumbers(input string) (WinningNumbers, error) {
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
			if (err == nil){
				err = newErr
			} else {
				err = fmt.Errorf("Combined error: %v %v", err, newErr)
			}
		}
		winningNumbers = append(winningNumbers, currentInt)
	}

	return WinningNumbers{ values: winningNumbers }, err
}

func setLatestWinningNumber(winningNumbers WinningNumbers, newLatest int) (WinningNumbers, error) {
	if winningNumbers.values == nil || len(winningNumbers.values) == 0 {
		return winningNumbers, errors.New("winningNumbers was nil, or len(winningNumbers.values) == 0")
	}
	return WinningNumbers{ values: winningNumbers.values, latestWinningNumber: newLatest}, nil
}

func getLatestWinningNumber(winningNumbers WinningNumbers) (int) {
	return winningNumbers.latestWinningNumber
}
