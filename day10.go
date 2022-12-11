package main

import (
	"bufio"
	"fmt"
	"os"
)

func day10() {

	readFile, err := os.Open("day10.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	x := 1
	strength := 0
	cycle := 21
	var image [6][40]bool

	for fileScanner.Scan() {
		line := fileScanner.Text()

		if len(line) == 0 {
			break
		}

		row := (cycle-21 % 240)/40
		pos := (cycle-21) % 40
		if line == "noop" {
			if cycle % 40 == 0 {
				strength += (cycle-20)*x
			}
			if Abs(x - pos) <= 1 {
				image[row][pos] = true
			}
			cycle++
		} else {
			if (cycle+1) % 40 <= 1 {
				strength += (cycle-20+(cycle%40%38))*x
			}
			if Abs(x - pos) <= 1 {
				image[row][pos] = true
			}
			row = (cycle-21+1 % 240)/40
			pos = (cycle-21+1) % 40
			if Abs(x - pos) <= 1 {
				image[row][pos] = true
			}
			var amount int
			_, err = fmt.Sscanf(line, "addx %d", &amount)
			x += amount
			cycle += 2
		}
	}
	fmt.Println(strength)

	for _, r := range image {
		for _, p := range  r {
			if p {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	err = readFile.Close()
	if err != nil {
		fmt.Println(err)
	}
}
