package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func day14() {

	readFile, err := os.Open("day14.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	solids := make(map[struct {x int; y int}]bool)
	bottom := 0

	for fileScanner.Scan() {
		line := strings.TrimSpace(fileScanner.Text())

		if len(line) == 0 {
			continue
		}

		var sections []struct {x int; y int}

		for _, coords := range strings.Split(line, " -> ") {
			parts := strings.Split(coords, ",")
			x, _ := strconv.Atoi(parts[0])
			y, _ := strconv.Atoi(parts[1])
			sections = append(sections, struct {x int; y int}{x: x, y: y})
		}

		first := sections[0]
		for i := 1; i < len(sections); i++ {
			second := sections[i]
			if first.x == second.x {
				// go in y direction
				for n := Min(first.y, second.y); n <= Max(first.y, second.y); n++ {
					solids[struct {
						x int
						y int
					}{x: first.x, y: n}] = true
				}
			} else {
				for n := Min(first.x, second.x); n <= Max(first.x, second.x); n++ {
					solids[struct {
						x int
						y int
					}{x: n, y: first.y}] = true
				}
			}
			if first.y > bottom {
				bottom = first.y
			} else if second.y > bottom {
				bottom = second.y
			}
			first = second
		}
	}
	err = readFile.Close()
	if err != nil {
		fmt.Println(err)
	}
	unitCount := 0
	stop := false
	part1 := true
	for !stop {
		x := 500
		placed := false
		for y := 1; y < bottom+2; y++ {
			if _, exists := solids[struct {
				x int
				y int
			}{x: x, y: y}]; !exists {
				// free
				continue
			}
			if _, exists := solids[struct {
				x int
				y int
			}{x: x-1, y: y}]; !exists {
				// free
				x--
				continue
			}
			if _, exists := solids[struct {
				x int
				y int
			}{x: x+1, y: y}]; !exists {
				// free
				x++
				continue
			}
			placed = true
			if y == 1 {
				stop = true
			}
			solids[struct {
				x int
				y int
			}{x: x, y: y-1}] = true
			break
		}
		if ! placed {
			if part1 {
				fmt.Println(unitCount)
				part1 = false
			}
			solids[struct {
				x int
				y int
			}{x: x, y: bottom+1}] = true
		}
		unitCount++
	}
	fmt.Println(unitCount)
}

func Min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
func Max(a int, b int) int {
	if a < b {
		return b
	}
	return a
}