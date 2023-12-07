package day07

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func getCardValue2(c rune) int {
	if c >= '2' && c <= '9' {
		return int(c - '0')
	}
	switch c {
	case 'T':
		return 10
	case 'J':
		return 0
	case 'Q':
		return 12
	case 'K':
		return 13
	case 'A':
		return 14
	}
	panic(fmt.Sprintf("unkown card %c", c))
}

// LOL BEST CODE EVER
func getHType2(cards []int) HType {
	amounts := make([]int, 15)
	for _, val := range cards {
		amounts[val]++
	}

	bestRepeat := 0
	for idx, val := range amounts {
		if idx == 0 {
			continue // Skip jokers
		}
		if val > bestRepeat {
			bestRepeat = val
		}
	}
	jAmount := amounts[0]

	switch bestRepeat {
	case 0:
		return FiveKind // All jokers
	case 1:
		switch jAmount {
		case 0:
			return High
		case 1:
			return OnePair
		case 2:
			return ThreeKind
		case 3:
			return FourKind
		case 4:
			return FiveKind
		}
	case 2:
		{
			pairAmount := 0
			for idx, val := range amounts {
				if idx != 0 && val == 2 { // Skip jokers
					pairAmount++
				}
			}
			if pairAmount == 1 {
				switch jAmount {
				case 0:
					return OnePair
				case 1:
					return ThreeKind
				case 2:
					return FourKind
				case 3:
					return FiveKind
				}
			} else if pairAmount == 2 {
				switch jAmount {
				case 0:
					return TwoPairs
				case 1:
					return FullHouse
				}
			}
		}
	case 3:
		{
			pairAmount := 0
			for idx, val := range amounts {
				if idx != 0 && val == 2 { // Skip jokers
					pairAmount++
				}
			}
			if pairAmount == 0 {
				switch jAmount {
				case 0:
					return ThreeKind
				case 1:
					return FourKind
				case 2:
					return FiveKind
				}
			} else if pairAmount == 1 {
				switch jAmount {
				case 0:
					return FullHouse
				}
			}
		}
	case 4:
		switch jAmount {
		case 0:
			return FourKind
		case 1:
			return FiveKind
		}
	case 5:
		switch jAmount {
		case 0:
			return FiveKind
		}
	}
	panic("No combination found")
}

func getHand2(line *string) Hand {
	hStr := strings.Split(*line, " ")
	hand := Hand{}

	for _, c := range hStr[0] {
		hand.cards = append(hand.cards, getCardValue2(c))
	}
	hand.htype = getHType2(hand.cards)

	bet, err := strconv.Atoi(hStr[1])
	if err != nil {
		panic(err)
	}
	hand.bet = bet

	return hand
}

func part2(lines *[]string) {
	hands := make([]Hand, len(*lines))
	for idx := range *lines {
		hand := getHand2(&(*lines)[idx])
		hands[idx] = hand
	}

	sort.SliceStable(hands, func(i, j int) bool {
		if hands[i].htype == hands[j].htype {
			for idx := range hands[i].cards {
				if hands[i].cards[idx] == hands[j].cards[idx] {
					continue
				}
				return hands[i].cards[idx] < hands[j].cards[idx]
			}
		}
		return hands[i].htype < hands[j].htype
	})

	total := 0
	for idx := range hands {
		total += hands[idx].bet * (idx + 1)
	}
	fmt.Println(total)
}
