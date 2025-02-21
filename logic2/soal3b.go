package logic2

import go_slice_fresh "github.com/djohanmirza/Slice-Fresh"

func Soal3b(n int) (result [][]int) {
	result = go_slice_fresh.CreateSlice(n)
	start := 1
	for i := start; i < n; i++ {
		for j := start; j < i; j++ {
			if i%2 == 0 {
				result[i][j] = start
				start += 2
			} else {
				result[i][j] = start
				start += 2
			}
		}
	}
	return result
}
