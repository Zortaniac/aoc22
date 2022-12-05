package main

import (
	"bufio"
	"fmt"
	"os"
)

func day5() {

	readFile, err := os.Open("day5.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var stacks []string
	for fileScanner.Scan() {
		l := fileScanner.Text()

		if len(l) == 0 {
			break
		}
		stacks = append(stacks, l)
	}

	shipCount := len(stacks[len(stacks)-1]) / 3
	shipsA := make([][]byte, shipCount)
	shipsB := make([][]byte, shipCount)
	for i := len(stacks) - 2; i >= 0; i-- {
		for n := 0; n <= shipCount; n++ {
			idx := (n * 4) + 1
			if idx >= len(stacks[i]) || stacks[i][idx] == ' ' {
				continue
			}
			shipsA[n] = append(shipsA[n], stacks[i][idx])
			shipsB[n] = append(shipsB[n], stacks[i][idx])
		}
	}
	for fileScanner.Scan() {
		l := fileScanner.Text()
		var count, from, to int
		_, err = fmt.Sscanf(l, "move %d from %d to %d", &count, &from, &to)
		if err != nil {
			fmt.Println(err)
			continue
		}
		from--
		to--
		for i := 0; i < count; i++ {
			shipsA[to] = append(shipsA[to], shipsA[from][len(shipsA[from])-1])
			shipsA[from] = shipsA[from][:len(shipsA[from])-1]
		}

		shipsB[to] = append(shipsB[to], shipsB[from][len(shipsB[from])-count:]...)
		shipsB[from] = shipsB[from][:len(shipsB[from])-count]

	}

	var resultA []byte
	var resultB []byte
	for i := 0; i < shipCount; i++ {
		if len(shipsA[i]) > 0 {
			resultA = append(resultA, shipsA[i][len(shipsA[i])-1])
		}
		if len(shipsB[i]) > 0 {
			resultB = append(resultB, shipsB[i][len(shipsB[i])-1])
		}
	}
	fmt.Println(string(resultA))
	fmt.Println(string(resultB))

	err = readFile.Close()
	if err != nil {
		fmt.Println(err)
	}
}
