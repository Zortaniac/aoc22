package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func day12() {

	readFile, err := os.Open("day12.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var heightMap [][]byte

	var startX, startY int
	for fileScanner.Scan() {
		line := strings.TrimSpace(fileScanner.Text())

		if len(line) == 0 {
			continue
		}

		heightMap = append(heightMap, []byte(line))
		y := len(heightMap) - 1
		for i := 0; i < len(heightMap[y]); i++ {
			if heightMap[y][i] == 'S' {
				heightMap[y][i] = 'a'
				startX = i
				startY = y
			}
		}
	}
	err = readFile.Close()
	if err != nil {
		fmt.Println(err)
	}
	lengthMap := make([][]int, len(heightMap))
	for i := 0; i < len(lengthMap); i++ {
		lengthMap[i] = make([]int, len(heightMap[i]))
	}

	changed := true
	for changed {
		changed = false
		for y := 0; y < len(lengthMap); y++ {
			for x := 0; x < len(lengthMap[y]); x++ {
				current := heightMap[y][x]
				if current == 'E' {
					continue
				}

				if test(x+1, y, current, heightMap, lengthMap) {
					tmp := lengthMap[y][x+1] + 1
					if lengthMap[y][x] == 0 || lengthMap[y][x] > tmp {
						lengthMap[y][x] = tmp
						changed = true
					}
				}
				if test(x-1, y, current, heightMap, lengthMap) {
					tmp := lengthMap[y][x-1] + 1
					if lengthMap[y][x] == 0 || lengthMap[y][x] > tmp {
						lengthMap[y][x] = tmp
						changed = true
					}
				}
				if test(x, y+1, current, heightMap, lengthMap) {
					tmp := lengthMap[y+1][x] + 1
					if lengthMap[y][x] == 0 || lengthMap[y][x] > tmp {
						lengthMap[y][x] = tmp
						changed = true
					}
				}
				if test(x, y-1, current, heightMap, lengthMap) {
					tmp := lengthMap[y-1][x] + 1
					if lengthMap[y][x] == 0 || lengthMap[y][x] > tmp {
						lengthMap[y][x] = tmp
						changed = true
					}
				}
			}
		}
	}

	fmt.Println(lengthMap[startY][startX])
	fmt.Println()
}

func test(x int, y int, current byte, heightMap [][]byte, lengthMap [][]int) bool {
	if y < 0 {
		return false
	}
	if x < 0 {
		return false
	}
	if y >= len(heightMap) {
		return false
	}
	if x >= len(heightMap[y]) {
		return false
	}
	if heightMap[y][x] == 'E' {
		return current == 'z' || current == 'y'
	}
	if lengthMap[y][x] == 0 {
		return false
	}
	if current+1 < heightMap[y][x] {
		return false
	}
	return true
}
