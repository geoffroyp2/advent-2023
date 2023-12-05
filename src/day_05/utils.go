package day05

import (
	"sort"
	"strconv"
	"strings"
)

type Mapping struct {
	start     int
	end       int
	transform int
}

type Block struct {
	mappings []Mapping
}

type SeedRange struct {
	start int
	end   int
}

func getBlock(str string) Block {
	blockStr := strings.Split(str, "\n")
	mappings := make([]Mapping, 0)
	for i := 1; i < len(blockStr); i++ {
		valuesStr := strings.Split(blockStr[i], " ")
		val1, err1 := strconv.Atoi(valuesStr[0])
		if err1 != nil {
			panic(err1)
		}
		val2, err2 := strconv.Atoi(valuesStr[1])
		if err2 != nil {
			panic(err2)
		}
		val3, err3 := strconv.Atoi(valuesStr[2])
		if err3 != nil {
			panic(err3)
		}
		mapping := Mapping{transform: val1 - val2, start: val2, end: val2 + val3}
		mappings = append(mappings, mapping)
	}
	return Block{mappings: mappings}
}

func getSeeds(str string) []int {
	lineStr := strings.Split(str, ": ")
	seedsStr := strings.Split(lineStr[1], " ")

	seeds := make([]int, 0)
	for idx := range seedsStr {
		val, err := strconv.Atoi(seedsStr[idx])
		if err != nil {
			panic(err)
		}
		seeds = append(seeds, val)
	}

	return seeds
}

func getSeedRanges(str string) []SeedRange {
	lineStr := strings.Split(str, ": ")
	seedsStr := strings.Split(lineStr[1], " ")

	seeds := make([]int, 0)
	for idx := range seedsStr {
		val, err := strconv.Atoi(seedsStr[idx])
		if err != nil {
			panic(err)
		}
		seeds = append(seeds, val)
	}

	seedRanges := make([]SeedRange, 0)
	for i := 0; i < len(seeds); i += 2 {
		sr := SeedRange{
			start: seeds[i],
			end:   seeds[i+1] + seeds[i],
		}
		seedRanges = append(seedRanges, sr)
	}

	sort.SliceStable(seedRanges, func(i, j int) bool {
		return seedRanges[i].start < seedRanges[j].start
	})
	return seedRanges
}

func getEndValue(blocks *[]Block, startIdx int) int {
	currValue := startIdx
	for blockIdx := range *blocks {
		for _, r := range (*blocks)[blockIdx].mappings {
			if currValue >= r.start && currValue < r.end {
				currValue = currValue + r.transform
				break
			}
		}
	}
	return currValue
}
