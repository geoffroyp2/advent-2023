package day02

import (
	"advent-2023/src/utils"
	"strings"
)

func Run() {
	input := utils.GetFileContent("./src/day_02/input")
	lines := strings.Split(strings.Trim(input, "\n "), "\n")

	total := 0
	for idx := range lines {
		game := NewGame(&lines[idx])

		red := game.GetMaxAmount("red")
		blue := game.GetMaxAmount("blue")
		green := game.GetMaxAmount("green")
		power := red * blue * green
		total += power
	}

	println(total)
}
