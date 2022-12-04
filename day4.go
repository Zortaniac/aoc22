package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func day4() {

	readFile, err := os.Open("day4.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	matchFullCount := 0
	matchCount := 0
	for fileScanner.Scan() {
		l := strings.TrimSpace(fileScanner.Text())
		var x1,x2,y1,y2 int
		_, err = fmt.Sscanf(l, "%d-%d,%d-%d", &x1, &x2, &y1, &y2)
		if err != nil {
			fmt.Println(err)
		}

		if overlapFull(x1, x2, y1, y2) {
			matchFullCount += 1
			matchCount += 1
		} else if overlap(x1, x2, y1, y2) {
			matchCount += 1
		}
	}
	err = readFile.Close()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(matchFullCount)
	fmt.Println(matchCount)
}

func overlap(x1 int, x2 int, y1 int, y2 int) bool {
	if x2 < y1 || y2 < x1 {
		return false
	}
	return true
}

func overlapFull(x1 int, x2 int, y1 int, y2 int) bool {
	if y1 <= x1 && x2 <= y2 {
		return true
	}
	if x1 <= y1 && y2 <= x2 {
		return true
	}

	return false
}
