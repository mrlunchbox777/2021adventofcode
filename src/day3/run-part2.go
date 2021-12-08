package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"math"
)

const MaxUint = ^uint(0) 
const MaxInt = int(MaxUint >> 1) 

func stringToIntArr(str string) ([]int) {
	chars := []rune(str)
	ints := []int{}

	for i := 0; i < len(chars); i++ {
		currentChar := string(chars[i])
		currentInt, err := strconv.Atoi(currentChar)
		if err != nil {
			panic(err)
		}
		ints = append(ints, currentInt)
	}

	return ints
}

func findMostCommonNumber(numbers map[int]int, keepMostCommon bool) int {
	numberLen := len(numbers)
	commonalityCounter := make(map[int]int)
	largestCount := 0
	mostCommonNumber := 0

	for i := 0; i < numberLen; i++ {
		commonalityCounter[numbers[i]]++
	}

	if keepMostCommon {
		largestCount = 0
	} else {
		largestCount = MaxInt
	}
	for i := 1; i < len(commonalityCounter); i++ {
		current := commonalityCounter[i]
		if (keepMostCommon) && (current > largestCount) {
			mostCommonNumber = current
		} else if (!keepMostCommon) && (current < largestCount) {
			mostCommonNumber = current
		} else if current == largestCount {
			if keepMostCommon {
				mostCommonNumber = 1
			} else {
				mostCommonNumber = 0
			}
		}
	}
	
	return mostCommonNumber
}

func invertMatrix(matrix map[int]map[int]int) map[int]map[int]int {
	if len(matrix) == 0 {
		panic("no values in the matrix")
	}

	columnLength := len(matrix)
	lineLength := len(matrix[0])
	invertedMatrix := make(map[int]map[int]int)
	for i := 0; i < lineLength; i++ {
		invertedMatrix[i] = make(map[int]int)
		for j := 0; j < columnLength; j ++ {
			invertedMatrix[i][j] = matrix[j][i]  
		}
	}

	return invertedMatrix
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

func main() {
	file, err := os.Open("./data/input test")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	matrix := make(map[int]map[int]int)
	scanner := bufio.NewScanner(file)
	lineCount := 0

	for scanner.Scan() {
		i := scanner.Text()
		ints := stringToIntArr(i)
		matrix[lineCount] = make(map[int]int)
		for j := 0; j < len(ints); j ++ {
			currentInt := ints[j]
			matrix[lineCount][j] = currentInt  
		}
		lineCount++
	}

	if len(matrix) == 0 {
		panic("no values in the matrix")
	}

	invertedMatrix := invertMatrix(matrix)

	o2 := make(map[int]int)
	cO2 := make(map[int]int)
	for i := 0; i < len(invertedMatrix); i++ {
		o2[i] = findMostCommonNumber(invertedMatrix[i], true)
		cO2[i] = findMostCommonNumber(invertedMatrix[i], false)
	}

	fmt.Println("o2")
	fmt.Println(o2)
	fmt.Println("cO2")
	fmt.Println(cO2)

	// gammaIntArr, epsilonIntArr := calcGammaEpsilon(lineCount, diagByColumn)
	// fmt.Println("gammaIntArr -", gammaIntArr, ", epsilonIntArr", epsilonIntArr)
	// gamma := getDecimalFromBinary(gammaIntArr)
	// epsilon := getDecimalFromBinary(epsilonIntArr)
	// product := gamma * epsilon
	// fmt.Println("gamma -", gamma, "epsilon -", epsilon, "product - ", product)
}
