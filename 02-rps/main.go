package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	totalScore := 0

	for s := bufio.NewScanner(os.Stdin); s.Scan(); {
		line := strings.TrimSpace(s.Text())
		action, outcome, err := ParseLine(line)
		if err != nil {
			log.Print(err)
			continue
		}

		var reaction = action
		if outcome == Win {
			reaction = action.Killer()
		} else if outcome == Lose {
			reaction = action.Killer().Killer()
		}

		totalScore += Score(action, reaction)
		totalScore += int(reaction)
	}

	fmt.Println(totalScore)

}

type Hand int

const (
	Rock Hand = iota + 1
	Paper
	Scissors
)

func (h Hand) Killer() Hand {
	switch h {
	case Rock:
		return Paper
	case Paper:
		return Scissors
	case Scissors:
		return Rock
	}
	panic("unknown hand")
}

type Outcome int

const (
	Lose Outcome = iota
	Draw
	Win
)

func Score(action, reaction Hand) int {
	if action == reaction {
		return 3
	}

	switch action {
	case Rock:
		if reaction == Paper {
			return 6
		}
	case Paper:
		if reaction == Scissors {
			return 6
		}
	case Scissors:
		if reaction == Rock {
			return 6
		}
	}
	return 0
}

func ParseLine(line string) (action Hand, outcome Outcome, err error) {
	var a, r rune

	_, err = fmt.Sscanf(line, "%c %c", &a, &r)
	if err != nil {
		return
	}

	switch a {
	case 'A':
		action = Rock
	case 'B':
		action = Paper
	case 'C':
		action = Scissors
	default:
		err = fmt.Errorf("parse error: unknown typo %v", a)
		return
	}

	switch r {
	case 'X':
		outcome = Lose
	case 'Y':
		outcome = Draw
	case 'Z':
		outcome = Win
	default:
		err = fmt.Errorf("parse error: unknown typo %v", a)
		return
	}

	return
}
