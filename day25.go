package main

import (
	"bufio"
	"fmt"
	"os"
)

func day25() {

	readFile, err := os.Open("day25.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	mapping := map[byte]int{'2': 2, '1': 1, '0': 0, '-': -1, '=': -2}
	pow5Table := make([]int, 20)
	x := 1
	for i := range pow5Table {
		pow5Table[i] = x
		x *= 5
	}
	sum := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()

		if len(line) == 0 {
			break
		}

		var numberParts []int
		for i := range line {
			numberParts = append([]int{mapping[line[i]]}, numberParts...)
		}
		number := 0
		for i, n := range numberParts {
			number += pow5Table[i] * n
		}
		sum += number
	}
	err = readFile.Close()
	if err != nil {
		fmt.Println(err)
	}
	firstPos := 0
	for i := 0; i < len(pow5Table); i++ {
		if pow5Table[i]*2 >= sum {
			firstPos = i
			break
		}
	}
	reverseMapping := make(map[int]byte, len(mapping))
	for k, v := range mapping {
		reverseMapping[v] = k
	}
	var bobsNumber []byte
	for ; firstPos >= 0; firstPos-- {
		abs := Abs(sum)
		n := 0
		for i := -2; i <= 2; i++ {
			a := Abs(sum - pow5Table[firstPos]*i)
			if abs > a {
				abs = a
				n = i
			}
		}

		bobsNumber = append(bobsNumber, reverseMapping[n])
		sum -= pow5Table[firstPos] * n
	}
	fmt.Println(string(bobsNumber))
}
