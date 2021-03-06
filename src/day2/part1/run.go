package day2part1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Main() (error) {
	file, err := os.Open("src/day2/data/input")
	if err != nil {
		return err
	}
	defer file.Close()

	x := 0
	y := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i := scanner.Text()
		i_array := strings.Split(i, " ")
		direction := i_array[0]
		value, _ := strconv.Atoi(i_array[1])
		fmt.Println("direction", direction, "- value", value)
		switch direction {
			case "up":
				y -= value
			case "down":
				y += value
			case "forward":
				x += value
		}
		fmt.Println("x", x, "- y", y)
	}
	final_product := x * y
	fmt.Println("final_product", final_product)

	return nil
}
