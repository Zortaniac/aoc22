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

	rowToCheck := 2000000
	checked := make(map[int]struct{})
	for sensor, _ := range sensors {
		dist := distances[sensor]
		for y := 0; y < dist; y++ {
			if sensor.y-dist+y == rowToCheck || sensor.y+dist-y == rowToCheck {
				// check upper
				for x := 0; x <= y; x++ {
					if _, exists := beacons[pos{x: sensor.x - x, y: rowToCheck}]; !exists {
						checked[sensor.x-x] = struct{}{}
					}
					if _, exists := beacons[pos{x: sensor.x + x, y: rowToCheck}]; !exists {
						checked[sensor.x+x] = struct{}{}
					}
				}
			}
		}
	}
	fmt.Println(len(checked))
}

func tDist(a pos, b pos) int {
	return Abs(a.x-b.x) + Abs(a.y-b.y)
}
