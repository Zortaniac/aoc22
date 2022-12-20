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

}
