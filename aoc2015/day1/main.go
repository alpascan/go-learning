package main

import (
	"fmt"
	"os"
)

func readFile(input string) string {
	data, err := os.ReadFile(input)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func getFloor(input string) {
	var floor int
	var position int
	for _, char := range input {
		position++
		switch char {
		case '(':
			floor++
		case ')':
			floor--
		}
		if floor == -1 {
			fmt.Println("Position: ", position)
			break
		}
	}
	fmt.Println("Floor: ", floor)
}

func main() {
	data := readFile("part1.txt")
	getFloor(data)
}
