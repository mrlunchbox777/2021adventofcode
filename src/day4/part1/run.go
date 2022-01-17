package day4part1

import (
	"bufio"
	d4 "github.com/mrlunchbox777/2021adventofcode/src/day4/models"
	"os"
)

func Main() (error) {
	file, err := os.Open("src/day4/data/input")
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	bingoGame, err := d4.PrepGame(scanner)
	if err != nil {
		return err
	}

	bingoGame, err = d4.CalcGame(bingoGame)
	if err != nil {
		return err
	}

	err = d4.PrintResults(bingoGame, false)
	if err != nil {
		return err
	}

	// d4.PrintWinningNumbers(bingoGame)

	return nil
}
