package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	monkeys, err := parse(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	numMonkeys := len(monkeys)

	sm := 1
	for _, m := range monkeys {
		sm *= m.test
	}

	for round := 0; round < 10000; round++ {
		for i := 0; i < numMonkeys; i++ {
			monkey := monkeys[i]
			for _, item := range monkey.items {
				item = monkey.op(item)
				monkey.numInspected++
				item = item % sm
				if item%monkey.test == 0 {
					monkeys[monkey.trueCase].items = append(monkeys[monkey.trueCase].items, item)
				} else {
					monkeys[monkey.falseCase].items = append(monkeys[monkey.falseCase].items, item)
				}
			}
			monkey.items = nil
		}
	}

	monkeySlice := make([]*monkey, 0, len(monkeys))
	for _, m := range monkeys {
		monkeySlice = append(monkeySlice, m)
	}

	sort.Slice(monkeySlice, func(i, j int) bool {
		return monkeySlice[i].numInspected > monkeySlice[j].numInspected
	})

	fmt.Println(monkeySlice[0].numInspected * monkeySlice[1].numInspected)
}

func parse(r io.Reader) (map[int]*monkey, error) {
	monkeys := make(map[int]*monkey)
	monkeyID := 0
	var currentMonkey *monkey
	for s := bufio.NewScanner(r); s.Scan(); {
		switch {
		case strings.HasPrefix(s.Text(), "Monkey"):
			_, err := fmt.Sscanf(s.Text(), "Monkey %d:", &monkeyID)
			if err != nil {
				log.Print(s.Text())
				return nil, err
			}
			currentMonkey = new(monkey)
		case strings.HasPrefix(s.Text(), "  Starting items:"):
			list := strings.TrimPrefix(s.Text(), "  Starting items: ")
			for _, s := range strings.Split(list, ", ") {
				item, err := strconv.ParseInt(s, 10, 64)
				if err != nil {
					log.Print(s)
					return nil, err
				}
				currentMonkey.items = append(currentMonkey.items, int(item))
			}
		case strings.HasPrefix(s.Text(), "  Operation:"):
			var (
				op rune
				by string
			)
			_, err := fmt.Sscanf(s.Text(), "  Operation: new = old %c %s", &op, &by)
			if err != nil {
				log.Print(s.Text())
				return nil, err
			}
			operation, err := getOp(op, by)
			if err != nil {
				return nil, err
			}
			currentMonkey.op = operation
		case strings.HasPrefix(s.Text(), "  Test:"):
			var num int
			_, err := fmt.Sscanf(s.Text(), "  Test: divisible by %d", &num)
			if err != nil {
				log.Print(s.Text())
				return nil, err
			}
			currentMonkey.test = num
		case strings.HasPrefix(s.Text(), "    If"):
			var (
				testOutcome bool
				num         int
			)
			_, err := fmt.Sscanf(s.Text(), "    If %t: throw to monkey %d", &testOutcome, &num)
			if err != nil {
				log.Print(s.Text())
				return nil, err
			}
			if testOutcome {
				currentMonkey.trueCase = num
			} else {
				currentMonkey.falseCase = num
			}
		default:
			if currentMonkey == nil {
				log.Print(s.Text())
				return nil, errors.New("parse error")
			}
			monkeys[monkeyID] = currentMonkey
		}
	}
	monkeys[monkeyID] = currentMonkey
	return monkeys, nil
}

func getOp(op rune, by string) (operation, error) {
	if by == "old" {
		if op == '*' {
			return sqOp, nil
		} else if op == '+' {
			return mulByOp(2), nil
		}
		return nil, errors.New("unknown op")
	}

	num, err := strconv.ParseInt(by, 10, 64)
	if err != nil {
		return nil, err
	}

	if op == '*' {
		return mulByOp(int(num)), nil
	} else if op == '+' {
		return addOp(int(num)), nil
	}

	return nil, errors.New("unknown op")
}

type operation func(int) int

func mulByOp(x int) operation {
	return func(i int) int {
		return x * i
	}
}

var sqOp operation = func(i int) int {
	return i * i
}

func addOp(x int) operation {
	return func(i int) int {
		return i + x
	}
}

type monkey struct {
	items               []int
	op                  operation
	test                int
	trueCase, falseCase int
	numInspected        int
}
