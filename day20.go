package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type item struct {
	x         int
	processed bool
}

func day20() {

	readFile, err := os.Open("day20.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	var numbers []item

	for fileScanner.Scan() {
		line := strings.TrimSpace(fileScanner.Text())

		if len(line) == 0 {
			continue
		}
		x, _ := strconv.Atoi(line)
		numbers = append(numbers, item{x: x})
	}
	err = readFile.Close()
	if err != nil {
		fmt.Println(err)
	}
	pos := 0
	for pos < len(numbers) {
		i := numbers[pos]
		if i.processed {
			pos++
			continue
		}
		if i.x == 0 {
			pos++
		} else {
			newPos := (pos + i.x) % (len(numbers) - 1)
			if newPos <= 0 {
				newPos = len(numbers) - 1 + newPos
			}
			without := append(numbers[:pos], numbers[pos+1:]...)
			newNumbers := make([]item, len(numbers))
			copy(newNumbers, without[:newPos])
			i.processed = true
			newNumbers[newPos] = i
			copy(newNumbers[newPos+1:], without[newPos:])
			numbers = newNumbers
			if newPos < pos {
				pos++
			}
		}
		//break
	}
	var nullIdx int
	for i := 0; i < len(numbers); i++ {
		if numbers[i].x != 0 {
			continue
		}
		nullIdx = i
		break
	}
	fmt.Println(numbers[(nullIdx+1000)%len(numbers)].x + numbers[(nullIdx+2000)%len(numbers)].x + numbers[(nullIdx+3000)%len(numbers)].x)

}
