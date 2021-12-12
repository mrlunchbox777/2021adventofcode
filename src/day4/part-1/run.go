package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func stringToIntArr(str string) ([]int, error) {
	chars := []rune(str)
	ints := []int{}
	var err error

	for i := 0; i < len(chars); i++ {
		currentChar := string(chars[i])
		currentInt, newErr := strconv.Atoi(currentChar)
		if newErr != nil {
			if (err != nil){
				err = newErr
			} else {
				err = fmt.Errorf("Combined error: %v %v", err, newErr)
			}
		}
		ints = append(ints, currentInt)
	}

	return ints, err
}

func getWinningNumbers(input string) ([]int, error) {
	var winningNumbers []int
	winningNumbersStringArr := strings.Split(input, ",")
	err := error(nil)

	for i := 0; i < len(winningNumbersStringArr); i++ {
		currentInt, newErr := strconv.Atoi(winningNumbersStringArr[i])
		if newErr != nil {
			if (err != nil){
				err = newErr
			} else {
				err = fmt.Errorf("Combined error: %v %v", err, newErr)
			}
		}
		winningNumbers = append(winningNumbers, currentInt)
	}

	return winningNumbers, err
}

func main() {
	file, err := os.Open("../data/input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	winningNumbers, bingoBoards, err := getBingoBoards(scanner)
	if (err != nil){
		panic(err)
	}

	printAllBingoBoards(bingoBoards)
	fmt.Println("winningNumbers - ", winningNumbers)
}
