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
	fmt.Println("winningNumbers - ", bingoGame.answers.values)
}

func getWinningNumbers(input string) (WinningNumbers, error) {
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

	return WinningNumbers{ values: winningNumbers }, err
}
