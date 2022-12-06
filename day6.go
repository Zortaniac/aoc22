package main

import (
	"bufio"
	"fmt"
	"os"
)

func day6() {

	readFile, err := os.Open("day6.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanRunes)

	var last4 [4]byte
	var last14 [14]byte
	for i := 0; i < 4; i++ {
		fileScanner.Scan()
		last4[i] = fileScanner.Text()[0]
		last14[i] = last4[i]
	}
	current := 4
	sopm := 0
	distinctCounter := 0
	for fileScanner.Scan() {
		c := fileScanner.Text()[0]
		last4[current%4] = c
		last14[current%14] = c
		current++

		if distinct4(last4) {
			if sopm == 0 {
				sopm = current
			}
			distinctCounter++
		} else {
			distinctCounter = 0
		}

		if distinctCounter >= 12 && distinct14(last14) {
			break
		}

	}

	fmt.Printf("start-of-packet marker: %d\n", sopm)
	fmt.Printf("start-of-message marker: %d\n", current)
	err = readFile.Close()
	if err != nil {
		fmt.Println(err)
	}
}

func distinct4(arr [4] byte) bool {
	for i := 0; i < 3; i++ {
		for n := i+1; n < 4; n++ {
			if arr[i] == arr[n] {
				return false
			}
		}
	}
	return true
}
func distinct14(arr [14] byte) bool {
	for i := 0; i < 13; i++ {
		for n := i+1; n < 14; n++ {
			if arr[i] == arr[n] {
				return false
			}
		}
	}
	return true
}