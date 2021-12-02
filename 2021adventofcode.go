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

func main() {
	rand.Seed(time.Now().UnixNano())
	file, err := os.Open("./data/data-day1")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	totalIncreases := int(0)
	currentDepth := 0
	firstRun := true
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		increased := false
		if firstRun {
			firstRun = false
		} else if i > currentDepth {
			increased = true
			totalIncreases++
		}
		currentDepth = i
		if increased {
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