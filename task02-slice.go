package homework

func reverse(input []int64) (result []int64) {
	slc := make([]int64, len(input))
	copy(slc, input)
	for i, j := 0, len(slc)-1; i < j; i, j = i+1, j-1 {
		slc[i], slc[j] = slc[j], slc[i]
	}
	return slc
}
