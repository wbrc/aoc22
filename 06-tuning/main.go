package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	runes := make([]rune, 0, len(data))
	for _, r := range string(data) {
		runes = append(runes, r)
	}

	for i := 0; i < len(runes)-14; i++ {
		if isMarker(runes[i : i+14]) {
			fmt.Println(i + 14)
			break
		}
	}
}

func isMarker(frame []rune) bool {
	for i := 0; i < len(frame)-1; i++ {
		for j := i + 1; j < len(frame); j++ {
			if frame[i] == frame[j] {
				return false
			}
		}
	}
	return true
}
