package day06

import (
	"advent-2023/src/utils"
	"fmt"
	"strconv"
	"strings"
)

type OP = int
type Item struct {
	op    OP
	label string
	val   int
	hash  int
	full  string
}
type Lens struct {
	label string
	val   int
}

const (
	REMOVE OP = iota
	ADD    OP = iota
)

func getHash(s string) int {
	val := 0
	for _, c := range s {
		val += int(c)
		val *= 17
		val %= 256
	}
	return val
}

func getItem(s string) Item {
	sIdx := 0
	op := ADD
	for ; sIdx < len(s); sIdx++ {
		if s[sIdx] == '-' {
			op = REMOVE
			break
		}
		if s[sIdx] == '=' {
			break
		}
	}
	id := s[:sIdx]
	hash := getHash(id)
	if op == ADD {
		value, err := strconv.Atoi(s[sIdx+1:])
		if err != nil {
			panic(err)
		}
		return Item{op: op, label: id, hash: hash, val: value, full: s}
	}
	return Item{op: op, label: id, hash: hash, val: -1, full: s}
}

func step1(instructions []string) {
	total := 0
	for idx := range instructions {
		total += getHash(instructions[idx])
	}
	fmt.Println(total)
}

func remove(slice []Lens, idx int) []Lens {
	return append(slice[:idx], slice[idx+1:]...)
}

func handleItem(item *Item, boxes *[][]Lens) {
	if item.op == ADD {
		for idx := range (*boxes)[item.hash] {
			if (*boxes)[item.hash][idx].label == item.label {
				(*boxes)[item.hash][idx] = Lens{label: item.label, val: item.val} // Replace existing
				return
			}
		}
		(*boxes)[item.hash] = append((*boxes)[item.hash], Lens{label: item.label, val: item.val}) // Add new item
	} else {
		idx := 0
		for ; idx < len((*boxes)[item.hash]); idx++ {
			if (*boxes)[item.hash][idx].label == item.label {
				(*boxes)[item.hash] = remove((*boxes)[item.hash], idx) // Remove existing
				return
			}
		}
	}
}

func step2(instructions []string) {
	boxes := make([][]Lens, 256)
	for i := 0; i < 256; i++ {
		boxes[i] = make([]Lens, 0)
	}

	for idx := range instructions {
		item := getItem(instructions[idx])
		handleItem(&item, &boxes)
	}

	total := 0
	for bidx := range boxes {
		for lidx := range boxes[bidx] {
			total += (bidx + 1) * (lidx + 1) * boxes[bidx][lidx].val
		}
	}
	fmt.Println(total)
}

func Run() {
	input := utils.GetFileContent("./src/day_15/input")
	lines := strings.Split(strings.Trim(input, "\n "), "\n")
	instructions := strings.Split(lines[0], ",")

	step1(instructions)
	step2(instructions)
}
