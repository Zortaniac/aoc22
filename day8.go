package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func day8() {

	readFile, err := os.Open("day8.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var grid [][]int

	for fileScanner.Scan() {
		l := fileScanner.Text()

		if len(l) == 0 {
			break
		}
		row := make([]int, len(l))
		for i := 0; i < len(l); i++ {
			row[i], _ = strconv.Atoi(string(l[i]))
		}
		grid = append(grid, row)
	}

	visible := make([][]bool, len(grid))
	viewingDistance := make([][][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		visible[i] = make([]bool, len(grid[i]))
		viewingDistance[i] = make([][]int, len(grid[i]))
		for n := 0; n < len(grid[i]); n++ {
			viewingDistance[i][n] = make([]int, 4)
		}
	}

	last := len(visible) - 1
	for i := 0; i < len(visible[0]); i++ {
		visible[0][i] = true
		visible[last][i] = true
	}

	for r := 0; r < len(grid); r++ {
		visible[r][0] = true
		visible[r][len(visible[r])-1] = true

		lenRow := len(grid[r])
		currentLeft := grid[r][0]
		currentRight := grid[r][lenRow-1]

		for c := 0; c < len(grid[r]); c++ {
			if grid[r][c] > currentLeft {
				visible[r][c] = true
				currentLeft = grid[r][c]
				viewingDistance[r][c][0] += c
			} else if c > 0 {
				viewingDistance[r][c][0] = 1
				for d := 1; d <= c; d++ {
					if grid[r][c] <= grid[r][c-d] {
						break
					}
					viewingDistance[r][c][0]++
				}
			}
			if grid[r][lenRow-1-c] > currentRight {
				visible[r][lenRow-1-c] = true
				currentRight = grid[r][lenRow-1-c]
				viewingDistance[r][lenRow-1-c][1] += c
			} else if c > 0 {
				viewingDistance[r][lenRow-1-c][1] = 1
				for d := 1; d <= c; d++ {
					if grid[r][lenRow-1-c] <= grid[r][lenRow-1-c+d] {
						break
					}
					viewingDistance[r][lenRow-1-c][1]++
				}
			}
		}
	}

	for c := 0; c < len(grid[0]); c++ {
		lenCol := len(grid)
		currentTop := grid[0][c]
		currentBottom := grid[lenCol-1][c]

		for r := 0; r < lenCol; r++ {
			if grid[r][c] > currentTop {
				visible[r][c] = true
				currentTop = grid[r][c]
				viewingDistance[r][c][2] += r
			} else if r > 0 {
				viewingDistance[r][c][2] = 1
				for d := 1; d <= r; d++ {
					if grid[r][c] <= grid[r-d][c] {
						break
					}
					viewingDistance[r][c][2]++
				}
			}
			if grid[lenCol-1-r][c] > currentBottom {
				visible[lenCol-1-r][c] = true
				currentBottom = grid[lenCol-1-r][c]
				viewingDistance[lenCol-1-r][c][3] += r
			} else if r > 0 {
				viewingDistance[lenCol-1-r][c][3] = 1
				for d := 1; d <= r; d++ {
					if grid[lenCol-1-r][c] <= grid[lenCol-1-r+d][c] {
						break
					}
					viewingDistance[lenCol-1-r][c][3]++
				}
			}
		}
	}

	visibilityCount := 0
	highestSpot := 0
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[r]); c++ {
			if visible[r][c] {
				visibilityCount++
			}
			sum := 1
			for _, v := range viewingDistance[r][c] {
				sum *= v
			}
			if sum > highestSpot {
				highestSpot = sum
			}
		}
	}

	fmt.Println(visibilityCount)
	fmt.Println(highestSpot)

	err = readFile.Close()
	if err != nil {
		fmt.Println(err)
	}
}
