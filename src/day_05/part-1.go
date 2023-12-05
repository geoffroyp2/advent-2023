package day05

func part1(blocks_str *[]string) {
	blocks := make([]Block, 0)
	for idx := range *blocks_str {
		if idx == 0 {
			continue
		}
		block := getBlock((*blocks_str)[idx])
		blocks = append(blocks, block)
	}

	seeds := getSeeds((*blocks_str)[0])
	min_val := -1
	for _, seed := range seeds {
		value := getEndValue(&blocks, seed)

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
