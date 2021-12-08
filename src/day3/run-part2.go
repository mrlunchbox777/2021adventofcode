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

func reduceDiagnosticMatrix(matrix map[int]map[int]int, keepMostCommon bool, targetColumn int) map[int]map[int]int {
	rowCount := len(matrix)
	fmt.Println("rowCount - ", rowCount)
	if rowCount == 1 {
		return matrix
	}

	halfRowCount := rowCount / 2
	newMatrix := make(map[int]map[int]int)
	commonalityCounter := 0
	valueToKeep := 0

	for i := 0; i < rowCount; i++ {
		commonalityCounter += matrix[i][targetColumn]
	}

	if commonalityCounter >= halfRowCount {
		if keepMostCommon {
			valueToKeep = 1
		} else {
			valueToKeep = 0
		}
	} else {
		if keepMostCommon {
			valueToKeep = 0
		} else {
			valueToKeep = 1
		}
	}
	fmt.Println("value to keep -", valueToKeep)

	for i := 0; i < rowCount; i++ {
		if matrix[i][targetColumn] == valueToKeep {
			newMatrixRow := len(newMatrix)
			newMatrix[newMatrixRow] = make(map[int]int)
			for j := 0; j < len(matrix[i]); j++ {
				newMatrix[newMatrixRow][j] = matrix[i][j]
			}
		}
	}
	fmt.Println("new matrix -", newMatrix)

	return reduceDiagnosticMatrix(matrix, keepMostCommon, (targetColumn + 1))
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

	o2 := reduceDiagnosticMatrix(matrix, true, 0)
	cO2 := reduceDiagnosticMatrix(matrix, false, 0)

	fmt.Println("o2")
	fmt.Println(o2)
	fmt.Println("cO2")
	fmt.Println(cO2)
}
