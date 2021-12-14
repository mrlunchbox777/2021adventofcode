package day3part1

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func stringToIntArr(str string) ([]int, error) {
	chars := []rune(str)
	ints := []int{}

	for i := 0; i < len(chars); i++ {
		currentChar := string(chars[i])
		currentInt, err := strconv.Atoi(currentChar)
		if err != nil {
			return nil, err
		}
		ints = append(ints, currentInt)
	}

	return ints, nil
}

func calcGammaEpsilon(lineCount int, diagByColumn map[int]int) (map[int]int, map[int]int) {
	halfLineCount := lineCount / 2
	gamma := make(map[int]int)
	epsilon := make(map[int]int)

	for j, i := range diagByColumn {
		fmt.Println("diag by column i -", i)
		if i < halfLineCount {
			gamma[j] = 0
			epsilon[j] = 1
		} else {
			gamma[j] = 1
			epsilon[j] = 0
		}
	}

	return gamma, epsilon
}

func reverse(numbers map[int]int) map[int]int {
	for i := 0; i < len(numbers)/2; i++ {
		j := len(numbers) - i - 1
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	return numbers
}

func getDecimalFromBinary(numbers map[int]int) float64 {
	retVal := float64(0)
	reverseNumbers := reverse(numbers)
	for i, j := range reverseNumbers {
		retVal += float64(j) * math.Pow(float64(2), float64(i))
	}
	return retVal
}

func Main() (error) {
	file, err := os.Open("src/day3/data/input")
	if err != nil {
		return err
	}
	defer file.Close()

	diagByColumn := make(map[int]int)  
	scanner := bufio.NewScanner(file)
	lineCount := 0

	for scanner.Scan() {
		lineCount++
		i := scanner.Text()
		ints, err := stringToIntArr(i)
		if err != nil {
			return err
		}
		for j := 0; j < len(ints); j ++ {
			currentInt := ints[j]
			diagByColumn[j] += currentInt  
		}
	}

	gammaIntArr, epsilonIntArr := calcGammaEpsilon(lineCount, diagByColumn)
	fmt.Println("gammaIntArr -", gammaIntArr, ", epsilonIntArr", epsilonIntArr)
	gamma := getDecimalFromBinary(gammaIntArr)
	epsilon := getDecimalFromBinary(epsilonIntArr)
	product := gamma * epsilon
	fmt.Println("gamma -", gamma, "epsilon -", epsilon, "product - ", product)

	return nil
}
