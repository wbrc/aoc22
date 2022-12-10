package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func decodeProgram(r io.Reader) <-chan instruction {
	instructions := make(chan instruction)

	go func() {
		for s := bufio.NewScanner(r); s.Scan(); {
			if s.Text() == "noop" {
				instructions <- instruction{isNoop: true}
			} else {
				var inc int
				_, err := fmt.Sscanf(s.Text(), "addx %d", &inc)
				if err != nil {
					log.Print(err)
				}
				instructions <- instruction{inc: inc}
			}
		}
		close(instructions)
	}()

	return instructions
}

func main() {

	f, err := os.Open("input.txt")
	if err != nil {
		log.Print(err)
	}
	defer f.Close()

	instructions := decodeProgram(f)

	var (
		regx  = 1
		cycle = 0
		pos   = 0
	)

	for currentInstruction := range instructions {
		cycle++
		pos++

		if pos == regx || pos == regx+1 || pos == regx+2 {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}
		if pos == 40 {
			fmt.Println()
			pos = 0
		}

		if currentInstruction.isNoop {
			continue
		}

		cycle++
		pos++

		if pos == regx || pos == regx+1 || pos == regx+2 {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}
		if pos == 40 {
			fmt.Println()
			pos = 0
		}

		regx += currentInstruction.inc
	}

}

type instruction struct {
	isNoop bool
	inc    int
}
