package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var (
		maxCalories         [3]int
		currentElveCalories int
	)

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		line := strings.TrimSpace(s.Text())

		if line == "" {
			maxCalories = pushVal(maxCalories, currentElveCalories)
			currentElveCalories = 0
			continue
		}

		calories, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		currentElveCalories += int(calories)
	}
	maxCalories = pushVal(maxCalories, currentElveCalories)

	fmt.Println(maxCalories[0] + maxCalories[1] + maxCalories[2])
}

func pushVal(top3 [3]int, newVal int) [3]int {
	if newVal > top3[2] {
		top3[2] = newVal
	}

	if top3[2] > top3[1] {
		top3[1], top3[2] = top3[2], top3[1]
	}

	if top3[1] > top3[0] {
		top3[0], top3[1] = top3[1], top3[0]
	}

	return top3
}
