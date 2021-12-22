package day4part1

import (
	"bufio"
	d4 "github.com/mrlunchbox777/2021adventofcode/src/day4"
	"fmt"
	"os"
	"strconv"
)

func stringToIntArr(str string) ([]int, error) {
	chars := []rune(str)
	ints := []int{}
	var err error

	for i := 0; i < len(chars); i++ {
		currentChar := string(chars[i])
		currentInt, newErr := strconv.Atoi(currentChar)
		if newErr != nil {
			if (err == nil){
				err = newErr
			} else {
				err = fmt.Errorf("Combined error: %v %v", err, newErr)
			}
		}
		ints = append(ints, currentInt)
	}

	return ints, err
}

func Main() (error) {
	file, err := os.Open("src/day4/data/input")
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	bingoGame, err := d4.PrepGame(scanner)
	if (err != nil){
		return err
	}
	// d4.PrintBingoBoards(bingoGame)
	// d4.PrintWinningNumbers(bingoGame)

	bingoGame, err = d4.CalcGame(bingoGame)
	if (err != nil){
		return err
	}

	d4.PrintBingoBoardsAnswers(bingoGame)
	d4.PrintWinningNumbers(bingoGame)
	return nil
}
