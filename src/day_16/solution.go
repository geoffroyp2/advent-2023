package day06

import (
	"advent-2023/src/utils"
	"errors"
	"fmt"
	"strings"
)

type Dir int

const (
	U Dir = iota
	D Dir = iota
	L Dir = iota
	R Dir = iota
)

type Position struct {
	x   int
	y   int
	dir Dir
}
type Cell struct {
	val     rune
	visited []Dir
}
type Grid [][]Cell

func getGrid(lines []string) Grid {
	grid := make([][]Cell, len(lines))
	for y := range lines {
		row := make([]Cell, len(lines[y]))
		for x := range lines[y] {
			row[x] = Cell{val: rune(lines[y][x])}
		}
		grid[y] = row
	}
	return grid
}

func getNext(pos Position, w, h int) (Position, error) {
	switch pos.dir {
	case U:
		if pos.y == 0 {
			return pos, errors.New("OOB")
		}
		return Position{x: pos.x, y: pos.y - 1, dir: pos.dir}, nil
	case D:
		if pos.y == h-1 {
			return pos, errors.New("OOB")
		}
		return Position{x: pos.x, y: pos.y + 1, dir: pos.dir}, nil
	case L:
		if pos.x == 0 {
			return pos, errors.New("OOB")
		}
		return Position{x: pos.x - 1, y: pos.y, dir: pos.dir}, nil
	case R:
		if pos.x == w-1 {
			return pos, errors.New("OOB")
		}
		return Position{x: pos.x + 1, y: pos.y, dir: pos.dir}, nil
	}
	panic("INVALID DIRECTION")
}

func walk(grid *Grid, start Position, queue *[]Position) {
	w := len((*grid)[0])
	h := len(*grid)
	pos := start

	for {
		cell := &(*grid)[pos.y][pos.x]

		for idx := range cell.visited {
			if cell.visited[idx] == pos.dir {
				return
			}
		}

		cell.visited = append(cell.visited, pos.dir)
		switch cell.val {
		case '.':
			next, err := getNext(pos, w, h)
			if err != nil {
				return
			}
			pos = next
		case '/':
			switch pos.dir {
			case U:
				pos.dir = R
			case D:
				pos.dir = L
			case R:
				pos.dir = U
			case L:
				pos.dir = D
			}
			next, err := getNext(pos, w, h)
			if err != nil {
				return
			}
			pos = next
		case '\\':
			switch pos.dir {
			case U:
				pos.dir = L
			case D:
				pos.dir = R
			case R:
				pos.dir = D
			case L:
				pos.dir = U
			}
			next, err := getNext(pos, w, h)
			if err != nil {
				return
			}
			pos = next
		case '|':
			if pos.dir == D || pos.dir == U {
				next, err := getNext(pos, w, h)
				if err != nil {
					return
				}
				pos = next
			} else {
				n1 := pos
				n1.dir = U
				n2 := pos
				n2.dir = D
				next1, err1 := getNext(n1, w, h)
				if err1 == nil {
					*queue = append(*queue, next1)
				}
				next2, err2 := getNext(n2, w, h)
				if err2 == nil {
					*queue = append(*queue, next2)
				}
				return
			}
		case '-':
			if pos.dir == L || pos.dir == R {
				next, err := getNext(pos, w, h)
				if err != nil {
					return
				}
				pos = next
			} else {
				n1 := pos
				n1.dir = L
				n2 := pos
				n2.dir = R
				next1, err1 := getNext(n1, w, h)
				if err1 == nil {
					*queue = append(*queue, next1)
				}
				next2, err2 := getNext(n2, w, h)
				if err2 == nil {
					*queue = append(*queue, next2)
				}
				return
			}
		}
	}
}

func getValue(lines []string, start Position) int {
	grid := getGrid(lines)
	queue := make([]Position, 0)
	queue = append(queue, start)

	for len(queue) != 0 {
		next := queue[0]
		queue = queue[1:]
		walk(&grid, next, &queue)
	}

	total := 0
	for y := range grid {
		for x := range grid[y] {
			if len(grid[y][x].visited) > 0 {
				total += 1
			}
		}
	}
	return total
}

func part2(lines []string) {
	maxValue := 0

	for y := 0; y < len(lines); y++ {
		rVal := getValue(lines, Position{x: 0, y: y, dir: R})
		maxValue = max(maxValue, rVal)
		lVal := getValue(lines, Position{x: len(lines[0]) - 1, y: y, dir: L})
		maxValue = max(maxValue, lVal)
	}
	for x := 0; x < len(lines[0]); x++ {
		dVal := getValue(lines, Position{x: x, y: 0, dir: D})
		maxValue = max(maxValue, dVal)
		uVal := getValue(lines, Position{x: x, y: len(lines) - 1, dir: U})
		maxValue = max(maxValue, uVal)
	}
	fmt.Println(maxValue)
}

func part1(lines []string) {
	grid := getGrid(lines)
	queue := make([]Position, 0)
	queue = append(queue, Position{x: 0, y: 0, dir: R})

	for len(queue) != 0 {
		next := queue[0]
		queue = queue[1:]
		walk(&grid, next, &queue)
	}

	total := 0
	for y := range grid {
		for x := range grid[y] {
			if len(grid[y][x].visited) > 0 {
				total += 1
			}
		}
	}
	fmt.Print(total)
}

func Run() {
	input := utils.GetFileContent("./src/day_16/input")
	lines := strings.Split(strings.Trim(input, "\n "), "\n")

	part1(lines)
	part2(lines)
}
