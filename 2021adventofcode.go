package main

import (
	"math/rand"
	"time"
	"fmt"
	"os"
	"log"
	"bufio"
	"strconv"
)

func getNextDepth(input int, a int, b int) (int, int, int) {
	total := input + a + b
	return input, a, total
}

func main() {
	rand.Seed(time.Now().UnixNano())
	file, err := os.Open("./data/data-day1")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	totalIncreases := int(0)
	var depths [2]int
	currentDepth := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		previousDepth := currentDepth
		depths[0], depths[1], currentDepth = getNextDepth(i, depths[0], depths[1])
		increased := currentDepth > previousDepth
		if increased {
			totalIncreases++
			fmt.Println(currentDepth, "(increased)")
		} else {
			fmt.Println(currentDepth, "(decreased)")
		}
	}
	fmt.Println("Increases", totalIncreases)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}