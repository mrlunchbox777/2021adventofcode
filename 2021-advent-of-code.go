package main

import (
	"fmt"
	// d1p1 "github.com/mrlunchbox777/2021adventofcode/src/day1/part1"
	// d2p1 "github.com/mrlunchbox777/2021adventofcode/src/day2/part1"
	// d2p2 "github.com/mrlunchbox777/2021adventofcode/src/day2/part2"
	// d3p1 "github.com/mrlunchbox777/2021adventofcode/src/day3/part1"
	// d3p2 "github.com/mrlunchbox777/2021adventofcode/src/day3/part2"
	// d4p1 "github.com/mrlunchbox777/2021adventofcode/src/day4/part1"
	"io/ioutil"
	// "os"
	"regexp"
	"sort"
	"strconv"
)

func IOReadDir(root string) ([]string, error) {
	var files []string
	fileInfo, err := ioutil.ReadDir(root)
	if err != nil {
			return files, err
	}

	for _, file := range fileInfo {
			files = append(files, file.Name())
	}
	return files, nil
}

func main() {
	regGetNum := regexp.MustCompile("[0-9]+")
	regNo := regexp.MustCompile("^[Nn][Oo]*$")
	regPartDir := regexp.MustCompile("^part[0-9]+$")

	dayDirs, err := IOReadDir("./src")
	if err != nil {
		panic(err)
	}
	sort.Strings(dayDirs)
	day, err := strconv.Atoi(regGetNum.FindString(dayDirs[len(dayDirs) - 1]))
	if err != nil {
		panic(err)
	}
	regDay := regexp.MustCompile("^[0-" + strconv.Itoa(day) + "]+$")

	var lastestDayString string
	fmt.Println("Use Latest Day?")
	fmt.Scanln(&lastestDayString)
	if regNo.MatchString(lastestDayString) {
		var dayString string
		fmt.Println("Which Day?")
		fmt.Scanln(&dayString)
		fmt.Println("dayString -" + dayString + "-")
		if regDay.MatchString(dayString) {
			day, err = strconv.Atoi(dayString)
			if day == 0 {
				panic("Error: Bad Day Selection")
			}
			if err != nil {
				panic(err)
			}
		} else {
			panic("Error: Bad Day Selection")
		}
	}

	daySubDirs, err := IOReadDir("./src/day" + strconv.Itoa(day))
	var partDirs []string
	if err != nil {
		panic(err)
	}
	fmt.Println(daySubDirs)
	for i := 0; i < len(daySubDirs); i++ {
		currentItem := daySubDirs[i]
		if regPartDir.MatchString(currentItem) {
			partDirs = append(partDirs, currentItem)
		}
	}
	sort.Strings(partDirs)
	fmt.Println(partDirs[len(partDirs) - 1])
	fmt.Println(regGetNum.FindString(partDirs[len(partDirs) - 1]))
	part, err := strconv.Atoi(regGetNum.FindString(partDirs[len(partDirs) - 1]))
	if err != nil {
		panic(err)
	}
	regPart := regexp.MustCompile("^[0-" + strconv.Itoa(part) + "]+$")

	var latestPartString string
	var partString string
	fmt.Println("Use Latest Part?")
	fmt.Scanln(&latestPartString)
	if regNo.MatchString(latestPartString) {
		fmt.Println("Which Part?")
		fmt.Scanln(&partString)
		if regPart.MatchString(partString) {
			part, err = strconv.Atoi(partString)
			if part == 0 {
				panic("Error: Bad Part Selection")
			}
			if err != nil {
				panic(err)
			}
		} else {
			panic("Error: Bad Part Selection")
		}
	}

	fmt.Println("Running for day", day, "part", partString)
	// d4p1.Main()
	fmt.Println("Ran for day", day)
}