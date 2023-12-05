package day05

import (
	"advent-2023/src/utils"
	"strings"
)

func Run() {
	input := utils.GetFileContent("./src/day_05/input")
	blocksStr := strings.Split(strings.Trim(input, "\n "), "\n\n")

	part1(&blocksStr)
	// part2Bruteforce(&blocksStr)
	part2(&blocksStr)
}
