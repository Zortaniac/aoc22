package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	packets := [][]interface{}{{[]interface{}{2}}, {[]interface{}{6}}}

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
		if r := compare(left, right); r == 1 {
			sum += pair
		}
		pair++
		packets = append(packets, left)
		packets = append(packets, right)
	}
	err = readFile.Close()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sum)
	sort.Slice(packets, func(i, j int) bool {
		return compare(packets[i], packets[j]) == 1
	})
	div := 1
	for k, p := range packets {
		if len(p) != 1 {
			continue
		}
		switch p[0].(type) {
		case []interface{}:
			v := p[0].([]interface{})
			if len(v) == 1 && (v[0] == 2 || v[0] == 6) {
				div *= k + 1
			}
		}
	}
	fmt.Println(div)
}

func compare(left []interface{}, right []interface{}) int {
	for i := 0; i < len(left); i++ {
		if len(right) <= i {
			// not right order
			return -1
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
			if left[i].(int) < right[i].(int) {
				return 1
			}
			return -1
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

		result := compare(newLeft, newRight)
		if result != 0 {
			return result
		}
	}
	if len(left) < len(right) {
		return 1
	}
	return 0
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
