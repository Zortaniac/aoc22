package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type monkey struct {
	items     []int
	operation func(v int) int
	test      int
	ifTrue    int
	ifFalse   int
	inspected int
}
type monkeyExtended struct {
	items     [][]int
	operation func(v int) int
	test      int
	ifTrue    int
	ifFalse   int
	inspected int
}

func day11() {

	readFile, err := os.Open("day11.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var monkeys []monkey

	var lines []string
	for fileScanner.Scan() {
		line := fileScanner.Text()

		if len(line) > 0 {
			lines = append(lines, line)
		}
		if len(lines) == 6 {
			// process
			mnkey := monkey{}

			for _, item := range strings.Split(lines[1][18:], ", ") {
				i, _ := strconv.Atoi(item)
				mnkey.items = append(mnkey.items, i)
			}

			var operation byte
			var valueString string
			_, err = fmt.Sscanf(lines[2], "  Operation: new = old %c %s", &operation, &valueString)
			if operation == '+' {
				value, err := strconv.Atoi(valueString)
				if err != nil {
					mnkey.operation = func(v int) int {
						return v + v
					}
				} else {
					mnkey.operation = func(v int) int {
						return value + v
					}
				}
			} else {
				value, err := strconv.Atoi(valueString)
				if err != nil {
					mnkey.operation = func(v int) int {
						return v * v
					}
				} else {
					mnkey.operation = func(v int) int {
						return value * v
					}
				}
			}

			_, err = fmt.Sscanf(lines[3], "  Test: divisible by %d", &mnkey.test)
			_, err = fmt.Sscanf(lines[4], "    If true: throw to monkey %d", &mnkey.ifTrue)
			_, err = fmt.Sscanf(lines[5], "    If false: throw to monkey %d", &mnkey.ifFalse)
			monkeys = append(monkeys, mnkey)
			lines = lines[:0]
			continue
		}
	}

	var monkeysExt []monkeyExtended
	for _, m := range monkeys {
		monkeyExt := monkeyExtended{}
		monkeyExt.test = m.test
		monkeyExt.operation = m.operation
		monkeyExt.ifTrue = m.ifTrue
		monkeyExt.ifFalse = m.ifFalse
		for _, i := range m.items {
			var item []int
			for _, n := range monkeys {
				item = append(item, i%n.test)
			}
			monkeyExt.items = append(monkeyExt.items, item)
		}
		monkeysExt = append(monkeysExt, monkeyExt)
	}

	for r := 0; r < 20; r++ {
		for i := 0; i < len(monkeys); i++ {
			for _, item := range monkeys[i].items {
				item = monkeys[i].operation(item)
				item /= 3
				idx := monkeys[i].ifFalse
				if item%monkeys[i].test == 0 {
					idx = monkeys[i].ifTrue
				}
				monkeys[idx].items = append(monkeys[idx].items, item)
				monkeys[i].inspected++
			}
			monkeys[i].items = monkeys[i].items[:0]
		}
	}

	for r := 0; r < 10000; r++ {
		for i := 0; i < len(monkeysExt); i++ {
			for _, item := range monkeysExt[i].items {
				newItem := make([]int, len(monkeysExt))
				for n := 0; n < len(monkeysExt); n++ {
					newItem[n] = monkeysExt[i].operation(item[n])
					newItem[n] %= monkeysExt[n].test
				}
				idx := monkeysExt[i].ifFalse
				if newItem[i] == 0 {
					idx = monkeys[i].ifTrue
				}
				monkeysExt[idx].items = append(monkeysExt[idx].items, newItem)
				monkeysExt[i].inspected++
			}
			monkeysExt[i].items = monkeysExt[i].items[:0]
		}
	}
	var inspected []int
	for _, m := range monkeys {
		inspected = append(inspected, m.inspected)
	}
	sort.Slice(inspected, func(i, j int) bool {
		return inspected[i] > inspected[j]
	})
	fmt.Println(inspected[0] * inspected[1])

	inspected = inspected[:0]
	for _, m := range monkeysExt {
		inspected = append(inspected, m.inspected)
	}
	sort.Slice(inspected, func(i, j int) bool {
		return inspected[i] > inspected[j]
	})
	fmt.Println(inspected[0] * inspected[1])

	err = readFile.Close()
	if err != nil {
		fmt.Println(err)
	}
}
