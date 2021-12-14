package day3part2

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
)

const MaxUint = ^uint(0) 
const MaxInt = int(MaxUint >> 1) 

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

func reduceDiagnosticMatrix(matrix map[int]map[int]int, keepMostCommon bool, targetColumn int) map[int]map[int]int {
	rowCount := len(matrix)
	if rowCount == 1 {
		return matrix
	}

	halfRowCount := int(math.Round(float64(rowCount) / 2.0))
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

	for i := 0; i < rowCount; i++ {
		if matrix[i][targetColumn] == valueToKeep {
			newMatrixRow := len(newMatrix)
			newMatrix[newMatrixRow] = make(map[int]int)
			for j := 0; j < len(matrix[i]); j++ {
				newMatrix[newMatrixRow][j] = matrix[i][j]
			}
		}
	}

	return reduceDiagnosticMatrix(newMatrix, keepMostCommon, (targetColumn + 1))
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

func Main() (error){
	file, err := os.Open("src/day3/data/input")
	if err != nil {
		return err
	}
	defer file.Close()

	matrix := make(map[int]map[int]int)
	scanner := bufio.NewScanner(file)
	lineCount := 0

	for scanner.Scan() {
		i := scanner.Text()
		ints, err := stringToIntArr(i)
		if err != nil {
			return err
		}
		matrix[lineCount] = make(map[int]int)
		for j := 0; j < len(ints); j ++ {
			currentInt := ints[j]
			matrix[lineCount][j] = currentInt
		}
		lineCount++
	}

	if len(matrix) == 0 {
		return errors.New("no values in the matrix")
	}

	o2Reduction := reduceDiagnosticMatrix(matrix, true, 0)
	cO2Reduction := reduceDiagnosticMatrix(matrix, false, 0)
	o2 := getDecimalFromBinary(o2Reduction[0])
	cO2 := getDecimalFromBinary(cO2Reduction[0])
	finalValue := o2 * cO2
	fmt.Printf("finalValue - %f\n", finalValue)
	return nil
}
