package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type monkeyInstruction struct {
	monkeyA string
	monkeyB string
	operation byte
}

func day21() {

	readFile, err := os.Open("day21.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	numbers := make(map[string]int)
	monkeys := make(map[string]monkeyInstruction)

	for fileScanner.Scan() {
		line := strings.TrimSpace(fileScanner.Text())

		if len(line) == 0 {
			continue
		}
		parts := strings.Split(line, ": ")

		monkey := parts[0]
		number, err := strconv.Atoi(strings.TrimSpace(parts[1]))
		if err == nil {
			numbers[monkey] = number
			continue
		}

		var instruction monkeyInstruction
		_, err = fmt.Sscanf(strings.TrimSpace(parts[1]), "%s %c %s", &instruction.monkeyA, &instruction.operation, &instruction.monkeyB)
		if err != nil {
			fmt.Println(err)
			continue
		}
		monkeys[monkey] = instruction
	}
	err = readFile.Close()
	if err != nil {
		fmt.Println(err)
	}
	{
		// part 1
		numbers1 := make(map[string]int, len(numbers))
		monkeys1 := make(map[string]monkeyInstruction, len(monkeys))
		for k, v := range monkeys {
			monkeys1[k] = v
		}
		for k, v := range numbers {
			numbers1[k] = v
		}

		for len(monkeys1) > 0 {
			for k, i := range monkeys1 {
				if _, exists := numbers1[i.monkeyA]; !exists {
					continue
				}
				if _, exists := numbers1[i.monkeyB]; !exists {
					continue
				}
				switch i.operation {
				case '+':
					numbers1[k] = numbers1[i.monkeyA] + numbers1[i.monkeyB]
				case '-':
					numbers1[k] = numbers1[i.monkeyA] - numbers1[i.monkeyB]
				case '*':
					numbers1[k] = numbers1[i.monkeyA] * numbers1[i.monkeyB]
				case '/':
					numbers1[k] = numbers1[i.monkeyA] / numbers1[i.monkeyB]
				}
				delete(monkeys1, k)
			}
		}
		fmt.Println(numbers1["root"])
	}

	changed := true
	for changed {
		changed = false
		for k, i := range monkeys {
			if k == "root" {
				continue
			}
			if i.monkeyA == "humn" || i.monkeyB == "humn" {
				continue
			}
			if _, exists := numbers[i.monkeyA]; !exists {
				continue
			}
			if _, exists := numbers[i.monkeyB]; !exists {
				continue
			}
			numbers[k] = execOperation(i.operation, numbers[i.monkeyA], numbers[i.monkeyB])
			delete(monkeys, k)
			changed = true
		}
	}
	var value int
	var instruction monkeyInstruction
	if v, exists := numbers[monkeys["root"].monkeyA]; exists {
		value = v
		instruction = monkeys[monkeys["root"].monkeyB]
	} else {
		value = numbers[monkeys["root"].monkeyB]
		instruction = monkeys[monkeys["root"].monkeyA]
	}
	delete(numbers, "humn")
	fmt.Println(evaluate(value, instruction, &numbers, &monkeys))
}

func execOperation(op byte, a int, b int) int {
	switch op {
	case '+':
		return a + b
	case '-':
		return a - b
	case '*':
		return a * b
	case '/':
		return a / b
	}
	return 0
}

func evaluate(value int, instruction monkeyInstruction, numbers *map[string]int, monkeys *map[string]monkeyInstruction) int {
	if _, exists := (*numbers)[instruction.monkeyA]; !exists {
		b := (*numbers)[instruction.monkeyB]
		var a int
		switch instruction.operation {
		case '+':
			a = value - b
		case '-':
			a = value + b
		case '*':
			a = value / b
		case '/':
			a = value * b
		}
		if instruction.monkeyA == "humn" {
			return a
		}
		return evaluate(a, (*monkeys)[instruction.monkeyA], numbers, monkeys)
	}
	if _, exists := (*numbers)[instruction.monkeyB]; !exists {
		a := (*numbers)[instruction.monkeyA]
		var b int
		switch instruction.operation {
		case '+':
			b = value - a
		case '-':
			b = a - value
		case '*':
			b = value / a
		case '/':
			b = a / value
		}
		if instruction.monkeyB == "humn" {
			return b
		}
		return evaluate(b, (*monkeys)[instruction.monkeyB], numbers, monkeys)
	}
	return value
}