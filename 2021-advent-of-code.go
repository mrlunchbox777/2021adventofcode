package main

import (
	"fmt"
	// d1p1 "github.com/mrlunchbox777/2021adventofcode/src/day1/part1"
	// d2p1 "github.com/mrlunchbox777/2021adventofcode/src/day2/part1"
	// d2p2 "github.com/mrlunchbox777/2021adventofcode/src/day2/part2"
	// d3p1 "github.com/mrlunchbox777/2021adventofcode/src/day3/part1"
	// d3p2 "github.com/mrlunchbox777/2021adventofcode/src/day3/part2"
	d4p1 "github.com/mrlunchbox777/2021adventofcode/src/day4/part1"
	"io/ioutil"
	// "os"
	"regexp"
	"sort"
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
	regGetNum := regexp.MustCompile(`[0-9]+`)
	regNo := regexp.MustCompile(`^[Nn][Oo]*$`)
	// regNum := regexp.MustCompile(`^[0-9]+$`)

	var dayString string
	fmt.Println("Use Latest Day?")
	fmt.Scanln(&dayString)
	if regNo.MatchString(dayString) {
		// do check for days
		// day := jkfdls;
	} else {
		dayDirs, err := IOReadDir("./src")
		if err != nil {
			panic(err)
		}
		sort.Strings(dayDirs)
		day := regGetNum.FindString(dayDirs[len(dayDirs) - 1])
		fmt.Println("day - " + day)
	}

	var partString string
	fmt.Println("Use Latest Part?")
	fmt.Scanln(&partString)


	// fmt.Println("Running For " + )
	d4p1.Main()
}