package main

import (
	"bufio"
	"container/list"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	var stacks []*list.List

	for s := bufio.NewScanner(os.Stdin); s.Scan(); {
		if strings.Contains(s.Text(), "[") {
			line := toRunes(s.Text())
			numStacks := numberOfStacks(line)

			if stacks == nil {
				stacks = initStacks(numStacks)
			}

			for i := 0; i < numStacks; i++ {
				if item := getItemAtPosition(line, i); item != ' ' {
					stacks[i].PushFront(item)
				}
			}
		} else if strings.Contains(s.Text(), "move") {
			var num, from, to int
			_, err := fmt.Sscanf(s.Text(), "move %d from %d to %d", &num, &from, &to)
			if err != nil {
				log.Print(err)
				continue
			}

			tmpList := list.New()
			for i := 0; i < num; i++ {
				tmpList.PushFront(stacks[from-1].Remove(stacks[from-1].Back()).(rune))
			}
			stacks[to-1].PushBackList(tmpList)
		}
	}

	for _, stack := range stacks {
		fmt.Printf("%c", stack.Back().Value.(rune))
	}
	fmt.Println()
}

func initStacks(num int) []*list.List {
	stacks := make([]*list.List, 0, num)
	for i := 0; i < num; i++ {
		stacks = append(stacks, list.New())
	}
	return stacks
}

func toRunes(line string) []rune {
	runes := make([]rune, 0, len(line))
	for _, r := range line {
		runes = append(runes, r)
	}
	return runes
}

func numberOfStacks(line []rune) int {
	return (len(line) + 1) / 4
}

func getItemAtPosition(line []rune, pos int) rune {
	return line[(pos)*4+1]
}

func isLetter(r rune) bool {
	return 'A' <= r && r <= 'Z'
}
