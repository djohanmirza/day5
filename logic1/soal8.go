package logic1

func Soal81(n int) []int {
	mid := n / 2
	num := 2
	result := make([]int, n)

	for i := 0; i < n; i++ {
		result[i] = num
		if i < mid {
			num += 2
		} else if i > mid {
			num -= 2
		}
	}
	return result
}

func Soal82(n int) []int {
	mid := n / 2
	num := 2
	result := make([]int, n)

	for i := 0; i < n; i++ {
		result[i] = num
		if i < mid {
			num += 2
		} else if i >= mid {
			num -= 2
		}
	}
	return result
}
