package utils

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func LCMForList(val []int) int {
	result := lcm(val[0], val[1])
	for i := 2; i < len(val); i++ {
		result = lcm(result, val[i])
	}
	return result
}
