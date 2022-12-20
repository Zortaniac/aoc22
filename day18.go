package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type drop struct {
	x int
	y int
	z int
}

func day18() {

	readFile, err := os.Open("day18.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	var drops []drop

	for fileScanner.Scan() {
		line := strings.TrimSpace(fileScanner.Text())

		if len(line) == 0 {
			continue
		}
		var d drop
		_, err = fmt.Sscanf(line, "%d,%d,%d", &d.x, &d.y, &d.z)
		if err != nil {
			fmt.Println(err)
			continue
		}
		drops = append(drops, d)
	}
	err = readFile.Close()
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(drops)
	sides := len(drops) * 6
	for i := 0; i < len(drops); i++ {
		for n := i + 1; n < len(drops); n++ {
			a := drops[i]
			b := drops[n]

			dist := Abs(a.x-b.x) + Abs(a.y-b.y) + Abs(a.z-b.z)
			if dist == 1 {
				sides -= 2
			}
		}
	}
	fmt.Println(sides)
	maxX := 0
	maxY := 0
	maxZ := 0
	for i := 0; i < len(drops); i++ {
		// remove existing drops from potential air bubbles
		a := drops[i]
		if a.x > maxX {
			maxX = a.x
		}
		if a.y > maxY {
			maxY = a.y
		}
		if a.z > maxZ {
			maxZ = a.z
		}
	}

	lava := make([][][]int8, maxX+1)

	for i := range lava {
		lava[i] = make([][]int8, maxY+1)
		for j := range lava[i] {
			lava[i][j] = make([]int8, maxZ+1)
		}
	}

	for _, d := range drops {
		lava[d.x][d.y][d.z] = 1
	}
	changed := true
	for changed {
		changed = false
		for x := range lava {
			for y := range lava[x] {
				for z := range lava[x][y] {
					if lava[x][y][z] > 0 {
						//fmt.Print(1)
						continue
					}
					if x == 0 || x == len(lava)-1 {
						lava[x][y][z] = 2
						changed = true
					} else if y == 0 || y == len(lava[x])-1 {
						lava[x][y][z] = 2
						changed = true
					} else if z == 0 || z == len(lava[x][y])-1 {
						lava[x][y][z] = 2
						changed = true
					} else if lava[x-1][y][z] == 2 || lava[x+1][y][z] == 2 ||
						lava[x][y-1][z] == 2 || lava[x][y+1][z] == 2 ||
						lava[x][y][z-1] == 2 || lava[x][y][z+1] == 2 {
						lava[x][y][z] = 2
						changed = true
					}
				}
				//fmt.Println()
			}
			//fmt.Println()
			//fmt.Println()
		}
	}
	sides = 0
	for x := range lava {
		for y := range lava[x] {
			for z := range lava[x][y] {
				if lava[x][y][z] != 1 {
					continue
				}
				if x == 0 || lava[x-1][y][z] == 2 {
					sides++
				}
				if x == len(lava)-1 || lava[x+1][y][z] == 2 {
					sides++
				}
				if y == 0 || lava[x][y-1][z] == 2 {
					sides++
				}
				if y == len(lava[x])-1 || lava[x][y+1][z] == 2 {
					sides++
				}
				if z == 0 || lava[x][y][z-1] == 2 {
					sides++
				}
				if z == len(lava[x][y])-1 || lava[x][y][z+1] == 2 {
					sides++
				}
			}
		}
	}
	fmt.Println(sides)
}
