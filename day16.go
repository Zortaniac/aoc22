package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func day16() {

	readFile, err := os.Open("day16.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	valves := make(map[string]int)
	distances := make(map[string]int)
	connections := make(map[string][]string)

	for fileScanner.Scan() {
		line := strings.TrimSpace(fileScanner.Text())

		if len(line) == 0 {
			continue
		}

		parts := strings.Split(line, "; ")

		var valve string
		var flow int
		_, err = fmt.Sscanf(strings.TrimSpace(parts[0]), "Valve %s has flow rate=%d", &valve, &flow)
		if err != nil {
			fmt.Println(err)
			continue
		}
		connections[valve] = strings.Split(strings.TrimSpace(strings.ReplaceAll(strings.ReplaceAll(parts[1], "tunnel leads to valve", ""), "tunnels lead to valves", "")), ", ")
		valves[valve] = flow
		for _, c := range connections[valve] {
			distances[c+valve] = 1
			distances[valve+c] = 1
		}
	}
	err = readFile.Close()
	if err != nil {
		fmt.Println(err)
	}
	changed := true
	for changed {
		changed = false
		for a := range valves {
			for b := range valves {
				if a == b {
					continue
				}

				var current int
				if d, exists := distances[a+b]; exists {
					if d == 1 {
						continue
					}
					current = d
				}
				for _, c := range connections[a] {
					if _, exists := distances[c+b]; exists {
						d := distances[c+b]+1
						if current == 0 || current > d {
							distances[a+b] = d
							changed = true
						}
					}
				}
			}
		}
	}

	for v, f := range valves {
		if v != "AA" && f == 0 {
			delete(valves, v)
		}
	}

	fmt.Println(findBest("AA", 30, valves, distances))
	fmt.Println(findBest2Gether("AA", "AA", 26, 0, 0, valves, distances))
}

func findBest(current string, remainingMinutes int, valves map[string]int, distances map[string]int) int {
	if remainingMinutes <=0 {
		return 0
	}

	bestFlow := 0
	for v, f := range valves {
		if f == 0 {
			continue
		}
		distance := distances[current+v]+1
		if distance > remainingMinutes {
			continue
		}
		flow := (remainingMinutes-distance)*f
		valves[v] = 0
		flow += findBest(v, remainingMinutes-distance, valves, distances)
		if flow > bestFlow {
			bestFlow = flow
		}
		valves[v] = f
	}

	return bestFlow
}

func findBest2Gether(currentA string, currentB string, remainingMinutes int, distA int, distB int, valves map[string]int, distances map[string]int) int {
	if remainingMinutes <=0 {
		return 0
	}

	bestFlow := 0
	for v, f := range valves {
		if f == 0 {
			delete(valves, v)
			continue
		}
		distanceA := distances[currentA+v]+1+distA
		distanceB := distances[currentB+v]+1+distB
		if distanceA >= remainingMinutes && distanceB >= remainingMinutes {
			continue
		}
		flowA := (remainingMinutes-distanceA)*f
		flowB := (remainingMinutes-distanceB)*f
		delete(valves, v)
		var flow int
		if flowA > flowB {
			move := Min(distanceA, distB)
			flow = flowA + findBest2Gether(v, currentB, remainingMinutes-move, distanceA-move, distB-move, valves, distances)
		} else {
			move := Min(distanceB, distA)
			flow = flowB + findBest2Gether(currentA, v, remainingMinutes-move, distA-move, distanceB-move, valves, distances)
		}
		if flow > bestFlow {
			bestFlow = flow
		}
		valves[v] = f
	}

	return bestFlow
}

func findBestTogether(currentA string, currentB string, remainingMinutes int, distA int, distB int, valves map[string]int, distances map[string]int, swap bool) (int, []string, []int) {
	if remainingMinutes <=0 {
		return 0, nil, nil
	}
	release := 0
	current := currentA
	if distB == 0 {
		current = currentB
	}
	myFlow := valves[current]
	valves[current] = 0
	if myFlow > 0 {
		release = (remainingMinutes-1)*myFlow
	}
	optimum := 0
	var bestPath []string
	var bestrema []int

	for v, f := range valves {
		if f == 0 {
			continue
		}
		distance := distances[current+v]
		if distB == 0 {
			currentB = v
			distB = distance
			if myFlow > 0 {
				distB++
			}
		} else {
			currentA = v
			distA = distance
			if myFlow > 0 {
				distA++
			}
		}
		moveDistance := Min(distA, distB)
		flow, path, rem := findBestTogether(currentA, currentB, remainingMinutes-moveDistance, Abs(distA-moveDistance), Abs(distB-moveDistance), valves, distances, true)
		if flow > optimum {
			optimum = flow
			bestPath = path
			bestrema = rem
		}
	}
	valves[current] = myFlow
	return optimum+release, append(bestPath, current), append(bestrema, remainingMinutes)
}