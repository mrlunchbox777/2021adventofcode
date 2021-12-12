package day4part1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	d4 "github.com/mrlunchbox777/2021adventofcode/src/day4"
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

func Main() {
	file, err := os.Open("src/day4/data/input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	winningNumbers, bingoBoards, err := d4.GetBingoBoards(scanner)
	if (err != nil){
		panic(err)
	}

	d4.PrintAllBingoBoards(bingoBoards)
	fmt.Println("winningNumbers - ", winningNumbers)
}
