package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("../data/input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	x := 0
	y := 0
	aim := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i := scanner.Text()
		i_array := strings.Split(i, " ")
		direction := i_array[0]
		value, _ := strconv.Atoi(i_array[1])
		fmt.Println("direction", direction, "- value", value)
		switch direction {
			case "up":
				aim -= value
			case "down":
				aim += value
			case "forward":
				x += value
				y += (aim * value)
		}
		fmt.Println("x", x, "- y", y, "- aim", aim)
	}
	final_product := x * y
	fmt.Println("final_product", final_product)
}
