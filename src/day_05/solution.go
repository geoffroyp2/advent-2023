package day05

import (
	"advent-2023/src/utils"
	"fmt"
	"strconv"
	"strings"
)

type Mapping struct {
	source      int
	destination int
	amount      int
}

type Block struct {
	sourceName string
	destName   string
	mappings   []Mapping
}

type SeedRange struct {
	start  int
	amount int
}

func get_block(str string) Block {
	block := Block{}

	block_str := strings.Split(str, "\n")

	// name
	name_str := strings.Split(block_str[0], " ")
	name_block_str := strings.Split(name_str[0], "-to-")
	block.sourceName = name_block_str[0]
	block.destName = name_block_str[1]

	// mappings
	mappings := make([]Mapping, 0)
	for i := 1; i < len(block_str); i++ {
		values_str := strings.Split(block_str[i], " ")
		val1, err1 := strconv.Atoi(values_str[0])
		if err1 != nil {
			panic(err1)
		}
		val2, err2 := strconv.Atoi(values_str[1])
		if err2 != nil {
			panic(err2)
		}
		val3, err3 := strconv.Atoi(values_str[2])
		if err3 != nil {
			panic(err3)
		}

		mapping := Mapping{}
		mapping.destination = val1
		mapping.source = val2
		mapping.amount = val3

		mappings = append(mappings, mapping)
	}

	block.mappings = mappings
	return block
}

func get_seeds(str string) []int {
	line_str := strings.Split(str, ": ")
	seeds_str := strings.Split(line_str[1], " ")

	seeds := make([]int, 0)
	for idx := range seeds_str {
		val, err := strconv.Atoi(seeds_str[idx])
		if err != nil {
			panic(err)
		}
		seeds = append(seeds, val)
	}

	return seeds
}

func get_seed_ranges(str string) []SeedRange {
	line_str := strings.Split(str, ": ")
	seeds_str := strings.Split(line_str[1], " ")

	seeds := make([]int, 0)
	for idx := range seeds_str {
		val, err := strconv.Atoi(seeds_str[idx])
		if err != nil {
			panic(err)
		}
		seeds = append(seeds, val)
	}

	seedRanges := make([]SeedRange, 0)
	for i := 0; i < len(seeds); i += 2 {
		sr := SeedRange{}
		sr.start = seeds[i]
		sr.amount = seeds[i+1]
		seedRanges = append(seedRanges, sr)
	}

	return seedRanges
}

func get_end_value(blocks *[]Block, start_idx int) int {
	cv := start_idx

	for block_idx := range *blocks {
		for _, r := range (*blocks)[block_idx].mappings {
			if cv >= r.source && cv < (r.source+r.amount) {
				cv = r.destination + (cv - r.source)
				break
			}
		}
	}

	return cv
}

func part1(blocks_str *[]string) {
	blocks := make([]Block, 0)
	for idx := range *blocks_str {
		if idx == 0 {
			continue
		}
		block := get_block((*blocks_str)[idx])
		blocks = append(blocks, block)
	}

	seeds := get_seeds((*blocks_str)[0])
	min_val := -1
	for _, seed := range seeds {
		value := get_end_value(&blocks, seed)

		if min_val == -1 {
			min_val = value
		} else {
			if min_val > value {
				min_val = value
			}
		}
	}

	println(min_val)
}

func bruteforce(blocks *[]Block, seedrange *SeedRange, resultChan *chan int) {
	fmt.Println(*seedrange)
	min_val := -1
	for idx := 0; idx < seedrange.amount; idx++ {
		value := get_end_value(blocks, seedrange.start+idx)
		if min_val == -1 {
			min_val = value
		} else {
			if min_val > value {
				min_val = value
			}
		}
	}
	fmt.Println(*seedrange, min_val)
	*resultChan <- min_val
}

func part2(blocks_str *[]string) {
	blocks := make([]Block, 0)
	for idx := range *blocks_str {
		if idx == 0 {
			continue
		}
		block := get_block((*blocks_str)[idx])
		blocks = append(blocks, block)
	}

	seeds := get_seed_ranges((*blocks_str)[0])
	results := make(chan int, len(seeds))

	// Bruteforce lol
	for idx := range seeds {
		go bruteforce(&blocks, &seeds[idx], &results)
	}

	min_val := -1
	for range seeds {
		val := <-results
		if min_val == -1 {
			min_val = val
		} else if val < min_val {
			min_val = val
		}
	}

	println(min_val)
}

func Run() {
	input := utils.GetFileContent("./src/day_05/input")
	blocks_str := strings.Split(strings.Trim(input, "\n "), "\n\n")

	part1(&blocks_str)
	part2(&blocks_str)
}
