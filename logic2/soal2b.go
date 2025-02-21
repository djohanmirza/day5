package logic2

import go_slice_fresh "github.com/djohanmirza/Slice-Fresh"

func Soal2b(n int) (result [][]int) {
	result = go_slice_fresh.CreateSlice(n)
	for i := 0; i < n; i++ {
		num := 2
		for j := 0; j < num; j++ {
			result[i][j] = num
			num += 2
		}
	}
	return result
}
