package day06

import (
	"advent-2023/src/utils"
	"fmt"
	"strconv"
	"strings"
)

type Record struct {
	s   string
	val []int
}

func getRecord(line string) Record {
	rStr := strings.Split(line, " ")
	vStr := strings.Split(rStr[1], ",")

	values := make([]int, 0)
	for _, v := range vStr {
		val, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		values = append(values, val)
	}
	return Record{s: rStr[0], val: values}
}

func getRecord2(line string) Record {
	rStr := strings.Split(line, " ")
	vStr := strings.Split(rStr[1], ",")

	values := make([]int, 0)
	for _, v := range vStr {
		val, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		values = append(values, val)
	}
	repeatedValues := make([]int, 0)
	for idx := 0; idx < 5; idx++ {
		repeatedValues = append(repeatedValues, values...)
	}

	repeatedString := rStr[0] + "?" + rStr[0] + "?" + rStr[0] + "?" + rStr[0] + "?" + rStr[0]

	return Record{s: repeatedString, val: repeatedValues}
}

func isValid(record Record) bool {
	vIdx := 0
	for sIdx := 0; sIdx < len(record.s); sIdx++ {
		char := record.s[sIdx]
		if char == '?' {
			return true // Always correct if we encounter first unkown
		}

		if char == '.' {
			continue // Nothing to check, go next
		}

		if char == '#' && vIdx >= len(record.val) {
			return false // Too many groups

		}

		// Check current group of #
		vAmount := 0
		for ; vAmount < record.val[vIdx]; vAmount++ {
			if sIdx+vAmount >= len(record.s) {
				return false // out of range while checking group
			}

			nChar := record.s[sIdx+vAmount]
			if nChar == '?' {
				return true // Always correct if we encounter first unkown
			}
			if nChar == '.' {
				return false // empty in the middle of a group
			}
		}
		sIdx += vAmount // jump to end of current group
		vIdx++          // Increment group id

		if sIdx >= len(record.s) {
			return vIdx == len(record.val) // Block is directly at the end, check that all group have been counted
		}
		if record.s[sIdx] == '#' {
			return false // Group should be followed by an empty space
		}
		if record.s[sIdx] == '?' {
			return true // Always correct if we encounter first unkown
		}
	}
	return vIdx == len(record.val) // check that all group have been counted
}

func computeValidPaths(record Record, valid *int) {
	if !isValid(record) {
		fmt.Println("[X]  ", record)
		return
	}

	firstUnkown := -1
	for idx, c := range record.s {
		if c == '?' {
			firstUnkown = idx
			break
		}
	}

	if firstUnkown == -1 {
		// No unknown left
		*valid += 1
		fmt.Println("[V]  ", record)
		return
	}

	present := record
	present.s = present.s[:firstUnkown] + "#" + present.s[firstUnkown+1:]
	computeValidPaths(present, valid)

	absent := record
	absent.s = absent.s[:firstUnkown] + "." + absent.s[firstUnkown+1:]
	computeValidPaths(absent, valid)
}

func getPossibilityAmount(record Record) int {
	valid := 0
	computeValidPaths(record, &valid)
	fmt.Println(record, valid)
	return valid
}

func part1(lines []string) {
	total := 0
	for idx := range lines {
		record := getRecord(lines[idx])
		total += getPossibilityAmount(record)
	}
	fmt.Println(total)
}

func part2(lines []string) {
	record := getRecord2(lines[0])
	getPossibilityAmount(record)

	// total := 0
	// for idx := range lines {
	// 	record := getRecord2(lines[idx])
	// 	total += getPossibilityAmount(record)
	// }
	// fmt.Println(total)
}

func Run() {
	input := utils.GetFileContent("./src/day_12/input")
	lines := strings.Split(strings.Trim(input, "\n "), "\n")
	// part1(lines)
	part2(lines)
}
