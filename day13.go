package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func day13() {

	readFile, err := os.Open("day13.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	pair := 1
	sum := 0
	for fileScanner.Scan() {
		line1 := strings.TrimSpace(fileScanner.Text())

		if len(line1) == 0 {
			continue
		}
		fileScanner.Scan()
		line2 := strings.TrimSpace(fileScanner.Text())

		if len(line2) == 0 {
			continue
		}
		i := 0
		left := parse(line1, &i)
		i = 0
		right := parse(line2, &i)
		if r, _ := compare(left, right); r {
			sum += pair
		}
		pair++
	}
	err = readFile.Close()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sum)
}

func compare(left []interface{}, right []interface{}) (bool, bool) {
	for i := 0; i < len(left); i++ {
		if len(right) <= i {
			// not right order
			return false, true
		}
		leftItemList := false
		rightItemList := false

		switch left[i].(type) {
		case []interface{}:
			leftItemList = true
		}
		switch right[i].(type) {
		case []interface{}:
			rightItemList = true
		}

		if !leftItemList && !rightItemList {
			// both are integers
			//fmt.Println("Compare", left[i].(int), right[i].(int))
			if left[i].(int) == right[i].(int) {
				continue
			}
			return left[i].(int) < right[i].(int), true
		}

		var newLeft []interface{}
		var newRight []interface{}
		if !leftItemList {
			newLeft = make([]interface{}, 1)
			newLeft[0] = left[i]
		} else {
			newLeft = left[i].([]interface{})
		}
		if !rightItemList {
			newRight = make([]interface{}, 1)
			newRight[0] = right[i]
		} else {
			newRight = right[i].([]interface{})
		}

		result, success := compare(newLeft, newRight)
		if success {
			return result, true
		}
	}
	if len(left) < len(right) {
		return true, true
	}
	return false, false
}

func parse(line string, idx *int) []interface{} {
	var current []interface{}
	for *idx < len(line)-1 {
		*idx++
		if line[*idx] == '[' {
			current = append(current, parse(line, idx))
		} else if line[*idx] == ']' {
			break
		} else if line[*idx] == ',' {
			continue
		} else {
			current = append(current, readNumber(line, idx))
		}
	}
	return current
}

func readNumber(line string, idx *int) int {
	start := *idx
	for line[*idx] >= '0' && line[*idx] <= '9' {
		*idx++
	}
	value, _ := strconv.Atoi(line[start:*idx])
	*idx--
	return value
}
