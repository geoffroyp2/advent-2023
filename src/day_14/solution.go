package day06

import (
	"advent-2023/src/utils"
	"fmt"
	"math/rand"
	"strings"
)

type Cell = int
type Grid = [][]Cell
type GridWithHash struct {
	grid Grid
	hash uint64
	idx  int
}

const (
	MovingCell Cell = iota
	FixedCell  Cell = iota
	EmptyCell  Cell = iota
)

func copyGrid(grid *Grid) Grid {
	copy := make([][]Cell, 0)
	for y := range *grid {
		l := make([]Cell, 0)
		for x := range (*grid)[0] {
			l = append(l, (*grid)[y][x])
		}
		copy = append(copy, l)
	}
	return copy
}

func printGrid(grid Grid) {
	fmt.Print("\n")
	for y := range grid {
		for x := range grid[0] {
			if grid[y][x] == MovingCell {
				fmt.Print("O")
			} else if grid[y][x] == FixedCell {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
}

func getGrid(lines []string) Grid {
	grid := make(Grid, 0)
	for y := range lines {
		l := make([]Cell, 0)
		for x := range lines[y] {
			if lines[y][x] == 'O' {
				l = append(l, MovingCell)
			} else if lines[y][x] == '#' {
				l = append(l, FixedCell)
			} else {
				l = append(l, EmptyCell)
			}
		}
		grid = append(grid, l)
	}
	return grid
}

func tiltNorth(grid *Grid) {
	for x := range (*grid)[0] {
		y := 0
		for y < len((*grid)) {
			if (*grid)[y][x] == MovingCell {
				for {
					if y-1 < 0 {
						break
					}
					if (*grid)[y-1][x] == EmptyCell {
						(*grid)[y-1][x] = MovingCell
						(*grid)[y][x] = EmptyCell
						y--
					} else {
						break
					}
				}
			}
			y++
		}
	}
}

func tiltSouth(grid *Grid) {
	for x := range (*grid)[0] {
		y := len((*grid)) - 1
		for y >= 0 {
			if (*grid)[y][x] == MovingCell {
				for {
					if y+1 >= len((*grid)) {
						break
					}
					if (*grid)[y+1][x] == EmptyCell {
						(*grid)[y+1][x] = MovingCell
						(*grid)[y][x] = EmptyCell
						y++
					} else {
						break
					}
				}
			}
			y--
		}
	}
}

func tiltWest(grid *Grid) {
	for y := range *grid {
		x := 0
		for x < len((*grid)[0]) {
			if (*grid)[y][x] == MovingCell {
				for {
					if x-1 < 0 {
						break
					}
					if (*grid)[y][x-1] == EmptyCell {
						(*grid)[y][x-1] = MovingCell
						(*grid)[y][x] = EmptyCell
						x--
					} else {
						break
					}
				}
			}
			x++
		}
	}
}

func tiltEast(grid *Grid) {
	for y := range *grid {
		x := len((*grid)[0]) - 1
		for x >= 0 {
			if (*grid)[y][x] == MovingCell {
				for {
					if x+1 >= len((*grid)[0]) {
						break
					}
					if (*grid)[y][x+1] == EmptyCell {
						(*grid)[y][x+1] = MovingCell
						(*grid)[y][x] = EmptyCell
						x++
					} else {
						break
					}
				}
			}
			x--
		}
	}
}

func getValue(grid Grid) int {
	total := 0
	for y := range grid {
		for x := range grid[0] {
			if grid[y][x] == MovingCell {
				total += len(grid) - y
			}
		}
	}
	return total
}

func initHash(x, y int) []uint64 {
	l := x * y
	val := make([]uint64, l)
	for i := 0; i < l; i++ {
		val[i] = rand.Uint64()
	}
	return val
}

func getHash(grid *Grid, val *[]uint64) uint64 {
	var total uint64 = 0
	l := len((*grid))
	for y := range *grid {
		for x := range (*grid)[0] {
			if ((*grid)[y][x]) == MovingCell {
				total ^= (*val)[y*l+x]
			}
		}
	}
	return total
}

func part1(lines []string) {
	grid := getGrid(lines)
	tiltNorth(&grid)
	total := getValue(grid)
	fmt.Println(total)
}

func part2(lines []string) {
	grid := getGrid(lines)
	hashValues := initHash(len(grid), len(grid[0]))
	hashes := make(map[uint64]GridWithHash)

	// Find the loop
	loopStart := -1
	loopEnd := -1
	for i := 0; i < 1000; i++ {
		tiltNorth(&grid)
		tiltWest(&grid)
		tiltSouth(&grid)
		tiltEast(&grid)
		hash := getHash(&grid, &hashValues)

		if hashes[hash].hash == 0 {
			hashes[hash] = GridWithHash{idx: i, grid: copyGrid(&grid), hash: hash}
		} else {
			loopStart = hashes[hash].idx
			loopEnd = i
			break
		}
	}

	loopLen := loopEnd - loopStart
	stepAtEnd := (1000000000-loopEnd)%loopLen - 1
	for _, g := range hashes {
		if g.idx == loopStart+stepAtEnd {
			fmt.Println(getValue(g.grid)) // Grid state at the end step
			break
		}
	}
}

func Run() {
	input := utils.GetFileContent("./src/day_14/input")
	lines := strings.Split(strings.Trim(input, "\n "), "\n")

	part1(lines) //108889
	part2(lines) // 104671
}
