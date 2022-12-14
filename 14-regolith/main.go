package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	cave := make(map[position]objType)
	deepest := 0

	for s := bufio.NewScanner(os.Stdin); s.Scan(); {
		positions, err := parseLine(s.Text())
		if err != nil {
			log.Fatal(err)
		}

		drawRocks(cave, positions)

		if d := deepestPoint(positions).depth(); d > deepest {
			deepest = d
		}
	}

	entrypoint := position{0, 500}

	count := 1
	for ; simulate(cave, entrypoint, deepest); count++ {
	}

	fmt.Println(count)

}

func simulate(cave map[position]objType, sandPosition position, baseDepth int) bool {
	entryPoint := sandPosition
	for {
		if sandPosition.below().depth() == baseDepth+2 {
			cave[sandPosition] = sand
			return true
		}

		if cave[sandPosition.below()] == air {
			sandPosition = sandPosition.below()
			continue
		}
		if cave[sandPosition.diagLeft()] == air {
			sandPosition = sandPosition.diagLeft()
			continue
		}
		if cave[sandPosition.diagRight()] == air {
			sandPosition = sandPosition.diagRight()
			continue
		}

		if sandPosition == entryPoint {
			return false
		}

		cave[sandPosition] = sand
		return true
	}
}

func deepestPoint(points []position) position {
	best := 0
	index := 0
	for i, p := range points {
		if p.depth() > best {
			best = p.depth()
			index = i
		}
	}
	return points[index]
}

func drawRocks(cave map[position]objType, points []position) {
	for i := 0; i < len(points)-1; i++ {
		from := points[i]
		to := points[i+1]

		if from.x > to.x || from.y > to.y {
			from, to = to, from
		}

		for x := from.x; x <= to.x; x++ {
			for y := from.y; y <= to.y; y++ {
				cave[position{x, y}] = rock
			}
		}
	}
}

func parseLine(line string) ([]position, error) {
	var positions []position
	tokens := strings.Split(line, " -> ")
	for _, s := range tokens {
		var x, y int
		_, err := fmt.Sscanf(s, "%d,%d", &y, &x)
		if err != nil {
			return nil, err
		}
		positions = append(positions, position{x, y})
	}
	return positions, nil
}

type objType int

const (
	air objType = iota
	rock
	sand
)

type position struct {
	x, y int
}

func (p position) depth() int {
	return p.x
}

func (p position) below() position {
	return position{p.x + 1, p.y}
}

func (p position) diagLeft() position {
	return position{p.x + 1, p.y - 1}
}

func (p position) diagRight() position {
	return position{p.x + 1, p.y + 1}
}
