package main

import (
	"errors"
	"fmt"
	d1p1 "github.com/mrlunchbox777/2021adventofcode/src/day1/part1"
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

func runDayPart(day int, part int) (error){
	fmt.Println("Running for day", day, "part", part)
	var err error
	switch day {
	case 1:
		err = day1(part)
	default:
		return errors.New("got an unknown day")
	}
	fmt.Println("Ran for day", day, "part", part)
	return err;
}

func day1(part int) (error) {
	switch part {
	case 1:
		d1p1.Main()
	default:
		return errors.New("got an unknown part")
	}
	return error(nil)
}

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

func getDay(regNo *regexp.Regexp, regGetNum *regexp.Regexp) (int) {
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

	return day
}

func getPart(regNo *regexp.Regexp, regGetNum *regexp.Regexp, regPartDir *regexp.Regexp, day int) (int) {
	daySubDirs, err := IOReadDir("./src/day" + strconv.Itoa(day))
	var partDirs []string
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(daySubDirs); i++ {
		currentItem := daySubDirs[i]
		if regPartDir.MatchString(currentItem) {
			partDirs = append(partDirs, currentItem)
		}
	}
	sort.Strings(partDirs)
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

	return part
}

func main() {
	regGetNum := regexp.MustCompile("[0-9]+")
	regNo := regexp.MustCompile("^[Nn][Oo]*$")
	regPartDir := regexp.MustCompile("^part[0-9]+$")
	day := getDay(regNo, regGetNum)
	part := getPart(regNo, regGetNum, regPartDir, day)

	err := runDayPart(day, part)
	if err != nil {
		panic(err)
	}
}