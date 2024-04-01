package main

import (
	"fmt"
	"os"
	"strings"
)

func readLines(input string) []string {
	data, err := os.ReadFile(input)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(data), "\n")
}

func isWordNice(word string) bool {
	vowels := "aeiou"
	hasDouble, hasForbidden := false, false
	numVowels := 0

	for i := 0; i < len(word); i++ {
		if strings.Contains(vowels, string(word[i])) {
			numVowels++
		}
		if i > 0 && word[i] == word[i-1] {
			hasDouble = true
		}
		if i > 0 && ((word[i-1] == 'a' && word[i] == 'b') || (word[i-1] == 'c' && word[i] == 'd') || (word[i-1] == 'p' && word[i] == 'q') || (word[i-1] == 'x' && word[i] == 'y')) {
			hasForbidden = true
		}
	}
	return (numVowels >= 3) && hasDouble && !hasForbidden
}

func isWordNice2(word string) bool {
	hasDoublePair, hasRepeatingLetter := false, false
	for i := 0; i < len(word)-1; i++ {
		if strings.Contains(word[i+2:], word[i:i+2]) {
			hasDoublePair = true
		}
		if i < len(word)-2 && word[i] == word[i+2] {
			hasRepeatingLetter = true
		}
	}
	return hasDoublePair && hasRepeatingLetter
}

func main() {
	data := readLines("input.txt")
	niceWords := 0
	for _, word := range data {
		if isWordNice2(word) {
			niceWords++
		}
	}

	fmt.Println(niceWords)
	fmt.Println(isWordNice2("aaa"))
}
