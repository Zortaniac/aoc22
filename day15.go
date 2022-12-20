package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func day15() {

	readFile, err := os.Open("day15.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	beacons := make(map[pos]pos)
	sensors := make(map[pos]pos)
	distances := make(map[pos]int)

	for fileScanner.Scan() {
		line := strings.TrimSpace(fileScanner.Text())

		if len(line) == 0 {
			continue
		}

		sensor := pos{}
		beacon := pos{}
		_, err = fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sensor.x, &sensor.y, &beacon.x, &beacon.y)
		if err != nil {
			fmt.Println(err)
			continue
		}

		beacons[beacon] = sensor
		sensors[sensor] = beacon
		distances[sensor] = tDist(sensor, beacon)
	}
	err = readFile.Close()
	if err != nil {
		fmt.Println(err)
	}

	{
		rowToCheck := 2000000
		checked := make(map[int]struct{})
		for sensor := range sensors {
			dist := distances[sensor]
			if sensor.y-rowToCheck > dist {
				continue
			}
			y := rowToCheck - sensor.y
			// check upper
			for x := -(dist - Abs(y)); x <= (dist - Abs(y)); x++ {
				if _, exists := beacons[pos{x: sensor.x + x, y: rowToCheck}]; !exists {
					checked[sensor.x+x] = struct{}{}
				}
			}
		}
		fmt.Println(len(checked))
	}
	{
		boundary := 4000000
		for y := 0; y <= boundary; y++ {
			for x := 0; x <= boundary; x++ {
				found := false
				for sensor := range sensors {
					distSensor := tDist(pos{x: x, y: y}, sensor)
					if distSensor <= distances[sensor] {
						x = xMove(sensor, y, distances[sensor])
						found = true
						break
					}
				}
				if x <= boundary && !found {
					fmt.Println(x*4000000 + y)
					return
				}
			}
		}
	}
}

func tDist(a pos, b pos) int {
	return Abs(a.x-b.x) + Abs(a.y-b.y)
}

func xMove(p pos, y int, d int) int {
	return d - Abs(p.y-y) + p.x
}
