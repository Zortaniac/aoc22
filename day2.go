package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func day2a() {

	readFile, err := os.Open("day2.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	score := 0
	for fileScanner.Scan() {
		l := strings.Fields(fileScanner.Text())
		if len(l) == 0 {
			continue
		}
		op := l[0]
		me := l[1]

		win := (me == "X" && op == "C") || (me == "Z" && op == "B") || (me == "Y" && op == "A")

		me_t := "A"
		if me == "Y" {
			me_t = "B"
		} else if me == "Z" {
			me_t = "C"
		}

		if me_t == op {
			// draw
			score += 3
		} else if win {
			score += 6
		}

		switch me {
		case "A":
		case "X":
			score += 1
			break
		case "B":
		case "Y":
			score += 2
			break
		case "C":
		case "Z":
			score += 3
			break
		}
	}
	err = readFile.Close()
	if err != nil {
		fmt.Println(err)
	}

	// part one
	fmt.Println(score)

}

func getScore(t string) int {
	switch t {
	case "A":
		return 1
	case "B":
		return 2
	case "C":
		return 3
	}
	return 0
}

func lose(t string) int {
	switch t {
	case "A": // paper
		return 3 // scissors
	case "B":
		return 1 // rock
	case "C":
		return 2 // paper
	}
	return 0
}

func win(t string) int {
	switch t {
	case "A": // rock
		return 2 // paper
	case "B": // paper
		return 3 // scissors
	case "C": // scissors
		return 1 // rock
	}
	return 0
}

func day2b() {

	readFile, err := os.Open("day2.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	score := 0
	for fileScanner.Scan() {
		l := strings.Fields(fileScanner.Text())
		if len(l) == 0 {
			continue
		}
		op := l[0]

		switch l[1] {
		case "X":
			// lose
			score += lose(op)
			break
		case "Y":
			// draw
			score += getScore(op)
			score += 3
		case "Z":
			// win
			score += win(op)
			score += 6
		}
	}
	err = readFile.Close()
	if err != nil {
		fmt.Println(err)
	}

	// part one
	fmt.Println(score)

}
