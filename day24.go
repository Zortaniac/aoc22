package main

import (
	"bufio"
	"fmt"
	"os"
)

type blizzard struct {
	pos       pos
	direction int
}

func day24() {

	readFile, err := os.Open("day24.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var blizzards []blizzard
	blizzardPos := make(map[pos]struct{})

	fileScanner.Scan()
	y := 0
	dimX := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()

		if len(line) == 0 {
			break
		}
		dimX = len(line) - 3
		for x := 1; x < len(line)-1; x++ {
			direction := -1
			switch line[x] {
			case '^':
				direction = 0
			case '>':
				direction = 1
			case 'v':
				direction = 2
			case '<':
				direction = 3
			}
			if direction > -1 {
				position := pos{x: x - 1, y: y}
				blizzards = append(blizzards, blizzard{
					pos:       position,
					direction: direction,
				})
				blizzardPos[position] = struct{}{}
			}
		}
		y++
	}
	dim := pos{x: dimX, y: y - 2}
	err = readFile.Close()
	if err != nil {
		fmt.Println(err)
	}
	minute := 0
	foundTarget := false
	start := pos{0, -1}
	target := pos{dim.x, dim.y + 1}
	secondTime := false
	movements := make(map[pos]struct{})
	movements[pos{x: 0, y: -1}] = struct{}{}
	for true {
		minute++
		blizzards, blizzardPos = updatePositions(blizzards, dim)
		newMovements := make(map[pos]struct{})
		for m := range movements {
			for i := 0; i < len(directions); i++ {
				n := m.Add(directions[i])
				if n != start && n != target {
					if n.x < 0 {
						continue
					}
					if n.y < 0 {
						continue
					}
					if n.x > dim.x {
						continue
					}
					if n.y > dim.y {
						continue
					}
				}
				if _, exists := blizzardPos[n]; !exists {
					newMovements[n] = struct{}{}
				}
			}
			if _, exists := blizzardPos[m]; !exists {
				newMovements[m] = struct{}{}
			}
		}
		movements = newMovements

		if !foundTarget || secondTime {
			if _, exists := movements[target]; exists {
				fmt.Println(minute)
				foundTarget = true
				movements = map[pos]struct{}{target: {}}

				if secondTime {
					break
				}
			}
		}
		if foundTarget && !secondTime {
			if _, exists := movements[start]; exists {
				secondTime = true
				movements = map[pos]struct{}{start: {}}
			}
		}
		/*
			fmt.Println(movements)

			fmt.Print("#.")
			for x := 0; x <= dim.x; x++ {
				fmt.Print("#")
			}
			fmt.Println()
			for y := 0; y <= dim.y; y++ {
				fmt.Print("#")
				for x := 0; x <= dim.x; x++ {
					if _, exists := blizzardPos[pos{x: x, y: y}]; exists {
						fmt.Print("Z")
					} else {
						fmt.Print(".")
					}
				}
				fmt.Println("#")
			}
			for x := 0; x <= dim.x; x++ {
				fmt.Print("#")
			}
			fmt.Print(".#")
			fmt.Println()*/
	}
}

var directions = [4]pos{{y: -1}, {x: 1}, {y: 1}, {x: -1}}

func updatePositions(blizzards []blizzard, dim pos) ([]blizzard, map[pos]struct{}) {
	newBlizzards := make([]blizzard, len(blizzards))
	blizzardPos := make(map[pos]struct{}, len(blizzards))

	for i, bz := range blizzards {
		b := bz.pos.Add(directions[bz.direction])
		if b.x < 0 {
			b.x = dim.x
		}
		if b.x > dim.x {
			b.x = 0
		}
		if b.y < 0 {
			b.y = dim.y
		}
		if b.y > dim.y {
			b.y = 0
		}
		newBlizzards[i] = blizzard{pos: b, direction: bz.direction}
		blizzardPos[b] = struct{}{}
	}
	return newBlizzards, blizzardPos
}

func (p pos) Add(b pos) pos {
	return pos{x: p.x + b.x, y: p.y + b.y}
}
