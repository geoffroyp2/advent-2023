package day05

import "sort"

type RangeMap struct {
	source      SeedRange
	destination SeedRange
}

// Get the intersection between 2 ranges
func getIntersection(A *SeedRange, B Mapping) *SeedRange {
	if A.start >= B.start && A.start < B.end {
		if A.end < B.end {
			return &SeedRange{start: A.start, end: A.end}
		} else {
			return &SeedRange{start: A.start, end: B.end}
		}
	}

	if B.start >= A.start && B.start < A.end {
		if B.end < A.end {
			return &SeedRange{start: B.start, end: B.end}
		} else {
			return &SeedRange{start: B.start, end: A.end}
		}
	}

	return nil
}

func getMinValueR(blocks *[]Block, blockIdx int, seedRange *SeedRange) int {
	if blockIdx >= len(*blocks) {
		// End of recursion
		return seedRange.start
	}

	nextBlocks := make([]RangeMap, 0)
	for _, m := range (*blocks)[blockIdx].mappings {
		// Get the intersection with current range
		source := getIntersection(seedRange, m)

		if source != nil {
			// If there is an intersection, calculate the resulting range for next iteration
			destination := SeedRange{start: source.start + m.transform, end: source.end + m.transform}
			nextBlocks = append(nextBlocks, RangeMap{source: *source, destination: destination})
		}
	}
	// Sort ranges in order to easily fill the gaps
	sort.SliceStable(nextBlocks, func(i, j int) bool {
		return nextBlocks[i].source.start < nextBlocks[j].source.start
	})

	missing := make([]SeedRange, 0)
	val := seedRange.start
	for _, next := range nextBlocks {
		// Find portions of original range that were not re-mapped
		if next.source.start > val {
			miss := SeedRange{start: val, end: next.source.start - 1}
			missing = append(missing, miss)
		}
		val = next.source.end
	}
	for idx := range missing {
		// Add missing blocks for next iteration
		nextBlocks = append(nextBlocks, RangeMap{source: missing[idx], destination: missing[idx]})
	}

	// If there is no next iteration identified, keep the original unmapped
	if len(nextBlocks) == 0 {
		nextBlocks = append(nextBlocks, RangeMap{source: *seedRange, destination: *seedRange})
	}

	// Recursively get the minimum value from the next iterations
	minVal := -1
	for _, next := range nextBlocks {
		nextVal := getMinValueR(blocks, blockIdx+1, &next.destination)
		if minVal == -1 {
			minVal = nextVal
		} else if nextVal < minVal {
			minVal = nextVal
		}
	}
	return minVal
}

func getMinValue(blocks *[]Block, seedrange *SeedRange, resultChan *chan int) {
	minVal := getMinValueR(blocks, 0, seedrange)
	*resultChan <- minVal
}

func part2(blocks_str *[]string) {
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

	for idx := range seeds {
		go getMinValue(&blocks, &seeds[idx], &results)
	}

	minVal := -1
	for range seeds {
		val := <-results
		if minVal == -1 {
			minVal = val
		} else if val < minVal {
			minVal = val
		}
	}

	println(minVal)
}
