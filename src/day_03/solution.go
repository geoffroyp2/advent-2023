package day03

import (
	"advent-2023/src/utils"
	"strings"
)

type Coord struct {
	x int
	y int
}

type Entry struct {
	value int
	start Coord
	end   Coord
}

func is_num(char byte) bool {
	return char >= '0' && char <= '9'
}

func is_symbol(char byte) bool {
	return !is_num(char) && char != '.'
}

func get_entry(lines *[]string, x int, y int) Entry {
	// Get bounds
	x_start := x
	x_end := x + 1
	for {
		if x_start < 0 || !is_num((*lines)[y][x_start]) {
			break
		}
		x_start--
	}
	for {
		if x_end == len((*lines)[y]) || !is_num((*lines)[y][x_end]) {
			break
		}
		x_end++
	}

	// Parse values
	factor := 1
	value := 0
	for i := x_end - 1; i > x_start; i-- {
		value += int(((*lines)[y][i] - '0')) * factor
		factor *= 10
	}
	return Entry{value: value, start: Coord{x: x_start, y: y}, end: Coord{x: x_end, y: y}}
}

func is_valid(lines *[]string, entry Entry) bool {
	y := entry.start.y
	start_x := entry.start.x
	end_x := entry.end.x

	for x := start_x - 1; x < end_x+1; x++ {
		if x < 0 || x >= len((*lines)[y]) {
			continue
		}
		// Check above
		if y > 0 && is_symbol((*lines)[y-1][x]) {
			return true
		}
		// Check below
		if y < len(*lines)-1 && is_symbol((*lines)[y+1][x]) {
			return true
		}
	}
	// Check left
	if start_x > 0 && is_symbol((*lines)[y][start_x-1]) {
		return true
	}
	// Check right
	if end_x < len((*lines)[y]) && is_symbol((*lines)[y][end_x]) {
		return true
	}

	return false
}

func part_1(lines *[]string) {
	total := 0
	for y := range *lines {
		for x := 0; x < len((*lines)[y]); x++ {
			if is_num((*lines)[y][x]) {
				entry := get_entry(lines, x, y)
				x = entry.end.x
				if is_valid(lines, entry) {
					total += entry.value
				}
			}
		}
	}
	println(total)
}

func get_adj_entries(lines *[]string, x int, y int) []Entry {
	entries := make([]Entry, 0)

	for dx := -1; dx < 2; dx++ {
		X := x + dx
		if X < 0 || X >= len((*lines)[y]) {
			continue
		}

		for dy := -1; dy < 2; dy++ {
			Y := y + dy
			if Y < 0 || Y >= len(*lines) {
				continue
			}

			if is_num((*lines)[Y][X]) {
				entry := get_entry(lines, X, Y)

				is_new_entry := true
				for idx := range entries {
					if entries[idx].start.x == entry.start.x && entries[idx].start.y == entry.start.y {
						is_new_entry = false
						break
					}
				}
				if is_new_entry {
					entries = append(entries, entry)
				}
			}
		}
	}

	return entries
}

func part_2(lines *[]string) {
	total := 0
	for y := range *lines {
		for x := range (*lines)[y] {
			if (*lines)[y][x] == '*' {
				adj_entries := get_adj_entries(lines, x, y)
				if len(adj_entries) == 2 {
					total += adj_entries[0].value * adj_entries[1].value
				}
			}
		}
	}

	println(total)
}

func Run() {
	input := utils.GetFileContent("./src/day_03/input")
	lines := strings.Split(strings.Trim(input, "\n "), "\n")
	part_1(&lines)
	part_2(&lines)
}
