package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func day3a() {

	readFile, err := os.Open("day3.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	score := 0
	for fileScanner.Scan() {
		l := strings.TrimSpace(fileScanner.Text())

		size := len(l) / 2
		first := l[0:size]
		second := l[size:]

		lookup := make(map[byte]struct{}, size)

		for i := 0; i < size; i++ {
			if _, exist := lookup[first[i]]; !exist {
				lookup[first[i]] = struct{}{}
			}
		}

		for i := 0; i < size; i++ {
			if _, exist := lookup[second[i]]; exist {
				c := int(second[i])
				if c-97 < 0 {
					// uppercase
					score += c - 38
				} else {
					// lowercase
					score += c - 96
				}
				delete(lookup, second[i])

			}
		}
	}
	err = readFile.Close()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(score)
}

func day3b() {

	readFile, err := os.Open("day3.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	score := 0
	groupCount := 0
	lookupGroup := make(map[byte]int, 26*2)
	for fileScanner.Scan() {
		l := strings.TrimSpace(fileScanner.Text())
		if groupCount == 0 {
			for i := 65; i < 91; i++ {
				lookupGroup[byte(i)] = 0
			}
			for i := 97; i < 123; i++ {
				lookupGroup[byte(i)] = 0
			}
		}

		lookup := make(map[byte]struct{}, len(l))

		for i := 0; i < len(l); i++ {
			if _, exist := lookup[l[i]]; !exist {
				lookup[l[i]] = struct{}{}
			}
		}
		for k, _ := range lookup {
			lookupGroup[k] += 1
		}
		groupCount += 1
		groupCount %= 3
		if groupCount == 0 {
			for c, a := range lookupGroup {

				if a == 3 {
					if int(c)-97 < 0 {
						// uppercase
						score += int(c) - 38
					} else {
						// lowercase
						score += int(c) - 96
					}
					break
				}
			}
		}

	}

	err = readFile.Close()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(score)
}
