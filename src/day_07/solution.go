package day07

import (
	"advent-2023/src/utils"
	"strings"
)

type HType int

const (
	High      HType = iota
	OnePair   HType = iota
	TwoPairs  HType = iota
	ThreeKind HType = iota
	FullHouse HType = iota
	FourKind  HType = iota
	FiveKind  HType = iota
)

type Hand struct {
	cards []int
	htype HType
	bet   int
}

func Run() {
	input := utils.GetFileContent("./src/day_07/input")
	lines := strings.Split(strings.Trim(input, "\n "), "\n")

	part1(&lines)
	part2(&lines)
}
