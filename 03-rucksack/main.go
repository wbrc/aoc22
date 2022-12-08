package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sum := 0
	var elveA, elveB, elveC []Item
	for s := bufio.NewScanner(os.Stdin); s.Scan(); {
		items := parseLine(s.Text())

		if elveA == nil {
			elveA = items
			continue
		}

		if elveB == nil {
			elveB = items
			continue
		}

		if elveC == nil {
			elveC = items
			sum += findDuplicateItem(elveA, elveB, elveC).Priority()

			elveA = nil
			elveB = nil
			elveC = nil
		}
	}
	fmt.Println(sum)
}

type Item rune

func (i Item) Priority() int {
	if 'a' <= i && i <= 'z' {
		return int(i) - 'a' + 1
	}
	return int(i) - 'A' + 27
}

func parseLine(line string) []Item {
	items := make([]Item, 0, len(line))

	for _, r := range line {
		items = append(items, Item(r))
	}

	return items
}

func removeDuplicates(items []Item) []Item {
	reduced := make([]Item, 0, len(items))
	seen := make(map[Item]struct{})

	for _, item := range items {
		if _, ok := seen[item]; ok {
			continue
		}

		reduced = append(reduced, item)
		seen[item] = struct{}{}
	}

	return reduced
}

func findDuplicateItem(a, b, c []Item) Item {
	a, b, c = removeDuplicates(a), removeDuplicates(b), removeDuplicates(c)

	seen := make(map[Item]int)

	for _, item := range a {
		seen[item]++
	}

	for _, item := range b {
		seen[item]++
	}

	for _, item := range c {
		if seen[item] == 2 {
			return item
		}
	}

	panic("no duplicate found")
}
