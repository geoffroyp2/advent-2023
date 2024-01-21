package day06

import (
	"advent-2023/src/utils"
	"fmt"
	"strings"
)

type Block struct {
	lines  []string
	vMatch []int
	hMatch []int
}

func getBlocks(lines []string) []Block {
	blocks := make([]Block, 0)

	start := 0
	end := 0
	for ; end < len(lines); end++ {
		if lines[end] == "" {
			blocks = append(blocks, Block{lines: lines[start:end]})
			start = end + 1
		}
	}
	blocks = append(blocks, Block{lines: lines[start:end]})
	return blocks
}

func addBlockSym(block *Block, print ...bool) {
	// Vertical
	vResults := make([][]bool, 0)
	for lidx := range block.lines {
		line := block.lines[lidx]
		lineResult := make([]bool, 0)
		for x := 1; x < len(line); x++ {
			dx := 0
			match := true
			for {
				if x+dx >= len(line) || x-dx-1 < 0 {
					break // this is a match
				}
				if line[x+dx] != line[x-dx-1] {
					match = false
					break
				}
				dx++
			}
			lineResult = append(lineResult, match)
		}
		vResults = append(vResults, lineResult)
	}
	vMatches := make([]int, 0)
	for x := range vResults[0] {
		isMatch := true
		for y := range block.lines {
			if !vResults[y][x] {
				isMatch = false
				break
			}
		}
		if isMatch {
			vMatches = append(vMatches, x+1)
		}
	}
	block.vMatch = vMatches

	// Horizontal
	hMatches := make([]int, 0)
	for y := 1; y < len(block.lines); y++ {
		dy := 0
		match := true
		for {
			if y+dy >= len(block.lines) || y-dy-1 < 0 {
				break // this is a match
			}
			if block.lines[y+dy] != block.lines[y-dy-1] {
				match = false
				break
			}
			dy++
		}
		if match {
			hMatches = append(hMatches, y)
		}
	}
	block.hMatch = hMatches
}

func getNewSym(block *Block) Block {
	for y := range block.lines {
		for x := range block.lines[y] {
			lineCopy := make([]string, len(block.lines))
			copy(lineCopy, block.lines)
			copy := Block{lines: lineCopy}
			isEmpty := copy.lines[y][x] == '.'
			if isEmpty {
				newLine := copy.lines[y][:x] + "#" + copy.lines[y][x+1:]
				copy.lines[y] = newLine
			} else {
				newLine := copy.lines[y][:x] + "." + copy.lines[y][x+1:]
				copy.lines[y] = newLine
			}
			addBlockSym(&copy)

			if len(copy.hMatch) > len(block.hMatch) { // New hmatch
				copy.vMatch = make([]int, 0) // Remove vmatch
				if len(copy.hMatch) == 1 {
					return copy
				}
				// 2 hmatches
				for _, v := range copy.hMatch {
					if v != block.hMatch[0] { // Find the new one
						copy.hMatch = make([]int, 0)
						copy.hMatch = append(copy.hMatch, v)
						return copy
					}
				}
			}
			if len(copy.hMatch) == 1 && len(block.hMatch) == 1 && copy.hMatch[0] != block.hMatch[0] { // hmatch changed
				copy.vMatch = make([]int, 0) // Remove vmatch
				return copy
			}

			if len(copy.vMatch) > len(block.vMatch) { // New vMatch
				copy.hMatch = make([]int, 0) // Remove hMatch

				if len(copy.vMatch) == 1 {
					return copy
				}
				for _, v := range copy.vMatch {
					if v != block.vMatch[0] {
						copy.vMatch = make([]int, 0)
						copy.vMatch = append(copy.vMatch, v)
						return copy
					}
				}
			}
			if len(copy.vMatch) == 1 && len(block.vMatch) == 1 && copy.vMatch[0] != block.vMatch[0] { // vMatch changed
				copy.hMatch = make([]int, 0) // Remove hMatch
				return copy
			}
		}
	}
	for i := range block.lines {
		fmt.Println(block.lines[i])
	}
	fmt.Println(block.vMatch, block.hMatch)
	panic("No Valid symetry found")
}

func part1(lines []string) {
	blocks := getBlocks(lines)

	total := 0
	for idx := range blocks {
		addBlockSym(&blocks[idx])
		for _, v := range blocks[idx].vMatch {
			total += v
		}
		for _, v := range blocks[idx].hMatch {
			total += 100 * v
		}
	}

	fmt.Println(total) // 39939
}

func part2(lines []string) {
	blocks := getBlocks(lines)
	total := 0
	for idx := range blocks {
		addBlockSym(&blocks[idx])
		newBlock := getNewSym(&blocks[idx])
		for _, v := range newBlock.vMatch {
			total += v
		}
		for _, v := range newBlock.hMatch {
			total += 100 * v
		}
	}
	fmt.Println(total)
}

func Run() {
	input := utils.GetFileContent("./src/day_13/input")
	lines := strings.Split(strings.Trim(input, "\n "), "\n")

	part1(lines)
	part2(lines)
}
