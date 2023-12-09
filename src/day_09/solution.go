package day06

import (
	"advent-2023/src/utils"
	"fmt"
	"strconv"
	"strings"
)

func getSequence(line string) [][]int {
	vStr := strings.Split(line, " ")
	values := make([]int, len(vStr))
	for idx := range vStr {
		val, err := strconv.Atoi(vStr[idx])
		if err != nil {
			panic(err)
		}
		values[idx] = val
	}

	sequence := make([][]int, 0)
	sequence = append(sequence, values)
	for {
		nextVal := make([]int, 0)
		allZero := true
		for idx := 0; idx < len(values)-1; idx++ {
			nextVal = append(nextVal, values[idx+1]-values[idx])
			if nextVal[idx] != 0 {
				allZero = false
			}
		}
		values = nextVal
		sequence = append(sequence, nextVal)
		if allZero {
			break
		}
	}
	return sequence
}

func getNextValue(line string) int {
	sequence := getSequence(line)
	for idx := len(sequence) - 1; idx > 0; idx-- {
		lastValCurrent := sequence[idx][len(sequence[idx])-1]
		lastValPrevious := sequence[idx-1][len(sequence[idx-1])-1]
		nextVal := lastValCurrent + lastValPrevious
		sequence[idx-1] = append(sequence[idx-1], nextVal)
	}
	return sequence[0][len(sequence[0])-1]
}

func part1(lines *[]string) {
	total := 0
	for idx := range *lines {
		total += getNextValue((*lines)[idx])
	}
	fmt.Println(total)
}

func getPreviousValue(line string) int {
	sequence := getSequence(line)
	for idx := len(sequence) - 1; idx > 0; idx-- {
		lastValCurrent := sequence[idx][0]
		lastValPrevious := sequence[idx-1][0]
		nextVal := lastValPrevious - lastValCurrent
		sequence[idx-1] = append([]int{nextVal}, sequence[idx-1]...)
	}
	return sequence[0][0]
}

func part2(lines *[]string) {
	total := 0
	for idx := range *lines {
		total += getPreviousValue((*lines)[idx])
	}
	fmt.Println(total)
}

func Run() {
	input := utils.GetFileContent("./src/day_09/input")
	lines := strings.Split(strings.Trim(input, "\n "), "\n")

	part1(&lines)
	part2(&lines)
}
