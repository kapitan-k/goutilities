package slice

// Uint64SameSliceMulti returns a 2 dimensional slice of len = cnt where
// all elements are slice
func Uint64SameSliceMulti(slice []uint64, cnt int) (result [][]uint64) {
	result = make([][]uint64, cnt)
	for i := range result {
		result[i] = slice
	}

	return
}
