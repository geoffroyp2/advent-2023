package day08

import (
	"advent-2023/src/utils"
	"strings"
)

type Node struct {
	name  string
	left  string
	right string
}

func Run() {
	input := utils.GetFileContent("./src/day_08/input")
	lines := strings.Split(strings.Trim(input, "\n "), "\n")

	part1(&lines)
	part2(&lines)
}
