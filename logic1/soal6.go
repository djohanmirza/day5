package logic1

func Soal6(n int) []int {
	num := 30
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = num
		num -= 3
	}
	return result
}
