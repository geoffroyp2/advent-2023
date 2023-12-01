package day01

import (
	"advent-2023/src/structures"
	"advent-2023/src/utils"
	"strconv"
	"strings"
)

func Run() {
	input := utils.GetFileContent("./src/day_01/input")
	lines := strings.Split(strings.Trim(input, "\n "), "\n")

	trie := structures.NewTrie([]string{
		"1", "2", "3", "4", "5", "6", "7", "8", "9", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
	})

	total := 0
	for idx := range lines {
		solutions := trie.GetMatches(lines[idx])
		value, err := strconv.Atoi(converter[solutions[0]] + converter[solutions[len(solutions)-1]])
		if err != nil {
			panic(err)
		}
		total += value
	}

	println(total)
}

var converter = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
	"1":     "1",
	"2":     "2",
	"3":     "3",
	"4":     "4",
	"5":     "5",
	"6":     "6",
	"7":     "7",
	"8":     "8",
	"9":     "9",
}
