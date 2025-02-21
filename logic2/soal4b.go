package logic2

import go_slice_fresh "github.com/djohanmirza/Slice-Fresh"

func Soal4b(n int) (result [][]int) {
	start := 1
	result = go_slice_fresh.CreateSlice(n)
	for i := start; i < n; i++ {
		for j := 0; j < i; j++ {
			if i%2 == 0 {
				result[i][j] = start
				start += 3
			} else {
				result[i][j] = start
				start += 3
			}
		}
	}
	return result
}
