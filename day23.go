package main

import (
	"bufio"
	"fmt"
	"os"
)

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

func day23() {

	readFile, err := os.Open("day23.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	elves := make(map[pos]struct{})

	y := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()

		if len(line) == 0 {
			break
		}
		for x := 0; x < len(line); x++ {
			if line[x] == '#' {
				elves[pos{x: x, y: y}] = struct{}{}
			}
		}
		y++
	}
	err = readFile.Close()
	if err != nil {
		fmt.Println(err)
	}
	currentDirection := 1
	r := 0
	for true {
		proposedMoves := make(map[pos]int, len(elves))
		moveMap := make(map[pos]pos, len(elves))
		for e := range elves {
			target := process(e, elves, currentDirection)
			proposedMoves[target]++
			moveMap[e] = target
		}
		currentDirection = (currentDirection % 4) + 1
		elves = make(map[pos]struct{}, len(elves))
		moved := 0
		for e, t := range moveMap {
			if e == t {
				elves[e] = struct{}{}
				continue
			}
			if proposedMoves[t] > 1 {
				elves[e] = struct{}{}
			} else {
				elves[t] = struct{}{}
				moved++
			}
		}
		if moved == 0 {
			fmt.Println(r + 1)
			break
		}
		if r == 9 {
			xMin, yMin := MaxInt, MaxInt
			xMax, yMax := MinInt, MinInt

			for e := range elves {
				if e.x < xMin {
					xMin = e.x
				}
				if e.x > xMax {
					xMax = e.x
				}
				if e.y < yMin {
					yMin = e.y
				}
				if e.y > yMax {
					yMax = e.y
				}
			}
			fmt.Println((Abs(xMin-xMax)+1)*(Abs(yMin-yMax)+1) - len(elves))
		}
		r++
	}
}

func process(elf pos, elves map[pos]struct{}, currentDirection int) pos {
	var posMap [9]bool
	count := -1
	for y := 0; y < 3; y++ {
		for x := 0; x < 3; x++ {

			if _, e := elves[pos{x: elf.x - 1 + x, y: elf.y - 1 + y}]; e {
				posMap[y*3+x] = true
				count++
			}
		}
	}
	if count == 0 || count == 8 {
		return elf
	}
	for i := -1; i < 3; i++ {
		d := ((i + currentDirection) % 4) + 1
		offset := 0
		if d == 2 {
			offset = 6
			d = 1
		} else if d == 4 {
			offset = 2
			d = 3
		}
		found := false
		for n := 0; n < 3; n++ {
			if posMap[offset+(n*d)] {
				found = true
				break
			}
		}
		if !found {
			// proposed direction
			switch ((i + currentDirection) % 4) + 1 {
			case 1:
				return pos{x: elf.x, y: elf.y - 1}
			case 2:
				return pos{x: elf.x, y: elf.y + 1}
			case 3:
				return pos{x: elf.x - 1, y: elf.y}
			case 4:
				return pos{x: elf.x + 1, y: elf.y}
			}
		}
	}
	return elf
}
