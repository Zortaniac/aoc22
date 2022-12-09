package main

import (
	"bufio"
	"fmt"
	"os"
)

func day9() {

	readFile, err := os.Open("day9.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var head [2]int
	var tail [2]int
	visited := make(map[[2]int]bool)
	var rope [10][2]int
	visitedRope := make(map[[2]int]bool)

	for fileScanner.Scan() {
		line := fileScanner.Text()

		if len(line) == 0 {
			break
		}
		var direction byte
		var steps int

		_, err = fmt.Sscanf(line, "%c %d", &direction, &steps)
		//fmt.Println("=========", line)
		for s := 0; s < steps; s++ {
			switch direction {
			case 'R':
				head[0]++
				rope[0][0]++
				break
			case 'L':
				head[0]--
				rope[0][0]--
				break
			case 'U':
				head[1]++
				rope[0][1]++
				break
			case 'D':
				head[1]--
				rope[0][1]--
				break
			}
			{
				// part A
				absX := Abs(head[0] - tail[0])
				absY := Abs(head[1] - tail[1])
				movX := absX > 1
				movY := absY > 1

				if (movX || movY) && absX >= 1 && absY >= 1 {
					// move diagonally
					movX = true
					movY = true
				}
				if movX {
					if head[0] < tail[0] {
						tail[0]--
					} else {
						tail[0]++
					}
				}
				if movY {
					if head[1] < tail[1] {
						tail[1]--
					} else {
						tail[1]++
					}
				}
				visited[tail] = true
			}

			for i := 1; i < len(rope); i++ {
				absX := Abs(rope[i-1][0] - rope[i][0])
				absY := Abs(rope[i-1][1] - rope[i][1])
				movX := absX > 1
				movY := absY > 1

				if (movX || movY) && absX >= 1 && absY >= 1 {
					// move diagonally
					movX = true
					movY = true
				}
				if movX {
					if rope[i-1][0] < rope[i][0] {
						rope[i][0]--
					} else {
						rope[i][0]++
					}
				}
				if movY {
					if rope[i-1][1] < rope[i][1] {
						rope[i][1]--
					} else {
						rope[i][1]++
					}
				}
			}
			visitedRope[rope[9]] = true
			//fmt.Println("Head", head[0], head[1])
			//fmt.Println("Tail", tail[0], tail[1])
			//fmt.Println("-------")
		}
	}

	counter := 0
	for range visited {
		counter++
	}
	fmt.Println(counter)

	counter = 0
	for range visitedRope {
		counter++
	}
	fmt.Println(counter)

	err = readFile.Close()
	if err != nil {
		fmt.Println(err)
	}
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
