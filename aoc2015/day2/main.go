package main

import (
	"fmt"
	"os"
	"strings"
)

func readFileLines(input string) []string {
	data, err := os.ReadFile(input)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(data), "\n")
}

func calculateWrappingPaper(boxSize string) (int, int) {
	var l, w, h int
	fmt.Sscanf(boxSize, "%dx%dx%d", &l, &w, &h)
	lw := l * w
	wh := w * h
	hl := h * l
	vol := l * w * h
	smallestPerimeter := min(2*l+2*w, 2*w+2*h, 2*h+2*l)
	return 2*lw + 2*wh + 2*hl + min(lw, wh, hl),
		vol + smallestPerimeter
}

func main() {
	sum := 0
	var val int
	for _, line := range readFileLines("input.txt") {
		_, val = calculateWrappingPaper(line)
		sum += val
	}
	fmt.Print(sum)
}
