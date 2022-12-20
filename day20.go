package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type item struct {
	x    int
	orig int
	pos  int
}

func day20() {

	readFile, err := os.Open("day20.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	var numbers []item

	pos := 0
	for fileScanner.Scan() {
		line := strings.TrimSpace(fileScanner.Text())

		if len(line) == 0 {
			continue
		}
		x, _ := strconv.Atoi(line)
		numbers = append(numbers, item{orig: x, pos: pos})
		pos++
	}
	err = readFile.Close()
	if err != nil {
		fmt.Println(err)
	}
	for n := 0; n < len(numbers); n++ {
		numbers[n].x = (numbers[n].orig * 811589153) % (len(numbers) - 1)
	}

	for t := 0; t < 10; t++ {
		pos = 0
		for pos < len(numbers) {
			for n := 0; n < len(numbers); n++ {
				if numbers[n].pos != pos {
					continue
				}
				i := numbers[n]
				if i.x != 0 {
					newPos := (n + i.x) % (len(numbers) - 1)
					if newPos <= 0 {
						newPos = len(numbers) - 1 + newPos
					}
					without := append(numbers[:n], numbers[n+1:]...)
					newNumbers := make([]item, len(numbers))
					copy(newNumbers, without[:newPos])
					newNumbers[newPos] = i
					copy(newNumbers[newPos+1:], without[newPos:])
					numbers = newNumbers
				}
				break
			}
			pos++
			//fmt.Println(numbers)
			//break
		}
		if t != 0 && t != 9 {
			continue
		}
		var nullIdx int
		for i := 0; i < len(numbers); i++ {
			if numbers[i].x != 0 {
				continue
			}
			nullIdx = i
			break
		}
		result := numbers[(nullIdx+1000)%len(numbers)].orig + numbers[(nullIdx+2000)%len(numbers)].orig + numbers[(nullIdx+3000)%len(numbers)].orig
		if t == 0 {
			//fmt.Println(result)
		} else {
			fmt.Println(result * 811589153)
		}
	}

}
