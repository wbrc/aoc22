package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	sum := 0
	for s := bufio.NewScanner(os.Stdin); s.Scan(); {
		rangeA, rangeB, err := parseLine(s.Text())
		if err != nil {
			log.Print(err)
			continue
		}

		if rangeA.OverlapsWith(rangeB) {
			sum++
		}
	}

	fmt.Println(sum)
}

type Range struct {
	Begin, End int
}

func (r Range) Contains(other Range) bool {
	return r.Begin <= other.Begin && r.End >= other.End
}

func (r Range) OverlapsWith(other Range) bool {
	for i := r.Begin; i <= r.End; i++ {
		if other.Contains(Range{i, i}) {
			return true
		}
	}

	return false
}

func parseLine(line string) (Range, Range, error) {
	var one, two Range
	_, err := fmt.Sscanf(line, "%d-%d,%d-%d", &one.Begin, &one.End, &two.Begin, &two.End)
	return one, two, err
}
