package day05

import "fmt"

func getMinValueBruteforce(blocks *[]Block, seedrange *SeedRange, resultChan *chan int) {
	fmt.Println(*seedrange)
	minVal := -1
	for idx := seedrange.start; idx < seedrange.end; idx++ {
		value := getEndValue(blocks, idx)
		if minVal == -1 {
			minVal = value
		} else {
			if minVal > value {
				minVal = value
			}
		}
	}
	fmt.Println(*seedrange, minVal)
	*resultChan <- minVal
}

func part2Bruteforce(blocks_str *[]string) {
	blocks := make([]Block, 0)
	for idx := range *blocks_str {
		if idx == 0 {
			continue
		}
		block := getBlock((*blocks_str)[idx])
		blocks = append(blocks, block)
	}

	seeds := getSeedRanges((*blocks_str)[0])
	results := make(chan int, len(seeds))

	// Bruteforce lol
	for idx := range seeds {
		go getMinValueBruteforce(&blocks, &seeds[idx], &results)
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
