package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func readFile(input string) []byte {
	data, err := os.ReadFile(input)
	if err != nil {
		panic(err)
	}
	return data
}

func main() {
	data := readFile("input.txt")
	houses := make(map[point]int)
	var currentSanta *point
	santaIndex := point{0, 0}
	robotIndex := point{0, 0}
	houses[santaIndex]++
	houses[robotIndex]++
	for index, direction := range data {
		if index%2 == 0 {
			currentSanta = &santaIndex
		} else {
			currentSanta = &robotIndex
		}
		switch direction {
		case '^':
			currentSanta.y++
		case 'v':
			currentSanta.y--
		case '>':
			currentSanta.x++
		case '<':
			currentSanta.x--
		}
		houses[point{currentSanta.x, currentSanta.y}]++
	}

	for k, v := range houses {
		fmt.Println(k, v)
	}

	fmt.Println("Houses visited: ", len(houses))
}
