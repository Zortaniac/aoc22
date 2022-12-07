package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type node struct {
	name     string
	children []*node
	size     int
	parent   *node
}

func day7() {

	readFile, err := os.Open("day7.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	root := node{}
	pwd := &root
	currentSize := 0
	for fileScanner.Scan() {
		line := strings.TrimSpace(fileScanner.Text())
		if line[0] == '$' {
			if currentSize > 0 {
				pwd.size = currentSize
				currentSize = 0
			}
			//command
			if line[2] == 'c' {
				// cd command
				dir := line[5:]
				if dir == "/" {
					pwd = &root
				} else if dir == ".." {
					pwd = pwd.parent
				} else {
					pwd = changeDir(pwd, dir)
				}
			}
		} else {
			// output
			if line[0] != 'd' {
				var size int
				var file string
				_, err = fmt.Sscanf(line, "%d %s", &size, &file)
				if err != nil {
					fmt.Println(err)
				}
				currentSize += size
			}
		}
	}

	if currentSize > 0 {
		pwd.size = currentSize
		currentSize = 0
	}

	err = readFile.Close()
	if err != nil {
		fmt.Println(err)
	}

	sum := 0
	total := 0
	findBelow100k(&root, &sum, &total)
	fmt.Println(sum)

	requiredSpace := 30000000 - (70000000 - total)
	current := math.MaxInt
	findSmallest(&root, &current, requiredSpace)
	fmt.Println(current)
}

func findBelow100k(pwd *node, sum *int, total *int) int {
	size := 0
	for i := range pwd.children {
		size += findBelow100k(pwd.children[i], sum, total)
	}
	size += pwd.size
	*total += pwd.size
	if size < 100000 {
		*sum += size
	}
	return size
}

func findSmallest(pwd *node, current *int, required int) int {
	size := 0
	for i := range pwd.children {
		size += findSmallest(pwd.children[i], current, required)
	}
	size += pwd.size
	if size > required && *current > size {
		*current = size
	}
	return size
}

func changeDir(pwd *node, dir string) *node {
	for i := range pwd.children {
		if dir == pwd.children[i].name {
			return pwd.children[i]
		}
	}
	n := node{parent: pwd, name: dir}
	pwd.children = append(pwd.children, &n)
	return &n
}
