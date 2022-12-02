package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func day1() {
	readFile, err := os.Open("day1.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	round := 0
	var blocks []int
	for fileScanner.Scan() {
		l := strings.TrimSpace(fileScanner.Text())
		if len(l) == 0 {
			// end of block
			blocks = append(blocks, round)
			round = 0
			continue
		}
		v, err := strconv.Atoi(l)
		if err != nil {
			fmt.Println(err)
		}
		round += v
	}
	if round > 0 {
		blocks = append(blocks, round)
	}
	err = readFile.Close()
	if err != nil {
		fmt.Println(err)
	}

	sort.Ints(blocks)

	// part one
	max := blocks[len(blocks)-1]
	fmt.Println(max)

	// part two
	sum := 0
	for _, x := range blocks[len(blocks)-3:] {
		sum += x
	}
	fmt.Println(sum)
}
