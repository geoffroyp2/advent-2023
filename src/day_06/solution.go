package day06

import (
	"advent-2023/src/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Race struct {
	time     int
	distance int
}

func getRaces(lines []string) []Race {
	spaces_re := regexp.MustCompile(" +")
	sep_re := regexp.MustCompile(": +")

	line1Str := sep_re.Split(lines[0], -1)
	line1Val := spaces_re.Split(line1Str[1], -1)
	line2Str := sep_re.Split(lines[1], -1)
	line2Val := spaces_re.Split(line2Str[1], -1)

	races := make([]Race, 0)
	for idx := range line1Val {
		time, err1 := strconv.Atoi(line1Val[idx])
		if err1 != nil {
			panic(err1)
		}
		dist, err2 := strconv.Atoi(line2Val[idx])
		if err2 != nil {
			panic(err2)
		}
		races = append(races, Race{time: time, distance: dist})
	}
	return races
}

func getRacePart2(lines []string) Race {
	spaces_re := regexp.MustCompile(" +")
	sep_re := regexp.MustCompile(": +")

	line1Str := sep_re.Split(lines[0], -1)
	line1Val := spaces_re.Split(line1Str[1], -1)
	line2Str := sep_re.Split(lines[1], -1)
	line2Val := spaces_re.Split(line2Str[1], -1)

	timeStr := ""
	distStr := ""
	for idx := range line1Val {
		timeStr += line1Val[idx]
		distStr += line2Val[idx]
	}

	time, err1 := strconv.Atoi(timeStr)
	if err1 != nil {
		panic(err1)
	}
	dist, err2 := strconv.Atoi(distStr)
	if err2 != nil {
		panic(err2)
	}
	return Race{time: time, distance: dist}
}

func getWinAmount(race Race) int {
	low := 0
	high := race.time/2 + 1

	for low <= high {
		middle := (low + high) / 2
		value := (race.time - middle) * middle
		if value <= race.distance {
			low = middle + 1
		} else {
			high = middle - 1
		}
	}

	matchVal := (race.time - low) * low
	if matchVal < race.distance {
		low += 1
	}

	return race.time - (2 * low) + 1
}

func part1(lines []string) {
	races := getRaces(lines)
	total := 1
	for idx := range races {
		amount := getWinAmount(races[idx])
		total *= amount
	}
	fmt.Println(total)
}

func part2(lines []string) {
	race := getRacePart2(lines)
	total := getWinAmount(race)
	fmt.Println(total)
}

func Run() {
	input := utils.GetFileContent("./src/day_06/input")
	lines := strings.Split(strings.Trim(input, "\n "), "\n")

	part1(lines)
	part2(lines)
}
