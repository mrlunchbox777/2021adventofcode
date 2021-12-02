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

func getNextDepth(a int, b int, input int) (int, int, int) {
	total := a + b + input
	return b, input, total
}

func main() {
	rand.Seed(time.Now().UnixNano())
	file, err := os.Open("./data/data-day1")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	totalIncreases := 0
	var depths [2]int
	currentDepth := 0
	had3answers := false
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		fmt.Println("before values: ", "a -", depths[0], "b -", depths[1], "i -", i)
		previousDepth := currentDepth
		depths[0], depths[1], currentDepth = getNextDepth(depths[0], depths[1], i)
		if ( ! had3answers ) && ( depths[0] > 0 && ( depths[1] > 0 && i > 0 )) {
			had3answers = true
		} else {
			if currentDepth > previousDepth {
				totalIncreases++
				fmt.Println(currentDepth, "(increased)")
			} else {
				fmt.Println(currentDepth, "(decreased)")
			}
		}
		fmt.Println("after values: ", "a -", depths[0], "b -", depths[1])
	}
	fmt.Println("Increases", totalIncreases)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}