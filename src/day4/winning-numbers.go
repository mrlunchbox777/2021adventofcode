package day4

import (
	"fmt"
	"strconv"
	"strings"
)

type WinningNumbers struct {
	values []int
}

func PrintWinningNumbers(bingoGame BingoGame) () {
	printWinningNumbersStruct(bingoGame.answers)
}

func printWinningNumbersStruct(winningNumbers WinningNumbers) () {
	fmt.Println("winningNumbersLen - ", len(winningNumbers.values))
	fmt.Println("winningNumbers - ", winningNumbers.values)
}

func getWinningNumbers(input string) (WinningNumbers, error) {
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
