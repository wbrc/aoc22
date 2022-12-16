package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

func main() {

	var minCol, maxCol int = math.MaxInt, math.MinInt
	var sensors []sensor

	occupiedPos := make(map[position]struct{})

	for s := bufio.NewScanner(os.Stdin); s.Scan(); {
		var sensorPos, beaconPos position
		_, err := fmt.Sscanf(s.Text(), "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sensorPos.col, &sensorPos.row, &beaconPos.col, &beaconPos.row)
		if err != nil {
			log.Fatal(err)
		}

		if sensorPos.col-sensorPos.distance(beaconPos) < minCol {
			minCol = sensorPos.col - sensorPos.distance(beaconPos)
		}
		if sensorPos.col+sensorPos.distance(beaconPos) > maxCol {
			maxCol = sensorPos.col + sensorPos.distance(beaconPos)
		}
		if beaconPos.col-sensorPos.distance(beaconPos) < minCol {
			minCol = beaconPos.col - sensorPos.distance(beaconPos)
		}
		if beaconPos.col+sensorPos.distance(beaconPos) > maxCol {
			maxCol = beaconPos.col + sensorPos.distance(beaconPos)
		}

		sensors = append(sensors, sensor{sensorPos, sensorPos.distance(beaconPos)})
		occupiedPos[sensorPos] = struct{}{}
		occupiedPos[beaconPos] = struct{}{}
	}

	go find(sensors, 0, 500000)
	go find(sensors, 500000, 1000000)
	go find(sensors, 1000000, 1500000)
	go find(sensors, 1500000, 2000000)
	go find(sensors, 2000000, 2500000)
	go find(sensors, 2500000, 3000000)
	go find(sensors, 3000000, 3500000)
	go find(sensors, 3500000, 4000000)

	<-make(chan struct{})
}

func find(sensors []sensor, startR, endR int) {

	for r := startR; r < endR; r++ {
	loop:
		for c := 0; c < 4000000; c++ {
			for _, s := range sensors {
				if (position{r, c}).within(s) {
					continue loop
				}
			}
			fmt.Println(c*4000000 + r)
			os.Exit(0)
		}
	}
}

type position struct {
	row, col int
}

type sensor struct {
	position
	radius int
}

func (p position) distance(to position) int {
	return abs(p.row-to.row) + abs(p.col-to.col)
}

func (p position) within(s sensor) bool {
	return p.distance(s.position) <= s.radius
}

func abs(a int) int {
	if a < 0 {
		a = -a
	}
	return a
}
