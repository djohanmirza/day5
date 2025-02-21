package logic2

import go_slice_fresh "github.com/djohanmirza/Slice-Fresh"

func Soal6b(n int) (result [][]int) {
	result = go_slice_fresh.CreateSlice(n)
	start := 1
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if i%2 == 0 {
				result[i][j] = start
				start += 3
			} else {
				result[i][j] = start
				start += 2
			}
		}
	}
	return result
}
