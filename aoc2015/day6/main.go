package main

import (
	"os"
	"strconv"
	"strings"
)

func readLines(input string) []string {
	data, err := os.ReadFile(input)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(data), "\n")
}

func main() {
	lights := [1000][1000]bool{}
	data := readLines("input.txt")
	for _, line := range data {
		words := strings.Split(line, " ")
		fromString := strings.Split(words[len(words)-3], ",")
		toString := strings.Split(words[len(words)-1], ",")
		fromX, _ := strconv.Atoi(fromString[0])
		fromY, _ := strconv.Atoi(fromString[1])
		toX, _ := strconv.Atoi(toString[0])
		toY, _ := strconv.Atoi(toString[1])

		for i := fromX; i <= toX; i++ {
			for j := fromY; j <= toY; j++ {
				switch words[1] {
				case "on":
					lights[i][j] = true
				case "off":
					lights[i][j] = false
				default:
					lights[i][j] = !lights[i][j]
				}
			}
		}
	}
	count := 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if lights[i][j] {
				count++
			}
		}
	}
}
